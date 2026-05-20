/// LoginPage 设计原型
///
/// 展示所有交互状态：
/// - 初始态
/// - 手机号输入中
/// - 倒计时中
/// - 可登录
/// - 登录加载中
/// - 验证码错误
/// - 手机号格式错误
/// - 频繁请求错误
///
/// 设计约束：
/// - 全部使用 DesignTokens
/// - 支持 SingleChildScrollView（键盘弹起适配）
/// - 支持安全区域

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'auth_components.dart';

// 内联 Token 定义（设计资产独立可预览）
class _Tokens {
  _Tokens._();
  static const Color primary = Color(0xFF1976D2);
  static const Color secondary = Color(0xFF9C27B0);
  static const Color background = Color(0xFFFFFFFF);
  static const Color textPrimary = Color(0xFF212121);
  static const Color textSecondary = Color(0xFF757575);
  static const Color error = Color(0xFFD32F2F);
  static const Color surface = Color(0xFFF5F5F5);
  static const Color border = Color(0xFFE0E0E0);

  static const double fontSizeBase = 16;
  static const double fontSizeLg = 18;
  static const double fontSizeXl = 24;
  static const double fontSizeXxl = 32;
  static const double fontSizeSm = 14;

  static const double spacingXs = 4;
  static const double spacingSm = 8;
  static const double spacingMd = 16;
  static const double spacingLg = 24;
  static const double spacingXl = 32;
  static const double spacingXxl = 48;

  static const double radiusMd = 8;

  static const FontWeight fontWeightBold = FontWeight.w700;
  static const FontWeight fontWeightMedium = FontWeight.w500;
}

// ───────────────────────────────────────────────
// LoginPage 设计原型
// ───────────────────────────────────────────────

class LoginPageDesign extends StatefulWidget {
  const LoginPageDesign({super.key});

  @override
  State<LoginPageDesign> createState() => _LoginPageDesignState();
}

class _LoginPageDesignState extends State<LoginPageDesign> {
  String _phone = '';
  String _code = '';
  bool _isCountingDown = false;
  bool _isLoading = false;
  String? _phoneError;
  String? _codeError;
  String? _globalError;

  bool get _isPhoneValid => _phone.length == 11;
  bool get _isCodeValid => _code.length == 6;
  bool get _canLogin => _isPhoneValid && _isCodeValid && !_isLoading;

  void _onPhoneChanged(String value) {
    setState(() {
      _phone = value;
      _phoneError = null;
      _globalError = null;
    });
  }

  void _onCodeChanged(String value) {
    setState(() {
      _code = value;
      _codeError = null;
      _globalError = null;
    });
  }

  void _onSendCode() {
    if (!_isPhoneValid) {
      setState(() => _phoneError = '请输入正确的手机号');
      return;
    }
    setState(() => _isCountingDown = true);
  }

  void _onLogin() {
    if (!_canLogin) return;
    setState(() => _isLoading = true);
    // 模拟登录结果
    Future.delayed(const Duration(seconds: 2), () {
      if (!mounted) return;
      setState(() {
        _isLoading = false;
        _codeError = '验证码错误，请重新输入';
      });
    });
  }

