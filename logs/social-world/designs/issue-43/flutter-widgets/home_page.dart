import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';

import 'sw_action_button.dart';
import 'swipe_card_stack.dart';
import 'empty_recommendation_view.dart';

/// 首页页面 Skeleton
///
/// 承载模式切换、卡片堆叠、操作按钮和底部导航。
/// 实际开发时需接入 MatchingBloc 状态管理。
class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    // TODO: 接入 MatchingBloc，根据状态渲染不同视图
    return const HomePageSuccess();
  }
}

// ============================================================================
// HomePage - Loading State
// ============================================================================
class HomePageLoading extends StatelessWidget {
  const HomePageLoading({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      body: SafeArea(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.all(DesignTokens.spacingMd),
              child: _ModeTabs(activeMode: 'nearby'),
            ),
            const Expanded(
              child: Center(
                child: CircularProgressIndicator(
                  color: DesignTokens.primary,
                ),
              ),
            ),
            const _BottomNav(selectedIndex: 0),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// HomePage - Empty State
// ============================================================================
class HomePageEmpty extends StatelessWidget {
  const HomePageEmpty({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      body: SafeArea(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.all(DesignTokens.spacingMd),
              child: _ModeTabs(activeMode: 'nearby'),
            ),
            const Expanded(
              child: EmptyRecommendationView(),
            ),
            const _BottomNav(selectedIndex: 0),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// HomePage - Error State
// ============================================================================
class HomePageError extends StatelessWidget {
  final VoidCallback? onRetry;

  const HomePageError({super.key, this.onRetry});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      body: SafeArea(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.all(DesignTokens.spacingMd),
              child: _ModeTabs(activeMode: 'nearby'),
            ),
            Expanded(
              child: Center(
                child: Padding(
                  padding: const EdgeInsets.all(DesignTokens.spacingXl),
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      const Icon(
                        Icons.error_outline,
                        size: 64,
                        color: DesignTokens.error,
                      ),
                      const SizedBox(height: DesignTokens.spacingMd),
                      const Text(
                        '加载失败',
                        style: TextStyle(
                          fontSize: DesignTokens.fontSizeTitle,
                          fontWeight: FontWeight.bold,
                          color: DesignTokens.textPrimary,
                        ),
                      ),
                      const SizedBox(height: DesignTokens.spacingXs),
                      const Text(
                        '网络异常，请检查网络后重试',
                        textAlign: TextAlign.center,
                        style: TextStyle(
                          fontSize: DesignTokens.fontSizeBody,
                          color: DesignTokens.secondaryColor,
                        ),
                      ),
                      const SizedBox(height: DesignTokens.spacingLg),
                      Semantics(
                        identifier: 'retry_button',
                        label: '重新加载',
                        button: true,
                        child: GestureDetector(
                          behavior: HitTestBehavior.opaque,
                          onTapDown: (_) => onRetry?.call(),
                          child: Container(
                            padding: const EdgeInsets.symmetric(
                              horizontal: DesignTokens.spacingXl,
                              vertical: DesignTokens.spacingMd,
                            ),
                            decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(
                                DesignTokens.radiusMd,
                              ),
                              border: Border.all(
                                color: DesignTokens.primary,
                              ),
                            ),
                            child: const Text(
                              '重新加载',
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
              ),
            ),
            const _BottomNav(selectedIndex: 0),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// HomePage - Success State (with cards)
// ============================================================================
class HomePageSuccess extends StatelessWidget {
  const HomePageSuccess({super.key});

  @override
  Widget build(BuildContext context) {
    final cards = [
      SwipeCardData(
        name: 'Alice',
        age: 24,
        distance: '2.5km',
        bio: '喜欢旅行、摄影、咖啡 ☕️',
      ),
      SwipeCardData(
        name: 'Bob',
        age: 26,
        distance: '5.1km',
        bio: '爱健身、看电影 🎬',
      ),
      SwipeCardData(
        name: 'Charlie',
        age: 23,
        distance: '1.2km',
        bio: '程序员，养了两只猫 🐱',
      ),
    ];

    return Scaffold(
      backgroundColor: DesignTokens.background,
      body: SafeArea(
        child: Column(
          children: [
            // Mode tabs
            Padding(
              padding: const EdgeInsets.all(DesignTokens.spacingMd),
              child: _ModeTabs(activeMode: 'nearby'),
            ),
            // Card stack
            Expanded(
              child: SwipeCardStack(cards: cards),
            ),
            // Action buttons
            Padding(
              padding: const EdgeInsets.symmetric(
                vertical: DesignTokens.spacingMd,
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  SwActionButton(
                    type: SwActionType.pass,
                    onTap: () {},
                  ),
                  const SizedBox(width: DesignTokens.spacingXl),
                  SwActionButton(
                    type: SwActionType.like,
                    onTap: () {},
                  ),
                ],
              ),
            ),
            // Bottom nav
            const _BottomNav(selectedIndex: 0),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// HomePage - Match Created State (overlay dialog)
// ============================================================================
class HomePageMatchCreated extends StatelessWidget {
  final String matchedUserName;
  final VoidCallback? onChat;
  final VoidCallback? onContinue;

  const HomePageMatchCreated({
    super.key,
    required this.matchedUserName,
    this.onChat,
    this.onContinue,
  });

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        const HomePageSuccess(),
        // Match success dialog overlay
        // TODO: import and use MatchSuccessDialog
      ],
    );
  }
}

// ============================================================================
// Internal: Mode Tabs
// ============================================================================
class _ModeTabs extends StatelessWidget {
  final String activeMode;

  const _ModeTabs({required this.activeMode});

  @override
  Widget build(BuildContext context) {
    final modes = [
      _ModeData('nearby', '附近'),
      _ModeData('city', '同城'),
      _ModeData('nationwide', '全国'),
    ];

    return Container(
      padding: const EdgeInsets.all(4),
      decoration: BoxDecoration(
        color: const Color(0xFFF3EDF7),
        borderRadius: BorderRadius.circular(24),
      ),
      child: Row(
        children: modes.map((mode) {
          final isActive = mode.key == activeMode;
          return Expanded(
            child: Semantics(
              identifier: 'mode_tab_${mode.key}',
              label: mode.label,
              button: true,
              child: GestureDetector(
                behavior: HitTestBehavior.opaque,
                onTapDown: (_) {},
                child: Container(
                  padding: const EdgeInsets.symmetric(vertical: 8),
                  decoration: BoxDecoration(
                    color: isActive ? DesignTokens.primary : Colors.transparent,
                    borderRadius: BorderRadius.circular(20),
                  ),
                  child: Text(
                    mode.label,
                    textAlign: TextAlign.center,
                    style: TextStyle(
                      fontSize: 14,
                      fontWeight: FontWeight.w500,
                      color: isActive
                          ? Colors.white
                          : DesignTokens.secondaryColor,
                    ),
                  ),
                ),
              ),
            ),
          );
        }).toList(),
      ),
    );
  }
}

class _ModeData {
  final String key;
  final String label;
  _ModeData(this.key, this.label);
}

// ============================================================================
// Internal: Bottom Navigation
// ============================================================================
class _BottomNav extends StatelessWidget {
  final int selectedIndex;

  const _BottomNav({required this.selectedIndex});

  @override
  Widget build(BuildContext context) {
    final items = [
      _NavItem(Icons.home, '首页'),
      _NavItem(Icons.chat_bubble, '消息'),
      _NavItem(Icons.person, '我的'),
    ];

    return Container(
      height: 80,
      decoration: const BoxDecoration(
        color: Colors.white,
        border: Border(top: BorderSide(color: Color(0xFFEEEEEE))),
      ),
      child: SafeArea(
        child: Row(
          children: items.asMap().entries.map((entry) {
            final index = entry.key;
            final item = entry.value;
            final isSelected = index == selectedIndex;
            return Expanded(
              child: Semantics(
                identifier: 'nav_${item.label}',
                label: item.label,
                button: true,
                child: GestureDetector(
                  behavior: HitTestBehavior.opaque,
                  onTapDown: (_) {},
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Icon(
                        item.icon,
                        color: isSelected
                            ? DesignTokens.primary
                            : DesignTokens.secondaryColor,
                        size: 24,
                      ),
                      const SizedBox(height: 4),
                      Text(
                        item.label,
                        style: TextStyle(
                          fontSize: DesignTokens.fontSizeCaption,
                          color: isSelected
                              ? DesignTokens.primary
                              : DesignTokens.secondaryColor,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            );
          }).toList(),
        ),
      ),
    );
  }
}

class _NavItem {
  final IconData icon;
  final String label;
  _NavItem(this.icon, this.label);
}
