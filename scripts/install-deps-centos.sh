#!/bin/bash
# =============================================================================
# New-World Ops 依赖一键安装脚本 (CentOS/RHEL/AlmaLinux/Rocky Linux)
# =============================================================================
# 用途: 在全新的 CentOS 环境中一键安装本项目所需的所有 CLI 工具依赖
# 运行: chmod +x scripts/install-deps-centos.sh && ./scripts/install-deps-centos.sh
# =============================================================================

set -euo pipefail

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info()  { echo -e "${BLUE}[INFO]${NC}  $1"; }
log_ok()    { echo -e "${GREEN}[OK]${NC}    $1"; }
log_warn()  { echo -e "${YELLOW}[WARN]${NC}  $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# 检测 OS 版本
 detect_os() {
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        OS_NAME=$NAME
        OS_VERSION=$VERSION_ID
        OS_ID=$ID
    else
        log_error "无法检测操作系统版本"
        exit 1
    fi
    log_info "检测到操作系统: $OS_NAME $OS_VERSION"
}

# 检查命令是否存在
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# 安装基础开发工具
install_base_tools() {
    log_info "安装基础开发工具..."
    
    if command_exists dnf; then
        PKG_MGR="dnf"
    elif command_exists yum; then
        PKG_MGR="yum"
    else
        log_error "未找到 dnf 或 yum 包管理器"
        exit 1
    fi

    sudo $PKG_MGR groupinstall -y "Development Tools" || true
    sudo $PKG_MGR install -y \
        curl \
        wget \
        git \
        which \
        epel-release \
        procps-ng \
        tmux \
        python3 \
        python3-pip \
        gcc \
        make \
        ca-certificates \
        2>/dev/null || log_warn "部分基础包可能已安装"
    
    log_ok "基础工具安装完成"
}

# 安装 Go (1.21+)
install_go() {
    if command_exists go && go version | grep -q "go1.2[1-9]\|go1.[3-9][0-9]"; then
        log_ok "Go 已安装: $(go version)"
        return
    fi

    log_info "安装 Go 1.24.2 (最新稳定版)..."
    GO_VERSION="1.24.2"
    GO_TAR="go${GO_VERSION}.linux-amd64.tar.gz"
    
    cd /tmp
    wget -q "https://go.dev/dl/${GO_TAR}" -O "${GO_TAR}" || {
        log_error "下载 Go 失败，尝试国内镜像..."
        wget -q "https://mirrors.aliyun.com/golang/${GO_TAR}" -O "${GO_TAR}"
    }
    
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf "${GO_TAR}"
    rm -f "${GO_TAR}"
    
    # 添加到 PATH
    if ! grep -q "/usr/local/go/bin" ~/.bashrc; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
        echo 'export GOPATH=$HOME/go' >> ~/.bashrc
        echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
    fi
    
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    
    log_ok "Go 安装完成: $(go version)"
}

# 安装 Node.js (LTS) 和 npm
install_nodejs() {
    if command_exists node && node --version | grep -q "v2[0-9]"; then
        log_ok "Node.js 已安装: $(node --version)"
        return
    fi

    log_info "安装 Node.js LTS..."
    
    # 使用 NodeSource 官方仓库
    curl -fsSL https://rpm.nodesource.com/setup_22.x | sudo bash - || {
        log_warn "NodeSource 设置失败，尝试备用方式..."
    }
    
    if command_exists dnf; then
        sudo dnf install -y nodejs || {
            log_error "dnf 安装 nodejs 失败，尝试手动安装..."
            install_nodejs_manual
        }
    else
        sudo yum install -y nodejs || install_nodejs_manual
    fi
    
    log_ok "Node.js 安装完成: $(node --version), npm: $(npm --version)"
}

install_nodejs_manual() {
    NODE_VERSION="v22.14.0"
    NODE_TAR="node-${NODE_VERSION}-linux-x64.tar.xz"
    cd /tmp
    wget -q "https://nodejs.org/dist/${NODE_VERSION}/${NODE_TAR}"
    sudo tar -C /usr/local --strip-components=1 -xJf "${NODE_TAR}"
    rm -f "${NODE_TAR}"
}

# 安装 pnpm
install_pnpm() {
    if command_exists pnpm; then
        log_ok "pnpm 已安装: $(pnpm --version)"
        return
    fi

    log_info "安装 pnpm..."
    curl -fsSL https://get.pnpm.io/install.sh | sh -
    
    # 加载 pnpm 环境
    export PNPM_HOME="$HOME/.local/share/pnpm"
    export PATH="$PNPM_HOME:$PATH"
    
    if ! grep -q "PNPM_HOME" ~/.bashrc; then
        echo 'export PNPM_HOME="$HOME/.local/share/pnpm"' >> ~/.bashrc
        echo 'export PATH="$PNPM_HOME:$PATH"' >> ~/.bashrc
    fi
    
    log_ok "pnpm 安装完成: $(pnpm --version)"
}