  void _setScenario(String scenario) {
    setState(() {
      _phone = '';
      _code = '';
      _isCountingDown = false;
      _isLoading = false;
      _phoneError = null;
      _codeError = null;
      _globalError = null;

      switch (scenario) {
        case 'initial':
          break;
        case 'phone_input':
          _phone = '13800138000';
          break;
        case 'countdown':
          _phone = '13800138000';
          _isCountingDown = true;
          break;
        case 'ready':
          _phone = '13800138000';
          _code = '123456';
          break;
        case 'loading':
          _phone = '13800138000';
          _code = '123456';
          _isLoading = true;
          break;
        case 'code_error':
          _phone = '13800138000';
          _code = '123456';
          _codeError = '验证码错误，请重新输入';
          break;
        case 'phone_error':
          _phone = '123';
          _phoneError = '请输入正确的手机号';
          break;
        case 'frequent':
          _phone = '13800138000';
          _globalError = '操作太频繁，请稍后再试';
          break;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Stack(
          children: [
            SingleChildScrollView(
              padding: const EdgeInsets.symmetric(
                horizontal: _Tokens.spacingLg,
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
                      const SizedBox(height: _Tokens.spacingXxl),

                      // 标题
                      const Text(
                        '欢迎登录',
                        style: TextStyle(
                          fontSize: _Tokens.fontSizeXxl,
                          fontWeight: _Tokens.fontWeightBold,
                          color: _Tokens.textPrimary,
                        ),
                        textAlign: TextAlign.start,
                      ),
                      const SizedBox(height: _Tokens.spacingSm),

                      // 副标题
                      const Text(
                        '请输入手机号登录',
                        style: TextStyle(
                          fontSize: _Tokens.fontSizeBase,
                          color: _Tokens.textSecondary,
                        ),
                      ),
                      const SizedBox(height: _Tokens.spacingXxl),

                      // 手机号输入框
                      ToTextField(
                        hint: '请输入手机号',
                        semanticsIdentifier: 'phone_input',
                        keyboardType: TextInputType.phone,
                        maxLength: 11,
                        isEnabled: !_isLoading,
                        errorText: _phoneError,
                        prefix: const Text('+86'),
                        inputFormatters: [
                          FilteringTextInputFormatter.digitsOnly,
                        ],
                        onChanged: _onPhoneChanged,
                      ),
                      const SizedBox(height: _Tokens.spacingLg),

                      // 验证码行
                      Row(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          // 验证码输入框
                          Expanded(
                            flex: 3,
                            child: ToTextField(
                              hint: '请输入验证码',
                              semanticsIdentifier: 'code_input',
                              keyboardType: TextInputType.number,
                              maxLength: 6,
                              isEnabled: !_isLoading && _isPhoneValid,
                              errorText: _codeError,
                              inputFormatters: [
                                FilteringTextInputFormatter.digitsOnly,
                              ],
                              onChanged: _onCodeChanged,
                            ),
                          ),
                          const SizedBox(width: _Tokens.spacingMd),

                          // 获取验证码按钮
                          SizedBox(
                            width: 110,
                            child: ToCountdownButton(
                              label: '获取验证码',
                              duration: 60,
                              isDisabled: !_isPhoneValid || _isLoading,
                              onPressed: _onSendCode,
                            ),
                          ),
                        ],
                      ),
                      const SizedBox(height: _Tokens.spacingMd),

                      // 全局错误提示（如频繁请求）
                      AnimatedSwitcher(
                        duration: const Duration(milliseconds: 200),
                        child: _globalError != null
                            ? Container(
                                key: ValueKey(_globalError),
                                padding: const EdgeInsets.symmetric(
                                  horizontal: _Tokens.spacingMd,
                                  vertical: _Tokens.spacingSm,
                                ),
                                decoration: BoxDecoration(
                                  color: _Tokens.error.withOpacity(0.08),
                                  borderRadius: BorderRadius.circular(
                                    _Tokens.radiusMd,
                                  ),
                                ),
                                child: Row(
                                  children: [
                                    const Icon(
                                      Icons.error_outline,
                                      size: 16,
                                      color: _Tokens.error,
                                    ),
                                    const SizedBox(width: 8),
                                    Expanded(
                                      child: Text(
                                        _globalError!,
                                        style: const TextStyle(
                                          fontSize: _Tokens.fontSizeSm,
                                          color: _Tokens.error,
                                        ),
                                      ),
                                    ),
                                  ],
                                ),
                              )
                            : const SizedBox.shrink(
                                key: Key('no_global_error'),
                              ),
                      ),

                      const Spacer(),

                      // 登录按钮
                      GestureDetector(
                        onTapDown: _canLogin ? (_) => _onLogin() : null,
                        child: AnimatedContainer(
                          duration: const Duration(milliseconds: 150),
                          padding: const EdgeInsets.symmetric(
                            vertical: _Tokens.spacingMd,
                          ),
                          decoration: BoxDecoration(
                            color: _canLogin
                                ? _Tokens.primary
                                : _Tokens.primary.withOpacity(0.4),
                            borderRadius: BorderRadius.circular(
                              _Tokens.radiusMd,
                            ),
                          ),
                          child: _isLoading
                              ? const SizedBox(
                                  height: 20,
                                  child: Center(
                                    child: SizedBox(
                                      width: 20,
                                      height: 20,
                                      child: CircularProgressIndicator(
                                        strokeWidth: 2,
                                        valueColor:
                                            AlwaysStoppedAnimation<Color>(
                                          _Tokens.background,
                                        ),
                                      ),
                                    ),
                                  ),
                                )
                              : const Text(
                                  '登录',
                                  style: TextStyle(
                                    fontSize: _Tokens.fontSizeLg,
                                    fontWeight: _Tokens.fontWeightMedium,
                                    color: _Tokens.background,
                                  ),
                                  textAlign: TextAlign.center,
                                ),
                        ),
                      ),
                      const SizedBox(height: _Tokens.spacingLg),

                      // 底部留白（键盘适配）
                      const SizedBox(height: _Tokens.spacingXl),
                    ],
                  ),
                ),
              ),
            ),

            // 场景切换浮层（仅设计预览用）
            Positioned(
              top: MediaQuery.of(context).padding.top + _Tokens.spacingMd,
              right: _Tokens.spacingMd,
              child: _ScenarioSwitcher(
                current: _currentScenario,
                onChanged: _setScenario,
              ),
            ),
          ],
        ),
      ),
    );
  }

  String get _currentScenario {
    if (_isLoading) return 'loading';
    if (_globalError != null) return 'frequent';
    if (_phoneError != null) return 'phone_error';
    if (_codeError != null) return 'code_error';
    if (_isCountingDown && _code.isEmpty) return 'countdown';
    if (_code.isNotEmpty) return 'ready';
    if (_phone.isNotEmpty) return 'phone_input';
    return 'initial';
  }
}

