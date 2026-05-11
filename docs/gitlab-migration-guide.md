# GitHub → GitLab 迁移改造清单

> 本文档列出将 new-world-ops 从 GitHub 迁移到 GitLab 所需的全部改造点。

---

## 一、概念映射对照表

| GitHub 概念 | GitLab 概念 | 备注 |
|------------|------------|------|
| PR (Pull Request) | MR (Merge Request) | 核心概念变化，需全局重命名 |
| Issue | Issue | 概念相同，API 不同 |
| PR Comment | MR Note / Comment | GitLab 叫 "note" 或 "discussion" |
| Label | Label | 概念相同 |
| `gh` CLI | `glab` CLI | GitLab 官方 CLI 工具 |
| `owner/repo` | `group/project` 或 `namespace/project` | repo 路径格式变化 |
| PR Diff | MR Changes | `glab mr diff` |

---

## 二、CLI 工具层改造

### 2.1 替换 `gh` → `glab`

**当前：** 所有 Git 平台操作通过 `gh` (GitHub CLI) 执行。

**改造：** 全局替换为 `glab` (GitLab CLI)。

```bash
# 安装 glab
# macOS
brew install glab

# CentOS/RHEL
curl -sL https://jfgreffier.dev/glab-installer.sh | sudo bash
# 或下载二进制: https://gitlab.com/gitlab-org/cli/-/releases
```

### 2.2 命令映射速查

| 功能 | GitHub (`gh`) | GitLab (`glab`) |
|------|--------------|-----------------|
| 认证状态 | `gh auth status` | `glab auth status` |
| 列出 PR/MR | `gh pr list --state open --json ...` | `glab mr list --state opened -F json` |
| 查看 PR/MR | `gh pr view <n> --json body` | `glab mr view <n> -F json` |
| PR/MR Diff | `gh pr diff <n> --name-only` | `glab mr diff <n>` (需配合 grep) |
| 创建评论 | `gh pr comment <n> --body "..."` | `glab mr note <n> -m "..."` |
| 编辑标签 | `gh pr edit <n> --add-label x` | `glab mr update <n> -l x` |
| 合并 | `gh pr merge <n> --squash` | `glab mr merge <n> -s` |
| 更新分支 | `gh pr update-branch <n>` | `glab mr rebase <n>` 或手动 rebase |
| 创建 PR/MR | `gh pr create --base x --title y` | `glab mr create -b x -t y` |
| 查看 Issue | `gh issue view <n> --json ...` | `glab issue view <n> -F json` |
| 创建 Issue | `gh issue create --title x --body y` | `glab issue create -t x -d y` |

**关键差异：**
- `gh` 的 `--json` 输出 vs `glab` 的 `-F json` 输出，JSON 结构完全不同，需重新解析
- `gh pr diff --name-only` 需替换为 `glab mr diff <n> | grep ...` 或 `git diff --name-only origin/main...HEAD`
- `gh pr comment --edit-last` 在 glab 中没有直接等价物，需手动查询最新 note 再编辑

---

## 三、Go 代码层改造

### 3.1 核心调用点清单

所有 `exec.Command("gh", ...)` 的调用点都需要改造，分布如下：

#### `cmd/dashboard/handlers_pr.go`（最密集，约 15 处）

| 行号范围 | 当前命令 | 改造说明 |
|---------|---------|---------|
| ~86 | `gh pr list --state open --json ...` | 改为 `glab mr list --state opened -F json`，并重写 JSON 解析 |
| ~177 | `gh pr view --json body` | 改为 `glab mr view -F json`，`Body` 字段可能叫 `Description` |
| ~375 | `gh pr view --json title,body,headRefName` | 同上，字段映射变化 |
| ~394 | `gh issue view --json title` | 改为 `glab issue view -F json` |
| ~615 | `gh pr merge --squash --delete-branch` | 改为 `glab mr merge -s -d` |
| ~622 | `gh pr update-branch` | 改为 `glab mr rebase` 或手动实现 |
| ~761 | `gh pr view --json title,body,baseRefName,headRefName` | 改为 `glab mr view -F json` |
| ~776 | `gh pr diff --name-only` | 改为 `glab mr diff` 或 `git diff` 方式 |
| ~842 | `gh pr view --json headRefName` | 改为 `glab mr view -F json` |
| ~850 | `gh pr diff --name-only` | 同上 |
| ~928 | `gh pr comment --edit-last` | **无直接等价**，需先 `glab mr note list` 获取最新 note ID，再 `glab mr note update` |
| ~938 | `gh pr comment --body-file` | 改为 `glab mr note -m "$(cat file)"` |