# 安装 GitHub CLI
install_gh() {
    if command_exists gh; then
        log_ok "GitHub CLI 已安装: $(gh --version | head -1)"
        return
    fi

    log_info "安装 GitHub CLI..."
    
    if command_exists dnf; then
        sudo dnf config-manager --add-repo https://cli.github.com/packages/rpm/gh-cli.repo 2>/dev/null || \
        sudo dnf config-manager --add-repo https://mirror.ghproxy.com/https://cli.github.com/packages/rpm/gh-cli.repo
        sudo dnf install -y gh
    else
        # 手动安装
        GH_VERSION=$(curl -s https://api.github.com/repos/cli/cli/releases/latest | grep '"tag_name"' | sed -E 's/.*"v([^"]+)".*/\1/')
        GH_TAR="gh_${GH_VERSION}_linux_amd64.tar.gz"
        cd /tmp
        wget -q "https://github.com/cli/cli/releases/download/v${GH_VERSION}/${GH_TAR}"
        tar -xzf "${GH_TAR}"
        sudo cp "gh_${GH_VERSION}_linux_amd64/bin/gh" /usr/local/bin/
        rm -rf "gh_${GH_VERSION}_linux_amd64" "${GH_TAR}"
    fi
    
    log_ok "GitHub CLI 安装完成"
}

# 安装 GitLab CLI (glab)
install_glab() {
    if command_exists glab; then
        log_ok "GitLab CLI 已安装"
        return
    fi

    log_info "安装 GitLab CLI (glab)..."
    
    # 方式1：通过官方安装脚本
    curl -sL https://jfgreffier.dev/glab-installer.sh | sudo bash - || {
        log_warn "官方安装脚本失败，尝试手动安装..."
        # 方式2：手动下载最新 release
        GLAB_VERSION=$(curl -s https://gitlab.com/api/v4/projects/34675721/releases | grep -o '"tag_name":"[^"]*"' | head -1 | sed 's/"tag_name":"//;s/"$//')
        if [ -z "$GLAB_VERSION" ]; then
            GLAB_VERSION="v1.53.0"
        fi
        GLAB_TAR="glab_${GLAB_VERSION#v}_Linux_x86_64.tar.gz"
        cd /tmp
        wget -q "https://gitlab.com/gitlab-org/cli/-/releases/${GLAB_VERSION}/downloads/${GLAB_TAR}" || \
        wget -q "https://github.com/profclems/glab/releases/download/${GLAB_VERSION}/${GLAB_TAR}"
        tar -xzf "${GLAB_TAR}" 2>/dev/null || true
        if [ -f "bin/glab" ]; then
            sudo mv bin/glab /usr/local/bin/
        elif [ -f "glab" ]; then
            sudo mv glab /usr/local/bin/
        fi
        rm -rf bin glab "${GLAB_TAR}"
    }
    
    if command_exists glab; then
        log_ok "GitLab CLI (glab) 安装完成"
    else
        log_warn "glab 安装可能未完成，如需 GitLab 支持请手动安装"
    fi
}

# 安装 Kimi CLI
install_kimi() {
    if command_exists kimi; then
        log_ok "Kimi CLI 已安装: $(kimi --version 2>/dev/null || echo 'installed')"
        return
    fi

    log_info "安装 Kimi CLI..."
    log_warn "Kimi CLI 需要 pip 安装，请确保有 Python 3.10+ 环境"
    
    # 尝试通过 uv 安装（推荐方式）
    if command_exists uv; then
        uv tool install kimi-cli || true
    elif command_exists pip3; then
        pip3 install --user kimi-cli || {
            log_warn "pip 安装失败，尝试使用 uv 安装..."
            install_uv
            uv tool install kimi-cli
        }
    else
        install_uv
        uv tool install kimi-cli
    fi
    
    if command_exists kimi; then
        log_ok "Kimi CLI 安装完成"
    else
        log_warn "Kimi CLI 安装可能未完成，请手动运行: pip3 install kimi-cli 或 uv tool install kimi-cli"
    fi
}

# 安装 uv (Python 包管理器)
install_uv() {
    if command_exists uv; then
        return
    fi
    
    log_info "安装 uv (Python 包管理器)..."
    curl -LsSf https://astral.sh/uv/install.sh | sh
    export PATH="$HOME/.cargo/bin:$PATH"
    if ! grep -q ".cargo/bin" ~/.bashrc; then
        echo 'export PATH="$HOME/.cargo/bin:$PATH"' >> ~/.bashrc
    fi
    log_ok "uv 安装完成"
}

# 安装 pm2
install_pm2() {
    if command_exists pm2; then
        log_ok "pm2 已安装: $(pm2 --version)"
        return
    fi

    log_info "安装 pm2..."
    npm install -g pm2
    log_ok "pm2 安装完成"
}

# 安装 lark-cli (飞书 CLI)
install_lark_cli() {
    if command_exists lark-cli; then
        log_ok "lark-cli 已安装"
        return
    fi

    log_info "安装 lark-cli (飞书 CLI)..."
    
    if command_exists pip3; then
        pip3 install --user lark-cli || {
            log_warn "pip 安装 lark-cli 失败，尝试 uv..."
            install_uv
            uv tool install lark-cli || true
        }
    else
        install_uv
        uv tool install lark-cli || true
    fi
    
    if command_exists lark-cli; then
        log_ok "lark-cli 安装完成"
    else
        log_warn "lark-cli 安装可能未完成。如需飞书通知功能，请手动安装: pip3 install lark-cli"
    fi
}

# 安装项目 npm 依赖 (Tailwind CSS)
install_project_npm_deps() {
    log_info "安装项目 npm 依赖 (Tailwind CSS)..."
    
    if [ -f "package.json" ]; then
        npm install
        log_ok "项目 npm 依赖安装完成"
    else
        log_warn "未找到 package.json，跳过项目 npm 依赖安装"
    fi
}

# 验证所有依赖
verify_deps() {
    log_info "验证所有依赖..."
    echo ""
    
    local all_ok=true
    local tools=("go" "git" "gh" "node" "npm" "tmux" "python3")
    
    printf "  ${GREEN}✓${NC} %-12s %s\n" "glab" "$(glab --version 2>/dev/null | head -1 || echo 'optional')" "${YELLOW}(可选)${NC}"
    
    for tool in "${tools[@]}"; do
        if command_exists "$tool"; then
            local version
            version=$($tool --version 2>/dev/null | head -1 || echo "installed")
            printf "  ${GREEN}✓${NC} %-12s %s\n" "$tool" "$version"
        else
            printf "  ${RED}✗${NC} %-12s ${RED}未安装${NC}\n" "$tool"
            all_ok=false
        fi
    done
    
    # 可选工具
    local optional_tools=("kimi" "pnpm" "pm2" "lark-cli")
    for tool in "${optional_tools[@]}"; do
        if command_exists "$tool"; then
            local version
            version=$($tool --version 2>/dev/null | head -1 || echo "installed")
            printf "  ${GREEN}✓${NC} %-12s %s ${YELLOW}(可选)${NC}\n" "$tool" "$version"
        else
            printf "  ${YELLOW}○${NC} %-12s ${YELLOW}未安装 (可选)${NC}\n" "$tool"
        fi
    done
    
    echo ""
    if $all_ok; then
        log_ok "核心依赖全部就绪！"
    else
        log_error "部分核心依赖缺失，请检查上方输出"
        return 1
    fi
}

# 主函数
main() {
    echo "======================================================================"
    echo "  New-World Ops 依赖安装脚本"
    echo "======================================================================"
    echo ""
    
    detect_os
    
    # 确认安装
    echo ""
    log_warn "此脚本将安装以下工具："
    echo "  • 基础: git, tmux, python3, gcc, make"
    echo "  • Go 1.21+ (当前安装 1.24.2)"
    echo "  • Node.js LTS + npm"
    echo "  • pnpm (Node 包管理器)"
    echo "  • GitHub CLI (gh)"
    echo "  • Kimi CLI (kimi)"
    echo "  • pm2 (进程管理器)"
    echo "  • lark-cli (飞书 CLI, 可选)"
    echo "  • 项目 npm 依赖 (Tailwind CSS)"
    echo ""
    read -p "是否继续? [Y/n]: " confirm
    if [[ "$confirm" =~ ^[Nn]$ ]]; then
        log_info "已取消安装"
        exit 0
    fi
    
    echo ""
    log_info "开始安装..."
    echo ""
    
    install_base_tools
    install_go
    install_nodejs
    install_pnpm
    install_gh
    install_glab
    install_kimi
    install_pm2
    install_lark_cli
    install_project_npm_deps
    
    echo ""
    echo "======================================================================"
    verify_deps
    echo "======================================================================"
    
    echo ""
    log_info "安装完成！请执行以下操作："
    echo ""
    echo "  1. 重新加载环境变量:"
    echo "     source ~/.bashrc"
    echo ""
    echo "  2. 配置 GitHub CLI 认证:"
    echo "     gh auth login"
    echo ""
    echo "  3. 配置 GitLab CLI 认证（如使用 GitLab）:"
    echo "     glab auth login"
    echo ""
    echo "  4. 配置 Kimi CLI 认证:"
    echo "     kimi auth login"
    echo ""
    echo "  5. (可选) 配置 lark-cli 飞书认证:"
    echo "     lark-cli config init"
    echo "     lark-cli auth login"
    echo ""
    echo "  5. 验证并构建项目:"
    echo "     make install-deps"
    echo "     make build"
    echo ""
}

main "$@"
