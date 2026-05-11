# PR #56 代码审查报告

> **分支**: `design/issue-5`  
> **类型**: Design-Only PR（IM 模块 UI 设计资产）  
> **关联 Issue**: #5  
> **审查日期**: 2026-04-24  
> **审查依据**: `AGENTS.md`、`docs/design-docs/flutter.md`、`docs/exec-plans/templates/TEMPLATE-design.md`、`prd/v1-im.md`

---

## 审查结论

**审查结论: PASS**

本 PR 是一份高质量的 Design-Only PR，完整产出了 Issue #5（IM 模块）所需的设计资产：
- 高保真原型 (`mockup.html`)
- 设计规范文档 (`design-spec.md`)
- 纯 UI Widget 实现（全部位于 `designs/issue-5/flutter-widgets/`，未污染 `packages/` 或 `apps/`）
- Widgetbook 用例代码（覆盖全部分子组件和页面骨架）
- Golden Test 代码 + 截图基线（10 张 PNG）
- 迁移清单 `MANIFEST.md`（含 Token 提案和迁移路径）
- 关联 PRD 链接更新（`prd/v1-mvp.md`、`prd/v1-matching.md`）
- UI 规范文档 (`docs/design-docs/ui-specs/im/im-wireframe.md`)

代码风格、Token 使用、Accessibility（Maestro 兼容性）均符合项目规范。仅存在 2 个轻微问题，不影响合并。

---

## 阻塞问题（Blocking）

无。

---

## 严重问题（Major）

无。

---

## 轻微问题（Minor）

### 1. `SwConversationListItem` 的 `Semantics` identifier 为固定值，列表场景下 Maestro 定位冲突
- **位置**: `designs/issue-5/flutter-widgets/molecules/sw_conversation_list_item.dart:45`
- **问题**: `identifier: 'conversation_list_item'` 是硬编码字符串。当会话列表中存在多个条目时，Maestro `tapOn: { id: "conversation_list_item" }` 永远只能匹配到第一个条目，无法定位后续条目。
- **影响**: 未来编写 E2E 测试（如 `tests/bdd/im.feature` 中 "Alice opens the conversation with Bob"）时，无法通过 `id` 精确定位特定会话。
- **修复**: 在 `SwConversationListItem` 增加可选的 `String? identifier` 参数，默认回退到 `'conversation_list_item'`，并在 `conversation_list_page.dart` 的 `ListView.itemBuilder` 中传入带索引的唯一标识，例如 `identifier: 'conversation_list_item_$index'`。
- **依据**: `docs/design-docs/flutter.md` 第 256–279 行（Semantics 三要素与 Maestro 定位规范）

### 2. `ConversationListPage` 空状态引导按钮未使用 `SwButton`
- **位置**: `designs/issue-5/flutter-widgets/pages/conversation_list_page.dart:119–139`
- **问题**: 空状态的「去首页发现」按钮使用裸 `GestureDetector` + `Container` 实现，而非 `SwButton(variant: SwButtonVariant.filled)`。
- **影响**: 与项目 design-system 一致性有偏差；后续 feat PR 迁移到 `apps/mobile/` 时需额外重构为 `SwButton`。
- **修复**: 替换为 `SwButton`：
  ```dart
  SwButton(
    text: '去首页发现',
    variant: SwButtonVariant.filled,
    onPressed: () {},
  )
  ```
  若需要大圆角（`radiusXl`）而 `SwButton` 不支持，可在 `MANIFEST.md` 中补充 "`SwButton` 需支持自定义圆角" 的迁移项。
- **依据**: `docs/design-docs/flutter.md` 第 90–133 行（页面级 UI 开发前置规范：禁止在 `pages/` 中使用原生 Material 组件替代 design-system）

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 功能完整性 | ✅ | 覆盖了 `prd/v1-im.md` 定义的两个核心页面（会话列表页、聊天页）及其全部关键状态：Loading、正常态、空态、错误态、发送中、发送失败、已读。 |
| 代码规范性 | ✅ | 文件名使用 `snake_case`，类名使用 `PascalCase`，Widget 类名以功能结尾；未使用 Legacy aliases（如 `surfaceColor`）。 |
| 全链路完整性 | N/A | Design-Only PR 不涉及 DB/API/Mobile 全链路，仅产出设计资产。后续 `feat` PR 需按 `MANIFEST.md` 迁移。 |
| 规范符合性 | ✅ | 颜色/字号/间距/圆角均来自 `DesignTokens` 或已在 `design-spec.md` / `MANIFEST.md` 中声明的临时 Token；`_tokens.dart` 的硬编码色值已作为 Token 提案记录。 |
| 测试覆盖 | ✅ | Golden Tests 覆盖 `SwChatBubble`（4 状态）、`SwConversationListItem`（2 状态）、`SwReadReceiptIndicator`（4 状态），共 10 张基线 PNG 已提交到 PR 中。 |
| 文档同步 | ✅ | `prd/v1-mvp.md`、`prd/v1-matching.md` 已更新设计资产引用；`docs/design-docs/ui-specs/im/im-wireframe.md` 已创建；`MANIFEST.md` 完整列出迁移清单和 Token 迁移表。 |

