import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';
import '../molecules/conversation_list_item.dart';
import '../_tokens.dart';

/// Pure UI skeleton of ConversationListPage for Widgetbook review.
/// Contains no BLoC, no repository, no network calls.
class ConversationListPage extends StatelessWidget {
  /// Mock conversations for preview
  final List<MockConversation> conversations;

  /// Whether data is loading
  final bool isLoading;

  /// Error message (null if no error)
  final String? errorMessage;

  /// Creates a [ConversationListPage] skeleton
  const ConversationListPage({
    super.key,
    this.conversations = const [],
    this.isLoading = false,
    this.errorMessage,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      appBar: AppBar(
        backgroundColor: DesignTokens.background,
        elevation: 0,
        title: const Text(
          '消息',
          style: TextStyle(
            color: DesignTokens.textPrimary,
            fontSize: DesignTokens.fontSizeTitle,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: _buildBody(),
    );
  }

  Widget _buildBody() {
    if (isLoading) {
      return const Center(
        child: CircularProgressIndicator(color: DesignTokens.primary),
      );
    }

    if (errorMessage != null) {
      return Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Icon(
              Icons.error_outline,
              size: 48,
              color: DesignTokens.error,
            ),
            const SizedBox(height: DesignTokens.spacingMd),
            Text(
              errorMessage!,
              style: const TextStyle(color: DesignTokens.error),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: DesignTokens.spacingMd),
            Semantics(
              identifier: 'conversation_list_retry_button',
              label: '重试',
              button: true,
              container: true,
              child: GestureDetector(
                behavior: HitTestBehavior.opaque,
                onTapDown: (_) {},
                child: Container(
                  padding: const EdgeInsets.symmetric(
                    horizontal: DesignTokens.spacingLg,
                    vertical: DesignTokens.spacingMd,
                  ),
                  decoration: BoxDecoration(
                    border: Border.all(color: DesignTokens.primary),
                    borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                  ),
                  child: const Text(
                    '重试',
                    style: TextStyle(color: DesignTokens.primary),
                  ),
                ),
              ),
            ),
          ],
        ),
      );
    }

    if (conversations.isEmpty) {
      return Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(
              Icons.chat_bubble_outline,
              size: 80,
              color: DesignTokens.textPrimary.withOpacity(0.3),
            ),
            const SizedBox(height: DesignTokens.spacingMd),
            const Text(
              '还没有匹配成功的好友',
              style: TextStyle(
                fontSize: DesignTokens.fontSizeBody,
                fontWeight: FontWeight.w600,
                color: DesignTokens.textPrimary,
              ),
            ),
            const SizedBox(height: DesignTokens.spacingSm),
            Text(
              '继续滑动，发现感兴趣的人吧',
              style: TextStyle(
                fontSize: Tokens.fontSizeSmall,
                color: Tokens.timestampText,
              ),
            ),
            const SizedBox(height: DesignTokens.spacingLg),
            Semantics(
              identifier: 'go_discover_button',
              label: '去首页发现',
              button: true,
              container: true,
              child: GestureDetector(
                behavior: HitTestBehavior.opaque,
                onTapDown: (_) {},
                child: Container(
                  padding: const EdgeInsets.symmetric(
                    horizontal: DesignTokens.spacingLg,
                    vertical: DesignTokens.spacingMd,
                  ),
                  decoration: BoxDecoration(
                    color: DesignTokens.primary,
                    borderRadius: BorderRadius.circular(24),
                  ),
                  child: const Text(
                    '去首页发现',
                    style: TextStyle(
                      color: DesignTokens.textOnDarkPrimary,
                      fontWeight: FontWeight.w600,
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      );
    }

    return ListView.separated(
      physics: const AlwaysScrollableScrollPhysics(),
      itemCount: conversations.length,
      separatorBuilder: (_, __) => const Divider(
        height: 1,
        color: Tokens.divider,
        indent: DesignTokens.spacingMd + 56 + DesignTokens.spacingMd,
      ),
      itemBuilder: (context, index) {
        final conv = conversations[index];
        return ConversationListItem(
          nickname: conv.nickname,
          avatarUrl: conv.avatarUrl,
          lastMessagePreview: conv.lastMessagePreview,
          lastMessageTime: conv.lastMessageTime,
          unreadCount: conv.unreadCount,
          onTap: () {},
        );
      },
    );
  }
}

/// Mock conversation data for UI preview
class MockConversation {
  final String nickname;
  final String? avatarUrl;
  final String lastMessagePreview;
  final String? lastMessageTime;
  final int unreadCount;

  const MockConversation({
    required this.nickname,
    this.avatarUrl,
    required this.lastMessagePreview,
    this.lastMessageTime,
    this.unreadCount = 0,
  });
}
