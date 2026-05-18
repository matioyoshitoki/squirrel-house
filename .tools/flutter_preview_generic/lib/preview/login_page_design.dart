//  LoginPage 页面设计
// 
//  登录页完整布局，包含标题区、表单区、底部提示。
//  纯 UI 展示，不含业务逻辑（无 BLoC、无 API 调用）。
// 
//  为便于设计预览，本组件内置 [ValueNotifier] 模拟表单交互，
//  实际实现中应将状态提升至 LoginCubit。

import 'package:flutter/material.dart';
import 'package:type_one/theme/tokens.dart';

import 'phone_login_form_design.dart';

//  LoginPage — 默认状态
// 
//  视觉：白色背景，安全区包裹，垂直居中布局（可滚动）。
//  标题「欢迎登录」+ 副标题 + PhoneLoginForm。
class LoginPageDefault extends StatelessWidget {
  final String headline;
  final String? subheadline;

  const LoginPageDefault({
    super.key,
    this.headline = '欢迎登录',
    this.subheadline,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      resizeToAvoidBottomInset: true,
      body: SafeArea(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(
            horizontal: DesignTokens.spacingLg,
          ),
          child: ConstrainedBox(
            constraints: BoxConstraints(
              minHeight: MediaQuery.of(context).size.height -
                  MediaQuery.of(context).padding.top -
                  MediaQuery.of(context).padding.bottom,
            ),
            child: IntrinsicHeight(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  const Spacer(flex: 2),

                  // ── 标题区 ──
                  Text(
                    headline,
                    style: const TextStyle(
                      fontSize: DesignTokens.fontSizeXxl,
                      fontWeight: DesignTokens.fontWeightBold,
                      color: DesignTokens.textPrimary,
                    ),
                    textAlign: TextAlign.center,
                  ),
                  if (subheadline != null) ...[
                    const SizedBox(height: DesignTokens.spacingSm),
                    Text(
                      subheadline!,
                      style: const TextStyle(
                        fontSize: DesignTokens.fontSizeBase,
                        fontWeight: DesignTokens.fontWeightRegular,
                        color: DesignTokens.textSecondary,
                      ),
                      textAlign: TextAlign.center,
                    ),
                  ],
                  const SizedBox(height: DesignTokens.spacingXxl),

                  // ── 表单区 ──
                  // 实际实现中这里放 PhoneLoginForm
                  // 设计预览中使用一个占位示意
                  const _FormPlaceholder(),

                  const Spacer(flex: 3),

                  // ── 底部提示（可选）─
                  const _BottomHint(),
                  const SizedBox(height: DesignTokens.spacingMd),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}

//  LoginPage — 带完整交互预览
// 
//  使用 [ValueNotifier] 模拟表单状态变化，用于设计评审和视觉预览。
//  实际开发时，状态由 LoginCubit 管理，本组件仅作设计参考。
class LoginPageInteractivePreview extends StatefulWidget {
  const LoginPageInteractivePreview({super.key});

  @override
  State<LoginPageInteractivePreview> createState() =>
      _LoginPageInteractivePreviewState();
}

class _LoginPageInteractivePreviewState
    extends State<LoginPageInteractivePreview> {
  String _phone = '';
  String _code = '';
  String? _phoneError;
  String? _codeError;
  int _countdown = 0;
  bool _isLoading = false;
  bool _isSendCodeLoading = false;

  void _onPhoneChanged(String value) {
    setState(() {
      _phone = value;
      _phoneError = null;
    });
  }

  void _onCodeChanged(String value) {
    setState(() {
      _code = value;
      _codeError = null;
    });
  }

  void _onSendCode() {
    if (_phone.length != 11) {
      setState(() => _phoneError = '请输入正确的手机号');
      return;
    }
    setState(() => _isSendCodeLoading = true);
    // 模拟请求
    Future.delayed(const Duration(seconds: 1), () {
      if (!mounted) return;
      setState(() {
        _isSendCodeLoading = false;
        _countdown = 60;
      });
      _startCountdown();
    });
  }

  void _startCountdown() {
    Future.doWhile(() async {
      await Future.delayed(const Duration(seconds: 1));
      if (!mounted) return false;
      setState(() => _countdown--);
      return _countdown > 0;
    });
  }

  void _onLogin() {
    setState(() => _isLoading = true);
    // 模拟请求
    Future.delayed(const Duration(seconds: 2), () {
      if (!mounted) return;
      setState(() {
        _isLoading = false;
        // 模拟 50% 失败率用于展示错误态
        if (_code != '123456') {
          _codeError = '验证码错误，请重新输入';
        }
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: DesignTokens.background,
      resizeToAvoidBottomInset: true,
      body: SafeArea(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(
            horizontal: DesignTokens.spacingLg,
          ),
          child: ConstrainedBox(
            constraints: BoxConstraints(
              minHeight: MediaQuery.of(context).size.height -
                  MediaQuery.of(context).padding.top -
                  MediaQuery.of(context).padding.bottom,
            ),
            child: IntrinsicHeight(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  const Spacer(flex: 2),
                  const Text(
                    '欢迎登录',
                    style: TextStyle(
                      fontSize: DesignTokens.fontSizeXxl,
                      fontWeight: DesignTokens.fontWeightBold,
                      color: DesignTokens.textPrimary,
                    ),
                    textAlign: TextAlign.center,
                  ),
                  const SizedBox(height: DesignTokens.spacingSm),
                  const Text(
                    '登录后体验更多功能',
                    style: TextStyle(
                      fontSize: DesignTokens.fontSizeBase,
                      fontWeight: DesignTokens.fontWeightRegular,
                      color: DesignTokens.textSecondary,
                    ),
                    textAlign: TextAlign.center,
                  ),
                  const SizedBox(height: DesignTokens.spacingXxl),
                  PhoneLoginForm(
                    phone: _phone,
                    code: _code,
                    phoneError: _phoneError,
                    codeError: _codeError,
                    countdown: _countdown,
                    isLoading: _isLoading,
                    isSendCodeLoading: _isSendCodeLoading,
                    onPhoneChanged: _onPhoneChanged,
                    onCodeChanged: _onCodeChanged,
                    onSendCode: _onSendCode,
                    onLogin: _onLogin,
                  ),
                  const Spacer(flex: 3),
                  const _BottomHint(),
                  const SizedBox(height: DesignTokens.spacingMd),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}

//  表单占位（设计示意用）
class _FormPlaceholder extends StatelessWidget {
  const _FormPlaceholder();

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        // 手机号输入框占位
        Container(
          height: 56,
          decoration: BoxDecoration(
            color: DesignTokens.surface,
            borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
            border: Border.all(color: DesignTokens.border),
          ),
          padding: const EdgeInsets.symmetric(horizontal: DesignTokens.spacingMd),
          alignment: Alignment.centerLeft,
          child: Text(
            '请输入手机号',
            style: TextStyle(
              fontSize: DesignTokens.fontSizeBase,
              color: DesignTokens.textDisabled,
            ),
          ),
        ),
        const SizedBox(height: DesignTokens.spacingMd),
        Row(
          children: [
            Expanded(
              child: Container(
                height: 56,
                decoration: BoxDecoration(
                  color: DesignTokens.surface,
                  borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
                  border: Border.all(color: DesignTokens.border),
                ),
                padding: const EdgeInsets.symmetric(
                  horizontal: DesignTokens.spacingMd,
                ),
                alignment: Alignment.centerLeft,
                child: Text(
                  '请输入验证码',
                  style: TextStyle(
                    fontSize: DesignTokens.fontSizeBase,
                    color: DesignTokens.textDisabled,
                  ),
                ),
              ),
            ),
            const SizedBox(width: DesignTokens.spacingMd),
            Container(
              width: 120,
              height: 56,
              decoration: BoxDecoration(
                color: DesignTokens.primary,
                borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
              ),
              alignment: Alignment.center,
              child: const Text(
                '获取验证码',
                style: TextStyle(
                  fontSize: DesignTokens.fontSizeSm,
                  fontWeight: DesignTokens.fontWeightMedium,
                  color: DesignTokens.background,
                ),
              ),
            ),
          ],
        ),
        const SizedBox(height: DesignTokens.spacingXl),
        Container(
          height: 48,
          decoration: BoxDecoration(
            color: DesignTokens.primary,
            borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
          ),
          alignment: Alignment.center,
          child: const Text(
            '登录',
            style: TextStyle(
              fontSize: DesignTokens.fontSizeBase,
              fontWeight: DesignTokens.fontWeightMedium,
              color: DesignTokens.background,
            ),
          ),
        ),
      ],
    );
  }
}

//  底部提示文案
class _BottomHint extends StatelessWidget {
  const _BottomHint();

  @override
  Widget build(BuildContext context) {
    return Text(
      '未注册手机号验证后将自动创建账号',
      style: TextStyle(
        fontSize: DesignTokens.fontSizeXs,
        fontWeight: DesignTokens.fontWeightRegular,
        color: DesignTokens.textSecondary.withValues(alpha: 0.7),
      ),
      textAlign: TextAlign.center,
    );
  }
}