#### `cmd/dashboard/handlers_dev.go`（约 3 处）

| 行号范围 | 当前命令 | 改造说明 |
|---------|---------|---------|
| ~198 | `gh issue view --json ...` | 改为 `glab issue view -F json` |
| ~332 | `prFinderImpl.findPR(branch)` → 内部调用 `gh pr list --head` | 改为 `glab mr list --source-branch` |
| ~543 | `gh pr view --json headRefName` | 改为 `glab mr view -F json` |

#### `cmd/dashboard/handlers_pm.go`（1 处）

| 行号范围 | 当前命令 | 改造说明 |
|---------|---------|---------|
| ~50 | `gh issue create --label ... --title ... --body ...` | 改为 `glab issue create -l ... -t ... -d ...` |

#### `cmd/dashboard/hooks.go`（1 处）

| 行号范围 | 当前命令 | 改造说明 |
|---------|---------|---------|
| ~341 | `gh pr list --head <branch> --json number` | 改为 `glab mr list --source-branch <branch> -F json` |

#### `cmd/dashboard/notify.go`（1 处）

| 行号范围 | 当前命令 | 改造说明 |
|---------|---------|---------|
| ~161 | `gh pr comment --body-file` | 改为 `glab mr note -m "$(cat file)"` |

### 3.2 数据结构改造

#### `PullRequest` struct (`handlers_pr.go:18`)

当前字段映射的是 GitHub API 返回的 JSON 字段名。GitLab 的 JSON 结构完全不同：

```go
// GitHub 版本（当前）
type PullRequest struct {
    Number    int    `json:"number"`
    Title     string `json:"title"`
    State     string `json:"state"`
    Branch    string `json:"branch"`        // 运行时填充
    BaseRef   string `json:"baseRefName"`   // GitHub 字段
    HeadRef   string `json:"headRefName"`   // GitHub 字段
    URL       string `json:"url"`
    CreatedAt string `json:"createdAt"`
    Author    struct {
        Login string `json:"login"`
    } `json:"author"`
}

// GitLab 版本（需改造）
type MergeRequest struct {
    IID        int    `json:"iid"`           // GitLab 用 iid，不是 id
    Title      string `json:"title"`
    State      string `json:"state"`
    Branch     string `json:"branch"`        // 运行时填充
    SourceBranch string `json:"source_branch"`
    TargetBranch string `json:"target_branch"`
    WebURL     string `json:"web_url"`       // 不是 url
    CreatedAt  string `json:"created_at"`    // 下划线命名
    Author     struct {
        Username string `json:"username"`     // 不是 login
    } `json:"author"`
}
```

**注意：** GitLab 的 MR 编号字段是 `iid`（project-level ID），不是全局 `id`。所有 `strconv.Itoa(prNumber)` 的地方需要确认是否使用 `iid`。

#### Issue struct

当前代码中 Issue 的解析分布在多处，GitLab Issue JSON 结构也需要适配：
- `number` → `iid`
- `title` → `title`（相同）
- `state` → `state`（相同）
- `labels` → `labels`（GitLab 是 `[]string`）
- `body` → `description`
- `updatedAt` → `updated_at`

### 3.3 新增抽象层（推荐方案）

直接在每处调用中硬编码 `glab` 命令会导致未来再次迁移时同样的麻烦。**建议**在 `cmd/dashboard/` 下新增一个平台抽象层：

```go
// cmd/dashboard/platform.go
package main

type Platform interface {
    ListMRs(state string) ([]MergeRequest, error)
    ViewMR(iid int) (*MergeRequest, error)
    CreateMRComment(iid int, body string) error
    EditMRComment(iid int, body string) error   // 或 UpdateLastMRComment
    GetMRDiffFiles(iid int) ([]string, error)
    MergeMR(iid int, squash bool) error
    RebaseMR(iid int) error
    CreateMR(sourceBranch, targetBranch, title, body string) error
    
    ListIssues(state string) ([]Issue, error)
    ViewIssue(iid int) (*Issue, error)
    CreateIssue(title, body string, labels []string) (int, error)
    
    FindMRByBranch(branch string) (int, error)
    AuthStatus() error
}

type GitHubPlatform struct{ projectPath string }
type GitLabPlatform struct{ projectPath string }

func NewPlatform(platformType, projectPath string) Platform {
    switch platformType {
    case "gitlab":
        return &GitLabPlatform{projectPath: projectPath}
    default:
        return &GitHubPlatform{projectPath: projectPath}
    }
}
```

然后在 `Project` 配置中新增 `Platform string `json:"platform"`` 字段（默认 `"github"`），所有 `exec.Command("gh", ...)` 的调用替换为 `platform.Xxx()` 方法调用。

