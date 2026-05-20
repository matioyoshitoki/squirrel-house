package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

//go:embed static/*
var staticFS embed.FS

// Config 项目配置
type Config struct {
	Port                   string    `json:"port"`
	LarkChatID             string    `json:"larkChatId"`
	Projects               []Project `json:"projects"`
	FlutterPreviewTemplate string    `json:"flutterPreviewTemplate,omitempty"`
}

type TechStack struct {
	Backend        string `json:"backend"`
	Mobile         string `json:"mobile"`
	Frontend       string `json:"frontend"`
	DB             string `json:"db"`
	PackageManager string `json:"packageManager"`
}

type Project struct {
	Name          string     `json:"name"`
	Path          string     `json:"path"`
	Description   string     `json:"description"`
	DefaultBranch string     `json:"defaultBranch"`
	Domain        string     `json:"domain"`
	TechStack     TechStack  `json:"techStack"`
	Repo          string     `json:"repo"`
	Platform      string     `json:"platform,omitempty"` // "github" | "gitlab"，空则自动推断
}

type Issue struct {
	Number         int     `json:"number"`
	Title          string  `json:"title"`
	State          string  `json:"state"`
	Labels         []Label `json:"labels"`
	Body           string  `json:"body"`
	UpdatedAt      string  `json:"updatedAt"`
	Classification string  `json:"classification"`
}

type Label struct {
	Name string `json:"name"`
}

type LogInfo struct {
	Filename    string    `json:"filename"`
	IssueNumber int       `json:"issueNumber"`
	Timestamp   time.Time `json:"timestamp"`
	Size        int64     `json:"size"`
	ModifiedAt  time.Time `json:"modifiedAt"`
}

var (
	config              Config
	currentProjectIndex int
	projectMutex        sync.RWMutex
	tasks               = make(map[string]*Task)
	tasksMutex          sync.RWMutex
	tmpl                *template.Template
	projectRoot         string // 项目根目录，解决 cd 到其他目录后相对路径失效的问题
)

func main() {
	// 定位项目根目录（不依赖当前工作目录）
	projectRoot = findProjectRoot()
	log.Printf("📁 项目根目录: %s", projectRoot)

	// 加载配置与状态
	loadConfig()
	loadState()

	// 初始化模板
	initTemplates()

	// 加载 CTX prompt 模板
	if err := loadCtxTemplates(); err != nil {
		log.Printf("⚠️ 无法加载 CTX 模板: %v", err)
	} else {
		log.Printf("✅ 已加载 %d 个 CTX 模板", len(ctxTemplates))
	}

	// 启动时从日志重建任务状态
	rebuildTaskStatus()

	// 启动僵尸任务清理器
	go cleanZombieTasks()

	// 启动任务进程健康检查
	go checkRunningTasksHealth()

	// 注册路由
	http.HandleFunc("/", handleDashboard)
	// 优先从文件系统读取 static 目录，开发时不需要重新编译
	staticDir := filepath.Join(projectRoot, "cmd", "dashboard", "static")
	if _, err := os.Stat(staticDir); err == nil {
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	} else {
		http.Handle("/static/", http.FileServer(http.FS(staticFS)))
	}
	
	// 静态文件服务：允许浏览器直接访问 logs 下的设计资产
	logsDir := filepath.Join(projectRoot, "logs")
	http.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir(logsDir))))

	// API 路由
	http.HandleFunc("/api/issues", handleAPIIssues)
	http.HandleFunc("/api/design-assets", handleDesignAssets)
	http.HandleFunc("/api/build-design-preview", handleBuildDesignPreview)
	http.HandleFunc("/api/design-feedback", handleDesignFeedback)
	http.HandleFunc("/api/start-dev", handleStartDev)
	http.HandleFunc("/api/resume-dev", handleResumeDev)
	http.HandleFunc("/api/start-design", handleStartDesign)
	http.HandleFunc("/api/running-tasks", handleRunningTasks)
	http.HandleFunc("/api/stop-task", handleStopTask)
	http.HandleFunc("/api/logs", handleLogs)
	http.HandleFunc("/api/log-content", handleLogContent)
	http.HandleFunc("/api/log-tail", handleLogTail)
	http.HandleFunc("/api/log-status", handleLogStatus)
	http.HandleFunc("/api/pull-requests", handlePullRequests)
	http.HandleFunc("/api/review-tasks", handleReviewTasks)
	http.HandleFunc("/api/review-pr", handleReviewPR)
	http.HandleFunc("/api/review-report", handleReviewReport)
	http.HandleFunc("/api/rework-pr", handleReworkPR)
	http.HandleFunc("/api/merge-pr", handleMergePR)
	http.HandleFunc("/api/doc-gardener", handleDocGardener)
	http.HandleFunc("/api/merged-doc-gardener", handleMergedDocGardener)
	http.HandleFunc("/api/doc-gardener-tasks", handleDocGardenerTasks)
	http.HandleFunc("/api/start-pm", handleStartPM)
	http.HandleFunc("/api/run-e2e", handleRunE2E)
	http.HandleFunc("/api/projects", handleProjects)
	http.HandleFunc("/api/switch-project", handleSwitchProject)
	http.HandleFunc("/api/events", handleSSEEvents)
	http.HandleFunc("/api/prompt-evolution-stats", handlePromptEvolutionStats)

	l, err := getListener(":" + config.Port)
	if err != nil {
		log.Fatalf("❌ 无法监听端口 %s: %v", config.Port, err)
	}

	// 注册路由（已在上面完成）
	log.Printf("🚀 Agent Dashboard 启动于 http://localhost:%s", config.Port)
	runServer(l)
}

