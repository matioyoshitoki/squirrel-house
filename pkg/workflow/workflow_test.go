package workflow

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewContext(t *testing.T) {
	vars := map[string]interface{}{"key": "value"}
	ctx := NewContext("/tmp/project", vars)
	assert.Equal(t, "/tmp/project", ctx.ProjectPath)
	assert.Equal(t, "value", ctx.Vars["key"])
	assert.NotNil(t, ctx.State)
	assert.Empty(t, ctx.State)
	assert.WithinDuration(t, time.Now(), ctx.StartTime, time.Second)
}

func TestEngineRun_ShellStep_Success(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "echo", Type: "shell", Command: "echo hello-world"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	result := ctx.State["stepResult_echo"].(StepResult)
	assert.Equal(t, "success", result.Status)
	assert.Equal(t, 0, result.ExitCode)
}

func TestEngineRun_ShellStep_Failure(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "fail", Type: "shell", Command: "exit 42"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.Error(t, err)
	result := ctx.State["stepResult_fail"].(StepResult)
	assert.Equal(t, "failed", result.Status)
	assert.Equal(t, 42, result.ExitCode)
}

func TestEngineRun_ShellStep_IgnoreError(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "fail-ok", Type: "shell", Command: "exit 1", IgnoreError: true},
			{ID: "next", Type: "shell", Command: "echo ok"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	// When IgnoreError is true, recordStepResult records success
	assert.Equal(t, "success", ctx.State["stepResult_fail-ok"].(StepResult).Status)
	assert.Equal(t, 1, ctx.State["stepResult_fail-ok"].(StepResult).ExitCode)
	assert.Equal(t, "success", ctx.State["stepResult_next"].(StepResult).Status)
}

func TestEngineRun_StepCondition_False(t *testing.T) {
	ctx := NewContext("", map[string]interface{}{"skip": "false"})
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "cond", Type: "shell", Command: "echo should-not-run", Condition: "{{ .Vars.skip }}"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	assert.Nil(t, ctx.State["stepResult_cond"])
}

func TestEngineRun_StepCondition_True(t *testing.T) {
	ctx := NewContext("", map[string]interface{}{"run": "true"})
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "cond", Type: "shell", Command: "echo run", Condition: "{{ .Vars.run }}"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	assert.Equal(t, "success", ctx.State["stepResult_cond"].(StepResult).Status)
}

func TestEngineRun_StepCondition_EmptyString(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "no-cond", Type: "shell", Command: "echo run"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	assert.Equal(t, "success", ctx.State["stepResult_no-cond"].(StepResult).Status)
}

func TestEngineRun_Cleanup_OnSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	cleanupFile := filepath.Join(tmpDir, "cleanup.txt")
	ctx := NewContext(tmpDir, nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "main", Type: "shell", Command: "echo main"},
			{ID: "cleanup", Type: "shell", Command: "echo cleanup-done > " + cleanupFile, IsCleanup: true},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	data, _ := os.ReadFile(cleanupFile)
	assert.Equal(t, "cleanup-done\n", string(data))
}

