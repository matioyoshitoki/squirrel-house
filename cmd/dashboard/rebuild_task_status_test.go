package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestRebuild(t *testing.T) (func(), func()) {
	// Change to temp dir for file operations
	tmpDir := t.TempDir()
	origWd, err := os.Getwd()
	require.NoError(t, err)
	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Save/restore globals
	oldConfig := config
	oldIndex := currentProjectIndex
	oldTasks := make(map[string]*Task)
	tasksMutex.Lock()
	for k, v := range tasks {
		oldTasks[k] = v
	}
	for k := range tasks {
		delete(tasks, k)
	}
	tasksMutex.Unlock()

	cleanupGlobals := func() {
		config = oldConfig
		projectMutex.Lock()
		currentProjectIndex = oldIndex
		projectMutex.Unlock()

		tasksMutex.Lock()
		for k := range tasks {
			delete(tasks, k)
		}
		for k, v := range oldTasks {
			tasks[k] = v
		}
		tasksMutex.Unlock()
	}

	cleanupDir := func() {
		os.Chdir(origWd)
	}

	return cleanupGlobals, cleanupDir
}

func TestRebuildTaskStatus_FromStatusFile_Success(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("state", 0755)
	os.MkdirAll("logs/test-proj", 0755)

	statusData, _ := json.Marshal(TaskStatusFile{
		TaskID:        "dev-5",
		Type:          TaskTypeDev,
		Status:        "success",
		ProjectName:   "test-proj",
		LastHeartbeat: time.Now(),
		StartedAt:     time.Now(),
	})
	os.WriteFile("state/task-dev-5.status.json", statusData, 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:dev-5"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeDev, task.Type)
	assert.Equal(t, "success", task.Status)
	assert.Equal(t, 5, task.TargetID)
	assert.Equal(t, "test-proj", task.ProjectName)
}

func TestRebuildTaskStatus_FromStatusFile_RunningTimeout(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("state", 0755)
	os.MkdirAll("logs/test-proj", 0755)

	// Heartbeat is 10 minutes ago (older than 3 minute threshold)
	statusData, _ := json.Marshal(TaskStatusFile{
		TaskID:        "dev-7",
		Type:          TaskTypeDev,
		Status:        "running",
		ProjectName:   "test-proj",
		LastHeartbeat: time.Now().Add(-10 * time.Minute),
		StartedAt:     time.Now().Add(-20 * time.Minute),
	})
	os.WriteFile("state/task-dev-7.status.json", statusData, 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:dev-7"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	// tmux session won't exist in test, so it should be interrupted
	assert.Equal(t, "interrupted", task.Status)
}

func TestRebuildTaskStatus_FromLogInference(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("logs/test-proj", 0755)

	// Create a log file with success marker
	logContent := "some output\n✅ 任务完成\n"
	os.WriteFile("logs/test-proj/dev-issue-42-20260425-120000.log", []byte(logContent), 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:dev-42"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeDev, task.Type)
	assert.Equal(t, "success", task.Status)
	assert.Equal(t, 42, task.TargetID)
}

func TestRebuildTaskStatus_FromLogInference_NoMarker(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("logs/test-proj", 0755)

	// Create a log file without any termination marker but with content > 100 bytes
	logContent := "some ongoing output without any marker. " + strings.Repeat("more text ", 20) + "\n"
	os.WriteFile("logs/test-proj/e2e-issue-10-20260425-120000.log", []byte(logContent), 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:e2e-10"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeE2E, task.Type)
	// Log has content (>100 bytes) and is recent, should be inferred as running
	assert.Equal(t, "running", task.Status)
}

func TestRebuildTaskStatus_FromLogInference_FailedMarker(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("logs/test-proj", 0755)

	// Create a log file with failed marker
	logContent := "some output\n\n❌ 失败: something went wrong\n"
	os.WriteFile("logs/test-proj/rework-pr-99-20260425-120000.log", []byte(logContent), 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:rework-99"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeRework, task.Type)
	assert.Equal(t, "failed", task.Status)
	assert.Equal(t, 99, task.TargetID)
}

func TestRebuildTaskStatus_ReviewReport(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("logs/test-proj", 0755)

	// Create review report with a pattern that extractReviewResult recognizes
	reportContent := "## Review Report\n\n审查结论: PASS\n\nLooks good!"
	os.WriteFile("logs/test-proj/review-report-58.md", []byte(reportContent), 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:review-58"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeReview, task.Type)
	assert.Equal(t, "success", task.Status)
	assert.Equal(t, 58, task.TargetID)
	assert.Equal(t, "passed", task.Result)
	assert.NotEmpty(t, task.Report)
}

func TestRebuildTaskStatus_ReviewReport_NeedsFix(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("logs/test-proj", 0755)

	reportContent := "## Review Report\n\n**Result:** NEEDS_FIX\n\nPlease fix the issues."
	os.WriteFile("logs/test-proj/review-report-60.md", []byte(reportContent), 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:review-60"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeReview, task.Type)
	assert.Equal(t, "success", task.Status)
	assert.Equal(t, "needs_fix", task.Result)
}

func TestRebuildTaskStatus_ReviewState(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("logs/test-proj", 0755)

	stateData, _ := json.Marshal(struct {
		Status  string    `json:"status"`
		Result  string    `json:"result"`
		EndTime time.Time `json:"endTime"`
	}{
		Status:  "failed",
		Result:  "needs_fix",
		EndTime: time.Now().Add(-1 * time.Hour),
	})
	os.WriteFile("logs/test-proj/review-state-77.json", stateData, 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:review-77"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeReview, task.Type)
	assert.Equal(t, "failed", task.Status)
	assert.Equal(t, "needs_fix", task.Result)
	assert.Equal(t, 77, task.TargetID)
}

func TestRebuildTaskStatus_StatusFileTakesPriorityOverLog(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("state", 0755)
	os.MkdirAll("logs/test-proj", 0755)

	// Status file says success
	statusData, _ := json.Marshal(TaskStatusFile{
		TaskID:        "dev-20",
		Type:          TaskTypeDev,
		Status:        "success",
		ProjectName:   "test-proj",
		LastHeartbeat: time.Now(),
		StartedAt:     time.Now(),
	})
	os.WriteFile("state/task-dev-20.status.json", statusData, 0644)

	// Log file says failed
	logContent := "\n❌ 失败: error\n"
	os.WriteFile("logs/test-proj/dev-issue-20-20260425-120000.log", []byte(logContent), 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:dev-20"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	// Status file takes priority
	assert.Equal(t, "success", task.Status)
}

func TestRebuildTaskStatus_LoadsPersistedTasks(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("state", 0755)
	os.MkdirAll("logs/test-proj", 0755)

	// Persisted tasks.json
	taskData, _ := json.Marshal(map[string]*Task{
		"dev-15": {
			ID:          "dev-15",
			Type:        TaskTypeDev,
			Status:      "running",
			ProjectName: "test-proj",
			TargetID:    15,
			TargetTitle: "Issue #15",
		},
	})
	os.WriteFile("state/tasks.json", taskData, 0644)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task, ok := tasks["test-proj:dev-15"]
	tasksMutex.RUnlock()
	require.True(t, ok)
	assert.Equal(t, TaskTypeDev, task.Type)
	assert.Equal(t, "running", task.Status)
	assert.Equal(t, 15, task.TargetID)
}

func TestRebuildTaskStatus_HeartbeatCheck(t *testing.T) {
	cleanupGlobals, cleanupDir := setupTestRebuild(t)
	defer cleanupGlobals()
	defer cleanupDir()

	config = Config{
		Projects: []Project{{Name: "test-proj", Path: "/tmp/test"}},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	os.MkdirAll("logs/test-proj", 0755)

	// Create a running task without .status.json, with old log file
	tasksMutex.Lock()
	tasks["test-proj:dev-30"] = &Task{
		ID:          "dev-30",
		Type:        TaskTypeDev,
		Status:      "running",
		ProjectName: "test-proj",
		TargetID:    30,
		LogFile:     filepath.Join("logs", "test-proj", "dev-issue-30-20260425-120000.log"),
	}
	tasksMutex.Unlock()

	// Create the log file with old modification time
	logContent := "some output\n"
	logPath := "logs/test-proj/dev-issue-30-20260425-120000.log"
	os.WriteFile(logPath, []byte(logContent), 0644)
	// Set modification time to 10 minutes ago
	oldTime := time.Now().Add(-10 * time.Minute)
	os.Chtimes(logPath, oldTime, oldTime)

	rebuildTaskStatus()

	tasksMutex.RLock()
	task := tasks["test-proj:dev-30"]
	tasksMutex.RUnlock()
	require.NotNil(t, task)
	// Heartbeat timeout (log file older than 3 minutes, no tmux)
	assert.Equal(t, "interrupted", task.Status)
}
