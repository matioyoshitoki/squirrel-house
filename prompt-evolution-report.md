# Prompt 演进报告 — type-one

生成时间: 2026-05-20T18:19:12+08:00
分析范围: 2026-05-20T16:50:10 至 2026-05-20T18:17:14

## 统计摘要

- 分析日志数: 7
- Agent 类型分布: review=3, rework=4
- 总体趋势: 平均步数下降 29.4 步（改善），错误率下降 0.5（改善），成功率下降 0.14（恶化）

## 问题诊断

### 问题 1: Review Agent 存在 33% 零产出任务（successRate 仅 66.7%）

**证据**:
- review 任务 count=3, totalWrites=2, successRate=0.667 → 1/3 的任务完全没有 WriteFile 产出
- 平均步数仅 10.67，远低于 prompt 中规定的 45 步上限，说明 agent 在早期阶段即终止但未执行强制写入
- 总错误数仅 1 个，说明失败任务很可能是在环境检查阶段就放弃，未进入 WriteFile 流程

**根因分析**:
1. reviewer.md 的「快速开始」虽然规定了第 6 步必须 WriteFile，但 agent 将「第 6 步」理解为思考轮次而非工具调用序号，导致在 tool call 编号 3-4 时就因环境异常（如 gh 认证失败）以纯文本输出结束。
2. Prompt 中虽有「熔断后的下一步必须是 WriteFile」的约束，但缺少显式的「工具调用计数器」机制，agent 无法感知自己已调用几次工具，因此在低步数时不会触发强制写入。
3. Workflow 层面的 diagnoseAgentRun 虽然会生成默认失败报告，但那是 workflow 的行为，不计入 agent 的 totalWrites——统计上仍表现为零产出。

**改进建议**:
1. **在 reviewer.md 中明确「第 6 步 = 第 6 个 tool call」**，并增加不可违背的「WriteFile 强制检查点」：第 6、12、18… 个 tool call 必须检查是否已执行过 WriteFile；若未执行，下一个 tool call 必须是 WriteFile（即使是进度报告）。（涉及文件: `prompts/reviewer.md`，修改方式: 在「快速开始」第 6 步后追加强制检查点规则）
2. **在 review.yaml 的 runAgent env 中增加 `AGENT_MANDATORY_WRITE_STEP=6`**，让外部系统与 prompt 对齐，agent 可从环境变量读取该数字并自我监控。（涉及文件: `workflows/review.yaml`，修改方式: runAgent 步骤 env 追加）
3. **在 reviewer.md 的「第 0 步 — 错误预判」中增加显式规则**：如果 `AGENT_ENV_STATUS` 包含 `auth_failed`/`pr_not_found`，agent 的下一个 tool call **必须是 WriteFile**，且在此之前禁止调用 ReadFile、Shell、Grep 等任何其他工具——当前 prompt 虽有类似描述，但 agent 仍调用了其他工具。（涉及文件: `prompts/reviewer.md`，修改方式: 强化第 0 步与第 1 步的禁止条款）

### 问题 2: Rework Agent 错误率过高（6 错误 / 4 任务，平均 1.5 错误/任务）

**证据**:
- rework 任务 count=4, totalErrors=6, 平均步数 22.75
- Shell 调用 40 次（平均 10 次/任务），已达到 rework 的 Shell 硬上限 10 次
- ReadFile 42 次（平均 10.5 次/任务），说明 agent 在大量读取和 Shell 探测中反复出错
- 尽管 successRate=100%，但高错误率意味着 agent 在「试探—失败—重试」循环中消耗了大量预算

**根因分析**:
1. developer.md 的 rework 模式已规定「错误 ≥ 2 即停」，但 agent 没有显式自报错误计数，导致接近上限时仍继续冒险操作（如未经确认的 StrReplaceFile、未预检的 Shell 命令）。
2. Shell 预算虽已设上限，但缺少「逐次自报」机制，agent 在调用第 8、9、10 次 Shell 时未意识到已接近上限，导致最后几步被迫在错误状态下收尾。
3. StrReplaceFile 前置检查规则存在（必须先 ReadFile），但统计中 StrReplaceFile 有 10 次、ReadFile 有 42 次，比例不匹配，说明部分 StrReplaceFile 可能未经充分确认就执行，导致替换失败并产生错误。