---

## 四、Workflow YAML 层改造

所有 `workflows/*.yaml` 中硬编码的 `gh` 命令都需要替换。注意 workflow 通过模板变量执行命令，改造需同时修改 Go 代码中的变量注入。

### 4.1 `workflows/review.yaml`

```yaml
# 改造前（review.yaml:29）
command: |
  gh pr diff {{ .Vars.prNumber }} --name-only | grep -iE '\.(png|jpg|jpeg|svg|webp|gif)$' > .design-assets.txt 2>/dev/null || true

# 改造后
command: |
  glab mr diff {{ .Vars.prNumber }} | grep -E '^diff --git' | sed 's|diff --git a/||;s| b/.*||' | grep -iE '\.(png|jpg|jpeg|svg|webp|gif)$' > .design-assets.txt 2>/dev/null || true
  # 或者更简单的：在 worktree 中直接 git diff
  git diff --name-only origin/{{ .Vars.branch }}...HEAD | grep -iE '\.(png|jpg|jpeg|svg|webp|gif)$' > .design-assets.txt 2>/dev/null || true
```

```yaml
# 改造前（review.yaml:34）
command: |
  links=$(gh pr view {{ .Vars.prNumber }} --json body -q '.body' | grep -oE 'https?://...')

# 改造后
command: |
  links=$(glab mr view {{ .Vars.prNumber }} -F json | jq -r '.description' | grep -oE 'https?://...')
```

```yaml
# 改造前（review.yaml:45）
command: gh auth status >/dev/null 2>&1 || (echo "❌ GitHub CLI 未认证...")

# 改造后
command: glab auth status >/dev/null 2>&1 || (echo "❌ GitLab CLI 未认证...")
```

```yaml
# 改造前（review.yaml:54-57）
gh pr diff {{ .Vars.prNumber }} --name-only

# 改造后
glab mr diff {{ .Vars.prNumber }} | grep -E '^diff --git' | sed 's|diff --git a/||;s| b/.*||'
# 或
git diff --name-only origin/{{ .Vars.branch }}...HEAD
```

### 4.2 `workflows/rework.yaml`

```yaml
# 改造前（rework.yaml:6）
command: gh pr comment {{ .Vars.prNumber }} --body {{ .Vars.commentBody | shellquote }}

# 改造后
command: glab mr note {{ .Vars.prNumber }} -m {{ .Vars.commentBody | shellquote }}
```

```yaml
# 改造前（rework.yaml:12）
command: gh pr edit {{ .Vars.prNumber }} --add-label needs-work

# 改造后
command: glab mr update {{ .Vars.prNumber }} -l needs-work
```

### 4.3 `workflows/design.yaml`

```yaml
# 改造前（design.yaml:140）
gh pr create --base {{ .Vars.baseBranch }} --head {{ .Vars.branch }} --title "..." --body "$body"

# 改造后
glab mr create -b {{ .Vars.baseBranch }} -s {{ .Vars.branch }} -t "..." -d "$body"
```

### 4.4 `workflows/pm.yaml`

PM workflow 本身没有直接 `gh` 命令，但 prompt 中会指示 agent 执行 `gh pr create`。需要在 prompt 模板中同步修改。

### 4.5 项目级 workflow 覆盖

`workflows/new-world-monorepo/` 和 `workflows/social-world/` 下的 project-specific workflow 也需要同步检查修改。

---

## 五、Prompt 模板层改造

### 5.1 `prompts/reviewer.md`

当前 reviewer prompt 中明确指示 agent 使用 `gh` 命令：

```markdown
1. **环境预检查**：运行任何 `gh` 命令前，先执行 `gh auth status` 确认 GitHub CLI 已认证...
```

需要替换为：
```markdown
1. **环境预检查**：运行任何 `glab` 命令前，先执行 `glab auth status` 确认 GitLab CLI 已认证...
```

同时 prompt 中所有 `gh pr view`、`gh pr diff`、`gh pr comment` 等示例都需要替换为 `glab mr view`、`glab mr diff`、`glab mr note`。

### 5.2 `prompts/product-manager.md`

```markdown
# 改造前
根据用户输入，判断是否需要创建 GitHub Issue...

# 改造后
根据用户输入，判断是否需要创建 GitLab Issue...
```

### 5.3 其他 prompt 文件

需要全局搜索所有 prompts/ 目录下的 markdown 文件，替换：
- `gh` → `glab`
- `GitHub` → `GitLab`
- `PR` → `MR` (在命令上下文中)
- `pull request` → `merge request`
- `gh pr create` → `glab mr create`

