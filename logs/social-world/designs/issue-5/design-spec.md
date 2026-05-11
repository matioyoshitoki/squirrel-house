# Issue #5 [IM] V1 Instant Messaging 设计规范

> 关联 PRD：[prd/v1-mvp.md](../../prd/v1-mvp.md)、[prd/v1-matching.md](../../prd/v1-matching.md)  
> 关联执行计划：[docs/exec-plans/TASK-im.md](../../docs/exec-plans/TASK-im.md)  
> 关联 Issue：#5  
> 技术栈：Flutter (Mobile) + BLoC/Cubit  
> 任务类型：Design-Only PR（设计资产）

---

## 1. 页面/场景清单

| 序号 | 页面/场景           | 说明                                              | 关键状态                                                    |
| ---- | ------------------- | ------------------------------------------------- | ----------------------------------------------------------- |
| 1    | **会话列表页**      | 展示所有匹配成功的会话，含最后消息预览与未读角标  | Loading、正常态（列表）、空态、错误态                       |
| 2    | **聊天页**          | 一对一实时聊天界面，消息气泡 + 底部输入栏         | Loading（历史消息）、正常态（消息列表）、空态（无历史消息） |
| 3    | **消息气泡**        | 文本/图片/系统消息三种类型的气泡展示              | 发送中、已发送、已送达、已读、发送失败                      |
| 4    | **底部输入栏**      | 文本输入 + 图片选择 + 发送按钮组合                | 空输入（发送禁用）、有内容（发送可用）、发送中              |
| 5    | **会话列表项**      | 列表中的单个会话条目                              | 有未读、无未读、最后消息预览                                |

### 场景流程图

```
ConversationListPage (Tab: 消息)
  ├── 加载中 → CircularProgressIndicator
  ├── 有数据 → ConversationListItem 列表
  │       └── 点击 → ChatPage
  ├── 无数据 → 空状态图标 + 文案「还没有匹配成功的好友」
  └── 错误 → 错误图标 + 文案 + 重试按钮

ChatPage
  ├── 加载历史消息 → 顶部 CircularProgressIndicator
  ├── 消息列表 → ChatBubble（对方左/自己右）+ 系统消息居中
  │       ├── 文本消息 → 圆角气泡 + 文字
  │       ├── 图片消息 → 圆角图片占位/缩略图
  │       └── 系统消息 → 居中灰色小字
  ├── 底部输入栏 → ChatInputBar（文本 + 图片 + 发送）
  └── 空历史 → 系统消息「你们已匹配成功，开始聊天吧」
```

---

## 2. 视觉系统

视觉系统严格复用项目已有 Design Token，**禁止新增硬编码色值**。

### 2.1 色板

| Token                                    | 色值      | 使用场景                                                |
| ---------------------------------------- | --------- | ------------------------------------------------------- |
| `DesignTokens.primary`                   | `#6750A4` | 自己发送的气泡背景、发送按钮、未读角标                  |
| `DesignTokens.background`                | `#FFFBFE` | 页面背景                                                |
| `DesignTokens.textPrimary`               | `#1C1B1F` | 标题、正文                                              |
| `DesignTokens.error`                     | `#B3261E` | 发送失败标识、重试按钮                                  |
| `ColorScheme.surface`                    | `#F3EDF7` | 对方气泡背景、输入框背景                                |
| `ColorScheme.outline`                    | `#79747E` | 辅助文字、时间戳、已读状态                              |
| `Tokens.bubbleSelfBackground`            | `#6750A4` | 自己消息气泡背景（与 primary 一致）                     |
| `Tokens.bubbleOtherBackground`           | `#F3EDF7` | 对方消息气泡背景（与 surface 一致）                     |
| `Tokens.bubbleSelfText`                  | `#FFFFFF` | 自己消息气泡文字色                                      |
| `Tokens.bubbleOtherText`                 | `#1C1B1F` | 对方消息气泡文字色（与 textPrimary 一致）               |
| `Tokens.systemMessageText`               | `#79747E` | 系统消息文字色                                          |
| `Tokens.timestampText`                   | `#79747E` | 时间戳文字色                                            |
| `Tokens.readReceiptText`                 | `#6750A4` | 已读回执文字色                                          |
| `Tokens.inputBarBackground`              | `#FFFFFF` | 底部输入栏背景                                          |
| `Tokens.divider`                         | `#EEEEEE` | 列表分割线、输入栏顶部边框                              |
| `Tokens.unreadBadgeBackground`           | `#B3261E` | 未读数角标背景                                          |
| `Tokens.unreadBadgeText`                 | `#FFFFFF` | 未读数角标文字                                          |

### 2.2 字体

| Token                           | 值     | 使用场景                             |
| ------------------------------- | ------ | ------------------------------------ |
| `DesignTokens.fontSizeHeadline` | `28.0` | 空状态大标题                         |
| `DesignTokens.fontSizeTitle`    | `22.0` | 页面标题（聊天页对方昵称）           |
| `DesignTokens.fontSizeBody`     | `16.0` | 消息正文、输入框文字、列表主文字     |
| `DesignTokens.fontSizeCaption`  | `12.0` | 时间戳、已读状态、系统消息、列表副文字 |

