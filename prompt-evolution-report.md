# Prompt 演进报告 — type-one

生成时间: 2026-05-22T14:46:58+08:00
分析范围: 2026-05-22T13:40:23 到 2026-05-22T14:41:37

## 统计摘要

- 分析日志数: 17
- Agent 类型分布:
  - `review`: 9 个任务
  - `rework`: 8 个任务

| 指标 | review | rework |
|------|--------|--------|
| 平均步数 | 9.11 | 6.375 |
| 成功率 | 100% | 87.5% |
| 总错误数 | 1 | 0 |
| 平均 ReadFile/任务 | ~9.4 | ~1.75 |
| 平均 Shell/任务 | ~3.9 | ~1.625 |
| 常用 Git 操作 | commands | checkout |

## 问题诊断

### 问题 1: reviewer.md 引用的历史统计数据与当前报告严重脱节
**证据**: 
- reviewer.md 第 303 行写道："当前统计平均 29 次/任务，严重超过上限"，而本次 review 的 Shell 平均仅 **3.9 次/任务**。
- reviewer.md 第 304 行写道："当前统计平均 16 次/任务"，而本次 review 的 ReadFile 平均仅 **9.4 次/任务**。
- reviewer.md 第 346 行写道："统计报告已记录 review 任务的 topGitOp 为 diff"，而本次报告 review 的 topGitOp 实际为 **commands**（且 review.yaml workflow 本身在执行 `git branch`、`git checkout`、`git worktree remove` 等命令，统计口径可能包含 workflow 层面的 git 操作）。

**根因分析**: 
prompt 中嵌入了具体的历史统计数据作为行为约束的背景论据，但当实际统计数据已经大幅改善后，这些过时的引用会产生两种负面效果：
1. **虚假紧迫感**：agent 可能因看到"严重超过上限"而过度紧缩行为，反而影响审查深度。
2. **认知失调**：agent 在执行 Git 禁令自检时，看到 prompt 说"topGitOp 为 diff"，但实际自身并未调用 git，可能产生困惑；或更糟，误以为 workflow 的 git 命令是自己调用的，导致错误的自我降级。

**改进建议**:
1. ✅ **已应用** — 将 reviewer.md 中的"当前统计平均 29/16 次"改为"历史统计曾平均...最新统计已降至约 4/9 次/任务"，保留上限约束的严肃性，同时让 agent 了解最新效率水平。（涉及文件: `prompts/reviewer.md`，修改方式: StrReplaceFile）
2. ✅ **已应用** — 在 Git 禁令自检段落中增加显式说明："统计报告中的 `topGitOp` 可能包含 workflow 自动执行的 git 命令，agent 自检时应严格依据自身实际调用的 Shell 命令判断"，避免误报导致的自我降级。（涉及文件: `prompts/reviewer.md`，修改方式: StrReplaceFile）
3. **后续建议** — 建立 prompt 中统计引用的版本化标注（如 `[统计截止: 2026-05-22]`），并在每次 prompt-evolution 分析时强制扫描所有具体数字引用，防止再次过期。

### 问题 2: developer.md 中 rework 模式的步骤编号存在双重定义冲突
**证据**: 
- rework successRate = **87.5%**，存在一个零产出任务；rework 平均步数仅 **6.375**，说明大部分任务快速结束。
- developer.md 第 181-182 行规定："第 3 步必须写 `logs/rework-start.txt`"。
- developer.md 第 187 行规定：如果 review report 无 Blocking/Major 问题，"第 3 步必须写 `logs/rework-noop.md`"。

**根因分析**: 
同一个"第 3 步"被赋予了两个互斥的 WriteFile 目标（rework-start.txt vs rework-noop.md）。当 review report 无 Blocking/Major 问题时，agent 读完 report（第 2 步）后面临二选一：
- 若遵守"第 3 步写 start.txt"，则 noop.md 被推迟到第 4 步，但 prompt 明确说"第 3 步必须写 noop.md"，导致 agent 可能困惑或遗漏。
- 若直接写 noop.md 而跳过 start.txt，则违反了"第 3 步必须写 start.txt"的铁律。
这种步骤冲突在 rework 这种低步数（平均 6.375 步）的场景下极易导致 agent 提前以纯文本结束任务，从而触发零产出统计。

