import 'package:flutter/material.dart';
import '../_tokens.dart';

/// 聊天室系统消息
///
/// 居中显示的系统提示，如 "你们已匹配成功，开始聊天吧"
class SystemMessage extends StatelessWidget {
  final String text;

  const SystemMessage({
    super.key,
    required this.text,
  });

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        margin: const EdgeInsets.symmetric(vertical: 8),
        padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 6),
        decoration: BoxDecoration(
          color: Tokens.systemMessageBg,
          borderRadius: BorderRadius.circular(12),
        ),
        child: Text(
          text,
          style: const TextStyle(
            fontSize: 12,
            color: Tokens.systemMessageText,
            fontWeight: FontWeight.w500,
          ),
        ),
      ),
    );
  }
}
