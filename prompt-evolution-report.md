# Prompt 演进报告 — type-one

生成时间: 2026-06-01 10:00:38+08:00
分析范围: 2026-05-29 13:52:38 至 2026-06-01 10:00:38

## 统计摘要

- 分析日志数: 4
- Agent 类型分布:
  - e2e: 4 任务，avgSteps=14.25，totalErrors=9，successRate=50%，Shell=34（8.5/任务），ReadFile=25（6.25/任务）

## 问题诊断

### 问题 1: e2e 错误率飙升至 50%，Shell 调用冗余严重
**证据**: 
- 4 个 e2e 任务共 9 个错误（平均 2.25 次/任务），successRate 仅 50%。
- errorRateDelta: +0.96，successRateDelta: -0.36，趋势显著恶化。
- Shell 调用 34 次（8.5/任务），占工具调用总量的 39%。而 workflow 的 `checkEnv` 步骤已详尽检查环境并输出 `.e2e-env-status`，agent 却仍需自行执行环境验证命令。

**根因分析**: 
1. **workflow 与 prompt 的职责重叠**：`workflows/e2e.yaml` 的 `checkEnv` 已执行 `adb`、`maestro`、`emulator` 检查并写入 `.e2e-env-status`，但 `prompts/e2e-tester.md` 第 7 步仍强制要求 agent 再次运行 `maestro --version` 和 `adb devices`。重复检查不仅浪费 Shell 预算（每个任务 2+ 次），还在环境已损坏时产生无意义的错误，直接消耗错误熔断配额。
2. **workflow 存在无效步骤**：`checkBackendPorts` 检查 3306/6379（MySQL/Redis）端口，但 mobile E2E（Maestro）完全不依赖本地数据库端口，且 prompt 中从未引用该检查结果。该步骤仅增加日志噪音和 workflow 执行时间，对 agent 决策零贡献。
3. **错误熔断机制过于粗粒度**：累计错误熔断未区分"环境预检错误"与"测试执行错误"。`maestro --version` 的 command_not_found 与 flow 断言失败被同等计数，导致 agent 可能在尚未开始实际测试时就因环境检查失败而熔断退出。

**改进建议**:
1. **已应用**：在 `workflows/e2e.yaml` 中删除 `checkBackendPorts` 步骤，并将 `MAESTRO_VERSION`、`ADB_DEVICES` 写入 `.e2e-env-status`，使 workflow 成为环境信息的唯一权威来源。（涉及文件: `workflows/e2e.yaml`）
2. **已应用**：修改 `prompts/e2e-tester.md` 第 7 步，明确当 `.e2e-env-status` 显示 `ENV_READY=true` 时，agent 必须直接信任 workflow 预检结果，**禁止**再次执行 `maestro --version` 或 `adb devices`；仅当 `ENV_READY=false` 或文件不存在时，直接跳至环境不可用报告。（涉及文件: `prompts/e2e-tester.md`）
3. **已应用**：修改"累计错误熔断"条款，增加例外：第 7 步读取 `.e2e-env-status` 返回的 file_not_found 不计入熔断计数，禁止通过重复执行环境检查命令来"验证"环境状态。（涉及文件: `prompts/e2e-tester.md`）

## 已应用改进

1. **`workflows/e2e.yaml`** — 消除 workflow 层无效步骤，增强环境状态传递
   - 删除 `checkBackendPorts` 步骤（检查 3306/6379 端口），消除与 mobile E2E 无关的检查开销。
   - 在 `checkEnv` 中把 `MAESTRO_VERSION` 和 `ADB_DEVICES` 追加写入 `.e2e-env-status`，使 agent 可直接读取而无需执行 Shell 命令。
   - 预期收益: 每个 e2e 任务减少 2-3 次冗余 Shell 调用（从 8.5 降至约 6），消除环境检查命令失败导致的错误来源。

2. **`prompts/e2e-tester.md`** — 消除 workflow/prompt 职责重叠，优化熔断规则
   - 第 7 步改为"条件信任"模式：`ENV_READY=true` 时直接跳过环境检查进入测试；`ENV_READY=false` 时直接生成环境不可用报告。
   - 累计错误熔断增加"环境预检例外"，`.e2e-env-status` 的 file_not_found 不计入熔断。
   - 预期收益: 减少因重复环境检查产生的错误（预计从 2.25 次/任务降至 <1 次/任务），提升 successRate；释放 Shell 预算用于实际测试执行。

## 趋势追踪

- 与上次对比:
  - **avgStepsDelta: -2.61** ✅ — 平均步数继续下降，说明整体操作在精简。
  - **errorRateDelta: +0.96** ❌ — 错误率大幅恶化，是本次重点关注的紧急问题。
  - **successRateDelta: -0.36** ❌ — e2e 成功率跌至 50%，零产出风险上升。
- 建议下次重点关注:
  - **e2e Shell 使用率**：验证本次修改后，Shell 调用是否从 8.5/任务降至 6 次以下。
  - **e2e 错误率回落**：观察 errorRate 是否从 2.25/任务下降，特别是 `command_not_found` 和 `adb` 相关错误是否消除。
  - **环境不可用报告的产出率**：如果环境确实经常未就绪，需进一步检查 workflow 的 `checkEnv` 逻辑（如 AVD 启动超时、emulator 稳定性），而非让 agent 承担环境修复责任。
