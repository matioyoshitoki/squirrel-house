#!/bin/bash
# 启动 social-world 项目的 Product Manager Agent
# 启动后直接在终端中与 Agent 交互，输入需求即可

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
AGENT_FILE="$SCRIPT_DIR/../agents/product-manager.yaml"
PROJECT_DIR="/Users/insulate/Desktop/social-world"

cd "$PROJECT_DIR"

echo "🚀 启动 social-world Product Manager Agent"
echo "📁 项目目录: $PROJECT_DIR"
echo ""

kimi --agent-file "$AGENT_FILE"
