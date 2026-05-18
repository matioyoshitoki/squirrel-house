/// 认证模块新增组件设计原型
///
/// 包含：
/// - ToTextField: 带错误态的文本输入框
/// - ToCountdownButton: 倒计时按钮
///
/// 设计约束：
/// - 所有颜色、间距、圆角使用 DesignTokens
/// - 支持 TalkBack (Semantics)
/// - 支持 E2E 测试 (identifier)

import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

// 内联 Token 定义（设计资产独立可预览）
class _Tokens {
  _Tokens._();
  static const Color primary = Color(0xFF1976D2);
  static const Color surface = Color(0xFFF5F5F5);
  static const Color error = Color(0xFFD32F2F);
  static const Color textPrimary = Color(0xFF212121);
  static const Color textSecondary = Color(0xFF757575);
  static const Color textDisabled = Color(0xFF9E9E9E);
  static const Color border = Color(0xFFE0E0E0);
  static const Color background = Color(0xFFFFFFFF);

  static const double fontSizeSm = 14;
  static const double fontSizeBase = 16;
  static const double spacingXs = 4;
  static const double spacingSm = 8;
  static const double spacingMd = 16;
  static const double radiusMd = 8;

  static const FontWeight fontWeightMedium = FontWeight.w500;
}

// ───────────────────────────────────────────────
// ToTextField — 带错误态的文本输入框（原子组件）
// ───────────────────────────────────────────────

class ToTextField extends StatelessWidget {
  final String hint;
  final String? errorText;
  final Widget? prefix;
  final TextInputType keyboardType;
  final int? maxLength;
  final bool isEnabled;
  final ValueChanged<String>? onChanged;
  final TextEditingController? controller;
  final String semanticsIdentifier;
  final List<TextInputFormatter>? inputFormatters;
  final FocusNode? focusNode;

  const ToTextField({
    super.key,
    this.hint = '',
    this.errorText,
    this.prefix,
    this.keyboardType = TextInputType.text,
    this.maxLength,
    this.isEnabled = true,
    this.onChanged,
    this.controller,
    this.semanticsIdentifier = '',
    this.inputFormatters,
    this.focusNode,
  });

  bool get _hasError => errorText != null && errorText!.isNotEmpty;

  @override
  Widget build(BuildContext context) {
    return Semantics(
      identifier: semanticsIdentifier.isNotEmpty ? semanticsIdentifier : null,
      label: hint,
      textField: true,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisSize: MainAxisSize.min,
        children: [
          AnimatedContainer(
            duration: const Duration(milliseconds: 150),
            decoration: BoxDecoration(
              color: isEnabled
                  ? _Tokens.surface
                  : _Tokens.surface.withOpacity(0.5),
              borderRadius: BorderRadius.circular(_Tokens.radiusMd),
              border: Border.all(
                color: _hasError ? _Tokens.error : _Tokens.border,
                width: 1,
              ),
            ),
            padding: const EdgeInsets.symmetric(
              horizontal: _Tokens.spacingMd,
              vertical: _Tokens.spacingMd,
            ),
            child: Row(
              children: [
                if (prefix != null) ...[
                  DefaultTextStyle(
                    style: const TextStyle(
                      fontSize: _Tokens.fontSizeBase,
                      color: _Tokens.textSecondary,
                      fontWeight: _Tokens.fontWeightMedium,
                    ),
                    child: prefix!,
                  ),
                  const SizedBox(width: _Tokens.spacingSm),
                ],
                Expanded(
                  child: TextField(
                    controller: controller,
                    focusNode: focusNode,
                    enabled: isEnabled,
                    keyboardType: keyboardType,
                    maxLength: maxLength,
                    inputFormatters: inputFormatters,
                    onChanged: onChanged,
                    style: const TextStyle(
                      fontSize: _Tokens.fontSizeBase,
                      color: _Tokens.textPrimary,
                    ),
                    decoration: InputDecoration(
                      isCollapsed: true,
                      border: InputBorder.none,
                      hintText: hint,
                      hintStyle: const TextStyle(
                        fontSize: _Tokens.fontSizeBase,
                        color: _Tokens.textDisabled,
                      ),
                      counterText: '',
                    ),
                  ),
                ),
              ],
            ),
          ),
          AnimatedSwitcher(
            duration: const Duration(milliseconds: 200),
            child: _hasError
                ? Padding(
                    key: ValueKey(errorText),
                    padding: const EdgeInsets.only(
                      top: _Tokens.spacingXs,
                      left: _Tokens.spacingSm,
                    ),
                    child: Text(
                      errorText!,
                      style: const TextStyle(
                        fontSize: _Tokens.fontSizeSm,
                        color: _Tokens.error,
                      ),
                    ),
                  )
                : const SizedBox.shrink(key: Key('no_error')),
          ),
        ],
      ),
    );
  }
}

// ───────────────────────────────────────────────
// ToCountdownButton — 倒计时按钮（分子组件）
// ───────────────────────────────────────────────

class ToCountdownButton extends StatefulWidget {
  final String label;
  final String countdownLabel;
  final int duration;
  final VoidCallback? onPressed;
  final bool isDisabled;

  const ToCountdownButton({
    super.key,
    this.label = '获取验证码',
    this.countdownLabel = '{s}s',
    this.duration = 60,
    this.onPressed,
    this.isDisabled = false,
  });

  @override
  State<ToCountdownButton> createState() => _ToCountdownButtonState();
}

class _ToCountdownButtonState extends State<ToCountdownButton> {
  Timer? _timer;
  int _remaining = 0;

  bool get _isRunning => _remaining > 0;
  bool get _isEffectiveDisabled => widget.isDisabled || _isRunning;

  String get _displayLabel {
    if (_isRunning) {
      return widget.countdownLabel.replaceFirst('{s}', '$_remaining');
    }
    return widget.label;
  }

  void _startCountdown() {
    if (_isRunning) return;
    setState(() => _remaining = widget.duration);
    _timer?.cancel();
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      if (!mounted) {
        timer.cancel();
        return;
      }
      setState(() {
        _remaining--;
        if (_remaining <= 0) {
          timer.cancel();
        }
      });
    });
  }

  void _handleTap() {
    if (_isEffectiveDisabled || widget.onPressed == null) return;
    _startCountdown();
    widget.onPressed!();
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Semantics(
      button: true,
      label: _displayLabel,
      child: GestureDetector(
        onTapDown: _isEffectiveDisabled ? null : (_) => _handleTap(),
        child: AnimatedContainer(
          duration: const Duration(milliseconds: 150),
          padding: const EdgeInsets.symmetric(
            horizontal: _Tokens.spacingMd,
            vertical: _Tokens.spacingMd,
          ),
          decoration: BoxDecoration(
            color: _isEffectiveDisabled
                ? _Tokens.primary.withOpacity(0.4)
                : _Tokens.primary,
            borderRadius: BorderRadius.circular(_Tokens.radiusMd),
          ),
          child: Text(
            _displayLabel,
            style: const TextStyle(
              fontSize: _Tokens.fontSizeSm,
              fontWeight: _Tokens.fontWeightMedium,
              color: _Tokens.background,
            ),
            textAlign: TextAlign.center,
          ),
        ),
      ),
    );
  }
}
