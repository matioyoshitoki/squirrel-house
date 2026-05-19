package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// PullRequest 表示 GitHub PR
type PullRequest struct {
	Number    int    `json:"number"`
	Title     string `json:"title"`
	State     string `json:"state"`
	Branch    string `json:"branch"`
	BaseRef   string `json:"baseRefName"`
	HeadRef   string `json:"headRefName"`
	URL       string `json:"url"`
	CreatedAt string `json:"createdAt"`
	Author    struct {
		Login string `json:"login"`
	} `json:"author"`
}

// buildReworkTestCommands 根据项目技术栈生成返工时的自测命令说明
func buildReworkTestCommands(ts TechStack) string {
	var parts []string
	if ts.Mobile == "flutter" {
		parts = append(parts,
			"   - 若修改了 Flutter 代码（apps/mobile/ 或 packages/design-system/）：",
			"     cd apps/mobile && flutter pub get",
			"     cd apps/mobile && flutter analyze",
			"     cd apps/mobile && flutter test",
		)
	}
	if ts.Frontend == "phaser3" || ts.Frontend == "react" || ts.Frontend == "vue" || ts.Frontend == "nextjs" {
		parts = append(parts,
			"   - 若修改了前端代码（apps/phaser3/ 或对应前端目录）：",
			"     cd 对应前端目录 && pnpm install（如 node_modules 缺失）",
			"     cd 对应前端目录 && npm run lint",
			"     cd 对应前端目录 && npm run test",
			"     cd 对应前端目录 && npm run build",
		)
	}
	if ts.Backend == "nestjs" || ts.Backend == "nodejs" || ts.Backend == "express" {
		parts = append(parts,
			"   - 若修改了 Node/NestJS 后端代码：",
			"     cd apps/api && npm test",
			"     cd apps/api && npm run build",
		)
	}
	if ts.Backend == "go" || ts.Backend == "golang" {
		parts = append(parts,
			"   - 若修改了 Go 代码：",
			"     go test ./...",
			"     go vet ./...",
		)
	}
	parts = append(parts,
		"   - 通用检查：",
		"     make check-docs（文档检查）",
	)
	if len(parts) == 2 { // 只有通用检查，说明技术栈未知
		parts = append([]string{
			"   - 若项目使用 Go：go test ./...、go vet ./...",
			"   - 若项目使用 Node：npm test、npm run build",
			"   - 若项目使用 Flutter：flutter analyze、flutter test",
		}, parts...)
	}
	return strings.Join(parts, "\n")
}

func handlePullRequests(w http.ResponseWriter, r *http.Request) {
	platform := getPlatform(getProjectPath())
	mrs, err := platform.ListOpenMRs()
	if err != nil {
		msg := fmt.Sprintf("获取 PR/MR 列表失败: %v", err)
		log.Printf("⚠️ %s", msg)
		writeJSONError(w, 500, msg)
		return
	}

	// 转换为前端兼容的 PullRequest 格式
	prs := make([]PullRequest, len(mrs))
	for i, mr := range mrs {
		prs[i] = PullRequest{
			Number:  mr.IID,
			Title:   mr.Title,
			State:   mr.State,
			Branch:  mr.SourceBranch,
			BaseRef: mr.TargetBranch,
			HeadRef: mr.SourceBranch,
			URL:     mr.URL,
			CreatedAt: mr.CreatedAt,
		}
		prs[i].Author.Login = mr.Author
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prs)
}

