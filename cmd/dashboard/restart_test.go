package main

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============================================================================
// waitForRunningTasksWithTimeout
// ============================================================================

func TestWaitForRunningTasks_NoTasks(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	// No running tasks — should return immediately
	done := make(chan struct{})
	go func() {
		waitForRunningTasksWithTimeout(30*time.Second, 50*time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		// PASS
	case <-time.After(2 * time.Second):
		t.Fatal("waitForRunningTasks should return immediately when no tasks are running")
	}
}

func TestWaitForRunningTasks_TasksFinish(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	tasksMutex.Lock()
	tasks["default:dev-1"] = &Task{ID: "dev-1", ProjectName: "default", Status: "running"}
	tasksMutex.Unlock()

	done := make(chan struct{})
	go func() {
		waitForRunningTasksWithTimeout(30*time.Second, 50*time.Millisecond)
		close(done)
	}()

	// Simulate task finishing after a short delay
	time.Sleep(150 * time.Millisecond)
	tasksMutex.Lock()
	tasks["default:dev-1"].Status = "success"
	tasksMutex.Unlock()

	select {
	case <-done:
		// PASS
	case <-time.After(2 * time.Second):
		t.Fatal("waitForRunningTasks should return after tasks finish")
	}
}

func TestWaitForRunningTasks_Timeout(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	tasksMutex.Lock()
	tasks["default:dev-1"] = &Task{ID: "dev-1", ProjectName: "default", Status: "running"}
	tasksMutex.Unlock()

	done := make(chan struct{})
	go func() {
		waitForRunningTasksWithTimeout(200*time.Millisecond, 50*time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		// PASS — should have timed out
	case <-time.After(2 * time.Second):
		t.Fatal("waitForRunningTasks should timeout")
	}

	// Task should still be running (we didn't change it)
	tasksMutex.RLock()
	status := tasks["default:dev-1"].Status
	tasksMutex.RUnlock()
	assert.Equal(t, "running", status)
}

func TestWaitForRunningTasks_MultipleTasks(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	tasksMutex.Lock()
	tasks["default:dev-1"] = &Task{ID: "dev-1", ProjectName: "default", Status: "running"}
	tasks["default:dev-2"] = &Task{ID: "dev-2", ProjectName: "default", Status: "running"}
	tasksMutex.Unlock()

	done := make(chan struct{})
	go func() {
		waitForRunningTasksWithTimeout(30*time.Second, 50*time.Millisecond)
		close(done)
	}()

	// Finish one task
	time.Sleep(100 * time.Millisecond)
	tasksMutex.Lock()
	tasks["default:dev-1"].Status = "success"
	tasksMutex.Unlock()

	// Should not return yet — dev-2 still running
	select {
	case <-done:
		t.Fatal("should not return while dev-2 is still running")
	case <-time.After(200 * time.Millisecond):
		// expected
	}

	// Finish the second task
	tasksMutex.Lock()
	tasks["default:dev-2"].Status = "success"
	tasksMutex.Unlock()

	select {
	case <-done:
		// PASS
	case <-time.After(2 * time.Second):
		t.Fatal("should return after all tasks finish")
	}
}

// ============================================================================
// isShuttingDown / saveTaskState
// ============================================================================

func TestGetShuttingDown(t *testing.T) {
	// Default state
	assert.False(t, getShuttingDown())

	setShuttingDown(true)
	assert.True(t, getShuttingDown())

	setShuttingDown(false)
	assert.False(t, getShuttingDown())
}

func TestSaveTaskState_ShuttingDown(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	// Change to temp dir
	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origWd)

	// Pre-create a state file
	os.MkdirAll("state", 0755)
	os.WriteFile("state/tasks.json", []byte("old-data"), 0644)

	// Populate tasks
	tasksMutex.Lock()
	tasks["default:dev-1"] = &Task{ID: "dev-1", ProjectName: "default", Status: "running"}
	tasksMutex.Unlock()

	// Set shutting down flag
	setShuttingDown(true)
	defer setShuttingDown(false)

	saveTaskState()

	// File should NOT have been overwritten
	data, _ := os.ReadFile("state/tasks.json")
	assert.Equal(t, "old-data", string(data))
}

func TestSaveTaskState_Normal(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	// Change to temp dir
	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origWd)

	// Populate tasks
	tasksMutex.Lock()
	tasks["default:dev-5"] = &Task{
		ID:          "dev-5",
		Type:        TaskTypeDev,
		Status:      "success",
		ProjectName: "default",
		TargetID:    5,
	}
	tasksMutex.Unlock()

	setShuttingDown(false)
	saveTaskState()

	data, err := os.ReadFile("state/tasks.json")
	require.NoError(t, err)
	assert.Contains(t, string(data), "default:dev-5")
	assert.Contains(t, string(data), "success")
}

