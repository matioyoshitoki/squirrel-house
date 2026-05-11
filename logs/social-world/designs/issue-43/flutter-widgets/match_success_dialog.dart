import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';

/// 配对成功全屏弹窗
///
/// 双向喜欢时触发，展示双方头像组合和 CTA 按钮。
class MatchSuccessDialog extends StatelessWidget {
  /// 当前用户头像 URL
  final String currentUserAvatar;

  /// 匹配用户头像 URL
  final String matchedUserAvatar;

  /// 匹配用户昵称
  final String matchedUserName;

  /// 点击「开始聊天」回调
  final VoidCallback? onChat;

  /// 点击「继续滑动」回调
  final VoidCallback? onContinue;

  const MatchSuccessDialog({
    super.key,
    required this.currentUserAvatar,
    required this.matchedUserAvatar,
    required this.matchedUserName,
    this.onChat,
    this.onContinue,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black.withOpacity(0.6),
      body: SafeArea(
        child: Center(
          child: AnimatedContainer(
            duration: const Duration(milliseconds: 300),
            curve: Curves.easeOutBack,
            margin: const EdgeInsets.all(DesignTokens.spacingMd),
            padding: const EdgeInsets.all(DesignTokens.spacingXl),
            decoration: BoxDecoration(
              color: DesignTokens.background,
              borderRadius: BorderRadius.circular(DesignTokens.radiusLg),
            ),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Text(
                  '配对成功！',
                  style: TextStyle(
                    fontSize: DesignTokens.fontSizeHeadline,
                    fontWeight: FontWeight.bold,
                    color: DesignTokens.primary,
                  ),
                ),
                const SizedBox(height: DesignTokens.spacingMd),
                Text(
                  '你和 $matchedUserName 互相喜欢了',
                  textAlign: TextAlign.center,
                  style: const TextStyle(
                    fontSize: DesignTokens.fontSizeBody,
                    color: DesignTokens.secondaryColor,
                  ),
                ),
                const SizedBox(height: DesignTokens.spacingXl),
                // Avatar combination
                Stack(
                  alignment: Alignment.center,
                  children: [
                    Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        _buildAvatar(currentUserAvatar, DesignTokens.spacingLg),
                        const SizedBox(width: DesignTokens.spacingMd),
                        _buildAvatar(matchedUserAvatar, 0),
                      ],
                    ),
                    Container(
                      width: 40,
                      height: 40,
                      decoration: BoxDecoration(
                        color: const Color(0xFF2E7D32),
                        shape: BoxShape.circle,
                        border: Border.all(color: Colors.white, width: 3),
                      ),
                      child: const Icon(
                        Icons.favorite,
                        color: Colors.white,
                        size: 20,
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: DesignTokens.spacingXl),
                SizedBox(
                  width: double.infinity,
                  child: Semantics(
                    identifier: 'match_chat_button',
                    label: '开始聊天',
                    button: true,
                    child: GestureDetector(
                      behavior: HitTestBehavior.opaque,
                      onTapDown: (_) => onChat?.call(),
                      child: Container(
                        alignment: Alignment.center,
                        padding: const EdgeInsets.symmetric(
                          vertical: DesignTokens.spacingMd,
                        ),
                        decoration: BoxDecoration(
                          color: DesignTokens.primary,
                          borderRadius: BorderRadius.circular(
                            DesignTokens.radiusMd,
                          ),
                        ),
                        child: const Text(
                          '开始聊天',
                          style: TextStyle(
                            color: Colors.white,
                            fontSize: DesignTokens.fontSizeBody,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: DesignTokens.spacingMd),
                SizedBox(
                  width: double.infinity,
                  child: Semantics(
                    identifier: 'match_continue_button',
                    label: '继续滑动',
                    button: true,
                    child: GestureDetector(
                      behavior: HitTestBehavior.opaque,
                      onTapDown: (_) => onContinue?.call(),
                      child: Container(
                        alignment: Alignment.center,
                        padding: const EdgeInsets.symmetric(
                          vertical: DesignTokens.spacingMd,
                        ),
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(
                            DesignTokens.radiusMd,
                          ),
                          border: Border.all(color: DesignTokens.primary),
                        ),
                        child: const Text(
                          '继续滑动',
                          style: TextStyle(
                            color: DesignTokens.primary,
                            fontSize: DesignTokens.fontSizeBody,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildAvatar(String url, double rightPadding) {
    return Padding(
      padding: EdgeInsets.only(right: rightPadding),
      child: Container(
        width: 80,
        height: 80,
        decoration: BoxDecoration(
          color: DesignTokens.primary.withOpacity(0.2),
          shape: BoxShape.circle,
          border: Border.all(color: Colors.white, width: 3),
        ),
        clipBehavior: Clip.antiAlias,
        child: url.isNotEmpty
            ? Image.network(url, fit: BoxFit.cover)
            : const Icon(
                Icons.person,
                size: 40,
                color: DesignTokens.primary,
              ),
      ),
    );
  }
}
