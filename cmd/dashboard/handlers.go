package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// rebuildTaskStatus 从 state/tasks.json 和 .status.json 恢复任务状态
// 优先信任结构化状态文件，fallback 到日志推断
func rebuildTaskStatus() {
	projectName := getCurrentProjectName()
	logsDir := getCurrentProjectLogsDir()
	log.Printf("🔍 开始重建任务状态 [%s]...", projectName)

	// 1. 从持久化状态文件恢复全局任务状态
	loadTaskState()

	// 1b. 过滤掉非当前项目的任务，防止项目切换后状态串了
	tasksMutex.Lock()
	for k, v := range tasks {
		if v.ProjectName != projectName {
			delete(tasks, k)
		}
	}
	tasksMutex.Unlock()

	// 2. 优先从结构化状态文件 (.status.json) 恢复
	stateDir := "state"
	statusEntries, _ := os.ReadDir(stateDir)
	for _, entry := range statusEntries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "task-") || !strings.HasSuffix(entry.Name(), ".status.json") {
			continue
		}
		taskID := strings.TrimSuffix(strings.TrimPrefix(entry.Name(), "task-"), ".status.json")
		status := readTaskStatusFile(taskID)
		if status == nil || status.ProjectName != projectName {
			continue
		}

		// 解析 taskID 格式：type-targetID
		parts := strings.SplitN(taskID, "-", 2)
		if len(parts) != 2 {
			continue
		}
		var taskType TaskType
		switch parts[0] {
		case "dev":
			taskType = TaskTypeDev
		case "design":
			taskType = TaskTypeDesign
		case "rework":
			taskType = TaskTypeRework
		case "review":
			taskType = TaskTypeReview
		case "doc-gardener":
			taskType = TaskTypeDocGardener
		case "e2e":
			taskType = TaskTypeE2E
		case "pm":
			taskType = TaskTypePM
		default:
			continue
		}
		targetID := toInt(parts[1])
		if targetID == 0 {
			continue
		}
		key := taskKey(projectName, taskType, targetID)

		tasksMutex.Lock()
		if existing, ok := tasks[key]; ok {
			// 内存中已存在，用 .status.json 补充信息
			// 优先信任 .status.json 中的最终状态（success/failed），防止内存中的 interrupted 覆盖真实结果
			if existing.Status == "running" && status.Status != "running" {
				existing.Status = status.Status
				existing.EndTime = status.LastHeartbeat
				log.Printf("📝 从状态文件更新: %s -> %s", key, status.Status)
			} else if (existing.Status == "interrupted" || existing.Status == "stopped") && (status.Status == "success" || status.Status == "failed") {
				existing.Status = status.Status
				existing.EndTime = status.LastHeartbeat
				log.Printf("📝 从状态文件修正: %s %s -> %s", key, existing.Status, status.Status)
			}
			tasksMutex.Unlock()
			continue
		}
		tasksMutex.Unlock()

		// 内存中没有，从 .status.json 重建
		finalStatus := status.Status
		if finalStatus == "running" && time.Since(status.LastHeartbeat) > 3*time.Minute {
			// dashboard 重启后，tmux 任务可能仍在运行但心跳已旧
			// 如果 tmux session 仍然存在，保持 running 状态
			tmuxSession := fmt.Sprintf("nwops-%s", taskID)
			if err := exec.Command("tmux", "has-session", "-t", tmuxSession).Run(); err == nil {
				log.Printf("📝 %s 心跳超时但 tmux session 仍在运行，重建为 running", taskID)
			} else {
				finalStatus = "interrupted"
				log.Printf("📝 %s 心跳超时，从状态文件重建为 interrupted", taskID)
			}
		}


		if finalStatus != status.Status {
			updateTaskStatusFileStatus(taskID, finalStatus)
		}
		tasksMutex.Lock()
		tasks[key] = &Task{
			ID:          taskID,
			Type:        taskType,
			Status:      finalStatus,
			ProjectName: projectName,
			TargetID:    targetID,
			TargetTitle: status.TaskID,
			StartTime:   status.StartedAt,
			LogFile:     filepath.Join(logsDir, fmt.Sprintf("%s.log", taskID)),
		}
		tasksMutex.Unlock()
		log.Printf("📝 从状态文件重建: %s -> %s", key, finalStatus)
	}

	// 3. Fallback：对没有 .status.json 的任务，从日志推断
	entries, err := os.ReadDir(logsDir)
	if err != nil {
		os.MkdirAll(logsDir, 0755)
		return
	}

	latestLogs := make(map[string]os.FileInfo)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		info, err := entry.Info()
		if err != nil {
			continue
		}

		var key string
		if strings.HasPrefix(name, "dev-issue-") && strings.HasSuffix(name, ".log") {
			key = "dev-" + extractIDFromFilename(name, "dev-issue-", ".log")
		} else if strings.HasPrefix(name, "design-issue-") && strings.HasSuffix(name, ".log") {
			key = "design-" + extractIDFromFilename(name, "design-issue-", ".log")
		} else if strings.HasPrefix(name, "rework-pr-") && strings.HasSuffix(name, ".log") {
			key = "rework-" + extractIDFromFilename(name, "rework-pr-", ".log")
		} else if strings.HasPrefix(name, "doc-gardener-pr-") && strings.HasSuffix(name, ".log") {
			key = "doc-gardener-" + extractIDFromFilename(name, "doc-gardener-pr-", ".log")
		} else if strings.HasPrefix(name, "review-pr-") && strings.HasSuffix(name, ".log") {
			key = "review-" + extractIDFromFilename(name, "review-pr-", ".log")
		} else if strings.HasPrefix(name, "e2e-issue-") && strings.HasSuffix(name, ".log") {
			key = "e2e-" + extractIDFromFilename(name, "e2e-issue-", ".log")
		} else if strings.HasPrefix(name, "pm-issue-") && strings.HasSuffix(name, ".log") {
			key = "pm-" + extractIDFromFilename(name, "pm-issue-", ".log")
		}

		if key != "" {
			if existing, ok := latestLogs[key]; !ok || info.ModTime().After(existing.ModTime()) {
				latestLogs[key] = info
			}
		}
	}

	for key, info := range latestLogs {
		filename := info.Name()
		content, _ := os.ReadFile(filepath.Join(logsDir, filename))
		contentStr := string(content)
		status := inferStatusFromLog(contentStr, info.ModTime())

		parts := strings.SplitN(key, "-", 2)
		if len(parts) != 2 {
			continue
		}
		typeStr, idStr := parts[0], parts[1]
		targetID := toInt(idStr)
		if targetID == 0 {
			continue
		}

		var taskType TaskType
		var taskIDStr string
		var title string
		var branch string

		switch typeStr {
		case "dev":
			taskType = TaskTypeDev
			taskIDStr = taskID(taskType, targetID)
			title = fmt.Sprintf("Issue #%d", targetID)
		case "design":
			taskType = TaskTypeDesign
			taskIDStr = taskID(taskType, targetID)
			title = fmt.Sprintf("Issue #%d", targetID)
		case "doc-gardener":
			taskType = TaskTypeDocGardener
			taskIDStr = taskID(taskType, targetID)
			title = fmt.Sprintf("PR #%d", targetID)
		case "review":
			taskType = TaskTypeReview
			taskIDStr = taskID(taskType, targetID)
			title = fmt.Sprintf("PR #%d", targetID)
			branch = extractBranchFromLog(contentStr)
		case "rework":
			taskType = TaskTypeRework
			taskIDStr = taskID(taskType, targetID)
			title = fmt.Sprintf("PR #%d", targetID)
		case "e2e":
			taskType = TaskTypeE2E
			taskIDStr = taskID(taskType, targetID)
			title = fmt.Sprintf("E2E Issue #%d", targetID)
		case "pm":
			taskType = TaskTypePM
			taskIDStr = taskID(taskType, targetID)
			title = fmt.Sprintf("Issue #%d", targetID)
		}

		// 跳过已有 .status.json 的任务
		taskKey := taskKey(projectName, taskType, targetID)
		tasksMutex.RLock()
		_, alreadyExists := tasks[taskKey]
		tasksMutex.RUnlock()
		if alreadyExists {
			continue
		}

		// 检查 tmux session 是否存在（长时间任务在 tmux 中运行）
		tmuxSession := fmt.Sprintf("nwops-%s", taskIDStr)
		tmuxAlive := isTmuxSessionAlive(tmuxSession)

		tasksMutex.Lock()
		if existing, ok := tasks[taskKey]; ok && existing.ProjectName == projectName {
			// 内存中已存在该任务，优先信任持久化状态
			if existing.Status == "running" {
				if tmuxAlive {
					// tmux session 还在运行，保持 running
					existing.LogFile = filepath.Join(logsDir, filename)
					if existing.Metadata == nil {
						existing.Metadata = make(map[string]interface{})
					}
					existing.Metadata["tmuxSession"] = tmuxSession
					log.Printf("📝 更新任务状态: %s -> running (tmux session 存活)", taskIDStr)
				} else if status == "success" || status == "failed" || status == "stopped" {
					// 只有当日志有明确的终止标记时才覆盖 running
					existing.Status = status
					existing.EndTime = info.ModTime()
					existing.LogFile = filepath.Join(logsDir, filename)
					log.Printf("📝 更新任务状态: %s -> %s (来自日志终止标记)", taskIDStr, status)
				}
				// 其他情况（interrupted/unknown）保持 running
			} else {
				// 非 running 状态可用日志推断补充
				existing.Status = status
				existing.LogFile = filepath.Join(logsDir, filename)
				log.Printf("📝 更新任务状态: %s -> %s (来自日志推断)", taskIDStr, status)
			}
		} else {
			// 内存中没有，从日志新建
			if tmuxAlive {
				status = "running" // tmux 还在运行，状态应为 running
			}
			tasks[taskKey] = &Task{
				ID:          taskIDStr,
				Type:        taskType,
				Status:      status,
				ProjectName: projectName,
				TargetID:    targetID,
				TargetTitle: title,
				Branch:      branch,
				StartTime:   info.ModTime(),
				LogFile:     filepath.Join(logsDir, filename),
				Metadata: map[string]interface{}{
					"tmuxSession": tmuxSession,
				},
			}
			log.Printf("📝 重建任务状态: %s -> %s", taskKey, status)
		}
		tasksMutex.Unlock()
	}

	// 4. 对仍然 running 的任务，检查心跳（优先 .status.json，fallback 日志修改时间）
	tasksMutex.Lock()
	for _, task := range tasks {
		if task.Status != "running" {
			continue
		}
		// 有 tmux session 的任务优先检查 tmux 存活
		if tmuxSession := getTaskTmuxSession(task); tmuxSession != "" {
			if isTmuxSessionAlive(tmuxSession) {
				continue
			}
		}
		// 尝试从 .status.json 读取心跳
		status := readTaskStatusFile(task.ID)
		if status != nil {
			if time.Since(status.LastHeartbeat) > 3*time.Minute {
				task.Status = "interrupted"
				appendToFile(task.LogFile, []byte("\n⏸️ 任务心跳超时，被标记为中断\n"))
				log.Printf("🔄 重启后恢复的任务心跳超时: %s", task.ID)
			}
			continue
		}
		// fallback：日志修改时间（旧逻辑，仅对无 .status.json 的任务）
		info, err := os.Stat(task.LogFile)
		if err == nil && time.Since(info.ModTime()) > 3*time.Minute {
			task.Status = "interrupted"
			appendToFile(task.LogFile, []byte("\n⏸️ 任务因 Dashboard 重启被标记为中断\n"))
			log.Printf("🔄 重启后恢复的任务标记为中断: %s", task.ID)
		}
	}
	tasksMutex.Unlock()

	// 5. 从 review-report-*.md 补充 review 任务的 report 和 result
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "review-report-") || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		info, _ := entry.Info()
		prNumStr := strings.TrimPrefix(strings.TrimSuffix(entry.Name(), ".md"), "review-report-")
		prNum := toInt(prNumStr)
		if prNum == 0 {
			continue
		}

		content, _ := os.ReadFile(filepath.Join(logsDir, entry.Name()))
		contentStr := string(content)
		result := extractReviewResult(contentStr)

		logFile := filepath.Join(logsDir, fmt.Sprintf("review-pr-%d-*.log", prNum))
		matches, _ := filepath.Glob(logFile)
		actualLogFile := logFile
		latestLogTime := time.Time{}
		if len(matches) > 0 {
			latestLog := matches[0]
			latestInfo, _ := os.Stat(latestLog)
			for _, m := range matches[1:] {
				if li, err := os.Stat(m); err == nil && li.ModTime().After(latestInfo.ModTime()) {
					latestLog = m
					latestInfo = li
				}
			}
			actualLogFile = latestLog
			latestLogTime = latestInfo.ModTime()
		}

		if !latestLogTime.IsZero() && info.ModTime().Before(latestLogTime) {
			log.Printf("📝 跳过过期的 Review Report: PR #%d", prNum)
			continue
		}

		tid := taskID(TaskTypeReview, prNum)
		tkey := taskKey(projectName, TaskTypeReview, prNum)
		tasksMutex.Lock()
		if t, ok := tasks[tkey]; ok && t.Type == TaskTypeReview {
			t.Report = contentStr
			t.Result = result
			if t.Status != "running" {
				t.Status = "success"
				t.EndTime = info.ModTime()
			}
		} else {
			tasks[tkey] = &Task{
				ID:          tid,
				Type:        TaskTypeReview,
				Status:      "success",
				ProjectName: projectName,
				TargetID:    prNum,
				TargetTitle: fmt.Sprintf("PR #%d", prNum),
				Result:      result,
				Report:      contentStr,
				StartTime:   info.ModTime().Add(-5 * time.Minute),
				EndTime:     info.ModTime(),
				LogFile:     actualLogFile,
			}
		}
		tasksMutex.Unlock()
		log.Printf("📝 重建 Review 任务: PR #%d -> %s", prNum, result)
	}

	// 6. 从 review-state-*.json 恢复返工状态
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "review-state-") || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		info, _ := entry.Info()
		prNumStr := strings.TrimPrefix(strings.TrimSuffix(entry.Name(), ".json"), "review-state-")
		prNum := toInt(prNumStr)
		if prNum == 0 {
			continue
		}

		stateData, _ := os.ReadFile(filepath.Join(logsDir, entry.Name()))
		var state struct {
			Status  string    `json:"status"`
			Result  string    `json:"result"`
			EndTime time.Time `json:"endTime"`
		}
		json.Unmarshal(stateData, &state)

		stateMTime := info.ModTime()
		reportPath := filepath.Join(logsDir, fmt.Sprintf("review-report-%d.md", prNum))
		if ri, err := os.Stat(reportPath); err == nil && ri.ModTime().After(stateMTime) {
			continue
		}
		logMatches, _ := filepath.Glob(filepath.Join(logsDir, fmt.Sprintf("review-pr-%d-*.log", prNum)))
		latestLogMTime := time.Time{}
		for _, lm := range logMatches {
			if li, err := os.Stat(lm); err == nil && li.ModTime().After(latestLogMTime) {
				latestLogMTime = li.ModTime()
			}
		}
		if latestLogMTime.After(stateMTime) {
			continue
		}

		tid := taskID(TaskTypeReview, prNum)
		tkey := taskKey(projectName, TaskTypeReview, prNum)
		tasksMutex.Lock()
		if t, ok := tasks[tkey]; ok && t.Type == TaskTypeReview {
			// 仅在持久化状态缺失或更旧时补充
			if t.Status == "running" || t.Status == "" {
				t.Status = state.Status
				t.Result = state.Result
				if !state.EndTime.IsZero() {
					t.EndTime = state.EndTime
				}
			}
		} else {
			tasks[tkey] = &Task{
				ID:          tid,
				Type:        TaskTypeReview,
				Status:      state.Status,
				ProjectName: projectName,
				TargetID:    prNum,
				TargetTitle: fmt.Sprintf("PR #%d", prNum),
				Result:      state.Result,
				EndTime:     state.EndTime,
				LogFile:     filepath.Join(logsDir, fmt.Sprintf("review-pr-%d-*.log", prNum)),
			}
		}
		tasksMutex.Unlock()
		log.Printf("📝 恢复 Review 返工状态: PR #%d -> %s", prNum, state.Status)
	}

	go saveTaskState()
	log.Printf("✅ 任务状态重建完成，当前 %d 个任务", len(tasks))
}