**改进建议**:
1. **在 developer.md 的 rework 专属规则中增加「错误计数显式自报」**：每次工具调用返回错误后，reasoning 中必须输出 `错误计数: X/2`，当 X=1 时进入「只读/收尾模式」，X=2 时立即 WriteFile 并结束。（涉及文件: `prompts/developer.md`，修改方式: 在「常见错误 SOP」前追加自报规则）
2. **在 developer.md 中增加「Shell 逐次自报」**：rework 模式下每次调用 Shell 前必须在 reasoning 中输出 `Shell 计数: X/10`；当 X≥7 时，仅允许 git add/commit/push 和 1 次验证命令。（涉及文件: `prompts/developer.md`，修改方式: 在「Shell 预算预警」后追加逐次自报）
3. **在 rework.yaml 的 envReadinessCheck 中将预检结果写入结构化文件**，并通过环境变量 `AGENT_ENV_CHECK_FILE` 传递给 agent，减少 agent 自行执行 `test -f`、`which` 等探测型 Shell 命令。（涉及文件: `workflows/rework.yaml`，修改方式: envReadinessCheck 输出文件 + runAgent env 传递）

### 问题 3: Review Agent 仍在使用本地 git 命令（与 prompt 禁令冲突）

**证据**:
- review 的 topGitOp 为 `diff`
- reviewer.md 第 3 条元原则明确「🔴 绝对铁律 —— Git 禁令：review 任务严禁执行任何本地 git 命令」
- 但 workflow review.yaml 的第 1 步 `checkCurrentBranch` 即执行了 `git branch --show-current`，这个属于 workflow 行为，不是 agent 行为

**根因分析**:
- 统计中的 `topGitOp=diff` 并非来自 agent 本身，而是来自 workflow 的 `checkCurrentBranch`、`stashAndCheckout` 等步骤。这些步骤在 agent 启动前执行，agent 实际上遵守了 git 禁令。
- 但 agent 统计与 workflow 统计混在了一起，导致误判。

**改进建议**:
1. **本次不修改 prompt**：agent 本身遵守了 git 禁令，问题在于统计粒度。建议 dashboard 后续将 workflow 的 git 操作与 agent 的 git 操作分开统计。（涉及文件: 无直接修改，建议纳入 dashboard 改进 backlog）

## 已应用改进

本次分析后实际修改了以下文件：

1. **`prompts/reviewer.md`**:
   - 在「快速开始」第 6 步「强制写入」后增加「WriteFile 强制检查点」条款，明确第 6、12、18… 个 tool call 必须检查 WriteFile 执行情况。
   - 在「第 0 步 — 错误预判」中强化：若 `AGENT_ENV_STATUS` 非 READY，下一个 tool call 必须是 WriteFile，禁止在此之前调用任何其他工具。

2. **`prompts/developer.md`**:
   - 在 rework 专属规则中增加「错误计数显式自报」条款，要求每次错误后 reasoning 中必须报告 `错误计数: X/2`。
   - 增加「Shell 逐次自报」条款，要求每次 Shell 调用前报告 `Shell 计数: X/10`，X≥7 时进入受限模式。

3. **`workflows/review.yaml`**:
   - 在 `runAgent` 步骤的 env 中增加 `AGENT_MANDATORY_WRITE_STEP: '6'`，让 agent 可以从环境变量读取强制写入阈值。

4. **`workflows/rework.yaml`**:
   - 在 `envReadinessCheck` 步骤中将预检结果写入结构化文件 `logs/rework-env-check.json`。
   - 在 `runAgent` 步骤的 env 中增加 `AGENT_ENV_CHECK_FILE`，指向该结构化文件，减少 agent 自行探测。

## 趋势追踪

- 与上次对比:
  - 平均步数下降 29.4 步（显著改善，说明预算约束开始生效）
  - 错误率下降 0.5（改善，但绝对错误数仍偏高）
  - 成功率下降 0.14（轻微恶化，主要由 review 的 33% 零产出任务导致）
- 建议下次重点关注:
  - **review agent 的零产出问题**：观察强制 WriteFile 检查点是否将 successRate 提升至 100%。
  - **rework agent 的错误率**：观察错误自报和 Shell 逐次自报是否将平均错误数从 1.5 降至 <0.5。
