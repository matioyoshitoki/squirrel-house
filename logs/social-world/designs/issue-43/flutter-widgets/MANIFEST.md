# Issue #43 Matching 设计资产迁移清单

> 本清单记录 `flutter-widgets/` 下各组件的迁移目标位置，供后续开发 Agent 参考。

## 组件清单

| 源文件 | 目标路径 | 组件层级 | 说明 |
|--------|----------|----------|------|
| `sw_action_button.dart` | `packages/design-system/lib/src/atoms/sw_action_button.dart` | atom | 圆形操作按钮（Pass/Like） |
| `swipe_card.dart` | `packages/design-system/lib/src/molecules/swipe_card.dart` | molecule | 可拖动推荐用户卡片 |
| `swipe_card_stack.dart` | `packages/design-system/lib/src/organisms/swipe_card_stack.dart` | organism | 卡片堆叠容器 |
| `match_success_dialog.dart` | `packages/design-system/lib/src/organisms/match_success_dialog.dart` | organism | 配对成功全屏弹窗 |
| `match_list_item.dart` | `packages/design-system/lib/src/molecules/match_list_item.dart` | molecule | 匹配列表单项 |
| `empty_recommendation_view.dart` | `apps/mobile/lib/presentation/widgets/empty_recommendation_view.dart` | page-widget | 推荐空状态视图 |
| `home_page.dart` | `apps/mobile/lib/presentation/pages/home_page.dart` | page | 首页（卡片堆叠） |
| `match_list_page.dart` | `apps/mobile/lib/presentation/pages/match_list_page.dart` | page | 匹配列表页 |

## 迁移步骤

1. **Design-System 组件先行**
   - 将 `sw_action_button.dart`、`swipe_card.dart`、`swipe_card_stack.dart`、`match_success_dialog.dart`、`match_list_item.dart` 复制到 `packages/design-system/lib/src/` 对应目录
   - 在 `packages/design-system/lib/social_world_design_system.dart` 中追加 export
   - 在 `packages/design-system/widgetbook/main.dart` 中注册各组件的 use-case

2. **页面级 Widget**
   - 将 `home_page.dart`、`match_list_page.dart` 复制到 `apps/mobile/lib/presentation/pages/`
   - 将 `empty_recommendation_view.dart` 复制到 `apps/mobile/lib/presentation/widgets/`
   - 在 `apps/mobile/lib/presentation/routes.dart` 中注册 `/home` 和 `/messages` 路由

3. **BLoC 实现**
   - 参考 `docs/modules/matching.md` 实现 `MatchingBloc` 和 `MatchListCubit`
   - 将 BLoC 注入到对应页面

## Token 提案

本设计使用了一个新的语义色 `likeActionColor`（`#2E7D32`），需在 `tokens.dart` 中注册：

```dart
static const Color likeActionColor = Color(0xFF2E7D32);
```

## Widgetbook 用例要求

| 组件 | 必须覆盖的用例 |
|------|----------------|
| `SwActionButton` | Pass 默认态、Like 默认态、禁用态 |
| `SwipeCard` | 默认态、拖动中 Like、拖动中 Pass |
| `SwipeCardStack` | 3 张卡片堆叠、1 张卡片、空态 |
| `MatchSuccessDialog` | 默认态 |
| `MatchListItem` | 默认态、长昵称截断 |
