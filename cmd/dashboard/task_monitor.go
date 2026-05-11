package main

import (
	"log"
	"os/exec"
	"time"
)

// getTaskTmuxSession 从 task metadata 中获取 tmux session 名称
func getTaskTmuxSession(task *Task) string {
	if task.Metadata == nil {
		return ""
	}
	if s, ok := task.Metadata["tmuxSession"].(string); ok {
		return s
	}
	return ""
}

// isTmuxSessionAlive 检查 tmux session 是否存在
func isTmuxSessionAlive(session string) bool {
	err := exec.Command("tmux", "has-session", "-t", session).Run()
	return err == nil
}

// checkRunningTasksHealth 定时检查所有 running 任务的心跳状态。
// 如果心跳超时（超过 3 分钟未更新），将任务标记为 interrupted。
// 不再检查 PID、tmux session 或日志关键字——所有状态以 .status.json 为准。
func checkRunningTasksHealth() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		projectName := getCurrentProjectName()
		tasksMutex.Lock()
		var deadTasks []*Task
		for _, task := range tasks {
			if task.Status != "running" || task.ProjectName != projectName {
				continue
			}

			// 新启动的任务留出 2 分钟缓冲，避免首次心跳还没写入就误判
			if time.Since(task.StartTime) < 2*time.Minute {
				continue
			}

			status := readTaskStatusFile(task.ID)
			if status == nil {
				// 状态文件不存在，说明 heartbeat 从未写入 → 标记为 interrupted
				log.Printf("🩺 %s 无状态文件，标记为 interrupted", task.ID)
				task.Status = "interrupted"
				task.EndTime = time.Now()
				deadTasks = append(deadTasks, task)
				continue
			}

			// 检查心跳时间
			if time.Since(status.LastHeartbeat) > 3*time.Minute {
				// dashboard 重启后，tmux 任务可能仍在运行但失去了心跳更新
				// 如果 tmux session 仍然存在，不标记为 interrupted
				session := getTaskTmuxSession(task)
				if session != "" && isTmuxSessionAlive(session) {
					log.Printf("🩺 %s 心跳超时但 tmux session %s 仍在运行，跳过", task.ID, session)
					continue
				}
				log.Printf("🩺 %s 心跳超时 (last=%s)，标记为 interrupted", task.ID, status.LastHeartbeat.Format("15:04:05"))
				task.Status = "interrupted"
				task.EndTime = time.Now()
				deadTasks = append(deadTasks, task)
			}
		}
		tasksMutex.Unlock()

		for _, task := range deadTasks {
			appendToFile(task.LogFile, []byte("\n⏸️ 任务心跳超时，被标记为中断\n"))
			go saveTaskState()
			go broadcastTaskChanged(task, "interrupted")
		}
	}
}
