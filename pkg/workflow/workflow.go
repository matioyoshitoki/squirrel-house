package workflow

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"
	"time"
)

// Step 表示工作流中的一个步骤
type Step struct {
	ID          string            `yaml:"id"`
	Label       string            `yaml:"label"`
	Type        string            `yaml:"type"` // shell, readFile, writeFile
	Condition   string            `yaml:"condition,omitempty"`
	Command     string            `yaml:"command,omitempty"`
	Dir         string            `yaml:"dir,omitempty"`
	Env         map[string]string `yaml:"env,omitempty"`
	Path        string            `yaml:"path,omitempty"`
	Content     string            `yaml:"content,omitempty"`
	OutputKey   string            `yaml:"outputKey,omitempty"`
	IgnoreError bool              `yaml:"ignoreError,omitempty"`
	IsCleanup   bool              `yaml:"isCleanup,omitempty"`
	LogFile     string            `yaml:"logFile,omitempty"`
	Tmux        bool              `yaml:"tmux,omitempty"` // 用 tmux session 运行，Dashboard 重启不中断
}

// Workflow 表示一个完整的工作流
type Workflow struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

// Context 保存工作流执行期间的共享状态
type Context struct {
	Vars        map[string]interface{}
	State       map[string]interface{}
	ProjectPath string
	StartTime   time.Time
}

// NewContext 创建一个新的执行上下文
func NewContext(projectPath string, vars map[string]interface{}) *Context {
	return &Context{
		Vars:        vars,
		State:       make(map[string]interface{}),
		ProjectPath: projectPath,
		StartTime:   time.Now(),
	}
}

// Engine 负责执行工作流
type Engine struct {
	cleanupSteps  []Step
	ctx           *Context
	currentCmd    *exec.Cmd
	OnProcessStart func(pid int)
}

// NewEngine 创建执行引擎
func NewEngine(ctx *Context) *Engine {
	return &Engine{ctx: ctx}
}

// CurrentPID 返回当前正在执行的 shell 步骤的进程 PID
func (e *Engine) CurrentPID() int {
	if e.currentCmd != nil && e.currentCmd.Process != nil {
		return e.currentCmd.Process.Pid
	}
	return 0
}

// Run 执行指定工作流
func (e *Engine) Run(ctx context.Context, wf *Workflow) error {
	e.cleanupSteps = nil
	for i, step := range wf.Steps {
		if step.IsCleanup {
			e.cleanupSteps = append(e.cleanupSteps, step)
			continue
		}
		if err := e.runStep(ctx, step, i+1); err != nil {
			// 执行 cleanup 后返回错误（cleanup 用 background context，避免被取消）
			e.runCleanups()
			return err
		}
	}
	// 正常结束时也执行 cleanup
	e.runCleanups()
	return nil
}

// Kill 强制终止当前正在执行的 shell 步骤及其所有子进程。
// 由于 runShellStep 使用 Setsid 创建独立 session，进程组 ID 等于进程 ID，
// 通过向负 PID 发送 SIGKILL 可以级联终止整个进程组，避免孤儿进程残留。
func (e *Engine) Kill() error {
	if e.currentCmd == nil || e.currentCmd.Process == nil {
		return nil
	}
	pid := e.currentCmd.Process.Pid
	// 先尝试 kill 整个进程组（Setsid 创建的 session leader 同时也是进程组长）
	if err := syscall.Kill(-pid, syscall.SIGKILL); err == nil {
		return nil
	}
	// 降级：只 kill 当前进程
	return e.currentCmd.Process.Kill()
}

func (e *Engine) runCleanups() {
	// cleanup 步骤使用 background context，不受外层 cancel 影响
	log.Printf("[workflow] runCleanups: %d steps", len(e.cleanupSteps))
	for _, step := range e.cleanupSteps {
		log.Printf("[workflow] runCleanups: running step %s", step.ID)
		err := e.runStep(context.Background(), step, 0)
		if err != nil {
			log.Printf("[workflow] runCleanups: step %s failed: %v", step.ID, err)
		}
	}
}

func (e *Engine) runStep(ctx context.Context, step Step, idx int) error {
	// condition 判断
	if step.Condition != "" {
		cond, err := e.renderString(step.Condition)
		if err != nil {
			return fmt.Errorf("step %s condition render error: %w", step.ID, err)
		}
		if strings.TrimSpace(cond) == "" || strings.TrimSpace(cond) == "false" {
			return nil
		}
	}

	switch step.Type {
	case "shell":
		return e.runShellStep(ctx, step)
	case "readFile":
		return e.runReadFileStep(step)
	case "writeFile":
		return e.runWriteFileStep(step)
	default:
		return fmt.Errorf("unknown step type: %s", step.Type)
	}
}

var funcMap = template.FuncMap{
	"shellquote": func(s string) string {
		return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
	},
	"trim":     strings.TrimSpace,
	"contains": strings.Contains,
	"now":      func() int64 { return time.Now().Unix() },
}