func extractIDFromFilename(name, prefix, suffix string) string {
	core := strings.TrimPrefix(name, prefix)
	core = strings.TrimSuffix(core, suffix)
	// core 格式可能是 "7-20260414-092845"，取第一部分
	parts := strings.Split(core, "-")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func matchIssueNumberInFilename(name string, issueNumber int) bool {
	numStr := strconv.Itoa(issueNumber)
	idx := strings.Index(name, numStr)
	if idx == -1 {
		return false
	}
	// 检查数字前后是否是非数字字符（或边界）
	before := idx > 0 && name[idx-1] >= '0' && name[idx-1] <= '9'
	afterEnd := idx + len(numStr)
	after := afterEnd < len(name) && name[afterEnd] >= '0' && name[afterEnd] <= '9'
	return !before && !after
}

func buildAgentContext(worktreePath string, issueNumber, prNumber int) map[string]string {
	ctx := make(map[string]string)
	ctx["AGENT_PROJECT_PATH"] = worktreePath

	// 静态项目文件（AGENTS.md、PRD、exec-plans、docs）从项目根目录读取，
	// 因为 worktree 临时目录（design/issue-N 分支）可能不包含这些文件。
	projectRoot := getProjectPath()
	// 确保项目根目录是绝对路径，避免 agent 在 worktree 中解析相对路径失败
	if absPath, err := filepath.Abs(projectRoot); err == nil {
		projectRoot = absPath
	}

	// AGENTS.md
	agentsMd := filepath.Join(projectRoot, "AGENTS.md")
	if _, err := os.Stat(agentsMd); err == nil {
		ctx["AGENT_AGENTS_MD"] = agentsMd
	}

	// PRD 文件
	prdDir := filepath.Join(projectRoot, "prd")
	if entries, err := os.ReadDir(prdDir); err == nil {
		var prdFiles []string
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if strings.HasSuffix(name, ".md") {
				if issueNumber > 0 && matchIssueNumberInFilename(name, issueNumber) {
					prdFiles = append(prdFiles, filepath.Join(prdDir, name))
				}
			}
		}
		// fallback: 最多 3 个最近的 prd md
		if len(prdFiles) == 0 {
			var allPrds []os.DirEntry
			for _, entry := range entries {
				if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
					allPrds = append(allPrds, entry)
				}
			}
			if len(allPrds) > 3 {
				sort.Slice(allPrds, func(i, j int) bool {
					fi, _ := allPrds[i].Info()
					fj, _ := allPrds[j].Info()
					return fi.ModTime().After(fj.ModTime())
				})
				allPrds = allPrds[:3]
			}
			for _, entry := range allPrds {
				prdFiles = append(prdFiles, filepath.Join(prdDir, entry.Name()))
			}
		}
		if len(prdFiles) > 0 {
			ctx["AGENT_PRD_FILES"] = strings.Join(prdFiles, ",")
		}
	}

	// 执行计划
	plansDir := filepath.Join(projectRoot, "docs", "exec-plans")
	if entries, err := os.ReadDir(plansDir); err == nil {
		var plans []string
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if strings.HasPrefix(name, "PLAN-") && strings.HasSuffix(name, ".md") {
				if issueNumber > 0 && matchIssueNumberInFilename(name, issueNumber) {
					plans = append(plans, filepath.Join(plansDir, name))
				}
			}
		}
		if len(plans) > 0 {
			ctx["AGENT_EXEC_PLANS"] = strings.Join(plans, ",")
		}
	}

	// 设计资产 — 扫描 worktree 中 designs/issue-N/ 目录下的所有文件
	// （设计资产是在 worktree 中生成的，优先从 worktreePath 读取）
	// 如果 worktree 中没有（如开发任务基于 main 分支），回退到 dashboard logs 目录
	if issueNumber > 0 {
		var designDir string
		worktreeDesignDir := filepath.Join(worktreePath, "designs", fmt.Sprintf("issue-%d", issueNumber))
		if info, err := os.Stat(worktreeDesignDir); err == nil && info.IsDir() {
			designDir = worktreeDesignDir
		} else {
			// fallback: 从 dashboard logs 目录读取（设计任务 copyDesignAssets 步骤复制过来的）
			logsDesignDir := filepath.Join(getProjectLogsDir(getCurrentProjectName()), "designs", fmt.Sprintf("issue-%d", issueNumber))
			if info, err := os.Stat(logsDesignDir); err == nil && info.IsDir() {
				designDir = logsDesignDir
			}
		}
		if designDir != "" {
			ctx["AGENT_DESIGN_DIR"] = designDir
			entries, _ := os.ReadDir(designDir)
			var assets []string
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				name := entry.Name()
				assets = append(assets, name)
				// 保留特定类型的 backward compatible keys
				switch name {
				case "design-spec.md":
					ctx["AGENT_DESIGN_SPEC"] = filepath.Join(designDir, name)
				case "mockup.html":
					ctx["AGENT_DESIGN_MOCKUP"] = filepath.Join(designDir, name)
				case "wireframe.svg":
					ctx["AGENT_DESIGN_WIREFRAME"] = filepath.Join(designDir, name)
				case "docs-to-update.txt":
					ctx["AGENT_DESIGN_DOC_LIST"] = filepath.Join(designDir, name)
				}
			}
			if len(assets) > 0 {
				ctx["AGENT_DESIGN_ASSETS"] = strings.Join(assets, ", ")
			}
		}
	}

	// 文档目录 — 从项目根目录读取
	docsDir := filepath.Join(projectRoot, "docs")
	if info, err := os.Stat(docsDir); err == nil && info.IsDir() {
		ctx["AGENT_DOCS_DIR"] = docsDir
	}

	// 设计反馈 — 从 dashboard logs 目录读取
	if issueNumber > 0 {
		logsDir := getProjectLogsDir(getCurrentProjectName())
		feedbackPath := filepath.Join(logsDir, "designs", fmt.Sprintf("issue-%d", issueNumber), "design-feedback.md")
		if _, err := os.Stat(feedbackPath); err == nil {
			ctx["AGENT_DESIGN_FEEDBACK"] = feedbackPath
		}
	}

	// 历史 review report — 从 dashboard logs 目录查找带时间戳的版本
	// 供后续 review agent 执行「验证上一轮问题是否修复」时读取
	if prNumber > 0 {
		logsDir := getProjectLogsDir(getCurrentProjectName())
		prefix := fmt.Sprintf("review-report-%d-", prNumber)
		entries, _ := os.ReadDir(logsDir)
		var latestReport string
		var latestTime time.Time
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasPrefix(entry.Name(), prefix) || !strings.HasSuffix(entry.Name(), ".md") {
				continue
			}
			info, _ := entry.Info()
			if info != nil && info.ModTime().After(latestTime) {
				latestTime = info.ModTime()
				latestReport = filepath.Join(logsDir, entry.Name())
			}
		}
		if latestReport != "" {
			ctx["AGENT_PREV_REVIEW_REPORT"] = latestReport
		}
	}

	return ctx
}

