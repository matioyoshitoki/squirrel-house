# Prompt 演进报告 — type-one

生成时间: 2026-05-19T08:05:32+08:00
分析范围: 2026-05-18T22:46:30+08:00 至 2026-05-19T08:03:55+08:00

## 统计摘要

- 分析日志数: 8
- Agent 类型分布:
  - review: 3 个任务
  - rework: 5 个任务
- 趋势变化 (与上期对比):
  - avgStepsDelta: -21.75
  - errorRateDelta: -1.75
  - successRateDelta: -0.375

## 问题诊断

### 问题 1: rework 任务零产出（successRate 0%）

**证据**: 5 个 rework 任务，avgSteps=1，totalWrites=0，successRate=0%。totalToolCalls 中 ReadFile 11 次、Shell 1 次，说明 agent 有少量文件读取活动，但未产生任何文件修改或写入。

**根因分析**:
- developer.md 中 rework 模式虽要求启动确认和强制产出，但 agent 未执行 WriteFile 即终止。统计表明 agent 在极早阶段（平均 1 步）就以纯文本输出结束任务，未触发任何文件写入。
- "必须有文件修改产出"的规则依赖 agent 自觉遵守，缺少更早的强制触发点；且"停止并报告"的表述易被 agent 误解为纯文本输出即可。

**改进建议**:
1. 在 `prompts/developer.md` 中将 rework 启动确认（写 `logs/rework-start.txt`）明确为不可跳过的第 2 步，并增加"提前终止也必须 WriteFile"的铁律，同时规定 3 步内无写操作即触发熔断。（已修改）
2. 在 `workflows/rework.yaml` 中增加 `diagnoseAgentRun` 步骤，检测日志中是否有 WriteFile/StrReplaceFile，零写操作时记录诊断信息，帮助定位 agent 提前终止根因。（已修改）

### 问题 2: review 任务产出率过低（successRate 33.3%）

**证据**: 3 个 review 任务，avgSteps=7，totalToolCalls 中 ReadFile 35 次、Shell 11 次，但 totalWrites 仅 1 次，successRate=33.3%。agent 投入大量步数阅读文件，却只有 1 个任务产出了审查报告。

**根因分析**:
- reviewer.md 已多次强调 WriteFile，但仍有 2/3 的任务未执行。说明 agent 在错误路径（gh 失败、PR 不存在等）或探索过程中以纯文本输出终止，绕过了强制写入约束。
- workflow 的诊断步骤 `diagnoseAgentRun` 仅检查是否有工具调用，不检查是否有 WriteFile，导致"有活动但无产出"的任务未被兜底。

**改进建议**:
1. 在 `prompts/reviewer.md` 的"快速开始"和"执行约束"中增加"提前终止也必须 WriteFile"的明确规则，并将零产出熔断中的"空报告"要求升级为必须包含时间戳和当前状态的正式报告。（已修改）
2. 在 `workflows/review.yaml` 的 `diagnoseAgentRun` 中增加 `WRITE_COUNT` 检测，当 agent 有工具调用但无写操作时，生成默认失败报告并标注"可能因 prompt 约束提前终止且未执行强制写入"。（已修改）

### 问题 3: 整体成功率趋势恶化（successRateDelta -0.375）

**证据**: 趋势指标显示 successRateDelta=-0.375，avgStepsDelta=-21.75，errorRateDelta=-1.75。

**根因分析**:
- 步数大幅下降（-21.75）伴随成功率下降，说明 agent 倾向于在更早阶段终止任务，但终止方式不是 WriteFile，而是纯文本输出。这与 prompt 中"停止并报告"的表述被 agent 误解为"纯文本报告"有关。
- errorRate 下降（-1.75）并非真正的改善，而是 agent 提前终止避免了后续错误，属于"以零产出换取零错误"的虚假改善。

**改进建议**:
1. 对所有 agent prompt 中的"停止并报告"类表述进行审计，统一替换为"停止并执行 WriteFile 写报告"。（本次已针对 review/developer 修改）
2. 在 dashboard 统计中增加"纯文本终止率"指标，区分"工具调用后无 WriteFile"和"完全无工具调用"两种零产出模式，帮助更精准定位问题。（建议下次迭代评估）

## 已应用改进

1. **`prompts/developer.md`**:
   - 强化 rework 模式启动确认：第 1 步读 review report，第 2 步必须 WriteFile 写 `logs/rework-start.txt`，禁止纯文本代替。
   - 新增"3 步无写操作即熔断"规则：如果任务已执行 3 步仍未调用任何写操作，立即 WriteFile 写 `logs/rework-halted.md` 并结束。
   - 强化"提前终止也必须 WriteFile"规则：任何停止方式都必须是 WriteFile。

2. **`prompts/reviewer.md`**:
   - 在"快速开始"中新增第 7 步："提前终止也必须 WriteFile"。
   - 在"执行约束"中升级零产出熔断要求：禁止以纯文本解释代替 WriteFile，报告必须包含时间戳和当前状态。
   - 更新统计引用为当前 33% 成功率，强化警示。

3. **`workflows/review.yaml`**:
   - `diagnoseAgentRun` 增加 `WRITE_COUNT` 检测（统计 WriteFile/StrReplaceFile 出现次数）。
   - 当 `STEP_COUNT>0` 但 `WRITE_COUNT=0` 时，生成兜底失败报告，明确标注"agent 执行了工具调用但未产生 WriteFile/StrReplaceFile 输出"。

4. **`workflows/rework.yaml`**:
   - 在 `runAgent` 与 `checkAgentOutput` 之间新增 `diagnoseAgentRun` 步骤。
   - 统计日志中的工具调用和写操作次数，写入 `logs/rework-diagnose.md`，为零产出任务提供诊断线索。

## 趋势追踪

- **与上次对比**: 恶化（successRateDelta=-0.375，rework 从低成功率跌至 0%，review 维持在 33.3%）
- **建议下次重点关注**:
  - rework 任务在 prompt 修改后是否恢复产出（目标：successRate > 0%）
  - review 任务的 WriteFile 执行率是否提升（目标：successRate > 50%）
  - 新增 diagnose 文件是否揭示了 agent 提前终止的具体原因
