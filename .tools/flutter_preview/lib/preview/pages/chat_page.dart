import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';
import '../molecules/chat_bubble.dart';
import '../molecules/chat_input_bar.dart';
import '../_tokens.dart';

/// Pure UI skeleton of ChatPage for Widgetbook review.
/// Contains no BLoC, no repository, no network calls.
class ChatPage extends StatefulWidget {
  /// Partner display name shown in app bar
  final String partnerNickname;

  /// Partner avatar URL
  final String? partnerAvatarUrl;

  /// Mock messages for preview
  final List<MockMessage> messages;

  /// Creates a [ChatPage] skeleton
  const ChatPage({
    super.key,
    required this.partnerNickname,
    this.partnerAvatarUrl,
    this.messages = const [],
  });

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
  final _controller = TextEditingController();
  final _scrollController = ScrollController();
  int _currentLength = 0;

  @override
  void initState() {
    super.initState();
    _controller.addListener(() {
      setState(() {
        _currentLength = _controller.text.length;
      });
    });
  }

  @override
  void dispose() {
    _controller.dispose();
    _scrollController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      appBar: AppBar(
        backgroundColor: DesignTokens.background,
        elevation: 0,
        leading: Semantics(
          identifier: 'chat_back_button',
          label: '返回',
          button: true,
          container: true,
          child: GestureDetector(
            behavior: HitTestBehavior.opaque,
            onTapDown: (_) => Navigator.of(context).maybePop(),
            child: const Icon(
              Icons.arrow_back_ios,
              color: DesignTokens.textPrimary,
              size: 20,
            ),
          ),
        ),
        title: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            SwAvatar(
              imageUrl: widget.partnerAvatarUrl,
              initials: widget.partnerNickname.isNotEmpty
                  ? widget.partnerNickname.substring(0, 1)
                  : '?',
              size: 36,
            ),
            const SizedBox(width: DesignTokens.spacingSm),
            Text(
              widget.partnerNickname,
              style: const TextStyle(
                fontSize: DesignTokens.fontSizeBody,
                fontWeight: FontWeight.w600,
                color: DesignTokens.textPrimary,
              ),
            ),
          ],
        ),
        centerTitle: true,
      ),
      body: Column(
        children: [
          Expanded(
            child: widget.messages.isEmpty
                ? _buildEmptyState()
                : ListView.separated(
                    controller: _scrollController,
                    padding: const EdgeInsets.all(DesignTokens.spacingMd),
                    reverse: false,
                    itemCount: widget.messages.length,
                    separatorBuilder: (_, __) =>
                        const SizedBox(height: DesignTokens.spacingMd),
                    itemBuilder: (context, index) {
                      final msg = widget.messages[index];
                      return ChatBubble(
                        content: msg.content,
                        type: msg.type,
                        status: msg.status,
                        timestamp: msg.timestamp,
                        onRetry: msg.status == MessageStatus.failed
                            ? () {}
                            : null,
                      );
                    },
                  ),
          ),
          ChatInputBar(
            controller: _controller,
            currentLength: _currentLength,
            onSend: () {
              // Pure UI skeleton: no actual send logic
              _controller.clear();
            },
            onPickImage: () {
              // Pure UI skeleton: no actual image pick
            },
          ),
        ],
      ),
    );
  }

  Widget _buildEmptyState() {
    return Center(
      child: ChatBubble(
        content: '你们已匹配成功，开始聊天吧',
        type: ChatBubbleType.system,
      ),
    );
  }
}

/// Mock message data for UI preview
class MockMessage {
  final String content;
  final ChatBubbleType type;
  final MessageStatus? status;
  final String? timestamp;

  const MockMessage({
    required this.content,
    required this.type,
    this.status,
    this.timestamp,
  });
}
