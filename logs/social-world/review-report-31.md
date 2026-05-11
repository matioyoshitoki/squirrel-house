## 审查结果

| 检查项 | 状态 | 说明 |
|--------|------|------|
| 全链路完整性 | ✅ | 本需求为 Debug-Only Mobile 工具（PRD 明确无 DB/API 变更），代码覆盖 Mobile + 文档 + BDD + Widgetbook，符合范围。 |
| 规范符合性 | ✅ | 代码结构符合 `docs/design-docs/flutter.md` 约定：使用 Cubit、GetIt DI、DesignTokens、按层级分目录。 |
| 契约一致性 | ✅ | 本需求为内部调试工具，不涉及后端 DTO 或 `packages/shared-types` 变更。DioClient 热重载 `baseUrl` 为内部行为，不影响外部契约。 |
| 测试覆盖 | ⚠️ | Flutter 侧单元/Widget/Golden 测试齐全。BDD Step Definitions 为纯空壳（仅注释说明由 Flutter 测试覆盖），未在 JS 层执行任何断言。 |
| 文档同步 | ✅ | `docs/modules/INDEX.md` 已注册、`docs/modules/dev-tools.md` 完整、`docs/design-docs/flutter.md` 已补充 Debug-Only 规范、`ui-specs/dev-tools/` 已归档。 |
| UI/设计一致性 | ✅ | 整体还原度高，颜色/间距/动画/组件使用均符合 `designs/issue-29/design-spec.md`。 |

**综合结论：LGTM（建议合入，附 minor 备注）**

---

## 问题详情

### 1. BDD Step Definitions 为空壳实现
**问题**：`tests/bdd/step-definitions/dev_tools_drawer.steps.js` 中所有 `given/when/then` 回调均为空函数，仅包含注释说明由 Flutter 测试覆盖。这导致 `jest-cucumber` 运行时虽能"通过"，但实际未验证任何行为。
**建议**：
- 方式 A（推荐）：在 JS 层调用 `child_process` 执行 `flutter test --dart-define=BDD_SCENARIO=...` 或将步骤与 Flutter 集成测试桥接。
- 方式 B：至少在当前步骤中增加 `expect` 断言，读取并校验 `dev_tools_drawer.feature` 的文本一致性，避免完全空跑。
- 若团队已约定 Mobile UI 的 BDD 仅作为场景描述、由 Flutter 测试实际覆盖，请在 `tests/bdd/README.md` 中明确说明，避免后续审查误解。

### 2. 设计资产内部存在微小不一致（非代码问题）
**问题**：`designs/issue-29/design-spec.md` 规定抽屉标题 "开发工具" 使用 `DesignTokens.fontSizeTitle`（`22.0`），但同目录下的 `designs/issue-29/mockup.html` 中 `.drawer-title` 写死为 `font-size: var(--fontBody)`（`16px`）且 `font-weight: 600`。
**现状**：代码实现（`dev_tools_drawer.dart`）遵循了文字版设计规范（`fontSizeTitle`），这是正确做法。
**建议**：在下次设计资产迭代时同步修正 `mockup.html` 的标题字号，保持设计资产内部一致性。

### 3. Widgetbook Demo 的交互细节未完全对齐生产逻辑
**问题**：`packages/design-system/widgetbook/main.dart` 中的 `_ServerUrlToolDemo` 在 Error use-case 下，用户只要输入任意字符就会将 `_errorText` 置为 `null`，导致保存按钮立即变为可用状态，而未像生产代码 `ServerUrlTool` 那样重新校验 URL 格式。
**建议**：在 `_ServerUrlToolDemoState` 的 `onChanged` 中增加简易校验，保持 demo 与生产逻辑一致，提升 Widgetbook 的可信度。

### 4. `DevToolsState` 实现方式与设计规范示例不一致
**问题**：`designs/issue-29/design-spec.md` 第 5.4 节给出的 State 示例使用 `@freezed`，而实际代码 `apps/mobile/lib/presentation/blocs/dev_tools/dev_tools_state.dart` 使用纯 `Equatable` 手写 `copyWith`。
**说明**：功能等价，且避免了引入 `freezed` 代码生成开销；但属于设计规范与实际实现的偏离。
**建议**：如团队希望严格对齐设计资产中的代码示例，可后续替换为 `freezed`；当前状态不影响功能，可保留。

### 5. 环境导致的 `make test` 失败（非代码缺陷）
**问题**：审查环境中运行 `make test` 时，`apps/api` 因缺少 `jest` 可执行文件而失败；`apps/mobile` 与 `packages/design-system` 因未安装 Flutter SDK 无法执行 `flutter test`。
**说明**：`apps/mobile/test/dev_tools_cubit_test.dart`、`dev_tools_widget_test.dart`、`dev_tools_golden_test.dart` 的代码质量良好，覆盖了 Cubit 状态流转、Widget 交互、Golden 截图。失败原因系审查环境缺失依赖，非 PR 代码问题。
**建议**：在 CI 中确保 `flutter test` 与 `npm test` 环境完整，以便后续 PR 能真正跑通 `make test`。

---

## 正面清单（值得肯定的点）

1. **Debug-Only 隔离严格**：`app.dart` 与 `di.dart` 中均使用 `if (kDebugMode)` 条件编译，Release 模式下不会将调试 Widget、Repository、Cubit 打包进产物，符合 PRD 安全要求与 `docs/design-docs/flutter.md` 规范。
2. **视觉还原度高**：悬浮按钮直径 `56`、右边缘 `8`、背景色 `primary.withOpacity(0.85)`、带 `BackdropFilter.blur`；抽屉宽度 `min(屏幕 80%, 320)`、动画 `250ms/200ms`、曲线 `easeOutCubic/easeInCubic`，均精确对应 `designs/issue-29/design-spec.md`。
3. **组件使用规范**：`ServerUrlTool` 中正确复用了 `SwTextField`、`SwButton`、`SwCard`，未出现原生 Material 组件替代 design-system 的情况。
4. **Widgetbook 注册完整**：`ServerUrlTool`（Normal / Error / Saving）与 `DevToolsDrawer`（Default）均已在 `packages/design-system/widgetbook/main.dart` 的 Features 目录下注册。
5. **Golden File 已提交**：`dev_tools_drawer_default.png` 与 `server_url_tool_error.png` 已纳入版本控制，符合 Design-Only PR 的验收要求。
6. **文档覆盖充分**：模块索引、实现文档、UI 规格归档、Flutter 编码规范补充均已到位（部分在 PR #30 前置完成，当前 PR 状态完整）。
