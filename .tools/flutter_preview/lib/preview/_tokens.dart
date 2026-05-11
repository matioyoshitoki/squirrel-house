import 'package:flutter/material.dart';

/// Temporary tokens for Issue #5 IM Design-Only PR.
/// These must be migrated to `packages/design-system/lib/src/theme/tokens.dart`
/// in the subsequent feat PR.
class Tokens {
  // Bubble colors
  static const Color bubbleSelfBackground = Color(0xFF6750A4);
  static const Color bubbleOtherBackground = Color(0xFFF3EDF7);
  static const Color bubbleSelfText = Color(0xFFFFFFFF);
  static const Color bubbleOtherText = Color(0xFF1C1B1F);
  static const Color systemMessageText = Color(0xFF79747E);
  static const Color timestampText = Color(0xFF79747E);
  static const Color readReceiptText = Color(0xFF6750A4);

  // Input
  static const Color inputBarBackground = Color(0xFFFFFFFF);
  static const Color divider = Color(0xFFEEEEEE);

  // Badge
  static const Color unreadBadgeBackground = Color(0xFFB3261E);
  static const Color unreadBadgeText = Color(0xFFFFFFFF);

  // Read receipt
  static const Color readReceiptSent = Color(0xFF958DA5);
  static const Color readReceiptDelivered = Color(0xFF958DA5);
  static const Color readReceiptRead = Color(0xFF6750A4);

  // System message
  static const Color systemMessageBg = Color(0xFFF3EDF7);

  // Typography supplement
  static const double fontSizeSmall = 14.0;
}
