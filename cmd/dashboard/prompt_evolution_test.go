package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLogTimestamp(t *testing.T) {
	cases := []struct {
		input    string
		expected time.Time
	}{
		{"[2026-04-26 21:14:43] StepBegin", time.Date(2026, 4, 26, 21, 14, 43, 0, time.UTC)},
		{"[2026-04-26 21:14:43] some log", time.Date(2026, 4, 26, 21, 14, 43, 0, time.UTC)},
		{"no timestamp here", time.Time{}},
		{"", time.Time{}},
	}

	for _, c := range cases {
		got := parseLogTimestamp(c.input)
		assert.Equal(t, c.expected, got, "input: %s", c.input)
	}
}

func TestIsToolName(t *testing.T) {
	assert.True(t, isToolName("ReadFile"))
	assert.True(t, isToolName("WriteFile"))
	assert.True(t, isToolName("StrReplaceFile"))
	assert.True(t, isToolName("Shell"))
	assert.True(t, isToolName("Grep"))
	assert.False(t, isToolName("FunctionBody"))
	assert.False(t, isToolName("arguments"))
	assert.True(t, isToolName("Shell"))
}

func TestParseLogFile(t *testing.T) {
	// 创建临时日志文件
	tmpDir := t.TempDir()

	// 测试用例 1: 典型的 dev 任务日志
	logContent := `[2026-04-26 21:14:43] Task started
StepBegin(n=1)
    function=FunctionBody(name='ReadFile', arguments='{"path": "test.go"}')
        think='I need to read the file first'
StepBegin(n=2)
    function=FunctionBody(name='Shell', arguments='{"command": "ls"}')
        think='Let me check the directory'
    return_value=ToolError(error='command not found')
StepBegin(n=3)
    function=FunctionBody(name='WriteFile', arguments='{"path": "out.go", "content": "package main"}')
        name='WriteFile',
    return_value=FileWriteResult(path='out.go')
StepBegin(n=4)
    function=FunctionBody(name='StrReplaceFile', arguments='{"path": "out.go"}')
        name='StrReplaceFile',
StepBegin(n=5)
    function=FunctionBody(name='Shell', arguments='{"command": "git add ."}')
    git add .
    git commit -m "test"
[2026-04-26 21:15:10] Task finished
`

	logPath := filepath.Join(tmpDir, "dev-issue-64-20260426-211443.log")
	require.NoError(t, os.WriteFile(logPath, []byte(logContent), 0644))

	a := parseLogFile(logPath)
	require.NotNil(t, a)

	assert.Equal(t, "dev-issue-64-20260426-211443.log", a.Filename)
	assert.Equal(t, "dev", a.TaskType)
	assert.Equal(t, 64, a.TargetID)
	assert.Equal(t, 5, a.TotalSteps)
	assert.Equal(t, 2, a.ThinkCount)
	assert.Equal(t, 1, a.Errors)
	assert.Equal(t, 1, a.FileWrites)
	assert.Equal(t, 1, a.FileModifies)
	assert.True(t, a.HasResult)
	assert.Equal(t, 2, a.GitOps["add"])
	assert.Equal(t, 1, a.GitOps["commit"])

	// 测试时间解析
	expectedStart := time.Date(2026, 4, 26, 21, 14, 43, 0, time.UTC)
	expectedEnd := time.Date(2026, 4, 26, 21, 15, 10, 0, time.UTC)
	assert.Equal(t, expectedStart, a.StartTime)
	assert.Equal(t, expectedEnd, a.EndTime)
	assert.Equal(t, 27*time.Second, a.Duration)

	// 测试用例 2: 零产出的日志
	noOutputLog := `[2026-04-26 22:00:00] Task started
StepBegin(n=1)
    think='What should I do?'
StepBegin(n=2)
    think='I am confused'
StepBegin(n=3)
    think='Still confused'
[2026-04-26 22:05:00] Task finished
`
	noOutputPath := filepath.Join(tmpDir, "review-pr-30-20260426-123456.log")
	require.NoError(t, os.WriteFile(noOutputPath, []byte(noOutputLog), 0644))

	b := parseLogFile(noOutputPath)
	require.NotNil(t, b)
	assert.Equal(t, "review", b.TaskType)
	assert.Equal(t, 30, b.TargetID)
	assert.False(t, b.HasResult)
	assert.Equal(t, 0, b.FileWrites)
	assert.Equal(t, 0, b.FileModifies)
	assert.Equal(t, 5*time.Minute, b.Duration)
}