---

## 六、配置层改造

### 6.1 `configs/dashboard.json`

```json
// 改造前
{
  "repo": "MatioYoshitoki/new-world"
}

// 改造后
{
  "repo": "group/subgroup/new-world",
  "platform": "gitlab"
}
```

GitLab 的 repo 路径可能是多层级的（`group/subgroup/project`），当前代码中解析 `repo` 字段的地方需要支持斜杠分割的多级路径。

### 6.2 Project struct (`cmd/dashboard/main.go`)

新增 `Platform` 字段：

```go
type Project struct {
    Name          string     `json:"name"`
    Path          string     `json:"path"`
    Description   string     `json:"description"`
    DefaultBranch string     `json:"defaultBranch"`
    Domain        string     `json:"domain"`
    TechStack     TechStack  `json:"techStack"`
    Repo          string     `json:"repo"`
    Platform      string     `json:"platform"`  // 新增: "github" | "gitlab"
}
```

---

## 七、前端/UI 层改造

### 7.1 `cmd/dashboard/static/index.html`

检查是否有硬编码的 GitHub 相关 URL 或文案：
- "GitHub" → "GitLab"
- "PR" → "MR"（界面文案中）
- GitHub 图标/logo → GitLab 图标

### 7.2 `cmd/dashboard/static/js/terminal.js`

搜索是否有 GitHub 相关的文本替换或解析逻辑。

---

## 八、安装脚本改造

### 8.1 `scripts/install-deps-centos.sh`

```bash
# 新增 glab 安装段（替代 gh）
install_glab() {
    if command_exists glab; then
        log_ok "GitLab CLI 已安装"
        return
    fi
    log_info "安装 GitLab CLI (glab)..."
    # 下载最新 release
    GLAB_VERSION=$(curl -s https://gitlab.com/api/v4/projects/34675721/releases | grep -o '"tag_name":"[^"]*"' | head -1 | sed 's/"tag_name":"//;s/"$//')
    curl -L "https://gitlab.com/gitlab-org/cli/-/releases/${GLAB_VERSION}/downloads/glab_${GLAB_VERSION#v}_Linux_x86_64.tar.gz" | tar -xz -C /tmp
    sudo mv /tmp/bin/glab /usr/local/bin/
    log_ok "glab 安装完成"
}
```

---

## 九、改造优先级建议

按**依赖关系**和**影响范围**排序：

| 优先级 | 改造项 | 预估工作量 | 风险 |
|-------|--------|----------|------|
| P0 | 新增 Platform 抽象接口 | 2-3h | 低（纯新增） |
| P0 | Go 代码所有 `gh` 调用改为抽象方法 | 4-6h | 中（逻辑密集） |
| P0 | `PullRequest` / `Issue` struct 适配 | 1-2h | 中（JSON 解析） |
| P1 | Workflow YAML 中的 `gh` → `glab` | 2-3h | 低（文本替换） |
| P1 | Prompt 模板中的 `gh` → `glab` | 1-2h | 低（文本替换） |
| P1 | `configs/dashboard.json` 加 `platform` 字段 | 0.5h | 低 |
| P2 | 前端 UI 文案/图标替换 | 1-2h | 低 |
| P2 | 安装脚本增加 glab | 0.5h | 低 |
| P2 | README 文档更新 | 0.5h | 低 |

**推荐路径：** 先实现 P0 的 Platform 抽象层（这是架构改进，未来再迁移到其他平台也不用改业务代码），再逐层改造外围。

---

## 十、关键风险点

1. **`glab` JSON 输出与 `gh` 完全不同**：所有 `json.Unmarshal` 的地方都需要重新对照 glab 的实际输出字段做适配，不能简单字符串替换。

2. **`gh pr comment --edit-last` 无直接等价物**：这是 GitLab CLI 的缺失功能，需要先用 `glab mr note list` 获取最新 note ID，再 `glab mr note update <note-id>`。或者改用直接调用 GitLab REST API。

3. **`gh pr diff --name-only` 在 glab 中没有**：需要改用 `git diff --name-only origin/target...HEAD` 在 worktree 中实现，或者解析 `glab mr diff` 的 patch 格式。

4. **GitLab Issue/MR 编号**：GitLab 同时有全局 `id` 和项目级 `iid`。API 中 MR 的 URL 编号是 `iid`，但某些 API 可能需要 `id`。需确保代码中一直使用 `iid`。

5. **Webhook / CI 触发**：如果目前依赖 GitHub Actions webhook，需要改为 GitLab CI webhook，URL 和 payload 结构完全不同。