func formatAgentContextPrompt(ctx map[string]string, issueNumber, prNumber int) string {
	data := map[string]interface{}{
		"IssueNumber":          issueNumber,
		"PRNumber":             prNumber,
		"ProjectPath":          ctx["AGENT_PROJECT_PATH"],
		"AgentsMd":             ctx["AGENT_AGENTS_MD"],
		"PRDFiles":             ctx["AGENT_PRD_FILES"],
		"ExecPlans":            ctx["AGENT_EXEC_PLANS"],
		"DesignDir":            ctx["AGENT_DESIGN_DIR"],
		"DesignAssets":         ctx["AGENT_DESIGN_ASSETS"],
		"DocsDir":              ctx["AGENT_DOCS_DIR"],
		"PrevReviewReport":     ctx["AGENT_PREV_REVIEW_REPORT"],
		"DesignFeedback":       ctx["AGENT_DESIGN_FEEDBACK"],
	}
	s, err := renderCtxTemplate("agent_context", data)
	if err != nil {
		log.Printf("⚠️ render agent_context failed: %v, using fallback", err)
		return fallbackAgentContextPrompt(ctx, issueNumber, prNumber)
	}
	return s + "\n"
}

// fallbackAgentContextPrompt 模板渲染失败时的备用实现
func fallbackAgentContextPrompt(ctx map[string]string, issueNumber, prNumber int) string {
	var b strings.Builder
	b.WriteString("【任务上下文】\n")
	if issueNumber > 0 {
		b.WriteString(fmt.Sprintf("- Issue 编号: #%d\n", issueNumber))
	}
	if prNumber > 0 {
		b.WriteString(fmt.Sprintf("- PR 编号: #%d\n", prNumber))
	}
	if p, ok := ctx["AGENT_PROJECT_PATH"]; ok {
		b.WriteString(fmt.Sprintf("- 项目路径: %s\n", p))
	}
	if p, ok := ctx["AGENT_AGENTS_MD"]; ok {
		b.WriteString(fmt.Sprintf("- 项目地图: %s\n", p))
	} else {
		b.WriteString("- 项目地图: 未找到 AGENTS.md\n")
	}
	if p, ok := ctx["AGENT_PRD_FILES"]; ok {
		b.WriteString(fmt.Sprintf("- 关联 PRD: %s\n", p))
	}
	if p, ok := ctx["AGENT_EXEC_PLANS"]; ok {
		b.WriteString(fmt.Sprintf("- 执行计划: %s\n", p))
	}
	if p, ok := ctx["AGENT_DESIGN_DIR"]; ok {
		parts := []string{fmt.Sprintf("设计资产目录: %s", p)}
		if assets, ok := ctx["AGENT_DESIGN_ASSETS"]; ok {
			parts = append(parts, fmt.Sprintf("包含: %s", assets))
		} else {
			parts = append(parts, "目录为空")
		}
		b.WriteString(fmt.Sprintf("- %s\n", strings.Join(parts, "，")))
	}
	if p, ok := ctx["AGENT_DOCS_DIR"]; ok {
		b.WriteString(fmt.Sprintf("- 文档目录: %s\n", p))
	}
	if p, ok := ctx["AGENT_PREV_REVIEW_REPORT"]; ok {
		b.WriteString(fmt.Sprintf("- 上一轮 review report: %s\n", p))
	}
	if issueNumber > 0 && ctx["AGENT_DESIGN_DIR"] == "" {
		b.WriteString("- 设计资产: 未找到\n")
	}
	b.WriteString("\n")
	return b.String()
}

