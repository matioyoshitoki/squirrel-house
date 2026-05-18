import 'package:flutter/material.dart';
import 'preview/auth_components.dart';
import 'preview/login_page_design.dart';
import 'preview/phone_login_form_design.dart';
import 'preview/splash_page_design.dart';

void main() {
  runApp(const DesignPreviewApp());
}

class DesignPreviewApp extends StatelessWidget {
  const DesignPreviewApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Design Preview',
      debugShowCheckedModeBanner: false,
      theme: ThemeData.dark(),
      home: Container(
        color: const Color(0xFF0f172a),
        padding: const EdgeInsets.all(16),
        child: ListView(
          children: [
            const Center(
              child: Padding(
                padding: EdgeInsets.symmetric(vertical: 16),
                child: Text('Design Assets Preview', style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
              ),
            ),
            _buildCard('ToTextField', Container(height: 400, decoration: BoxDecoration(borderRadius: BorderRadius.circular(8), border: Border.all(color: Colors.grey)), clipBehavior: Clip.antiAlias, child: ToTextField())),
            _buildCard('ToCountdownButton', ToCountdownButton()),
            _buildCard('LoginPageDefault', Container(height: 400, decoration: BoxDecoration(borderRadius: BorderRadius.circular(8), border: Border.all(color: Colors.grey)), clipBehavior: Clip.antiAlias, child: LoginPageDefault())),
            _buildCard('LoginPageInteractivePreview', Container(height: 400, decoration: BoxDecoration(borderRadius: BorderRadius.circular(8), border: Border.all(color: Colors.grey)), clipBehavior: Clip.antiAlias, child: LoginPageInteractivePreview())),
            _buildCard('PhoneLoginForm', Container(height: 400, decoration: BoxDecoration(borderRadius: BorderRadius.circular(8), border: Border.all(color: Colors.grey)), clipBehavior: Clip.antiAlias, child: PhoneLoginForm())),
            _buildCard('SplashPageChecking', Container(height: 400, decoration: BoxDecoration(borderRadius: BorderRadius.circular(8), border: Border.all(color: Colors.grey)), clipBehavior: Clip.antiAlias, child: SplashPageChecking())),
            _buildCard('SplashPageError', Container(height: 400, decoration: BoxDecoration(borderRadius: BorderRadius.circular(8), border: Border.all(color: Colors.grey)), clipBehavior: Clip.antiAlias, child: SplashPageError())),
          ],
        ),
      ),
    );
  }

  Widget _buildCard(String title, Widget child) {
    return Card(
      margin: const EdgeInsets.only(bottom: 16),
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(title, style: const TextStyle(fontSize: 16, fontWeight: FontWeight.bold)),
            const SizedBox(height: 12),
            Container(
              padding: const EdgeInsets.all(12),
              decoration: BoxDecoration(
                color: const Color(0xFF1a1a2e),
                borderRadius: BorderRadius.circular(8),
              ),
              child: child is Container ? child : Center(child: child),
            ),
          ],
        ),
      ),
    );
  }
}
