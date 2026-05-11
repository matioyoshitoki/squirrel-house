# New-World Ops

New-World 项目的运维工具集，包含 Agent Dashboard 用于任务管理和代码审查。

## 架构

Dashboard 通过 **YAML 工作流引擎** (`pkg/workflow`) 编排任务，Agent 配置和 Prompt 模板已统一集中到本仓库：

```
new-world-ops/
├── agents/              ← 集中式 Agent YAML 配置
├── prompts/             ← Go template 化的 Prompt 模板
├── workflows/           ← 工作流定义（review.yaml / dev.yaml）
├── pkg/workflow/        ← YAML 工作流执行引擎
├── cmd/dashboard/       ← HTTP 服务 + 任务调度
│   ├── handlers_dev.go
│   ├── handlers_pr.go
│   ├── agent.go         # 动态渲染 agent 配置
│   └── static/
├── configs/
│   └── dashboard.json   # 项目配置（路径、分支、技术栈）
└── logs/                # 任务日志与审查报告
```

**关键分离点：**
1. **Agent 配置集中化**：`agents/` 和 `prompts/` 位于 `new-world-ops` 中，使用 Go `text/template` 根据项目变量（`domain`、`techStack`、`defaultBranch`）动态渲染
2. **Ops 工具独立**：Dashboard 只负责启动 agent 和展示结果，不包含任何项目特定的业务逻辑
3. **通过配置关联**：`configs/dashboard.json` 配置项目路径、分支、领域和技术栈
4. **不侵入项目代码**：Ops 通过 Platform 抽象层操作 PR/MR（自动支持 GitHub / GitLab），不引用被管理项目的任何 internal/ 包
5. **平台自动推断**：根据项目 git remote URL 自动识别 GitHub 或 GitLab，无需手动配置

## 快速开始

### 前置要求

- Go 1.21+
- GitHub CLI (`gh`) **或** GitLab CLI (`glab`) —— 系统会根据项目 git remote URL 自动推断使用哪个平台
- Kimi CLI (`kimi`)

### 安装

```bash
# 克隆仓库
cd ~/Desktop/new-world/new-world-ops

# 检查依赖
make install-deps

# 配置项目路径
vim configs/dashboard.json
```

### 运行

```bash
# 开发模式（推荐）
make dev

# 或构建后运行
make build
make run
```

访问 http://localhost:8080

## 配置

### Dashboard 配置

编辑 `configs/dashboard.json`：

```json
{
  "port": "8080",
  "projects": [
    {
      "name": "new-world-monorepo",
      "path": "../../new-world-monorepo",
      "defaultBranch": "dev/2.1.0",
      "domain": "game-city-builder",
      "techStack": {
        "backend": "nestjs",
        "frontend": "phaser3",
        "db": "postgres+typeorm",
        "packageManager": "pnpm"
      },
      "repo": "MatioYoshitoki/new-world"
    }
  ]
}
```

### Agent 与 Prompt 配置

Agent 配置已统一集中到 `new-world-ops/agents/` 和 `new-world-ops/prompts/`。Dashboard 启动任务时会：

1. 读取 `prompts/{agentType}.md` 模板
2. 注入项目变量（`Name`、`Domain`、`TechStack`、`DefaultBranch`）渲染出最终 prompt
3. 生成临时 agent YAML 文件供 `kimi --agent-file` 使用

无需在每个项目中维护 `.kimi/agents/` 和 `.kimi/prompts/`。

### UI 设计资产规范

项目引入了独立的 **UI 设计 Agent (`ui-designer`)**，为每个 Issue 产出标准设计资产：

```
designs/issue-{N}/
├── wireframe.svg      # 线框图
├── mockup.html        # 高保真原型
└── design-spec.md     # 设计规范文档
```

- **Harness 层规范**：[`docs/ui-design-spec.md`](docs/ui-design-spec.md) 定义了设计资产的统一目录结构和文件格式（接口契约）
- **Project 层规范**：Agent 启动后会优先读取被管理项目自身的设计系统文档（如 `docs/design-system/README.md`），项目特定的色板、组件、状态管理方案等应放在那里

这符合 Harness 架构原则：Ops 工具保持通用薄层，项目特定知识下沉到项目自身文档中。

## 功能

### 1. Issue 开发任务
- 查看带标签的 Issues
- 启动 Agent 开发任务（Web 界面）
- 实时查看日志（SSE 流）
- 支持任务中断恢复

**Web 界面启动开发：**
访问 http://localhost:8080，选择项目后点击对应 Issue 的「开始开发」按钮。

### 2. PR Review
- 查看打开的 PRs
- 启动 Agent 代码审查
- 生成结构化审查报告
- 自动添加 PR 评论
- 标记返工/合并 PR

**Web 界面启动审查：**
访问 http://localhost:8080，选择项目后点击对应 PR 的「Review」按钮。

### 3. 返工修复
- 审查未通过的 PR 自动标记为返工
- **使用专门的返工 Agent（rework-agent）** 自动修复问题
- 实时查看修复进度
- 修复完成后自动推送代码并添加评论
- 支持手动触发返工修复

**返工 Agent 与新开发 Agent 的区别：**

| 对比项 | 新开发 (developer) | 返工修复 (rework) |
|--------|-------------------|-------------------|
| **目标** | 从零实现功能 | 修复审查发现的问题 |
| **输入** | PRD/需求文档 | 审查报告 (review-report) |
| **流程** | 设计 → 实现 → 测试 → PR | 阅读报告 → 逐项修复 → 验证 |
| **分支** | 创建新分支 | 在现有 PR 分支上修复 |
| **提交** | 多个功能 commit | 一个 fix commit |

**Web 界面启动返工修复：**
在 PR Review 结果为 "needs_fix" 时，点击「开始返工」按钮触发修复任务。

## 日志

所有日志保存在 `logs/` 目录：

- `dev-issue-{number}-{timestamp}.log` - 开发任务日志
- `review-pr-{number}-{timestamp}.log` - 审查任务日志
- `review-report-{number}.md` - 审查报告

## 开发

### 添加新的 API 端点

1. 在 `handlers_*.go` 中添加处理器函数
2. 在 `main.go` 的 `main()` 中注册路由

### 修改前端

编辑 `cmd/dashboard/static/index.html`，刷新页面即可生效。

## 部署

### 独立部署

```bash
make build
./build/dashboard
```

### Docker

```bash
make docker-build
make docker-run
```

## 迁移记录

**Agent 配置集中化：**
- `social-world/.kimi/` 和 `new-world-monorepo/.kimi/` 中的 agent/prompt 文件已迁移到 `new-world-ops/agents/` 和 `new-world-ops/prompts/`
- Dashboard 通过 `prepareAgentConfig()` 动态渲染临时 agent YAML，不再依赖项目目录下的 `.kimi/`

**工作流引擎：**
- 新增 `pkg/workflow/` YAML 工作流引擎，逐步替代 `handlers_pr.go` / `handlers_dev.go` 中的硬编码逻辑
- `workflows/review.yaml` 已接入 `handleReviewPR`
- `workflows/dev.yaml` 已定义，待接入 `handlers_dev.go`

**改进点：**
- 前后端分离（HTML 独立为 static 文件）
- 配置文件化（支持 JSON 配置）
- 代码模块化（按功能分文件）
- 与原项目完全解耦
- Agent 配置统一维护，避免多项目重复