func clearReviewFailedState(prNumber int, projectName string) {
	logsDir := getProjectLogsDir(projectName)
	stateFile := filepath.Join(logsDir, fmt.Sprintf("review-state-%d.json", prNumber))
	os.Remove(stateFile)

	// 删除当前 review report，避免被 rebuildTaskStatus 恢复为旧状态
	// 但保留带时间戳的历史版本，供下一轮 review agent 读取以验证修复情况
	reportFile := filepath.Join(logsDir, fmt.Sprintf("review-report-%d.md", prNumber))
	os.Remove(reportFile)

	// 注意：不删除 review-report-<pr>-<timestamp>.md 历史文件
	// 这些历史报告是后续 review agent 执行「验证上一轮问题是否修复」的关键上下文
	// rebuildTaskStatus 已有「过期跳过」机制（report 修改时间早于最新 log 则跳过），
	// 不会将历史 report 错误恢复为当前状态。

	tid := taskID(TaskTypeReview, prNumber)
	tasksMutex.Lock()
	delete(tasks, tid)
	tasksMutex.Unlock()
	go saveTaskState()
	log.Printf("✅ PR #%d 的 review 状态已重置（返工完成，需重新 Review），历史 review report 已保留", prNumber)
}

func inferStatusFromLog(content string, modTime time.Time) string {
	// 1. 从文件末尾反向搜索终止状态标记，找到最后一次出现的标记（最准确）
	// 避免只检查最后 1000 字节（agent 输出可能很长，把标记推到前面）
	markers := []struct {
		mark   string
		status string
	}{
		{"✅ 任务完成", "success"},
		{"✅ Review 完成", "success"},
		{"✅ 开发完成", "success"},
		{"✅ 设计完成", "success"},
		{"✅ 返工完成", "success"},
		{"✅ 文档维护完成", "success"},
		{"⏹️ 任务已手动停止", "stopped"},
		{"⏹️ 已停止", "stopped"},
		{"⏸️ 任务因长时间无响应被标记为中断", "interrupted"},
		{"⏸️ 中断", "interrupted"},
	}

	lastPos := -1
	lastStatus := ""
	for _, m := range markers {
		if idx := strings.LastIndex(content, m.mark); idx != -1 {
			if idx > lastPos {
				lastPos = idx
				lastStatus = m.status
			}
		}
	}
	if lastStatus != "" {
		return lastStatus
	}

	// 2. 检查失败标记（必须是 dashboard 显式追加的终止行）
	failMarkers := []string{"\n❌ ", "Max number of steps reached", "步骤限制"}
	for _, m := range failMarkers {
		if strings.Contains(content, m) {
			// 确认包含 "失败:" 才判定为 failed（避免代码 diff 中的 ❌ 误判）
			if m == "\n❌ " {
				if strings.Contains(content, "失败:") {
					return "failed"
				}
				continue
			}
			return "failed"
		}
	}

	// 3. 无明确终止标记时的推断（kimi 任务可能运行很久，日志写入不频繁）
	if len(content) > 100 {
		if time.Since(modTime) < 30*time.Minute {
			return "running"
		}
		return "interrupted"
	}
	return "unknown"
}

