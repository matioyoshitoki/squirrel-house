# Prompt 演进报告 — type-one

生成时间: 2026-05-18 17:26:01
分析范围: 0001-01-01T00:00:00Z 至 2026-05-18T17:19:59+08:00

## 统计摘要

- 分析日志数: 9
- Agent 类型分布:
  - design: 3 任务（avgSteps 49, successRate 100%, totalErrors 15）
  - dev: 2 任务（avgSteps 40, successRate 50%, totalErrors 7）
  - review: 4 任务（avgSteps 5.75, successRate 25%, totalErrors 1）

## 问题诊断

### 问题 1: review 任务零产出率极高（75%）
**证据**: review 执行 4 次，successRate 仅 0.25，totalWrites 仅 1，avgSteps 仅 5.75。
**根因分析**:
- reviewer prompt 要求 gh 认证失败后"用 WriteFile 输出极简失败报告"，但统计数据表明 3/4 的任务没有产生任何 WriteFile 调用。
- `workflows/review.yaml` 没有向 agent 传递 `AGENT_REVIEW_REPORT` 环境变量，agent 可能因路径不确定而未写入文件。
- avgSteps 5.75 说明大量任务在第 1-2 步 gh 认证失败后就以纯文本输出结束，未执行 WriteFile，被统计系统标记为零产出。
**改进建议**:
1. 在 reviewer prompt 中明确：gh 失败后**必须** WriteFile 到 `logs/review-report-<pr>-fail.md`，禁止以纯文本输出代替文件写入。（已修改 `prompts/reviewer.md`）
2. 在 review workflow 中显式传递 `AGENT_REVIEW_REPORT` 环境变量，消除路径不确定性。（已修改 `workflows/review.yaml`）
3. 增加"零产出熔断"约束：任务结束前若未产生 WriteFile，必须强制写入空报告说明原因。（已修改 `prompts/reviewer.md`）

### 问题 2: design 任务高错误率 + 高步数
**证据**: design 3 次任务，totalErrors 15（平均 5 次错误/任务），avgSteps 49，2 次超过 50 步（最高 65 步），Shell 调用 99 次（33 次/任务），ReadFile 48 次（16 次/任务）。
**根因分析**:
- `prompts/ui-designer.md` 相比 `prompts/developer.md` **完全缺少**效率约束章节（无步数上限、无 Think 限制、无错误熔断、无 Shell 预算）。
- agent 在探索阶段无节制地调用 Shell（33 次/任务）和 ReadFile（16 次/任务），遇到错误后缺乏恢复指引，反复试错导致 totalErrors 高达 15。
- 没有"快速开始"检查清单，agent 在环境确认和地图读取上消耗过多步数，缺乏明确的执行路径。
**改进建议**:
1. 在 ui-designer prompt 中增加与 developer 对标的效率约束：步数上限 60、Think 上限 3 次、错误上限 5 次、Shell 预算 20 次、git status 上限 3 次。（已修改 `prompts/ui-designer.md`）
2. 增加"40 步预警"和"设计探索预算 10 步"机制，强制 agent 在探索超支时切换到执行模式。（已修改 `prompts/ui-designer.md`）
3. 增加快速开始检查清单（第 1-3 步完成环境确认、地图读取、待办清单），减少前期犹豫。（已修改 `prompts/ui-designer.md`）

### 问题 3: dev 任务成功率 50% + Shell 调用失控
**证据**: dev 2 次任务，successRate 0.5，avgSteps 40，最高 77 步（突破 60 步上限），Shell 调用 47 次（23.5 次/任务），topGitOp 为 `status`。
**根因分析**:
- `prompts/developer.md` 虽有 60 步上限，但 **dev 模式缺少独立的 Shell 预算上限**（仅 rework 模式限 10 次）。agent 将大量 Shell 调用用于本可用 ReadFile/Grep/StrReplaceFile 完成的操作。
- topGitOp 为 `status` 直接印证了 prompt 中提到的症状："agent 在用 `git status` 代替有效的进度判断——这是步数失控的典型症状"。
- 最高 77 步突破 60 步上限，说明 prompt 中的熔断机制执行不力，agent 对预算缺乏感知。
**改进建议**:
1. 在 developer prompt 中增加 dev 模式 Shell 预算上限 25 次，与 rework 模式 10 次并列，并要求每次 Shell 前更新 SetTodoList 计数器。（已修改 `prompts/developer.md`）
2. 强化"超过 60 步后调用工具属于严重违规"的措辞，增加步数超限后的具体处理 SOP。
3. （建议纳入下次迭代评估）`workflows/dev.yaml` 中 `preCommitChecks` 的 `ignoreError: false` 可能在 agent 完成代码修改后因 lint/test 失败而阻断提交，导致 successRate 降低。需评估是否应改为 `ignoreError: true` 仅输出警告。

## 已应用改进

本次分析后直接修改了以下文件：

1. **`prompts/ui-designer.md`** — 新增完整"效率约束（不可违背）"章节
   - 步数硬上限 60 步，40 步预警进入收尾阶段
   - Think 上限 3 次，第 4 次触发熔断
   - 错误上限 5 次，连续 2 次错误熔断
   - Shell 预算上限 20 次
   - git status 上限 3 次
   - 设计探索预算 10 步
   - 进度检查点（每 10 步自评）
   - 快速开始检查清单（第 1-3 步）

2. **`prompts/reviewer.md`** — 强化零产出保护
   - 环境预检查（第 1 步）失败后强制要求 WriteFile 到指定路径，并引用统计数据警示
   - 审查结论写入文件条款增加 `AGENT_REVIEW_REPORT` 未设置时的兜底路径
   - 新增"纯文本输出绝对不算完成"和"零产出熔断"条款

3. **`workflows/review.yaml`** — 消除环境变量盲区
   - `runAgent` 步骤 env 中新增 `AGENT_REVIEW_REPORT: '{{ .Vars.tmpDir }}/logs/review-report-{{ .Vars.prNumber }}.md'`

4. **`prompts/developer.md`** — 补齐 dev 模式 Shell 预算
   - 铁律中 Shell 预算上限扩展为"dev 限额 25 次 / rework 限额 10 次"
   - 增加统计数据引用（dev 平均 23.5 次 Shell）作为警示
   - 明确倡导"能用 ReadFile/Grep/StrReplaceFile 完成的工作绝不使用 Shell"

## 趋势追踪

- **与上次对比**: 无历史报告可供对比，本次为基线分析。
- **建议下次重点关注**:
  - review agent 的 successRate 是否从 25% 提升至 75% 以上
  - design agent 的 avgSteps 是否降至 40 以下，totalErrors 是否降至 2 以下/任务
  - dev agent 的 Shell 调用次数是否降至 25 次以下，maxSteps 是否不再突破 60
