# Prompt 演进报告 — type-one

生成时间: 2026-05-29 13:52:38+08:00
分析范围: 2026-05-29 11:32:14 至 2026-05-29 13:52:38

## 统计摘要

- 分析日志数: 7
- Agent 类型分布:
  - review: 4 任务，avgSteps=16.5，totalErrors=8，successRate=75%，Shell=45（11.25/任务）
  - rework: 3 任务，avgSteps=17.33，totalErrors=1，successRate=100%，Shell=32（10.67/任务）

## 问题诊断

### 问题 1: review 任务错误率仍高（每任务 2 个错误），且出现 1 例零产出
**证据**: review 4 个任务共 8 个错误（平均 2 次/任务），successRate 从上次 100% 降至 75%。topIssues 标记 `review_high_error_rate` 为 critical，`review_no_output` 为 warning。Shell 调用 45 次（11.25/任务），接近 15 次上限。
**根因分析**: `prompts/reviewer.md` 内部存在**三处自相矛盾**，导致 agent 在"遵守禁令"和"执行建议"之间摇摆，产生不必要的错误和 Shell 浪费：
1. **错误预判 vs 执行约束的矛盾**：第 280 条错误预判建议"读取前先用 `test -f` 确认"，但第 318 条执行约束明确"禁止先用 Shell `test -f`"——agent 可能先执行 test -f 浪费 Shell，或直接 ReadFile 触发 file_not_found 错误。
2. **gh diff 双次调用**：第 3 步要求执行 `gh pr diff --name-only` 和 `gh pr diff | head -100` 两次调用，而 `gh pr diff` 本身已是主要错误来源（网络/认证/chunk exceed）。
3. **第 1 步的冗余 `echo $AGENT_PR_NUMBER`**：workflow 的 `envReadinessCheck` 已验证 PR 可访问性，agent 仍被建议用 Shell echo 确认环境变量，浪费 1 次 Shell/任务。
4. **第 4 步的 `ls` 搜索历史报告**：用 Shell `ls` 搜索历史报告消耗 1 次 Shell，且当 logs 目录不存在时会报错。
**改进建议**:
1. **已应用**：删除 reviewer.md 第 280 条错误预判中的"读取前先用 `test -f` 确认"，统一为"直接执行 ReadFile，若返回 file_not_found 则标记未验证并跳过，禁止先用 test -f"。消除 prompt 内部矛盾，降低 agent 困惑度。（涉及文件: `prompts/reviewer.md`）
2. **已应用**：合并第 3 步的两次 `gh pr diff` 为单次调用 `gh pr diff <number> | head -200`，从 diff 内容中自行解析文件名列表。每个任务减少 1 次 Shell 调用和 1 个潜在错误来源。（涉及文件: `prompts/reviewer.md`）
3. **已应用**：删除 READY 状态下"`echo $AGENT_PR_NUMBER` 确认"的冗余指令，改为"禁止执行 echo，直接进入第 2 步"；删除第 4 步的 `ls` 搜索历史报告，改为"检查上下文是否提供路径，未提供则视为首次审查，禁止执行 ls"。每个任务减少 2 次 Shell 调用。（涉及文件: `prompts/reviewer.md`）