func extractBranchFromLog(content string) string {
	if idx := strings.Index(content, "分支:"); idx != -1 {
		line := content[idx:]
		if endIdx := strings.Index(line, "\n"); endIdx != -1 {
			return strings.TrimSpace(strings.TrimPrefix(line[:endIdx], "分支:"))
		}
	}
	return "unknown"
}

func extractIssueNumberFromLog(content string) int {
	if idx := strings.Index(content, "Issue: #"); idx != -1 {
		after := content[idx+len("Issue: #"):]
		var n int
		fmt.Sscanf(after, "%d", &n)
		return n
	}
	return 0
}


func handleDesignAssets(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	issueNumber := r.URL.Query().Get("issueNumber")
	if issueNumber == "" {
		writeJSONError(w, http.StatusBadRequest, "Missing issueNumber")
		return
	}

	logsDir := getCurrentProjectLogsDir()
	assetsDir := filepath.Join(logsDir, "designs", "issue-"+issueNumber)

	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"issueNumber": issueNumber,
			"projectName": getCurrentProjectName(),
			"assets":      []interface{}{},
		})
		return
	}

	projectName := getCurrentProjectName()
	var assets []map[string]string

	// 优先检查是否已有编译好的 Flutter Web 预览
	flutterWebIndex := filepath.Join(assetsDir, "web", "index.html")
	if _, err := os.Stat(flutterWebIndex); err == nil {
		assets = append(assets, map[string]string{
			"name": "Flutter Web 预览",
			"url":  fmt.Sprintf("/logs/%s/designs/issue-%s/web/", projectName, issueNumber),
			"type": "flutter-web",
		})
	}

	// 检查 Phaser3 / Web 预览
	phaserWebIndex := filepath.Join(assetsDir, "preview", "index.html")
	if _, err := os.Stat(phaserWebIndex); err == nil {
		assets = append(assets, map[string]string{
			"name": "Phaser3 预览",
			"url":  fmt.Sprintf("/logs/%s/designs/issue-%s/preview/", projectName, issueNumber),
			"type": "phaser3-web",
		})
	}

	// 递归扫描 designs/issue-N/ 目录下的所有文件（包括子目录）
	filepath.WalkDir(assetsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		rel, _ := filepath.Rel(assetsDir, path)
		name := filepath.ToSlash(rel) // 统一使用正斜杠作为路径分隔符
		// 跳过 web/ 目录及其所有子目录（Flutter Web 编译产物已单独处理）
		if name == "web" || strings.HasPrefix(name, "web/") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		// 跳过 preview/ 目录及其所有子目录（Phaser3 Web 预览已单独处理）
		if name == "preview" || strings.HasPrefix(name, "preview/") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(name))
		assetType := "file"
		switch ext {
		case ".svg", ".png", ".jpg", ".jpeg", ".gif", ".webp":
			assetType = "image"
		case ".html", ".htm":
			assetType = "html"
		case ".md":
			assetType = "markdown"
		case ".mp3", ".ogg", ".wav":
			assetType = "audio"
		case ".fnt":
			assetType = "font"
		case ".json":
			assetType = "json"
		case ".dart", ".ts", ".tsx", ".js", ".jsx", ".css", ".scss", ".less", ".yaml", ".yml", ".py", ".go", ".rs", ".java", ".kt", ".swift":
			assetType = "code"
		}
		urlPath := fmt.Sprintf("/logs/%s/designs/issue-%s/%s", projectName, issueNumber, name)
		assets = append(assets, map[string]string{
			"name": name,
			"url":  urlPath,
			"type": assetType,
		})
		return nil
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"issueNumber": issueNumber,
		"projectName": projectName,
		"assets":      assets,
	})
}

func writeJSONError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func cleanZombieTasks() {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		tasksMutex.Lock()
		for id, task := range tasks {
			if task.Status != "running" || task.LogFile == "" {
				continue
			}
			if info, err := os.Stat(task.LogFile); err == nil {
				if time.Since(info.ModTime()) > 30*time.Minute {
					task.Status = "interrupted"
					updateTaskStatusFileStatus(task.ID, "interrupted")
					appendToFile(task.LogFile, []byte("\n⏸️ 任务因长时间无响应被标记为中断\n"))
					if task.Cancel != nil {
						task.Cancel()
					}
					if task.Kill != nil {
						task.Kill()
					}
					log.Printf("🧹 清理僵尸任务: %s", id)
					go broadcastTaskChanged(task, "interrupted")
				}
			}
		}
		tasksMutex.Unlock()
	}
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":    "Agent Control Center",
		"Projects": config.Projects,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func handleAPIIssues(w http.ResponseWriter, r *http.Request) {
	platform := getPlatform(getProjectPath())
	issues, err := platform.ListOpenIssues()
	if err != nil {
		log.Printf("⚠️ 获取 issues 失败: %v", err)
		writeJSONError(w, 500, fmt.Sprintf("获取 issues 失败: %v", err))
		return
	}

	for i := range issues {
		issues[i].Classification = classifyIssue(issues[i])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issues)
}

func classifyIssue(issue Issue) string {
	labelNames := make([]string, 0, len(issue.Labels))
	for _, l := range issue.Labels {
		labelNames = append(labelNames, strings.ToLower(l.Name))
	}
	labels := strings.Join(labelNames, ",")
	titleBody := strings.ToLower(issue.Title + " " + issue.Body)

	backendKeywords := []string{"backend", "api", "server", "db", "database", "infra", "nestjs", "migration"}
	frontendKeywords := []string{"frontend", "ui", "mobile", "design", "flutter", "react", "vue", "component", "screen", "page", "widget"}

	backendScore := 0
	frontendScore := 0

	for _, kw := range backendKeywords {
		if strings.Contains(labels, kw) || strings.Contains(titleBody, kw) {
			backendScore++
		}
	}
	for _, kw := range frontendKeywords {
		if strings.Contains(labels, kw) || strings.Contains(titleBody, kw) {
			frontendScore++
		}
	}

	if frontendScore > 0 && backendScore == 0 {
		return "frontend-only"
	}
	if backendScore > 0 && frontendScore == 0 {
		return "backend-only"
	}
	return "full-stack"
}

