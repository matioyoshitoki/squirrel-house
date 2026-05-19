# Prompt 演进报告 — type-one

生成时间: 2026-05-19T15:45:00+08:00
分析范围: 2026-05-19T14:19:09 至 2026-05-19T15:40:05

## 统计摘要

- 分析日志数: 10
- Agent 类型分布:
  - review: 4 个任务
  - rework: 6 个任务
- 整体平均步数: 14.9 步（较上次 +8.1 步）
- 整体错误率: 上升 30%（errorRateDelta +0.3）
- 整体成功率: 持平（successRateDelta 0）

## 问题诊断

### 问题 1: review agent 零产出率过高（Critical）
**证据**:
- review successRate = 25%（4 个任务中只有 1 个有产出）
- totalErrors = 3，错误率 75%
- totalWrites（WriteFile）仅 1 次，但有 4 个任务
- avgSteps = 11.5，远低于正常审查所需步数，说明任务在早期快速失败
- topIssues 中 `review_high_error_rate` 出现 3 次，`review_no_output` 出现 1 次

**根因分析**:
1. **workflow 缺少前置环境检查**：review 任务在 runAgent 前没有检查 `gh auth status` 和 PR 可访问性。agent 启动后直接在 gh 环节失败，浪费步数且触发错误。
2. **prompt 的"强制写入"规则执行不到位**：虽然 reviewer.md 明确要求"gh 失败时必须 WriteFile"，但 agent 在遭遇错误后仍以纯文本输出结束任务。规则给 agent 留下了"理解后自由裁量"的空间。
3. **错误熔断后的行为未完全机械化**：prompt 说"熔断后停止并报告"，但 agent 对"停止"的理解可能是"不再调用工具，直接文本回复"。

**改进建议**:
1. **在 `workflows/review.yaml` 中增加 `envReadinessCheck` 步骤**（已应用）：在 runAgent 前运行 `gh auth status` 和 `gh pr view`，如果失败直接由 workflow 生成默认失败报告并跳过 agent，避免 agent 在已知错误环境中空转。
2. **在 `prompts/reviewer.md` 中将 gh 失败后的 WriteFile 改为"机械下一步"**（已应用）：将"如果失败，必须立即 WriteFile"改为"如果任一失败，你的下一步（第 2 步）必须是 WriteFile，禁止执行任何其他工具调用"。消除 agent 的自由裁量空间。
3. **在 `prompts/reviewer.md` 中增加"任务结束前强制检查"**（已应用）：新增规则 15，要求 agent 在最后一个思考轮次中必须回答"我是否已经执行了 WriteFile？"，如果否则立即执行。禁止在没有 WriteFile 的情况下结束任务。

### 问题 2: rework agent 成功率偏低且步数膨胀趋势（Warning）
**证据**:
- rework successRate = 50%（6 个任务中 3 个有产出）
- avgSteps = 17.17 步，且 avgStepsDelta = +8.1（步数显著膨胀）
- totalErrors = 6（平均每个任务 1 次错误）
- topIssues 中 `rework_high_error_rate` 出现 2 次

**根因分析**:
1. **workflow 的产出检查过于依赖 git diff**：`workflows/rework.yaml` 的 `checkAgentOutput` 仅检查 `git diff --quiet`。agent 按 prompt 要求写入 `logs/rework-noop.md` 或 `logs/rework-halted.md` 时，这些文件可能被 `.gitignore` 忽略，导致 workflow 误判为"零产出"而失败。
2. **步数膨胀但成功率未提升**：avgSteps 增加了 8.1 步，但 successRate 持平，说明额外步数消耗在无效探索上。rework 预算 40 步，当前 17.17 步虽未超限，但趋势令人担忧。
3. **进度检查频率过低**：developer.md 要求每 10 步检查一次进度，但 rework 预算只有 40 步，意味着最多只能检查 3 次，无法及时发现无效探索。

**改进建议**:
1. **修改 `workflows/rework.yaml` 的 `checkAgentOutput`**（已应用）：将产出检查逻辑从仅依赖 `git diff` 改为"先检查日志中的 WriteFile/StrReplaceFile 调用次数，再检查 git diff"。如果 agent 有写入操作即视为有产出，避免 logs/ 目录被 gitignore 导致的误判。
2. **在 `prompts/developer.md` 中增加 rework 专属步数追踪**（已应用）：在 rework 专属规则中新增"每 5 步必须在思考中自评一次当前进度"，如果过去 5 步没有产生任何文件修改，立即停止并报告。将通用"进度检查点"也注明 rework 模式为每 5 步检查一次。