### Design-Only PR 专项检查项

| 检查项 | 状态 | 证据 |
|--------|------|------|
| 产出仅位于 `designs/issue-XXX/`，未写入 `packages/` 或 `apps/` | ✅ | `git diff HEAD~1 --name-only` 中无 `packages/` 或 `apps/` 路径 |
| 包含设计规范文档 `design-spec.md` | ✅ | `designs/issue-5/design-spec.md` 存在，含色板、字体、间距、组件清单、交互说明 |
| 包含迁移清单 `MANIFEST.md` | ✅ | `designs/issue-5/flutter-widgets/MANIFEST.md` 存在，含组件清单、Token Migrations、Widgetbook/Golden 要求 |
| 包含 Widgetbook 用例代码 | ✅ | `widgetbook/main.dart` 注册了 4 个组件 + 2 个页面骨架，共 18 个 use-case |
| 包含 Golden Test 代码 + 截图 | ✅ | `goldens/` 目录下 10 张 PNG + 3 个 `_test.dart` |
| 代码中无未声明的硬编码色值 | ✅ | 全局搜索 `Color(0xFF` 仅在 `_tokens.dart` 中命中，且已在 `design-spec.md` 和 `MANIFEST.md` 中声明 |
| 所有 Maestro 可点击组件使用 `onTapDown` | ✅ | `sw_chat_bubble.dart`、 `sw_chat_input_bar.dart`、 `sw_conversation_list_item.dart`、 `chat_page.dart`、 `conversation_list_page.dart` 均使用 `onTapDown` |
| `Stack` 中未使用 `Align` 包裹点击组件 | ✅ | `sw_read_receipt_indicator.dart` 的 `_DoubleCheck` 使用 `Positioned` + 显式 `width`/`height` |
| 输入框暴露 `Semantics(identifier: 'xxx')` | ✅ | `sw_chat_input_bar.dart` 中 `TextField` 外层包裹 `Semantics(identifier: 'chat_input_field', ...)` |
| 输入框获得焦点时自动全选旧内容 | ✅ | `chat_page.dart` 中 `_focusNode.addListener` 实现了全选逻辑 |
| PRD 链接已更新 | ✅ | `prd/v1-mvp.md`、`prd/v1-matching.md` 顶部新增 `设计资产参见 designs/issue-5/` |
| `docs/design-docs/ui-specs/` 已更新 | ✅ | 新增 `docs/design-docs/ui-specs/im/im-wireframe.md` |

---

## 风险与建议

1. **Token 迁移追踪**: `MANIFEST.md` 中列出了 11 个临时 Token 需迁移至 `packages/design-system/lib/src/theme/tokens.dart`。建议在后续 `feat` PR 中将 `_tokens.dart` 的内容一次性迁入，并删除临时文件，避免设计资产与正式代码长期并存导致 Token 来源混乱。

2. **`SwChatInputBar` 直接使用 `TextField` 的备注**: `MANIFEST.md` 已充分说明使用 `TextField` 而非 `SwTextField` 的理由（胶囊形无框输入框视觉需求）。建议在后续 feat PR 中评估是否将 `SwTextField` 增强为支持 `borderSide: BorderSide.none` 和自定义圆角，以便 `SwChatInputBar` 能回退到 design-system 组件。

3. **页面骨架的导航方式**: `chat_page.dart` 当前使用 `Navigator.of(context).maybePop()` 返回，后续迁移到 `apps/mobile/` 时需替换为 `GoRouter` 的 `context.pop()` 或指定路由返回，以符合 `apps/mobile/lib/presentation/routes.dart` 的路由管理规范。

4. **建议保持当前高质量标准**: 本 PR 的 Widgetbook 用例覆盖度和 Golden Test 完整度优于部分历史 Design-Only PR（如 Issue #43），可作为后续 design 任务的标杆。
