# Prompt 演进报告 — type-one

生成时间: 2026-05-29 11:32:14+08:00
分析范围: 2026-05-29 10:28:17 至 2026-05-29 11:32:14

## 统计摘要

- 分析日志数: 10
- Agent 类型分布:
  - review: 5 任务，avgSteps=18.2，totalErrors=10，successRate=100%
  - rework: 5 任务，avgSteps=18.8，totalErrors=4，successRate=100%

## 问题诊断

### 问题 1: review 任务错误率 80%，Shell 调用逼近上限
**证据**: review agent 5 个任务共 10 个错误，topIssues 标记为 `critical`；Shell 调用 68 次（平均 13.6 次/任务），距 15 次上限仅剩 1.4 次余量；topGitOp 为 `commands`。
**根因分析**: `prompts/reviewer.md` 第 1 步环境预检查存在**自相矛盾**：前文明确声明 "workflow 的 `envReadinessCheck` 已验证 `gh` 认证和 PR 可访问性，agent 无需再次执行 `gh auth status` 或 `gh pr view`"，但同一段落紧接着又指令 "然后运行 `gh auth status` 和 `gh pr view`"。这导致 READY 状态下的 review 任务**固定浪费 2 次 Shell 调用**，且在网络波动时引入不必要的 `is_error`（401/403/404/超时）。
**改进建议**:
1. **已应用**：删除 `reviewer.md` 第 1 步中 READY 状态下 "然后运行 `gh auth status` 和 `gh pr view`" 的冗余指令，改为 "禁止再次执行……直接进入第 2 步"。每个 READY 任务节省 2 次 Shell，直接消除一个主要错误来源。（涉及文件: `prompts/reviewer.md`，修改方式: StrReplaceFile）
2. 建议后续迭代评估：将 `gh pr diff <number> --name-only` 与 `gh pr diff <number> | head -100` 的获取逻辑合并为单次调用（先取完整 diff 再解析文件名），可再节省 1 次 Shell/任务。（涉及文件: `prompts/reviewer.md`）

### 问题 2: rework 任务 Shell 预算突破，prompt 存在内部矛盾
**证据**: rework agent 平均 Shell 调用 10.4 次/任务，已突破 `developer.md` 规定的 10 次上限；5 个任务共 4 个错误，topIssues 标记为 `warning`。
**根因分析**: `prompts/developer.md` 中 rework 模式存在**三处互相冲突的约束**：
- A. "rework 任务一律直接使用 ReadFile/StrReplaceFile，让错误处理捕获缺失，**禁止用 Shell 做文件存在性检查**"
- B. "rework 模式下，额外执行 `Shell test -f "$AGENT_REVIEW_REPORT"` 确认 review report 存在"
- C. "rework 启动确认：第 1 步必须先确认 review report 可访问：**运行 Shell test -f**"

约束 A 明确禁止，而 B、C 强制要求执行。这导致每个 rework 任务**固定浪费 1 次 Shell** 在 `test -f` 上，加上 `git status`（1 次）、git add/commit/push（3 次）、验证命令（1 次），剩余用于修复的 Shell 预算不足 4 次，极易触碰 10 次上限。
**改进建议**:
1. **已应用**：将 B、C 两处的 `Shell test -f "$AGENT_REVIEW_REPORT"` 统一改为 `ReadFile "$AGENT_REVIEW_REPORT"`，让 ReadFile 的错误处理捕获文件缺失，彻底消除矛盾并节省 1 次 Shell/任务。（涉及文件: `prompts/developer.md`，修改方式: StrReplaceFile，共 2 处）
2. 建议将 rework Shell 上限从 10 次提升至 12 次，或把 `git status` 的启动确认下沉到 workflow 中（由 `workflows/rework.yaml` 的 `envReadinessCheck` 统一输出 worktree 状态到 JSON），进一步释放 agent 预算。（涉及文件: `prompts/developer.md` / `workflows/rework.yaml`）

### 问题 3: 趋势恶化 —— 错误率_delta +0.6，步数_delta +5.3
**证据**: trends 显示 errorRateDelta = +0.6，avgStepsDelta = +5.3，说明相比上一轮统计，错误率和平均步数均在显著恶化。
**根因分析**: 即使 successRate 提升了 0.4（得益于 WriteFile 强制兜底机制），agent 仍在执行大量无产出的试探性调用。步数增加 5.3 意味着 agent 可能在错误发生后尝试替代路径或重试，而不是按 prompt 规定的"单次错误后禁止立即重试，连续 2 次错误立即熔断"快速退出。
**改进建议**:
1. 在 `prompts/developer.md` 和 `prompts/reviewer.md` 的"错误预判"章节中，增加一条**强制诊断轮次**：工具调用返回 `is_error=true` 后，必须在下一个 reasoning 中显式输出 "错误诊断：失败原因 = [file_not_found / command_not_found / chunk_exceed / timeout / auth]，下一步动作 = [skip / retry_once / halt]"，禁止不分析直接重试。（涉及文件: `prompts/developer.md`、`prompts/reviewer.md`）
2. 建议 dashboard 在生成统计报告时，增加 `errorTypeBreakdown` 维度（按 file_not_found / command_not_found / chunk_exceed / timeout / auth 分类），以便下次演进精确定位主要错误来源。（涉及文件: `cmd/dashboard/*.go`）

## 已应用改进

1. **`prompts/reviewer.md`** — 删除 READY 状态下的冗余 `gh auth status` + `gh pr view` 指令
   - 旧文本: "然后运行 `gh auth status` 和 `gh pr view`……"
   - 新文本: "**禁止**再次执行 `gh auth status` 或 `gh pr view`——workflow 的 `envReadinessCheck` 已验证……直接进入第 2 步。"
   - 预期收益: 每个 READY review 任务减少 2 次 Shell 调用，消除一个高频错误来源。

2. **`prompts/developer.md`** — 消除 rework 模式的 `test -f` 矛盾（2 处）
   - 第 1 处（任务启动检查清单）: 将 `Shell test -f "$AGENT_REVIEW_REPORT"` 改为 `ReadFile "$AGENT_REVIEW_REPORT"`。
   - 第 2 处（rework 启动确认）: 将 "运行 `Shell test -f`……第 2 步 `ReadFile`" 改为 "第 1 步 `ReadFile`……第 2 步分析 review report"。
   - 预期收益: 每个 rework 任务减少 1 次 Shell 调用，消除 prompt 内部矛盾，降低 agent 困惑度。

## 趋势追踪

- 与上次对比:
  - **avgStepsDelta: +5.3** — 步数恶化，说明 agent 仍在试探或重试。
  - **errorRateDelta: +0.6** — 错误率显著恶化，是当前最紧迫的问题。
  - **successRateDelta: +0.4** — 产出率改善，主要得益于 WriteFile 强制兜底机制生效，但"有产出"不等于"高质量"。
- 建议下次重点关注:
  - **review agent 的错误类型分布**：10 个错误具体是 file_not_found、gh 命令失败、还是 chunk exceed？需要更细粒度的错误分类数据。
  - **rework Shell 使用率**：本次修改后预计平均 Shell 从 10.4 降至 9.4 次/任务，需验证是否还有突破 10 次上限的任务。
  - **reviewer.md 的 diff 获取策略**：`gh pr diff --name-only` + `gh pr diff` 两次调用是否可以合并，以进一步降低 Shell 压力和错误率。
