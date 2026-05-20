# Prompt 演进报告 — type-one

生成时间: 2026-05-20T08:25:00+08:00
分析范围: 2026-05-19T16:40:39 至 2026-05-20T08:20:58

## 统计摘要

- 分析日志数: 7
- Agent 类型分布:
  - design: 3 个任务
  - review: 2 个任务
  - rework: 2 个任务
- 整体趋势（相较上次）:
  - avgStepsDelta: **+9.38**（步数膨胀）
  - errorRateDelta: **+1.69**（错误率恶化）
  - successRateDelta: **+0.36**（成功率改善）

## 问题诊断

### 问题 1: review agent 错误率 100%（2/2 任务均遇错误）

**证据**:
- review: 2 任务, successRate = **100%**, 但 errorRate = **100%**（2/2 任务均出现错误）
- totalErrors = **6**, 平均每个任务 **3 次错误**
- avgSteps = **19.5**, Shell = **15 次/任务**, ReadFile = **12.5 次/任务**
- 对比上次（3 任务, successRate = 0%, avgSteps = 0），虽然 successRate 恢复到了 100%，但错误率极端恶化

**根因分析**:
1. **prompt 预算上限与实际严重脱节**：`prompts/reviewer.md` 中规定 Shell ≤ 8 次、ReadFile ≤ 12 次，但实际统计分别为 15 次和 12.5 次/任务。agent 在实际运行中必然突破上限，导致"预算约束完全失效"的心理暗示，反而不再关注任何约束。
2. **workflow 与 prompt 重复执行已知的失败命令**：`workflows/review.yaml` 的 `envReadinessCheck` 已执行 `gh auth status` 检查，但结果未传递给 agent。agent prompt 第 1 步再次执行 `gh auth status`，如果认证失败则产生重复错误。统计中的 6 个错误很可能大量来自这种"已知失败命令的重复执行"。
3. **过时的统计描述误导行为**：prompt 中声称"当前统计平均 6.7 次 Shell"，而实际已达 15 次。agent 基于错误的心理预期运行，可能在"接近 8 次上限"时提前 panic，或因上限明显不可达而彻底忽略约束。

**改进建议**:
1. **修改 `workflows/review.yaml`**：将 `envReadinessCheck` 的结果通过环境变量 `AGENT_ENV_STATUS` 传递给 agent，避免 agent 重复执行已确认失败的 `gh` 命令（已应用）。
2. **修改 `prompts/reviewer.md`**：
   - 更新第 1 步环境预检查，优先读取 `AGENT_ENV_STATUS`。若包含 `auth_failed`/`pr_not_found`，直接 `WriteFile` 失败报告，不再执行任何 `gh` 命令（已应用）。
   - 将 Shell 预算上限从 8 次调整至 **15 次**（匹配实际需求），ReadFile 上限从 12 次调整至 **15 次**，并同步更新统计描述（已应用）。
3. **基础设施层面**：如果 `gh` 认证问题持续存在，建议在 dashboard 启动或 workflow 执行前预配置 GitHub CLI 认证，从根本上消除 review agent 的高频错误源。

### 问题 2: design agent 错误率 67% 且 successRate 仅 67%

**证据**:
- design: 3 任务, successRate = **67%**, errorRate = **67%**（2/3 任务出现错误）
- totalErrors = **6**, 平均每个任务 **2 次错误**
- avgSteps = **20**, topGitOp = `status`
- 1 个任务零产出（successRate < 100% 意味着存在零产出）

**根因分析**:
1. **prompt 中包含严重过时的统计数据**：`prompts/ui-designer.md` 中声称"design 任务平均已达 49 步且频繁超支"、"平均 8 次 think 的严重失控"、"错误率高达 100%（平均 5 次错误/任务）"。这些描述与当前实际数据（20 步、think 次数未统计、2 次错误/任务）严重不符，会扭曲 agent 的风险感知和行为策略。
2. **缺少前置检查导致可预见的错误**：design prompt 中没有类似 reviewer/developer 的"文档存在性检查"和"第 0 步错误预判"。agent 可能在未确认文件存在的情况下直接 `ReadFile`，产生本可避免的错误。
3. **workflow 缺少 agent 产出诊断**：`workflows/design.yaml` 没有类似 review/rework 的 `diagnoseAgentRun` 步骤。当 agent 零产出时，workflow 无法区分"agent 未启动"、"agent 提前终止"还是"确实无需修改"，导致统计失真且不利于故障排查。

**改进建议**:
1. **修改 `prompts/ui-designer.md`**：删除/修正过时的统计描述（49 步、8 次 think、5 次错误），替换为基于最新数据的客观描述；增加"第 0 步 — 错误预判"和"文档存在性检查"要求，与 review/developer 对齐（已应用）。
2. **修改 `workflows/design.yaml`**：在 `runAgent` 前增加 `validateAgentInputs` 校验，在 `runAgent` 后增加 `diagnoseAgentRun` 诊断步骤，检测 agent 是否有 `WriteFile` 产出及 designs 目录是否生成资产（已应用）。

### 问题 3: 多个 prompt 中存在严重过时的统计引用

**证据**:
- `prompts/ui-designer.md`：引用"49 步、8 次 think、5 次错误/任务"
- `prompts/developer.md`：引用"rework 任务已达 50 步、7 次错误、28 次 Shell/任务"
- `prompts/reviewer.md`：引用"Shell 平均 6.7 次、ReadFile 平均 16.5 次"

