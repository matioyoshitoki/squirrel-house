# Prompt 演进报告 — type-one

生成时间: 2026-05-22T10:56:37+08:00
分析范围: 2026-05-21T15:56:07 至 2026-05-22T10:56:37

## 统计摘要

- 分析日志数: 4
- Agent 类型分布:
  - design: 1 任务
  - dev: 2 任务
  - e2e: 1 任务

## 问题诊断

### 问题 1: dev 任务高步数 + 高错误率

**证据**:
- dev count=2, avgSteps=55, totalErrors=13（6.5 errors/任务）
- dev 有 2 个任务超过 50 步，占比 100%
- successRate=1 但错误率 100%，说明 agent 在反复试错后才完成产出
- trends.avgStepsDelta=12.25, errorRateDelta=2.3，整体恶化主要由 dev 拉动

**根因分析**:
1. `workflows/dev.yaml` 中 `--max-steps-per-turn=55` 与 developer.md 的 60 步上限过于接近，agent 几乎耗尽全部预算，缺少提前收尾的硬边界。
2. `developer.md` 的 dev 模式缺少像 design/rework 那样的"启动前错误预判"强制检查，agent 对常见错误（文件不存在、命令不存在）没有提前防御，导致错误频发。
3. 缺少"文档存在性检查"的硬性要求，agent 频繁直接 `ReadFile` 不存在的文件，浪费步数并触发错误。
4. 虽然有"零进度熔断"（前 5 步），但缺少更宽松的"早期产出压力"（如 10 步内必须完成第一个文件修改），导致探索期容易膨胀。

**改进建议**:
1. **已应用**：将 `workflows/dev.yaml` 的 `--max-steps-per-turn` 从 55 降至 45，强制 agent 提前进入收尾阶段。
2. **已应用**：在 `prompts/developer.md` 的"任务启动检查清单"中新增"第 0 步 — 错误预判"，要求 agent 在启动时勾选 4 项高频错误预判（文件不存在、StrReplaceFile 不匹配、命令输出过大、命令不存在）。
3. **已应用**：在 `prompts/developer.md` 的"效率约束"中新增"早期产出压力"（10 步内必须完成第一个文件修改）和"文档存在性检查"（ReadFile 前必须 `test -f`）。
4. **已应用**：在 `prompts/developer.md` 中新增"常见错误快速诊断表（dev 模式）"，要求出现错误后 1 个思考轮次内完成诊断，禁止不经分析直接重试。

### 问题 2: design 任务错误率持续 100%

**证据**:
- design count=1, totalErrors=5, successRate=1
- 与上次对比：design 平均步数从 41.5 降至 34，有改善；但错误率仍然是 100%（5 errors/任务）
- 上次修改的强制报数机制、Shell 预算收紧（15 次）已部分生效（步数下降），但错误预判的执行效果不明显

**根因分析**:
1. `ui-designer.md` 中的"错误预判与强制 SOP"被放在"效率约束"末尾（第 12 条），agent 启动时容易忽略，未能真正在"第 1 步"完成预判。
2. "早期产出压力"为 12 步，对于平均 34 步的任务来说仍然偏宽松，agent 可能在探索阶段消耗过多预算后才被迫产出。

**改进建议**:
1. **已应用**：将 `prompts/ui-designer.md` 中的"错误预判与强制 SOP"从效率约束末尾抽出，提升为独立的"启动前错误预判"章节，放在"效率约束"和"快速开始"之间，提高可见性和执行优先级。
2. **已应用**：将 `ui-designer.md` 的"早期产出压力"从 12 步收紧到 10 步，并增加数据引用提示"当前统计 design 错误率 100%，早期产出是降低错误预算消耗的关键"。

### 问题 3: 整体趋势恶化（avgSteps +12.25, errorRate +2.3）

**证据**:
- trends.avgStepsDelta=12.25（步数大幅上升）
- trends.errorRateDelta=2.3（错误率大幅上升）
- 虽然 trends.successRateDelta=+0.2（成功率上升），但这是以更高错误和更多步数为代价的

**根因分析**:
- 新增 dev 任务数据暴露了 developer prompt 的防御不足。design 的步数虽然下降，但 dev 的高步数和高错误拉高了整体均值。
- 说明上次对 design 的改进有局部效果，但 dev 模式是新的瓶颈。

**改进建议**:
1. 重点观察本次 dev prompt 和 workflow 修改的效果。
2. 下次迭代若 dev 的 avgSteps 仍 > 50 或 errorRate 仍高，考虑进一步收紧 `max-steps-per-turn` 至 40，或在 workflow 层面增加前置环境检查（如预装依赖、预检文件存在性）。

## 已应用改进

本次分析后直接修改了以下文件：

1. **`workflows/dev.yaml`**
   - `--max-steps-per-turn` 从 55 降至 45，强制 agent 提前收尾，避免接近 60 步上限才停止。

2. **`prompts/developer.md`**
   - 在"任务启动检查清单"中新增"第 0 步 — 错误预判"，要求启动时勾选 4 项高频错误预判（文件路径不存在、StrReplaceFile 不匹配、Shell 输出过大、命令不存在）。
   - 在"效率约束"中新增"早期产出压力"：启动后 10 步内必须完成第一个 WriteFile/StrReplaceFile。
   - 在"效率约束"中新增"文档存在性检查"：ReadFile 前必须先用 `test -f` 确认文件存在。
   - 新增"常见错误快速诊断表（dev 模式）"，要求错误后 1 个思考轮次内完成诊断，禁止不经分析直接重试。
   - 调整相关章节编号。

3. **`prompts/ui-designer.md`**
   - 将"错误预判与强制 SOP"从效率约束第 12 条抽出，提升为独立的"启动前错误预判"章节，置于"效率约束"与"快速开始"之间。
   - "早期产出压力"从 12 步收紧到 10 步，增加数据引用强调紧迫性。

## 趋势追踪

- 与上次对比:
  - avgSteps: +12.25（恶化，主要由 dev 任务拉动）
  - errorRate: +2.3（恶化，dev 和 design 错误率均为 100%）
  - successRate: +0.2（改善，但代价高）
- 建议下次重点关注:
  - **dev agent 的 avgSteps 和 errorRate**：本次已做多项 prompt/workflow 修改，需验证效果。
  - **design agent 的 errorRate**：步数已改善（41.5→34），需观察错误率是否随"启动前错误预判"前置而下降。
  - **e2e agent**：表现优异（17 步，0 错误，successRate=1），可保持当前策略。
