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
	"syscall"
	"time"

	"github.com/MatioYoshitoki/new-world-ops/pkg/workflow"
	"gopkg.in/yaml.v3"
)

func loadWorkflow(name, projectName string) (*workflow.Workflow, error) {
	// 优先查找项目级 workflow 覆盖
	if projectName != "" {
		projectPath := filepath.Join("workflows", projectName, name+".yaml")
		if data, err := os.ReadFile(projectPath); err == nil {
			var wf workflow.Workflow
			if err := yaml.Unmarshal(data, &wf); err == nil {
				log.Printf("[loadWorkflow] using project-specific workflow: %s", projectPath)
				return &wf, nil
			}
		}
	}
	// 回退到默认 workflow
	data, err := os.ReadFile(filepath.Join("workflows", name+".yaml"))
	if err != nil {
		return nil, err
	}
	var wf workflow.Workflow
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return nil, err
	}
	return &wf, nil
}

// resolveBaseRef 验证 origin/{branch} 是否存在，不存在则 fallback 到 origin/main
func resolveBaseRef(projectPath, branch string) string {
	candidate := "origin/" + branch
	cmd := exec.Command("git", "rev-parse", "--verify", candidate)
	cmd.Dir = projectPath
	if err := cmd.Run(); err == nil {
		return candidate
	}
	log.Printf("[baseRef] %s 不存在，fallback 到 origin/main", candidate)
	return "origin/main"
}

// cleanWorktreeForBranch 清理与指定分支关联的旧 worktree 和分支
func cleanWorktreeForBranch(projectPath, branchName string) error {
	// 将 projectPath 转为绝对路径，避免与 git worktree list 返回的绝对路径比较失败
	absProjectPath, err := filepath.Abs(projectPath)
	if err != nil {
		absProjectPath = projectPath
	}

	// 先 prune 失效的 worktree
	pruneCmd := exec.Command("git", "worktree", "prune")
	pruneCmd.Dir = projectPath
	pruneCmd.Run()

	// 找到并移除占用该分支的现有 worktree
	listCmd := exec.Command("git", "worktree", "list", "--porcelain")
	listCmd.Dir = projectPath
	listOutput, _ := listCmd.CombinedOutput()
	lines := strings.Split(string(listOutput), "\n")
	var currentWorktreePath string
	for _, line := range lines {
		if strings.HasPrefix(line, "worktree ") {
			currentWorktreePath = strings.TrimPrefix(line, "worktree ")
		}
		if strings.HasPrefix(line, "branch refs/heads/") && strings.TrimPrefix(line, "branch refs/heads/") == branchName && currentWorktreePath != "" {
			absWorktreePath, _ := filepath.Abs(currentWorktreePath)
			if absWorktreePath == absProjectPath {
				// 主 worktree 占用了目标分支，stash 后切回 main
				stashCmd := exec.Command("git", "stash", "push", "-m", "dashboard-auto-stash")
				stashCmd.Dir = projectPath
				if err := stashCmd.Run(); err != nil {
					// stash 失败时尝试强制 checkout（可能没什么可 stash 的）
					log.Printf("[cleanWorktree] stash failed for %s: %v, attempting force checkout", branchName, err)
				}
				checkoutCmd := exec.Command("git", "checkout", "-f", "main")
				checkoutCmd.Dir = projectPath
				if err := checkoutCmd.Run(); err != nil {
					return fmt.Errorf("checkout main failed: %w", err)
				}
			} else {
				removeCmd := exec.Command("git", "worktree", "remove", "--force", currentWorktreePath)
				removeCmd.Dir = projectPath
				if err := removeCmd.Run(); err != nil {
					// 如果 --force 移除失败，尝试直接 rm -rf 后 prune
					log.Printf("[cleanWorktree] git worktree remove --force %s failed: %v, trying rm -rf", currentWorktreePath, err)
					// 安全检查：绝不删除项目根目录
					if absWorktreePath == absProjectPath {
						log.Printf("[cleanWorktree] REFUSING to rm -rf project root: %s", absProjectPath)
					} else {
						os.RemoveAll(currentWorktreePath)
					}
					pruneCmd2 := exec.Command("git", "worktree", "prune")
					pruneCmd2.Dir = projectPath
					pruneCmd2.Run()
					// 再次检查是否还有 worktree 占用该分支
					listCmd2 := exec.Command("git", "worktree", "list", "--porcelain")
					listCmd2.Dir = projectPath
					listOutput2, _ := listCmd2.CombinedOutput()
					if strings.Contains(string(listOutput2), "branch refs/heads/"+branchName) {
						return fmt.Errorf("failed to remove worktree for branch %s at %s: %w", branchName, currentWorktreePath, err)
					}
				}
			}
			currentWorktreePath = ""
		}
	}

	// 删除已有分支（如果存在）
	delBranchCmd := exec.Command("git", "branch", "-D", branchName)
	delBranchCmd.Dir = projectPath
	if err := delBranchCmd.Run(); err != nil {
		// 分支可能不存在，忽略错误
		log.Printf("[cleanWorktree] delete branch %s: %v", branchName, err)
	}
	return nil
}

