# Prompt 演进报告 — type-one

生成时间: 2026-05-22T12:40:33+08:00
分析范围: 2026-05-22T10:56:37 至 2026-05-22T12:36:49

## 统计摘要

- 分析日志数: 4
- Agent 类型分布:
  - dev: 1 任务
  - review: 2 任务
  - rework: 1 任务

## 问题诊断

### 问题 1: dev Think 严重超限导致高步数

**证据**:
- dev count=1, avgSteps=57, totalErrors=6, successRate=1
- 该 dev 任务平均 think 高达 **19 次**，而 `developer.md` 明确规定 Think 上限为 **2 次**
- topGitOp 为 `add`，说明 agent 执行了大量 git add 操作（可能伴随反复修改）

**根因分析**:
1. `developer.md` 中虽然规定了 Think 上限（dev 2 次），但缺乏**逐次显式自报**机制。agent 在 reasoning 中没有持续追踪 Think 计数，导致无意识超限。
2. 现有熔断规则（"达到上限后严禁继续"）执行不力，agent 未能在第 2 次 Think 后切换到行动模式。
3. Think 超限直接贡献了约 19 步（占 57 步的 33%），是高步数的首要原因。

**改进建议**:
1. 在 `prompts/developer.md` 的"Think 使用限制"中新增**逐次显式自报**机制：每次 Think 前必须在 reasoning 中输出 `Think 计数: X/2`，不报数不得调用。达到上限后，严禁再次调用 Think，改用 WriteFile 写简要分析后立刻执行。（已应用）
2. 在 `prompts/developer.md` 的"错误预警与熔断"中新增**错误计数显式自报**机制：每次错误后必须在 reasoning 中输出 `错误计数: X/5`，X≥3 时进入只读/收尾模式，X=5 时立即停止。与 rework 模式对齐，提高 agent 对错误的敏感度。（已应用）
3. 在 `workflows/dev.yaml` 中新增 `diagnoseAgentRun` 步骤（仿照 review/rework workflow），在 agent 运行结束后检测日志中的 Think 调用次数和错误次数，超限则生成预警报告，为下次迭代提供数据。（已应用）

### 问题 2: review 违反 Git 禁令且错误率 50%

**证据**:
- review count=2, avgSteps=11, totalErrors=1, successRate=1
- review 任务的 topGitOp 为 `diff`，而 `reviewer.md` 明确**严禁**执行任何本地 `git` 命令
- 错误率 50%（1/2），说明有 1 个 review 任务在执行中遇到了错误

**根因分析**:
1. `reviewer.md` 中的 Git 禁令 enforcement 完全依赖**任务结束前自检**。agent 可能在任务过程中使用了 `git diff`，但在最终报告中未诚实声明。
2. 现有 prompt 的惩罚机制（降级为 NEEDS_MAJOR_FIX）是事后性的，无法在 agent 执行过程中阻止违规行为。
3. workflow 层面缺少对 agent 日志的自动化检测，未能发现违规使用 git 命令的行为。

**改进建议**:
1. 在 `prompts/reviewer.md` 的"快速开始"中将 Git 禁令提升为**前置确认项**：第 0 步之前，agent 必须在 reasoning 中声明"已理解 Git 禁令：本任务严禁任何本地 git 命令"。（已应用）
2. 在 `workflows/review.yaml` 的 `diagnoseAgentRun` 步骤中，增加对 agent 日志的**自动化 Git 禁令检测**：统计 `git status|diff|log|show|blame` 等命令的出现次数，若检测到违规使用，自动将审查报告结论降级为 NEEDS_MAJOR_FIX，并追加违规声明。实现 workflow 层面的强制 enforcement。（已应用）
3. 在 `prompts/reviewer.md` 的"任务结束前强制检查"中，将 Git 禁令自检从"是否使用过"改为"本任务工具调用中是否包含任何 git 命令"，并增加"若已使用，无论审查质量如何，结论强制降级"的明确措辞。（已应用）

