# Prompt 演进报告 — type-one

生成时间: 2026-05-18 20:57:22
分析范围: 2026-05-18T17:19:59+08:00 至 2026-05-18T20:57:22+08:00

## 统计摘要

- 分析日志数: 5
- Agent 类型分布:
  - dev: 1 任务（avgSteps 69, successRate 100%, totalErrors 0, Shell 34 次）
  - review: 2 任务（avgSteps 0, successRate 0%, totalErrors 0, totalWrites 0）
  - rework: 2 任务（avgSteps 32, successRate 100%, totalErrors 1, ReadFile 32 次）

## 问题诊断

### 问题 1: review agent 完全零执行（successRate 0%, avgSteps 0）

**证据**: review 执行 2 次，avgSteps=0，totalToolCalls={}，totalWrites=0。与上次报告（avgSteps 5.75, successRate 25%）相比，情况**恶化**。

**根因分析**:
- 上次已强化 reviewer prompt 的零产出保护和 workflow 的 `AGENT_REVIEW_REPORT` 环境变量传递，但本次 review agent 连最基本的工具调用都没有，说明问题不在"agent 忘记写文件"，而在 **agent 根本没有启动或启动后立即崩溃**。
- `workflows/review.yaml` 的 `runAgent` 步骤使用 tmux 模式运行 kimi，但**没有任何前置校验和后置诊断**确认 agent 是否真实执行。如果 agent 文件路径、prompt 路径或环境变量配置有误，agent 会静默失败。
- `reviewer.md` 第 270 行仍有"如果失败，立即停止"表述，与后续"WriteFile"指令存在潜在的执行顺序歧义——agent 可能将"停止"理解为"结束任务"，从而跳过 WriteFile。

**改进建议**:
1. 在 `workflows/review.yaml` 的 `runAgent` 前增加 `validateAgentInputs` 步骤，校验 `agentFile` 和 `prompt` 文件存在性，提前暴露配置错误；在 `runAgent` 后增加 `diagnoseAgentRun` 步骤，检查日志是否存在、是否有工具调用痕迹，若 agent 完全未执行则生成默认失败报告，阻断零产出。（已修改 `workflows/review.yaml`）
2. 在 `prompts/reviewer.md` 中将"如果失败，立即停止"改为"如果失败，禁止继续审查分析，必须立即执行 WriteFile"，彻底消除"停止即结束"的歧义；并在"快速开始"中新增第 6 步"强制写入"，要求无论前 5 步成败，第 6 步必须是 WriteFile。（已修改 `prompts/reviewer.md`）

### 问题 2: dev agent 单任务全面预算失控（69 步 / 34 Shell / 16 Think）

**证据**: dev 1 个任务 avgSteps=69（突破 60 步上限 15%），Shell=34（突破 25 次上限 36%），Think=16（突破 3 次上限 433%）。虽然 successRate=100%，但资源效率极差。

**根因分析**:
- `prompts/developer.md` 虽在"铁律"中设定了 60 步 / 3 Think / 25 Shell 的上限，但 **agent 完全未内化这些约束**。16 次 think 表明 agent 将 think 当作了常规思考工具；34 次 Shell 表明 agent 仍在用 Shell 完成本可用 ReadFile/Grep 完成的操作。
- 上次报告已将 dev Shell 预算设为 25 次，但本次单任务即突破到 34 次，说明 25 次的上限对当前 agent 模型的约束力不足。同理，3 次 Think 上限被突破到 16 次，完全失效。
- 40 步预警未能阻止 69 步的失控，说明预警阈值太晚，agent 在 40 步时已完全失去预算感知。

**改进建议**:
1. 将 `prompts/developer.md` 的 Think 上限从 3 次收紧至 2 次（第 3 次触发熔断），并增加"Think 滥用快速自检"：每次 think 前必须在思考中回答"是否涉及 3 个以上文件协调或复杂架构决策？如否，此为滥用，必须跳过 think 直接执行"。（已修改 `prompts/developer.md`）
2. 将 dev 模式 Shell 预算从 25 次收紧至 20 次，并新增"Shell 使用审批"机制：超过 15 次后，每次 Shell 必须在思考中说明"为什么当前操作不能用 ReadFile/Grep/StrReplaceFile 替代"。（已修改 `prompts/developer.md`）
3. 将"40 步预警"提前至"35 步预警"，因为在 69 步的失控案例中，40 步时 agent 已完全失去预算感知，需要更早的强制干预；同时更新所有预算引用数字以匹配新限制。（已修改 `prompts/developer.md`）

### 问题 3: rework agent 错误率 50%（1/2 任务出错）