### 问题 2: rework Shell 预算持续突破（10.67/任务 > 上限 10）
**证据**: rework 3 个任务共 32 次 Shell（平均 10.67/任务），突破 developer.md 规定的 10 次上限。虽然 totalErrors 仅 1 个（显著改善），但 Shell 超标意味着 agent 在压缩其他操作空间，或触发隐性错误。
**根因分析**: `prompts/developer.md` 存在**通用规则与 rework 专属规则的系统性冲突**：
1. 通用规则第 7 条要求"读取任何文件前先用 `test -f` 确认"，但 rework 专属规则第 28 条明确"严禁用 Shell 做文件存在性检查"。
2. 常见错误 SOP 第 201、203 条在 rework 模式下仍要求"用 `test -f` 确认路径"和"用 `which` 确认命令存在性"，与第 189 条"rework 模式下 grep/find/ls/test -f 等探索命令总计不得超过 3 次"直接冲突。
3. 第 1 步"确认环境"要求执行 `git status`，但 rework 模式下该操作可由 workflow 预检替代。每个 rework 任务固定消耗 1 次 `git status` + 3 次 git add/commit/push + 1 次验证命令 = 5 次 Shell，剩余预算仅 5 次用于实际修复，极易触碰上限。
**改进建议**:
1. **已应用**：将通用规则第 7 条拆分为 dev/rework 两个独立子条款：dev 模式允许 test -f，rework 模式"严禁先用 Shell 确认文件存在性，一律直接 ReadFile/StrReplaceFile"。同时修改常见错误 SOP（第 176、201、203 条），在 rework 模式下删除 `test -f` 和 `which` 的强制要求。消除系统性冲突，预计每个 rework 任务减少 1-2 次 Shell。（涉及文件: `prompts/developer.md`）
2. **已应用**：将 rework Shell 上限从 10 次提升至 12 次，匹配当前实际消耗（10.67/任务）并预留安全余量。（涉及文件: `prompts/developer.md`）
3. **已应用**：将 `git status` 的启动确认下沉到 workflow：在 `workflows/rework.yaml` 的 `envReadinessCheck` 中增加 `git_branch` 和 `git_status_lines` 到 `rework-env-check.json`。在 `developer.md` 中修改 rework 启动流程：第 1 步先 `ReadFile "$AGENT_ENV_CHECK_FILE"` 读取预检信息，**禁止**执行 `git status`。预计每个 rework 任务减少 1 次 Shell。（涉及文件: `workflows/rework.yaml`、`prompts/developer.md`）

## 已应用改进

1. **`prompts/reviewer.md`** — 消除内部矛盾 + 优化 Shell 使用
   - 删除错误预判中的 `test -f` 建议，与执行约束保持一致。
   - 合并 `gh pr diff --name-only` + `gh pr diff` 为单次调用，减少 1 次 Shell/任务。
   - 删除 READY 状态下的 `echo $AGENT_PR_NUMBER` 和第 4 步的 `ls` 搜索，减少 2 次 Shell/任务。
   - 预期收益: 每个 review 任务减少 3 次 Shell 调用（从 11.25 降至约 8.25），消除 2 个高频错误来源。

2. **`prompts/developer.md`** — 消除 dev/rework 规则冲突 + 释放 Shell 预算
   - 通用规则第 7 条拆分为 dev/rework 双轨制，rework 模式严禁 `test -f`/`which`。
   - 常见错误 SOP（第 176、201、203 条）增加 dev/rework 区分，rework 模式删除 `test -f` 强制要求。
   - rework Shell 上限从 10 提升至 12。
   - rework 启动流程改为先读取 `AGENT_ENV_CHECK_FILE` 预检信息，跳过 `git status`。
   - 预期收益: 每个 rework 任务减少 2-3 次 Shell 调用（从 10.67 降至约 8），消除规则冲突导致的 agent 困惑。

3. **`workflows/rework.yaml`** — git 状态预检下沉
   - 在 `envReadinessCheck` 的 JSON 输出中增加 `git_branch` 和 `git_status_lines`。
   - 预期收益: rework agent 无需自行执行 `git status`，节省 1 次 Shell/任务。

## 趋势追踪

- 与上次对比:
  - **avgStepsDelta: -1.64** ✅ — 平均步数继续下降，说明 prompt 精简措施生效。
  - **errorRateDelta: -0.11** ✅ — 错误率继续下降，上次修改（删除 review 冗余 gh 命令）的收益在持续释放。
  - **successRateDelta: -0.14** ❌ — review 成功率从 100% 降至 75%，出现 1 例零产出。虽然整体趋势改善，但零产出任务是个别异常，需观察下轮是否复现。
- 建议下次重点关注:
  - **review 的零产出根因**：本次 1 例零产出是否与 agent 未遵守 WriteFile 强制检查点有关？需要更细粒度的错误类型分布（file_not_found / gh 失败 / chunk exceed）。
  - **rework Shell 使用率**：本次修改后预计平均 Shell 从 10.67 降至 8-9 次/任务，需验证是否还有突破 12 次上限的任务。
  - **reviewer.md 的 diff 获取策略**：合并 `gh pr diff` 后，chunk exceed 的风险是否增加？`head -200` 是否足够覆盖小型 PR 的完整 diff？
