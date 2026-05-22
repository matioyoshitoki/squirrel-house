package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/MatioYoshitoki/new-world-ops/pkg/workflow"
)

// TaskType 任务类型
type TaskType string

const (
	TaskTypeReview         TaskType = "review"
	TaskTypeRework         TaskType = "rework"
	TaskTypeDev            TaskType = "dev"
	TaskTypeDesign         TaskType = "design"
	TaskTypeDocGardener    TaskType = "doc-gardener"
	TaskTypePM             TaskType = "pm"
	TaskTypeE2E            TaskType = "e2e"
	TaskTypePromptEvolution TaskType = "prompt-evolution"
)

// Task 统一任务模型
type Task struct {
	ID          string                 `json:"id"`
	Type        TaskType               `json:"type"`
	Status      string                 `json:"status"` // running/success/failed/stopped/interrupted
	ProjectName string                 `json:"projectName"`
	TargetID    int                    `json:"targetId"`    // PR # 或 Issue #
	TargetTitle string                 `json:"targetTitle"` // PR 标题或 Issue 标题
	PRNumber    int                    `json:"prNumber,omitempty"`
	Merged      bool                   `json:"merged,omitempty"`
	Branch      string                 `json:"branch,omitempty"`
	Kind        string                 `json:"kind,omitempty"`   // frontend-only/backend-only/full-stack
	Result      string                 `json:"result,omitempty"` // passed/needs_fix (review)
	Report      string                 `json:"report,omitempty"`
	LogFile     string                 `json:"logFile"`
	Cancel      func()                `json:"-"`
	Kill        func() error          `json:"-"`
	PID         int                    `json:"pid"`           // 当前正在执行的物理进程 PID
	StartTime   time.Time              `json:"startTime"`
	EndTime     time.Time              `json:"endTime,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

var (
	isShuttingDown bool
	shuttingDownMu sync.RWMutex
)

// TaskStatusFile 结构化任务状态文件（取代脆弱的日志关键字推断）
type TaskStatusFile struct {
	TaskID        string                       `json:"taskID"`
	Type          TaskType                     `json:"type"`
	Status        string                       `json:"status"`
	ProjectName   string                       `json:"projectName"`
	CurrentStep   string                       `json:"currentStep,omitempty"`
	StartedAt     time.Time                    `json:"startedAt"`
	LastHeartbeat time.Time                    `json:"lastHeartbeat"`
	StepResults   map[string]workflow.StepResult `json:"stepResults,omitempty"`
	Error         string                       `json:"error,omitempty"`
}

// taskStatusFilePath 返回任务状态文件的绝对路径
func taskStatusFilePath(taskID string) string {
	return filepath.Join("state", fmt.Sprintf("task-%s.status.json", taskID))
}

// writeTaskStatusFile 将当前任务状态写入结构化状态文件
func writeTaskStatusFile(task *Task, wfCtx *workflow.Context) {
	status := TaskStatusFile{
		TaskID:        task.ID,
		Type:          task.Type,
		Status:        task.Status,
		ProjectName:   task.ProjectName,
		StartedAt:     task.StartTime,
		LastHeartbeat: time.Now(),
		StepResults:   make(map[string]workflow.StepResult),
	}

	if wfCtx != nil {
		for k, v := range wfCtx.State {
			if strings.HasPrefix(k, "stepResult_") {
				if sr, ok := v.(workflow.StepResult); ok {
					stepID := strings.TrimPrefix(k, "stepResult_")
					status.StepResults[stepID] = sr
					if sr.Status == "failed" && status.CurrentStep == "" {
						status.CurrentStep = stepID
					}
				}
			}
		}
	}

	path := taskStatusFilePath(task.ID)
	os.MkdirAll(filepath.Dir(path), 0755)
	data, _ := json.MarshalIndent(status, "", "  ")
	os.WriteFile(path, data, 0644)
}

// readTaskStatusFile 从结构化状态文件读取任务状态
func readTaskStatusFile(taskID string) *TaskStatusFile {
	path := taskStatusFilePath(taskID)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var status TaskStatusFile
	if err := json.Unmarshal(data, &status); err != nil {
		return nil
	}
	return &status
}

// updateTaskStatusFileStatus 仅更新状态文件中的 status 字段（用于重建和清理时持久化）
func updateTaskStatusFileStatus(taskID string, newStatus string) {
	path := taskStatusFilePath(taskID)
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	var status TaskStatusFile
	if err := json.Unmarshal(data, &status); err != nil {
		return
	}
	if status.Status == newStatus {
		return
	}
	status.Status = newStatus
	status.LastHeartbeat = time.Now()
	data, _ = json.MarshalIndent(status, "", "  ")
	os.WriteFile(path, data, 0644)
}

// taskRegistry 任务类型策略表
// 使用 taskSpec（定义在 hooks.go）替代匿名 struct，支持 AfterSuccess / AfterFailure hooks
var taskRegistry = map[TaskType]taskSpec{
	TaskTypeReview: {
		WorkflowName: "review",
		AgentType:    "reviewer",
		OnStarted:    func(t *Task) { notifyReviewStarted(t.TargetID, t.Branch) },
		OnFinished:   func(t *Task) { notifyReviewFinished(t.TargetID, t.Result) },
		AfterSuccess: []HookConfig{
			{
				// Review 未通过时自动触发 Rework
				TargetType: TaskTypeRework,
				Condition: func(t *Task, sourceVars map[string]interface{}, wfCtx *workflow.Context) bool {
					if t.Result == "passed" {
						log.Printf("[hook] PR #%d Review 通过，无需 Rework", t.TargetID)
						return false
					}
					log.Printf("[hook] PR #%d Review 未通过，将自动触发 Rework", t.TargetID)
					return true
				},
				Vars: map[string]string{},
			},
		},
	},
	TaskTypeRework: {
		WorkflowName: "rework",
		AgentType:    "rework",
		OnStarted:    func(t *Task) { notifyDevTaskStarted(t.TargetID, t.TargetTitle, "rework") },
		OnFinished: func(t *Task) {
			notifyDevTaskFinished(t.TargetID, t.TargetTitle, t.Status, "rework")
			if t.Status == "success" && t.PRNumber > 0 {
				clearReviewFailedState(t.PRNumber, t.ProjectName)
				createNewPRComment(getProjectPathByName(t.ProjectName), t.PRNumber, "## 修复完成\n\n已根据审查报告修复问题并推送代码，请重新 Review。\n\n---\n*由 Agent Control Center 自动生成*")
			}
		},
		AfterSuccess: []HookConfig{
			{
				// Rework 成功后自动触发 Review
				TargetType: TaskTypeReview,
				Condition: func(t *Task, sourceVars map[string]interface{}, wfCtx *workflow.Context) bool {
					prNumber := prFinderImpl.findPR(t.Branch)
					if prNumber == 0 {
						log.Printf("[hook] Rework 任务分支 %s 未找到关联 PR，跳过自动 Review", t.Branch)
						return false
					}
					log.Printf("[hook] Rework 成功，将自动触发 PR #%d Review", prNumber)
					return true
				},
				Vars: map[string]string{},
			},
		},
	},
	TaskTypeDev: {
		WorkflowName: "dev",
		AgentType:    "developer",
		OnStarted:    func(t *Task) { notifyDevTaskStarted(t.TargetID, t.TargetTitle, "developer") },
		OnFinished:   func(t *Task) { notifyDevTaskFinished(t.TargetID, t.TargetTitle, t.Status, "developer") },
		AfterSuccess: []HookConfig{
			{
				// Dev 成功后自动触发 Review
				TargetType: TaskTypeReview,
				Condition: func(t *Task, sourceVars map[string]interface{}, wfCtx *workflow.Context) bool {
					prNumber := prFinderImpl.findPR(t.Branch)
					if prNumber == 0 {
						log.Printf("[hook] Dev 任务分支 %s 未找到关联 PR，跳过自动 Review", t.Branch)
						return false
					}
					log.Printf("[hook] Dev 成功，将自动触发 PR #%d Review", prNumber)
					return true
				},
				Vars: map[string]string{},
			},
		},
	},
	TaskTypeDesign: {
		WorkflowName: "design",
		AgentType:    "ui-designer",
		OnStarted:    func(t *Task) { notifyDevTaskStarted(t.TargetID, t.TargetTitle, "ui-designer") },
		OnFinished: func(t *Task) {
			notifyDevTaskFinished(t.TargetID, t.TargetTitle, t.Status, "ui-designer")
			if t.Status == "success" {
				projectPath := getProjectPathByName(t.ProjectName)
				platform := getPlatform(projectPath)
				baseBranch := getProjectDefaultBranch()
				if baseBranch == "" {
					baseBranch = "main"
				}
				body := fmt.Sprintf("由 ui-designer agent 自动生成。\n\n- 相关 Issue: #%d\n- 设计资产目录: designs/issue-%d/\n- 包含: wireframe.svg / mockup.html / design-spec.md / preview.dart (Flutter)", t.TargetID, t.TargetID)
				if _, err := platform.CreateMR(t.Branch, baseBranch, fmt.Sprintf("design: assets for issue #%d", t.TargetID), body); err != nil {
					log.Printf("[design] 创建 MR 失败: %v", err)
				} else {
					log.Printf("[design] 已为 Issue #%d 创建设计 MR", t.TargetID)
				}
			}
		},
	},
	TaskTypeDocGardener: {
		WorkflowName: "doc-gardener",
		AgentType:    "doc-gardener",
		OnStarted:    func(t *Task) {},
		OnFinished:   func(t *Task) {},
	},
	TaskTypePM: {
		WorkflowName: "pm",
		AgentType:    "product-manager",
		OnStarted:    func(t *Task) { notifyDevTaskStarted(t.TargetID, t.TargetTitle, "product-manager") },
		OnFinished:   func(t *Task) { notifyDevTaskFinished(t.TargetID, t.TargetTitle, t.Status, "product-manager") },
	},
	TaskTypeE2E: {
		WorkflowName: "e2e",
		OnStarted:    func(t *Task) { log.Printf("E2E 任务启动: e2e-%d", t.TargetID) },
		OnFinished: func(t *Task) {
			log.Printf("E2E 任务结束: e2e-%d, 状态: %s", t.TargetID, t.Status)

			// 1. 生成 E2E 报告（返回路径和内容）
			_, reportContent := generateE2EReport(t)
			if reportContent == "" {
				reportContent = fmt.Sprintf("E2E 测试完成，状态: %s\n\n日志: %s", t.Status, filepath.Base(t.LogFile))
			}

			// 2. 飞书通知（推送报告原文）
			// E2E 是独立任务，不再关联 PR 评论
			notifyE2EFinished(t.TargetID, t.TargetTitle, t.Status, reportContent, 0)
		},
		// E2E 是独立任务，成功后不再自动触发 Review
		AfterSuccess: []HookConfig{},
	},
	TaskTypePromptEvolution: {
		WorkflowName: "prompt-evolution",
		AgentType:    "prompt-evolution",
		OnStarted:    func(t *Task) { log.Printf("[prompt-evolution] 任务 #%d 启动", t.TargetID) },
		OnFinished:   func(t *Task) { log.Printf("[prompt-evolution] 任务 #%d 结束: %s", t.TargetID, t.Status) },
	},
}

func taskStateFile() string {
	return filepath.Join("state", "tasks.json")
}

// loadTaskState 从磁盘加载任务状态到内存
func loadTaskState() {
	stateFile := taskStateFile()
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return
	}
	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	var loaded map[string]*Task
	if err := json.Unmarshal(data, &loaded); err != nil {
		log.Printf("⚠️ 加载任务状态文件失败: %v", err)
		return
	}
	for k, v := range loaded {
		if v != nil {
			v.Cancel = nil // func 不可序列化，恢复后清空
			// 为 running 且 PID > 0 的任务恢复 Kill 能力
			if v.Status == "running" && v.PID > 0 {
				pid := v.PID
				v.Kill = func() error {
					proc, err := os.FindProcess(pid)
					if err != nil {
						return err
					}
					return proc.Kill()
				}
			}
			// 键迁移：旧格式 "dev-5" → 新格式 "projectName:dev-5"
			key := k
			if !strings.Contains(k, ":") {
				projectName := v.ProjectName
				if projectName == "" {
					projectName = getCurrentProjectName()
				}
				parts := strings.SplitN(k, "-", 2)
				if len(parts) == 2 {
					key = fmt.Sprintf("%s:%s", projectName, k)
				}
			}
			tasks[key] = v
		}
	}
	log.Printf("📝 已从 %s 恢复 %d 个任务状态", stateFile, len(tasks))
}

// saveTaskState 将内存任务状态持久化到磁盘
func getShuttingDown() bool {
	shuttingDownMu.RLock()
	defer shuttingDownMu.RUnlock()
	return isShuttingDown
}

func saveTaskState() {
	if getShuttingDown() {
		return
	}

	stateFile := taskStateFile()
	if err := os.MkdirAll(filepath.Dir(stateFile), 0755); err != nil {
		log.Printf("⚠️ 创建状态目录失败: %v", err)
		return
	}

	// 先加载现有状态，避免覆盖其他项目的任务
	var allTasks map[string]*Task
	if data, err := os.ReadFile(stateFile); err == nil {
		json.Unmarshal(data, &allTasks)
	}
	if allTasks == nil {
		allTasks = make(map[string]*Task)
	}

	tasksMutex.RLock()
	// 更新所有任务（键已经是 project-scoped 格式）
	for k, v := range tasks {
		allTasks[k] = v
	}
	tasksMutex.RUnlock()

	data, err := json.MarshalIndent(allTasks, "", "  ")
	if err != nil {
		log.Printf("⚠️ 序列化任务状态失败: %v", err)
		return
	}
	if err := os.WriteFile(stateFile, data, 0644); err != nil {
		log.Printf("⚠️ 保存任务状态文件失败: %v", err)
	}
}

func taskID(taskType TaskType, targetID int) string {
	return fmt.Sprintf("%s-%d", taskType, targetID)
}

// taskKey 生成带项目作用域的任务键，解决跨项目 taskID 冲突问题
func taskKey(projectName string, taskType TaskType, targetID int) string {
	return fmt.Sprintf("%s:%s-%d", projectName, taskType, targetID)
}

// getTask 从统一任务表中读取任务（当前项目作用域）
func getTask(taskType TaskType, targetID int) (*Task, bool) {
	tasksMutex.RLock()
	defer tasksMutex.RUnlock()
	t, ok := tasks[taskKey(getCurrentProjectName(), taskType, targetID)]
	return t, ok
}

// saveTaskStateAsync 允许测试替换为同步版本，避免 goroutine 在 t.Cleanup 恢复工作目录后才执行
var saveTaskStateAsync = func() { go saveTaskState() }

// setTask 向统一任务表中写入任务（如果已存在 running 任务，返回 false）
func setTask(task *Task) bool {
	tasksMutex.Lock()
	key := taskKey(task.ProjectName, task.Type, task.TargetID)
	if existing, ok := tasks[key]; ok && existing.Status == "running" {
		tasksMutex.Unlock()
		return false
	}
	tasks[key] = task
	tasksMutex.Unlock()
	saveTaskStateAsync()
	go broadcastTaskChanged(task, "started")
	return true
}

// countReviewFailuresFromLogs 扫描日志目录中的 review report 文件，统计该 PR 的非 PASS review 次数
func countReviewFailuresFromLogs(prNumber int, projectName string) int {
	logsDir := getProjectLogsDir(projectName)
	entries, err := os.ReadDir(logsDir)
	if err != nil {
		return 0
	}
	count := 0
	prefix := fmt.Sprintf("review-report-%d-", prNumber)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), prefix) || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		path := filepath.Join(logsDir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil || len(data) == 0 {
			continue
		}
		content := string(data)
		if strings.Contains(content, "NEEDS_FIX") || strings.Contains(content, "NEEDS_MAJOR_FIX") {
			count++
		}
	}
	return count
}

// updateTaskStatus 更新任务状态
func updateTaskStatus(taskType TaskType, targetID int, status string) *Task {
	tasksMutex.Lock()
	if t, ok := tasks[taskKey(getCurrentProjectName(), taskType, targetID)]; ok {
		oldStatus := t.Status
		t.Status = status
		if status != "running" {
			t.EndTime = time.Now()
		}
		tasksMutex.Unlock()
		saveTaskStateAsync()
		if oldStatus != status {
			go broadcastTaskChanged(t, "status_changed")
		}
		return t
	}
	tasksMutex.Unlock()
	return nil
}

// runTaskWorkflow 统一的任务执行入口
func runTaskWorkflow(task *Task, vars map[string]interface{}, preRun func(*Task) error) {
	go func() {
		projectPath := getProjectPathByName(task.ProjectName)
		logsDir := getProjectLogsDir(task.ProjectName)
		spec, ok := taskRegistry[task.Type]
		if !ok {
			appendToFile(task.LogFile, []byte(fmt.Sprintf("\n❌ 未知任务类型: %s\n", task.Type)))
			updateTaskStatus(task.Type, task.TargetID, "failed")
			if spec.OnFinished != nil {
				spec.OnFinished(task)
			}
			return
		}

		if preRun != nil {
			if err := preRun(task); err != nil {
				appendToFile(task.LogFile, []byte(fmt.Sprintf("\n❌ 前置准备失败: %v\n", err)))
				updateTaskStatus(task.Type, task.TargetID, "failed")
				if spec.OnFinished != nil {
					spec.OnFinished(task)
				}
				return
			}
		}

		if spec.OnStarted != nil {
			spec.OnStarted(task)
		}

		wf, err := loadWorkflow(spec.WorkflowName, task.ProjectName)
		if err != nil {
			appendToFile(task.LogFile, []byte(fmt.Sprintf("\n❌ 加载 workflow 失败: %v\n", err)))
			updateTaskStatus(task.Type, task.TargetID, "failed")
			if spec.OnFinished != nil {
				spec.OnFinished(task)
			}
			return
		}

		// 传递 taskID 给 workflow，用于 tmux session 命名
		vars["taskID"] = task.ID

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		wfCtx := workflow.NewContext(projectPath, vars)
		engine := workflow.NewEngine(wfCtx)
		engine.OnProcessStart = func(pid int) {
			tasksMutex.Lock()
			key := taskKey(task.ProjectName, task.Type, task.TargetID)
			if t, ok := tasks[key]; ok {
				t.PID = pid
			}
			tasksMutex.Unlock()
			go saveTaskState() // PID 更新后立即持久化
		}

		// 构造 tmux session 名称（用于 Kill 和监控，与 workflow 中的命名一致）
		tmuxSession := fmt.Sprintf("nwops-%s", task.ID)

		tasksMutex.Lock()
		key := taskKey(task.ProjectName, task.Type, task.TargetID)
		if t, ok := tasks[key]; ok {
			t.Cancel = cancel
			t.Kill = func() error {
				// 先尝试 kill tmux session（长时间任务在 tmux 中运行）
				if err := exec.Command("tmux", "kill-session", "-t", tmuxSession).Run(); err == nil {
					return nil
				}
				// fallback：kill 当前 shell 进程
				return engine.Kill()
			}
			if t.Metadata == nil {
				t.Metadata = make(map[string]interface{})
			}
			t.Metadata["tmuxSession"] = tmuxSession
		}
		tasksMutex.Unlock()

		// 启动 heartbeat：定期刷新结构化状态文件
		heartbeatTicker := time.NewTicker(30 * time.Second)
		defer heartbeatTicker.Stop()
		heartbeatDone := make(chan struct{})
		stopHeartbeat := make(chan struct{})
		go func() {
			defer close(heartbeatDone)
			for {
				select {
				case <-heartbeatTicker.C:
					writeTaskStatusFile(task, wfCtx)
				case <-stopHeartbeat:
					return
				case <-ctx.Done():
					return
				}
			}
		}()

		wfErr := engine.Run(ctx, wf)

		// 停止 heartbeat
		heartbeatTicker.Stop()
		close(stopHeartbeat)
		<-heartbeatDone

		// 最终状态写入（步骤级结果）
		writeTaskStatusFile(task, wfCtx)

		// 特定类型后处理（更新 task.Status 为 success/failed）
		if task.Type == TaskTypeReview {
			postProcessReviewTask(task, wfCtx, wfErr, logsDir)
		} else {
			postProcessGenericTask(task, wfErr)
		}

		// 再次写入状态文件，确保最终状态（success/failed）被持久化
		writeTaskStatusFile(task, wfCtx)

		// 触发 hooks（在 OnFinished 之前，让 hook 任务和通知并行）
		triggerHooks(task, spec, vars, wfCtx, task.Status == "success")

		// cleanup 由 workflow isCleanup 步骤自动处理
		if spec.OnFinished != nil {
			spec.OnFinished(task)
		}

		// 检查是否需要触发 prompt 演进
		go checkAndTriggerPromptEvolution(task.ProjectName)
	}()
}

// postProcessGenericTask 通用任务后处理
// 优先信任结构化状态文件中的步骤级结果，解决 wfErr 为 nil 但步骤实际失败的问题
func postProcessGenericTask(task *Task, wfErr error) {
	if task.Status == "stopped" {
		appendToFile(task.LogFile, []byte("\n⏹️ 任务已手动停止\n"))
	} else {
		// 从结构化状态文件读取步骤级结果
		statusFile := readTaskStatusFile(task.ID)
		failedStep := ""
		failedExitCode := 0
		if statusFile != nil {
			for stepID, sr := range statusFile.StepResults {
				if sr.Status == "failed" {
					failedStep = stepID
					failedExitCode = sr.ExitCode
					break
				}
			}
		}

		if failedStep != "" {
			task.Status = "failed"
			appendToFile(task.LogFile, []byte(fmt.Sprintf("\n❌ %s失败: 步骤 %s 返回 exit code %d\n", task.Type, failedStep, failedExitCode)))
		} else if wfErr != nil {
			task.Status = "failed"
			appendToFile(task.LogFile, []byte(fmt.Sprintf("\n❌ %s失败: %v\n", task.Type, wfErr)))
		} else {
			task.Status = "success"
			appendToFile(task.LogFile, []byte("\n✅ 任务完成\n"))
		}

		// E2E 任务特殊处理：agent 总是返回 exit code 0，需检查报告中的实际测试结果
		if task.Type == TaskTypeE2E && task.Status == "success" {
			_, reportContent := generateE2EReport(task)
			if strings.Contains(reportContent, "❌ 失败") || strings.Contains(reportContent, "Exit Code: 1") {
				log.Printf("[e2e] e2e-%d 报告检测到失败，将状态修正为 failed", task.TargetID)
				task.Status = "failed"
				appendToFile(task.LogFile, []byte("\n❌ E2E 测试报告检测到失败，将状态修正为 failed\n"))
			}
		}
	}
	task.EndTime = time.Now()

	// 同步到 tasks map（如果存在），保证前端状态一致
	tasksMutex.Lock()
	key := taskKey(task.ProjectName, task.Type, task.TargetID)
	if t, ok := tasks[key]; ok {
		t.Status = task.Status
		t.EndTime = task.EndTime
	}
	tasksMutex.Unlock()

	go broadcastTaskChanged(task, "finished")
	go saveTaskState()
}

// postProcessReviewTask Review 任务特有后处理
func postProcessReviewTask(task *Task, wfCtx *workflow.Context, wfErr error, logsDir string) {
	var output string
	if v, ok := wfCtx.State["agentOutput"]; ok {
		output = v.(string)
	}

	workDir := ""
	if v, ok := wfCtx.Vars["tmpDir"]; ok {
		workDir = v.(string)
	}

	// 读取报告：优先从 workflow state 读取（确保拿到当前 review 生成的最新报告）
	var reportStr string
	reportFromFile := false
	dashboardReportPath := filepath.Join(logsDir, fmt.Sprintf("review-report-%d.md", task.TargetID))
	worktreeReportPath := filepath.Join(workDir, fmt.Sprintf("logs/review-report-%d.md", task.TargetID))

	if v, ok := wfCtx.State["reviewReportContent"]; ok && v != "" {
		reportStr = v.(string)
		reportFromFile = true
		os.WriteFile(dashboardReportPath, []byte(reportStr), 0644)
		log.Printf("📄 PR #%d Review 报告来自 workflow state (readReport 步骤)", task.TargetID)
	} else if data, err := os.ReadFile(worktreeReportPath); err == nil && len(data) > 0 {
		reportStr = string(data)
		reportFromFile = true
		os.WriteFile(dashboardReportPath, data, 0644)
		log.Printf("📄 PR #%d Review 报告来自 worktree 文件", task.TargetID)
	} else if data, err := os.ReadFile(dashboardReportPath); err == nil && len(data) > 0 {
		reportStr = string(data)
		reportFromFile = true
		log.Printf("📄 PR #%d Review 报告来自 dashboard 日志目录（旧报告）", task.TargetID)
	} else if output != "" {
		reportStr = unescapeNewlines(output)
		os.WriteFile(dashboardReportPath, []byte(reportStr), 0644)
		log.Printf("⚠️ PR #%d Review 报告来自 agent stdout fallback", task.TargetID)
	}

	reportStr = unwrapHardWraps(unescapeNewlines(reportStr))
	result := extractReviewResult(reportStr)

	// 保存带时间戳的 review report，用于历史聚合分析
	if reportStr != "" {
		timestamp := time.Now().Format("20060102-150405")
		historicalReportPath := filepath.Join(logsDir, fmt.Sprintf("review-report-%d-%s.md", task.TargetID, timestamp))
		os.WriteFile(historicalReportPath, []byte(reportStr), 0644)
		log.Printf("📄 PR #%d Review 报告已保存到历史文件: %s", task.TargetID, historicalReportPath)
	}

	if task.Status == "stopped" {
		appendToFile(task.LogFile, []byte("\n⏹️ Review 已被手动停止\n"))
	} else if wfErr != nil {
		task.Status = "failed"
		appendToFile(task.LogFile, []byte(fmt.Sprintf("\n❌ Review 失败: %v\n", wfErr)))
	} else {
		task.Status = "success"
		appendToFile(task.LogFile, []byte("\n✅ Review 完成\n"))
	}
	task.Result = result
	task.Report = reportStr
	task.EndTime = time.Now()

	// 同步到 tasks map（如果存在），保证前端状态一致
	tasksMutex.Lock()
	key := taskKey(task.ProjectName, task.Type, task.TargetID)
	if t, ok := tasks[key]; ok {
		t.Status = task.Status
		t.Result = task.Result
		t.Report = task.Report
		t.EndTime = task.EndTime
	}
	tasksMutex.Unlock()

	go broadcastTaskChanged(task, "finished")

	go saveTaskState()

	log.Printf("Review 任务结束: PR #%d, 结果: %s, 错误: %v", task.TargetID, result, wfErr)

	if reportStr != "" && reportFromFile {
		addPRComment(getProjectPath(), task.TargetID, reportStr)
	} else {
		logContent, _ := os.ReadFile(task.LogFile)
		content := string(logContent)
		if reportStr != "" {
			content = reportStr
		}
		if len(content) > 5000 {
			content = "...\n" + content[len(content)-5000:]
		}
		fallbackReport := fmt.Sprintf("## 🤖 Agent Code Review Report\n\n⚠️ 本次 Review 未生成结构化报告。\n\n### 审查日志摘要\n\n```\n%s\n```\n\n---\n*由 Agent Control Center 自动生成*", content)
		addPRComment(getProjectPath(), task.TargetID, fallbackReport)
	}
}

// tasksByType 获取指定类型的所有任务（按项目过滤）
func tasksByType(taskType TaskType, projectName string) []*Task {
	tasksMutex.RLock()
	defer tasksMutex.RUnlock()
	result := make([]*Task, 0)
	for _, t := range tasks {
		if t.Type == taskType && t.ProjectName == projectName {
			result = append(result, t)
		}
	}
	return result
}

// runningTasksByProject 获取当前项目所有非 success 状态的任务（用于前端展示 running / interrupted / failed）
func runningTasksByProject(projectName string) []*Task {
	tasksMutex.RLock()
	defer tasksMutex.RUnlock()
	result := make([]*Task, 0)
	for _, t := range tasks {
		if t.ProjectName == projectName && t.Status == "running" {
			result = append(result, t)
		}
	}
	return result
}

// appendToFile 追加内容到文件
func appendToFile(filename string, data []byte) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	f.Write(data)
}

func setShuttingDown(v bool) {
	shuttingDownMu.Lock()
	isShuttingDown = v
	shuttingDownMu.Unlock()
}

func waitForRunningTasksWithTimeout(maxWait time.Duration, pollInterval time.Duration) {
	deadline := time.Now().Add(maxWait)
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tasksMutex.RLock()
			runningCount := 0
			for _, task := range tasks {
				if task.Status == "running" {
					runningCount++
				}
			}
			tasksMutex.RUnlock()

			if runningCount == 0 {
				log.Printf("✅ 所有 running 任务已完成，旧进程准备退出")
				return
			}

			if time.Now().After(deadline) {
				log.Printf("⚠️ 等待 running 任务超时（%v），强制退出", maxWait)
				return
			}

			log.Printf("⏳ 旧进程等待 %d 个 running 任务完成...", runningCount)
		}
	}
}

func waitForRunningTasks() {
	waitForRunningTasksWithTimeout(30*time.Minute, 5*time.Second)
}
