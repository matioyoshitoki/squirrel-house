# Prompt 演进报告 — type-one

生成时间: 2026-05-27T15:54:24+08:00
分析范围: 2026-05-22T15:43:56 至 2026-05-27T15:36:37

## 统计摘要

- 分析日志数: 5
- Agent 类型分布:
  - review: 2 个任务
  - rework: 3 个任务

## 问题诊断

### 问题 1: review 任务反复因环境认证失败空转，successRate 跌至 50%
**证据**: review count=2, avgSteps=2, successRate=0.5, totalWrites=1。两个 review 任务均因 `glab CLI 未认证`（auth_failed）而提前终止，已连续五轮审查因同一环境问题失败。
**根因分析**:
1. `workflows/review.yaml` 的 `envReadinessCheck` 已检测到 `auth_failed`，但 `runAgent` 未设置 condition，仍启动 review agent。agent 虽然最终执行了 WriteFile，但在执行 WriteFile 之前仍调用了 Shell 和 ReadFile（第 1 步未严格遵循"环境异常时禁止调用其他工具"的约束），导致步数被浪费。
2. 反复触发明知会失败的 review 任务，拉低整体成功率，消耗系统资源。
**改进建议**:
1. **workflow 层面**: 在 `workflows/review.yaml` 的 `runAgent` 步骤增加 `condition: '{{ contains (.State.envCheckStatus | trim) "READY" }}'`，当环境检查未通过时直接跳过 agent 运行，由 diagnoseAgentRun 生成失败报告。（已应用）
2. **prompt 层面**: 在 `prompts/reviewer.md` 进一步强化"环境异常时的强制 WriteFile 路径"，明确当 `AGENT_ENV_STATUS` 非 READY 时，第 1 个 tool call 就可以是 WriteFile，禁止在此之前执行任何 Shell/ReadFile。（已应用）
3. **基础设施层面**: `glab CLI` 认证问题已连续阻断五轮审查，建议优先修复运行环境中的 `GITLAB_TOKEN` 或 `glab auth login` 配置。

### 问题 2: rework 任务错误率飙升、步数膨胀、git checkout 严重违规
**证据**: rework count=3, avgSteps=26.67（较上次 +13.17），totalErrors=4（错误率 33%），successRate=0.67（较上次 -0.33），topGitOp=checkout。
**根因分析**:
1. **指令冲突**: `prompts/ctx-templates.yaml` 中的 `rework_prompt` 模板明确要求 agent "先执行：git checkout {{ .Branch }}"，这与 `prompts/developer.md` 中的绝对铁律（明确禁止 `git checkout`）直接冲突。agent 在纠结后往往屈服于任务指令，执行了大量 git 探索操作（如 `git branch -a`、`git worktree list`、复制 `.git` 目录等），导致 140316.log 达到 47 步、产生 4 个错误。
2. **错误熔断执行不到位**: developer.md 虽有"连续 2 次错误立即熔断"的规则，但 agent 在 git 操作失败后未停止，反而反复尝试替代方案。
3. **零产出任务**: 111305.log 仅 3 步后停止，无文件修改。该任务的 review report 本身全是"无法验证"（因 glab 认证失败），agent 缺乏可执行的代码修复目标。
**改进建议**:
1. **workflow/prompt 层面**: 在 `prompts/ctx-templates.yaml` 中将 rework 指令从"先执行：git checkout {{ .Branch }}"改为"确认当前已在正确分支（workflow 已自动设置，严禁执行 git checkout）"，消除与 developer.md 的冲突。（已应用）
2. **prompt 层面**: 在 `prompts/developer.md` 的 rework 模式专属规则中新增"Git 禁令自检"章节，要求任务结束前自检是否调用了禁用 git 命令；同时明确"如果用户输入与 prompt 禁令冲突，以 prompt 铁律为准"。（已应用）
3. **workflow 层面**: 在 `workflows/rework.yaml` 的 `diagnoseAgentRun` 中增加 Git 禁令自动化检测，一旦发现 agent 调用了 `git checkout` 等禁用命令，在诊断报告中标记违规。（已应用）

