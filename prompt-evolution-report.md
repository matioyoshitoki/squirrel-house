# Prompt 演进报告 — type-one

生成时间: 2026-05-21T12:00:00+08:00
分析范围: 2026-05-20T18:17:14 至 2026-05-21T11:59:03

## 统计摘要

- 分析日志数: 4
- Agent 类型分布:
  - design: 1 任务
  - dev: 1 任务
  - review: 2 任务

| Agent | 任务数 | 平均步数 | 总错误数 | 成功率 | 常用 Git 操作 |
|-------|--------|----------|----------|--------|---------------|
| design | 1 | 42 | 4 | 1.0 | status |
| dev | 1 | 55 | 3 | 1.0 | status |
| review | 2 | 11 | 2 | 1.0 | diff |

趋势变化（与上次对比）:
- 平均步数变化: +12.18（恶化）
- 错误率变化: +1.25（恶化）
- 成功率变化: +0.14（改善）

## 问题诊断

### 问题 1: design 与 dev 错误率 100%，步数逼近上限

**证据**:
- design: 1/1 任务出现 4 个错误，avgSteps=42（接近 60 步上限）
- dev: 1/1 任务出现 3 个错误，avgSteps=55（接近 60 步上限）
- topIssues 中 `design_high_error_rate` 和 `dev_high_error_rate` 均标记为 critical

**根因分析**:
1. **外部硬性边界缺失**: `workflows/design.yaml` 中 `--max-steps-per-turn 10000` 与 `ui-designer.md` 的 60 步上限严重脱节；`workflows/dev.yaml` 中 `--max-steps-per-turn 65` 也高于 `developer.md` 的 60 步上限。外部参数未对 agent 形成有效约束。
2. **错误预判未真正前置**: prompt 中虽有"错误预判"段落，但未要求在任务启动后的前 3 步内**强制确认**错误预判清单。agent 在遇到文件不存在、StrReplaceFile 失败时仍然反复试探，消耗步数并累积错误。

**改进建议**:
1. **统一 workflow 步数上限**（已应用）: 将 `design.yaml` 的 `--max-steps-per-turn` 从 `10000` 降至 `55`，`dev.yaml` 从 `65` 降至 `55`，`review.yaml` 从 `50` 降至 `40`，使外部参数与 prompt 的内部预算一致并预留收尾余量。
2. **强化错误预判为强制检查**（已应用）: 在 `prompts/ui-designer.md` 中将"错误预判"升级为启动后必须完成的强制 SOP，并新增错误类型快速诊断表（file_not_found → 确认路径后跳过；StrReplaceFile 失败 → ReadFile 确认后重试 1 次；Shell 超时 → 禁止重试）。
3. **在 developer.md 中前置环境确认**: 要求 dev 任务在第 1-2 步内用 `test -f` 批量确认关键路径（AGENTS.md、package.json、Makefile 等）是否存在，避免因路径假设错误导致早期步数浪费。

### 问题 2: dev 任务 Think 次数严重超标（11 次 vs 限额 2 次）

**证据**:
- topIssues 中 `dev_high_steps` 描述: dev 有 1 个任务超过 50 步，平均 think 11.0 次
- `developer.md` 明确设定 Think 上限为 dev **2 次**
- 实际 think 次数是限额的 **5.5 倍**

**根因分析**:
- prompt 中的"达到上限后严禁继续使用"属于**软性禁令**，agent 在面临复杂决策时仍然习惯性调用 Think，因为 prompt 没有给出**超限后可执行的替代路径**。agent 不知道"不能 think 了，我应该做什么"。
- `developer.md` 第 21 行的统计警告（"预算约束被全面突破"）虽提及超标，但未在 Think 上限条款中直接引用 think 超标数据，agent 对 think 失控的严重性认知不足。

**改进建议**:
1. **增加超限后的强制降级路径**（已应用）: 在 `developer.md` 中明确规定：达到 2 次 Think 后，下一次 tool call 必须是 ReadFile/StrReplaceFile/WriteFile/Shell（仅限验证/git）之一，严禁再次 Think。如确实需要思考，改用 WriteFile 写简要分析到 `logs/dev-think-dump.md`（计入 WriteFile，不额外消耗 think 预算），然后立即继续执行。
2. **在铁律中引用超标数据强化认知**（已应用）: 在 `developer.md` 的 Think 上限描述中补充引用当前统计（"当前统计 dev 任务平均 think 高达 11 次/任务"），使 agent 明确感知到违规的普遍性。

