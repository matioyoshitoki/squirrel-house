import 'package:flutter/material.dart';
import 'package:widgetbook/widgetbook.dart';
import 'package:social_world_design_system/social_world_design_system.dart';
import '../molecules/chat_bubble.dart';
import '../molecules/chat_input_bar.dart';
import '../molecules/conversation_list_item.dart';
import '../pages/chat_page.dart';
import '../pages/conversation_list_page.dart';

void main() {
  runApp(const WidgetbookApp());
}

class WidgetbookApp extends StatelessWidget {
  const WidgetbookApp({super.key});

  @override
  Widget build(BuildContext context) {
    return Widgetbook.material(
      directories: [
        WidgetbookFolder(
          name: 'Molecules',
          children: [
            WidgetbookComponent(
              name: 'ChatBubble',
              useCases: [
                WidgetbookUseCase(
                  name: 'Self Text',
                  builder: (context) => const Center(
                    child: ChatBubble(
                      content: '嗨！很高兴认识你 😊',
                      type: ChatBubbleType.selfText,
                      status: MessageStatus.read,
                      timestamp: '10:26',
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Other Text',
                  builder: (context) => const Center(
                    child: ChatBubble(
                      content: '你好呀 👋',
                      type: ChatBubbleType.otherText,
                      timestamp: '10:25',
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'System Message',
                  builder: (context) => const Center(
                    child: ChatBubble(
                      content: '你们已匹配成功，开始聊天吧',
                      type: ChatBubbleType.system,
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Sending',
                  builder: (context) => const Center(
                    child: ChatBubble(
                      content: '正在发送的消息...',
                      type: ChatBubbleType.selfText,
                      status: MessageStatus.sending,
                      timestamp: '10:28',
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Failed',
                  builder: (context) => Center(
                    child: ChatBubble(
                      content: '发送失败的消息',
                      type: ChatBubbleType.selfText,
                      status: MessageStatus.failed,
                      timestamp: '10:27',
                      onRetry: () {},
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Long Self Text',
                  builder: (context) => const Center(
                    child: ChatBubble(
                      content: '真的吗！太巧了，我也是！下次可以一起去拍照，我知道一个很棒的拍摄地点 📸',
                      type: ChatBubbleType.selfText,
                      status: MessageStatus.read,
                      timestamp: '10:30',
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Self Image',
                  builder: (context) => const Center(
                    child: ChatBubble(
                      content: 'https://example.com/image.jpg',
                      type: ChatBubbleType.selfImage,
                      status: MessageStatus.sent,
                      timestamp: '10:35',
                    ),
                  ),
                ),
              ],
            ),
            WidgetbookComponent(
              name: 'ChatInputBar',
              useCases: [
                WidgetbookUseCase(
                  name: 'Empty',
                  builder: (context) => const Column(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      ChatInputBar(
                        currentLength: 0,
                      ),
                    ],
                  ),
                ),
                WidgetbookUseCase(
                  name: 'With Text',
                  builder: (context) => const Column(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      ChatInputBar(
                        currentLength: 12,
                      ),
                    ],
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Sending',
                  builder: (context) => const Column(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      ChatInputBar(
                        currentLength: 8,
                        isSending: true,
                      ),
                    ],
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Over Limit',
                  builder: (context) => const Column(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      ChatInputBar(
                        currentLength: 520,
                      ),
                    ],
                  ),
                ),
              ],
            ),
            WidgetbookComponent(
              name: 'ConversationListItem',
              useCases: [
                WidgetbookUseCase(
                  name: 'Default',
                  builder: (context) => const Center(
                    child: ConversationListItem(
                      nickname: '阿杰',
                      lastMessagePreview: '好的，周末见！',
                      lastMessageTime: '昨天',
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'With Unread',
                  builder: (context) => const Center(
                    child: ConversationListItem(
                      nickname: '小雨',
                      lastMessagePreview: '今晚有空一起喝咖啡吗？☕️',
                      lastMessageTime: '10:30',
                      unreadCount: 3,
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Long Message',
                  builder: (context) => const Center(
                    child: ConversationListItem(
                      nickname: 'Alice',
                      lastMessagePreview:
                          '这是一个非常长的消息预览，需要被截断显示...',
                      lastMessageTime: '周一',
                      unreadCount: 1,
                    ),
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Image Preview',
                  builder: (context) => const Center(
                    child: ConversationListItem(
                      nickname: 'Bob',
                      lastMessagePreview: '[图片]',
                      lastMessageTime: '3天前',
                    ),
                  ),
                ),
              ],
            ),
          ],
        ),
        WidgetbookFolder(
          name: 'Pages',
          children: [
            WidgetbookComponent(
              name: 'ChatPage Skeleton',
              useCases: [
                WidgetbookUseCase(
                  name: 'Normal',
                  builder: (context) => ChatPage(
                    partnerNickname: '小雨',
                    partnerAvatarUrl: 'https://i.pravatar.cc/150?img=5',
                    messages: [
                      const MockMessage(
                        content: '你们已匹配成功，开始聊天吧',
                        type: ChatBubbleType.system,
                      ),
                      const MockMessage(
                        content: '你好呀 👋',
                        type: ChatBubbleType.otherText,
                        timestamp: '10:25',
                      ),
                      const MockMessage(
                        content: '嗨！很高兴认识你 😊',
                        type: ChatBubbleType.selfText,
                        status: MessageStatus.read,
                        timestamp: '10:26',
                      ),
                      const MockMessage(
                        content: '看你资料说喜欢摄影，我也超爱！',
                        type: ChatBubbleType.otherText,
                        timestamp: '10:27',
                      ),
                      const MockMessage(
                        content: '真的吗！太巧了，我也是！',
                        type: ChatBubbleType.selfText,
                        status: MessageStatus.sending,
                        timestamp: '10:28',
                      ),
                    ],
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Empty History',
                  builder: (context) => const ChatPage(
                    partnerNickname: '阿杰',
                    messages: [],
                  ),
                ),
                WidgetbookUseCase(
                  name: 'With Failed Message',
                  builder: (context) => ChatPage(
                    partnerNickname: '阿杰',
                    messages: [
                      const MockMessage(
                        content: '你们已匹配成功，开始聊天吧',
                        type: ChatBubbleType.system,
                      ),
                      const MockMessage(
                        content: '最近有看什么好看的电影吗？',
                        type: ChatBubbleType.otherText,
                        timestamp: '昨天 15:20',
                      ),
                      MockMessage(
                        content: '看了《奥本海默》，强烈推荐！',
                        type: ChatBubbleType.selfText,
                        status: MessageStatus.failed,
                        timestamp: '昨天 15:22',
                      ),
                      const MockMessage(
                        content: '看了《奥本海默》，强烈推荐！',
                        type: ChatBubbleType.selfText,
                        status: MessageStatus.read,
                        timestamp: '昨天 15:25',
                      ),
                    ],
                  ),
                ),
              ],
            ),
            WidgetbookComponent(
              name: 'ConversationListPage Skeleton',
              useCases: [
                WidgetbookUseCase(
                  name: 'Normal',
                  builder: (context) => const ConversationListPage(
                    conversations: [
                      MockConversation(
                        nickname: '小雨',
                        lastMessagePreview: '今晚有空一起喝咖啡吗？☕️',
                        lastMessageTime: '10:30',
                        unreadCount: 3,
                      ),
                      MockConversation(
                        nickname: '阿杰',
                        lastMessagePreview: '好的，周末见！',
                        lastMessageTime: '昨天',
                      ),
                      MockConversation(
                        nickname: 'Alice',
                        lastMessagePreview: '[图片]',
                        lastMessageTime: '周一',
                        unreadCount: 1,
                      ),
                      MockConversation(
                        nickname: 'Bob',
                        lastMessagePreview: '哈哈哈太有趣了',
                        lastMessageTime: '3天前',
                      ),
                    ],
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Loading',
                  builder: (context) => const ConversationListPage(
                    isLoading: true,
                  ),
                ),
                WidgetbookUseCase(
                  name: 'Empty',
                  builder: (context) => const ConversationListPage(),
                ),
                WidgetbookUseCase(
                  name: 'Error',
                  builder: (context) => const ConversationListPage(
                    errorMessage: '网络连接失败，请检查网络后重试',
                  ),
                ),
              ],
            ),
          ],
        ),
      ],
      addons: [],
    );
  }
}
