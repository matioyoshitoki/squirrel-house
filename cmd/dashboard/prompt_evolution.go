package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// PromptEvolutionState 记录 prompt 演进的状态
type PromptEvolutionState struct {
	LastEvolutionTime    time.Time `json:"lastEvolutionTime"`
	TaskCountSinceLast   int       `json:"taskCountSinceLast"`
	TotalAnalyzed        int       `json:"totalAnalyzed"`
	LastReportPath       string    `json:"lastReportPath,omitempty"`
}

// LogAnalysis 单个日志文件的分析结果
type LogAnalysis struct {
	Filename       string            `json:"filename"`
	TaskType       string            `json:"taskType"`
	TargetID       int               `json:"targetID"`
	StartTime      time.Time         `json:"startTime"`
	EndTime        time.Time         `json:"endTime"`
	Duration       time.Duration     `json:"duration"`
	TotalSteps     int               `json:"totalSteps"`
	ToolCalls      map[string]int    `json:"toolCalls"`
	ThinkCount     int               `json:"thinkCount"`
	ThinkTotalLen  int               `json:"thinkTotalLen"`
	GitOps         map[string]int    `json:"gitOps"`
	Errors         int               `json:"errors"`
	FileWrites     int               `json:"fileWrites"`
	FileModifies   int               `json:"fileModifies"`
	HasResult      bool              `json:"hasResult"` // 是否有实际产出
}

// PromptEvolutionReport 统计报告
type PromptEvolutionReport struct {
	GeneratedAt   time.Time       `json:"generatedAt"`
	ProjectName   string          `json:"projectName"`
	Since         time.Time       `json:"since"`
	TotalLogs     int             `json:"totalLogs"`
	AgentStats    map[string]*AgentStat `json:"agentStats"` // key: agent type
	TopIssues     []IssuePattern  `json:"topIssues"`
	Trends        *TrendComparison `json:"trends,omitempty"`
}

// AgentStat 按 agent 类型聚合的统计
type AgentStat struct {
	Count          int               `json:"count"`
	AvgSteps       float64           `json:"avgSteps"`
	AvgDuration    time.Duration     `json:"avgDuration"`
	TotalToolCalls map[string]int    `json:"totalToolCalls"`
	TotalErrors    int               `json:"totalErrors"`
	TotalWrites    int               `json:"totalWrites"`
	SuccessRate    float64           `json:"successRate"` // hasResult / count
	TopGitOp       string            `json:"topGitOp"`
}

// IssuePattern 常见失败模式
type IssuePattern struct {
	Pattern     string `json:"pattern"`
	Count       int    `json:"count"`
	Description string `json:"description"`
	RootCause   string `json:"rootCause,omitempty"`
	Suggestion  string `json:"suggestion,omitempty"`
	Severity    string `json:"severity,omitempty"` // critical/warning/info
}

// TrendComparison 与上次对比的趋势
type TrendComparison struct {
	AvgStepsDelta     float64 `json:"avgStepsDelta"`
	AvgDurationDelta  float64 `json:"avgDurationDelta"` // seconds
	ErrorRateDelta    float64 `json:"errorRateDelta"`
	SuccessRateDelta  float64 `json:"successRateDelta"`
}

// loadPromptEvolutionState 读取状态文件
func loadPromptEvolutionState(projectName string) *PromptEvolutionState {
	statePath := filepath.Join(getProjectLogsDir(projectName), "prompt-evolution-state.json")
	data, err := os.ReadFile(statePath)
	if err != nil {
		return &PromptEvolutionState{
			LastEvolutionTime:  time.Time{},
			TaskCountSinceLast: 0,
			TotalAnalyzed:      0,
		}
	}
	var state PromptEvolutionState
	if err := json.Unmarshal(data, &state); err != nil {
		return &PromptEvolutionState{}
	}
	return &state
}

