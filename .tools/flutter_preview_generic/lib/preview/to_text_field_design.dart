//  ToTextField 原子组件设计
// 
//  统一文本输入框，支持前缀、错误态、禁用态、清除按钮。
//  所有视觉属性来自 DesignTokens，零硬编码。
// 
//  使用位置：PhoneLoginForm（手机号输入、验证码输入）

import 'package:flutter/material.dart';
import 'package:type_one/theme/tokens.dart';

//  文本输入原子组件
// 
//  设计系统标准输入框。禁止在业务代码中直接使用裸 [TextField]。
class ToTextField extends StatefulWidget {
  final TextEditingController? controller;
  final String? hintText;
  final String? errorText;
  final TextInputType keyboardType;
  final bool autofocus;
  final bool obscureText;
  final int? maxLength;
  final Widget? prefix;
  final ValueChanged<String>? onChanged;
  final VoidCallback? onClear;
  final String? semanticsIdentifier;
  final FocusNode? focusNode;
  final TextInputAction? textInputAction;
  final VoidCallback? onEditingComplete;

  const ToTextField({
    super.key,
    this.controller,
    this.hintText,
    this.errorText,
    this.keyboardType = TextInputType.text,
    this.autofocus = false,
    this.obscureText = false,
    this.maxLength,
    this.prefix,
    this.onChanged,
    this.onClear,
    this.semanticsIdentifier,
    this.focusNode,
    this.textInputAction,
    this.onEditingComplete,
  });

  @override
  State<ToTextField> createState() => _ToTextFieldState();
}

class _ToTextFieldState extends State<ToTextField> {
  late final FocusNode _internalFocusNode;
  FocusNode get _effectiveFocusNode => widget.focusNode ?? _internalFocusNode;
  bool _hasFocus = false;

  @override
  void initState() {
    super.initState();
    _internalFocusNode = FocusNode();
    _effectiveFocusNode.addListener(_onFocusChange);
  }

  @override
  void dispose() {
    _effectiveFocusNode.removeListener(_onFocusChange);
    _internalFocusNode.dispose();
    super.dispose();
  }

  void _onFocusChange() {
    setState(() => _hasFocus = _effectiveFocusNode.hasFocus);
  }

  @override
  Widget build(BuildContext context) {
    final hasError = widget.errorText != null && widget.errorText!.isNotEmpty;
    final borderColor = hasError
        ? DesignTokens.error
        : _hasFocus
            ? DesignTokens.primary
            : DesignTokens.border;

    return Semantics(
      identifier: widget.semanticsIdentifier,
      label:
          '${widget.hintText ?? "文本输入框"}${hasError ? '，${widget.errorText}' : ''}',
      textField: true,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisSize: MainAxisSize.min,
        children: [
          AnimatedContainer(
            duration: const Duration(milliseconds: 150),
            decoration: BoxDecoration(
              color: DesignTokens.surface,
              borderRadius: BorderRadius.circular(DesignTokens.radiusMd),
              border: Border.all(color: borderColor, width: 1),
            ),
            child: TextField(
              controller: widget.controller,
              focusNode: _effectiveFocusNode,
              keyboardType: widget.keyboardType,
              autofocus: widget.autofocus,
              obscureText: widget.obscureText,
              maxLength: widget.maxLength,
              textInputAction: widget.textInputAction,
              onChanged: widget.onChanged,
              onEditingComplete: widget.onEditingComplete,
              style: const TextStyle(
                fontSize: DesignTokens.fontSizeBase,
                fontWeight: DesignTokens.fontWeightRegular,
                color: DesignTokens.textPrimary,
              ),
              decoration: InputDecoration(
                hintText: widget.hintText,
                hintStyle: const TextStyle(
                  fontSize: DesignTokens.fontSizeBase,
                  fontWeight: DesignTokens.fontWeightRegular,
                  color: DesignTokens.textDisabled,
                ),
                prefixIcon: widget.prefix != null
                    ? Padding(
                        padding: const EdgeInsets.only(
                          left: DesignTokens.spacingMd,
                          right: DesignTokens.spacingSm,
                        ),
                        child: widget.prefix,
                      )
                    : null,
                prefixIconConstraints: widget.prefix != null
                    ? const BoxConstraints(
                        minWidth: 0,
                        minHeight: 0,
                      )
                    : null,
                suffixIcon: _ClearButton(
                  controller: widget.controller,
                  onClear: widget.onClear,
                ),
                border: InputBorder.none,
                contentPadding: const EdgeInsets.symmetric(
                  horizontal: DesignTokens.spacingMd,
                  vertical: DesignTokens.spacingMd,
                ),
                counterText: '',
              ),
            ),
          ),
          if (hasError) ...[
            const SizedBox(height: DesignTokens.spacingXs),
            Text(
              widget.errorText!,
              style: const TextStyle(
                fontSize: DesignTokens.fontSizeSm,
                fontWeight: DesignTokens.fontWeightRegular,
                color: DesignTokens.error,
              ),
            ),
          ],
        ],
      ),
    );
  }
}

//  清除按钮（内部组件）
class _ClearButton extends StatefulWidget {
  final TextEditingController? controller;
  final VoidCallback? onClear;

  const _ClearButton({this.controller, this.onClear});

  @override
  State<_ClearButton> createState() => _ClearButtonState();
}

class _ClearButtonState extends State<_ClearButton> {
  late final TextEditingController _fallbackController;
  TextEditingController get _effectiveController =>
      widget.controller ?? _fallbackController;

  @override
  void initState() {
    super.initState();
    _fallbackController = TextEditingController();
    _effectiveController.addListener(_onTextChanged);
  }

  @override
  void dispose() {
    _effectiveController.removeListener(_onTextChanged);
    _fallbackController.dispose();
    super.dispose();
  }

  void _onTextChanged() => setState(() {});

  @override
  Widget build(BuildContext context) {
    final hasText = _effectiveController.text.isNotEmpty;
    if (!hasText) return const SizedBox.shrink();

    return Semantics(
      identifier: widget.controller != null
          ? '${widget.controller.hashCode}_clear'
          : 'clear_button',
      label: '清除',
      button: true,
      child: GestureDetector(
        onTap: () {
          _effectiveController.clear();
          widget.onClear?.call();
        },
        child: Container(
          width: 40,
          height: 40,
          alignment: Alignment.center,
          child: const Icon(
            Icons.close,
            size: 18,
            color: DesignTokens.textSecondary,
          ),
        ),
      ),
    );
  }
}