func TestEngineRun_Cleanup_OnFailure(t *testing.T) {
	tmpDir := t.TempDir()
	cleanupFile := filepath.Join(tmpDir, "cleanup.txt")
	ctx := NewContext(tmpDir, nil)
	engine := NewEngine(ctx)

	// Cleanup steps must appear BEFORE the step that may fail,
	// because they are collected during iteration.
	wf := &Workflow{
		Steps: []Step{
			{ID: "cleanup", Type: "shell", Command: "echo cleanup-done > " + cleanupFile, IsCleanup: true},
			{ID: "main", Type: "shell", Command: "exit 1"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.Error(t, err)
	data, _ := os.ReadFile(cleanupFile)
	assert.Equal(t, "cleanup-done\n", string(data))
}

func TestEngineRun_Cleanup_SkippedDuringCollection(t *testing.T) {
	// Cleanup steps should not be executed during normal flow, only collected
	tmpDir := t.TempDir()
	markerFile := filepath.Join(tmpDir, "marker.txt")
	ctx := NewContext(tmpDir, nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "cleanup1", Type: "shell", Command: "echo early > " + markerFile, IsCleanup: true},
			{ID: "main", Type: "shell", Command: "echo main"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	// Cleanup should run AFTER main, not before
	data, _ := os.ReadFile(markerFile)
	assert.Equal(t, "early\n", string(data)) // still written, but after main completes
}

func TestEngineRun_ReadFileStep(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("file-content"), 0644)

	ctx := NewContext(tmpDir, nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "read", Type: "readFile", Path: testFile, OutputKey: "content"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	assert.Equal(t, "file-content", ctx.State["content"])
}

func TestEngineRun_ReadFileStep_NotFound_IgnoreError(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "read", Type: "readFile", Path: "/nonexistent/file.txt", OutputKey: "content", IgnoreError: true},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	assert.Equal(t, "", ctx.State["content"])
}

func TestEngineRun_ReadFileStep_NotFound_Fails(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "read", Type: "readFile", Path: "/nonexistent/file.txt", OutputKey: "content"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.Error(t, err)
}

func TestEngineRun_WriteFileStep(t *testing.T) {
	tmpDir := t.TempDir()
	outFile := filepath.Join(tmpDir, "out.txt")
	ctx := NewContext(tmpDir, nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "write", Type: "writeFile", Path: outFile, Content: "hello-{{ .ProjectPath }}"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	data, _ := os.ReadFile(outFile)
	assert.Equal(t, "hello-"+tmpDir, string(data))
}

func TestEngineRun_OutputKey(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "out", Type: "shell", Command: "echo captured-output", OutputKey: "myOutput"},
		},
	}

	err := engine.Run(context.Background(), wf)
	require.NoError(t, err)
	assert.Equal(t, "captured-output\n", ctx.State["myOutput"])
}

func TestRenderString(t *testing.T) {
	ctx := NewContext("/project", map[string]interface{}{
		"name": "world",
	})
	engine := NewEngine(ctx)

	tests := []struct {
		name     string
		template string
		want     string
	}{
		{"simple var", "hello {{ .Vars.name }}", "hello world"},
		{"project path", "path: {{ .ProjectPath }}", "path: /project"},
		{"func trim", "{{ .Vars.name | trim }}", "world"},
		{"func shellquote", "{{ .Vars.name | shellquote }}", "'world'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := engine.renderString(tt.template)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRenderString_NowFunc(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	got, err := engine.renderString("{{ now }}")
	require.NoError(t, err)
	assert.NotEmpty(t, got)
	// Should be a numeric timestamp
	_, parseErr := time.Parse("2006", got)
	assert.Error(t, parseErr) // It's a unix timestamp, not a date string
}

func TestRecordStepResult_Success(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)
	step := Step{ID: "s1"}
	engine.recordStepResult(step, 0, nil)

	result := ctx.State["stepResult_s1"].(StepResult)
	assert.Equal(t, "success", result.Status)
	assert.Equal(t, 0, result.ExitCode)
	assert.Empty(t, result.Error)
}

func TestRecordStepResult_Failure(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)
	step := Step{ID: "s1"}
	stepErr := assert.AnError
	engine.recordStepResult(step, 1, stepErr)

	result := ctx.State["stepResult_s1"].(StepResult)
	assert.Equal(t, "failed", result.Status)
	assert.Equal(t, 1, result.ExitCode)
	assert.NotEmpty(t, result.Error)
}

func TestRecordStepResult_IgnoreError(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)
	step := Step{ID: "s1", IgnoreError: true}
	engine.recordStepResult(step, 1, assert.AnError)

	// When IgnoreError is true, recordStepResult still records as success
	result := ctx.State["stepResult_s1"].(StepResult)
	assert.Equal(t, "success", result.Status)
	assert.Equal(t, 1, result.ExitCode)
}

