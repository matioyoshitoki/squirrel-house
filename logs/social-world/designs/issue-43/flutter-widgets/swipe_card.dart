import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';

/// 可拖动/可点击的推荐用户卡片
///
/// 展示用户头像、昵称、年龄、距离和个性签名。
/// 支持拖动旋转和 Like/Pass badge 透明度渐变。
class SwipeCard extends StatelessWidget {
  /// 用户昵称
  final String name;

  /// 年龄
  final int age;

  /// 距离文本，如 "2.5km"
  final String distance;

  /// 个性签名
  final String bio;

  /// 头像/照片 URL
  final String? imageUrl;

  /// 拖动旋转角度（弧度）
  final double rotation;

  /// Like badge 透明度 [0, 1]
  final double likeOpacity;

  /// Pass badge 透明度 [0, 1]
  final double passOpacity;

  const SwipeCard({
    super.key,
    required this.name,
    required this.age,
    required this.distance,
    required this.bio,
    this.imageUrl,
    this.rotation = 0,
    this.likeOpacity = 0,
    this.passOpacity = 0,
  });

  @override
  Widget build(BuildContext context) {
    return Transform.rotate(
      angle: rotation,
      child: Container(
        margin: const EdgeInsets.symmetric(
          horizontal: DesignTokens.spacingMd,
        ),
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(DesignTokens.radiusLg),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withOpacity(0.08),
              blurRadius: 20,
              offset: const Offset(0, 4),
            ),
          ],
        ),
        clipBehavior: Clip.antiAlias,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            // Image area
            AspectRatio(
              aspectRatio: 1.0,
              child: Stack(
                fit: StackFit.expand,
                children: [
                  Container(
                    decoration: const BoxDecoration(
                      gradient: LinearGradient(
                        begin: Alignment.topLeft,
                        end: Alignment.bottomRight,
                        colors: [Color(0xFFE8E0EB), Color(0xFFD4C8DB)],
                      ),
                    ),
                    child: imageUrl != null && imageUrl!.isNotEmpty
                        ? Image.network(
                            imageUrl!,
                            fit: BoxFit.cover,
                            errorBuilder: (_, __, ___) => _buildPlaceholder(),
                          )
                        : _buildPlaceholder(),
                  ),
                  // Like badge
                  Positioned(
                    top: DesignTokens.spacingLg,
                    left: DesignTokens.spacingLg,
                    child: Opacity(
                      opacity: likeOpacity,
                      child: _buildBadge('LIKE', const Color(0xFF2E7D32)),
                    ),
                  ),
                  // Pass badge
                  Positioned(
                    top: DesignTokens.spacingLg,
                    right: DesignTokens.spacingLg,
                    child: Opacity(
                      opacity: passOpacity,
                      child: _buildBadge('PASS', DesignTokens.error),
                    ),
                  ),
                ],
              ),
            ),
            // Info area
            Padding(
              padding: const EdgeInsets.all(DesignTokens.spacingMd),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    '$name, $age',
                    style: const TextStyle(
                      fontSize: DesignTokens.fontSizeTitle,
                      fontWeight: FontWeight.bold,
                      color: DesignTokens.textPrimary,
                    ),
                  ),
                  const SizedBox(height: DesignTokens.spacingXs),
                  Row(
                    children: [
                      const Icon(
                        Icons.location_on,
                        size: 14,
                        color: DesignTokens.secondaryColor,
                      ),
                      const SizedBox(width: DesignTokens.spacingXxs),
                      Text(
                        distance,
                        style: const TextStyle(
                          fontSize: 14,
                          color: DesignTokens.secondaryColor,
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(height: DesignTokens.spacingXs),
                  Text(
                    bio,
                    maxLines: 2,
                    overflow: TextOverflow.ellipsis,
                    style: const TextStyle(
                      fontSize: 14,
                      color: DesignTokens.secondaryColor,
                      height: 1.4,
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildPlaceholder() {
    return const Center(
      child: Icon(
        Icons.person,
        size: 80,
        color: Color(0xFFBBB5C0),
      ),
    );
  }

  Widget _buildBadge(String text, Color color) {
    return Container(
      padding: const EdgeInsets.symmetric(
        horizontal: DesignTokens.spacingMd,
        vertical: DesignTokens.spacingXs,
      ),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(DesignTokens.radiusSm),
        border: Border.all(color: color, width: 3),
      ),
      child: Text(
        text,
        style: TextStyle(
          fontSize: 20,
          fontWeight: FontWeight.w800,
          color: color,
          letterSpacing: 2,
        ),
      ),
    );
  }
}
