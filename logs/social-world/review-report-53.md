# PR #53 审查报告 — design/issue-43

> **审查分支**: `design/issue-43`  
> **关联 Issue**: #43  
> **关联 PRD**: `prd/v1-matching.md`, `prd/v1-mvp.md`, `prd/v1-profile.md`  
> **审查日期**: 2025-04-21  
> **审查依据**: `AGENTS.md` → `docs/exec-plans/templates/TEMPLATE-design.md`, `docs/design-docs/flutter.md`

---

## 审查结论

**审查结论: PASS**

本 PR 作为 Design-Only PR，整体符合项目设计任务模板（`TEMPLATE-design.md`）与 Flutter 编码规范（`flutter.md`）的要求。设计资产完整、Widget 纯 UI 无业务逻辑、Accessibility 合规、Golden File 已提交。存在 2 个 Minor 文档/工程问题，建议修复后合并，但不阻塞合并。

---

## 阻塞问题（Blocking）

**无**

---

## 严重问题（Major）

**无**

---

## 轻微问题（Minor）

### 1. `widgetbook` 依赖缺失导致 `widgetbook/main.dart` 无法独立编译
- **位置**: `designs/issue-43/flutter-widgets/pubspec.yaml`
- **问题**: `widgetbook/main.dart` 导入了 `package:widgetbook/widgetbook.dart`，但同目录 `pubspec.yaml` 未声明 `widgetbook` 依赖。在该目录下直接执行 `flutter run -t widgetbook/main.dart` 会报依赖缺失错误。
- **影响**: 人类 reviewer 无法在不迁移代码到 `packages/design-system/` 的情况下直接运行 Widgetbook 预览。
- **修复**: 在 `designs/issue-43/flutter-widgets/pubspec.yaml` 的 `dev_dependencies` 中追加：
  ```yaml
  widgetbook: ^3.10.0
  ```
- **依据**: `TEMPLATE-design.md` 将 Widgetbook 用例列为强制产出；`flutter.md` 要求新增组件必须在 Widgetbook 注册。虽然 Golden File 已提供视觉验收（方案 A），但 Widgetbook 代码应保证可编译。

### 2. `docs/api/matching.md` 仍引用旧设计资产路径 `designs/issue-4/`
- **位置**: `docs/api/matching.md:4`
- **问题**: 文件头部注释指向 `../../designs/issue-4/design-spec.md`，而本次 PR 已将正式设计资产迁移至 `designs/issue-43/` 并同步更新了 `prd/v1-matching.md` 的链接。该遗留链接会给后续开发 Agent 造成歧义。
- **影响**: 后续开发 Agent 可能误入已废弃的 `issue-4` 设计资产。
- **修复**: 将 `docs/api/matching.md` 第 4 行修改为：
  ```markdown
  > 关联设计规范: [design-spec.md](../../designs/issue-43/design-spec.md)
  ```
- **依据**: PR 中 `designs/issue-43/design-spec.md` 第 7.1 节明确列出 `docs/api/matching.md` 为需同步更新的文档；`prd/v1-matching.md` 已在本 PR 中完成同类修正。

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 功能完整性 | ✅ | 覆盖 PRD `v1-matching.md` 定义的全部页面与状态：首页（加载/空态/错误/正常/配对成功弹窗）、匹配列表页（加载/空态/错误/正常）、滑动操作反馈、配对成功弹窗。 |
| 代码规范性 | ✅ | 无硬编码 `Color(0xFF...)`（预提交钩子已豁免 `_tokens.dart` 与 `preview.dart`）；组件命名以 `Sw` 为前缀；文件名为 `snake_case.dart`；Token 引用正确，未使用 Legacy alias。 |
| 全链路完整性 | N/A | 本 PR 为 Design-Only，不涉及 DB/API/Mobile 全链路。MANIFEST.md 已列出完整迁移路径。 |
| 规范符合性 | ✅ | 符合 `TEMPLATE-design.md` 的交付范围与文件位置约定；符合 `flutter.md` 的 Accessibility 规范（`onTapDown`、`Semantics` 三要素、`container: true`）。 |
| 测试覆盖 | ✅ | 5 组 Golden Test 已提交 PNG（`sw_action_button_variants`、`swipe_card_default/like/pass`、`swipe_card_stack_3_cards`、`match_list_item_default/long_name`、`match_success_dialog_default`）。 |
| 文档同步 | ⚠️ | `prd/v1-matching.md` 已更新链接；`docs/design-docs/ui-specs/matching/matching-wireframe.md` 已新增；`docs/api/matching.md` 遗留旧链接未修（见 Minor #2）。 |

---

## 详细审查记录