func handleRunningTasks(w http.ResponseWriter, r *http.Request) {
	currentProject := getCurrentProjectName()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(runningTasksByProject(currentProject))
}

func handleReviewTasks(w http.ResponseWriter, r *http.Request) {
	currentProject := getCurrentProjectName()
	reviewTasks := tasksByType(TaskTypeReview, currentProject)
	// 重新解析 report 中的 result，修复旧版本代码错误解析导致的历史数据
	needSave := false
	for _, t := range reviewTasks {
		if t.Report != "" {
			parsed := extractReviewResult(t.Report)
			if parsed != t.Result {
				t.Result = parsed
				needSave = true
			}
		}
	}
	if needSave {
		go saveTaskState()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviewTasks)
}

func handleDocGardenerTasks(w http.ResponseWriter, r *http.Request) {
	currentProject := getCurrentProjectName()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasksByType(TaskTypeDocGardener, currentProject))
}

func handleLogs(w http.ResponseWriter, r *http.Request) {
	logs := make([]LogInfo, 0)
	seen := make(map[string]bool)

	// 辅助函数：解析文件名中的 issue/pr 编号
	parseLogInfo := func(name string, info os.FileInfo) LogInfo {
		logInfo := LogInfo{
			Filename:   name,
			Size:       info.Size(),
			Timestamp:  info.ModTime(),
			ModifiedAt: info.ModTime(),
		}
		if strings.HasPrefix(name, "dev-issue-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				fmt.Sscanf(parts[2], "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "pm-issue-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				fmt.Sscanf(parts[2], "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "review-pr-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				fmt.Sscanf(parts[2], "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "review-report-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				numStr := strings.TrimSuffix(parts[2], ".md")
				fmt.Sscanf(numStr, "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "review-state-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				numStr := strings.TrimSuffix(parts[2], ".json")
				fmt.Sscanf(numStr, "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "doc-gardener-pr-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 4 {
				numStr := parts[3]
				fmt.Sscanf(numStr, "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "e2e-issue-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				fmt.Sscanf(parts[2], "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "design-issue-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				fmt.Sscanf(parts[2], "%d", &logInfo.IssueNumber)
			}
		} else if strings.HasPrefix(name, "rework-pr-") {
			parts := strings.Split(name, "-")
			if len(parts) >= 3 {
				fmt.Sscanf(parts[2], "%d", &logInfo.IssueNumber)
			}
		}
		return logInfo
	}

	// 1. 优先读取当前项目日志目录
	logsDir := getCurrentProjectLogsDir()
	if entries, err := os.ReadDir(logsDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			info, err := entry.Info()
			if err != nil {
				continue
			}
			seen[entry.Name()] = true
			logs = append(logs, parseLogInfo(entry.Name(), info))
		}
	}

	// 2. 兼容旧日志：读取 logs/ 根目录下的文件（排除子目录和 dashboard.log）
	if entries, err := os.ReadDir("logs"); err == nil {
		for _, entry := range entries {
			if entry.IsDir() || seen[entry.Name()] || entry.Name() == "dashboard.log" {
				continue
			}
			info, err := entry.Info()
			if err != nil {
				continue
			}
			// 只包含已知的日志文件前缀
			name := entry.Name()
			if strings.HasPrefix(name, "dev-issue-") || strings.HasPrefix(name, "pm-issue-") ||
				strings.HasPrefix(name, "review-pr-") ||
				strings.HasPrefix(name, "review-report-") || strings.HasPrefix(name, "review-state-") ||
				strings.HasPrefix(name, "doc-gardener-pr-") || strings.HasPrefix(name, "e2e-issue-") {
				logs = append(logs, parseLogInfo(name, info))
			}
		}
	}

	// 按修改时间降序排列，最新的在前
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].ModifiedAt.After(logs[j].ModifiedAt)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func handleLogContent(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "Missing file parameter", 400)
		return
	}

	if strings.Contains(filename, "..") {
		http.Error(w, "Invalid file", 400)
		return
	}

	path := getLogPath(filename)
	if path == "" {
		http.Error(w, "Log file not found", 404)
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(content)
}

func handleLogTail(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "Missing file parameter", 400)
		return
	}

	if strings.Contains(filename, "..") {
		http.Error(w, "Invalid file", 400)
		return
	}

	path := getLogPath(filename)
	if path == "" {
		http.Error(w, "Log file not found", 404)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", 500)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(w, "data: Error opening file: %s\n\n", err.Error())
		flusher.Flush()
		return
	}
	defer file.Close()

	// 默认先发送最后 2000 行历史日志，可通过 ?n= 调整
	n := 2000
	if nStr := r.URL.Query().Get("n"); nStr != "" {
		if parsed, err := strconv.Atoi(nStr); err == nil && parsed > 0 {
			n = parsed
		}
	}

	// 获取文件大小，先读取末尾 64KB 作为历史缓冲区
	stat, _ := file.Stat()
	if stat != nil && stat.Size() > 0 {
		offset := stat.Size() - 65536
		if offset < 0 {
			offset = 0
		}
		file.Seek(offset, 0)
		reader := bufio.NewReader(file)
		// 如果不是从文件头开始，丢弃第一行（可能不完整）
		if offset > 0 {
			reader.ReadString('\n')
		}
		var lines []string
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			lines = append(lines, strings.TrimRight(line, "\n"))
		}
		if len(lines) > n {
			lines = lines[len(lines)-n:]
		}
		for _, line := range lines {
			fmt.Fprintf(w, "data: %s\n\n", line)
		}
		if len(lines) > 0 {
			fmt.Fprintf(w, "data: \"--- 以上为最近 %d 行历史日志 ---\"\n\n", len(lines))
		}
	}

	// 跳到文件末尾，进入实时跟随模式
	file.Seek(0, 2)

	// 发送初始连接成功消息
	fmt.Fprint(w, "data: \"📡 日志流已连接...\"\n\n")
	flusher.Flush()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	reader := bufio.NewReader(file)
	done := r.Context().Done()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				line = strings.TrimRight(line, "\n")
				fmt.Fprintf(w, "data: %s\n\n", line)
			}
			flusher.Flush()
		}
	}
}

func handleLogStatus(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		writeJSONError(w, 400, "Missing file parameter")
		return
	}

	if strings.Contains(filename, "..") {
		writeJSONError(w, 400, "Invalid file")
		return
	}

	path := getLogPath(filename)
	if path == "" {
		writeJSONError(w, 404, "Log file not found")
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		writeJSONError(w, 500, err.Error())
		return
	}

	content, _ := os.ReadFile(path)
	contentStr := string(content)

	status := "unknown"
	if strings.Contains(contentStr, "开发完成") || strings.Contains(contentStr, "✅ Issue") || strings.Contains(contentStr, "任务已完成") || strings.Contains(contentStr, "✅ 任务完成") {
		status = "success"
	} else if strings.Contains(contentStr, "🛑 任务已停止") || strings.Contains(contentStr, "⏹️ 已停止") {
		status = "stopped"
	} else if (strings.Contains(contentStr, "\n❌ ") && strings.Contains(contentStr, "失败:")) || strings.Contains(contentStr, "Max number of steps reached") || strings.Contains(contentStr, "步骤限制") {
		status = "failed"
	} else if len(contentStr) > 100 {
		if time.Since(info.ModTime()) < 10*time.Minute {
			status = "running"
		} else {
			status = "interrupted"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"size":       info.Size(),
		"modifiedAt": info.ModTime(),
		"status":     status,
	})
}