// startReviewTask 启动 PR Review 任务（HTTP handler 和 hook 共用）
func startReviewTask(prNumber int, branch string, projectName string) string {
	tid := taskID(TaskTypeReview, prNumber)
	tkey := taskKey(projectName, TaskTypeReview, prNumber)
	tasksMutex.RLock()
	if t, ok := tasks[tkey]; ok && t.Status == "running" {
		tasksMutex.RUnlock()
		log.Printf("[review] PR #%d 正在 Review 中，跳过", prNumber)
		return ""
	}
	tasksMutex.RUnlock()

	logsDir := getProjectLogsDir(projectName)
	os.MkdirAll(logsDir, 0755)
	timestamp := time.Now().Format("20060102-150405")
	logFileName := filepath.Join(logsDir, fmt.Sprintf("review-pr-%d-%s.log", prNumber, timestamp))

	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Printf("[review] 创建日志文件失败: %v", err)
		return ""
	}
	logFile.Close()

	logMsg := fmt.Sprintf("[%s] PR #%d Review 启动\n分支: %s\n项目路径: %s\n日志文件: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), prNumber, branch, getProjectPathByName(projectName), logFileName)
	os.WriteFile(logFileName, []byte(logMsg), 0644)

	task := &Task{
		ID:          tid,
		Type:        TaskTypeReview,
		Status:      "running",
		ProjectName: projectName,
		TargetID:    prNumber,
		TargetTitle: fmt.Sprintf("PR #%d", prNumber),
		Branch:      branch,
		StartTime:   time.Now(),
		LogFile:     logFileName,
		Metadata: map[string]interface{}{
			"tmuxSession": fmt.Sprintf("nwops-%s", tid),
		},
	}
	if !setTask(task) {
		log.Printf("[review] PR #%d 正在 Review 中，跳过", prNumber)
		return ""
	}

	go func() {
		projectPath := getProjectPathByName(projectName)
		tmpDir, _ := os.MkdirTemp("", fmt.Sprintf("review-pr-%d-*", prNumber))
		if err := cleanWorktreeForBranch(projectPath, branch); err != nil {
			appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ clean worktree failed: %v\n", err)))
			updateTaskStatus(TaskTypeReview, prNumber, "failed")
			return
		}

		fetchCmd := exec.Command("git", "fetch", "origin")
		fetchCmd.Dir = projectPath
		fetchCmd.Run()

		baseRef := fmt.Sprintf("origin/%s", branch)
		worktreeCmd := fmt.Sprintf("git worktree add -B %s %s %s", branch, tmpDir, baseRef)
		wtCmd := exec.Command("bash", "-c", worktreeCmd)
		wtCmd.Dir = projectPath
		wtCmd.Run()

		platform := getPlatform(projectPath)
		if err := platform.AuthStatus(); err != nil {
			appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ %s CLI 未认证: %v\n", platform.CLIName(), err)))
			updateTaskStatus(TaskTypeReview, prNumber, "failed")
			return
		}

		prBody := ""
		if mr, err := platform.ViewMR(prNumber); err == nil {
			prBody = mr.Body
		}
		issueNumber := parseRelatedIssue(prBody)

		// 预生成平台相关数据文件（使 workflow 平台无关）
		changedFiles, _ := platform.GetMRDiffFiles(prNumber)
		docCoverage := "=== PR 变更文件 ===\n" + strings.Join(changedFiles, "\n") + "\n\n=== docs/ 目录变更 ===\n"
		designAssets := []string{}
		for _, f := range changedFiles {
			if strings.HasPrefix(f, "docs/") {
				docCoverage += f + "\n"
			}
			lower := strings.ToLower(f)
			if strings.HasSuffix(lower, ".png") || strings.HasSuffix(lower, ".jpg") ||
				strings.HasSuffix(lower, ".jpeg") || strings.HasSuffix(lower, ".svg") ||
				strings.HasSuffix(lower, ".webp") || strings.HasSuffix(lower, ".gif") {
				designAssets = append(designAssets, f)
			}
		}
		// 补充 designs/ 目录下的文件
		if _ , err := os.ReadDir(filepath.Join(tmpDir, "designs")); err == nil {
			var walk func(string)
			walk = func(dir string) {
				entries, _ := os.ReadDir(dir)
				for _, e := range entries {
					if e.IsDir() {
						walk(filepath.Join(dir, e.Name()))
					} else {
						designAssets = append(designAssets, filepath.Join(dir, e.Name()))
					}
				}
			}
			walk(filepath.Join(tmpDir, "designs"))
		}
		os.WriteFile(filepath.Join(tmpDir, ".doc-coverage.txt"), []byte(docCoverage), 0644)
		os.WriteFile(filepath.Join(tmpDir, ".design-assets.txt"), []byte(strings.Join(designAssets, "\n")), 0644)

		// 提取 Figma / 设计链接
		figmaLinks := extractFigmaLinks(prBody)
		os.WriteFile(filepath.Join(tmpDir, ".figma-links.txt"), []byte(strings.Join(figmaLinks, "\n")), 0644)

		agentCtx := buildAgentContext(tmpDir, issueNumber, prNumber)
		agentCtx["AGENT_DOC_COVERAGE"] = filepath.Join(tmpDir, ".doc-coverage.txt")
		agentCtx["AGENT_DESIGN_ASSETS_LIST"] = filepath.Join(tmpDir, ".design-assets.txt")
		agentCtx["AGENT_FIGMA_LINKS"] = filepath.Join(tmpDir, ".figma-links.txt")

		ctxPrompt := formatAgentContextPrompt(agentCtx, issueNumber, prNumber)
		auxFiles, _ := renderCtxTemplate("review_aux_files", map[string]interface{}{
			"DocCoverage":      agentCtx["AGENT_DOC_COVERAGE"],
			"DesignAssetsList": agentCtx["AGENT_DESIGN_ASSETS_LIST"],
			"FigmaLinks":       agentCtx["AGENT_FIGMA_LINKS"],
		})
		ctxPrompt += auxFiles

		agentFile := prepareAgentConfig("reviewer")
		prompt, _ := renderCtxTemplate("review_main", map[string]interface{}{
			"ContextPrompt": ctxPrompt,
			"PRNumber":      prNumber,
			"Branch":        branch,
		})

		vars := map[string]interface{}{
			"prNumber":     prNumber,
			"branch":       branch,
			"tmpDir":       tmpDir,
			"logsDir":      logsDir,
			"logFileName":  logFileName,
			"agentFile":    agentFile,
			"prompt":       prompt,
			"worktreeCmd":  worktreeCmd,
			"platformType": platform.Name(),
			"platformCLI":  platform.CLIName(),
		}
		for k, v := range agentCtx {
			vars[k] = v
		}

		runTaskWorkflow(task, vars, nil)
	}()

	return logFileName
}

