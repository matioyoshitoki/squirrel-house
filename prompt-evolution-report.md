# Prompt 演进报告 — type-one

生成时间: 2026-05-20 09:53:35
分析范围: 2026-05-20T08:20:58 到 2026-05-20T09:51:59

## 统计摘要

- 分析日志数: 4
- Agent 类型分布: dev: 1, review: 2, rework: 1

## 问题诊断

### 问题 1: dev 任务步数严重超标（134 步 vs 60 步限制）
**证据**: 
- dev 任务平均 **134 步**，超出 prompt 规定的 60 步上限 **123%**
- Shell 调用 **64 次**，超出 20 次上限 **220%**
- 累计错误 **6 次**，超出 5 次错误熔断线
- topGitOp 为 `commit`，但 Shell 调用总量异常高

**根因分析**:
1. **缺乏外部硬约束**: `workflows/dev.yaml` 中 `--max-steps-per-turn 10000` 给了 agent 极大的自由度，当 agent 未严格遵守 prompt 中的内部预算时，没有任何外部机制兜底。
2. **预算自我监控失效**: prompt 中虽然规定了 60 步/20 Shell/5 错误的上限，但 agent 在每次工具调用前没有强制报数的习惯，导致逐步失控。
3. **Shell 使用边界模糊**: prompt 未明确列出 Shell 的常见误用场景，agent 可能用 `ls`、`find`、`cat`、`git status` 等本可用其他工具替代的命令消耗了大量 Shell 预算。

**改进建议**:
1. **workflow 硬兜底**: 将 `workflows/dev.yaml` 的 `--max-steps-per-turn` 从 `10000` 降至 `75`，作为 prompt 60 步上限的外部保险。（已应用）
2. **强制报数机制**: 在 `prompts/developer.md` 的任务启动检查清单中新增规则：每次调用工具前必须在 reasoning 中显式报数 `[预算] 步数 X/60, Think Y/2, Shell Z/20, 错误 W/5`，不报数不得执行。（已应用）
3. **Shell 白名单**: 在 `prompts/developer.md` 的效率约束中新增第 9 条，明确禁止使用 Shell 的场景（`ls`/`find` 浏览目录、`cat` 读文件、反复 `git status`、未修复就重试同一命令），将 Shell 严格限制在构建/测试/安装依赖/git 操作四类。（已应用）

### 问题 2: dev 任务 StrReplaceFile 效率低下
**证据**:
- `StrReplaceFile` 调用 **33 次**，`ReadFile` 调用 **51 次**
- 6 个错误中有相当一部分可能来自未经确认的 StrReplaceFile 替换失败

**根因分析**:
- 虽然 prompt 要求"执行 StrReplaceFile 前必须先 ReadFile"，但对于小文件，多次 ReadFile → StrReplaceFile 的往返效率极低，且容易因微小差异导致替换失败。
- agent 缺乏"小文件直接重写"的策略指引，倾向于逐段替换，增加了步骤和错误概率。

**改进建议**:
1. **小文件优先重写**: 在 `prompts/developer.md` 的 StrReplaceFile 规则中增加：对于小于 100 行的文件，优先使用 `WriteFile` 重写整个文件，减少往返。（已应用）
2. **下次迭代评估**: 如果下次统计仍显示 ReadFile/StrReplaceFile 比例低于 1:1（每次替换前都读取），考虑在 prompt 中增加"批量替换"指引，要求一次 ReadFile 后尽量在同一回合完成该文件的所有修改。

## 已应用改进

### 1. `workflows/dev.yaml`
- 修改: `--max-steps-per-turn 10000` → `--max-steps-per-turn 75`
- 原因: 作为 dev agent 60 步预算的外部硬兜底，防止 agent 自我监控失效时无限消耗步骤。

### 2. `prompts/developer.md`
- **新增"强制报数机制"**（任务启动检查清单第 4 条）: 要求每次调用工具前在 reasoning 中报数，格式固定，不报数不得执行。
- **新增"Shell 白名单与常见误用"**（效率约束第 9 条）: 明确禁止使用 Shell 的四种典型误用，将 Shell 严格限定在四类合法用途。
- **优化"StrReplaceFile 前置检查与失败处理"**: 增加"小于 100 行的文件优先用 WriteFile 重写"指引，减少低效往返。

## 趋势追踪

- 与上次对比:
  - `avgStepsDelta`: **+19.5**（平均步数增加，趋势恶化）
  - `errorRateDelta`: **-0.36**（错误率下降，趋势改善）
  - `successRateDelta`: **+0.14**（成功率上升，趋势改善）
- 综合判断: 成功率在提升，但 dev 任务的步数失控是本次最突出的问题。本次改进的核心目标是将 dev 任务的步数从 134 拉回到 60 以内。

- 建议下次重点关注:
  1. **dev agent 的预算遵守率**: 观察强制报数机制和 workflow 硬兜底是否有效将步数控制在 75 以下。
  2. **Shell 调用次数**: 观察 Shell 白名单是否能将 Shell 调用从 64 次降至 20 次以内。
  3. **StrReplaceFile 成功率**: 观察小文件重写指引是否能减少 StrReplaceFile 相关的错误和往返。
