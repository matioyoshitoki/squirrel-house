package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/MatioYoshitoki/new-world-ops/pkg/workflow"
)

// HookCondition 判断是否应该触发 hook
// sourceTask: 触发 hook 的源任务
// sourceVars: 源任务启动时传入的变量
// wfCtx: 源任务 workflow 执行完成后的上下文（包含 State 和 Vars）
type HookCondition func(sourceTask *Task, sourceVars map[string]interface{}, wfCtx *workflow.Context) bool

// HookConfig 定义 workflow 完成后的 hook
// 当源 workflow 达到指定状态（成功/失败）时，异步启动目标 Task
type HookConfig struct {
	TargetType TaskType          // 目标任务类型，如 TaskTypeE2E
	Condition  HookCondition     // 触发条件，nil 表示无条件触发
	Vars       map[string]string // key=目标 task 变量名, value=Go template 表达式（渲染自 wfCtx）
	Delay      time.Duration     // 延迟触发（可选）
}

// taskSpec 任务类型规范
// 替代原有的匿名 struct，增加 AfterSuccess / AfterFailure hook 配置
type taskSpec struct {
	WorkflowName string
	AgentType    string
	OnStarted    func(*Task)
	OnFinished   func(*Task)
	AfterSuccess []HookConfig // 源 workflow 成功后触发
	AfterFailure []HookConfig // 源 workflow 失败后触发
}

// hookAsyncMode 控制 hook 触发是否异步。生产环境为 true，测试中设为 false 以保证确定性。
var hookAsyncMode = true

// triggerHooks 根据任务执行结果触发对应的 hooks
// 在 runTaskWorkflow 的 workflow 执行完成后、OnFinished 回调之前调用
func triggerHooks(task *Task, spec taskSpec, sourceVars map[string]interface{}, wfCtx *workflow.Context, success bool) {
	var hooks []HookConfig
	if success {
		hooks = spec.AfterSuccess
	} else {
		hooks = spec.AfterFailure
	}

	if len(hooks) == 0 {
		return
	}

	for _, hook := range hooks {
		if hook.Condition != nil && !hook.Condition(task, sourceVars, wfCtx) {
			log.Printf("[hook] 跳过 %s-%d → %s (条件不满足)", task.Type, task.TargetID, hook.TargetType)
			continue
		}

		// 渲染变量模板
		hookVars := make(map[string]interface{})
		for k, tmplStr := range hook.Vars {
			rendered, err := renderHookTemplate(tmplStr, wfCtx)
			if err != nil {
				log.Printf("[hook] 变量渲染失败 %s=%s: %v", k, tmplStr, err)
				continue
			}
			hookVars[k] = rendered
		}

		log.Printf("[hook] 触发 %s-%d → %s (vars keys=%v)", task.Type, task.TargetID, hook.TargetType, mapKeys(hookVars))

		if hook.Delay > 0 {
			time.AfterFunc(hook.Delay, func() {
				startHookTask(hook.TargetType, task, hookVars)
			})
		} else if hookAsyncMode {
			go startHookTask(hook.TargetType, task, hookVars)
		} else {
			startHookTask(hook.TargetType, task, hookVars)
		}
	}
}

// renderHookTemplate 使用 workflow context 渲染模板字符串
func renderHookTemplate(tmplStr string, wfCtx *workflow.Context) (string, error) {
	funcMap := template.FuncMap{
		"shellquote": func(s string) string {
			return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
		},
		"trim": strings.TrimSpace,
	}
	tmpl, err := template.New("hook").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, wfCtx); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

// startHookTask 根据目标类型启动对应的 hook 任务
func startHookTask(targetType TaskType, sourceTask *Task, vars map[string]interface{}) {
	hookDispatcher.Dispatch(targetType, sourceTask, vars)
}

