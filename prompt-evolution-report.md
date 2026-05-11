# Prompt 演进报告 — social-world

生成时间: 2026-05-07 12:00:00+08:00
分析范围: 2026-05-07T10:45:50 到 2026-05-07T11:56:21

## 统计摘要

- 分析日志数: 4
- Agent 类型分布: rework=4

| 指标 | 数值 |
|------|------|
| 平均步数 | 23.25 |
| 成功率 | 50% |
| 错误率（任务级） | 75% (3/4) |
| 总错误数 | 5 |
| 总文件写入/替换 | 10 |
| Shell 调用总数 | 59（平均 14.75 次/任务） |
| Top Git Op | Shell（表明大量 Shell 调用未产生有效 git 操作） |

## 问题诊断

### 问题 1: rework Shell 预算严重超支
**证据**: 4 个 rework 任务共调用 Shell 59 次，平均 14.75 次/任务；developer.md 规定 rework 非 git Shell 限额 10 次，但 agent 实际严重失控。
**根因分析**:
1. prompt 中 Shell 预算定义为"除 git 操作外"，agent 难以准确区分 git Shell 与非 git Shell，导致计数混乱、实际失控。
2. `rework_prompt` 要求"按项目技术栈运行自测"，并列出大量测试/构建命令（flutter analyze/test、npm run lint/test/build、go test/vet 等），agent 倾向于顺序执行多个命令。
3. 缺少 Shell 预算的强制外化要求，agent 未在 `SetTodoList` 中实时跟踪 Shell 消耗。

**改进建议**:
1. **修改 developer.md 的 Shell 预算定义**：将"除 git 操作外"改为"**所有 Shell 工具调用（含 git）**总计不超过 10 次"，并要求在 `SetTodoList` 中外化 `Shell 预算: N/10`。（涉及文件: `prompts/developer.md`）
2. **修改 rework_prompt 的测试策略**：将"运行自测"改为"**仅从命令列表中选择 1 个与修改最直接相关的命令执行验证**，严禁顺序执行多个命令或运行全量测试"。（涉及文件: `prompts/ctx-templates.yaml`）
3. **在 developer.md 中增加 rework 验证限制条款**：明确 rework 模式只允许 1 个验证命令，全量测试/构建留给 dev 模式。（涉及文件: `prompts/developer.md`）

### 问题 2: rework 错误率 75%，成功率仅 50%
**证据**: 4 个任务中 3 个出现错误（累计 5 次工具错误），successRate=0.5；其中 1 个任务零产出。
**根因分析**:
1. 错误发生后，agent 缺少针对 rework 场景的"快速止损"指引。现有错误熔断规则（累计 3 次停止）虽存在，但 agent 在达到熔断前已消耗大量步数试错。
2. `StrReplaceFile` 失败是常见错误源之一（10 次写入对应若干次错误），虽然 prompt 要求先 ReadFile，但执行不到位。
3. 零产出任务说明"必须有文件修改"的约束在边界情况下执行不严格。

**改进建议**:
1. **在 developer.md 的 rework 模式段落中增加"常见错误快速诊断"子项**：将已有的 `StrReplaceFile` 失败处理、`Shell` 超时禁止重试、`git commit` hook 失败处理等规则，在 rework 预算段落内以清单形式再次强调，降低认知负荷。（涉及文件: `prompts/developer.md`）
2. **强化零产出熔断**：明确如果第 5 步前无文件修改，必须停止并报告；如果 review report 中问题已修复或无需修改，必须在第 5 步前用 `WriteFile` 生成 `logs/rework-noop-<timestamp>.md` 说明文件。（涉及文件: `prompts/developer.md`，该规则已存在但可前置到启动检查清单中强调）

### 问题 3: 成功率较上次下降 30%
**证据**: successRateDelta = -0.3。
**根因分析**: 平均步数减少 10.15、错误数也有所下降，说明熔断机制部分生效（agent 遇到问题时更快停止）。但成功产出比例反而降低，表明 agent 在紧缩预算后缺乏"有效完成"的策略——它知道何时停止，但不知道如何在预算内稳定成功。
**改进建议**:
1. **在 rework_prompt 中增加修复范围控制**：明确"优先修复所有 Blocking 问题，最多再修复 3 个 Major 问题，其余报告为留待后续迭代"。防止 agent 试图修复所有问题而耗尽预算。（涉及文件: `prompts/ctx-templates.yaml`）

## 已应用改进

本次分析产生以下直接修改：

1. **`prompts/developer.md`**:
   - 将 rework Shell 预算从"除 git 外 10 次"改为"**所有 Shell 调用（含 git）总计 10 次**"，消除计数歧义。
   - 在 `SetTodoList` 启动检查清单中，将计数器从 3 项扩展为包含 `"Shell 预算: 0/10"` 的 4 项，并要求每次 Shell 调用后同步更新。
   - 在 rework 模式预算段落中新增**"rework 验证限制"**：明确只允许执行 1 个验证命令，禁止全量测试/构建。
   - 在 rework 模式预算段落中新增**"rework 常见错误快速诊断清单"**，将已有规则以清单形式集中呈现，降低 rework 场景下的认知负荷。

2. **`prompts/ctx-templates.yaml`**:
   - 修改 `rework_prompt` 步骤 3 的表述，从"按项目技术栈运行自测"改为"**仅从以下命令中选择 1 个与修改最直接相关的命令执行验证**"，并增加"严禁顺序执行多个命令或运行全量测试"的警告。
   - 在步骤 4 的 `git commit` 提示中增加"同一 hook 错误不要重复 commit 超过 2 次"的限制。
   - 在审查报告段落前增加修复范围控制说明：优先 Blocking，最多 3 个 Major。

## 趋势追踪

- 与上次对比:
  - 平均步数: ↓ 10.15（改善，说明步数控制生效）
  - 错误数/日志: ↓ 0.95（改善，错误熔断部分起作用）
  - 成功率: ↓ 30%（恶化，agent 在预算内完成修复的能力下降）
- 建议下次重点关注: **rework 的成功率与 Shell 调用次数**。如果下次报告 Shell 调用仍 > 10 次/任务，需进一步在 workflow 层面限制 `TestCommands` 的注入范围，或设置 `--max-steps-per-turn` 更严格的硬件限制。