### 2.3 间距与圆角

| Token                    | 值     | 使用场景                          |
| ------------------------ | ------ | --------------------------------- |
| `DesignTokens.spacingMd` | `16.0` | 页面边距、气泡与屏幕边缘间距      |
| `DesignTokens.spacingSm` | `12.0` | 气泡内部内边距、列表项内间距      |
| `DesignTokens.spacingXs` | `8.0`  | 头像与气泡间距、行内小间距        |
| `DesignTokens.spacingXxs`| `4.0`  | 状态图标与文字间距                |
| `DesignTokens.radiusLg`  | `16.0` | 图片气泡圆角                      |
| `DesignTokens.radiusMd`  | `12.0` | 文本气泡圆角                      |
| `DesignTokens.radiusSm`  | `8.0`  | 系统消息气泡圆角                  |

### 2.4 布局尺寸

- **手机画布**：按 375×812pt（iPhone X 基准）设计
- **消息气泡最大宽度**：屏幕宽度的 70%
- **头像尺寸**：聊天页 40pt，会话列表 56pt
- **输入栏高度**：56pt（输入框）+ 安全区
- **未读角标**：最小 20×20pt，圆角 10pt
- **系统消息**：水平内边距 16pt，垂直 8pt，圆角 12pt

---

## 3. 组件清单

### 3.1 需新增 design-system 组件

| 层级     | 组件名                  | 文件路径                                                              | 说明                                         | Widgetbook 用例要求                                      |
| -------- | ----------------------- | --------------------------------------------------------------------- | -------------------------------------------- | -------------------------------------------------------- |
| molecule | `ChatBubble`            | `packages/design-system/lib/src/molecules/chat_bubble.dart`           | 消息气泡（文本/图片/系统三种类型）           | 自己文本、对方文本、系统消息、图片消息、发送失败态       |
| molecule | `ChatInputBar`          | `packages/design-system/lib/src/molecules/chat_input_bar.dart`        | 底部输入栏（文本输入+图片选择+发送按钮）     | 空输入、有内容、发送中                                   |
| molecule | `ConversationListItem`  | `packages/design-system/lib/src/molecules/conversation_list_item.dart`| 会话列表单项（头像+昵称+最后消息+未读+时间） | 默认态、有未读、长消息截断                               |
| organism | `ChatMessageList`       | `packages/design-system/lib/src/organisms/chat_message_list.dart`     | 消息列表容器（处理滚动、分组、加载更多）     | 混合消息列表、空列表                                     |

### 3.2 页面级 Widget

| Widget 名                 | 路径                                                                  | 说明                                    |
| ------------------------- | --------------------------------------------------------------------- | --------------------------------------- |
| `ConversationListPage`    | `apps/mobile/lib/presentation/pages/conversation_list_page.dart`      | 会话列表页（更新现有 MatchListPage）    |
| `ChatPage`                | `apps/mobile/lib/presentation/pages/chat_page.dart`                   | 聊天页面                                |

### 3.3 禁止直接使用原生组件替代

- 禁止用 `TextField` / `ElevatedButton` 直接替代 `ChatInputBar`
- 禁止在 `pages/` 中写死 `Color(0xFF...)` 或 `BoxDecoration` 颜色值
- 消息气泡阴影、圆角、间距必须来自 `DesignTokens`

---

## 4. 图片/图标资源

本设计全部使用 Material 矢量图标 / Emoji / CSS 渐变占位，**无需额外图片资源**。

| 资源位置       | 说明                                                      |
| -------------- | --------------------------------------------------------- |
| 用户头像       | 网络图 `Image.network(...)`，fallback 到 `SwAvatar`       |
| 发送按钮       | Material `Icons.send`，主色                               |
| 图片选择       | Material `Icons.image`，outline 色                        |
| 发送失败重试   | Material `Icons.error_outline`，error 色                  |
| 空状态插图     | `Icons.chat_bubble_outline`，placeholder 色               |
| 返回按钮       | Material `Icons.arrow_back_ios`，textPrimary 色           |

---

## 5. 交互说明

### 5.1 手势与动画

| 交互           | 触发                           | 动画表现                                                          | 时长  | 曲线           |
| -------------- | ------------------------------ | ----------------------------------------------------------------- | ----- | -------------- |
| 发送消息       | 点击发送按钮                   | 气泡从底部滑入（opacity 0→1, translateY 20→0）                    | 200ms | `easeOutCubic` |
| 收到消息       | Socket 推送                    | 气泡从底部滑入，同时播放轻微震动反馈                              | 200ms | `easeOutCubic` |
| 历史消息加载   | 上拉至顶部                     | 顶部显示 loading，加载完成后列表向下展开                          | 200ms | `easeIn`       |
| 图片预览       | 点击图片消息                   | 全屏 Modal 淡入，支持左右滑动切换、捏合缩放（V1 简化：仅全屏展示）| 300ms | `easeOutCubic` |
| 输入框聚焦     | 点击输入框                     | 键盘弹出，消息列表自动滚动至底部                                  | 即时  | —              |
| 已读回执更新   | 对方打开聊天页                 | 状态文字从「已发送」渐变过渡到「已读」（opacity 闪烁一次）        | 150ms | `easeInOut`    |

