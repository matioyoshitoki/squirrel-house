//  PhoneLoginForm 有机体组件设计
// 
//  手机号验证码登录表单，包含手机号输入、验证码输入、获取验证码按钮、登录按钮。
//  纯 UI 展示，不含业务逻辑（无 API 调用、无状态管理实现）。
// 
//  使用位置：LoginPage

import 'package:flutter/material.dart';
import 'package:type_one/components/atoms/to_button.dart';
import 'package:type_one/theme/tokens.dart';

import 'to_text_field_design.dart';

//  手机号验证码登录表单
// 
//  设计参数说明：
//  - [phone] / [code]：当前输入值（用于控制按钮启用状态）
//  - [phoneError] / [codeError]：错误提示文案
//  - [countdown]：验证码倒计时剩余秒数，>0 表示倒计时中
//  - [isLoading]：登录按钮加载态
//  - [isSendCodeLoading]：获取验证码按钮加载态
//  - [onPhoneChanged] / [onCodeChanged] / [onSendCode] / [onLogin]：交互回调
class PhoneLoginForm extends StatelessWidget {
  final String? phone;
  final String? code;
  final String? phoneError;
  final String? codeError;
  final int countdown;
  final bool isLoading;
  final bool isSendCodeLoading;
  final ValueChanged<String>? onPhoneChanged;
  final ValueChanged<String>? onCodeChanged;
  final VoidCallback? onSendCode;
  final VoidCallback? onLogin;

  const PhoneLoginForm({
    super.key,
    this.phone,
    this.code,
    this.phoneError,
    this.codeError,
    this.countdown = 0,
    this.isLoading = false,
    this.isSendCodeLoading = false,
    this.onPhoneChanged,
    this.onCodeChanged,
    this.onSendCode,
    this.onLogin,
  });

  bool get _canSendCode =>
      countdown == 0 && !isSendCodeLoading && !isLoading;

  bool get _canLogin =>
      (phone?.isNotEmpty ?? false) &&
      (code?.isNotEmpty ?? false) &&
      !isLoading &&
      !isSendCodeLoading;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        // ── 手机号输入 ──
        ToTextField(
          semanticsIdentifier: 'phone_input',
          hintText: '请输入手机号',
          keyboardType: TextInputType.phone,
          autofocus: true,
          maxLength: 11,
          errorText: phoneError,
          prefix: const _PhonePrefix(),
          onChanged: onPhoneChanged,
          textInputAction: TextInputAction.next,
        ),
        const SizedBox(height: DesignTokens.spacingMd),

        // ── 验证码输入 + 获取验证码按钮 ──
        Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Expanded(
              child: ToTextField(
                semanticsIdentifier: 'code_input',
                hintText: '请输入验证码',
                keyboardType: TextInputType.number,
                maxLength: 6,
                errorText: codeError,
                onChanged: onCodeChanged,
                textInputAction: TextInputAction.done,
              ),
            ),
            const SizedBox(width: DesignTokens.spacingMd),
            _CountdownButton(
              countdown: countdown,
              isLoading: isSendCodeLoading,
              isEnabled: _canSendCode,
              onPressed: onSendCode,
            ),
          ],
        ),
        const SizedBox(height: DesignTokens.spacingXl),

        // ── 登录按钮 ──
        ToButton(
          label: '登录',
          isLoading: isLoading,
          isDisabled: !_canLogin,
          onPressed: _canLogin ? onLogin : null,
        ),
      ],
    );
  }
}

//  手机号前缀（+86 / 国旗图标）
// 
//  设计说明：首期简化，仅展示 +86 文本，后续可扩展为国家选择器。
class _PhonePrefix extends StatelessWidget {
  const _PhonePrefix();

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        Text(
          '+86',
          style: TextStyle(
            fontSize: DesignTokens.fontSizeBase,
            fontWeight: DesignTokens.fontWeightMedium,
            color: DesignTokens.textPrimary.withValues(alpha: 0.8),
          ),
        ),
        const SizedBox(width: DesignTokens.spacingSm),
        Container(
          width: 1,
          height: 20,
          color: DesignTokens.border,
        ),
        const SizedBox(width: DesignTokens.spacingSm),
      ],
    );
  }
}

//  获取验证码按钮（带倒计时）
// 
//  固定宽度 120dp，高度与输入框一致（56dp 内容区）。
class _CountdownButton extends StatelessWidget {
  final int countdown;
  final bool isLoading;
  final bool isEnabled;
  final VoidCallback? onPressed;

  const _CountdownButton({
    required this.countdown,
    required this.isLoading,
    required this.isEnabled,
    this.onPressed,
  });

  @override
  Widget build(BuildContext context) {
    final isCounting = countdown > 0;
    final isDisabled = !isEnabled || isCounting || isLoading;

    return Semantics(
      identifier: 'send_code_button',
      label: isCounting ?  '$countdown秒后重新获取' : '获取验证码',
      button: true,
      child: GestureDetector(
        onTapDown: isDisabled ? null : (_) => onPressed?.call(),
        child: AnimatedContainer(
          duration: const Duration(milliseconds: 150),
          width: 120,
          height: 56,
          alignment: Alignment.center,
          decoration: BoxDecoration(
            color: isCounting
                ? DesignTokens.surface
                : isDisabled
                    ? DesignTokens.primary.withValues(alpha: 0.5)
                    : DesignTokens.primary,
            borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
          ),
          child: isLoading
              ? const SizedBox(
                  width: 20,
                  height: 20,
                  child: CircularProgressIndicator(
                    strokeWidth: 2,
                    valueColor:
                        AlwaysStoppedAnimation<Color>(DesignTokens.background),
                  ),
                )
              : Text(
                  isCounting ?  '$countdown s' : '获取验证码',
                  style: TextStyle(
                    fontSize: DesignTokens.fontSizeSm,
                    fontWeight: DesignTokens.fontWeightMedium,
                    color: isCounting
                        ? DesignTokens.textDisabled
                        : DesignTokens.background,
                  ),
                ),
        ),
      ),
    );
  }
}