func (e *Engine) renderString(s string) (string, error) {
	tmpl, err := template.New("wf").Funcs(funcMap).Parse(s)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, e.ctx); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ensureToolPaths 确保 PATH 包含常用开发工具（flutter、homebrew、Android SDK）
func ensureToolPaths(env []string) []string {
	extraPaths := []string{
		"/opt/flutter/bin",
		"/usr/local/bin",
		"/opt/android-sdk/platform-tools",
		"/opt/android-sdk/emulator",
	}
	for i, e := range env {
		if strings.HasPrefix(e, "PATH=") {
			pathVal := strings.TrimPrefix(e, "PATH=")
			parts := strings.Split(pathVal, ":")
			existing := make(map[string]bool)
			for _, p := range parts {
				existing[p] = true
			}
			for _, p := range extraPaths {
				if !existing[p] {
					parts = append(parts, p)
				}
			}
			env[i] = "PATH=" + strings.Join(parts, ":")
			return env
		}
	}
	// PATH not found, append it
	return append(env, "PATH="+strings.Join(extraPaths, ":"))
}

// StepResult 记录单个步骤的执行结果
type StepResult struct {
	Status   string `json:"status"`
	ExitCode int    `json:"exitCode"`
	Error    string `json:"error,omitempty"`
}

// recordStepResult 将步骤执行结果写入 workflow context
func (e *Engine) recordStepResult(step Step, exitCode int, err error) {
	key := "stepResult_" + step.ID
	if err != nil && !step.IgnoreError {
		e.ctx.State[key] = StepResult{
			Status:   "failed",
			ExitCode: exitCode,
			Error:    err.Error(),
		}
	} else {
		e.ctx.State[key] = StepResult{
			Status:   "success",
			ExitCode: exitCode,
		}
	}
}

// getExitCode 从 exec error 中提取 exit code
func getExitCode(err error) int {
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus()
		}
	}
	return -1
}

func (e *Engine) runShellStep(ctx context.Context, step Step) error {
	cmdStr, err := e.renderString(step.Command)
	if err != nil {
		return fmt.Errorf("render command: %w", err)
	}
	if strings.TrimSpace(cmdStr) == "" {
		return nil
	}

	dir := e.ctx.ProjectPath
	if step.Dir != "" {
		r, err := e.renderString(step.Dir)
		if err != nil {
			return fmt.Errorf("render dir: %w", err)
		}
		if r != "" {
			dir = r
		}
	}

	logFile, _ := e.renderString(step.LogFile)

	if step.Tmux {
		return e.runTmuxStep(ctx, step, dir, cmdStr, logFile)
	}

	cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true} // 独立 session，父进程退出不影响子进程
	cmd.Dir = dir
	cmd.Env = ensureToolPaths(os.Environ())
	for k, v := range step.Env {
		rv, _ := e.renderString(v)
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, rv))
	}
	// 自动注入以 AGENT_ 开头的上下文变量作为环境变量
	for k, v := range e.ctx.Vars {
		if strings.HasPrefix(k, "AGENT_") {
			if s, ok := v.(string); ok {
				cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, s))
			}
		}
	}

	var outputBuf bytes.Buffer
	var writers []io.Writer
	writers = append(writers, &outputBuf)
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			defer f.Close()
			writers = append(writers, f)
		}
	}
	mw := io.MultiWriter(writers...)
	cmd.Stdout = mw
	cmd.Stderr = mw

	e.currentCmd = cmd
	if err := cmd.Start(); err != nil {
		e.currentCmd = nil
		return fmt.Errorf("shell step %s start failed: %v", step.ID, err)
	}
	if e.OnProcessStart != nil && cmd.Process != nil {
		e.OnProcessStart(cmd.Process.Pid)
	}
	err = cmd.Wait()
	e.currentCmd = nil

	output := outputBuf.String()
	if step.OutputKey != "" {
		e.ctx.State[step.OutputKey] = output
	}

	exitCode := 0
	if err != nil {
		exitCode = getExitCode(err)
	}
	e.recordStepResult(step, exitCode, err)

	if err != nil && !step.IgnoreError {
		return fmt.Errorf("shell step %s failed: %v, output: %s", step.ID, err, output)
	}
	return nil
}

