## 🤖 E2E 测试报告

| 项目 | 内容 |
|------|------|
| Issue | #5 |
| 分支 | `feat/issue-5` |
| 结果 | 🛑 已停止 |
| 时间 | 2026-04-24 09:53:34 |
| 日志 | `e2e-issue-5-20260424-095305.log` |

### 详细日志

[2026-04-24 09:53:05] E2E 任务启动 (由 manual)
分支: feat/issue-5
项目: social-world
日志: /Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-issue-5-20260424-095305.log

=== E2E 环境检查 ===
任务类型: full-stack
✅ adb: Android Debug Bridge version 1.0.41
✅ maestro: 1.39.13

=== 检查后端正端 ===
✅ 后端 API 已运行
✅ Android Emulator 已启动:
emulator-5556	device

🔧 关闭 ANR 弹窗...
✅ ANR 弹窗已关闭

=== 构建 Flutter Debug APK ===
Deleting build...                                                  192ms
Deleting .dart_tool...                                              42ms
Deleting .flutter-plugins-dependencies...                            1ms
Deleting .flutter-plugins...                                         0ms
Resolving dependencies...
Downloading packages...
  _fe_analyzer_shared 72.0.0 (100.0.0 available)
  accessibility_tools 2.5.0 (2.8.0 available)
  analyzer 6.7.0 (13.0.0 available)
  async 2.11.0 (2.13.1 available)
  bloc 8.1.4 (9.2.0 available)
  boolean_selector 2.1.1 (2.1.2 available)
  build 2.4.1 (4.0.5 available)
  build_config 1.1.1 (1.3.0 available)
  build_daemon 4.0.2 (4.1.1 available)
  build_resolvers 2.4.2 (3.0.4 available)
  build_runner 2.4.13 (2.13.1 available)
  build_runner_core 7.3.2 (9.3.2 available)
  characters 1.3.0 (1.4.1 available)
  checked_yaml 2.0.3 (2.0.4 available)
  clock 1.1.1 (1.1.2 available)
  code_builder 4.10.1 (4.11.1 available)
  collection 1.18.0 (1.19.1 available)
  cross_file 0.3.4+2 (0.3.5+2 available)
  cupertino_icons 1.0.8 (1.0.9 available)
  dart_style 2.3.7 (3.1.8 available)
  fake_async 1.3.1 (1.3.3 available)
  ffi 2.1.3 (2.2.0 available)
  file_selector_linux 0.9.3+2 (0.9.4 available)
  file_selector_macos 0.9.4+2 (0.9.5 available)
  file_selector_platform_interface 2.6.2 (2.7.0 available)
  file_selector_windows 0.9.3+4 (0.9.3+5 available)
  flutter_bloc 8.1.6 (9.1.1 available)
  flutter_gen_core 5.10.0 (5.14.0 available)
  flutter_gen_runner 5.10.0 (5.14.0 available)
  flutter_lints 4.0.0 (6.0.0 available)
  flutter_plugin_android_lifecycle 2.0.26 (2.0.34 available)
  flutter_secure_storage 9.2.4 (10.0.0 available)
  flutter_secure_storage_linux 1.2.3 (3.0.0 available)
  flutter_secure_storage_macos 3.1.3 (4.0.0 available)
  flutter_secure_storage_platform_interface 1.1.2 (2.0.1 available)
  flutter_secure_storage_web 1.2.1 (2.1.0 available)
  flutter_secure_storage_windows 3.1.2 (4.1.0 available)
  freezed 2.5.7 (3.2.5 available)
  freezed_annotation 2.4.4 (3.1.0 available)
  get_it 7.7.0 (9.2.1 available)
  go_router 14.8.1 (17.2.2 available)
  http_parser 4.0.2 (4.1.2 available)
  image_picker 1.1.2 (1.2.1 available)
  image_picker_android 0.8.12+21 (0.8.13+16 available)
  image_picker_for_web 3.0.6 (3.1.1 available)
  image_picker_ios 0.8.12+2 (0.8.13+6 available)
  image_picker_linux 0.2.1+2 (0.2.2 available)
  image_picker_macos 0.2.1+2 (0.2.2+1 available)
  image_picker_platform_interface 2.10.1 (2.11.1 available)
  image_picker_windows 0.2.1+1 (0.2.2 available)
  injectable 2.5.1 (3.0.0 available)
  injectable_generator 2.6.2 (3.0.2 available)
  inspector 3.1.0 (4.0.0 available)
  intl 0.19.0 (0.20.2 available)
  js 0.6.7 (0.7.2 available)
  json_annotation 4.9.0 (4.11.0 available)
  json_serializable 6.9.0 (6.13.1 available)
  leak_tracker 10.0.5 (11.0.2 available)
  leak_tracker_flutter_testing 3.0.5 (3.0.10 available)
  leak_tracker_testing 3.0.1 (3.0.2 available)
  lints 4.0.0 (6.1.0 available)
  macros 0.1.2-main.4 (0.1.3-main.0 available)
  matcher 0.12.16+1 (0.12.19 available)
  material_color_utilities 0.11.1 (0.13.0 available)
  meta 1.15.0 (1.18.2 available)