// startE2ETask 启动 E2E 任务（统一入口：手动触发）
// E2E 是独立任务，由专门的 E2E issue 定义测试范围，不再依附于功能 issue 的附属流程。
func startE2ETask(issueNumber int, issueTitle, issueBody, branch string, taskKind string, triggeredBy string, worktreePath string, projectName string) {
	// 检查是否已有 running 的 E2E 任务
	if existing, ok := getTask(TaskTypeE2E, issueNumber); ok && existing.Status == "running" {
		log.Printf("[e2e] E2E 任务已在运行: e2e-%d", issueNumber)
		return
	}

	projectPath := getProjectPathByName(projectName)
	// 自动推断 Flutter 项目：如果指定项目没有 Flutter mobile，遍历其他项目
	if _, err := os.Stat(filepath.Join(projectPath, "apps", "mobile", "pubspec.yaml")); os.IsNotExist(err) {
		projectMutex.RLock()
		for _, p := range config.Projects {
			if _, err := os.Stat(filepath.Join(p.Path, "apps", "mobile", "pubspec.yaml")); err == nil {
				projectName = p.Name
				projectPath = p.Path
				log.Printf("[e2e] 自动推断项目: %s (路径: %s)", projectName, projectPath)
				break
			}
		}
		projectMutex.RUnlock()
	}

	logsDir := getProjectLogsDir(projectName)
	os.MkdirAll(logsDir, 0755)

	timestamp := time.Now().Format("20060102-150405")
	logFileName := filepath.Join(logsDir, fmt.Sprintf("e2e-issue-%d-%s.log", issueNumber, timestamp))

	f, err := os.Create(logFileName)
	if err != nil {
		log.Printf("[e2e] 创建 E2E 日志文件失败: %v", err)
		return
	}
	f.Close()

	logMsg := fmt.Sprintf("[%s] E2E 任务启动 (由 %s)\nIssue: #%d %s\n分支: %s\n项目: %s\n日志: %s\n\n",
		time.Now().Format("2006-01-02 15:04:05"), triggeredBy, issueNumber, issueTitle, branch, projectName, logFileName)
	os.WriteFile(logFileName, []byte(logMsg), 0644)

	if issueTitle == "" {
		issueTitle = fmt.Sprintf("E2E Issue #%d", issueNumber)
	}

	e2eTask := &Task{
		ID:          taskID(TaskTypeE2E, issueNumber),
		Type:        TaskTypeE2E,
		Status:      "running",
		ProjectName: projectName,
		TargetID:    issueNumber,
		TargetTitle: issueTitle,
		Branch:      branch,
		StartTime:   time.Now(),
		LogFile:     logFileName,
		Metadata: map[string]interface{}{
			"triggeredBy": triggeredBy,
			"branch":      branch,
			"taskKind":    taskKind,
			"issueBody":   issueBody,
		},
	}

	if !setTask(e2eTask) {
		log.Printf("[e2e] E2E 任务注册失败（可能已存在）: e2e-%d", issueNumber)
		return
	}

	debugDir := filepath.Join(logsDir, fmt.Sprintf("e2e-screenshots-%s", filepath.Base(logFileName)))
	debugDir = strings.TrimSuffix(debugDir, ".log")
	reportFile := filepath.Join(logsDir, fmt.Sprintf("e2e-report-%s.md", strings.TrimSuffix(filepath.Base(logFileName), ".log")))

	// 准备 E2E Agent 配置
	agentFile := prepareAgentConfig("e2e-tester")
	if agentFile == "" {
		agentFile = prepareAgentConfig("developer")
	}

	agentCtx := buildAgentContext(worktreePath, issueNumber, 0)
	ctxPrompt := formatAgentContextPrompt(agentCtx, issueNumber, 0)

	// 构建 E2E prompt，包含 issue body 中定义的测试范围
	var scopeSection string
	if issueBody != "" {
		scopeSection = fmt.Sprintf("\n## E2E 测试范围（来自 Issue #%d）\n\n%s\n", issueNumber, issueBody)
	}
	e2ePrompt := fmt.Sprintf("%s\n\n**你的任务类型：E2E 测试任务（E2E Tester Agent）**\n目标：执行 Maestro E2E 测试，分析失败原因，生成详细报告。\n\nIssue #%d: %s\n分支: %s\n%s", ctxPrompt, issueNumber, issueTitle, branch, scopeSection)

	e2eVars := map[string]interface{}{
		"issueNumber":  issueNumber,
		"issueTitle":   issueTitle,
		"branch":       branch,
		"logFileName":  logFileName,
		"taskKind":     taskKind,
		"worktreePath": worktreePath,
		"agentFile":    agentFile,
		"debugDir":     debugDir,
		"reportFile":   reportFile,
		"prompt":       e2ePrompt,
	}

	runTaskWorkflow(e2eTask, e2eVars, nil)
	log.Printf("[e2e] E2E 任务已启动: e2e-%d", issueNumber)
}

// startE2ETaskFromHook 从 hook 启动 E2E 任务（保留兼容性，但 E2E 已改为独立 issue 驱动）
func startE2ETaskFromHook(sourceTask *Task, vars map[string]interface{}) {
	branch, _ := vars["branch"].(string)
	taskKind, _ := vars["taskKind"].(string)
	worktreePath, _ := vars["tmpDir"].(string)

	// 直接从 sourceTask 取 issueNumber，不依赖模板渲染（更可靠）
	issueNum := sourceTask.TargetID
	if issueNum == 0 {
		// fallback：尝试从 vars 中读取
		issueNumberStr, _ := vars["issueNumber"].(string)
		if issueNumberStr == "" {
			log.Printf("[hook] E2E 任务缺少 issueNumber")
			return
		}
		var err error
		issueNum, err = strconv.Atoi(issueNumberStr)
		if err != nil || issueNum == 0 {
			log.Printf("[hook] E2E 任务 issueNumber 无效: %s", issueNumberStr)
			return
		}
	}

	triggeredBy := fmt.Sprintf("%s-%d", sourceTask.Type, sourceTask.TargetID)
	// Hook 触发时无 issue title/body，传空值（E2E 已主要走独立 issue 手动触发）
	startE2ETask(issueNum, "", "", branch, taskKind, triggeredBy, worktreePath, sourceTask.ProjectName)
}