// issueTypePrompt 根据 Issue 类型返回附加指令
func issueTypePrompt(issueType string) string {
	switch issueType {
	case "bug":
		return `
【任务类型：Bug Fix】
- 这是一个 Bug 修复任务。先仔细阅读 Issue 中的复现步骤和环境信息。
- 判断影响范围：≤3 个文件且根因明确可直接修复；涉及多模块或根因不明必须先写简短 RCA/PLAN。
- 所有 bug fix 必须附带回归测试（单元测试或 E2E 测试）。
- 如果 bug 发生在已有 Maestro 流程覆盖的页面，必须更新/新增对应的 E2E 断言。
- Commit 信息建议使用 "fix: ..." 格式。
`
	case "feature":
		return `
【任务类型：Feature】
- 这是一个功能需求任务。如有 PRD，优先读取 PRD 再开始实现。
- 涉及新的用户流程（登录、注册、资料设置、新页面交互）时，必须新建 .maestro/flows/<flow>.yaml。
- 修改已有 UI 导致定位失效时，必须同步更新 .maestro/flows/ 和 .maestro/adb_flows/。
- **规范强制检查**：
  - 开发前必须读取 docs/design-docs/flutter.md（Flutter 项目）或对应技术栈规范文档
  - Design-System 组件禁止 Color(0xFF...) 硬编码，必须使用 DesignTokens
  - 需要 Maestro E2E 点击的组件必须使用 GestureDetector(onTapDown: ...) + Semantics，禁止 InkWell(onTap: ...)
  - 页面中禁止直接使用原生 Material 组件替代 design-system 组件
- **全链路验证**：
  - 写了 Socket emit 必须验证客户端能否收到（检查房间加入逻辑）
  - 写了缓存必须验证多参数场景（如 mode 切换后缓存是否隔离）
  - 写了状态机必须验证每个状态转换路径
- 确保代码、测试、文档三同步。
- Commit 信息建议使用 "feat: ..." 格式。
`
	case "performance":
		return `
【任务类型：Performance】
- 这是一个性能优化任务。先定位瓶颈，再实施优化。
- 优化前后应提供基准数据对比（QPS / P99 / 内存占用等）。
- 如涉及数据库查询优化，请同时更新相关测试和文档。
- Commit 信息建议使用 "perf: ..." 格式。
`
	case "tech-debt":
		return `
【任务类型：Tech Debt】
- 这是一个技术债务任务。先评估影响范围，再决定直接修复还是录入 tech-debt-tracker.md 排期。
- 如果是重构，确保不破坏现有接口契约；如有破坏，必须同步修改所有调用方和测试。
- Commit 信息建议使用 "refactor: ..." 或 "chore: ..." 格式。
`
	case "ui-feedback":
		return `
【任务类型：UI Feedback】
- 这是一个 UI 反馈任务。请对比设计资产（design-spec.md / Figma 等）判断是否需要修改。
- 只做 UI/交互层面的调整，不要引入无关的业务逻辑变更。
- 如涉及组件样式修改，请同步更新 Golden Test 或相关截图测试。
- Commit 信息建议使用 "style: ..." 或 "ui: ..." 格式。
`
	default:
		return `
【任务类型：通用开发】
-- 涉及新的用户流程、新页面交互或新 UI 组件时，必须新建对应的 E2E 测试（.maestro/flows/ 或相应项目的 E2E 目录）。
-- 修改已有 UI 导致定位失效时，必须同步更新 E2E 测试。
-- 用户可见的核心 UI 改动必须被对应层级的测试覆盖到（E2E > Widget Test > Unit Test）。
-- 确保代码、测试、文档三同步。
`
	}
}

