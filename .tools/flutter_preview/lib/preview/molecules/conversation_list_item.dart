import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';
import '../_tokens.dart';

/// A list item showing a conversation with last message preview and unread count.
///
/// Use-cases:
/// - Default state (no unread)
/// - With unread badge
/// - Long message truncation
/// - Image message preview
class ConversationListItem extends StatelessWidget {
  /// Display name of the conversation partner
  final String nickname;

  /// Avatar URL of the conversation partner
  final String? avatarUrl;

  /// Preview text of the last message
  final String lastMessagePreview;

  /// Formatted time string (e.g. "10:30", "昨天")
  final String? lastMessageTime;

  /// Number of unread messages
  final int unreadCount;

  /// Callback when the item is tapped
  final VoidCallback? onTap;

  /// Creates a [ConversationListItem]
  const ConversationListItem({
    super.key,
    required this.nickname,
    this.avatarUrl,
    required this.lastMessagePreview,
    this.lastMessageTime,
    this.unreadCount = 0,
    this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return Semantics(
      identifier: 'conversation_list_item',
      label: '$nickname, 最新消息: $lastMessagePreview',
      button: true,
      container: true,
      child: GestureDetector(
        behavior: HitTestBehavior.opaque,
        onTapDown: (_) => onTap?.call(),
        child: Container(
          padding: const EdgeInsets.symmetric(
            horizontal: DesignTokens.spacingMd,
            vertical: DesignTokens.spacingSm,
          ),
          child: Row(
            children: [
              SwAvatar(
                imageUrl: avatarUrl,
                initials: nickname.isNotEmpty ? nickname.substring(0, 1) : '?',
                size: 56,
              ),
              const SizedBox(width: DesignTokens.spacingMd),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            nickname,
                            style: const TextStyle(
                              fontSize: DesignTokens.fontSizeBody,
                              fontWeight: FontWeight.w600,
                              color: DesignTokens.textPrimary,
                            ),
                            maxLines: 1,
                            overflow: TextOverflow.ellipsis,
                          ),
                        ),
                        if (lastMessageTime != null) ...[
                          const SizedBox(width: DesignTokens.spacingXs),
                          Text(
                            lastMessageTime!,
                            style: const TextStyle(
                              fontSize: DesignTokens.fontSizeCaption,
                              color: Tokens.timestampText,
                            ),
                          ),
                        ],
                      ],
                    ),
                    const SizedBox(height: 4),
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            lastMessagePreview,
                            style: TextStyle(
                              fontSize: Tokens.fontSizeSmall,
                              color: unreadCount > 0
                                  ? DesignTokens.textPrimary
                                  : Tokens.timestampText,
                              fontWeight: unreadCount > 0
                                  ? FontWeight.w500
                                  : FontWeight.normal,
                            ),
                            maxLines: 1,
                            overflow: TextOverflow.ellipsis,
                          ),
                        ),
                        if (unreadCount > 0) ...[
                          const SizedBox(width: DesignTokens.spacingXs),
                          Container(
                            constraints: const BoxConstraints(
                              minWidth: 20,
                              minHeight: 20,
                            ),
                            padding: const EdgeInsets.symmetric(horizontal: 6),
                            decoration: const BoxDecoration(
                              color: Tokens.unreadBadgeBackground,
                              borderRadius: BorderRadius.all(
                                Radius.circular(10),
                              ),
                            ),
                            alignment: Alignment.center,
                            child: Text(
                              unreadCount > 99 ? '99+' : '$unreadCount',
                              style: const TextStyle(
                                fontSize: 11,
                                fontWeight: FontWeight.w600,
                                color: Tokens.unreadBadgeText,
                              ),
                            ),
                          ),
                        ],
                      ],
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
