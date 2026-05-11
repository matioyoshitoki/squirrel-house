package main

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/MatioYoshitoki/new-world-ops/pkg/workflow"
)

// ============================================================================
// extractReviewResult
// ============================================================================

func TestExtractReviewResult(t *testing.T) {
	tests := []struct {
		name   string
		report string
		want   string
	}{
		{"explicit PASS", "## 审查结论: PASS\n", "passed"},
		{"bold PASS", "审查结论：**PASS**", "passed"},
		{"colon bold PASS", "审查结论: **PASS**", "passed"},
		{"bold colon PASS", "**审查结论**: PASS", "passed"},
		{"bold both", "**审查结论**: **PASS**", "passed"},
		{"checkmark pass", "审查结论: ✅ 通过", "passed"},
		{"simple pass", "审查结论: 通过", "passed"},
		{"fullwidth colon pass", "审查结论：通过", "passed"},
		{"checkmark no space", "审查结论: ✅通过", "passed"},
		{"standalone checkmark", "✅ 通过", "passed"},
		{"standalone checkmark no space", "✅通过", "passed"},
		{"LGTM uppercase", "LGTM", "passed"},
		{"lgtm lowercase", "lgtm", "passed"},
		{"no match defaults to needs_fix", "some random content", "needs_fix"},
		{"empty report", "", "needs_fix"},
		{"NEEDS_FIX in content", "Result: NEEDS_FIX", "needs_fix"},
		{"NEEDS_MAJOR_FIX in content", "审查结论: NEEDS_MAJOR_FIX", "needs_fix"},
		{"NEEDS_MINOR_FIX in content", "审查结论: NEEDS_MINOR_FIX", "needs_fix"},
		{"reject with checkmark in body", "审查结论: NEEDS_MAJOR_FIX\n\n自检结果\n- State: ✅ 通过\n- 回归: ✅ 通过", "needs_fix"},
		{"reject 未通过", "审查结论: 未通过", "needs_fix"},
		{"reject 需修改", "❌ 需修改", "needs_fix"},
		{"mixed with PASS keyword in text", "This is not a PASS but a normal sentence.", "needs_fix"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractReviewResult(tt.report)
			if got != tt.want {
				t.Errorf("extractReviewResult(%q) = %q, want %q", tt.report, got, tt.want)
			}
		})
	}
}

// ============================================================================
// extractBlockingIssues
// ============================================================================

func TestExtractBlockingIssues(t *testing.T) {
	tests := []struct {
		name   string
		report string
		want   string
	}{
		{
			name: "blocking section",
			report: "必须修复\n- Issue 1\n- Issue 2\n建议优化\n- Issue 3",
			want:   "- Issue 1\n- Issue 2",
		},
		{
			name: "blocking header",
			report: "### 🔴 Blocking\n- Bug A\n- Bug B\n### Non-blocking",
			want:   "- Bug A\n- Bug B",
		},
		{
			name:   "no blocking section",
			report: "建议优化\n- Issue 1",
			want:   "请根据审查报告中的问题进行修改。",
		},
		{
			name:   "empty report",
			report: "",
			want:   "请根据审查报告中的问题进行修改。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractBlockingIssues(tt.report)
			if got != tt.want {
				t.Errorf("extractBlockingIssues mismatch:\ngot:\n%s\nwant:\n%s", got, tt.want)
			}
		})
	}
}

// ============================================================================
// parseRelatedIssue
// ============================================================================

func TestParseRelatedIssue(t *testing.T) {
	tests := []struct {
		name string
		body string
		want int
	}{
		{"Fixes #42", "Fixes #42", 42},
		{"Closes #123", "Closes #123", 123},
		{"Resolves #999", "Resolves #999", 999},
		{"refs #7", "refs #7", 7},
		{"References #88", "References #88", 88},
		{"no issue", "just some text", 0},
		{"empty", "", 0},
		{"multiple matches first wins", "Closes #123 and fixes #456", 123},
		{"case insensitive", "fixes #77", 77},
		{"number in text", "Issue #42 is fixed", 0}, // only matches specific keywords
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseRelatedIssue(tt.body)
			if got != tt.want {
				t.Errorf("parseRelatedIssue(%q) = %d, want %d", tt.body, got, tt.want)
			}
		})
	}
}