// 场景切换器（设计预览调试组件）
class _ScenarioSwitcher extends StatelessWidget {
  final String current;
  final ValueChanged<String> onChanged;

  const _ScenarioSwitcher({required this.current, required this.onChanged});

  @override
  Widget build(BuildContext context) {
    final scenarios = [
      'initial',
      'phone_input',
      'countdown',
      'ready',
      'loading',
      'code_error',
      'phone_error',
      'frequent',
    ];

    return Container(
      decoration: BoxDecoration(
        color: _Tokens.background.withOpacity(0.95),
        borderRadius: BorderRadius.circular(_Tokens.radiusMd),
        border: Border.all(color: _Tokens.border),
        boxShadow: const [
          BoxShadow(
            color: Color(0x1A000000),
            blurRadius: 6,
            offset: Offset(0, 2),
          ),
        ],
      ),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: scenarios.map((s) {
          final isActive = s == current;
          return GestureDetector(
            onTap: () => onChanged(s),
            child: Container(
              padding: const EdgeInsets.symmetric(
                horizontal: 12,
                vertical: 8,
              ),
              decoration: BoxDecoration(
                color: isActive ? _Tokens.primary.withOpacity(0.1) : null,
                borderRadius: BorderRadius.circular(_Tokens.radiusMd),
              ),
              child: Text(
                s,
                style: TextStyle(
                  fontSize: 11,
                  color: isActive ? _Tokens.primary : _Tokens.textSecondary,
                  fontWeight:
                      isActive ? _Tokens.fontWeightMedium : FontWeight.normal,
                ),
              ),
            ),
          );
        }).toList(),
      ),
    );
  }
}
