//  Issue #1 设计资产预览入口
// 
//  本文件用于独立编译和预览所有认证登录体系相关的设计状态。
//  运行方式（在 designs/issue-1/flutter-widgets/ 目录下）：
//    flutter run -d chrome preview.dart
// 
//  预览内容：
//  - SplashPage 各状态（检查中、网络错误）
//  - LoginPage 默认布局
//  - LoginPage 完整交互预览（可点击、可输入）
//  - PhoneLoginForm 错误态预览
//  - ToTextField 各变体预览

import 'package:flutter/material.dart';
import 'package:type_one/theme/tokens.dart';

import 'login_page_design.dart';
import 'phone_login_form_design.dart';
import 'splash_page_design.dart';
import 'to_text_field_design.dart';

void main() {
  runApp(const _PreviewApp());
}

class _PreviewApp extends StatelessWidget {
  const _PreviewApp();

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Issue #1 设计预览',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        brightness: Brightness.light,
        scaffoldBackgroundColor: DesignTokens.background,
        fontFamily: DesignTokens.fontFamily,
      ),
      home: const _PreviewHomePage(),
    );
  }
}

class _PreviewHomePage extends StatelessWidget {
  const _PreviewHomePage();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Issue #1 认证登录体系设计预览'),
        backgroundColor: DesignTokens.background,
        foregroundColor: DesignTokens.textPrimary,
      ),
      body: ListView(
        padding: const EdgeInsets.all(DesignTokens.spacingMd),
        children: [
          _SectionTitle('SplashPage'),
          _PreviewCard(
            label: '检查中状态',
            child: SizedBox(
              height: 400,
              child: ClipRRect(
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                child: const SplashPageChecking(),
              ),
            ),
          ),
          _PreviewCard(
            label: '网络错误提示',
            child: SizedBox(
              height: 400,
              child: ClipRRect(
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                child: const SplashPageError(),
              ),
            ),
          ),
          _SectionTitle('LoginPage'),
          _PreviewCard(
            label: '默认布局（静态示意）',
            child: SizedBox(
              height: 600,
              child: ClipRRect(
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                child: const LoginPageDefault(
                  subheadline: '登录后体验更多功能',
                ),
              ),
            ),
          ),
          _PreviewCard(
            label: '完整交互预览（可点击）',
            child: SizedBox(
              height: 600,
              child: ClipRRect(
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                child: const LoginPageInteractivePreview(),
              ),
            ),
          ),
          _SectionTitle('PhoneLoginForm'),
          _PreviewCard(
            label: '错误态预览',
            child: Container(
              padding: const EdgeInsets.all(DesignTokens.spacingLg),
              decoration: BoxDecoration(
                color: DesignTokens.background,
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                border: Border.all(color: DesignTokens.border),
              ),
              child: const PhoneLoginForm(
                phone: '13800138000',
                code: '000000',
                phoneError: '请输入正确的手机号',
                codeError: '验证码错误，请重新输入',
                countdown: 45,
              ),
            ),
          ),
          _PreviewCard(
            label: '倒计时中状态',
            child: Container(
              padding: const EdgeInsets.all(DesignTokens.spacingLg),
              decoration: BoxDecoration(
                color: DesignTokens.background,
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                border: Border.all(color: DesignTokens.border),
              ),
              child: const PhoneLoginForm(
                phone: '13812345678',
                code: '123',
                countdown: 32,
              ),
            ),
          ),
          _SectionTitle('ToTextField（新增 Atom）'),
          _PreviewCard(
            label: '默认态 / 聚焦态 / 错误态',
            child: Container(
              padding: const EdgeInsets.all(DesignTokens.spacingLg),
              decoration: BoxDecoration(
                color: DesignTokens.background,
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                border: Border.all(color: DesignTokens.border),
              ),
              child: const Column(
                children: [
                  ToTextField(
                    hintText: '默认状态',
                  ),
                  SizedBox(height: DesignTokens.spacingMd),
                  ToTextField(
                    hintText: '带前缀',
                    prefix: Text('+86 '),
                  ),
                  SizedBox(height: DesignTokens.spacingMd),
                  ToTextField(
                    hintText: '错误状态',
                    errorText: '请输入正确的手机号',
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

class _SectionTitle extends StatelessWidget {
  final String text;

  const _SectionTitle(this.text);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(
        top: DesignTokens.spacingXl,
        bottom: DesignTokens.spacingSm,
      ),
      child: Text(
        text,
        style: const TextStyle(
          fontSize: DesignTokens.fontSizeLg,
          fontWeight: DesignTokens.fontWeightBold,
          color: DesignTokens.textPrimary,
        ),
      ),
    );
  }
}

class _PreviewCard extends StatelessWidget {
  final String label;
  final Widget child;

  const _PreviewCard({
    required this.label,
    required this.child,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.only(
            top: DesignTokens.spacingMd,
            bottom: DesignTokens.spacingSm,
          ),
          child: Text(
            label,
            style: const TextStyle(
              fontSize: DesignTokens.fontSizeSm,
              fontWeight: DesignTokens.fontWeightMedium,
              color: DesignTokens.textSecondary,
            ),
          ),
        ),
        child,
      ],
    );
  }
}
