# Issue #43 [Matching] V1 Matching 设计规范

> 关联 PRD：[prd/v1-matching.md](../../prd/v1-matching.md)、[prd/v1-mvp.md](../../prd/v1-mvp.md)、[prd/v1-profile.md](../../prd/v1-profile.md)  
> 关联 Issue：#43  
> 技术栈：Flutter (Mobile) + BLoC/Cubit  
> 任务类型：Design-Only PR（设计资产重建）

---

## 1. 页面/场景清单

| 序号 | 页面/场景 | 说明 | 关键状态 |
|------|-----------|------|----------|
| 1 | **首页-推荐卡片** | 用户进入 App 后默认页面，展示可滑动的推荐卡片堆叠 | Loading、正常态（卡片堆叠）、空数据态、错误态 |
| 2 | **滑动操作反馈** | 用户左滑 Pass / 右滑 Like 时的实时视觉反馈 | 拖动中、释放动画、操作成功 |
| 3 | **配对成功弹窗** | 双向喜欢时触发的全屏 Modal，引导用户聊天或继续滑动 | 动画进入、可关闭 |
| 4 | **匹配列表页** | 底部导航「消息」入口内的配对成功用户列表 | Loading、列表态、空态、错误态 |

### 场景流程图

```
HomePage
  ├── 加载中 → CircularProgressIndicator
  ├── 有数据 → SwipeCardStack + 操作按钮
  │       ├── 左滑/点击 ✕ → Pass
  │       ├── 右滑/点击 ♥ → Like
  │       └── 双向 Like → MatchSuccessDialog
  ├── 无数据 → EmptyRecommendationView
  └── 错误 → 错误图标 + 文案 + 重试按钮

MatchListPage (Tab: 消息)
  ├── 加载中 → CircularProgressIndicator
  ├── 有数据 → MatchListItem 列表
  ├── 无数据 → 空状态图标 + 文案
  └── 错误 → 错误图标 + 文案 + 重试按钮
```

---

## 2. 视觉系统

视觉系统严格复用项目已有 Design Token，**禁止新增硬编码色值**。

### 2.1 色板

| Token | 色值 | 使用场景 |
|-------|------|----------|
| `DesignTokens.primary` | `#6750A4` | 主按钮、激活 Tab、底部导航选中态、Like 按钮 hover 强调 |
| `DesignTokens.background` | `#FFFBFE` | 页面背景、Card 背景 |
| `DesignTokens.textPrimary` | `#1C1B1F` | 标题、正文、图标主色 |
| `DesignTokens.error` | `#B3261E` | Pass 操作色、错误提示 |
| `ColorScheme.surface` | `#F3EDF7` | Tab 背景、头像 fallback 背景 |
| `ColorScheme.outline` | `#79747E` | 辅助文字、分割线、次要按钮边框 |

额外语义色（需在 `tokens.dart` / `app_theme.dart` 中注册）：
- `--success` / `likeActionColor`: `#2E7D32` — Like 按钮、配对成功状态

### 2.2 字体

| Token | 值 | 使用场景 |
|-------|-----|----------|
| `DesignTokens.fontSizeHeadline` | `28.0` | 空状态大标题 |
| `DesignTokens.fontSizeTitle` | `22.0` | 卡片姓名、弹窗标题、页面标题 |
| `DesignTokens.fontSizeBody` | `16.0` | 按钮文字、列表主文字 |
| `DesignTokens.fontSizeCaption` | `12.0` | 距离/时间、Bio、底部导航标签 |

### 2.3 间距与圆角

| Token | 值 | 使用场景 |
|-------|-----|----------|
| `DesignTokens.spacingMd` | `16.0` | 页面边距、Card 内边距、列表项间距 |
| `DesignTokens.spacingSm` | `12.0` | Tab 内间距、操作按钮与卡片间距 |
| `DesignTokens.spacingXs` | `8.0` | 行内小间距、文字与图标间距 |
| `DesignTokens.radiusLg` | `16.0` | 推荐卡片圆角、弹窗圆角 |
| `DesignTokens.radiusMd` | `12.0` | 按钮、列表项、小卡片圆角 |
| `DesignTokens.radiusSm` | `8.0` | Tag、Badge 圆角 |

### 2.4 布局尺寸

- **手机画布**：按 375×812pt（iPhone X 基准）设计
- **推荐卡片**：水平边距 16pt，圆角 16pt，图片区高度 340pt，信息区高度自适应
- **卡片堆叠视觉差**：下层卡片 scale(0.95) + translateY(12pt)
- **操作按钮**：直径 64pt，与卡片底部间距 16pt
- **底部导航栏**：高度 80pt（含 Home Indicator 安全区）

---

## 3. 组件清单

### 3.1 需新增 design-system 组件