// ============================================================================
// getListenerWithEnv
// ============================================================================

func TestGetListenerWithEnv_Fresh(t *testing.T) {
	envLookup := func(string) string { return "" }
	ln, err := getListenerWithEnv("127.0.0.1:0", envLookup)
	require.NoError(t, err)
	defer ln.Close()

	assert.NotNil(t, ln.Addr())
	assert.NotEmpty(t, ln.Addr().String())
}

func TestGetListenerWithEnv_FdInheritance(t *testing.T) {
	// Create a TCP listener and get its underlying file descriptor
	parentLn, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer parentLn.Close()

	tcpLn := parentLn.(*net.TCPListener)
	file, err := tcpLn.File()
	require.NoError(t, err)
	// Do NOT defer file.Close() here — we need the fd to remain valid
	// for getListenerWithEnv. net.FileListener will dup it internally.

	fd := int(file.Fd())
	envLookup := func(key string) string {
		if key == "DASHBOARD_FD" {
			return strconv.Itoa(fd)
		}
		return ""
	}

	// Create listener from inherited fd
	childLn, err := getListenerWithEnv("127.0.0.1:0", envLookup)
	require.NoError(t, err)
	require.NotNil(t, childLn)
	defer childLn.Close()

	// Verify the inherited listener is usable by accepting a connection
	acceptDone := make(chan struct{})
	go func() {
		conn, err := childLn.Accept()
		if err == nil && conn != nil {
			conn.Close()
		}
		close(acceptDone)
	}()

	conn, err := net.Dial("tcp", childLn.Addr().String())
	require.NoError(t, err)
	conn.Close()

	select {
	case <-acceptDone:
		// PASS — inherited listener accepted the connection
	case <-time.After(2 * time.Second):
		t.Fatal("inherited listener did not accept connection")
	}
}

func TestGetListenerWithEnv_InvalidFd(t *testing.T) {
	envLookup := func(key string) string {
		if key == "DASHBOARD_FD" {
			return "not-a-number"
		}
		return ""
	}

	// Invalid DASHBOARD_FD should fall back to creating a new listener
	ln, err := getListenerWithEnv("127.0.0.1:0", envLookup)
	require.NoError(t, err)
	defer ln.Close()

	assert.NotNil(t, ln.Addr())
}

func TestGetListenerWithEnv_NonexistentFd(t *testing.T) {
	envLookup := func(key string) string {
		if key == "DASHBOARD_FD" {
			return "99999"
		}
		return ""
	}

	// Non-existent fd should fall back to creating a new listener
	ln, err := getListenerWithEnv("127.0.0.1:0", envLookup)
	require.NoError(t, err)
	defer ln.Close()

	assert.NotNil(t, ln.Addr())
}

// ============================================================================
// taskStateFile (helper used by saveTaskState)
// ============================================================================

func TestTaskStateFile(t *testing.T) {
	// Verify taskStateFile returns expected path
	path := taskStateFile()
	assert.Equal(t, filepath.Join("state", "tasks.json"), path)
}