// handleReviewPR 启动 PR Review（HTTP 入口）
func handleReviewPR(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		PRNumber int    `json:"prNumber"`
		Branch   string `json:"branch"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	logFileName := startReviewTask(req.PRNumber, req.Branch, getCurrentProjectName())
	if logFileName == "" {
		writeJSONError(w, 409, fmt.Sprintf("PR #%d 正在 Review 中", req.PRNumber))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "started",
		"message": fmt.Sprintf("PR #%d Review 已启动", req.PRNumber),
		"logFile": logFileName,
	})
}

// handleReviewReport 获取审查报告
func handleReviewReport(w http.ResponseWriter, r *http.Request) {
	prNumber := r.URL.Query().Get("pr")
	if prNumber == "" {
		writeJSONError(w, 400, "Missing pr parameter")
		return
	}

	prNum := toInt(prNumber)

	tasksMutex.RLock()
	task, taskExists := tasks[taskKey(getCurrentProjectName(), TaskTypeReview, prNum)]
	tasksMutex.RUnlock()

	reportPaths := []string{
		filepath.Join(getCurrentProjectLogsDir(), fmt.Sprintf("review-report-%s.md", prNumber)),
		fmt.Sprintf("logs/review-report-%s.md", prNumber),
		filepath.Join(getProjectPath(), fmt.Sprintf("logs/review-report-%s.md", prNumber)),
	}

	// 优先使用内存中的 task.Report（始终是最新结果），回退到文件读取
	var report string
	if taskExists && task.Report != "" {
		report = task.Report
	} else {
		for _, rp := range reportPaths {
			if data, err := os.ReadFile(rp); err == nil {
				report = string(data)
				break
			}
		}
	}

	if report == "" && !taskExists {
		writeJSONError(w, 404, "Report not found")
		return
	}

	// 读取日志内容
	logContent := ""
	if taskExists && task.LogFile != "" && !strings.Contains(task.LogFile, "*") {
		if lc, err := os.ReadFile(task.LogFile); err == nil {
			logContent = string(lc)
		}
	}

	resp := map[string]interface{}{
		"status": "success",
		"report": report,
	}
	if taskExists {
		// 重新解析 report 中的 result，修复旧版本代码错误解析导致的历史数据
		result := task.Result
		if report != "" {
			result = extractReviewResult(report)
		}
		resp["result"] = result
		resp["startTime"] = task.StartTime
		resp["endTime"] = task.EndTime
	}
	if logContent != "" {
		resp["log"] = logContent
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// relatedIssueRe 用于从 PR body 中解析关联的 Issue 编号
var relatedIssueRe = regexp.MustCompile(`(?i)(?:Fixes|Closes|Resolves|refs|References?)\s*#(\d+)`)

// figmaLinkRe 用于从文本中提取 Figma / 设计相关链接
var figmaLinkRe = regexp.MustCompile(`(?i)https?://[^\s"<>\[\]^` + "`" + `{|}]*(?:figma\.com|design)[^\s"<>\[\]^` + "`" + `{|}]*`)

// extractFigmaLinks 从文本中提取设计相关链接
func extractFigmaLinks(text string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, m := range figmaLinkRe.FindAllString(text, -1) {
		if !seen[m] {
			seen[m] = true
			result = append(result, m)
		}
	}
	return result
}

// parseRelatedIssue 从 PR body 中解析关联的 Issue 编号
func parseRelatedIssue(body string) int {
	matches := relatedIssueRe.FindStringSubmatch(body)
	if len(matches) >= 2 {
		var num int
		fmt.Sscanf(matches[1], "%d", &num)
		if num > 0 {
			return num
		}
	}
	return 0
}

// extractBlockingIssues 从审查报告中提取"必须修复"的问题清单
func extractBlockingIssues(report string) string {
	var issues []string
	lines := strings.Split(report, "\n")
	inBlocking := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "必须修复") || strings.Contains(trimmed, "Blocking") {
			inBlocking = true
			continue
		}
		if inBlocking {
			if strings.Contains(trimmed, "建议优化") || strings.Contains(trimmed, "Non-blocking") {
				break
			}
			if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") || strings.HasPrefix(trimmed, "1. ") || strings.HasPrefix(trimmed, "2. ") {
				issues = append(issues, trimmed)
			} else if strings.HasPrefix(trimmed, "### 🔴") {
				issues = append(issues, trimmed)
			}
		}
	}
	if len(issues) == 0 {
		return "请根据审查报告中的问题进行修改。"
	}
	return strings.Join(issues, "\n")
}

