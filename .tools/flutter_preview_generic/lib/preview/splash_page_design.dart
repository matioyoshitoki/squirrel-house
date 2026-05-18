//  SplashPage 页面设计
// 
//  启动页各状态的可视化参考。不含业务逻辑（无 Token 校验、无路由跳转）。
// 
//  状态列表：
//  - [SplashPageChecking]：检查中（默认态）
//  - [SplashPageError]：底部网络错误提示

import 'package:flutter/material.dart';
import 'package:type_one/theme/tokens.dart';

//  SplashPage — 检查中状态
// 
//  视觉：全屏白色背景，居中加载指示器 + App 名称。
class SplashPageChecking extends StatelessWidget {
  final String appName;

  const SplashPageChecking({
    super.key,
    this.appName = 'Type One',
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const CircularProgressIndicator(
              color: DesignTokens.primary,
              strokeWidth: 3,
            ),
            const SizedBox(height: DesignTokens.spacingMd),
            Text(
              appName,
              style: const TextStyle(
                fontSize: DesignTokens.fontSizeXl,
                fontWeight: DesignTokens.fontWeightBold,
                color: DesignTokens.textPrimary,
              ),
            ),
          ],
        ),
      ),
    );
  }
}

//  SplashPage — 带网络错误提示
// 
//  视觉：在 [SplashPageChecking] 基础上，底部出现 SnackBar 风格轻提示。
//  注意：实际实现中应使用 [ScaffoldMessenger.showSnackBar]，
//  本设计组件仅用于视觉预览。
class SplashPageError extends StatelessWidget {
  final String appName;
  final String errorMessage;

  const SplashPageError({
    super.key,
    this.appName = 'Type One',
    this.errorMessage = '网络异常，请检查网络连接',
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      body: Stack(
        children: [
          // 底层：检查中状态
          const Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                CircularProgressIndicator(
                  color: DesignTokens.primary,
                  strokeWidth: 3,
                ),
                SizedBox(height: DesignTokens.spacingMd),
                Text(
                  'Type One',
                  style: TextStyle(
                    fontSize: DesignTokens.fontSizeXl,
                    fontWeight: DesignTokens.fontWeightBold,
                    color: DesignTokens.textPrimary,
                  ),
                ),
              ],
            ),
          ),

          // 顶层：底部错误提示（模拟 SnackBar 视觉）
          Positioned(
            left: DesignTokens.spacingLg,
            right: DesignTokens.spacingLg,
            bottom: DesignTokens.spacingXl,
            child: Container(
              padding: const EdgeInsets.symmetric(
                horizontal: DesignTokens.spacingMd,
                vertical: DesignTokens.spacingSm,
              ),
              decoration: BoxDecoration(
                color: DesignTokens.textPrimary.withValues(alpha: 0.8),
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
              ),
              child: Row(
                children: [
                  const Icon(
                    Icons.wifi_off,
                    size: 18,
                    color: DesignTokens.background,
                  ),
                  const SizedBox(width: DesignTokens.spacingSm),
                  Expanded(
                    child: Text(
                      errorMessage,
                      style: const TextStyle(
                        fontSize: DesignTokens.fontSizeSm,
                        fontWeight: DesignTokens.fontWeightRegular,
                        color: DesignTokens.background,
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
