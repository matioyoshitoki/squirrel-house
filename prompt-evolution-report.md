# Prompt 演进报告 — type-one

生成时间: 2026-05-20T16:50:10+08:00
分析范围: 2026-05-20T15:23:53 到 2026-05-20T16:50:10

## 统计摘要

- 分析日志数: 4
- Agent 类型分布: dev × 1, review × 2, rework × 1

## 问题诊断

### 问题 1: dev agent 预算约束全面失效（高步数 + 高 Shell + 高错误）

**证据**:
- avgSteps = 72，超过 prompt 中 60 步硬上限 20%
- Shell 调用 34 次/任务，超过 20 次上限 70%
- totalErrors = 2，错误率 100%（1/1 个任务发生错误）
- topGitOp = `status`，说明 git status 被反复调用，直接违反 prompt 中"git status 最多 3 次"的铁律

**根因分析**:
1. **workflow 与 prompt 步数上限不匹配**：`workflows/dev.yaml` 中 `--max-steps-per-turn 75` 高于 prompt 的 60 步，agent 感知到的外部限制比内部约束更宽松，prompt 的熔断规则被架空。
2. **Shell 预算缺乏预警机制**：prompt 中只规定了 20 次 Shell 硬上限，但没有设置中间预警线（如 15 次预警）。agent 在步数尚未触顶时，认为 Shell 预算"还有余量"，用 `git status`、`ls`、`grep` 等小命令反复确认状态。
3. **过时的统计锚点失效**：developer.md 第 138 行引用了一个旧统计（"Shell 调用高达 76 次"），agent 看到当前 Shell 调用 34 次反而认为"比历史峰值低"，降低了紧迫感。

**改进建议**:
1. **对齐 workflow 与 prompt 的步数上限**：将 `workflows/dev.yaml` 的 `--max-steps-per-turn` 从 75 下调至 65，使其略低于 prompt 的 60 步熔断线，给 agent 一个不可逾越的外部硬边界。（涉及文件: `workflows/dev.yaml`，修改方式: `--max-steps-per-turn 75` → `--max-steps-per-turn 65`）
2. **增加 Shell 预警线并更新统计锚点**：在 `prompts/developer.md` 的预算约束中增加 "Shell 15 次预警"，达到后禁止运行任何新的构建/测试/探索命令，只能执行代码修改、1 次 git add/commit/push 和写报告。同时将第 21 行和第 138 行的统计引用更新为当前真实数据（72 步、34 次 Shell、2 次错误）。（涉及文件: `prompts/developer.md`）
3. **将 git status 禁令提升为熔断项**：当前 git status 限制放在效率约束中，agent 将其视为"建议"。应将其与"步数上限"并列，明确"超过 3 次 git status → 立即停止并报告"。（涉及文件: `prompts/developer.md`）

### 问题 2: review agent 错误率与 Shell 预算双高，且违反 git 禁令

**证据**:
- totalErrors = 4，2 个任务均有错误，错误率 100%
- Shell 调用 58 次 / 2 任务 = 29 次/任务，超过 15 次上限 93%
- topGitOp = `diff`，直接违反 reviewer.md 第 327 行"严禁执行任何本地 `git` 命令"的铁律

**根因分析**:
1. **"先 test 再 ReadFile"策略浪费 Shell 预算**：reviewer.md 第 292 行要求"读取规范文档前，先用 Shell `ls`/`test -f` 确认文件存在"。这种策略将每个文档读取拆成 2 步（1 次 Shell + 1 次 ReadFile），当审查涉及多个规范文档时，Shell 调用迅速膨胀。更糟的是，如果路径本身错误，test 也会失败，既浪费步数又增加错误数。
2. **git 禁令位置靠后、威慑不足**：git 禁令被放在执行约束第 13 条（第 327 行），而不是元原则。agent 在启动初期容易忽略靠后的约束，在遇到 `gh` 命令失败时本能地 fallback 到 `git diff`。
3. **workflow 步数限制过于宽松**：`workflows/review.yaml` 的 `--max-steps-per-turn 10000` 形同虚设，agent 完全依赖 prompt 自我约束，而 prompt 中的 45 步上限对 LLM 来说不够具象。

**改进建议**:
1. **消除冗余的 Shell 预检**：将 `prompts/reviewer.md` 第 292 行的"先用 Shell 确认文件存在"改为"直接用 ReadFile 读取，若返回文件不存在则标记为'未验证'并跳过"。仅此一项预计可将 Shell 调用降低 30-40%。（涉及文件: `prompts/reviewer.md`）
2. **将 git 禁令前移至元原则并增加后果**：在 reviewer.md 的元原则中新增一条绝对铁律："review 任务严禁执行任何本地 `git` 命令。所有审查依据必须通过 `gh pr diff/view` 获取。违反此条将直接导致审查结论不可信并触发熔断。"（涉及文件: `prompts/reviewer.md`）
3. **收紧 workflow 步数上限**：将 `workflows/review.yaml` 的 `--max-steps-per-turn` 从 10000 降至 50，使其略高于 prompt 的 45 步上限，形成外部硬边界。（涉及文件: `workflows/review.yaml`）

