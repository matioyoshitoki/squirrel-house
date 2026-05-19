# Prompt 演进报告 — type-one

生成时间: 2026-05-19T12:00:13+08:00
分析范围: 2026-05-19T08:03:55+08:00 至 2026-05-19T12:00:13+08:00

## 统计摘要

- 分析日志数: 7
- Agent 类型分布:
  - review: 2 个任务
  - rework: 5 个任务
- 趋势变化 (与上期对比):
  - avgStepsDelta: +5.18
  - errorRateDelta: +0.14
  - successRateDelta: +0.30

## 问题诊断

### 问题 1: rework 任务产出率仍过低（successRate 20%）

**证据**: 5 个 rework 任务，successRate=0.2，avgSteps=4.8。totalWrites=7 全部集中在 1 个任务中，其余 4 个任务零产出。对比上期（successRate=0%, avgSteps=1），虽有改善但仍不可接受。

**根因分析**:
- developer.md 上期已增加"3 步无写操作即熔断"和"强制 WriteFile"规则，但 80% 的任务仍在极早阶段终止。
- avgSteps=4.8 接近启动确认所需的 2 步（读 report + 写 start.txt），说明 agent 在完成启动后就退出了，没有进入实际修复阶段。
- 4 个零产出任务可能的原因是：(a) review report 不存在或无法读取，agent 按"停止并报告"规则以纯文本终止；(b) review report 无 Blocking/Major 问题，agent 直接纯文本输出"无需修复"；(c) agent 在启动确认后因步数/错误顾虑提前放弃。
- "停止并报告"的表述在 developer.md 中被 agent 理解为纯文本输出，绕过了 WriteFile 强制约束。

**改进建议**:
1. **统一终止规则为 WriteFile**: 在 `prompts/developer.md` 的"效率约束"开头增加统一规则：所有"停止并报告"一律指执行 WriteFile 写入报告文件后结束。特别针对 rework 模式，将"review report 无法读取，立即停止并报告"明确改为"立即停止并执行 WriteFile 写 `logs/rework-halted.md`"。（已修改）
2. **增加无问题路径和步数下限**: 在 `prompts/developer.md` 的 rework 专属规则中增加：(a) review report 无 Blocking/Major 问题时必须 WriteFile 写 `logs/rework-noop.md`；(b) rework 任务少于 5 步结束时必须回退检查。防止 agent 在读完 report 后未经 WriteFile 直接结束。（已修改）
3. **workflow 前置拦截**: 在 `workflows/rework.yaml` 的 `envReadinessCheck` 中，如果 review report 不存在，直接生成默认的 `logs/rework-noop.md` 并跳过 agent 运行，避免启动无意义的零产出任务。（已修改）

### 问题 2: review 任务步数膨胀（avgSteps 17.5，+10.5 vs 上期）

**证据**: 2 个 review 任务，avgSteps=17.5，ReadFile 33 次（16.5/任务），Shell 23 次（11.5/任务），totalErrors=1。虽然 successRate 从 33.3% 提升到 100%，但步数翻倍，Shell 调用远超合理范围。

**根因分析**:
- reviewer.md 上期修改成功解决了零产出问题（successRate 33%→100%），但 agent 在审查过程中过度探索。
- Shell 23 次/2 任务 = 11.5 次每任务，显著超出合理范围（正常 4-6 次）。可能是 `gh pr diff` 重复调用、测试命令多次执行或错误后重试。
- ReadFile 33 次/2 任务 = 16.5 次每任务，说明 agent 读取了大量规范文档，缺少读取预算约束。
- 与 developer.md 不同，reviewer.md 没有 Shell/ReadFile 的硬性上限，agent 缺乏预算意识。

**改进建议**:
1. **增加工具调用预算**: 在 `prompts/reviewer.md` 中增加硬性预算：Shell 每任务最多 8 次（当前 11.5），ReadFile 每任务最多 12 次（当前 16.5）。达到预警线时强制进入收尾阶段。（已修改）
2. **缓存 gh diff 结果**: 在 reviewer.md 的"快速开始"第 3 步中，要求 agent 将 `gh pr diff` 结果写入临时文件，后续引用时直接 ReadFile，禁止重复调用 `gh pr diff`。（已修改）

### 问题 3: 整体错误率上升（errorRateDelta +0.14）

**证据**: trends.errorRateDelta=0.142857，review 任务错误率 50%（1/2）。

**根因分析**:
- review 任务的 1 个错误可能来自 gh 命令失败、文档路径不存在或 Shell 输出过大。
- 虽然 reviewer.md 已有错误处理规则，但 agent 在遇到错误时仍可能消耗额外步数尝试恢复（avgSteps 17.5 佐证）。
- Shell/ReadFile 预算的缺失使 agent 没有足够强的约束来避免错误路径上的反复试探。

**改进建议**:
1. 通过 reviewer.md 新增的 Shell/ReadFile 预算约束，强制 agent 在工具调用达到上限前进入收尾阶段，减少错误发生的机会。
2. 下次迭代建议 dashboard 增加"错误类型分布"统计（如 gh 失败、文件不存在、输出过大等），帮助更精准定位 review 错误的根因。

## 已应用改进

1. **`prompts/developer.md`**:
   - 新增"统一终止规则"：所有"停止并报告"一律指执行 WriteFile 写入报告文件，禁止纯文本输出。
   - rework 专属规则：review report 无法读取时，明确必须 WriteFile 写 `logs/rework-halted.md`。
   - 新增"review report 无 Blocking/Major 问题"处理路径：必须 WriteFile 写 `logs/rework-noop.md`。
   - 新增"rework 步数下限"：少于 5 步结束时必须回退检查。

2. **`prompts/reviewer.md`**:
   - 新增"工具调用预算"：Shell 每任务最多 8 次（预警线 6 次），ReadFile 每任务最多 12 次（预警线 10 次）。
   - 要求 `gh pr diff` 结果写入临时文件缓存，禁止重复调用。

3. **`workflows/rework.yaml`**:
   - `envReadinessCheck` 增加 `outputKey: envCheckStatus`，当 review report 不存在时直接生成 `logs/rework-noop.md` 并输出 `SKIP_AGENT`。
   - `runAgent`、`checkAgentOutput`、`preCommitChecks` 增加 `condition`，仅在 `envCheckStatus == READY` 时执行，避免无 review report 时的无效启动和失败。

## 趋势追踪

- **与上次对比**: 改善（successRateDelta=+0.30，review 从 33%→100%，rework 从 0%→20%），但代价是 avgStepsDelta=+5.18 和 errorRateDelta=+0.14。review 的零产出问题已根治，rework 仍需观察。
- **建议下次重点关注**:
  - rework 任务在 prompt 修改后是否突破 20% 产出率（目标：>50%）
  - review 任务在 Shell/ReadFile 预算约束下 avgSteps 是否回落到 12 以下
  - review 任务的 errorRate 是否因预算约束而下降