### 问题 3: 整体错误率恶化（Warning）
**证据**:
- errorRateDelta = +0.3（整体错误率上升 30%）
- successRateDelta = 0（成功率没有随步数增加而改善）

**根因分析**:
- 步数增加但成功率持平，说明 agent 正在做更多的"无产出探索"——读取更多文件、执行更多 Shell 命令，但这些额外操作没有转化为有效产出。
- 缺乏 workflow 层面的前置过滤，导致 agent 反复在已知错误环境（如 gh 未认证）中尝试。

**改进建议**:
1. **对所有涉及外部服务（gh、git、make）的 workflow 增加前置就绪检查**：已在 review workflow 中应用，建议后续对 dev、doc-gardener 等 workflow 也做类似审计。
2. **在 dashboard 或日志收集中增加"错误来源分类"**：当前 topIssues 只能笼统地识别"high_error_rate"，建议按错误类型（gh 认证失败、文件不存在、StrReplaceFile 失败、Shell 超时等）细分，以便更精确地优化 prompt。

## 已应用改进

本次分析后已直接修改以下文件：

### 1. `workflows/review.yaml`
- **新增 `envReadinessCheck` 步骤**：在 runAgent 之前检查 `gh auth status` 和 `gh pr view`。如果 GitHub CLI 未认证或 PR 不存在，直接生成默认失败报告到 `logs/review-report-<pr>-fail.md`，并输出 `SKIP_AGENT` 跳过 agent 执行。
- **为 `runAgent` 增加 condition**：`condition: '{{ contains (.State.envCheckStatus | trim) "READY" }}'`，确保只有环境就绪时才启动 agent。

### 2. `prompts/reviewer.md`
- **强化错误熔断后的行为**：将"单次错误后禁止立即重试，连续 2 次错误立即熔断"改为"熔断后的下一步必须是 WriteFile，禁止以纯文本输出作为任务终点"。
- **机械化 gh 失败后的步骤**：将第 1 步环境预检查的结果分支明确化——如果 gh 检查失败，下一步（第 2 步）**必须是** WriteFile，禁止执行任何其他工具调用。
- **新增规则 15（任务结束前强制检查）**：要求 agent 在最后一个思考轮次中必须确认"我是否已经执行了至少一次 WriteFile？"，如果否则立即执行。禁止在没有 WriteFile 的情况下结束任务。

### 3. `workflows/rework.yaml`
- **优化 `checkAgentOutput` 逻辑**：不再仅依赖 `git diff --quiet` 判断产出。改为先读取 agent 日志，统计 `WriteFile` 和 `StrReplaceFile` 调用次数（`WRITE_COUNT`）。如果 `WRITE_COUNT > 0` 即视为有产出；如果为 0 再检查 git diff。这解决了 agent 写入 logs/ 目录但被 gitignore 忽略导致的误判问题。

### 4. `prompts/developer.md`
- **增加 rework 步数追踪规则**：在 rework 专属规则中新增"每 5 步必须在思考中自评一次当前进度，如果过去 5 步没有文件修改则停止"。
- **调整通用进度检查点**：注明 rework 模式为每 5 步检查一次（dev 仍为每 10 步）。

## 趋势追踪

- 与上次对比: 错误率上升 30%（恶化），平均步数增加 8.1 步（恶化），成功率持平（无改善）
- 建议下次重点关注:
  1. **review agent 的 successRate**：本次修改的核心目标是将其从 25% 提升到 60% 以上。如果仍然低于 50%，需要进一步审查 workflow 的 diagnoseAgentRun 兜底逻辑和 agent 的实际工具调用序列。
  2. **rework agent 的 avgSteps 趋势**：当前 17.17 步且有 +8.1 的上升趋势。如果继续膨胀，需要收紧 prompt 中的探索预算（如将"定位修改点不得超过 10 步"改为 6 步）。
  3. **错误来源细分**：建议在日志收集中增加错误类型标签（gh/文件/替换/命令/超时），以便更精准地定位问题。