### 问题 3: review 任务违反 Git 禁令（topGitOp = diff）

**证据**:
- review 统计的 `topGitOp` 为 `diff`
- `reviewer.md` 元原则第 3 条明确声明: "review 任务**严禁执行任何本地 `git` 命令**（包括 `git status`、`git diff`...）"
- review 任务错误率 50%（1/2 任务出错）

**根因分析**:
- 虽然 prompt 将 Git 禁令列为"绝对铁律"，但缺少**任务结束前的强制自检环节**，agent 未在最终阶段回顾自己是否违规。
- 违规后果描述（"直接导致审查结论不可信并触发熔断"）停留在抽象层面，没有要求 agent 在**输出物（审查报告）中显式声明**合规性，导致违规难以被后续流程发现。

**改进建议**:
1. **追加显式合规声明要求**（已应用）: 在 `reviewer.md` 的 Git 禁令条款中追加："任务结束前必须在最终报告中显式声明：'本审查全程未使用任何本地 git 命令。'"
2. **增加 Git 禁令自检条目**（已应用）: 在 `reviewer.md` 执行约束中新增第 16 条：任务结束前必须自检是否调用过本地 git 命令；如果违规，必须在报告「风险与建议」章节追加 Major 级别声明，并将审查结论强制降级为 NEEDS_MAJOR_FIX。同时直接引用统计证据（"统计报告已记录 review 任务的 topGitOp 为 diff，说明此禁令被违反"），形成数据驱动的约束压力。

## 已应用改进

本次分析后已直接修改以下文件：

1. **`workflows/design.yaml`**
   - 修改: `--max-steps-per-turn 10000` → `--max-steps-per-turn 55`
   - 原因: 与 `ui-designer.md` 的 60 步上限保持一致并预留 5 步收尾余量。

2. **`workflows/dev.yaml`**
   - 修改: `--max-steps-per-turn 65` → `--max-steps-per-turn 55`
   - 原因: 与 `developer.md` 的 60 步上限保持一致并预留余量。

3. **`workflows/review.yaml`**
   - 修改: `--max-steps-per-turn 50` → `--max-steps-per-turn 40`
   - 原因: 与 `reviewer.md` 的 45 步上限保持一致并预留余量。

4. **`prompts/reviewer.md`**
   - 修改 1: 在"绝对铁律 —— Git 禁令"中追加任务结束前的显式合规声明要求。
   - 修改 2: 在执行约束中新增第 16 条"Git 禁令自检"，要求任务结束前自检是否调用过本地 git 命令，违规时强制降级审查结论并引用统计证据。

5. **`prompts/developer.md`**
   - 修改 1: 在"铁律"的 Think 上限描述中补充超标数据引用（11 次/任务），强化严重性认知。
   - 修改 2: 将"Think 使用限制"升级为"Think 使用限制与超限熔断"，明确超限后的 tool call 白名单（ReadFile/StrReplaceFile/WriteFile/Shell）以及 WriteFile 降级路径（`logs/dev-think-dump.md`）。

6. **`prompts/ui-designer.md`**
   - 修改: 将"错误预判"升级为启动后强制检查清单，并新增错误类型快速诊断表，涵盖 file_not_found、StrReplaceFile 失败、command not found、Shell 超时、chunk exceed the limit 等 5 类高频错误的强制处理 SOP。

## 趋势追踪

- **与上次对比**: 平均步数（+12.18）和错误率（+1.25）均呈恶化趋势，说明 prompt 中的软性预算约束在 agent 实际执行中持续被突破。成功率（+0.14）略有改善，但样本量仅 4 个日志，统计意义有限。
- **建议下次重点关注**:
  1. **dev 任务的 Think 合规性**: 已应用外部参数收紧（55 步/turn）和 prompt 降级路径，需观察下一周期 think 次数是否从 11 次降至 2 次以内。
  2. **design 任务的错误预判执行效果**: 已应用错误类型快速诊断 SOP，需观察 design 错误率是否从 100% 下降。
  3. **review 任务的 Git 禁令合规性**: 已增加强制自检和显式声明，需验证下一周期 topGitOp 是否从 `diff` 变为 `gh` 相关命令或消失。
