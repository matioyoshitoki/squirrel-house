// ignore_for_file: avoid_print
// Social World - Issue #43 Matching Design Preview
// 自包含 Flutter Web 预览，不依赖项目私有包

import 'package:flutter/material.dart';

void main() {
  runApp(const MatchingPreviewApp());
}

// ============================================================================
// Local Design Tokens (不依赖外部包)
// ============================================================================
class _Tokens {
  static const Color primary = Color(0xFF6750A4);
  static const Color background = Color(0xFFFFFBFE);
  static const Color textPrimary = Color(0xFF1C1B1F);
  static const Color error = Color(0xFFB3261E);
  static const Color success = Color(0xFF2E7D32);
  static const Color surface = Color(0xFFF3EDF7);
  static const Color outline = Color(0xFF79747E);
  static const Color onSurfaceVariant = Color(0xFF49454F);

  static const double fontSizeHeadline = 28.0;
  static const double fontSizeTitle = 22.0;
  static const double fontSizeBody = 16.0;
  static const double fontSizeCaption = 12.0;

  static const double spacingXxs = 4.0;
  static const double spacingXs = 8.0;
  static const double spacingSm = 12.0;
  static const double spacingMd = 16.0;
  static const double spacingLg = 24.0;
  static const double spacingXl = 32.0;

  static const double radiusSm = 8.0;
  static const double radiusMd = 12.0;
  static const double radiusLg = 16.0;
}

// ============================================================================
// Preview App
// ============================================================================
class MatchingPreviewApp extends StatelessWidget {
  const MatchingPreviewApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Issue #43 Matching Preview',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: _Tokens.primary),
        useMaterial3: true,
      ),
      home: const PreviewHomePage(),
    );
  }
}

class PreviewHomePage extends StatelessWidget {
  const PreviewHomePage({super.key});

  @override
  Widget build(BuildContext context) {
    final items = [
      _PreviewItem('首页 - 加载态', const HomePageLoading()),
      _PreviewItem('首页 - 空状态', const HomePageEmpty()),
      _PreviewItem('首页 - 错误态', const HomePageError()),
      _PreviewItem('首页 - 推荐卡片', const HomePageSuccess()),
      _PreviewItem('配对成功弹窗', const MatchSuccessDialogPreview()),
      _PreviewItem('匹配列表 - 加载态', const MatchListPageLoading()),
      _PreviewItem('匹配列表 - 空状态', const MatchListPageEmpty()),
      _PreviewItem('匹配列表 - 错误态', const MatchListPageError()),
      _PreviewItem('匹配列表 - 正常态', const MatchListPageSuccess()),
      _PreviewItem('原子组件 - SwActionButton', const SwActionButtonPreview()),
      _PreviewItem('分子组件 - SwipeCard', const SwipeCardPreview()),
      _PreviewItem('分子组件 - MatchListItem', const MatchListItemPreview()),
    ];

    return Scaffold(
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.primary,
        foregroundColor: Colors.white,
        title: const Text('Issue #43 Matching 设计预览'),
      ),
      body: ListView.builder(
        padding: const EdgeInsets.all(_Tokens.spacingMd),
        itemCount: items.length,
        itemBuilder: (context, index) {
          final item = items[index];
          return Card(
            margin: const EdgeInsets.only(bottom: _Tokens.spacingSm),
            child: ListTile(
              title: Text(item.title),
              trailing: const Icon(Icons.chevron_right),
              onTap: () {
                Navigator.of(context).push(
                  MaterialPageRoute(builder: (_) => item.widget),
                );
              },
            ),
          );
        },
      ),
    );
  }
}

class _PreviewItem {
  final String title;
  final Widget widget;
  _PreviewItem(this.title, this.widget);
}

// ============================================================================
// Atom: SwActionButton
// ============================================================================
enum SwActionType { pass, like }

class SwActionButton extends StatelessWidget {
  final SwActionType type;
  final VoidCallback? onTap;
  final double size;

