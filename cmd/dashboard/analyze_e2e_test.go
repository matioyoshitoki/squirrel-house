package main

import (
	"fmt"
	"os"
	"testing"
)

func TestAnalyzeE2ELog(t *testing.T) {
	logPath := "logs/test-project/e2e-issue-64-20260427-003255.log"
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Skipf("日志文件不存在: %s", logPath)
	}
	a := parseLogFile(logPath)
	if a == nil {
		t.Fatal("failed to parse")
	}
	fmt.Printf("Task: %s #%d\n", a.TaskType, a.TargetID)
	fmt.Printf("Steps: %d\n", a.TotalSteps)
	fmt.Printf("Errors: %d\n", a.Errors)
	fmt.Printf("Think: %d (len=%d)\n", a.ThinkCount, a.ThinkTotalLen)
	fmt.Printf("Writes: %d, Modifies: %d\n", a.FileWrites, a.FileModifies)
	fmt.Printf("HasResult: %v\n", a.HasResult)
	fmt.Println("Tools:")
	for tool, count := range a.ToolCalls {
		fmt.Printf("  %s: %d\n", tool, count)
	}
	fmt.Println("GitOps:")
	for op, count := range a.GitOps {
		fmt.Printf("  %s: %d\n", op, count)
	}
}
