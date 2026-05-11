/// Social World — Issue #4 Matching 设计预览
/// 
/// 可独立编译的 Flutter Web 预览文件。
/// 展示：推荐卡片堆叠、滑动匹配、配对成功弹窗、空状态、匹配列表。
///
/// 编译命令：flutter build web --target=preview.dart
import 'package:flutter/material.dart';

void main() {
  runApp(const MatchingPreviewApp());
}

// ============================================================
// Design Tokens (本地副本，确保 standalone 编译)
// ============================================================
class _Tokens {
  static const Color primary = Color(0xFF6750A4);
  static const Color background = Color(0xFFFFFBFE);
  static const Color textPrimary = Color(0xFF1C1B1F);
  static const Color error = Color(0xFFB3261E);
  static const Color surface = Color(0xFFF3EDF7);
  static const Color outline = Color(0xFF79747E);
  static const Color success = Color(0xFF2E7D32);
  static const Color successContainer = Color(0xFFE5F0E5);
  static const Color errorContainer = Color(0xFFF9DEDC);

  static const double fontSizeTitle = 22.0;
  static const double fontSizeBody = 16.0;
  static const double fontSizeCaption = 12.0;

  static const double spacingXs = 8.0;
  static const double spacingSm = 12.0;
  static const double spacingMd = 16.0;
  static const double spacingLg = 24.0;
  static const double spacingXl = 32.0;

  static const double radiusSm = 8.0;
  static const double radiusMd = 12.0;
  static const double radiusLg = 16.0;
}

// ============================================================
// App
// ============================================================
class MatchingPreviewApp extends StatelessWidget {
  const MatchingPreviewApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Social World - Matching Preview',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: _Tokens.primary),
        useMaterial3: true,
        fontFamily: '-apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica Neue, Arial, sans-serif',
      ),
      home: const _PreviewHome(),
    );
  }
}

class _PreviewHome extends StatefulWidget {
  const _PreviewHome();

  @override
  State<_PreviewHome> createState() => _PreviewHomeState();
}

class _PreviewHomeState extends State<_PreviewHome> {
  int _selectedIndex = 0;

  final _screens = const [
    _HomeScreen(),
    _EmptyScreen(),
    _MatchSuccessOverlayScreen(),
    _MatchListScreen(),
    _LoadingScreen(),
  ];

