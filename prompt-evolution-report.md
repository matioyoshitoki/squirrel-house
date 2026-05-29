# Prompt 演进报告 — type-one

生成时间: 2026-05-29
分析范围: 0001-01-01T00:00:00Z 至 2026-05-29T10:28:17+08:00

## 统计摘要

- 分析日志数: 5
- Agent 类型分布:
  - review: 4 个任务
  - rework: 1 个任务

## 问题诊断

### 问题 1: review 任务零产出率过高（successRate 仅 50%）
**证据**: review agent `count=4`, `successRate=0.5`, `totalWrites=2`, `totalErrors=3`。4 个任务中仅有 2 个产生了 WriteFile 产出，其余 2 个完全零产出。
**根因分析**: review prompt 中虽设置了多处 WriteFile 强制检查点（第 3/6/8/12 步），但存在以下设计缺陷导致 agent 遗漏：
1. 第 319 行包含过时的 `"successRate 为 0%"` 历史注释，与当前 50% 的数据脱节，可能削弱 agent 对零产出风险的警觉；
2. "报告"一词在多处被使用（如"停止并报告网络问题"），agent 可能将其误解为 reasoning 文本输出而非强制 WriteFile；
3. prompt 要求 agent 精确计数"第 X 个 tool call"，这对 LLM 较困难，容易错过检查点。
**改进建议**:
1. 在 `prompts/reviewer.md` "快速开始"章节开头增加**全局 WriteFile 铁律**，明确声明"本任务不存在以纯文本输出作为终点的选项"，消除语义歧义（修改方式: 插入新段落）。
2. 将 PR 信息获取失败等场景中的"停止并报告"统一改为"**执行 WriteFile** 写入失败报告"，用动词替换名词，降低误解概率（修改方式: 替换 3 处表述）。
3. 更新第 319 行过时统计注释，将 `"successRate 为 0%"` 改为 `"successRate 约 50%，仍有 50% 任务零产出"`，使 agent 感知到当前风险（修改方式: 字符串替换）。
4. 明确 `test -f` 的使用边界：项目地图（AGENTS.md）的初始确认允许 `test -f`，规范文档禁止 `test -f`，消除第 283 行与第 309 行的规则矛盾（修改方式: 在第 309 行增加例外说明）。

### 问题 2: rework 任务违反 Git 禁令（topGitOp = checkout）
**证据**: rework agent `count=1`, `topGitOp="checkout"`, `totalErrors=1`, `avgSteps=25`。developer.md 明确将 `git checkout` 列为 rework 模式绝对禁止的命令，但 agent 仍然执行。
**根因分析**: developer.md 虽有详细的 Git 禁令和自检要求，但缺乏**执行前强制确认**机制。agent 在遇到环境异常或任务指令冲突时，没有在进入工具调用前完成"我不会执行被禁命令"的显式声明，导致禁令被突破。
**改进建议**:
1. 在 `prompts/developer.md` 的 rework 专属规则中增加**第 0.5 步 — Git 禁令前置确认**，要求在任何 tool call 前在 reasoning 中显式声明：「已确认本任务不会执行任何被禁的 git 命令（特别是 `git checkout`）」（修改方式: 在"第 0 步"后插入新步骤）。
2. 更新 developer.md 第 21 行的过时统计注释，将 `"24.2 步、0.22 次错误、10.6 次 Shell"` 更新为 `"约 25 步、约 1 次错误/任务"`，保留 Git 禁令警告（修改方式: 字符串替换）。
3. 强化 rework 错误计数达到 1 时的「只读/收尾模式」，将允许范围从"ReadFile、WriteFile、git 操作"收紧为"ReadFile、WriteFile、`git add`/`commit`/`push`"，**明确禁止 Think 和未经确认的 StrReplaceFile 尝试**（修改方式: 替换第 198 行附近描述）。

### 问题 3: rework workflow 对 Git 禁令违规缺乏硬阻断
**证据**: `workflows/rework.yaml` 的 `diagnoseAgentRun` 检测到 git usage 后仅追加声明到报告，未阻止后续 `ensureCommitAndPush` 步骤执行。
**根因分析**: workflow 层面的违规处理停留在"记录日志"，没有形成熔断机制。即使 agent 违规使用了 `git checkout`，workflow 仍可能将修改提交到远程，导致不可控的代码变更。
**改进建议**:
1. 在 `workflows/rework.yaml` 的 `diagnoseAgentRun` 中，检测到 git usage 时除追加声明外，额外创建标记文件 `logs/rework-git-violation.flag`（修改方式: 在现有 if 块内增加 `touch` 命令）。
2. 在 `preCommitChecks` 之前新增 `checkGitViolation` 步骤，检测到标记文件存在时直接 `exit 1`，以 `ignoreError: false` 强制终止 workflow，阻止违规代码提交（修改方式: 插入新步骤）。

## 已应用改进

本次分析后，已直接修改以下文件：

1. **`prompts/reviewer.md`**
   - 在"快速开始"开头插入全局 WriteFile 铁律，明确禁止纯文本输出作为终点；增加"步数不确定时立即 WriteFile"的保守策略。
   - 将 PR 获取失败场景中的"停止并报告"统一改为"执行 WriteFile 写入失败报告"（3 处）。
   - 更新第 319 行过时统计：`successRate 为 0%` → `successRate 约 50%，仍有 50% 任务零产出`。
   - 统一文档存在性检查规则：明确第 2 步确认项目地图时的 `test -f` 为例外。

2. **`prompts/developer.md`**
   - 更新第 21 行 rework 统计注释，移除已改善的 Shell 数据，保留 Git 禁令警告。
   - 在 rework 专属规则中增加"第 0.5 步 — Git 禁令前置确认"，要求执行任何工具前显式声明遵守禁令。
   - 收紧错误计数达到 1 时的只读模式，明确禁止 Think 和未经确认的 StrReplaceFile。

3. **`workflows/rework.yaml`**
   - 在 `diagnoseAgentRun` 的 Git 检测逻辑中增加 `touch logs/rework-git-violation.flag`。
   - 在 `preCommitChecks` 前新增 `checkGitViolation` 步骤，检测到标记文件时 `exit 1` 硬阻断提交。

## 趋势追踪

- **与上次对比**: review 任务 successRate 从历史 0% 提升至 50%，说明前期对 WriteFile 的约束已产生效果，但仍有 50% 任务零产出，需继续收紧。rework 任务的步数（25 步）与历史均值（24.2 步）持平，Git 禁令违规（checkout）仍未消除。
- **建议下次重点关注**:
  - review 任务的 WriteFile 执行率是否从 50% 提升至更高水平；
  - rework 任务在 prompt 和 workflow 双重加固后，git checkout 违规是否被消除；
  - review 任务的 `totalErrors` 能否从 3 次/4 任务降至更低。
