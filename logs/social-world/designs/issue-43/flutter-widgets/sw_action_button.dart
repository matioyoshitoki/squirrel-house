import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';

/// 圆形操作按钮（Pass / Like 两种变体）
///
/// 用于首页卡片堆叠下方的操作区，支持点击和手势反馈。
enum SwActionType { pass, like }

class SwActionButton extends StatelessWidget {
  /// 按钮类型：pass 为跳过，like 为喜欢
  final SwActionType type;

  /// 点击回调
  final VoidCallback? onTap;

  /// 按钮直径，默认 64
  final double size;

  const SwActionButton({
    super.key,
    required this.type,
    this.onTap,
    this.size = 64,
  });

  @override
  Widget build(BuildContext context) {
    final color = type == SwActionType.pass
        ? DesignTokens.error
        : const Color(0xFF2E7D32); // likeActionColor
    final icon = type == SwActionType.pass ? Icons.close : Icons.favorite;

    return Semantics(
      identifier: type == SwActionType.pass ? 'pass_button' : 'like_button',
      label: type == SwActionType.pass ? '跳过' : '喜欢',
      button: true,
      container: true,
      child: GestureDetector(
        behavior: HitTestBehavior.opaque,
        onTapDown: (_) => onTap?.call(),
        child: Container(
          width: size,
          height: size,
          decoration: BoxDecoration(
            color: Colors.white,
            shape: BoxShape.circle,
            boxShadow: [
              BoxShadow(
                color: color.withOpacity(0.25),
                blurRadius: 12,
                offset: const Offset(0, 4),
              ),
            ],
            border: Border.all(
              color: color.withOpacity(0.3),
              width: 2,
            ),
          ),
          child: Icon(
            icon,
            color: color,
            size: size * 0.4,
            semanticLabel: type == SwActionType.pass ? '跳过' : '喜欢',
          ),
        ),
      ),
    );
  }
}
