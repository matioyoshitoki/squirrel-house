import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';
import '../_tokens.dart';

/// Message bubble types
enum ChatBubbleType {
  /// Text message sent by current user
  selfText,

  /// Text message received from other user
  otherText,

  /// Image message sent by current user
  selfImage,

  /// Image message received from other user
  otherImage,

  /// System notification message (centered, e.g. "You matched")
  system,
}

/// Message send status
enum MessageStatus {
  /// Message is being sent
  sending,

  /// Message has been sent to server
  sent,

  /// Message has been delivered to recipient
  delivered,

  /// Recipient has read the message
  read,

  /// Message failed to send
  failed,
}

/// A chat bubble molecule supporting text/image/system message types.
///
/// Use-cases:
/// - Self text message (right-aligned, primary background)
/// - Other text message (left-aligned, surface background)
/// - System message (centered, subtle background)
/// - Image messages (rounded rectangle with optional loading state)
/// - Failed state (red error icon + retry)
class ChatBubble extends StatelessWidget {
  /// The message content (text or image url)
  final String content;

  /// Bubble type determining alignment and styling
  final ChatBubbleType type;

  /// Optional send status shown for self messages
  final MessageStatus? status;

  /// Optional timestamp shown at bottom of bubble
  final String? timestamp;

  /// Callback when failed message is tapped to retry
  final VoidCallback? onRetry;

  /// Maximum width ratio of screen width
  final double maxWidthRatio;

  /// Creates a [ChatBubble]
  const ChatBubble({
    super.key,
    required this.content,
    required this.type,
    this.status,
    this.timestamp,
    this.onRetry,
    this.maxWidthRatio = 0.7,
  });

  bool get isSelf =>
      type == ChatBubbleType.selfText || type == ChatBubbleType.selfImage;

  bool get isSystem => type == ChatBubbleType.system;

  @override
  Widget build(BuildContext context) {
    if (isSystem) {
      return _buildSystemBubble();
    }

    final screenWidth = MediaQuery.of(context).size.width;
    final maxBubbleWidth = screenWidth * maxWidthRatio;

    return Row(
      mainAxisAlignment:
          isSelf ? MainAxisAlignment.end : MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.end,
      children: [
        if (!isSelf && status == MessageStatus.failed) ...[
          _buildFailIcon(),
          const SizedBox(width: DesignTokens.spacingXxs),
        ],
        ConstrainedBox(
          constraints: BoxConstraints(maxWidth: maxBubbleWidth),
          child: Container(
            padding: type == ChatBubbleType.selfImage ||
                    type == ChatBubbleType.otherImage
                ? const EdgeInsets.all(DesignTokens.spacingXxs)
                : const EdgeInsets.symmetric(
                    horizontal: DesignTokens.spacingMd,
                    vertical: DesignTokens.spacingSm,
                  ),
            decoration: BoxDecoration(
              color: isSelf
                  ? Tokens.bubbleSelfBackground
                  : Tokens.bubbleOtherBackground,
              borderRadius: BorderRadius.only(
                topLeft: const Radius.circular(DesignTokens.radiusMd),
                topRight: const Radius.circular(DesignTokens.radiusMd),
                bottomLeft: Radius.circular(isSelf ? DesignTokens.radiusMd : 4),
                bottomRight:
                    Radius.circular(isSelf ? 4 : DesignTokens.radiusMd),
              ),
            ),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.end,
              mainAxisSize: MainAxisSize.min,
              children: [
                _buildContent(),
                const SizedBox(height: 4),
                _buildMeta(),
              ],
            ),
          ),
        ),
        if (isSelf && status == MessageStatus.failed) ...[
          const SizedBox(width: DesignTokens.spacingXxs),
          _buildFailIcon(),
        ],
      ],
    );
  }

  Widget _buildContent() {
    final isImage =
        type == ChatBubbleType.selfImage || type == ChatBubbleType.otherImage;

    if (isImage) {
      return ClipRRect(
        borderRadius: BorderRadius.circular(DesignTokens.radiusMd - 2),
        child: Container(
          width: 200,
          height: 200,
          color: DesignTokens.placeholderGradientStart,
          child: const Center(
            child: Icon(
              Icons.image,
              size: 48,
              color: DesignTokens.placeholderIcon,
            ),
          ),
        ),
      );
    }

    return Text(
      content,
      style: TextStyle(
        fontSize: DesignTokens.fontSizeBody,
        color: isSelf ? Tokens.bubbleSelfText : Tokens.bubbleOtherText,
        height: 1.4,
      ),
    );
  }

  Widget _buildMeta() {
    final children = <Widget>[];

    if (timestamp != null) {
      children.add(
        Text(
          timestamp!,
          style: TextStyle(
            fontSize: 10,
            color: isSelf
                ? Tokens.bubbleSelfText.withOpacity(0.7)
                : Tokens.timestampText,
          ),
        ),
      );
    }

    if (isSelf && status != null) {
      children.add(const SizedBox(width: 4));
      children.add(_buildStatusIndicator());
    }

    return Row(
      mainAxisSize: MainAxisSize.min,
      children: children,
    );
  }

  Widget _buildStatusIndicator() {
    switch (status) {
      case MessageStatus.sending:
        return SizedBox(
          width: 12,
          height: 12,
          child: CircularProgressIndicator(
            strokeWidth: 2,
            color: Tokens.bubbleSelfText.withOpacity(0.7),
          ),
        );
      case MessageStatus.read:
        return const Text(
          '已读',
          style: TextStyle(
            fontSize: 10,
            color: Tokens.readReceiptText,
            fontWeight: FontWeight.w500,
          ),
        );
      case MessageStatus.failed:
        return const SizedBox.shrink();
      default:
        return const SizedBox.shrink();
    }
  }

  Widget _buildFailIcon() {
    return Semantics(
      identifier: 'message_retry_button',
      label: '重发消息',
      button: true,
      container: true,
      child: GestureDetector(
        behavior: HitTestBehavior.opaque,
        onTapDown: (_) => onRetry?.call(),
        child: const Icon(
          Icons.error_outline,
          size: 20,
          color: DesignTokens.error,
        ),
      ),
    );
  }

  Widget _buildSystemBubble() {
    return Center(
      child: Container(
        padding: const EdgeInsets.symmetric(
          horizontal: DesignTokens.spacingMd,
          vertical: 6,
        ),
        decoration: BoxDecoration(
          color: DesignTokens.textPrimary.withOpacity(0.06),
          borderRadius: BorderRadius.circular(DesignTokens.radiusSm),
        ),
        child: Text(
          content,
          style: const TextStyle(
            fontSize: DesignTokens.fontSizeCaption,
            color: Tokens.systemMessageText,
          ),
        ),
      ),
    );
  }
}