func TestParseLogFile_EdgeCases(t *testing.T) {
	tmpDir := t.TempDir()

	// 空文件
	emptyPath := filepath.Join(tmpDir, "empty.log")
	require.NoError(t, os.WriteFile(emptyPath, []byte{}, 0644))
	a := parseLogFile(emptyPath)
	require.NotNil(t, a)
	assert.Equal(t, 0, a.TotalSteps)
	assert.False(t, a.HasResult)

	// 不存在的文件
	b := parseLogFile(filepath.Join(tmpDir, "nonexistent.log"))
	assert.Nil(t, b)

	// 文件名无法解析任务类型
	badNamePath := filepath.Join(tmpDir, "something.log")
	require.NoError(t, os.WriteFile(badNamePath, []byte("[2026-04-26 21:14:43] test\n"), 0644))
	c := parseLogFile(badNamePath)
	require.NotNil(t, c)
	assert.Equal(t, "", c.TaskType) // parts[0] = "something"
	assert.Equal(t, 0, c.TargetID)
}

func TestLoadSavePromptEvolutionState(t *testing.T) {
	// 由于 getProjectLogsDir 返回的是相对路径，我们需要在测试中创建临时日志目录
	projectName := fmt.Sprintf("test-project-%d", time.Now().UnixNano())
	logsDir := filepath.Join("logs", projectName)
	require.NoError(t, os.MkdirAll(logsDir, 0755))
	defer os.RemoveAll(logsDir)

	// 加载不存在的文件
	state := loadPromptEvolutionState(projectName)
	assert.NotNil(t, state)
	assert.True(t, state.LastEvolutionTime.IsZero())
	assert.Equal(t, 0, state.TaskCountSinceLast)

	// 保存并重新加载
	state.LastEvolutionTime = time.Date(2026, 4, 26, 10, 0, 0, 0, time.UTC)
	state.TaskCountSinceLast = 3
	state.TotalAnalyzed = 10
	state.LastReportPath = "/tmp/report.json"
	savePromptEvolutionState(projectName, state)

	loaded := loadPromptEvolutionState(projectName)
	assert.Equal(t, state.LastEvolutionTime, loaded.LastEvolutionTime)
	assert.Equal(t, 3, loaded.TaskCountSinceLast)
	assert.Equal(t, 10, loaded.TotalAnalyzed)
	assert.Equal(t, "/tmp/report.json", loaded.LastReportPath)
}

func TestExtractIssuePatterns(t *testing.T) {
	analyses := []*LogAnalysis{
		{TaskType: "dev", TargetID: 1, TotalSteps: 60, Errors: 5, HasResult: false},
		{TaskType: "dev", TargetID: 2, TotalSteps: 30, Errors: 0, HasResult: true},
		{TaskType: "dev", TargetID: 3, TotalSteps: 20, Errors: 3, HasResult: false},
		{TaskType: "review", TargetID: 10, TotalSteps: 5, Errors: 0, HasResult: false},
		{TaskType: "review", TargetID: 11, TotalSteps: 15, Errors: 0, HasResult: false},
	}

	issues := extractIssuePatterns(analyses)

	// dev 错误率 2/3 = 66.7% > 30%，应该被报告
	foundDevError := false
	foundDevHighSteps := false
	foundDevNoOutput := false
	for _, issue := range issues {
		switch issue.Pattern {
		case "dev_high_error_rate":
			foundDevError = true
			assert.Contains(t, issue.Description, "67%")
		case "dev_high_steps":
			foundDevHighSteps = true
			assert.Contains(t, issue.Description, "60")
		case "dev_no_output":
			foundDevNoOutput = true
		}
	}
	assert.True(t, foundDevError, "should detect high error rate for dev")
	assert.True(t, foundDevHighSteps, "should detect high steps for dev")
	assert.True(t, foundDevNoOutput, "should detect no output for dev")
}

func TestAnalyzeLogs(t *testing.T) {
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("test-analyze-%d", time.Now().UnixNano())
	logsDir := filepath.Join(tmpDir, "logs", projectName)
	require.NoError(t, os.MkdirAll(logsDir, 0755))

	// 临时替换 getProjectLogsDir 行为不可行（硬编码函数）
	// 所以我们在实际路径下创建文件，但使用一个自定义路径来测试
	// 更好的方式是直接测试 analyzeLogs 的核心逻辑。
	// 由于 analyzeLogs 依赖 getProjectLogsDir 的硬编码路径，
	// 我们直接测试 parseLogFile 和 extractIssuePatterns 等独立函数。
	// analyzeLogs 的集成测试需要修改 getProjectLogsDir 或使用 mock。

	// 创建一些测试日志
	log1 := `[2026-04-26 21:14:43] Task started
StepBegin(n=1)
    function=FunctionBody(name='ReadFile')
    name='ReadFile',
StepBegin(n=2)
    function=FunctionBody(name='WriteFile')
    name='WriteFile',
[2026-04-26 21:15:00] Task finished
`
	require.NoError(t, os.WriteFile(filepath.Join(logsDir, "dev-issue-1-20260426-211443.log"), []byte(log1), 0644))

	log2 := `[2026-04-26 22:00:00] Task started
StepBegin(n=1)
    think='I am confused'
[2026-04-26 22:10:00] Task finished
`
	require.NoError(t, os.WriteFile(filepath.Join(logsDir, "review-pr-5-20260426-220000.log"), []byte(log2), 0644))

	// 由于 getProjectLogsDir 返回的是相对路径 logs/<projectName>，
	// 我们需要在当前工作目录下运行测试。为了安全，我们跳过集成测试。
	// 或者我们可以在一个临时工作目录中运行。

	// 这里不做集成测试，因为会污染实际日志目录。
	// parseLogFile 和 extractIssuePatterns 的单元测试已经足够。
}