// ============================================================================
// hasMergeConflict
// ============================================================================

func TestHasMergeConflict(t *testing.T) {
	tests := []struct {
		name   string
		output string
		want   bool
	}{
		{"conflict keyword", "error: conflict detected", true},
		{"has conflicts", "PR has conflicts", true},
		{"merge conflict", "merge conflict in file.go", true},
		{"dirty", "working tree dirty", true},
		{"out of date", "branch is out of date", true},
		{"no conflict", "everything is fine", false},
		{"empty", "", false},
		{"case insensitive", "CONFLICT", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasMergeConflict(tt.output)
			if got != tt.want {
				t.Errorf("hasMergeConflict(%q) = %v, want %v", tt.output, got, tt.want)
			}
		})
	}
}

// ============================================================================
// unescapeNewlines
// ============================================================================

func TestUnescapeNewlines(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"header", "n## Title", "\n## Title"},
		{"subheader", "n### Sub", "\n### Sub"},
		{"blockquote", "n> quote", "\n> quote"},
		{"table", "n| a | b |", "\n| a | b |"},
		{"bullet", "n- item", "\n- item"},
		{"numbered", "n1. item", "\n1. item"},
		{"bold", "n**text**", "\n**text**"},
		{"code block", "n```go", "\n```go"},
		{"double newline", "n\n", "\n\n"},
		{"no replacement", "normal text", "normal text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unescapeNewlines(tt.in)
			if got != tt.want {
				t.Errorf("unescapeNewlines(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// ============================================================================
// unwrapHardWraps
// ============================================================================

func TestUnwrapHardWraps(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "simple paragraph join",
			in:   "This is a line.\nAnd another line.",
			want: "This is a line. And another line.",
		},
		{
			name: "preserve code block",
			in:   "```\ncode line 1\ncode line 2\n```",
			want: "```\ncode line 1\ncode line 2\n```",
		},
		{
			name: "preserve heading",
			in:   "# Heading\nBody text.",
			want: "# Heading\nBody text.",
		},
		{
			name: "CJK no space join",
			in:   "这是一个长句\n被截断了",
			want: "这是一个长句被截断了",
		},
		{
			name: "table preservation",
			in:   "| a | b |\n| c | d |",
			want: "| a | b |\n| c | d |",
		},
		{
			name: "empty string",
			in:   "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unwrapHardWraps(tt.in)
			if got != tt.want {
				t.Errorf("unwrapHardWraps mismatch:\ngot:  %q\nwant: %q", got, tt.want)
			}
		})
	}
}

// ============================================================================
// buildReworkTestCommands
// ============================================================================

func TestBuildReworkTestCommands(t *testing.T) {
	tests := []struct {
		name string
		ts   TechStack
		want []string
	}{
		{
			name: "flutter mobile",
			ts:   TechStack{Mobile: "flutter"},
			want: []string{"flutter pub get", "flutter analyze", "flutter test"},
		},
		{
			name: "nestjs backend",
			ts:   TechStack{Backend: "nestjs"},
			want: []string{"npm test", "npm run build"},
		},
		{
			name: "go backend",
			ts:   TechStack{Backend: "go"},
			want: []string{"go test ./...", "go vet ./..."},
		},
		{
			name: "phaser3 frontend",
			ts:   TechStack{Frontend: "phaser3"},
			want: []string{"npm run lint", "npm run test", "npm run build"},
		},
		{
			name: "unknown stack",
			ts:   TechStack{},
			want: []string{"make check-docs"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildReworkTestCommands(tt.ts)
			for _, w := range tt.want {
				if !strings.Contains(got, w) {
					t.Errorf("buildReworkTestCommands() missing %q in output:\n%s", w, got)
				}
			}
		})
	}
}

// ============================================================================
// getResultEmoji / toInt
// ============================================================================

func TestGetResultEmoji(t *testing.T) {
	if getResultEmoji("passed") != "✅ 通过" {
		t.Error("getResultEmoji(passed) mismatch")
	}
	if getResultEmoji("needs_fix") != "❌ 需修改" {
		t.Error("getResultEmoji(needs_fix) mismatch")
	}
	if getResultEmoji("unknown") != "❌ 需修改" {
		t.Error("getResultEmoji(unknown) mismatch")
	}
}