  final _labels = [
    '首页卡片',
    '空状态',
    '配对成功',
    '匹配列表',
    '加载中',
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFFE8E8E8),
      body: Center(
        child: Container(
          width: 390,
          height: 844,
          decoration: BoxDecoration(
            color: _Tokens.background,
            borderRadius: BorderRadius.circular(40),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.25),
                blurRadius: 50,
                offset: const Offset(0, 25),
              ),
            ],
          ),
          clipBehavior: Clip.antiAlias,
          child: Column(
            children: [
              // Status bar
              Container(
                height: 44,
                color: _Tokens.background,
                alignment: Alignment.center,
                child: const Text(
                  '9:41',
                  style: TextStyle(
                    fontSize: 14,
                    fontWeight: FontWeight.w600,
                    color: _Tokens.textPrimary,
                  ),
                ),
              ),
              // Screen content
              Expanded(child: _screens[_selectedIndex]),
              // Bottom tab selector (for preview navigation)
              Container(
                height: 56,
                decoration: BoxDecoration(
                  color: Colors.white,
                  border: Border(
                    top: BorderSide(color: Colors.black.withOpacity(0.05)),
                  ),
                ),
                child: SingleChildScrollView(
                  scrollDirection: Axis.horizontal,
                  child: Row(
                    children: List.generate(_labels.length, (index) {
                      final isActive = index == _selectedIndex;
                      return GestureDetector(
                        onTap: () => setState(() => _selectedIndex = index),
                        child: Container(
                          padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                          margin: const EdgeInsets.symmetric(horizontal: 4),
                          decoration: BoxDecoration(
                            color: isActive ? _Tokens.primary : _Tokens.surface,
                            borderRadius: BorderRadius.circular(16),
                          ),
                          child: Text(
                            _labels[index],
                            style: TextStyle(
                              fontSize: 12,
                              fontWeight: isActive ? FontWeight.w600 : FontWeight.normal,
                              color: isActive ? Colors.white : _Tokens.textPrimary,
                            ),
                          ),
                        ),
                      );
                    }),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

// ============================================================
// Screen 1: Home / Card Stack
// ============================================================
class _HomeScreen extends StatefulWidget {
  const _HomeScreen();

  @override
  State<_HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<_HomeScreen> {
  final _cards = [
    _CardData(emoji: '👩🏻', name: '小雨', age: 24, distance: '2.5km', active: '刚刚活跃', bio: '喜欢旅行、摄影、咖啡 ☕️ 周末常去探店'),
    _CardData(emoji: '👨🏻', name: '阿杰', age: 27, distance: '5.1km', active: '2小时前活跃', bio: '程序员、吉他手、铲屎官 🐱'),
    _CardData(emoji: '👩🏼', name: 'Alice', age: 25, distance: '1.2km', active: '5分钟前活跃', bio: '瑜伽教练，爱吃甜品 🍰'),
  ];

  int _currentIndex = 0;
  String _mode = 'nearby';

  void _onSwipe(String direction) {
    setState(() {
      _currentIndex++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Mode tabs
        Padding(
          padding: const EdgeInsets.all(_Tokens.spacingMd),
          child: Container(
            decoration: BoxDecoration(
              color: _Tokens.surface,
              borderRadius: BorderRadius.circular(24),
            ),
            padding: const EdgeInsets.all(4),
            child: Row(
              children: [
                _buildModeTab('附近', 'nearby'),
                _buildModeTab('同城', 'city'),
                _buildModeTab('全国', 'nationwide'),
              ],
            ),
          ),
        ),
        // Card stack
        Expanded(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: _Tokens.spacingMd),
            child: _currentIndex < _cards.length
                ? _CardStack(
                    cards: _cards.sublist(_currentIndex),
                    onSwipe: _onSwipe,
                  )
                : const _EmptyBody(),
          ),
        ),
        // Action buttons
        if (_currentIndex < _cards.length)
          Padding(
            padding: const EdgeInsets.all(_Tokens.spacingMd),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                _ActionButton(
                  icon: Icons.close,
                  color: _Tokens.error,
                  backgroundColor: _Tokens.errorContainer,
                  onPressed: () => _onSwipe('pass'),
                ),
                const SizedBox(width: 48),
                _ActionButton(
                  icon: Icons.favorite,
                  color: _Tokens.success,
                  backgroundColor: _Tokens.successContainer,
                  onPressed: () => _onSwipe('like'),
                ),
              ],
            ),
          ),
        // Bottom nav
        _buildBottomNav(activeIndex: 0),
      ],
    );
  }

  Widget _buildModeTab(String label, String mode) {
    final isActive = _mode == mode;
    return Expanded(
      child: GestureDetector(
        onTap: () => setState(() => _mode = mode),
        child: Container(
          padding: const EdgeInsets.symmetric(vertical: 8),
          decoration: BoxDecoration(
            color: isActive ? _Tokens.primary : Colors.transparent,
            borderRadius: BorderRadius.circular(20),
          ),
          alignment: Alignment.center,
          child: Text(
            label,
            style: TextStyle(
              fontSize: 14,
              fontWeight: FontWeight.w500,
              color: isActive ? Colors.white : _Tokens.textPrimary.withOpacity(0.7),
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildBottomNav({required int activeIndex}) {
    final items = [
      (Icons.home, '首页'),
      (Icons.chat_bubble_outline, '消息'),
      (Icons.person_outline, '我的'),
    ];
    return Container(
      height: 80,
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border(
          top: BorderSide(color: Colors.black.withOpacity(0.05)),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: List.generate(items.length, (i) {
          final isActive = i == activeIndex;
          return Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                width: isActive ? 32 : 24,
                height: isActive ? 32 : 24,
                decoration: isActive
                    ? const BoxDecoration(
                        color: _Tokens.primary,
                        shape: BoxShape.circle,
                      )
                    : null,
                child: Icon(
                  items[i].$1,
                  size: isActive ? 18 : 20,
                  color: isActive ? Colors.white : _Tokens.outline,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                items[i].$2,
                style: TextStyle(
                  fontSize: _Tokens.fontSizeCaption,
                  color: isActive ? _Tokens.primary : _Tokens.outline,
                  fontWeight: isActive ? FontWeight.w500 : FontWeight.normal,
                ),
              ),
            ],
          );
        }),
      ),
    );
  }
}

class _CardData {
  final String emoji;
  final String name;
  final int age;
  final String distance;
  final String active;
  final String bio;

  const _CardData({
    required this.emoji,
    required this.name,
    required this.age,
    required this.distance,
    required this.active,
    required this.bio,
  });
}

class _CardStack extends StatelessWidget {
  final List<_CardData> cards;
  final ValueChanged<String> onSwipe;

  const _CardStack({required this.cards, required this.onSwipe});

  @override
  Widget build(BuildContext context) {
    return Stack(
      clipBehavior: Clip.none,
      children: List.generate(cards.length, (i) {
        final card = cards[i];
        if (i == 0) {
          return _SwipeableCard(
            data: card,
            onSwipe: onSwipe,
          );
        }
        final scale = 1.0 - (i * 0.05);
        final translateY = i * 12.0;
        final opacity = 1.0 - (i * 0.2);
        return Transform.translate(
          offset: Offset(0, translateY),
          child: Transform.scale(
            scale: scale,
            child: Opacity(
              opacity: opacity.clamp(0.0, 1.0),
              child: _StaticCard(data: card),
            ),
          ),
        );
      }).reversed.toList(),
    );
  }
}

class _StaticCard extends StatelessWidget {
  final _CardData data;
  const _StaticCard({required this.data});

  @override
  Widget build(BuildContext context) {
    return Container(
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
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          Expanded(
            child: Container(
              decoration: const BoxDecoration(
                gradient: LinearGradient(
                  begin: Alignment.topLeft,
                  end: Alignment.bottomRight,
                  colors: [Color(0xFFE8E0EB), Color(0xFFD4C8DB)],
                ),
              ),
              alignment: Alignment.center,
              child: Text(data.emoji, style: const TextStyle(fontSize: 80)),
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(_Tokens.spacingMd),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  '${data.name}, ${data.age}',
                  style: const TextStyle(
                    fontSize: _Tokens.fontSizeTitle,
                    fontWeight: FontWeight.bold,
                    color: _Tokens.textPrimary,
                  ),
                ),
                const SizedBox(height: 4),
                Text(
                  '📍 ${data.distance} · ${data.active}',
                  style: const TextStyle(
                    fontSize: 14,
                    color: _Tokens.textPrimary,
                  ),
                ),
                const SizedBox(height: _Tokens.spacingXs),
                Text(
                  data.bio,
                  style: const TextStyle(
                    fontSize: 14,
                    color: _Tokens.outline,
                  ),
                  maxLines: 2,
                  overflow: TextOverflow.ellipsis,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

class _SwipeableCard extends StatefulWidget {
  final _CardData data;
  final ValueChanged<String> onSwipe;

  const _SwipeableCard({required this.data, required this.onSwipe});

  @override
  State<_SwipeableCard> createState() => _SwipeableCardState();
}

class _SwipeableCardState extends State<_SwipeableCard>
    with SingleTickerProviderStateMixin {
  double _dx = 0;
  bool _isDragging = false;

  void _onDragUpdate(DragUpdateDetails details) {
    setState(() {
      _dx += details.delta.dx;
      _isDragging = true;
    });
  }

  void _onDragEnd(DragEndDetails details) {
    if (_dx > 120) {
      _animateSwipe('like', 400);
    } else if (_dx < -120) {
      _animateSwipe('pass', -400);
    } else {
      setState(() {
        _dx = 0;
        _isDragging = false;
      });
    }
  }

  void _animateSwipe(String direction, double target) {
    // In a real app we'd use AnimationController; here we just snap for preview
    setState(() {
      _dx = target;
    });
    Future.delayed(const Duration(milliseconds: 300), () {
      widget.onSwipe(direction);
    });
  }

  double get _rotation => _dx * 0.001;
  double get _likeOpacity => _dx > 60 ? (_dx / 120).clamp(0.0, 1.0) : 0.0;
  double get _passOpacity => _dx < -60 ? (_dx.abs() / 120).clamp(0.0, 1.0) : 0.0;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onHorizontalDragUpdate: _onDragUpdate,
      onHorizontalDragEnd: _onDragEnd,
      child: Transform.translate(
        offset: Offset(_dx, 0),
        child: Transform.rotate(
          angle: _rotation,
          child: Stack(
            children: [
              _StaticCard(data: widget.data),
              // Pass badge
              Positioned(
                top: 24,
                left: 24,
                child: Opacity(
                  opacity: _passOpacity,
                  child: Container(
                    padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                    decoration: BoxDecoration(
                      border: Border.all(color: _Tokens.error, width: 3),
                      borderRadius: BorderRadius.circular(_Tokens.radiusSm),
                    ),
                    child: const Text(
                      'PASS',
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.w800,
                        color: _Tokens.error,
                        letterSpacing: 2,
                      ),
                    ),
                  ),
                ),
              ),
              // Like badge
              Positioned(
                top: 24,
                right: 24,
                child: Opacity(
                  opacity: _likeOpacity,
                  child: Container(
                    padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                    decoration: BoxDecoration(
                      border: Border.all(color: _Tokens.success, width: 3),
                      borderRadius: BorderRadius.circular(_Tokens.radiusSm),
                    ),
                    child: const Text(
                      'LIKE',
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.w800,
                        color: _Tokens.success,
                        letterSpacing: 2,
                      ),
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class _ActionButton extends StatelessWidget {
  final IconData icon;
  final Color color;
  final Color backgroundColor;
  final VoidCallback onPressed;

  const _ActionButton({
    required this.icon,
    required this.color,
    required this.backgroundColor,
    required this.onPressed,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onPressed,
      child: Container(
        width: 64,
        height: 64,
        decoration: BoxDecoration(
          color: Colors.white,
          shape: BoxShape.circle,
          border: Border.all(color: backgroundColor, width: 2),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withOpacity(0.1),
              blurRadius: 8,
              offset: const Offset(0, 2),
            ),
          ],
        ),
        alignment: Alignment.center,
        child: Icon(icon, size: 28, color: color),
      ),
    );
  }
}

// ============================================================
// Screen 2: Empty State
// ============================================================
class _EmptyScreen extends StatelessWidget {
  const _EmptyScreen();

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Padding(
          padding: const EdgeInsets.all(_Tokens.spacingMd),
          child: Container(
            decoration: BoxDecoration(
              color: _Tokens.surface,
              borderRadius: BorderRadius.circular(24),
            ),
            padding: const EdgeInsets.all(4),
            child: Row(
              children: [
                _buildTab('附近', true),
                _buildTab('同城', false),
                _buildTab('全国', false),
              ],
            ),
          ),
        ),
        const Expanded(child: _EmptyBody()),
        _buildBottomNav(activeIndex: 0),
      ],
    );
  }

  Widget _buildTab(String label, bool isActive) {
    return Expanded(
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 8),
        decoration: BoxDecoration(
          color: isActive ? _Tokens.primary : Colors.transparent,
          borderRadius: BorderRadius.circular(20),
        ),
        alignment: Alignment.center,
        child: Text(
          label,
          style: TextStyle(
            fontSize: 14,
            fontWeight: FontWeight.w500,
            color: isActive ? Colors.white : _Tokens.textPrimary.withOpacity(0.7),
          ),
        ),
      ),
    );
  }

  Widget _buildBottomNav({required int activeIndex}) {
    final items = [
      (Icons.home, '首页'),
      (Icons.chat_bubble_outline, '消息'),
      (Icons.person_outline, '我的'),
    ];
    return Container(
      height: 80,
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border(
          top: BorderSide(color: Colors.black.withOpacity(0.05)),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: List.generate(items.length, (i) {
          final isActive = i == activeIndex;
          return Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                width: isActive ? 32 : 24,
                height: isActive ? 32 : 24,
                decoration: isActive
                    ? const BoxDecoration(
                        color: _Tokens.primary,
                        shape: BoxShape.circle,
                      )
                    : null,
                child: Icon(
                  items[i].$1,
                  size: isActive ? 18 : 20,
                  color: isActive ? Colors.white : _Tokens.outline,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                items[i].$2,
                style: TextStyle(
                  fontSize: _Tokens.fontSizeCaption,
                  color: isActive ? _Tokens.primary : _Tokens.outline,
                  fontWeight: isActive ? FontWeight.w500 : FontWeight.normal,
                ),
              ),
            ],
          );
        }),
      ),
    );
  }
}

class _EmptyBody extends StatelessWidget {
  const _EmptyBody();

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(_Tokens.spacingXl),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Container(
              width: 120,
              height: 120,
              decoration: const BoxDecoration(
                color: _Tokens.surface,
                shape: BoxShape.circle,
              ),
              alignment: Alignment.center,
              child: const Icon(
                Icons.search_off,
                size: 56,
                color: Color(0xFFBBB),
              ),
            ),
            const SizedBox(height: _Tokens.spacingLg),
            const Text(
              '附近暂无可推荐用户',
              style: TextStyle(
                fontSize: _Tokens.fontSizeBody,
                fontWeight: FontWeight.w600,
                color: _Tokens.textPrimary,
              ),
            ),
            const SizedBox(height: _Tokens.spacingXs),
            const Text(
              '扩大范围或稍后再来看看吧',
              style: TextStyle(
                fontSize: 14,
                color: _Tokens.outline,
              ),
            ),
            const SizedBox(height: _Tokens.spacingLg),
            OutlinedButton(
              onPressed: () {},
              style: OutlinedButton.styleFrom(
                foregroundColor: _Tokens.primary,
                side: const BorderSide(color: _Tokens.outline),
                padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 12),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(24),
                ),
              ),
              child: const Text('切换至全国模式'),
            ),
          ],
        ),
      ),
    );
  }
}

// ============================================================
// Screen 3: Match Success Overlay
// ============================================================
class _MatchSuccessOverlayScreen extends StatelessWidget {
  const _MatchSuccessOverlayScreen();

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        // Background: dimmed home screen
        Opacity(
          opacity: 0.4,
          child: Column(
            children: [
              Padding(
                padding: const EdgeInsets.all(_Tokens.spacingMd),
                child: Container(
                  decoration: BoxDecoration(
                    color: _Tokens.surface,
                    borderRadius: BorderRadius.circular(24),
                  ),
                  padding: const EdgeInsets.all(4),
                  child: Row(
                    children: [
                      _buildTab('附近', true),
                      _buildTab('同城', false),
                      _buildTab('全国', false),
                    ],
                  ),
                ),
              ),
              Expanded(
                child: Padding(
                  padding: const EdgeInsets.symmetric(horizontal: _Tokens.spacingMd),
                  child: _StaticCard(
                    data: _CardData(
                      emoji: '👩🏻',
                      name: '小雨',
                      age: 24,
                      distance: '2.5km',
                      active: '刚刚活跃',
                      bio: '喜欢旅行、摄影、咖啡 ☕️',
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 96),
              _buildBottomNav(activeIndex: 0),
            ],
          ),
        ),
        // Match success modal
        Container(
          color: Colors.black.withOpacity(0.6),
          alignment: Alignment.center,
          padding: const EdgeInsets.all(_Tokens.spacingMd),
          child: Container(
            width: double.infinity,
            constraints: const BoxConstraints(maxWidth: 320),
            decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(_Tokens.radiusLg),
            ),
            padding: const EdgeInsets.all(_Tokens.spacingXl),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Text(
                  '🎉 配对成功！',
                  style: TextStyle(
                    fontSize: _Tokens.fontSizeTitle,
                    fontWeight: FontWeight.bold,
                    color: _Tokens.textPrimary,
                  ),
                ),
                const SizedBox(height: _Tokens.spacingLg),
                Stack(
                  alignment: Alignment.center,
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Container(
                          width: 72,
                          height: 72,
                          decoration: BoxDecoration(
                            color: _Tokens.surface,
                            shape: BoxShape.circle,
                            border: Border.all(color: Colors.white, width: 3),
                            boxShadow: [
                              BoxShadow(
                                color: Colors.black.withOpacity(0.1),
                                blurRadius: 8,
                              ),
                            ],
                          ),
                          alignment: Alignment.center,
                          child: const Text('😊', style: TextStyle(fontSize: 32)),
                        ),
                        const SizedBox(width: 24),
                        Container(
                          width: 72,
                          height: 72,
                          decoration: BoxDecoration(
                            color: _Tokens.surface,
                            shape: BoxShape.circle,
                            border: Border.all(color: Colors.white, width: 3),
                            boxShadow: [
                              BoxShadow(
                                color: Colors.black.withOpacity(0.1),
                                blurRadius: 8,
                              ),
                            ],
                          ),
                          alignment: Alignment.center,
                          child: const Text('👩🏻', style: TextStyle(fontSize: 32)),
                        ),
                      ],
                    ),
                    Container(
                      width: 32,
                      height: 32,
                      decoration: const BoxDecoration(
                        color: _Tokens.primary,
                        shape: BoxShape.circle,
                        border: Border.fromBorderSide(
                          BorderSide(color: Colors.white, width: 2),
                        ),
                      ),
                      alignment: Alignment.center,
                      child: const Text('💕', style: TextStyle(fontSize: 14)),
                    ),
                  ],
                ),
                const SizedBox(height: _Tokens.spacingLg),
                const Text(
                  '你们互相喜欢了对方',
                  style: TextStyle(
                    fontSize: 14,
                    color: _Tokens.textPrimary,
                  ),
                ),
                const SizedBox(height: 4),
                const Text(
                  '现在可以开始聊天啦',
                  style: TextStyle(
                    fontSize: 13,
                    color: _Tokens.outline,
                  ),
                ),
                const SizedBox(height: _Tokens.spacingLg),
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    onPressed: () {},
                    style: ElevatedButton.styleFrom(
                      backgroundColor: _Tokens.primary,
                      foregroundColor: Colors.white,
                      padding: const EdgeInsets.symmetric(vertical: 14),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(24),
                      ),
                    ),
                    child: const Text(
                      '立即聊天',
                      style: TextStyle(fontSize: 14, fontWeight: FontWeight.w600),
                    ),
                  ),
                ),
                const SizedBox(height: _Tokens.spacingSm),
                SizedBox(
                  width: double.infinity,
                  child: OutlinedButton(
                    onPressed: () {},
                    style: OutlinedButton.styleFrom(
                      foregroundColor: _Tokens.textPrimary,
                      side: const BorderSide(color: _Tokens.outline),
                      padding: const EdgeInsets.symmetric(vertical: 14),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(24),
                      ),
                    ),
                    child: const Text(
                      '继续滑动',
                      style: TextStyle(fontSize: 14),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }

  Widget _buildTab(String label, bool isActive) {
    return Expanded(
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 8),
        decoration: BoxDecoration(
          color: isActive ? _Tokens.primary : Colors.transparent,
          borderRadius: BorderRadius.circular(20),
        ),
        alignment: Alignment.center,
        child: Text(
          label,
          style: TextStyle(
            fontSize: 14,
            fontWeight: FontWeight.w500,
            color: isActive ? Colors.white : _Tokens.textPrimary.withOpacity(0.7),
          ),
        ),
      ),
    );
  }

  Widget _buildBottomNav({required int activeIndex}) {
    final items = [
      (Icons.home, '首页'),
      (Icons.chat_bubble_outline, '消息'),
      (Icons.person_outline, '我的'),
    ];
    return Container(
      height: 80,
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border(
          top: BorderSide(color: Colors.black.withOpacity(0.05)),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: List.generate(items.length, (i) {
          final isActive = i == activeIndex;
          return Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                width: isActive ? 32 : 24,
                height: isActive ? 32 : 24,
                decoration: isActive
                    ? const BoxDecoration(
                        color: _Tokens.primary,
                        shape: BoxShape.circle,
                      )
                    : null,
                child: Icon(
                  items[i].$1,
                  size: isActive ? 18 : 20,
                  color: isActive ? Colors.white : _Tokens.outline,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                items[i].$2,
                style: TextStyle(
                  fontSize: _Tokens.fontSizeCaption,
                  color: isActive ? _Tokens.primary : _Tokens.outline,
                  fontWeight: isActive ? FontWeight.w500 : FontWeight.normal,
                ),
              ),
            ],
          );
        }),
      ),
    );
  }
}