**证据**: rework 2 次任务中有 1 次错误（50% 错误率），平均 32 步，Shell 调用总计 19 次（平均 9.5 次）。

**根因分析**:
- rework 与 dev 共用 `prompts/developer.md`，而 developer.md 的"rework 常见错误快速诊断清单"虽已覆盖 StrReplaceFile / Shell / 超时等场景，但 **缺少"环境预检查"指引**。agent 可能在环境未就绪时执行命令，导致早期错误。
- `workflows/rework.yaml` 的 `installDeps` 步骤使用 `|| true` 忽略错误，agent 可能误以为依赖已就绪，后续命令因环境缺失而失败。
- 50% 的错误率意味着每 2 个任务就有 1 个遇到工具调用错误，而 rework 的累计错误熔断阈值原为 3 次，对于仅有 40 步预算的 rework 模式过于宽松，未能及时止损。

**改进建议**:
1. 在 `workflows/rework.yaml` 的 `installDeps` 后增加 `envReadinessCheck` 步骤：确认关键文件（AGENTS.md、reviewReport 等）存在，并将检查结果写入 `logs/rework-env-check.txt` 供 agent 读取，减少 agent 因环境假设错误导致的工具调用失败。（已修改 `workflows/rework.yaml`）
2. 在 `prompts/developer.md` 的 rework 章节中，将"累计错误熔断"从 3 次强化为"累计 2 次错误即熔断"，因为 rework 预算仅 40 步，容错空间比 dev 更小。（已修改 `prompts/developer.md`）

## 已应用改进

本次分析后直接修改了以下文件：

1. **`prompts/developer.md`** — 全面收紧预算约束与预警阈值
   - Think 上限从 3 次收紧至 2 次（rework 保持 1 次），第 3 次触发熔断
   - 增加"Think 滥用快速自检"：每次 think 前必须回答是否为复杂架构决策
   - dev 模式 Shell 预算从 25 次收紧至 20 次，更新所有 SetTodoList 引用
   - 新增"Shell 使用审批"机制：超过 15 次后必须说明不能用非 Shell 工具替代的理由
   - "40 步预警"提前至"35 步预警"，更早强制进入收尾阶段
   - rework 累计错误熔断从 3 次收紧至 2 次
   - 更新所有统计数据引用（69 步、16 Think、34 Shell），以最新失控数据警示 agent

2. **`prompts/reviewer.md`** — 彻底消除零产出歧义
   - 将"如果失败，立即停止，在第 2 步 WriteFile"改为"如果失败，禁止继续审查分析，必须立即执行 WriteFile"
   - 在"快速开始"中新增第 6 步"强制写入"：无论前 5 步成败，第 6 步必须是 WriteFile
   - 明确警示 avgSteps=0 的零产出任务是不可接受的

3. **`workflows/review.yaml`** — 增加 agent 启动与执行诊断
   - 新增 `validateAgentInputs` 步骤（`runAgent` 前）：校验 `agentFile` 和 `prompt` 文件存在性，配置错误时提前失败
   - 新增 `diagnoseAgentRun` 步骤（`runAgent` 后）：检查日志文件是否存在、是否有工具调用痕迹；若 agent 未执行，自动生成默认失败报告到 `review-report-<pr>.md`，阻断零产出统计

4. **`workflows/rework.yaml`** — 增加环境就绪检查
   - 新增 `envReadinessCheck` 步骤（`installDeps` 后）：检查 AGENTS.md、Makefile、reviewReport 等关键文件存在性
   - 将检查结果写入 `logs/rework-env-check.txt`，供 rework agent 启动时读取，减少因环境假设错误导致的工具调用失败

## 趋势追踪

- **与上次对比**:
  - review **恶化**（successRate 25% → 0%，avgSteps 5.75 → 0），从"低产出"恶化为"完全零执行"，是最紧急的回归问题。
  - dev **成功率改善**（successRate 50% → 100%），但**效率恶化**（avgSteps 40 → 69，且单任务 Shell 34 次、Think 16 次），说明 agent 在有产出的同时完全无视了预算约束。
  - rework **新增问题**（错误率 50%），此前无 rework 统计数据，本次首次暴露环境错误和容错阈值问题。
  - 整体趋势数据（avgStepsDelta -1.18, errorRateDelta -2.36）主要受历史数据稀释，本次窗口内的实际表现并不乐观。

- **建议下次重点关注**:
  - review agent 的启动问题是否解决：successRate 是否从 0% 回升至 50% 以上，avgSteps 是否大于 0
  - dev agent 的预算遵守情况：步数是否回归 60 以内，Shell 是否控制在 20 次以内，Think 是否在 2 次以内
  - rework agent 的错误率是否从 50% 降至 20% 以下
