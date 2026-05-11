import 'package:flutter/material.dart';
import 'swipe_card.dart';

/// 卡片堆叠容器，管理多层卡片渲染与手势冲突
///
/// 下层卡片呈现 scale 缩小和向下偏移的视觉差效果。
class SwipeCardStack extends StatelessWidget {
  /// 卡片数据列表，按从底到顶排序
  final List<SwipeCardData> cards;

  const SwipeCardStack({
    super.key,
    required this.cards,
  });

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: Alignment.center,
      children: [
        // Background cards (behind the front card)
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

/// 卡片数据模型
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