// handleReworkPR 标记 PR 需要返工，并启动自动修复
// startReworkTask 启动 Rework 任务（HTTP handler 和 hook 共用）
func startReworkTask(prNumber int, branch string, reportStr string, projectName string) (string, int) {
	projectPath := getProjectPathByName(projectName)
	var prTitle, prBody string
	platform := getPlatform(projectPath)
	if mr, err := platform.ViewMR(prNumber); err == nil {
		prTitle = mr.Title
		prBody = mr.Body
		if branch == "" {
			branch = mr.SourceBranch
		}
	}

	issueNumber := parseRelatedIssue(prBody)
	issueTitle := ""
	if issueNumber > 0 {
		if issue, err := platform.ViewIssue(issueNumber); err == nil {
			issueTitle = issue.Title
		}
	}

	blockingIssues := extractBlockingIssues(reportStr)

	logsDir := getProjectLogsDir(projectName)
	timestamp := time.Now().Format("20060102-150405")
	logFileName := filepath.Join(logsDir, fmt.Sprintf("rework-pr-%d-%s.log", prNumber, timestamp))
	logContent := fmt.Sprintf("[%s] PR #%d Rework 启动\n分支: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), prNumber, branch)
	if issueNumber > 0 {
		logContent += fmt.Sprintf("关联 Issue: #%d\n", issueNumber)
	}
	os.WriteFile(logFileName, []byte(logContent), 0644)

	targetTitle := prTitle
	if issueTitle != "" {
		targetTitle = fmt.Sprintf("Issue #%d: %s", issueNumber, issueTitle)
	}

	taskKind := "full-stack"
	if issueNumber > 0 {
		taskKind = getIssueClassification(issueNumber)
	}

	task := &Task{
		ID:          taskID(TaskTypeRework, prNumber),
		Type:        TaskTypeRework,
		Status:      "running",
		ProjectName: projectName,
		TargetID:    prNumber,
		TargetTitle: targetTitle,
		PRNumber:    prNumber,
		Branch:      branch,
		Kind:        taskKind,
		StartTime:   time.Now(),
		LogFile:     logFileName,
		Metadata: map[string]interface{}{
			"issueNumber": issueNumber,
			"taskKind":    taskKind,
			"tmuxSession": fmt.Sprintf("nwops-%s", taskID(TaskTypeRework, prNumber)),
		},
	}
	commentBody, _ := renderCtxTemplate("rework_comment", map[string]interface{}{
		"BlockingIssues": blockingIssues,
	})

	if !setTask(task) {
		appendToFile(logFileName, []byte("\n❌ Rework 任务已在运行中\n"))
		return "", issueNumber
	}

	go func() {
		// 前置平台操作：添加返工评论和标签（使 workflow 平台无关）
		if commentBody != "" {
			if err := platform.CreateMRComment(prNumber, commentBody); err != nil {
				log.Printf("[rework] 添加返工评论失败: %v", err)
			}
		}
		if err := platform.AddMRLabel(prNumber, "needs-work"); err != nil {
			log.Printf("[rework] 添加 needs-work 标签失败: %v", err)
		}

		tmpDir, _ := os.MkdirTemp("", fmt.Sprintf("rework-pr-%d-*", prNumber))
		if err := cleanWorktreeForBranch(projectPath, branch); err != nil {
			appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ clean worktree failed: %v\n", err)))
			updateTaskStatus(TaskTypeRework, prNumber, "failed")
			return
		}

		fetchCmd := exec.Command("git", "fetch", "origin")
		fetchCmd.Dir = projectPath
		fetchCmd.Run()

		worktreeCmd := fmt.Sprintf("git worktree add -B %s %s origin/%s", branch, tmpDir, branch)
		wtCmd := exec.Command("bash", "-c", worktreeCmd)
		wtCmd.Dir = projectPath
		wtCmd.Run()

		agentFile := prepareAgentConfig("rework")
		if agentFile == "" {
			agentFile = prepareAgentConfig("developer")
		}

		promptTitle := issueTitle
		if promptTitle == "" {
			promptTitle = prTitle
		}
		promptIssueNumber := issueNumber
		if promptIssueNumber == 0 {
			promptIssueNumber = prNumber
		}
		agentCtx := buildAgentContext(tmpDir, issueNumber, prNumber)
		ctxPrompt := formatAgentContextPrompt(agentCtx, issueNumber, prNumber)

		projectMutex.RLock()
		prj := config.Projects[currentProjectIndex]
		projectMutex.RUnlock()
		testCmds := buildReworkTestCommands(prj.TechStack)

		prompt, _ := renderCtxTemplate("rework_prompt", map[string]interface{}{
			"ContextPrompt": ctxPrompt,
			"IssueNumber":   promptIssueNumber,
			"IssueTitle":    promptTitle,
			"PRNumber":      prNumber,
			"Branch":        branch,
			"TestCommands":  testCmds,
			"ReviewReport":  reportStr,
		})

		reviewReportPath := filepath.Join(logsDir, fmt.Sprintf("review-report-%d.md", prNumber))
		if _, err := os.Stat(reviewReportPath); os.IsNotExist(err) {
			// 非时间戳文件被 clearReviewFailedState 删除，回退到最新的带时间戳文件
			matches, _ := filepath.Glob(filepath.Join(logsDir, fmt.Sprintf("review-report-%d-*.md", prNumber)))
			if len(matches) > 0 {
				// 按修改时间排序，取最新
				latest := matches[0]
				latestMod := time.Time{}
				for _, m := range matches {
					if fi, err := os.Stat(m); err == nil && fi.ModTime().After(latestMod) {
						latest = m
						latestMod = fi.ModTime()
					}
				}
				reviewReportPath = latest
			}
		}
		vars := map[string]interface{}{
			"prNumber":         prNumber,
			"branch":           branch,
			"issueNumber":      issueNumber,
			"issueTitle":       issueTitle,
			"tmpDir":           tmpDir,
			"worktreeCmd":      worktreeCmd,
			"logFileName":      logFileName,
			"agentFile":        agentFile,
			"prompt":           prompt,
			"reviewReport":     reportStr,
			"reviewReportPath": reviewReportPath,
			"commentBody":      commentBody,
			"taskKind":         taskKind,
		}
		for k, v := range agentCtx {
			vars[k] = v
		}
		runTaskWorkflow(task, vars, nil)
	}()

	return logFileName, issueNumber
}

func handleReworkPR(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		PRNumber int `json:"prNumber"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	// 更新 review 任务状态为 failed
	projectName := getCurrentProjectName()
	tkey := taskKey(projectName, TaskTypeReview, req.PRNumber)
	tasksMutex.Lock()
	if t, ok := tasks[tkey]; ok && t.Type == TaskTypeReview {
		t.Status = "failed"
		t.Result = "needs_fix"
		t.EndTime = time.Now()
	}
	tasksMutex.Unlock()

	// 持久化返工状态
	stateFile := filepath.Join(getCurrentProjectLogsDir(), fmt.Sprintf("review-state-%d.json", req.PRNumber))
	stateData, _ := json.Marshal(map[string]interface{}{
		"prNumber": req.PRNumber,
		"status":   "failed",
		"result":   "needs_fix",
		"endTime":  time.Now(),
	})
	os.WriteFile(stateFile, stateData, 0644)

	var reportStr string
	tasksMutex.RLock()
	if t, ok := tasks[tkey]; ok && t.Type == TaskTypeReview {
		reportStr = t.Report
	}
	tasksMutex.RUnlock()

	logFileName, issueNumber := startReworkTask(req.PRNumber, "", reportStr, projectName)
	if logFileName == "" {
		writeJSONError(w, 409, "Rework 任务已在运行中")
		return
	}

	notifyReworkMarked(req.PRNumber, issueNumber, extractBlockingIssues(reportStr))

	msg := fmt.Sprintf("PR #%d 已标记为需要返工，自动修复已启动", req.PRNumber)
	if issueNumber > 0 {
		msg = fmt.Sprintf("PR #%d 已标记为需要返工，关联 Issue #%d 并启动自动修复", req.PRNumber, issueNumber)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":      "marked",
		"message":     msg,
	})
}

// hasMergeConflict 判断 merge 错误输出是否由冲突导致
func hasMergeConflict(output string) bool {
	lower := strings.ToLower(output)
	keywords := []string{"conflict", "has conflicts", "merge conflict", "dirty", "branch is out of date"}
	for _, k := range keywords {
		if strings.Contains(lower, k) {
			return true
		}
	}
	return false
}

// handleMergePR 合并 PR
func handleMergePR(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		PRNumber int `json:"prNumber"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	// 执行 squash 合并
	platform := getPlatform(getProjectPath())
	err := platform.MergeMR(req.PRNumber)

	// 如果合并失败且可能是冲突导致，尝试自动更新分支
	if err != nil && hasMergeConflict(err.Error()) {
		log.Printf("⚠️ PR/MR #%d 合并失败（检测到冲突），尝试自动更新分支...", req.PRNumber)
		if rebaseErr := platform.RebaseMR(req.PRNumber); rebaseErr != nil {
			writeJSONError(w, 409, fmt.Sprintf("PR/MR #%d 存在合并冲突，自动更新分支失败，请手动解决冲突后再合并。\n原始错误: %v\n更新分支错误: %v", req.PRNumber, err, rebaseErr))
			return
		}
		// 更新分支成功后再次尝试合并
		err = platform.MergeMR(req.PRNumber)
	}

	if err != nil {
		writeJSONError(w, 500, fmt.Sprintf("合并失败: %v", err))
		return
	}

	notifyPRMerged(req.PRNumber)
	go broadcastEvent("pr:merged", map[string]interface{}{
		"prNumber": req.PRNumber,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "merged",
		"message": fmt.Sprintf("PR #%d 已合并", req.PRNumber),
	})
}

// startDocGardenerTask 启动 Doc Gardener 工作流
func startDocGardenerTask(prNumber int, merged bool, baseBranch, headBranch, changedFiles, prompt string, worktreePath string, projectName string) string {
	agentFile := prepareAgentConfig("pr-doc-gardener")
	if agentFile == "" {
		return ""
	}

	projectPath := getProjectPathByName(projectName)
	logsDir := getProjectLogsDir(projectName)
	os.MkdirAll(logsDir, 0755)
	timestamp := time.Now().Format("20060102-150405")
	logFileName := filepath.Join(logsDir, fmt.Sprintf("doc-gardener-pr-%d-%s.log", prNumber, timestamp))
	logFile, err := os.Create(logFileName)
	if err != nil {
		return ""
	}
	logFile.Close()

	logMsg := fmt.Sprintf("[%s] PR #%d Doc Gardener 启动\nMerged: %v\nBase: %s\nHead: %s\n日志: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), prNumber, merged, baseBranch, headBranch, logFileName)
	os.WriteFile(logFileName, []byte(logMsg), 0644)

	task := &Task{
		ID:          taskID(TaskTypeDocGardener, prNumber),
		Type:        TaskTypeDocGardener,
		Status:      "running",
		ProjectName: projectName,
		TargetID:    prNumber,
		TargetTitle: fmt.Sprintf("PR #%d", prNumber),
		PRNumber:    prNumber,
		Merged:      merged,
		Branch:      headBranch,
		StartTime:   time.Now(),
		LogFile:     logFileName,
		Metadata: map[string]interface{}{
			"tmuxSession": fmt.Sprintf("nwops-%s", taskID(TaskTypeDocGardener, prNumber)),
		},
	}
	if !setTask(task) {
		log.Printf("❌ Doc Gardener 任务已在运行中: PR #%d", prNumber)
		return ""
	}

	// 后台执行 git 准备和 workflow 启动
	go func() {
		var tmpDir string
		var worktreeCmd string
		if worktreePath != "" {
			// 复用 dev/rework 的 worktree，演进效果直接生效，且多次演进可叠加
			tmpDir = worktreePath
			log.Printf("[doc-gardener] 复用 dev worktree: %s", tmpDir)
			worktreeCmd = fmt.Sprintf("git worktree add -B %s %s origin/%s", headBranch, tmpDir, headBranch)
		} else {
			tmpDir, _ = os.MkdirTemp("", fmt.Sprintf("doc-gardener-%d-*", prNumber))
			if err := cleanWorktreeForBranch(projectPath, headBranch); err != nil {
				appendToFile(logFileName, []byte(fmt.Sprintf("\n❌ clean worktree failed: %v\n", err)))
				updateTaskStatus(TaskTypeDocGardener, prNumber, "failed")
				return
			}

			fetchCmd := exec.Command("git", "fetch", "origin")
			fetchCmd.Dir = projectPath
			fetchCmd.Run()

			worktreeCmd = fmt.Sprintf("git worktree add -B %s %s origin/%s", headBranch, tmpDir, headBranch)
			// 同步创建 worktree，确保 buildAgentContext 能扫描到文件
			wtCmd := exec.Command("bash", "-c", worktreeCmd)
			wtCmd.Dir = projectPath
			wtCmd.Run()
		}

		agentCtx := buildAgentContext(tmpDir, 0, prNumber)
		ctxPrompt := formatAgentContextPrompt(agentCtx, 0, prNumber)
		fullPrompt := ctxPrompt + prompt

		vars := map[string]interface{}{
			"prNumber":    prNumber,
			"tmpDir":      tmpDir,
			"worktreeCmd": worktreeCmd,
			"branch":      headBranch,
			"logFileName": logFileName,
			"agentFile":   agentFile,
			"prompt":      fullPrompt,
		}
		for k, v := range agentCtx {
			vars[k] = v
		}
		runTaskWorkflow(task, vars, nil)
	}()

	return logFileName
}

// handleMergedDocGardener 对已合并 PR 执行文档检查并自动创建 PR
func handleMergedDocGardener(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		PRNumber int `json:"prNumber"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	projectPath := getProjectPath()
	platform := getPlatform(projectPath)
	mr, err := platform.ViewMR(req.PRNumber)
	if err != nil {
		writeJSONError(w, 500, fmt.Sprintf("获取 MR 信息失败: %v", err))
		return
	}

	changedFilesList, err := platform.GetMRDiffFiles(req.PRNumber)
	if err != nil {
		writeJSONError(w, 500, fmt.Sprintf("获取 MR diff 失败: %v", err))
		return
	}
	changedFiles := strings.Join(changedFilesList, "\n")

	prompt := fmt.Sprintf("请对以下已合并的 PR 进行回顾性文档协同检查，并直接产出可合并的文档 PR。\n\n"+
		"【上下文】\n"+
		"- 这是一个已合并的 PR\n"+
		"- PR 编号: #%d\n"+
		"- PR 标题: %s\n"+
		"- Base 分支: %s\n"+
		"- 原 Head 分支: %s\n\n"+
		"【变更文件列表】\n%s\n\n"+
		"【必须执行的操作】\n"+
		"1. 分析上述变更，检查对应文档是否已补充\n"+
		"2. 对遗漏项，直接创建或修改对应的 .md 文档文件\n"+
		"3. 检查并更新 CHANGELOG.md：\n"+
		"   - 如果仓库根目录没有 CHANGELOG.md，创建一个\n"+
		"   - 在最新版本区块下追加条目，格式：`- #%d %s`\n"+
		"4. 如果 PR 包含接口破坏性变更（删除/重命名字段、修改 URL、修改返回值结构），必须创建或更新迁移指南 docs/migration/README.md\n"+
		"5. 输出结构化报告到日志\n"+
		"6. 创建新分支：git checkout -b docs/pr-%d-doc-gardener\n"+
		"7. 将文档变更提交：git add -A && git commit -m 'docs: 补充 PR #%d 文档 (by PR Doc Gardener)'\n"+
		"8. push 新分支：git push origin docs/pr-%d-doc-gardener\n"+
		"9. 创建 PR/MR：%s mr create -b %s -t 'docs: 补充 MR #%d 文档' -d '由 PR Doc Gardener 自动生成的文档补充。'\n\n"+
		"注意：只修改文档文件（.md），不要修改业务代码。", req.PRNumber, mr.Title, mr.TargetBranch, mr.SourceBranch, changedFiles, req.PRNumber, mr.Title, req.PRNumber, req.PRNumber, req.PRNumber, platform.CLIName(), mr.TargetBranch, req.PRNumber)

	logFileName := startDocGardenerTask(req.PRNumber, true, mr.TargetBranch, mr.TargetBranch, changedFiles, prompt, "", getCurrentProjectName())
	if logFileName == "" {
		writeJSONError(w, 500, "启动 Doc Gardener 失败")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "started",
		"message": fmt.Sprintf("PR #%d（已合并）文档检查已启动", req.PRNumber),
		"logFile": logFileName,
	})
}