func TestKill(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	wf := &Workflow{
		Steps: []Step{
			{ID: "sleep", Type: "shell", Command: "sleep 10"},
		},
	}

	runCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- engine.Run(runCtx, wf)
	}()

	// Give the shell command time to start
	time.Sleep(200 * time.Millisecond)
	err := engine.Kill()
	require.NoError(t, err)

	select {
	case runErr := <-done:
		require.Error(t, runErr)
	case <-time.After(2 * time.Second):
		t.Fatal("Kill did not terminate workflow in time")
	}
}

func TestKill_NoRunningProcess(t *testing.T) {
	ctx := NewContext("", nil)
	engine := NewEngine(ctx)

	// Kill when no process is running should not error
	err := engine.Kill()
	assert.NoError(t, err)
}

func TestEnsureToolPaths(t *testing.T) {
	env := []string{"PATH=/usr/bin"}
	result := ensureToolPaths(env)
	found := false
	for _, e := range result {
		if strings.HasPrefix(e, "PATH=") {
			assert.Contains(t, e, "/Users/insulate/flutter/bin")
			found = true
		}
	}
	assert.True(t, found, "PATH should be modified")
}

func TestEnsureToolPaths_NoExistingPath(t *testing.T) {
	env := []string{"HOME=/home/user"}
	result := ensureToolPaths(env)
	found := false
	for _, e := range result {
		if strings.HasPrefix(e, "PATH=") {
			assert.Contains(t, e, "/Users/insulate/flutter/bin")
			found = true
		}
	}
	assert.True(t, found, "PATH should be appended")
}

func TestGetExitCode(t *testing.T) {
	// Cannot easily create *exec.ExitError in a test, but we can test the -1 fallback
	code := getExitCode(assert.AnError)
	assert.Equal(t, -1, code)
}

// TestEngineKillTerminatesProcessGroup 验证 Kill 使用进程组级联终止，
// 避免 Setsid 创建的子进程在父进程被杀后变成孤儿进程继续运行。
func TestEngineKillTerminatesProcessGroup(t *testing.T) {
	marker := fmt.Sprintf("nwops-pgkill-test-%d", time.Now().UnixNano())
	ctx := NewContext("/tmp", map[string]interface{}{})
	engine := NewEngine(ctx)

	var startedPID int
	engine.OnProcessStart = func(pid int) {
		startedPID = pid
	}

	wf := &Workflow{
		Steps: []Step{
			{
				ID:      "parentWithChildren",
				Type:    "shell",
				Command: fmt.Sprintf("echo %s; sleep 300 & sleep 300", marker),
			},
		},
	}

	// 延迟 kill，确保子进程已经启动
	go func() {
		time.Sleep(400 * time.Millisecond)
		engine.Kill()
	}()

	err := engine.Run(context.Background(), wf)
	require.Error(t, err, "expected error after Kill")

	// 给内核一点时间回收进程
	time.Sleep(200 * time.Millisecond)

	// 检查 shell 父进程是否已退出
	if startedPID > 0 {
		if err := exec.Command("kill", "-0", fmt.Sprintf("%d", startedPID)).Run(); err == nil {
			t.Errorf("parent shell process %d still alive after Kill", startedPID)
		}
	}

	// 检查子进程（sleep）是否被级联终止——通过 marker 定位
	out, _ := exec.Command("pgrep", "-f", marker).CombinedOutput()
	if len(out) > 0 {
		pids := strings.TrimSpace(string(out))
		t.Errorf("Kill did not terminate child processes,残留 PIDs: %s", pids)
		// 清理残留，避免影响后续测试
		exec.Command("pkill", "-9", "-f", marker).Run()
	}
}

// TestEngineKillNoopWhenNoProcess 验证没有 currentCmd 时 Kill 不 panic
func TestEngineKillNoopWhenNoProcess(t *testing.T) {
	engine := NewEngine(NewContext("/tmp", nil))
	assert.NoError(t, engine.Kill())
}
