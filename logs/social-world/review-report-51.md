## 审查结论

审查结论: **NEEDS_FIX**

PR #51 是一个 `design/issue-43` 分支的 Design-Only PR，产出 Matching 模块的 Flutter 纯 UI 设计资产。代码质量整体较高：组件结构清晰、Accessibility（Semantics）覆盖完整、Golden File 截图已提交、Widgetbook 用例覆盖较全。但存在 **4 个阻塞级范围违规** 和 **2 个严重级代码缺陷**，必须修复后才能合并。

---

## 阻塞问题（Blocking）

### 1. Design-Only PR 直接修改了 `packages/design-system/lib/src/theme/tokens.dart`
- **位置**: `packages/design-system/lib/src/theme/tokens.dart:50-56`
- **问题**: PR 直接向 `packages/` 写入了 5 个新的 Design Token（`likeActionColor`、`likeGradientStart`、`likeGradientEnd`、`overlay`、`onPrimary`）。
- **影响**: 违反了 Design-Only PR 的边界约束。AGENTS.md 明确声明：
  > "`.dart` 资产必须放在 `designs/issue-XXX/flutter-widgets/`，**禁止直接写入 `packages/` 或 `apps/`**"
  
  以及 TEMPLATE-design.md 的「范围边界」：
  > "Flutter Widget 代码: `designs/issue-XXX/flutter-widgets/`（**禁止**直接写入 `packages/` 或 `apps/`）"
- **修复**: 
  1. `git checkout a5f6ca6 -- packages/design-system/lib/src/theme/tokens.dart` 恢复原文件。
  2. 在 `MANIFEST.md` 的「设计 Token 变更」章节中保留迁移说明（已有），由后续 feat PR 统一写入 `packages/design-system/`、`docs/design-docs/ui-style-guide/tokens.md` 和 `assets/tokens.json`。
- **依据**: AGENTS.md 第 37 行 / `docs/exec-plans/templates/TEMPLATE-design.md` 第 17、25 行

### 2. Design-Only PR 直接修改了 `packages/design-system/widgetbook/main.dart`
- **位置**: `packages/design-system/widgetbook/main.dart:4,329`
- **问题**: PR 向生产包 `packages/design-system/widgetbook/main.dart` 中注入了设计资产导入：
  ```dart
  import '../../../designs/issue-43/flutter-widgets/widgetbook/widgetbook_use_cases.dart';
  ...
  ...matchingDesignDirectories,
  ```
  这使生产包产生了对 `designs/` 目录的跨包相对路径依赖，破坏了包边界隔离。
- **影响**: 同上，违反 Design-Only PR 禁止写入 `packages/` 的强制规范。该导入在 pub publish 或 CI 隔离构建时可能失败。
- **修复**: `git checkout a5f6ca6 -- packages/design-system/widgetbook/main.dart` 恢复原文件。`matchingDesignDirectories` 的合并应在后续 feat PR 中完成。
- **依据**: AGENTS.md 第 37 行 / TEMPLATE-design.md 第 25 行

### 3. 提前提交了无法运行的 Maestro E2E Flow
- **位置**: `.maestro/flows/matching/swipe_and_match.yaml`（新增）
- **问题**: Design-Only PR 的 `.dart` 资产全部存放在 `designs/issue-43/flutter-widgets/`，**并未集成到 `apps/mobile/` 的真实页面中**。此时提交 Maestro Flow 会导致以下问题：
  - `id: "pass_button"`、`id: "like_button"`、`id: "mode_tab_nationwide"` 等定位目标在 App 中不存在。
  - Flow 中 Step 5（进入匹配列表）和边界空态全部被注释掉，说明作者也意识到 UI 尚未接入。
- **影响**: E2E 资产属于「尚未就绪的代码」，混入 Design-Only PR 会增加维护负担，且无法在 CI 中验证。
- **修复**: 删除 `.maestro/flows/matching/swipe_and_match.yaml`，移至后续 feat PR（接入真实路由和 BottomNav 后）再提交。
- **依据**: `docs/design-docs/flutter.md` 第 455-461 行规定 E2E 针对「Mobile 变更」；但 Design-Only PR 的产出位置在 `designs/`，不属于 Mobile 生产代码。

### 4. Design-Only PR 触碰了 `apps/mobile/`
- **位置**: `apps/mobile/assets/images/.gitkeep`（新增）、`apps/mobile/pubspec.lock`（修改）
- **问题**: `apps/mobile/` 是生产 App 目录，Design-Only PR 禁止写入。`.gitkeep` 文件完全无必要（设计规范明确说明「本设计全部使用 Emoji / 矢量图标 / CSS 渐变占位，**无需额外图片资源**」）。`pubspec.lock` 的变更与设计资产无关。
- **修复**: 
  1. 删除 `apps/mobile/assets/images/.gitkeep`
  2. `git checkout a5f6ca6 -- apps/mobile/pubspec.lock`