### 1. 范围与边界合规性
- ✅ 所有 `.dart` 文件均位于 `designs/issue-43/flutter-widgets/`，未直接写入 `packages/` 或 `apps/`
- ✅ 无 BLoC、Repository、Dio、网络请求等业务逻辑
- ✅ `MANIFEST.md` 完整列出 8 个文件的迁移目标路径及 Token Migrations 清单

### 2. 视觉系统与 Token
- ✅ `design-spec.md` 以表格形式声明了全部 Token 提案（色值、使用场景、变更类型）
- ✅ `_tokens.dart` 集中定义了所有临时 Token，并标注后续迁移目标
- ✅ `flutter-widgets/` 目录下（除 `_tokens.dart`）无 `Color(0xFF` 硬编码
- ✅ 未使用 Legacy alias（如 `surfaceColor`）
- ⚠️ `preview.dart` 内含硬编码色值，但已被预提交钩子排除，且 `MANIFEST.md` 已声明其自包含特性

### 3. Accessibility / Maestro 兼容性
- ✅ 所有需 Maestro 点击的交互组件均使用 `GestureDetector(onTapDown: ...)`，未使用 `onTap`
- ✅ 所有 `Semantics` 均设置了 `identifier`、`label`、`button: true`、`container: true`
- ✅ `Stack` 中无使用 `Align` 包裹可点击组件的情况；悬浮/覆盖元素使用 `Positioned` 或 `Positioned.fill`
- ✅ 输入框不在本次 PR 范围内，无需检查

### 4. Widgetbook 用例
- ✅ `SwActionButton`: Pass / Like / Disabled
- ✅ `SwipeCard`: Default / Dragging Like / Dragging Pass
- ✅ `SwipeCardStack`: 3 Cards / 1 Card / Empty
- ✅ `MatchSuccessDialog`: Default
- ✅ `MatchListItem`: Default / Long Name
- ⚠️ `widgetbook` 包未在 `pubspec.yaml` 中声明，导致无法独立编译（见 Minor #1）

### 5. Golden Tests
- ✅ `sw_action_button_test.dart` → `sw_action_button_variants.png`
- ✅ `swipe_card_test.dart` → `swipe_card_default.png`, `swipe_card_like.png`, `swipe_card_pass.png`
- ✅ `swipe_card_stack_test.dart` → `swipe_card_stack_3_cards.png`
- ✅ `match_list_item_test.dart` → `match_list_item_default.png`, `match_list_item_long_name.png`
- ✅ `match_success_dialog_test.dart` → `match_success_dialog_default.png`

### 6. 设计资产完整性
- ✅ `wireframe.svg`（线框图）
- ✅ `mockup.html`（高保真原型）
- ✅ `design-spec.md`（设计规范 + Token 提案）
- ✅ `preview.dart`（独立 Flutter Web 预览）
- ✅ `MANIFEST.md`（迁移清单）

---

## 风险与建议

1. **Token 迁移风险**：`MANIFEST.md` 列出了 13 个新增 Token（如 `likeActionColor`、`surface`、`outline` 等），后续 `feat` PR 需逐一迁入 `packages/design-system/lib/src/theme/tokens.dart`。建议后续开发 Agent 在迁移完成后在 `MANIFEST.md` 中打勾确认。
2. **底部导航重复代码**：`home_page.dart` 与 `match_list_page.dart` 均内联了 `_BottomNav` 私有组件。`MANIFEST.md` 已要求在 feat PR 中提取为公共 `bottom_nav.dart`，设计阶段无需处理。
3. **Dart SDK 约束漂移**：`packages/design-system/pubspec.lock` 中 Dart SDK 下限从 `>=3.5.0` 变为 `>=3.9.0-0`，系本地 `flutter pub get` 的副作用。若 CI 环境 Flutter 版本低于 3.24.x，可能导致构建失败。建议在合并前确认 CI 镜像版本，或在 PR 中剔除无关的 `pubspec.lock` 变更。
4. **配对成功弹窗动画**：当前 `MatchSuccessDialog` 的 `TweenAnimationBuilder` 对整个 overlay（含背景遮罩）做了 `scale` 变换，而 `design-spec.md` 要求“背景遮罩 fadeIn + 内容 scale”。建议在 feat PR 迁移时将动画拆分为：外层 `AnimatedOpacity`（背景）+ 内层 `ScaleTransition`（内容卡片）。

---

## 附：引用规范来源

- `AGENTS.md` 第 37 行：Design-Only PR 产出位置约定
- `docs/exec-plans/templates/TEMPLATE-design.md`：Design-Only PR 范围边界、Token 变更交付规范、验收 Checklist
- `docs/design-docs/flutter.md` 第 48 行：组件 `Sw` 前缀命名；第 92–133 行：Design-Only PR 前置规范；第 154–172 行：Token 使用指南；第 252–320 行：Accessibility / Maestro 规范
