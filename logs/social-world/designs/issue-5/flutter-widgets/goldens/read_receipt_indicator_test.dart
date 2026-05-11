import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import '../atoms/read_receipt_indicator.dart';

void main() {
  group('ReadReceiptIndicator Golden', () {
    testWidgets('sending', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: Center(
              child: ReadReceiptIndicator(status: ReadReceiptStatus.sending),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ReadReceiptIndicator),
        matchesGoldenFile('read_receipt_sending.png'),
      );
    });

    testWidgets('sent', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: Center(
              child: ReadReceiptIndicator(status: ReadReceiptStatus.sent),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ReadReceiptIndicator),
        matchesGoldenFile('read_receipt_sent.png'),
      );
    });

    testWidgets('delivered', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: Center(
              child: ReadReceiptIndicator(status: ReadReceiptStatus.delivered),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ReadReceiptIndicator),
        matchesGoldenFile('read_receipt_delivered.png'),
      );
    });

    testWidgets('read', (tester) async {
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: Center(
              child: ReadReceiptIndicator(status: ReadReceiptStatus.read),
            ),
          ),
        ),
      );
      await expectLater(
        find.byType(ReadReceiptIndicator),
        matchesGoldenFile('read_receipt_read.png'),
      );
    });
  });
}
