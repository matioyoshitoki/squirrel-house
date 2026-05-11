import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';

/// 推荐为空时的占位视图
///
/// 展示图标、文案和切换范围的引导按钮。
class EmptyRecommendationView extends StatelessWidget {
  /// 点击「切换全国」回调
  final VoidCallback? onSwitchNationwide;

  const EmptyRecommendationView({
    super.key,
    this.onSwitchNationwide,
  });

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(DesignTokens.spacingXl),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Container(
              width: 120,
              height: 120,
              decoration: BoxDecoration(
                color: DesignTokens.surfaceColor,
                shape: BoxShape.circle,
              ),
              child: const Icon(
                Icons.search_off,
                size: 48,
                color: DesignTokens.secondaryColor,
              ),
            ),
            const SizedBox(height: DesignTokens.spacingLg),
            const Text(
              '附近暂无可推荐用户',
              style: TextStyle(
                fontSize: DesignTokens.fontSizeTitle,
                fontWeight: FontWeight.bold,
                color: DesignTokens.textPrimary,
              ),
            ),
            const SizedBox(height: DesignTokens.spacingXs),
            const Text(
              '扩大范围或稍后再来看看吧',
              textAlign: TextAlign.center,
              style: TextStyle(
                fontSize: DesignTokens.fontSizeBody,
                color: DesignTokens.secondaryColor,
              ),
            ),
            const SizedBox(height: DesignTokens.spacingLg),
            Semantics(
              identifier: 'switch_nationwide_button',
              label: '切换全国',
              button: true,
              child: GestureDetector(
                behavior: HitTestBehavior.opaque,
                onTapDown: (_) => onSwitchNationwide?.call(),
                child: Container(
                  padding: const EdgeInsets.symmetric(
                    horizontal: DesignTokens.spacingXl,
                    vertical: DesignTokens.spacingMd,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                    border: Border.all(color: DesignTokens.primary),
                  ),
                  child: const Text(
                    '切换全国',
                    style: TextStyle(
                      color: DesignTokens.primary,
                      fontSize: DesignTokens.fontSizeBody,
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
