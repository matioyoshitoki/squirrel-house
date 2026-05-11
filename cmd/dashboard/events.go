package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// SSEEvent 服务端推送事件
type SSEEvent struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
	Time    int64                  `json:"time"`
}

// sseClient 单个 SSE 连接客户端
type sseClient struct {
	id     string
	ch     chan SSEEvent
	closed uint32 // atomic
}

var (
	sseClients   = make(map[string]*sseClient)
	sseClientsMu sync.RWMutex
	sseClientID  uint64
)

// broadcastEvent 向所有连接的 SSE 客户端广播事件（非阻塞）
func broadcastEvent(eventType string, payload map[string]interface{}) {
	ev := SSEEvent{
		Type:    eventType,
		Payload: payload,
		Time:    time.Now().UnixMilli(),
	}

	sseClientsMu.RLock()
	clients := make([]*sseClient, 0, len(sseClients))
	for _, c := range sseClients {
		clients = append(clients, c)
	}
	sseClientsMu.RUnlock()

	for _, c := range clients {
		if atomic.LoadUint32(&c.closed) == 1 {
			continue
		}
		select {
		case c.ch <- ev:
		default:
			// 客户端缓冲区满，丢弃事件并标记清理
			atomic.StoreUint32(&c.closed, 1)
			closeClient(c)
		}
	}
}

// broadcastTaskChanged 广播任务状态变化事件
func broadcastTaskChanged(task *Task, action string) {
	broadcastEvent("task:changed", map[string]interface{}{
		"id":         task.ID,
		"type":       task.Type,
		"status":     task.Status,
		"action":     action,
		"targetId":   task.TargetID,
		"targetTitle": task.TargetTitle,
		"projectName": task.ProjectName,
	})
}

func closeClient(c *sseClient) {
	if atomic.CompareAndSwapUint32(&c.closed, 0, 1) {
		close(c.ch)
	}
	sseClientsMu.Lock()
	delete(sseClients, c.id)
	sseClientsMu.Unlock()
}

// handleSSEEvents SSE 事件流端点
func handleSSEEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	id := fmt.Sprintf("%d", atomic.AddUint64(&sseClientID, 1))
	client := &sseClient{
		id: id,
		ch: make(chan SSEEvent, 64),
	}

	sseClientsMu.Lock()
	sseClients[id] = client
	sseClientsMu.Unlock()

	// 清理
	defer closeClient(client)

	// 发送初始连接成功事件
	fmt.Fprint(w, "event: connected\n")
	fmt.Fprintf(w, "data: %s\n\n", `{"type":"connected","payload":{}}`)
	flusher.Flush()

	done := r.Context().Done()
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			fmt.Fprint(w, ":ping\n\n")
			flusher.Flush()
		case ev, ok := <-client.ch:
			if !ok {
				return
			}
			data, err := json.Marshal(ev)
			if err != nil {
				continue
			}
			fmt.Fprintf(w, "event: %s\ndata: %s\n\n", ev.Type, string(data))
			flusher.Flush()
		}
	}
}

