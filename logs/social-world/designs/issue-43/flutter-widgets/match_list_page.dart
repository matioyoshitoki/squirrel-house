import 'package:flutter/material.dart';
import 'package:social_world_design_system/social_world_design_system.dart';

import 'match_list_item.dart';

/// 匹配列表页 Skeleton
///
/// 底部导航「消息」Tab 内容，展示配对成功用户列表。
/// 实际开发时需接入 MatchListCubit 状态管理。
class MatchListPage extends StatelessWidget {
  const MatchListPage({super.key});

  @override
  Widget build(BuildContext context) {
    // TODO: 接入 MatchListCubit，根据状态渲染不同视图
    return const MatchListPageSuccess();
  }
}

// ============================================================================
// MatchListPage - Loading State
// ============================================================================
class MatchListPageLoading extends StatelessWidget {
  const MatchListPageLoading({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      appBar: _buildAppBar(),
      body: const Center(
        child: CircularProgressIndicator(
          color: DesignTokens.primary,
        ),
      ),
      bottomNavigationBar: const _BottomNav(selectedIndex: 1),
    );
  }
}

// ============================================================================
// MatchListPage - Empty State
// ============================================================================
class MatchListPageEmpty extends StatelessWidget {
  const MatchListPageEmpty({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      appBar: _buildAppBar(),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(DesignTokens.spacingXl),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Container(
                width: 120,
                height: 120,
                decoration: BoxDecoration(
                  color: const Color(0xFFF3EDF7),
                  shape: BoxShape.circle,
                ),
                child: const Icon(
                  Icons.chat_bubble_outline,
                  size: 48,
                  color: DesignTokens.secondaryColor,
                ),
              ),
              const SizedBox(height: DesignTokens.spacingLg),
              const Text(
                '暂无匹配记录',
                style: TextStyle(
                  fontSize: DesignTokens.fontSizeTitle,
                  fontWeight: FontWeight.bold,
                  color: DesignTokens.textPrimary,
                ),
              ),
              const SizedBox(height: DesignTokens.spacingXs),
              const Text(
                '去首页滑动卡片，发现更多有趣的灵魂',
                textAlign: TextAlign.center,
                style: TextStyle(
                  fontSize: DesignTokens.fontSizeBody,
                  color: DesignTokens.secondaryColor,
                ),
              ),
            ],
          ),
        ),
      ),
      bottomNavigationBar: const _BottomNav(selectedIndex: 1),
    );
  }
}

// ============================================================================
// MatchListPage - Error State
// ============================================================================
class MatchListPageError extends StatelessWidget {
  final VoidCallback? onRetry;

  const MatchListPageError({super.key, this.onRetry});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      appBar: _buildAppBar(),
      body: Center(
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
                      border: Border.all(color: DesignTokens.primary),
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
      bottomNavigationBar: const _BottomNav(selectedIndex: 1),
    );
  }
}

// ============================================================================
// MatchListPage - Success State
// ============================================================================
class MatchListPageSuccess extends StatelessWidget {
  const MatchListPageSuccess({super.key});

  @override
  Widget build(BuildContext context) {
    final matches = [
      _MatchData('Alice', '2小时前匹配'),
      _MatchData('Bob', '昨天匹配'),
      _MatchData('Charlie', '3天前匹配'),
      _MatchData('Diana', '1周前匹配'),
      _MatchData('Eve', '2周前匹配'),
    ];

    return Scaffold(
      backgroundColor: DesignTokens.background,
      appBar: _buildAppBar(),
      body: ListView.builder(
        itemCount: matches.length,
        itemBuilder: (context, index) {
          final m = matches[index];
          return MatchListItem(
            name: m.name,
            matchTime: m.time,
            onTap: () {},
          );
        },
      ),
      bottomNavigationBar: const _BottomNav(selectedIndex: 1),
    );
  }
}

// ============================================================================
// Internal helpers
// ============================================================================
AppBar _buildAppBar() {
  return AppBar(
    backgroundColor: DesignTokens.background,
    elevation: 0,
    centerTitle: true,
    title: const Text(
      '消息',
      style: TextStyle(
        color: DesignTokens.textPrimary,
        fontSize: DesignTokens.fontSizeTitle,
        fontWeight: FontWeight.bold,
      ),
    ),
  );
}

class _MatchData {
  final String name;
  final String time;
  _MatchData(this.name, this.time);
}

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
