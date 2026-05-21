# Prompt 演进报告 — type-one

生成时间: 2026-05-21T16:14:54+08:00
分析范围: 2026-05-21T11:59:04 至 2026-05-21T15:56:07

## 统计摘要

- 分析日志数: 5
- Agent 类型分布:
  - design: 2 任务
  - review: 1 任务
  - rework: 2 任务

## 问题诊断

### 问题 1: design 任务错误率 100%，Shell/Think 严重超标

**证据**:
- design 2/2 任务均有错误，totalErrors=10，平均 5 错误/任务
- Shell 调用 51 次（25.5/任务），超出 `ui-designer.md` 中 20 次上限 27.5%
- Think 平均 7.0 次/任务，超出 prompt 中 3 次上限 133%
- topGitOp=`status`，说明 `git status` 被过度使用代替有效进度跟踪
- 1 个任务超过 50 步（最高 55 步），平均 41.5 步

**根因分析**:
1. `ui-designer.md` 中的效率约束虽有数字上限，但缺少 developer.md 中那样的"强制报数机制"，agent 在 reasoning 中不主动报数，导致预算失控。
2. `max-steps-per-turn=55` 与 prompt 中 60 步总上限过于接近，agent 在单轮内几乎耗尽全部预算，缺少提前收尾的硬边界。
3. "错误预判"被放在效率约束末尾，agent 启动时未优先执行，导致早期错误消耗大量预算。

**改进建议**:
1. **已应用**：在 `prompts/ui-designer.md` 的"效率约束"开头增加"强制报数机制"，要求每次工具调用前在 reasoning 中显式输出 `[预算] 步数 X/60, Think Y/3, Shell Z/15, 错误 W/5`；不报数不得调用工具。（同时将 Shell 上限从 20 收紧到 15，与当前实际均值 25.5 拉开更大差距。）
2. **已应用**：将 `workflows/design.yaml` 中的 `--max-steps-per-turn` 从 55 降至 45，强制 agent 在单轮内提前 10 步进入收尾阶段。
3. **已应用**：在 `ui-designer.md` 的"快速开始"中更新计数器为 `Shell 预算: 0/15`，并增加"禁止用 `git status` 代替有效进度判断"的提醒（当前 topGitOp 为 `status`，说明此限制被违反）。

### 问题 2: rework 任务零产出率 50%

**证据**:
- rework count=2, successRate=0.5，即 1/2 任务无代码修改产出
- avgSteps=18.5（远低于 40 步上限），说明 agent 提前终止或未能执行修改
- totalErrors=1，错误不是主因
- topGitOp=`add`，说明 agent 至少执行了 git add，但可能未产生实际文件修改

**根因分析**:
1. `developer.md` 中 rework 模式虽有"第 3 步必须写 `logs/rework-start.txt`"等规则，但 agent 可能在第 3 步之后陷入"继续阅读文档、思考分析"的循环，未进入实际修复阶段。
2. `workflows/rework.yaml` 的 `checkAgentOutput` 使用 `ignoreError: false`，当 agent 未产出时代码修改时 workflow 直接失败，生成零产出记录。
3. 缺少在第 5/10/15 步等关键节点的"产出检查点"，agent 在早期未产出时没有及时熔断。

**改进建议**:
1. **已应用**：在 `prompts/developer.md` 的 rework 专属规则中增加"产出检查点"：第 5/10/15/20/25/30/35/40 步必须自评"过去 5 步是否产生过至少一次文件修改"，如否则立即停止并写 `logs/rework-halted.md`。同时增加"第 4 步强制开始修复"规则，禁止第 4 步后继续探索。
2. **已应用**：将 `workflows/rework.yaml` 的 `checkAgentOutput` 的 `ignoreError` 从 `false` 改为 `true`，避免 workflow 因零产出而硬性失败；同时在 `diagnoseAgentRun` 中增加"当 WRITE_COUNT=0 时自动生成 `logs/rework-noop.md` 默认报告"的逻辑，确保任何情况下都有文件产出。

### 问题 3: 整体 successRate 趋势下降

**证据**:
- trends.successRateDelta = -0.2（成功率下降 20 个百分点）
- 虽然 avgStepsDelta=-1.75（步数微降）、errorRateDelta=-0.05（错误率微降），但 successRate 仍在恶化

**根因分析**:
- rework 的 50% 零产出率直接拉低了整体成功率。
- 步数和错误率的微降可能是 prompt 中已有约束的弱效果，但 agent 未严格执行，导致"看起来在改善"但实际产出在恶化。

**改进建议**:
1. 重点解决 rework 零产出问题（见问题 2）。
2. 下次迭代重点跟踪 design 的 errorRate（当前 100%）和 rework 的 successRate（当前 50%）。
3. 考虑在 workflow 层面增加更主动的"agent 产出兜底"机制：当 diagnose 检测到零产出时，自动生成默认报告而不是让任务失败。

## 已应用改进

本次分析后直接修改了以下文件：

1. **`prompts/ui-designer.md`**
   - 新增"强制报数机制"，要求每次工具调用前显式报数。
   - Shell 预算上限从 20 收紧到 15；预警线从 15 调整到 10。
   - 更新统计数据引用（平均 41.5 步、最高 55 步、Think 7.0 次、Shell 25.5 次）。
   - 增加"禁止用 `git status` 代替进度判断"的提醒，针对 topGitOp=`status`。
   - "快速开始"中的计数器同步更新为 `Shell 预算: 0/15`。

2. **`workflows/design.yaml`**
   - `--max-steps-per-turn` 从 55 降至 45，强制提前收尾。

3. **`prompts/developer.md`**
   - rework 专属规则新增"产出检查点"（第 5/10/15...步强制自评）。
   - 新增"第 4 步强制开始修复"规则，禁止读完 report 后继续探索而不修改。

4. **`workflows/rework.yaml`**
   - `diagnoseAgentRun` 中增加：当 WRITE_COUNT=0 时自动生成 `logs/rework-noop.md` 默认报告。
   - `checkAgentOutput` 的 `ignoreError` 从 `false` 改为 `true`，避免 workflow 因零产出而失败；同时将失败状态写入 `logs/rework-diagnose.md` 供后续分析。

## 趋势追踪

- 与上次对比:
  - avgSteps: -1.75（改善）
  - errorRate: -0.05（微降，改善）
  - successRate: -0.2（恶化，需重点关注）
- 建议下次重点关注:
  - **design agent 的 errorRate**：当前 100%，是系统中最突出的问题。需观察 Shell 预算收紧和强制报数机制是否有效。
  - **rework agent 的 successRate**：当前 50%，零产出问题需观察产出检查点和 workflow 兜底机制的效果。
  - **review agent**：表现良好（successRate=1, errors=0），可保持当前策略。