// savePromptEvolutionState 保存状态文件
func savePromptEvolutionState(projectName string, state *PromptEvolutionState) {
	statePath := filepath.Join(getProjectLogsDir(projectName), "prompt-evolution-state.json")
	data, _ := json.MarshalIndent(state, "", "  ")
	os.WriteFile(statePath, data, 0644)
}

// parseLogFile 解析单个日志文件
func parseLogFile(path string) *LogAnalysis {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	content := string(data)
	lines := strings.Split(content, "\n")

	a := &LogAnalysis{
		Filename:     filepath.Base(path),
		ToolCalls:    make(map[string]int),
		GitOps:       make(map[string]int),
		StartTime:    time.Time{},
		EndTime:      time.Time{},
	}

	// 从文件名解析任务类型和目标ID
	// 格式: dev-issue-64-20260426-211443.log
	//       review-pr-30-20260426-123456.log
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, ".log")
	parts := strings.Split(base, "-")
	if len(parts) >= 3 {
		a.TaskType = parts[0]
		if id, err := strconv.Atoi(parts[2]); err == nil {
			a.TargetID = id
		}
	}

	// 解析时间戳（从首行和尾行）
	if len(lines) > 0 {
		a.StartTime = parseLogTimestamp(lines[0])
	}
	if len(lines) > 1 {
		// 从后往前找最后一个带时间戳的行
		for i := len(lines) - 1; i >= 0; i-- {
			if t := parseLogTimestamp(lines[i]); !t.IsZero() {
				a.EndTime = t
				break
			}
		}
	}
	if !a.StartTime.IsZero() && !a.EndTime.IsZero() {
		a.Duration = a.EndTime.Sub(a.StartTime)
	}

	// 正则表达式
	stepRe := regexp.MustCompile(`StepBegin\(n=(\d+)\)`)
	// 匹配缩进后的 name='...'（包括多行和单行 FunctionBody 格式）
	toolRe := regexp.MustCompile(`^\s{4,}name=['"](\w+)['"]`)
	thinkRe := regexp.MustCompile(`think='([^']+)'`)
	// 匹配 ToolError 以及 ToolReturnValue(is_error=True) 两种错误形态
	toolErrorRe := regexp.MustCompile(`(?m)^\s+is_error=True,$`)
	writeFileRe := regexp.MustCompile(`^\s{4,}name=['"]WriteFile['"]`)
	strReplaceRe := regexp.MustCompile(`^\s{4,}name=['"]StrReplaceFile['"]`)
	gitRe := regexp.MustCompile(`git\s+(\w+)`)
	// 匹配 git diff --stat 输出中的文件变更统计（如 "3 files changed, 87 insertions(+), 1 deletion(-)"）
	gitDiffStatRe := regexp.MustCompile(`(\d+)\s+files?\s+changed,\s+(\d+)\s+insertions?\(\+\)(?:,\s+(\d+)\s+deletions?\(-\))?`)

	maxStep := 0
	hasGitDiffChanges := false
	for _, line := range lines {
		// Step 计数
		if m := stepRe.FindStringSubmatch(line); m != nil {
			if n, err := strconv.Atoi(m[1]); err == nil && n > maxStep {
				maxStep = n
			}
		}

		// 工具调用
		if m := toolRe.FindStringSubmatch(line); m != nil {
			toolName := m[1]
			// 排除非工具名称的匹配
			if isToolName(toolName) {
				a.ToolCalls[toolName]++
			}
		}

		// Think
		if m := thinkRe.FindStringSubmatch(line); m != nil {
			a.ThinkCount++
			a.ThinkTotalLen += len(m[1])
		}

		// 错误
		if toolErrorRe.MatchString(line) {
			a.Errors++
		}

		// 文件写入
		if writeFileRe.MatchString(line) {
			a.FileWrites++
		}
		if strReplaceRe.MatchString(line) {
			a.FileModifies++
		}

		// Git 操作
		if m := gitRe.FindStringSubmatch(line); m != nil {
			// 排除 diff --git 输出中的误匹配（如 diff --git a/... b/...）
			if strings.Contains(line, "diff --git") || strings.Contains(line, "--git") {
				continue
			}
			a.GitOps[m[1]]++
		}

		// 检测 git diff --stat 中的实际文件修改（兜底：agent 可能通过 Shell 命令修改文件）
		if m := gitDiffStatRe.FindStringSubmatch(line); m != nil {
			insertions, _ := strconv.Atoi(m[2])
			deletions := 0
			if len(m) > 3 && m[3] != "" {
				deletions, _ = strconv.Atoi(m[3])
			}
			if insertions > 0 || deletions > 0 {
				hasGitDiffChanges = true
			}
		}
	}
	a.TotalSteps = maxStep
	a.HasResult = a.FileWrites > 0 || a.FileModifies > 0 || hasGitDiffChanges

	return a
}