// ============================================================
// Screen 4: Match List
// ============================================================
class _MatchListScreen extends StatelessWidget {
  const _MatchListScreen();

  @override
  Widget build(BuildContext context) {
    final matches = [
      (emoji: '👩🏻', name: '小雨', time: '2小时前配对'),
      (emoji: '👨🏻', name: '阿杰', time: '昨天配对'),
      (emoji: '👩🏼', name: 'Alice', time: '3天前配对'),
      (emoji: '👨🏽', name: 'Bob', time: '上周配对'),
    ];

    return Column(
      children: [
        Container(
          padding: const EdgeInsets.symmetric(vertical: _Tokens.spacingSm),
          alignment: Alignment.center,
          child: const Text(
            '我的配对',
            style: TextStyle(
              fontSize: _Tokens.fontSizeTitle,
              fontWeight: FontWeight.bold,
              color: _Tokens.textPrimary,
            ),
          ),
        ),
        Expanded(
          child: ListView.builder(
            padding: const EdgeInsets.symmetric(horizontal: _Tokens.spacingMd),
            itemCount: matches.length,
            itemBuilder: (context, i) {
              final match = matches[i];
              return Container(
                margin: const EdgeInsets.only(bottom: _Tokens.spacingSm),
                padding: const EdgeInsets.all(_Tokens.spacingMd),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(_Tokens.radiusMd),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black.withOpacity(0.05),
                      blurRadius: 3,
                      offset: const Offset(0, 1),
                    ),
                  ],
                ),
                child: Row(
                  children: [
                    Container(
                      width: 48,
                      height: 48,
                      decoration: const BoxDecoration(
                        color: _Tokens.surface,
                        shape: BoxShape.circle,
                      ),
                      alignment: Alignment.center,
                      child: Text(match.emoji, style: const TextStyle(fontSize: 22)),
                    ),
                    const SizedBox(width: _Tokens.spacingMd),
                    Expanded(
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            match.name,
                            style: const TextStyle(
                              fontSize: _Tokens.fontSizeBody,
                              fontWeight: FontWeight.w600,
                              color: _Tokens.textPrimary,
                            ),
                          ),
                          const SizedBox(height: 2),
                          Text(
                            match.time,
                            style: const TextStyle(
                              fontSize: 13,
                              color: _Tokens.outline,
                            ),
                          ),
                        ],
                      ),
                    ),
                    Container(
                      width: 40,
                      height: 40,
                      decoration: BoxDecoration(
                        color: _Tokens.primary.withOpacity(0.08),
                        shape: BoxShape.circle,
                      ),
                      alignment: Alignment.center,
                      child: const Icon(
                        Icons.chat_bubble_outline,
                        size: 18,
                        color: _Tokens.primary,
                      ),
                    ),
                  ],
                ),
              );
            },
          ),
        ),
        _buildBottomNav(activeIndex: 1),
      ],
    );
  }

  Widget _buildBottomNav({required int activeIndex}) {
    final items = [
      (Icons.home, '首页'),
      (Icons.chat_bubble, '消息'),
      (Icons.person_outline, '我的'),
    ];
    return Container(
      height: 80,
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border(
          top: BorderSide(color: Colors.black.withOpacity(0.05)),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: List.generate(items.length, (i) {
          final isActive = i == activeIndex;
          return Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                width: isActive ? 32 : 24,
                height: isActive ? 32 : 24,
                decoration: isActive
                    ? const BoxDecoration(
                        color: _Tokens.primary,
                        shape: BoxShape.circle,
                      )
                    : null,
                child: Icon(
                  items[i].$1,
                  size: isActive ? 18 : 20,
                  color: isActive ? Colors.white : _Tokens.outline,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                items[i].$2,
                style: TextStyle(
                  fontSize: _Tokens.fontSizeCaption,
                  color: isActive ? _Tokens.primary : _Tokens.outline,
                  fontWeight: isActive ? FontWeight.w500 : FontWeight.normal,
                ),
              ),
            ],
          );
        }),
      ),
    );
  }
}

