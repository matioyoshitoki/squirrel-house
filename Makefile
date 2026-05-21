# =============================================================================
# New-World-Ops Dashboard — 全环境 Makefile
# 用途：自动安装缺失依赖、构建、重启、管理 Dashboard 服务
# =============================================================================

.PHONY: all setup install-deps install-system-deps install-node-deps install-go-deps \
        build build-css run dev stop restart reload reload-hot status logs monit save delete clean test

# -----------------------------------------------------------------------------
# 配置变量
# -----------------------------------------------------------------------------
BINARY_NAME   := dashboard
BUILD_DIR     := ./build
LOG_DIR       := ./logs
PM2_ECOSYSTEM := ./ecosystem.config.js
CSS_INPUT     := cmd/dashboard/static/input.css
CSS_OUTPUT    := cmd/dashboard/static/style.css

# 检测操作系统 (Darwin / Linux)
UNAME_S := $(shell uname -s)

# 颜色输出
COLOR_RESET  := \033[0m
COLOR_GREEN  := \033[32m
COLOR_YELLOW := \033[33m
COLOR_RED    := \033[31m
COLOR_CYAN   := \033[36m

# -----------------------------------------------------------------------------
# 默认目标
# -----------------------------------------------------------------------------
all: build

# -----------------------------------------------------------------------------
# 一键 setup：全新环境从零到运行
# -----------------------------------------------------------------------------
setup: install-deps build
	@echo "$(COLOR_GREEN)✅ 环境就绪！使用以下命令管理服务：$(COLOR_RESET)"
	@echo "   make start   — 启动 Dashboard (pm2)"
	@echo "   make stop    — 停止 Dashboard"
	@echo "   make restart — 重启 Dashboard"
	@echo "   make status  — 查看运行状态"
	@echo "   make logs    — 查看实时日志"

# -----------------------------------------------------------------------------
# 依赖安装：自动检测并安装缺失依赖
# -----------------------------------------------------------------------------
install-deps: install-system-deps install-node-deps install-go-deps
	@echo "$(COLOR_GREEN)✅ 所有依赖检查完毕$(COLOR_RESET)"

# --- 系统级依赖检查与安装引导 ---
install-system-deps:
	@echo "$(COLOR_CYAN)🔍 检查系统级依赖...$(COLOR_RESET)"

# Go
	@which go > /dev/null 2>&1 || { \
		echo "$(COLOR_RED)❌ Go 未安装$(COLOR_RESET)"; \
		if [ "$(UNAME_S)" = "Darwin" ]; then \
			echo "   macOS: brew install go"; \
		elif [ "$(UNAME_S)" = "Linux" ]; then \
			echo "   Ubuntu/Debian: sudo apt update && sudo apt install -y golang-go"; \
			echo "   CentOS/RHEL:  sudo yum install -y golang"; \
		fi; \
		echo "   或访问 https://go.dev/dl/ 手动安装 (要求 Go 1.21+)"; \
		exit 1; \
	}
	@GO_VERSION=$$(go version | awk '{print $$3}' | sed 's/go//'); \
	MAJOR=$$(echo $$GO_VERSION | cut -d. -f1); \
	MINOR=$$(echo $$GO_VERSION | cut -d. -f2); \
	if [ "$$MAJOR" -lt 1 ] || { [ "$$MAJOR" -eq 1 ] && [ "$$MINOR" -lt 21 ]; }; then \
		echo "$(COLOR_RED)❌ Go 版本过低: $$GO_VERSION (需要 ≥ 1.21)$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_GREEN)   ✓ Go $$(go version | awk '{print $$3}')$(COLOR_RESET)"

# Node.js & npm
	@which node > /dev/null 2>&1 || { \
		echo "$(COLOR_RED)❌ Node.js 未安装$(COLOR_RESET)"; \
		if [ "$(UNAME_S)" = "Darwin" ]; then \
			echo "   macOS: brew install node"; \
		elif [ "$(UNAME_S)" = "Linux" ]; then \
			echo "   Ubuntu/Debian: sudo apt update && sudo apt install -y nodejs npm"; \
			echo "   CentOS/RHEL:  sudo yum install -y nodejs npm"; \
		fi; \
		echo "   或访问 https://nodejs.org/ 手动安装 (要求 Node.js 18+)"; \
		exit 1; \
	}
	@NODE_MAJOR=$$(node --version | sed 's/v//' | cut -d. -f1); \
	if [ "$$NODE_MAJOR" -lt 18 ]; then \
		echo "$(COLOR_RED)❌ Node.js 版本过低: $$(node --version) (需要 ≥ 18)$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_GREEN)   ✓ Node.js $$(node --version)$(COLOR_RESET)"

# PM2
	@which pm2 > /dev/null 2>&1 || { \
		echo "$(COLOR_YELLOW)⚠️ PM2 未安装，正在全局安装...$(COLOR_RESET)"; \
		npm install -g pm2 || { echo "$(COLOR_RED)❌ PM2 安装失败，请手动运行: npm install -g pm2$(COLOR_RESET)"; exit 1; }; \
	}
	@echo "$(COLOR_GREEN)   ✓ PM2 $$(pm2 --version | head -1)$(COLOR_RESET)"

# 可选：GitHub CLI / GitLab CLI（仅警告，不阻断）
	@which gh > /dev/null 2>&1 || which glab > /dev/null 2>&1 || \
		echo "$(COLOR_YELLOW)   ⚠ gh / glab 未安装（可选，用于 PR/MR 操作）$(COLOR_RESET)"

# 可选：Kimi CLI（仅警告，不阻断）
	@which kimi > /dev/null 2>&1 || \
		echo "$(COLOR_YELLOW)   ⚠ kimi 未安装（可选，用于 AI Agent 任务）$(COLOR_RESET)"

# --- Node.js 依赖 ---
install-node-deps:
	@echo "$(COLOR_CYAN)📦 检查 Node.js 依赖...$(COLOR_RESET)"
	@if [ ! -d "node_modules" ] || [ "package.json" -nt "node_modules/.package-lock.json" ]; then \
		echo "   → npm install"; \
		npm install; \
	else \
		echo "$(COLOR_GREEN)   ✓ node_modules 已是最新$(COLOR_RESET)"; \
	fi

# --- Go 依赖 ---
install-go-deps:
	@echo "$(COLOR_CYAN)📦 检查 Go 依赖...$(COLOR_RESET)"
	@go mod download
	@echo "$(COLOR_GREEN)   ✓ Go modules 就绪$(COLOR_RESET)"

# -----------------------------------------------------------------------------
# 构建
# -----------------------------------------------------------------------------
build: build-css build-go
	@echo "$(COLOR_GREEN)✅ 构建完成: $(BUILD_DIR)/$(BINARY_NAME)$(COLOR_RESET)"

build-go:
	@echo "$(COLOR_CYAN)🔨 构建 Go 后端...$(COLOR_RESET)"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME).tmp ./cmd/dashboard
	@mv $(BUILD_DIR)/$(BINARY_NAME).tmp $(BUILD_DIR)/$(BINARY_NAME)

build-css:
	@echo "$(COLOR_CYAN)🎨 构建 CSS...$(COLOR_RESET)"
	@node_modules/.bin/tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT)

# -----------------------------------------------------------------------------
# 服务管理（PM2）
# -----------------------------------------------------------------------------
start: build
	@echo "$(COLOR_CYAN)🚀 启动 Dashboard...$(COLOR_RESET)"
	@mkdir -p $(LOG_DIR)/pm2
	@if pm2 describe $(BINARY_NAME) > /dev/null 2>&1; then \
		echo "$(COLOR_YELLOW)   Dashboard 已在运行，执行重启...$(COLOR_RESET)"; \
		pm2 restart $(PM2_ECOSYSTEM); \
	else \
		pm2 start $(PM2_ECOSYSTEM); \
	fi

stop:
	@echo "$(COLOR_CYAN)🛑 停止 Dashboard...$(COLOR_RESET)"
	@pm2 stop $(BINARY_NAME) 2>/dev/null || echo "$(COLOR_YELLOW)   Dashboard 未运行$(COLOR_RESET)"

restart: stop start

reload: build
	@echo "$(COLOR_CYAN)🔄 PM2 reload Dashboard...$(COLOR_RESET)"
	@pm2 reload $(BINARY_NAME) 2>/dev/null || { \
		echo "$(COLOR_YELLOW)   Dashboard 未运行，改为启动...$(COLOR_RESET)"; \
		pm2 start $(PM2_ECOSYSTEM); \
	}

# SIGHUP 热重启（零停机，旧进程优雅退出）
reload-hot:
	@echo "$(COLOR_CYAN)📡 发送 SIGHUP 热重启信号...$(COLOR_RESET)"
	@pm2 sendSignal SIGHUP $(BINARY_NAME)

status:
	@pm2 status

logs:
	@pm2 logs $(BINARY_NAME)

monit:
	@pm2 monit

# 保存 PM2 进程列表（用于开机自启）
save:
	@pm2 save

# 从 PM2 中移除并清理
delete: stop
	@echo "$(COLOR_CYAN)🗑️ 从 PM2 中删除 Dashboard...$(COLOR_RESET)"
	@pm2 delete $(BINARY_NAME) 2>/dev/null || true

# -----------------------------------------------------------------------------
# 开发模式
# -----------------------------------------------------------------------------
dev:
	@echo "$(COLOR_CYAN)🚀 开发模式启动（Go run）...$(COLOR_RESET)"
	@mkdir -p $(LOG_DIR)
	go run ./cmd/dashboard

# -----------------------------------------------------------------------------
# 清理与测试
# -----------------------------------------------------------------------------
clean:
	@echo "$(COLOR_CYAN)🧹 清理构建文件...$(COLOR_RESET)"
	rm -rf $(BUILD_DIR)
	@echo "$(COLOR_GREEN)✅ 已清理$(COLOR_RESET)"

test:
	@echo "$(COLOR_CYAN)🧪 运行测试...$(COLOR_RESET)"
	npx vitest run
	go test -v ./...

# -----------------------------------------------------------------------------
# Docker（可选）
# -----------------------------------------------------------------------------
docker-build:
	docker build -t new-world-ops:latest .

docker-run:
	docker run -p 8080:8080 -v $(PWD)/logs:/app/logs new-world-ops:latest

# -----------------------------------------------------------------------------
# 帮助
# -----------------------------------------------------------------------------
help:
	@echo ""
	@echo "$(COLOR_CYAN)New-World-Ops Dashboard — 可用命令$(COLOR_RESET)"
	@echo "============================================"
	@echo "$(COLOR_GREEN)环境初始化$(COLOR_RESET)"
	@echo "  make setup          一键完整初始化（安装依赖 + 构建）"
	@echo "  make install-deps   检测并安装所有缺失依赖"
	@echo ""
	@echo "$(COLOR_GREEN)构建$(COLOR_RESET)"
	@echo "  make build          完整构建（Go + CSS）"
	@echo "  make build-go       仅构建 Go 后端"
	@echo "  make build-css      仅构建 Tailwind CSS"
	@echo ""
	@echo "$(COLOR_GREEN)服务管理 (PM2)$(COLOR_RESET)"
	@echo "  make start          构建并启动 Dashboard"
	@echo "  make stop           停止 Dashboard"
	@echo "  make restart        完全重启（stop + start）"
	@echo "  make reload         平滑重载（重新构建后 reload）"
	@echo "  make reload-hot     SIGHUP 热重启（零停机）"
	@echo "  make status         查看 PM2 进程状态"
	@echo "  make logs           查看 Dashboard 实时日志"
	@echo "  make monit          打开 PM2 监控面板"
	@echo "  make save           保存 PM2 进程列表（开机自启）"
	@echo "  make delete         从 PM2 中移除 Dashboard"
	@echo ""
	@echo "$(COLOR_GREEN)开发$(COLOR_RESET)"
	@echo "  make dev            开发模式（go run，不经过 PM2）"
	@echo "  make test           运行 Go 测试"
	@echo "  make clean          清理构建产物"
	@echo ""
	@echo "$(COLOR_GREEN)Docker$(COLOR_RESET)"
	@echo "  make docker-build   构建 Docker 镜像"
	@echo "  make docker-run     运行 Docker 容器"
	@echo ""