**改进建议**:
1. ✅ **已应用** — 将"review report 为空/无问题"的处理从"第 3 步"改为"第 4 步"，明确顺序：第 1 步确认 report 存在 → 第 2 步读取 report → 第 3 步写 start.txt → 第 4 步写 noop.md（如无需修复）。（涉及文件: `prompts/developer.md`，修改方式: StrReplaceFile）
2. **后续建议** — 在 rework.yaml 的 `diagnoseAgentRun` 步骤中，对"agent 日志存在但 WriteFile 次数为 0"的情况增加更详细的分类诊断（如区分"未启动修复"、"report 无问题但未写 noop"、"StrReplaceFile 全部失败"），帮助精确定位零产出根因。（涉及文件: `workflows/rework.yaml`，修改方式: 增强 diagnose 脚本）

### 问题 3: topIssues 为 null，失败模式捕获机制未生效
**证据**: 
- 报告 `topIssues: null`，但 rework 存在一个零产出任务，review 也存在 1 个总错误。

**根因分析**: 
失败模式聚合逻辑可能未在日志中提取和归类常见的错误/失败模式（如 `old string not found`、`file_not_found`、`outside the workspace`、`exit code` 等）。没有 topIssues，prompt-evolution agent 难以做跨任务的根因关联。

**改进建议**:
1. **后续建议** — 检查 dashboard 的统计生成逻辑（`cmd/dashboard/*.go`），确保在汇总日志时提取并归类常见的 `is_error=true` 模式、零产出原因、以及 workflow 层面的 SKIP_AGENT / 失败原因，填充 `topIssues` 字段。（涉及文件: `cmd/dashboard/*.go`，修改方式: 增强日志模式提取）
2. **后续建议** — 作为兜底，在 rework.yaml 的 `diagnoseAgentRun` 和 review.yaml 的 `diagnoseAgentRun` 中增加结构化的失败原因标签（如 `FAIL_REASON=str_replace_fail` / `fail_reason=no_write_op`），便于下游统计直接读取。

## 已应用改进

1. **`prompts/reviewer.md`**:
   - 更新 Shell 预算说明：将"当前统计平均 29 次/任务"改为"历史统计曾平均 29 次/任务，最新统计已降至约 4 次/任务"。
   - 更新 ReadFile 预算说明：将"当前统计平均 16 次/任务"改为"历史统计曾平均 16 次/任务，最新统计约 9 次/任务"。
   - 更新 Git 禁令自检段落：增加说明，提示 agent 统计报告中的 `topGitOp` 可能包含 workflow 自动执行的 git 命令，自检应基于自身实际调用记录。

2. **`prompts/developer.md`**:
   - 修正 rework 模式"review report 为空/无问题"的处理步骤：将"第 3 步必须写 `logs/rework-noop.md`"改为"完成第 3 步写 `logs/rework-start.txt` 之后的下一步（第 4 步）必须写 `logs/rework-noop.md`"，消除步骤编号冲突。

## 趋势追踪

- **与上次对比**: 
  - 平均步数下降 **8.53**（大幅优化，说明预算约束生效）
  - 错误率下降 **0.37**（显著改善）
  - 成功率提升 **1.26%**（温和上升，整体稳定）
  - 趋势全面向好，prompt 中施加的预算硬上限和步数预警机制正在产生效果。

- **建议下次重点关注**:
  1. **dev agent**: 本次统计周期内无 dev 任务数据，无法评估 developer.md 中 dev 模式约束的有效性。建议下次关注 dev 的 avgSteps、successRate 和 topGitOp。
  2. **rework 零产出根因**: 虽然已修复步骤冲突，但仍需观察下一轮统计中 rework successRate 是否回升至 100%。
  3. **topIssues 数据完整性**: 推动 dashboard 统计逻辑补充失败模式提取，使 prompt-evolution 分析能基于更细粒度的失败根因做关联。