// getIssueClassification 通过 gh issue view 获取 Issue 详情并分类
func getIssueClassification(issueNumber int) string {
	platform := getPlatform(getProjectPath())
	issue, err := platform.ViewIssue(issueNumber)
	if err != nil {
		log.Printf("[getIssueClassification] 获取 Issue #%d 失败: %v", issueNumber, err)
		return "full-stack"
	}
	classification := classifyIssue(*issue)
	log.Printf("[getIssueClassification] Issue #%d 分类为: %s", issueNumber, classification)
	return classification
}

// startDevTask 启动开发任务的核心逻辑
// agentType: "developer" 或 "rework"
// kind: frontend-only / backend-only / full-stack（从 Issue 标签推断）
func startDevTask(issueNumber int, issueTitle string, resume bool, prompt string, reviewReport string, agentType string, issueType string, overrideBranch string, kind string, projectName string) string {
	if agentType == "" {
		agentType = "developer"
	}

	if projectName == "" {
		projectName = getCurrentProjectName()
	}
	projectPath := getProjectPathByName(projectName)
	logsDir := getProjectLogsDir(projectName)
	os.MkdirAll(logsDir, 0755)

	timestamp := time.Now().Format("20060102-150405")
	suffix := ""
	if resume {
		suffix = "-resumed"
	}
	logFileName := filepath.Join(logsDir, fmt.Sprintf("dev-issue-%d%s-%s.log", issueNumber, suffix, timestamp))

	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Printf("❌ 创建日志文件失败: %v", err)
		return ""
	}
	logFile.Close()

	action := "启动开发"
	if agentType == "rework" {
		action = "启动返工修复"
	}
	if resume {
		action = "恢复开发"
	}

	logMsg := fmt.Sprintf("[%s] Issue #%d: %s\nAgent类型: %s\n项目路径: %s\n日志文件: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), issueNumber, action, agentType, projectPath, logFileName)
	os.WriteFile(logFileName, []byte(logMsg), 0644)

	// 后台执行 git 准备和 workflow 启动，让 API 立即返回
	go func() {
		// 确定当前分支并构造 worktree 命令
		currentBranchCmd := exec.Command("git", "branch", "--show-current")
		currentBranchCmd.Dir = projectPath
		currentBranchBytes, _ := currentBranchCmd.Output()
		currentBranch := strings.TrimSpace(string(currentBranchBytes))
		if currentBranch == "" {
			currentBranch = "main"
		}

		fetchCmd := exec.Command("git", "fetch", "origin")
		fetchCmd.Dir = projectPath
		fetchCmd.Run()

		baseRef := resolveBaseRef(projectPath, currentBranch)

		tmpDir, _ := os.MkdirTemp("", fmt.Sprintf("dev-issue-%d-*", issueNumber))
		var worktreeCmd string
		var branchName string
		if overrideBranch != "" {
			branchName = overrideBranch
			if err := cleanWorktreeForBranch(projectPath, branchName); err != nil {
				log.Printf("[startDevTask] cleanWorktreeForBranch failed: %v", err)
				os.RemoveAll(tmpDir)
				return
			}
			worktreeCmd = fmt.Sprintf("git worktree add -B %s %s %s", branchName, tmpDir, resolveBaseRef(projectPath, branchName))
		} else if agentType == "developer" {
			branchName = fmt.Sprintf("feat/issue-%d", issueNumber)
			if err := cleanWorktreeForBranch(projectPath, branchName); err != nil {
				log.Printf("[startDevTask] cleanWorktreeForBranch failed: %v", err)
				os.RemoveAll(tmpDir)
				return
			}
			worktreeCmd = fmt.Sprintf("git worktree add -b %s %s %s", branchName, tmpDir, baseRef)
		} else {
			branchName = currentBranch
			if err := cleanWorktreeForBranch(projectPath, branchName); err != nil {
				log.Printf("[startDevTask] cleanWorktreeForBranch failed: %v", err)
				os.RemoveAll(tmpDir)
				return
			}
			worktreeCmd = fmt.Sprintf("git worktree add -B %s %s %s", currentBranch, tmpDir, baseRef)
		}
		// 同步创建 worktree，确保 buildAgentContext 能扫描到文件
		wtCmd := exec.Command("bash", "-c", worktreeCmd)
		wtCmd.Dir = projectPath
		wtCmd.Run()

		// 设置 per-worktree git author，避免所有 commit 都显示为全局默认作者
		agentName := strings.Title(agentType) + " Agent"
		_ = exec.Command("git", "-C", tmpDir, "config", "user.name", agentName).Run()
		_ = exec.Command("git", "-C", tmpDir, "config", "user.email", agentType+"-agent@example.com").Run()

		agentFile := prepareAgentConfig(agentType)
		if agentFile == "" {
			agentFile = prepareAgentConfig("developer")
		}

		agentCtx := buildAgentContext(tmpDir, issueNumber, 0)
		ctxPrompt := formatAgentContextPrompt(agentCtx, issueNumber, 0)
		var taskPrompt string
		if prompt == "" {
			typePrompt := issueTypePrompt(issueType)
			commitPrefix := "feat"
			switch issueType {
			case "bug":
				commitPrefix = "fix"
			case "performance":
				commitPrefix = "perf"
			case "tech-debt":
				commitPrefix = "refactor"
			case "ui-feedback":
				commitPrefix = "style"
			}
			existingPR := prFinderImpl.findPR(branchName)
			taskPrompt, _ = renderCtxTemplate("dev_task", map[string]interface{}{
				"ContextPrompt": ctxPrompt + typePrompt,
				"IssueNumber":   issueNumber,
				"IssueTitle":    issueTitle,
				"TmpDir":        tmpDir,
				"BaseBranch":    currentBranch,
				"CommitPrefix":  commitPrefix,
				"BranchName":    branchName,
				"ExistingPR":    existingPR,
			})
		} else {
			taskPrompt = ctxPrompt + prompt
		}

		// 恢复任务时，在 prompt 前加上恢复上下文
		if resume {
			truncatedReport := ""
			if reviewReport != "" {
				truncatedReport = reviewReport[:min(len(reviewReport), 2000)]
			}
			resumeHint, _ := renderCtxTemplate("resume_hint", map[string]interface{}{
				"ReviewReport": truncatedReport,
			})
			taskPrompt = resumeHint + "\n\n" + taskPrompt
		}

		taskType := TaskTypeDev
		if agentType == "rework" {
			taskType = TaskTypeRework
		}

		task := &Task{
			ID:          taskID(taskType, issueNumber),
			Type:        taskType,
			Status:      "running",
			ProjectName: projectName,
			TargetID:    issueNumber,
			TargetTitle: issueTitle,
			Branch:      branchName,
			Kind:        kind,
			StartTime:   time.Now(),
			LogFile:     logFileName,
			Metadata: map[string]interface{}{
				"agentType":    agentType,
				"issueType":    issueType,
				"resume":       resume,
				"prompt":       taskPrompt,
				"reviewReport": reviewReport,
				"taskKind":     kind,
				"tmuxSession":  fmt.Sprintf("nwops-%s", taskID(taskType, issueNumber)),
				"worktreePath": tmpDir,
			},
		}
		if !setTask(task) {
			log.Printf("❌ 任务已在运行中: Issue #%d", issueNumber)
			return
		}

		vars := map[string]interface{}{
			"issueNumber":  issueNumber,
			"issueTitle":   issueTitle,
			"tmpDir":       tmpDir,
			"worktreeCmd":  worktreeCmd,
			"branch":       branchName,
			"logFileName":  logFileName,
			"agentFile":    agentFile,
			"prompt":       taskPrompt,
			"agentMode":    "dev",
			"reviewReport": reviewReport,
			"taskKind":     kind,
		}
		for k, v := range agentCtx {
			vars[k] = v
		}

		// rework workflow 需要额外的模板变量
		if agentType == "rework" {
			vars["prNumber"] = issueNumber
			vars["commentBody"] = fmt.Sprintf("## 🔧 返工通知\n\n此 PR 已被标记为需要返工。\n\n---\n*由 Agent Control Center 自动标记*")
		}

		runTaskWorkflow(task, vars, nil)
	}()

	return logFileName
}

