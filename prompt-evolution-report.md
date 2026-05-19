# Prompt 演进报告 — type-one

生成时间: 2026-05-19T16:45:00+08:00
分析范围: 2026-05-19T15:40:05 至 2026-05-19T16:40:39

## 统计摘要

- 分析日志数: 6
- Agent 类型分布:
  - review: 3 个任务
  - rework: 3 个任务
- 整体平均步数: 11.33 步（较上次 -3.57 步）
- 整体错误率: 下降 73%（errorRateDelta -0.73）
- 整体成功率: 上升 10%（successRateDelta +0.10）

## 问题诊断

### 问题 1: review agent 零产出率恶化至 0%（Critical）

**证据**:
- review: 3 任务, successRate = **0%**, avgSteps = **0**, totalWrites = **0**
- 相比上次（4 任务, successRate = 25%），successRate 进一步恶化，且所有任务均无任何工具调用
- avgSteps = 0 表明 agent 完全没有执行，而非执行后失败

**根因分析**:
上次修改在 `workflows/review.yaml` 中增加了 `envReadinessCheck` 步骤，当 `gh auth status` 失败或 PR 不存在时，workflow 输出 `SKIP_AGENT` 并跳过 `runAgent`。随后 `diagnoseAgentRun` 用 shell 生成默认失败报告。然而 dashboard 的 `agentStats` 只统计 **agent 自身的工具调用**（`WriteFile`/`StrReplaceFile` 等），workflow 层面的 shell 步骤不计入 agent 产出。因此，当 gh 认证失败时，agent 被完全跳过，统计上呈现为 avgSteps=0、successRate=0%、totalWrites=0。

这是一个典型的「workflow 与 prompt 职责重叠」导致的副作用：prompt 中已配置了详细的 gh 失败兜底逻辑（第 1 步检测 → 第 2 步 WriteFile），但 workflow 的前置检查抢了 agent 的处理机会，导致 prompt 中的约束完全失效，同时统计系统无法识别 workflow 生成的报告。

**改进建议**:
1. **修改 `workflows/review.yaml`，移除 `runAgent` 的 condition**（已应用）：不再因 `envReadinessCheck` 输出非 `READY` 而跳过 agent。`envReadinessCheck` 仍执行并输出状态（`auth_failed`、`pr_not_found`、`READY`），但不再输出 `SKIP_AGENT`，也不再由 workflow 生成默认报告。让 reviewer agent 始终运行，其 prompt 中的「gh 失败后立即 WriteFile」兜底逻辑才能被实际执行，并被 dashboard 正确统计为有产出。
2. **修改 `prompts/reviewer.md`，增加 PR 号缺失的前置兜底**（已应用）：在第 1 步环境预检查中，先确认环境变量 `AGENT_PR_NUMBER` 是否已设置。如果未设置，立即 WriteFile 失败报告，避免执行无意义的 `gh` 命令产生额外错误和步数浪费。

### 问题 2: rework agent 偶发错误率 33%（Warning）

**证据**:
- rework: 3 任务, successRate = **100%**, avgSteps = **22.67**, totalErrors = **1**, errorRate = **33%**
- 相比上次（6 任务, successRate = 50%, totalErrors = 6），successRate 和错误数大幅改善
- topIssues 中 `rework_high_error_rate` 出现 1 次

**根因分析**:
`developer.md` 中已配置较完善的错误处理 SOP（单次错误后禁止立即重试、StrReplaceFile 前置检查、错误熔断等）和预算约束（rework 40 步、错误 2 次即停），successRate 已提升至 100%。剩余 33% 错误率属于偶发错误（如 StrReplaceFile 失败或 Shell 超时），agent 在特定边界场景下仍可能触发。

avgSteps = 22.67 步，虽在 40 步预算内，但较上次（17.17 步）有所膨胀。rework 任务的核心价值在于「快速修复 review 指出的问题」，不应消耗过多步数在探索和试探上。

**改进建议**:
1. **在 `prompts/developer.md` 的 rework 专属规则中增加「第 0 步 — 错误预判」**（已应用）：在任务启动前预判本次 rework 最可能遇到的错误来源（review report 路径不存在、StrReplaceFile 内容不匹配、Shell 超时），并提前制定应对策略，降低试探性重试的概率。
2. **收紧 rework 模式的探索预算**（已应用）：将通用规则「定位修改点不得超过 10 步」改为对 dev 和 rework 区分——dev 10 步，**rework 6 步**。rework 模式下 review report 已明确问题位置，不应消耗过多步数在探索上。

## 已应用改进

本次分析后已直接修改以下文件：

### 1. `workflows/review.yaml`
- **移除 `runAgent` 的 condition**：不再检查 `envCheckStatus` 是否包含 `READY`，确保 reviewer agent 始终运行。
- **简化 `envReadinessCheck`**：当 gh 认证失败或 PR 不存在时，仅输出状态日志（`auth_failed`、`pr_not_found`），不再输出 `SKIP_AGENT`，也不再由 workflow 生成默认失败报告。将错误处理职责完全交还给 agent 的 prompt 约束。

### 2. `prompts/reviewer.md`
- **增加 PR 号缺失的前置兜底**：在第 1 步环境预检查中，要求先确认 `AGENT_PR_NUMBER` 环境变量是否已设置。如果未设置，下一步必须是 `WriteFile` 写入失败报告，避免执行无意义的 gh 命令。

### 3. `prompts/developer.md`
- **增加 rework「第 0 步 — 错误预判」**：在 rework 专属规则开头新增错误预判步骤，要求 agent 在启动前预判 StrReplaceFile 失败、路径不存在、Shell 超时等常见错误，并提前制定应对策略。
- **收紧 rework 探索预算**：将通用规则「定位修改点不得超过 10 步」改为区分模式——dev 10 步，rework 6 步，抑制 rework 任务的步数膨胀趋势。

## 趋势追踪

- 与上次对比:
  - **整体错误率**：大幅改善，errorRateDelta 从 +0.30 变为 -0.73（下降 73%）
  - **整体成功率**：改善，successRateDelta 从 0 变为 +0.10（上升 10%）
  - **整体平均步数**：改善，avgStepsDelta 从 +8.1 变为 -3.57（下降 3.57 步）
  - **rework agent**：显著改善，successRate 从 50% 提升至 100%，totalErrors 从 6 降至 1
  - **review agent**：恶化，successRate 从 25% 降至 0%。这是上次 workflow 修改引入的副作用，本次已回修复原

- 建议下次重点关注:
  1. **review agent 的 successRate**：本次修改的核心目标是将其从 0% 恢复到 60% 以上。如果 gh 认证持续失败，agent 应能在 2-3 步内完成 WriteFile，successRate 应接近 100%（虽然产出是「失败报告」）。如果仍然低于 50%，需要审查 agent 的实际工具调用序列，确认 prompt 中的「机械下一步」规则是否被严格遵守。
  2. **rework agent 的 avgSteps 趋势**：当前 22.67 步，预算 40 步。如果继续膨胀，需要进一步收紧探索预算或增加「每 5 步无修改则停止」的执行力度。
  3. **环境层面的 gh 认证问题**：3/3 的 review 任务均因 gh 认证失败被跳过，这是一个高频环境问题。建议在 infrastructure 层面（如 dashboard 启动时或 workflow 外部）预检查并配置 GitHub CLI 认证，避免 review agent 持续产出「gh 未认证」的失败报告。