  const SwActionButton({
    super.key,
    required this.type,
    this.onTap,
    this.size = 64,
  });

  @override
  Widget build(BuildContext context) {
    final color = type == SwActionType.pass ? _Tokens.error : _Tokens.success;
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
            border: Border.all(color: color.withOpacity(0.3), width: 2),
          ),
          child: Icon(icon, color: color, size: size * 0.4),
        ),
      ),
    );
  }
}

class SwActionButtonPreview extends StatelessWidget {
  const SwActionButtonPreview({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.primary,
        foregroundColor: Colors.white,
        title: const Text('SwActionButton'),
      ),
      body: Center(
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SwActionButton(
              type: SwActionType.pass,
              onTap: () => print('Pass tapped'),
            ),
            const SizedBox(width: _Tokens.spacingXl),
            SwActionButton(
              type: SwActionType.like,
              onTap: () => print('Like tapped'),
            ),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// Molecule: SwipeCard
// ============================================================================
class SwipeCard extends StatelessWidget {
  final String name;
  final int age;
  final String distance;
  final String bio;
  final String? imageUrl;
  final double rotation;
  final double likeOpacity;
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
        margin: const EdgeInsets.symmetric(horizontal: _Tokens.spacingMd),
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(_Tokens.radiusLg),
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
                    top: _Tokens.spacingLg,
                    left: _Tokens.spacingLg,
                    child: Opacity(
                      opacity: likeOpacity,
                      child: _buildBadge('LIKE', _Tokens.success),
                    ),
                  ),
                  // Pass badge
                  Positioned(
                    top: _Tokens.spacingLg,
                    right: _Tokens.spacingLg,
                    child: Opacity(
                      opacity: passOpacity,
                      child: _buildBadge('PASS', _Tokens.error),
                    ),
                  ),
                ],
              ),
            ),
            // Info area
            Padding(
              padding: const EdgeInsets.all(_Tokens.spacingMd),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    '$name, $age',
                    style: const TextStyle(
                      fontSize: _Tokens.fontSizeTitle,
                      fontWeight: FontWeight.bold,
                      color: _Tokens.textPrimary,
                    ),
                  ),
                  const SizedBox(height: _Tokens.spacingXs),
                  Row(
                    children: [
                      const Icon(
                        Icons.location_on,
                        size: 14,
                        color: _Tokens.onSurfaceVariant,
                      ),
                      const SizedBox(width: _Tokens.spacingXxs),
                      Text(
                        distance,
                        style: const TextStyle(
                          fontSize: 14,
                          color: _Tokens.onSurfaceVariant,
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(height: _Tokens.spacingXs),
                  Text(
                    bio,
                    maxLines: 2,
                    overflow: TextOverflow.ellipsis,
                    style: const TextStyle(
                      fontSize: 14,
                      color: _Tokens.outline,
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
      child: Icon(Icons.person, size: 80, color: Color(0xFFBBB5C0)),
    );
  }

  Widget _buildBadge(String text, Color color) {
    return Container(
      padding: const EdgeInsets.symmetric(
        horizontal: _Tokens.spacingMd,
        vertical: _Tokens.spacingXs,
      ),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(_Tokens.radiusSm),
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

class SwipeCardPreview extends StatelessWidget {
  const SwipeCardPreview({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.primary,
        foregroundColor: Colors.white,
        title: const Text('SwipeCard'),
      ),
      body: ListView(
        padding: const EdgeInsets.all(_Tokens.spacingMd),
        children: const [
          Text('默认态', style: TextStyle(fontWeight: FontWeight.bold)),
          SizedBox(height: _Tokens.spacingSm),
          SwipeCard(
            name: 'Alice',
            age: 24,
            distance: '2.5km',
            bio: '喜欢旅行、摄影、咖啡 ☕️',
          ),
          SizedBox(height: _Tokens.spacingXl),
          Text('拖动中 - Like', style: TextStyle(fontWeight: FontWeight.bold)),
          SizedBox(height: _Tokens.spacingSm),
          SwipeCard(
            name: 'Bob',
            age: 26,
            distance: '5.1km',
            bio: '爱健身、看电影 🎬',
            rotation: 0.15,
            likeOpacity: 1.0,
          ),
          SizedBox(height: _Tokens.spacingXl),
          Text('拖动中 - Pass', style: TextStyle(fontWeight: FontWeight.bold)),
          SizedBox(height: _Tokens.spacingSm),
          SwipeCard(
            name: 'Charlie',
            age: 23,
            distance: '1.2km',
            bio: '程序员，养了两只猫 🐱',
            rotation: -0.15,
            passOpacity: 1.0,
          ),
        ],
      ),
    );
  }
}

// ============================================================================
// Organism: SwipeCardStack
// ============================================================================
class SwipeCardStack extends StatelessWidget {
  final List<SwipeCardData> cards;

  const SwipeCardStack({super.key, required this.cards});

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: Alignment.center,
      children: [
        // Background cards
        for (int i = cards.length - 1; i > 0; i--)
          Positioned(
            top: (cards.length - i) * 12.0,
            left: 0,
            right: 0,
            child: Transform.scale(
              scale: 1.0 - (cards.length - i) * 0.05,
              child: Opacity(
                opacity: 0.7 - (cards.length - i) * 0.15,
                child: SwipeCard(
                  name: cards[i].name,
                  age: cards[i].age,
                  distance: cards[i].distance,
                  bio: cards[i].bio,
                  imageUrl: cards[i].imageUrl,
                ),
              ),
            ),
          ),
        // Front card
        if (cards.isNotEmpty)
          SwipeCard(
            name: cards.first.name,
            age: cards.first.age,
            distance: cards.first.distance,
            bio: cards.first.bio,
            imageUrl: cards.first.imageUrl,
          ),
      ],
    );
  }
}

class SwipeCardData {
  final String name;
  final int age;
  final String distance;
  final String bio;
  final String? imageUrl;

  SwipeCardData({
    required this.name,
    required this.age,
    required this.distance,
    required this.bio,
    this.imageUrl,
  });
}

// ============================================================================
// Organism: MatchSuccessDialog
// ============================================================================
class MatchSuccessDialog extends StatelessWidget {
  final String currentUserAvatar;
  final String matchedUserAvatar;
  final String matchedUserName;
  final VoidCallback? onChat;
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
            margin: const EdgeInsets.all(_Tokens.spacingMd),
            padding: const EdgeInsets.all(_Tokens.spacingXl),
            decoration: BoxDecoration(
              color: _Tokens.background,
              borderRadius: BorderRadius.circular(_Tokens.radiusLg),
            ),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Text(
                  '配对成功！',
                  style: TextStyle(
                    fontSize: _Tokens.fontSizeHeadline,
                    fontWeight: FontWeight.bold,
                    color: _Tokens.primary,
                  ),
                ),
                const SizedBox(height: _Tokens.spacingMd),
                Text(
                  '你和 $matchedUserName 互相喜欢了',
                  textAlign: TextAlign.center,
                  style: const TextStyle(
                    fontSize: _Tokens.fontSizeBody,
                    color: _Tokens.onSurfaceVariant,
                  ),
                ),
                const SizedBox(height: _Tokens.spacingXl),
                // Avatar combination
                Stack(
                  alignment: Alignment.center,
                  children: [
                    Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        _buildAvatar(currentUserAvatar, _Tokens.spacingLg),
                        const SizedBox(width: _Tokens.spacingMd),
                        _buildAvatar(matchedUserAvatar, 0),
                      ],
                    ),
                    Container(
                      width: 40,
                      height: 40,
                      decoration: BoxDecoration(
                        color: _Tokens.success,
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
                const SizedBox(height: _Tokens.spacingXl),
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    style: ElevatedButton.styleFrom(
                      backgroundColor: _Tokens.primary,
                      foregroundColor: Colors.white,
                      padding: const EdgeInsets.symmetric(
                        vertical: _Tokens.spacingMd,
                      ),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(_Tokens.radiusMd),
                      ),
                    ),
                    onPressed: onChat,
                    child: const Text('开始聊天'),
                  ),
                ),
                const SizedBox(height: _Tokens.spacingMd),
                SizedBox(
                  width: double.infinity,
                  child: OutlinedButton(
                    style: OutlinedButton.styleFrom(
                      foregroundColor: _Tokens.primary,
                      padding: const EdgeInsets.symmetric(
                        vertical: _Tokens.spacingMd,
                      ),
                      side: const BorderSide(color: _Tokens.primary),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(_Tokens.radiusMd),
                      ),
                    ),
                    onPressed: onContinue,
                    child: const Text('继续滑动'),
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
          color: _Tokens.primary.withOpacity(0.2),
          shape: BoxShape.circle,
          border: Border.all(color: Colors.white, width: 3),
        ),
        clipBehavior: Clip.antiAlias,
        child: url.isNotEmpty
            ? Image.network(url, fit: BoxFit.cover)
            : const Icon(Icons.person, size: 40, color: _Tokens.primary),
      ),
    );
  }
}

