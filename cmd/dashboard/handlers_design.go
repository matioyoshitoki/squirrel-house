package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// handleStartDesign 启动 UI 设计任务
func handleStartDesign(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		IssueNumber int    `json:"issueNumber"`
		IssueTitle  string `json:"issueTitle"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	// 检查是否已有运行中的任务
	projectName := getCurrentProjectName()
	tkey := taskKey(projectName, TaskTypeDesign, req.IssueNumber)
	tasksMutex.RLock()
	if existingTask, ok := tasks[tkey]; ok && existingTask.Status == "running" {
		tasksMutex.RUnlock()
		writeJSONError(w, 409, fmt.Sprintf("Issue #%d 已在运行中", req.IssueNumber))
		return
	}
	tasksMutex.RUnlock()

	logFileName := startDesignTask(req.IssueNumber, req.IssueTitle)
	if logFileName == "" {
		writeJSONError(w, 500, "创建日志文件失败")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "started",
		"message": fmt.Sprintf("Issue #%d 设计任务已启动", req.IssueNumber),
		"logFile": logFileName,
	})
}

// startDesignTask 启动设计任务的核心逻辑
func startDesignTask(issueNumber int, issueTitle string) string {
	projectName := getCurrentProjectName()
	logsDir := getCurrentProjectLogsDir()
	os.MkdirAll(logsDir, 0755)

	timestamp := time.Now().Format("20060102-150405")
	logFileName := filepath.Join(logsDir, fmt.Sprintf("design-issue-%d-%s.log", issueNumber, timestamp))

	logFile, err := os.Create(logFileName)
	if err != nil {
		return ""
	}
	logFile.Close()

	go func() {
		projectPath := getProjectPath()
		logMsg := fmt.Sprintf("[%s] Issue #%d: 启动 UI 设计\n项目路径: %s\n日志文件: %s\n",
			time.Now().Format("2006-01-02 15:04:05"), issueNumber, projectPath, logFileName)
		os.WriteFile(logFileName, []byte(logMsg), 0644)

		// 确定基础分支
		currentBranch := getProjectDefaultBranch()
		if currentBranch == "" {
			currentBranch = "main"
		}

		branchName := fmt.Sprintf("design/issue-%d", issueNumber)
		tmpDir, _ := os.MkdirTemp("", fmt.Sprintf("design-issue-%d-*", issueNumber))
		if err := cleanWorktreeForBranch(projectPath, branchName); err != nil {
			appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ clean worktree failed: %v\n", err)))
			return
		}

		fetchCmd := exec.Command("git", "fetch", "origin")
		fetchCmd.Dir = projectPath
		fetchCmd.Run()

		baseRef := resolveBaseRef(projectPath, currentBranch)
		worktreeCmd := fmt.Sprintf("git worktree add -B %s %s %s", branchName, tmpDir, baseRef)
		wtCmd := exec.Command("bash", "-c", worktreeCmd)
		wtCmd.Dir = projectPath
		if output, err := wtCmd.CombinedOutput(); err != nil {
			appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ git worktree add failed: %v\n%s\n", err, string(output))))
			return
		}

		agentFile := prepareAgentConfig("ui-designer")
		if agentFile == "" {
			appendToFile(logFileName, []byte("\n❌ 未找到 ui-designer agent 配置\n"))
			return
		}

		agentCtx := buildAgentContext(tmpDir, issueNumber, 0)
		ctxPrompt := formatAgentContextPrompt(agentCtx, issueNumber, 0)
		prompt, _ := renderCtxTemplate("design_task", map[string]interface{}{
			"ContextPrompt": ctxPrompt,
			"IssueNumber":   issueNumber,
			"IssueTitle":    issueTitle,
			"Branch":        branchName,
		})

		task := &Task{
			ID:          taskID(TaskTypeDesign, issueNumber),
			Type:        TaskTypeDesign,
			Status:      "running",
			ProjectName: projectName,
			TargetID:    issueNumber,
			TargetTitle: issueTitle,
			Branch:      branchName,
			StartTime:   time.Now(),
			LogFile:     logFileName,
			Metadata: map[string]interface{}{
				"tmuxSession": fmt.Sprintf("nwops-%s", taskID(TaskTypeDesign, issueNumber)),
			},
		}
		if !setTask(task) {
			appendToFile(logFileName, []byte("\n❌ 设计任务已在运行中\n"))
			return
		}

		vars := map[string]interface{}{
			"issueNumber": issueNumber,
			"issueTitle":  issueTitle,
			"tmpDir":      tmpDir,
			"worktreeCmd": worktreeCmd,
			"baseBranch":  currentBranch,
			"branch":      branchName,
			"Branch":      branchName,
			"logFileName": logFileName,
			"agentFile":   agentFile,
			"prompt":      prompt,
			"logsDir":     logsDir,
			"projectName": projectName,
			"projectPath": projectPath,
		}
		for k, v := range agentCtx {
			vars[k] = v
		}
		runTaskWorkflow(task, vars, nil)
	}()

	return logFileName
}

// getProjectDefaultBranch 获取当前项目的默认分支
func getProjectDefaultBranch() string {
	projectMutex.RLock()
	idx := currentProjectIndex
	projects := config.Projects
	projectMutex.RUnlock()

	if idx < 0 || idx >= len(projects) {
		return "main"
	}
	b := strings.TrimSpace(projects[idx].DefaultBranch)
	if b == "" {
		return "main"
	}
	return b
}