### 5.2 状态与边界

| 状态                 | 视觉处理                                                                   |
| -------------------- | -------------------------------------------------------------------------- |
| 历史消息加载中       | 列表顶部显示 `CircularProgressIndicator`                                   |
| 消息发送中           | 自己气泡右下角显示小型 `CircularProgressIndicator`（12pt）                 |
| 消息发送失败         | 气泡左侧显示红色 `Icons.error_outline`，点击可重发                         |
| 消息已读             | 气泡右下角显示「已读」小字（caption，primary 色）                          |
| 会话列表空数据       | 居中图标 + 文案「还没有匹配成功的好友」+ 引导按钮「去首页发现」            |
| 聊天页空历史         | 仅显示系统消息「你们已匹配成功，开始聊天吧」                               |
| 网络错误             | `SnackBar` 提示 + 重试按钮                                                 |
| 输入框为空           | 发送按钮置灰（opacity 0.4），不可点击                                      |
| 输入文字超过 500 字  | 输入框底部显示红色提示「已超过 500 字限制」，发送按钮禁用                  |

### 5.3 消息分组规则

- 同一发送者连续消息，头像仅在第一条显示
- 同一分钟内发送的消息，时间戳仅在最后一条显示
- 跨天消息自动插入日期分割线（居中灰色小字「2024年6月1日」）

### 5.4 已读回执规则

- 自己发送的消息：显示状态指示器（发送中 → 已发送 → 已送达 → 已读）
- MVP 阶段简化为：发送中（spinner）→ 已发送（无标识）→ 已读（「已读」文字）
- 对方发送的消息：不显示状态标识

---

## 6. BLoC / Event / State 设计（供开发参考）

> 本章节仅作为开发 Agent 的参考，Design-Only PR 不产出 BLoC 代码。

### 6.1 ImBloc（聊天页 + Socket 连接）

```dart
// im_event.dart
abstract class ImEvent {}
class ConnectSocket extends ImEvent {}
class DisconnectSocket extends ImEvent {}
class SendMessage extends ImEvent { final String conversationId; final String content; final String type; }
class ReceiveMessage extends ImEvent { final Message message; }
class MarkAsRead extends ImEvent { final String conversationId; final String lastMessageId; }
class LoadMessages extends ImEvent { final String conversationId; final String? cursor; }

// im_state.dart
abstract class ImState {}
class ImInitial extends ImState {}
class ImConnected extends ImState {}
class ImDisconnected extends ImState {}
class MessagesLoaded extends ImState { final List<Message> messages; final bool hasMore; }
class MessageSending extends ImState { final List<Message> messages; final String clientMessageId; }
class MessageSent extends ImState { final List<Message> messages; }
class MessageRead extends ImState { final String conversationId; final String lastReadMessageId; }
```

### 6.2 ConversationListCubit（会话列表）

```dart
class ConversationListCubit extends Cubit<ConversationListState> {
  Future<void> fetchConversations({String? cursor}) async { ... }
}

abstract class ConversationListState {}
class ConversationListInitial extends ConversationListState {}
class ConversationListLoadInProgress extends ConversationListState {}
class ConversationListLoadSuccess extends ConversationListState {
  final List<Conversation> conversations;
  final bool hasMore;
  final String? nextCursor;
}
class ConversationListLoadFailure extends ConversationListState { final String message; }
```

---

## 7. 需更新文档

### 7.1 API 文档

| 文档路径                    | 更新说明                                                    |
| --------------------------- | ----------------------------------------------------------- |
| `docs/api/im.md`            | 新增 IM 模块 API 文档（Socket.io 事件 + HTTP API）          |
| `docs/api/socket-events.md` | 更新 Socket.io 事件文档，记录 IM 相关事件 Payload           |

### 7.2 架构/流程文档

| 文档路径                      | 更新说明                                               |
| ----------------------------- | ------------------------------------------------------ |
| `docs/modules/im.md`          | 更新 IM 模块实现文档，补全 Flutter 侧的 BLoC、Widget、路由 |
| `docs/design-docs/flutter.md` | 在「Design-Only PR 验收 Checklist」中补充 IM 相关组件要求  |

### 7.3 PRD 注释

- 在 `prd/v1-mvp.md` 顶部添加链接注释：`设计资产参见 designs/issue-5/`

---

## 8. 验收 Checklist（供 Design-Only PR 使用）

- [x] 低保真/高保真原型图已产出（本文件 + mockup.html）
- [ ] 所有新增组件已在 `widgetbook/main.dart` 中注册 use-case
- [ ] PR 中附带 Golden File 截图或 Widgetbook 运行截图
- [ ] `pages/` 中未使用 `Color(0xFF...)`、`TextField`、`ElevatedButton` 等硬编码或原生 Material 组件替代 design-system
- [ ] 页面中的状态（加载、空态、错误、发送中、发送失败）在原型图中有明确定义
