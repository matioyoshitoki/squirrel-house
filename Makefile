.PHONY: all build run clean test build-css

BINARY_NAME=dashboard
BUILD_DIR=./build
LOG_DIR=./logs

all: build

build:
	@echo "🔨 构建 Dashboard..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME).tmp ./cmd/dashboard
	@mv $(BUILD_DIR)/$(BINARY_NAME).tmp $(BUILD_DIR)/$(BINARY_NAME)

run: build
	@echo "🚀 启动 Dashboard..."
	@mkdir -p $(LOG_DIR)
	$(BUILD_DIR)/$(BINARY_NAME)

dev:
	@echo "🚀 以开发模式启动..."
	@mkdir -p $(LOG_DIR)
	go run ./cmd/dashboard

clean:
	@echo "🧹 清理构建文件..."
	rm -rf $(BUILD_DIR)

test:
	@echo "🧪 运行测试..."
	go test -v ./...

install-deps:
	@echo "📦 检查依赖..."
	@which gh > /dev/null || which glab > /dev/null || (echo "❌ 请先安装 GitHub CLI (gh) 或 GitLab CLI (glab)" && exit 1)
	@which kimi > /dev/null || (echo "❌ 请先安装 Kimi CLI" && exit 1)
	@echo "✅ 所有依赖已安装"

# 使用 pm2 启动（如果未启动）
run: build
	@mkdir -p $(LOG_DIR)
	@if pm2 describe dashboard >/dev/null 2>&1; then \
		echo "🚀 pm2 restart dashboard..."; \
		pm2 restart dashboard; \
	else \
		echo "🚀 pm2 start dashboard..."; \
		pm2 start ecosystem.config.js; \
	fi

# 停止 dashboard
stop:
	@echo "🛑 停止 Dashboard..."
	@pm2 stop dashboard 2>/dev/null || true

# pm2 重启（推荐：先 kill 旧进程，再启动新进程，可靠无僵尸）
reload: build
	@echo "🔄 pm2 restart dashboard..."
	@pm2 restart dashboard

# 保留 SIGHUP 热重启（零停机，但可能有僵尸进程风险）
reload-hot:
	@echo "📡 发送 SIGHUP 给 dashboard..."
	@pm2 sendSignal SIGHUP dashboard

# 普通重启（kill + 重新启动）
restart: stop run

# pm2 状态
status:
	@pm2 status

# pm2 日志
logs:
	@pm2 logs dashboard

# 从 pm2 中删除 dashboard
delete:
	@echo "🗑️ 从 pm2 中删除 dashboard..."
	@pm2 delete dashboard 2>/dev/null || true

# 保存 pm2 配置（开机自启）
save:
	@pm2 save

# pm2 监控面板
monit:
	@pm2 monit

# 构建 CSS（Tailwind v4 CLI）
build-css:
	@echo "🎨 构建 CSS..."
	@node_modules/.bin/tailwindcss -i cmd/dashboard/static/input.css -o cmd/dashboard/static/style.css

# Docker 支持
docker-build:
	docker build -t new-world-ops:latest .

docker-run:
	docker run -p 8080:8080 -v $(PWD)/logs:/app/logs new-world-ops:latest