class MatchSuccessDialogPreview extends StatelessWidget {
  const MatchSuccessDialogPreview({super.key});

  @override
  Widget build(BuildContext context) {
    return MatchSuccessDialog(
      currentUserAvatar: '',
      matchedUserAvatar: '',
      matchedUserName: 'Alice',
      onChat: () => print('Chat tapped'),
      onContinue: () => Navigator.of(context).pop(),
    );
  }
}

// ============================================================================
// Molecule: MatchListItem
// ============================================================================
class MatchListItem extends StatelessWidget {
  final String name;
  final String? avatarUrl;
  final String matchTime;
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
            horizontal: _Tokens.spacingMd,
            vertical: _Tokens.spacingSm,
          ),
          decoration: const BoxDecoration(
            border: Border(
              bottom: BorderSide(color: Color(0xFFEEEEEE)),
            ),
          ),
          child: Row(
            children: [
              Container(
                width: 48,
                height: 48,
                decoration: BoxDecoration(
                  color: _Tokens.primary.withOpacity(0.2),
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
              const SizedBox(width: _Tokens.spacingMd),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      name,
                      style: const TextStyle(
                        fontSize: _Tokens.fontSizeBody,
                        fontWeight: FontWeight.w600,
                        color: _Tokens.textPrimary,
                      ),
                    ),
                    const SizedBox(height: _Tokens.spacingXxs),
                    Text(
                      matchTime,
                      style: const TextStyle(
                        fontSize: _Tokens.fontSizeCaption,
                        color: _Tokens.outline,
                      ),
                    ),
                  ],
                ),
              ),
              const Icon(
                Icons.chevron_right,
                color: _Tokens.outline,
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildFallback() {
    return const Center(
      child: Icon(Icons.person, size: 24, color: _Tokens.primary),
    );
  }
}

