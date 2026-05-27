# Prompt 演进报告 — type-one

生成时间: 2026-05-27 17:42:35 +0800
分析范围: 2026-05-27T16:37:24 至 2026-05-27T17:42:35

## 统计摘要

- 分析日志数: 19
- Agent 类型分布:
  - review: 10 次
  - rework: 9 次

| Agent 类型 | 任务数 | 平均步数 | 成功率 | 总错误数 | Top Git Op |
|-----------|--------|---------|--------|---------|-----------|
| review    | 10     | 0.0     | 0%     | 0       | —         |
| rework    | 9      | 24.2    | 100%   | 2       | checkout  |

趋势对比（与上次统计）:
- 平均步数变化: -6.3（改善）
- 错误率变化: -0.13（改善）
- 成功率变化: -0.14（恶化，主要受 review 零产出拖累）

## 问题诊断

### 问题 1: Review Agent 完全零产出
**证据**: review 任务 count=10, avgSteps=0, successRate=0%, totalWrites=0。日志文件存在（平均 194 字节），但内容仅包含 Dashboard 写入的启动/结束信息，无任何 `StepBegin` 或工具调用记录。

**根因分析**:
1. `review.yaml` 的 `runAgent` 步骤带有 condition `contains (.State.envCheckStatus | trim) "READY"`。当 `envReadinessCheck` 步骤因 worktree 目录无效、`gh` CLI 异常或 PR 不可访问等原因输出非 READY 时，agent 被静默跳过。
2. Workflow 的 `diagnoseAgentRun` 虽然能检测到零产出并生成默认失败报告，但 `ignoreError: true` 导致整个 workflow 仍返回 success，任务被错误标记为 "✅ Review 完成"。
3. Reviewer prompt 中复杂的强制写入机制（第 3/6/8/12 步 WriteFile 检查点）因 agent 根本未被执行而完全失效。

**改进建议**:
1. **移除 `runAgent` 的 condition**（已应用）：Reviewer prompt 已内置环境异常处理逻辑（第 1 步检查 `AGENT_ENV_STATUS`，非 READY 时直接 WriteFile 失败报告）。让 agent 始终运行，避免静默跳过。修改文件: `workflows/review.yaml`
2. **零产出时 workflow 失败**（已应用）：将 `diagnoseAgentRun` 的 `ignoreError` 改为 `false`，并在检测到 `STEP_COUNT=0 || WRITE_COUNT=0` 时 `exit 1`，使零产出任务被正确标记为 failed。修改文件: `workflows/review.yaml`
3. **更新统计引用并加强警示**（已应用）：Reviewer.md 中原引用 "successRate 仅 33%" 已严重过时，更新为 "successRate 为 0%，avgSteps 为 0——agent 完全未执行或零产出"，并增加 "任务结束时未执行 WriteFile 将被 workflow 判定为失败"。修改文件: `prompts/reviewer.md`

### 问题 2: Rework Agent Shell 预算突破与 Git 禁令违反
**证据**: rework Shell=95 次 / 9 任务 = 10.6 次/任务，超过 10 次硬性上限。`topGitOp=checkout`，而 `developer.md` 明确禁止 rework 模式执行 `git checkout`。

**根因分析**:
1. `developer.md` 中 rework 的 Shell 上限为 10 次，但历史统计引用（36.8 步、17.3 次 Shell）已过时，agent 可能因缺乏最新数据驱动警示而未严格执行。
2. Git 禁令自检虽然存在，但对 `git checkout` 的违规仅要求在报告中声明，违规成本过低，agent 在任务上下文冲突时选择服从外部指令而非 prompt 铁律。