// handleDocGardener 启动 PR Doc Gardener 文档检查（针对 open PR）
func handleDocGardener(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		PRNumber   int    `json:"prNumber"`
		BaseBranch string `json:"baseBranch"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	projectPath := getProjectPath()
	baseBranch := req.BaseBranch
	if baseBranch == "" {
		baseBranch = resolveBaseRef(projectPath, getProjectDefaultBranch())
	}
	platform := getPlatform(projectPath)
	mr, _ := platform.ViewMR(req.PRNumber)
	headRefName := ""
	if mr != nil {
		headRefName = mr.SourceBranch
	}

	changedFilesList, _ := platform.GetMRDiffFiles(req.PRNumber)
	changedFiles := strings.Join(changedFilesList, "\n")

	prompt := fmt.Sprintf("请对当前 open 的 PR 进行文档协同检查，并直接补充遗漏的文档。\n\n"+
		"【上下文】\n"+
		"- 这是一个尚未合并的 open PR\n"+
		"- PR 编号: #%d\n"+
		"- 当前分支: %s\n"+
		"- Base 分支: %s\n\n"+
		"【变更文件列表】\n%s\n\n"+
		"【必须执行的操作】\n"+
		"1. 分析上述变更，检查对应文档是否已补充\n"+
		"2. 对遗漏项，直接创建或修改对应的 .md 文档文件\n"+
		"3. 检查 CHANGELOG.md：\n"+
		"   - 如果 PR 包含用户可见的功能变更或接口变更，而 CHANGELOG.md 中尚未记录，必须在最新版本区块下追加条目 `- #%d`\n"+
		"4. 如果 PR 包含接口破坏性变更，必须创建或更新迁移指南 docs/migration/README.md\n"+
		"5. 输出结构化报告到日志\n"+
		"6. 将文档变更提交：git add -A && git commit -m 'docs: 补充 MR #%d 文档 (by PR Doc Gardener)'\n"+
		"7. push 到当前分支：git push origin %s\n\n"+
		"注意：只修改文档文件（.md），不要修改业务代码。", req.PRNumber, headRefName, baseBranch, changedFiles, req.PRNumber, req.PRNumber, headRefName)

	logFileName := startDocGardenerTask(req.PRNumber, false, baseBranch, headRefName, changedFiles, prompt, "", getCurrentProjectName())
	if logFileName == "" {
		writeJSONError(w, 500, "启动 Doc Gardener 失败")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "started",
		"message": fmt.Sprintf("PR #%d 文档检查已启动", req.PRNumber),
		"logFile": logFileName,
	})
}