**根因分析**:
Prompt 中嵌入统计数据的初衷是让 agent "感知"当前系统状态并调整行为。但当这些数据严重偏离实际时，会产生两种负面效应：
- **过度悲观**：agent 认为自己处于"严重超支"状态，可能过早收缩范围或 panic，导致产出质量下降。
- **约束失效**：agent 发现实际使用量远超上限后，会将所有预算约束视为"不可执行的建议"，从而彻底无视它们。

**改进建议**:
1. **建立 prompt 统计数据的定期刷新机制**：每次 prompt-evolution agent 运行后，同步更新所有 prompt 中的引用统计。本次已直接修正 `ui-designer.md`、`developer.md`、`reviewer.md` 中的三处过时数据（已应用）。
2. **保守使用统计引用**：将具体的数字描述改为趋势描述（如"最新统计显示 rework 预算约束整体有效"），降低因数据过期导致的失真风险。

## 已应用改进

本次分析后已直接修改以下文件：

### 1. `workflows/review.yaml`
- **增加 `AGENT_ENV_STATUS` 环境变量传递**：在 `runAgent` 步骤的 env 中新增 `AGENT_ENV_STATUS: '{{ .State.envCheckStatus | trim }}'`，将 workflow 层面 `envReadinessCheck` 的结果（`READY` / `auth_failed` / `pr_not_found`）传递给 reviewer agent，避免 agent 重复执行已确认失败的 `gh` 命令。

### 2. `prompts/reviewer.md`
- **优化第 1 步环境预检查**：改为优先读取 `AGENT_ENV_STATUS`。若其值包含 `auth_failed` 或 `pr_not_found`，直接 `WriteFile` 失败报告并结束，不再执行任何 `gh` 命令。
- **更新工具调用预算**：Shell 上限从 8 次提升至 **15 次**（匹配实际平均 15 次），ReadFile 上限从 12 次提升至 **15 次**（匹配实际平均 12.5 次）。同步更新统计描述，删除过时的"6.7 次 Shell"引用。
- **更新错误预判**：明确说明 workflow 已通过 `AGENT_ENV_STATUS` 传递预检结果，agent 不得重复执行 `gh auth status`。

### 3. `prompts/ui-designer.md`
- **修正三处过时统计描述**：将"49 步、8 次 think、5 次错误/任务"更新为基于最新数据的客观描述（20 步、错误率 67%、平均 2 次错误/任务）。
- **新增错误预判**：增加独立的"错误预判"段落，列举 design 任务的常见错误来源（`AGENTS.md` 不存在、`StrReplaceFile` 内容不匹配、Shell 输出过大、Flutter 预览编译失败）。
- **新增文档存在性检查**：要求读取任何规范文档前先用 `test -f` 确认，避免 `ReadFile` 直接报错浪费步数。

### 4. `prompts/developer.md`
- **修正 rework 过时统计**：将"50 步、7 次错误、28 次 Shell/任务"更新为"23 步、0.5 次错误、9 次 Shell/任务"，并调整语气为"预算约束整体有效，但仍需严格执行"。

### 5. `workflows/design.yaml`
- **增加 `validateAgentInputs` 步骤**：在 `runAgent` 前校验 `agentFile` 和 `prompt` 是否存在且非空，避免启动参数缺失导致的运行失败。
- **增加 `diagnoseAgentRun` 步骤**：在 `runAgent` 后诊断 agent 执行状态，统计工具调用和写操作次数，并检查 `designs/issue-<N>/` 目录是否有实际产出，提升可观测性。

## 趋势追踪

- 与上次对比:
  - **review agent**：successRate 从 0% **恢复至 100%**，验证了"移除 workflow 跳过逻辑、让 agent 自行兜底"策略的有效性。但 errorRate 恶化至 100%（2/2 任务均有错误），说明 gh 认证/环境问题是持续性的高频错误源。
  - **design agent**：successRate **下降至 67%**（上次未在统计范围内，无直接对比），出现零产出任务；errorRate 67%，提示前置检查和错误预判不足。
  - **整体错误率**：errorRateDelta **+1.69**，趋势恶化。主要驱动因素是 review 和 design 的频繁工具调用错误。
  - **整体步数**：avgStepsDelta **+9.38**，步数膨胀。review 从 0 恢复至 19.5 步是预期内的，但 design（20 步）和 rework（23 步）也处于高位。

- 建议下次重点关注:
  1. **review agent 的错误细分**：本次修改通过 `AGENT_ENV_STATUS` 消除了重复的 `gh auth status` 失败。下次需要观察 errorRate 是否从 100% 下降。如果仍然很高，需要深入分析剩余错误的具体来源（是 `gh pr diff` 输出过大？还是文档读取失败？）。
  2. **design agent 的零产出根因**：当前统计仅能识别"有 1 个任务零产出"，但无法区分是 agent 未启动、步数耗尽未 WriteFile，还是探索后认为无需修改但忘记写 noop 报告。`diagnoseAgentRun` 步骤将在下次提供诊断日志。
  3. **prompt 中统计数据的保鲜度**：建议建立机制，每次 prompt-evolution 报告生成后，自动扫描所有 prompt 中的数字引用并与最新统计对比，标记偏差超过 50% 的描述。
