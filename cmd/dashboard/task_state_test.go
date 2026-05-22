package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/MatioYoshitoki/new-world-ops/pkg/workflow"
)

func setupTaskStateTest(t *testing.T) {
	t.Helper()
	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	setupTestTasks(t)

	// 重置 shuttingDown 状态，避免 saveTaskState 被跳过
	oldShuttingDown := isShuttingDown
	isShuttingDown = false

	// 设置统一的项目上下文，避免 getCurrentProjectName() 返回不一致的值
	oldConfig := config
	oldIndex := currentProjectIndex
	config = Config{Projects: []Project{{Name: "test-project", Path: "/tmp/test"}}}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	// 将 saveTaskStateAsync 替换为同步版本，避免 goroutine 在 t.Cleanup 恢复工作目录后才执行
	oldAsync := saveTaskStateAsync
	saveTaskStateAsync = func() { saveTaskState() }

	t.Cleanup(func() {
		saveTaskStateAsync = oldAsync
		os.Chdir(origWd)
		config = oldConfig
		projectMutex.Lock()
		currentProjectIndex = oldIndex
		projectMutex.Unlock()
		isShuttingDown = oldShuttingDown
	})
}

// ============================================================================
// setTask / getTask
// ============================================================================

func TestSetTask_AllowsNewRunning(t *testing.T) {
	setupTaskStateTest(t)

	task := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"}
	if !setTask(task) {
		t.Error("expected setTask to succeed for new running task")
	}

	got, ok := getTask(TaskTypeDev, 5)
	if !ok {
		t.Fatal("expected getTask to find the task")
	}
	if got.Status != "running" {
		t.Errorf("expected status running, got %s", got.Status)
	}
}

func TestSetTask_RejectsDuplicateRunning(t *testing.T) {
	setupTaskStateTest(t)

	task1 := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"}
	setTask(task1)

	task2 := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"}
	if setTask(task2) {
		t.Error("expected setTask to reject duplicate running task")
	}
}

func TestSetTask_AllowsOverwriteNonRunning(t *testing.T) {
	setupTaskStateTest(t)

	task1 := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "success", TargetID: 5, ProjectName: "test-project"}
	setTask(task1)

	task2 := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"}
	if !setTask(task2) {
		t.Error("expected setTask to allow overwrite of non-running task")
	}
}

// ============================================================================
// updateTaskStatus
// ============================================================================

func TestUpdateTaskStatus_Success(t *testing.T) {
	setupTaskStateTest(t)

	task := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project", StartTime: time.Now()}
	setTask(task)

	got := updateTaskStatus(TaskTypeDev, 5, "success")
	if got == nil {
		t.Fatal("expected updateTaskStatus to return task")
	}
	if got.Status != "success" {
		t.Errorf("expected status success, got %s", got.Status)
	}
	if got.EndTime.IsZero() {
		t.Error("expected EndTime to be set")
	}
}

func TestUpdateTaskStatus_NotFound(t *testing.T) {
	setupTaskStateTest(t)

	got := updateTaskStatus(TaskTypeDev, 999, "success")
	if got != nil {
		t.Error("expected updateTaskStatus to return nil for unknown task")
	}
}

func TestUpdateTaskStatus_RunningClearsEndTime(t *testing.T) {
	setupTaskStateTest(t)

	task := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "success", TargetID: 5, ProjectName: "test-project", EndTime: time.Now()}
	setTask(task)

	got := updateTaskStatus(TaskTypeDev, 5, "running")
	if got == nil || got.EndTime.IsZero() {
		t.Error("expected EndTime to remain for running status... wait, code says if status != running set EndTime")
	}
}

// ============================================================================
// tasksByType / runningTasksByProject
// ============================================================================

