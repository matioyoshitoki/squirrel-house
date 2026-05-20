/// SplashPage 设计原型
///
/// 展示所有状态：
/// - 检查中（默认）
/// - 网络异常提示
///
/// 设计约束：
/// - 无 AppBar，全屏展示
/// - 内容垂直居中
/// - 总耗时 ≤ 2 秒
/// - 跳转无转场动画

import 'package:flutter/material.dart';

// 内联 Token 定义（设计资产独立可预览）
class _Tokens {
  _Tokens._();
  static const Color primary = Color(0xFF1976D2);
  static const Color background = Color(0xFFFFFFFF);
  static const Color textPrimary = Color(0xFF212121);
  static const Color error = Color(0xFFD32F2F);

  static const double fontSizeXl = 24;
  static const double fontSizeSm = 14;
  static const double spacingMd = 16;
  static const double spacingLg = 24;
  static const double radiusMd = 8;

  static const FontWeight fontWeightBold = FontWeight.w700;
}

// ───────────────────────────────────────────────
// SplashPage 设计原型
// ───────────────────────────────────────────────

class SplashPageDesign extends StatefulWidget {
  const SplashPageDesign({super.key});

  @override
  State<SplashPageDesign> createState() => _SplashPageDesignState();
}

class _SplashPageDesignState extends State<SplashPageDesign> {
  bool _showNetworkError = false;

  void _toggleNetworkError() {
    setState(() => _showNetworkError = !_showNetworkError);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        fit: StackFit.expand,
        children: [
          // 主内容：居中加载指示器 + 品牌名
          Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              mainAxisSize: MainAxisSize.min,
              children: [
                // 品牌图标占位（圆形背景 + 首字母）
                Container(
                  width: 80,
                  height: 80,
                  decoration: BoxDecoration(
                    color: _Tokens.primary.withOpacity(0.1),
                    shape: BoxShape.circle,
                  ),
                  child: const Icon(
                    Icons.all_inclusive,
                    size: 40,
                    color: _Tokens.primary,
                  ),
                ),
                const SizedBox(height: _Tokens.spacingMd),
                const CircularProgressIndicator(
                  color: _Tokens.primary,
                  strokeWidth: 3,
                ),
                const SizedBox(height: _Tokens.spacingMd),
                const Text(
                  'Type One',
                  style: TextStyle(
                    fontSize: _Tokens.fontSizeXl,
                    fontWeight: _Tokens.fontWeightBold,
                    color: _Tokens.textPrimary,
                  ),
                ),
              ],
            ),
          ),

          // 网络异常提示条（底部）
          AnimatedPositioned(
            duration: const Duration(milliseconds: 250),
            curve: Curves.easeOut,
            bottom: _showNetworkError
                ? MediaQuery.of(context).padding.bottom + _Tokens.spacingLg
                : -60,
            left: _Tokens.spacingLg,
            right: _Tokens.spacingLg,
            child: AnimatedOpacity(
              duration: const Duration(milliseconds: 200),
              opacity: _showNetworkError ? 1.0 : 0.0,
              child: Container(
                padding: const EdgeInsets.symmetric(
                  horizontal: _Tokens.spacingLg,
                  vertical: _Tokens.spacingMd,
                ),
                decoration: BoxDecoration(
                  color: _Tokens.error.withOpacity(0.9),
                  borderRadius: BorderRadius.circular(_Tokens.radiusMd),
                ),
                child: const Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Icon(
                      Icons.wifi_off,
                      size: 16,
                      color: _Tokens.background,
                    ),
                    SizedBox(width: 8),
                    Text(
                      '网络异常，请检查网络连接',
                      style: TextStyle(
                        fontSize: _Tokens.fontSizeSm,
                        color: _Tokens.background,
                      ),
                    ),
                  ],
                ),
              ),
            ),
          ),

          // 调试按钮：切换网络异常提示（仅设计预览用）
          Positioned(
            top: MediaQuery.of(context).padding.top + _Tokens.spacingMd,
            right: _Tokens.spacingMd,
            child: GestureDetector(
              onTap: _toggleNetworkError,
              child: Container(
                padding: const EdgeInsets.symmetric(
                  horizontal: 12,
                  vertical: 6,
                ),
                decoration: BoxDecoration(
                  color: _Tokens.primary.withOpacity(0.1),
                  borderRadius: BorderRadius.circular(_Tokens.radiusMd),
                ),
                child: const Text(
                  '切换网络提示',
                  style: TextStyle(
                    fontSize: 12,
                    color: _Tokens.primary,
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