// getListenerWithEnv 获取 listener，支持从父进程热重启继承 fd
// envLookup 参数用于测试注入环境变量
func getListenerWithEnv(addr string, envLookup func(string) string) (net.Listener, error) {
	if fdStr := envLookup("DASHBOARD_FD"); fdStr != "" {
		fd, err := strconv.Atoi(fdStr)
		if err == nil {
			f := os.NewFile(uintptr(fd), "listener")
			if ln, err := net.FileListener(f); err == nil {
				log.Printf("🔄 从父进程继承 listener (fd=%d)", fd)
				return ln, nil
			}
		}
	}
	return net.Listen("tcp", addr)
}

// getListener 获取 listener，支持从父进程热重启继承 fd
func getListener(addr string) (net.Listener, error) {
	return getListenerWithEnv(addr, os.Getenv)
}

// forkChild fork 新进程，把 listener fd 传递过去
func forkChild(ln net.Listener) (*exec.Cmd, error) {
	tcpLn, ok := ln.(*net.TCPListener)
	if !ok {
		return nil, fmt.Errorf("listener is not TCP")
	}
	f, err := tcpLn.File()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	exe, err := os.Executable()
	if err != nil {
		exe = os.Args[0]
	}

	cmd := exec.Command(exe, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("DASHBOARD_FD=%d", 3))
	cmd.ExtraFiles = []*os.File{f}
	// 子进程独立进程组，避免继承父进程的终端信号
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	log.Printf("🚀 子进程已启动 (pid=%d)", cmd.Process.Pid)
	return cmd, nil
}

// runServer 运行 HTTP server，支持 SIGHUP 热重启
func runServer(ln net.Listener) {
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: nil, // 使用 http.DefaultServeMux
	}

	go func() {
		if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Printf("❌ Server error: %v", err)
		}
	}()

	// 保存 pid 供外部 reload 使用
	os.WriteFile(filepath.Join(projectRoot, ".dashboard.pid"), []byte(strconv.Itoa(os.Getpid())), 0644)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	switch sig {
	case syscall.SIGHUP:
		log.Printf("📡 收到 SIGHUP，启动热重启...")
		childCmd, err := forkChild(ln)
		if err != nil {
			log.Printf("⚠️ fork 子进程失败: %v", err)
		} else {
			// 给子进程一点时间接管 listener
			time.Sleep(500 * time.Millisecond)
		}

		// 关闭 HTTP 服务，不再接收新请求
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("⚠️ 优雅关闭失败: %v", err)
		}
		cancel()

		// 设置关闭标志，避免与新进程竞争写入状态文件
		setShuttingDown(true)

		// 忽略后续信号，专心等待任务完成
		signal.Ignore(syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		// 等待所有 running 任务完成
		waitForRunningTasks()

		// 回收子进程状态，避免其成为僵尸
		if childCmd != nil && childCmd.Process != nil {
			done := make(chan error, 1)
			go func() { done <- childCmd.Wait() }()
			select {
			case <-done:
				// 子进程已退出
			case <-time.After(5 * time.Second):
				// 子进程仍在运行（正常情况，新进程接管服务）
			}
		}

		log.Printf("👋 旧进程退出")
		return
	case syscall.SIGINT, syscall.SIGTERM:
		log.Printf("📡 收到 %v，优雅关闭...", sig)

		// 关闭 HTTP 服务，不再接收新请求
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("⚠️ 优雅关闭失败: %v", err)
		}
		cancel()

		// 设置关闭标志，避免竞争写入状态文件
		setShuttingDown(true)

		// 忽略后续信号，专心等待任务完成
		signal.Ignore(syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		// 等待所有 running 任务完成（PM2 kill_timeout 为 5min，预留 1min 缓冲）
		waitForRunningTasksWithTimeout(4*time.Minute, 5*time.Second)

		log.Printf("👋 进程退出")
		return
	}
}


func handleProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}
	projectMutex.RLock()
	idx := currentProjectIndex
	projectMutex.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"projects":        config.Projects,
		"currentIndex":    idx,
		"currentProject":  config.Projects[idx],
	})
}

func handleSwitchProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeJSONError(w, 405, "Method not allowed")
		return
	}
	var req struct {
		ProjectName string `json:"projectName"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, 400, err.Error())
		return
	}
	targetIdx := -1
	for i, p := range config.Projects {
		if p.Name == req.ProjectName {
			targetIdx = i
			break
		}
	}
	if targetIdx == -1 {
		writeJSONError(w, 404, "Project not found: "+req.ProjectName)
		return
	}
	projectMutex.Lock()
	currentProjectIndex = targetIdx
	projectMutex.Unlock()

	// 清空内存中的任务状态，并重建当前项目状态
	tasksMutex.Lock()
	for k := range tasks {
		delete(tasks, k)
	}
	tasksMutex.Unlock()

	log.Printf("🔄 已切换项目到: %s (%s)", config.Projects[targetIdx].Name, config.Projects[targetIdx].Path)

	go broadcastEvent("project:switched", map[string]interface{}{
		"projectName": config.Projects[targetIdx].Name,
		"path":        config.Projects[targetIdx].Path,
	})

	// 异步重建当前项目的任务状态
	go rebuildTaskStatus()

	// 持久化当前项目索引
	saveState()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "switched",
		"message": fmt.Sprintf("已切换到项目: %s", config.Projects[targetIdx].Name),
		"path":    config.Projects[targetIdx].Path,
	})
}

func stateFilePath() string {
	return filepath.Join(projectRoot, "configs", "dashboard-state.json")
}

func loadState() {
	if data, err := os.ReadFile(stateFilePath()); err == nil {
		var state struct {
			CurrentProjectIndex int `json:"currentProjectIndex"`
		}
		if err := json.Unmarshal(data, &state); err == nil {
			projectMutex.Lock()
			if state.CurrentProjectIndex >= 0 && state.CurrentProjectIndex < len(config.Projects) {
				currentProjectIndex = state.CurrentProjectIndex
				log.Printf("📍 已恢复上次选择的项目: %s", config.Projects[currentProjectIndex].Name)
			}
			projectMutex.Unlock()
		}
	}
}

func saveState() {
	projectMutex.RLock()
	idx := currentProjectIndex
	projectMutex.RUnlock()
	stateData, _ := json.Marshal(map[string]interface{}{
		"currentProjectIndex": idx,
	})
	os.MkdirAll(filepath.Join(projectRoot, "configs"), 0755)
	if err := os.WriteFile(stateFilePath(), stateData, 0644); err != nil {
		log.Printf("⚠️ 保存状态失败: %v", err)
	}
}

func findProjectRoot() string {
	// 1. 当前工作目录
	markers := []string{"go.mod", "configs/dashboard.json"}
	for _, marker := range markers {
		if _, err := os.Stat(marker); err == nil {
			return "."
		}
	}

	// 2. 可执行文件目录的父目录（针对 build/dashboard）
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		parent := filepath.Join(exeDir, "..")
		for _, marker := range markers {
			if _, err := os.Stat(filepath.Join(parent, marker)); err == nil {
				abs, _ := filepath.Abs(parent)
				return abs
			}
		}
	}

	return "."
}

func loadConfig() {
	// 默认配置
	config = Config{
		Port:     "9090",
		Projects: []Project{},
	}

	// 尝试从配置文件加载
	configPath := filepath.Join(projectRoot, "configs", "dashboard.json")
	if data, err := os.ReadFile(configPath); err == nil {
		json.Unmarshal(data, &config)
		log.Printf("📋 已加载配置文件: %s", configPath)
	} else {
		log.Printf("⚠️ 无法加载配置文件 %s: %v，使用默认配置", configPath, err)
	}

	// 默认值填充
	if config.FlutterPreviewTemplate == "" {
		config.FlutterPreviewTemplate = ".tools/flutter_preview_generic"
	}

	// 环境变量覆盖
	if port := os.Getenv("DASHBOARD_PORT"); port != "" {
		config.Port = port
	}
}

func initTemplates() {
	// 优先从文件系统读取，开发时不需要重新编译
	staticIndexPath := filepath.Join(projectRoot, "cmd", "dashboard", "static", "index.html")
	if _, statErr := os.Stat(staticIndexPath); statErr == nil {
		t, err := template.ParseFiles(staticIndexPath)
		if err != nil {
			log.Printf("⚠️ 无法解析磁盘模板: %v, 回退到内嵌模板", err)
			t, err = template.ParseFS(staticFS, "static/index.html")
			if err != nil {
				log.Printf("⚠️ 无法解析内嵌模板: %v, 使用硬编码回退", err)
				t = template.Must(template.New("dashboard").Parse(fallbackTemplate))
			}
		}
		tmpl = t
	} else {
		t, err := template.ParseFS(staticFS, "static/index.html")
		if err != nil {
			log.Printf("⚠️ 无法解析模板: %v, 使用内嵌模板", err)
			t = template.Must(template.New("dashboard").Parse(fallbackTemplate))
		}
		tmpl = t
	}
}

func getProjectPath() string {
	projectMutex.RLock()
	defer projectMutex.RUnlock()
	idx := currentProjectIndex
	if idx < 0 || idx >= len(config.Projects) {
		idx = 0
	}
	if len(config.Projects) > 0 {
		return config.Projects[idx].Path
	}
	return "."
}

// getCurrentProjectName 获取当前项目名称
func getCurrentProjectName() string {
	projectMutex.RLock()
	defer projectMutex.RUnlock()
	idx := currentProjectIndex
	if idx < 0 || idx >= len(config.Projects) {
		idx = 0
	}
	if len(config.Projects) > 0 {
		return config.Projects[idx].Name
	}
	return "default"
}

// getProjectPathByName 根据项目名称返回项目路径
func getProjectPathByName(name string) string {
	projectMutex.RLock()
	defer projectMutex.RUnlock()
	for _, p := range config.Projects {
		if p.Name == name {
			return p.Path
		}
	}
	return getProjectPath()
}

// getProjectLogsDir 获取指定项目的日志目录
func getProjectLogsDir(projectName string) string {
	absPath, _ := filepath.Abs(filepath.Join("logs", projectName))
	return absPath
}

// getCurrentProjectLogsDir 获取当前项目的日志目录
func getCurrentProjectLogsDir() string {
	return getProjectLogsDir(getCurrentProjectName())
}

// getLogPath 解析日志文件的真实路径
// 兼容三种输入：纯文件名、logs/前缀的完整路径、旧格式
func getLogPath(filename string) string {
	if strings.Contains(filename, "..") {
		return ""
	}
	
	// 如果 filename 已经是绝对路径，直接返回（兼容 getCurrentProjectLogsDir 改为绝对路径后的情况）
	if filepath.IsAbs(filename) {
		return filename
	}
	
	filename = strings.TrimPrefix(filename, "logs/")
	
	// 如果 filename 已包含项目目录（如 new-world-monorepo/dev-issue-1.log），直接尝试
	if strings.Contains(filename, "/") {
		if _, err := os.Stat(filepath.Join("logs", filename)); err == nil {
			return filepath.Join("logs", filename)
		}
	}
	
	// 优先在当前项目日志目录查找
	currentDir := getCurrentProjectLogsDir()
	if _, err := os.Stat(filepath.Join(currentDir, filename)); err == nil {
		return filepath.Join(currentDir, filename)
	}
	
	// 回退到旧的全局 logs 目录（兼容已有日志）
	if _, err := os.Stat(filepath.Join("logs", filename)); err == nil {
		return filepath.Join("logs", filename)
	}
	
	// 默认返回当前项目日志目录下的路径（用于新建）
	return filepath.Join(currentDir, filename)
}