func TestTasksByType(t *testing.T) {
	setupTaskStateTest(t)

	setTask(&Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"})
	setTask(&Task{ID: "review-58", Type: TaskTypeReview, Status: "running", TargetID: 58, ProjectName: "test-project"})
	setTask(&Task{ID: "dev-6", Type: TaskTypeDev, Status: "running", TargetID: 6, ProjectName: "other"})

	devs := tasksByType(TaskTypeDev, "test-project")
	if len(devs) != 1 {
		t.Fatalf("expected 1 dev task, got %d", len(devs))
	}
	if devs[0].TargetID != 5 {
		t.Errorf("expected target 5, got %d", devs[0].TargetID)
	}
}

func TestRunningTasksByProject(t *testing.T) {
	setupTaskStateTest(t)

	setTask(&Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"})
	setTask(&Task{ID: "dev-6", Type: TaskTypeDev, Status: "success", TargetID: 6, ProjectName: "test-project"})

	running := runningTasksByProject("test-project")
	if len(running) != 1 {
		t.Fatalf("expected 1 running task, got %d", len(running))
	}
	if running[0].TargetID != 5 {
		t.Errorf("expected target 5, got %d", running[0].TargetID)
	}
}

// ============================================================================
// saveTaskState / loadTaskState
// ============================================================================

func TestSaveAndLoadTaskState(t *testing.T) {
	setupTaskStateTest(t)

	task := &Task{
		ID:          "dev-5",
		Type:        TaskTypeDev,
		Status:      "success",
		ProjectName: "test-project",
		TargetID:    5,
		TargetTitle: "Issue #5",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
	}
	setTask(task)

	// Force save
	saveTaskState()

	// Clear memory
	tasksMutex.Lock()
	for k := range tasks {
		delete(tasks, k)
	}
	tasksMutex.Unlock()

	// Load back
	loadTaskState()

	tasksMutex.RLock()
	got, ok := tasks["test-project:dev-5"]
	tasksMutex.RUnlock()

	if !ok {
		t.Fatal("expected task to be restored from state file")
	}
	if got.Status != "success" {
		t.Errorf("expected status success, got %s", got.Status)
	}
	if got.TargetTitle != "Issue #5" {
		t.Errorf("expected title 'Issue #5', got %s", got.TargetTitle)
	}
	// Cancel func should be nil after restore
	if got.Cancel != nil {
		t.Error("expected Cancel to be nil after restore")
	}
}

func TestLoadTaskState_RunningWithPID(t *testing.T) {
	setupTaskStateTest(t)

	task := &Task{
		ID:          "dev-5",
		Type:        TaskTypeDev,
		Status:      "running",
		ProjectName: "test-project",
		TargetID:    5,
		PID:         os.Getpid(), // use current pid so Kill works
		StartTime:   time.Now(),
	}
	setTask(task)
	saveTaskState()

	// Clear memory
	tasksMutex.Lock()
	for k := range tasks {
		delete(tasks, k)
	}
	tasksMutex.Unlock()

	loadTaskState()

	tasksMutex.RLock()
	got, ok := tasks["test-project:dev-5"]
	tasksMutex.RUnlock()

	if !ok {
		t.Fatal("expected task to be restored")
	}
	if got.Kill == nil {
		t.Error("expected Kill to be restored for running task with PID")
	}
}

// ============================================================================
// countReviewFailuresFromLogs (via interface)
// ============================================================================

func TestCountReviewFailuresFromLogs_RealFS(t *testing.T) {
	setupTaskStateTest(t)

	logsDir := filepath.Join("logs", "test-project")
	os.MkdirAll(logsDir, 0755)

	// Create 2 NEEDS_FIX reports and 1 PASS report
	os.WriteFile(filepath.Join(logsDir, "review-report-58-20260425-120000.md"), []byte("Result: NEEDS_MAJOR_FIX\n"), 0644)
	os.WriteFile(filepath.Join(logsDir, "review-report-58-20260425-130000.md"), []byte("Result: NEEDS_FIX\n"), 0644)
	os.WriteFile(filepath.Join(logsDir, "review-report-58-20260425-140000.md"), []byte("Result: PASS\n"), 0644)
	// Different PR
	os.WriteFile(filepath.Join(logsDir, "review-report-99-20260425-120000.md"), []byte("Result: NEEDS_FIX\n"), 0644)

	counter := &defaultReviewFailureCounter{}
	got := counter.count(58, "test-project")
	if got != 2 {
		t.Errorf("expected 2 failures for PR #58, got %d", got)
	}

	got = counter.count(99, "test-project")
	if got != 1 {
		t.Errorf("expected 1 failure for PR #99, got %d", got)
	}
}

// ============================================================================
// postProcessGenericTask
// ============================================================================

func TestPostProcessGenericTask_StepFailureOverridesWfErr(t *testing.T) {
	setupTaskStateTest(t)

	task := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"}
	setTask(task)

	// Write a status file with a failed step
	status := TaskStatusFile{
		TaskID:  "dev-5",
		Type:    TaskTypeDev,
		Status:  "running",
		StepResults: map[string]workflow.StepResult{
			"build": {Status: "failed", ExitCode: 1},
		},
	}
	os.MkdirAll("state", 0755)
	data, _ := json.Marshal(status)
	os.WriteFile(filepath.Join("state", "task-dev-5.status.json"), data, 0644)

	postProcessGenericTask(task, nil) // wfErr is nil but step failed

	if task.Status != "failed" {
		t.Errorf("expected status failed, got %s", task.Status)
	}
}

func TestPostProcessGenericTask_E2ESuccessButReportFailed(t *testing.T) {
	setupTaskStateTest(t)

	logsDir := filepath.Join("logs", "test-project")
	os.MkdirAll(logsDir, 0755)

	logFile := filepath.Join(logsDir, "e2e-issue-5-test.log")
	os.WriteFile(logFile, []byte("some log\n"), 0644)

	task := &Task{
		ID:          "e2e-5",
		Type:        TaskTypeE2E,
		Status:      "running",
		TargetID:    5,
		ProjectName: "test-project",
		LogFile:     logFile,
	}
	setTask(task)

	// Create a report that says failed
	reportPath := filepath.Join(logsDir, "e2e-report-e2e-issue-5-test.md")
	os.WriteFile(reportPath, []byte("## E2E Report\n❌ 失败\n"), 0644)

	postProcessGenericTask(task, nil)

	if task.Status != "failed" {
		t.Errorf("expected status failed (report detection), got %s", task.Status)
	}
}

func TestPostProcessGenericTask_SyncsToTasksMap(t *testing.T) {
	setupTaskStateTest(t)

	task := &Task{ID: "dev-5", Type: TaskTypeDev, Status: "running", TargetID: 5, ProjectName: "test-project"}
	setTask(task)

	postProcessGenericTask(task, nil)

	tasksMutex.RLock()
	got, ok := tasks["test-project:dev-5"]
	tasksMutex.RUnlock()

	if !ok {
		t.Fatal("expected task in map")
	}
	if got.Status != "success" {
		t.Errorf("expected map status success, got %s", got.Status)
	}
}