- **依据**: AGENTS.md 第 37 行 / TEMPLATE-design.md 第 25 行

---

## 严重问题（Major）

### 1. `SwipeCardStack` 动画期间未屏蔽拖拽手势
- **位置**: `designs/issue-43/flutter-widgets/organisms/swipe_card_stack.dart:176-189`
- **问题**: `_onDragStart` 和 `_onDragUpdate` 没有检查 `_isAnimating` 标志。当卡片正在执行飞出动画时，用户再次触碰会重置 `_dragStartX` 和 `_dragPosition`，导致：
  - 动画中的卡片位置被意外覆盖
  - 可能产生重复 swipe 回调
  - 手势状态机混乱
- **修复**: 在 `_onDragStart` 和 `_onDragUpdate` 入口处增加防护：
  ```dart
  void _onDragStart(DragStartDetails details) {
    if (widget.cards.isEmpty || _isAnimating) return;
    ...
  }
  
  void _onDragUpdate(DragUpdateDetails details) {
    if (!_isDragging || _isAnimating) return;
    ...
  }
  ```
- **依据**: 代码 diff 中的状态机逻辑；通用 Flutter 动画手势最佳实践。

### 2. `SwipeCardStack` 的 `onNeedMoreCards` 触发存在 1 张卡片延迟
- **位置**: `designs/issue-43/flutter-widgets/organisms/swipe_card_stack.dart:55-78`
- **问题**: `_onAnimationStatusChanged` 中先调用 `widget.onSwipe`（通知父级移除卡片），再调用 `_maybeRequestMoreCards`。但 `widget.onSwipe` 内部的 `setState` **仅标记脏状态**，在当前帧结束前的同步代码中，`widget.cards` 仍是旧列表。因此：
  - 初始 4 张卡片 → 第 1 次 swipe 时 `widget.cards.length == 4`，不触发
  - 第 2 次 swipe 时 `widget.cards.length == 3`，不触发
  - 第 3 次 swipe 时 `widget.cards.length == 2`，触发（此时用户只剩 1 张可见卡片）
  
  这导致「剩余卡片 < 3 时自动加载」的逻辑延迟了 1 张卡片，与设计规范「若剩余卡片 < 3 张，自动触发下一页加载」不符。
- **修复**: 将 `_maybeRequestMoreCards` 的调用时机移到父级完成 `setState` 之后。推荐方案：不在 `SwipeCardStack` 内部判断阈值，而是让父级在收到 `onSwipe` 后自行根据剩余数量决定是否加载；或者将阈值判断逻辑改为基于「本次 swipe 前的数量 - 1」：
  ```dart
  void _maybeRequestMoreCards() {
    // 父级尚未移除，因此用当前长度 - 1 预估移除后的数量
    if (widget.cards.length - 1 < 3) {
      widget.onNeedMoreCards?.call();
    }
  }
  ```
  并在动画开始前（而非结束后）调用，避免状态不同步。
- **依据**: `designs/issue-43/design-spec.md` 第 5.2 节「滑动后缓存刷新」要求 / `git diff` 中的动画状态机代码。

---

## 轻微问题（Minor）

### 1. `SwActionButton` Widgetbook 缺少「按下态」用例
- **位置**: `designs/issue-43/flutter-widgets/widgetbook/widgetbook_use_cases.dart:20-59`
- **问题**: `docs/design-docs/flutter.md` 第 131 行明确要求 `SwActionButton` 必须覆盖：
  > "默认态（Pass / Like）、禁用态、**按下态**"
  
  当前 Widgetbook 用例只有 Default 和 Disabled，缺少 Pressed。虽然 `SwActionButton` 是 StatelessWidget 且使用 `GestureDetector(onTapDown)` 无内置按下视觉反馈，但设计规范仍要求展示该状态。建议要么补充按下态视觉（如 scale 0.95），要么在 Widgetbook 中增加一个模拟按下的 Stateful 用例。
- **依据**: `docs/design-docs/flutter.md` 第 127-137 行 Matching 模块 Widgetbook 覆盖要求。

### 2. `MatchSuccessDialog` 中 `AnimatedContainer` 无实际动画作用
- **位置**: `designs/issue-43/flutter-widgets/organisms/match_success_dialog.dart:41-43`
- **问题**: `AnimatedContainer` 的 `color` 属性是常量 `DesignTokens.overlay.withOpacity(0.6)`，没有任何状态变化会导致它产生动画。真正的进入动画由内部的 `TweenAnimationBuilder`（scale 动画）完成。`AnimatedContainer` 在此处是冗余的，增加了一层无意义的动画开销。
- **修复**: 将 `AnimatedContainer` 替换为普通 `Container`。

