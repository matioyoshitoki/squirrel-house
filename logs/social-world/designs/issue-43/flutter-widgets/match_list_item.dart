import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';

/// 匹配列表单项
///
/// 展示匹配用户的头像、昵称、匹配时间和聊天入口箭头。
class MatchListItem extends StatelessWidget {
  /// 用户昵称
  final String name;

  /// 头像 URL
  final String? avatarUrl;

  /// 匹配时间文本，如 "2小时前匹配"
  final String matchTime;

  /// 点击回调
  final VoidCallback? onTap;

  const MatchListItem({
    super.key,
    required this.name,
    this.avatarUrl,
    required this.matchTime,
    this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return Semantics(
      identifier: 'match_item_$name',
      label: '匹配用户 $name',
      button: true,
      child: GestureDetector(
        behavior: HitTestBehavior.opaque,
        onTapDown: (_) => onTap?.call(),
        child: Container(
          padding: const EdgeInsets.symmetric(
            horizontal: DesignTokens.spacingMd,
            vertical: DesignTokens.spacingSm,
          ),
          decoration: const BoxDecoration(
            border: Border(
              bottom: BorderSide(color: Color(0xFFEEEEEE)),
            ),
          ),
          child: Row(
            children: [
              // Avatar
              Container(
                width: 48,
                height: 48,
                decoration: BoxDecoration(
                  color: DesignTokens.primary.withOpacity(0.2),
                  shape: BoxShape.circle,
                ),
                clipBehavior: Clip.antiAlias,
                child: avatarUrl != null && avatarUrl!.isNotEmpty
                    ? Image.network(
                        avatarUrl!,
                        fit: BoxFit.cover,
                        errorBuilder: (_, __, ___) => _buildFallback(),
                      )
                    : _buildFallback(),
              ),
              const SizedBox(width: DesignTokens.spacingMd),
              // Info
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      name,
                      style: const TextStyle(
                        fontSize: DesignTokens.fontSizeBody,
                        fontWeight: FontWeight.w600,
                        color: DesignTokens.textPrimary,
                      ),
                    ),
                    const SizedBox(height: DesignTokens.spacingXxs),
                    Text(
                      matchTime,
                      style: const TextStyle(
                        fontSize: DesignTokens.fontSizeCaption,
                        color: DesignTokens.secondaryColor,
                      ),
                    ),
                  ],
                ),
              ),
              // Arrow
              const Icon(
                Icons.chevron_right,
                color: DesignTokens.secondaryColor,
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildFallback() {
    return const Center(
      child: Icon(
        Icons.person,
        size: 24,
        color: DesignTokens.primary,
      ),
    );
  }
}
