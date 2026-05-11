# Issue #5 IM 设计资产迁移清单

> 本清单记录 `flutter-widgets/` 下各组件的迁移目标位置，供后续开发 Agent 参考。

## 组件清单

| 源文件                           | 目标路径                                                                  | 组件层级    | 说明                       |
| -------------------------------- | ------------------------------------------------------------------------- | ----------- | -------------------------- |
| `chat_bubble.dart`               | `packages/design-system/lib/src/molecules/chat_bubble.dart`               | molecule    | 消息气泡（文本/图片/系统） |
| `chat_input_bar.dart`            | `packages/design-system/lib/src/molecules/chat_input_bar.dart`            | molecule    | 底部输入栏                 |
| `conversation_list_item.dart`    | `packages/design-system/lib/src/molecules/conversation_list_item.dart`    | molecule    | 会话列表项                 |
| `chat_page.dart`                 | `apps/mobile/lib/presentation/pages/chat_page.dart`                       | page        | 聊天页面                   |
| `conversation_list_page.dart`    | `apps/mobile/lib/presentation/pages/conversation_list_page.dart`          | page        | 会话列表页（更新现有）     |

## 迁移步骤

1. **Design-System 组件先行**
   - 将 `chat_bubble.dart`、`chat_input_bar.dart`、`conversation_list_item.dart` 复制到 `packages/design-system/lib/src/molecules/` 对应目录
   - 在 `packages/design-system/lib/social_world_design_system.dart` 中追加 export
   - 在 `packages/design-system/widgetbook/main.dart` 中注册各组件的 use-case

2. **页面级 Widget**
   - 将 `chat_page.dart` 复制到 `apps/mobile/lib/presentation/pages/`
   - 将 `conversation_list_page.dart` 中的内容合并到现有的 `apps/mobile/lib/presentation/pages/conversation_list_page.dart`
   - 在 `apps/mobile/lib/presentation/routes.dart` 中注册 `/chat/:conversationId` 路由

3. **BLoC 实现**
   - 参考 `docs/exec-plans/TASK-im.md` 实现 `ImBloc` 和 `ConversationListCubit`
   - 将 BLoC 注入到对应页面

## Token 提案

本设计使用了以下临时 Token（定义于 `_tokens.dart`），需在后续 feat PR 中迁移至 `packages/design-system/lib/src/theme/tokens.dart`：

```dart
static const Color bubbleSelfBackground = Color(0xFF6750A4);
static const Color bubbleOtherBackground = Color(0xFFF3EDF7);
static const Color bubbleSelfText = Color(0xFFFFFFFF);
static const Color bubbleOtherText = Color(0xFF1C1B1F);
static const Color systemMessageText = Color(0xFF79747E);
static const Color timestampText = Color(0xFF79747E);
static const Color readReceiptText = Color(0xFF6750A4);
static const Color inputBarBackground = Color(0xFFFFFFFF);
static const Color divider = Color(0xFFEEEEEE);
static const Color unreadBadgeBackground = Color(0xFFB3261E);
static const Color unreadBadgeText = Color(0xFFFFFFFF);
static const double fontSizeSmall = 14.0;
```

## Token Migrations

| Token 名                | 变更类型 | 目标文件                                           | 变更内容                                                      | 影响组件                                                          |
| ----------------------- | -------- | -------------------------------------------------- | ------------------------------------------------------------- | ----------------------------------------------------------------- |
| `bubbleSelfBackground`  | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color bubbleSelfBackground = Color(0xFF6750A4);`   | `ChatBubble`                                                      |
| `bubbleOtherBackground` | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color bubbleOtherBackground = Color(0xFFF3EDF7);`  | `ChatBubble`、`ChatInputBar`                                      |
| `bubbleSelfText`        | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color bubbleSelfText = Color(0xFFFFFFFF);`         | `ChatBubble`                                                      |
| `bubbleOtherText`       | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color bubbleOtherText = Color(0xFF1C1B1F);`        | `ChatBubble`                                                      |
| `systemMessageText`     | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color systemMessageText = Color(0xFF79747E);`      | `ChatBubble`                                                      |
| `timestampText`         | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color timestampText = Color(0xFF79747E);`          | `ChatBubble`、`ConversationListItem`                              |
| `readReceiptText`       | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color readReceiptText = Color(0xFF6750A4);`        | `ChatBubble`                                                      |
| `inputBarBackground`    | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color inputBarBackground = Color(0xFFFFFFFF);`     | `ChatInputBar`                                                    |
| `divider`               | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color divider = Color(0xFFEEEEEE);`                | `ChatInputBar`、`ConversationListItem`、`ConversationListPage`    |
| `unreadBadgeBackground` | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color unreadBadgeBackground = Color(0xFFB3261E);` | `ConversationListItem`                                            |
| `unreadBadgeText`       | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const Color unreadBadgeText = Color(0xFFFFFFFF);`        | `ConversationListItem`                                            |
| `fontSizeSmall`         | 新增     | `packages/design-system/lib/src/theme/tokens.dart` | 新增 `static const double fontSizeSmall = 14.0;`                      | `ConversationListItem`、`ConversationListPage`                    |

## Widgetbook 用例要求

| 组件                    | 必须覆盖的用例                                   |
| ----------------------- | ------------------------------------------------ |
| `ChatBubble`            | 自己文本、对方文本、系统消息、发送中、发送失败   |
| `ChatInputBar`          | 空输入、有内容、发送中、超出字数限制             |
| `ConversationListItem`  | 默认态、有未读、长消息截断、图片预览             |

## Golden Test 要求

| 组件                    | 必须覆盖的 Golden File                                   |
| ----------------------- | -------------------------------------------------------- |
| `ChatBubble`            | `chat_bubble_self_text_read.png`                         |
| `ChatBubble`            | `chat_bubble_other_text.png`                             |
| `ChatBubble`            | `chat_bubble_system.png`                                 |
| `ChatBubble`            | `chat_bubble_failed.png`                                 |
| `ConversationListItem`  | `conversation_list_item_default.png`                     |
| `ConversationListItem`  | `conversation_list_item_unread.png`                      |
