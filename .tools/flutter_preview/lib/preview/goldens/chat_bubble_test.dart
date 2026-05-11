import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import '../../molecules/chat_bubble.dart';

void main() {
  group('ChatBubble Golden Tests', () {
    testWidgets('self text read', (tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: Center(
              child: ChatBubble(
                content: '嗨！很高兴认识你',
                type: ChatBubbleType.selfText,
                status: MessageStatus.read,
                timestamp: '10:26',
              ),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ChatBubble),
        matchesGoldenFile('chat_bubble_self_text_read.png'),
      );
    });

    testWidgets('other text', (tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: Center(
              child: ChatBubble(
                content: '你好呀',
                type: ChatBubbleType.otherText,
                timestamp: '10:25',
              ),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ChatBubble),
        matchesGoldenFile('chat_bubble_other_text.png'),
      );
    });

    testWidgets('system message', (tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: Center(
              child: ChatBubble(
                content: '你们已匹配成功，开始聊天吧',
                type: ChatBubbleType.system,
              ),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ChatBubble),
        matchesGoldenFile('chat_bubble_system.png'),
      );
    });

    testWidgets('failed', (tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: Center(
              child: ChatBubble(
                content: '发送失败',
                type: ChatBubbleType.selfText,
                status: MessageStatus.failed,
                timestamp: '10:27',
                onRetry: () {},
              ),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ChatBubble),
        matchesGoldenFile('chat_bubble_failed.png'),
      );
    });
  });
}