// addPRComment 添加或更新 PR 评论
func addPRComment(projectPath string, prNumber int, report string) {
	addOrEditPRComment(projectPath, prNumber, report, true)
}

// createNewPRComment 新建一条 PR 评论（不编辑已有评论）
func createNewPRComment(projectPath string, prNumber int, body string) {
	addOrEditPRComment(projectPath, prNumber, body, false)
}

func addOrEditPRComment(projectPath string, prNumber int, report string, tryEditLast bool) {
	result := extractReviewResult(report)

	// GitHub/GitLab 评论长度限制约 65536 字符，留一些余量
	maxLen := 60000
	content := report
	if len(content) > maxLen {
		content = content[:maxLen] + "\n\n---\n*报告内容过长，已在上方截断。请在 Dashboard 中查看完整报告。*"
	}

	commentBody := fmt.Sprintf(`## 🤖 Agent Code Review Report

### 审查结果: %s

%s

---
*由 Agent Control Center 自动生成*`,
		getResultEmoji(result),
		content)

	platform := getPlatform(projectPath)
	if tryEditLast {
		if err := platform.UpdateLastMRComment(prNumber, commentBody); err == nil {
			log.Printf("✅ PR/MR #%d 评论已更新", prNumber)
			return
		}
		log.Printf("ℹ️ PR/MR #%d 更新最新评论失败，尝试新建", prNumber)
	}

	if err := platform.CreateMRComment(prNumber, commentBody); err != nil {
		log.Printf("⚠️ 添加 PR/MR 评论失败: %v", err)
	} else {
		log.Printf("✅ PR/MR #%d 评论已添加", prNumber)
	}
}