// ============================================================
// Screen 5: Loading
// ============================================================
class _LoadingScreen extends StatelessWidget {
  const _LoadingScreen();

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Padding(
          padding: const EdgeInsets.all(_Tokens.spacingMd),
          child: Container(
            decoration: BoxDecoration(
              color: _Tokens.surface,
              borderRadius: BorderRadius.circular(24),
            ),
            padding: const EdgeInsets.all(4),
            child: Row(
              children: [
                _buildTab('附近', true),
                _buildTab('同城', false),
                _buildTab('全国', false),
              ],
            ),
          ),
        ),
        const Expanded(
          child: Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                CircularProgressIndicator(
                  valueColor: AlwaysStoppedAnimation<Color>(_Tokens.primary),
                ),
                SizedBox(height: _Tokens.spacingMd),
                Text(
                  '正在寻找附近的人…',
                  style: TextStyle(
                    fontSize: 14,
                    color: _Tokens.outline,
                  ),
                ),
              ],
            ),
          ),
        ),
        _buildBottomNav(activeIndex: 0),
      ],
    );
  }

  Widget _buildTab(String label, bool isActive) {
    return Expanded(
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 8),
        decoration: BoxDecoration(
          color: isActive ? _Tokens.primary : Colors.transparent,
          borderRadius: BorderRadius.circular(20),
        ),
        alignment: Alignment.center,
        child: Text(
          label,
          style: TextStyle(
            fontSize: 14,
            fontWeight: FontWeight.w500,
            color: isActive ? Colors.white : _Tokens.textPrimary.withOpacity(0.7),
          ),
        ),
      ),
    );
  }

  Widget _buildBottomNav({required int activeIndex}) {
    final items = [
      (Icons.home, '首页'),
      (Icons.chat_bubble_outline, '消息'),
      (Icons.person_outline, '我的'),
    ];
    return Container(
      height: 80,
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border(
          top: BorderSide(color: Colors.black.withOpacity(0.05)),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: List.generate(items.length, (i) {
          final isActive = i == activeIndex;
          return Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                width: isActive ? 32 : 24,
                height: isActive ? 32 : 24,
                decoration: isActive
                    ? const BoxDecoration(
                        color: _Tokens.primary,
                        shape: BoxShape.circle,
                      )
                    : null,
                child: Icon(
                  items[i].$1,
                  size: isActive ? 18 : 20,
                  color: isActive ? Colors.white : _Tokens.outline,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                items[i].$2,
                style: TextStyle(
                  fontSize: _Tokens.fontSizeCaption,
                  color: isActive ? _Tokens.primary : _Tokens.outline,
                  fontWeight: isActive ? FontWeight.w500 : FontWeight.normal,
                ),
              ),
            ],
          );
        }),
      ),
    );
  }
}