// handleRunE2E 手动触发 E2E 测试运行
// 统一走 workflow 引擎，和 hook 自动触发同一路径：startE2ETask → runTaskWorkflow → workflows/{project}/e2e.yaml
func handleDesignFeedback(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		IssueNumber int    `json:"issueNumber"`
		Feedback    string `json:"feedback"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.IssueNumber == 0 {
		writeJSONError(w, http.StatusBadRequest, "issueNumber 必填")
		return
	}
	if strings.TrimSpace(req.Feedback) == "" {
		writeJSONError(w, http.StatusBadRequest, "反馈内容不能为空")
		return
	}

	projectName := getCurrentProjectName()
	logsDir := getProjectLogsDir(projectName)
	designDir := filepath.Join(logsDir, "designs", fmt.Sprintf("issue-%d", req.IssueNumber))
	feedbackPath := filepath.Join(designDir, "design-feedback.md")

	os.MkdirAll(designDir, 0755)
	content := fmt.Sprintf("# 设计反馈 (Issue #%d)\n\n%s\n\n---\n*反馈时间: %s*\n",
		req.IssueNumber, req.Feedback, time.Now().Format("2006-01-02 15:04:05"))
	if err := os.WriteFile(feedbackPath, []byte(content), 0644); err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("保存反馈失败: %v", err))
		return
	}

	log.Printf("[design-feedback] Issue #%d 设计反馈已保存: %s", req.IssueNumber, feedbackPath)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"issueNumber": req.IssueNumber,
		"message":     "反馈已保存",
	})
}

func handleBuildDesignPreview(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		IssueNumber int `json:"issueNumber"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.IssueNumber == 0 {
		writeJSONError(w, http.StatusBadRequest, "issueNumber 必填")
		return
	}

	projectName := getCurrentProjectName()
	logsDir := getProjectLogsDir(projectName)
	inputDir := filepath.Join(logsDir, "designs", fmt.Sprintf("issue-%d", req.IssueNumber))

	// 检测项目类型：Flutter 或 Phaser3
	isFlutter := isFlutterProject(projectName)
	isPhaser3 := isPhaser3Project(projectName)

	if isFlutter {
		buildFlutterDesignPreview(w, r, req.IssueNumber, projectName, logsDir, inputDir)
		return
	}

	if isPhaser3 {
		buildPhaser3DesignPreview(w, r, req.IssueNumber, projectName, logsDir, inputDir)
		return
	}

	writeJSONError(w, http.StatusBadRequest, "不支持的项目类型，无法生成预览")
}

// isFlutterProject 判断当前项目是否为 Flutter 项目
func isFlutterProject(projectName string) bool {
	projectPath := getProjectPathByName(projectName)
	pubspecPath := filepath.Join(projectPath, "pubspec.yaml")
	if _, err := os.Stat(pubspecPath); err == nil {
		return true
	}
	projectMutex.RLock()
	idx := currentProjectIndex
	if idx >= 0 && idx < len(config.Projects) {
		p := config.Projects[idx]
		if p.TechStack.Mobile == "flutter" || p.TechStack.Frontend == "flutter" {
			projectMutex.RUnlock()
			return true
		}
	}
	projectMutex.RUnlock()
	return false
}

// isPhaser3Project 判断当前项目是否为 Phaser3 项目
func isPhaser3Project(projectName string) bool {
	projectMutex.RLock()
	idx := currentProjectIndex
	if idx >= 0 && idx < len(config.Projects) {
		p := config.Projects[idx]
		if p.TechStack.Mobile == "phaser3" || p.TechStack.Frontend == "phaser3" {
			projectMutex.RUnlock()
			return true
		}
	}
	projectMutex.RUnlock()
	return false
}

