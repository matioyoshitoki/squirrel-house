import 'package:flutter/material.dart';
import '../_tokens.dart';

/// 消息已读回执指示器
///
/// 展示单条消息的送达/已读状态：
/// - [sending] 发送中：旋转的 CircularProgressIndicator
/// - [sent] 已发送（服务器已接收）：单勾
/// - [delivered] 已送达（对方客户端已收到）：双勾（灰色）
/// - [read] 已读（对方已打开会话）：双勾（主色）
enum ReadReceiptStatus { sending, sent, delivered, read }

class ReadReceiptIndicator extends StatelessWidget {
  final ReadReceiptStatus status;
  final double size;

  const ReadReceiptIndicator({
    super.key,
    required this.status,
    this.size = 16,
  });

  @override
  Widget build(BuildContext context) {
    return Semantics(
      identifier: 'read_receipt_${status.name}',
      label: _label,
      container: true,
      child: SizedBox(
        width: size,
        height: size,
        child: switch (status) {
          ReadReceiptStatus.sending => Center(
              child: SizedBox(
                width: size * 0.75,
                height: size * 0.75,
                child: CircularProgressIndicator(
                  strokeWidth: 2,
                  color: Tokens.readReceiptSent,
                ),
              ),
            ),
          ReadReceiptStatus.sent => Icon(
              Icons.check,
              size: size,
              color: Tokens.readReceiptSent,
            ),
          ReadReceiptStatus.delivered => _DoubleCheck(
              size: size,
              color: Tokens.readReceiptDelivered,
            ),
          ReadReceiptStatus.read => _DoubleCheck(
              size: size,
              color: Tokens.readReceiptRead,
            ),
        },
      ),
    );
  }

  String get _label => switch (status) {
        ReadReceiptStatus.sending => '发送中',
        ReadReceiptStatus.sent => '已发送',
        ReadReceiptStatus.delivered => '已送达',
        ReadReceiptStatus.read => '已读',
      };
}

class _DoubleCheck extends StatelessWidget {
  final double size;
  final Color color;

  const _DoubleCheck({required this.size, required this.color});

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Positioned(
          left: 0,
          top: 0,
          width: size,
          height: size,
          child: Icon(Icons.check, size: size, color: color),
        ),
        Positioned(
          left: size * 0.25,
          top: 0,
          width: size,
          height: size,
          child: Icon(Icons.check, size: size, color: color),
        ),
      ],
    );
  }
}
