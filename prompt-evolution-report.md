# Prompt 演进报告 — type-one

生成时间: 2026-05-27T18:55:00+08:00
分析范围: 2026-05-27T17:42:36 至 2026-05-27T18:47:48

## 统计摘要

- 分析日志数: 11
- Agent 类型分布:
  - `review`: 6 个任务，avgSteps=9.17，successRate=66.7%，totalErrors=0，topGitOp="commands"
  - `rework`: 5 个任务，avgSteps=26.2，successRate=100%，totalErrors=2（错误率 40%），topGitOp="checkout"

## 问题诊断

### 问题 1: Workflow git 命令检测误报导致合规 agent 被错误降级

**证据**:
- 对 review 日志（review-pr-10-20260527-*.log）进行精确提取，只扫描 `FunctionBody(name='Shell'...)` 的 arguments 后，真实 git 命令调用为 **0 次**；但 workflow 的 `grep -cE '\bgit\s+...'` 匹配到了 11–22 次 "git 命令" 文本——这些全部来自 agent reasoning 中的声明（如"本审查全程未使用任何本地 git 命令"）。
- 对 23 个 rework 日志的同样精确提取显示，仅 **1 个任务**存在真实 `git checkout` 调用，其余 22 个任务的 `topGitOp="checkout"` 均为误报。
- review.yaml 的 diagnoseAgentRun 在误匹配后会强制将审查结论降级为 `NEEDS_MAJOR_FIX`，直接影响 review 任务质量评估。

**根因分析**:
- review.yaml 和 rework.yaml 的 diagnoseAgentRun 使用全文 grep 匹配 git 命令，未区分 agent 实际 Shell 工具调用与 reasoning/prompt/result 中的文本引用。
- 这导致 agent 的合规自检声明反而触发违规检测，产生大量假阳性，掩盖了真实违规信号。

**改进建议**:
1. **review.yaml / rework.yaml**: 将宽松 grep 替换为 Python 精确提取脚本，只统计 `FunctionBody(name='Shell'...)` 的 arguments 中的真实 git 命令。（**已应用**）
2. **dashboard 统计逻辑**: 建议同步修正报告生成器中的 `topGitOp` 统计方式，采用同样的精确提取逻辑，避免后续 prompt 基于错误统计数据做出过度反应（如 developer.md 中引用的 "topGitOp 为 checkout" 警示语即来自误报数据）。

### 问题 2: rework 任务在 worktree 环境异常时失控探索，步数膨胀

**证据**:
- rework avgSteps=26.2，接近 40 步上限；Shell 平均 8.4 次/任务，接近 10 次上限。
- 逐 Shell 分析 rework-pr-10-20260527-140316.log 发现：agent 在发现 `.git` 缺失后，连续执行了 `git branch`、`git checkout`、手动复制 `.git` 目录、`git worktree list` 等 **30+ 次 Shell 调用**，严重违反 Git 禁令并耗尽预算。
- 该任务是 5 个 rework 任务中 2 个错误来源之一，直接拉高错误率至 40%。

**根因分析**:
- developer.md 虽禁止 `git checkout`，但未明确禁止 agent 在发现 worktree 损坏时"自行诊断并修复"环境。agent 将 worktree 异常视为可修复的本地问题，触发大量探索性 Shell 命令，完全偏离 rework 的核心目标（按 review report 修复代码）。
- 缺少针对 worktree infrastructure 故障的熔断规则，使 agent 在错误路径上持续消耗步数，直至触达上限或最终失败。

**改进建议**:
1. **developer.md（rework 模式）**: 新增 "Worktree 环境异常熔断" 铁律——发现 `.git` 缺失、目录损坏或不在正确分支时，**禁止**自行复制 `.git`、执行 `git checkout` 或手动修复 worktree；正确做法是立即停止修复，WriteFile 报告环境异常并结束任务。（**已应用**）
2. **developer.md（错误预判）**: 将 "Worktree 环境异常" 列为第 5 项错误预判，在任务启动时即建立熔断意识。（**已应用**）

### 问题 3: review 任务空日志率过高，successRate 仅 66.7%

**证据**:
- review avgSteps=9.17，远低于正常审查所需的 25–35 步。
- 在统计周期内，review-pr-10 系列存在大量 **194 字节空日志**（0 tool calls），约占 review 任务的 1/3。
- successRate=66.7% 意味着 2/6 的任务被 workflow 判定为零产出（agent 未产生有效工具调用）。

**根因分析**:
- 空日志表明 agent 未成功启动或立即退出，属于 **基础设施/调度层问题**，非 prompt 设计缺陷。
- review prompt 已具备完善的快速失败和 WriteFile 兜底机制（环境异常时第 1–2 步即写入失败报告），且正常任务的审查流程完整（50–108 次 tool calls）。
- diagnoseAgentRun 对空日志的兜底（生成默认失败报告并 exit 1）导致这些任务被计入失败，拉低 successRate。

**改进建议**:
1. **dashboard/workflow 层面**: 建议对空日志（0 tool calls 且日志大小 < 1KB）任务单独标记为 `INFRA_FAILURE`，不计入 agent successRate，或触发自动重试，避免将基础设施问题误判为 agent 能力问题。
2. **review prompt**: 当前已有充足的环境预检查和 WriteFile 兜底，暂不做额外修改；待基础设施稳定后再评估 avgSteps 是否回升至健康区间（25–35 步）。

## 已应用改进

1. **`workflows/review.yaml` + `workflows/rework.yaml` — 精确化 git 命令检测**
   - 将 diagnoseAgentRun 中的宽松 `grep -cE '\bgit\s+...'` 替换为调用独立 Python 脚本 `scripts/detect-git-usage.py`。
   - 该脚本精确提取 `FunctionBody(name='Shell'...)` 的 arguments 后再匹配 git 命令，消除因 agent reasoning 中出现 "git 命令" 声明文本而导致的假阳性降级。

3. **`prompts/developer.md` — 新增 Worktree 环境异常熔断规则**
   - 在 rework 模式专属规则中增加铁律：发现 worktree 异常（`.git` 缺失、目录损坏、不在正确分支）时，**禁止**自行修复，必须立即停止并 WriteFile 报告。
   - 在错误预判清单中新增第 5 项 "Worktree 环境异常（rework 模式）"，提前建立熔断意识。

## 趋势追踪

- **与上次对比（16:37–17:42 → 17:42–18:47）**:
  - `review` successRate 从 **0% → 66.7%**：显著改善，主要得益于 review.yaml 移除了 `envCheckStatus` 非 READY 时的 condition 阻断，避免 agent 被静默跳过。
  - `rework` avgSteps 从 **24.2 → 26.2**：小幅回升，仍高于健康值（建议 < 20 步），需持续观察 worktree 熔断规则的效果。
  - `rework` errorRate 从 **22.2%（2/9）→ 40%（2/5）**：样本量小导致波动，但真实违规和错误模式已明确。
  - `review` 仍存在约 1/3 的任务为空日志（基础设施问题），拉低 avgSteps 和 successRate。

- **建议下次重点关注**:
  1. **rework 的步数控制**：在 worktree 熔断规则生效后，观察 avgSteps 是否从 26.2 降至 20 以下。
  2. **真实 git 违规率**：在修复检测误报后，统计真正的 `git checkout` 违规次数，评估 developer.md Git 禁令的实际执行效果。
  3. **review 空日志根因**：协同 infrastructure 团队定位 agent 未启动或零产出的原因（tmux/kimi 调度层），将空日志任务排除在 agent 统计之外。