**改进建议**:
1. **更新统计引用并维持高压线**（已应用）：将 `developer.md` 第 21 行统计更新为最新数据 "rework 任务平均 24.2 步、0.22 次错误、10.6 次 Shell/任务"，并强调 "Shell 预算仍被突破，Git 禁令仍被违反"。修改文件: `prompts/developer.md`
2. **升级 Git 禁令违规代价**（已应用）：将 `git checkout` 的违规后果从"报告中声明"升级为"立即 WriteFile 写入 `logs/rework-halted.md` 并结束任务"，提高拒止执行的确定性。修改文件: `prompts/developer.md`
3. **workflow 侧增加 Shell 审计**：在 `rework.yaml` 的 `diagnoseAgentRun` 中增加 Shell 调用次数统计，若超过 10 次在诊断报告中记录警告，供下次演进分析使用。修改文件: `workflows/rework.yaml`（待下次迭代评估）

### 问题 3: Review Workflow 环境检查与 worktree 可靠性
**证据**: 10 次 review 任务全部因 `envCheckStatus` 非 READY 导致 agent 被跳过，但 Dashboard 层面的 `platform.AuthStatus()` 已通过。

**根因分析**:
- `startReviewTask` 中 `git worktree add` 的错误被忽略（`wtCmd.Run()` 无错误处理）。如果 worktree 创建失败（分支已被占用、origin/branch 不存在等），`tmpDir` 可能不是一个有效的 git 仓库，导致 workflow 中的 `gh pr view` 失败。
- `envReadinessCheck` 在 `tmpDir` 中执行，而 `tmpDir` 的可靠性未经校验。

**改进建议**:
1. 在 `review.yaml` 的 `envReadinessCheck` 之前增加 `validateWorktree` 步骤：检查 `tmpDir` 是否为有效的 git 仓库（`test -d .git || git rev-parse --git-dir`）。如果无效，workflow 应立即失败并记录原因。
2. 在 Dashboard 侧（`handlers_pr.go`）对 `git worktree add` 的返回值进行检查，如果失败则终止任务并记录错误，而非忽略错误继续执行。

## 已应用改进

本次分析后直接修改了以下文件：

1. **`workflows/review.yaml`**
   - 移除了 `runAgent` 步骤的 `condition`，确保 reviewer agent 不再被静默跳过。
   - 修改 `diagnoseAgentRun`：当检测到零产出（`STEP_COUNT=0` 或 `WRITE_COUNT=0`）时，生成默认失败报告后执行 `exit 1`。
   - 将 `diagnoseAgentRun` 的 `ignoreError` 从 `true` 改为 `false`，确保零产出任务被 workflow 引擎标记为失败。

2. **`prompts/reviewer.md`**
   - 更新第 319 行统计引用：将 "successRate 仅 33%" 更新为 "successRate 为 0%，avgSteps 为 0——agent 完全未执行或零产出"。
   - 增加警示语："任务结束时未执行 WriteFile 将被 workflow 判定为失败。"

3. **`prompts/developer.md`**
   - 更新第 21 行统计引用：反映最新 rework 数据（24.2 步、10.6 次 Shell、checkout 违规）。
   - 强化 Git 禁令自检（第 206 行）：将 `git checkout` 的违规后果从报告声明升级为强制写 halted 报告并结束任务。

## 趋势追踪

- **与上次对比**:
  - rework 平均步数下降（约 -6.3），错误率下降（-0.13），说明前期 prompt 改进（预算限制、错误熔断）已产生效果。
  - 成功率下降（-0.14），主要受 review agent 零产出拖累。review 任务在统计期间全部失败，是本次最突出的恶化点。

- **建议下次重点关注**:
  1. **review agent 成功率**：本次修改已移除 condition 并增加零产出熔断，下次统计应关注 review 的 avgSteps 是否从 0 提升到正常范围（25-35 步），successRate 是否恢复到合理水平。
  2. **rework Shell 次数与 git checkout**：关注 developer.md 强化后的执行效果，特别是 `topGitOp` 是否不再是 `checkout`，以及 Shell/任务是否降至 10 次以下。
  3. **worktree 可靠性**：如果 review 继续出现大量零步数任务，需进一步检查 Dashboard 侧 `git worktree add` 的错误处理逻辑。
