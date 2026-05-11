import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';
import '../_tokens.dart';

/// Input bar for chat page combining text field, image picker, and send button.
///
/// Use-cases:
/// - Empty input (send button disabled)
/// - Text entered (send button enabled)
/// - Sending state (send button shows spinner)
class ChatInputBar extends StatelessWidget {
  /// Text controller for the input field
  final TextEditingController? controller;

  /// Callback when send button is tapped
  final VoidCallback? onSend;

  /// Callback when image picker button is tapped
  final VoidCallback? onPickImage;

  /// Whether a message is being sent
  final bool isSending;

  /// Hint text shown when empty
  final String hintText;

  /// Maximum length of input text
  final int maxLength;

  /// Current length of input text (for validation display)
  final int currentLength;

  /// Creates a [ChatInputBar]
  const ChatInputBar({
    super.key,
    this.controller,
    this.onSend,
    this.onPickImage,
    this.isSending = false,
    this.hintText = '输入消息...',
    this.maxLength = 500,
    this.currentLength = 0,
  });

  bool get _isOverLimit => currentLength > maxLength;

  bool get _canSend =>
      !isSending &&
      currentLength > 0 &&
      !_isOverLimit;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.only(
        left: DesignTokens.spacingMd,
        right: DesignTokens.spacingMd,
        top: DesignTokens.spacingSm,
        bottom: DesignTokens.spacingSm +
            MediaQuery.of(context).padding.bottom,
      ),
      decoration: const BoxDecoration(
        color: Tokens.inputBarBackground,
        border: Border(
          top: BorderSide(color: Tokens.divider),
        ),
      ),
      child: SafeArea(
        top: false,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Row(
              children: [
                // Image picker button
                Semantics(
                  identifier: 'chat_image_picker_button',
                  label: '选择图片',
                  button: true,
                  container: true,
                  child: GestureDetector(
                    behavior: HitTestBehavior.opaque,
                    onTapDown: (_) => onPickImage?.call(),
                    child: Container(
                      width: 40,
                      height: 40,
                      alignment: Alignment.center,
                      child: const Icon(
                        Icons.image_outlined,
                        color: DesignTokens.placeholderIcon,
                        size: 24,
                      ),
                    ),
                  ),
                ),
                const SizedBox(width: DesignTokens.spacingSm),
                // Text field
                Expanded(
                  child: Semantics(
                    identifier: 'chat_input_field',
                    label: '消息输入框',
                    child: TextField(
                      controller: controller,
                      maxLines: null,
                      maxLength: null,
                      textInputAction: TextInputAction.send,
                      onSubmitted: (_) => _canSend ? onSend?.call() : null,
                      decoration: InputDecoration(
                        hintText: hintText,
                        hintStyle: const TextStyle(
                          color: DesignTokens.placeholderIcon,
                        ),
                        filled: true,
                        fillColor: Tokens.bubbleOtherBackground,
                        contentPadding: const EdgeInsets.symmetric(
                          horizontal: DesignTokens.spacingMd,
                          vertical: 10,
                        ),
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(20),
                          borderSide: BorderSide.none,
                        ),
                        enabledBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(20),
                          borderSide: BorderSide.none,
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(20),
                          borderSide: const BorderSide(
                            color: DesignTokens.primary,
                            width: 1,
                          ),
                        ),
                        counterText: '',
                      ),
                    ),
                  ),
                ),
                const SizedBox(width: DesignTokens.spacingSm),
                // Send button
                Semantics(
                  identifier: 'chat_send_button',
                  label: '发送',
                  button: true,
                  container: true,
                  child: GestureDetector(
                    behavior: HitTestBehavior.opaque,
                    onTapDown: (_) => _canSend ? onSend?.call() : null,
                    child: Container(
                      width: 40,
                      height: 40,
                      decoration: BoxDecoration(
                        color: _canSend
                            ? DesignTokens.primary
                            : DesignTokens.primary.withOpacity(0.3),
                        shape: BoxShape.circle,
                      ),
                      alignment: Alignment.center,
                      child: isSending
                          ? const SizedBox(
                              width: 20,
                              height: 20,
                              child: CircularProgressIndicator(
                                strokeWidth: 2,
                                color: DesignTokens.textOnDarkPrimary,
                              ),
                            )
                          : const Icon(
                              Icons.send,
                              color: DesignTokens.textOnDarkPrimary,
                              size: 20,
                            ),
                    ),
                  ),
                ),
              ],
            ),
            if (_isOverLimit) ...[
              const SizedBox(height: 4),
              Text(
                '已超过 $maxLength 字限制',
                style: const TextStyle(
                  fontSize: DesignTokens.fontSizeCaption,
                  color: DesignTokens.error,
                ),
              ),
            ],
          ],
        ),
      ),
    );
  }
}