func TestToInt(t *testing.T) {
	if toInt("42") != 42 {
		t.Error("toInt(42) mismatch")
	}
	if toInt("abc") != 0 {
		t.Error("toInt(abc) should be 0")
	}
	if toInt("") != 0 {
		t.Error("toInt(empty) should be 0")
	}
}

// ============================================================================
// taskID
// ============================================================================

func TestTaskID(t *testing.T) {
	if got := taskID(TaskTypeReview, 58); got != "review-58" {
		t.Errorf("taskID(review, 58) = %q, want review-58", got)
	}
	if got := taskID(TaskTypeDev, 5); got != "dev-5" {
		t.Errorf("taskID(dev, 5) = %q, want dev-5", got)
	}
}

// ============================================================================
// renderHookTemplate
// ============================================================================

func TestRenderHookTemplate(t *testing.T) {
	wfCtx := workflow.NewContext("/tmp/project", map[string]interface{}{
		"branch": "feat/issue-5",
	})
	wfCtx.Vars["issueNumber"] = "42"

	tests := []struct {
		tmplStr string
		want    string
		wantErr bool
	}{
		{"{{ .Vars.issueNumber }}", "42", false},
		{"{{ .Vars.branch }}", "feat/issue-5", false},
		{"{{ .Vars.missing }}", "<no value>", false},
		{"{{ .ProjectPath }}", "/tmp/project", false},
		{"invalid {{ .Bad", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.tmplStr, func(t *testing.T) {
			got, err := renderHookTemplate(tt.tmplStr, wfCtx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("renderHookTemplate error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("renderHookTemplate(%q) = %q, want %q", tt.tmplStr, got, tt.want)
			}
		})
	}
}

// ============================================================================
// mapKeys
// ============================================================================

func TestMapKeys(t *testing.T) {
	m := map[string]interface{}{"a": 1, "b": 2, "c": 3}
	keys := mapKeys(m)
	if len(keys) != 3 {
		t.Fatalf("mapKeys length = %d, want 3", len(keys))
	}
	for _, k := range []string{"a", "b", "c"} {
		found := false
		for _, got := range keys {
			if got == k {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("mapKeys missing %q", k)
		}
	}
}

// ============================================================================
// inferStatusFromLog
// ============================================================================

func TestInferStatusFromLog(t *testing.T) {
	now := time.Now()
	old := now.Add(-time.Hour)

	tests := []struct {
		name   string
		content string
		modTime time.Time
		want   string
	}{
		{"success marker", "✅ 任务完成", now, "success"},
		{"review success", "✅ Review 完成", now, "success"},
		{"stopped marker", "⏹️ 任务已手动停止", now, "stopped"},
		{"interrupted marker", "⏸️ 任务因长时间无响应被标记为中断", now, "interrupted"},
		{"failed marker", "\n❌ 失败: something", now, "failed"},
		{"max steps", "Max number of steps reached", now, "failed"},
		{"步骤限制", "步骤限制 exceeded", now, "failed"},
		{"no marker old", strings.Repeat("some log content that is long enough to exceed the threshold of one hundred characters for infer status logic to work properly. ", 2), old, "interrupted"},
		{"no marker recent", strings.Repeat("some log content that is long enough to exceed the threshold of one hundred characters for infer status logic to work properly. ", 2), now, "running"},
		{"no marker short", "hi", old, "unknown"},
		{"last marker wins", "✅ 任务完成\n⏹️ 任务已手动停止", now, "stopped"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferStatusFromLog(tt.content, tt.modTime)
			if got != tt.want {
				t.Errorf("inferStatusFromLog() = %q, want %q", got, tt.want)
			}
		})
	}
}

// ============================================================================
// extractBranchFromLog / extractIssueNumberFromLog
// ============================================================================

func TestExtractBranchFromLog(t *testing.T) {
	if got := extractBranchFromLog("分支: feat/issue-5\n其他"); got != "feat/issue-5" {
		t.Errorf("extractBranchFromLog = %q, want feat/issue-5", got)
	}
	if got := extractBranchFromLog("no branch"); got != "unknown" {
		t.Errorf("extractBranchFromLog(no branch) = %q, want unknown", got)
	}
}

func TestExtractIssueNumberFromLog(t *testing.T) {
	if got := extractIssueNumberFromLog("Issue: #42\n"); got != 42 {
		t.Errorf("extractIssueNumberFromLog = %d, want 42", got)
	}
	if got := extractIssueNumberFromLog("no issue"); got != 0 {
		t.Errorf("extractIssueNumberFromLog(no issue) = %d, want 0", got)
	}
}

// ============================================================================
// extractIDFromFilename / matchIssueNumberInFilename
// ============================================================================

func TestExtractIDFromFilename(t *testing.T) {
	if got := extractIDFromFilename("dev-issue-5-20260425.log", "dev-issue-", ".log"); got != "5" {
		t.Errorf("extractIDFromFilename = %q, want 5", got)
	}
	if got := extractIDFromFilename("review-pr-58-20260425.log", "review-pr-", ".log"); got != "58" {
		t.Errorf("extractIDFromFilename = %q, want 58", got)
	}
}

func TestMatchIssueNumberInFilename(t *testing.T) {
	if !matchIssueNumberInFilename("issue-5.md", 5) {
			t.Error("matchIssueNumberInFilename(issue-5.md, 5) should be true")
	}
	if matchIssueNumberInFilename("issue-15.md", 5) {
		t.Error("matchIssueNumberInFilename(issue-15.md, 5) should be false")
	}
	if !matchIssueNumberInFilename("PLAN-42-v1.md", 42) {
		t.Error("matchIssueNumberInFilename(PLAN-42-v1.md, 42) should be true")
	}
}

// ============================================================================
// classifyIssue
// ============================================================================

func TestClassifyIssue(t *testing.T) {
	tests := []struct {
		name   string
		issue  Issue
		want   string
	}{
		{
			name: "frontend only",
			issue: Issue{
				Title:  "Add login screen",
				Labels: []Label{{Name: "frontend"}},
			},
			want: "frontend-only",
		},
		{
			name: "backend only",
			issue: Issue{
				Title:  "Fix API auth",
				Labels: []Label{{Name: "backend"}},
			},
			want: "backend-only",
		},
		{
			name: "full stack",
			issue: Issue{
				Title:  "Implement chat feature",
				Labels: []Label{{Name: "frontend"}, {Name: "backend"}},
			},
			want: "full-stack",
		},
		{
			name: "no labels",
			issue: Issue{
				Title: "Update README",
			},
			want: "full-stack",
		},
		{
			name: "flutter label",
			issue: Issue{
				Title:  "Build flutter widget",
				Labels: []Label{{Name: "flutter"}},
			},
			want: "frontend-only",
		},
		{
			name: "nestjs label",
			issue: Issue{
				Title:  "Setup NestJS module",
				Labels: []Label{{Name: "nestjs"}},
			},
			want: "backend-only",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := classifyIssue(tt.issue)
			if got != tt.want {
				t.Errorf("classifyIssue() = %q, want %q", got, tt.want)
			}
		})
	}
}

// ============================================================================
// countReviewFailuresFromLogs
// ============================================================================

func TestCountReviewFailuresFromLogs(t *testing.T) {
	// 通过临时目录 + 覆盖 getProjectLogsDir 行为来测试
	// 由于 getProjectLogsDir 是硬编码的，我们在测试中使用实际文件系统
	// 但放在 t.TempDir() 下，通过 monkey-patch 的替代方式：
	// 这里我们直接测试扫描逻辑，构造一个 mock 场景

	// 实际上 countReviewFailuresFromLogs 调用 getProjectLogsDir(projectName)
	// getProjectLogsDir 返回 filepath.Abs(filepath.Join("logs", projectName))
	// 这在测试中不可控。为了测试它，我们需要重构或接受无法单元测试的现实。
	// 目前先标记为需要接口化。
	
	// 创建一个临时 logs 目录结构来测试
	tmpDir := t.TempDir()
	logsDir := filepath.Join(tmpDir, "test-project")
	_ = logsDir // will use once we refactor

	// TODO: after extracting FileSystem interface, add real tests here
	t.Skip("countReviewFailuresFromLogs depends on getProjectLogsDir hardcoded path; needs FileSystem interface extraction")
}
