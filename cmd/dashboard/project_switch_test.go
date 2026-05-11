package main

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestProjectSwitch(t *testing.T) func() {
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

	return func() {
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
}

func TestGetProjectPath(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "p1", Path: "/path/one"},
			{Name: "p2", Path: "/path/two"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = 1
	projectMutex.Unlock()

	assert.Equal(t, "/path/two", getProjectPath())
}

func TestGetProjectPath_OutOfBounds(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "p1", Path: "/path/one"},
			{Name: "p2", Path: "/path/two"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = 99
	projectMutex.Unlock()

	assert.Equal(t, "/path/one", getProjectPath())
}

func TestGetProjectPath_EmptyProjects(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{Projects: []Project{}}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	assert.Equal(t, ".", getProjectPath())
}

func TestGetCurrentProjectName(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "alpha", Path: "/a"},
			{Name: "beta", Path: "/b"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	assert.Equal(t, "alpha", getCurrentProjectName())
}

func TestGetCurrentProjectName_OutOfBounds(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "alpha", Path: "/a"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = -1
	projectMutex.Unlock()

	assert.Equal(t, "alpha", getCurrentProjectName())
}

func TestGetCurrentProjectName_EmptyProjects(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{Projects: []Project{}}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	assert.Equal(t, "default", getCurrentProjectName())
}

func TestGetProjectPathByName(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "alpha", Path: "/a"},
			{Name: "beta", Path: "/b"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	assert.Equal(t, "/b", getProjectPathByName("beta"))
}

func TestGetProjectPathByName_NotFound(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "alpha", Path: "/a"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	assert.Equal(t, "/a", getProjectPathByName("nonexistent"))
}

func TestLoadState(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "p0", Path: "/0"},
			{Name: "p1", Path: "/1"},
			{Name: "p2", Path: "/2"},
		},
	}

	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origWd)

	os.MkdirAll("configs", 0755)
	stateData, _ := json.Marshal(map[string]interface{}{
		"currentProjectIndex": 2,
	})
	os.WriteFile("configs/dashboard-state.json", stateData, 0644)

	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	loadState()

	projectMutex.RLock()
	idx := currentProjectIndex
	projectMutex.RUnlock()
	assert.Equal(t, 2, idx)
}

func TestLoadState_InvalidIndex(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "p0", Path: "/0"},
		},
	}

	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origWd)

	os.MkdirAll("configs", 0755)
	stateData, _ := json.Marshal(map[string]interface{}{
		"currentProjectIndex": 99,
	})
	os.WriteFile("configs/dashboard-state.json", stateData, 0644)

	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	loadState()

	projectMutex.RLock()
	idx := currentProjectIndex
	projectMutex.RUnlock()
	assert.Equal(t, 0, idx) // should remain unchanged
}

func TestLoadState_MalformedJSON(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{{Name: "p0", Path: "/0"}},
	}

	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origWd)

	os.MkdirAll("configs", 0755)
	os.WriteFile("configs/dashboard-state.json", []byte("not-json"), 0644)

	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	// Should not panic
	loadState()

	projectMutex.RLock()
	idx := currentProjectIndex
	projectMutex.RUnlock()
	assert.Equal(t, 0, idx)
}

func TestLoadState_FileNotFound(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{{Name: "p0", Path: "/0"}},
	}

	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origWd)

	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	// No state file exists
	loadState()

	projectMutex.RLock()
	idx := currentProjectIndex
	projectMutex.RUnlock()
	assert.Equal(t, 0, idx)
}

func TestSaveState(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "p0", Path: "/0"},
			{Name: "p1", Path: "/1"},
		},
	}

	tmpDir := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origWd)

	projectMutex.Lock()
	currentProjectIndex = 1
	projectMutex.Unlock()

	saveState()

	data, err := os.ReadFile("configs/dashboard-state.json")
	require.NoError(t, err)

	var state struct {
		CurrentProjectIndex int `json:"currentProjectIndex"`
	}
	err = json.Unmarshal(data, &state)
	require.NoError(t, err)
	assert.Equal(t, 1, state.CurrentProjectIndex)
}

func TestHandleSwitchProject(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "alpha", Path: "/a"},
			{Name: "beta", Path: "/b"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = 0
	projectMutex.Unlock()

	// Pre-populate a task
	tasksMutex.Lock()
	tasks["alpha:dev-1"] = &Task{ID: "dev-1", ProjectName: "alpha", Status: "running"}
	tasksMutex.Unlock()

	body := `{"projectName":"beta"}`
	req := httptest.NewRequest("POST", "/api/switch-project", strings.NewReader(body))
	w := httptest.NewRecorder()

	handleSwitchProject(w, req)

	assert.Equal(t, 200, w.Code)

	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, "switched", resp["status"])

	projectMutex.RLock()
	idx := currentProjectIndex
	projectMutex.RUnlock()
	assert.Equal(t, 1, idx)

	tasksMutex.RLock()
	count := len(tasks)
	tasksMutex.RUnlock()
	assert.Equal(t, 0, count)

	// Wait for async rebuildTaskStatus goroutine to finish to avoid
	// interfering with subsequent tests that manipulate state/tasks.json.
	time.Sleep(100 * time.Millisecond)
}

func TestHandleSwitchProject_NotFound(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{{Name: "alpha", Path: "/a"}},
	}

	body := `{"projectName":"nonexistent"}`
	req := httptest.NewRequest("POST", "/api/switch-project", strings.NewReader(body))
	w := httptest.NewRecorder()

	handleSwitchProject(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestHandleSwitchProject_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/switch-project", nil)
	w := httptest.NewRecorder()

	handleSwitchProject(w, req)

	assert.Equal(t, 405, w.Code)
}

func TestHandleSwitchProject_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/switch-project", strings.NewReader("not-json"))
	w := httptest.NewRecorder()

	handleSwitchProject(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestHandleProjects(t *testing.T) {
	cleanup := setupTestProjectSwitch(t)
	defer cleanup()

	config = Config{
		Projects: []Project{
			{Name: "alpha", Path: "/a"},
			{Name: "beta", Path: "/b"},
		},
	}
	projectMutex.Lock()
	currentProjectIndex = 1
	projectMutex.Unlock()

	req := httptest.NewRequest("GET", "/api/projects", nil)
	w := httptest.NewRecorder()

	handleProjects(w, req)

	assert.Equal(t, 200, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, float64(1), resp["currentIndex"])
}

func TestHandleProjects_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/projects", nil)
	w := httptest.NewRecorder()

	handleProjects(w, req)

	assert.Equal(t, 405, w.Code)
}