### 3. `MatchListItem` 依赖的 `SwCard` 存在已知 Maestro 兼容性隐患
- **位置**: `designs/issue-43/flutter-widgets/molecules/match_list_item.dart:34`（使用 `SwCard`）
- **问题**: `packages/design-system/lib/src/atoms/sw_card.dart` 内部使用 `InkWell(onTap: onTap)`。根据 `flutter.md`「禁止的写法」第 3 条，`InkWell(onTap)` 在 Maestro E2E 场景中被明确禁止。当前 `MatchListItem` 按设计规范必须使用 `SwCard`，这意味着迁移到生产代码后，匹配列表项的点击在 Maestro 中可能失效。
- **建议**: 在 `MANIFEST.md` 中增加一条迁移注意：「`SwCard` 的 `InkWell(onTap)` 需替换为 `GestureDetector(onTapDown)` 或在外层包裹兼容层，确保 Maestro 可点击」。或者单独提一个 fix PR 改造 `SwCard`。
- **依据**: `docs/design-docs/flutter.md` 第 314-318 行及「禁止的写法」列表。

### 4. PR 中混入大量纯格式化的 Markdown 变更
- **位置**: `prd/v1-matching.md`、`docs/api/matching.md` 等
- **问题**: 多个 Markdown 文件出现了无意义的空行插入和表格列对齐格式化（例如 `prd/v1-matching.md` 在各级标题后插入空行）。这些变更不增加任何语义信息，增大了 diff 噪音，增加了 review 成本。
- **建议**: 回退纯格式化变更，或使用项目统一的 Markdown formatter 在独立 PR 中处理。

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 功能完整性 | ✅ | 覆盖了 PRD / 设计规范中定义的所有页面和状态：HomePage（loading/empty/success/matchCreated）、MatchListPage（loading/empty/success/error）、SwipeCardStack、MatchSuccessDialog、EmptyRecommendationView |
| 代码规范性 | ⚠️ | 组件命名符合 `Sw` 前缀规范；无硬编码颜色（除新增 Token 本身）；但存在范围违规（见 Blocking 1-4） |
| 全链路完整性 | N/A | Design-Only PR 不包含 DB/API/Mobile 全链路，符合预期 |
| 规范符合性 | ❌ | 违反 Design-Only PR 边界规范（直接写入 `packages/` 和 `apps/`） |
| 测试覆盖 | ✅ | Golden Tests 代码完整，6 张 golden PNG 已提交；Widgetbook 用例覆盖了所有新增组件 |
| 文档同步 | ✅ | `docs/api/matching.md`、`docs/modules/matching.md`、`docs/design-docs/flutter.md`、`docs/design-docs/ui-style-guide/components.md`、`docs/design-docs/ui-style-guide/tokens.md`、`prd/v1-matching.md` 均已更新 |

### 特别说明：项目文档内部冲突
`docs/design-docs/flutter.md` 中「Design-Only PR 流程」一节（第 176-183 行）允许修改 `packages/design-system/` 和 `apps/mobile/lib/presentation/`，但 **AGENTS.md**（第 37 行）和 **TEMPLATE-design.md**（第 25 行）均明确 **禁止** 直接写入 `packages/` 或 `apps/`。结合 commit 历史（`a5f6ca6 docs: clarify .dart asset location for Design-Only PRs (#50)`），AGENTS.md 的约束是更新后的权威规则。本次审查以 AGENTS.md 为准。

---

## 风险与建议

1. **范围违规风险**：如果允许 Design-Only PR 直接修改 `packages/design-system/`，后续回滚或调整设计方向时将污染生产包的历史记录。建议严格执行「设计资产 → `designs/` → feat PR 迁移 → `packages/`」的两阶段流程。
2. **手势冲突风险**：`SwipeCardStack` 动画期间未屏蔽手势的问题（Major #1）在快速滑动场景下会导致卡片位置异常，建议在迁移到生产代码前修复。
3. **预加载延迟风险**：`onNeedMoreCards` 的 1 张卡片延迟（Major #2）会导致用户在剩余 1-2 张卡片时才看到 loading，体验不佳。
4. **Maestro 兼容性债务**：`SwCard` 使用 `InkWell(onTap)` 是一个已存在的设计系统兼容性问题，建议在组件迁移前优先修复，避免 MatchListItem 上线后 E2E 无法通过。
5. **文档一致性债务**：`flutter.md` 与 AGENTS.md 在 Design-Only PR 范围上存在矛盾，建议后续提一个 `docs` PR 统一措辞，消除冲突。