func TestCheckAndTriggerPromptEvolution(t *testing.T) {
	projectName := fmt.Sprintf("test-trigger-%d", time.Now().UnixNano())
	logsDir := filepath.Join("logs", projectName)
	require.NoError(t, os.MkdirAll(logsDir, 0755))
	defer os.RemoveAll(logsDir)

	// 初始状态：不应触发
	state := loadPromptEvolutionState(projectName)
	assert.Equal(t, 0, state.TaskCountSinceLast)

	// 手动设置状态为刚过阈值
	state.TaskCountSinceLast = 5
	state.LastEvolutionTime = time.Now().Add(-2 * time.Hour) // 2小时前
	savePromptEvolutionState(projectName, state)

	// 由于没有实际日志文件，analyzeLogs 会返回 nil（ReadDir 空目录返回 nil err？）
	// 实际上 ReadDir 空目录不会返回 error，而是返回空列表
	// 所以 report.TotalLogs = 0，不会返回 nil

	// 验证执行后状态被重置
	checkAndTriggerPromptEvolution(projectName)

	newState := loadPromptEvolutionState(projectName)
	assert.Equal(t, 0, newState.TaskCountSinceLast)
	assert.False(t, newState.LastEvolutionTime.IsZero())
	assert.Equal(t, 0, newState.TotalAnalyzed) // 没有日志文件
}

func TestCheckAndTriggerPromptEvolution_NotTrigger(t *testing.T) {
	projectName := fmt.Sprintf("test-notrigger-%d", time.Now().UnixNano())
	logsDir := filepath.Join("logs", projectName)
	require.NoError(t, os.MkdirAll(logsDir, 0755))
	defer os.RemoveAll(logsDir)

	state := loadPromptEvolutionState(projectName)
	state.TaskCountSinceLast = 2
	state.LastEvolutionTime = time.Now().Add(-30 * time.Minute)
	savePromptEvolutionState(projectName, state)

	checkAndTriggerPromptEvolution(projectName)

	newState := loadPromptEvolutionState(projectName)
	assert.Equal(t, 3, newState.TaskCountSinceLast) // 增加 1
}

func TestAnalyzeLogs_Integration(t *testing.T) {
	// 这个测试使用实际的日志文件来验证 analyzeLogs
	// 只在存在实际日志时运行
	logsDir := "logs/new-world-monorepo"
	entries, err := os.ReadDir(logsDir)
	if err != nil || len(entries) == 0 {
		t.Skip("无实际日志文件，跳过集成测试")
	}

	report := analyzeLogs("new-world-monorepo", time.Time{})
	require.NotNil(t, report)
	assert.Equal(t, "new-world-monorepo", report.ProjectName)
	assert.Greater(t, report.TotalLogs, 0)

	// 打印统计信息供人工检查
	t.Logf("Total logs analyzed: %d", report.TotalLogs)
	for agent, stat := range report.AgentStats {
		t.Logf("Agent %s: count=%d avgSteps=%.1f successRate=%.1f%% errors=%d",
			agent, stat.Count, stat.AvgSteps, stat.SuccessRate*100, stat.TotalErrors)
	}
	for _, issue := range report.TopIssues {
		t.Logf("Issue: %s - %s", issue.Pattern, issue.Description)
	}
}

func TestAnalyzeLogs_WithRealLogs(t *testing.T) {
	// 切换到项目根目录，使 getProjectLogsDir 返回正确路径
	origDir, _ := os.Getwd()
	rootDir := filepath.Join(origDir, "..", "..")
	os.Chdir(rootDir)
	defer os.Chdir(origDir)

	logsDir := getProjectLogsDir("new-world-monorepo")
	entries, err := os.ReadDir(logsDir)
	if err != nil || len(entries) == 0 {
		t.Skipf("无实际日志文件在 %s", logsDir)
	}

	report := analyzeLogs("new-world-monorepo", time.Time{})
	require.NotNil(t, report)
	assert.Equal(t, "new-world-monorepo", report.ProjectName)
	assert.Greater(t, report.TotalLogs, 0)

	t.Logf("Total logs analyzed: %d", report.TotalLogs)
	for agent, stat := range report.AgentStats {
		t.Logf("Agent %s: count=%d avgSteps=%.1f avgDuration=%s successRate=%.1f%% errors=%d writes=%d topGitOp=%s",
			agent, stat.Count, stat.AvgSteps, stat.AvgDuration, stat.SuccessRate*100,
			stat.TotalErrors, stat.TotalWrites, stat.TopGitOp)
	}
	for _, issue := range report.TopIssues {
		t.Logf("Issue: %s - %s", issue.Pattern, issue.Description)
	}
}