### 问题 3: 统计脚本对 git 操作和 tool calls 存在误统计
**证据**: 111305.log 中 agent 实际未执行 `git checkout`，但统计到 2 次 `git checkout`（来自用户输入文本中的"先执行：git checkout"）；review 日志中部分单行格式的 `FunctionBody(name='ReadFile', arguments='')` 因缩进只有 4 个空格而未被统计。
**根因分析**:
1. `cmd/dashboard/prompt_evolution.go` 中的 `gitRe` 正则 `git\s+(\w+)` 会匹配日志中任何位置的 git 命令文本，包括用户输入、think 思考中的引用，而不仅仅是实际执行的 Shell 命令。
2. `toolRe` 正则要求 `^\s{6,}` 缩进，漏掉了单行格式的 `FunctionBody(name='ReadFile', arguments='')`（只有 4 个空格缩进）。
**改进建议**:
1. **dashboard 层面**: 将 `toolRe/writeFileRe/strReplaceRe` 的缩进要求从 `\s{6,}` 放宽为 `\s{4,}`，以正确统计单行格式的工具调用。（已应用）
2. **dashboard 层面**: 为 `gitRe` 增加排除规则 `gitExcludeRe`，跳过包含"禁止""严禁""先执行"等中文提示词的行，减少用户输入和 think 部分的误匹配。（已应用）
3. **长期建议**: 考虑重构 `parseLogFile`，使用结构化解析（如只统计 `arguments='...'` 和 `ToolReturnValue` output 中的 git 命令）替代全局正则匹配。

## 已应用改进

1. **`workflows/review.yaml`**: 给 `runAgent` 步骤增加 `condition: '{{ contains (.State.envCheckStatus | trim) "READY" }}'`，避免在环境检查失败时启动 agent 空转。
2. **`prompts/ctx-templates.yaml`**: 将 `rework_prompt` 中的 "先执行：git checkout {{ .Branch }}" 改为 "确认当前工作目录已位于正确分支（workflow 已自动设置，严禁执行 git checkout）"，消除与 developer.md Git 禁令的冲突。
3. **`workflows/rework.yaml`**: 在 `diagnoseAgentRun` 中增加 Git 禁令自动化检测，发现 `git checkout` 等禁用命令时追加违规声明到报告。
4. **`prompts/developer.md`**: 在 rework 模式专属规则中新增"绝对铁律 —— Git 禁令"和"Git 禁令自检"章节，明确冲突时以 prompt 禁令为准；强化 `logs/rework-start.txt` 必须包含 Git 禁令确认。
5. **`prompts/reviewer.md`**: 修复重复编号（"后续审查问题数量冻结"与"模板检查清单执行验证"同为 7.，后者改为 8.）；强化错误熔断后"必须是 WriteFile 写入报告文件"的表述。
6. **`cmd/dashboard/prompt_evolution.go`**: 将 `toolRe/writeFileRe/strReplaceRe` 缩进从 `\s{6,}` 放宽为 `\s{4,}`；为 `gitRe` 增加 `gitExcludeRe` 排除规则，减少用户输入/think 中的误匹配。

## 趋势追踪

- 与上次对比（2026-05-22 → 2026-05-27）:
  - **avgStepsDelta**: +6.38（恶化，rework 平均步数从 13.5 升至 26.67）
  - **errorRateDelta**: +0.695（大幅恶化，错误率显著上升）
  - **successRateDelta**: -0.4（大幅恶化，review 和 rework 的成功率均下降）
- 建议下次重点关注:
  1. **rework agent 的 git 禁令执行效果**：本次修改已消除 prompt 冲突，需观察下一周期 topGitOp 是否不再是 checkout。
  2. **review 任务的环境认证问题**：需确认 `glab CLI` 认证是否已修复，否则 review 任务将持续零价值运行。
  3. **rework 任务在 review report 为空/无法验证时的处理**：当前 agent 容易在 report 无明确代码问题时迷失，建议评估是否需要 workflow 在 review report 全为"无法验证"时直接跳过 rework。