// handleStartDev 启动开发任务
func handleStartDev(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		IssueNumber  int    `json:"issueNumber"`
		IssueTitle   string `json:"issueTitle"`
		Resume       bool   `json:"resume"`
		Prompt       string `json:"prompt"`
		ReviewReport string `json:"reviewReport"`
		IssueType    string `json:"issueType"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	// 检查是否已有运行中的任务
	projectName := getCurrentProjectName()
	tkey := taskKey(projectName, TaskTypeDev, req.IssueNumber)
	tasksMutex.RLock()
	if existingTask, ok := tasks[tkey]; ok && existingTask.Status == "running" {
		tasksMutex.RUnlock()
		writeJSONError(w, 409, fmt.Sprintf("Issue #%d 已在运行中", req.IssueNumber))
		return
	}
	tasksMutex.RUnlock()

	kind := getIssueClassification(req.IssueNumber)
	logFileName := startDevTask(req.IssueNumber, req.IssueTitle, req.Resume, req.Prompt, req.ReviewReport, "developer", req.IssueType, "", kind, getCurrentProjectName())
	if logFileName == "" {
		writeJSONError(w, 500, "创建日志文件失败")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "started",
		"message": fmt.Sprintf("Issue #%d 开发任务已启动", req.IssueNumber),
		"logFile": logFileName,
	})
}

// handleResumeDev 恢复开发/返工/文档检查任务
func handleResumeDev(w http.ResponseWriter, r *http.Request) {
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

	// 查找可恢复的任务（支持 dev、rework、doc-gardener 等多种类型）
	// 优先匹配 rework/doc-gardener（更具体的类型），最后匹配 dev（默认类型）
	var task *Task
	var taskType TaskType
	for _, tt := range []TaskType{TaskTypeRework, TaskTypeDocGardener, TaskTypeDev} {
		tkey := taskKey(getCurrentProjectName(), tt, req.IssueNumber)
		tasksMutex.RLock()
		if t, ok := tasks[tkey]; ok {
			task = t
			taskType = tt
		}
		tasksMutex.RUnlock()
		if task != nil {
			break
		}
	}

	if task == nil {
		writeJSONError(w, 404, fmt.Sprintf("Issue/PR #%d 没有可恢复的任务", req.IssueNumber))
		return
	}

	if task.Status == "running" {
		writeJSONError(w, 409, fmt.Sprintf("任务 #%d 正在运行中", req.IssueNumber))
		return
	}

	// 从 taskRegistry 获取对应 agent 类型
	registry, ok := taskRegistry[taskType]
	if !ok {
		writeJSONError(w, 500, fmt.Sprintf("未知的任务类型: %s", taskType))
		return
	}

	// 从已有任务恢复上下文
	prompt := ""
	reviewReport := ""
	issueType := ""
	kind := task.Kind
	if task.Metadata != nil {
		if p, ok := task.Metadata["prompt"].(string); ok {
			prompt = p
		}
		if rv, ok := task.Metadata["reviewReport"].(string); ok {
			reviewReport = rv
		}
		if it, ok := task.Metadata["issueType"].(string); ok {
			issueType = it
		}
		if k, ok := task.Metadata["taskKind"].(string); ok && kind == "" {
			kind = k
		}
	}
	if kind == "" {
		kind = "full-stack"
	}

	// 确定恢复时使用的分支
	// 对于 rework 任务，branch 应该是 PR 的 head 分支，而不是 main/master
	restoreBranch := task.Branch
	if taskType == TaskTypeRework && (restoreBranch == "main" || restoreBranch == "master" || restoreBranch == "") {
		platform := getPlatform(getProjectPath())
		if mr, err := platform.ViewMR(req.IssueNumber); err == nil && mr.SourceBranch != "" {
			restoreBranch = mr.SourceBranch
		}
	}

	// 直接启动恢复任务（复用原任务的类型和 agent）
	logFileName := startDevTask(req.IssueNumber, req.IssueTitle, true, prompt, reviewReport, registry.AgentType, issueType, restoreBranch, kind, getCurrentProjectName())
	if logFileName == "" {
		writeJSONError(w, 500, "创建恢复任务失败")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "started",
		"message": fmt.Sprintf("任务 #%d 已恢复", req.IssueNumber),
		"logFile": logFileName,
	})
}

// handleStopTask 停止运行中的任务（开发 / Review / Doc Gardener）
func handleStopTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}

	var req struct {
		IssueNumber int    `json:"issueNumber"`
		PRNumber    int    `json:"prNumber"`
		TaskType    string `json:"taskType"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}

	taskType := req.TaskType
	if taskType == "" {
		taskType = "dev"
	}

	targetID := req.IssueNumber
	if targetID == 0 {
		targetID = req.PRNumber
	}

	var tvar TaskType
	switch taskType {
	case "review":
		tvar = TaskTypeReview
	case "doc-gardener":
		tvar = TaskTypeDocGardener
	case "design":
		tvar = TaskTypeDesign
	case "rework":
		tvar = TaskTypeRework
	case "e2e":
		tvar = TaskTypeE2E
	case "pm":
		tvar = TaskTypePM
	default:
		tvar = TaskTypeDev
	}

	tkey := taskKey(getCurrentProjectName(), tvar, targetID)
	tasksMutex.Lock()
	var stopped bool
	var logFile string
	var taskKeyStr string
	var task *Task

	if t, ok := tasks[tkey]; ok && t.Status == "running" {
		task = t
	} else if req.TaskType == "" {
		// 未指定 taskType，遍历所有 running 任务按 targetId 匹配
		for _, t := range tasks {
			if t.Status == "running" && (t.TargetID == targetID || t.PRNumber == targetID) {
				task = t
				break
			}
		}
	}

	if task != nil {
		if task.Cancel != nil {
			task.Cancel()
		}
		if task.Kill != nil {
			task.Kill()
		} else if task.PID > 0 {
			// 降级：重启后恢复的任务 Kill 为 nil，直接用 PID kill
			// 由于 Setsid 创建独立进程组，向负 PID 发送 SIGKILL 级联终止所有子进程
			proc, err := os.FindProcess(task.PID)
			if err == nil && proc != nil {
				if err := syscall.Kill(-task.PID, syscall.SIGKILL); err != nil {
					proc.Kill()
				}
				appendToFile(task.LogFile, []byte("\n⚠️ 任务在 Dashboard 重启后恢复，已通过 PID 终止进程\n"))
			} else {
				appendToFile(task.LogFile, []byte("\n⚠️ 任务在 Dashboard 重启后恢复，进程可能已退出\n"))
			}
		}
		task.Status = "stopped"
		logFile = task.LogFile
		stopped = true
		go broadcastTaskChanged(task, "stopped")
	}
	tasksMutex.Unlock()

	if stopped {
		go saveTaskState()
	}

	switch tvar {
	case TaskTypeReview:
		taskKeyStr = fmt.Sprintf("PR #%d", targetID)
	case TaskTypeDocGardener:
		taskKeyStr = fmt.Sprintf("Doc Gardener PR #%d", targetID)
	case TaskTypeDesign:
		taskKeyStr = fmt.Sprintf("Design Issue #%d", targetID)
	case TaskTypeRework:
		taskKeyStr = fmt.Sprintf("Rework PR #%d", targetID)
	case TaskTypeE2E:
		taskKeyStr = fmt.Sprintf("E2E Issue #%d", targetID)
	case TaskTypePM:
		taskKeyStr = fmt.Sprintf("PM Issue #%d", targetID)
	default:
		taskKeyStr = fmt.Sprintf("Issue #%d", targetID)
	}

	if stopped && logFile != "" {
		appendToFile(logFile, []byte("\n🛑 任务已停止\n"))
	}

	if !stopped {
		writeJSONError(w, 404, "Task not found or not running")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "stopped",
		"message": fmt.Sprintf("%s 已停止", taskKeyStr),
	})
}