// unescapeNewlines 对 Agent 偶尔输出的字面量 n/nn 做保守修复
func unescapeNewlines(s string) string {
	replacements := []struct{ old, new string }{
		{"n---", "\n---"},
		{"n## ", "\n## "},
		{"n### ", "\n### "},
		{"n#### ", "\n#### "},
		{"n> ", "\n> "},
		{"n|", "\n|"},
		{"n- ", "\n- "},
		{"n1. ", "\n1. "},
		{"n2. ", "\n2. "},
		{"n3. ", "\n3. "},
		{"n4. ", "\n4. "},
		{"n5. ", "\n5. "},
		{"n6. ", "\n6. "},
		{"n7. ", "\n7. "},
		{"n8. ", "\n8. "},
		{"n9. ", "\n9. "},
		{"n**", "\n**"},
		{"n```", "\n```"},
		{"n\n", "\n\n"},
	}
	for _, r := range replacements {
		s = strings.ReplaceAll(s, r.old, r.new)
	}
	return s
}

// unwrapHardWraps 消除 kimi 因终端宽度产生的强制折行，恢复自然段落。
// 代码块（```）和表格（|）内部会保留原有换行结构。
func unwrapHardWraps(s string) string {
	var result []string
	inCodeBlock := false
	var paraBuffer []string
	var tableBuffer []string

	flushTable := func() {
		if len(tableBuffer) == 0 {
			return
		}
		result = append(result, strings.Join(tableBuffer, " "))
		tableBuffer = nil
	}
	flushPara := func() {
		if len(paraBuffer) > 0 {
			result = append(result, strings.Join(paraBuffer, " "))
			paraBuffer = nil
		}
	}

	isBlockElement := func(line string) bool {
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "- ") ||
			strings.HasPrefix(line, "* ") || strings.HasPrefix(line, "> ") ||
			strings.HasPrefix(line, "+ ") || strings.HasPrefix(line, "---") ||
			strings.HasPrefix(line, "===") {
			return true
		}
		for _, c := range line {
			if c >= '0' && c <= '9' {
				idx := strings.Index(line, ". ")
				if idx > 0 {
					_, err := strconv.Atoi(line[:idx])
					return err == nil
				}
				return false
			}
			break
		}
		return false
	}

	isCJK := func(r rune) bool {
		return (r >= 0x4E00 && r <= 0x9FFF) ||
			(r >= 0x3400 && r <= 0x4DBF) ||
			(r >= 0x20000 && r <= 0x2A6DF) ||
			(r >= 0x2A700 && r <= 0x2B73F) ||
			(r >= 0x2B740 && r <= 0x2B81F) ||
			(r >= 0xF900 && r <= 0xFAFF) ||
			(r >= 0x3040 && r <= 0x309F) ||
			(r >= 0x30A0 && r <= 0x30FF) ||
			(r >= 0xAC00 && r <= 0xD7AF)
	}

	isInlineToken := func(r rune) bool {
		return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') ||
			r == '_' || r == '/' || r == '.' || r == '-' || r == '`'
	}

	joinWrapped := func(prev, next string) string {
		if prev == "" {
			return next
		}
		prevRunes := []rune(prev)
		lastRune := prevRunes[len(prevRunes)-1]
		firstRune := []rune(next)[0]
		if isCJK(lastRune) && isCJK(firstRune) {
			return prev + next
		}
		// 代码路径或英文标识符被硬截断时，直接拼接
		if isInlineToken(lastRune) && isInlineToken(firstRune) {
			return prev + next
		}
		return prev + " " + next
	}

	shouldMergeWithNext := func(prev, next string) bool {
		if prev == "" || next == "" {
			return false
		}
		lastChar := prev[len(prev)-1]
		if strings.ContainsRune("。！？；：.!?;:\n`", rune(lastChar)) {
			return false
		}
		if isBlockElement(next) {
			return false
		}
		return true
	}

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// 代码块边界
		if strings.HasPrefix(trimmed, "```") {
			flushTable()
			flushPara()
			inCodeBlock = !inCodeBlock
			result = append(result, line)
			continue
		}
		if inCodeBlock {
			result = append(result, line)
			continue
		}

		// 表格行检测：以 | 开头则视为表格开始/新行；如果已在表格中，持续收集直到遇到空行或块级元素
		if strings.HasPrefix(trimmed, "|") {
			if len(tableBuffer) > 0 {
				last := tableBuffer[len(tableBuffer)-1]
				if strings.HasSuffix(last, "|") {
					flushTable()
				}
			}
			tableBuffer = append(tableBuffer, trimmed)
			continue
		}
		if len(tableBuffer) > 0 {
			if trimmed == "" || isBlockElement(trimmed) {
				flushTable()
				result = append(result, line)
				continue
			}
			// 可能是表格单元格内容被折断后不含 | 的延续行
			tableBuffer = append(tableBuffer, trimmed)
			continue
		}

		// 空行
		if trimmed == "" {
			flushPara()
			result = append(result, line)
			continue
		}

		// Markdown 块级元素
		if isBlockElement(trimmed) {
			flushPara()
			result = append(result, line)
			continue
		}

		// 普通文本：尝试与上一行合并
		if len(paraBuffer) > 0 {
			last := paraBuffer[len(paraBuffer)-1]
			if shouldMergeWithNext(last, trimmed) {
				paraBuffer[len(paraBuffer)-1] = joinWrapped(last, trimmed)
				continue
			}
		}
		paraBuffer = append(paraBuffer, trimmed)
	}
	flushTable()
	flushPara()
	return strings.Join(result, "\n")
}

