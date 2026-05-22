# Prompt 演进报告 — type-one

生成时间: 2026-05-22T15:49:48+08:00
分析范围: 2026-05-22T14:41:37+08:00 至 2026-05-22T15:43:56+08:00

## 统计摘要

- 分析日志数: 19
- Agent 类型分布:
  - review: 9 个任务，平均 7 步，成功率 100%
  - rework: 10 个任务，平均 13.5 步，成功率 100%
- 整体错误数: 2 个（review 1，rework 1）
- 趋势变化:
  - avgStepsDelta: +2.597（平均步数增加）
  - errorRateDelta: +0.046（错误率上升 4.6%）
  - successRateDelta: +0.059（成功率上升 5.9%）

## 问题诊断

### 问题 1: Git 操作白名单执行偏差
**证据**: rework agent 的 `topGitOp` 为 `"checkout"`，但 `developer.md` 的 Git 白名单仅明确允许 `status/diff/add/commit/push`，未将 `checkout` 列入禁止清单；review agent 的 `topGitOp` 为 `"diff"`，而 `reviewer.md` 严禁任何本地 `git` 命令。
**根因分析**: `developer.md` 的白名单使用"只允许...禁止历史浏览命令"的表述，边界不够清晰，agent 在需要确认状态或切换上下文时可能下意识调用 `git checkout` 等非白名单命令。review agent 的 `git diff` 则可能是对 `gh pr diff` 与本地 `git diff` 的混淆。
**改进建议**:
1. 在 `developer.md` 的 Git 白名单章节中，以独立列表形式明确禁止 `git checkout`、`git reset`、`git merge` 等 8 个高危操作，并强调"你已在临时 worktree 的正确分支上，任何分支切换都是多余且危险的"。（已应用）
2. 在 `reviewer.md` 第 0 步错误预判中，增加一条显式区分：`gh pr diff` 是允许的远程 PR 交互命令，`git diff` 是严格禁止的本地工作树命令，防止 agent 将 `gh` 子命令误写为 `git` 子命令。

### 问题 2: 错误率与步数双升趋势
**证据**: `errorRateDelta = +0.046`（+4.6%），`avgStepsDelta = +2.597`（+2.6 步/任务），当前批次 `successRate = 1`（100%）。
**根因分析**: 任务最终均通过 WriteFile 产出了报告（workflow 兜底机制也起到作用），但执行过程中遇到的错误变多，agent 花费更多步数恢复或绕过错误。rework 模式下 `test -f` 等前置确认 Shell 调用可能消耗了本应用于修复的预算。rework 平均 Shell 调用 4.2 次/任务，在 10 次上限内但仍有优化空间。
**改进建议**:
1. 在 `developer.md` 的 rework 模式规则中，将 `test -f` 计入探索预算（≤3 次），并允许对 review report 中给出的相对路径直接 `ReadFile`，用错误处理替代前置存在性确认。（已应用）
2. 在 `developer.md` 的"常见错误快速诊断表"中增加"趋势感知"提醒：如果前 5 步内累计遇到 1 次错误，下一步必须是 `ReadFile` 确认状态或 `WriteFile` 记录错误，禁止新的 Shell 探索；前 10 步内累计 2 次错误立即进入收尾模式。

### 问题 3: topIssues 缺失导致无法定位重复失败模式
**证据**: 报告中 `topIssues: null`，未能从 19 条日志中提取出常见失败模式。
**根因分析**: dashboard 的日志分析模块可能因错误样本不足或分类阈值过高，未能输出 Top Issues。没有分类数据，prompt 演进就失去了针对性。
**改进建议**:
1. 建议检查 `cmd/dashboard/` 中的日志分类逻辑，降低 Top Issues 提取阈值（如从"出现 2 次以上"降至"出现 1 次以上"，保留 Top 5），并增加基于错误关键词的自动分类标签（如 `file_not_found`、`old_string_not_found`、`command_not_found`、`chunk_exceed`、`auth_failed` 等）。
2. 在 `workflows/review.yaml` 和 `workflows/rework.yaml` 的 `diagnoseAgentRun` 步骤中，将错误关键词分类结果写入结构化 JSON（如 `logs/diagnose.json`），供 dashboard 直接消费，减少日志正则解析的不确定性。

## 已应用改进

1. **`prompts/developer.md`** — Git 操作白名单强化
   - 将原来的"只允许...禁止历史浏览命令"改为"允许列表 + 明确禁止列表"双栏格式，新增禁止：`git checkout`、`git reset`、`git revert`、`git merge`、`git rebase`。
   - 强调"你已在临时 worktree 的正确分支上，任何分支切换都是多余且危险的"。

2. **`prompts/developer.md`** — rework 模式路径确认优化
   - 在"禁止 explore"规则中，将 `test -f` 明确计入探索预算（总计 ≤3 次）。
   - 新增"路径确认优化"规则：review report 中的相对路径可直接 `ReadFile`，用错误处理捕获文件缺失，无需前置 `test -f`。
   - 在任务启动检查清单的"文件路径不存在"错误预判中，标注"rework 模式例外"。

## 趋势追踪

- **与上次对比**: 成功率提升 +5.9%（受益于 workflow 的 diagnoseAgentRun 兜底和 prompt 中的强制写入检查点），但错误率上升 +4.6%、平均步数增加 +2.6，说明 agent 的**执行韧性增强但执行效率在轻微恶化**。任务在出错后仍能完成，但消耗了更多步数。
- **建议下次重点关注**:
  1. **dev agent**: 本次统计无 dev 任务数据，而历史数据显示 dev 平均 72 步、错误率 100%，是系统最大的效率黑洞。建议下次确保 dev 任务被纳入统计。
  2. **rework Shell 分布**: 验证本次改进后，`test -f` 前置调用是否减少，Shell 调用次数是否从 4.2 次/任务下降。
  3. **review Git 禁令**: 监控 review agent 的 `topGitOp` 是否仍为 `diff`，如果是，需进一步区分 `gh pr diff` 与 `git diff` 的统计口径。
