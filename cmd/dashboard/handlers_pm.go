package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// handleStartPM 启动 Product Manager 任务（创建 Issue + 产出 PRD）
func handleStartPM(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		IssueNumber int    `json:"issueNumber"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		writeJSONError(w, 400, "标题不能为空")
		return
	}

	issueNumber := req.IssueNumber
	// 如果未指定 issueNumber，先创建 GitHub Issue
	if issueNumber == 0 {
		title := req.Title
		if !strings.HasPrefix(title, "[") {
			title = "[FEAT] " + title
		}
		body := req.Description
		if body == "" {
			body = "由 Product Manager Agent 自动创建，即将产出 PRD。"
		}

		platform := getPlatform(getProjectPath())
		num, err := platform.CreateIssue(title, body, []string{"feature", "triage"})
		if err != nil {
			log.Printf("❌ 创建 Issue 失败: %v", err)
			writeJSONError(w, 500, fmt.Sprintf("创建 Issue 失败: %v", err))
			return
		}
		issueNumber = num
		log.Printf("✅ 已创建 Issue #%d", issueNumber)
	}

	// 检查是否已有运行中的 PM 任务
	projectName := getCurrentProjectName()
	tkey := taskKey(projectName, TaskTypePM, issueNumber)
	tasksMutex.RLock()
	if existingTask, ok := tasks[tkey]; ok && existingTask.Status == "running" {
		tasksMutex.RUnlock()
		writeJSONError(w, 409, fmt.Sprintf("Issue #%d 的 PRD 任务已在运行中", issueNumber))
		return
	}
	tasksMutex.RUnlock()

	logFileName := startPMTask(issueNumber, req.Title, req.Description)
	if logFileName == "" {
		writeJSONError(w, 500, "创建日志文件失败")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":      "started",
		"message":     fmt.Sprintf("Issue #%d PRD 任务已启动", issueNumber),
		"issueNumber": issueNumber,
		"logFile":     logFileName,
	})
}

// startPMTask 启动 PRD 产出任务的核心逻辑
func startPMTask(issueNumber int, issueTitle, issueDesc string) string {
	projectName := getCurrentProjectName()
	logsDir := getCurrentProjectLogsDir()
	os.MkdirAll(logsDir, 0755)

	timestamp := time.Now().Format("20060102-150405")
	logFileName := filepath.Join(logsDir, fmt.Sprintf("pm-issue-%d-%s.log", issueNumber, timestamp))

	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Printf("❌ 创建日志文件失败: %v", err)
		return ""
	}
	logFile.Close()

	logMsg := fmt.Sprintf("[%s] Issue #%d: 启动 PRD 产出\n项目路径: %s\n日志文件: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), issueNumber, getProjectPath(), logFileName)
	os.WriteFile(logFileName, []byte(logMsg), 0644)

	// 后台执行 git 准备和 workflow 启动，让 API 立即返回
	go func() {
		// 确定基础分支
		currentBranch := getProjectDefaultBranch()
		if currentBranch == "" {
			currentBranch = "main"
		}

		branchName := fmt.Sprintf("prd/issue-%d", issueNumber)
		tmpDir, _ := os.MkdirTemp("", fmt.Sprintf("pm-issue-%d-*", issueNumber))
		if err := cleanWorktreeForBranch(getProjectPath(), branchName); err != nil {
			appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ clean worktree failed: %v\n", err)))
			return
		}

		fetchCmd := exec.Command("git", "fetch", "origin")
		fetchCmd.Dir = getProjectPath()
		fetchCmd.Run()

		baseRef := resolveBaseRef(getProjectPath(), currentBranch)
		worktreeCmd := fmt.Sprintf("git worktree add -B %s %s %s", branchName, tmpDir, baseRef)
		wtCmd := exec.Command("bash", "-c", worktreeCmd)
		wtCmd.Dir = getProjectPath()
		wtCmd.Run()

		agentFile := prepareAgentConfig("product-manager")
		if agentFile == "" {
			appendToFile(logFileName, []byte("\n❌ 未找到 product-manager agent 配置\n"))
			return
		}

		project := getCurrentProject()
		basePrompt, err := renderPrompt("product-manager", project)
		if err != nil {
			appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ 渲染 prompt 失败: %v\n", err)))
			return
		}
		prompt := strings.ReplaceAll(basePrompt, "__ISSUE_NUMBER__", fmt.Sprintf("%d", issueNumber))
		prompt = strings.ReplaceAll(prompt, "__BRANCH__", branchName)

		ctxPrompt := formatAgentContextPrompt(buildAgentContext(tmpDir, issueNumber, 0), issueNumber, 0)
		userPrompt, _ := renderCtxTemplate("pm_task", map[string]interface{}{
			"ContextPrompt": ctxPrompt,
			"IssueNumber":   issueNumber,
			"IssueTitle":    issueTitle,
			"IssueDesc":     issueDesc,
		})
		fullPrompt := prompt + "\n\n" + userPrompt

		task := &Task{
			ID:          taskID(TaskTypePM, issueNumber),
			Type:        TaskTypePM,
			Status:      "running",
			ProjectName: projectName,
			TargetID:    issueNumber,
			TargetTitle: issueTitle,
			Branch:      branchName,
			StartTime:   time.Now(),
			LogFile:     logFileName,
			Metadata: map[string]interface{}{
				"issueTitle":  issueTitle,
				"issueDesc":   issueDesc,
				"tmuxSession": fmt.Sprintf("nwops-%s", taskID(TaskTypePM, issueNumber)),
			},
		}
		if !setTask(task) {
			appendToFile(logFileName, []byte("\n❌ PRD 任务已在运行中\n"))
			return
		}

		vars := map[string]interface{}{
			"issueNumber": issueNumber,
			"issueTitle":  issueTitle,
			"tmpDir":      tmpDir,
			"worktreeCmd": worktreeCmd,
			"branch":      branchName,
			"Branch":      branchName,
			"logFileName": logFileName,
			"agentFile":   agentFile,
			"prompt":      fullPrompt,
		}
		for k, v := range buildAgentContext(tmpDir, issueNumber, 0) {
			vars[k] = v
		}
		runTaskWorkflow(task, vars, nil)
	}()

	return logFileName
}

// getCurrentProject 获取当前项目配置
func getCurrentProject() Project {
	projectMutex.RLock()
	defer projectMutex.RUnlock()
	idx := currentProjectIndex
	if idx < 0 || idx >= len(config.Projects) {
		idx = 0
	}
	if len(config.Projects) > 0 {
		return config.Projects[idx]
	}
	return Project{Name: "default", Path: "."}
}