class MatchListItemPreview extends StatelessWidget {
  const MatchListItemPreview({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.primary,
        foregroundColor: Colors.white,
        title: const Text('MatchListItem'),
      ),
      body: ListView(
        children: [
          MatchListItem(
            name: 'Alice',
            matchTime: '2小时前匹配',
            onTap: () => print('Alice tapped'),
          ),
          MatchListItem(
            name: 'Bob',
            matchTime: '昨天匹配',
            onTap: () => print('Bob tapped'),
          ),
          MatchListItem(
            name: 'Charlie',
            matchTime: '3天前匹配',
            onTap: () => print('Charlie tapped'),
          ),
        ],
      ),
    );
  }
}

// ============================================================================
// Page: HomePage - Loading
// ============================================================================
class HomePageLoading extends StatelessWidget {
  const HomePageLoading({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      body: SafeArea(
        child: Column(
          children: [
            // Mode tabs placeholder
            Padding(
              padding: const EdgeInsets.all(_Tokens.spacingMd),
              child: Container(
                height: 40,
                decoration: BoxDecoration(
                  color: _Tokens.surface,
                  borderRadius: BorderRadius.circular(20),
                ),
              ),
            ),
            const Expanded(
              child: Center(
                child: CircularProgressIndicator(color: _Tokens.primary),
              ),
            ),
            // Bottom nav placeholder
            Container(
              height: 80,
              decoration: const BoxDecoration(
                border: Border(top: BorderSide(color: Color(0xFFEEEEEE))),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// Page: HomePage - Empty
// ============================================================================
class HomePageEmpty extends StatelessWidget {
  const HomePageEmpty({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      body: SafeArea(
        child: Column(
          children: [
            // Mode tabs
            Padding(
              padding: const EdgeInsets.all(_Tokens.spacingMd),
              child: _buildModeTabs('nearby'),
            ),
            const Expanded(
              child: EmptyRecommendationView(),
            ),
            // Bottom nav
            _buildBottomNav(0),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// Page: HomePage - Error
// ============================================================================
class HomePageError extends StatelessWidget {
  const HomePageError({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      body: SafeArea(
        child: Column(
          children: [
            // Mode tabs
            Padding(
              padding: const EdgeInsets.all(_Tokens.spacingMd),
              child: _buildModeTabs('nearby'),
            ),
            Expanded(
              child: Center(
                child: Padding(
                  padding: const EdgeInsets.all(_Tokens.spacingXl),
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      const Icon(
                        Icons.error_outline,
                        size: 64,
                        color: _Tokens.error,
                      ),
                      const SizedBox(height: _Tokens.spacingMd),
                      const Text(
                        '加载失败',
                        style: TextStyle(
                          fontSize: _Tokens.fontSizeTitle,
                          fontWeight: FontWeight.bold,
                          color: _Tokens.textPrimary,
                        ),
                      ),
                      const SizedBox(height: _Tokens.spacingXs),
                      const Text(
                        '网络异常，请检查网络后重试',
                        textAlign: TextAlign.center,
                        style: TextStyle(
                          fontSize: _Tokens.fontSizeBody,
                          color: _Tokens.outline,
                        ),
                      ),
                      const SizedBox(height: _Tokens.spacingLg),
                      OutlinedButton(
                        style: OutlinedButton.styleFrom(
                          foregroundColor: _Tokens.primary,
                          side: const BorderSide(color: _Tokens.primary),
                          padding: const EdgeInsets.symmetric(
                            horizontal: _Tokens.spacingXl,
                            vertical: _Tokens.spacingMd,
                          ),
                          shape: RoundedRectangleBorder(
                            borderRadius:
                                BorderRadius.circular(_Tokens.radiusMd),
                          ),
                        ),
                        onPressed: () => print('Retry'),
                        child: const Text('重新加载'),
                      ),
                    ],
                  ),
                ),
              ),
            ),
            // Bottom nav
            _buildBottomNav(0),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// Page: HomePage - Success
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
      backgroundColor: _Tokens.background,
      body: SafeArea(
        child: Column(
          children: [
            // Mode tabs
            Padding(
              padding: const EdgeInsets.all(_Tokens.spacingMd),
              child: _buildModeTabs('nearby'),
            ),
            // Card stack
            Expanded(
              child: SwipeCardStack(cards: cards),
            ),
            // Action buttons
            Padding(
              padding: const EdgeInsets.symmetric(
                vertical: _Tokens.spacingMd,
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  SwActionButton(
                    type: SwActionType.pass,
                    onTap: () => print('Pass'),
                  ),
                  const SizedBox(width: _Tokens.spacingXl),
                  SwActionButton(
                    type: SwActionType.like,
                    onTap: () => print('Like'),
                  ),
                ],
              ),
            ),
            // Bottom nav
            _buildBottomNav(0),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// Widget: EmptyRecommendationView
// ============================================================================
class EmptyRecommendationView extends StatelessWidget {
  const EmptyRecommendationView({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(_Tokens.spacingXl),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Container(
              width: 120,
              height: 120,
              decoration: BoxDecoration(
                color: _Tokens.surface,
                shape: BoxShape.circle,
              ),
              child: const Icon(
                Icons.search_off,
                size: 48,
                color: _Tokens.outline,
              ),
            ),
            const SizedBox(height: _Tokens.spacingLg),
            const Text(
              '附近暂无可推荐用户',
              style: TextStyle(
                fontSize: _Tokens.fontSizeTitle,
                fontWeight: FontWeight.bold,
                color: _Tokens.textPrimary,
              ),
            ),
            const SizedBox(height: _Tokens.spacingXs),
            const Text(
              '扩大范围或稍后再来看看吧',
              textAlign: TextAlign.center,
              style: TextStyle(
                fontSize: _Tokens.fontSizeBody,
                color: _Tokens.outline,
              ),
            ),
            const SizedBox(height: _Tokens.spacingLg),
            OutlinedButton(
              style: OutlinedButton.styleFrom(
                foregroundColor: _Tokens.primary,
                side: const BorderSide(color: _Tokens.primary),
                padding: const EdgeInsets.symmetric(
                  horizontal: _Tokens.spacingXl,
                  vertical: _Tokens.spacingMd,
                ),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(_Tokens.radiusMd),
                ),
              ),
              onPressed: () => print('Switch nationwide'),
              child: const Text('切换全国'),
            ),
          ],
        ),
      ),
    );
  }
}

// ============================================================================
// Page: MatchListPage - Loading
// ============================================================================
class MatchListPageLoading extends StatelessWidget {
  const MatchListPageLoading({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.background,
        elevation: 0,
        centerTitle: true,
        title: const Text(
          '消息',
          style: TextStyle(
            color: _Tokens.textPrimary,
            fontSize: _Tokens.fontSizeTitle,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: const Center(
        child: CircularProgressIndicator(color: _Tokens.primary),
      ),
      bottomNavigationBar: _buildBottomNav(1),
    );
  }
}

// ============================================================================
// Page: MatchListPage - Empty
// ============================================================================
class MatchListPageEmpty extends StatelessWidget {
  const MatchListPageEmpty({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.background,
        elevation: 0,
        centerTitle: true,
        title: const Text(
          '消息',
          style: TextStyle(
            color: _Tokens.textPrimary,
            fontSize: _Tokens.fontSizeTitle,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(_Tokens.spacingXl),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Container(
                width: 120,
                height: 120,
                decoration: BoxDecoration(
                  color: _Tokens.surface,
                  shape: BoxShape.circle,
                ),
                child: const Icon(
                  Icons.chat_bubble_outline,
                  size: 48,
                  color: _Tokens.outline,
                ),
              ),
              const SizedBox(height: _Tokens.spacingLg),
              const Text(
                '暂无匹配记录',
                style: TextStyle(
                  fontSize: _Tokens.fontSizeTitle,
                  fontWeight: FontWeight.bold,
                  color: _Tokens.textPrimary,
                ),
              ),
              const SizedBox(height: _Tokens.spacingXs),
              const Text(
                '去首页滑动卡片，发现更多有趣的灵魂',
                textAlign: TextAlign.center,
                style: TextStyle(
                  fontSize: _Tokens.fontSizeBody,
                  color: _Tokens.outline,
                ),
              ),
            ],
          ),
        ),
      ),
      bottomNavigationBar: _buildBottomNav(1),
    );
  }
}

// ============================================================================
// Page: MatchListPage - Error
// ============================================================================
class MatchListPageError extends StatelessWidget {
  const MatchListPageError({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.background,
        elevation: 0,
        centerTitle: true,
        title: const Text(
          '消息',
          style: TextStyle(
            color: _Tokens.textPrimary,
            fontSize: _Tokens.fontSizeTitle,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(_Tokens.spacingXl),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              const Icon(
                Icons.error_outline,
                size: 64,
                color: _Tokens.error,
              ),
              const SizedBox(height: _Tokens.spacingMd),
              const Text(
                '加载失败',
                style: TextStyle(
                  fontSize: _Tokens.fontSizeTitle,
                  fontWeight: FontWeight.bold,
                  color: _Tokens.textPrimary,
                ),
              ),
              const SizedBox(height: _Tokens.spacingXs),
              const Text(
                '网络异常，请检查网络后重试',
                textAlign: TextAlign.center,
                style: TextStyle(
                  fontSize: _Tokens.fontSizeBody,
                  color: _Tokens.outline,
                ),
              ),
              const SizedBox(height: _Tokens.spacingLg),
              OutlinedButton(
                style: OutlinedButton.styleFrom(
                  foregroundColor: _Tokens.primary,
                  side: const BorderSide(color: _Tokens.primary),
                  padding: const EdgeInsets.symmetric(
                    horizontal: _Tokens.spacingXl,
                    vertical: _Tokens.spacingMd,
                  ),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(_Tokens.radiusMd),
                  ),
                ),
                onPressed: () => print('Retry match list'),
                child: const Text('重新加载'),
              ),
            ],
          ),
        ),
      ),
      bottomNavigationBar: _buildBottomNav(1),
    );
  }
}

// ============================================================================
// Page: MatchListPage - Success
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
      backgroundColor: _Tokens.background,
      appBar: AppBar(
        backgroundColor: _Tokens.background,
        elevation: 0,
        centerTitle: true,
        title: const Text(
          '消息',
          style: TextStyle(
            color: _Tokens.textPrimary,
            fontSize: _Tokens.fontSizeTitle,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: ListView.builder(
        itemCount: matches.length,
        itemBuilder: (context, index) {
          final m = matches[index];
          return MatchListItem(
            name: m.name,
            matchTime: m.time,
            onTap: () => print('Open chat with ${m.name}'),
          );
        },
      ),
      bottomNavigationBar: _buildBottomNav(1),
    );
  }
}

class _MatchData {
  final String name;
  final String time;
  _MatchData(this.name, this.time);
}

// ============================================================================
// Shared helpers
// ============================================================================
Widget _buildModeTabs(String activeMode) {
  final modes = [
    _ModeData('nearby', '附近'),
    _ModeData('city', '同城'),
    _ModeData('nationwide', '全国'),
  ];

  return Container(
    padding: const EdgeInsets.all(_Tokens.spacingXxs),
    decoration: BoxDecoration(
      color: _Tokens.surface,
      borderRadius: BorderRadius.circular(24),
    ),
    child: Row(
      children: modes.map((mode) {
        final isActive = mode.key == activeMode;
        return Expanded(
          child: Container(
            padding: const EdgeInsets.symmetric(vertical: 8),
            decoration: BoxDecoration(
              color: isActive ? _Tokens.primary : Colors.transparent,
              borderRadius: BorderRadius.circular(20),
            ),
            child: Text(
              mode.label,
              textAlign: TextAlign.center,
              style: TextStyle(
                fontSize: 14,
                fontWeight: FontWeight.w500,
                color: isActive ? Colors.white : _Tokens.onSurfaceVariant,
              ),
            ),
          ),
        );
      }).toList(),
    ),
  );
}

class _ModeData {
  final String key;
  final String label;
  _ModeData(this.key, this.label);
}

Widget _buildBottomNav(int selectedIndex) {
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
                onTapDown: (_) => print('Nav ${item.label}'),
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Icon(
                      item.icon,
                      color: isSelected ? _Tokens.primary : _Tokens.outline,
                      size: 24,
                    ),
                    const SizedBox(height: _Tokens.spacingXxs),
                    Text(
                      item.label,
                      style: TextStyle(
                        fontSize: _Tokens.fontSizeCaption,
                        color: isSelected ? _Tokens.primary : _Tokens.outline,
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

class _NavItem {
  final IconData icon;
  final String label;
  _NavItem(this.icon, this.label);
}