// isToolName 判断是否为有效的工具名称
func isToolName(name string) bool {
	validTools := map[string]bool{
		"ReadFile": true, "WriteFile": true, "StrReplaceFile": true,
		"Shell": true, "Grep": true, "Glob": true,
		"SearchWeb": true, "FetchURL": true,
		"AskUserQuestion": true, "ReadMediaFile": true,
	}
	return validTools[name]
}

// parseLogTimestamp 从日志行解析时间戳
// 格式: [2026-04-26 21:14:43] ...
func parseLogTimestamp(line string) time.Time {
	re := regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2})\]`)
	if m := re.FindStringSubmatch(line); m != nil {
		if t, err := time.Parse("2006-01-02 15:04:05", m[1]); err == nil {
			return t
		}
	}
	return time.Time{}
}

// analyzeLogs 扫描并分析日志
func analyzeLogs(projectName string, since time.Time) *PromptEvolutionReport {
	logsDir := getProjectLogsDir(projectName)
	entries, err := os.ReadDir(logsDir)
	if err != nil {
		return nil
	}

	report := &PromptEvolutionReport{
		GeneratedAt: time.Now(),
		ProjectName: projectName,
		Since:       since,
		AgentStats:  make(map[string]*AgentStat),
	}

	var analyses []*LogAnalysis
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".log") {
			continue
		}
		// 跳过 prompt-evolution 自身的日志，避免递归
		if strings.Contains(entry.Name(), "prompt-evolution") {
			continue
		}
		path := filepath.Join(logsDir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			continue
		}
		// 只分析 since 之后的日志
		if !since.IsZero() && info.ModTime().Before(since) {
			continue
		}
		a := parseLogFile(path)
		if a != nil {
			analyses = append(analyses, a)
		}
	}

	report.TotalLogs = len(analyses)

	// 按 agent 类型聚合
	// 临时聚合 git ops 以确定 top git op
	gitOpCounts := make(map[string]map[string]int)

	for _, a := range analyses {
		stat, ok := report.AgentStats[a.TaskType]
		if !ok {
			stat = &AgentStat{
				TotalToolCalls: make(map[string]int),
			}
			report.AgentStats[a.TaskType] = stat
		}
		stat.Count++
		stat.TotalErrors += a.Errors
		stat.TotalWrites += a.FileWrites + a.FileModifies
		for tool, count := range a.ToolCalls {
			stat.TotalToolCalls[tool] += count
		}
		if gitOpCounts[a.TaskType] == nil {
			gitOpCounts[a.TaskType] = make(map[string]int)
		}
		for gitOp, count := range a.GitOps {
			gitOpCounts[a.TaskType][gitOp] += count
		}
	}

	for agentType, stat := range report.AgentStats {
		ops := gitOpCounts[agentType]
		maxCount := 0
		for op, count := range ops {
			if count > maxCount {
				maxCount = count
				stat.TopGitOp = op
			}
		}
		_ = agentType
	}

	// 计算平均值
	for agentType, stat := range report.AgentStats {
		var totalSteps int
		var totalDuration time.Duration
		var successCount int
		for _, a := range analyses {
			if a.TaskType == agentType {
				totalSteps += a.TotalSteps
				totalDuration += a.Duration
				if a.HasResult {
					successCount++
				}
			}
		}
		if stat.Count > 0 {
			stat.AvgSteps = float64(totalSteps) / float64(stat.Count)
			stat.AvgDuration = totalDuration / time.Duration(stat.Count)
			stat.SuccessRate = float64(successCount) / float64(stat.Count)
		}
		_ = agentType // avoid unused if we extend later
	}

	// 提取常见失败模式
	report.TopIssues = extractIssuePatterns(analyses)

	// 填充趋势对比
	report.Trends = computeTrends(projectName, report)

	return report
}

// extractIssuePatterns 从日志分析中提取常见失败模式（聚合+根因+建议）
func extractIssuePatterns(analyses []*LogAnalysis) []IssuePattern {
	var issues []IssuePattern

	// ── 1. 按 agent 类型聚合统计 ──
	typeStats := make(map[string]*struct {
		count       int
		errCount    int
		highSteps   int    // >50 步的任务数
		maxSteps    int    // 最高步数
		noOutput    int    // 零产出任务数
		thinkTotal  int    // 总 think 次数
		thinkLen    int    // 总 think 长度
		totalSteps  int    // 总步数
		targetIDs   map[int]bool
	})

	for _, a := range analyses {
		s := typeStats[a.TaskType]
		if s == nil {
			s = &struct {
				count     int
				errCount  int
				highSteps int
				maxSteps  int
				noOutput  int
				thinkTotal int
				thinkLen   int
				totalSteps int
				targetIDs  map[int]bool
			}{
				targetIDs: make(map[int]bool),
			}
			typeStats[a.TaskType] = s
		}
		s.count++
		s.totalSteps += a.TotalSteps
		s.thinkTotal += a.ThinkCount
		s.thinkLen += a.ThinkTotalLen
		s.targetIDs[a.TargetID] = true
		if a.Errors > 0 {
			s.errCount++
		}
		if a.TotalSteps > 50 {
			s.highSteps++
			if a.TotalSteps > s.maxSteps {
				s.maxSteps = a.TotalSteps
			}
		}
		if !a.HasResult && a.TotalSteps > 10 {
			s.noOutput++
		}
	}

	// ── 2. 生成聚合后的问题 ──
	for agentType, s := range typeStats {
		if s.count == 0 {
			continue
		}
		avgSteps := float64(s.totalSteps) / float64(s.count)
		errRate := float64(s.errCount) / float64(s.count)
		avgThink := 0.0
		if s.count > 0 {
			avgThink = float64(s.thinkTotal) / float64(s.count)
		}

		// 2a. 高错误率
		if errRate > 0.3 {
			severity := "warning"
			if errRate > 0.6 {
				severity = "critical"
			}
			rootCause := "agent 在执行工具调用时频繁遇到错误，可能缺少错误处理指引或环境预检查步骤"
			suggestion := fmt.Sprintf("在 %s 的 prompt 中增加：1) 执行命令前检查环境状态；2) 遇到错误时的重试策略；3) 常见错误的快速诊断清单", agentType)
			if agentType == "doc" {
				rootCause = "doc-gardener 任务虽然执行了步骤，但未能成功写入文件，可能是 worktree 路径问题或权限问题"
				suggestion = "检查 doc-gardener workflow 的 worktree 创建逻辑，确保 agent 在正确的目录下执行文件写入"
			}
			issues = append(issues, IssuePattern{
				Pattern:     fmt.Sprintf("%s_high_error_rate", agentType),
				Count:       s.errCount,
				Description: fmt.Sprintf("%s 任务错误率 %.0f%% (%d/%d)，平均 %.1f 步", agentType, errRate*100, s.errCount, s.count, avgSteps),
				RootCause:   rootCause,
				Suggestion:  suggestion,
				Severity:    severity,
			})
		}

		// 2b. 高步数（聚合为一条）
		if s.highSteps > 0 {
			severity := "warning"
			if avgSteps > 80 {
				severity = "critical"
			}
			rootCause := "agent 对任务目标理解不清，在试探性操作中消耗了大量步骤"
			suggestion := fmt.Sprintf("在 %s 的 prompt 开头增加「快速开始」段落：明确第一步该做什么、预期产出是什么、什么情况下应该停止", agentType)
			if avgThink > 5 {
				rootCause = fmt.Sprintf("agent 过度思考（平均 %.1f 次 think），在分析阶段消耗了过多步骤，且缺乏明确的执行路径", avgThink)
				suggestion = fmt.Sprintf("在 %s 的 prompt 中限制 think 的使用：只在需要复杂决策时才 think，常规操作直接执行", agentType)
			}
			issues = append(issues, IssuePattern{
				Pattern:     fmt.Sprintf("%s_high_steps", agentType),
				Count:       s.highSteps,
				Description: fmt.Sprintf("%s 有 %d 个任务超过 50 步（最高 %d 步），平均 %.1f 步", agentType, s.highSteps, s.maxSteps, avgSteps),
				RootCause:   rootCause,
				Suggestion:  suggestion,
				Severity:    severity,
			})
		}

		// 2c. 零产出（聚合为一条）
		if s.noOutput > 0 {
			issues = append(issues, IssuePattern{
				Pattern:     fmt.Sprintf("%s_no_output", agentType),
				Count:       s.noOutput,
				Description: fmt.Sprintf("%s 有 %d 个任务执行了步骤但没有文件修改产出", agentType, s.noOutput),
				RootCause:   "agent 没有明确被要求「必须有文件修改产出」，或者任务边界不清晰导致 agent 在空转",
				Suggestion:  fmt.Sprintf("在 %s 的 prompt 中增加强制要求：「本任务必须有至少一次 WriteFile 或 StrReplaceFile 调用，如果没有找到可修改的内容，请明确报告原因并停止」", agentType),
				Severity:    "warning",
			})
		}
	}

	// 按严重程度排序：critical > warning > info
	severityOrder := map[string]int{"critical": 0, "warning": 1, "info": 2}
	sort.Slice(issues, func(i, j int) bool {
		return severityOrder[issues[i].Severity] < severityOrder[issues[j].Severity]
	})

	return issues
}

// computeTrends 与历史报告对比计算趋势
func computeTrends(projectName string, current *PromptEvolutionReport) *TrendComparison {
	// 查找最近的历史报告
	logsDir := getProjectLogsDir(projectName)
	entries, err := os.ReadDir(logsDir)
	if err != nil {
		return nil
	}

	var latestReportPath string
	var latestTime time.Time
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "prompt-evolution-report-") || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.ModTime().After(latestTime) && info.ModTime().Before(current.GeneratedAt.Add(-1*time.Minute)) {
			latestTime = info.ModTime()
			latestReportPath = filepath.Join(logsDir, entry.Name())
		}
	}
	if latestReportPath == "" {
		return nil
	}

	data, err := os.ReadFile(latestReportPath)
	if err != nil {
		return nil
	}
	var prev PromptEvolutionReport
	if err := json.Unmarshal(data, &prev); err != nil {
		return nil
	}

	trends := &TrendComparison{}

	// 计算全局平均值
	var curAvgSteps, prevAvgSteps float64
	var curErrRate, prevErrRate float64
	var curSuccessRate, prevSuccessRate float64
	var curCount, prevCount int

	for _, s := range current.AgentStats {
		curAvgSteps += s.AvgSteps * float64(s.Count)
		curErrRate += float64(s.TotalErrors)
		curSuccessRate += s.SuccessRate * float64(s.Count)
		curCount += s.Count
	}
	for _, s := range prev.AgentStats {
		prevAvgSteps += s.AvgSteps * float64(s.Count)
		prevErrRate += float64(s.TotalErrors)
		prevSuccessRate += s.SuccessRate * float64(s.Count)
		prevCount += s.Count
	}

	if curCount > 0 {
		curAvgSteps /= float64(curCount)
		curSuccessRate /= float64(curCount)
	}
	if prevCount > 0 {
		prevAvgSteps /= float64(prevCount)
		prevSuccessRate /= float64(prevCount)
	}
	if current.TotalLogs > 0 {
		curErrRate /= float64(current.TotalLogs)
	}
	if prev.TotalLogs > 0 {
		prevErrRate /= float64(prev.TotalLogs)
	}

	trends.AvgStepsDelta = curAvgSteps - prevAvgSteps
	trends.AvgDurationDelta = 0 // duration 目前不准确，暂不对比
	trends.ErrorRateDelta = curErrRate - prevErrRate
	trends.SuccessRateDelta = curSuccessRate - prevSuccessRate

	return trends
}

// generatePromptEvolutionReport 生成 JSON 报告文件
func generatePromptEvolutionReport(projectName string, report *PromptEvolutionReport) string {
	logsDir := getProjectLogsDir(projectName)
	timestamp := time.Now().Format("20060102-150405")
	reportPath := filepath.Join(logsDir, fmt.Sprintf("prompt-evolution-report-%s.json", timestamp))
	data, _ := json.MarshalIndent(report, "", "  ")
	os.WriteFile(reportPath, data, 0644)
	return reportPath
}

// checkAndTriggerPromptEvolution 检查是否需要触发 prompt 演进
func checkAndTriggerPromptEvolution(projectName string) {
	state := loadPromptEvolutionState(projectName)
	state.TaskCountSinceLast++

	thresholdCount := 5
	thresholdDuration := 1 * time.Hour

	shouldTrigger := state.TaskCountSinceLast >= thresholdCount &&
		time.Since(state.LastEvolutionTime) >= thresholdDuration

	if !shouldTrigger {
		savePromptEvolutionState(projectName, state)
		log.Printf("[prompt-evolution] %s: 计数 %d/%d, 上次分析 %s 前, 暂不触发",
			projectName, state.TaskCountSinceLast, thresholdCount,
			time.Since(state.LastEvolutionTime).Round(time.Minute))
		return
	}

	// 触发分析
	log.Printf("[prompt-evolution] %s: 触发分析 (计数 %d, 间隔 %s)",
		projectName, state.TaskCountSinceLast, time.Since(state.LastEvolutionTime).Round(time.Minute))

	// 分析日志
	report := analyzeLogs(projectName, state.LastEvolutionTime)
	if report == nil {
		log.Printf("[prompt-evolution] %s: 分析失败", projectName)
		return
	}

	reportPath := generatePromptEvolutionReport(projectName, report)

	// 更新状态
	state.LastEvolutionTime = time.Now()
	state.TaskCountSinceLast = 0
	state.TotalAnalyzed += report.TotalLogs
	state.LastReportPath = reportPath
	savePromptEvolutionState(projectName, state)

	log.Printf("[prompt-evolution] %s: 报告已生成 %s, 分析了 %d 个日志",
		projectName, reportPath, report.TotalLogs)

	// 启动 prompt-evolution agent 任务
	startPromptEvolutionTask(projectName, reportPath)
}

// startPromptEvolutionTask 启动 prompt 演进 agent 任务
func startPromptEvolutionTask(projectName string, reportPath string) {
	// 检查是否已有正在运行的 prompt-evolution 任务
	tasksMutex.RLock()
	for key, t := range tasks {
		if t.Type == TaskTypePromptEvolution && t.Status == "running" {
			tasksMutex.RUnlock()
			log.Printf("[prompt-evolution] %s: 已有演进任务在运行 (%s)，跳过", projectName, key)
			return
		}
	}
	tasksMutex.RUnlock()

	agentFile := prepareAgentConfig("prompt-evolution")
	if agentFile == "" {
		log.Printf("[prompt-evolution] %s: 准备 agent 配置失败", projectName)
		return
	}

	logsDir := getProjectLogsDir(projectName)
	os.MkdirAll(logsDir, 0755)
	timestamp := time.Now().Format("20060102-150405")
	logFileName := filepath.Join(logsDir, fmt.Sprintf("prompt-evolution-%s.log", timestamp))

	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Printf("[prompt-evolution] %s: 创建日志文件失败: %v", projectName, err)
		return
	}
	logFile.Close()

	targetID := int(time.Now().Unix())
	task := &Task{
		ID:          taskID(TaskTypePromptEvolution, targetID),
		Type:        TaskTypePromptEvolution,
		Status:      "running",
		ProjectName: projectName,
		TargetID:    targetID,
		TargetTitle: fmt.Sprintf("Prompt Evolution (%s)", projectName),
		StartTime:   time.Now(),
		LogFile:     logFileName,
		Metadata: map[string]interface{}{
			"tmuxSession": fmt.Sprintf("nwops-%s", taskID(TaskTypePromptEvolution, targetID)),
			"reportPath":  reportPath,
		},
	}

	if !setTask(task) {
		log.Printf("[prompt-evolution] %s: 演进任务已在运行中", projectName)
		return
	}

	// 获取 dashboard 项目根目录
	baseDir, err := os.Getwd()
	if err != nil {
		baseDir = "."
	}

	prompt := fmt.Sprintf(
		"请分析以下统计报告，识别 prompt 和 workflow 的改进机会。\n\n"+
			"报告文件路径: %s\n\n"+
			"分析完成后，将改进建议写入项目根目录的 prompt-evolution-report.md，\n"+
			"并根据数据驱动的判断，直接修改 prompts/*.md 或 workflows/*.yaml 文件。\n",
		reportPath,
	)

	vars := map[string]interface{}{
		"projectDir":  baseDir,
		"reportPath":  reportPath,
		"logFileName": logFileName,
		"agentFile":   agentFile,
		"prompt":      prompt,
	}

	runTaskWorkflow(task, vars, nil)
}

// handlePromptEvolutionStats 返回 prompt 演进统计数据供 Dashboard 预览
func handlePromptEvolutionStats(w http.ResponseWriter, r *http.Request) {
	projectName := getCurrentProjectName()
	if projectName == "" {
		writeJSONError(w, 400, "未选择项目")
		return
	}

	// 加载状态
	state := loadPromptEvolutionState(projectName)

	// 实时分析（从上次分析时间到现在）
	var report *PromptEvolutionReport
	if !state.LastEvolutionTime.IsZero() {
		report = analyzeLogs(projectName, state.LastEvolutionTime)
	} else {
		report = analyzeLogs(projectName, time.Time{})
	}

	// 查找历史报告文件
	logsDir := getProjectLogsDir(projectName)
	var historyReports []map[string]interface{}
	entries, _ := os.ReadDir(logsDir)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "prompt-evolution-report-") || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		info, _ := entry.Info()
		historyReports = append(historyReports, map[string]interface{}{
			"filename":   entry.Name(),
			"createdAt":  info.ModTime(),
			"path":       filepath.Join(logsDir, entry.Name()),
			"url":        "/logs/" + projectName + "/" + entry.Name(),
		})
	}
	// 按时间倒序
	for i, j := 0, len(historyReports)-1; i < j; i, j = i+1, j-1 {
		historyReports[i], historyReports[j] = historyReports[j], historyReports[i]
	}

	// 查找当前运行的 prompt-evolution 任务
	var runningTask *Task
	tasksMutex.RLock()
	for _, t := range tasks {
		if t.Type == TaskTypePromptEvolution && t.ProjectName == projectName && t.Status == "running" {
			runningTask = t
			break
		}
	}
	tasksMutex.RUnlock()

	result := map[string]interface{}{
		"projectName":    projectName,
		"state": map[string]interface{}{
			"taskCountSinceLast": state.TaskCountSinceLast,
			"lastEvolutionTime":  state.LastEvolutionTime,
			"totalAnalyzed":      state.TotalAnalyzed,
			"lastReportPath":     state.LastReportPath,
		},
		"report":         report,
		"historyReports": historyReports,
		"runningTask":    runningTask,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