! package_info_plus 8.0.2 (overridden) (10.1.0 available)
  package_info_plus_platform_interface 3.2.1 (4.1.0 available)
  path 1.9.0 (1.9.1 available)
  path_provider_android 2.2.15 (2.3.1 available)
  path_provider_foundation 2.4.1 (2.6.0 available)
  petitparser 6.0.2 (7.0.2 available)
  pubspec_parse 1.4.0 (1.5.0 available)
  sentry 8.14.2 (9.19.0 available)
  sentry_flutter 8.14.2 (9.19.0 available)
  shared_preferences 2.5.3 (2.5.5 available)
  shared_preferences_android 2.4.7 (2.4.23 available)
  shared_preferences_foundation 2.5.4 (2.5.6 available)
  shared_preferences_platform_interface 2.4.1 (2.4.2 available)
  shelf 1.4.1 (1.4.2 available)
  shelf_web_socket 2.0.1 (3.0.0 available)
  source_gen 1.5.0 (4.2.2 available)
  source_helper 1.3.5 (1.3.11 available)
  source_span 1.10.0 (1.10.2 available)
  stack_trace 1.11.1 (1.12.1 available)
  stream_channel 2.1.2 (2.1.4 available)
  string_scanner 1.2.0 (1.4.1 available)
  term_glyph 1.2.1 (1.2.2 available)
  test_api 0.7.2 (0.7.11 available)
  url_launcher 6.3.1 (6.3.2 available)
  url_launcher_android 6.3.14 (6.3.29 available)
  url_launcher_ios 6.3.3 (6.4.1 available)
  url_launcher_linux 3.2.1 (3.2.2 available)
  url_launcher_macos 3.2.2 (3.2.5 available)
  url_launcher_web 2.3.3 (2.4.2 available)
  url_launcher_windows 3.1.4 (3.1.5 available)
  vector_graphics_compiler 1.1.16 (1.2.0 available)
  vector_math 2.1.4 (2.3.0 available)
  vm_service 14.2.5 (15.1.0 available)
  widgetbook 3.15.0 (3.22.0 available)
  widgetbook_annotation 3.6.0 (3.11.0 available)
  win32 5.10.1 (6.1.0 available)
  xml 6.5.0 (6.6.1 available)
Got dependencies!
102 packages have newer versions incompatible with dependency constraints.
Try `flutter pub outdated` for more information.

Running Gradle task 'assembleDebug'...                          
Your project is configured with Android NDK 23.1.7779620, but the following plugin(s) depend on a different Android NDK version:
- flutter_plugin_android_lifecycle requires Android NDK 26.1.10909125
- flutter_secure_storage requires Android NDK 26.1.10909125
- image_picker_android requires Android NDK 26.1.10909125
- package_info_plus requires Android NDK 26.1.10909125
- path_provider_android requires Android NDK 26.1.10909125
- sentry_flutter requires Android NDK 26.1.10909125
- shared_preferences_android requires Android NDK 26.1.10909125
- url_launcher_android requires Android NDK 26.1.10909125
Fix this issue by using the highest Android NDK version (they are backward compatible).
Add the following to /Users/insulate/Desktop/social-world/apps/mobile/android/app/build.gradle:

    android {
        ndkVersion = "26.1.10909125"
        ...
    }


🛑 任务已停止
Running Gradle task 'assembleDebug'...                             23.5s
✓ Built build/app/outputs/flutter-apk/app-debug.apk

⏹️ 任务已手动停止


---
*由 Agent Control Center 自动生成*