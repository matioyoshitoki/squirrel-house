# Prompt 演进报告 — type-one

生成时间: 2026-05-22T13:42:20+08:00
分析范围: 2026-05-22T12:36:49 到 2026-05-22T13:40:22

## 统计摘要

- 分析日志数: 14
- Agent 类型分布: review 7, rework 7

| 指标 | review | rework |
|------|--------|--------|
| 任务数 | 7 | 7 |
| 平均步数 | 14.29 | 18.43 |
| 成功率 | 85.7% | 100% |
| 总错误数 | 1 | 5 |
| 平均 ReadFile/任务 | 16.0 | 7.0 |
| 平均 Shell/任务 | 5.0 | 6.86 |
| 写入操作总计 | 9 (WriteFile) | 38 (WriteFile 14 + StrReplaceFile 24) |
| topGitOp | commands | status |

## 问题诊断

### 问题 1: rework 高错误率

**证据**: 7 个 rework 任务共产生 5 个工具调用错误，错误涉及任务比例 71%（topIssues 中标记为 43% 的任务级错误率）。尽管 successRate 为 100%（最终均有文件修改产出），但高错误率表明 agent 在执行过程中频繁遭遇工具调用失败，消耗了额外的步数和预算。

**根因分析**:
- `developer.md` 虽已设定 rework 错误上限为 2 次，但统计数据表明该约束未被严格遵守。
- rework 任务中 `StrReplaceFile` 使用频繁（24 次），而 "old string not found" 是常见失败模式。agent 可能在未充分 `ReadFile` 确认的情况下直接执行替换，或在遇到路径错误后尝试替代路径。
- rework 平均 Shell 调用 6.86 次/任务，接近 10 次上限，表明环境探索和验证步骤仍偏多。

**改进建议**:
1. **在 `developer.md` 中增加 StrReplaceFile 强制声明**: 每次 StrReplaceFile 前必须在 reasoning 中显式声明「已 ReadFile 确认目标内容存在，准备替换」。如果 ReadFile 后发现内容已变化，必须重新评估修复策略，禁止强行替换。（已应用）
2. **强化 rework 错误熔断刚性**: `developer.md` 中已更新：rework 错误达到 1 次时，下一个 tool call 必须是 ReadFile 确认状态或 WriteFile 记录错误，禁止新的 Shell 探索或未经确认的 StrReplaceFile；达到 2 次时，立即停止所有操作，下一个且唯一的 tool call 必须是 WriteFile 写入 `logs/rework-halted.md`。（已应用）
3. **在 `workflows/rework.yaml` 中增加错误分类统计**: `diagnoseAgentRun` 步骤已增强，现在会统计总错误数、StrReplaceFile 失败次数、文件缺失次数和 Shell 失败次数，便于下次分析精准定位主要错误来源。（已应用）

### 问题 2: review ReadFile 预算不足与零产出风险

**证据**: review 任务平均 ReadFile 16.0 次/任务，超过了 `reviewer.md` 中设定的 15 次上限。successRate 为 85.7%，即 7 个任务中有 1 个零产出。totalWrites 仅 9 次，平均 1.29 次/任务。

**根因分析**:
- `reviewer.md` 中的 ReadFile 预算（15 次）已低于实际平均值（16 次）。审查任务需要读取 AGENTS.md、规范文档、PR diff、历史 review report 等多个文件，预算紧张可能导致 agent 在后期被迫跳过必要阅读，或以其他方式（如多次 Shell 调用）变相补偿。
- 零产出任务的存在说明，尽管 prompt 中已有多层 WriteFile 强制要求（第 6 步、最后一步等），在某些异常路径下（如 gh 命令失败、环境状态非 READY）agent 仍可能未正确执行 WriteFile。

**改进建议**:
1. **提升 ReadFile 预算并引入文档缓存策略**: `reviewer.md` 中 ReadFile 上限已从 15 次提升到 20 次，并增加「文档缓存策略」：对需要多次引用的规范文档，首次 ReadFile 后将关键摘要写入临时文件，后续引用该摘要而非重复 ReadFile。（已应用）
2. **增加第 8 步 WriteFile 二次检查点**: `reviewer.md` 的快速开始流程中新增第 7 步：如果第 8 个 tool call 仍未完成最终审查报告写入，下一个 tool call 必须是 WriteFile 写入进度报告，禁止继续探索新文件或执行新命令。（已应用）

### 问题 3: review 与 rework 的 Git 操作统计异常

**证据**: review 的 topGitOp 为 `commands`，rework 的 topGitOp 为 `status`。reviewer.md 明确禁止 review 任务执行任何本地 git 命令。

**根因分析**:
- review 的 `commands` 可能包含 `gh pr diff` / `gh pr view` 等 GitHub CLI 命令，这些命令在日志解析中可能被归入 git 相关操作类别。
- rework 的 `status` 直接表明 developer.md 中 "git status 全任务最多 3 次" 的约束仍被突破（7 个任务共 35 次 Shell 调用，部分为 git status）。

**改进建议**:
1. **在 `developer.md` 中强化 git status 预算执行**: 当前虽已设定 "git status 全任务最多 3 次"，但统计表明 agent 仍在用 `git status` 代替有效的进度判断。建议在 reasoning 中增加每次 git status 调用前的强制声明：「git status 计数 X/3，调用理由：...」。
2. **优化日志解析的 Git 操作分类**: 建议 `cmd/dashboard/*.go` 中将 `gh` 命令与本地 `git` 命令分开统计，避免 review 任务的合法 `gh` 命令被误判为 Git 禁令违规。

> 注：问题 3 的建议 1 涉及 reasoning 习惯，建议 2 涉及 dashboard 统计逻辑，需更多数据支撑后再实施。本次暂不修改。

## 已应用改进

本次分析后，已直接修改以下文件：

1. **`prompts/developer.md`**:
   - 在 rework 第 0 步错误预判中，增加 StrReplaceFile 强制声明要求：每次替换前必须在 reasoning 中声明「已 ReadFile 确认目标内容存在，准备替换」。
   - 在铁律部分强化错误上限执行：rework 错误达到 1 次时进入只读/收尾模式，达到 2 次时下一个且唯一的 tool call 必须是 WriteFile 写入 halted 报告。
   - 在 rework 错误计数显式自报部分，明确 X=2 时「立即停止所有操作」，严禁继续其他操作。

2. **`prompts/reviewer.md`**:
   - ReadFile 上限从 15 次提升到 20 次，并增加「文档缓存策略」指引。
   - 在快速开始流程中新增第 7 步：第 8 个 tool call 仍未 WriteFile 时必须写入进度报告。

3. **`workflows/rework.yaml`**:
   - 在 `diagnoseAgentRun` 步骤中增加错误分类统计：总错误数、StrReplaceFile 失败次数、文件缺失次数、Shell 失败次数。

## 趋势追踪

- 与上次对比:
  - 平均步数减少 4.64（改善）
  - 错误率减少 1.32（改善）
  - 成功率下降 0.07（轻微恶化，可能受 review 零产出任务影响）
- 建议下次重点关注:
  - **rework 错误分类数据**：观察 `diagnoseAgentRun` 新增的错误分类统计，确认 StrReplaceFile 失败是否为首要错误来源。
  - **review ReadFile 与产出**：观察提升 ReadFile 预算至 20 次后，review 的平均 ReadFile 次数是否下降（因缓存策略）以及 successRate 是否回升至 100%。
  - **git status 频率**：观察 developer.md 强化约束后，rework 的 topGitOp 是否从 `status` 转变为构建/测试命令。
