import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import '../../molecules/conversation_list_item.dart';

void main() {
  group('ConversationListItem Golden Tests', () {
    testWidgets('default', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: ConversationListItem(
              nickname: '阿杰',
              lastMessagePreview: '好的，周末见！',
              lastMessageTime: '昨天',
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ConversationListItem),
        matchesGoldenFile('conversation_list_item_default.png'),
      );
    });

    testWidgets('with unread', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: ConversationListItem(
              nickname: '小雨',
              lastMessagePreview: '今晚有空一起喝咖啡吗？',
              lastMessageTime: '10:30',
              unreadCount: 3,
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ConversationListItem),
        matchesGoldenFile('conversation_list_item_unread.png'),
      );
    });
  });
}