### 问题 3: rework 步数接近上限，workflow 外部约束过于宽松

**证据**:
- avgSteps = 49，接近 prompt 中 rework 40 步上限（超出 22%）
- workflow 中 `--max-steps-per-turn 10000`，对 rework 完全无外部约束

**根因分析**:
rework 的 prompt 内嵌在 developer.md 中，虽有 40 步上限，但 workflow 没有配合的外部限制。agent 在修复多个 review 问题时容易逐条修复、逐条验证，导致步数膨胀。

**改进建议**:
1. **对齐 rework workflow 的步数上限**：将 `workflows/rework.yaml` 的 `--max-steps-per-turn` 从 10000 下调至 50，略高于 prompt 的 40 步上限，防止 agent 透支预算。（涉及文件: `workflows/rework.yaml`，修改方式: `--max-steps-per-turn 10000` → `--max-steps-per-turn 50`）

## 已应用改进

### 1. `workflows/dev.yaml` — 下调外部步数限制
- `--max-steps-per-turn` 从 `75` 下调至 `65`，使外部限制严格低于 prompt 的 60 步熔断线，消除 agent 的"余量幻觉"。

### 2. `workflows/review.yaml` — 收紧外部步数限制
- `--max-steps-per-turn` 从 `10000` 下调至 `50`，与 prompt 的 45 步上限对齐，防止 review agent 无限制探索。

### 3. `workflows/rework.yaml` — 收紧外部步数限制
- `--max-steps-per-turn` 从 `10000` 下调至 `50`，与 prompt 的 40 步上限对齐。

### 4. `prompts/developer.md` — 强化预算约束与统计锚点
- **更新统计锚点**（第 21 行）：将引用从"rework 任务平均 23 步..."更新为"dev 任务平均 **72 步、2 次错误、34 次 Shell/任务**，topGitOp 为 `status`——预算约束被全面突破，必须严格执行"。
- **更新 git status 引用**（第 138 行）：将旧统计"Shell 调用高达 76 次"更新为"当前统计 dev 任务 Shell 调用 34 次/任务且 topGitOp 为 `status`——预算约束仍被突破"。
- **增加 Shell 预警线**：在"步数硬上限与预警"章节中新增"Shell 预警：dev 达到 **15 次**时进入收尾阶段，禁止运行新的构建/测试/探索命令"。
- **强化错误预警**：增加"错误 3 次预警"，第 3 次错误后进入只读/收尾模式（只允许代码修改、1 次验证、git 操作和写报告），第 5 次才完全熔断。

### 5. `prompts/reviewer.md` — 减少 Shell 浪费并前置 git 禁令
- **前置 git 禁令**：在元原则中新增第 3 条绝对铁律："review 任务严禁执行任何本地 `git` 命令（包括 `git status`、`git diff`、`git log` 等）。所有代码审查依据必须通过 `gh pr diff` 和 `gh pr view` 获取。违反此条将直接导致审查结论不可信。"
- **消除冗余 Shell 预检**：将第 292 行的"先用 Shell `ls`/`test -f` 确认文件存在"改为"直接用 `ReadFile` 读取目标文件。如果返回文件不存在错误，将该维度标记为'未验证（文件缺失）'并跳过，不得尝试替代路径。反复用 Shell 预检是 review 任务 Shell 预算超标的首要原因。"
- **更新 Shell 统计锚点**：在 Shell 预算说明中将"当前统计平均 15 次"更新为"当前统计平均 29 次/任务，严重超过 15 次上限"。

## 趋势追踪

- 与上次对比: 上期（design）avgSteps 上升 22.5 步、errorRate 上升 4.3；本期（dev/review/rework）errorRateDelta -3.5、successRateDelta +0.25，整体错误率和成功率呈改善趋势。但 dev/review 的步数、Shell 调用和错误率仍然显著高于 prompt 预算，说明预算约束的**执行力**仍是最大短板。
- 建议下次重点关注: **dev agent 的 Shell 调用分布**（具体是哪些 Shell 命令占用了 34 次预算）和 **review agent 的错误类型分类**（当前仅知 4 个错误，但缺少是路径类、gh 认证类还是 ReadFile 类的细分）。建议在日志采集侧增加错误类型和 Shell 命令分类标签，以便精准定位。