### 问题 3: dev 错误率 100%（6 errors/任务）

**证据**:
- dev 任务 totalErrors=6，错误率 100%
- 虽然 successRate=1（最终有产出），但 agent 经历了 6 次错误才完成，试错成本高

**根因分析**:
1. `developer.md` 已有"错误预判"和"常见错误快速诊断表"，但本次任务仍然发生了 6 次错误，说明错误预判的执行效果不够理想。
2. 缺少像 rework 模式那样的**错误计数显式自报**机制，agent 对累计错误的感知不足，容易在错误模式下持续试探。

**改进建议**:
1. 已在问题 1 的改进建议 2 中覆盖：新增错误计数显式自报机制，与 rework 模式对齐。
2. 在 `prompts/developer.md` 的"任务启动检查清单"中，将"错误预判"从"待办清单勾选"改为"必须在 reasoning 中逐条确认已完成"，增加执行刚性。（已应用）

## 已应用改进

本次分析后直接修改了以下文件：

1. **`prompts/developer.md`**
   - 在"Think 使用限制与超限熔断"段落中新增**逐次显式自报**要求：每次 Think 前必须在 reasoning 中输出 `Think 计数: X/2`，不报数不得调用；达到上限后严禁再次调用 Think。
   - 在"错误预警与熔断"段落中新增**错误计数显式自报**要求：每次错误后必须在 reasoning 中输出 `错误计数: X/5`；X=3 时进入只读/收尾模式，X=5 时立即停止并输出最终报告。
   - 在"任务启动检查清单"中，将第 0 步的 4 项错误预判从"待办清单勾选"提升为"必须在 reasoning 中逐条声明已确认"，增强执行刚性。

2. **`prompts/reviewer.md`**
   - 在"快速开始"第 0 步之前新增**Git 禁令前置确认**：agent 必须在执行第一个工具调用前，在 reasoning 中显式声明"已理解 Git 禁令：本任务严禁执行任何本地 git 命令（git status/diff/log/show/blame/checkout/add/commit/push 等）"。
   - 在"任务结束前强制检查"中强化 Git 禁令措辞：将自检问题改为"本任务的所有工具调用中，是否包含任何以 `git` 为命令的 Shell 调用？"，并增加"若答案为是，必须在「风险与建议」章节开头追加 Major 声明，且审查结论强制降级为 NEEDS_MAJOR_FIX"的强调。

3. **`workflows/review.yaml`**
   - 在 `diagnoseAgentRun` 步骤中新增**Git 禁令自动化检测**：统计日志中 `git status|diff|log|show|blame|checkout|add|commit|push` 的出现次数。若检测到违规使用，自动追加审查报告违规声明并将结论降级为 NEEDS_MAJOR_FIX。

4. **`workflows/dev.yaml`**
   - 在 `runAgent` 步骤之后新增 `diagnoseAgentRun` 步骤（仿照 review/rework workflow），检测 agent 日志中的工具调用次数、写操作次数、Think 调用次数和错误次数。若 Think 超过 2 次或错误超过 5 次，生成预警诊断报告，为后续 prompt 迭代提供数据支撑。

## 趋势追踪

- 与上次对比:
  - avgStepsDelta: **-19.25**（显著改善，说明上次对 design 和 dev 的 prompt 收紧生效）
  - errorRateDelta: **-2.75**（显著改善，整体错误率下降）
  - successRateDelta: 0（保持 100% 有产出）
- 建议下次重点关注:
  - **dev agent 的 Think 次数**：本次修改了逐次显式自报机制，需验证 Think 是否被控制在 2 次以内。
  - **review agent 的 Git 禁令遵守情况**：本次在 prompt 和 workflow 层面双管齐下，需验证 topGitOp 是否消失。
  - **dev agent 的 errorRate**：虽然整体 errorRate 下降，但 dev 单任务仍有 6 次错误，需观察错误计数自报机制是否有效减少试探性错误。