// runTmuxStep 用 tmux session 运行命令，Dashboard 重启不中断
func (e *Engine) runTmuxStep(ctx context.Context, step Step, dir, cmdStr, logFile string) error {
	// 生成 session 名称（只用 taskID，不依赖 stepID，确保 Kill/监控能匹配）
	taskID := "unknown"
	if v, ok := e.ctx.Vars["taskID"].(string); ok && v != "" {
		taskID = v
	}
	sessionName := fmt.Sprintf("nwops-%s", taskID)

	// 创建临时文件记录 exit code
	exitCodeFile := filepath.Join(os.TempDir(), fmt.Sprintf("nwops-exit-%s-%d", taskID, time.Now().UnixNano()))
	defer os.Remove(exitCodeFile)

	// 构建环境变量（必须在包装命令之前，以便注入到命令中）
	env := ensureToolPaths(os.Environ())
	for k, v := range step.Env {
		rv, _ := e.renderString(v)
		env = append(env, fmt.Sprintf("%s=%s", k, rv))
	}
	for k, v := range e.ctx.Vars {
		if strings.HasPrefix(k, "AGENT_") {
			if s, ok := v.(string); ok {
				env = append(env, fmt.Sprintf("%s=%s", k, s))
			}
		}
	}

	// 收集需要传递给 tmux session 的 AGENT_ 环境变量（用 -e 选项）
	var tmuxEnv []string
	for _, e := range env {
		if strings.Contains(e, "=") && strings.HasPrefix(strings.SplitN(e, "=", 2)[0], "AGENT_") {
			tmuxEnv = append(tmuxEnv, e)
		}
	}

	// 包装命令：输出同时写到日志文件，最后记录 exit code
	wrappedCmd := cmdStr
	if logFile != "" {
		wrappedCmd = fmt.Sprintf("(%s) 2>&1 | tee -a %s; echo $? > %s", cmdStr, logFile, exitCodeFile)
	} else {
		wrappedCmd = fmt.Sprintf("(%s); echo $? > %s", cmdStr, exitCodeFile)
	}

	// 清理可能存在的同名 session
	_ = exec.Command("tmux", "kill-session", "-t", sessionName).Run()

	// 如果命令太长，写入临时脚本文件执行（避免 tmux "command too long"）
	execCmd := wrappedCmd
	if len(wrappedCmd) > 8000 {
		scriptFile := filepath.Join(os.TempDir(), fmt.Sprintf("nwops-script-%s-%d.sh", taskID, time.Now().UnixNano()))
		if err := os.WriteFile(scriptFile, []byte("#!/bin/bash\n"+wrappedCmd+"\n"), 0755); err == nil {
			defer os.Remove(scriptFile)
			execCmd = scriptFile
		}
	}

	// 启动 tmux session，使用 -e 传递环境变量（避免 bash 字符串注入问题）
	args := []string{"new-session", "-d", "-s", sessionName, "-c", dir}
	for _, e := range tmuxEnv {
		args = append(args, "-e", e)
	}
	args = append(args, execCmd)
	startCmd := exec.Command("tmux", args...)
	startCmd.Env = env
	var startStderr bytes.Buffer
	startCmd.Stderr = &startStderr
	log.Printf("[workflow] tmux new-session: session=%s dir=%s cmdLen=%d", sessionName, dir, len(execCmd))
	if err := startCmd.Run(); err != nil {
		log.Printf("[workflow] tmux step %s start failed: %v, stderr: %s", step.ID, err, startStderr.String())
		return fmt.Errorf("tmux step %s start failed: %v", step.ID, err)
	}

	// 通知 process start（tmux 用 PID 0 表示）
	if e.OnProcessStart != nil {
		e.OnProcessStart(0)
	}

	// 轮询等待 session 结束
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Dashboard 重启或任务被停止：kill tmux session
			_ = exec.Command("tmux", "kill-session", "-t", sessionName).Run()
			return fmt.Errorf("tmux step %s cancelled", step.ID)
		case <-ticker.C:
			if err := exec.Command("tmux", "has-session", "-t", sessionName).Run(); err != nil {
				// session 不存在了，任务完成
				goto done
			}
		}
	}

done:
	// 读取 exit code（session 已结束，跳过 capture-pane 避免 tmux 挂起）
	exitCode := 0
	if data, err := os.ReadFile(exitCodeFile); err == nil {
		fmt.Sscanf(string(data), "%d", &exitCode)
	}
	// 清理 session（幂等）
	_ = exec.Command("tmux", "kill-session", "-t", sessionName).Run()

	var stepErr error
	if exitCode != 0 && !step.IgnoreError {
		stepErr = fmt.Errorf("tmux step %s exited with code %d", step.ID, exitCode)
	}
	e.recordStepResult(step, exitCode, stepErr)

	if stepErr != nil {
		return stepErr
	}
	return nil
}

func (e *Engine) runReadFileStep(step Step) error {
	p, err := e.renderString(step.Path)
	if err != nil {
		return fmt.Errorf("render path: %w", err)
	}
	data, err := os.ReadFile(p)
	if err != nil {
		if step.OutputKey != "" {
			e.ctx.State[step.OutputKey] = ""
		}
		if !step.IgnoreError {
			return fmt.Errorf("read file %s failed: %w", p, err)
		}
		return nil
	}
	if step.OutputKey != "" {
		e.ctx.State[step.OutputKey] = string(data)
	}
	return nil
}

func (e *Engine) runWriteFileStep(step Step) error {
	p, err := e.renderString(step.Path)
	if err != nil {
		return fmt.Errorf("render path: %w", err)
	}
	content, err := e.renderString(step.Content)
	if err != nil {
		return fmt.Errorf("render content: %w", err)
	}
	if err := os.WriteFile(p, []byte(content), 0644); err != nil {
		return fmt.Errorf("write file %s failed: %w", p, err)
	}
	return nil
}