// readPubspecPackageName 从项目 pubspec.yaml 中读取 Dart 包名
func readPubspecPackageName(projectPath string) string {
	data, err := os.ReadFile(filepath.Join(projectPath, "pubspec.yaml"))
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "name:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

// buildFlutterDesignPreview 为 Flutter 项目生成设计预览
func buildFlutterDesignPreview(w http.ResponseWriter, r *http.Request, issueNumber int, projectName, logsDir, inputDir string) {
	widgetsDir := filepath.Join(inputDir, "flutter-widgets")
	if _, err := os.Stat(widgetsDir); os.IsNotExist(err) {
		writeJSONError(w, http.StatusBadRequest, "未找到 Flutter 设计代码 (flutter-widgets/)")
		return
	}

	// 创建临时 Flutter 预览项目
	tmpDir, err := os.MkdirTemp("", fmt.Sprintf("flutter-preview-issue-%d-*", issueNumber))
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("创建临时目录失败: %v", err))
		return
	}
	defer os.RemoveAll(tmpDir)

	// 读取当前项目的 Dart 包名
	projectPath := getProjectPathByName(projectName)
	dartPkgName := readPubspecPackageName(projectPath)

	// 检查是否存在 design/issue-N 分支，如果有则导出到临时目录作为依赖
	// 这样预览编译能使用 agent 生成的完整代码（包括 app_strings.dart、di.config.dart 等）
	designBranch := fmt.Sprintf("design/issue-%d", issueNumber)
	useDesignBranch := false
	var designBranchPath string
	checkBranchCmd := exec.Command("git", "-C", projectPath, "rev-parse", "--verify", designBranch)
	if err := checkBranchCmd.Run(); err == nil {
		designBranchPath, err = os.MkdirTemp("", fmt.Sprintf("type-one-design-%d-*", issueNumber))
		if err == nil {
			archiveCmd := exec.Command("git", "-C", projectPath, "archive", designBranch)
			archiveCmd.Dir = projectPath
			tarCmd := exec.Command("tar", "-x", "-C", designBranchPath)
			archiveOut, _ := archiveCmd.StdoutPipe()
			tarCmd.Stdin = archiveOut
			_ = tarCmd.Start()
			_ = archiveCmd.Run()
			_ = tarCmd.Wait()

			// 尝试生成 injectable 的 di.config.dart（如果项目使用 injectable）
			if _, err := os.Stat(filepath.Join(designBranchPath, "lib", "core", "di", "di.dart")); err == nil {
				if _, err := os.Stat(filepath.Join(designBranchPath, "lib", "core", "di", "di.config.dart")); err != nil {
					log.Printf("[build-preview] 尝试生成 di.config.dart...")
					getCmd := exec.Command("flutter", "pub", "get")
					getCmd.Dir = designBranchPath
					_ = getCmd.Run()
					genCmd := exec.Command("flutter", "pub", "run", "build_runner", "build", "--delete-conflicting-outputs")
					genCmd.Dir = designBranchPath
					_ = genCmd.Run()
				}
			}

			useDesignBranch = true
			projectPath = designBranchPath
			log.Printf("[build-preview] 使用 %s 分支代码作为依赖: %s", designBranch, designBranchPath)
		}
	}
	if !useDesignBranch {
		log.Printf("[build-preview] %s 分支不存在，使用当前项目代码", designBranch)
	}


	// 复制通用 Flutter 预览模板（使用 templateDir+"/." 避免 filepath.Join 清理掉 .）
	templateDir := filepath.Join(projectRoot, ".tools", "flutter_preview_generic")
	cpCmd := exec.Command("cp", "-R", templateDir+"/.", tmpDir)
	if out, err := cpCmd.CombinedOutput(); err != nil {
		log.Printf("[build-preview] 复制模板失败: %v\n%s", err, string(out))
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("复制模板失败: %v", err))
		return
	}

	// 动态生成 pubspec.yaml，注入当前项目依赖（使 widget 中的 package:xxx 导入可解析）
	var pubspecContent string
	if dartPkgName != "" {
		pubspecContent = fmt.Sprintf(`name: flutter_preview_generic

description: "Flutter preview"

publish_to: 'none'

version: 1.0.0+1


environment:

  sdk: '>=3.4.0 <4.0.0'


dependencies:

  flutter:

    sdk: flutter

  cupertino_icons: ^1.0.8

  %s:

    path: %s


dev_dependencies:

  flutter_test:

    sdk: flutter

  flutter_lints: ^5.0.0


flutter:

  uses-material-design: true

`, dartPkgName, projectPath)
	} else {
		basePubspec, _ := os.ReadFile(filepath.Join(templateDir, "pubspec.yaml"))
		pubspecContent = string(basePubspec)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "pubspec.yaml"), []byte(pubspecContent), 0644); err != nil {
		log.Printf("[build-preview] 生成 pubspec.yaml 失败: %v", err)
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("生成 pubspec.yaml 失败: %v", err))
		return
	}

	// 复制 flutter-widgets/*.dart 到临时项目的 lib/preview/
	previewDir := filepath.Join(tmpDir, "lib", "preview")
	os.MkdirAll(previewDir, 0755)
	cpWidgetsCmd := exec.Command("cp", "-R", widgetsDir+"/.", previewDir)
	if out, err := cpWidgetsCmd.CombinedOutput(); err != nil {
		log.Printf("[build-preview] 复制 widgets 失败: %v\n%s", err, string(out))
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("复制 widgets 失败: %v", err))
		return
	}

	// 生成 main.dart
	mainDartPath := filepath.Join(tmpDir, "lib", "main.dart")
	genScript := filepath.Join(projectRoot, ".tools", "generate_preview_main.py")
	genCmd := exec.Command("python3", genScript, previewDir, mainDartPath)
	if out, err := genCmd.CombinedOutput(); err != nil {
		log.Printf("[build-preview] 生成 main.dart 失败: %v\n%s", err, string(out))
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("生成预览入口失败: %v", err))
		return
	}

	// 编译 Flutter Web
	outputWebDir := filepath.Join(inputDir, "web")
	os.RemoveAll(outputWebDir)
	baseHref := fmt.Sprintf("/logs/%s/designs/issue-%d/web/", projectName, issueNumber)
	buildCmd := exec.Command("flutter", "build", "web", "--release", "--base-href", baseHref)
	buildCmd.Dir = tmpDir
	buildOut, buildErr := buildCmd.CombinedOutput()
	if buildErr != nil {
		log.Printf("[build-preview] Flutter Web 编译失败: %v\n%s", buildErr, string(buildOut))
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Flutter Web 编译失败: %v\n%s", buildErr, string(buildOut)))
		return
	}

	// 复制编译产物到日志目录
	buildWebDir := filepath.Join(tmpDir, "build", "web")
	if _, err := os.Stat(buildWebDir); os.IsNotExist(err) {
		writeJSONError(w, http.StatusInternalServerError, "未找到 Flutter Web 编译产物")
		return
	}
	cpOutCmd := exec.Command("cp", "-R", buildWebDir+"/.", outputWebDir)
	if out, err := cpOutCmd.CombinedOutput(); err != nil {
		log.Printf("[build-preview] 复制产物失败: %v\n%s", err, string(out))
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("复制产物失败: %v", err))
		return
	}

	log.Printf("[build-preview] Issue #%d Flutter 预览生成成功", issueNumber)

	previewURL := fmt.Sprintf("/logs/%s/designs/issue-%d/web/", projectName, issueNumber)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"issueNumber": issueNumber,
		"url":         previewURL,
		"message":     "Flutter 预览生成成功",
	})
}

// buildPhaser3DesignPreview 为 Phaser3 项目生成设计预览
func buildPhaser3DesignPreview(w http.ResponseWriter, r *http.Request, issueNumber int, projectName, logsDir, inputDir string) {
	outputDir := filepath.Join(inputDir, "preview")

	// 检查设计代码是否存在
	codeDir := filepath.Join(inputDir, "src", "ui")
	if _, err := os.Stat(codeDir); os.IsNotExist(err) {
		writeJSONError(w, http.StatusBadRequest, "未找到设计代码 (src/ui/)")
		return
	}

	// 调用 generate_preview.py
	scriptPath := filepath.Join(projectRoot, ".tools", "phaser3_preview", "generate_preview.py")
	cmd := exec.Command("python3", scriptPath,
		"--input", inputDir,
		"--output", outputDir,
		"--issue", strconv.Itoa(issueNumber))
	cmd.Dir = projectRoot

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("[build-preview] 生成失败: %v\n%s", err, string(output))
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("预览生成失败: %v", err))
		return
	}

	log.Printf("[build-preview] Issue #%d 预览生成成功\n%s", issueNumber, string(output))

	previewURL := fmt.Sprintf("/logs/%s/designs/issue-%d/preview/", projectName, issueNumber)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"issueNumber": issueNumber,
		"url":         previewURL,
		"message":     "预览生成成功",
	})
}

func handleRunE2E(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		IssueNumber int    `json:"issueNumber"`
		Branch      string `json:"branch"`
		TaskKind    string `json:"taskKind"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.IssueNumber == 0 {
		writeJSONError(w, http.StatusBadRequest, "issueNumber 必填")
		return
	}

	// 1. 推断分支
	branch := req.Branch
	if branch == "" {
		// 优先从已有的 dev task 获取分支
		if devTask, ok := getTask(TaskTypeDev, req.IssueNumber); ok && devTask.Branch != "" {
			branch = devTask.Branch
		} else {
			branch = fmt.Sprintf("feat/issue-%d", req.IssueNumber)
		}
	}

	// 2. 推断任务类型
	taskKind := req.TaskKind
	if taskKind == "" {
		taskKind = getIssueClassification(req.IssueNumber)
	}

	// 3. 检查是否已有 running 的 E2E 任务
	if existing, ok := getTask(TaskTypeE2E, req.IssueNumber); ok && existing.Status == "running" {
		writeJSONError(w, http.StatusConflict, fmt.Sprintf("E2E 任务 e2e-%d 已在运行中", req.IssueNumber))
		return
	}

	// 5. 启动 E2E 任务（统一入口，和 hook 自动触发同一路径）
	worktreePath := getProjectPath()
	startE2ETask(req.IssueNumber, branch, taskKind, "manual", worktreePath, getCurrentProjectName())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "started",
		"message": fmt.Sprintf("E2E 任务已启动: e2e-%d", req.IssueNumber),
	})
}