// startReviewTaskFromHook E2E 成功后自动触发 Review
func startReviewTaskFromHook(sourceTask *Task, vars map[string]interface{}) {
	branch := sourceTask.Branch
	if branch == "" {
		log.Printf("[hook] E2E 任务缺少 branch，无法触发 Review")
		return
	}

	prNumber := findPRForBranch(branch)
	if prNumber == 0 {
		log.Printf("[hook] E2E 任务分支 %s 未找到关联 PR，跳过自动 Review", branch)
		return
	}

	log.Printf("[hook] E2E 任务 e2e-%d 结束，将自动触发 PR #%d Review", sourceTask.TargetID, prNumber)
	startReviewTask(prNumber, branch, sourceTask.ProjectName)
}

// startReworkTaskFromHook Review 未通过时自动触发 Rework
func startReworkTaskFromHook(sourceTask *Task, vars map[string]interface{}) {
	prNumber := sourceTask.TargetID
	if prNumber == 0 {
		log.Printf("[hook] 任务缺少 TargetID，无法触发 Rework")
		return
	}

	branch := sourceTask.Branch
	reportStr := sourceTask.Report

	// 如果 sourceTask 是 DocGardener，从 Metadata 读取 review report
	if reportStr == "" && sourceTask.Type == TaskTypeDocGardener {
		if m, ok := sourceTask.Metadata["reviewReport"].(string); ok && m != "" {
			reportStr = m
			log.Printf("[hook] 从 DocGardener Metadata 读取 review report (%d bytes)", len(reportStr))
		}
	}

	log.Printf("[hook] PR #%d 将自动触发 Rework (来源: %s)", prNumber, sourceTask.Type)
	startReworkTask(prNumber, branch, reportStr, sourceTask.ProjectName)
}

// mapKeys 返回 map 的所有 key（用于日志）
func mapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// generateE2EReport 生成 E2E 测试报告
// 优先使用 e2e.yaml 生成的详细报告，回退到旧版日志报告
func generateE2EReport(e2eTask *Task) (string, string) {
	logsDir := filepath.Dir(e2eTask.LogFile)
	logBase := strings.TrimSuffix(filepath.Base(e2eTask.LogFile), ".log")

	// 1. 优先读取 e2e.yaml 生成的详细步骤级报告
	detailedReportPath := filepath.Join(logsDir, fmt.Sprintf("e2e-report-%s.md", logBase))
	if data, err := os.ReadFile(detailedReportPath); err == nil && len(data) > 0 {
		log.Printf("[e2e] 使用详细报告: %s", detailedReportPath)
		return detailedReportPath, string(data)
	}

	// 2. 回退：读取旧版日志内容生成简单报告
	logContent, err := os.ReadFile(e2eTask.LogFile)
	if err != nil || len(logContent) == 0 {
		return "", ""
	}

	fallbackPath := filepath.Join(logsDir, fmt.Sprintf("e2e-report-%d.md", e2eTask.TargetID))

	var statusEmoji string
	switch e2eTask.Status {
	case "success":
		statusEmoji = "✅ 通过"
	case "failed":
		statusEmoji = "❌ 失败"
	case "stopped":
		statusEmoji = "🛑 已停止"
	default:
		statusEmoji = "⚠️ " + e2eTask.Status
	}

	report := fmt.Sprintf("## 🤖 E2E 测试报告\n\n| 项目 | 内容 |\n|------|------|\n| Issue | #%d |\n| 分支 | `%s` |\n| 结果 | %s |\n| 时间 | %s |\n| 日志 | `%s` |\n\n### 详细日志\n\n%s\n\n---\n*由 Agent Control Center 自动生成*",
		e2eTask.TargetID,
		e2eTask.Branch,
		statusEmoji,
		e2eTask.EndTime.Format("2006-01-02 15:04:05"),
		filepath.Base(e2eTask.LogFile),
		string(logContent),
	)

	if err := os.WriteFile(fallbackPath, []byte(report), 0644); err != nil {
		log.Printf("[e2e] 生成报告失败: %v", err)
		return "", report
	}
	log.Printf("[e2e] 报告已生成: %s", fallbackPath)
	return fallbackPath, report
}

// findPRForBranch 通过分支名查找关联的 PR
func findPRForBranch(branch string) int {
	if branch == "" {
		return 0
	}
	platform := getPlatform(getProjectPath())
	iid, err := platform.FindMRByBranch(branch)
	if err != nil {
		log.Printf("[e2e] 查找 PR/MR 失败: %v", err)
		return 0
	}
	return iid
}

// hasMaestroFlows 检查指定分支是否包含 Maestro flow 文件
// 用于 dev → e2e hook 的条件判断
func hasMaestroFlows(branch string) bool {
	if branch == "" {
		return false
	}
	projectPath := getProjectPath()
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf("git ls-tree -r --name-only origin/%s .maestro/flows/ 2>/dev/null | head -1", branch))
	cmd.Dir = projectPath
	out, _ := cmd.Output()
	result := strings.TrimSpace(string(out)) != ""
	if result {
		log.Printf("[hook] 分支 origin/%s 包含 Maestro flows", branch)
	} else {
		log.Printf("[hook] 分支 origin/%s 无 Maestro flows", branch)
	}
	return result
}