| 层级 | 组件名 | 文件路径 | 说明 | Widgetbook 用例要求 |
|------|--------|----------|------|---------------------|
| atom | `SwActionButton` | `packages/design-system/lib/src/atoms/sw_action_button.dart` | 圆形操作按钮（Pass/Like 两种变体） | 默认态、禁用态、按下态 |
| molecule | `SwipeCard` | `packages/design-system/lib/src/molecules/swipe_card.dart` | 可拖动/可点击的推荐用户卡片 | 默认态、拖动中(show like badge)、拖动中(show pass badge) |
| organism | `SwipeCardStack` | `packages/design-system/lib/src/organisms/swipe_card_stack.dart` | 卡片堆叠容器，管理多层卡片渲染与手势冲突 | 3张卡片堆叠、1张卡片、空态 |
| organism | `MatchSuccessDialog` | `packages/design-system/lib/src/organisms/match_success_dialog.dart` | 配对成功全屏弹窗，含头像组合、CTA 按钮 | 默认态 |
| molecule | `MatchListItem` | `packages/design-system/lib/src/molecules/match_list_item.dart` | 匹配列表单项，含头像、昵称、时间、聊天入口 | 默认态 |

### 3.2 页面级 Widget

| Widget 名 | 路径 | 说明 |
|-----------|------|------|
| `HomePage` | `apps/mobile/lib/presentation/pages/home_page.dart` | 首页，承载模式切换、卡片堆叠、操作按钮、底部导航 |
| `MatchListPage` | `apps/mobile/lib/presentation/pages/match_list_page.dart` | 匹配列表页，底部导航「消息」Tab 内容 |
| `EmptyRecommendationView` | `apps/mobile/lib/presentation/widgets/empty_recommendation_view.dart` | 推荐为空时的占位视图（可复用） |

### 3.3 禁止直接使用原生组件替代

- 禁止用 `ElevatedButton` / `IconButton` 直接替代 `SwActionButton`
- 禁止在 `pages/` 中写死 `Color(0xFF...)` 或 `BoxDecoration` 颜色值
- 卡片阴影、圆角、间距必须来自 `DesignTokens`

---

## 4. 图片/图标资源

本设计全部使用 Emoji / 矢量图标 / CSS 渐变占位，**无需额外图片资源**。实际开发时：

| 资源位置 | 说明 |
|----------|------|
| 用户头像 | 网络图 `Image.network(card.avatarUrl)`，fallback 到 `SwAvatar` |
| 空状态插图 | 使用 `Icons.search_off` 或 Lottie 动画（待产品确认），MVP 阶段可用 `Icons.person_off` 替代 |
| 操作按钮图标 | Material `Icons.close` (Pass) / `Icons.favorite` (Like)，颜色语义化 |
| 底部导航图标 | Material `Icons.home` / `Icons.chat_bubble` / `Icons.person` |

> **待产品确认**：配对成功弹窗是否使用 Lottie 庆祝动画？当前原型使用 CSS confetti 演示，Flutter 侧可用 `AnimatedContainer` + `Particle` 简易实现，或引入 `lottie` 包。

---

## 5. 交互说明

### 5.1 手势与动画

| 交互 | 触发 | 动画表现 | 时长 | 曲线 |
|------|------|----------|------|------|
| 卡片拖动 | 手指按住并水平移动 | 卡片跟随手指，旋转 angle = dx × 0.05°，Like/Pass badge 透明度渐变 | 实时 | — |
| 释放-Pass | 左滑超过阈值（≈120pt）或点击 ✕ | 卡片向左飞出屏幕，下层卡片 scale 恢复 1.0 | 300ms | `easeOutCubic` |
| 释放-Like | 右滑超过阈值或点击 ♥ | 卡片向右飞出屏幕，触发 API；若双向匹配则弹 MatchSuccessDialog | 300ms | `easeOutCubic` |
| 弹窗进入 | 双向匹配成功 | 背景遮罩 fadeIn + 内容 scale(0.9→1.0) | 300ms | `easeOutBack`（轻微回弹） |
| 弹窗关闭 | 点击「继续滑动」或遮罩 | 背景 fadeOut + 内容 scale(0.95) + opacity(0) | 200ms | `easeInCubic` |
| 列表加载 | 进入 MatchListPage | 顶部 CircularProgressIndicator，数据返回后列表淡入 | 200ms | `easeIn` |

### 5.2 状态与边界

| 状态 | 视觉处理 |
|------|----------|
| 推荐加载中 | 页面居中 `CircularProgressIndicator(color: DesignTokens.primary)` |
| 推荐空数据 | `EmptyRecommendationView`：图标 + 文案「附近暂无可推荐用户」+ 切换全国按钮 |
| 滑动后缓存刷新 | 卡片飞出后，若剩余卡片 < 3 张，自动触发下一页加载（底部小 loading） |
| 网络错误 | `SnackBar` 提示 + 重试按钮（`SwButton.outlined`） |
| 滑动自己（极端边界） | 后端返回 code=1000，前端展示 `SnackBar`「操作异常，请稍后再试」 |

