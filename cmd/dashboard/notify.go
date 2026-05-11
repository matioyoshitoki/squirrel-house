package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// getLarkChatID 获取飞书群聊 ID（优先环境变量，其次配置文件）
func getLarkChatID() string {
	if id := os.Getenv("LARK_CHAT_ID"); id != "" {
		return id
	}
	return config.LarkChatID
}

// notifyLark 发送飞书通知（阻塞调用，内部已做错误兜底）
func notifyLark(title string, content string) {
	chatID := getLarkChatID()
	if chatID == "" {
		return
	}

	// 截断过长内容，避免消息发送失败
	maxContentLen := 4000
	if len(content) > maxContentLen {
		content = content[:maxContentLen] + "\n\n...（内容已截断）"
	}

	markdown := fmt.Sprintf("## 🤖 %s\n\n%s\n\n---\n*Agent Control Center · %s*",
		title, content, time.Now().Format("01/02 15:04"))

	cmd := exec.Command("lark-cli", "im", "+messages-send",
		"--as", "bot",
		"--chat-id", chatID,
		"--markdown", markdown)
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Printf("⚠️ 飞书通知发送失败: %v, output: %s", err, string(output))
	} else {
		log.Printf("✅ 飞书通知已发送: %s", title)
	}
}

// notifyLarkAsync 异步发送飞书通知，不阻塞主流程
func notifyLarkAsync(title string, content string) {
	go notifyLark(title, content)
}

// notifyDevTaskStarted 开发任务启动通知
func notifyDevTaskStarted(issueNumber int, issueTitle string, agentType string) {
	action := "开发任务"
	if agentType == "rework" {
		action = "返工修复任务"
	}
	notifyLarkAsync(
		fmt.Sprintf("%s已启动 · Issue #%d", action, issueNumber),
		fmt.Sprintf("**Issue**: #%d %s\n**动作**: %s开始执行", issueNumber, issueTitle, action),
	)
}

// notifyDevTaskFinished 开发任务结束通知
func notifyDevTaskFinished(issueNumber int, issueTitle string, status string, agentType string) {
	action := "开发任务"
	if agentType == "rework" {
		action = "返工修复任务"
	}

	var emoji, result string
	switch status {
	case "success":
		emoji = "✅"
		result = "执行成功"
	case "failed":
		emoji = "❌"
		result = "执行失败"
	case "stopped":
		emoji = "🛑"
		result = "已手动停止"
	case "interrupted":
		emoji = "⏸️"
		result = "已中断（可能超时）"
	default:
		emoji = "⚠️"
		result = "状态异常: " + status
	}

	notifyLarkAsync(
		fmt.Sprintf("%s %s · Issue #%d", emoji, result, issueNumber),
		fmt.Sprintf("**Issue**: #%d %s\n**动作**: %s\n**结果**: %s", issueNumber, issueTitle, action, result),
	)
}

// notifyReviewStarted Review 任务启动通知
func notifyReviewStarted(prNumber int, branch string) {
	notifyLarkAsync(
		fmt.Sprintf("🔍 PR Review 已启动 · PR #%d", prNumber),
		fmt.Sprintf("**PR**: #%d\n**分支**: `%s`", prNumber, branch),
	)
}

// notifyReviewFinished Review 完成通知
func notifyReviewFinished(prNumber int, result string) {
	var emoji, resultText string
	if result == "passed" {
		emoji = "✅"
		resultText = "审查通过，可以合并"
	} else {
		emoji = "❌"
		resultText = "审查未通过，需要返工"
	}
	notifyLarkAsync(
		fmt.Sprintf("%s PR Review 完成 · PR #%d", emoji, prNumber),
		fmt.Sprintf("**PR**: #%d\n**审查结论**: %s", prNumber, resultText),
	)
}

// notifyReworkMarked PR 标记返工通知
func notifyReworkMarked(prNumber int, issueNumber int, blockingIssues string) {
	content := fmt.Sprintf("**PR**: #%d", prNumber)
	if issueNumber > 0 {
		content += fmt.Sprintf("\n**关联 Issue**: #%d", issueNumber)
	}
	if strings.TrimSpace(blockingIssues) != "" && blockingIssues != "请根据审查报告中的问题进行修改。" {
		content += fmt.Sprintf("\n\n**必须修复的问题**:\n%s", blockingIssues)
	}
	content += "\n\n已自动启动 Developer Agent 进行修复。"
	notifyLarkAsync(
		fmt.Sprintf("🔧 PR 已标记返工 · PR #%d", prNumber),
		content,
	)
}

// notifyPRMerged PR 合并通知
func notifyPRMerged(prNumber int) {
	notifyLarkAsync(
		fmt.Sprintf("🎉 PR 已合并 · PR #%d", prNumber),
		fmt.Sprintf("**PR**: #%d\n\n已通过 squash 方式合并到主分支。", prNumber),
	)
}

// addE2EPRComment 在 PR 上添加 E2E 测试报告评论
func addE2EPRComment(projectPath string, prNumber int, body string) {
	// 截断过长内容
	maxLen := 60000
	if len(body) > maxLen {
		body = body[:maxLen] + "\n\n---\n*报告内容过长，已在上方截断。请在 Dashboard 中查看完整报告。*"
	}

	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("e2e-pr-comment-%d-%d.md", prNumber, time.Now().Unix()))
	if err := os.WriteFile(tmpFile, []byte(body), 0644); err != nil {
		log.Printf("[e2e] 写入 PR 评论临时文件失败: %v", err)
		return
	}
	defer os.Remove(tmpFile)

	platform := getPlatform(projectPath)
	if err := platform.CreateMRComment(prNumber, body); err != nil {
		log.Printf("[e2e] PR/MR #%d 评论添加失败: %v", prNumber, err)
	} else {
		log.Printf("[e2e] PR/MR #%d 评论已添加", prNumber)
	}
}

// notifyE2EFinished E2E 测试完成通知（推送报告原文到飞书）
func notifyE2EFinished(issueNumber int, issueTitle string, status string, reportContent string, prNumber int) {
	var emoji, result string
	switch status {
	case "success":
		emoji = "✅"
		result = "E2E 测试通过"
	case "failed":
		emoji = "❌"
		result = "E2E 测试失败"
	case "stopped":
		emoji = "🛑"
		result = "E2E 已手动停止"
	default:
		emoji = "⚠️"
		result = "E2E 状态异常: " + status
	}

	// 报告原文直接作为飞书通知正文
	content := fmt.Sprintf("**Issue**: #%d %s\n**结果**: %s", issueNumber, issueTitle, result)
	if prNumber > 0 {
		content += fmt.Sprintf("\n**PR**: #%d", prNumber)
	}
	content += "\n\n---\n\n" + reportContent

	notifyLarkAsync(
		fmt.Sprintf("%s %s · Issue #%d", emoji, result, issueNumber),
		content,
	)
}
