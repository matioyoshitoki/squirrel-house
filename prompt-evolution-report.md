# Prompt 演进报告 — type-one

生成时间: 2026-05-20T11:01:37+08:00
分析范围: 2026-05-20T09:51:59 到 2026-05-20T10:57:55

## 统计摘要

- 分析日志数: 10
- Agent 类型分布: review 5, rework 5

## 问题诊断

### 问题 1: rework 高工具调用错误率

**证据**: rework 任务共 5 个，其中 2 个任务出现错误（40% 任务错误率），累计 4 个工具调用错误。平均每个任务 0.8 个错误，已接近 developer.md 规定的 rework 错误上限（2 次）。对比 review agent 的 5 任务 3 错误，rework 的单位任务错误密度显著更高。

**根因分析**: developer.md 中虽已包含 rework 专属错误 SOP（第 0 步错误预判、常见错误处理），但存在三个缺口：
1. **SetTodoList 计数器缺失错误维度**——agent 对错误预算缺乏显性跟踪，容易在不知不觉中触及熔断线；
2. **缺少系统化的"错误前预防"机制**——agent 常在未确认文件存在性时直接执行 `ReadFile`/`StrReplaceFile`，或未检查命令可用性就执行 `Shell`；
3. **常见错误 SOP 未按错误类型强制要求诊断步骤**——agent 可能在遇到路径错误时仍尝试替代路径，而非按规则直接跳过。

**改进建议**:
1. 在 `prompts/developer.md` 的 SetTodoList 中增加"错误"计数器，并在强制报数格式中同步体现；同时在"确认环境"步骤增加 `test -f "$AGENT_REVIEW_REPORT"` 的硬性预检要求。（涉及文件: `prompts/developer.md`）
2. 扩展 `developer.md` 的 rework 常见错误 SOP，按"路径类/StrReplaceFile/Shell/git"四类错误分别规定强制诊断步骤，明确"命令执行前存在性检查"为前置动作。（涉及文件: `prompts/developer.md`）
3. 在 `workflows/rework.yaml` 的 envReadinessCheck 中增加 review report 内容非空检查（< 50 字节视为异常），提前拦截空报告导致的 agent 困惑。（涉及文件: `workflows/rework.yaml`）

### 问题 2: review Shell 调用密度偏高且逼近预算上限

**证据**: review 任务平均 Shell 调用 15.6 次/任务（78/5），已触及 reviewer.md 规定的 15 次 Shell 上限。avgSteps=25.4，虽在合理范围，但 Shell 密度高意味着 agent 在用命令行做大量本可用 ReadFile/Grep 完成的工作。

**根因分析**: reviewer.md 虽有"将 diff 内容写入临时文件"的要求，但缺乏对"diff 获取后禁止再用 Shell 过滤/浏览"的明确约束。agent 可能在获取 diff 后反复使用 `grep`、`head`、`cat` 等 Shell 命令处理输出，而非将 diff 保存后用 ReadFile/Grep 工具处理。

**改进建议**:
1. 在 `prompts/reviewer.md` 的"Shell 命令输出硬限制"条款后，增加一条"diff 处理规范"：获取 `gh pr diff` 后必须立即写入临时文件，后续所有对 diff 内容的引用必须通过 `ReadFile` 或 `Grep` 完成，禁止再用 Shell 过滤 diff。（涉及文件: `prompts/reviewer.md`）

## 已应用改进

### 1. `prompts/developer.md` — 强化 rework 错误预防机制
- **SetTodoList 计数器扩展**: 将初始计数器从三项扩展为四项，增加 `"错误: 0/5（rework 0/2）"`，使 agent 对错误预算有显性感知。
- **环境确认增加预检**: 在"确认环境"步骤中，明确要求 rework 模式下额外执行 `test -f "$AGENT_REVIEW_REPORT"`，确保 review report 存在且可访问后再读取。
- **rework 启动确认增加路径验证**: 将原来的"第 1 步 ReadFile"改为"第 1 步先确认路径存在，第 2 步 ReadFile，第 3 步 WriteFile"，避免在文件不存在时直接 ReadFile 导致错误。
- **常见错误 SOP 结构化扩展**: 按错误类型（路径类、StrReplaceFile、Shell command not found、Shell 超时/非零、git commit hook）分别规定强制诊断步骤，并增加"命令执行前存在性检查"要求。

### 2. `workflows/rework.yaml` — 增加 review report 非空检查
- 在 `envReadinessCheck` 步骤中，当 review report 存在时，额外检查其字节数是否小于 50 字节。若过短，在环境检查日志中标记 `review_report_empty: true`，供 agent 和人类快速识别空报告场景。

## 趋势追踪

- 与上次对比: avgSteps 下降 18.25 步，errorRate 下降 0.8，整体呈改善趋势。successRate 维持 100%，零产出问题已得到控制。
- 建议下次重点关注: **rework agent 的错误类型分布**。当前统计仅知 rework 有 4 个错误，但缺少具体错误分类（路径类、StrReplaceFile、Shell、网络等）。建议在日志采集侧增加错误类型标签，以便 prompt-evolution agent 做更精准的根因定位。