### 5.3 模式切换

- 顶部 Tab 切换「附近 / 同城 / 全国」时：
  1. 当前卡片堆叠淡出（150ms）
  2. 触发 `GET /matching/recommendations?mode=xxx`
  3. 新卡片堆叠淡入（200ms）

---

## 6. BLoC / Event / State 设计

### 6.1 MatchingBloc（首页推荐 + 滑动）

```dart
// matching_bloc.dart
class MatchingBloc extends Bloc<MatchingEvent, MatchingState> { ... }

// matching_event.dart
abstract class MatchingEvent {}
class RecommendationsRequested extends MatchingEvent {
  final String mode; // 'nearby' | 'city' | 'nationwide'
}
class SwipeActionSubmitted extends MatchingEvent {
  final String targetId;
  final SwipeAction action; // like | pass
}
class MatchModalDismissed extends MatchingEvent {}

// matching_state.dart
abstract class MatchingState {}
class MatchingInitial extends MatchingState {}
class MatchingLoadInProgress extends MatchingState {}
class MatchingLoadSuccess extends MatchingState {
  final List<RecommendationCard> cards;
  final String mode;
}
class MatchingLoadEmpty extends MatchingState {
  final String mode;
}
class MatchingSwipeInProgress extends MatchingState {
  final List<RecommendationCard> cards;
  final String targetId;
  final SwipeAction action;
}
class MatchingMatchCreated extends MatchingState {
  final MatchSummary match; // 用于弹窗展示
}
class MatchingFailure extends MatchingState {
  final String message;
}
```

### 6.2 MatchListCubit（匹配列表）

```dart
// match_list_cubit.dart
class MatchListCubit extends Cubit<MatchListState> {
  Future<void> fetchMatches({String? cursor}) async { ... }
}

// match_list_state.dart
abstract class MatchListState {}
class MatchListInitial extends MatchListState {}
class MatchListLoadInProgress extends MatchListState {}
class MatchListLoadSuccess extends MatchListState {
  final List<MatchItem> matches;
  final bool hasMore;
  final String? nextCursor;
}
class MatchListLoadFailure extends MatchListState {
  final String message;
}
```

### 6.3 Socket.io 事件监听

- 在 `HomePage` 的 `initState` 中建立 Socket 连接（或通过全局 BLoC）
- 监听 `match:created` 事件，接收到后向 `MatchingBloc` dispatch `ExternalMatchReceived(matchSummary)`
- `MatchingBloc` 状态流转至 `MatchingMatchCreated`，触发 `MatchSuccessDialog`

---

## 7. 需更新文档

本设计产物完成后，以下文档必须同步新增或更新，以保证后续开发 Agent 有唯一真实来源。

### 7.1 API 文档

| 文档路径 | 更新说明 |
|----------|----------|
| `docs/api/matching.md` | 新增 Matching 模块 API 文档，包含 `GET /matching/recommendations` 与 `POST /matching/actions` 的请求/响应示例、错误码表 |
| `docs/api/socket-events.md` | 新增 Socket.io 事件文档，记录 `match:created` 的 Payload 结构 |

### 7.2 架构/流程文档

| 文档路径 | 更新说明 |
|----------|----------|
| `docs/modules/matching.md` | 更新 Matching 模块实现文档，补全 Flutter 侧的 BLoC、Widget、路由映射关系 |
| `docs/design-docs/flutter.md` | 在「Design-Only PR 验收 Checklist」中补充 Matching 相关组件的 Widgetbook 用例要求（已存在通用规范，本 Issue 执行即可） |

### 7.3 模块索引

| 文档路径 | 更新说明 |
|----------|----------|
| `docs/modules/INDEX.md` | 确认 `matching.md` 已在模块清单中注册（已注册，无需修改） |

### 7.4 PRD 注释

- 在 `prd/v1-matching.md` 顶部添加链接注释：`设计资产参见 designs/issue-43/`

---

## 8. 验收 Checklist（供 Design-Only PR 使用）

- [ ] 低保真/高保真原型图已产出（本文件 + wireframe.svg + mockup.html）
- [ ] 所有新增组件已在 `widgetbook/main.dart` 中注册 use-case
- [ ] PR 中附带 Golden File 截图或 Widgetbook 运行截图
- [ ] `pages/` 中未使用 `Color(0xFF...)`、`TextField`、`ElevatedButton` 等硬编码或原生 Material 组件替代 design-system
- [ ] 页面中的状态（加载、空态、错误、配对成功弹窗）在原型图中有明确定义