// extractReviewResult 严格从报告中提取审查结论
func extractReviewResult(report string) string {
	// 1. 优先匹配明确的失败结论，避免正文中的 "✅ 通过" 子串误判
	failPatterns := []string{
		"NEEDS_MAJOR_FIX",
		"NEEDS_MINOR_FIX",
		"NEEDS_FIX",
		"审查结论: 未通过",
		"审查结论：未通过",
		"审查结论: ❌",
		"审查结论：❌",
		"❌ 需修改",
		"❌需修改",
	}
	for _, p := range failPatterns {
		if strings.Contains(report, p) {
			return "needs_fix"
		}
	}

	// 2. 匹配明确的通过结论
	passPatterns := []string{
		"审查结论: PASS",
		"审查结论：**PASS**",
		"审查结论: **PASS**",
		"**审查结论**: PASS",
		"**审查结论**: **PASS**",
		"审查结论: ✅ 通过",
		"审查结论: 通过",
		"审查结论：通过",
		"审查结论: ✅通过",
		"✅ 通过",
		"✅通过",
		"LGTM",
		"lgtm",
	}
	for _, p := range passPatterns {
		if strings.Contains(report, p) {
			return "passed"
		}
	}
	// 若未匹配到明确的 NEEDS_FIX 或 PASS，默认保守返回 needs_fix
	return "needs_fix"
}

func getResultEmoji(result string) string {
	if result == "passed" {
		return "✅ 通过"
	}
	return "❌ 需修改"
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
