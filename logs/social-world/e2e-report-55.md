## 🤖 E2E 测试报告

| 项目 | 内容 |
|------|------|
| Issue | #55 |
| 分支 | `feat/issue-43` |
| 结果 | ❌ 失败 |
| 时间 | 2026-04-24 02:54:28 |
| 日志 | `e2e-issue-55-20260424-012302.log` |

### 详细日志

[2026-04-24 01:23:02] E2E 任务启动 (由 rework-55)
分支: feat/issue-43
项目: social-world
日志: /Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-issue-55-20260424-012302.log

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
Deleting .dart_tool...                                               0ms
Deleting .flutter-plugins-dependencies...                            0ms
Deleting .flutter-plugins...                                         0ms
Resolving dependencies...
Downloading packages...
  _fe_analyzer_shared 72.0.0 (99.0.0 available)
  accessibility_tools 2.5.0 (2.8.0 available)
  analyzer 6.7.0 (12.1.0 available)
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
! package_info_plus 8.0.2 (overridden) (10.0.0 available)
  package_info_plus_platform_interface 3.2.1 (4.0.0 available)
  path 1.9.0 (1.9.1 available)
  path_provider_android 2.2.15 (2.3.1 available)
  path_provider_foundation 2.4.1 (2.6.0 available)
  petitparser 6.0.2 (7.0.2 available)
  pubspec_parse 1.4.0 (1.5.0 available)
  sentry 8.14.2 (9.18.0 available)
  sentry_flutter 8.14.2 (9.18.0 available)
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
  win32 5.10.1 (6.0.1 available)
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
Add the following to /private/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/apps/mobile/android/app/build.gradle:

    android {
        ndkVersion = "26.1.10909125"
        ...
    }

Running Gradle task 'assembleDebug'...                             15.7s
✓ Built build/app/outputs/flutter-apk/app-debug.apk
✅ APK 构建完成

=== 安装 APK ===
Performing Streamed Install
Success
✅ APK 安装完成
Success

=== 修复 Maestro includes 文件格式 ===
✅ includes 文件格式检查完成

=== 运行 E2E 测试 Agent ===
Agent 文件: '/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dashboard-agent-social-world-e2e-tester-68689.yaml'
Debug 目录: '/Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302'
报告文件: '/Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-report-e2e-issue-55-20260424-012302.md'

【任务上下文】
- Issue 编号: #55
- 项目路径: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001
100
- 项目地图: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001
100/AGENTS.md
- 关联 PRD: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001
100/prd/v1-profile.md,/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/prd/v1-mvp.md,/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev
-issue-55-815001100/prd/v1-matching.md
- 设计资产: 未找到
- 文档目录: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001
100/docs



**你的任务类型：E2E 测试任务（E2E Tester Agent）**
目标：执行 Maestro E2E 测试，分析失败原因，生成详细报告。

Issue #55，分支: feat/issue-43
TurnBegin(
    user_input='【任务上下文】\n- Issue 编号: #55\n- 项目路径: /var/folders/lt/w
3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100\n- 项目地图: /var/folders
/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/AGENTS.md\n- 关联 PR
D: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/prd/v
1-profile.md,/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-81500
1100/prd/v1-mvp.md,/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100/prd/v1-matching.md\n- 设计资产: 未找到\n- 文档目录: /var/folders/lt/w
3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/docs\n\n\n\n**你的任务类
型：E2E 测试任务（E2E Tester Agent）**\n目标：执行 Maestro E2E 测试，分析失败原
因，生成详细报告。\n\nIssue #55，分支: feat/issue-43'
)
StepBegin(n=1)
ThinkPart(
    type='think',
    think='让我开始执行 E2E 测试任务。按照流程，我需要：\n\n1. 先读取项目地图 `A
GENTS.md`\n2. 读取 E2E 规范 `.maestro/README.md`\n3. 扫描 flow 文件\n4. 阅读 PRD
了解功能需求\n5. 执行测试、分析结果、生成报告\n\n先并行读取项目地图和相关文件。'
,
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_Id9UQcMvo6o5L4BYEr7XObAR',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/AGENTS.md"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_cT3vW70ne5oRUQECaLv9EvYr',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_Id9UQcMvo6o5L4BYEr7XObAR',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# Social World 智能体导航地图\n     2\t\n     3\t> 仓库
即唯一真实来源。不在仓库里的知识等于不存在。\n     4\t\n     5\t## 文档地图\n   
 6\t\n     7\t| 你要找什么         | 去哪里                                     
                                         |\n     8\t| ------------------ | -----
------------------------------------------------------------------------------- 
|\n     9\t| 项目背景           | [README.md](README.md)                        
                                      |\n    10\t| 编码与架构规范     | [docs/de
sign-docs/INDEX.md](docs/design-docs/INDEX.md)                               |\n
   11\t| Issue 处理工作流   | [docs/design-docs/issue-workflow.md](docs/design-d
ocs/issue-workflow.md)             |\n    12\t| UI 风格规范        | [docs/desig
n-docs/ui-style-guide/INDEX.md](docs/design-docs/ui-style-guide/INDEX.md) |\n   
13\t| 业务术语           | [docs/product-specs/glossary.md](docs/product-specs/g
lossary.md)                     |\n    14\t| 当前需求           | [prd/v1-mvp.md
](prd/v1-mvp.md)                                                       |\n    15
\t| 模块实现           | [docs/modules/INDEX.md](docs/modules/INDEX.md)（含 prof
ile）                         |\n    16\t| **执行计划与模板** | **[docs/exec-pla
ns/INDEX.md](docs/exec-plans/INDEX.md)** ⬅️ **开发前必读**           |\n    17\t|
技术债务           | [docs/exec-plans/tech-debt-tracker.md](docs/exec-plans/tech
-debt-tracker.md)         |\n    18\t| Maestro E2E 规范   | [.maestro/README.md]
(.maestro/README.md)                                             |\n    19\t\n  
 20\t## 目录结构\n    21\t\n    22\t```\n    23\tdocs/design-docs/    # 架构规范
、编码约定、数据库/API 规范\n    24\tdocs/product-specs/  # 业务术语、领域规则\n
   25\tdocs/modules/        # 模块实现文档（跨代码层级）\n    26\tdocs/exec-plan
s/     # 执行计划 + 按任务类型拆分的专用模板 + 技术债务\n    27\tprd/           
     # 产品需求文档\n    28\tapps/                # mobile (Flutter), api (NestJ
S), admin (Next.js)\n    29\tpackages/            # shared-types, design-system\
n    30\t```\n    31\t\n    32\t## 按任务类型的必读文档速查\n    33\t\n    34\t|
任务类型                             | 强制第 1 步                              
      | 强制第 2 步                   | 产出位置约定                            
                                                          |\n    35\t| ---------
--------------------------- | ----------------------------------------------- | 
----------------------------- | ------------------------------------------------
-------------------------------------------------- |\n    36\t| 新功能开发 (`fea
t`)                  | `docs/exec-plans/TEMPLATE-fullstack-feature.md` | `docs/d
esign-docs/flutter.md` | 代码直接写入 `apps/` / `packages/`                     
                                           |\n    37\t| **设计与 Design-Only PR 
(`design`)** | `docs/exec-plans/TEMPLATE-design.md`            | `docs/design-do
cs/flutter.md` | **`.dart` 资产必须放在 `designs/issue-XXX/flutter-widgets/`，禁
止直接写入 `packages/` 或 `apps/`** |\n    38\t| Bug 修复 (`fix`)               
     | `docs/exec-plans/TEMPLATE-fullstack-bugfix.md`  | —                      
      | 代码直接写入对应模块                                                    
                          |\n    39\t| 文档任务 (`docs`)                    | `d
ocs/exec-plans/TEMPLATE-documentation.md`     | —                             | 
`docs/` / `prd/`                                                                
                  |\n    40\t\n    41\t> 设计任务 ≠ 美术任务。Design-Only PR 必
须产出 `.dart` 纯 UI Widget + `MANIFEST.md` 迁移清单 + Widgetbook 用例代码 + Gol
den Test 代码。\n    42\t\n    43\t## 黄金原则\n    44\t\n    45\t1. **计划先行*
*：复杂任务（多模块协调、架构变更、新用户旅程）必须先写 `docs/exec-plans/PLAN-XX
X.md`，按模板执行\n    46\t2. **全链路交付**：一个需求必须同时包含 DB + API + Mo
bile 的实现\n    47\t3. **契约同步**：`packages/shared-types/` 是前后端的唯一真
实来源\n    48\t4. **约束即代码**：linter、hooks、CI 会强制执行规则，不在规则里
的约束等于不存在\n    49\t5. **效果优先**：先跑通端到端，再考虑优雅性\n    50\t\
n    51\t## 常见任务\n    52\t\n    53\t| 任务                  | 命令          
                                      |\n    54\t| --------------------- | -----
----------------------------------------------- |\n    55\t| 查看执行计划模板   
  | [docs/exec-plans/INDEX.md](docs/exec-plans/INDEX.md) |\n    56\t| 启动开发环
境          | `make dev`                                           |\n    57\t| 
检查文档健康          | `make check-docs`                                    |\n
   58\t| 运行测试              | `make test`                                    
     |\n    59\t| Mobile E2E (login)    | `make test-mobile-e2e-login`          
              |\n    60\t| Mobile E2E (devtools) | `make test-mobile-e2e-devtool
s`                      |\n    61\t| Mobile E2E (完整)     | `make test-mobile-e
2e`                               |\n    62\t| 启动 E2E 后端环境     | `make e2e
-env-up`                                    |\n    63\t| 安装 hooks            |
`make hooks`                                         |\n    64\t| 架构约束检查  
       | `make arch-check`                                    |\n    65\t| Agent
配置检查        | `make agent-check`                                   |\n',
        message='65 lines read from file starting from line 1. Total lines in fi
le: 65. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='n00000gn/T/dev-issue-55-815001100/prd/v1-profile.md"}'
)
ToolCall(
    type='function',
    id='tool_dk3Kipt7j5hxNsC3LT8nLfcm',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_cT3vW70ne5oRUQECaLv9EvYr',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# V1 Profile 产品需求文档\n     2\t\n     3\t> 状态：待
评审  \n     4\t> 版本：1.0  \n     5\t> 日期：2026-04-10  \n     6\t> 关联 PRD
：[prd/v1-mvp.md](./v1-mvp.md)\n     7\t\n     8\t## 1. 背景与目标\n     9\t\n  
 10\t### 背景\n    11\tSocial World 的核心社交体验建立在用户个人资料之上。资料的
质量直接影响匹配成功率和用户留存。V1 MVP 阶段要求用户首次注册后必须完成「最小资
料」，才能进入首页的卡片推荐和滑动匹配流程。\n    12\t\n    13\t### 目标\n    14
\t1. 支持用户上传头像和相册照片，并存储于对象存储（OSS）。\n    15\t2. 支持用户
填写基础资料（昵称、性别、出生日期、当前位置）。\n    16\t3. 支持选填扩展资料（
个人签名、兴趣标签、匹配偏好）。\n    17\t4. 在用户保存资料时校验年龄必须 ≥ 18 
周岁，否则阻断。\n    18\t5. 资料不完整时，系统应阻止用户进入匹配首页并引导其补
全。\n    19\t6. 支持用户随时编辑自己的资料。\n    20\t7. 为内容审核预留接口：照
片上传后默认进入 `pending_review` 状态。\n    21\t\n    22\t---\n    23\t\n    2
4\t## 2. 用户故事与验收标准\n    25\t\n    26\t### US-PROF-1：完成最小资料\n    
27\t\n    28\t**作为** 新注册用户，  \n    29\t**我希望** 上传头像并填写基础资料
，  \n    30\t**以便** 进入首页开始匹配。\n    31\t\n    32\t#### 验收标准\n    
33\t> 以上验收标准的完整 Gherkin 用例详见 `tests/bdd/profile.feature`。\n    34\
t\n    35\t```gherkin\n    36\tFeature: Complete Minimum Profile\n    37\t  Scen
ario: User completes minimum profile successfully\n    38\t    Given the user ha
s just registered and has no profile\n    39\t    When the user uploads an avata
r, enters nickname "Alice", selects gender "female", enters birth_date "1998-05-
20", and allows location access\n    40\t    Then the profile should be marked a
s complete\n    41\t    And the user should enter the home screen\n    42\t\n   
43\t  Scenario: Missing required field blocks completion\n    44\t    Given the 
user has just registered\n    45\t    When the user tries to save profile withou
t uploading an avatar\n    46\t    Then the system should reject with code 1000\
n    47\t    And the profile should remain incomplete\n    48\t\n    49\t  Scena
rio: Underage user is blocked\n    50\t    Given the user has just registered\n 
  51\t    When the user enters birth_date "2010-01-01"\n    52\t    Then the sys
tem should reject with code 1000\n    53\t    And the response message should in
dicate age restriction\n    54\t    And no profile should be saved\n    55\t```\
n    56\t\n    57\t#### 技术要求\n    58\t- 必填字段：头像 `avatar_url`、昵称 `n
ickname`、性别 `gender`、出生日期 `birth_date`、经纬度 `latitude` / `longitude`
。\n    59\t- 性别枚举：`male`、`female`、`other`。\n    60\t- 年龄计算：根据 `b
irth_date` 与当前 UTC 时间精确计算周岁，`< 18` 则返回 `code=1000`。\n    61\t- 
位置获取：Mobile 端调用系统定位 API，经纬度上传至后端。\n    62\t\n    63\t---\n
   64\t\n    65\t### US-PROF-2：上传头像与相册\n    66\t\n    67\t**作为** 用户
，  \n    68\t**我希望** 上传多张照片展示自己，  \n    69\t**以便** 提高被匹配的
概率。\n    70\t\n    71\t#### 验收标准\n    72\t> 以上验收标准的完整 Gherkin 用
例详见 `tests/bdd/profile.feature`。\n    73\t\n    74\t```gherkin\n    75\tFeat
ure: Photo Upload\n    76\t  Scenario: User uploads avatar successfully\n    77\
t    Given the user is on the profile setup page\n    78\t    When the user sele
cts a JPG image within 10MB as avatar\n    79\t    Then the image should be uplo
aded to OSS\n    80\t    And the avatar_url should be saved in user_profiles\n  
 81\t    And the photo should trigger a content moderation event with status "pe
nding_review"\n    82\t\n    83\t  Scenario: User uploads album photos\n    84\t
   Given the user has completed minimum profile\n    85\t    When the user uploa
ds 3 additional photos to the album\n    86\t    Then each photo should be saved
in user_photos with sort_order 1 to 3\n    87\t    And each photo should have mo
deration_status "pending_review"\n    88\t    And the user should not exceed the
maximum of 6 album photos\n    89\t\n    90\t  Scenario: Upload oversized image 
is rejected\n    91\t    Given the user selects an image larger than 10MB\n    9
2\t    When the upload request is sent\n    93\t    Then the API should reject w
ith code 1000\n    94\t```\n    95\t\n    96\t#### 技术要求\n    97\t- 头像限制
：JPG/PNG，最大 10MB。\n    98\t- 相册限制：最多 6 张，单张最大 10MB。\n    99\t
- OSS 上传：MVP 阶段后端生成临时 STS 凭证或直接提供预签名 URL 给 Mobile 端直传。
\n   100\t- 上传完成后，Mobile 将最终 URL 回传后端保存。\n   101\t- `user_photos
` 表字段：`id`, `user_id`, `url`, `sort_order`, `moderation_status`（默认 `pendi
ng_review`）。\n   102\t- 内容审核：MVP 阶段仅触发 Webhook 事件或打印日志，不做
实时拦截。\n   103\t\n   104\t---\n   105\t\n   106\t### US-PROF-3：填写扩展资料
\n   107\t\n   108\t**作为** 用户，  \n   109\t**我希望** 添加签名、标签和匹配偏
好，  \n   110\t**以便** 让系统推荐更符合我口味的人。\n   111\t\n   112\t#### 验
收标准\n   113\t> 以上验收标准的完整 Gherkin 用例详见 `tests/bdd/profile.feature
`。\n   114\t\n   115\t```gherkin\n   116\tFeature: Extended Profile Information
\n   117\t  Scenario: User adds bio and tags\n   118\t    Given the user has com
pleted minimum profile\n   119\t    When the user enters bio "Love hiking and co
ffee" and selects tags "hiking,coffee,music"\n   120\t    Then the bio and tags 
should be saved in user_profiles\n   121\t    And the tags should be parsed as a
JSON array\n   122\t\n   123\t  Scenario: User sets matching preferences\n   124
\t    Given the user has completed minimum profile\n   125\t    When the user se
ts preferred_gender to "male" and age_range to "20-30"\n   126\t    Then the mat
ching preferences should be saved in user_profiles\n   127\t    And future recom
mendations should respect these preferences\n   128\t\n   129\t  Scenario: Tag l
imit exceeded\n   130\t    Given the user tries to add 10 tags\n   131\t    When
the profile is saved\n   132\t    Then the system should reject with code 1000\n
  133\t    And the maximum tag count should be 8\n   134\t```\n   135\t\n   136\
t#### 技术要求\n   137\t- 个人签名 `bio`：限制 200 字符。\n   138\t- 兴趣标签 `t
ags`：JSON 数组格式，最多 8 个，每个标签最长 16 字符。\n   139\t- 匹配偏好 `pref
erred_gender`：`male`、`female`、`both`。\n   140\t- 匹配偏好 `preferred_age_min
` / `preferred_age_max`：整数，范围 18-80。\n   141\t- 扩展资料不影响「资料完整
性」判定，但影响推荐算法。\n   142\t\n   143\t---\n   144\t\n   145\t### US-PROF
-4：资料完整性校验\n   146\t\n   147\t**作为** 平台运营者，  \n   148\t**我希望*
* 未完善资料的用户无法进入匹配流程，  \n   149\t**以便** 保证推荐池的用户质量。\
n   150\t\n   151\t#### 验收标准\n   152\t> 以上验收标准的完整 Gherkin 用例详见 
`tests/bdd/profile.feature`。\n   153\t\n   154\t```gherkin\n   155\tFeature: Pr
ofile Completeness Guard\n   156\t  Scenario: Incomplete profile blocks home scr
een access\n   157\t    Given the user has not completed minimum profile\n   158
\t    When the user tries to navigate to the home screen\n   159\t    Then the a
pp should redirect the user to profile setup\n   160\t    And the API should rej
ect any matching-related requests with code 1000\n   161\t\n   162\t  Scenario: 
Incomplete profile blocks swipe action\n   163\t    Given the user has incomplet
e profile\n   164\t    When the user sends a swipe request via API\n   165\t    
Then the API should return code 1000\n   166\t    And the message should prompt 
completing profile first\n   167\t```\n   168\t\n   169\t#### 技术要求\n   170\t
- 后端 JWT Guard 之后增加 `ProfileCompletionGuard`：检查 `user_profiles.is_compl
ete`。\n   171\t- `GET /auth/me` 返回 `profile_complete: boolean`，供 Mobile 在
启动时判断跳转目标。\n   172\t- 未完善资料的用户调用 `GET /matching/recommendati
ons` 或 `POST /matching/actions` 时，统一返回 `code=1000`。\n   173\t\n   174\t-
--\n   175\t\n   176\t### US-PROF-5：编辑资料\n   177\t\n   178\t**作为** 已完善
资料的用户，  \n   179\t**我希望** 随时修改我的头像、照片和资料信息，  \n   180\
t**以便** 保持资料的时效性。\n   181\t\n   182\t#### 验收标准\n   183\t> 以上验
收标准的完整 Gherkin 用例详见 `tests/bdd/profile.feature`。\n   184\t\n   185\t`
``gherkin\n   186\tFeature: Edit Profile\n   187\t  Scenario: User edits nicknam
e and bio\n   188\t    Given the user has a complete profile\n   189\t    When t
he user updates nickname to "Alice2" and bio to "New bio"\n   190\t    Then the 
updated fields should be saved immediately\n   191\t    And the updated_at times
tamp should be refreshed\n   192\t\n   193\t  Scenario: User reorders album phot
os\n   194\t    Given the user has 3 album photos\n   195\t    When the user dra
gs photo 3 to position 1\n   196\t    Then the sort_order in user_photos should 
be updated to reflect the new order\n   197\t\n   198\t  Scenario: User deletes 
an album photo\n   199\t    Given the user has 3 album photos\n   200\t    When 
the user deletes the second photo\n   201\t    Then the photo should be soft-del
eted (deleted_at set)\n   202\t    And the remaining photos should maintain vali
d sort_order\n   203\t```\n   204\t\n   205\t#### 技术要求\n   206\t- 资料编辑 A
PI：`PATCH /users/me/profile`（支持部分更新）。\n   207\t- 相册照片重排：`PATCH 
/users/me/photos/reorder`（接收 `photo_ids` 数组，按索引作为新 `sort_order`）。\
n   208\t- 相册照片删除：`DELETE /users/me/photos/:id`（软删除，设置 `deleted_at
`）。\n   209\t- 头像变更后，旧头像 OSS 对象可异步清理（MVP 阶段可暂不实现物理删
除）。\n   210\t\n   211\t---\n   212\t\n   213\t## 3. 非功能需求\n   214\t\n   
215\t### 性能\n   216\t- 头像上传响应时间（从选择图片到收到 OSS URL）< 3 秒（P95
，普通 4G 网络）。\n   217\t- `PATCH /users/me/profile` 响应时间 < 200ms。\n   2
18\t- 相册照片列表加载时间 < 500ms。\n   219\t\n   220\t### 安全\n   221\t- OSS 
上传使用临时凭证或预签名 URL，避免长期 AccessKey 泄露。\n   222\t- 用户只能修改
自己的资料（通过 JWT `user_id` 校验）。\n   223\t- 照片上传后默认 `moderation_st
atus=pending_review`，不直接展示给全部用户（MVP 阶段可先展示，预留状态字段）。\n
  224\t\n   225\t### 兼容性\n   226\t- 图片压缩：Mobile 上传前自动压缩图片至宽度
1080px，减少流量和上传时间。\n   227\t- 支持相册从相机拍摄或图库选择。\n   228\t
\n   229\t---\n   230\t\n   231\t## 4. 里程碑\n   232\t\n   233\t| 阶段 | 目标 |
预计工期 |\n   234\t|------|------|----------|\n   235\t| M-PROF-1 | 最小资料 AP
I + 年龄校验 + Mobile 资料设置页（头像/昵称/性别/生日/位置） | 2 天 |\n   236\t|
M-PROF-2 | 相册上传/重排/删除 + OSS 集成 + 扩展资料（签名/标签/偏好） | 2 天 |\n
  237\t| M-PROF-3 | 资料完整性校验 Guard + 资料编辑页 + BDD 测试闭环 | 1 天 |\n 
 238\t\n   239\t**总预计工期：5 天。**\n   240\t\n   241\t---\n   242\t\n   243\
t## 5. 风险与回滚策略\n   244\t\n   245\t| 风险 | 影响 | 缓解措施 |\n   246\t|--
----|------|----------|\n   247\t| OSS 上传凭证过期导致用户上传失败 | 中 | 预签
名 URL 有效期设置为 15 分钟，失败时后端重新生成 |\n   248\t| 年龄计算边界错误导
致合法用户被误封 | 高 | 增加专门单元测试覆盖闰年、生日当天、跨时区场景 |\n   249
\t| 相册照片排序并发冲突 | 低 | 重排接口接收完整 `photo_ids` 数组，服务端事务更
新 |\n   250\t| 未完善资料用户绕过 Guard 进入首页 | 高 | 后端所有匹配相关 API 统
一增加 `ProfileCompletionGuard` |\n   251\t\n   252\t### 回滚策略\n   253\t- **
资料字段变更回滚**：通过数据库 Migration `down` 方法回退 Schema 变更。\n   254\t
- **OSS 配置错误**：切换为本地存储或开发环境 Mock，保证测试可用。\n   255\t- **G
uard 误拦截**：临时移除 `ProfileCompletionGuard` 的匹配 API 拦截，保留日志告警。
\n',
        message='255 lines read from file starting from line 1. Total lines in f
ile: 255. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/prd
/v1-mvp.md"}'
)
ToolCall(
    type='function',
    id='tool_X0uDsKvQ4IlPsAemPDNQk0Cj',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_dk3Kipt7j5hxNsC3LT8nLfcm',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# V1 MVP 产品需求文档\n     2\t\n     3\t> 状态：执行中 
\n     4\t> 版本：1.0  \n     5\t> 日期：2026-04-09\n     6\t\n     7\t## 1. 背
景与目标\n     8\t\n     9\t### 背景\n    10\tSocial World 是一款面向年轻人群的
陌生人社交应用，核心定位是通过高效的滑动匹配和即时通讯帮助用户建立新的社交关系。
\n    11\t\n    12\t### 目标\n    13\t在 V1 MVP 阶段，实现最小可用的陌生人社交闭
环：\n    14\t- 用户可以注册登录并完善个人资料\n    15\t- 系统为用户推荐附近的异
性/同性用户卡片\n    16\t- 用户可以通过左滑/右滑表达喜好\n    17\t- 双向喜欢的用
户可以开始一对一聊天\n    18\t\n    19\t## 2. 用户故事与验收标准\n    20\t\n    
21\t### US-1：注册与登录\n    22\t\n    23\t**作为** 新用户，**我希望** 通过手机
号快速注册登录，**以便** 开始使用 App。\n    24\t\n    25\t#### 验收标准\n    26
\t```gherkin\n    27\tFeature: User Authentication\n    28\t  Scenario: New user
registers with phone number\n    29\t    Given the user opens the app for the fi
rst time\n    30\t    When the user enters a valid phone number and SMS code\n  
 31\t    Then an account should be created\n    32\t    And the user should rece
ive an access token\n    33\t    And the user should be redirected to profile se
tup\n    34\t\n    35\t  Scenario: Existing user logs in\n    36\t    Given the 
user already has an account\n    37\t    When the user enters the phone number a
nd correct SMS code\n    38\t    Then the user should receive new tokens\n    39
\t    And the user should land on the home screen if profile is complete\n    40
\t```\n    41\t\n    42\t#### 技术要求\n    43\t- 使用 JWT Token（Access 15min +
Refresh 7天）\n    44\t- 短信验证码 MVP 阶段可先用固定码（如 `123456`）或接入阿
里云短信\n    45\t- 手机号作为唯一用户标识\n    46\t\n    47\t---\n    48\t\n   
49\t### US-2：完善个人资料\n    50\t\n    51\t**作为** 用户，**我希望** 上传照片
并填写个人资料，**以便** 被别人发现和匹配。\n    52\t\n    53\t#### 验收标准\n  
 54\t```gherkin\n    55\tFeature: User Profile\n    56\t  Scenario: User complet
es minimum profile\n    57\t    Given the user has just registered\n    58\t    
When the user uploads an avatar, enters nickname, selects gender and birth date\
n    59\t    Then the profile should be marked as complete\n    60\t    And the 
user should enter the home screen\n    61\t\n    62\t  Scenario: Incomplete prof
ile blocks matching\n    63\t    Given the user has not completed minimum profil
e\n    64\t    When the user tries to swipe cards\n    65\t    Then the system s
hould prompt the user to complete profile first\n    66\t```\n    67\t\n    68\t
#### 资料字段\n    69\t- 必填：头像、昵称、性别、出生日期、当前位置\n    70\t- 
选填：个人签名、相册（最多6张）、兴趣标签（最多8个）、匹配偏好（性别、年龄范围）
\n    71\t\n    72\t#### 技术要求\n    73\t- 头像和照片上传至对象存储（阿里云 OS
S / AWS S3）\n    74\t- 照片上传后经过内容安全审核（MVP 阶段可先人工延后审核，预
留接口）\n    75\t- 用户需满 18 周岁（根据 birth_date 计算）\n    76\t\n    77\t
---\n    78\t\n    79\t### US-3：推荐卡片与滑动匹配\n    80\t\n    81\t**作为** 
用户，**我希望** 看到附近有趣的异性/同性用户卡片，**以便** 通过滑动选择喜欢或跳
过。\n    82\t\n    83\t#### 验收标准\n    84\t```gherkin\n    85\tFeature: Card
Swipe Matching\n    86\t  Scenario: User sees recommendation cards\n    87\t    
Given the user is on the home screen with complete profile\n    88\t    When the
app loads\n    89\t    Then the user should see a stack of profile cards sorted 
by distance and activity\n    90\t\n    91\t  Scenario: User likes a profile\n  
 92\t    Given the user sees a card\n    93\t    When the user swipes right or t
aps the like button\n    94\t    Then a like action should be recorded\n    95\t
   And if the other user already liked this user, a match should be created\n   
96\t\n    97\t  Scenario: Mutual like creates a match\n    98\t    Given Alice l
iked Bob and Bob liked Alice\n    99\t    When the system detects the mutual lik
e\n   100\t    Then a match record should be created\n   101\t    And a conversa
tion should be auto-created\n   102\t    And both users should receive a match n
otification\n   103\t```\n   104\t\n   105\t#### 推荐规则（MVP）\n   106\t1. **
过滤条件**：\n   107\t   - 根据用户设置的偏好性别筛选\n   108\t   - 根据当前位置
，默认半径 50km（可切换「同城」「全国」）\n   109\t   - 排除已经 Like / Pass / B
lock 过的用户\n   110\t   - 排除已经匹配成功的用户\n   111\t   - 排除资料不完整
的用户\n   112\t2. **排序规则**：\n   113\t   - 距离越近越优先（权重 50%）\n   1
14\t   - 近期活跃用户优先（权重 30%）\n   115\t   - 新用户有额外曝光加成（权重 2
0%）\n   116\t3. **卡片展示**：\n   117\t   - 首图（头像或相册第一张）占卡片主要
区域\n   118\t   - 叠加显示：昵称、年龄、距离、个性签名\n   119\t\n   120\t#### 
技术要求\n   121\t- 推荐池缓存于 Redis（TTL 5 分钟），减少数据库查询压力\n   122
\t- 地理位置使用 MySQL `decimal` 存储经纬度，通过 Haversine 公式计算距离（MVP 简
化方案）\n   123\t- 滑动操作幂等处理（防止重复提交）\n   124\t\n   125\t---\n   
126\t\n   127\t### US-4：即时通讯（IM）\n   128\t\n   129\t**作为** 匹配成功的用
户，**我希望** 与对方实时聊天，**以便** 加深了解。\n   130\t\n   131\t#### 验收
标准\n   132\t```gherkin\n   133\tFeature: Instant Messaging\n   134\t  Scenario
: Matched users start chatting\n   135\t    Given Alice and Bob have matched\n  
136\t    When Alice opens the conversation\n   137\t    Then Alice should see th
e system message "你们已匹配成功，开始聊天吧"\n   138\t    And Alice should be a
ble to send a text message\n   139\t\n   140\t  Scenario: Real-time message deli
very\n   141\t    Given Alice and Bob are both online\n   142\t    When Alice se
nds "你好"\n   143\t    Then Bob should receive the message within 1 second\n   
144\t    And Bob\'s unread count for this conversation should increase by 1\n   
145\t\n   146\t  Scenario: Read receipt\n   147\t    Given Bob has received a me
ssage from Alice\n   148\t    When Bob opens the conversation\n   149\t    Then 
Alice should see the message status change to "已读"\n   150\t    And Bob\'s unr
ead count should be reset to 0\n   151\t```\n   152\t\n   153\t#### 消息类型\n  
154\t- MVP 阶段支持：文本、图片\n   155\t- 单条文本限制 500 字符\n   156\t- 图片
限制 10MB（JPG/PNG）\n   157\t\n   158\t#### 技术要求\n   159\t- 使用 Socket.io 
实现实时消息收发\n   160\t- 消息持久化到 MySQL\n   161\t- 未读数存储于 Redis\n  
162\t- 离线用户通过极光推送/Firebase 触发新消息提醒\n   163\t\n   164\t---\n   1
65\t\n   166\t## 3. 非功能需求\n   167\t\n   168\t### 性能\n   169\t- 首页推荐卡
片加载时间 < 800ms（P95）\n   170\t- 消息发送到达时间 < 1s（在线用户）\n   171\t
\n   172\t### 安全\n   173\t- 所有 API 通信使用 HTTPS\n   174\t- 用户密码/Token 
不存储于本地明文\n   175\t- 聊天内容预留敏感词过滤接口（MVP 可先用简单关键词拦截
）\n   176\t\n   177\t### 兼容性\n   178\t- 支持 Android 10+ 和 iOS 14+\n   179\
t\n   180\t---\n   181\t\n   182\t## 4. 里程碑\n   183\t\n   184\t| 阶段 | 目标 
| 预计工期 |\n   185\t|------|------|----------|\n   186\t| M1 | 用户系统：注册
登录 + 资料完善 | 1 周 |\n   187\t| M2 | 匹配系统：推荐算法 + 滑动交互 + 匹配成
功 | 1 周 |\n   188\t| M3 | IM 系统：Socket.io 消息收发 + 会话列表 + 推送 | 1 周
|\n   189\t| M4 | 联调优化：E2E 测试、UI 打磨、Bug 修复 | 1 周 |\n   190\t\n   1
91\t---\n   192\t\n   193\t## 5. 后期拓展（非 V1 范围）\n   194\t\n   195\t以下
功能在 V1 中仅预留接口，不实现完整逻辑：\n   196\t- AI 破冰话术推荐\n   197\t- 
语音/视频通话\n   198\t- 会员订阅与虚拟货币\n   199\t- 礼物系统\n   200\t- 内容
审核全自动流水线\n   201\t- 精准的推荐算法模型（DSSM / DeepFM）\n',
        message='201 lines read from file starting from line 1. Total lines in f
ile: 201. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='00gn/T/dev-issue-55-815001100/prd/v1-matching.md"}'
)
ToolCall(
    type='function',
    id='tool_GNbXXCkJzzSkZqFIBAy03gYb',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_X0uDsKvQ4IlPsAemPDNQk0Cj',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# V1 Matching 产品需求文档\n     2\t\n     3\t> 状态：待
评审  \n     4\t> 版本：1.0  \n     5\t> 日期：2026-04-10  \n     6\t> 关联 PRD
：[prd/v1-mvp.md](./v1-mvp.md)\n     7\t>\n     8\t> 设计资产参见 [designs/issue
-43/](../designs/issue-43/)\n     9\t\n    10\t## 1. 背景与目标\n    11\t\n    1
2\t### 背景\n    13\t\n    14\tSocial World 的核心体验是「滑动匹配」。用户在首页
看到基于地理位置和偏好筛选推荐的卡片，通过左滑（Pass）或右滑（Like）表达喜好。当
两个用户互相喜欢时，系统创建一次 `Match` 并自动开启 `Conversation`，促成后续的即
时通讯。\n    15\t\n    16\t### 目标\n    17\t\n    18\t1. 为用户推荐附近符合偏
好条件的异性/同性用户卡片。\n    19\t2. 支持左滑 Pass 和右滑 Like，操作幂等。\n 
  20\t3. 双向喜欢时自动创建 `Match` 和 `Conversation`，并实时推送通知。\n    21\
t4. 排除已交互、已匹配、已拉黑和资料不完整的用户。\n    22\t5. 推荐结果缓存于 Re
dis，减少数据库查询压力。\n    23\t6. 支持「同城」「全国」两种距离范围切换。\n  
 24\t\n    25\t---\n    26\t\n    27\t## 2. 用户故事与验收标准\n    28\t\n    29
\t### US-MATCH-1：获取推荐卡片\n    30\t\n    31\t**作为** 已完善资料的用户，  \
n    32\t**我希望** 看到附近符合条件的用户卡片，  \n    33\t**以便** 发现感兴趣
的社交对象。\n    34\t\n    35\t#### 验收标准\n    36\t\n    37\t> 以上验收标准
的完整 Gherkin 用例详见 `tests/bdd/matching.feature`。\n    38\t\n    39\t```ghe
rkin\n    40\tFeature: Card Recommendations\n    41\t  Scenario: User sees recom
mendation cards\n    42\t    Given the user is on the home screen with a complet
e profile\n    43\t    And there are 20 eligible candidates nearby\n    44\t    
When the app requests GET /matching/recommendations\n    45\t    Then the user s
hould receive a stack of up to 20 profile cards\n    46\t    And each card shoul
d contain: user_id, nickname, age, distance, bio, avatar_url\n    47\t    And th
e cards should be sorted by composite score\n    48\t\n    49\t  Scenario: Cache
hit returns quickly\n    50\t    Given the user has requested recommendations wi
thin the last 5 minutes\n    51\t    When the app requests GET /matching/recomme
ndations again\n    52\t    Then the API should return cached results from Redis
within 100ms\n    53\t\n    54\t  Scenario: No eligible candidates\n    55\t    
Given the user is in a remote area with no other users\n    56\t    When the app
requests recommendations\n    57\t    Then the API should return an empty list\n
   58\t    And the app should show an empty state "附近暂无可推荐用户"\n    59\t
```\n    60\t\n    61\t#### 技术要求\n    62\t\n    63\t- 接口：`GET /matching/r
ecommendations?mode=nearby&limit=20`。\n    64\t- `mode` 参数：`nearby`（默认 50
km）、`city`（同城市，根据用户当前位置行政边界）、`nationwide`（全国）。\n    65
\t- 过滤条件：\n    66\t  1. 目标用户资料完整。\n    67\t  2. 目标用户性别符合当
前用户的 `preferred_gender`（若 `both` 则不限制）。\n    68\t  3. 当前用户性别符
合目标用户的 `preferred_gender`（双向匹配偏好）。\n    69\t  4. 目标用户在设定距
离范围内（Haversine 公式计算）。\n    70\t  5. 排除已存在 `match_actions`（Like 
或 Pass）的记录。\n    71\t  6. 排除已存在 `matches`（无论状态）的记录。\n    72
\t  7. 排除 `blocks` 表中的双向屏蔽记录（blocker_id 或 blocked_id 为当前用户）。
\n    73\t- 排序规则（综合得分）：\n    74\t  - 距离越近越优先（权重 50%）：`sco
re += (1 - distance / max_distance) * 0.5`\n    75\t  - 近期活跃用户优先（权重 3
0%）：根据 `last_active_at` 分档（24h 内 1.0，7 天内 0.7，30 天内 0.4，其余 0.1
）。\n    76\t  - 新用户曝光加成（权重 20%）：注册 7 天内加 0.2。\n    77\t- Red
is 缓存：`recommendations:{user_id}`，String（JSON）存储候选 `user_id` 列表和得
分，TTL 5 分钟。\n    78\t- 缓存击穿防护：使用互斥锁（Redis `SET recommendations
:lock:{user_id} NX EX 10`）。\n    79\t\n    80\t---\n    81\t\n    82\t### US-M
ATCH-2：滑动操作\n    83\t\n    84\t**作为** 用户，  \n    85\t**我希望** 通过左
滑或右滑表达喜好，  \n    86\t**以便** 快速筛选感兴趣的人。\n    87\t\n    88\t#
### 验收标准\n    89\t\n    90\t> 以上验收标准的完整 Gherkin 用例详见 `tests/bdd
/matching.feature`。\n    91\t\n    92\t```gherkin\n    93\tFeature: Swipe Actio
ns\n    94\t  Scenario: User likes a profile\n    95\t    Given the user sees a 
card for target user "Bob"\n    96\t    When the user swipes right or taps the l
ike button\n    97\t    Then a like action should be recorded in match_actions\n
   98\t    And the target user should be removed from the user\'s recommendation
cache\n    99\t\n   100\t  Scenario: User passes a profile\n   101\t    Given th
e user sees a card for target user "Charlie"\n   102\t    When the user swipes l
eft or taps the pass button\n   103\t    Then a pass action should be recorded i
n match_actions\n   104\t    And the target user should be removed from the user
\'s recommendation cache\n   105\t\n   106\t  Scenario: Duplicate swipe is idemp
otent\n   107\t    Given the user has already liked target user "Bob"\n   108\t 
  When the user likes "Bob" again within 1 second\n   109\t    Then the API shou
ld return success without creating a duplicate match_actions record\n   110\t   
And no duplicate match should be created\n   111\t\n   112\t  Scenario: User can
not swipe themselves\n   113\t    Given the user sees their own profile in the c
andidate list by accident\n   114\t    When the user tries to like themselves\n 
 115\t    Then the API should reject with code 1000\n   116\t```\n   117\t\n   1
18\t#### 技术要求\n   119\t\n   120\t- 接口：`POST /matching/actions`（`target_i
d`, `action: \'like\' | \'pass\'`）。\n   121\t- 幂等性：`match_actions` 表对 `(
actor_id, target_id)` 建立唯一索引，重复插入被数据库拦截，后端捕获异常返回成功。
\n   122\t- 自身滑动校验：若 `target_id == actor_id`，返回 `code=1000`。\n   123
\t- 滑动后清除 Redis 缓存 `recommendations:{user_id}`，确保下次请求不再包含该用
户。\n   124\t- `match_actions` 表字段：`id`, `actor_id`, `target_id`, `action`
（`like`/`pass`/`super_like`）, `created_at`。\n   125\t\n   126\t---\n   127\t\
n   128\t### US-MATCH-3：双向匹配成功\n   129\t\n   130\t**作为** 用户，  \n   1
31\t**我希望** 当我和对方互相喜欢时自动配对成功，  \n   132\t**以便** 立即开始聊
天。\n   133\t\n   134\t#### 验收标准\n   135\t\n   136\t> 以上验收标准的完整 Gh
erkin 用例详见 `tests/bdd/matching.feature`。\n   137\t\n   138\t```gherkin\n   
139\tFeature: Mutual Like Creates Match\n   140\t  Scenario: Mutual like trigger
s match and conversation\n   141\t    Given Alice liked Bob\n   142\t    And Bob
has not yet liked Alice\n   143\t    When Bob likes Alice\n   144\t    Then a ma
tch record should be created with status "matched"\n   145\t    And a conversati
on should be auto-created for Alice and Bob\n   146\t    And both users should r
eceive a match notification via Socket.io\n   147\t    And both users should see
the new match in their match list\n   148\t\n   149\t  Scenario: One-way like re
mains pending\n   150\t    Given Alice liked Bob\n   151\t    And Bob has not li
ked Alice\n   152\t    When the system checks for mutual like\n   153\t    Then 
no match record should exist\n   154\t    And no conversation should be created\
n   155\t```\n   156\t\n   157\t#### 技术要求\n   158\t\n   159\t- 匹配判定：在 
`POST /matching/actions` 处理 `like` 时，查询 `match_actions` 是否存在 `actor_id
= target_id AND target_id = actor_id AND action = \'like\'`。\n   160\t- 若存在
双向喜欢：\n   161\t  1. 创建 `matches` 记录：`user_a_id`（较小的 UUID）、`user_
b_id`（较大的 UUID），`status = \'matched\'`。\n   162\t  2. 创建 `conversations
` 记录，关联 `match_id`。\n   163\t  3. 创建两条 `conversation_participants` 记
录，`last_read_message_id = NULL`。\n   164\t  4. 通过 Socket.io 向双方 emit `ma
tch:created` 事件，Payload 包含 `match_id`、`conversation_id`、`matched_user` 资
料摘要。\n   165\t  5. 若接收方离线，触发推送通知（极光/Firebase）。\n   166\t- 
`matches` 表字段：`id`, `user_a_id`, `user_b_id`, `status`（`pending`/`matched`/
`unmatched`），`created_at`, `updated_at`。\n   167\t- 事务：匹配创建必须使用数
据库事务，确保 `matches` + `conversations` + `conversation_participants` 原子性
。\n   168\t\n   169\t---\n   170\t\n   171\t### US-MATCH-4：匹配列表查看\n   17
2\t\n   173\t**作为** 已匹配成功的用户，  \n   174\t**我希望** 查看所有匹配成功
的人列表，  \n   175\t**以便** 快速找到想聊天的人。\n   176\t\n   177\t#### 验收
标准\n   178\t\n   179\t> 以上验收标准的完整 Gherkin 用例详见 `tests/bdd/matchin
g.feature`。\n   180\t\n   181\t```gherkin\n   182\tFeature: Match List\n   183\
t  Scenario: User sees their matches\n   184\t    Given Alice has matched with B
ob and Charlie\n   185\t    When Alice opens the match list\n   186\t    Then sh
e should see Bob and Charlie ordered by match creation time descending\n   187\t
   And each item should show the matched user\'s nickname, avatar, and match tim
e\n   188\t```\n   189\t\n   190\t#### 技术要求\n   191\t\n   192\t- 接口：`GET 
/matches?limit=20&cursor=`（游标分页）。\n   193\t- 查询当前用户参与的 `matches`
中 `status = \'matched\'` 的记录，按 `created_at DESC` 排序。\n   194\t- 返回体
包含对方用户的 `user_id`, `nickname`, `avatar_url`, `match_id`, `conversation_id
`, `matched_at`。\n   195\t- 列表中的用户若被拉黑，则自动过滤（`unmatched` 状态
）。\n   196\t\n   197\t---\n   198\t\n   199\t## 3. 非功能需求\n   200\t\n   20
1\t### 性能\n   202\t\n   203\t- 首页推荐卡片加载时间 < 800ms（P95）。\n   204\t
- 滑动操作响应时间 < 200ms。\n   205\t- 双向匹配判定和创建 < 300ms（含事务）。\n
  206\t- Redis 缓存命中时响应时间 < 100ms。\n   207\t\n   208\t### 安全\n   209\
t\n   210\t- 推荐池必须排除资料不完整用户（后端通过 JOIN `user_profiles` 的 `is_
complete` 字段过滤）。\n   211\t- 用户不能看到自己（`target_id != actor_id`）。\
n   212\t- 已拉黑用户双向不可见、不可匹配。\n   213\t- 匹配成功后若一方拉黑另一
方，`match.status` 应更新为 `unmatched`（拉黑模块实现，Matching 侧读取时过滤）。
\n   214\t\n   215\t### 兼容性\n   216\t\n   217\t- 距离计算在 MySQL 中使用 Have
rsine 公式，不依赖 Spatial 扩展，保证跨环境兼容。\n   218\t- 推荐列表为空时 Mobi
le 需展示空状态引导页。\n   219\t\n   220\t---\n   221\t\n   222\t## 4. 里程碑\n
  223\t\n   224\t| 阶段      | 目标                                             
             | 预计工期 |\n   225\t| --------- | -------------------------------
------------------------------- | -------- |\n   226\t| M-MATCH-1 | 推荐 API + R
edis 缓存 + Haversine 排序 + Mobile 首页卡片       | 3 天     |\n   227\t| M-MAT
CH-2 | 滑动 API + 幂等性 + 双向匹配判定 + Match/Conversation 自动创建 | 2 天    
|\n   228\t| M-MATCH-3 | Socket.io 匹配通知 + 推送 + 匹配列表 + BDD 测试闭环    
       | 2 天     |\n   229\t\n   230\t**总预计工期：7 天。**\n   231\t\n   232\
t---\n   233\t\n   234\t## 5. 风险与回滚策略\n   235\t\n   236\t| 风险          
                                      | 影响 | 缓解措施                         
                        |\n   237\t| -------------------------------------------
--------- | ---- | --------------------------------------------------------- |\n
  238\t| 推荐算法性能不足，首页加载超过 800ms                 | 高   | 加 Redis 
缓存、建立复合索引、MVP 阶段限制候选池为 1000 人 |\n   239\t| 并发滑动导致重复创
建 Match                           | 中   | 数据库唯一索引 `(user_a_id, user_b_i
d)` + 事务处理        |\n   240\t| 缓存与数据库状态不一致（用户被拉黑后仍出现在
推荐中） | 中   | 拉黑操作时主动清除双方 Redis 推荐缓存                     |\n 
 241\t| 地理位置隐私合规风险                                 | 中   | 用户协议中
明确告知位置用途，提供「全国」模式淡化精确位置  |\n   242\t\n   243\t### 回滚策
略\n   244\t\n   245\t- **推荐算法故障**：关闭 Redis 缓存，直接查询数据库（响应
变慢但可用）。\n   246\t- **匹配创建事务死锁**：重试 3 次，失败后返回临时错误并
记录日志，不丢失滑动记录。\n   247\t- **Socket.io 通知失败**：降级为仅推送通知（
若用户离线）或记录待补发事件。\n',
        message='247 lines read from file starting from line 1. Total lines in f
ile: 247. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='/T/dev-issue-55-815001100 && find .maestro -name \\"*.yaml\\
" -o -name \\"*.yml\\" | sort"}'
)
StatusUpdate(
    context_usage=0.024929046630859375,
    context_tokens=6535,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=3463,
        output=378,
        input_cache_read=3072,
        input_cache_creation=0
    ),
    message_id='chatcmpl-BwvACltKaowoQGcMwH8UyLDi',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_GNbXXCkJzzSkZqFIBAy03gYb',
    return_value=ToolReturnValue(
        is_error=False,
        output='.maestro/adb_flows/login.yaml\n.maestro/adb_flows/onboarding.yam
l\n.maestro/adb_flows/profile_setup.yaml\n.maestro/config.yaml\n.maestro/flows/a
uth/login.yaml\n.maestro/flows/devtools/edit_server_url.yaml\n.maestro/flows/dev
tools/open_drawer.yaml\n.maestro/flows/matching/swipe_recommendations.yaml\n.mae
stro/flows/matching/view_match_list.yaml\n.maestro/flows/profile/setup.yaml\n.ma
estro/includes/login-as-test-user.yaml\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=2)
ThinkPart(
    type='think',
    think='让我继续读取 `.maestro/README.md` 和现有的 flow 文件，了解 E2E 规范和
已有测试范围。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_J4jLFZJ1FU6Oe0bdngaWOkJH',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/README.md"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_TjXlMPDOzi3hYVMtNxkhgMJk',
    function=FunctionBody(name='ReadFile', arguments='{"path": "/'),
    extras=None
)
ToolResult(
    tool_call_id='tool_J4jLFZJ1FU6Oe0bdngaWOkJH',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# Mobile E2E 测试（Maestro + ADB Fallback）\n     2\t\n 
   3\t本目录包含 Social World Mobile 端的端到端 UI 自动化测试。\n     4\t\n     
5\t## 目录结构\n     6\t\n     7\t```\n     8\t.maestro/\n     9\t├── flows/    
         # Maestro 标准测试流（推荐）\n    10\t│   ├── login.yaml      # 登录 UI
流（无需后端，可独立运行）\n    11\t│   └── profile_setup.yaml  # 完善资料完整流
（需要后端服务）\n    12\t├── adb_flows/          # ADB fallback 测试流（当 Maes
tro 不可用时）\n    13\t│   ├── login.yaml\n    14\t│   ├── onboarding.yaml\n   
15\t│   └── profile_setup.yaml\n    16\t└── scripts/\n    17\t    ├── run-e2e.sh
     # 统一 E2E 入口脚本（自动处理 emulator / build / install / 端口检查）\n    
18\t    ├── adb_runner.py   # ADB fallback runner（uiautomator dump + adb input
）\n    19\t    └── adb_flow.py     # 旧版 adb runner（已弃用，请使用 adb_runner
.py）\n    20\t```\n    21\t\n    22\t## 快速开始\n    23\t\n    24\t### 前置要
求\n    25\t\n    26\t- Android Emulator 已启动（推荐 API 34，如 `Pixel_8_Pro`）
\n    27\t- Maestro 1.39.13+ 已安装\n    28\t- **端口 7001 必须空闲**（Maestro 
硬编码端口，会与 Dashboard `new_world_ws` 冲突）\n    29\t\n    30\t### 常用命令
\n    31\t\n    32\t```bash\n    33\t# 仅跑 login 流（纯 UI，无需后端）\n    34\
tmake test-mobile-e2e-login\n    35\t\n    36\t# 跑所有 flows（需要后端 + emulat
or）\n    37\tmake test-mobile-e2e\n    38\t\n    39\t# CI 模式：自动构建 APK、
启动后端、运行测试\n    40\tmake test-mobile-e2e-ci\n    41\t\n    42\t# ADB fal
lback（当 Maestro 不可用时）\n    43\tmake test-mobile-e2e-adb\n    44\t```\n   
45\t\n    46\t### 手动运行\n    47\t\n    48\t```bash\n    49\t# 使用统一脚本\n 
  50\tbash .maestro/scripts/run-e2e.sh --flow login\n    51\tbash .maestro/scrip
ts/run-e2e.sh --flow login --build\n    52\tbash .maestro/scripts/run-e2e.sh --c
i\n    53\t\n    54\t# 直接使用 maestro\n    55\tmaestro --device emulator-5556 
test .maestro/flows/login.yaml\n    56\t```\n    57\t\n    58\t## 环境说明\n    
59\t\n    60\t### login.yaml\n    61\t\n    62\t- **不依赖后端**：仅验证登录页面
的 UI 交互（输入手机号、验证码、点击按钮）。\n    63\t- 适合作为日常开发和 CI 的
快速回归测试。\n    64\t\n    65\t### profile_setup.yaml\n    66\t\n    67\t- **
依赖后端**：登录后会调用 `/auth/login` 和 `/auth/me`，需要后端返回 `isNewUser=tr
ue` 或 `profileComplete=false` 才能跳转到完善资料页。\n    68\t- 运行前请确保后
端已启动：\n    69\t  ```bash\n    70\t  make e2e-env-up   # 启动 mysql + redis 
+ api\n    71\t  make test-mobile-e2e\n    72\t  make e2e-env-down # 测试结束后
关闭\n    73\t  ```\n    74\t\n    75\t## 故障排查\n    76\t\n    77\t### `io.gr
pc.StatusRuntimeException: INTERNAL: http2 exception`\n    78\t\n    79\t**原因*
*：端口 7001 被其他进程占用（通常是 `new_world_ws` Dashboard）。  \n    80\t**解
决**：终止占用进程后再运行测试。\n    81\t\n    82\t```bash\n    83\tlsof -ti :7
001 | xargs kill\n    84\t```\n    85\t\n    86\t### `adb shell input text` 无法
输入到 Flutter 自定义输入框\n    87\t\n    88\t**原因**：ADB 的 `input text` 依
赖系统 IME，某些自定义 `TextField` 不响应。  \n    89\t**解决**：Maestro 的 `inp
utText` 可以正常工作；如果只能用 ADB fallback，目前只能绕过文字输入步骤，仅做 ta
p / assert。\n    90\t\n    91\t### Emulator 未找到\n    92\t\n    93\t**原因**
：脚本通过 `adb devices` 查找 emulator，没有检测到已启动的设备。  \n    94\t**解
决**：先启动模拟器，例如：\n    95\t\n    96\t```bash\n    97\temulator -avd Pix
el_8_Pro\n    98\t```\n',
        message='98 lines read from file starting from line 1. Total lines in fi
le: 98. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100/.maestro/flows/auth/login.yaml"}'
)
ToolCall(
    type='function',
    id='tool_5RPEgB3c9JuSSoxBCZ2Dndpk',
    function=FunctionBody(name='ReadFile', arguments='{"path":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_TjXlMPDOzi3hYVMtNxkhgMJk',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @flow\n     2\t# @name: auth/login\n     3\t# @display
Name: 手机号登录\n     4\t# @module: auth\n     5\t# @dependencies: []\n     6\t
# @regression: smoke\n     7\t# @author: agent\n     8\t# @lastUpdated: 2025-04-
16\n     9\t# @syncWith: tests/bdd/auth_login_phone_number_login.feature\n    10
\t\n    11\t# 运行方式: maestro test .maestro/flows/auth/login.yaml\n    12\t\n 
  13\tappId: com.example.mobile\n    14\t---\n    15\t- launchApp:\n    16\t    
clearState: true\n    17\t- runFlow: includes/login-as-test-user.yaml\n    18\t\
n    19\t# 断言：登录成功后离开登录页（新用户进入完善资料，老用户进入首页）\n   
20\t- assertNotVisible: "手机号一键登录"\n    21\t- assertVisible: "完善资料"\n 
  22\t  optional: true\n    23\t- assertVisible:\n    24\t    id: "tab_home"\n  
 25\t    optional: true\n',
        message='25 lines read from file starting from line 1. Total lines in fi
le: 25. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/flows/profile/setup.yaml"}'
)
ToolCall(
    type='function',
    id='tool_rGaOWeFMG1gIetreMBsvqcC3',
    function=FunctionBody(name='ReadFile', arguments='{"path":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_5RPEgB3c9JuSSoxBCZ2Dndpk',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @flow\n     2\t# @name: profile/setup\n     3\t# @disp
layName: 完善资料\n     4\t# @module: profile\n     5\t# @dependencies: [e2e-bac
kend]\n     6\t# @regression: full\n     7\t# @author: agent\n     8\t# @lastUpd
ated: 2025-04-16\n     9\t# @syncWith: tests/bdd/profile_complete_minimum_profil
e.feature\n    10\t\n    11\t# 运行方式: maestro test .maestro/flows/profile/set
up.yaml\n    12\t\n    13\tappId: com.example.mobile\n    14\t---\n    15\t- lau
nchApp:\n    16\t    clearState: true\n    17\t\n    18\t# 先完成登录（clearStat
e 会清除本地 token）\n    19\t- runFlow: includes/login-as-test-user.yaml\n    2
0\t\n    21\t# 等待页面跳转后断言完善资料页\n    22\t- assertNotVisible: "手机号
一键登录"\n    23\t- assertVisible: "完善资料"\n    24\t\n    25\t# ===== Step 1
: 上传头像 =====\n    26\t- assertVisible: "完善资料"\n    27\t- assertVisible: 
"上传头像"\n    28\t- tapOn:\n    29\t    point: "50%,45%"\n    30\t- assertVisi
ble: "从相册选择"\n    31\t- tapOn: "从相册选择"\n    32\t- assertVisible: "头像
已选择"\n    33\t- tapOn: "下一步"\n    34\t\n    35\t# ===== Step 2: 基础资料 =
====\n    36\t- assertVisible: "基础资料"\n    37\t\n    38\t# 输入昵称\n    39\
t- tapOn: "昵称"\n    40\t- inputText: "E2E User"\n    41\t\n    42\t# 选择性别\
n    43\t- tapOn: "男"\n    44\t\n    45\t# 选择出生日期\n    46\t- tapOn: "出生
日期"\n    47\t- tapOn: "OK"\n    48\t\n    49\t# 进入下一步\n    50\t- tapOn: "
下一步"\n    51\t\n    52\t# ===== Step 3: 位置授权 =====\n    53\t- assertVisib
le: "位置授权"\n    54\t- tapOn: "允许位置授权"\n    55\t- assertVisible: "位置
已获取"\n    56\t\n    57\t# 完成资料设置\n    58\t- tapOn: "完成"\n',
        message='58 lines read from file starting from line 1. Total lines in fi
le: 58. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/flows/matching/swipe_recommendations.yaml"}'
)
ToolCall(
    type='function',
    id='tool_CvYYVEMlE8fK4qwCIqKqgDM4',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_rGaOWeFMG1gIetreMBsvqcC3',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @flow\n     2\t# @name: matching/swipe-recommendations
\n     3\t# @displayName: 推荐卡片滑动\n     4\t# @module: matching\n     5\t# @
dependencies: [auth/login, e2e-backend]\n     6\t# @regression: full\n     7\t# 
@author: agent\n     8\t# @lastUpdated: 2026-04-22\n     9\t# @syncWith: tests/b
dd/matching.feature\n    10\t\n    11\tappId: com.example.mobile\n    12\t---\n 
  13\t- launchApp:\n    14\t    clearState: true\n    15\t- runFlow: includes/lo
gin-as-test-user.yaml\n    16\t\n    17\t# 断言首页加载（通过模式切换 Tab 判断）
\n    18\t- assertVisible:\n    19\t    id: "mode_tab_nearby"\n    20\t- assertV
isible: "附近"\n    21\t\n    22\t# 切换同城模式\n    23\t- tapOn:\n    24\t    
id: "mode_tab_city"\n    25\t- assertVisible: "同城"\n    26\t\n    27\t# 切换全
国模式\n    28\t- tapOn:\n    29\t    id: "mode_tab_nationwide"\n    30\t- asser
tVisible: "全国"\n    31\t\n    32\t# 切回附近\n    33\t- tapOn:\n    34\t    id
: "mode_tab_nearby"\n    35\t\n    36\t# 如果有卡片，测试左滑（Pass）手势\n    3
7\t- swipe:\n    38\t    direction: LEFT\n    39\t    optional: true\n    40\t\n
   41\t# 如果有卡片，测试右滑（Like）手势\n    42\t- swipe:\n    43\t    directi
on: RIGHT\n    44\t    optional: true\n    45\t\n    46\t# 如果有卡片，点击 Pass
按钮\n    47\t- tapOn:\n    48\t    id: "pass_button"\n    49\t    optional: tru
e\n    50\t\n    51\t# 如果有卡片，点击 Like 按钮\n    52\t- tapOn:\n    53\t   
id: "like_button"\n    54\t    optional: true\n    55\t\n    56\t# 验证卡片区域
是否存在（optional，后端可能没有返回卡片）\n    57\t- assertVisible:\n    58\t  
 id: "card_stack"\n    59\t    optional: true\n    60\t\n    61\t# 如果进入空状
态，验证切换全国模式按钮可见\n    62\t- assertVisible:\n    63\t    id: "switch_
nationwide_button"\n    64\t    optional: true\n    65\t\n    66\t# 如果触发了双
向匹配，验证配对成功弹窗\n    67\t- assertVisible:\n    68\t    id: "match_succe
ss_dialog"\n    69\t    optional: true\n    70\t\n    71\t# 如果弹窗出现，点击继
续滑动关闭弹窗\n    72\t- tapOn:\n    73\t    id: "match_dialog_continue_button"
\n    74\t    optional: true\n',
        message='74 lines read from file starting from line 1. Total lines in fi
le: 74. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='001100/.maestro/flows/matching/view_match_list.yaml"}'
)
ToolCall(
    type='function',
    id='tool_j0g4CDpH164w17f44t678otE',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_CvYYVEMlE8fK4qwCIqKqgDM4',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @flow\n     2\t# @name: matching/view-match-list\n    
3\t# @displayName: 查看匹配列表\n     4\t# @module: matching\n     5\t# @depende
ncies: [auth/login, e2e-backend]\n     6\t# @regression: full\n     7\t# @author
: agent\n     8\t# @lastUpdated: 2026-04-22\n     9\t# @syncWith: tests/bdd/matc
hing.feature\n    10\t\n    11\tappId: com.example.mobile\n    12\t---\n    13\t
- launchApp:\n    14\t    clearState: true\n    15\t- runFlow: includes/login-as
-test-user.yaml\n    16\t\n    17\t# 切换到消息 Tab（使用稳定 id 断言）\n    18\
t- tapOn:\n    19\t    id: "tab_messages"\n    20\t\n    21\t# 断言页面标题\n   
22\t- assertVisible: "消息"\n    23\t\n    24\t# 断言空状态或列表\n    25\t- ass
ertVisible:\n    26\t    text: "还没有匹配成功的好友"\n    27\t    optional: tru
e\n    28\t- assertVisible:\n    29\t    text: "继续滑动，发现感兴趣的人吧"\n   
30\t    optional: true\n    31\t\n    32\t# 如果列表中有数据，验证列表项可见\n  
 33\t- assertVisible:\n    34\t    id: "match_list_item"\n    35\t    optional: 
true\n    36\t\n    37\t# 测试下拉刷新\n    38\t- swipe:\n    39\t    direction:
DOWN\n',
        message='39 lines read from file starting from line 1. Total lines in fi
le: 39. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='/dev-issue-55-815001100/.maestro/includes/login-as-test-user
.yaml"}'
)
StatusUpdate(
    context_usage=0.068878173828125,
    context_tokens=18056,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=11656,
        output=378,
        input_cache_read=6400,
        input_cache_creation=0
    ),
    message_id='chatcmpl-O9Rfa80IEdFWtFXPBL18vc4n',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_j0g4CDpH164w17f44t678otE',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @include\n     2\t# @name: login-as-test-user\n     3\
t# @displayName: 使用测试用户登录\n     4\t# @description: 输入测试手机号和固定
验证码，完成登录前置步骤\n     5\t# @lastUpdated: 2025-04-16\n     6\t\n     7\t
appId: com.example.mobile\n     8\t---\n     9\t\n    10\t- assertVisible: "欢迎
使用 Social World"\n    11\t- assertVisible: "手机号一键登录"\n    12\t\n    13\
t# 输入手机号\n    14\t- tapOn: "手机号"\n    15\t- inputText: "13800138000"\n  
 16\t\n    17\t# 获取验证码\n    18\t- tapOn: "获取验证码"\n    19\t# 注：MVP 阶
段 _onSendCode 只启动本地倒计时，不调用后端 API\n    20\t# 原 assertVisible: "60
s" 因设备性能波动导致秒数变化而失败，此处不再严格断言\n    21\t\n    22\t# 输入
验证码（MVP 固定验证码为任意 6 位）\n    23\t- tapOn: "验证码"\n    24\t- inputT
ext: "123456"\n    25\t\n    26\t# 收起软键盘，避免遮挡登录按钮\n    27\t- hideK
eyboard\n    28\t\n    29\t# 点击登录\n    30\t- tapOn: "登录"\n',
        message='30 lines read from file starting from line 1. Total lines in fi
le: 30. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=3)
ThinkPart(
    type='think',
    think='让我继续检查其他 flow 文件、adb_flows 和 devtools flow，以及检查 mobi
le app 的代码，了解 UI 元素 ID 和页面结构。\n\n同时我需要设置 todo list 来跟踪任
务进度。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_qEFENq8lpX73H8KSgEWsPtI7',
    function=FunctionBody(
        name='SetTodoList',
        arguments='{"todos": [{"title":"阅读项目地图与E2E规范","status":"done"},
{"title":"扫描已有flow并分析覆盖范围","status":"in_progress"},{"title":"检查mobi
le代码中的UI元素和页面结构","status":"pending"},{"title":"修复includes格式问题",
"status":"pending"},{"title":"执行Maestro测试","status":"pending"},{"title":"分
析失败根因并修复flow","status":"pending"},{"title":"生成E2E测试报告","status":"p
ending"}]}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_Papj72JuPlUhymm3mbQaocbn',
    function=FunctionBody(name='ReadFile', arguments='{"path": "/'),
    extras=None
)
ToolResult(
    tool_call_id='tool_qEFENq8lpX73H8KSgEWsPtI7',
    return_value=ToolReturnValue(
        is_error=False,
        output='Todo list updated',
        message='Todo list updated',
        display=[
            TodoDisplayBlock(
                type='todo',
                items=[
                    TodoDisplayItem(
                        title='阅读项目地图与E2E规范',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='扫描已有flow并分析覆盖范围',
                        status='in_progress'
                    ),
                    TodoDisplayItem(
                        title='检查mobile代码中的UI元素和页面结构',
                        status='pending'
                    ),
                    TodoDisplayItem(
                        title='修复includes格式问题',
                        status='pending'
                    ),
                    TodoDisplayItem(title='执行Maestro测试', status='pending'),
                    TodoDisplayItem(
                        title='分析失败根因并修复flow',
                        status='pending'
                    ),
                    TodoDisplayItem(title='生成E2E测试报告', status='pending')
                ]
            )
        ],
        extras=None
    )
)
ToolCallPart(
    arguments_part='var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100/.maestro/flows/devtools/edit_server_url.yaml"}'
)
ToolCall(
    type='function',
    id='tool_zThmFBjaVo6f30oVqr1hIP5d',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_Papj72JuPlUhymm3mbQaocbn',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @flow\n     2\t# @name: devtools/edit_server_url\n    
3\t# @displayName: 编辑服务器 Base URL\n     4\t# @module: devtools\n     5\t# @
dependencies: [devtools/open_drawer]\n     6\t# @regression: regression\n     7\
t# @author: agent\n     8\t# @lastUpdated: 2026-04-09\n     9\t# @syncWith: test
s/bdd/dev_tools_drawer.feature\n    10\t\n    11\t# 运行方式: maestro test .maes
tro/flows/devtools/edit_server_url.yaml\n    12\t# 说明：验证 ServerUrlTool 可以
正确编辑和保存 Base URL\n    13\t\n    14\tappId: com.example.mobile\n    15\t--
-\n    16\t- launchApp:\n    17\t    clearState: true\n    18\t\n    19\t# 打开 
DevTools 抽屉\n    20\t- tapOn:\n    21\t    id: "dev_tools_toggle_button"\n    
22\t- assertVisible: "服务器地址"\n    23\t\n    24\t# 点击输入框获得焦点（自动
全选旧内容）\n    25\t- tapOn:\n    26\t    id: "server_base_url_field"\n    27\
t\n    28\t# 输入新 URL（旧内容已被全选，新输入自动替换）\n    29\t- inputText: 
"http://192.168.1.100:3000/api/v1"\n    30\t\n    31\t# 保存\n    32\t- tapOn: "
保存配置"\n    33\t- assertVisible: "保存成功，新配置已生效"\n    34\t\n    35\t
# 验证二次编辑也能正确替换\n    36\t- tapOn:\n    37\t    id: "server_base_url_f
ield"\n    38\t- inputText: "http://example.com:8080/api/v2"\n    39\t- tapOn: "
保存配置"\n    40\t- assertVisible: "保存成功，新配置已生效"\n    41\t\n    42\t
# 最终恢复为正确的 E2E 服务器地址，避免污染后续 flow\n    43\t- tapOn:\n    44\t
   id: "server_base_url_field"\n    45\t- eraseText\n    46\t- inputText: "http:
//10.0.2.2:3001/api/v1"\n    47\t- tapOn: "保存配置"\n    48\t- assertVisible: "
保存成功，新配置已生效"\n',
        message='48 lines read from file starting from line 1. Total lines in fi
le: 48. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(arguments_part='maestro/flows/devtools/open_drawer.yaml"}')
ToolCall(
    type='function',
    id='tool_jyEzuTk9KDe4qUusCJEPjbwr',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_zThmFBjaVo6f30oVqr1hIP5d',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @flow\n     2\t# @name: devtools/open_drawer\n     3\t
# @displayName: 开发工具抽屉定位\n     4\t# @module: devtools\n     5\t# @depend
encies: []\n     6\t# @regression: smoke\n     7\t# @author: agent\n     8\t# @l
astUpdated: 2025-04-17\n     9\t# @syncWith: tests/bdd/dev_tools_drawer.feature\
n    10\t\n    11\t# 运行方式: maestro test .maestro/flows/devtools/open_drawer.
yaml\n    12\t# 说明：验证 DevTools 悬浮按钮可被 Maestro 定位并点击打开抽屉\n   
13\t\n    14\tappId: com.example.mobile\n    15\t---\n    16\t- launchApp:\n    
17\t    clearState: true\n    18\t\n    19\t# 断言：DevTools 悬浮按钮的 accessib
ility label 可见\n    20\t- assertVisible: "开发工具"\n    21\t\n    22\t# 点击
悬浮按钮\n    23\t- tapOn:\n    24\t    id: "dev_tools_toggle_button"\n    25\t\
n    26\t# 断言：抽屉已打开（ServerUrlTool 中的输入框提示可见）\n    27\t- asser
tVisible: "服务器地址"\n',
        message='27 lines read from file starting from line 1. Total lines in fi
le: 27. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/.maestro/adb_f
lows/login.yaml"}'
)
ToolCall(
    type='function',
    id='tool_9xH48AyMpMW9I0spEjVStPSx',
    function=FunctionBody(name='ReadFile', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_jyEzuTk9KDe4qUusCJEPjbwr',
    return_value=ToolOk(
        is_error=False,
        output='     1\tname: 登录流程\n     2\tsteps:\n     3\t  - action: laun
chApp\n     4\t    wait: 2\n     5\t\n     6\t  - action: assertVisible\n     7\
t    contentDesc: "欢迎使用 Social World"\n     8\t\n     9\t  - action: assertV
isible\n    10\t    contentDesc: "手机号一键登录"\n    11\t\n    12\t  - action:
tap\n    13\t    contentDesc: "手机号"\n    14\t\n    15\t  - action: inputText\
n    16\t    text: "13800138000"\n    17\t\n    18\t  - action: tap\n    19\t   
contentDesc: "获取验证码"\n    20\t    wait: 1\n    21\t\n    22\t  - action: ta
p\n    23\t    contentDesc: "验证码"\n    24\t\n    25\t  - action: inputText\n 
  26\t    text: "123456"\n    27\t\n    28\t  - action: tap\n    29\t    content
Desc: "登录"\n    30\t    wait: 2\n    31\t\n    32\t  # 登录成功后通常跳转首页
，可根据实际项目调整断言\n    33\t  # - action: assertVisible\n    34\t  #   tex
t: "首页"\n    35\t  #   timeout: 5\n',
        message='35 lines read from file starting from line 1. Total lines in fi
le: 35. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/d
ev-issue-55-815001100/.maestro/adb_flows/profile_setup.yaml"}'
)
ToolCall(
    type='function',
    id='tool_bp0wwH5dxaGDXyi2V7IfBGdw',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_9xH48AyMpMW9I0spEjVStPSx',
    return_value=ToolOk(
        is_error=False,
        output='     1\tname: 完善资料流程\n     2\tsteps:\n     3\t  - action: 
launchApp\n     4\t    wait: 2\n     5\t\n     6\t  - action: assertVisible\n   
 7\t    text: "完善资料"\n     8\t\n     9\t  # ===== Step 1: 上传头像 =====\n  
 10\t  - action: assertVisible\n    11\t    text: "上传头像"\n    12\t\n    13\t
 - action: tap\n    14\t    text: "点击上传头像"\n    15\t\n    16\t  - action: 
assertVisible\n    17\t    text: "从相册选择"\n    18\t\n    19\t  - action: tap
\n    20\t    text: "从相册选择"\n    21\t\n    22\t  - action: assertVisible\n 
  23\t    text: "头像已选择"\n    24\t\n    25\t  - action: tap\n    26\t    tex
t: "下一步"\n    27\t\n    28\t  # ===== Step 2: 基础资料 =====\n    29\t  - act
ion: assertVisible\n    30\t    text: "基础资料"\n    31\t\n    32\t  - action: 
tap\n    33\t    contentDesc: "昵称"\n    34\t\n    35\t  - action: inputText\n 
  36\t    text: "测试用户"\n    37\t\n    38\t  - action: tap\n    39\t    text:
"男"\n    40\t\n    41\t  - action: tap\n    42\t    contentDesc: "出生日期"\n  
 43\t\n    44\t  - action: tap\n    45\t    text: "确定"\n    46\t    wait: 1\n 
  47\t\n    48\t  - action: tap\n    49\t    text: "下一步"\n    50\t\n    51\t 
# ===== Step 3: 位置授权 =====\n    52\t  - action: assertVisible\n    53\t    t
ext: "位置授权"\n    54\t\n    55\t  - action: tap\n    56\t    text: "允许位置
授权"\n    57\t\n    58\t  - action: assertVisible\n    59\t    text: "位置已获
取"\n    60\t\n    61\t  - action: tap\n    62\t    text: "完成"\n    63\t    wa
it: 2\n    64\t\n    65\t  # 完成后应跳转到首页\n    66\t  # - action: assertVis
ible\n    67\t  #   text: "首页"\n    68\t  #   timeout: 5\n',
        message='68 lines read from file starting from line 1. Total lines in fi
le: 68. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000
gn/T/dev-issue-55-815001100 && find apps/mobile -type f -name \\"*.dart\\" | hea
d -50"}'
)
StatusUpdate(
    context_usage=0.084625244140625,
    context_tokens=22184,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=4264,
        output=465,
        input_cache_read=17920,
        input_cache_creation=0
    ),
    message_id='chatcmpl-L7N4p50JeQkh5y9qUeY9tyty',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_bp0wwH5dxaGDXyi2V7IfBGdw',
    return_value=ToolReturnValue(
        is_error=False,
        output='apps/mobile/test/dev_tools_widget_test.dart\napps/mobile/test/lo
gin_page_golden_test.dart\napps/mobile/test/dev_tools_golden_test.dart\napps/mob
ile/test/presentation/blocs/matching_bloc_test.dart\napps/mobile/test/presentati
on/blocs/match_list_cubit_test.dart\napps/mobile/test/dev_tools_cubit_test.dart\
napps/mobile/lib/core/di/di.dart\napps/mobile/lib/core/constants/error_codes.dar
t\napps/mobile/lib/core/constants/error_messages.dart\napps/mobile/lib/core/netw
ork/auth_interceptor.dart\napps/mobile/lib/core/network/dio_client.dart\napps/mo
bile/lib/core/network/matching_socket_client.dart\napps/mobile/lib/core/storage/
secure_storage.dart\napps/mobile/lib/models/generated/matching/test/record_actio
n200_response_test.dart\napps/mobile/lib/models/generated/matching/test/recommen
dation_card_test.dart\napps/mobile/lib/models/generated/matching/test/default_ap
i_test.dart\napps/mobile/lib/models/generated/matching/test/api_response_test.da
rt\napps/mobile/lib/models/generated/matching/test/match_created_event_matched_u
ser_test.dart\napps/mobile/lib/models/generated/matching/test/match_list_item_te
st.dart\napps/mobile/lib/models/generated/matching/test/swipe_action_response_te
st.dart\napps/mobile/lib/models/generated/matching/test/get_matches200_response_
all_of_data_test.dart\napps/mobile/lib/models/generated/matching/test/get_matche
s200_response_test.dart\napps/mobile/lib/models/generated/matching/test/match_cr
eated_event_test.dart\napps/mobile/lib/models/generated/matching/test/get_recomm
endations200_response_test.dart\napps/mobile/lib/models/generated/matching/test/
swipe_action_request_test.dart\napps/mobile/lib/models/generated/matching/lib/au
th/http_bearer_auth.dart\napps/mobile/lib/models/generated/matching/lib/auth/oau
th.dart\napps/mobile/lib/models/generated/matching/lib/auth/api_key_auth.dart\na
pps/mobile/lib/models/generated/matching/lib/auth/authentication.dart\napps/mobi
le/lib/models/generated/matching/lib/auth/http_basic_auth.dart\napps/mobile/lib/
models/generated/matching/lib/api_exception.dart\napps/mobile/lib/models/generat
ed/matching/lib/model/swipe_action_request.dart\napps/mobile/lib/models/generate
d/matching/lib/model/get_matches200_response.dart\napps/mobile/lib/models/genera
ted/matching/lib/model/api_response.dart\napps/mobile/lib/models/generated/match
ing/lib/model/get_recommendations200_response.dart\napps/mobile/lib/models/gener
ated/matching/lib/model/match_created_event_matched_user.dart\napps/mobile/lib/m
odels/generated/matching/lib/model/match_list_item.dart\napps/mobile/lib/models/
generated/matching/lib/model/recommendation_card.dart\napps/mobile/lib/models/ge
nerated/matching/lib/model/get_matches200_response_all_of_data.dart\napps/mobile
/lib/models/generated/matching/lib/model/match_created_event.dart\napps/mobile/l
ib/models/generated/matching/lib/model/swipe_action_response.dart\napps/mobile/l
ib/models/generated/matching/lib/model/record_action200_response.dart\napps/mobi
le/lib/models/generated/matching/lib/api/default_api.dart\napps/mobile/lib/model
s/generated/matching/lib/api_client.dart\napps/mobile/lib/models/generated/match
ing/lib/api_helper.dart\napps/mobile/lib/models/generated/matching/lib/api.dart\
napps/mobile/lib/models/generated/auth/test/login_response_test.dart\napps/mobil
e/lib/models/generated/auth/test/login_request_test.dart\napps/mobile/lib/models
/generated/auth/test/error_response_test.dart\napps/mobile/lib/models/generated/
auth/test/default_api_test.dart\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=4)
ThinkPart(
    type='think',
    think='我需要继续查看 mobile app 的 UI 代码，特别是登录页、资料设置页、首页
推荐卡片、匹配列表等页面的 Widget，以便了解实际的 UI 元素 ID 和文字。\n\n让我搜
索关键页面文件。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_Zn54nYL2Ez7CWSAMzOyvUWzH',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100 && find apps/mobile/lib -type f -name \\"*.dart\\" | 
grep -v generated | grep -v test | sort"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_FBhR5DAGuwy3xTkxahCM5GR6',
    function=FunctionBody(name='Shell', arguments='{"command":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_Zn54nYL2Ez7CWSAMzOyvUWzH',
    return_value=ToolReturnValue(
        is_error=False,
        output='apps/mobile/lib/app.dart\napps/mobile/lib/core/constants/error_c
odes.dart\napps/mobile/lib/core/constants/error_messages.dart\napps/mobile/lib/c
ore/di/di.dart\napps/mobile/lib/core/network/auth_interceptor.dart\napps/mobile/
lib/core/network/dio_client.dart\napps/mobile/lib/core/network/matching_socket_c
lient.dart\napps/mobile/lib/core/storage/secure_storage.dart\napps/mobile/lib/da
ta/models/match_summary.dart\napps/mobile/lib/data/models/recommendation_card.da
rt\napps/mobile/lib/data/models/swipe_result.dart\napps/mobile/lib/data/reposito
ries/auth_repository_impl.dart\napps/mobile/lib/data/repositories/dev_tools_repo
sitory_impl.dart\napps/mobile/lib/data/repositories/matching_repository_impl.dar
t\napps/mobile/lib/data/repositories/profile_repository_impl.dart\napps/mobile/l
ib/domain/exceptions/profile_incomplete_exception.dart\napps/mobile/lib/domain/r
epositories/auth_repository.dart\napps/mobile/lib/domain/repositories/dev_tools_
repository.dart\napps/mobile/lib/domain/repositories/matching_repository.dart\na
pps/mobile/lib/domain/repositories/profile_repository.dart\napps/mobile/lib/main
.dart\napps/mobile/lib/presentation/blocs/auth/auth_bloc.dart\napps/mobile/lib/p
resentation/blocs/auth/auth_event.dart\napps/mobile/lib/presentation/blocs/auth/
auth_state.dart\napps/mobile/lib/presentation/blocs/dev_tools/dev_tools_cubit.da
rt\napps/mobile/lib/presentation/blocs/dev_tools/dev_tools_state.dart\napps/mobi
le/lib/presentation/blocs/matching/match_list_cubit.dart\napps/mobile/lib/presen
tation/blocs/matching/match_list_state.dart\napps/mobile/lib/presentation/blocs/
matching/matching_bloc.dart\napps/mobile/lib/presentation/blocs/matching/matchin
g_event.dart\napps/mobile/lib/presentation/blocs/matching/matching_state.dart\na
pps/mobile/lib/presentation/blocs/profile/profile_bloc.dart\napps/mobile/lib/pre
sentation/blocs/profile/profile_event.dart\napps/mobile/lib/presentation/blocs/p
rofile/profile_state.dart\napps/mobile/lib/presentation/pages/home_page.dart\nap
ps/mobile/lib/presentation/pages/login_page.dart\napps/mobile/lib/presentation/p
ages/match_list_page.dart\napps/mobile/lib/presentation/pages/profile_edit_page.
dart\napps/mobile/lib/presentation/pages/profile_setup_page.dart\napps/mobile/li
b/presentation/widgets/dev_tools_drawer.dart\napps/mobile/lib/presentation/widge
ts/dev_tools_toggle_button.dart\napps/mobile/lib/presentation/widgets/dev_tools_
wrapper.dart\napps/mobile/lib/presentation/widgets/empty_recommendation_view.dar
t\napps/mobile/lib/presentation/widgets/server_url_tool.dart\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-is
sue-55-815001100 && grep -rl \\"tab_home\\\\|tab_messages\\\\|mode_tab_\\\\|pass
_button\\\\|like_button\\\\|card_stack\\\\|match_success_dialog\\\\|match_dialog
_continue_button\\\\|dev_tools_toggle_button\\\\|server_base_url_field\\\\|switc
h_nationwide_button\\\\|match_list_item\\" apps/mobile/lib --include=\\"*.dart\\
" | sort"}'
)
StatusUpdate(
    context_usage=0.09650421142578125,
    context_tokens=25298,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=3282,
        output=253,
        input_cache_read=22016,
        input_cache_creation=0
    ),
    message_id='chatcmpl-J3fjofqxzD6WORqpQzXT7EEY',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_FBhR5DAGuwy3xTkxahCM5GR6',
    return_value=ToolReturnValue(
        is_error=False,
        output='apps/mobile/lib/app.dart\napps/mobile/lib/models/generated/match
ing/lib/api.dart\napps/mobile/lib/presentation/pages/home_page.dart\napps/mobile
/lib/presentation/widgets/dev_tools_toggle_button.dart\napps/mobile/lib/presenta
tion/widgets/dev_tools_wrapper.dart\napps/mobile/lib/presentation/widgets/empty_
recommendation_view.dart\napps/mobile/lib/presentation/widgets/server_url_tool.d
art\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=5)
ThinkPart(
    type='think',
    think='让我读取关键页面文件，了解实际的 UI 元素和文字，以便验证 flow 的准确
性。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_tOEMFdz5UIZaheJeCidUshTZ',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/apps/mobile/lib/presentation/pages/login_page.dart"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_4dOcw0JtYuR8n3f2z2bRHxmr',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_tOEMFdz5UIZaheJeCidUshTZ',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:flutter_bloc/flutter_bloc.dart';\n     3\timport 'package:social_world_
design_system/social_world_design_system.dart';\n     4\timport '../blocs/auth/a
uth_bloc.dart';\n     5\timport '../blocs/auth/auth_event.dart';\n     6\timport
'../blocs/auth/auth_state.dart';\n     7\t\n     8\tclass LoginPage extends Stat
efulWidget {\n     9\t  const LoginPage({super.key});\n    10\t\n    11\t  @over
ride\n    12\t  State<LoginPage> createState() => _LoginPageState();\n    13\t}\
n    14\t\n    15\tclass _LoginPageState extends State<LoginPage> {\n    16\t  f
inal _phoneController = TextEditingController();\n    17\t  final _codeControlle
r = TextEditingController();\n    18\t  bool _isCountingDown = false;\n    19\t 
int _countdown = 60;\n    20\t  bool _isLoading = false;\n    21\t\n    22\t  @o
verride\n    23\t  void dispose() {\n    24\t    _phoneController.dispose();\n  
 25\t    _codeController.dispose();\n    26\t    super.dispose();\n    27\t  }\n
   28\t\n    29\t  void _startCountdown() {\n    30\t    setState(() {\n    31\t
     _isCountingDown = true;\n    32\t      _countdown = 60;\n    33\t    });\n 
  34\t\n    35\t    Future.doWhile(() async {\n    36\t      await Future.delaye
d(const Duration(seconds: 1));\n    37\t      if (!mounted) return false;\n    3
8\t      setState(() {\n    39\t        _countdown--;\n    40\t      });\n    41
\t      if (_countdown <= 0) {\n    42\t        setState(() {\n    43\t         
_isCountingDown = false;\n    44\t        });\n    45\t        return false;\n  
 46\t      }\n    47\t      return true;\n    48\t    });\n    49\t  }\n    50\t
\n    51\t  void _onSendCode() {\n    52\t    final phone = _phoneController.tex
t.trim();\n    53\t    if (phone.length != 11) {\n    54\t      ScaffoldMessenge
r.of(context).showSnackBar(\n    55\t        const SnackBar(content: Text('请输
入正确的手机号')),\n    56\t      );\n    57\t      return;\n    58\t    }\n    
59\t    // MVP fixed code; just start countdown for UX\n    60\t    _startCountd
own();\n    61\t  }\n    62\t\n    63\t  void _onLogin() {\n    64\t    final ph
one = _phoneController.text.trim();\n    65\t    final code = _codeController.te
xt.trim();\n    66\t    if (phone.length != 11 || code.length != 6) {\n    67\t 
    ScaffoldMessenger.of(context).showSnackBar(\n    68\t        const SnackBar(
content: Text('请输入手机号和6位验证码')),\n    69\t      );\n    70\t      retu
rn;\n    71\t    }\n    72\t    context.read<AuthBloc>().add(LoginRequested(phon
e: phone, code: code));\n    73\t  }\n    74\t\n    75\t  @override\n    76\t  W
idget build(BuildContext context) {\n    77\t    return Scaffold(\n    78\t     
backgroundColor: DesignTokens.background,\n    79\t      body: SafeArea(\n    80
\t        child: Padding(\n    81\t          padding: const EdgeInsets.all(Desig
nTokens.spacingMd),\n    82\t          child: Column(\n    83\t            cross
AxisAlignment: CrossAxisAlignment.stretch,\n    84\t            children: [\n   
85\t              const SizedBox(height: DesignTokens.spacingXl * 2),\n    86\t 
            const Text(\n    87\t                '欢迎使用 Social World',\n    8
8\t                style: TextStyle(\n    89\t                  fontSize: Design
Tokens.fontSizeHeadline,\n    90\t                  fontWeight: FontWeight.bold,
\n    91\t                  color: DesignTokens.textPrimary,\n    92\t          
     ),\n    93\t                textAlign: TextAlign.center,\n    94\t         
    ),\n    95\t              const SizedBox(height: DesignTokens.spacingLg),\n 
  96\t              Text(\n    97\t                '手机号一键登录',\n    98\t  
             style: TextStyle(\n    99\t                  fontSize: DesignTokens
.fontSizeBody,\n   100\t                  color: DesignTokens.textPrimary.withOp
acity(0.7),\n   101\t                ),\n   102\t                textAlign: Text
Align.center,\n   103\t              ),\n   104\t              const SizedBox(he
ight: DesignTokens.spacingXl * 2),\n   105\t              SwTextField(\n   106\t
               controller: _phoneController,\n   107\t                keyboardTy
pe: TextInputType.phone,\n   108\t                labelText: '手机号',\n   109\t
               hintText: '请输入11位手机号',\n   110\t                prefixIcon
: const Icon(Icons.phone),\n   111\t              ),\n   112\t              cons
t SizedBox(height: DesignTokens.spacingMd),\n   113\t              Row(\n   114\
t                children: [\n   115\t                  Expanded(\n   116\t     
              flex: 2,\n   117\t                    child: SwTextField(\n   118\
t                      controller: _codeController,\n   119\t                   
  keyboardType: TextInputType.number,\n   120\t                      labelText: 
'验证码',\n   121\t                      hintText: '请输入验证码',\n   122\t    
                 prefixIcon: const Icon(Icons.lock_outline),\n   123\t          
         ),\n   124\t                  ),\n   125\t                  const Sized
Box(width: DesignTokens.spacingMd),\n   126\t                  Expanded(\n   127
\t                    flex: 1,\n   128\t                    child: SizedBox(\n  
129\t                      height: 56,\n   130\t                      child: SwB
utton(\n   131\t                        text: _isCountingDown ? '$_countdown s' 
: '获取验证码',\n   132\t                        variant: SwButtonVariant.outlin
ed,\n   133\t                        isDisabled: _isCountingDown,\n   134\t     
                  onPressed: _isCountingDown ? null : _onSendCode,\n   135\t    
                 ),\n   136\t                    ),\n   137\t                  )
,\n   138\t                ],\n   139\t              ),\n   140\t              c
onst SizedBox(height: DesignTokens.spacingXl),\n   141\t              BlocListen
er<AuthBloc, AuthState>(\n   142\t                listener: (context, state) {\n
  143\t                  if (state is AuthAuthenticating) {\n   144\t           
        setState(() => _isLoading = true);\n   145\t                  } else if 
(state is AuthAuthenticated || state is AuthFailure) {\n   146\t                
   setState(() => _isLoading = false);\n   147\t                  }\n   148\t   
              if (state is AuthFailure) {\n   149\t                    ScaffoldM
essenger.of(context).showSnackBar(\n   150\t                      SnackBar(conte
nt: Text(state.message)),\n   151\t                    );\n   152\t             
    }\n   153\t                },\n   154\t                child: SwButton(\n   
155\t                  text: '登录',\n   156\t                  variant: SwButto
nVariant.filled,\n   157\t                  isLoading: _isLoading,\n   158\t    
             onPressed: _isLoading ? null : _onLogin,\n   159\t                )
,\n   160\t              ),\n   161\t            ],\n   162\t          ),\n   16
3\t        ),\n   164\t      ),\n   165\t    );\n   166\t  }\n   167\t}\n",
        message='167 lines read from file starting from line 1. Total lines in f
ile: 167. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/app
s/mobile/lib/presentation/pages/home_page.dart"}'
)
ToolCall(
    type='function',
    id='tool_ncDtdXCRuMyuSkY3XgFngNKm',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_4dOcw0JtYuR8n3f2z2bRHxmr',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:flutter_bloc/flutter_bloc.dart';\n     3\timport 'package:go_router/go_
router.dart';\n     4\timport 'package:social_world_design_system/social_world_d
esign_system.dart';\n     5\timport '../../core/di/di.dart';\n     6\timport '..
/../core/network/dio_client.dart';\n     7\timport '../../core/network/matching_
socket_client.dart';\n     8\timport '../../core/storage/secure_storage.dart';\n
    9\timport '../blocs/auth/auth_bloc.dart';\n    10\timport '../blocs/auth/aut
h_state.dart';\n    11\timport '../blocs/matching/matching_bloc.dart';\n    12\t
import '../blocs/matching/matching_event.dart';\n    13\timport '../blocs/matchi
ng/matching_state.dart';\n    14\timport '../widgets/empty_recommendation_view.d
art';\n    15\t\n    16\tclass HomePage extends StatefulWidget {\n    17\t  cons
t HomePage({super.key});\n    18\t\n    19\t  @override\n    20\t  State<HomePag
e> createState() => _HomePageState();\n    21\t}\n    22\t\n    23\tclass _HomeP
ageState extends State<HomePage> {\n    24\t  String _currentMode = 'nearby';\n 
  25\t  double _dragOffsetX = 0;\n    26\t  double _dragRotation = 0;\n    27\t 
bool _showLikeBadge = false;\n    28\t  bool _showPassBadge = false;\n    29\t  
MatchingSocketClient? _socketClient;\n    30\t\n    31\t  static const double _s
wipeThreshold = 120;\n    32\t\n    33\t  @override\n    34\t  void initState() 
{\n    35\t    super.initState();\n    36\t    context.read<MatchingBloc>().add(
const RecommendationsRequested());\n    37\t    _initSocket();\n    38\t  }\n   
39\t\n    40\t  Future<void> _initSocket() async {\n    41\t    if (_socketClien
t != null) return; // idempotent\n    42\t    final authState = context.read<Aut
hBloc>().state;\n    43\t    if (authState is AuthAuthenticated) {\n    44\t    
 final token = await SecureStorage().readAccessToken();\n    45\t      if (token
== null || token.isEmpty) return;\n    46\t      _socketClient = MatchingSocketC
lient(\n    47\t        token: token,\n    48\t        dioClient: getIt<DioClien
t>(),\n    49\t        onMatchCreated: (summary) {\n    50\t          if (mounte
d) {\n    51\t            context.read<MatchingBloc>().add(ExternalMatchReceived
(summary));\n    52\t          }\n    53\t        },\n    54\t      );\n    55\t
     _socketClient?.connect();\n    56\t    }\n    57\t  }\n    58\t\n    59\t  
@override\n    60\t  void dispose() {\n    61\t    _socketClient?.disconnect();\
n    62\t    super.dispose();\n    63\t  }\n    64\t\n    65\t  void _onModeChan
ged(String mode) {\n    66\t    setState(() {\n    67\t      _currentMode = mode
;\n    68\t      _dragOffsetX = 0;\n    69\t      _dragRotation = 0;\n    70\t  
   _showLikeBadge = false;\n    71\t      _showPassBadge = false;\n    72\t    }
);\n    73\t    context.read<MatchingBloc>().add(RecommendationsRequested(mode: 
mode));\n    74\t  }\n    75\t\n    76\t  void _onHorizontalDragUpdate(DragUpdat
eDetails details) {\n    77\t    setState(() {\n    78\t      _dragOffsetX += de
tails.delta.dx;\n    79\t      _dragRotation = _dragOffsetX * 0.001;\n    80\t  
   _showLikeBadge = _dragOffsetX > _swipeThreshold * 0.5;\n    81\t      _showPa
ssBadge = _dragOffsetX < -_swipeThreshold * 0.5;\n    82\t    });\n    83\t  }\n
   84\t\n    85\t  void _onHorizontalDragEnd(DragEndDetails details, String targ
etId) {\n    86\t    if (_dragOffsetX > _swipeThreshold) {\n    87\t      _submi
tSwipe(targetId, SwipeAction.like);\n    88\t    } else if (_dragOffsetX < -_swi
peThreshold) {\n    89\t      _submitSwipe(targetId, SwipeAction.pass);\n    90\
t    } else {\n    91\t      setState(() {\n    92\t        _dragOffsetX = 0;\n 
  93\t        _dragRotation = 0;\n    94\t        _showLikeBadge = false;\n    9
5\t        _showPassBadge = false;\n    96\t      });\n    97\t    }\n    98\t  
}\n    99\t\n   100\t  void _submitSwipe(String targetId, SwipeAction action) {\
n   101\t    setState(() {\n   102\t      _dragOffsetX = action == SwipeAction.l
ike ? 500 : -500;\n   103\t      _dragRotation = action == SwipeAction.like ? 0.
3 : -0.3;\n   104\t    });\n   105\t\n   106\t    Future.delayed(const Duration(
milliseconds: 200), () {\n   107\t      if (!mounted) return;\n   108\t      con
text.read<MatchingBloc>().add(\n   109\t            SwipeActionSubmitted(targetI
d: targetId, action: action),\n   110\t          );\n   111\t      setState(() {
\n   112\t        _dragOffsetX = 0;\n   113\t        _dragRotation = 0;\n   114\
t        _showLikeBadge = false;\n   115\t        _showPassBadge = false;\n   11
6\t      });\n   117\t    });\n   118\t  }\n   119\t\n   120\t  void _dismissMat
chModal() {\n   121\t    context.read<MatchingBloc>().add(const MatchModalDismis
sed());\n   122\t  }\n   123\t\n   124\t  @override\n   125\t  Widget build(Buil
dContext context) {\n   126\t    return BlocListener<AuthBloc, AuthState>(\n   1
27\t      listenWhen: (prev, curr) =>\n   128\t          curr is AuthAuthenticat
ed && prev is! AuthAuthenticated,\n   129\t      listener: (context, state) => _
initSocket(),\n   130\t      child: Scaffold(\n   131\t        backgroundColor: 
DesignTokens.background,\n   132\t        body: SafeArea(\n   133\t          chi
ld: BlocConsumer<MatchingBloc, MatchingState>(\n   134\t            listener: (c
ontext, state) {\n   135\t              if (state is MatchingFailure) {\n   136\
t                ScaffoldMessenger.of(context).showSnackBar(\n   137\t          
       SnackBar(content: Text(state.message)),\n   138\t                );\n   1
39\t              }\n   140\t              if (state is MatchingProfileIncomplet
e) {\n   141\t                ScaffoldMessenger.of(context).showSnackBar(\n   14
2\t                  SnackBar(content: Text(state.message)),\n   143\t          
     );\n   144\t                context.go('/profile-setup');\n   145\t        
     }\n   146\t            },\n   147\t          builder: (context, state) {\n 
 148\t            return Stack(\n   149\t              children: [\n   150\t    
           Column(\n   151\t                  children: [\n   152\t             
      _buildModeTabs(),\n   153\t                    Expanded(\n   154\t        
             child: _buildContent(state),\n   155\t                    ),\n   15
6\t                    if (state is MatchingLoadSuccess)\n   157\t              
       _buildActionButtons(state),\n   158\t                    const SizedBox(h
eight: DesignTokens.spacingMd),\n   159\t                  ],\n   160\t         
      ),\n   161\t                if (state is MatchingMatchCreated)\n   162\t  
               MatchSuccessDialog(\n   163\t                    matchedUserNickn
ame: state.match.matchedUser.nickname,\n   164\t                    matchedUserA
vatarUrl: state.match.matchedUser.avatarUrl,\n   165\t                    onChat
: () {\n   166\t                      _dismissMatchModal();\n   167\t           
          // TODO: Navigate to chat\n   168\t                    },\n   169\t   
                onContinue: _dismissMatchModal,\n   170\t                  ),\n 
 171\t              ],\n   172\t            );\n   173\t          },\n   174\t  
     ),\n   175\t      ),\n   176\t    ),\n   177\t    );\n   178\t  }\n   179\t
\n   180\t  Widget _buildModeTabs() {\n   181\t    final modes = [\n   182\t    
 {'key': 'nearby', 'label': '附近'},\n   183\t      {'key': 'city', 'label': '同
城'},\n   184\t      {'key': 'nationwide', 'label': '全国'},\n   185\t    ];\n  
186\t\n   187\t    return Padding(\n   188\t      padding: const EdgeInsets.all(
DesignTokens.spacingMd),\n   189\t      child: Container(\n   190\t        decor
ation: BoxDecoration(\n   191\t          color: Theme.of(context).colorScheme.su
rface,\n   192\t          borderRadius: BorderRadius.circular(DesignTokens.radiu
sXl),\n   193\t        ),\n   194\t        padding: const EdgeInsets.all(DesignT
okens.spacingXs),\n   195\t        child: Row(\n   196\t          children: mode
s.map((mode) {\n   197\t            final isActive = _currentMode == mode['key']
;\n   198\t            return Expanded(\n   199\t              child: Semantics(
\n   200\t                identifier: 'mode_tab_${mode['key']}',\n   201\t      
         label: mode['label'] as String,\n   202\t                button: true,\
n   203\t                child: GestureDetector(\n   204\t                  beha
vior: HitTestBehavior.opaque,\n   205\t                  onTapDown: (_) => _onMo
deChanged(mode['key'] as String),\n   206\t                  child: Container(\n
  207\t                    padding: const EdgeInsets.symmetric(\n   208\t       
              vertical: DesignTokens.spacingSm,\n   209\t                    ),\
n   210\t                    decoration: BoxDecoration(\n   211\t               
      color: isActive\n   212\t                          ? DesignTokens.primary\
n   213\t                          : DesignTokens.background.withOpacity(0),\n  
214\t                      borderRadius:\n   215\t                          Bord
erRadius.circular(DesignTokens.radiusLg),\n   216\t                    ),\n   21
7\t                    child: Center(\n   218\t                      child: Text
(\n   219\t                        mode['label'] as String,\n   220\t           
            style: TextStyle(\n   221\t                          color: isActive
\n   222\t                              ? DesignTokens.textOnDarkPrimary\n   223
\t                              : DesignTokens.textPrimary,\n   224\t           
              fontSize: DesignTokens.fontSizeBody,\n   225\t                    
     fontWeight:\n   226\t                              isActive ? FontWeight.bo
ld : FontWeight.normal,\n   227\t                        ),\n   228\t           
          ),\n   229\t                    ),\n   230\t                  ),\n   2
31\t                ),\n   232\t              ),\n   233\t            );\n   234
\t          }).toList(),\n   235\t        ),\n   236\t      ),\n   237\t    );\n
  238\t  }\n   239\t\n   240\t  Widget _buildContent(MatchingState state) {\n   
241\t    if (state is MatchingLoadInProgress) {\n   242\t      return const Cent
er(\n   243\t        child: CircularProgressIndicator(color: DesignTokens.primar
y),\n   244\t      );\n   245\t    }\n   246\t\n   247\t    if (state is Matchin
gLoadEmpty) {\n   248\t      return EmptyRecommendationView(\n   249\t        on
SwitchNationwide: state.mode != 'nationwide'\n   250\t            ? () => _onMod
eChanged('nationwide')\n   251\t            : null,\n   252\t      );\n   253\t 
  }\n   254\t\n   255\t    if (state is MatchingLoadSuccess || state is Matching
SwipeInProgress) {\n   256\t      final cards = state is MatchingLoadSuccess\n  
257\t          ? state.cards\n   258\t          : (state as MatchingSwipeInProgr
ess).cards;\n   259\t\n   260\t      if (cards.isEmpty) {\n   261\t        retur
n EmptyRecommendationView(\n   262\t          onSwitchNationwide: _currentMode !
= 'nationwide'\n   263\t              ? () => _onModeChanged('nationwide')\n   2
64\t              : null,\n   265\t        );\n   266\t      }\n   267\t\n   268
\t      return Semantics(\n   269\t        identifier: 'card_stack',\n   270\t  
     label: '推荐卡片区域',\n   271\t        container: true,\n   272\t        c
hild: GestureDetector(\n   273\t          onHorizontalDragUpdate: _onHorizontalD
ragUpdate,\n   274\t          onHorizontalDragEnd: (details) =>\n   275\t       
      _onHorizontalDragEnd(details, cards[0].userId),\n   276\t          child: 
SwipeCardStack(\n   277\t            cards: cards\n   278\t                .map(
(c) => SwipeCardData(\n   279\t                      userId: c.userId,\n   280\t
                     nickname: c.nickname,\n   281\t                      age: c
.age,\n   282\t                      distance: c.distance,\n   283\t            
         bio: c.bio,\n   284\t                      avatarUrl: c.avatarUrl,\n   
285\t                    ))\n   286\t                .toList(),\n   287\t       
    topCardOffsetX: _dragOffsetX,\n   288\t            topCardRotation: _dragRot
ation,\n   289\t            showLikeBadge: _showLikeBadge,\n   290\t            
showPassBadge: _showPassBadge,\n   291\t          ),\n   292\t        ),\n   293
\t      );\n   294\t    }\n   295\t\n   296\t    if (state is MatchingFailure) {
\n   297\t      return Center(\n   298\t        child: Column(\n   299\t        
 mainAxisAlignment: MainAxisAlignment.center,\n   300\t          children: [\n  
301\t            const Icon(Icons.error_outline,\n   302\t                size: 
48, color: DesignTokens.error),\n   303\t            const SizedBox(height: Desi
gnTokens.spacingMd),\n   304\t            Text(\n   305\t              state.mes
sage,\n   306\t              style: const TextStyle(color: DesignTokens.error),\
n   307\t              textAlign: TextAlign.center,\n   308\t            ),\n   
309\t            const SizedBox(height: DesignTokens.spacingMd),\n   310\t      
     Semantics(\n   311\t              identifier: 'retry_button',\n   312\t    
         label: '重试',\n   313\t              button: true,\n   314\t          
   child: GestureDetector(\n   315\t                behavior: HitTestBehavior.op
aque,\n   316\t                onTapDown: (_) => context.read<MatchingBloc>().ad
d(\n   317\t                      RecommendationsRequested(mode: _currentMode),\
n   318\t                    ),\n   319\t                child: Container(\n   3
20\t                  padding: const EdgeInsets.symmetric(\n   321\t            
       horizontal: DesignTokens.spacingLg,\n   322\t                    vertical
: DesignTokens.spacingMd,\n   323\t                  ),\n   324\t               
  decoration: BoxDecoration(\n   325\t                    border: Border.all(col
or: DesignTokens.primary),\n   326\t                    borderRadius: BorderRadi
us.circular(DesignTokens.radiusMd),\n   327\t                  ),\n   328\t     
            child: const Text(\n   329\t                    '重试',\n   330\t   
                style: TextStyle(color: DesignTokens.primary),\n   331\t        
         ),\n   332\t                ),\n   333\t              ),\n   334\t     
      ),\n   335\t          ],\n   336\t        ),\n   337\t      );\n   338\t  
 }\n   339\t\n   340\t    return const SizedBox.shrink();\n   341\t  }\n   342\t
\n   343\t  Widget _buildActionButtons(MatchingLoadSuccess state) {\n   344\t   
if (state.cards.isEmpty) return const SizedBox.shrink();\n   345\t\n   346\t    
final targetId = state.cards[0].userId;\n   347\t\n   348\t    return Padding(\n
  349\t      padding: const EdgeInsets.symmetric(\n   350\t        horizontal: D
esignTokens.spacingLg,\n   351\t      ),\n   352\t      child: Row(\n   353\t   
    mainAxisAlignment: MainAxisAlignment.spaceEvenly,\n   354\t        children:
[\n   355\t          SwActionButton(\n   356\t            type: SwActionButtonTy
pe.pass,\n   357\t            onTap: () => _submitSwipe(targetId, SwipeAction.pa
ss),\n   358\t          ),\n   359\t          SwActionButton(\n   360\t         
  type: SwActionButtonType.like,\n   361\t            onTap: () => _submitSwipe(
targetId, SwipeAction.like),\n   362\t          ),\n   363\t        ],\n   364\t
     ),\n   365\t    );\n   366\t  }\n   367\t}\n",
        message='367 lines read from file starting from line 1. Total lines in f
ile: 367. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/apps/mobile/li
b/presentation/pages/profile_setup_page.dart"}'
)
ToolCall(
    type='function',
    id='tool_6cznT4zLnggjFZTxLPS41Cti',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_ncDtdXCRuMyuSkY3XgFngNKm',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:flutter_bloc/flutter_bloc.dart';\n     3\timport 'package:go_router/go_
router.dart';\n     4\timport 'package:social_world_design_system/social_world_d
esign_system.dart';\n     5\timport '../../core/di/di.dart';\n     6\timport 'pa
ckage:social_world_profile_models/api.dart';\n     7\timport '../blocs/profile/p
rofile_bloc.dart';\n     8\timport '../blocs/profile/profile_event.dart';\n     
9\timport '../blocs/profile/profile_state.dart';\n    10\t\n    11\tclass Profil
eSetupPage extends StatelessWidget {\n    12\t  const ProfileSetupPage({super.ke
y});\n    13\t\n    14\t  @override\n    15\t  Widget build(BuildContext context
) {\n    16\t    return BlocProvider(\n    17\t      create: (_) => getIt<Profil
eBloc>(),\n    18\t      child: const _ProfileSetupPageBody(),\n    19\t    );\n
   20\t  }\n    21\t}\n    22\t\n    23\tclass _ProfileSetupPageBody extends Sta
tefulWidget {\n    24\t  const _ProfileSetupPageBody();\n    25\t\n    26\t  @ov
erride\n    27\t  State<_ProfileSetupPageBody> createState() => _ProfileSetupPag
eBodyState();\n    28\t}\n    29\t\n    30\tclass _ProfileSetupPageBodyState ext
ends State<_ProfileSetupPageBody> {\n    31\t  int _currentStep = 0;\n    32\t  
final _nicknameController = TextEditingController();\n    33\t  String? _gender;
\n    34\t  DateTime? _birthDate;\n    35\t  String? _avatarUrl;\n    36\t  doub
le? _latitude;\n    37\t  double? _longitude;\n    38\t  bool _isLoading = false
;\n    39\t\n    40\t  @override\n    41\t  void dispose() {\n    42\t    _nickn
ameController.dispose();\n    43\t    super.dispose();\n    44\t  }\n    45\t\n 
  46\t  void _onNext() {\n    47\t    if (_currentStep < 2) {\n    48\t      set
State(() => _currentStep++);\n    49\t    } else {\n    50\t      _onSave();\n  
 51\t    }\n    52\t  }\n    53\t\n    54\t  void _onSave() {\n    55\t    if (_
nicknameController.text.trim().isEmpty ||\n    56\t        _gender == null ||\n 
  57\t        _birthDate == null ||\n    58\t        _latitude == null ||\n    5
9\t        _longitude == null) {\n    60\t      ScaffoldMessenger.of(context).sh
owSnackBar(\n    61\t        const SnackBar(content: Text('请填写完整资料')),\n 
  62\t      );\n    63\t      return;\n    64\t    }\n    65\t\n    66\t    fina
l request = UpdateProfileRequest(\n    67\t      nickname: _nicknameController.t
ext.trim(),\n    68\t      gender: _gender == 'male'\n    69\t          ? Update
ProfileRequestGenderEnum.male\n    70\t          : _gender == 'female'\n    71\t
             ? UpdateProfileRequestGenderEnum.female\n    72\t              : Up
dateProfileRequestGenderEnum.other,\n    73\t      birthDate: _birthDate!,\n    
74\t      latitude: _latitude!,\n    75\t      longitude: _longitude!,\n    76\t
     avatarUrl: _avatarUrl,\n    77\t    );\n    78\t\n    79\t    context.read<
ProfileBloc>().add(SaveProfile(request));\n    80\t  }\n    81\t\n    82\t  Futu
re<void> _pickBirthDate() async {\n    83\t    final now = DateTime.now();\n    
84\t    final picked = await showDatePicker(\n    85\t      context: context,\n 
  86\t      initialDate: DateTime(now.year - 20),\n    87\t      firstDate: Date
Time(1900),\n    88\t      lastDate: now,\n    89\t    );\n    90\t    if (picke
d != null) {\n    91\t      setState(() => _birthDate = picked);\n    92\t    }\
n    93\t  }\n    94\t\n    95\t  Future<void> _pickLocation() async {\n    96\t
   // MVP: simulate location permission and use mock coordinates\n    97\t    se
tState(() {\n    98\t      _latitude = 31.2304;\n    99\t      _longitude = 121.
4737;\n   100\t    });\n   101\t  }\n   102\t\n   103\t  void _showAvatarSheet()
{\n   104\t    SwBottomSheet.show(\n   105\t      context,\n   106\t      title:
'上传头像',\n   107\t      actions: [\n   108\t        SwBottomSheetAction(\n   
109\t          label: '从相册选择',\n   110\t          icon: Icons.photo_library
,\n   111\t          onTap: () {\n   112\t            setState(() {\n   113\t   
          _avatarUrl = 'https://i.pravatar.cc/300?img=${DateTime.now().milliseco
nd % 70}';\n   114\t            });\n   115\t          },\n   116\t        ),\n 
 117\t        SwBottomSheetAction(\n   118\t          label: '拍照',\n   119\t  
       icon: Icons.camera_alt,\n   120\t          onTap: () {\n   121\t         
  setState(() {\n   122\t              _avatarUrl = 'https://i.pravatar.cc/300?i
mg=${DateTime.now().millisecond % 70}';\n   123\t            });\n   124\t      
   },\n   125\t        ),\n   126\t      ],\n   127\t    );\n   128\t  }\n   129
\t\n   130\t  @override\n   131\t  Widget build(BuildContext context) {\n   132\
t    return BlocListener<ProfileBloc, ProfileState>(\n   133\t      listener: (c
ontext, state) {\n   134\t        if (state is ProfileLoading) {\n   135\t      
   setState(() => _isLoading = true);\n   136\t        } else if (state is Profi
leComplete || state is ProfileError) {\n   137\t          setState(() => _isLoad
ing = false);\n   138\t        }\n   139\t        if (state is ProfileComplete) 
{\n   140\t          context.go('/home');\n   141\t        } else if (state is P
rofileError) {\n   142\t          ScaffoldMessenger.of(context).showSnackBar(\n 
 143\t            SnackBar(content: Text(state.message)),\n   144\t          );\
n   145\t        }\n   146\t      },\n   147\t      child: Scaffold(\n   148\t  
     backgroundColor: DesignTokens.background,\n   149\t        appBar: AppBar(\
n   150\t          backgroundColor: DesignTokens.background,\n   151\t          
elevation: 0,\n   152\t          title: const Text('完善资料', style: TextStyle(
color: DesignTokens.textPrimary)),\n   153\t        ),\n   154\t        body: Sa
feArea(\n   155\t          child: Padding(\n   156\t            padding: const E
dgeInsets.all(DesignTokens.spacingMd),\n   157\t            child: Column(\n   1
58\t              crossAxisAlignment: CrossAxisAlignment.stretch,\n   159\t     
        children: [\n   160\t                LinearProgressIndicator(\n   161\t 
                value: (_currentStep + 1) / 3,\n   162\t                  backgr
oundColor: DesignTokens.primary.withOpacity(0.2),\n   163\t                  val
ueColor: const AlwaysStoppedAnimation<Color>(DesignTokens.primary),\n   164\t   
            ),\n   165\t                const SizedBox(height: DesignTokens.spac
ingLg),\n   166\t                Expanded(child: _buildStep()),\n   167\t       
        const SizedBox(height: DesignTokens.spacingMd),\n   168\t               
SwButton(\n   169\t                  text: _currentStep < 2 ? '下一步' : '完成',
\n   170\t                  variant: SwButtonVariant.filled,\n   171\t          
       isLoading: _isLoading,\n   172\t                  onPressed: _isLoading ?
null : _onNext,\n   173\t                ),\n   174\t              ],\n   175\t 
          ),\n   176\t          ),\n   177\t        ),\n   178\t      ),\n   179
\t    );\n   180\t  }\n   181\t\n   182\t  Widget _buildStep() {\n   183\t    sw
itch (_currentStep) {\n   184\t      case 0:\n   185\t        return _buildAvata
rStep();\n   186\t      case 1:\n   187\t        return _buildBasicInfoStep();\n
  188\t      case 2:\n   189\t        return _buildLocationStep();\n   190\t    
 default:\n   191\t        return const SizedBox.shrink();\n   192\t    }\n   19
3\t  }\n   194\t\n   195\t  Widget _buildAvatarStep() {\n   196\t    return Colu
mn(\n   197\t      mainAxisAlignment: MainAxisAlignment.center,\n   198\t      c
hildren: [\n   199\t        const Text(\n   200\t          '上传头像',\n   201\t
         style: TextStyle(\n   202\t            fontSize: DesignTokens.fontSizeT
itle,\n   203\t            fontWeight: FontWeight.bold,\n   204\t            col
or: DesignTokens.textPrimary,\n   205\t          ),\n   206\t        ),\n   207\
t        const SizedBox(height: DesignTokens.spacingLg),\n   208\t        Gestur
eDetector(\n   209\t          onTap: _showAvatarSheet,\n   210\t          child:
SwAvatar(\n   211\t            imageUrl: _avatarUrl,\n   212\t            initia
ls: _nicknameController.text.isNotEmpty\n   213\t                ? _nicknameCont
roller.text.substring(0, 1)\n   214\t                : '?',\n   215\t           
size: 120,\n   216\t          ),\n   217\t        ),\n   218\t        const Size
dBox(height: DesignTokens.spacingMd),\n   219\t        Text(\n   220\t          
_avatarUrl == null ? '点击上传头像' : '头像已选择',\n   221\t          style: Te
xtStyle(\n   222\t            color: DesignTokens.textPrimary.withOpacity(0.7),\
n   223\t          ),\n   224\t        ),\n   225\t      ],\n   226\t    );\n   
227\t  }\n   228\t\n   229\t  Widget _buildBasicInfoStep() {\n   230\t    return
SingleChildScrollView(\n   231\t      child: Column(\n   232\t        crossAxisA
lignment: CrossAxisAlignment.stretch,\n   233\t        children: [\n   234\t    
     const Text(\n   235\t            '基础资料',\n   236\t            style: Te
xtStyle(\n   237\t              fontSize: DesignTokens.fontSizeTitle,\n   238\t 
            fontWeight: FontWeight.bold,\n   239\t              color: DesignTok
ens.textPrimary,\n   240\t            ),\n   241\t          ),\n   242\t        
 const SizedBox(height: DesignTokens.spacingLg),\n   243\t          SwTextField(
\n   244\t            controller: _nicknameController,\n   245\t            labe
lText: '昵称',\n   246\t            hintText: '请输入昵称',\n   247\t          )
,\n   248\t          const SizedBox(height: DesignTokens.spacingMd),\n   249\t  
       Column(\n   250\t            crossAxisAlignment: CrossAxisAlignment.start
,\n   251\t            children: [\n   252\t              const Text('性别', sty
le: TextStyle(color: DesignTokens.textPrimary)),\n   253\t              const Si
zedBox(height: DesignTokens.spacingXs),\n   254\t              SwGenderSelector(
\n   255\t                selectedGender: _gender,\n   256\t                onCh
anged: (v) => setState(() => _gender = v),\n   257\t              ),\n   258\t  
         ],\n   259\t          ),\n   260\t          const SizedBox(height: Desi
gnTokens.spacingMd),\n   261\t          SwTextField(\n   262\t            labelT
ext: '出生日期',\n   263\t            hintText: _birthDate == null\n   264\t    
           ? '请选择出生日期'\n   265\t                : '${_birthDate!.year}-${
_birthDate!.month.toString().padLeft(2, '0')}-${_birthDate!.day.toString().padLe
ft(2, '0')}',\n   266\t            readOnly: true,\n   267\t            onTap: _
pickBirthDate,\n   268\t          ),\n   269\t        ],\n   270\t      ),\n   2
71\t    );\n   272\t  }\n   273\t\n   274\t  Widget _buildLocationStep() {\n   2
75\t    return Column(\n   276\t      mainAxisAlignment: MainAxisAlignment.cente
r,\n   277\t      children: [\n   278\t        const Text(\n   279\t          '
位置授权',\n   280\t          style: TextStyle(\n   281\t            fontSize: D
esignTokens.fontSizeTitle,\n   282\t            fontWeight: FontWeight.bold,\n  
283\t            color: DesignTokens.textPrimary,\n   284\t          ),\n   285\
t        ),\n   286\t        const SizedBox(height: DesignTokens.spacingLg),\n  
287\t        Icon(\n   288\t          Icons.location_on,\n   289\t          size
: 64,\n   290\t          color: _latitude != null ? DesignTokens.primary : Desig
nTokens.textPrimary.withOpacity(0.3),\n   291\t        ),\n   292\t        const
SizedBox(height: DesignTokens.spacingMd),\n   293\t        Text(\n   294\t      
   _latitude == null ? '允许获取位置以开始匹配' : '位置已获取',\n   295\t       
  style: TextStyle(\n   296\t            color: DesignTokens.textPrimary.withOpa
city(0.7),\n   297\t          ),\n   298\t        ),\n   299\t        if (_latit
ude != null) ...[\n   300\t          const SizedBox(height: DesignTokens.spacing
Sm),\n   301\t          Text(\n   302\t            '纬度: ${_latitude!.toStringA
sFixed(4)}, 经度: ${_longitude!.toStringAsFixed(4)}',\n   303\t            style
: const TextStyle(color: DesignTokens.textPrimary),\n   304\t          ),\n   30
5\t        ],\n   306\t        const SizedBox(height: DesignTokens.spacingLg),\n
  307\t        if (_latitude == null)\n   308\t          SwButton(\n   309\t    
       text: '允许位置授权',\n   310\t            variant: SwButtonVariant.outli
ned,\n   311\t            onPressed: _pickLocation,\n   312\t          ),\n   31
3\t      ],\n   314\t    );\n   315\t  }\n   316\t}\n",
        message='316 lines read from file starting from line 1. Total lines in f
ile: 316. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='00000gn/T/dev-issue-55-815001100/apps/mobile/lib/presentatio
n/pages/match_list_page.dart"}'
)
ToolCall(
    type='function',
    id='tool_lM65u18gO1rrM4KqDtF3Ah6F',
    function=FunctionBody(name='ReadFile', arguments='{"path":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_6cznT4zLnggjFZTxLPS41Cti',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:flutter_bloc/flutter_bloc.dart';\n     3\timport 'package:social_world_
design_system/social_world_design_system.dart';\n     4\timport '../blocs/matchi
ng/match_list_cubit.dart';\n     5\timport '../blocs/matching/match_list_state.d
art';\n     6\t\n     7\tclass MatchListPage extends StatefulWidget {\n     8\t 
const MatchListPage({super.key});\n     9\t\n    10\t  @override\n    11\t  Stat
e<MatchListPage> createState() => _MatchListPageState();\n    12\t}\n    13\t\n 
  14\tclass _MatchListPageState extends State<MatchListPage> {\n    15\t  @overr
ide\n    16\t  void initState() {\n    17\t    super.initState();\n    18\t    c
ontext.read<MatchListCubit>().fetchMatches();\n    19\t  }\n    20\t\n    21\t  
@override\n    22\t  Widget build(BuildContext context) {\n    23\t    return Sc
affold(\n    24\t      backgroundColor: DesignTokens.background,\n    25\t      
appBar: AppBar(\n    26\t        backgroundColor: DesignTokens.background,\n    
27\t        elevation: 0,\n    28\t        title: const Text(\n    29\t         
'消息',\n    30\t          style: TextStyle(\n    31\t            color: DesignT
okens.textPrimary,\n    32\t            fontSize: DesignTokens.fontSizeTitle,\n 
  33\t            fontWeight: FontWeight.bold,\n    34\t          ),\n    35\t  
     ),\n    36\t      ),\n    37\t      body: BlocBuilder<MatchListCubit, Match
ListState>(\n    38\t        builder: (context, state) {\n    39\t          if (
state is MatchListLoadInProgress) {\n    40\t            return const Center(\n 
  41\t              child: CircularProgressIndicator(color: DesignTokens.primary
),\n    42\t            );\n    43\t          }\n    44\t\n    45\t          if 
(state is MatchListLoadFailure) {\n    46\t            return Center(\n    47\t 
            child: Column(\n    48\t                mainAxisAlignment: MainAxisA
lignment.center,\n    49\t                children: [\n    50\t                 
const Icon(\n    51\t                    Icons.error_outline,\n    52\t         
          size: 48,\n    53\t                    color: DesignTokens.error,\n   
54\t                  ),\n    55\t                  const SizedBox(height: Desig
nTokens.spacingMd),\n    56\t                  Text(\n    57\t                  
 state.message,\n    58\t                    style: const TextStyle(color: Desig
nTokens.error),\n    59\t                    textAlign: TextAlign.center,\n    6
0\t                  ),\n    61\t                  const SizedBox(height: Design
Tokens.spacingMd),\n    62\t                  Semantics(\n    63\t              
     identifier: 'match_list_retry_button',\n    64\t                    label: 
'重试',\n    65\t                    button: true,\n    66\t                    
child: GestureDetector(\n    67\t                      behavior: HitTestBehavior
.opaque,\n    68\t                      onTapDown: (_) =>\n    69\t             
            context.read<MatchListCubit>().fetchMatches(),\n    70\t            
         child: Container(\n    71\t                        padding: const EdgeI
nsets.symmetric(\n    72\t                          horizontal: DesignTokens.spa
cingLg,\n    73\t                          vertical: DesignTokens.spacingMd,\n  
 74\t                        ),\n    75\t                        decoration: Box
Decoration(\n    76\t                          border: Border.all(color: DesignT
okens.primary),\n    77\t                          borderRadius:\n    78\t      
                       BorderRadius.circular(DesignTokens.radiusMd),\n    79\t  
                     ),\n    80\t                        child: const Text(\n   
81\t                          '重试',\n    82\t                          style: 
TextStyle(color: DesignTokens.primary),\n    83\t                        ),\n   
84\t                      ),\n    85\t                    ),\n    86\t          
       ),\n    87\t                ],\n    88\t              ),\n    89\t       
    );\n    90\t          }\n    91\t\n    92\t          if (state is MatchListL
oadSuccess) {\n    93\t            if (state.matches.isEmpty) {\n    94\t       
      return Center(\n    95\t                child: Column(\n    96\t          
       mainAxisAlignment: MainAxisAlignment.center,\n    97\t                  c
hildren: [\n    98\t                    Icon(\n    99\t                      Ico
ns.chat_bubble_outline,\n   100\t                      size: 80,\n   101\t      
               color: DesignTokens.textPrimary.withOpacity(0.3),\n   102\t      
             ),\n   103\t                    const SizedBox(height: DesignTokens
.spacingMd),\n   104\t                    const Text(\n   105\t                 
    '还没有匹配成功的好友',\n   106\t                      style: TextStyle(\n  
107\t                        fontSize: DesignTokens.fontSizeTitle,\n   108\t    
                   fontWeight: FontWeight.bold,\n   109\t                       
color: DesignTokens.textPrimary,\n   110\t                      ),\n   111\t    
               ),\n   112\t                    const SizedBox(height: DesignToke
ns.spacingSm),\n   113\t                    Text(\n   114\t                     
'继续滑动，发现感兴趣的人吧',\n   115\t                      style: TextStyle(\n
  116\t                        fontSize: DesignTokens.fontSizeBody,\n   117\t   
                    color: DesignTokens.textPrimary.withOpacity(0.6),\n   118\t 
                    ),\n   119\t                    ),\n   120\t                
 ],\n   121\t                ),\n   122\t              );\n   123\t            }
\n   124\t\n   125\t            return RefreshIndicator(\n   126\t              
onRefresh: () => context.read<MatchListCubit>().fetchMatches(),\n   127\t       
      color: DesignTokens.primary,\n   128\t              child: ListView.separa
ted(\n   129\t                physics: const AlwaysScrollableScrollPhysics(),\n 
 130\t                itemCount: state.matches.length + (state.hasMore ? 1 : 0),
\n   131\t                separatorBuilder: (_, __) => const Divider(\n   132\t 
                height: 1,\n   133\t                  color: DesignTokens.divide
r,\n   134\t                  indent: DesignTokens.spacingMd + 56 + DesignTokens
.spacingMd,\n   135\t                ),\n   136\t                itemBuilder: (c
ontext, index) {\n   137\t                  if (index >= state.matches.length) {
\n   138\t                    context.read<MatchListCubit>().fetchMatches(\n   1
39\t                          cursor: state.nextCursor,\n   140\t               
        );\n   141\t                    return const Center(\n   142\t          
           child: Padding(\n   143\t                        padding: EdgeInsets.
all(DesignTokens.spacingMd),\n   144\t                        child: CircularPro
gressIndicator(\n   145\t                          color: DesignTokens.primary,\
n   146\t                        ),\n   147\t                      ),\n   148\t 
                  );\n   149\t                  }\n   150\t\n   151\t           
      final match = state.matches[index];\n   152\t                  return Matc
hListItem(\n   153\t                    nickname: match.nickname,\n   154\t     
              avatarUrl: match.avatarUrl,\n   155\t                    matchedAt
: match.matchedAt,\n   156\t                    onTap: () {\n   157\t           
          // TODO: Navigate to chat page\n   158\t                    },\n   159
\t                  );\n   160\t                },\n   161\t              ),\n  
162\t            );\n   163\t          }\n   164\t\n   165\t          return con
st SizedBox.shrink();\n   166\t        },\n   167\t      ),\n   168\t    );\n   
169\t  }\n   170\t}\n",
        message='170 lines read from file starting from line 1. Total lines in f
ile: 170. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/apps/mobile/lib/presentation/widgets/dev_tools_toggle_button.dart"
}'
)
ToolCall(
    type='function',
    id='tool_bik6xRAANhfvdP5SqnoOM3gE',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_lM65u18gO1rrM4KqDtF3Ah6F',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:social_world_design_system/social_world_design_system.dart';\n     3\t\
n     4\tclass DevToolsToggleButton extends StatelessWidget {\n     5\t  final V
oidCallback onTap;\n     6\t\n     7\t  const DevToolsToggleButton({\n     8\t  
 super.key,\n     9\t    required this.onTap,\n    10\t  });\n    11\t\n    12\t
 @override\n    13\t  Widget build(BuildContext context) {\n    14\t    final sc
reenHeight = MediaQuery.of(context).size.height;\n    15\t    return Positioned(
\n    16\t      right: DesignTokens.spacingXs,\n    17\t      top: screenHeight 
/ 2 - 28,\n    18\t      width: 56,\n    19\t      height: 56,\n    20\t      ch
ild: Semantics(\n    21\t        identifier: 'dev_tools_toggle_button',\n    22\
t        label: '开发工具',\n    23\t        button: true,\n    24\t        cont
ainer: true,\n    25\t        child: GestureDetector(\n    26\t          behavio
r: HitTestBehavior.opaque,\n    27\t          onTapDown: (_) => onTap(),\n    28
\t          child: Container(\n    29\t            decoration: const BoxDecorati
on(\n    30\t              color: Colors.black26,\n    31\t              shape: 
BoxShape.circle,\n    32\t            ),\n    33\t            child: const Icon(
\n    34\t              Icons.construction,\n    35\t              color: Colors
.white,\n    36\t              size: 24,\n    37\t            ),\n    38\t      
   ),\n    39\t        ),\n    40\t      ),\n    41\t    );\n    42\t  }\n    43
\t}\n",
        message='43 lines read from file starting from line 1. Total lines in fi
le: 43. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='00gn/T/dev-issue-55-815001100/apps/mobile/lib/presentation/w
idgets/dev_tools_drawer.dart"}'
)
StatusUpdate(
    context_usage=0.09985733032226562,
    context_tokens=26177,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1089,
        output=380,
        input_cache_read=25088,
        input_cache_creation=0
    ),
    message_id='chatcmpl-FAiBrJmp5TzBUuTL5vohHy2D',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_bik6xRAANhfvdP5SqnoOM3gE',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:social_world_design_system/social_world_design_system.dart';\n     3\ti
mport 'server_url_tool.dart';\n     4\t\n     5\tclass DevToolsDrawer extends St
atefulWidget {\n     6\t  final bool isOpen;\n     7\t  final VoidCallback onClo
se;\n     8\t\n     9\t  const DevToolsDrawer({\n    10\t    super.key,\n    11\
t    required this.isOpen,\n    12\t    required this.onClose,\n    13\t  });\n 
  14\t\n    15\t  @override\n    16\t  State<DevToolsDrawer> createState() => _D
evToolsDrawerState();\n    17\t}\n    18\t\n    19\tclass _DevToolsDrawerState e
xtends State<DevToolsDrawer>\n    20\t    with SingleTickerProviderStateMixin {\
n    21\t  late final AnimationController _controller;\n    22\t  late final Ani
mation<Offset> _slideAnimation;\n    23\t\n    24\t  static const _maxWidth = 32
0.0;\n    25\t\n    26\t  @override\n    27\t  void initState() {\n    28\t    s
uper.initState();\n    29\t    _controller = AnimationController(\n    30\t     
vsync: this,\n    31\t      duration: const Duration(milliseconds: 250),\n    32
\t      reverseDuration: const Duration(milliseconds: 200),\n    33\t    );\n   
34\t    _slideAnimation = Tween<Offset>(\n    35\t      begin: const Offset(1, 0
),\n    36\t      end: Offset.zero,\n    37\t    ).animate(\n    38\t      Curve
dAnimation(\n    39\t        parent: _controller,\n    40\t        curve: Curves
.easeOutCubic,\n    41\t        reverseCurve: Curves.easeInCubic,\n    42\t     
),\n    43\t    );\n    44\t    if (widget.isOpen) {\n    45\t      _controller.
value = 1.0;\n    46\t    }\n    47\t  }\n    48\t\n    49\t  @override\n    50\
t  void didUpdateWidget(covariant DevToolsDrawer oldWidget) {\n    51\t    super
.didUpdateWidget(oldWidget);\n    52\t    if (widget.isOpen && !oldWidget.isOpen
) {\n    53\t      _controller.forward();\n    54\t    } else if (!widget.isOpen
&& oldWidget.isOpen) {\n    55\t      _controller.reverse();\n    56\t    }\n   
57\t  }\n    58\t\n    59\t  @override\n    60\t  void dispose() {\n    61\t    
_controller.dispose();\n    62\t    super.dispose();\n    63\t  }\n    64\t\n   
65\t  @override\n    66\t  Widget build(BuildContext context) {\n    67\t    fin
al screenWidth = MediaQuery.of(context).size.width;\n    68\t    final drawerWid
th = screenWidth * 0.8 > _maxWidth\n    69\t        ? _maxWidth\n    70\t       
: screenWidth * 0.8;\n    71\t\n    72\t    return AnimatedBuilder(\n    73\t   
  animation: _controller,\n    74\t      builder: (context, child) {\n    75\t  
     if (_controller.value == 0.0) {\n    76\t          return const SizedBox.sh
rink();\n    77\t        }\n    78\t        return Stack(\n    79\t          chi
ldren: [\n    80\t            GestureDetector(\n    81\t              onTap: wid
get.onClose,\n    82\t              child: Container(\n    83\t                c
olor: Colors.black.withOpacity(0.15 * _controller.value),\n    84\t             
),\n    85\t            ),\n    86\t            SlideTransition(\n    87\t      
       position: _slideAnimation,\n    88\t              child: Align(\n    89\t
               alignment: Alignment.centerRight,\n    90\t                child:
Container(\n    91\t                  width: drawerWidth,\n    92\t             
    color: DesignTokens.background,\n    93\t                  child: SafeArea(\
n    94\t                    child: Column(\n    95\t                      cross
AxisAlignment: CrossAxisAlignment.stretch,\n    96\t                      childr
en: [\n    97\t                        Container(\n    98\t                     
    height: 56,\n    99\t                          padding: const EdgeInsets.sym
metric(\n   100\t                            horizontal: DesignTokens.spacingMd,
\n   101\t                          ),\n   102\t                          alignm
ent: Alignment.centerLeft,\n   103\t                          child: Row(\n   10
4\t                            children: [\n   105\t                            
 const Text(\n   106\t                                '开发工具',\n   107\t     
                          style: TextStyle(\n   108\t                           
      fontSize: DesignTokens.fontSizeTitle,\n   109\t                           
      fontWeight: FontWeight.bold,\n   110\t                                  co
lor: DesignTokens.textPrimary,\n   111\t                                ),\n   1
12\t                              ),\n   113\t                              cons
t Spacer(),\n   114\t                              IconButton(\n   115\t        
                       onPressed: widget.onClose,\n   116\t                     
          icon: const Icon(Icons.close),\n   117\t                              
 color: DesignTokens.textPrimary,\n   118\t                              ),\n   
119\t                            ],\n   120\t                          ),\n   12
1\t                        ),\n   122\t                        const Expanded(\n
  123\t                          child: SingleChildScrollView(\n   124\t        
                   padding: EdgeInsets.all(DesignTokens.spacingMd),\n   125\t   
                        child: ServerUrlTool(),\n   126\t                       
  ),\n   127\t                        ),\n   128\t                      ],\n   1
29\t                    ),\n   130\t                  ),\n   131\t              
 ),\n   132\t              ),\n   133\t            ),\n   134\t          ],\n   
135\t        );\n   136\t      },\n   137\t    );\n   138\t  }\n   139\t}\n",
        message='139 lines read from file starting from line 1. Total lines in f
ile: 139. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=6)
ThinkPart(
    type='think',
    think='现在我已经读取了主要的 flow 文件和 mobile 代码。让我继续检查一些关键
的 Widget 文件，比如 empty_recommendation_view.dart、server_url_tool.dart、app.d
art（了解底部导航栏），以及 design-system 中的 SwActionButton 和 MatchListItem 
的 semantics ID。\n\n然后我就可以开始分析已有 flow 是否覆盖了当前功能，并准备执
行测试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_42pmCPVerXhqVl5bb0NSVLV5',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/apps/mobile/lib/presentation/widgets/empty_recommendation_v
iew.dart"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_GJWUVPr9zZo00eXrVSFlhU4e',
    function=FunctionBody(name='ReadFile', arguments='{"path": "/'),
    extras=None
)
ToolResult(
    tool_call_id='tool_42pmCPVerXhqVl5bb0NSVLV5',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:social_world_design_system/social_world_design_system.dart';\n     3\t\
n     4\t/// Empty state when no recommendations are available\n     5\tclass Em
ptyRecommendationView extends StatelessWidget {\n     6\t  final VoidCallback? o
nSwitchNationwide;\n     7\t\n     8\t  const EmptyRecommendationView({\n     9\
t    super.key,\n    10\t    this.onSwitchNationwide,\n    11\t  });\n    12\t\n
   13\t  @override\n    14\t  Widget build(BuildContext context) {\n    15\t    
return Center(\n    16\t      child: Padding(\n    17\t        padding: const Ed
geInsets.all(DesignTokens.spacingMd),\n    18\t        child: Column(\n    19\t 
        mainAxisAlignment: MainAxisAlignment.center,\n    20\t          children
: [\n    21\t            Icon(\n    22\t              Icons.search_off,\n    23\
t              size: 80,\n    24\t              color: DesignTokens.placeholderI
con.withOpacity(0.5),\n    25\t            ),\n    26\t            const SizedBo
x(height: DesignTokens.spacingMd),\n    27\t            const Text(\n    28\t   
          '附近暂无可推荐用户',\n    29\t              style: TextStyle(\n    30
\t                fontSize: DesignTokens.fontSizeTitle,\n    31\t               
fontWeight: FontWeight.bold,\n    32\t                color: DesignTokens.textPr
imary,\n    33\t              ),\n    34\t              textAlign: TextAlign.cen
ter,\n    35\t            ),\n    36\t            const SizedBox(height: DesignT
okens.spacingSm),\n    37\t            Text(\n    38\t              '尝试切换模
式或稍后再试',\n    39\t              style: TextStyle(\n    40\t               
fontSize: DesignTokens.fontSizeBody,\n    41\t                color: DesignToken
s.textPrimary.withOpacity(0.6),\n    42\t              ),\n    43\t             
textAlign: TextAlign.center,\n    44\t            ),\n    45\t            if (on
SwitchNationwide != null) ...[\n    46\t              const SizedBox(height: Des
ignTokens.spacingLg),\n    47\t              Semantics(\n    48\t               
identifier: 'switch_nationwide_button',\n    49\t                label: '切换全
国模式',\n    50\t                button: true,\n    51\t                child: 
GestureDetector(\n    52\t                  behavior: HitTestBehavior.opaque,\n 
  53\t                  onTapDown: (_) => onSwitchNationwide!(),\n    54\t      
           child: Container(\n    55\t                    padding: const EdgeIns
ets.symmetric(\n    56\t                      horizontal: DesignTokens.spacingLg
,\n    57\t                      vertical: DesignTokens.spacingMd,\n    58\t    
               ),\n    59\t                    decoration: BoxDecoration(\n    6
0\t                      color: DesignTokens.primary,\n    61\t                 
    borderRadius:\n    62\t                          BorderRadius.circular(Desig
nTokens.radiusMd),\n    63\t                    ),\n    64\t                    
child: const Text(\n    65\t                      '切换全国模式',\n    66\t     
                style: TextStyle(\n    67\t                        color: Design
Tokens.textOnDarkPrimary,\n    68\t                        fontSize: DesignToken
s.fontSizeBody,\n    69\t                        fontWeight: FontWeight.bold,\n 
  70\t                      ),\n    71\t                    ),\n    72\t        
         ),\n    73\t                ),\n    74\t              ),\n    75\t     
      ],\n    76\t          ],\n    77\t        ),\n    78\t      ),\n    79\t  
 );\n    80\t  }\n    81\t}\n",
        message='81 lines read from file starting from line 1. Total lines in fi
le: 81. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100/apps/mobile/lib/presentation/widgets/server_url_tool.dart"}'
)
ToolCall(
    type='function',
    id='tool_40vY1m88v0Gn7kkMK82Fh3WJ',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_GJWUVPr9zZo00eXrVSFlhU4e',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:flutter_bloc/flutter_bloc.dart';\n     3\timport 'package:social_world_
design_system/social_world_design_system.dart';\n     4\timport '../blocs/dev_to
ols/dev_tools_cubit.dart';\n     5\timport '../blocs/dev_tools/dev_tools_state.d
art';\n     6\t\n     7\tclass ServerUrlTool extends StatefulWidget {\n     8\t 
const ServerUrlTool({super.key});\n     9\t\n    10\t  @override\n    11\t  Stat
e<ServerUrlTool> createState() => _ServerUrlToolState();\n    12\t}\n    13\t\n 
  14\tclass _ServerUrlToolState extends State<ServerUrlTool> {\n    15\t  late f
inal TextEditingController _controller;\n    16\t  late final FocusNode _focusNo
de;\n    17\t\n    18\t  @override\n    19\t  void initState() {\n    20\t    su
per.initState();\n    21\t    _controller = TextEditingController(\n    22\t    
 text: context.read<DevToolsCubit>().state.serverBaseUrl,\n    23\t    );\n    2
4\t    _focusNode = FocusNode();\n    25\t    _controller.addListener(() {\n    
26\t      context.read<DevToolsCubit>().onUrlChanged(_controller.text);\n    27\
t    });\n    28\t    _focusNode.addListener(() {\n    29\t      if (_focusNode.
hasFocus) {\n    30\t        _controller.selection = TextSelection(\n    31\t   
      baseOffset: 0,\n    32\t          extentOffset: _controller.text.length,\n
   33\t        );\n    34\t      }\n    35\t    });\n    36\t  }\n    37\t\n    
38\t  @override\n    39\t  void dispose() {\n    40\t    _controller.dispose();\
n    41\t    _focusNode.dispose();\n    42\t    super.dispose();\n    43\t  }\n 
  44\t\n    45\t  @override\n    46\t  Widget build(BuildContext context) {\n   
47\t    return BlocListener<DevToolsCubit, DevToolsState>(\n    48\t      listen
When: (previous, current) =>\n    49\t          previous.serverBaseUrl.isEmpty &
&\n    50\t          current.serverBaseUrl.isNotEmpty &&\n    51\t          _con
troller.text.isEmpty &&\n    52\t          !_focusNode.hasFocus,\n    53\t      
listener: (context, state) {\n    54\t        _controller.text = state.serverBas
eUrl;\n    55\t      },\n    56\t      child: BlocBuilder<DevToolsCubit, DevTool
sState>(\n    57\t        builder: (context, state) {\n    58\t          return 
SwCard(\n    59\t            child: Column(\n    60\t              crossAxisAlig
nment: CrossAxisAlignment.stretch,\n    61\t              mainAxisSize: MainAxis
Size.min,\n    62\t              children: [\n    63\t                const Row(
\n    64\t                  children: [\n    65\t                    Icon(\n    
66\t                      Icons.public,\n    67\t                      color: De
signTokens.primary,\n    68\t                      size: 20,\n    69\t          
         ),\n    70\t                    SizedBox(width: DesignTokens.spacingXs)
,\n    71\t                    Flexible(\n    72\t                      child: T
ext(\n    73\t                        'Server Base URL',\n    74\t              
         style: TextStyle(\n    75\t                          fontSize: DesignTo
kens.fontSizeBody,\n    76\t                          fontWeight: FontWeight.w60
0,\n    77\t                          color: DesignTokens.textPrimary,\n    78\t
                       ),\n    79\t                        overflow: TextOverflo
w.ellipsis,\n    80\t                      ),\n    81\t                    ),\n 
  82\t                  ],\n    83\t                ),\n    84\t                
const SizedBox(height: DesignTokens.spacingSm),\n    85\t                SwTextF
ield(\n    86\t                  controller: _controller,\n    87\t             
    focusNode: _focusNode,\n    88\t                  hintText: 'http://localhos
t:3000/api/v1',\n    89\t                  labelText: '服务器地址',\n    90\t   
              identifier: 'server_base_url_field',\n    91\t                  ke
yboardType: TextInputType.url,\n    92\t                  prefixIcon: const Icon
(Icons.link),\n    93\t                  errorText: state.urlError.isEmpty ? nul
l : state.urlError,\n    94\t                ),\n    95\t                const S
izedBox(height: DesignTokens.spacingMd),\n    96\t                SizedBox(\n   
97\t                  height: 48,\n    98\t                  child: SwButton(\n 
  99\t                    text: '保存配置',\n   100\t                    variant
: SwButtonVariant.filled,\n   101\t                    isLoading: state.isSaving
,\n   102\t                    isDisabled: state.urlError.isNotEmpty ||\n   103\
t                        state.serverBaseUrl.trim().isEmpty,\n   104\t          
         onPressed: () async {\n   105\t                      try {\n   106\t   
                    await context.read<DevToolsCubit>().save();\n   107\t       
                if (context.mounted) {\n   108\t                          final 
currentState =\n   109\t                              context.read<DevToolsCubit
>().state;\n   110\t                          if (currentState.saveSuccess) {\n 
 111\t                            ScaffoldMessenger.of(context).showSnackBar(\n 
 112\t                              const SnackBar(\n   113\t                   
            content: Text('保存成功，新配置已生效'),\n   114\t                  
           ),\n   115\t                            );\n   116\t                 
        }\n   117\t                        }\n   118\t                      } ca
tch (e) {\n   119\t                        if (context.mounted) {\n   120\t     
                    ScaffoldMessenger.of(context).showSnackBar(\n   121\t       
                    SnackBar(content: Text('保存失败: $e')),\n   122\t          
               );\n   123\t                        }\n   124\t                  
   }\n   125\t                    },\n   126\t                  ),\n   127\t    
           ),\n   128\t              ],\n   129\t            ),\n   130\t       
  );\n   131\t        },\n   132\t      ),\n   133\t    );\n   134\t  }\n   135\
t}\n",
        message='135 lines read from file starting from line 1. Total lines in f
ile: 135. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='00gn/T/dev-issue-55-815001100/apps/mobile/lib/app.dart"}'
)
ToolCall(
    type='function',
    id='tool_F5x7jJ9aRGG3UXlbeZiww4Xb',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_40vY1m88v0Gn7kkMK82Fh3WJ',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/foundation.dart';\n     2\timpor
t 'package:flutter/material.dart';\n     3\timport 'package:flutter_bloc/flutter
_bloc.dart';\n     4\timport 'package:go_router/go_router.dart';\n     5\timport
'package:social_world_design_system/social_world_design_system.dart';\n     6\ti
mport 'core/di/di.dart';\n     7\timport 'presentation/blocs/auth/auth_bloc.dart
';\n     8\timport 'presentation/blocs/auth/auth_event.dart';\n     9\timport 'p
resentation/blocs/auth/auth_state.dart';\n    10\timport 'presentation/blocs/mat
ching/matching_bloc.dart';\n    11\timport 'presentation/blocs/matching/match_li
st_cubit.dart';\n    12\timport 'presentation/pages/login_page.dart';\n    13\ti
mport 'presentation/pages/home_page.dart';\n    14\timport 'presentation/pages/m
atch_list_page.dart';\n    15\timport 'presentation/pages/profile_setup_page.dar
t';\n    16\timport 'presentation/pages/profile_edit_page.dart';\n    17\timport
'presentation/widgets/dev_tools_wrapper.dart';\n    18\t\n    19\tfinal _rootNav
igatorKey = GlobalKey<NavigatorState>();\n    20\t\n    21\tclass SocialWorldApp
extends StatelessWidget {\n    22\t  const SocialWorldApp({super.key});\n    23\
t\n    24\t  @override\n    25\t  Widget build(BuildContext context) {\n    26\t
   return BlocProvider(\n    27\t      create: (_) => getIt<AuthBloc>()..add(con
st AuthCheckRequested()),\n    28\t      child: const _SocialWorldAppBody(),\n  
 29\t    );\n    30\t  }\n    31\t}\n    32\t\n    33\tclass _SocialWorldAppBody
extends StatefulWidget {\n    34\t  const _SocialWorldAppBody();\n    35\t\n    
36\t  @override\n    37\t  State<_SocialWorldAppBody> createState() => _SocialWo
rldAppBodyState();\n    38\t}\n    39\t\n    40\tclass _SocialWorldAppBodyState 
extends State<_SocialWorldAppBody> {\n    41\t  late final GoRouter _router;\n  
 42\t\n    43\t  @override\n    44\t  void initState() {\n    45\t    super.init
State();\n    46\t    _router = GoRouter(\n    47\t      navigatorKey: _rootNavi
gatorKey,\n    48\t      initialLocation: '/login',\n    49\t      routes: [\n  
 50\t        GoRoute(\n    51\t          path: '/login',\n    52\t          buil
der: (context, state) => const LoginPage(),\n    53\t        ),\n    54\t       
GoRoute(\n    55\t          path: '/home',\n    56\t          builder: (context,
state) => const MainPage(),\n    57\t        ),\n    58\t        GoRoute(\n    5
9\t          path: '/profile-setup',\n    60\t          builder: (context, state
) => const ProfileSetupPage(),\n    61\t        ),\n    62\t        GoRoute(\n  
 63\t          path: '/profile-edit',\n    64\t          builder: (context, stat
e) => const ProfileEditPage(),\n    65\t        ),\n    66\t      ],\n    67\t  
 );\n    68\t  }\n    69\t\n    70\t  @override\n    71\t  Widget build(BuildCon
text context) {\n    72\t    return BlocListener<AuthBloc, AuthState>(\n    73\t
     listener: (context, state) {\n    74\t        if (state is AuthAuthenticate
d) {\n    75\t          if (state.isNewUser || !state.profileComplete) {\n    76
\t            _router.go('/profile-setup');\n    77\t          } else {\n    78\
t            _router.go('/home');\n    79\t          }\n    80\t        } else i
f (state is AuthUnauthenticated || state is AuthFailure) {\n    81\t          _r
outer.go('/login');\n    82\t        }\n    83\t      },\n    84\t      child: M
aterialApp.router(\n    85\t        title: 'Social World',\n    86\t        them
e: AppTheme.lightTheme,\n    87\t        darkTheme: AppTheme.darkTheme,\n    88\
t        routerConfig: _router,\n    89\t        builder: (context, child) {\n  
 90\t          if (kDebugMode) {\n    91\t            return DevToolsWrapper(chi
ld: child!);\n    92\t          }\n    93\t          return child!;\n    94\t   
    },\n    95\t      ),\n    96\t    );\n    97\t  }\n    98\t}\n    99\t\n   1
00\tclass MainPage extends StatefulWidget {\n   101\t  const MainPage({super.key
});\n   102\t\n   103\t  @override\n   104\t  State<MainPage> createState() => _
MainPageState();\n   105\t}\n   106\t\n   107\tclass _MainPageState extends Stat
e<MainPage> {\n   108\t  int _currentIndex = 0;\n   109\t\n   110\t  final _page
s = const [\n   111\t    HomePage(),\n   112\t    MatchListPage(),\n   113\t  ];
\n   114\t\n   115\t  void _onTabTapped(int index) {\n   116\t    setState(() =>
_currentIndex = index);\n   117\t  }\n   118\t\n   119\t  @override\n   120\t  W
idget build(BuildContext context) {\n   121\t    return MultiBlocProvider(\n   1
22\t      providers: [\n   123\t        BlocProvider(create: (_) => getIt<Matchi
ngBloc>()),\n   124\t        BlocProvider(create: (_) => getIt<MatchListCubit>()
),\n   125\t      ],\n   126\t      child: Scaffold(\n   127\t        body: _pag
es[_currentIndex],\n   128\t        bottomNavigationBar: Container(\n   129\t   
      decoration: const BoxDecoration(\n   130\t            color: DesignTokens.
background,\n   131\t            border: Border(\n   132\t              top: Bor
derSide(color: DesignTokens.divider),\n   133\t            ),\n   134\t         
),\n   135\t          child: SafeArea(\n   136\t            child: Row(\n   137\
t              children: [\n   138\t                Expanded(\n   139\t         
        child: Semantics(\n   140\t                    identifier: 'tab_home',\n
  141\t                    label: '首页',\n   142\t                    button: t
rue,\n   143\t                    container: true,\n   144\t                    
selected: _currentIndex == 0,\n   145\t                    child: GestureDetecto
r(\n   146\t                      behavior: HitTestBehavior.opaque,\n   147\t   
                  onTapDown: (_) => _onTabTapped(0),\n   148\t                  
   child: Container(\n   149\t                        padding: const EdgeInsets.
symmetric(vertical: 12),\n   150\t                        child: Column(\n   151
\t                          mainAxisSize: MainAxisSize.min,\n   152\t           
              children: [\n   153\t                            Icon(\n   154\t  
                           Icons.home,\n   155\t                              co
lor: _currentIndex == 0\n   156\t                                  ? DesignToken
s.primary\n   157\t                                  : DesignTokens.textPrimary.
withOpacity(0.5),\n   158\t                            ),\n   159\t             
              const SizedBox(height: 4),\n   160\t                            Te
xt(\n   161\t                              '首页',\n   162\t                    
         style: TextStyle(\n   163\t                                fontSize: 12
,\n   164\t                                color: _currentIndex == 0\n   165\t  
                                 ? DesignTokens.primary\n   166\t               
                    : DesignTokens.textPrimary.withOpacity(0.5),\n   167\t      
                       ),\n   168\t                            ),\n   169\t     
                    ],\n   170\t                        ),\n   171\t            
         ),\n   172\t                    ),\n   173\t                  ),\n   17
4\t                ),\n   175\t                Expanded(\n   176\t              
   child: Semantics(\n   177\t                    identifier: 'tab_messages',\n 
 178\t                    label: '消息',\n   179\t                    button: tr
ue,\n   180\t                    container: true,\n   181\t                    s
elected: _currentIndex == 1,\n   182\t                    child: GestureDetector
(\n   183\t                      behavior: HitTestBehavior.opaque,\n   184\t    
                 onTapDown: (_) => _onTabTapped(1),\n   185\t                   
  child: Container(\n   186\t                        padding: const EdgeInsets.s
ymmetric(vertical: 12),\n   187\t                        child: Column(\n   188\
t                          mainAxisSize: MainAxisSize.min,\n   189\t            
             children: [\n   190\t                            Icon(\n   191\t   
                          Icons.chat_bubble,\n   192\t                          
   color: _currentIndex == 1\n   193\t                                  ? Design
Tokens.primary\n   194\t                                  : DesignTokens.textPri
mary.withOpacity(0.5),\n   195\t                            ),\n   196\t        
                   const SizedBox(height: 4),\n   197\t                         
  Text(\n   198\t                              '消息',\n   199\t                
             style: TextStyle(\n   200\t                                fontSize
: 12,\n   201\t                                color: _currentIndex == 1\n   202
\t                                    ? DesignTokens.primary\n   203\t          
                         : DesignTokens.textPrimary.withOpacity(0.5),\n   204\t 
                            ),\n   205\t                            ),\n   206\t
                         ],\n   207\t                        ),\n   208\t       
              ),\n   209\t                    ),\n   210\t                  ),\n
  211\t                ),\n   212\t              ],\n   213\t            ),\n   
214\t          ),\n   215\t        ),\n   216\t      ),\n   217\t    );\n   218\
t  }\n   219\t}\n",
        message='219 lines read from file starting from line 1. Total lines in f
ile: 219. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='5n00000gn/T/dev-issue-55-815001100 && find packages/design-s
ystem -type f -name \\"*.dart\\" | sort"}'
)
StatusUpdate(
    context_usage=0.14647293090820312,
    context_tokens=38397,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=12285,
        output=324,
        input_cache_read=26112,
        input_cache_creation=0
    ),
    message_id='chatcmpl-903EMDYmUrsdieQMAj914ux2',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_F5x7jJ9aRGG3UXlbeZiww4Xb',
    return_value=ToolReturnValue(
        is_error=False,
        output='packages/design-system/lib/social_world_design_system.dart\npack
ages/design-system/lib/src/atoms/sw_action_button.dart\npackages/design-system/l
ib/src/atoms/sw_avatar.dart\npackages/design-system/lib/src/atoms/sw_button.dart
\npackages/design-system/lib/src/atoms/sw_card.dart\npackages/design-system/lib/
src/atoms/sw_text_field.dart\npackages/design-system/lib/src/molecules/match_lis
t_item.dart\npackages/design-system/lib/src/molecules/sw_bottom_sheet.dart\npack
ages/design-system/lib/src/molecules/sw_gender_selector.dart\npackages/design-sy
stem/lib/src/molecules/sw_photo_grid.dart\npackages/design-system/lib/src/molecu
les/sw_range_slider.dart\npackages/design-system/lib/src/molecules/sw_tag_chip.d
art\npackages/design-system/lib/src/molecules/swipe_card.dart\npackages/design-s
ystem/lib/src/organisms/match_success_dialog.dart\npackages/design-system/lib/sr
c/organisms/swipe_card_stack.dart\npackages/design-system/lib/src/theme/app_them
e.dart\npackages/design-system/lib/src/theme/tokens.dart\npackages/design-system
/lib/widgetbook/main.dart\npackages/design-system/test/design_system_golden_test
.dart\npackages/design-system/test/matching_components_golden_test.dart\npackage
s/design-system/test/widget_test.dart\npackages/design-system/widgetbook/main.da
rt\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=7)
ThinkPart(
    type='think',
    think='让我继续查看 design-system 中的关键组件，特别是 SwActionButton 和 Mat
chSuccessDialog，看看它们是否有 semantics identifier，以及 MatchListItem 的 iden
tifier。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_6EqolGIX0swg0xf1n4eDw8Bj',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/packages/design-system/lib/src/atoms/sw_action_button.dart"
}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_pu3lEVt7JmDPQjYjzz6eob4G',
    function=FunctionBody(name='ReadFile', arguments='{"path":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_6EqolGIX0swg0xf1n4eDw8Bj',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'../theme/tokens.dart';\n     3\t\n     4\t/// Circular action button for Pass/L
ike actions\n     5\tclass SwActionButton extends StatelessWidget {\n     6\t  f
inal VoidCallback? onTap;\n     7\t  final SwActionButtonType type;\n     8\t  f
inal double size;\n     9\t\n    10\t  const SwActionButton({\n    11\t    super
.key,\n    12\t    this.onTap,\n    13\t    required this.type,\n    14\t    thi
s.size = 64,\n    15\t  });\n    16\t\n    17\t  @override\n    18\t  Widget bui
ld(BuildContext context) {\n    19\t    final color = type == SwActionButtonType
.like\n    20\t        ? DesignTokens.success\n    21\t        : DesignTokens.er
ror;\n    22\t    final icon = type == SwActionButtonType.like ? Icons.favorite 
: Icons.close;\n    23\t\n    24\t    return Semantics(\n    25\t      identifie
r:\n    26\t          type == SwActionButtonType.like ? 'like_button' : 'pass_bu
tton',\n    27\t      label: type == SwActionButtonType.like ? '喜欢' : '跳过',\
n    28\t      button: true,\n    29\t      container: true,\n    30\t      chil
d: GestureDetector(\n    31\t        behavior: HitTestBehavior.opaque,\n    32\t
       onTapDown: (_) => onTap?.call(),\n    33\t        child: Container(\n    
34\t          width: size,\n    35\t          height: size,\n    36\t          d
ecoration: BoxDecoration(\n    37\t            shape: BoxShape.circle,\n    38\t
           color: color.withOpacity(0.1),\n    39\t            border: Border.al
l(color: color, width: 2),\n    40\t          ),\n    41\t          child: Icon(
\n    42\t            icon,\n    43\t            color: color,\n    44\t        
   size: size * 0.4,\n    45\t          ),\n    46\t        ),\n    47\t      ),
\n    48\t    );\n    49\t  }\n    50\t}\n    51\t\n    52\tenum SwActionButtonT
ype { like, pass }\n",
        message='52 lines read from file starting from line 1. Total lines in fi
le: 52. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/packages/design-system/lib/src/organisms/match_success_dialog.dart
"}'
)
ToolCall(
    type='function',
    id='tool_a2OuNXlSIz1oMNz6JRmIJWRc',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_pu3lEVt7JmDPQjYjzz6eob4G',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'../theme/tokens.dart';\n     3\timport '../atoms/sw_avatar.dart';\n     4\t\n  
  5\t/// Full-screen dialog shown when a mutual match occurs\n     6\tclass Matc
hSuccessDialog extends StatelessWidget {\n     7\t  final String matchedUserNick
name;\n     8\t  final String? matchedUserAvatarUrl;\n     9\t  final String? cu
rrentUserAvatarUrl;\n    10\t  final VoidCallback? onChat;\n    11\t  final Void
Callback? onContinue;\n    12\t\n    13\t  const MatchSuccessDialog({\n    14\t 
  super.key,\n    15\t    required this.matchedUserNickname,\n    16\t    this.m
atchedUserAvatarUrl,\n    17\t    this.currentUserAvatarUrl,\n    18\t    this.o
nChat,\n    19\t    this.onContinue,\n    20\t  });\n    21\t\n    22\t  @overri
de\n    23\t  Widget build(BuildContext context) {\n    24\t    return Semantics
(\n    25\t        identifier: 'match_success_dialog',\n    26\t        label: '
配对成功弹窗',\n    27\t        container: true,\n    28\t        child: Materia
l(\n    29\t          color: DesignTokens.overlayLight,\n    30\t          child
: SafeArea(\n    31\t            child: Padding(\n    32\t              padding:
const EdgeInsets.all(DesignTokens.spacingMd),\n    33\t              child: Colu
mn(\n    34\t                mainAxisAlignment: MainAxisAlignment.center,\n    3
5\t                children: [\n    36\t                  const Spacer(),\n    3
7\t                  const Text(\n    38\t                    '配对成功！',\n   
39\t                    style: TextStyle(\n    40\t                      fontSiz
e: DesignTokens.fontSizeHeadline,\n    41\t                      fontWeight: Fon
tWeight.bold,\n    42\t                      color: DesignTokens.textOnDarkPrima
ry,\n    43\t                    ),\n    44\t                  ),\n    45\t     
            const SizedBox(height: DesignTokens.spacingSm),\n    46\t           
      Text(\n    47\t                    '你和 $matchedUserNickname 互相喜欢了',
\n    48\t                    style: const TextStyle(\n    49\t                 
    fontSize: DesignTokens.fontSizeBody,\n    50\t                      color: D
esignTokens.textOnDarkSecondary,\n    51\t                    ),\n    52\t      
             textAlign: TextAlign.center,\n    53\t                  ),\n    54\
t                  const SizedBox(height: DesignTokens.spacingXl),\n    55\t    
             Row(\n    56\t                    mainAxisAlignment: MainAxisAlignm
ent.center,\n    57\t                    children: [\n    58\t                  
   SwAvatar(\n    59\t                        imageUrl: currentUserAvatarUrl,\n 
  60\t                        initials: '我',\n    61\t                        s
ize: 80,\n    62\t                      ),\n    63\t                      const 
SizedBox(width: DesignTokens.spacingMd),\n    64\t                      const Ic
on(\n    65\t                        Icons.favorite,\n    66\t                  
     color: DesignTokens.success,\n    67\t                        size: 40,\n  
 68\t                      ),\n    69\t                      const SizedBox(widt
h: DesignTokens.spacingMd),\n    70\t                      SwAvatar(\n    71\t  
                     imageUrl: matchedUserAvatarUrl,\n    72\t                  
     initials: matchedUserNickname.isNotEmpty\n    73\t                         
  ? matchedUserNickname.substring(0, 1)\n    74\t                            : '
?',\n    75\t                        size: 80,\n    76\t                      ),
\n    77\t                    ],\n    78\t                  ),\n    79\t        
         const Spacer(),\n    80\t                  Semantics(\n    81\t        
           identifier: 'match_dialog_chat_button',\n    82\t                    
label: '开始聊天',\n    83\t                    button: true,\n    84\t         
          child: GestureDetector(\n    85\t                      behavior: HitTe
stBehavior.opaque,\n    86\t                      onTapDown: (_) => onChat?.call
(),\n    87\t                      child: Container(\n    88\t                  
     width: double.infinity,\n    89\t                        padding: const Edg
eInsets.symmetric(\n    90\t                          vertical: DesignTokens.spa
cingMd,\n    91\t                        ),\n    92\t                        dec
oration: BoxDecoration(\n    93\t                          color: DesignTokens.p
rimary,\n    94\t                          borderRadius:\n    95\t              
               BorderRadius.circular(DesignTokens.radiusMd),\n    96\t          
             ),\n    97\t                        child: const Center(\n    98\t 
                        child: Text(\n    99\t                            '开始
聊天',\n   100\t                            style: TextStyle(\n   101\t         
                    color: DesignTokens.textOnDarkPrimary,\n   102\t            
                 fontSize: DesignTokens.fontSizeBody,\n   103\t                 
            fontWeight: FontWeight.bold,\n   104\t                            ),
\n   105\t                          ),\n   106\t                        ),\n   1
07\t                      ),\n   108\t                    ),\n   109\t          
       ),\n   110\t                  const SizedBox(height: DesignTokens.spacing
Md),\n   111\t                  Semantics(\n   112\t                    identifi
er: 'match_dialog_continue_button',\n   113\t                    label: '继续滑
动',\n   114\t                    button: true,\n   115\t                    chi
ld: GestureDetector(\n   116\t                      behavior: HitTestBehavior.op
aque,\n   117\t                      onTapDown: (_) => onContinue?.call(),\n   1
18\t                      child: Container(\n   119\t                        wid
th: double.infinity,\n   120\t                        padding: const EdgeInsets.
symmetric(\n   121\t                          vertical: DesignTokens.spacingMd,\
n   122\t                        ),\n   123\t                        decoration:
BoxDecoration(\n   124\t                          color: Colors.transparent,\n  
125\t                          borderRadius:\n   126\t                          
   BorderRadius.circular(DesignTokens.radiusMd),\n   127\t                      
   border: Border.all(\n   128\t                              color: DesignToken
s.textOnDarkTertiary),\n   129\t                        ),\n   130\t            
           child: const Center(\n   131\t                          child: Text(\
n   132\t                            '继续滑动',\n   133\t                      
     style: TextStyle(\n   134\t                              color: DesignToken
s.textOnDarkPrimary,\n   135\t                              fontSize: DesignToke
ns.fontSizeBody,\n   136\t                            ),\n   137\t              
           ),\n   138\t                        ),\n   139\t                     
),\n   140\t                    ),\n   141\t                  ),\n   142\t      
         ],\n   143\t              ),\n   144\t            ),\n   145\t         
),\n   146\t        ));\n   147\t  }\n   148\t}\n",
        message='148 lines read from file starting from line 1. Total lines in f
ile: 148. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='001100/packages/design-system/lib/src/molecules/match_list_i
tem.dart"}'
)
ToolCall(
    type='function',
    id='tool_DrOTU6KX8h9vLdDW5RbdEkAN',
    function=FunctionBody(name='ReadFile', arguments='{"path":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_a2OuNXlSIz1oMNz6JRmIJWRc',
    return_value=ToolOk(
        is_error=False,
        output='     1\timport \'package:flutter/material.dart\';\n     2\timpor
t \'../theme/tokens.dart\';\n     3\timport \'../atoms/sw_avatar.dart\';\n     4
\t\n     5\t/// List item showing a matched user\n     6\tclass MatchListItem ex
tends StatelessWidget {\n     7\t  final String nickname;\n     8\t  final Strin
g? avatarUrl;\n     9\t  final String? matchedAt;\n    10\t  final VoidCallback?
onTap;\n    11\t\n    12\t  const MatchListItem({\n    13\t    super.key,\n    1
4\t    required this.nickname,\n    15\t    this.avatarUrl,\n    16\t    this.ma
tchedAt,\n    17\t    this.onTap,\n    18\t  });\n    19\t\n    20\t  @override\
n    21\t  Widget build(BuildContext context) {\n    22\t    return Semantics(\n
   23\t      identifier: \'match_list_item\',\n    24\t      label: \'$nickname,
匹配时间: ${matchedAt ?? ""}\',\n    25\t      button: true,\n    26\t      cont
ainer: true,\n    27\t      child: GestureDetector(\n    28\t        behavior: H
itTestBehavior.opaque,\n    29\t        onTapDown: (_) => onTap?.call(),\n    30
\t        child: Container(\n    31\t          padding: const EdgeInsets.symmetr
ic(\n    32\t            horizontal: DesignTokens.spacingMd,\n    33\t          
 vertical: DesignTokens.spacingSm,\n    34\t          ),\n    35\t          chil
d: Row(\n    36\t            children: [\n    37\t              SwAvatar(\n    3
8\t                imageUrl: avatarUrl,\n    39\t                initials: nickn
ame.isNotEmpty ? nickname.substring(0, 1) : \'?\',\n    40\t                size
: 56,\n    41\t              ),\n    42\t              const SizedBox(width: Des
ignTokens.spacingMd),\n    43\t              Expanded(\n    44\t                
child: Column(\n    45\t                  crossAxisAlignment: CrossAxisAlignment
.start,\n    46\t                  children: [\n    47\t                    Text
(\n    48\t                      nickname,\n    49\t                      style:
const TextStyle(\n    50\t                        fontSize: DesignTokens.fontSiz
eBody,\n    51\t                        fontWeight: FontWeight.w600,\n    52\t  
                     color: DesignTokens.textPrimary,\n    53\t                 
    ),\n    54\t                    ),\n    55\t                    if (matchedA
t != null) ...[\n    56\t                      const SizedBox(height: 4),\n    5
7\t                      Text(\n    58\t                        matchedAt!,\n   
59\t                        style: const TextStyle(\n    60\t                   
      fontSize: DesignTokens.fontSizeCaption,\n    61\t                         
color: DesignTokens.placeholderIcon,\n    62\t                        ),\n    63
\t                      ),\n    64\t                    ],\n    65\t            
     ],\n    66\t                ),\n    67\t              ),\n    68\t         
    const Icon(\n    69\t                Icons.chevron_right,\n    70\t         
      color: DesignTokens.placeholderIcon,\n    71\t              ),\n    72\t  
         ],\n    73\t          ),\n    74\t        ),\n    75\t      ),\n    76\
t    );\n    77\t  }\n    78\t}\n',
        message='78 lines read from file starting from line 1. Total lines in fi
le: 78. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/packages/design-system/lib/src/molecules/sw_gender_selector.dart"}
'
)
StatusUpdate(
    context_usage=0.16501235961914062,
    context_tokens=43257,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=5113,
        output=284,
        input_cache_read=38144,
        input_cache_creation=0
    ),
    message_id='chatcmpl-7WFrTVgqdSbuYHwHXitknF1P',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_DrOTU6KX8h9vLdDW5RbdEkAN',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'../theme/tokens.dart';\n     3\timport '../atoms/sw_button.dart';\n     4\t\n  
  5\t/// A gender selector chip group using design-system styling.\n     6\tclas
s SwGenderSelector extends StatelessWidget {\n     7\t  final String? selectedGe
nder;\n     8\t  final ValueChanged<String?> onChanged;\n     9\t  final List<St
ring> options;\n    10\t\n    11\t  const SwGenderSelector({\n    12\t    super.
key,\n    13\t    required this.selectedGender,\n    14\t    required this.onCha
nged,\n    15\t    this.options = const ['male', 'female', 'other'],\n    16\t  
});\n    17\t\n    18\t  static String _label(String gender) {\n    19\t    swit
ch (gender) {\n    20\t      case 'male':\n    21\t        return '男';\n    22\
t      case 'female':\n    23\t        return '女';\n    24\t      case 'all':\n
   25\t        return '全部';\n    26\t      case 'other':\n    27\t      defaul
t:\n    28\t        return '其他';\n    29\t    }\n    30\t  }\n    31\t\n    32
\t  @override\n    33\t  Widget build(BuildContext context) {\n    34\t    retur
n Row(\n    35\t      children: options.map((gender) {\n    36\t        final is
Selected = selectedGender == gender;\n    37\t        return Expanded(\n    38\t
         child: Padding(\n    39\t            padding: const EdgeInsets.symmetri
c(horizontal: DesignTokens.spacingXs),\n    40\t            child: SwButton(\n  
 41\t              text: _label(gender),\n    42\t              variant: isSelec
ted ? SwButtonVariant.filled : SwButtonVariant.outlined,\n    43\t              
onPressed: () => onChanged(isSelected ? null : gender),\n    44\t            ),\
n    45\t          ),\n    46\t        );\n    47\t      }).toList(),\n    48\t 
  );\n    49\t  }\n    50\t}\n",
        message='50 lines read from file starting from line 1. Total lines in fi
le: 50. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=8)
ThinkPart(
    type='think',
    think='现在我已经对项目有了全面的了解。让我总结一下：\n\n**已有 flows：**\n1
. `auth/login.yaml` - 登录流程（smoke）\n2. `profile/setup.yaml` - 完善资料流程
（full）\n3. `matching/swipe_recommendations.yaml` - 推荐卡片滑动（full）\n4. `m
atching/view_match_list.yaml` - 查看匹配列表（full）\n5. `devtools/open_drawer.y
aml` - 开发工具抽屉（smoke）\n6. `devtools/edit_server_url.yaml` - 编辑服务器 UR
L（regression）\n\n**PRD 需求覆盖分析：**\n- US-1 注册登录: ✅ `auth/login.yaml`
\n- US-2 完善资料: ✅ `profile/setup.yaml`\n- US-3 推荐卡片与滑动匹配: ✅ `match
ing/swipe_recommendations.yaml`\n- US-4 即时通讯: 部分（匹配列表）✅ `matching/v
iew_match_list.yaml`\n- devtools: ✅ `devtools/open_drawer.yaml`, `devtools/edit
_server_url.yaml`\n\n看起来已有的 flow 基本覆盖了 PRD 中的主要用户旅程。但是有一
些问题需要注意：\n\n1. `profile/setup.yaml` 中：\n   - `tapOn: "昵称"` - 但 `SwT
extField` 的 labelText 是 "昵称"，这个可能在 accessibility tree 中显示为 label\n
  - `tapOn: "男"` - SwGenderSelector 中的按钮文字是 "男"\n   - `tapOn: "出生日期
"` - SwTextField 的 labelText\n   - `tapOn: "OK"` - 但 Flutter 的 showDatePicker
在中文环境下按钮可能是 "确定" 而不是 "OK"\n\n2. `matching/swipe_recommendations.
yaml` 中：\n   - 使用 swipe LEFT/RIGHT，这看起来没问题\n   - 但 `assertVisible: 
"附近"` 等模式 Tab 的文字在 home_page.dart 中确实是这些文字\n\n3. `matching/view
_match_list.yaml` 中：\n   - 空状态文字 "还没有匹配成功的好友" 和 "继续滑动，发
现感兴趣的人吧" 与代码一致\n\n现在让我执行测试。根据 README，环境应该已经准备好
了（APK 已构建并安装到模拟器）。\n\n让我先检查模拟器是否就绪，然后执行测试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_KgIfZGSZitj62JK8gQ2ymSUt',
    function=FunctionBody(
        name='SetTodoList',
        arguments='{"todos": [{"title":"阅读项目地图与E2E规范","status":"done"},
{"title":"扫描已有flow并分析覆盖范围","status":"done"},{"title":"检查mobile代码
中的UI元素和页面结构","status":"done"},{"title":"修复includes格式问题","status":
"in_progress"},{"title":"执行Maestro测试","status":"pending"},{"title":"分析失败
根因并修复flow","status":"pending"},{"title":"生成E2E测试报告","status":"pending
"}]}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_8hj2TvkAlCbaNsjJixe5eD1r',
    function=FunctionBody(name='Shell', arguments='{"command": "'),
    extras=None
)
ToolResult(
    tool_call_id='tool_KgIfZGSZitj62JK8gQ2ymSUt',
    return_value=ToolReturnValue(
        is_error=False,
        output='Todo list updated',
        message='Todo list updated',
        display=[
            TodoDisplayBlock(
                type='todo',
                items=[
                    TodoDisplayItem(
                        title='阅读项目地图与E2E规范',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='扫描已有flow并分析覆盖范围',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='检查mobile代码中的UI元素和页面结构',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='修复includes格式问题',
                        status='in_progress'
                    ),
                    TodoDisplayItem(title='执行Maestro测试', status='pending'),
                    TodoDisplayItem(
                        title='分析失败根因并修复flow',
                        status='pending'
                    ),
                    TodoDisplayItem(title='生成E2E测试报告', status='pending')
                ]
            )
        ],
        extras=None
    )
)
ToolCallPart(
    arguments_part='cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issu
e-55-815001100 && adb devices"}'
)
ToolCall(
    type='function',
    id='tool_kvENXnCs1zmDS2d0uSoNRY7Z',
    function=FunctionBody(name='Shell', arguments='{"command": "cd /'),
    extras=None
)
ToolResult(
    tool_call_id='tool_8hj2TvkAlCbaNsjJixe5eD1r',
    return_value=ToolReturnValue(
        is_error=False,
        output='List of devices attached\nemulator-5556\tdevice\n\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100 && echo \\"AGENT_E2E_DEBUG_DIR=${AGENT_E2E_DEBUG_DIR}\\" && echo \\"A
GENT_E2E_REPORT_FILE=${AGENT_E2E_REPORT_FILE}\\""}'
)
StatusUpdate(
    context_usage=0.17818069458007812,
    context_tokens=46709,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=3701,
        output=743,
        input_cache_read=43008,
        input_cache_creation=0
    ),
    message_id='chatcmpl-4QI7CXfI7qf10RPHeuMRuxsv',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_kvENXnCs1zmDS2d0uSoNRY7Z',
    return_value=ToolReturnValue(
        is_error=False,
        output='AGENT_E2E_DEBUG_DIR=/Users/insulate/Desktop/new-world/new-world-
ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302\nAGENT_E2E_RE
PORT_FILE=/Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-
report-e2e-issue-55-20260424-012302.md\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=9)
ThinkPart(
    type='think',
    think='模拟器已就绪，环境变量也设置了。让我先检查 includes 文件是否有 appId 
和 `---` 分隔符。`login-as-test-user.yaml` 看起来格式正确。\n\n不过让我再检查 `.
maestro/config.yaml`，然后执行测试。先确认 APK 是否已安装。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_PTLBQYcpHbTTrU8lPxFhpzGh',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/config.yaml"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_xNxSUxW6llHYa6RHlwg3Ow4x',
    function=FunctionBody(name='Shell', arguments='{"command": "'),
    extras=None
)
ToolResult(
    tool_call_id='tool_PTLBQYcpHbTTrU8lPxFhpzGh',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# Maestro 全局配置\n     2\t# 本文件是 Maestro 子系统的
唯一真相来源。\n     3\t\n     4\tappId: "com.example.mobile"\n     5\t\n     6\
t# 默认超时配置\n     7\textendedWaitUntil:\n     8\t  timeout: 10000    # 10 秒
\n     9\t\n    10\t# 环境变量默认值\n    11\tenv:\n    12\t  TEST_PHONE: "13800
138000"\n    13\t  TEST_SMS_CODE: "123456"\n    14\t  API_BASE_URL: "http://loca
lhost:3000"\n',
        message='14 lines read from file starting from line 1. Total lines in fi
le: 14. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issu
e-55-815001100 && adb shell pm list packages | grep com.example.mobile"}'
)
ToolCall(
    type='function',
    id='tool_MN5rlildqhoncGhHojykEIzw',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "which maestro && maestro --version"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.18166351318359375,
    context_tokens=47622,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1030,
        output=203,
        input_cache_read=46592,
        input_cache_creation=0
    ),
    message_id='chatcmpl-tYwZSD38E4DvV0ATcmaaqbqw',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_xNxSUxW6llHYa6RHlwg3Ow4x',
    return_value=ToolReturnValue(
        is_error=False,
        output='package:com.example.mobile\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolResult(
    tool_call_id='tool_MN5rlildqhoncGhHojykEIzw',
    return_value=ToolReturnValue(
        is_error=False,
        output='/Users/insulate/.maestro/maestro/bin/maestro\n1.39.13\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=10)
ThinkPart(
    type='think',
    think='一切就绪：\n- 模拟器 emulator-5556 在线\n- APK 已安装 (com.example.mo
bile)\n- Maestro 1.39.13 已安装\n- DEBUG_DIR 和 REPORT_FILE 环境变量已设置\n\n现
在我需要执行 Maestro 测试。我应该运行 `.maestro/flows/` 下的所有 yaml 文件。但需
要考虑：\n1. `profile/setup.yaml`, `matching/swipe_recommendations.yaml`, `match
ing/view_match_list.yaml` 依赖后端\n2. `auth/login.yaml` 是纯 UI，不依赖后端\n3.
`devtools/edit_server_url.yaml` 依赖 `devtools/open_drawer`\n\n让我先运行所有 fl
ows，看看哪些通过哪些失败。\n\n按照 README 的说法，先启动后端环境：\n```bash\nma
ke e2e-env-up\n```\n\n但环境变量中 API_BASE_URL 是 `http://localhost:3000`，而 M
aestro 运行在模拟器里，服务器地址应该是 `http://10.0.2.2:3001/api/v1`（从 edit_s
erver_url.yaml 的恢复地址推断）。\n\n让我先检查后端是否已经在运行。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_zJOs1y359LqTVuPsevQHQR3g',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100 && curl -s http://localhost:3000/health || curl -s ht
tp://localhost:3001/health || echo \\"Backend not running on 3000/3001\\""}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_pwPRlMaKZcWWFruaBmHtqUw8',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_zJOs1y359LqTVuPsevQHQR3g',
    return_value=ToolReturnValue(
        is_error=False,
        output='{"code":2000,"message":"Cannot GET /health","data":null}',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000
gn/T/dev-issue-55-815001100 && docker ps 2>/dev/null || echo \\"Docker not avail
able or no containers\\""}'
)
StatusUpdate(
    context_usage=0.18343353271484375,
    context_tokens=48086,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=470,
        output=396,
        input_cache_read=47616,
        input_cache_creation=0
    ),
    message_id='chatcmpl-m9QDG7Ep0RYeFCwsJSBdmhQV',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_pwPRlMaKZcWWFruaBmHtqUw8',
    return_value=ToolReturnValue(
        is_error=False,
        output='CONTAINER ID   IMAGE            COMMAND                  CREATED
     STATUS                  PORTS                                         NAMES
\nc857fae87156   redis:7-alpine   "docker-entrypoint.s…"   7 days ago   Up 32 ho
urs (healthy)   0.0.0.0:6379->6379/tcp, [::]:6379->6379/tcp   sw-redis\n30bb29bd
5b31   mysql:8.0        "docker-entrypoint.s…"   7 days ago   Up 32 hours (healt
hy)   0.0.0.0:3306->3306/tcp, [::]:3306->3306/tcp   sw-mysql\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=11)
ThinkPart(
    type='think',
    think='后端在端口 3000 上似乎有服务在运行（返回了 Cannot GET /health 的 API 
错误格式，说明 NestJS 在运行但 /health 路由不存在）。让我检查 API 的实际可用端口
和路由。同时 MySQL 和 Redis 也在运行。\n\n让我检查 API 的可用性，以及端点。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_59eRYpmJXkBdJoj1ITPgmJX0',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100 && curl -s http://localhost:3000/api/v1/auth/me -H \\
"Authorization: Bearer test\\" | head -c 200"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_DkEXbErI4gp4o1sb8VJdGkLK',
    function=FunctionBody(name='Shell', arguments='{"command": "cd /'),
    extras=None
)
ToolResult(
    tool_call_id='tool_59eRYpmJXkBdJoj1ITPgmJX0',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100 && ps aux | grep -E \\"(node|nest)\\" | grep -v grep | head -10"}'
)
StatusUpdate(
    context_usage=0.1857452392578125,
    context_tokens=48692,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=820,
        output=230,
        input_cache_read=47872,
        input_cache_creation=0
    ),
    message_id='chatcmpl-jIv91RQxIESqrrZ9JuIuPDcK',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_DkEXbErI4gp4o1sb8VJdGkLK',
    return_value=ToolReturnValue(
        is_error=False,
        output='insulate         15660  99.1  0.0 452668304   7856 s005  R+   Tu
e02PM 2398:12.89 node (vitest 1)     \ninsulate         14199  98.9  0.0 4526713
60   7904 s005  R+   Tue02PM 2400:17.86 node (vitest 1)     \ninsulate         4
6260   0.5  0.1 436147952  15808   ??  S     2:34PM   1:19.30 /opt/homebrew/bin/
node --max_old_space_size=3072 /Users/insulate/Applications/WebStorm.app/Content
s/plugins/javascript-plugin/jsLanguageServicesImpl/js-language-service.js -id=17
76926085051 -debug-name=typescript --pluginProbeLocations /Users/insulate/Deskto
p/agflux,/Users/insulate/Applications/WebStorm.app/Contents/plugins/javascript-p
lugin/jsLanguageServicesImpl/typescript --globalPlugins tsc-ide-plugin --cancell
ationPipeName /private/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/tsServerC
ancellationPipes8/cmd*\ninsulate         85414   0.1  0.1 444282512  26720   ?? 
Ss    9:26AM   0:33.68 node /Users/insulate/Desktop/agflux/apps/server/dist/serv
er.js \ninsulate         67041   0.0  0.2 444311008  43776   ??  Ss   11:15PM   
0:04.41 node /Users/insulate/Desktop/games/packages/server/dist/index.js \ninsul
ate         42241   0.0  0.1 444253264  30768 s025  S+    9:40PM   0:00.38 node 
/Users/insulate/Desktop/games/node_modules/.bin/prisma studio --port 5556\ninsul
ate         59623   0.0  0.0 410990576   8448 s018  S     3:42PM   0:06.87 /User
s/insulate/Desktop/games/node_modules/@esbuild/darwin-arm64/bin/esbuild --servic
e=0.21.5 --ping\ninsulate         59621   0.0  0.1 452676208  26960 s018  S     
3:42PM   0:11.86 node /Users/insulate/Desktop/games/node_modules/.bin/vite\ninsu
late         46261   0.0  0.0 435740832   7888   ??  S     2:34PM   0:00.11 /opt
/homebrew/Cellar/node/23.11.0/bin/node /Users/insulate/Desktop/agflux/node_modul
es/.pnpm/typescript@5.9.3/node_modules/typescript/lib/typingsInstaller.js --glob
alTypingsCacheLocation /Users/insulate/Library/Caches/typescript/5.9 --typesMapL
ocation /Users/insulate/Desktop/agflux/node_modules/.pnpm/typescript@5.9.3/node_
modules/typescript/lib/typesMap.json\ninsulate         46225   0.0  0.0 43585356
8   9520   ??  S     2:34PM   0:02.76 /opt/homebrew/bin/node /Users/insulate/App
lications/WebStorm.app/Contents/plugins/tailwindcss/server/tailwindcss-language-
server\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=12)
ThinkPart(
    type='think',
    think='看起来有一个 Node.js 服务器在运行：`node /Users/insulate/Desktop/agfl
ux/apps/server/dist/server.js`。但 curl 返回空，可能端口不是 3000 或者 API 路径
不同。这其实是另一个项目（agflux），不是我们的 social-world。\n\n让我检查 social
-world 的后端是否在运行。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_XlM3nXW2T1NSqh1WzAfFRtbj',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100 && ps aux | grep -E \\"social|world|issue-55\\" | gre
p -v grep"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_7SK3BdWm7as5fPUzTMmbqDX1',
    function=FunctionBody(name='Shell', arguments='{"command": "'),
    extras=None
)
ToolResult(
    tool_call_id='tool_XlM3nXW2T1NSqh1WzAfFRtbj',
    return_value=ToolReturnValue(
        is_error=False,
        output='insulate         68689   0.0  0.1 411403024  17344   ??  Ss   11
:18PM   0:03.36 /Users/insulate/Desktop/new-world/new-world-ops/build/dashboard\
ninsulate         21980   0.0  0.0 436142048  11984   ??  S    Wed10PM   3:29.75
node /Users/insulate/Desktop/social-world/apps/api/node_modules/.bin/../@nestjs/
cli/bin/nest.js start --watch\ninsulate         92754   0.0  0.1 436136368  1870
4 s003  S+   Wed05PM   3:52.88 node /Users/insulate/Desktop/social-world/apps/ap
i/node_modules/.bin/../@nestjs/cli/bin/nest.js start --watch\ninsulate         4
2178   0.0  0.0 435681168   8720   ??  S    Wed08AM   0:24.02 /Users/insulate/fl
utter/bin/cache/artifacts/engine/darwin-x64/flutter_tester --disable-vm-service 
--enable-checked-mode --verify-entry-points --enable-software-rendering --skia-d
eterministic-rendering --enable-dart-profiling --non-interactive --use-test-font
s --disable-asset-fonts --packages=/private/var/folders/lt/w3sl6z8j5v50vt2s6knvf
5n00000gn/T/dev-issue-55-600964399/packages/design-system/.dart_tool/package_con
fig.json --flutter-assets-dir=/private/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000
00gn/T/dev-issue-55-600964399/packages/design-system/build/unit_test_assets /var
/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/flutter_tools.3PfsEm/flutter_test_l
istener.g2VZtq/listener.dart.dill\ninsulate         69385   0.0  0.0 435299728  
 512 s012  S+   Tue06PM   0:00.00 /bin/bash ./pm-social-world.sh\ninsulate      
  11645   0.0  0.0 435299840    512 s016  S+   Tue02PM   0:00.00 /bin/bash ./pm-
new-world.sh\ninsulate          5845   0.0  0.0 435299328    528 s010  S+    1:2
3AM   0:00.01 tee -a /Users/insulate/Desktop/new-world/new-world-ops/logs/social
-world/e2e-issue-55-20260424-012302.log\ninsulate          5842   0.0  0.0 43530
8208   1104 s010  Ss+   1:23AM   0:00.01 zsh -c (echo ""\\012echo "=== 运行 E2E 
测试 Agent ==="\\012echo "Agent 文件: \'/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n0
0000gn/T/dashboard-agent-social-world-e2e-tester-68689.yaml\'"\\012echo "Debug 
目录: \'/Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-sc
reenshots-e2e-issue-55-20260424-012302\'"\\012echo "报告文件: \'/Users/insulate/
Desktop/new-world/new-world-ops/logs/social-world/e2e-report-e2e-issue-55-202604
24-012302.md\'"\\012echo ""\\012mkdir -p "\'/Users/insulate/Desktop/new-world/ne
w-world-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302\'"\\0
12kimi --agent-file \'/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dashboard
-agent-social-world-e2e-tester-68689.yaml\' \\\\012     --max-steps-per-turn 100
00 \\\\012     --print \\\\012     -p \'【任务上下文】\\012- Issue 编号: #55\\01
2- 项目路径: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-81500
1100\\012- 项目地图: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-
55-815001100/AGENTS.md\\012- 关联 PRD: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00
000gn/T/dev-issue-55-815001100/prd/v1-profile.md,/var/folders/lt/w3sl6z8j5v50vt2
s6knvf5n00000gn/T/dev-issue-55-815001100/prd/v1-mvp.md,/var/folders/lt/w3sl6z8j5
v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/prd/v1-matching.md\\012- 设计资产
: 未找到\\012- 文档目录: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-is
sue-55-815001100/docs\\012\\012\\012\\012**你的任务类型：E2E 测试任务（E2E Teste
r Agent）**\\012目标：执行 Maestro E2E 测试，分析失败原因，生成详细报告。\\012\\
012Issue #55，分支: feat/issue-43\\012\'\\012) 2>&1 | tee -a /Users/insulate/Des
ktop/new-world/new-world-ops/logs/social-world/e2e-issue-55-20260424-012302.log;
echo $? > /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/nwops-exit-e2e-55-177
6965003291486000\ninsulate          5841   0.0  0.0 435310000   4448   ??  Ss   
1:23AM   0:00.03 tmux new-session -d -s nwops-e2e-55 -c /var/folders/lt/w3sl6z8j
5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100 (echo ""\\012echo "=== 运行 E2E 
测试 Agent ==="\\012echo "Agent 文件: \'/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n0
0000gn/T/dashboard-agent-social-world-e2e-tester-68689.yaml\'"\\012echo "Debug 
目录: \'/Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-sc
reenshots-e2e-issue-55-20260424-012302\'"\\012echo "报告文件: \'/Users/insulate/
Desktop/new-world/new-world-ops/logs/social-world/e2e-report-e2e-issue-55-202604
24-012302.md\'"\\012echo ""\\012mkdir -p "\'/Users/insulate/Desktop/new-world/ne
w-world-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302\'"\\0
12kimi --agent-file \'/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dashboard
-agent-social-world-e2e-tester-68689.yaml\' \\\\012     --max-steps-per-turn 100
00 \\\\012     --print \\\\012     -p \'【任务上下文】\\012- Issue 编号: #55\\01
2- 项目路径: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-81500
1100\\012- 项目地图: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-
55-815001100/AGENTS.md\\012- 关联 PRD: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00
000gn/T/dev-issue-55-815001100/prd/v1-profile.md,/var/folders/lt/w3sl6z8j5v50vt2
s6knvf5n00000gn/T/dev-issue-55-815001100/prd/v1-mvp.md,/var/folders/lt/w3sl6z8j5
v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/prd/v1-matching.md\\012- 设计资产
: 未找到\\012- 文档目录: /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-is
sue-55-815001100/docs\\012\\012\\012\\012**你的任务类型：E2E 测试任务（E2E Teste
r Agent）**\\012目标：执行 Maestro E2E 测试，分析失败原因，生成详细报告。\\012\\
012Issue #55，分支: feat/issue-43\\012\'\\012) 2>&1 | tee -a /Users/insulate/Des
ktop/new-world/new-world-ops/logs/social-world/e2e-issue-55-20260424-012302.log;
echo $? > /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/nwops-exit-e2e-55-177
6965003291486000\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issu
e-55-815001100 && ls apps/api/"}'
)
StatusUpdate(
    context_usage=0.1897735595703125,
    context_tokens=49748,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1108,
        output=198,
        input_cache_read=48640,
        input_cache_creation=0
    ),
    message_id='chatcmpl-TnirLEDEk462uAxCFaHNya6I',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_7SK3BdWm7as5fPUzTMmbqDX1',
    return_value=ToolReturnValue(
        is_error=False,
        output='Dockerfile.dev\nnode_modules\npackage-lock.json\npackage.json\ns
rc\ntsconfig.json\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=13)
ThinkPart(
    type='think',
    think='看起来 social-world 的后端（NestJS）已经在运行了，有两个进程：\n- `no
de /Users/insulate/Desktop/social-world/apps/api/node_modules/.bin/../@nestjs/cl
i/bin/nest.js start --watch`\n- 另一个相同的进程\n\n但注意：实际工作目录是 `/var
/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100`，而运行中的
NestJS 是在 `/Users/insulate/Desktop/social-world/apps/api`。这可能是另一个实例
。\n\n让我检查当前项目（dev-issue-55-815001100）的后端是否在运行，以及端口。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_2uSLmnrmYuJEtAsGwHplCvGp',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100 && lsof -ti :3000 | xargs -I{} ps -p {} -o pid,args |
grep -v PID || echo \\"Port 3000 not used by expected process\\""}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_CB5SWs2yXLH6iQlIz36gXJ8T',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_2uSLmnrmYuJEtAsGwHplCvGp',
    return_value=ToolReturnValue(
        is_error=False,
        output='Port 3000 not used by expected process\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000
gn/T/dev-issue-55-815001100 && lsof -ti :3001 | xargs -I{} ps -p {} -o pid,args 
| grep -v PID || echo \\"Port 3001 not used\\""}'
)
StatusUpdate(
    context_usage=0.19921493530273438,
    context_tokens=52223,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=2559,
        output=328,
        input_cache_read=49664,
        input_cache_creation=0
    ),
    message_id='chatcmpl-enEO9ZYtapSwUeDu8HmkE5f9',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_CB5SWs2yXLH6iQlIz36gXJ8T',
    return_value=ToolReturnValue(
        is_error=False,
        output='14715 node dist/main\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=14)
ThinkPart(
    type='think',
    think='端口 3001 上有 `node dist/main` 在运行，这很可能是当前项目的后端。让
我验证一下这个进程的工作目录。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_nEidmosJSwpjOAWjTNxvOlgP',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "ps -p 14715 -o pid,args,cwd | grep -v PID"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_U88zwpQWF6gjJUSq6tNPyHox',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_nEidmosJSwpjOAWjTNxvOlgP',
    return_value=ToolReturnValue(
        is_error=False,
        output='ps: cwd: keyword not found\n14715 node dist/main\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"command": "curl -s http://localhost:3001/api/v1/auth/me -H
\\"Authorization: Bearer test\\" | head -c 200"}'
)
StatusUpdate(
    context_usage=0.20071029663085938,
    context_tokens=52615,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=391,
        output=105,
        input_cache_read=52224,
        input_cache_creation=0
    ),
    message_id='chatcmpl-0HSLgv3dbaketibiqdoSgWyi',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_U88zwpQWF6gjJUSq6tNPyHox',
    return_value=ToolReturnValue(
        is_error=False,
        output='{"code":2001,"message":"Invalid or expired token","data":null}',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=15)
ThinkPart(
    type='think',
    think='后端在 3001 上运行，返回了标准的 API 错误格式（Invalid or expired tok
en），说明 NestJS 后端是正常运行的。这很好。\n\n现在我可以执行 Maestro 测试了。
让我运行所有 flow。为了安全起见，我先创建 debug 目录。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_af5BBoZ4RKbejrCTxVmV8Mgl',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "mkdir -p /Users/insulate/Desktop/new-world/new-w
orld-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_ZS8r3zynxRNQF5TgJFKgRM7C',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_af5BBoZ4RKbejrCTxVmV8Mgl',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000
gn/T/dev-issue-55-815001100/.maestro && maestro test --debug-output=\\"/Users/in
sulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-iss
ue-55-20260424-012302\\" --flatten-debug-output flows/auth/login.yaml flows/prof
ile/setup.yaml flows/matching/swipe_recommendations.yaml flows/matching/view_mat
ch_list.yaml flows/devtools/open_drawer.yaml flows/devtools/edit_server_url.yaml
", "timeout": 300}'
)
StatusUpdate(
    context_usage=0.20141220092773438,
    context_tokens=52799,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=319,
        output=259,
        input_cache_read=52480,
        input_cache_creation=0
    ),
    message_id='chatcmpl-QQZErJzXx8vRJ8ycxikNVunk',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_ZS8r3zynxRNQF5TgJFKgRM7C',
    return_value=ToolReturnValue(
        is_error=True,
        output='Failed to parse file: /private/var/folders/lt/w3sl6z8j5v50vt2s6k
nvf5n00000gn/T/dev-issue-55-815001100/.maestro/flows/auth/login.yaml:22\nUnrecog
nized field "optional" (class maestro.orchestra.yaml.YamlFluentCommand), not mar
ked as ignorable\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=16)
ThinkPart(
    type='think',
    think='Maestro 1.39.13 不支持 `optional` 参数在 `assertVisible` 上？这很奇怪
。让我查看 Maestro 文档。\n\n实际上，`optional` 在 Maestro 中确实是支持的，但可
能是在较新版本中。Maestro 1.39.13 可能不支持 `optional` 字段在 `assertVisible` 
命令中。\n\n让我检查一下其他 flow 是否也有 `optional` 的使用。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_x6TcDJmtHZ5FNjBBBBvpVbHN',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && grep -rn \\"optional:\\" flows/ includes/
"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_gel3UD6tDr0DTVfKszyInYoo',
    function=FunctionBody(name='Shell', arguments='{"command": "maestro test'),
    extras=None
)
ToolResult(
    tool_call_id='tool_x6TcDJmtHZ5FNjBBBBvpVbHN',
    return_value=ToolReturnValue(
        is_error=False,
        output='flows/matching/swipe_recommendations.yaml:39:    optional: true\
nflows/matching/swipe_recommendations.yaml:44:    optional: true\nflows/matching
/swipe_recommendations.yaml:49:    optional: true\nflows/matching/swipe_recommen
dations.yaml:54:    optional: true\nflows/matching/swipe_recommendations.yaml:59
:    optional: true\nflows/matching/swipe_recommendations.yaml:64:    optional: 
true\nflows/matching/swipe_recommendations.yaml:69:    optional: true\nflows/mat
ching/swipe_recommendations.yaml:74:    optional: true\nflows/matching/view_matc
h_list.yaml:27:    optional: true\nflows/matching/view_match_list.yaml:30:    op
tional: true\nflows/matching/view_match_list.yaml:35:    optional: true\nflows/a
uth/login.yaml:22:  optional: true\nflows/auth/login.yaml:25:    optional: true\
n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' --help | grep -i option || echo \\"No option flag\\""}'
)
StatusUpdate(
    context_usage=0.202911376953125,
    context_tokens=53192,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=456,
        output=190,
        input_cache_read=52736,
        input_cache_creation=0
    ),
    message_id='chatcmpl-nNo12Mqc2ycaU6FBPwWYSZdG',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_gel3UD6tDr0DTVfKszyInYoo',
    return_value=ToolReturnValue(
        is_error=False,
        output="                           Optional YAML configuration file for 
the workspace. If not provided, Maestro will look for a config.yaml file in the 
workspace's\n",
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=17)
ThinkPart(
    type='think',
    think='Maestro 1.39.13 似乎不支持 `optional` 在 `assertVisible` 上。在 Maest
ro 中，`optional` 通常用于 `tapOn` 或 `assertVisible`，但可能需要不同的语法。\n\
n让我检查一下 Maestro 版本是否太旧。Maestro 1.39.13 应该是相当新的版本（2024年中
左右）。`optional` 在 Maestro 文档中是支持的，但可能只适用于某些命令类型，或者语
法不同。\n\n实际上，在 Maestro 中，`optional` 参数确实在较新版本中添加的。让我尝
试另一种方式：查看是否可以在 Maestro 中使用 `runScript` 或者将 `optional` 改为 `
enabled: false` 的条件执行。\n\n等等，我回忆一下 Maestro 的 `optional` 语法。实
际上在 Maestro 中，`optional` 参数是用于 `tapOn` 的，而 `assertVisible` 可能不支
持 `optional`。\n\n让我查看 Maestro 官方文档的语法。对于 `assertVisible`，如果元
素可能不存在，通常的做法是使用 `runFlow` 配合条件。\n\n实际上，经过查证，Maestro
的 `assertVisible` 从某个版本开始支持 `optional: true`。如果 1.39.13 不支持，那
可能需要升级 Maestro，或者修改 flow。\n\n但我的任务不是升级 Maestro，而是让测试
跑起来。所以我需要修改 flow 来移除不兼容的 `optional` 用法。\n\n对于 `assertVisi
ble` 的 `optional`，可以用 `runFlow` 的方式，或者干脆移除这些断言（因为它们是辅
助性的）。\n\n让我先尝试只运行 `login.yaml`，把 `optional` 去掉试试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_41ZFstxxnRDkYyvat2Euxtr0',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && cat > /tmp/test_optional.yaml << \'EOF\'\
\nappId: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- asse
rtVisible:\\n    text: \\"欢迎使用 Social World\\"\\n    optional: true\\nEOF\\n
maestro test /tmp/test_optional.yaml", "timeout": 60}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.20470428466796875,
    context_tokens=53662,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=670,
        output=478,
        input_cache_read=52992,
        input_cache_creation=0
    ),
    message_id='chatcmpl-zRw8IyDb4xEIJuSGPrMYy47v',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_41ZFstxxnRDkYyvat2Euxtr0',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_optional\nLaunch app "com
.example.mobile" with clear state... COMPLETED\nAssert that (Optional) "欢迎使用
Social World" is visible... COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=18)
ThinkPart(
    type='think',
    think='有意思，`optional` 在 `assertVisible` 的 `text` 形式下是工作的，但在 
`id` 形式下可能有问题？或者 `assertVisible` 后面直接跟字符串时不行？\n\n让我再看
看错误信息：\n`Unrecognized field "optional" (class maestro.orchestra.yaml.YamlF
luentCommand), not marked as ignorable`\n\n这发生在 `auth/login.yaml:22`，第22行
：\n```yaml\n- assertVisible: "手机号一键登录"\n```\n\n不对，第22行应该是：\n```
yaml\n- assertVisible: "完善资料"\n  optional: true\n```\n\n等等，让我重新数行号
：\n```\n14\tappId: com.example.mobile\n15\t---\n16\t- launchApp:\n17\t    clear
State: true\n18\t- runFlow: includes/login-as-test-user.yaml\n19\t\n20\t# 断言：
登录成功后离开登录页（新用户进入完善资料，老用户进入首页）\n21\t- assertNotVisib
le: "手机号一键登录"\n22\t- assertVisible: "完善资料"\n23\t  optional: true\n24\
t- assertVisible:\n25\t    id: "tab_home"\n26\t    optional: true\n```\n\n哦，我
明白了。YAML 的缩进问题！`- assertVisible: "完善资料"` 是一个 inline mapping，下
一行 `optional: true` 被解析为这个 mapping 的一部分。但实际上对于 inline mapping
，如果下一行缩进不同，YAML 解析器的行为可能因版本而异。\n\n等等，实际上这个 YAML
看起来是正确的：\n```yaml\n- assertVisible: "完善资料"\n  optional: true\n```\n\
n这在 YAML 中是合法的。但错误说 `Unrecognized field "optional"`。\n\n让我测试一
下这种精确的语法。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_qkvybbxXVYIdPfUd1FSEA3WM',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /tmp/test_optional2.yaml << \'EOF\'\\nappI
d: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- assertVisi
ble: \\"完善资料\\"\\n  optional: true\\n- assertVisible:\\n    id: \\"tab_home\
\"\\n    optional: true\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test /tmp/test_optional2.yaml", "
timeout": 60}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.206787109375,
    context_tokens=54208,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=704,
        output=499,
        input_cache_read=53504,
        input_cache_creation=0
    ),
    message_id='chatcmpl-aAEHgO4r4E4Ys5rGnjfNzEn8',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_qkvybbxXVYIdPfUd1FSEA3WM',
    return_value=ToolReturnValue(
        is_error=True,
        output='Failed to parse file: /tmp/test_optional2.yaml:6\nUnrecognized f
ield "optional" (class maestro.orchestra.yaml.YamlFluentCommand), not marked as 
ignorable\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=19)
ThinkPart(
    type='think',
    think='Aha! 问题在于：\n1. `assertVisible: "string"` 后面跟 `optional: true`
会被解析成同一个对象，而 Maestro 1.39.13 可能不支持 `optional` 在 `assertVisible
` 的这种语法上\n2. 但 `assertVisible: {text: "string", optional: true}` 或 `asse
rtVisible:\\n    text: ...\\n    optional: true` 可能可以工作\n\n等等，之前的测
试：\n```yaml\n- assertVisible:\n    text: "欢迎使用 Social World"\n    optional
: true\n```\n这是成功的！\n\n但：\n```yaml\n- assertVisible: "完善资料"\n  optio
nal: true\n```\n这是失败的！\n\n区别在于：当 `assertVisible` 的值是字符串时，Mae
stro 内部解析为 `YamlFluentCommand`（可能是一个简化的命令对象），这个类不支持 `o
ptional` 字段。但当值是一个 mapping（如 `text:` 和 `optional:`）时，它可能被解析
为另一个类。\n\n所以修复方案是：将所有 `assertVisible: "string"` 后面跟 `optiona
l: true` 改为 `assertVisible:\\n    text: "string"\\n    optional: true`。\n\n另
外 `assertNotVisible: "手机号一键登录"` 不支持 `optional`，这没问题。\n\n让我检
查一下 `matching/swipe_recommendations.yaml` 和 `view_match_list.yaml` 中的 `opt
ional` 用法。\n\n`swipe_recommendations.yaml`:\n```yaml\n- assertVisible: "附近"
  # 第20行，没有optional\n...\n- swipe:\n    direction: LEFT\n    optional: true
       # 第39行\n```\n\n`swipe` 命令的 `optional` 可能是支持的。\n\n`tapOn` 的 `
optional` 可能也支持。但 `assertVisible: "string"` 后面跟 `optional` 不行。\n\n
让我再测试一下 `tapOn` 的 `optional`。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_0USyXH7ApECYHZD3xUMHlLHY',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /tmp/test_optional3.yaml << \'EOF\'\\nappI
d: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- tapOn:\\n 
  id: \\"nonexistent\\"\\n    optional: true\\n- swipe:\\n    direction: LEFT\\n
   optional: true\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro && maestro test /tmp/test_optional3.yaml", "timeou
t": 60}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.20896530151367188,
    context_tokens=54779,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=763,
        output=549,
        input_cache_read=54016,
        input_cache_creation=0
    ),
    message_id='chatcmpl-NXzDlCDwLbViNPrDRA7y1q7J',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_0USyXH7ApECYHZD3xUMHlLHY',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_optional3\nLaunch app "co
m.example.mobile" with clear state... COMPLETED\nTap on (Optional) id: nonexiste
nt... WARNED\n\n\n Warning: Element not found: Id matching regex: nonexistent\nS
wiping in LEFT direction in 400 ms... COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=20)
ThinkPart(
    type='think',
    think='好的，`tapOn` 和 `swipe` 的 `optional` 都支持。问题只在 `assertVisibl
e: "string"` 这种简写形式后面跟 `optional: true`。\n\n修复方案：\n1. `auth/login
.yaml` 第22-26行：\n```yaml\n- assertVisible: "完善资料"\n  optional: true\n- as
sertVisible:\n    id: "tab_home"\n    optional: true\n```\n改为：\n```yaml\n- as
sertVisible:\n    text: "完善资料"\n    optional: true\n- assertVisible:\n    id
: "tab_home"\n    optional: true\n```\n\n2. `matching/view_match_list.yaml` 第25
-35行：\n```yaml\n- assertVisible:\n    text: "还没有匹配成功的好友"\n    option
al: true\n- assertVisible:\n    text: "继续滑动，发现感兴趣的人吧"\n    optional
: true\n```\n这里已经是 mapping 形式了，应该没问题。等等，让我再看看报错。不对，
报错只发生在 `auth/login.yaml`。让我重新运行测试看看是否还有其他文件报错。\n\n实
际上之前的测试只报了 `auth/login.yaml` 的错误，因为 Maestro 可能在解析阶段就退出
了。让我修复 `auth/login.yaml` 后再运行。\n\n但 `matching/view_match_list.yaml` 
中的 `assertVisible` 已经是 mapping 形式，应该没问题。`swipe_recommendations.yam
l` 中的 `assertVisible: "附近"` 没有 `optional`，也没问题。\n\n让我修复 `auth/lo
gin.yaml`，然后重新运行所有 flow。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_Kj9vm8FOpxIhMF9L79gQkGSo',
    function=FunctionBody(
        name='StrReplaceFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/flows/auth/login.yaml", "edit": {"old":"- assertVi
sible: \\"完善资料\\"\\n  optional: true\\n- assertVisible:\\n    id: \\"tab_hom
e\\"\\n    optional: true","new":"- assertVisible:\\n    text: \\"完善资料\\"\\n
   optional: true\\n- assertVisible:\\n    id: \\"tab_home\\"\\n    optional: tr
ue"}}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.2114105224609375,
    context_tokens=55420,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=892,
        output=457,
        input_cache_read=54528,
        input_cache_creation=0
    ),
    message_id='chatcmpl-R7TAFqlkhOyjodgqIQcSWU4D',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_Kj9vm8FOpxIhMF9L79gQkGSo',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/flows/auth/login.yaml',
                old_text='\n# 断言：登录成功后离开登录页（新用户进入完善资料，老
用户进入首页）\n- assertNotVisible: "手机号一键登录"\n- assertVisible: "完善资料
"\n  optional: true\n- assertVisible:\n    id: "tab_home"\n    optional: true',
                new_text='\n# 断言：登录成功后离开登录页（新用户进入完善资料，老
用户进入首页）\n- assertNotVisible: "手机号一键登录"\n- assertVisible:\n    text
: "完善资料"\n    optional: true\n- assertVisible:\n    id: "tab_home"\n    opti
onal: true',
                old_start=18,
                new_start=18,
                is_summary=False
            )
        ],
        extras=None
    )
)
StepBegin(n=21)
ThinkPart(
    type='think',
    think='现在让我重新运行所有 flow 测试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_mxpjf8Crmb9ByyMoMH07hHV6',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test --debug-output=\\"/Users/ins
ulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-issu
e-55-20260424-012302\\" --flatten-debug-output flows/auth/login.yaml flows/profi
le/setup.yaml flows/matching/swipe_recommendations.yaml flows/matching/view_matc
h_list.yaml flows/devtools/open_drawer.yaml flows/devtools/edit_server_url.yaml"
, "timeout": 300}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.21330642700195312,
    context_tokens=55917,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=621,
        output=159,
        input_cache_read=55296,
        input_cache_creation=0
    ),
    message_id='chatcmpl-HrETB09JNwV225b8fZlzKmto',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_mxpjf8Crmb9ByyMoMH07hHV6',
    return_value=ToolReturnValue(
        is_error=True,
        output='\nWaiting for flows to complete...\n[Failed] login (36s) (Assert
ion is false: "手机号一键登录" is not visible)\n[Failed] setup (35s) (Assertion 
is false: "手机号一键登录" is not visible)\n[Failed] swipe_recommendations (34s)
(Assertion is false: id: mode_tab_nearby is visible)\n[Failed] view_match_list (
34s) (Element not found: Id matching regex: tab_messages)\n[Passed] open_drawer 
(5s)\n[Passed] edit_server_url (37s)\n\n4/6 Flows Failed\n\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=22)
ThinkPart(
    type='think',
    think='测试运行了，4个失败，2个通过。让我分析每个失败的原因。\n\n**通过的：*
*\n- `open_drawer` - DevTools 抽屉可以正常打开\n- `edit_server_url` - 服务器 URL
编辑正常\n\n**失败的：**\n1. `login` - `Assertion is false: "手机号一键登录" is 
not visible`\n   - 这是因为 `login.yaml` 先运行了 `includes/login-as-test-user.y
aml`，然后断言 `assertNotVisible: "手机号一键登录"`。但 `assertNotVisible` 失败
了，说明登录页面上仍然可见 "手机号一键登录"。这意味着登录可能失败了，或者页面没
有跳转。\n   \n2. `setup` - 同样失败在 `assertNotVisible: "手机号一键登录"`，说
明登录没有成功。\n\n3. `swipe_recommendations` - `Assertion is false: id: mode_t
ab_nearby is visible`，说明首页没有加载。\n\n4. `view_match_list` - `Element not
found: Id matching regex: tab_messages`，说明没有进入首页。\n\n这些失败有一个共
同根因：**登录流程可能失败了**。让我查看 debug 输出和截图。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_UbqCmJ9d3LGN5Vo3nXcIWsS9',
    function=FunctionBody(
        name='SetTodoList',
        arguments='{"todos": [{"title":"阅读项目地图与E2E规范","status":"done"},
{"title":"扫描已有flow并分析覆盖范围","status":"done"},{"title":"检查mobile代码
中的UI元素和页面结构","status":"done"},{"title":"修复includes格式问题","status":
"done"},{"title":"执行Maestro测试","status":"done"},{"title":"分析失败根因并修复
flow","status":"in_progress"},{"title":"生成E2E测试报告","status":"pending"}]}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_rYWppdiUkIOTczQ8jsnjkql8',
    function=FunctionBody(name='Shell', arguments='{"command": "'),
    extras=None
)
ToolResult(
    tool_call_id='tool_UbqCmJ9d3LGN5Vo3nXcIWsS9',
    return_value=ToolReturnValue(
        is_error=False,
        output='Todo list updated',
        message='Todo list updated',
        display=[
            TodoDisplayBlock(
                type='todo',
                items=[
                    TodoDisplayItem(
                        title='阅读项目地图与E2E规范',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='扫描已有flow并分析覆盖范围',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='检查mobile代码中的UI元素和页面结构',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='修复includes格式问题',
                        status='done'
                    ),
                    TodoDisplayItem(title='执行Maestro测试', status='done'),
                    TodoDisplayItem(
                        title='分析失败根因并修复flow',
                        status='in_progress'
                    ),
                    TodoDisplayItem(title='生成E2E测试报告', status='pending')
                ]
            )
        ],
        extras=None
    )
)
ToolCallPart(
    arguments_part='ls -la /Users/insulate/Desktop/new-world/new-world-ops/logs/
social-world/e2e-screenshots-e2e-issue-55-20260424-012302/"}'
)
StatusUpdate(
    context_usage=0.21451187133789062,
    context_tokens=56233,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=425,
        output=405,
        input_cache_read=55808,
        input_cache_creation=0
    ),
    message_id='chatcmpl-KHWnnuhyZxjNNIYcrthu9pj1',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_rYWppdiUkIOTczQ8jsnjkql8',
    return_value=ToolReturnValue(
        is_error=False,
        output='total 1472\ndrwxr-xr-x@  25 insulate  staff    800 Apr 24 01:29 
.\ndrwxr-xr-x@ 383 insulate  staff  12256 Apr 24 01:25 ..\n-rw-r--r--@   1 insul
ate  staff    186 Apr 24 01:29 ai-(edit_server_url).json\n-rw-r--r--@   1 insula
te  staff    162 Apr 24 01:29 ai-(login).json\n-rw-r--r--@   1 insulate  staff  
 178 Apr 24 01:29 ai-(open_drawer).json\n-rw-r--r--@   1 insulate  staff    165 
Apr 24 01:29 ai-(setup).json\n-rw-r--r--@   1 insulate  staff    198 Apr 24 01:2
9 ai-(swipe_recommendations).json\n-rw-r--r--@   1 insulate  staff    186 Apr 24
01:29 ai-(view_match_list).json\n-rw-r--r--@   1 insulate  staff   3034 Apr 24 0
1:29 ai-report-edit_server_url.html\n-rw-r--r--@   1 insulate  staff   3000 Apr 
24 01:29 ai-report-login.html\n-rw-r--r--@   1 insulate  staff   3022 Apr 24 01:
29 ai-report-open_drawer.html\n-rw-r--r--@   1 insulate  staff   3003 Apr 24 01:
29 ai-report-setup.html\n-rw-r--r--@   1 insulate  staff   3052 Apr 24 01:29 ai-
report-swipe_recommendations.html\n-rw-r--r--@   1 insulate  staff   3034 Apr 24
01:29 ai-report-view_match_list.html\n-rw-r--r--@   1 insulate  staff   5672 Apr
24 01:29 commands-(edit_server_url).json\n-rw-r--r--@   1 insulate  staff  61554
Apr 24 01:27 commands-(login).json\n-rw-r--r--@   1 insulate  staff   1808 Apr 2
4 01:29 commands-(open_drawer).json\n-rw-r--r--@   1 insulate  staff  61552 Apr 
24 01:28 commands-(setup).json\n-rw-r--r--@   1 insulate  staff  61540 Apr 24 01
:28 commands-(swipe_recommendations).json\n-rw-r--r--@   1 insulate  staff  6204
1 Apr 24 01:29 commands-(view_match_list).json\n-rw-r--r--@   1 insulate  staff 
51287 Apr 24 01:29 maestro.log\n-rw-r--r--@   1 insulate  staff  92039 Apr 24 01
:27 screenshot-❌-1776965248733-(login).png\n-rw-r--r--@   1 insulate  staff  92
435 Apr 24 01:28 screenshot-❌-1776965283601-(setup).png\n-rw-r--r--@   1 insula
te  staff  92281 Apr 24 01:28 screenshot-❌-1776965317641-(swipe_recommendations
).png\n-rw-r--r--@   1 insulate  staff  92233 Apr 24 01:29 screenshot-❌-1776965
351172-(view_match_list).png\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=23)
ThinkPart(
    type='think',
    think='让我查看 login flow 的 commands JSON 和失败截图，了解登录失败的具体原
因。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_IXXpvAiFUYLcs1tKiiOm46LP',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/Users/insulate/Desktop/new-world/new-world-ops/log
s/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/commands-(login).jso
n", "n_lines": 100}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_5Sx8PRusLminSbhJ9zHphQit',
    function=FunctionBody(name='ReadMediaFile', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_IXXpvAiFUYLcs1tKiiOm46LP',
    return_value=ToolOk(
        is_error=False,
        output='     1\t[ {\n     2\t  "command" : {\n     3\t    "defineVariabl
esCommand" : {\n     4\t      "env" : {\n     5\t        "MAESTRO_FILENAME" : "l
ogin"\n     6\t      },\n     7\t      "optional" : false\n     8\t    }\n     9
\t  },\n    10\t  "metadata" : {\n    11\t    "status" : "COMPLETED",\n    12\t 
  "timestamp" : 1776965213286,\n    13\t    "duration" : 4\n    14\t  }\n    15\
t}, {\n    16\t  "command" : {\n    17\t    "runFlowCommand" : {\n    18\t      
"commands" : [ {\n    19\t        "applyConfigurationCommand" : {\n    20\t     
    "config" : {\n    21\t            "appId" : "com.example.mobile"\n    22\t  
       },\n    23\t          "optional" : false\n    24\t        }\n    25\t    
 }, {\n    26\t        "assertConditionCommand" : {\n    27\t          "conditio
n" : {\n    28\t            "visible" : {\n    29\t              "textRegex" : "
欢迎使用 Social World",\n    30\t              "optional" : false\n    31\t     
      }\n    32\t          },\n    33\t          "optional" : false\n    34\t   
    }\n    35\t      }, {\n    36\t        "assertConditionCommand" : {\n    37\
t          "condition" : {\n    38\t            "visible" : {\n    39\t         
    "textRegex" : "手机号一键登录",\n    40\t              "optional" : false\n 
  41\t            }\n    42\t          },\n    43\t          "optional" : false\
n    44\t        }\n    45\t      }, {\n    46\t        "tapOnElement" : {\n    
47\t          "selector" : {\n    48\t            "textRegex" : "手机号",\n    4
9\t            "optional" : false\n    50\t          },\n    51\t          "retr
yIfNoChange" : true,\n    52\t          "waitUntilVisible" : false,\n    53\t   
      "longPress" : false,\n    54\t          "optional" : false\n    55\t      
 }\n    56\t      }, {\n    57\t        "inputTextCommand" : {\n    58\t        
 "text" : "13800138000",\n    59\t          "optional" : false\n    60\t        
}\n    61\t      }, {\n    62\t        "tapOnElement" : {\n    63\t          "se
lector" : {\n    64\t            "textRegex" : "获取验证码",\n    65\t          
 "optional" : false\n    66\t          },\n    67\t          "retryIfNoChange" :
true,\n    68\t          "waitUntilVisible" : false,\n    69\t          "longPre
ss" : false,\n    70\t          "optional" : false\n    71\t        }\n    72\t 
    }, {\n    73\t        "tapOnElement" : {\n    74\t          "selector" : {\n
   75\t            "textRegex" : "验证码",\n    76\t            "optional" : fal
se\n    77\t          },\n    78\t          "retryIfNoChange" : true,\n    79\t 
        "waitUntilVisible" : false,\n    80\t          "longPress" : false,\n   
81\t          "optional" : false\n    82\t        }\n    83\t      }, {\n    84\
t        "inputTextCommand" : {\n    85\t          "text" : "123456",\n    86\t 
        "optional" : false\n    87\t        }\n    88\t      }, {\n    89\t     
  "hideKeyboardCommand" : {\n    90\t          "optional" : false\n    91\t     
  }\n    92\t      }, {\n    93\t        "tapOnElement" : {\n    94\t          "
selector" : {\n    95\t            "textRegex" : "登录",\n    96\t            "o
ptional" : false\n    97\t          },\n    98\t          "retryIfNoChange" : tr
ue,\n    99\t          "waitUntilVisible" : false,\n   100\t          "longPress
" : false,\n',
        message='100 lines read from file starting from line 1. Total lines in f
ile: 1538.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"path": "/Users/insulate/Desktop/new-world/new-world-ops/lo
gs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌-17769
65248733-(login).png"}'
)
StatusUpdate(
    context_usage=0.2194061279296875,
    context_tokens=57516,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1452,
        output=145,
        input_cache_read=56064,
        input_cache_creation=0
    ),
    message_id='chatcmpl-SgpB1OBTNWXssD26HcI06q6w',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_5Sx8PRusLminSbhJ9zHphQit',
    return_value=ToolOk(
        is_error=False,
        output=[
            TextPart(
                type='text',
                text='<image path="/Users/insulate/Desktop/new-world/new-world-o
ps/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌-
1776965248733-(login).png">'
            ),
            ImageURLPart(
                type='image_url',
                image_url=ImageURL(
                    url='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABDgAAAlgC
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
gvwffWxM4lSKRtoiohv+tt4skq+BJUzmf3bb4bMoqLiyqrKgoKCkpJi99m+vv6hoaGO9o6Bgf54LG6vw
zCNowUAADidpH6wSi5plfhqHybWvrIThh0xEmtkWcnNSJxooi09iqyiZcGGsy+76Sb75uKbPCYiIkXV1
QUVKUvjdnZ2Hj58ePfPft78cKdeAAAgAElEQVSy7QVxwoE/q/Q3N790111KUpYcMxcvDmXIKoapDHefG
GVMe1YhqOQRpZRpmiLiXWbBLp7E43HyCQAAwGQYf1bRWsfjlptV3LqKU1zJoa5ib1uS/Mf+VkRElmze3
HDZZcEjj8X2/PJXgVklPjy87+c/b9+9J1JcHBsYyC2rGPYDKrH4coasYtjhJLFNjEhAw8IEZBXzPe++c
Syvw+Rwt+Fx2Y9M97gAAABOW+kfo70lAn+1QJRSdk95ysnJpxK9Gu4aWSLibEkfDhv1dfGWkyoaVSJaE
qer1NPsf0QkKKXE+vv7WloKystFRBlG9YoVg51dfU1N9s0GTp3qa26qOOusfb/8Vfuu3XVnr191ww2D7
R0DJ1sTJRDRnZ06GjXqaj1vNXF753stWqX+EPw/Bc+Pzt9TPVFZhaACAACAM51v33rJLaukf0b3ZpVEr
/losopKnqhELU5LKdH+/p3f//7h3/+hdM7s4ro6CcgqauDUqeZnt/afPFm74eylb32rWVBQuXz5QNupg
ZOt4g4nQ1ZJVEeUKC2JuDRdWYWgAgAAAExMVlFK6UT4sF80urqKdgODWnjV6+Zfcbl3hHZK6Tp4SLRu3
bkzS1bRVrxuw9nL3/Y2IxQSEWWa3qzi3CQoqyQqOe62GNNZVyGoAAAAACLjyCrp3eTjqKs4n/p1LF61c
qUZSa4Ku//X97e9+qpTFbFSsopY1qndu3tPNNlDqN+wcfm1TkpxbuvJKirXrCLOLLZpqqsQVAAAAADH2
LJKln6V0ddVnHQw3NXVfeRo/XnnuTcrX7ig+8jhoa5uZ3SJrFJUVbXnl79qfuF5++qzN25Yfu213pTi3
NY0K5cv729r82UViUZV5rqKPZqsdRX3B5V8eEKyCkEFAAAASJqMfpWxzQEb6uoa6uioXr3afpEZiVSvW
pXMKiL2HLCuo0faduy0L1u/ceOK1JSiLWu4qytUWCgiyjRnLVvW29g41NHh7a0fOatMR289QQUAAABIk
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
TFLckqKeIRDCPHp2pX/4c6duxkZGVQXNZFYA5jy0fzz588R9++Lb2/Xpi3z2X55eXnfJ0+ZWtm9qEyhp
06FnjpVzpMkJCR07NxFLv0hiiz4WFBQwDCphhDi6ODg6ODA/oQd2revXbt2YmIiXYN7ERHFxSX8JS4Ab
FjUr6+jrc1QY/Hd+/cMgUpcfDxlXjVKz549Z55LHMlYbsXeDvP3ADgHgQoQQkg1Q8NVy5cHDBggvutLZ
ubU6TMO/LNPVVWV8tg3b6hXOvLfThUXl2wM3kzZYMigwRLn+ZRRjVfQTcpv5fZbZ5YsKiqaMXMmm6EAz
lZUxCp8qeTm5jI3sLe3l2oNlZqaWnd/v42bqP+HJYR8+/YtJSW5Tp067M8pUFZWlpb2KTUt9fPnz/n5+
TweT11dXVNT09jY2NTERNHFK4qLi5OSklPTUrOysvLzC0pLSzQ1tXR0tGuamtaqZcZmJV4Fy8zMSkxMz
MnNyc3NzcnNraJURVNTw8DAwNTEtFatmnS/jTlITU2tZcuW4Zcv0zWIjY1lKCofzWIRi8Dtu3cY9n7Jz
Ix/x7RGRTj7i1Qq994WVlpaGh8f/+FDwrfs7CpVqujoaNcwrtG4sXS/B2RWUFDw4UNCUnJybm5OUVGRh
oZmDWPjevXqVq9evQKuDr8qBCrwL3f3VqNHjfx723bxXfciIoK3bJ38xyTKA+n+kPBXl4aeCqWr1eXTV
caX1tdv3BTf2NDWtkGDBrKd8NewaMkShtpqwmTIjFwxfslcSYqb+lVSUsLcQIbnb3Nzc+YG2axfb/MVF
RU9evz48uUrV69de/f+PV0za2srby+vdm3aOjg4yHHEJjcv79atW+GXL1+7fiMrK4uuWbu2bbxat27fr
p1Uz5QLgoJ279lLt3fokMFBCxZI1dvi4uKIiPu37tyOuH+focShgYGBV2tPby+v1p6eUuWYrizNHB0ZA
hXm6VhPnjxlf6GYmDdpaWl0RZzeMUYppqYmbCq6CKuYe9vKtmF+fj7lrq5dOv+1ZQv/M4/HO3327PYdO
8STpMXGvFZomsfS0tIHDx+Gnjp14eIlyhmw7u6t+vbu3aF9e3V1dcV1A35VCFTgP1P++OPBg4dPnz0T3
7Vuw4bmzZ3dqIphXb1GvZzXxsY6MzNr2YqVlHu7dO7E5q+C+IS079+/U2Y67u7vT3eSjIyM8MtXJF5LR
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
CznpdXV1devXUuXppzO6FEjlywKKuelCSHjx46dO3uOfCd9CevYocP8uXNkO9bRwYH5XyTmTUUncJeBi
YmJJWOxeb4WLlKkw7JrxGqZisylFSv33uaUnj26y/xSYPmypSMZE/YACMMald9dRsaXtLQ0hgZGRkb/7
Nldt67kmt/iv7Y2bNy0dv16hkOmTZkyIDBQ4pknTRh/9tw5un6GX77cu9/Hv7ZsEVnlL27u/AX7/vlH4
uUYxMS8mfjHZEJIkyaNW7q6OjRtamVlVdvMTO5rMzIzs3Jzc7JzcnJycr59+5adnf3169ev37KzMjPzC
/LV1dWXLFok88xm5gfxpKSknbt3lz+v2ps3b5csW7Zk2TK7hg07d+rk5NSscePG7FP9iDh0+DDdBPF2b
duUo5u/DgMDgyWLgrr5+mzaHCztSgxKAwIDJ04Yb2IiXaJqQkg3X19LS8v5Cxc+fBgp26WbOTouDloob
f1yQoiSktKggQOtra3nzp/PnMuLjrGx8eKghZ07MVXPkIsRw4fr6+vPnjuPfT4DQkhrT8+1q1fFv3v39
7btdG1u3b6Tl/ed/UBEZXFt0ULiiwyppt7VNDW1sbGWmFeK5WAOpUq8t7lmzOhRBgZV585fwP4GVlFRW
bZkSb++fXai5iOwhkDldxcY0E9XV4dhdcr+fXsbNGgg7Wm/ffsWtHjJseNMVa6HDhk8aSLTlDABfX39r
Zs3d+/Vi65BTMwbrzZtD+7/x8Odeo4yj8ebM2/+4SNMhdKWL1vasUOH1JSUbTt2nj4joUKccBolQkhDW
1tbW5u65nVNTU00NDTU1dU1NDT4H1RVVU+dPMHj8QoKCng8XiEfj/c5IyN469b8/PzCwsL8/IL8/Pxv3
75lZGSkpKYyh46EkGlTpihi/eWjx48PHzkqsWKXtKJfvRKUHWzXrq2nu7ujo6M96/SghJD09PTVa0TTy
vFVMzRs4+0th17+KpydnP7Zu+fR48cnT4YeO3FCqodgPl1d3Z49uvt161aeRGp2DRseOnDg3PkLwVu30
hV4pWRRv/64cWO7+fiUJ/hv4eJy6sSJI0ePBW/dylxoT5iuru6QwYOGDRlqaFhByTN69+rV2L7xqjWr2
ZTA09TUnDBu7KiRI9XU1DQ0NHS0telC98LCwhdRL7hfmqNJY/v9ktpIO/rh6eHBHKg0tLVlKGjLRuXe2
5zSt0+fRo0aLV22/M7duxIbN25sv2jhQn5+SJMaUr/+gN8WAhUgvj4+qqqq4un569ert2Pb3zJU8Yt6+
XLS5CnMv8H9unWbO5u2poo4Z2enVStWzGAcaw4cMHDY0CGjRowUKVZYUFDw56xZJ0NPMRx7/OgRl+bNC
SFG1aoFb9ro3spt2ow/2Xfv1evXElPEylGP7v5yPNuHDwlPnj49eCiE5TvC+vXqrVq54vPnjLXr10s7s
+vy5Sv8HLLVDA29vbzMzc1r1qppUsPEpIZxjRo19PT0KI9at2EDXSKjPyZN0tHRkaoPEgkvtF2zbp2er
q6OjuSpZc9fPKfcrq1V0S+2lZSUnJ2cnJ2cZs3882Hko8hHkXfv3WOoeMjXzNHR0cHByamZe6tWcplKp
6am1t3fz6drl4eRkVeuXL12/TpD2Xtzc3Nvr9bt2rR1cWkul8c4HR2dYUOH9OnT+9atW/yrf8nMpGypr
q7erm0bby+vdm3bVq1atfyXloqNjfWuHTuePnt2+syZCxcvUb6ksLS07Nqlc7++fWuamvK36Onpde/e/
Z/9tM/5p06f5n6gIjEIaeXmRvc7gY7E6NrVVQ6qcy98AAAgAElEQVQ/lsq9tzmlkZ3dgX/23Yu4Hxoae
vHSJcrg2d29Vd/evTu0by9Ygl9D1nSd8BtSKitSYOpM+Ins3fePcF0qTw+PjRvWy5Ca5mDIoZmSIhBvL
69tf22VIZ/s/IUL9+zdJ7HZnzOm9+vbl995Ho83dfoM5lyKF86eER+Ij4uPD1q0WC5TaOSrQ/v2O7dvY
9MyL++7DVVtgR7d/ZcsXhwZGRlx/35YWDjDn1hx3l5e69eu5b9yLiwsDD11avXadenp8lm8a2xsbGtjU
6dOnRrGxj26+5uZmRFCrt+4MXDwEMr2RkZGt65fk/mpOv7du9beFNPGbGysL1+6xP+8/8ABGRbqCIt7E
8OFDDl5ed9TUpLTP3/OyvpaUFBQUlKspaWlra2tra2tq6tbp04dmWflsVRWVpaW9ik1LfXz58/fv3/n8
XiqqmpaWprVq1c3NTU1NTFR6Kqt4uKSpKSktE9pmZmZ+fkFJSXFGhoaurq6ZrVqmZmZlX/VgVwUFxcnf
PyYmJj47du3wsJCHR0dPT09U1PT+vXqVeKSNpCocu9t7sjPz3//4UNyUnJObm5xcZGmplYNY+P69euVc
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
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUA
AAAACgOQQUAAACA4hBUAAAAACgOQQUAAACA4hBUAAAAACiOxtEduEVYLBaVSlWRmiaTKT09QwgREOCvV
quruV8AAABAjURQuX4mk+nixfjTp88cOBB95szZjz+eHhwcVO6r0tMzRox4RgixZMn8kJDg6u8mAAAAU
PMQVCrKYDCkpqYlJV1KSEiMj0+Ii4u7cCGusLBQrvDhhzM+/vgjNze3634LnU536VKynQrh4eFabTntz
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
mzFmjVrt2/feejQf4mJ5cxEqlnkGVbC6ja9LJs2b5YKXl5ekZH1Sjzr7+cnD2hs3bbN/oIQ6wX0No8H6
dyxk1zeurXMWXxGo+nff9fJ7URF2hgR6te3r1RISEg4fKTMm+aLFy/Kx1neffdd5R4GfyOsv++PPnTIT
s20tDT599K3T28bTVntEx0dba8pi8UinYcjCQsrOfoUGBjYpXNnqbxqzRpdXp4ow549V/9UunXtUrpCF
TbVpnUbufz32rVltSOu/Ttp19ZG5gEAADcBQcWGgAD/Dz+cVPoQlZUrV3/44Yz8fBvziyRSSlm7dl2J6
yaT6fjxkytXrl669Od58+ZPn/7xO+9MeuaZF6ZM+ajc/aNqilYtW8rpYtkvyy9evFhWzd179uzfX3wff
9+9A0sf6KFSqQYOLF7MnZKSsnLVqrKaMhgM3373vfywS5fOpeu0b98+KKg4ds5fuMB6BMbauvXr5G/ZH
xs82OYMMTmoCCHmL1hY1oLs735YKpfvufvusjpfJbp36yqXFy1eYifU/frb7/IfW79+/UpX6Njh6uTGH
5YutbOKY/OWLUeOFI9d3DtwgKenjel5Dz5QvIVxdnb2L8uX22zHYDB8s2C+VNZqtb179bJZraqa8vf3k
zdW/nfdugMHD9ps6uLFi0t/+kkq9+7Vy/4kRgAAUH0IKraFhYXOmPFhQEDJDX/27t0/ceLb8qoAa2WlF
Dv27Nm3cePmG+qoYmg0GvmkjszMzFdefc1mVjl56tTEt96SHw57uuTuBZLBgx6Vy+99MEkegbGWn58/d
dq0HTt3Sg979erZpnXr0tU8PNxHjxwplS9dSn7nvfdLL2U5cfLk2++8K5VdXV0HDhggbGndulWf3sVjE
Wv/+eeb+QtKZ5X/++PPhYsWSeWWLVpY3/1Xh0aNGskbcO3es+eDyZN1upKr5IUQf69d+9GMGVI5MjJyw
D33lK4THh4+eNAgqXzi5Mn/vfOuzVB34ODBN968+hsc+uRQmx27d+C98qS+GTNn7dy1q0SFoqKi6TNmH
j58RHr41JNPlp4EWOVNPTX0Sbk88a3/lV70kpGR+frEN+VVN08OGWKzHQAAcBOoJ73/gQPfPiUlRQih1
WrtLGx1FF9fnx49uh08GJ2Tc81Xy5mZWWvXrsvPL2jatImLS/GK5OtIKZLg4KCOHZW4TF+v18/75hup3
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
+G6lu5AUAEAAAAnUHah6WxOTeHZupLCJqVCx3U5ADzl7eMem+CfMDh4cEYo/4fLdw5BBQAAAJxMc5Omr
kbV3KhRKfQ6nZFdXRtgYHJxEbi7i7183HwDPIJDvXz9PbiuqMcgqAAAAAAAAO9g8BkAAAAAAPAOggoAA
AAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOg
goAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAA
PAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAA
AAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOg
goAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAA
PAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAA
AAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOg
goAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAA
PAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAAAAAAPAOggoAA
AAAAPAOggoAAAAAAPAOggoAAAAAAPCOiOsCAAAAAAD6OZPJpNfrDQaD0WQ0mUxms9litlgYhoghErgIB
C5CF6FQKBKJxCKxq6urm5ubSDTQ/1Af6D8/AAAAAEDP0mg0SpVKrVZpNBqNRqvTaU0mk6M3EYlE7u4SD
w+Jh4eHp6eXt5eXh4dHb1TLWwLGaOG6BgAAAAAAJ8YwjKy5Wd7cLFfIFQpFN2KJPUQikY+Pj9RHKvX19
fP1FQgEvfEU/kBQAQAAAADoDq1W29DY0NTU1NQkY5g+/aNaIHDx9/fz9/cPDAiUSCR9+eg+g6ACAAAAA
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
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCE
SoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAA
OEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUA
ACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAA4P+3d3/HTVwLHMcNc9+zdKBJA6zdgNYVR
IF3LDeARAXgBrCcArBEAYxpwFryHhANBKWBSG7Avg+aYbjcBGywdX7Cn88jzHDOPhyx3/1zNo5QAQAA4
ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAI
I5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAA
OIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAA
CCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABA
ADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVA
AAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QA
QAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIF
QAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOU
AEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiC
BUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgj
lABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4
ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAI
I5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAA
OIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAA
CCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABA
ADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVA
AAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QA
QAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIF
QAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOU
AEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiCBUAACCOUAEAAOIIFQAAII5QAQAA4ggVAAAgjlABAADiC
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',
                    id=None
                )
            ),
            TextPart(type='text', text='</image>')
        ],
        message='Loaded image file `/Users/insulate/Desktop/new-world/new-world-
ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌
-1776965248733-(login).png` (image/png, 92039 bytes, original size 1080x2400px).
If you need to output coordinates, output relative coordinates first and compute
absolute coordinates using the original image size; if you generate or edit imag
es/videos via commands or scripts, read the result back immediately before conti
nuing.',
        display=[],
        extras=None
    )
)
StepBegin(n=24)
ThinkPart(
    type='think',
    think='从截图可以看到，登录页面仍然显示，手机号和验证码都已输入，但登录没有
成功。原因是：\n1. 点击了"登录"按钮后，页面没有跳转\n2. 后端可能返回了错误\n\n从
login_page.dart 的代码看，登录调用 `AuthBloc.add(LoginRequested(...))`。这可能涉
及到后端 API 调用。\n\n等等，实际上这是一个关键观察：截图显示登录按钮上没有任何
加载状态，而且倒计时还在 "35 s"。这说明 `login-as-test-user.yaml` 在点击登录后，
可能登录请求正在发送但失败了，或者网络请求超时。\n\n但更可能的问题是：**后端 API
在 3001 端口，而 App 配置的服务器地址可能不对**。从 `edit_server_url.yaml` 的恢
复地址看，E2E 服务器地址应该是 `http://10.0.2.2:3001/api/v1`（模拟器访问宿主机的
地址）。\n\n但 `config.yaml` 中 `API_BASE_URL` 是 `http://localhost:3000`。如果 
App 默认使用 `localhost:3000`，那在模拟器里是无法访问的（模拟器里的 localhost 是
它自己，不是宿主机）。\n\n而且从之前的 `curl` 测试看，后端实际上在 `localhost:30
01` 运行。\n\n所以问题可能是：\n1. App 默认配置的服务器地址不正确\n2. 登录请求发
送到了错误的地址，导致登录失败\n\n等等，但 `open_drawer` 和 `edit_server_url` 测
试通过了。`edit_server_url` 可以编辑服务器地址，说明 DevTools 抽屉是可以打开的。
但这不依赖后端。\n\n让我看看 App 是否默认配置了正确的服务器地址。我需要查看 DioC
lient 或相关配置。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_BrIO9ilUV8J8K7QyMfvxzIfZ',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100 && grep -rn \\"localhost\\\\|10.0.2.2\\\\|3000\\\\|30
01\\" apps/mobile/lib --include=\\"*.dart\\" | grep -v generated | head -20"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_bxluNdgirpruEw13zDayG4bU',
    function=FunctionBody(name='ReadFile', arguments='{"path":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_BrIO9ilUV8J8K7QyMfvxzIfZ',
    return_value=ToolReturnValue(
        is_error=False,
        output="apps/mobile/lib/core/network/dio_client.dart:14:              de
faultValue: 'http://10.0.2.2:3001/api/v1',\napps/mobile/lib/data/repositories/de
v_tools_repository_impl.dart:15:    defaultValue: 'http://10.0.2.2:3001/api/v1',
\napps/mobile/lib/presentation/widgets/server_url_tool.dart:88:                 
hintText: 'http://localhost:3000/api/v1',\n",
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/apps/mobile/lib/core/network/dio_client.dart"}'
)
StatusUpdate(
    context_usage=0.23712158203125,
    context_tokens=62160,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=4816,
        output=515,
        input_cache_read=57344,
        input_cache_creation=0
    ),
    message_id='chatcmpl-TsQD3xlwNSjFGoJw3GCTdVuL',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_bxluNdgirpruEw13zDayG4bU',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:dio/dio.dart';\n     2\timport 'auth_int
erceptor.dart';\n     3\timport '../storage/secure_storage.dart';\n     4\t\n   
 5\tclass DioClient {\n     6\t  late final Dio dio;\n     7\t\n     8\t  DioCli
ent({String? baseUrl, SecureStorage? secureStorage}) {\n     9\t    dio = Dio(\n
   10\t      BaseOptions(\n    11\t        baseUrl: baseUrl ??\n    12\t        
   const String.fromEnvironment(\n    13\t              'API_BASE_URL',\n    14\
t              defaultValue: 'http://10.0.2.2:3001/api/v1',\n    15\t           
),\n    16\t        connectTimeout: const Duration(seconds: 10),\n    17\t      
 receiveTimeout: const Duration(seconds: 10),\n    18\t        contentType: 'app
lication/json',\n    19\t      ),\n    20\t    );\n    21\t    dio.interceptors.
add(LogInterceptor(requestBody: true, responseBody: true));\n    22\t    dio.int
erceptors.add(\n    23\t      AuthInterceptor(\n    24\t        dio: dio,\n    2
5\t        secureStorage: secureStorage ?? SecureStorage(),\n    26\t      ),\n 
  27\t    );\n    28\t  }\n    29\t\n    30\t  void updateBaseUrl(String url) {\
n    31\t    dio.options.baseUrl = url;\n    32\t  }\n    33\t}\n",
        message='33 lines read from file starting from line 1. Total lines in fi
le: 33. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=25)
ThinkPart(
    type='think',
    think='DioClient 的默认 baseUrl 是 `http://10.0.2.2:3001/api/v1`，这是正确的
（模拟器访问宿主机 3001 端口）。\n\n那登录为什么失败？让我查看后端 API 的登录接
口实现，以及截图中是否有错误提示。\n\n从截图看，登录页面上没有显示任何 SnackBar 
错误消息。这说明：\n1. 登录请求可能还没完成（网络超时）\n2. 或者登录请求成功了但
跳转逻辑有问题\n\n等等，从 commands JSON 来看，截图是在 `assertNotVisible: "手机
号一键登录"` 失败时截取的。这个断言是在点击登录按钮之后执行的。让我看看 commands
JSON 中点击登录后到断言之间的时间间隔。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_BGvac08GLY8uGaR2KPA6pR2D',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/Users/insulate/Desktop/new-world/new-world-ops/log
s/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/commands-(login).jso
n", "line_offset": 90, "n_lines": 60}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.24088668823242188,
    context_tokens=63147,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=995,
        output=225,
        input_cache_read=62152,
        input_cache_creation=0
    ),
    message_id='chatcmpl-E3L4mqpjfF9IdrvGpEFeMDly',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_BGvac08GLY8uGaR2KPA6pR2D',
    return_value=ToolOk(
        is_error=False,
        output='    90\t          "optional" : false\n    91\t        }\n    92\
t      }, {\n    93\t        "tapOnElement" : {\n    94\t          "selector" : 
{\n    95\t            "textRegex" : "登录",\n    96\t            "optional" : f
alse\n    97\t          },\n    98\t          "retryIfNoChange" : true,\n    99\
t          "waitUntilVisible" : false,\n   100\t          "longPress" : false,\n
  101\t          "optional" : false\n   102\t        }\n   103\t      } ],\n   1
04\t      "sourceDescription" : "includes/login-as-test-user.yaml",\n   105\t   
  "config" : {\n   106\t        "appId" : "com.example.mobile"\n   107\t      },
\n   108\t      "optional" : false\n   109\t    }\n   110\t  },\n   111\t  "meta
data" : {\n   112\t    "status" : "COMPLETED",\n   113\t    "timestamp" : 177696
5214127,\n   114\t    "duration" : 17453\n   115\t  }\n   116\t}, {\n   117\t  "
command" : {\n   118\t    "tapOnElement" : {\n   119\t      "selector" : {\n   1
20\t        "textRegex" : "验证码",\n   121\t        "optional" : false\n   122\
t      },\n   123\t      "retryIfNoChange" : true,\n   124\t      "waitUntilVisi
ble" : false,\n   125\t      "longPress" : false,\n   126\t      "optional" : fa
lse\n   127\t    }\n   128\t  },\n   129\t  "metadata" : {\n   130\t    "status"
: "COMPLETED",\n   131\t    "timestamp" : 1776965223426,\n   132\t    "duration"
: 3060\n   133\t  }\n   134\t}, {\n   135\t  "command" : {\n   136\t    "inputTe
xtCommand" : {\n   137\t      "text" : "123456",\n   138\t      "optional" : fal
se\n   139\t    }\n   140\t  },\n   141\t  "metadata" : {\n   142\t    "status" 
: "COMPLETED",\n   143\t    "timestamp" : 1776965226486,\n   144\t    "duration"
: 1788\n   145\t  }\n   146\t}, {\n   147\t  "command" : {\n   148\t    "assertC
onditionCommand" : {\n   149\t      "condition" : {\n',
        message='60 lines read from file starting from line 90. Total lines in f
ile: 1538.',
        display=[],
        extras=None
    )
)
StepBegin(n=26)
ThinkPart(
    type='think',
    think='让我继续查看 login flow 的 commands JSON，找到 `assertNotVisible` 失
败的具体步骤和时间戳。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_90C6ePz5xPTG8EebTLDPzRxi',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /Users/insulate/Desktop/new-world/new-world-o
ps/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302 && python3 -c 
\\"\\nimport json\\nwith open(\'commands-(login).json\') as f:\\n    data = json
.load(f)\\nfor i, item in enumerate(data):\\n    meta = item[\'metadata\']\\n   
cmd = json.dumps(item[\'command\'], ensure_ascii=False)[:200]\\n    print(f\'{i:
3d} | {meta[\\\\\\"status\\\\\\"]:12s} | {meta[\\\\\\"duration\\\\\\"]:6d}ms | {
cmd}\')\\n\\" | tail -20"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.24404144287109375,
    context_tokens=63974,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=998,
        output=185,
        input_cache_read=62976,
        input_cache_creation=0
    ),
    message_id='chatcmpl-kPtGdjQoR6p4hzpN37i1Suax',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_90C6ePz5xPTG8EebTLDPzRxi',
    return_value=ToolReturnValue(
        is_error=False,
        output='  0 | COMPLETED    |      4ms | {"defineVariablesCommand": {"env
": {"MAESTRO_FILENAME": "login"}, "optional": false}}\n  1 | COMPLETED    |  174
53ms | {"runFlowCommand": {"commands": [{"applyConfigurationCommand": {"config":
{"appId": "com.example.mobile"}, "optional": false}}, {"assertConditionCommand":
{"condition": {"visible": {"textRegex": "欢迎使用\n  2 | COMPLETED    |   3060ms
| {"tapOnElement": {"selector": {"textRegex": "验证码", "optional": false}, "ret
ryIfNoChange": true, "waitUntilVisible": false, "longPress": false, "optional": 
false}}\n  3 | COMPLETED    |   1788ms | {"inputTextCommand": {"text": "123456",
"optional": false}}\n  4 | COMPLETED    |   1716ms | {"assertConditionCommand": 
{"condition": {"visible": {"textRegex": "欢迎使用 Social World", "optional": fal
se}}, "optional": false}}\n  5 | COMPLETED    |    550ms | {"assertConditionComm
and": {"condition": {"visible": {"textRegex": "手机号一键登录", "optional": fals
e}}, "optional": false}}\n  6 | COMPLETED    |   1064ms | {"hideKeyboardCommand"
: {"optional": false}}\n  7 | FAILED       |  17082ms | {"assertConditionCommand
": {"condition": {"notVisible": {"textRegex": "手机号一键登录", "optional": fals
e}}, "optional": false}}\n  8 | COMPLETED    |      0ms | {"applyConfigurationCo
mmand": {"config": {"appId": "com.example.mobile"}, "optional": false}}\n  9 | C
OMPLETED    |    837ms | {"launchAppCommand": {"appId": "com.example.mobile", "c
learState": true, "optional": false}}\n 10 | COMPLETED    |   2035ms | {"inputTe
xtCommand": {"text": "13800138000", "optional": false}}\n 11 | COMPLETED    |   
2273ms | {"tapOnElement": {"selector": {"textRegex": "获取验证码", "optional": f
alse}, "retryIfNoChange": true, "waitUntilVisible": false, "longPress": false, "
optional": false}}\n 12 | COMPLETED    |   2724ms | {"tapOnElement": {"selector"
: {"textRegex": "手机号", "optional": false}, "retryIfNoChange": true, "waitUnti
lVisible": false, "longPress": false, "optional": false}}\n 13 | COMPLETED    | 
    0ms | {"applyConfigurationCommand": {"config": {"appId": "com.example.mobile
"}, "optional": false}}\n 14 | COMPLETED    |   2242ms | {"tapOnElement": {"sele
ctor": {"textRegex": "登录", "optional": false}, "retryIfNoChange": true, "waitU
ntilVisible": false, "longPress": false, "optional": false}}\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=27)
ThinkPart(
    type='think',
    think='从 commands JSON 可以看到一个奇怪的现象：\n\n`login.yaml` 的步骤顺序
看起来不对！\n- 步骤 1: runFlow (includes/login-as-test-user.yaml) - 耗时 17453m
s，COMPLETED\n- 步骤 2: tapOn "验证码" - 3060ms\n- 步骤 3: inputText "123456" - 
1788ms\n- 步骤 4: assertVisible "欢迎使用 Social World" - 1716ms\n- 步骤 5: asse
rtVisible "手机号一键登录" - 550ms\n- 步骤 6: hideKeyboard - 1064ms\n- 步骤 7: a
ssertNotVisible "手机号一键登录" - FAILED 17082ms\n\n等等，这顺序完全乱了！runFl
ow 包含了登录步骤，但之后又在重复执行 tapOn "验证码" 和 inputText？不对，让我仔
细看。\n\n实际上 commands JSON 的排序可能不是执行顺序。Maestro 的 commands JSON 
可能按照某种内部顺序记录。\n\n等等，看 `sourceDescription`：\n- 步骤 1 的 `sourc
eDescription` 是 `includes/login-as-test-user.yaml`\n- 步骤 2-7 没有 sourceDescr
iption\n\n不对，让我重新看原始数据。实际上 commands JSON 可能不是按时间顺序排列
的，或者 Maestro 的 flatten 输出把多个 flow 的步骤混在一起了。\n\n但关键是：`ass
ertNotVisible "手机号一键登录"` 失败了，说明页面仍然显示登录页。\n\n从截图来看，
手机号和验证码都输入了，"登录"按钮可见，但没有点击？不对，步骤 14 显示 `tapOn "
登录"` 是 COMPLETED。\n\n等等，login.yaml 的内容是：\n```\n- launchApp: clearSta
te: true\n- runFlow: includes/login-as-test-user.yaml\n- assertNotVisible: "手机
号一键登录"\n- assertVisible: "完善资料" optional: true\n- assertVisible: id: "t
ab_home" optional: true\n```\n\n而 includes/login-as-test-user.yaml 的内容是：\n
```\n- assertVisible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"
\n- tapOn: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- tapOn:
"验证码"\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n```\n\n所以执
行顺序应该是：\n1. launchApp\n2. assertVisible "欢迎使用 Social World"\n3. asser
tVisible "手机号一键登录"\n4. tapOn "手机号"\n5. inputText "13800138000"\n6. tap
On "获取验证码"\n7. tapOn "验证码"\n8. inputText "123456"\n9. hideKeyboard\n10. 
tapOn "登录"\n11. assertNotVisible "手机号一键登录"\n\n但截图显示页面仍然在登录
页，而且倒计时是 "35 s"。这意味着在 `tapOn "登录"` 之后，页面没有跳转。可能的原
因：\n1. 登录 API 请求失败（网络错误、服务器未响应）\n2. 登录 API 返回了错误（用
户不存在等）\n3. 登录成功后应用崩溃或未正确跳转\n\n让我检查 API 的后端日志。但更
重要的是，我需要看看 `login-as-test-user.yaml` 中点击"登录"后，是否有什么错误提
示。\n\n等等，截图是在 `assertNotVisible` 失败时截取的。截图中没有任何 SnackBar 
错误消息，这说明：\n- 登录请求可能还在 pending 状态（网络超时）\n- 或者后端返回
了成功，但 Bloc 状态没有正确更新\n\n让我检查一下 auth_bloc.dart 的实现，看看登录
成功后的状态流转。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_fqHH8iTJUHivhTpOJQtSVbLp',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/apps/mobile/lib/presentation/blocs/auth/auth_bloc.dart"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_aBMD0SOpEgT2QGpziV3vERy3',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_fqHH8iTJUHivhTpOJQtSVbLp',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter_bloc/flutter_bloc.dart';\n     2
\timport '../../../core/storage/secure_storage.dart';\n     3\timport '../../../
domain/repositories/auth_repository.dart';\n     4\timport 'auth_event.dart';\n 
   5\timport 'auth_state.dart';\n     6\t\n     7\tclass AuthBloc extends Bloc<A
uthEvent, AuthState> {\n     8\t  final AuthRepository _authRepository;\n     9\
t  final SecureStorage _secureStorage;\n    10\t\n    11\t  AuthBloc({\n    12\t
   required AuthRepository authRepository,\n    13\t    required SecureStorage s
ecureStorage,\n    14\t  })  : _authRepository = authRepository,\n    15\t      
 _secureStorage = secureStorage,\n    16\t        super(const AuthUnauthenticate
d()) {\n    17\t    on<LoginRequested>(_onLoginRequested);\n    18\t    on<Logou
tRequested>(_onLogoutRequested);\n    19\t    on<AuthCheckRequested>(_onAuthChec
kRequested);\n    20\t    on<TokenRefreshed>(_onTokenRefreshed);\n    21\t  }\n 
  22\t\n    23\t  Future<void> _onLoginRequested(\n    24\t    LoginRequested ev
ent,\n    25\t    Emitter<AuthState> emit,\n    26\t  ) async {\n    27\t    emi
t(const AuthAuthenticating());\n    28\t\n    29\t    try {\n    30\t      final
loginResponse =\n    31\t          await _authRepository.login(event.phone, even
t.code);\n    32\t\n    33\t      final accessToken = loginResponse.accessToken;
\n    34\t      final refreshToken = loginResponse.refreshToken;\n    35\t      
final isNewUser = loginResponse.isNewUser;\n    36\t\n    37\t      await _secur
eStorage.writeAccessToken(accessToken);\n    38\t      await _secureStorage.writ
eRefreshToken(refreshToken);\n    39\t\n    40\t      final me = await _authRepo
sitory.getMe();\n    41\t\n    42\t      emit(AuthAuthenticated(\n    43\t      
 userId: me.id,\n    44\t        isNewUser: isNewUser || me.isNewUser,\n    45\t
       profileComplete: me.profileComplete,\n    46\t      ));\n    47\t    } on
Exception catch (e) {\n    48\t      emit(AuthFailure(e.toString()));\n    49\t 
  }\n    50\t  }\n    51\t\n    52\t  Future<void> _onLogoutRequested(\n    53\t
   LogoutRequested event,\n    54\t    Emitter<AuthState> emit,\n    55\t  ) asy
nc {\n    56\t    try {\n    57\t      final refreshToken = await _secureStorage
.readRefreshToken();\n    58\t      if (refreshToken != null) {\n    59\t       
await _authRepository.logout(refreshToken);\n    60\t      }\n    61\t    } catc
h (_) {\n    62\t      // Ignore logout errors\n    63\t    }\n    64\t\n    65\
t    await _secureStorage.clearAll();\n    66\t    emit(const AuthUnauthenticate
d());\n    67\t  }\n    68\t\n    69\t  Future<void> _onAuthCheckRequested(\n   
70\t    AuthCheckRequested event,\n    71\t    Emitter<AuthState> emit,\n    72\
t  ) async {\n    73\t    final accessToken = await _secureStorage.readAccessTok
en();\n    74\t    if (accessToken == null || accessToken.isEmpty) {\n    75\t  
   emit(const AuthUnauthenticated());\n    76\t      return;\n    77\t    }\n   
78\t\n    79\t    try {\n    80\t      final me = await _authRepository.getMe();
\n    81\t\n    82\t      emit(AuthAuthenticated(\n    83\t        userId: me.id
,\n    84\t        isNewUser: me.isNewUser,\n    85\t        profileComplete: me
.profileComplete,\n    86\t      ));\n    87\t    } on Exception {\n    88\t    
 emit(const AuthUnauthenticated());\n    89\t    }\n    90\t  }\n    91\t\n    9
2\t  Future<void> _onTokenRefreshed(\n    93\t    TokenRefreshed event,\n    94\
t    Emitter<AuthState> emit,\n    95\t  ) async {\n    96\t    // Handled by in
terceptor; this event can be used for explicit refresh triggers.\n    97\t  }\n 
  98\t}\n",
        message='98 lines read from file starting from line 1. Total lines in fi
le: 98. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='/apps/mobile/lib/data/repositories/auth_repository_impl.dart
"}'
)
StatusUpdate(
    context_usage=0.24736404418945312,
    context_tokens=64845,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1101,
        output=891,
        input_cache_read=63744,
        input_cache_creation=0
    ),
    message_id='chatcmpl-FBYUPt2ydsTLjKySO2nyeg0c',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_aBMD0SOpEgT2QGpziV3vERy3',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:dio/dio.dart';\n     2\timport '../../do
main/repositories/auth_repository.dart';\n     3\timport '../../core/di/di.dart'
;\n     4\timport '../../core/network/dio_client.dart';\n     5\timport 'package
:social_world_auth_models/api.dart';\n     6\t\n     7\tclass AuthRepositoryImpl
implements AuthRepository {\n     8\t  final Dio _dio;\n     9\t\n    10\t  Auth
RepositoryImpl({Dio? dio}) : _dio = dio ?? getIt<DioClient>().dio;\n    11\t\n  
 12\t  @override\n    13\t  Future<LoginResponse> login(String phone, String cod
e) async {\n    14\t    final response = await _dio.post<Map<String, dynamic>>(\
n    15\t      '/auth/login',\n    16\t      data: LoginRequest(phone: phone, sm
sCode: code).toJson(),\n    17\t    );\n    18\t    final data = response.data;\
n    19\t    if (data == null) {\n    20\t      throw Exception('Invalid respons
e from server');\n    21\t    }\n    22\t    final login200 = Login200Response.f
romJson(data);\n    23\t    if (login200 == null) {\n    24\t      throw Excepti
on('Invalid response from server');\n    25\t    }\n    26\t    return login200.
data;\n    27\t  }\n    28\t\n    29\t  @override\n    30\t  Future<LoginRespons
e> refresh(String refreshToken) async {\n    31\t    final response = await _dio
.post<Map<String, dynamic>>(\n    32\t      '/auth/refresh',\n    33\t      data
: RefreshTokenRequest(refreshToken: refreshToken).toJson(),\n    34\t      optio
ns: Options(\n    35\t        headers: {'Authorization': null},\n    36\t      )
,\n    37\t    );\n    38\t    final data = response.data;\n    39\t    if (data
== null) {\n    40\t      throw Exception('Invalid response from server');\n    
41\t    }\n    42\t    final login200 = Login200Response.fromJson(data);\n    43
\t    if (login200 == null) {\n    44\t      throw Exception('Invalid response f
rom server');\n    45\t    }\n    46\t    return login200.data;\n    47\t  }\n  
 48\t\n    49\t  @override\n    50\t  Future<void> logout(String refreshToken) a
sync {\n    51\t    await _dio.post<Map<String, dynamic>>(\n    52\t      '/auth
/logout',\n    53\t      data: {'refresh_token': refreshToken},\n    54\t    );\
n    55\t  }\n    56\t\n    57\t  @override\n    58\t  Future<UserResponse> getM
e() async {\n    59\t    final response = await _dio.get<Map<String, dynamic>>('
/auth/me');\n    60\t    final data = response.data;\n    61\t    if (data == nu
ll) {\n    62\t      throw Exception('Invalid response from server');\n    63\t 
  }\n    64\t    final me200 = GetMe200Response.fromJson(data);\n    65\t    if 
(me200 == null) {\n    66\t      throw Exception('Invalid response from server')
;\n    67\t    }\n    68\t    return me200.data;\n    69\t  }\n    70\t}\n",
        message='70 lines read from file starting from line 1. Total lines in fi
le: 70. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=28)
ThinkPart(
    type='think',
    think='从代码看，登录流程是：\n1. `login` 调用 POST `/auth/login`\n2. 成功后
调用 `getMe()` 调用 GET `/auth/me`\n3. 然后 emit `AuthAuthenticated`\n\n如果 `ge
tMe()` 失败，会 catch 并 emit `AuthFailure`。\n\n截图中没有 SnackBar 错误，说明
：\n- 可能 `getMe()` 卡住了（网络超时）\n- 或者登录请求本身卡住了\n\n但 `assertN
otVisible` 等待了 17082ms（约17秒），这接近 `extendedWaitUntil: timeout: 10000` 
加上其他延迟。这说明 Maestro 等了很长时间但页面还是没变。\n\n让我检查一下后端 AP
I 在 3001 端口上是否有 `/auth/login` 路由。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_Z4PLkgKAtp5mxQaheTkU7oce',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "curl -s -X POST http://localhost:3001/api/v1/aut
h/login -H \\"Content-Type: application/json\\" -d \'{\\"phone\\":\\"13800138000
\\",\\"smsCode\\":\\"123456\\"}\' | head -c 300"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_cAD0qxsyAapnnciXjzEGsXw4',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_Z4PLkgKAtp5mxQaheTkU7oce',
    return_value=ToolReturnValue(
        is_error=False,
        output='{"code":1000,"message":"property smsCode should not exist,sms_co
de must be longer than or equal to 6 characters,sms_code must be a string","data
":null}',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"command": "curl -s http://localhost:3001/api/v1/auth/me -H
\\"Authorization: Bearer test\\" | head -c 300"}'
)
StatusUpdate(
    context_usage=0.2573127746582031,
    context_tokens=67453,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=2685,
        output=270,
        input_cache_read=64768,
        input_cache_creation=0
    ),
    message_id='chatcmpl-BhoXBJsBhNRini6KsxkjcB4y',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_cAD0qxsyAapnnciXjzEGsXw4',
    return_value=ToolReturnValue(
        is_error=False,
        output='{"code":2001,"message":"Invalid or expired token","data":null}',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=29)
ThinkPart(
    type='think',
    think='Aha! 找到了根本原因！\n\n后端 API 期望的请求体字段是 `sms_code`（snak
e_case），但前端发送的是 `smsCode`（camelCase）。\n\n`LoginRequest` 模型生成的 `
.toJson()` 可能输出的是 `smsCode`，但 NestJS 后端（使用 class-validator）期望的
是 `sms_code`。\n\n这是一个 **产品 bug**：前端和后端的请求体字段名不一致。\n\n但
是，等等。`login-as-test-user.yaml` 中的测试在之前的运行中是否曾经成功过？从 PRD
和代码来看，这个 bug 应该是之前就存在的。\n\n让我验证一下后端实际期望的字段名。'
,
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_gtA53G8IjWOV3sT0qEUFSdGi',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100 && grep -rn \\"smsCode\\\\|sms_code\\" apps/api/src -
-include=\\"*.ts\\" | head -20"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_MuJ1Iutvht61h17j4zPNPhct',
    function=FunctionBody(name='Shell', arguments='{"command": "cd /'),
    extras=None
)
ToolResult(
    tool_call_id='tool_gtA53G8IjWOV3sT0qEUFSdGi',
    return_value=ToolReturnValue(
        is_error=False,
        output='apps/api/src/modules/auth/dto/login.dto.ts:11:  sms_code: string
;\napps/api/src/modules/auth/auth.controller.ts:27:    const data = await this.a
uthService.login(dto.phone, dto.sms_code);\napps/api/src/modules/auth/auth.servi
ce.ts:37:    const storedCode = await this.redis.get(`sms_code:${phone}`);\napps
/api/src/modules/auth/auth.service.ts:122:  async login(phone: string, smsCode: 
string): Promise<AuthResponseDto> {\napps/api/src/modules/auth/auth.service.ts:1
23:    await this.validateSmsCode(phone, smsCode);\napps/api/src/modules/auth/au
th.service.ts:136:      await this.redis.del(`sms_code:${phone}`);\napps/api/src
/modules/auth/auth.service.ts:283:    await this.redis.setex(`sms_code:${phone}`
, 300, code);\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100 && grep -rn \\"smsCode\\\\|sms_code\\" apps/mobile/lib --include=\\"*
.dart\\" | head -20"}'
)
StatusUpdate(
    context_usage=0.2587394714355469,
    context_tokens=67827,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=499,
        output=289,
        input_cache_read=67328,
        input_cache_creation=0
    ),
    message_id='chatcmpl-oJBGK0siZqCumAEeiiZy0Zna',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_MuJ1Iutvht61h17j4zPNPhct',
    return_value=ToolReturnValue(
        is_error=False,
        output='apps/mobile/lib/models/generated/auth/test/login_request_test.da
rt:24:    // String smsCode\napps/mobile/lib/models/generated/auth/test/login_re
quest_test.dart:25:    test(\'to test the property `smsCode`\', () async {\napps
/mobile/lib/models/generated/auth/lib/model/login_request.dart:17:    required t
his.smsCode,\napps/mobile/lib/models/generated/auth/lib/model/login_request.dart
:22:  String smsCode;\napps/mobile/lib/models/generated/auth/lib/model/login_req
uest.dart:27:    other.smsCode == smsCode;\napps/mobile/lib/models/generated/aut
h/lib/model/login_request.dart:33:    (smsCode.hashCode);\napps/mobile/lib/model
s/generated/auth/lib/model/login_request.dart:36:  String toString() => \'LoginR
equest[phone=$phone, smsCode=$smsCode]\';\napps/mobile/lib/models/generated/auth
/lib/model/login_request.dart:41:      json[r\'sms_code\'] = this.smsCode;\napps
/mobile/lib/models/generated/auth/lib/model/login_request.dart:58:        assert
(json.containsKey(r\'sms_code\'), \'Required key "LoginRequest[sms_code]" is mis
sing from JSON.\');\napps/mobile/lib/models/generated/auth/lib/model/login_reque
st.dart:59:        assert(json[r\'sms_code\'] != null, \'Required key "LoginRequ
est[sms_code]" has a null value in JSON.\');\napps/mobile/lib/models/generated/a
uth/lib/model/login_request.dart:65:        smsCode: mapValueOfType<String>(json
, r\'sms_code\')!,\napps/mobile/lib/models/generated/auth/lib/model/login_reques
t.dart:114:    \'sms_code\',\napps/mobile/lib/data/repositories/auth_repository_
impl.dart:16:      data: LoginRequest(phone: phone, smsCode: code).toJson(),\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=30)
ThinkPart(
    type='think',
    think="非常有趣！`login_request.dart` 的 `toJson()` 方法实际上输出的是 `sms_
code`（第41行：`json[r'sms_code'] = this.smsCode;`）。所以前端发送的请求体应该是
正确的。\n\n但为什么后端返回了 `property smsCode should not exist`？这说明后端接
收到的请求体中包含了 `smsCode` 字段。\n\n等等，让我再仔细看看。`login_request.da
rt` 第41行确实是 `json[r'sms_code'] = this.smsCode;`。但 Dio 可能添加了额外的序
列化逻辑？\n\n或者，`LoginRequest` 生成的代码有问题？让我读取完整的 `login_reque
st.dart` 文件。",
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_RMpfZRqkFBILyLhryOyG7uZb',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/apps/mobile/lib/models/generated/auth/lib/model/login_reque
st.dart"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.2620735168457031,
    context_tokens=68701,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1117,
        output=205,
        input_cache_read=67584,
        input_cache_creation=0
    ),
    message_id='chatcmpl-C1dDXNWEnVMSodZyeMasjbGk',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_RMpfZRqkFBILyLhryOyG7uZb',
    return_value=ToolOk(
        is_error=False,
        output='     1\t//\n     2\t// AUTO-GENERATED FILE, DO NOT MODIFY!\n    
3\t//\n     4\t// @dart=2.18\n     5\t\n     6\t// ignore_for_file: unused_eleme
nt, unused_import\n     7\t// ignore_for_file: always_put_required_named_paramet
ers_first\n     8\t// ignore_for_file: constant_identifier_names\n     9\t// ign
ore_for_file: lines_longer_than_80_chars\n    10\t\n    11\tpart of openapi.api;
\n    12\t\n    13\tclass LoginRequest {\n    14\t  /// Returns a new [LoginRequ
est] instance.\n    15\t  LoginRequest({\n    16\t    required this.phone,\n    
17\t    required this.smsCode,\n    18\t  });\n    19\t\n    20\t  String phone;
\n    21\t\n    22\t  String smsCode;\n    23\t\n    24\t  @override\n    25\t  
bool operator ==(Object other) => identical(this, other) || other is LoginReques
t &&\n    26\t    other.phone == phone &&\n    27\t    other.smsCode == smsCode;
\n    28\t\n    29\t  @override\n    30\t  int get hashCode =>\n    31\t    // i
gnore: unnecessary_parenthesis\n    32\t    (phone.hashCode) +\n    33\t    (sms
Code.hashCode);\n    34\t\n    35\t  @override\n    36\t  String toString() => \
'LoginRequest[phone=$phone, smsCode=$smsCode]\';\n    37\t\n    38\t  Map<String
, dynamic> toJson() {\n    39\t    final json = <String, dynamic>{};\n    40\t  
   json[r\'phone\'] = this.phone;\n    41\t      json[r\'sms_code\'] = this.smsC
ode;\n    42\t    return json;\n    43\t  }\n    44\t\n    45\t  /// Returns a n
ew [LoginRequest] instance and imports its values from\n    46\t  /// [value] if
it\'s a [Map], null otherwise.\n    47\t  // ignore: prefer_constructors_over_st
atic_methods\n    48\t  static LoginRequest? fromJson(dynamic value) {\n    49\t
   if (value is Map) {\n    50\t      final json = value.cast<String, dynamic>()
;\n    51\t\n    52\t      // Ensure that the map contains the required keys.\n 
  53\t      // Note 1: the values aren\'t checked for validity beyond being non-
null.\n    54\t      // Note 2: this code is stripped in release mode!\n    55\t
     assert(() {\n    56\t        assert(json.containsKey(r\'phone\'), \'Require
d key "LoginRequest[phone]" is missing from JSON.\');\n    57\t        assert(js
on[r\'phone\'] != null, \'Required key "LoginRequest[phone]" has a null value in
JSON.\');\n    58\t        assert(json.containsKey(r\'sms_code\'), \'Required ke
y "LoginRequest[sms_code]" is missing from JSON.\');\n    59\t        assert(jso
n[r\'sms_code\'] != null, \'Required key "LoginRequest[sms_code]" has a null val
ue in JSON.\');\n    60\t        return true;\n    61\t      }());\n    62\t\n  
 63\t      return LoginRequest(\n    64\t        phone: mapValueOfType<String>(j
son, r\'phone\')!,\n    65\t        smsCode: mapValueOfType<String>(json, r\'sms
_code\')!,\n    66\t      );\n    67\t    }\n    68\t    return null;\n    69\t 
}\n    70\t\n    71\t  static List<LoginRequest> listFromJson(dynamic json, {boo
l growable = false,}) {\n    72\t    final result = <LoginRequest>[];\n    73\t 
  if (json is List && json.isNotEmpty) {\n    74\t      for (final row in json) 
{\n    75\t        final value = LoginRequest.fromJson(row);\n    76\t        if
(value != null) {\n    77\t          result.add(value);\n    78\t        }\n    
79\t      }\n    80\t    }\n    81\t    return result.toList(growable: growable)
;\n    82\t  }\n    83\t\n    84\t  static Map<String, LoginRequest> mapFromJson
(dynamic json) {\n    85\t    final map = <String, LoginRequest>{};\n    86\t   
if (json is Map && json.isNotEmpty) {\n    87\t      json = json.cast<String, dy
namic>(); // ignore: parameter_assignments\n    88\t      for (final entry in js
on.entries) {\n    89\t        final value = LoginRequest.fromJson(entry.value);
\n    90\t        if (value != null) {\n    91\t          map[entry.key] = value
;\n    92\t        }\n    93\t      }\n    94\t    }\n    95\t    return map;\n 
  96\t  }\n    97\t\n    98\t  // maps a json object with a list of LoginRequest
-objects as value to a dart map\n    99\t  static Map<String, List<LoginRequest>
> mapListFromJson(dynamic json, {bool growable = false,}) {\n   100\t    final m
ap = <String, List<LoginRequest>>{};\n   101\t    if (json is Map && json.isNotE
mpty) {\n   102\t      // ignore: parameter_assignments\n   103\t      json = js
on.cast<String, dynamic>();\n   104\t      for (final entry in json.entries) {\n
  105\t        map[entry.key] = LoginRequest.listFromJson(entry.value, growable:
growable,);\n   106\t      }\n   107\t    }\n   108\t    return map;\n   109\t  
}\n   110\t\n   111\t  /// The list of required keys that must be present in a J
SON.\n   112\t  static const requiredKeys = <String>{\n   113\t    \'phone\',\n 
 114\t    \'sms_code\',\n   115\t  };\n   116\t}\n   117\t\n',
        message='117 lines read from file starting from line 1. Total lines in f
ile: 117. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=31)
ThinkPart(
    type='think',
    think='`LoginRequest.toJson()` 确实输出 `sms_code`。但后端报错说 `property s
msCode should not exist`。这说明后端接收到的 JSON 里同时有 `sms_code` 和 `smsCod
e`？或者只有 `smsCode`？\n\n等等，Dio 的 `LogInterceptor` 会打印请求体。让我检查
Maestro 测试运行时的 logcat，看看 App 实际发送了什么。\n\n或者，问题可能在于：`a
uth_repository_impl.dart` 中的 `toJson()` 方法返回 `Map<String, dynamic>`，然后 
Dio 发送这个 map。但 `LoginRequest` 的 `toJson()` 确实输出 `sms_code`。\n\n让我
用 curl 直接测试后端，使用正确的 `sms_code` 字段，看看后端是否正常工作。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_BMOvupaTN72PR9evcUyRnh0v',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "curl -s -X POST http://localhost:3001/api/v1/aut
h/login -H \\"Content-Type: application/json\\" -d \'{\\"phone\\":\\"13800138000
\\",\\"sms_code\\":\\"123456\\"}\' | head -c 300"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.2679100036621094,
    context_tokens=70231,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1623,
        output=227,
        input_cache_read=68608,
        input_cache_creation=0
    ),
    message_id='chatcmpl-Xq135Uk8P8u3tnEJUnYrdlwh',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_BMOvupaTN72PR9evcUyRnh0v',
    return_value=ToolReturnValue(
        is_error=False,
        output='{"code":1001,"message":"Invalid or expired SMS code","data":null
}',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=32)
ThinkPart(
    type='think',
    think='好的，使用正确的 `sms_code` 字段名后，后端返回 `Invalid or expired SM
S code`（验证码错误或过期）。这说明后端是正常运行的。\n\n但问题是：前端 `LoginRe
quest.toJson()` 确实输出 `sms_code`，为什么后端会收到 `smsCode`？\n\n让我查看一
下 `LogInterceptor` 的日志。我可以从模拟器的 logcat 中获取。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_rN9a3AZhXnZzjP9N0xnfXyNw',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "adb -s emulator-5556 logcat -d | grep -E \\"smsC
ode|sms_code|/auth/login|DioError\\" | tail -20"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.26894378662109375,
    context_tokens=70502,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=358,
        output=130,
        input_cache_read=70144,
        input_cache_creation=0
    ),
    message_id='chatcmpl-JlhvFLIl8c3Sr5H5qpXpu8sW',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_rN9a3AZhXnZzjP9N0xnfXyNw',
    return_value=ToolReturnValue(
        is_error=False,
        output='04-24 01:00:56.991 10661 10688 I flutter : uri: http://10.0.2.2:
3001/api/v1/auth/login\n04-24 01:02:18.393 10915 10939 I flutter : uri: http://1
0.0.2.2:3001/api/v1/auth/login\n04-24 01:02:18.394 10915 10939 I flutter : {phon
e: 13800138000, sms_code: 123456}\n04-24 01:02:18.499 10915 10939 I flutter : ur
i: http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:27:10.433 12206 12232 I flut
ter : uri: http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:27:10.433 12206 1223
2 I flutter : {phone: 13800138000, sms_code: 123456}\n04-24 01:27:10.513 12206 1
2232 I flutter : uri: http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:27:10.513
12206 12232 I flutter : uri: http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:27
:44.563 12321 12347 I flutter : uri: http://10.0.2.2:3001/api/v1/auth/login\n04-
24 01:27:44.564 12321 12347 I flutter : {phone: 13800138000, sms_code: 123456}\n
04-24 01:27:44.653 12321 12347 I flutter : uri: http://10.0.2.2:3001/api/v1/auth
/login\n04-24 01:27:44.653 12321 12347 I flutter : uri: http://10.0.2.2:3001/api
/v1/auth/login\n04-24 01:28:19.325 12435 12460 I flutter : uri: http://10.0.2.2:
3001/api/v1/auth/login\n04-24 01:28:19.326 12435 12460 I flutter : {phone: 13800
138000, sms_code: 123456}\n04-24 01:28:19.398 12435 12460 I flutter : uri: http:
//10.0.2.2:3001/api/v1/auth/login\n04-24 01:28:19.398 12435 12460 I flutter : ur
i: http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:28:52.885 12545 12570 I flut
ter : uri: http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:28:52.886 12545 1257
0 I flutter : {phone: 13800138000, sms_code: 123456}\n04-24 01:28:52.969 12545 1
2570 I flutter : uri: http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:28:52.971
12545 12570 I flutter : uri: http://10.0.2.2:3001/api/v1/auth/login\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=33)
ThinkPart(
    type='think',
    think='logcat 显示前端确实发送了 `{phone: 13800138000, sms_code: 123456}`，
这是正确的。但后端返回了 `Invalid or expired SMS code`（code: 1001）。\n\n这说明
后端没有为测试手机号 `13800138000` 设置正确的验证码。MVP 阶段 PRD 说"短信验证码 
MVP 阶段可先用固定码（如 `123456`）"，但后端实现中可能是从 Redis 读取验证码。\n\
n让我查看后端的 auth service，了解验证码是如何验证的。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_owaYCDQvk0gHcHNb5x9tLDJg',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/apps/api/src/modules/auth/auth.service.ts", "n_lines": 60}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.27251434326171875,
    context_tokens=71438,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1038,
        output=181,
        input_cache_read=70400,
        input_cache_creation=0
    ),
    message_id='chatcmpl-lwncnqlu404OITzHqZ42D7pa',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_owaYCDQvk0gHcHNb5x9tLDJg',
    return_value=ToolOk(
        is_error=False,
        output='     1\timport { Injectable, HttpStatus } from "@nestjs/common";
\n     2\timport { InjectRepository } from "@nestjs/typeorm";\n     3\timport { 
Repository } from "typeorm";\n     4\timport Redis from "ioredis";\n     5\timpo
rt * as jwt from "jsonwebtoken";\n     6\timport { v7 as uuidv7 } from "uuid";\n
    7\timport { UserEntity } from "../../database/entities/user.entity";\n     8
\timport { UserProfileEntity } from "../../database/entities/user-profile.entity
";\n     9\timport { AuthResponseDto } from "./dto/auth-response.dto";\n    10\t
import { UserResponseDto } from "./dto/user-response.dto";\n    11\timport { Bus
inessException } from "../../common/exceptions/business.exception";\n    12\timp
ort { jwtConfig } from "../../config/jwt.config";\n    13\timport { redisConfig 
} from "../../config/redis.config";\n    14\t\n    15\tinterface JwtPayload {\n 
  16\t  sub: string;\n    17\t  jti: string;\n    18\t  type: "access" | "refres
h";\n    19\t  iat: number;\n    20\t  exp: number;\n    21\t}\n    22\t\n    23
\t@Injectable()\n    24\texport class AuthService {\n    25\t  private readonly 
redis: Redis;\n    26\t\n    27\t  constructor(\n    28\t    @InjectRepository(U
serEntity)\n    29\t    private readonly userRepo: Repository<UserEntity>,\n    
30\t    @InjectRepository(UserProfileEntity)\n    31\t    private readonly profi
leRepo: Repository<UserProfileEntity>,\n    32\t  ) {\n    33\t    this.redis = 
new Redis(redisConfig);\n    34\t  }\n    35\t\n    36\t  async validateSmsCode(
phone: string, code: string): Promise<void> {\n    37\t    const storedCode = aw
ait this.redis.get(`sms_code:${phone}`);\n    38\t    if (!storedCode || storedC
ode !== code) {\n    39\t      throw new BusinessException(\n    40\t        100
1,\n    41\t        "Invalid or expired SMS code",\n    42\t        HttpStatus.B
AD_REQUEST,\n    43\t      );\n    44\t    }\n    45\t  }\n    46\t\n    47\t  a
sync findOrCreateUser(\n    48\t    phone: string,\n    49\t  ): Promise<{ user:
UserEntity; isNew: boolean }> {\n    50\t    let user = await this.userRepo.find
One({ where: { phone } });\n    51\t    let isNew = false;\n    52\t\n    53\t  
 if (!user) {\n    54\t      const userId = uuidv7();\n    55\t      user = this
.userRepo.create({ id: userId, phone });\n    56\t      await this.userRepo.save
(user);\n    57\t\n    58\t      const profile = this.profileRepo.create({\n    
59\t        user_id: userId,\n    60\t        is_complete: false,\n',
        message='60 lines read from file starting from line 1. Total lines in fi
le: 285.',
        display=[],
        extras=None
    )
)
StepBegin(n=34)
ThinkPart(
    type='think',
    think='后端 `validateSmsCode` 从 Redis 读取 `sms_code:${phone}`，如果不存在
或不匹配就抛出错误。这说明测试前需要先把验证码写入 Redis。\n\n这是一个**环境问题
/测试前置条件缺失**：测试流程没有预先将验证码写入 Redis。\n\n根据 PRD 和 login-a
s-test-user.yaml 的注释：\n> MVP 阶段 _onSendCode 只启动本地倒计时，不调用后端 A
PI\n\n所以前端 `_onSendCode` 没有调用后端发送验证码的 API。后端 Redis 中不会有这
个验证码。\n\n这意味着在运行 E2E 测试之前，需要手动将测试验证码写入 Redis：\n```
bash\nredis-cli SET sms_code:13800138000 123456\n```\n\n这是**环境问题**，不是产
品 bug，也不是 flow 过时。因为 flow 假设后端有可用的测试验证码，但环境没有预先设
置。\n\n让我设置 Redis 中的验证码，然后重新运行测试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_aP8OtIzKJXvrUvBS0b8lYIjv',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "redis-cli SET sms_code:13800138000 123456 EX 300
"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_p8IwWisDdsSjxRueccTWMSGk',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_aP8OtIzKJXvrUvBS0b8lYIjv',
    return_value=ToolReturnValue(
        is_error=True,
        output='/bin/bash: redis-cli: command not found\n',
        message='Command failed with exit code: 127.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 127')
        ],
        extras=None
    )
)
ToolCallPart(arguments_part='{"command": "redis-cli GET sms_code:13800138000"}')
StatusUpdate(
    context_usage=0.27579498291015625,
    context_tokens=72298,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=874,
        output=243,
        input_cache_read=71424,
        input_cache_creation=0
    ),
    message_id='chatcmpl-0rbWB3Q4C3EF3m2vzsXAKRYR',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_p8IwWisDdsSjxRueccTWMSGk',
    return_value=ToolReturnValue(
        is_error=True,
        output='/bin/bash: redis-cli: command not found\n',
        message='Command failed with exit code: 127.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 127')
        ],
        extras=None
    )
)
StepBegin(n=35)
ThinkPart(
    type='think',
    think='redis-cli 不可用，但 Docker 里有 redis。我可以用 docker exec 来执行 r
edis-cli。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_nAd5lSF46KsDEQuxYQBMWZg3',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "docker exec sw-redis redis-cli SET sms_code:1380
0138000 123456 EX 300"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_Bp9NjQ3Z5mdKTxF0mHdy974G',
    function=FunctionBody(name='Shell', arguments='{"command":'),
    extras=None
)
ToolResult(
    tool_call_id='tool_nAd5lSF46KsDEQuxYQBMWZg3',
    return_value=ToolReturnValue(
        is_error=False,
        output='OK\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part=' "docker exec sw-redis redis-cli GET sms_code:13800138000"}'
)
StatusUpdate(
    context_usage=0.2770347595214844,
    context_tokens=72623,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=431,
        output=87,
        input_cache_read=72192,
        input_cache_creation=0
    ),
    message_id='chatcmpl-7HK04kSkQHFrw9T32FFkzZbe',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_Bp9NjQ3Z5mdKTxF0mHdy974G',
    return_value=ToolReturnValue(
        is_error=False,
        output='123456\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=36)
ThinkPart(
    type='think',
    think='验证码已写入 Redis。现在让我重新运行所有失败的 flow。但等等，`open_dr
awer` 和 `edit_server_url` 已经通过了，我可以只运行失败的 4 个 flow。不过为了完
整性，我可以重新运行所有 flow。\n\n但在重新运行之前，让我先确认后端 API 现在能正
常登录。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_K6ksKwP4qPwwuHQshOCLlIzt',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "curl -s -X POST http://localhost:3001/api/v1/aut
h/login -H \\"Content-Type: application/json\\" -d \'{\\"phone\\":\\"13800138000
\\",\\"sms_code\\":\\"123456\\"}\' | head -c 300"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.2775688171386719,
    context_tokens=72763,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=315,
        output=134,
        input_cache_read=72448,
        input_cache_creation=0
    ),
    message_id='chatcmpl-RUtUZJ4dkcUBJwBAN9DtWmpg',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_K6ksKwP4qPwwuHQshOCLlIzt',
    return_value=ToolReturnValue(
        is_error=False,
        output='{"code":0,"message":"success","data":{"access_token":"eyJhbGciOi
JIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwMTlkYjU2ZC01ODk3LTc0NDAtOGVjNy0xMDNhZDAwO
GIwNDMiLCJqdGkiOiIwMTlkYmI2NS1lNTRhLTc0NDAtOGVjOS0wYjU5ODc1NjBiMmEiLCJ0eXBlIjoiY
WNjZXNzIiwiaWF0IjoxNzc2OTY1NTEwLCJleHAiOjE3NzY5NjY0MTB9.rhjPGfxx2ZXKdlv2IuHl',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=37)
ThinkPart(
    type='think',
    think='登录成功了！现在让我重新运行所有测试 flow。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_J93nqV1dKHxPD5G9Jcfxk5Qg',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test --debug-output=\\"/Users/ins
ulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-issu
e-55-20260424-012302\\" --flatten-debug-output flows/auth/login.yaml flows/profi
le/setup.yaml flows/matching/swipe_recommendations.yaml flows/matching/view_matc
h_list.yaml flows/devtools/open_drawer.yaml flows/devtools/edit_server_url.yaml"
, "timeout": 300}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.27886962890625,
    context_tokens=73104,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=400,
        output=161,
        input_cache_read=72704,
        input_cache_creation=0
    ),
    message_id='chatcmpl-Dbi3e8MLKdfAreY5fiqD4kKt',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_J93nqV1dKHxPD5G9Jcfxk5Qg',
    return_value=ToolReturnValue(
        is_error=True,
        output='\nWaiting for flows to complete...\n[Failed] login (18s) (Assert
ion is false: "欢迎使用 Social World" is visible)\n[Failed] setup (34s) (Asserti
on is false: "手机号一键登录" is not visible)\n[Failed] swipe_recommendations (3
5s) (Assertion is false: id: mode_tab_nearby is visible)\n[Failed] view_match_li
st (34s) (Element not found: Id matching regex: tab_messages)\n[Passed] open_dra
wer (5s)\n[Passed] edit_server_url (38s)\n\n4/6 Flows Failed\n\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=38)
ThinkPart(
    type='think',
    think='还是有4个失败。让我分析新的失败情况：\n\n1. `login` - 现在失败在 `Ass
ertion is false: "欢迎使用 Social World" is visible`。这说明 `login-as-test-user
.yaml` 的第一步 `assertVisible: "欢迎使用 Social World"` 就失败了。可能的原因是 
`clearState: true` 把应用状态清除了，但应用没有正确重启到登录页？或者应用崩溃了
？\n\n等等，`open_drawer` 也是用 `launchApp clearState: true`，但它通过了。区别
在于 `open_drawer` 没有 `runFlow: includes/login-as-test-user.yaml`。\n\n让我查
看 login flow 的截图和 commands JSON，看看第一次断言失败时的截图。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_1QHHafkhZgzqYPtPCteMqYY0',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "ls -lt /Users/insulate/Desktop/new-world/new-wor
ld-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot
-* | head -20"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_VWj1yvuJqZV5bobolrg9029x',
    function=FunctionBody(name='Shell', arguments='{"command": "cd /'),
    extras=None
)
ToolResult(
    tool_call_id='tool_1QHHafkhZgzqYPtPCteMqYY0',
    return_value=ToolReturnValue(
        is_error=False,
        output='-rw-r--r--@ 1 insulate  staff  92124 Apr 24 01:34 /Users/insulat
e/Desktop/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-issue-55
-20260424-012302/screenshot-❌-1776965640395-(view_match_list).png\n-rw-r--r--@ 
1 insulate  staff  91989 Apr 24 01:33 /Users/insulate/Desktop/new-world/new-worl
d-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-
❌-1776965606708-(swipe_recommendations).png\n-rw-r--r--@ 1 insulate  staff  922
48 Apr 24 01:32 /Users/insulate/Desktop/new-world/new-world-ops/logs/social-worl
d/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌-1776965572100-(setu
p).png\n-rw-r--r--@ 1 insulate  staff  19796 Apr 24 01:32 /Users/insulate/Deskto
p/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-issue-55-2026042
4-012302/screenshot-❌-1776965538151-(login).png\n-rw-r--r--@ 1 insulate  staff 
92233 Apr 24 01:29 /Users/insulate/Desktop/new-world/new-world-ops/logs/social-w
orld/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌-1776965351172-(v
iew_match_list).png\n-rw-r--r--@ 1 insulate  staff  92281 Apr 24 01:28 /Users/in
sulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-iss
ue-55-20260424-012302/screenshot-❌-1776965317641-(swipe_recommendations).png\n-
rw-r--r--@ 1 insulate  staff  92435 Apr 24 01:28 /Users/insulate/Desktop/new-wor
ld/new-world-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/
screenshot-❌-1776965283601-(setup).png\n-rw-r--r--@ 1 insulate  staff  92039 Ap
r 24 01:27 /Users/insulate/Desktop/new-world/new-world-ops/logs/social-world/e2e
-screenshots-e2e-issue-55-20260424-012302/screenshot-❌-1776965248733-(login).pn
g\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='Users/insulate/Desktop/new-world/new-world-ops/logs/social-w
orld/e2e-screenshots-e2e-issue-55-20260424-012302 && python3 -c \\"\\nimport jso
n\\nwith open(\'commands-(login).json\') as f:\\n    data = json.load(f)\\nfor i
, item in enumerate(data):\\n    meta = item[\'metadata\']\\n    cmd = json.dump
s(item[\'command\'], ensure_ascii=False)[:150]\\n    print(f\'{i:3d} | {meta[\\\
\\\"status\\\\\\"]:12s} | {meta[\\\\\\"duration\\\\\\"]:6d}ms | {cmd}\')\\n\\" |
head -20"}'
)
StatusUpdate(
    context_usage=0.2800788879394531,
    context_tokens=73421,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=461,
        output=362,
        input_cache_read=72960,
        input_cache_creation=0
    ),
    message_id='chatcmpl-CoGkPaVg4AwDT4QMKC7GOhEQ',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_VWj1yvuJqZV5bobolrg9029x',
    return_value=ToolReturnValue(
        is_error=False,
        output='  0 | COMPLETED    |      1ms | {"defineVariablesCommand": {"env
": {"MAESTRO_FILENAME": "login"}, "optional": false}}\n  1 | COMPLETED    |   10
41ms | {"launchAppCommand": {"appId": "com.example.mobile", "clearState": true, 
"optional": false}}\n  2 | COMPLETED    |      1ms | {"applyConfigurationCommand
": {"config": {"appId": "com.example.mobile"}, "optional": false}}\n  3 | FAILED
      |  17124ms | {"assertConditionCommand": {"condition": {"visible": {"textRe
gex": "欢迎使用 Social World", "optional": false}}, "optional": false}}\n  4 | C
OMPLETED    |      0ms | {"applyConfigurationCommand": {"config": {"appId": "com
.example.mobile"}, "optional": false}}\n  5 | FAILED       |  17216ms | {"runFlo
wCommand": {"commands": [{"applyConfigurationCommand": {"config": {"appId": "com
.example.mobile"}, "optional": false}}, {"assertConditionComma\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=39)
ThinkPart(
    type='think',
    think='`login` 的 `assertVisible: "欢迎使用 Social World"` 失败了。让我查看
失败截图，看看应用显示的是什么。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_GB7zitJfP1zDINNEpvIas52o',
    function=FunctionBody(
        name='ReadMediaFile',
        arguments='{"path": "/Users/insulate/Desktop/new-world/new-world-ops/log
s/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌-177696
5538151-(login).png"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.2848701477050781,
    context_tokens=74677,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1461,
        output=94,
        input_cache_read=73216,
        input_cache_creation=0
    ),
    message_id='chatcmpl-D9odqIqfOv1U2BR0iORZHoLl',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_GB7zitJfP1zDINNEpvIas52o',
    return_value=ToolOk(
        is_error=False,
        output=[
            TextPart(
                type='text',
                text='<image path="/Users/insulate/Desktop/new-world/new-world-o
ps/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌-
1776965538151-(login).png">'
            ),
            ImageURLPart(
                type='image_url',
                image_url=ImageURL(
                    url='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABDgAAAlgC
AIAAADieBCCAAAAAXNSR0IArs4c6QAAAANzQklUCAgI2+FP4AAAIABJREFUeJzs3VuQl/Vh//FnT8ruA
supC6yLiMIWiA3riRSoR+wQDY62TUjTSXTL9BC8SutFbad3aWZ6Y3OZpNNpxmkzYx3TqdXEtCmxasREx
AAG0fWABhZYYBfY5bDssRe///83Owh44Ad8CK/X1e/5/Z7n+T2Le7Fvv8/z/VaNjY0VAAAASaov9AUAA
ACcTKgAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoA
AAAcYQKAAAQp/ZCXwAAAPw/x44d27Vr186dO48fP75z585jx46V3m9oaJgzZ059ff2cOXNaW1sbGhou7
HWeZNu2bZ/61KcqcqrOzs6nnnrqlB/dc889bW1tFfmWi4JQSTE6Onro0KH33nuvr6+vp6en/P706dMnT
5581VVXTZkypbraCBgA8Guop6dn8+bNmzdv7uzsPN0+mzdvLr9ua2trb29vb2+fPn36ebnAM3nqqac6O
zsrFSpFUZzhH+GSIlQuvGPHjm3atOntt9/u7e3t7+8fGho6fvx4+dP6+vq6urpJkyZNmzZtwYIF119/f
dr/QgAA+MQ6OzvXr18/PkI+4lGdnZ2PP/54e3v7ypUrL+A4w1NPPfX0009fUgMd541QuZCOHz++ffv2z
Zs379ix4+DBg0NDQx/c5+jRo0VR7Nmzp66u7t133+3s7Gxvb1+0aFF9ff15v14AgIop3eN0lqMHpXGYt
ra2C3JbVKlSKnWqj7JP+Wf8zd/8zV/7OqoaGxu70NdwiTp48OCLL774i1/8oru7u5woVVWn+C8y/s26u
rqZM2ded911K1asmDp16nm9YgCASjh27Nijjz76cUdRPtTy5cu/8IUvnLd7T8ZXSltb20MPPXT25/lYV
q9efc8993yyL70onI8RlWPHjnV1dS1YsODsTzU8PPzuu+92d3cfPnx4ZGRk0qRJU6ZMufrqq6dMmfIRD
3///fe7uroOHTo0Ojra2Ng4adKkuXPnzpo1q6qq6uwv7yMaHR3ds2fPs88++9prrx06dGj8R6fsxvFvD
g0N7dq168iRI4cOHbr99ttnz57twRUA4CKyc+fOf/iHfyg/JV9BGzZs2Lx581/+5V/OmTOn4ic/SQXHU
j7xmNLTTz8tVD6h4eHht95669VXX/3Zz37W2tr6V3/1V2d5tueff/7HP/5xb2/vBz9dsmTJfffd19LSc
rrDx8bGNm7c+PTTT3d3d3/w03nz5t13330LFy48myv86Lq7u//7v/978+bNAwMDH/fY0ujKoUOHNm7cO
DQ09NnPfnb27Nnn4iIBACpuw4YNjz766Lk7/7Fjx/7u7/7ugQceWL58+bn7lgpWSuHR+dOrfKiMjY3t2
rVr8+bNGzZsOGVUfAJ9fX3f+c533n777dPtsGXLll/+8pd//Md/fNNNN33w04GBgSeeeOKFF1443eE7d
uz45je/ee+99959990VueAz6OnpKeX+J6iU8QYGBrZs2dLU1HTrrbcmzHcBAHBmjz766IYNG87PFx0/f
nzlypXn4uSVrRTOoJKh0tvbu3Xr1pdeeum9996r4GkHBgb+6Z/+aXyl1NTUzJs377LLLtu7d2+5hUZGR
r773e82NTWd9FzR2NjYSZVSU1PT2tra2Ni4Z8+egwcPlt9/8sknZ86cecMNN1Tw4k9y5MiRrVu3vvLKK
wMDA3V1dQ0NDYcPH/5YZyjdBjZp0qSBgYHjx49v3Lhx6tSpN91008SJE8/NJQMAVMCGDRvOT6WUPP744
/X19RUfV7kgldLW1nZpjrpULFT++Z//+ec//3mlzjbes88+++abb5Y377777pUrV5b+Lh8dHX399df/9
V//tdQbIyMjTzzxxMMPPzz+sY1XXnllfKXceuutd911V+kx9JGRkS1btvzLv/xL+S7JJ598sr29vaam5
lz8IEVR7N2795VXXunt7a2trb3yyisXL1788ssvn/JutDOYOXPmjTfe+MYbb7z//vu9vb2vvPLKnDlz5
s+ff46uGQDg7J3/P7U7OzsrGyoXpFIaGhra29svzVCp2HPY77zzzvjNmpqamTNnnv1ph4eHn3322fLmn
Xfeee+995ZHD6qrq6+99tovf/nL5R3ef//9Xbt2lTdHR0d/8IMflDeXLVv2pS99qTxZVk1NzfXXX/+Vr
3ylvEN3d/dJP0gFHT58ePv27b/61a+KopgxY8Ztt912880333nnnVdfffVHP8nVV19955133nrrrbfdd
tuMGTOKovjVr37V2dnZ399/ji4bAODsdXR0tLa2nreva21t7ejoqOAJL9QdX6tXrz7/Xxqi8hNGtbe3r
1279u///u8r8rzH/v37x98cdfPNN39wn8WLF4+fqHfPnj3l111dXeM3V61a9cHZvZYsWTJt2rTy5oEDB
87+sk9p9+7dr7322uDgYH19/YIFC5YsWdLU1LR06dI77rijubn5o5yhsbHxtttuW7p0aVNT05IlSxYsW
FBfXz84OPjaa6/t3r37HF02AEBFPPTQQ+fnwdrp06d/4smCT+lCVUpDQ8M5etLmolCxW7+uueaaz33uc
4sXL/6IMwWfpLe39/XXX6+qqrr22mubmprK70+YMOEP//APy5ulMYSTVFdXt7S0lJ82GT/b3d69e8uvp
02bNmvWrA8eXlNTM3fu3PKzLudoaGJ4eHj//v2lapoxY0Z7e/vll19eFEVtbe2UKVNqa2tra2tbWlpGR
0cPHTp05MiR8oETJ06cMmVKdXX17t27q6urm5qaamtri6K4/PLL29vb33vvvZ07d3Z1dXV3d19zzTWlj
wAAAjU0NKxbt+6RRx45fvz4ufuW+vr6devWVXA1lQv49PylPJxSVDBU1q5d+4mP7enp+cY3vlFagn369
OkPP/zw5MmTSx9NnTr19ttv/9AzjP91nzBhQvn1xIkT77jjjtLryZMnn26xlPEPpZT6oeL6+/v37ds3O
DhYU1PT3Nw8b96848eP19fXHzlyZOPGjXv27JkxY8Zdd901Ojq6f//+8YurTJky5Td+4zeqq6v//d///
cCBAxs3bpw5c2ZpBGn+/PnNzc27d+8+ceJEd3d3f3+/JSABgGRz5szp6Oj41re+de6+oqOjo4LrqFzAS
rnEh1OK87Pg44fatm1bqVKKoujp6XnzzTdPOcvw6Rw6dOjdd98tb86dO7f8etGiRYsWLTrz4aOjo+OnF
DvDYixn4/Dhw/v27SuKoqGhYebMmcPDw93d3W1tbX19fb/4xS8mTJjw6U9/+vrrrz/DGXbs2PHCCy9s2
rTpt3/7t5uamqqrqydMmDBr1qyGhob+/v7SPXJCBQAI197evnr16nP01//q1avb29srdbYLOxPxJT6cU
oSEykljc/X19R/92JGRkf/8z/8sb954440fdwHE119/vTyC0dTUNL5zKuj48eOlm8oaGxtnzpw5NjY2O
jpaFEVdXd2sWbMmTpy4bNmynp6eU9642dPTMzo6umLFikOHDh0+fLimpqa8Vn1zc3NjY2N/f39/f/85H
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
AAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4Q
gUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII
1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAO
EIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAi
CNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAg
DhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAA
IgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAA
IA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFA
ACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUA
ACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCB
QAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjV
AAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4Q
gUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII
1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAO
EIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAi
CNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAg
DhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAA
IgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAA
IA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFA
ACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUA
ACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCB
QAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjV
AAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4Q
gUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII
1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAO
EIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAi
CNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAg
DhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAA
IgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAA
IA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFA
ACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUA
ACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCB
QAAiCNUAACAOEIFAACII1RVDp+YAAAgAElEQVQAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII
1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAO
EIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAi
CNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAg
DhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAA
IgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAA
IA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFA
ACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUA
ACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCB
QAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjV
AAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4Q
gUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII
1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAO
EIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAi
CNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAg
DhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAA
IgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAA
IA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFA
ACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUA
ACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCB
QAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjV
AAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4Q
gUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII
1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAO
EIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAi
CNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAg
DhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAA
IgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAA
IA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFA
ACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUA
ACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCB
QAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjV
AAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4Q
gUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII
1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAO
EIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAi
CNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAg
DhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAA
IgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAA
IA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFA
ACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUAACAOEIFAACII1QAAIA4QgUAAIgjVAAAgDhCBQAAiCNUA
ACAOEIFAACII1QAAIA4QgUAAIgjVAAA/o/9OhYAAAAAGORvPY0dZRGwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
M+HfBEAAAznSURBVAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
MCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwI
yoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKA
ACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7
IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiA
gAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAA
DuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOq
AAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAAMCOqAAAADuiAgAA7IgKAACwIyoAA
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',
                    id=None
                )
            ),
            TextPart(type='text', text='</image>')
        ],
        message='Loaded image file `/Users/insulate/Desktop/new-world/new-world-
ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-❌
-1776965538151-(login).png` (image/png, 19796 bytes, original size 1080x2400px).
If you need to output coordinates, output relative coordinates first and compute
absolute coordinates using the original image size; if you generate or edit imag
es/videos via commands or scripts, read the result back immediately before conti
nuing.',
        display=[],
        extras=None
    )
)
StepBegin(n=40)
ThinkPart(
    type='think',
    think='截图显示应用是空白的白屏！这说明 `launchApp clearState: true` 后应用
启动失败了，或者正在加载中但 Maestro 没有等待足够时间。\n\n`open_drawer` 也使用
了 `launchApp clearState: true`，但它通过了。让我看看 `open_drawer` 的截图。但等
等，`open_drawer` 的测试也通过了，说明它应该能正常看到 UI。\n\n那为什么 `login` 
在 `clearState: true` 后是白屏？可能的原因是：\n1. 应用首次启动需要初始化时间（
如 Flutter 引擎预热）\n2. `clearState: true` 清除了应用的 SharedPreferences/Secu
reStorage，导致某些初始化逻辑卡住\n3. 应用有冷启动耗时\n\n等等，我注意到 `login`
的 `launchApp` 耗时只有 1041ms，而 `assertVisible` 等待了 17124ms（约17秒）。这
说明应用在启动后17秒内都没有显示 "欢迎使用 Social World"。\n\n从 logcat 看看应用
启动时是否有错误。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_q4h1o9K8T4egWrt06KQVHvXq',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "adb -s emulator-5556 logcat -d | grep -E \\"flut
ter|AndroidRuntime|FATAL|ERROR\\" | tail -30"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.2987823486328125,
    context_tokens=78324,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=3828,
        output=267,
        input_cache_read=74496,
        input_cache_creation=0
    ),
    message_id='chatcmpl-tsO69dZx0mko4ak7BN9C0i6T',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_q4h1o9K8T4egWrt06KQVHvXq',
    return_value=ToolReturnValue(
        is_error=False,
        output='04-24 01:34:38.992 13502 13527 I flutter : [DevToolsCubit] onUrl
Changed emitted: serverBaseUrl="http://192.168.1.100:3000/api/v1http://example.c
om:8080/api/v2http://10.0.2.2:300"\n04-24 01:34:38.994 13502 13502 D InputMethod
Manager: showSoftInput() view=io.flutter.embedding.android.FlutterView{3041f23 V
FED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPU
T\n04-24 01:34:39.121 13502 13527 I flutter : [DevToolsCubit] onUrlChanged calle
d: url="http://192.168.1.100:3000/api/v1http://example.com:8080/api/v2http://10.
0.2.2:3001", currentState.serverBaseUrl="http://192.168.1.100:3000/api/v1http://
example.com:8080/api/v2http://10.0.2.2:300"\n04-24 01:34:39.121 13502 13527 I fl
utter : [DevToolsCubit] onUrlChanged emitted: serverBaseUrl="http://192.168.1.10
0:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001"\n04-24 01:34:39
.123 13502 13502 D InputMethodManager: showSoftInput() view=io.flutter.embedding
.android.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824}
flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:34:39.255 13502 13527 I flutter : [DevT
oolsCubit] onUrlChanged called: url="http://192.168.1.100:3000/api/v1http://exam
ple.com:8080/api/v2http://10.0.2.2:3001/", currentState.serverBaseUrl="http://19
2.168.1.100:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001"\n04-2
4 01:34:39.255 13502 13527 I flutter : [DevToolsCubit] onUrlChanged emitted: ser
verBaseUrl="http://192.168.1.100:3000/api/v1http://example.com:8080/api/v2http:/
/10.0.2.2:3001/"\n04-24 01:34:39.257 13502 13502 D InputMethodManager: showSoftI
nput() view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F...... 
0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:34:39.
385 13502 13527 I flutter : [DevToolsCubit] onUrlChanged called: url="http://192
.168.1.100:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001/a", cur
rentState.serverBaseUrl="http://192.168.1.100:3000/api/v1http://example.com:8080
/api/v2http://10.0.2.2:3001/"\n04-24 01:34:39.385 13502 13527 I flutter : [DevTo
olsCubit] onUrlChanged emitted: serverBaseUrl="http://192.168.1.100:3000/api/v1h
ttp://example.com:8080/api/v2http://10.0.2.2:3001/a"\n04-24 01:34:39.393 13502 1
3502 D InputMethodManager: showSoftInput() view=io.flutter.embedding.android.Flu
tterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 rea
son=SHOW_SOFT_INPUT\n04-24 01:34:39.510 13502 13527 I flutter : [DevToolsCubit] 
onUrlChanged called: url="http://192.168.1.100:3000/api/v1http://example.com:808
0/api/v2http://10.0.2.2:3001/ap", currentState.serverBaseUrl="http://192.168.1.1
00:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001/a"\n04-24 01:34
:39.510 13502 13527 I flutter : [DevToolsCubit] onUrlChanged emitted: serverBase
Url="http://192.168.1.100:3000/api/v1http://example.com:8080/api/v2http://10.0.2
.2:3001/ap"\n04-24 01:34:39.511 13502 13502 D InputMethodManager: showSoftInput(
) view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F...... 0,0-1
080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:34:39.628 1
3502 13527 I flutter : [DevToolsCubit] onUrlChanged called: url="http://192.168.
1.100:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001/api", curren
tState.serverBaseUrl="http://192.168.1.100:3000/api/v1http://example.com:8080/ap
i/v2http://10.0.2.2:3001/ap"\n04-24 01:34:39.628 13502 13527 I flutter : [DevToo
lsCubit] onUrlChanged emitted: serverBaseUrl="http://192.168.1.100:3000/api/v1ht
tp://example.com:8080/api/v2http://10.0.2.2:3001/api"\n04-24 01:34:39.630 13502 
13502 D InputMethodManager: showSoftInput() view=io.flutter.embedding.android.Fl
utterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 re
ason=SHOW_SOFT_INPUT\n04-24 01:34:39.770 13502 13527 I flutter : [DevToolsCubit]
onUrlChanged called: url="http://192.168.1.100:3000/api/v1http://example.com:808
0/api/v2http://10.0.2.2:3001/api/", currentState.serverBaseUrl="http://192.168.1
.100:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001/api"\n04-24 0
1:34:39.771 13502 13527 I flutter : [DevToolsCubit] onUrlChanged emitted: server
BaseUrl="http://192.168.1.100:3000/api/v1http://example.com:8080/api/v2http://10
.0.2.2:3001/api/"\n04-24 01:34:39.776 13502 13502 D InputMethodManager: showSoft
Input() view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F......
0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:34:39.
939 13502 13527 I flutter : [DevToolsCubit] onUrlChanged called: url="http://192
.168.1.100:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001/api/v",
currentState.serverBaseUrl="http://192.168.1.100:3000/api/v1http://example.com:8
080/api/v2http://10.0.2.2:3001/api/"\n04-24 01:34:39.942 13502 13527 I flutter :
[DevToolsCubit] onUrlChanged emitted: serverBaseUrl="http://192.168.1.100:3000/a
pi/v1http://example.com:8080/api/v2http://10.0.2.2:3001/api/v"\n04-24 01:34:39.9
43 13502 13502 D InputMethodManager: showSoftInput() view=io.flutter.embedding.a
ndroid.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} f
lags=0 reason=SHOW_SOFT_INPUT\n04-24 01:34:40.073 13502 13527 I flutter : [DevTo
olsCubit] onUrlChanged called: url="http://192.168.1.100:3000/api/v1http://examp
le.com:8080/api/v2http://10.0.2.2:3001/api/v1", currentState.serverBaseUrl="http
://192.168.1.100:3000/api/v1http://example.com:8080/api/v2http://10.0.2.2:3001/a
pi/v"\n04-24 01:34:40.074 13502 13527 I flutter : [DevToolsCubit] onUrlChanged e
mitted: serverBaseUrl="http://192.168.1.100:3000/api/v1http://example.com:8080/a
pi/v2http://10.0.2.2:3001/api/v1"\n04-24 01:34:40.078 13502 13502 D InputMethodM
anager: showSoftInput() view=io.flutter.embedding.android.FlutterView{3041f23 VF
ED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT
\n04-24 01:34:42.222 13502 13527 I flutter : [DevToolsCubit] save() called with 
serverBaseUrl="http://192.168.1.100:3000/api/v1http://example.com:8080/api/v2htt
p://10.0.2.2:3001/api/v1"\n04-24 01:34:42.222 13502 13527 I flutter : [DevToolsC
ubit] save() saving url="http://192.168.1.100:3000/api/v1http://example.com:8080
/api/v2http://10.0.2.2:3001/api/v1"\n04-24 01:34:42.274 13502 13527 I flutter : 
[DevToolsCubit] save() success\n04-24 01:34:48.133 12905 12905 I AndroidRuntime:
VM exiting with result code 0.\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=41)
ThinkPart(
    type='think',
    think='logcat 显示了 `edit_server_url` 测试的一些问题：DevToolsCubit 的 URL 
变成了累积字符串 `http://192.168.1.3000/api/v1http://example.com:8080/api/v2http
://10.0.2.2:3001/api/v1`，说明 `eraseText` 没有正常工作或者 `inputText` 在 TextF
ield 全选后没有正确替换。\n\n但这可能不是当前 login 失败的原因。\n\n让我看看 log
in 测试时间点附近（01:32 左右）的 logcat。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_vtRIepgCpED3WRnkzUX2OBsw',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "adb -s emulator-5556 logcat -d | grep -E \\"01:3
2.*flutter|01:32.*AndroidRuntime\\" | head -30"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.30939483642578125,
    context_tokens=81106,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=2790,
        output=166,
        input_cache_read=78316,
        input_cache_creation=0
    ),
    message_id='chatcmpl-Y53BM5quoYJMXiRZIPXAl480',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_vtRIepgCpED3WRnkzUX2OBsw',
    return_value=ToolReturnValue(
        is_error=False,
        output='04-24 01:32:18.891 13060 13086 I ResourceExtractor: Extracted ba
seline resource assets/flutter_assets/kernel_blob.bin\n04-24 01:32:18.892 13060 
13086 I ResourceExtractor: Extracted baseline resource assets/flutter_assets/vm_
snapshot_data\n04-24 01:32:18.922 13060 13086 I ResourceExtractor: Extracted bas
eline resource assets/flutter_assets/isolate_snapshot_data\n04-24 01:32:18.975 1
3060 13098 I flutter : The Dart VM service is listening on http://127.0.0.1:4444
7/7AmirO7m4OA=/\n04-24 01:32:19.146 13060 13060 D Sentry  : \tat io.sentry.flutt
er.SentryFlutterPlugin.initNativeSdk(SentryFlutterPlugin.kt:152)\n04-24 01:32:19
.146 13060 13060 D Sentry  : \tat io.sentry.flutter.SentryFlutterPlugin.onMethod
Call(SentryFlutterPlugin.kt:86)\n04-24 01:32:19.146 13060 13060 D Sentry  : \tat
io.flutter.plugin.common.MethodChannel$IncomingMethodCallHandler.onMessage(Metho
dChannel.java:267)\n04-24 01:32:19.146 13060 13060 D Sentry  : \tat io.flutter.e
mbedding.engine.dart.DartMessenger.invokeHandler(DartMessenger.java:292)\n04-24 
01:32:19.146 13060 13060 D Sentry  : \tat io.flutter.embedding.engine.dart.DartM
essenger.lambda$dispatchMessageToQueue$0$io-flutter-embedding-engine-dart-DartMe
ssenger(DartMessenger.java:319)\n04-24 01:32:19.146 13060 13060 D Sentry  : \tat
io.flutter.embedding.engine.dart.DartMessenger$$ExternalSyntheticLambda0.run(D8$
$SyntheticClass:0)\n04-24 01:32:22.009 13060 13060 D InputMethodManager: showSof
tInput() view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F....I
D 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:2
3.237 13060 13060 D InputMethodManager: showSoftInput() view=io.flutter.embeddin
g.android.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824
} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:23.348 13060 13060 D InputMethodMa
nager: showSoftInput() view=io.flutter.embedding.android.FlutterView{3041f23 VFE
D..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\
n04-24 01:32:23.474 13060 13060 D InputMethodManager: showSoftInput() view=io.fl
utter.embedding.android.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 
aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:23.601 13060 13060 D
InputMethodManager: showSoftInput() view=io.flutter.embedding.android.FlutterVie
w{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHO
W_SOFT_INPUT\n04-24 01:32:23.730 13060 13060 D InputMethodManager: showSoftInput
() view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F...... 0,0-
1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:23.834 
13060 13060 D InputMethodManager: showSoftInput() view=io.flutter.embedding.andr
oid.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flag
s=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:23.951 13060 13060 D InputMethodManager:
showSoftInput() view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... 
.F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 
01:32:24.114 13060 13060 D InputMethodManager: showSoftInput() view=io.flutter.e
mbedding.android.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=107
3741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:24.227 13060 13060 D InputM
ethodManager: showSoftInput() view=io.flutter.embedding.android.FlutterView{3041
f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT
_INPUT\n04-24 01:32:24.352 13060 13060 D InputMethodManager: showSoftInput() vie
w=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2
274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:24.494 13060 
13060 D InputMethodManager: showSoftInput() view=io.flutter.embedding.android.Fl
utterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 re
ason=SHOW_SOFT_INPUT\n04-24 01:32:28.715 13060 13060 D InputMethodManager: showS
oftInput() view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F...
... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32
:29.412 13060 13060 D InputMethodManager: showSoftInput() view=io.flutter.embedd
ing.android.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=10737418
24} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:29.544 13060 13060 D InputMethod
Manager: showSoftInput() view=io.flutter.embedding.android.FlutterView{3041f23 V
FED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPU
T\n04-24 01:32:29.670 13060 13060 D InputMethodManager: showSoftInput() view=io.
flutter.embedding.android.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #
1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:29.826 13060 13060
D InputMethodManager: showSoftInput() view=io.flutter.embedding.android.FlutterV
iew{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} flags=0 reason=S
HOW_SOFT_INPUT\n04-24 01:32:29.961 13060 13060 D InputMethodManager: showSoftInp
ut() view=io.flutter.embedding.android.FlutterView{3041f23 VFED..... .F...... 0,
0-1080,2274 #1 aid=1073741824} flags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:30.07
9 13060 13060 D InputMethodManager: showSoftInput() view=io.flutter.embedding.an
droid.FlutterView{3041f23 VFED..... .F...... 0,0-1080,2274 #1 aid=1073741824} fl
ags=0 reason=SHOW_SOFT_INPUT\n04-24 01:32:33.701 13060 13088 I flutter : *** Req
uest ***\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=42)
ThinkPart(
    type='think',
    think='logcat 显示应用在 01:32:18 启动（Flutter VM service 启动），然后 01:3
2:33 才有第一个网络请求。这中间有大约 15 秒的空白时间。而 Maestro 的 `assertVisi
ble` 在 17 秒后超时。\n\n等等，logcat 中在 01:32:33 之前有很多 `showSoftInput` 
的日志，这说明软键盘一直在弹出，可能 Maestro 已经开始尝试交互了，但 Flutter 应用
可能还在初始化。\n\n关键是：01:32:18 应用启动，01:32:33 才有第一个网络请求。为什
么中间有 15 秒？\n\n可能是：\n1. Flutter 应用在冷启动时需要很长时间\n2. `clearSt
ate: true` 清除数据后，某些初始化需要更长时间\n3. 应用有 `AuthCheckRequested` 事
件，可能在等待网络请求\n\n让我看看 `main.dart` 是否有耗时的初始化。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_YbfE2tCKUTXwpImCIQlQvShJ',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/apps/mobile/lib/main.dart"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_XAjBDj15ny3JuKZztClEB4fr',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_YbfE2tCKUTXwpImCIQlQvShJ',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/material.dart';\n     2\timport 
'package:sentry_flutter/sentry_flutter.dart';\n     3\timport 'app.dart';\n     
4\timport 'core/di/di.dart';\n     5\t\n     6\tFuture<void> main() async {\n   
 7\t  WidgetsFlutterBinding.ensureInitialized();\n     8\t  configureDependencie
s();\n     9\t  await SentryFlutter.init(\n    10\t    (options) {\n    11\t    
 options.dsn = const String.fromEnvironment('SENTRY_DSN_MOBILE');\n    12\t     
options.tracesSampleRate = 1.0;\n    13\t    },\n    14\t    appRunner: () => ru
nApp(const SocialWorldApp()),\n    15\t  );\n    16\t}\n",
        message='16 lines read from file starting from line 1. Total lines in fi
le: 16. End of file reached.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/a
pps/mobile/lib/core/di/di.dart"}'
)
StatusUpdate(
    context_usage=0.3173637390136719,
    context_tokens=83195,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=2299,
        output=315,
        input_cache_read=80896,
        input_cache_creation=0
    ),
    message_id='chatcmpl-WrZ6zBzV5VHtV3dZoS3Ziy3Z',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_XAjBDj15ny3JuKZztClEB4fr',
    return_value=ToolOk(
        is_error=False,
        output="     1\timport 'package:flutter/foundation.dart';\n     2\timpor
t 'package:get_it/get_it.dart';\n     3\timport '../../domain/repositories/auth_
repository.dart';\n     4\timport '../../data/repositories/auth_repository_impl.
dart';\n     5\timport '../../domain/repositories/profile_repository.dart';\n   
 6\timport '../../data/repositories/profile_repository_impl.dart';\n     7\timpo
rt '../../domain/repositories/dev_tools_repository.dart';\n     8\timport '../..
/data/repositories/dev_tools_repository_impl.dart';\n     9\timport '../../domai
n/repositories/matching_repository.dart';\n    10\timport '../../data/repositori
es/matching_repository_impl.dart';\n    11\timport '../network/dio_client.dart';
\n    12\timport '../storage/secure_storage.dart';\n    13\timport '../../presen
tation/blocs/auth/auth_bloc.dart';\n    14\timport '../../presentation/blocs/pro
file/profile_bloc.dart';\n    15\timport '../../presentation/blocs/matching/matc
hing_bloc.dart';\n    16\timport '../../presentation/blocs/matching/match_list_c
ubit.dart';\n    17\timport '../../presentation/blocs/dev_tools/dev_tools_cubit.
dart';\n    18\t\n    19\tfinal getIt = GetIt.instance;\n    20\t\n    21\tvoid 
configureDependencies() {\n    22\t  getIt\n    23\t    ..registerLazySingleton<
DioClient>(DioClient.new)\n    24\t    ..registerLazySingleton<AuthRepository>(A
uthRepositoryImpl.new)\n    25\t    ..registerLazySingleton<ProfileRepository>(P
rofileRepositoryImpl.new)\n    26\t    ..registerLazySingleton<MatchingRepositor
y>(MatchingRepositoryImpl.new)\n    27\t    ..registerLazySingleton<SecureStorag
e>(SecureStorage.new)\n    28\t    ..registerFactory<AuthBloc>(\n    29\t      (
) => AuthBloc(\n    30\t        authRepository: getIt<AuthRepository>(),\n    31
\t        secureStorage: getIt<SecureStorage>(),\n    32\t      ),\n    33\t    
)\n    34\t    ..registerFactory<ProfileBloc>(\n    35\t      () => ProfileBloc(
\n    36\t        profileRepository: getIt<ProfileRepository>(),\n    37\t      
),\n    38\t    )\n    39\t    ..registerFactory<MatchingBloc>(\n    40\t      (
) => MatchingBloc(\n    41\t        matchingRepository: getIt<MatchingRepository
>(),\n    42\t      ),\n    43\t    )\n    44\t    ..registerFactory<MatchListCu
bit>(\n    45\t      () => MatchListCubit(\n    46\t        matchingRepository: 
getIt<MatchingRepository>(),\n    47\t      ),\n    48\t    );\n    49\t\n    50
\t  if (kDebugMode) {\n    51\t    getIt\n    52\t      ..registerLazySingleton<
DevToolsRepository>(\n    53\t        () => DevToolsRepositoryImpl(),\n    54\t 
    )\n    55\t      ..registerLazySingleton<DevToolsCubit>(\n    56\t        ()
=> DevToolsCubit(\n    57\t          repository: getIt<DevToolsRepository>(),\n 
  58\t          dioClient: getIt<DioClient>(),\n    59\t        ),\n    60\t    
 );\n    61\t  }\n    62\t}\n",
        message='62 lines read from file starting from line 1. Total lines in fi
le: 62. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=43)
ThinkPart(
    type='think',
    think='`main.dart` 里有 `SentryFlutter.init`，如果 `SENTRY_DSN_MOBILE` 环境
变量没有设置，Sentry 可能会尝试初始化但失败？不过通常 Sentry 在没有 DSN 时只是不
发送事件。\n\n但更关键的是：logcat 中 01:32:18 到 01:32:33 之间有 15 秒的空白。
而在 01:32:33 才有第一个网络请求。这说明应用在启动后做了一些事情。\n\n实际上，我
注意到 `open_drawer` 的测试也用了 `launchApp clearState: true`，但它通过了。让我
检查一下 `open_drawer` 的测试时间线。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_bZureqJ8QKrOObZakdKlQiDy',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /Users/insulate/Desktop/new-world/new-world-o
ps/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302 && python3 -c 
\\"\\nimport json\\nwith open(\'commands-(open_drawer).json\') as f:\\n    data 
= json.load(f)\\nfor i, item in enumerate(data):\\n    meta = item[\'metadata\']
\\n    cmd = json.dumps(item[\'command\'], ensure_ascii=False)[:150]\\n    print
(f\'{i:3d} | {meta[\\\\\\"status\\\\\\"]:12s} | {meta[\\\\\\"duration\\\\\\"]:6d
}ms | {cmd}\')\\n\\""}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.322021484375,
    context_tokens=84416,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1472,
        output=298,
        input_cache_read=82944,
        input_cache_creation=0
    ),
    message_id='chatcmpl-8x4Ktd1hCbOoiP0uq0I5VjsC',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_bZureqJ8QKrOObZakdKlQiDy',
    return_value=ToolReturnValue(
        is_error=False,
        output='  0 | COMPLETED    |      0ms | {"applyConfigurationCommand": {"
config": {"appId": "com.example.mobile"}, "optional": false}}\n  1 | COMPLETED  
 |   1544ms | {"assertConditionCommand": {"condition": {"visible": {"textRegex":
"开发工具", "optional": false}}, "optional": false}}\n  2 | COMPLETED    |   268
1ms | {"tapOnElement": {"selector": {"idRegex": "dev_tools_toggle_button", "opti
onal": false}, "retryIfNoChange": true, "waitUntilVisible": false, "longPres\n  
3 | COMPLETED    |     59ms | {"assertConditionCommand": {"condition": {"visible
": {"textRegex": "服务器地址", "optional": false}}, "optional": false}}\n  4 | C
OMPLETED    |      1ms | {"defineVariablesCommand": {"env": {"MAESTRO_FILENAME":
"open_drawer"}, "optional": false}}\n  5 | COMPLETED    |    710ms | {"launchApp
Command": {"appId": "com.example.mobile", "clearState": true, "optional": false}
}\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=44)
ThinkPart(
    type='think',
    think='`open_drawer` 的 commands JSON 也很奇怪：它的顺序是先 `assertConditio
nCommand`（步骤1）然后才是 `launchAppCommand`（步骤5）。这说明 Maestro 的 comman
ds JSON 不是按执行顺序排列的，或者 flatten 输出导致了顺序混乱。\n\n但关键是 `ope
n_drawer` 通过了，说明 `launchApp clearState: true` 后应用能正常显示。\n\n那为什
么 `login` 在同样的 `launchApp clearState: true` 后是白屏？\n\n让我重新看一下 `l
ogin` 的截图。截图是空白的，但有一个重要的事情：`open_drawer` 没有 `runFlow: inc
ludes/login-as-test-user.yaml`，而 `login` 有。\n\n等等，实际上 `login.yaml` 的
结构是：\n```yaml\n- launchApp:\n    clearState: true\n- runFlow: includes/login
-as-test-user.yaml\n```\n\n而 `runFlow` 里面的第一个命令是 `assertVisible: "欢迎
使用 Social World"`。\n\n但截图是白屏。这可能是因为：\n1. `clearState: true` 清
除了应用数据，应用重启后 Flutter 引擎需要预热\n2. 但 `launchApp` 返回太快（1秒）
，应用还没渲染完，Maestro 就开始执行 `runFlow` 里的断言了\n\n解决方案：在 `launc
hApp` 后添加一个 `wait` 或 `assertVisible` 给应用一些启动时间。\n\n但 `open_draw
er` 没有添加 wait 也通过了啊。让我看看 `open_drawer` 的截图确认一下。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_nN48AnufzuH08ytMbs0DyCjf',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "ls /Users/insulate/Desktop/new-world/new-world-o
ps/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-*op
en_drawer*"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3242149353027344,
    context_tokens=84991,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=767,
        output=385,
        input_cache_read=84224,
        input_cache_creation=0
    ),
    message_id='chatcmpl-AmJ7MDGQlyjr26l5JKpdmVFe',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_nN48AnufzuH08ytMbs0DyCjf',
    return_value=ToolReturnValue(
        is_error=True,
        output='ls: /Users/insulate/Desktop/new-world/new-world-ops/logs/social-
world/e2e-screenshots-e2e-issue-55-20260424-012302/screenshot-*open_drawer*: No 
such file or directory\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=45)
ThinkPart(
    type='think',
    think='`open_drawer` 没有截图，说明它没有失败。但我需要理解为什么 `login` 在
`launchApp clearState: true` 后是白屏。\n\n等等，让我想想。`login` 失败后，Maest
ro 会继续执行下一个 flow `setup`。`setup` 也是 `launchApp clearState: true`，然
后 `runFlow: includes/login-as-test-user.yaml`。`setup` 的失败信息是 `Assertion 
is false: "手机号一键登录" is not visible`，这说明 `setup` 的登录成功了（因为 `a
ssertNotVisible` 失败意味着页面还在登录页，但等等...）\n\n不对，`setup` 的 `asse
rtNotVisible: "手机号一键登录"` 失败，说明页面仍然显示登录页。这和第一次运行一样
。\n\n但 `login` 现在是 `assertVisible: "欢迎使用 Social World"` 失败，这说明在 
`login-as-test-user.yaml` 的第一步就失败了。\n\n让我重新思考。可能 `login` 的 `l
aunchApp clearState: true` 导致应用崩溃了，或者应用正在加载中。\n\n等等，实际上
我忽略了一个关键点：`login` 是 `auth/login.yaml`，它先 `launchApp clearState: tr
ue`，然后 `runFlow: includes/login-as-test-user.yaml`。\n\n但在 `runFlow` 的第一
步 `assertVisible: "欢迎使用 Social World"` 之前，应用可能还在启动中。\n\n`open_
drawer` 也用了 `launchApp clearState: true`，但 `open_drawer` 的第一步是 `assert
Visible: "开发工具"`，而不是 `runFlow`。\n\n让我尝试手动在 Maestro 中测试一下：`
launchApp clearState: true` 后立刻 `assertVisible: "欢迎使用 Social World"` 是否
成功。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_b5uZDbOORiTyK8l3Nif4Rynh',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /tmp/test_launch.yaml << \'EOF\'\\nappId: 
com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- assertVisible
: \\"欢迎使用 Social World\\"\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n
00000gn/T/dev-issue-55-815001100/.maestro && maestro test /tmp/test_launch.yaml"
, "timeout": 60}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3260002136230469,
    context_tokens=85459,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=467,
        output=477,
        input_cache_read=84992,
        input_cache_creation=0
    ),
    message_id='chatcmpl-OGBlZSrGwTcJqgxnouwTic1h',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_b5uZDbOORiTyK8l3Nif4Rynh',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_launch\nLaunch app "com.e
xample.mobile" with clear state... COMPLETED\nAssert that "欢迎使用 Social World
" is visible... COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=46)
ThinkPart(
    type='think',
    think='单独测试 `launchApp clearState: true` 然后 `assertVisible: "欢迎使用 
Social World"` 是成功的！\n\n那问题可能出在 **同时运行多个 flows 时的并发问题**
。Maestro 在同时运行多个 flows 时，可能一个 flow 的 `clearState` 会影响另一个 fl
ow 的状态，或者多个 flows 共享同一个模拟器实例导致冲突。\n\n实际上，Maestro 的 `
test` 命令在同时指定多个 flow 文件时，可能是**串行执行**的，但也可能是**并行执行
**的。从输出看：\n```\nWaiting for flows to complete...\n[Failed] login (18s)\n[
Failed] setup (34s)\n[Failed] swipe_recommendations (35s)\n[Failed] view_match_l
ist (34s)\n[Passed] open_drawer (5s)\n[Passed] edit_server_url (37s)\n```\n\n这
些似乎是并行运行的！因为时间几乎是同时完成的。\n\n当多个 flows 同时运行时，一个 
flow 执行 `clearState: true` 会清除应用数据，而另一个 flow 可能正在运行中，这会
导致冲突。\n\n**修复方案**：不要同时运行多个需要 `clearState: true` 的 flows，或
者使用 `maestro test` 的 `--continuous` 模式串行执行。\n\n但 Maestro 1.39.13 的 
`test` 命令默认是串行还是并行？让我查看一下帮助文档。\n\n实际上，Maestro 的 `tes
t` 命令当给定多个文件时，默认是**串行执行**在一个设备上。但如果有多个设备，可能
会并行。\n\n等等，从时间戳来看：\n- login 开始时间大约是 01:32:18（从 logcat）\n
- setup 的 logcat 时间我没有查，但从截图时间看：\n  - login 截图：01:32:18（1776
965538）\n  - setup 截图：01:32:37（1776965572）\n  - swipe 截图：01:33:06（1776
965606）\n  - view_match_list 截图：01:33:40（1776965640）\n\n等等，让我计算一下
时间戳：\n- 1776965538 = 2025-04-24 01:32:18\n- 1776965572 = 2025-04-24 01:32:52
\n- 1776965606 = 2025-04-24 01:33:26\n- 1776965640 = 2025-04-24 01:34:00\n\n不对
，这些 Unix 时间戳不对。让我重新计算：\n1776965538 / 1000 = 1776965.538 秒？不对
，文件名中的时间戳是毫秒级的。\n\n1776965538151 毫秒 = 1776965538 秒 ≈ 2025-04-2
4 01:32:18 UTC？不，macOS 的毫秒时间戳：\n1776965538151 ms = 2026-04-24 01:32:18
？让我验证。\n\n2026-04-24 01:32:18 UTC 的 Unix 时间戳大约是 1776965538。是的，
文件名时间戳是 2026-04-24。\n\n所以：\n- login 失败截图：01:32:18\n- setup 失败
截图：01:32:52\n- swipe 失败截图：01:33:26\n- view_match_list 失败截图：01:34:00
\n\n时间间隔大约是 34 秒一个。而 Maestro 输出显示：\n- login: 18s\n- setup: 34s\
n- swipe_recommendations: 35s\n- view_match_list: 34s\n\n这说明 flows 是**串行执
行**的！每个 flow 完成后才执行下一个。\n\n那为什么 `login` 在 `launchApp clearSt
ate: true` 后是白屏？让我再看一下 test_launch 的结果，它是成功的。\n\n等等，`tes
t_launch` 只运行了 `launchApp` 和 `assertVisible`，没有 `runFlow`。让我测试一下
包含 `runFlow` 的情况。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_vjPmskygEELPH8c4Lqg2T3wY',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /tmp/test_login_flow.yaml << \'EOF\'\\napp
Id: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- runFlow: 
includes/login-as-test-user.yaml\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6knv
f5n00000gn/T/dev-issue-55-815001100/.maestro && maestro test /tmp/test_login_flo
w.yaml", "timeout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3280677795410156,
    context_tokens=86001,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=753,
        output=987,
        input_cache_read=85248,
        input_cache_creation=0
    ),
    message_id='chatcmpl-f4Fqq14yuwu32H3rnfGyno3H',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_vjPmskygEELPH8c4Lqg2T3wY',
    return_value=ToolReturnValue(
        is_error=True,
        output='Failed to parse file: /tmp/test_login_flow.yaml\nFlow file does 
not exist: file:///tmp/includes/login-as-test-user.yaml\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=47)
ThinkPart(
    type='think',
    think='`runFlow` 的相对路径是相对于当前 flow 文件的位置。因为我在 `/tmp/test
_login_flow.yaml` 中引用 `includes/login-as-test-user.yaml`，Maestro 会在 `/tmp/
includes/` 下查找，找不到。\n\n让我在 `.maestro` 目录下创建测试文件。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_Qted3eFfgNzrMTHCPET2IMIV',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000
00gn/T/dev-issue-55-815001100/.maestro/test_login_flow.yaml << \'EOF\'\\nappId: 
com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- runFlow: incl
udes/login-as-test-user.yaml\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n0
0000gn/T/dev-issue-55-815001100/.maestro && maestro test test_login_flow.yaml", 
"timeout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3320655822753906,
    context_tokens=87049,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1289,
        output=220,
        input_cache_read=85760,
        input_cache_creation=0
    ),
    message_id='chatcmpl-HLdh7rWFEbxDNtqjBWd5GWPw',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_Qted3eFfgNzrMTHCPET2IMIV',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_login_flow\nLaunch app "c
om.example.mobile" with clear state...Run includes/login-as-test-user.yaml...\nR
un includes/login-as-test-user.yaml...\n  Assert that "欢迎使用 Social World" is
visible... COMPLETED\n  Assert that "手机号一键登录" is visible... COMPLETED\n  
Tap on "手机号"... COMPLETED\n  Input text 13800138000... COMPLETED\n  Tap on "
获取验证码"... COMPLETED\n  Tap on "验证码"... COMPLETED\n  Input text 123456...
COMPLETED\n  Hide Keyboard... COMPLETED\n  Tap on "登录"...Run includes/login-as
-test-user.yaml... RUNNING\nRun includes/login-as-test-user.yaml... RUNNING\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=48)
ThinkPart(
    type='think',
    think='测试卡在了 `Tap on "登录"` 这一步，没有完成也没有失败。这说明登录按钮
点击后应用可能卡住了，或者 Maestro 在等待什么。\n\n等等，`Run includes/login-as-
test-user.yaml... RUNNING` 重复出现，可能 Maestro 的输出在刷新。让我等一会儿看看
结果。\n\n实际上命令已经返回了 "Command executed successfully"，但输出显示 `RUNN
ING`。这可能是因为 `maestro test` 在等待某个操作完成，但命令本身已经超时返回了？
\n\n不，我的 timeout 是 120 秒，但输出在 60 秒内就返回了。让我检查一下是否测试实
际上还在后台运行。\n\n等等，Shell 工具的输出显示 "Command executed successfully"
，但最后的输出是 `RUNNING`。这说明测试可能实际上还在运行，但 Shell 命令已经完成
了？不，Shell 命令会等待 `maestro test` 完成。\n\n让我再运行一次，并增加更长的超
时，同时捕获完整输出。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_6NiTyQ5itai4hzehJWKei5aI',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test test_login_flow.yaml 2>&1", 
"timeout": 180}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.333587646484375,
    context_tokens=87448,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=408,
        output=274,
        input_cache_read=87040,
        input_cache_creation=0
    ),
    message_id='chatcmpl-2THPR0EQc6eYx5kg3hPylFTG',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_6NiTyQ5itai4hzehJWKei5aI',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_login_flow\nLaunch app "c
om.example.mobile" with clear state...Run includes/login-as-test-user.yaml...\nR
un includes/login-as-test-user.yaml...\n  Assert that "欢迎使用 Social World" is
visible... COMPLETED\n  Assert that "手机号一键登录" is visible... COMPLETED\n  
Tap on "手机号"... COMPLETED\n  Input text 13800138000... COMPLETED\n  Tap on "
获取验证码"... COMPLETED\n  Tap on "验证码"... COMPLETED\n  Input text 123456...
COMPLETED\n  Hide Keyboard... COMPLETED\n  Tap on "登录"...Run includes/login-as
-test-user.yaml... RUNNING\nRun includes/login-as-test-user.yaml... RUNNING\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=49)
ThinkPart(
    type='think',
    think='这很奇怪。`maestro test` 的输出在 `Tap on "登录"...` 后就停止了，但命
令成功返回了（exit code 0）。这意味着测试可能是成功的？但输出看起来不完整。\n\n
让我运行一个更简单的测试，只点击登录按钮，然后断言页面跳转。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_0JHXMeqxMv9VKATnJDvUtRVo',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000
00gn/T/dev-issue-55-815001100/.maestro/test_simple_login.yaml << \'EOF\'\\nappId
: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- assertVisib
le: \\"欢迎使用 Social World\\"\\n- tapOn: \\"手机号\\"\\n- inputText: \\"138001
38000\\"\\n- tapOn: \\"获取验证码\\"\\n- tapOn: \\"验证码\\"\\n- inputText: \\"1
23456\\"\\n- hideKeyboard\\n- tapOn: \\"登录\\"\\n- wait: 5\\n- takeScreenshot: 
login_after_click\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro && maestro test test_simple_login.yaml", "timeout"
: 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3353157043457031,
    context_tokens=87901,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=605,
        output=288,
        input_cache_read=87296,
        input_cache_creation=0
    ),
    message_id='chatcmpl-fMYZ1Qegzo2uzcFREhvAHhgx',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_0JHXMeqxMv9VKATnJDvUtRVo',
    return_value=ToolReturnValue(
        is_error=True,
        output='Failed to parse file: /private/var/folders/lt/w3sl6z8j5v50vt2s6k
nvf5n00000gn/T/dev-issue-55-815001100/.maestro/test_simple_login.yaml:13\nUnreco
gnized field "wait" (class maestro.orchestra.yaml.YamlFluentCommand), not marked
as ignorable\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=50)
ThinkPart(
    type='think',
    think='Maestro 1.39.13 不支持 `wait` 命令？这也很奇怪。让我用 `extendedWaitU
ntil` 代替。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_SlwyMsA7FWfRJ6JX7muflVAV',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000
00gn/T/dev-issue-55-815001100/.maestro/test_simple_login.yaml << \'EOF\'\\nappId
: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- assertVisib
le: \\"欢迎使用 Social World\\"\\n- tapOn: \\"手机号\\"\\n- inputText: \\"138001
38000\\"\\n- tapOn: \\"获取验证码\\"\\n- tapOn: \\"验证码\\"\\n- inputText: \\"1
23456\\"\\n- hideKeyboard\\n- tapOn: \\"登录\\"\\n- extendedWaitUntil:\\n    vis
ible: \\"完善资料\\"\\n    timeout: 10000\\n    optional: true\\n- takeScreensho
t: login_after_click\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T
/dev-issue-55-815001100/.maestro && maestro test test_simple_login.yaml", "timeo
ut": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3368377685546875,
    context_tokens=88300,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=492,
        output=278,
        input_cache_read=87808,
        input_cache_creation=0
    ),
    message_id='chatcmpl-HRRqm9iajbeqJgTp77ZP9N31',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_SlwyMsA7FWfRJ6JX7muflVAV',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_simple_login\nLaunch app 
"com.example.mobile" with clear state... COMPLETED\nAssert that "欢迎使用 Social
World" is visible... COMPLETED\nTap on "手机号"... COMPLETED\nInput text 1380013
8000... COMPLETED\nTap on "获取验证码"... COMPLETED\nTap on "验证码"... COMPLETE
D\nInput text 123456... COMPLETED\nHide Keyboard... COMPLETED\nTap on "登录"... 
COMPLETED\nAssert that (Optional) "完善资料" is visible... WARNED\n\n\n Warning:
Assertion is false: "完善资料" is visible\nTake screenshot login_after_click... 
COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=51)
ThinkPart(
    type='think',
    think='`test_simple_login` 成功运行了，但 `assertVisible: "完善资料" optiona
l: true` 返回了 WARNED，说明登录后没有跳转到 "完善资料" 页面。\n\n那登录后跳转到
了哪里？让我查看截图。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_8qBGvQAuidbN3S0cRjJEMFIP',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "find /Users/insulate/Desktop/new-world/new-world
-ops/logs/social-world/e2e-screenshots-e2e-issue-55-20260424-012302 -name \\"*lo
gin_after_click*\\" -o -name \\"*test_simple_login*\\" 2>/dev/null"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_0DRx44vGcbdUnEc7Q3diqZUF',
    function=FunctionBody(name='Shell', arguments=''),
    extras=None
)
ToolResult(
    tool_call_id='tool_8qBGvQAuidbN3S0cRjJEMFIP',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(
    arguments_part='{"command": "find /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000
00gn/T -name \\"*login_after_click*\\" -o -name \\"*test_simple_login*\\" 2>/dev
/null"}'
)
StatusUpdate(
    context_usage=0.3385162353515625,
    context_tokens=88740,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=676,
        output=199,
        input_cache_read=88064,
        input_cache_creation=0
    ),
    message_id='chatcmpl-vM8F3L33l4O0ZVdsjOI5IcCD',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_0DRx44vGcbdUnEc7Q3diqZUF',
    return_value=ToolReturnValue(
        is_error=True,
        output='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-81
5001100/.maestro/login_after_click.png\n/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n0
0000gn/T/dev-issue-55-815001100/.maestro/test_simple_login.yaml\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=52)
ThinkPart(
    type='think',
    think='截图在 `.maestro/login_after_click.png`。让我查看它。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_trQQ7LPJmG5CmFAGrMjAdgHb',
    function=FunctionBody(
        name='ReadMediaFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/login_after_click.png"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.33982086181640625,
    context_tokens=89082,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=506,
        output=75,
        input_cache_read=88576,
        input_cache_creation=0
    ),
    message_id='chatcmpl-bRXlAW7N55cPJsyHxyBu9dzC',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_trQQ7LPJmG5CmFAGrMjAdgHb',
    return_value=ToolOk(
        is_error=False,
        output=[
            TextPart(
                type='text',
                text='<image path="/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro/login_after_click.png">'
            ),
            ImageURLPart(
                type='image_url',
                image_url=ImageURL(
                    url='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABDgAAAlgC
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
+3bum2bTpdveX3ggAFDBg+KjIy0/XI/P78pL77w5VdfSZsNbNmy9dixYw89+GDXLl04RwWA8hFUAOC6n
D17Tl7M3aNHjyq+ymQ2S4Xi4tIZX7V6EKG8Y1jzZs3c3Nxq741qnNlslg52TEhImDbjfelUx3Zt295z1
91CiN//+EO6g5/07MS9e/fuP3Bwz959Ol3+Y4+OrqXTbIQQwcHBvXr13LdvvxDi8H+Hb73lFqvVog4fl
o7i6dihQ/Nmzar1Fnq9/uixY9u2bz9+/ITlda1We/PQmwb071/1T+fl5fXCc8+t+uUXafJbenrGlwu+a
hIZefuIEZ07d1LX5jgeAFwnggoAXDuj0fjr//0mlQcO6F/prCr5vtBQUjpDSZdfOhXHpdYGAQwGgxylp
A1865HYuLjc3FwhRH5BgXTl/vvuve3WWzUazZEjR6Qhiy5dOnfu1Kl1q1b5+QUnTp48dvz4e9Om33vPP
X379pF2l3Z1dW3bpo0Qoqa2O+vRrbsUVHbs3DXs5pvLj04YDAb5uJU7bh9R3fctKipe/sOP0geXBDRqd
MuwYb169byGraXd3NzGjR3btUuXn1b8nJqWJoS4GBPzxYIFISEh/fr26dSxY70bZAPQQBBUANwgpLO36
9j6DRukSTVOTk533HFHpfXlGWVJSUlNmkSaTKYzZ85IV2pvoOPEiRPSV/tCiPbt2tXSu9QGk8n0559/y
g+DAgPHjR0rLZSPiYlZuHiJEEKr1T44cqRKpdJqtc9NnvTDjz/t3LWrsLDwpxUr1qxd071b95YtWjRpE
vnKyy/V4KbMbdu28ff369O7d7du3aw2G3X4cFx8vBCia5cuLVqUXdlfKU9Pj3vvuVtaiN+iefObbx7au
VOn6s4qLKNTp07NmjVbvWbtxk2bpCuJiYmrfvl11S+/Pv3Uk7179bqexgGgNhBUANQnBQUFKpXKwcHB8
u7QbDYnJyefuXzeomtdHbt+7PjxVb/8KpXvuuOOqpz5LW8StWTZst//+KPEYJC+NQ8PC7ue7WsLCgqSU
1K0LqUcHR3VarVKpTIajQkJCT/+tEKqFhEe3qRJJasaFOXPv/46FHVYKvfu1euRUQ97eHgIIc6dP//V1
wsLCwuFEGMeezQoqHQzYmdn5/HjxrZv327lylVZ2dk6Xf7Wbduk9STeXl7+/v7u7u5ubq5arevlvx+z2
SzMZlOLFi169ewpv+/OXbuWLvum0u6t/fOvtX/+ZbtO1OHD4yY8YbvO42MeGzxoUJmL/fr2TUpK6tq1a
6uWLSuNWMXFxTmXtyq2sfjE3d199COjBg0csGfv3i1bt0k/wMGDBll+dgBQDoIKgPpk/foNq9eutV0nI
CCgDnpiNpt37Cjdf6ltmzZVXC3dtm0b+fvsjMxM+Xq/vn2vpzNOTk7ffPuttPOYDffee0/9WkLdtm3bN
Wv/NBqNj4waNfSmIVLnT5469cmnn0kDaI+OfqTMUIBKperVs2eH9u0PHDy4fceO6OiL0vXsnBzpRE6rB
vTvX5uf41o4Ojo+MmqU1acKCgrefW+qyWRydHLSqNUGozE7O1t/eXtrTy9P2y2HhoY+MHLkiOHDDx46t
G/f/vvvu/eGOQAUwA2GoAKgPhk0aKDtoHL7iOG2j7aoKSqV6qknn/Dz8zt46OCE8eOqOC2nS+fOk56du
HnzlosxMUVFRWq1Ojg4uHevXkOH3nQ9ndFoNAMHDPhpxc826jwwcmSnjh2v513qXvNmze67997QkJCOH
TtYXmzRvPnpM2eefGJC3z59rL7Q1dV10MCBgwYOzM3NjYuLi754MSnpUn5+vk6ny8nJKXNOYiN//7CwM
MsrPt4+AwfUXXTx9fWtVn1XV9f+/fuvXrOm/FM3DRniV7XW3NzcpB9Rtd4aAOqSylxisncfAKAaFi5aL
C8Nd3JycnFx8fby8vPzCwsLbdu2bcsWLWrp6+GY2NiU5GQhRHh4eOPGpWcmmkym1LQ0JZxMEp+QsHHjx
mK9Xq/X64v10pa+jo4ab28ff3+/rl26hoTUywXTZrO5/C80Ni4uIyOjIR9fGB8f/+dffwkhHBzUjo6Ob
m6uvr6+TZs0bdq0CcMjAG4YBBUA9Y/58t6+3JMBAHCjYuoXgPqHfAIAwA2vPq2qBAAAANBAEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
A5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAKA5BBQAAAIDiEFQAAAAAK
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
5eS33rn3fJLWU6eOvXmW29LZWdn5xHDhwtrOnbsMGhg6VjEun/++XrhovJZ5f9+/2PxkiVSuX27dpZ3/
7WhRYsW8gZce/bufW/aNJ2u7Cp5IcTf69Z9MHu2VI6IiBh+223l64SGhj4wcqRUPnnq1P/eettqqDt46
NCrr1/5DY5+ZLTVjt0+4nZ5Ut/sOXN37d5dpkJJScms2XOOHDkqPXz0kUfKTwKs8aYeHf2IXH7tjf+VX
/SSmZn1ymuvy6tuHhk1ymo7AACgDqinvvueHd8+NTVVCKHVam0sbLUXb2+vfv36HDoUlZt71VfLWVnZ6
9atLygobN26lZNT6Yrka0gpksDAgO7dlbhMv7i4eMHXX0vlTp06Db3JynfwZbRu1ergwUPSzV9SUtLvf
6zO0+kcHByKi4rS09NPnjy1/Mcf/vfmW+mXt8x6btKke++5x2pTjfz9PTw8tm3fLoQwGAx/rF4dExPr4
OBQXFycr8u/GHNx/YYN77z73oaNG6X6QUFBC76Y7+drfSfZ1q1bbdy8KSMjQwhx5uzZ7Tt3uDi7qFQqX
Z7u/IXzv/zy68uvviYvJX/91VeH3Ty0os/Ypk2bFT+vlDYD2LFz59lzZ11ctCaTMTMr6/iJE18s+Oqjj
z+WK3/5xfyI8LIT22zbf+CAfCM+9vHHfX18bNdXqVTt2rb76++/8/PzhRDHjh/fuGmTwWA0C3NhUVHSp
UtRUVGffj7v408+lV/y6ccft2ndymprnTt12rJtW2ZmphDizJkz//67Xq/Xm83moqKilJSU48ePL16y9
M23386/PP9q4jNPjx41yurcNq1W6+fnu37DRiGE0Wj8Y/UanU7n6OhoKDEkpyTv27d/+oz3/7g8VhYUF
PT5p5+U31C4xpsKDQlJunTpxIkTQojMzMy1f/7l4OCgUWuKiori4+M3bNr0ymuvHT9+Qqp889CbXn5pS
q3O3AMAQJlyc3OlaRq1utS2UipzianyWrVGWvzq4+NjucJBUXJycqdOff/MmbPln/L19Xn66ScGDOh3z
SlFCHHHHcOfffbp6+5mzcvJyWnfqbNUHvPYYzNnTK/Kq9IzMp565hl5ZpcNDz/00KyZ72sqPizcbDbPm
jPnq68XVtqUj4/PN0uXyCvdrTp1+vSYseMqXZp/7z33zJ09q/xsNEtr1q6d9Nzzlfbq3XfefnLChEqrl
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
AdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAA
AAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBB
QAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAe
AdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAA
AAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBB
QAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAe
AdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAA
AAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBB
QAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAeAdBBQAAAAAAe
EfMdQEAAAAAAL2c2Ww2GAxGo9FkNpnNZovFYrVYrQxDxBAJhAKBUCQUiURisVgilri4uLi6uorFff2De
l//+QEAAAAAulZ9fb1ao9FqNfX19fX1Or1eZzabHb2JWCx2c3P38HD38PDw9PTy9vLy8PDojmp5S8CYr
FzXAAAAAADgxBiGUdTVKevqlCqlSqXqQCyxh1gs9vHxkfpIpTKZr0wmEAi64yn8gaACAAAAANAROp2up
rZGLpfL5QqG6dEP1QKB0M/P18/Pr59/P3d39558dI9BUAEAAAAAcIBOp6uqrqqurlapVFzXQkTk4+MTE
BAQGBDYyxILggoAAAAAgF0qKioqKivkcjnXhbTMz88vOCg4ODiY60K6BoIKAAAAAEBb9Hp9aVlZeXmZ0
Wjkupb2ubi4hISEhoWGurm5cV1LpyCoAAAAAAC0TK3RlJQUl5eXc11IR4SEhISHR3h7eXFdSAchqAAAA
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
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAO
EIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAA
CAcoQIAAIQjVAAAgHCECgAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAADhCBUAACAcoQIAAIQjVAAAgHCEC
gAAEI5QAQAAwhEqAABAOEIFAAAIR6gAAAD/v737O27iWuA4bpj7nqUDTRpg7Qa0riAKvGO5ASQqADeA5
RSAJQpgTAPWkveAaCAoDURyA/Z90AzD5SZgg63zE/58HmGGc/bhiP3un7NxhAoAABBHqAAAAHGECgAAE
EeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAA
HGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAA
BBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAA
ABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKA
AAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoA
AAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGEC
gAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHq
AAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxh
AoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR
6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAc
YQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAE
EeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAA
HGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAA
BBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAA
ABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKA
AAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoA
AAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGEC
gAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHq
AAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxh
AoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR
6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAc
YQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAE
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
1UAAAAASUVORK5CYII=',
                    id=None
                )
            ),
            TextPart(type='text', text='</image>')
        ],
        message='Loaded image file `/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000
gn/T/dev-issue-55-815001100/.maestro/login_after_click.png` (image/png, 91898 by
tes, original size 1080x2400px). If you need to output coordinates, output relat
ive coordinates first and compute absolute coordinates using the original image 
size; if you generate or edit images/videos via commands or scripts, read the re
sult back immediately before continuing.',
        display=[],
        extras=None
    )
)
StepBegin(n=53)
ThinkPart(
    type='think',
    think='截图显示点击"登录"后，页面仍然在登录页！没有跳转。\n\n从 logcat 之前
看到的信息，前端发送了 `{phone: 13800138000, sms_code: 123456}` 到 `http://10.0.
2.2:3001/api/v1/auth/login`，而且 curl 测试显示后端返回了成功。\n\n那为什么 App 
没有跳转？可能的原因：\n1. 后端返回了成功，但 `getMe()` 调用失败\n2. `AuthBloc` 
没有正确 emit `AuthAuthenticated`\n3. `GoRouter` 没有正确监听状态变化\n\n等等，
从 logcat 看到，点击登录后没有后续的 flutter 日志输出。让我再仔细看一下 logcat。
之前我看到的网络请求日志是在之前运行的测试中。\n\n让我手动触发一次登录，并查看 l
ogcat。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_UzmxuTfyuGKNbIejyC9TbJrX',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "adb -s emulator-5556 logcat -c && cd /var/folder
s/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/.maestro && maestro
test test_simple_login.yaml 2>&1 | tail -5 &\\nsleep 25\\nadb -s emulator-5556 l
ogcat -d | grep -E \\"flutter|DioError|ERROR\\" | tail -30", "timeout": 60}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.35361480712890625,
    context_tokens=92698,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=3866,
        output=301,
        input_cache_read=88832,
        input_cache_creation=0
    ),
    message_id='chatcmpl-NaONk09J8yDbNjAsUY4VDI9O',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_UzmxuTfyuGKNbIejyC9TbJrX',
    return_value=ToolReturnValue(
        is_error=False,
        output='04-24 01:39:19.617 14674 14712 I flutter : receiveTimeout: 0:00:
10.000000\n04-24 01:39:19.617 14674 14712 I flutter : receiveDataWhenStatusError
: true\n04-24 01:39:19.617 14674 14712 I flutter : extra: {}\n04-24 01:39:19.617
14674 14712 I flutter : headers:\n04-24 01:39:19.617 14674 14712 I flutter :  co
ntent-type: application/json\n04-24 01:39:19.617 14674 14712 I flutter : data:\n
04-24 01:39:19.617 14674 14712 I flutter : {phone: 13800138000, sms_code: 123456
}\n04-24 01:39:19.617 14674 14712 I flutter : \n04-24 01:39:19.787 14674 14712 I
flutter : *** DioException ***:\n04-24 01:39:19.787 14674 14712 I flutter : uri:
http://10.0.2.2:3001/api/v1/auth/login\n04-24 01:39:19.787 14674 14712 I flutter
: DioException [bad response]: This exception was thrown because the response ha
s a status code of 400 and RequestOptions.validateStatus was configured to throw
for this status code.\n04-24 01:39:19.787 14674 14712 I flutter : The status cod
e of 400 has the following meaning: "Client error - the request contains bad syn
tax or cannot be fulfilled"\n04-24 01:39:19.787 14674 14712 I flutter : Read mor
e about status codes at https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
\n04-24 01:39:19.787 14674 14712 I flutter : In order to resolve this exception 
you typically have either to verify and fix your request code or you have to fix
the server code.\n04-24 01:39:19.788 14674 14712 I flutter : *** Response ***\n0
4-24 01:39:19.788 14674 14712 I flutter : uri: http://10.0.2.2:3001/api/v1/auth/
login\n04-24 01:39:19.788 14674 14712 I flutter : statusCode: 400\n04-24 01:39:1
9.788 14674 14712 I flutter : statusMessage: Bad Request\n04-24 01:39:19.788 146
74 14712 I flutter : headers:\n04-24 01:39:19.788 14674 14712 I flutter :  x-pow
ered-by: Express\n04-24 01:39:19.788 14674 14712 I flutter :  connection: keep-a
live\n04-24 01:39:19.788 14674 14712 I flutter :  keep-alive: timeout=5\n04-24 0
1:39:19.788 14674 14712 I flutter :  date: Thu, 23 Apr 2026 17:39:20 GMT\n04-24 
01:39:19.788 14674 14712 I flutter :  content-length: 65\n04-24 01:39:19.788 146
74 14712 I flutter :  etag: W/"41-fw+mOf61n6usbocCK5t6YiASr2Y"\n04-24 01:39:19.7
88 14674 14712 I flutter :  content-type: application/json; charset=utf-8\n04-24
01:39:19.788 14674 14712 I flutter : Response Text:\n04-24 01:39:19.788 14674 14
712 I flutter : {"code":1001,"message":"Invalid or expired SMS code","data":null
}\n04-24 01:39:19.789 14674 14712 I flutter : \n04-24 01:39:19.789 14674 14712 I
flutter : \nAssert that (Optional) "完善资料" is visible... WARNED\n\n\n Warning
: Assertion is false: "完善资料" is visible\nTake screenshot login_after_click..
. COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=54)
ThinkPart(
    type='think',
    think='Aha! 问题找到了！\n\nApp 发送请求到 `http://10.0.2.2:3001/api/v1/auth
/login`，但后端返回了 `{"code":1001,"message":"Invalid or expired SMS code","dat
a":null}`。\n\n这说明 Redis 中的验证码在 App 测试时已经过期了！我之前设置的 `EX 
300`（5分钟）可能已经过了。\n\n让我重新设置 Redis 验证码，并设置更长的过期时间，
比如 1 小时。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_gfwcXGzOdvplwBQyWmYXvynu',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "docker exec sw-redis redis-cli SET sms_code:1380
0138000 123456 EX 3600"}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_Kyo1LfP0Tf9TduocNa1LiQsy',
    function=FunctionBody(name='Shell', arguments='{"command": "docker exec'),
    extras=None
)
ToolResult(
    tool_call_id='tool_gfwcXGzOdvplwBQyWmYXvynu',
    return_value=ToolReturnValue(
        is_error=False,
        output='OK\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
ToolCallPart(arguments_part=' sw-redis redis-cli GET sms_code:13800138000"}')
StatusUpdate(
    context_usage=0.3585357666015625,
    context_tokens=93988,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1298,
        output=170,
        input_cache_read=92690,
        input_cache_creation=0
    ),
    message_id='chatcmpl-CsHZK2GkB9dwDMIDpsolA1SF',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_Kyo1LfP0Tf9TduocNa1LiQsy',
    return_value=ToolReturnValue(
        is_error=False,
        output='123456\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=55)
ThinkPart(
    type='think',
    think='验证码已重新设置，过期时间 1 小时。现在让我重新运行测试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_Vbu5DX1076kOcAyX1qNXQ4XO',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test test_simple_login.yaml", "ti
meout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3593864440917969,
    context_tokens=94211,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=259,
        output=90,
        input_cache_read=93952,
        input_cache_creation=0
    ),
    message_id='chatcmpl-IcjQDtjYLImgP3S5GizBIu8W',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_Vbu5DX1076kOcAyX1qNXQ4XO',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_simple_login\nLaunch app 
"com.example.mobile" with clear state... COMPLETED\nAssert that "欢迎使用 Social
World" is visible... COMPLETED\nTap on "手机号"... COMPLETED\nInput text 1380013
8000... COMPLETED\nTap on "获取验证码"... COMPLETED\nTap on "验证码"... COMPLETE
D\nInput text 123456... COMPLETED\nHide Keyboard... COMPLETED\nTap on "登录"... 
COMPLETED\nAssert that (Optional) "完善资料" is visible... COMPLETED\nTake scree
nshot login_after_click... COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=56)
ThinkPart(
    type='think',
    think='太好了！`test_simple_login` 通过了！`assertVisible: "完善资料" option
al: true` 也 COMPLETED 了，说明登录成功后跳转到了完善资料页面。\n\n现在让我运行
完整的 `login.yaml` 和 `setup.yaml`。但等等，`login.yaml` 中的 `assertVisible: "
完善资料" optional: true` 在我之前的修复后已经是 mapping 形式了。\n\n让我运行 `l
ogin.yaml`。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_YDVH2FHQXw2FaCbUkcqqNUKT',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test flows/auth/login.yaml", "tim
eout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.36029815673828125,
    context_tokens=94450,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=242,
        output=161,
        input_cache_read=94208,
        input_cache_creation=0
    ),
    message_id='chatcmpl-4t2ARgERnNeZ4CoXGcl1bM6a',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_YDVH2FHQXw2FaCbUkcqqNUKT',
    return_value=ToolReturnValue(
        is_error=True,
        output='Running on emulator-5556\n > Flow login\nLaunch app "com.example
.mobile" with clear state...Run includes/login-as-test-user.yaml...\nRun include
s/login-as-test-user.yaml...\n  Assert that "欢迎使用 Social World" is visible..
. COMPLETED\n  Assert that "手机号一键登录" is visible... COMPLETED\n  Tap on "
手机号"... COMPLETED\n  Input text 13800138000... COMPLETED\n  Tap on "获取验证
码"... COMPLETED\n  Tap on "验证码"... COMPLETED\n  Input text 123456... COMPLET
ED\n  Hide Keyboard... COMPLETED\n  Tap on "登录"...Run includes/login-as-test-u
ser.yaml... RUNNING\nRun includes/login-as-test-user.yaml... RUNNING\nAssert tha
t "手机号一键登录" is not visible... FAILED\n\nAssertion is false: "手机号一键登
录" is not visible\n\n==== Debug output (logs & screenshots) ====\n\n/Users/insu
late/.maestro/tests/2026-04-24_014022\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=57)
ThinkPart(
    type='think',
    think='`login.yaml` 仍然失败了！错误是 `Assertion is false: "手机号一键登录"
is not visible`，说明点击登录后页面仍然在登录页。\n\n但 `test_simple_login.yaml`
刚刚成功了！区别在于：\n- `test_simple_login.yaml` 没有使用 `runFlow`，步骤是直
接在文件里的\n- `login.yaml` 使用了 `runFlow: includes/login-as-test-user.yaml`\
n\n而且 output 很奇怪：`Tap on "登录"...Run includes/login-as-test-user.yaml... 
RUNNING` 重复出现。这可能意味着 `runFlow` 内部的步骤在 Maestro 的输出中有某种问
题。\n\n等等，让我重新看一下 `login.yaml` 的内容。我之前修改了 `assertVisible` 
的格式，但没有改其他部分。\n\n```yaml\nappId: com.example.mobile\n---\n- launchA
pp:\n    clearState: true\n- runFlow: includes/login-as-test-user.yaml\n\n# 断言
：登录成功后离开登录页（新用户进入完善资料，老用户进入首页）\n- assertNotVisible
: "手机号一键登录"\n- assertVisible:\n    text: "完善资料"\n    optional: true\n
- assertVisible:\n    id: "tab_home"\n    optional: true\n```\n\n`test_simple_lo
gin.yaml` 的内容是：\n```yaml\n- launchApp:\n    clearState: true\n- assertVisib
le: "欢迎使用 Social World"\n- tapOn: "手机号"\n- inputText: "13800138000"\n- ta
pOn: "获取验证码"\n- tapOn: "验证码"\n- inputText: "123456"\n- hideKeyboard\n- t
apOn: "登录"\n- extendedWaitUntil:\n    visible: "完善资料"\n    timeout: 10000\
n    optional: true\n- takeScreenshot: login_after_click\n```\n\n`test_simple_lo
gin.yaml` 有一个 `extendedWaitUntil` 来等待页面跳转，而 `login.yaml` 在 `runFlow
` 后立刻执行 `assertNotVisible`。\n\n但 `login.yaml` 的 `assertNotVisible` 等待
了 10 秒（默认 extendedWaitUntil timeout），应该足够页面跳转了。\n\n等等，可能问
题在于 `runFlow` 的语法在 Maestro 1.39.13 中有 bug？或者 `runFlow` 后的步骤执行
时机有问题？\n\n让我测试一下：在 `runFlow` 后添加一个 `extendedWaitUntil` 来等待
跳转。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_u89gETLuNAEIbmx64DHRnB58',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000
00gn/T/dev-issue-55-815001100/.maestro/test_runflow_login.yaml << \'EOF\'\\nappI
d: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- runFlow: i
ncludes/login-as-test-user.yaml\\n- extendedWaitUntil:\\n    notVisible: \\"手机
号一键登录\\"\\n    timeout: 15000\\nEOF\\ncd /var/folders/lt/w3sl6z8j5v50vt2s6k
nvf5n00000gn/T/dev-issue-55-815001100/.maestro && maestro test test_runflow_logi
n.yaml", "timeout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.36182403564453125,
    context_tokens=94850,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=642,
        output=680,
        input_cache_read=94208,
        input_cache_creation=0
    ),
    message_id='chatcmpl-VU8tuoVZsTcDPnh3S3q3ZDlf',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_u89gETLuNAEIbmx64DHRnB58',
    return_value=ToolReturnValue(
        is_error=True,
        output='Running on emulator-5556\n > Flow test_runflow_login\nLaunch app
"com.example.mobile" with clear state...Run includes/login-as-test-user.yaml...\
nRun includes/login-as-test-user.yaml...\n  Assert that "欢迎使用 Social World" 
is visible... COMPLETED\n  Assert that "手机号一键登录" is visible... COMPLETED\
n  Tap on "手机号"... COMPLETED\n  Input text 13800138000... COMPLETED\n  Tap on
"获取验证码"... COMPLETED\n  Tap on "验证码"... COMPLETED\n  Input text 123456..
. COMPLETED\n  Hide Keyboard... COMPLETED\n  Tap on "登录"...Run includes/login-
as-test-user.yaml... RUNNING\nRun includes/login-as-test-user.yaml... RUNNING\nA
ssert that "手机号一键登录" is not visible... FAILED\n\nAssertion is false: "手
机号一键登录" is not visible\n\n==== Debug output (logs & screenshots) ====\n\n/
Users/insulate/.maestro/tests/2026-04-24_014122\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=58)
ThinkPart(
    type='think',
    think='这很奇怪！`test_simple_login.yaml`（没有 `runFlow`）成功了，但 `test_
runflow_login.yaml`（有 `runFlow`）失败了。而且两者的步骤完全一样！\n\n这说明 Ma
estro 的 `runFlow` 命令在这个版本中有 bug，可能导致包含 `runFlow` 的 flow 在 `ru
nFlow` 内部执行时出现问题。\n\n等等，让我再仔细看输出：\n```\nTap on "登录"...Ru
n includes/login-as-test-user.yaml... RUNNING\nRun includes/login-as-test-user.y
aml... RUNNING\nAssert that "手机号一键登录" is not visible... FAILED\n```\n\n在
`test_simple_login.yaml` 中，输出是：\n```\nTap on "登录"... COMPLETED\nAssert t
hat (Optional) "完善资料" is visible... COMPLETED\n```\n\n但在 `test_runflow_log
in.yaml` 中，点击登录后显示 `RUNNING`，然后立刻跳到 `Assert that "手机号一键登录
" is not visible... FAILED`。\n\n这似乎说明 `runFlow` 内部的 `tapOn "登录"` 没有
正确完成，Maestro 可能提前退出了 `runFlow` 的执行。\n\n让我检查 `login-as-test-u
ser.yaml` 的内容。它前面没有 `appId` 行吗？不，它有：\n```yaml\nappId: com.examp
le.mobile\n---\n```\n\n等等！`login-as-test-user.yaml` 的第一行是注释，但第二行
开始有 `appId: com.example.mobile` 和 `---`。这在 Maestro 的 include 文件中可能
是错误的！\n\n根据 Maestro 文档，include 文件**不应该**有 `appId` 和 `---` 分隔
符。它们应该只是步骤列表。\n\n让我检查 Maestro 文档中关于 `runFlow` 的 include 
文件格式。\n\n实际上，在 Maestro 中，被 `runFlow` 引用的文件可以有 `appId` 和 `-
--`，但通常不推荐。但如果说这是问题所在，那去掉它们试试。\n\n但等等，`open_drawe
r.yaml` 没有使用 `runFlow`，`edit_server_url.yaml` 也没有使用 `runFlow`（虽然它
依赖 `devtools/open_drawer`，但实际文件里没有 `runFlow`）。\n\n让我验证一下：去
掉 `login-as-test-user.yaml` 中的 `appId` 和 `---`。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_HnuLYLrpuiHy2aa9XbSG5MaA',
    function=FunctionBody(
        name='ReadFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/includes/login-as-test-user.yaml"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3653411865234375,
    context_tokens=95772,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=1052,
        output=537,
        input_cache_read=94720,
        input_cache_creation=0
    ),
    message_id='chatcmpl-m26UGRIcIYUdi4ptS48Rfpb3',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_HnuLYLrpuiHy2aa9XbSG5MaA',
    return_value=ToolOk(
        is_error=False,
        output='     1\t# @include\n     2\t# @name: login-as-test-user\n     3\
t# @displayName: 使用测试用户登录\n     4\t# @description: 输入测试手机号和固定
验证码，完成登录前置步骤\n     5\t# @lastUpdated: 2025-04-16\n     6\t\n     7\t
appId: com.example.mobile\n     8\t---\n     9\t\n    10\t- assertVisible: "欢迎
使用 Social World"\n    11\t- assertVisible: "手机号一键登录"\n    12\t\n    13\
t# 输入手机号\n    14\t- tapOn: "手机号"\n    15\t- inputText: "13800138000"\n  
 16\t\n    17\t# 获取验证码\n    18\t- tapOn: "获取验证码"\n    19\t# 注：MVP 阶
段 _onSendCode 只启动本地倒计时，不调用后端 API\n    20\t# 原 assertVisible: "60
s" 因设备性能波动导致秒数变化而失败，此处不再严格断言\n    21\t\n    22\t# 输入
验证码（MVP 固定验证码为任意 6 位）\n    23\t- tapOn: "验证码"\n    24\t- inputT
ext: "123456"\n    25\t\n    26\t# 收起软键盘，避免遮挡登录按钮\n    27\t- hideK
eyboard\n    28\t\n    29\t# 点击登录\n    30\t- tapOn: "登录"\n',
        message='30 lines read from file starting from line 1. Total lines in fi
le: 30. End of file reached.',
        display=[],
        extras=None
    )
)
StepBegin(n=59)
ThinkPart(
    type='think',
    think='让我尝试去掉 `login-as-test-user.yaml` 中的 `appId` 和 `---`，看看是
否是格式问题导致 `runFlow` 异常。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_fFSRicJ598EKk6q6HQ0IbPOm',
    function=FunctionBody(
        name='StrReplaceFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/includes/login-as-test-user.yaml", "edit": {"old":
"appId: com.example.mobile\\n---\\n\\n- assertVisible: \\"欢迎使用 Social World\
\"","new":"- assertVisible: \\"欢迎使用 Social World\\""}}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.36887359619140625,
    context_tokens=96698,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=954,
        output=135,
        input_cache_read=95744,
        input_cache_creation=0
    ),
    message_id='chatcmpl-hjkAUcBEP19TGEJgjK11KHeG',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_fFSRicJ598EKk6q6HQ0IbPOm',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/includes/login-as-test-user.yaml',
                old_text='# @displayName: 使用测试用户登录\n# @description: 输入
测试手机号和固定验证码，完成登录前置步骤\n# @lastUpdated: 2025-04-16\n\nappId: c
om.example.mobile\n---\n\n- assertVisible: "欢迎使用 Social World"\n- assertVisi
ble: "手机号一键登录"',
                new_text='# @displayName: 使用测试用户登录\n# @description: 输入
测试手机号和固定验证码，完成登录前置步骤\n# @lastUpdated: 2025-04-16\n\n- assert
Visible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"',
                old_start=3,
                new_start=3,
                is_summary=False
            )
        ],
        extras=None
    )
)
StepBegin(n=60)
ThinkPart(
    type='think',
    think='让我重新运行 `test_runflow_login.yaml` 测试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_YTNJeZfOSVtIuI7MmiE1ejX0',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test test_runflow_login.yaml", "t
imeout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3695411682128906,
    context_tokens=96873,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=361,
        output=86,
        input_cache_read=96512,
        input_cache_creation=0
    ),
    message_id='chatcmpl-p9Fj0r3Tlfd0UA21lR84TOxH',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_YTNJeZfOSVtIuI7MmiE1ejX0',
    return_value=ToolReturnValue(
        is_error=True,
        output='Failed to parse file: /private/var/folders/lt/w3sl6z8j5v50vt2s6k
nvf5n00000gn/T/dev-issue-55-815001100/.maestro/test_runflow_login.yaml\nFailed t
o parse file: /private/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/includes/login-as-test-user.yaml\nFlow files must contain
a config section and a commands section separated by "---". For example:\n\nappI
d: com.example\n---\n- launchApp\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=61)
ThinkPart(
    type='think',
    think='Maestro 要求所有 flow 文件（包括 includes）都必须包含 `appId` 和 `---
`。那我之前的格式是正确的。\n\n问题可能在于 `runFlow` 在 Maestro 1.39.13 中有 bu
g，或者与 `tapOn` 的交互有问题。\n\n让我尝试另一种方法：不用 `runFlow`，而是把登
录步骤内联到每个 flow 中。但这会重复代码。\n\n或者，我可以尝试用 `runScript` 或
其他方式。\n\n等等，让我再想想。`test_simple_login.yaml` 成功了，它的步骤和 `log
in-as-test-user.yaml` 完全一样，只是没有 `runFlow`。那如果我把 `login.yaml` 中的
`runFlow` 替换为内联步骤呢？\n\n但这会破坏复用性。不过如果 `runFlow` 在这个版本
中有 bug，那可能只能内联。\n\n让我再做一个实验：在 `login-as-test-user.yaml` 末
尾添加一个 `wait` 或 `assertVisible` 来看看 `runFlow` 是否能完整执行。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_3GkAgSuYtTCRFC3NgR9poTHy',
    function=FunctionBody(
        name='StrReplaceFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/includes/login-as-test-user.yaml", "edit": {"old":
"# @lastUpdated: 2025-04-16\\n\\n- assertVisible: \\"欢迎使用 Social World\\"","
new":"# @lastUpdated: 2025-04-16\\n\\nappId: com.example.mobile\\n---\\n\\n- ass
ertVisible: \\"欢迎使用 Social World\\""}}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.37050628662109375,
    context_tokens=97126,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=358,
        output=346,
        input_cache_read=96768,
        input_cache_creation=0
    ),
    message_id='chatcmpl-nW3XLPOW1p4uxg2vpwClFmG1',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_3GkAgSuYtTCRFC3NgR9poTHy',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/includes/login-as-test-user.yaml',
                old_text='# @displayName: 使用测试用户登录\n# @description: 输入
测试手机号和固定验证码，完成登录前置步骤\n# @lastUpdated: 2025-04-16\n\n- assert
Visible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"',
                new_text='# @displayName: 使用测试用户登录\n# @description: 输入
测试手机号和固定验证码，完成登录前置步骤\n# @lastUpdated: 2025-04-16\n\nappId: c
om.example.mobile\n---\n\n- assertVisible: "欢迎使用 Social World"\n- assertVisi
ble: "手机号一键登录"',
                old_start=3,
                new_start=3,
                is_summary=False
            )
        ],
        extras=None
    )
)
StepBegin(n=62)
ThinkPart(
    type='think',
    think='让我尝试在 `login-as-test-user.yaml` 末尾添加一个 `extendedWaitUntil`
来等待页面跳转，看看 `runFlow` 是否能完整执行。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_tUbZEmPom98cV9fQAxGNa3mH',
    function=FunctionBody(
        name='StrReplaceFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/includes/login-as-test-user.yaml", "edit": {"old":
"# 点击登录\\n- tapOn: \\"登录\\"","new":"# 点击登录\\n- tapOn: \\"登录\\"\\n- e
xtendedWaitUntil:\\n    visible: \\"完善资料\\"\\n    timeout: 15000\\n    optio
nal: true"}}'
    ),
    extras=None
)
ToolResult(
    tool_call_id='tool_tUbZEmPom98cV9fQAxGNa3mH',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/includes/login-as-test-user.yaml',
                old_text='\n# 点击登录\n- tapOn: "登录"',
                new_text='\n# 点击登录\n- tapOn: "登录"\n- extendedWaitUntil:\n 
  visible: "完善资料"\n    timeout: 15000\n    optional: true',
                old_start=28,
                new_start=28,
                is_summary=False
            )
        ],
        extras=None
    )
)
StatusUpdate(
    context_usage=0.371978759765625,
    context_tokens=97512,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=488,
        output=158,
        input_cache_read=97024,
        input_cache_creation=0
    ),
    message_id='chatcmpl-OVAsWG2XXCxhv0vCAQLotIKl',
    plan_mode=False,
    mcp_status=None
)
StepBegin(n=63)
ThinkPart(
    type='think',
    think='让我重新运行 `test_runflow_login.yaml`。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_wucFNBH5c7ns6hY82pgeGFxi',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test test_runflow_login.yaml", "t
imeout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.37273406982421875,
    context_tokens=97710,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=430,
        output=83,
        input_cache_read=97280,
        input_cache_creation=0
    ),
    message_id='chatcmpl-GLCxleuoVS5Ed1lHU2WoxdDe',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_wucFNBH5c7ns6hY82pgeGFxi',
    return_value=ToolReturnValue(
        is_error=True,
        output='Running on emulator-5556\n > Flow test_runflow_login\nLaunch app
"com.example.mobile" with clear state...Run includes/login-as-test-user.yaml...\
nRun includes/login-as-test-user.yaml...\n  Assert that "欢迎使用 Social World" 
is visible... COMPLETED\n  Assert that "手机号一键登录" is visible... COMPLETED\
n  Tap on "手机号"... COMPLETED\n  Input text 13800138000... COMPLETED\n  Tap on
"获取验证码"... COMPLETED\n  Tap on "验证码"... COMPLETED\n  Input text 123456..
. COMPLETED\n  Hide Keyboard... COMPLETED\n  Tap on "登录"... COMPLETED\n  Asser
t that (Optional) "完善资料" is visible...Run includes/login-as-test-user.yaml..
. RUNNING\nRun includes/login-as-test-user.yaml... RUNNING\nAssert that "手机号
一键登录" is not visible... FAILED\n\nAssertion is false: "手机号一键登录" is no
t visible\n\n==== Debug output (logs & screenshots) ====\n\n/Users/insulate/.mae
stro/tests/2026-04-24_014252\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=64)
ThinkPart(
    type='think',
    think='这很奇怪！`runFlow` 内部的 `extendedWaitUntil` 显示 `RUNNING`，但然后
外层 flow 的 `assertNotVisible` 立即执行并失败了。\n\n这说明 Maestro 1.39.13 的 
`runFlow` 有严重 bug：`runFlow` 内部的异步等待命令（如 `extendedWaitUntil`）似乎
没有阻塞外层 flow 的执行。外层的 `assertNotVisible` 在 `runFlow` 还没完成时就开
始执行了。\n\n实际上，仔细看输出：\n```\nTap on "登录"... COMPLETED\nAssert that
(Optional) "完善资料" is visible...Run includes/login-as-test-user.yaml... RUNNI
NG\nRun includes/login-as-test-user.yaml... RUNNING\nAssert that "手机号一键登录
" is not visible... FAILED\n```\n\n`extendedWaitUntil` 的输出被截断了，显示为 `R
UNNING`。然后外层命令立即执行。\n\n这可能意味着 `runFlow` 确实没有正确等待内部命
令完成。这是一个 **Maestro 版本 bug** 或 **runFlow 使用方式问题**。\n\n让我查看 
Maestro 1.39.13 的 release notes 或已知问题。\n\n一个可能的解决方法是：不使用 `r
unFlow`，而是将登录步骤直接复制到需要登录的 flow 中。但这很丑陋。\n\n另一个方法
是：使用 `runScript` 来运行子 flow，但这更复杂。\n\n或者，我可以用 `evalScript` 
或 `repeat` 等命令。\n\n等等，也许 `runFlow` 的问题只出现在 `optional: true` 或 
`extendedWaitUntil` 上？让我测试一个更简单的 `runFlow`，只包含 `tapOn` 和 `asser
tVisible`，没有 `extendedWaitUntil`。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_1s9v4L2rZGLTDmb7G6Oi78ps',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cat > /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n000
00gn/T/dev-issue-55-815001100/.maestro/test_simple_runflow.yaml << \'EOF\'\\napp
Id: com.example.mobile\\n---\\n- launchApp:\\n    clearState: true\\n- runFlow: 
includes/login-as-test-user.yaml\\n- takeScreenshot: after_runflow\\nEOF\\ncd /v
ar/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55-815001100/.maestro &
& maestro test test_simple_runflow.yaml", "timeout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3740348815917969,
    context_tokens=98051,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=515,
        output=523,
        input_cache_read=97536,
        input_cache_creation=0
    ),
    message_id='chatcmpl-MQGCxIignEo6lWUW8bKYWqYE',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_1s9v4L2rZGLTDmb7G6Oi78ps',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_simple_runflow\nLaunch ap
p "com.example.mobile" with clear state...Run includes/login-as-test-user.yaml..
.\nRun includes/login-as-test-user.yaml...\n  Assert that "欢迎使用 Social World
" is visible... COMPLETED\n  Assert that "手机号一键登录" is visible... COMPLETE
D\n  Tap on "手机号"... COMPLETED\n  Input text 13800138000... COMPLETED\n  Tap 
on "获取验证码"... COMPLETED\n  Tap on "验证码"... COMPLETED\n  Input text 12345
6... COMPLETED\n  Hide Keyboard... COMPLETED\n  Tap on "登录"... COMPLETED\n  As
sert that (Optional) "完善资料" is visible...Run includes/login-as-test-user.yam
l... RUNNING\nRun includes/login-as-test-user.yaml... RUNNING\nTake screenshot a
fter_runflow... COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=65)
ThinkPart(
    type='think',
    think='`test_simple_runflow.yaml` 通过了！而且 `takeScreenshot after_runflow
` 也 COMPLETED。这说明 `runFlow` 实际上可能完成了，只是输出格式有问题（`RUNNING`
重复显示）。\n\n但 `test_runflow_login.yaml` 失败了，它的区别是在 `runFlow` 后有
一个 `extendedWaitUntil`（或 `assertNotVisible`）。\n\n等等，`test_simple_runflo
w.yaml` 在 `runFlow` 后只有 `takeScreenshot`，没有等待。它通过了。\n\n让我查看 `
after_runflow` 截图，看看 `runFlow` 完成后页面是什么状态。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_MlCMrO1xnMdVdl57aMD9S7I4',
    function=FunctionBody(
        name='ReadMediaFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/after_runflow.png"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.3768119812011719,
    context_tokens=98779,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=731,
        output=190,
        input_cache_read=98048,
        input_cache_creation=0
    ),
    message_id='chatcmpl-mMOQ9pFaYgVkKXQiWCvDHVNa',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_MlCMrO1xnMdVdl57aMD9S7I4',
    return_value=ToolOk(
        is_error=False,
        output=[
            TextPart(
                type='text',
                text='<image path="/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro/after_runflow.png">'
            ),
            ImageURLPart(
                type='image_url',
                image_url=ImageURL(
                    url='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABDgAAAlgC
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
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAA
AAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2lI7uwE3CbDYrFIqq1DQajVlZ2
UKIgAB/Z2fnWu4XAAAAUCcRVK6d0Wi8cCHp1KnTBw/GnD595uOPZwYHB1X6qqys7DFjnhZCLFu2MCQku
Pa7CQAAANQ9BJWq0uv1GRmZqakXk5NTkpKSExMTz59PLC4utlT44INZH3/8oUqluua30Gg0Fy+m2akQH
h6uVlfS/uzZc7Zt23FtHSA7AQAAQCYIKqW0Wm1eXn5eXn5+fn5eXl5ubl5eXl5OTm5OTm52dnZ2drZWW
1jRa6Ojozp0aN+5c8cqzv6qyIkTpyZPnm6nwhdffNKsWdPreQsAAACgTiColJo7d8HOnbsrraZQKBo2D
ImMjIyICAsLCw0LC4uOjvTw8LgBPayWXr16PPfc+CpWPnz4yJw5c2u1PwAAAEC1EFRKdenS6cCBQ56eH
u7u7h4e7u7u7l5eXr6+Pj4+Pt7e3t7eXp9//qVWqx07dvTw4ffXak/c3Nx++22V9ZWcnJxRo8ZWt5GAg
IAqVvbx8alW4wAAAEBtI6iUuv32226//TY7FRYsWKTVaq9zchcAAACAquAcFQAAAACyw4hKrdu5c/fMm
R+Xvz5u3DOWcsuWLebM+agG37SkpESj0VSxclFRUQ2+NQAAAHD9CCo3p9279+zevcfRvQAAAACuEUGl1
Isvvnb27LlKqy1d+u3Spd/aqdCnT6933nnT+kq7dm0++miG5eH69X9L55xMmvSqZb17paejAAAAAPUKQ
aXW+fr6+vr6Wh4uW/adVGjZskVtnK44adKrkya9WuPNAgAAADfSTRJUTCZTYWGhwWAwGAxGo/Hy/xtdX
V3Dw8Oq0sILL0woLtbZqTB9+kyNRjN06B0DBvS3U83Ly9POs9nZOadOna5Kf+qQixcvssoFAADgxlMIh
bPS2dnZWalUSv+vVCrd3d2dnG6GHbPqdlApKSnJy8srKCjQarVms7l8BaOxqlOqKj3x3cXFRQjRsGHDd
u3aVLefFocOHbaUT506XWZERfoITk51bAfk4uJirVbr6F4AAACglKenp7e3t7e3t3QHW0fVyaBiNBozM
zPz8vJ0OntjIDK0des2S3nOnLkNGgS2bt3KckX6OCqVurrNrly5evnyH2ukh5KwsNDFixfUYIMAAAC4Y
TQajUajSU1NVbmpvH18GjQIrItjLHUvqGRkZKSnp5tMJuuLbq5uXl5eLq4uyqs5qpM2paSkHjlyNDS0Y
WrqRSFESUnJtGkfzpkzOzS0oVShpMQg6uDC+kaNGjm6CwAAAPWU4Wp6vb6gQKPXl36bX6wrLk4vzsrKD
AoKCgwMdGxXq0tet/L25ebmpqVdKinRSw+dnJw8PT09Pb28vb3qxKjW5s1bhRBDh965ZMk3Qoh+/frs3
Ll7ypQP5sz5yNPTUwgh/UmpVG7VbXnIkDv69u1tfeXSpfT3358mhBg58qGBA+0tqrHJ2dm5ui8BAADAj
Wfz23m9Xl9QUKAp0Gi0GpPJZDQaL168mJWVFRwc4uvr45B+XoO6EVRMJtOFxAsFmgLpoZPCKbBBYIMGD
erQGFZxcfGGDRt9fX26d+8qBZXHHhuZkZF58uSpmTM/mTr1PaXSWVrN7+ZW7REVX1+fMn9zR44clQqDB
g2o4nYCAAAAuDm4uroGBAQEBASYTKaMjIzMjEyT2aTX65OSLuTmekdGRtSJu+g6EFR0Ol1CQoJeXzqQ4
uvrFxISXINDKI888nh+fkEVK1d6jorkySfHPPjgfdZXNmzYlJeX//DDw11dXaUrLi4ur7/+8sSJrxw+/
N+iRUufe278pUuXhBDBwQ2q+Qls2LXrXyFEREQ4KQUAAKDecnJyCg4O9vf3T0tLy83NFUIUFOSfOXMmO
jraza3as3huMLlnqYKCgrNnz0opxd3dvVnTZhER4XViope1oqLi1at/EULceusg6+uhoQ3Hjx8nhPDx8
RZCJCenCCHCwq43WmRlZR87FieEKDMfDAAAAPWQi4tLRERE06ZN3d09hBB6vf7s2bMFBVX9pt5RZD2ik
paWlpGRIYQQQhESEtygQQ0MNZT3+uuvGAyGSqt9+ukXWq321lsH9enTq9LKERHh1g/XrfszNzevfft24
eFh6ekZ1k/dccfg5OSUhx8eIYRISkoWQoSFhVbvA5SzY8cuaafjlStXr1y5uuov/O67JQ0a1LFVVgAAA
KgKtVrdpEnjjIyMtLRLJpMpISEhKCgoOLjmzx+vKfINKgkJCVLOc3ZyjoiM8PLyqqU36tq1c1Wqubq6a
rXa6Oionj27V6v9pKTkH35YKYR4+OEHyz+rUCieemqsEEKn00kZpmHDhtVqv7wNGzZew6uUSueAAP/rf
GsAAADIWYMGDVQq1YULF0wmU3p6enFxcVRUlKM7ZZtMg0pSUpKUUlxdXRs1amRZ11HnmEymzz+fZzAY2
rZt06lTRzs1z5w5JxUaNgyptNmEhMToaNt/UiaTaerU96rVyYULl+zbdyAsLKxOLKsCAADA9fDy8mrat
GlCQqJer8vPz09KSoqIiHB0p2yQY1DJzMyU1vqoVKomTZrU6btnhUJhMpkUCsWTT46xX/O//44IIZo2b
eLtXfnY0YsvvvrQQ8Mfe2xk+aecnJzKHHhfqUuX0oUQkZFy/AMFAABAjXNzc2vatEn8ufhiXXFubq5ar
ZbhKSuyywAajebixTQhhNJZGR0dXadTihBCoVCMHfvE8OH3t2jRzE41s9m8bdsOIUT37l0rbdNsNhsMx
sOHj9RID/V6fWLiBSFEo0bRNdIgAAAA5M/Z2Tm6UbR0et7Fi2lajdbRPSpLXjFAp9MlJiYKYVYoFNGNo
uvc7l42tW/f9vHHH7Vf59ixWOm4+t69K1+pbzQahRA1NR3u7Nlz0sr75s3tRSkAAADcZFxcXKKjoxUKh
RDmhMREnU7n6B5dRV5BJTHxgslkEkKEh4er1WpHd6fGlD8utIzVq38VQrRs2aJxY3vDGlKi0On0QggXl
5qZtidtZCyEaN68aY00CAAAgLrC3d09PDxcCGEyGZOSkhzdnavIKKjk5OTodMVCiAYNGvj6+jq6O2WZT
MZaavngwZiYmP+EEA88cK/NCtKQnBBC2kZZ2mbAw8Pj+t/abDZv2rRFCNGiRXNPT8/rbxAAAAB1i6+vr
7RApaioSFooLhNyWUxvNpsvpV0SQri6uoaEVL7tVa06e/ZcQUGBQuGkUAiFQmE2m5OTU/Py8oUQHh7uN
fte+fn58+Z9JYRo1aplRSe0WM4NPXDgkFqt3r17jxDCz696WS4rKzs/P9/T08Pd3d3V1dXZ2Vmn0//22
xppytmddw6+ro8BAACAOiskJCQ/P1+v16elpfn4+CgUCkf3SAj5BJWMjIwSQ4kQwuEpRQhx4ULSJ598X
v66SqVq375tzb7X559/mZGRoVQ6T5jwdEV/E25ubiEhwWlpl1at+mXVql+ki61atazWG+n1+okTX5Emj
5URERHev3+/6vYcAAAANweFQhEcHJyUlFRSUpKZmVlLx6xXlyymfhkMBukEepVK5ePj4+juiL59e/v6l
nZDoVB4eLhHRkYMHnzLl19+Fhp6vcfGlzF69KigoAYTJ05o2rSJnWqvvfZy8+bNpDlgoaGhDz00vKLhl
4o0bBjSrFnZVShKpbJPn15Tp76nVquq23MAAADcNHx9faVZPOnp6dJyA4dTmEtMju6DSE5OycnJFkJER
0fX3gn01WU2m81mc43vj6zX60+dOiOEaN68qfTXkJ9fUJWzU66f2WzW6XQ6nV6v15vNZldXF29v77q+A
TQAAABqRH5+fmJiohDC398/LCzM0d2RQVAxmUzH446bhVmtVjdtysZTAAAAgGOcO3uusKhQoVC0adPG4
StVHP9tukajMQuzECIgIMDRfQEAAADqL39/fyGE2WzWaDSO7osMgoq02a4QQj6TvgAAAIB6yNOr9LwKy
y26A8kgqOQXCCFUKlWlpyICAAAAqD0uLi5ubiohBCMqQqfTSbsSM5wCAAAAOJyXl6eQ7tJLShzbEwcHl
aKiIqnAsegAAACAw1luy4sKixzbEwcHFUtQU6k4xwMAAABwMLVaLRWkeU8O5OCgYigpPU2GBSoAAACAw
1luy+v71C8pqJFSAAAAAJlwdlYKgor0+QkqAAAAgEwolc6CoFJSYhAEFQAAAEA2pJtzyxoNR3F4UNELg
goAAAAgG9LNeX1fTC9RKBSO7gIAAAAAIYRwUjgJIUwmk4O74di3BwAAAIDyCCoAAAAAZIegAgAAAEB2C
CoAAAAAZIegAgAAAEB2CCoAAAAAZIegAgAAAEB2CCoAAAAAZKf+Hgmv1+vz8wsqetbDw12tVt/I/gAAA
ACwqI9BxWw2f/fdiv/7v98MBmNFdRQKRf/+fV9+eaKbm9uN7BsAAAAAUT+nfv37797Vq3+xk1KEEGaze
fv2nWvW/HHDegUAAADAoj4GlePHT1ax5rFjsbXaEwAAAAA21cegolQ6V7FmZmZWrfYEAAAAgE31MahUf
ZV8Tk5OrfYEAAAAgE31M6ioqlgzP7/A/lIWAAAAALWhPgYVlaoa+w5rNBVuYQwAAACgltTHoOLn51v1y
ixTAQAAAG68+hhUGjYMqXrlU6fO1F5PAAAAANhUHw98DA4OUigUZrO5KpXj4uKGDbuztrskT0VFRbt2/
7tm7VpfX98Ppk2t1mtTUlO3/vPPvn37T546mZp60dvLKyw8LCoyqm+fPj179ght2LDqTen1+kMxMTt27
DwWG5ucnJydkxMRER4ZEdG7V6++ffo2ahRd9aYMBuPuf3fv3bvvUExMckpKbm5uUIMGrVu36tK5y6233
BIdHVX1pjQazZatW/ft2x8bF5eSmlpSUhLasGH79u26d+t+y6BB/v5+VW+qjHPx8Rs3bvrxp5XfLlvWp
HHjqr/QbDbHxsVt37EjJibm9JmzWZmZQUFBoaGhLVo079unb4/u3by8vKreWnJKyu7du3fu2p2UlJSYm
Ojp5RURHt6qZcuBAwd069q16jtSiMt/DHv37jt37lxySoparY4ID+/YoUPfvn169+qlUlV1zVjNNnXs2
LGdu3bv27//QlJS2sWLfv7+TRo37tSp46CBAzt26KBQKKreFAAAqA0Kc4nJgW9/7NgxIYSfn194ePiNf
N+nnnouNTW1KjX9/PxWrPimtvsjK2azOe748b/Wb1i1enV6eroQYvTjj8+YPq2KL9dotQsXLfr8i7kVV
VAqlS9OnPjUk+Oqct984MDBKdOnHT16rKIKT40b98LEiVUJBrFxce9PmXLgwMGKKjw34dkXJ77g4eFea
VN//PnXjJkzk5OTbT7r5+f3zv/eemjEiGrd7Obl5W3bvv2XX/9v2/bt0pVtW7dUPagkJyd/OOujdX9Ue
ERpw4Yhb74x6d577ql0e26NRrNw8WI7v8GmTZtOfu/dgQMGVNqrkpKSBV99Pf+rr4qKimxWaN2q1fSpU
7t373Yjm8rMzJw+Y8b//fZ7RRXuuP32Ke+/d4P/pQQAgHwkJyXn5OYIIdq1a+fAbtTToPLhh7N37fq3i
pWXLVsYEhJcq/2RifT09M1btq5avTrm8GHr61UPKpcuXXpi7Li448crrdmje/dlSxZ7e3vbqfP7mjUvv
PRypU21aNH8x+XLg4KC7NQ5cODg2KeeysvLs9/UbbfeMm/uXE8PDzt1vv3u+/cmT660Vy9MnPjGa69Wm
lUMBmNMTMy6P//45ZdfNVqt9VNVDyr79u9/8unxlX46IcRT48a99+47Tk4VzvnMy8sb+9RTduKcxZT33
39y3Fg7FfR6/Vv/e/vnX3+1345SqVz09deDb7v1xjSVlpb25PjxdqKvJCoq6vtvv2ncqJH9agAA3JRkE
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
WEWdJMDAABoIicnxzSlODo6hoaGurq4GAyG+npJQUGB3qBn30pOSRaL3by9vdmXAoFg+LDh5tyitLS0p
LSE3Y6JjrZoulSANvT1oHL27Pn1699XtP4km2EYhaKhM7dQKhubWdPTr6SnX9m795/nnnt62LAhnblm1
3px1Uu/bN/eZOeA/v0zrlyx6Dp7//77199+Z7fj4mI3fPBhXFys8d2ofv3GjhkzZ/bsBx9+hA0Yx44d/
2vX7gXz/6/5pb757rtz586z2zNnzFj/1ptNJn0eOWLE/Hnzltx999WrWUS0a/ee02fOjB41yvxqlUrly
69aNvCGZRpUHnrwgU7Gzi++/Mo0pdy5ZMkjDz1o0RgeAACA1qjV6tS0y8aXwUHBQ4cOsbe3N+6Jjo46e
uyYsddxdk6OaVCJiAhv9xZ6vT4zM5PdthHY4PkadKE+HXnLyyveeuudNlJKd5DL5evWra+oqOzJm7bN8
O+jFCIKCQlZsXz533t2/7zlJ4suwjDMJ59+xm6LRKL3333XNKUYDRs6dM2rrxhf/vr7b82P0Wg033z7H
bvt7+//zttvNUkpxrdeWfWS8eWhw4ctKvjLzZvT0tKIKCwsbOgQC3JjVWUVu+Hh4dFiYeZLS0t7/Y032
G0XZ+efvv/+jddfQ0oBAICuUl1drdc39pUQiWybpBQiEovFgwYONL4sKS7R6/VkidLSUpm8cVRqZGSko
6Nj50oGuK5Pt6j8+ONWY3NHT1IqlVu3/vLkk4/3/K1b4+LsPH/+/BtnzRw+bBg7U4dEYtk6lTk5ORcuX
mS3F926cGB8fGtHzpwxIyQkpKioiIiOHTuuVCqb/FIrKChgB/ET0cIFC9oYcz9q1EiRSMT2Jbucnm5+t
VcyM997/wN2e9WLL2zf8av551ZUVrAbcbGxlnY2a2LDx41tKSKRaPNXX1rUIgQAANAutVpj/Br1cPdok
lJYXl5exm0DY1Cr1U5OTmZen2GYrKxs48uIyIgOFGkwGMorKirKK2pqahQNCqFQ6Ozk7O7u7h/g7+vjg
xEvfVmfDiqXL1vwp22vuXVzd99112urV3d44ixW5tXr/V8nTZzYxpG2trbjx439actW9qVEImkSVJxdX
F5b3dgpa3BSUtuXCgsNzb12jYhqa8ydjEur1b66eg0bb265+eYZ06dbFFRKSkrZjfDwMPPPau7cufN79
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
6SW51ypzM+plUlVXJcDwFOubg7hUZ5R/X37J/jzf7h82xBUAAAAwMrU1zZUlsvraxrkUrVKpWVX1wbom
2xsBA4Oti5u9u5eTr7+Lu6eTlxX1GUQVAAAAAAAgHcw+AwAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQ
QUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAA
HgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAA
AAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQ
QUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAA
HgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAA
AAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQ
QUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAA
HgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAA
AAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQQUAAAAAAHgHQ
QUAAAAAAHhHxHUBAAAAAAC9nE6nU6vVGo1Gq9PqdDq9Xm/QGwwMQ8QQCWwEAhuhjVAoFIlEtiJbOzs7e
3t7kaiv/6He139+AAAAAICu1dDQIJPLFQp5Q0NDQ4NSpVLqdDpLLyISiRwcHJ2cHJ2cnJydXVxdXJycn
LqjWt4SMFoD1zUAAAAAAFgxhmHq6usl9fUSqUQqlXYglphDJBK5ubmJ3cRid3cPd3eBQNAdd+EPBBUAA
AAAgI5QKpXVNdW1tbW1tXUM06N/VAsENp6eHp6ent5e3o6Ojj156x6DoAIAAAAAYAGlUllZVVlVVSWVS
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
LIn68ierNbevWfZ6FGTI4nowF+ZbAI5ezR//PTo6HjfkZMiZBLVhBnRDk62Oq1h3+8ZO3+6ZDzR2dWeb
W9p0aWzRS5uFn8xm4qI9vIJcCUinRb9vgD6Lobh3W+Aq+lF61b99OIbi8XuLbcbV1dKXnn6m7m3j/vu0
306nf7tV7Y+t/o2i7KKSqV5/7XtJ4+kE1HO1RK62jWV9xh9p5d2QVABAOgr/IPFoyZH/L0jXdnwn6dcQ
8eGjZwUQUQVpdKDf17/Jty55dLy1VNthILp8wYQUUWJ9KfPzl5NqzAesODeIWNu6PfHDxeP7M1u8Y7bv
kruZM2PvTiJDSo9MB4GAHjLwL+gQkQZqQXrVv704huLWxtDX1sj27xxD7udcibboqyiUmref70xpVgpf
afbwTDrFwBAn+Dl6/zkK1Nm/l/8Sx/MNp22y8ffZeG9QwQCAcPQ7l8umw5Yz06v3LPjMrstl6rfe2m/M
aXY2Ysefm781Jv7Oznbzb9nSESMdzeVbWsnZDd6uKcZAPAMw3UBLUu/lP/mqp9kErNWNmSzikLe/qS9K
qXmvdd+seqUQkQGQ2fjJYIKAECfUFOp2LUtTVqv8vRxvvuJ0fc+OcbOXuTqZv/I8xPdvZyI6OBfmWcO5
zU5688tqWeP5hORi5v9g8/+P3t3Hl3nXR54/L37vdo3W7bl3bITJ15jQuLsSVPSAANpp9Ck0Hba6dBCW
zqdaSEshTalQ4DQdqbTmZYOlB5CQhbOgZIEkrRZSCD76qzOVifY8RLHsiVr19X8IUcEIy9Yku8j+fP5w
8e+vnrfJ8fnRPer3/u+vzOLpVySJAvamz96+QVr189PkmRPR+8/fO6ulza+Nklj5wv7QqWv14oKHLuGg
3ZKkiTJhkdeuvxTV3d1Htbmhg/f99xVX/63g7+nt6f/i5ddd8/3n5qI6SppeNz/bC79AjhW3H3r808+v
OX9v3fKiWvnnHLOollz6wYHy20LGpIkeeKhLdd9ZezLtL72v+9tmlHdvnxG+/IZH/r42Y/e96N3XbKqW
JVLkmTzpo4v/9UPtrzccZCTfvyKC+cvaRr/8Hs7+8d/EIDJ8PhDL37uU9/46GUX19SWkiR5/bUDbr+46
qTFF/+ncw9yqJFKufeuKV8pE8KKCsAxZNfO7r+97Pbbbnx2eDhZ0N685PgZSZJs3tTxlb854H5hAwND/
++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
Hwhdfu/bB0cGK70dAAATBO1VTOymfyuzi3Dw8f0h8xMNllzWtPiZS3FUrFULK5csbKtrS2dTguVMfxEq
yTDyXCyct3cmvrcnTdsfX17f6WnAwBgmigV6rOZYkfXlv6BnkrPUhl1jdlVpzbNbmssFYstzS0rTlzR0
NAQqlKSJEkND5QrPcNPKJfLQ0NDAwMDff19PT093d3de/bsueOGH218vKvSowEAMK3s7tryn1qPAAAgA
ElEQVTa1fN6pac42uYtLq5eP6OqqrpYLCxauOiE5SeM7J2SToe4fX1UuFBJRltlcKC/r7+3t7e7p2fv3
q6nHt3+8F07d78+WOnpAACYPnr69uzu2jpUPiY+ZFbXZZatrFuwpKFYKtXW1CxfvrytrS2XjVgpScxQS
d5olcHBwf6B/r7evt7e3u7u7r3dex+6a9uTD+3p64k4MwAAU1F5eGh317bu3o5KDzKJ8oXUgmVVy1Y2F
gulYrEwb+6845cfX1WqymazMSslCRsqSZKUy+VyuTw4ODgwMNDf39/b19fb29Pb09vZuffx+3c8t2FvZ
8cxEb4AABwFvX2de7q3DQxOt1ujq2oybYuLi4+vryoV84V8fX3DcUuXzZ49O5fLZbPZdDods1KSyKGSJ
Mnw8PCbl1b6+/r7+vr6+vp6+3p7e3pfeLZj07Ndm/+91wILAAATorN7R+feHdPgcWD5QqplVr51Xmn23
Jp8IZ/P5fP53JIlS9rb2/O5/OhCSqi75/cTOlSSJBl5VvHQ0NBIroysrvT39/f19w/09/f19fUP9G/e1
Ln1lZ6d2/p37xzs2j14bD9oDgCAcRkqD3R279jbM8WuBEulklJ1prouXdeYa5pZaGktZTO5fD6XyWbzu
dzctrlLliypqakZSZRMJpNKpSJXShI/VEaMLq0MlYcGBwcHBwYH3mTkhcHBwZGY6doz0LN3cHCgXJ4C/
2UAAETU29uzY+f23bt3VXqQQ0ilkkw2VShmStXZTDqTzqQzmUw2M1Ij2Ww2M3v27IULFzY0NGSz2Uw6E
38hZdTUCJXkjaWVkdWVcrk8ODQ4NDg0OGJoaGhwcGTVZahcHi6Xy8PDyTG+hQ8AAOPW3d29ddvWHTvi7
mSfSpIklUolqVQ6lU6lM5l0Op1OZzKZdHrOnDnz582vb6jPZrIjr6XeUOmpD8uUCZURo7lSfsNIt+z7t
TxULpeHy/veU+lhAQCYDvr6+rZu27pt27aBgYFKzzKGkU5JpZJ0Kp1OpwuFwpw5c9rmtFVXV6ffMLUSZ
cQUC5VR+3JleF+WlMvl0V9HK2U40SoAAEyY7du3b9u+bdeuQNeDpZKR9kilUklTU9PsWbNnzZq1r0zSq
XQqPeX6ZNRUDZUR+4JktEze9JsKTwYAwDTV09uzY/uOHa/t6OzsrPQsSZIktbW1M2fMnDFzRlWpaqRJR
uNkiibKiKkdKvsZ7ROhAgDAZOvp6dm5c+fO13fu2rXrKH/+TKVSjY2NzU3Nzc3NpVJpNEimdJnsZ1qFC
gAAHH3Dw8O7Ojp2d3Ts3rN7z549g4OTsi95Nputq6urr6uvb2hobGiYTk0yJqECAAATqbu7u7Ora+/er
u7u7u7unt7eniNIl2w2WyyWqqpKVVVV1dU1tTU1VVVVkzFtWEIFAAAm1+DgYF9fX39//+juf+Whcnl4O
EmGkySVTqX27X+SzeayuXw+XygUstlspaeuMKECAACEk670AAAAAPsTKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVA
AAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwh
AoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQ
DhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAAP+/vfs7buJa4DhumPuepQNNGmDtBrSuIAq8Y
7kBJCoAN4DlFIAlCmBMA9aS94BoICgNRHID9n3QDMPlJmCDrfMT/nweYYZz9uGI/e6fsxBHqAAAAHGEC
gAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHq
AAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxh
AoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR
6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAc
YQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAE
EeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAA
HGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAA
BBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAA
ABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKA
AAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoA
AAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGEC
gAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHq
AAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxh
AoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR
6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAc
YQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAEEeoAAAAcYQKAAAQR6gAAABxhAoAABBHqAAAAHGECgAAE
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
vQAAAAASUVORK5CYII=',
                    id=None
                )
            ),
            TextPart(type='text', text='</image>')
        ],
        message='Loaded image file `/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000
gn/T/dev-issue-55-815001100/.maestro/after_runflow.png` (image/png, 91058 bytes,
original size 1080x2400px). If you need to output coordinates, output relative c
oordinates first and compute absolute coordinates using the original image size;
if you generate or edit images/videos via commands or scripts, read the result b
ack immediately before continuing.',
        display=[],
        extras=None
    )
)
StepBegin(n=66)
ThinkPart(
    type='think',
    think='截图显示 `runFlow` 完成后页面仍然在登录页，倒计时显示 "37 s"。这说明 
`tapOn "登录"` 实际上**没有触发**登录操作！\n\n但 `test_simple_login.yaml`（没有
`runFlow`）成功了。这太奇怪了。让我比较两者的差异。\n\n`test_simple_login.yaml`:
\n```yaml\n- launchApp:\n    clearState: true\n- assertVisible: "欢迎使用 Social
World"\n- tapOn: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- 
tapOn: "验证码"\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n```\n\n
`login-as-test-user.yaml` (runFlow):\n```yaml\nappId: com.example.mobile\n---\n-
assertVisible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"\n- tap
On: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- tapOn: "验证
码"\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n```\n\n唯一区别是 `
login-as-test-user.yaml` 有 `appId` 和 `---`，以及多了一个 `assertVisible: "手机
号一键登录"`。\n\n但 `test_simple_login.yaml` 在同样的设备上成功了！\n\n等等，让
我检查 `test_simple_login.yaml` 的截图，确认它确实跳转到了完善资料页。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_ZUomYgY8rtpIUMbBObBWZHIF',
    function=FunctionBody(
        name='ReadMediaFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/login_after_click.png"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.39105224609375,
    context_tokens=102512,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=3952,
        output=363,
        input_cache_read=98560,
        input_cache_creation=0
    ),
    message_id='chatcmpl-dZhR7Pl7jBn1zXqiBZ5EBS5g',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_ZUomYgY8rtpIUMbBObBWZHIF',
    return_value=ToolOk(
        is_error=False,
        output=[
            TextPart(
                type='text',
                text='<image path="/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro/login_after_click.png">'
            ),
            ImageURLPart(
                type='image_url',
                image_url=ImageURL(
                    url='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABDgAAAlgC
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
iXJlihZK/+VK39lSpK1UGq7/yuWofLZ559nLdesWXMTzmTL8c6772Yt791wg24+Wa+P5s4dP+GvpxKf0
rVrlSpV1rVm0yZNTuzSZeKkSZkfbx8zNitUMjIyvv322y8WLfp84cL58z/5YPYHs/NfGlWqVGnapEmD+
vVr1669yy475/9HAQDY3ORyDVjmCrnfYZ89UbKWs5dM1v8Xy1CZNWtW1nL9PfbYhDPZQnz//fcz3n476
2OR/jfPyMi4deSorI+djz8u9/XP6Xf2lKeeWr169XaVKw++/LJvv/32kccemzfv45mzZv3nP/8p2Bw6d
uhwSteTa9WqVb1aNY8zBgBIkf0asOyVkr6QfZ0UmatlX0h5UFjxC5Ufli79fOHCrI916mzUN6znYszYO
669/vqNdriieFTXukx96aWsl042bLhX7dq1i+5YM95++4WpUzOX27Zts8f67rqpWbPmGb1OHznqtlEjR
9auXXvJkiU3DLsxj8eqVKlS82ZNGzRosHudOm//850HH3ooc3zPBg3223ffAv8IAACbvfRWSdYRLeuS/
caVlIXM5eIXKrNnz87+cbfddt1UM9lCrF27NutCrCRJunTuXHQnGX755Zdr/nFt1sfTTj01L1v16NatZ
o0aLQ88IEmS6tWrb7vttpmPLU632667Nm3atF69unVq195l55132GHH0v//oXBLlny1wdMHANiCpLRKk
tNlYOvaal3nVZJifUbl2eeez1quU6dO9WrrfIMKheKVV16dM2dO1sc2bdoU3bHuuffej+bOzVw+5OCDD
9h//7xsVbVq1RO7dMlcLlWqVPNmzaa++GKSJGXLlm3SuPFee+1Zt27d3XbddedaO1euvG0RzRwAYAuU8
hyw9d5Yn3LuJfstK0m2RCmW96gsX778qaefzvrY+bjjiuhf9x98+OGmTZqs97qjdTn4oIPatmlduFNav
vzH4bfcUrj7XK9ly5ZfOXRo1sfOxx+/Q/XqRXSsd9+dOeym/770s//ZZxfsD7fbKV07H3/crrvsstNOO
3ltDgBAkcq9VTIHc7xlJftCSqhkKmahcv8D43777besj21atyqiA73zzruDLrv86E6dOnXseFDLA/P1P
sokSRo2bFjojzBesmTJRg6VjIyMm0eM+PLLLzM//q1ChQHnnZs5PmTo0N9++611q9b7tmhesWLFQjncn
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
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
QAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgH
KECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAA
BCOUAEAAMIRKgAAQDhCBQAACEeoAAAA4QgVAAAgHKECAACEI1QAAIBwhAoAABCOUAEAAMIRKgAAQDhCB
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
OEm++OOPWbyxAm77bprjk118003HnVk++wjny9c+OabbxXVdAGKG6ECsCWaMHHSiSd3nfPRR+lfDbvxp
omTJqWP3zJ8+P4FfZN93r355luDLrs8fbxZ06aTJk9OGXxo/PhDWrUePGTIv//970KcQ+4vfPz111/fe
PPNSwddttc+jQZdfsWiRTm8SWa7ypVvG3nriJtvrlix4tCrrxl20/D0dRru1bDLCSekDI4eM2ZDZg6wO
REqAFuWNWvWDL3mHwMvuWT6jBntj+owavTo3377Levbh8aPHzV6dPpWFw4YcPxxxxb13D799NOTu3VLH
69SpUouzwW+7/4HDjjo4FtG3Przzz8XyjRyfNDzzytWTHvjjSFDh9apt0fXbt2z7qdP1/Xkk15+6cVOH
TuuWvVLv3POveuee3Jc7eP5H++/334pL6qfPmPGo489voHzB9g8CBWALcjSZcv69O171913Z40Mu/GmE
048ae68eUmSPP/CC5cOuix9q25du557Tv+inttHc+f2OO30HL8a/9CD9997T+NcH0x884gRhx951NPPP
FM0s0sWLFhw4cCLU17gmOLQQw6Z8sTj1197bZXttvto7tyTunad8tRT61p59uwPy5Qp0/Xkk1LGr/nHP
35YurRwJg1QnAkVgC3Fp59+evwJXV5+5dWU8ffef//w9keef8EFvc/sm75Vm9atrhx8RVHP7bXXXz/iy
KOWLFmS/tXwG4ftUa9e61atpjzx+N133lG37u7r2smXX355Vr/+p3Tv8fH8+YU+wyaNG08c/3C1atVy/
LbRPvs8/OC4cfff17hRozVr1tw+ZuwRRx71/gcf5LLDt2ZMT5KkS+fOKeNLly1Lf80lwBZIqAAUS3l/M
32Wrbfe+v+xd99Rct31/f9vmzt9tmi1TaterLqSwcYNFwwWzWCbZnAwhC+hpJDwI5iSQhJICISEQMAQe
jG92AZjOzaYZhLjbjVbkmWrS7vStulz+++PkYUsbZvZmfncO/N8HB2d8cwtb+vunb2vez9lzeqzpvp00
hZHZ2/a9KlPfjIcDldR4SyVSqX//NSn3/THb5n001ddc/XrTrmUf/HmzXfcdtsn/+PfF/T3T7XB3/z2t
5tf8tLP3HijaZrVlTRVH5Xly5Z95YtfPO3Nc88956tf+tJPb73lkosvliTpiZ07X3vt6//14x+fcS87d
+6SJKmrq+uf/uEfTvvomzfdNGk3IQBoKQQVAAgk1z39Ynr6LuCSJA0MDHzus5/9xMc/NstdLFy48HM3f
rY823qd7N69+9o3XPfJT31q0k8vveSSj/7zv5z2pq7rr331q+/5+c8//I//OK+zc6ot/9sn/v31f/RHu
3fPPHtJRQYHN/zrR0+UdMnFF9/0ja/f/MMfXnHFiyRJOnDgwD9/9KObX/LSRx599MwV/+5v/ubMN0ulk
iRJb3j9tWeOpXbD+z/w45tvqW3xABAsBBUACKQqnqiUvf7aa391zy/OPfecGZf8+le+fOqEhrU1PDz8i
f/4jxdufvGkl/WSJG2+4oovfP7zU80sGY/H3vLHb77nFz+//o1vnGoXDz740As3v7iKyUmm/5d843XX3
fDXf33zD3/47Zu+edmll0qStG/f/n/68EcuuuTSL3zxS2cuv3jx4jtv/9k73v62Mz+yHUeSpGg0+sH3v
e/MT9/9nvd87vP/XWnxANA0CCoAEEhVPFE5acXy5d+56ab3vPvd0yxz7rnndE79vGIuJiYmPnPjjZdd/
sL/+sxnp1rmihe98LP/9empUspJ8zo7P/rPH/nB9767ft26qZb5uw/9w3XXX3/w4MHZVzjpZJqn/vP+5
bv+opz0Dh48+OGP/PPFl1325a9+ddJNvf7aa2+79Zb169ZNTEyc+altWeUXV1zxouve8PozF/jXj3/8/
73tbcePH5998QDQNGTPquPcxgDgT//5qU9P1dzIJ770hf9+yYtfPM0Cb3nrW0/rFv+5z37mFVdeWdFeH
nzwofe+731P79076aeLFy/+/I2f3bB+fUXbnMbQ0NDNt9w6Y/+Nd/35n//VX76roo4xpVLpmzfd9JF/+
eg0y3zz6197wWWXzWZrIyMjZ59z7mlvPr5tazKZLL/et2//7+///d0///nPf3HPVBvp6+v9lw9/pNwqT
JKkhx955OpXvfrUBbq6uh596MGT/5nPF970lj9+4IEHpTMMDAz8y0c+fPkLXjCb4gGgafBEBUArcj2/3
6OZ8fGIbTuVrnKmc88956e33nLmtINl+/fvf/VrX/fT226rdLNnevChh254/wfOPf+CGVPKV7/0pffd8
N5Ku+9HIpG3v+1tv/z53Wdv2jTVMm/647dM9dxjNhzH2b5jx5e+/OWrrnnVxZdddsP7PzBNSnnta15zx
223nUwpkiT98Ec/Om2ZSy+5+NT/jMdjn/7kJ7u7u8/c2qFDh978lv+3fceOqosHgCDSRBcAABCpra3t3
//t45decsnf/v3fj4+Pn/ZpsVj883f95cFDh/7sne+cZR+Yk9Lp9JatW39777133X33pNO3n+aC88//5
L9/YmBgoKK9nGrlypXf+863//PTn/7vL5w+NlfZP334I7t27f6nf/hQLDZDo7IzXXf9m7Zt2zbjYhdec
MG7/+ovLzj//FPf/PHNt3z7O989bclzzjm9m9DAwMCNn/mv1157ehuwrq6u73zrpjWrV1daMwAEGkEFA
AJJUWr5SPwVV778rFUr//itfzJpX46Pffzf8rn8+25472w2dfDgwW9/57s7Hn/817/5zSz33tfXe8N73
3v1K18ZCoUqKHoysVjsbz/4wQsvuOD9H/zg0aNDZy7wve9//+jRo9/65jcq3fIXPnfje9/3/v+7776pF
li/bt3/9+6/uvUnP3nd69+wbOnSNWtWRyIRx3EffPDBw0eOnLn8hnWTtKk7/7zzvnPTTdddf/2pb379K
18hpQBoQTT9AtCKFNnv336VPr6Yu1WrVn3v299asWLFpJ9+5sYbP/1fn5nNdhYuXJhOp2efUt733r++5
+67X/vqV889pZz0gssuu+3WW6fq5LP5lBZZs9fe3v75G28cHNxw5kfLli79zKc/deftP9t8xRV/+s53S
pL09N69t99x549vvuXWn/xk0pTyiiuv3LhxcNIdXXzx83/587tPDlj8xf/+/FRLAkBz8/uvagBAwyxat
OjbN31z3dq1k346fef+U33o7//utLZPZ1rQ3/8PH/r7Rx968F1/8RcnO6nXUE9Pz5e+8N9/+s53nPb+O
97+tjc9+3nF7HV2dnzi2X1sLrv00i9/8Qt3/8+dV191VfmdDevXv+iFl0+/nQX9/R/+x9MneTzVypUrf
/j9722+4orr3/jGl77kJdVVCwBBR9MvAK3o+Rdd1NPTI7qK6Zw5A2Bj9Pf1fftbN73jnX96/wMPnPr+z
37yk7POWjXLjUSj0U/95ydfcdXVx44dO/PTiy9+/puvv/4Fl12m63oNKp7W33zgAx3tHR/92Ik5Ll/+s
pe+/4Yb5rLBtWvWfOjv/vbD//wv73zH2191zTWTtsh6+cteftqAbKf5z0/+R1dX1/Q7mtfZ+ZUvfTGXz
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
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
QAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAc
IQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAA
EA4QgUAAAhHqAAAAOEIFQAAIByhAgAAhCNUAACAcIQKAAAQjlABAADCESoAAEA4QgUAAAhHqAAAAOEIF
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
AAAwDn9F2ZCV1bvEjGgAAAAAElFTkSuQmCC',
                    id=None
                )
            ),
            TextPart(type='text', text='</image>')
        ],
        message='Loaded image file `/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000
gn/T/dev-issue-55-815001100/.maestro/login_after_click.png` (image/png, 67131 by
tes, original size 1080x2400px). If you need to output coordinates, output relat
ive coordinates first and compute absolute coordinates using the original image 
size; if you generate or edit images/videos via commands or scripts, read the re
sult back immediately before continuing.',
        display=[],
        extras=None
    )
)
StepBegin(n=67)
ThinkPart(
    type='think',
    think='是的！`test_simple_login.yaml` 成功跳转到了"完善资料"页面。但 `test_s
imple_runflow.yaml`（使用 `runFlow`）没有跳转。\n\n这说明 **Maestro 的 `runFlow`
在 1.39.13 版本中有 bug**，导致 `tapOn` 等交互命令在 `runFlow` 内部没有正确执行
。\n\n具体表现：\n- `runFlow` 内部的 `tapOn "登录"` 报告 COMPLETED，但实际上没有
触发点击\n- 页面仍然停留在登录页，倒计时还在走\n\n这是一个**Maestro 版本 bug / 
环境问题**。由于我的任务不是修复 Maestro，而是让 E2E 测试跑通，我需要修改 flow 
文件来绕过这个问题。\n\n**修复方案**：将所有使用 `runFlow: includes/login-as-tes
t-user.yaml` 的地方，替换为内联的登录步骤。\n\n但这样会有代码重复。不过根据任务
要求，我需要先跑通测试。\n\n让我先检查一下所有需要修改的文件：\n1. `flows/auth/l
ogin.yaml`\n2. `flows/profile/setup.yaml`\n3. `flows/matching/swipe_recommendati
ons.yaml`\n4. `flows/matching/view_match_list.yaml`\n\n这些文件都使用了 `runFlow
: includes/login-as-test-user.yaml`。\n\n由于 `login-as-test-user.yaml` 有 10 行
左右的步骤，我可以把它内联到每个文件中。\n\n或者，我可以保留 `includes/login-as-
test-user.yaml` 作为文档，但在 flow 文件中内联步骤。\n\n但等等，我注意到 `test_s
imple_login.yaml` 没有 `assertVisible: "手机号一键登录"` 这一步，而 `login-as-te
st-user.yaml` 有。但这不是问题。\n\n让我把所有需要 `runFlow` 的 flow 文件修改为
内联步骤。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_krNy8BnpVBbK8AXElKl3j2FN',
    function=FunctionBody(
        name='SetTodoList',
        arguments='{"todos": [{"title":"阅读项目地图与E2E规范","status":"done"},
{"title":"扫描已有flow并分析覆盖范围","status":"done"},{"title":"检查mobile代码
中的UI元素和页面结构","status":"done"},{"title":"修复includes格式问题","status":
"done"},{"title":"执行Maestro测试","status":"done"},{"title":"分析失败根因并修复
flow","status":"in_progress"},{"title":"生成E2E测试报告","status":"pending"}]}'
    ),
    extras=None
)
ToolCall(
    type='function',
    id='tool_GHaCJMWd5F1qNFlKDtFUPFjH',
    function=FunctionBody(name='StrReplaceFile', arguments='{"path": "/'),
    extras=None
)
ToolResult(
    tool_call_id='tool_krNy8BnpVBbK8AXElKl3j2FN',
    return_value=ToolReturnValue(
        is_error=False,
        output='Todo list updated',
        message='Todo list updated',
        display=[
            TodoDisplayBlock(
                type='todo',
                items=[
                    TodoDisplayItem(
                        title='阅读项目地图与E2E规范',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='扫描已有flow并分析覆盖范围',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='检查mobile代码中的UI元素和页面结构',
                        status='done'
                    ),
                    TodoDisplayItem(
                        title='修复includes格式问题',
                        status='done'
                    ),
                    TodoDisplayItem(title='执行Maestro测试', status='done'),
                    TodoDisplayItem(
                        title='分析失败根因并修复flow',
                        status='in_progress'
                    ),
                    TodoDisplayItem(title='生成E2E测试报告', status='pending')
                ]
            )
        ],
        extras=None
    )
)
ToolCallPart(
    arguments_part='var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue-55
-815001100/.maestro/flows/auth/login.yaml", "edit": {"old":"- launchApp:\\n    c
learState: true\\n- runFlow: includes/login-as-test-user.yaml\\n\\n# 断言：登录
成功后离开登录页（新用户进入完善资料，老用户进入首页）\\n- assertNotVisible: \\"
手机号一键登录\\"\\n- assertVisible:\\n    text: \\"完善资料\\"\\n    optional: 
true\\n- assertVisible:\\n    id: \\"tab_home\\"\\n    optional: true","new":"- 
launchApp:\\n    clearState: true\\n\\n# ===== 登录步骤（原 includes/login-as-te
st-user.yaml，因 Maestro runFlow bug 内联） =====\\n- assertVisible: \\"欢迎使用
Social World\\"\\n- assertVisible: \\"手机号一键登录\\"\\n- tapOn: \\"手机号\\"\
\n- inputText: \\"13800138000\\"\\n- tapOn: \\"获取验证码\\"\\n- tapOn: \\"验证
码\\"\\n- inputText: \\"123456\\"\\n- hideKeyboard\\n- tapOn: \\"登录\\"\\n# ===
===============================================================================\
\n\\n# 断言：登录成功后离开登录页（新用户进入完善资料，老用户进入首页）\\n- asse
rtNotVisible: \\"手机号一键登录\\"\\n- assertVisible:\\n    text: \\"完善资料\\"
\\n    optional: true\\n- assertVisible:\\n    id: \\"tab_home\\"\\n    optional
: true"}}'
)
StatusUpdate(
    context_usage=0.40594482421875,
    context_tokens=106416,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=3912,
        output=834,
        input_cache_read=102504,
        input_cache_creation=0
    ),
    message_id='chatcmpl-UKS7FN7t8j3hQ97WiGW6I082',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_GHaCJMWd5F1qNFlKDtFUPFjH',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/flows/auth/login.yaml',
                old_text='---\n- launchApp:\n    clearState: true\n- runFlow: in
cludes/login-as-test-user.yaml\n\n# 断言：登录成功后离开登录页（新用户进入完善资
料，老用户进入首页）\n- assertNotVisible: "手机号一键登录"',
                new_text='---\n- launchApp:\n    clearState: true\n\n# ===== 登
录步骤（原 includes/login-as-test-user.yaml，因 Maestro runFlow bug 内联） =====
\n- assertVisible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"\n-
tapOn: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- tapOn: "验
证码"\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n# ===============
===================================================================\n\n# 断言：
登录成功后离开登录页（新用户进入完善资料，老用户进入首页）\n- assertNotVisible: 
"手机号一键登录"',
                old_start=14,
                new_start=14,
                is_summary=False
            )
        ],
        extras=None
    )
)
StepBegin(n=68)
ThinkPart(
    type='think',
    think='现在让我继续修改其他使用了 `runFlow: includes/login-as-test-user.yaml
` 的文件。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_9pZFJAqDMd5zI9RcHqhsP2UH',
    function=FunctionBody(
        name='StrReplaceFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/flows/profile/setup.yaml", "edit": {"old":"- launc
hApp:\\n    clearState: true\\n\\n# 先完成登录（clearState 会清除本地 token）\\n
- runFlow: includes/login-as-test-user.yaml\\n\\n# 等待页面跳转后断言完善资料页\
\n- assertNotVisible: \\"手机号一键登录\\"\\n- assertVisible: \\"完善资料\\"","n
ew":"- launchApp:\\n    clearState: true\\n\\n# ===== 登录步骤（原 includes/logi
n-as-test-user.yaml，因 Maestro runFlow bug 内联） =====\\n- assertVisible: \\"
欢迎使用 Social World\\"\\n- assertVisible: \\"手机号一键登录\\"\\n- tapOn: \\"
手机号\\"\\n- inputText: \\"13800138000\\"\\n- tapOn: \\"获取验证码\\"\\n- tapOn
: \\"验证码\\"\\n- inputText: \\"123456\\"\\n- hideKeyboard\\n- tapOn: \\"登录\\
"\\n# ==========================================================================
========\\n\\n# 等待页面跳转后断言完善资料页\\n- assertNotVisible: \\"手机号一键
登录\\"\\n- assertVisible: \\"完善资料\\""}}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.4093780517578125,
    context_tokens=107316,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=908,
        output=314,
        input_cache_read=106408,
        input_cache_creation=0
    ),
    message_id='chatcmpl-B6dnkKgl528GiljTjy087A85',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_9pZFJAqDMd5zI9RcHqhsP2UH',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/flows/profile/setup.yaml',
                old_text='- launchApp:\n    clearState: true\n\n# 先完成登录（cl
earState 会清除本地 token）\n- runFlow: includes/login-as-test-user.yaml\n\n# 等
待页面跳转后断言完善资料页\n- assertNotVisible: "手机号一键登录"',
                new_text='- launchApp:\n    clearState: true\n\n# ===== 登录步骤
（原 includes/login-as-test-user.yaml，因 Maestro runFlow bug 内联） =====\n- as
sertVisible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"\n- tapOn
: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- tapOn: "验证码"
\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n# ====================
==============================================================\n\n# 等待页面跳转
后断言完善资料页\n- assertNotVisible: "手机号一键登录"',
                old_start=15,
                new_start=15,
                is_summary=False
            )
        ],
        extras=None
    )
)
StepBegin(n=69)
ThinkPart(
    type='think',
    think='继续修改 `swipe_recommendations.yaml` 和 `view_match_list.yaml`。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_mhGrkJRnqMxocPzeaooyY507',
    function=FunctionBody(
        name='StrReplaceFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/flows/matching/swipe_recommendations.yaml", "edit"
: {"old":"- launchApp:\\n    clearState: true\\n- runFlow: includes/login-as-tes
t-user.yaml\\n\\n# 断言首页加载（通过模式切换 Tab 判断）\\n- assertVisible:\\n  
 id: \\"mode_tab_nearby\\"","new":"- launchApp:\\n    clearState: true\\n\\n# ==
=== 登录步骤（原 includes/login-as-test-user.yaml，因 Maestro runFlow bug 内联）
=====\\n- assertVisible: \\"欢迎使用 Social World\\"\\n- assertVisible: \\"手机
号一键登录\\"\\n- tapOn: \\"手机号\\"\\n- inputText: \\"13800138000\\"\\n- tapOn
: \\"获取验证码\\"\\n- tapOn: \\"验证码\\"\\n- inputText: \\"123456\\"\\n- hideK
eyboard\\n- tapOn: \\"登录\\"\\n# ==============================================
====================================\\n\\n# 断言首页加载（通过模式切换 Tab 判断
）\\n- assertVisible:\\n    id: \\"mode_tab_nearby\\""}}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.41072845458984375,
    context_tokens=107670,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=406,
        output=298,
        input_cache_read=107264,
        input_cache_creation=0
    ),
    message_id='chatcmpl-GquVI2NMnw7FH0NkEyoN8bIN',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_mhGrkJRnqMxocPzeaooyY507',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/flows/matching/swipe_recommendations.yaml',
                old_text='---\n- launchApp:\n    clearState: true\n- runFlow: in
cludes/login-as-test-user.yaml\n\n# 断言首页加载（通过模式切换 Tab 判断）\n- ass
ertVisible:',
                new_text='---\n- launchApp:\n    clearState: true\n\n# ===== 登
录步骤（原 includes/login-as-test-user.yaml，因 Maestro runFlow bug 内联） =====
\n- assertVisible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"\n-
tapOn: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- tapOn: "验
证码"\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n# ===============
===================================================================\n\n# 断言首
页加载（通过模式切换 Tab 判断）\n- assertVisible:',
                old_start=12,
                new_start=12,
                is_summary=False
            )
        ],
        extras=None
    )
)
StepBegin(n=70)
ThinkPart(
    type='think',
    think='继续修改 `view_match_list.yaml`。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_QiXiGrSNIxelytkpG3EcwfPE',
    function=FunctionBody(
        name='StrReplaceFile',
        arguments='{"path": "/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/de
v-issue-55-815001100/.maestro/flows/matching/view_match_list.yaml", "edit": {"ol
d":"- launchApp:\\n    clearState: true\\n- runFlow: includes/login-as-test-user
.yaml\\n\\n# 切换到消息 Tab（使用稳定 id 断言）\\n- tapOn:\\n    id: \\"tab_mess
ages\\"","new":"- launchApp:\\n    clearState: true\\n\\n# ===== 登录步骤（原 in
cludes/login-as-test-user.yaml，因 Maestro runFlow bug 内联） =====\\n- assertVi
sible: \\"欢迎使用 Social World\\"\\n- assertVisible: \\"手机号一键登录\\"\\n- t
apOn: \\"手机号\\"\\n- inputText: \\"13800138000\\"\\n- tapOn: \\"获取验证码\\"\
\n- tapOn: \\"验证码\\"\\n- inputText: \\"123456\\"\\n- hideKeyboard\\n- tapOn: 
\\"登录\\"\\n# =================================================================
=================\\n\\n# 切换到消息 Tab（使用稳定 id 断言）\\n- tapOn:\\n    id:
\\"tab_messages\\""}}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.412017822265625,
    context_tokens=108008,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=488,
        output=280,
        input_cache_read=107520,
        input_cache_creation=0
    ),
    message_id='chatcmpl-YmE7WtOb9L71yRusH9ZQgF3w',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_QiXiGrSNIxelytkpG3EcwfPE',
    return_value=ToolReturnValue(
        is_error=False,
        output='',
        message='File successfully edited. Applied 1 edit(s) with 1 total replac
ement(s).',
        display=[
            DiffDisplayBlock(
                type='diff',
                path='/var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000gn/T/dev-issue
-55-815001100/.maestro/flows/matching/view_match_list.yaml',
                old_text='---\n- launchApp:\n    clearState: true\n- runFlow: in
cludes/login-as-test-user.yaml\n\n# 切换到消息 Tab（使用稳定 id 断言）\n- tapOn:
',
                new_text='---\n- launchApp:\n    clearState: true\n\n# ===== 登
录步骤（原 includes/login-as-test-user.yaml，因 Maestro runFlow bug 内联） =====
\n- assertVisible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"\n-
tapOn: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- tapOn: "验
证码"\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n# ===============
===================================================================\n\n# 切换到
消息 Tab（使用稳定 id 断言）\n- tapOn:',
                old_start=12,
                new_start=12,
                is_summary=False
            )
        ],
        extras=None
    )
)
StepBegin(n=71)
ThinkPart(
    type='think',
    think='现在让我重新运行所有 flow 测试。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_Ba8TV7rnMbJuB71gL40TjUMp',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test --debug-output=\\"/Users/ins
ulate/Desktop/new-world/new-world-ops/logs/social-world/e2e-screenshots-e2e-issu
e-55-20260424-012302\\" --flatten-debug-output flows/auth/login.yaml flows/profi
le/setup.yaml flows/matching/swipe_recommendations.yaml flows/matching/view_matc
h_list.yaml flows/devtools/open_drawer.yaml flows/devtools/edit_server_url.yaml"
, "timeout": 300}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.413238525390625,
    context_tokens=108328,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=552,
        output=159,
        input_cache_read=107776,
        input_cache_creation=0
    ),
    message_id='chatcmpl-gHmu5nzsNrqTlwnR8YlfC1bL',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_Ba8TV7rnMbJuB71gL40TjUMp',
    return_value=ToolReturnValue(
        is_error=True,
        output='\nWaiting for flows to complete...\n[Failed] login (35s) (Assert
ion is false: "手机号一键登录" is not visible)\n[Failed] setup (34s) (Assertion 
is false: "手机号一键登录" is not visible)\n[Failed] swipe_recommendations (35s)
(Assertion is false: id: mode_tab_nearby is visible)\n[Failed] view_match_list (
34s) (Element not found: Id matching regex: tab_messages)\n[Passed] open_drawer 
(5s)\n[Passed] edit_server_url (37s)\n\n4/6 Flows Failed\n\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=72)
ThinkPart(
    type='think',
    think='仍然有4个失败。但现在的错误信息和之前不同：\n\n1. `login` - `Assertio
n is false: "手机号一键登录" is not visible`\n2. `setup` - `Assertion is false: 
"手机号一键登录" is not visible`\n3. `swipe_recommendations` - `Assertion is fal
se: id: mode_tab_nearby is visible`\n4. `view_match_list` - `Element not found: 
Id matching regex: tab_messages`\n\n等等，`login` 和 `setup` 的 `assertNotVisibl
e: "手机号一键登录"` 失败了，这说明登录**成功**了！因为如果登录失败，页面还在登
录页，那 `assertNotVisible` 应该检查的是登录后页面是否不再显示登录文字。\n\n不对
，`assertNotVisible: "手机号一键登录"` 失败意味着 "手机号一键登录" 仍然可见，说
明登录**没有成功**，页面还在登录页。\n\n但为什么 `test_simple_login.yaml` 成功了
呢？\n\n让我看看这些 flow 是否真的是串行执行的。从时间戳来看，它们可能还是串行的
，每个大约 35 秒。\n\n等等，我发现一个关键问题：`login.yaml` 现在内联了登录步骤
，但 `setup.yaml` 也内联了登录步骤，而且两者都用了 `clearState: true`。当 Maestr
o 串行执行时：\n1. `login` 用 `clearState: true` 清除数据，然后登录\n2. `setup` 
用 `clearState: true` 再次清除数据，然后登录\n\n如果两者都各自独立运行，那应该没
问题。\n\n但 `login` 失败在 `assertNotVisible: "手机号一键登录"`。让我查看 login
flow 的截图。\n\n等等，让我先单独运行 `login.yaml`，不要和其他 flow 一起。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_fwCQ2ETrR0mT5HHOpOwaVMIw',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test flows/auth/login.yaml", "tim
eout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.4144439697265625,
    context_tokens=108644,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=356,
        output=446,
        input_cache_read=108288,
        input_cache_creation=0
    ),
    message_id='chatcmpl-KLL02jTVOAu0sN69aMEQC4iv',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_fwCQ2ETrR0mT5HHOpOwaVMIw',
    return_value=ToolReturnValue(
        is_error=True,
        output='Running on emulator-5556\n > Flow login\nLaunch app "com.example
.mobile" with clear state... COMPLETED\nAssert that "欢迎使用 Social World" is v
isible... FAILED\n\nAssertion is false: "欢迎使用 Social World" is visible\n\n==
== Debug output (logs & screenshots) ====\n\n/Users/insulate/.maestro/tests/2026
-04-24_014920\n',
        message='Command failed with exit code: 1.',
        display=[
            BriefDisplayBlock(type='brief', text='Failed with exit code: 1')
        ],
        extras=None
    )
)
StepBegin(n=73)
ThinkPart(
    type='think',
    think='`login.yaml` 现在失败在 `assertVisible: "欢迎使用 Social World"`，这
说明 `launchApp clearState: true` 后应用启动但显示白屏！\n\n但 `test_simple_logi
n.yaml` 在单独运行时成功了。让我比较两者的差异。\n\n`login.yaml`:\n```yaml\nappI
d: com.example.mobile\n---\n- launchApp:\n    clearState: true\n\n# ===== 登录步
骤（原 includes/login-as-test-user.yaml，因 Maestro runFlow bug 内联） =====\n- 
assertVisible: "欢迎使用 Social World"\n- assertVisible: "手机号一键登录"\n- tap
On: "手机号"\n- inputText: "13800138000"\n- tapOn: "获取验证码"\n- tapOn: "验证
码"\n- inputText: "123456"\n- hideKeyboard\n- tapOn: "登录"\n```\n\n`test_simple
_login.yaml`:\n```yaml\nappId: com.example.mobile\n---\n- launchApp:\n    clearS
tate: true\n- assertVisible: "欢迎使用 Social World"\n- tapOn: "手机号"\n- input
Text: "13800138000"\n- tapOn: "获取验证码"\n- tapOn: "验证码"\n- inputText: "123
456"\n- hideKeyboard\n- tapOn: "登录"\n- extendedWaitUntil:\n    visible: "完善
资料"\n    timeout: 10000\n    optional: true\n- takeScreenshot: login_after_cli
ck\n```\n\n区别是 `login.yaml` 多了 `- assertVisible: "手机号一键登录"`。\n\n但
这不可能是原因。让我测试一下 `login.yaml` 的内容。等等，实际上 `login.yaml` 的 `
launchApp clearState: true` 后失败在白屏，这和之前批量运行时的现象一样。\n\n但 `
test_simple_login.yaml` 单独运行时是成功的！\n\n唯一的区别是：我之前运行 `test_s
imple_login.yaml` 时，可能应用已经在运行状态（不是冷启动）。而 `login.yaml` 运行
时是 `clearState: true`，这会强制冷启动。\n\n让我测试一下：`test_simple_login.ya
ml` 在 `clearState: true` 后是否也白屏。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_VcyhPZcHBejSupI6JVynpmlW',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "cd /var/folders/lt/w3sl6z8j5v50vt2s6knvf5n00000g
n/T/dev-issue-55-815001100/.maestro && maestro test test_simple_login.yaml", "ti
meout": 120}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.4165687561035156,
    context_tokens=109201,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=657,
        output=512,
        input_cache_read=108544,
        input_cache_creation=0
    ),
    message_id='chatcmpl-YzAW3oM4NsGFLuePcANCWxz2',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_VcyhPZcHBejSupI6JVynpmlW',
    return_value=ToolReturnValue(
        is_error=False,
        output='Running on emulator-5556\n > Flow test_simple_login\nLaunch app 
"com.example.mobile" with clear state... COMPLETED\nAssert that "欢迎使用 Social
World" is visible... COMPLETED\nTap on "手机号"... COMPLETED\nInput text 1380013
8000... COMPLETED\nTap on "获取验证码"... COMPLETED\nTap on "验证码"... COMPLETE
D\nInput text 123456... COMPLETED\nHide Keyboard... COMPLETED\nTap on "登录"... 
COMPLETED\nAssert that (Optional) "完善资料" is visible... WARNED\n\n\n Warning:
Assertion is false: "完善资料" is visible\nTake screenshot login_after_click... 
COMPLETED\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=74)
ThinkPart(
    type='think',
    think='`test_simple_login.yaml` 现在失败了！`assertVisible: "完善资料" optio
nal: true` 返回了 WARNED。这说明登录没有成功，页面没有跳转。\n\n但之前同样的文件
是成功的！唯一的区别是时间：之前我运行成功时验证码刚写入 Redis，现在可能又过期了
？\n\n让我检查 Redis。',
    encrypted=None
)
ToolCall(
    type='function',
    id='tool_3ebu4ua3TutUL0mD6GSKk1Lk',
    function=FunctionBody(
        name='Shell',
        arguments='{"command": "docker exec sw-redis redis-cli GET sms_code:1380
0138000"}'
    ),
    extras=None
)
StatusUpdate(
    context_usage=0.4191398620605469,
    context_tokens=109875,
    max_context_tokens=262144,
    token_usage=TokenUsage(
        input_other=819,
        output=100,
        input_cache_read=109056,
        input_cache_creation=0
    ),
    message_id='chatcmpl-6kMEwAekXa6gUra2tKUpT3SC',
    plan_mode=False,
    mcp_status=None
)
ToolResult(
    tool_call_id='tool_3ebu4ua3TutUL0mD6GSKk1Lk',
    return_value=ToolReturnValue(
        is_error=False,
        output='\n',
        message='Command executed successfully.',
        display=[],
        extras=None
    )
)
StepBegin(n=75)

⏸️ 任务因长时间无响应被标记为中断

❌ e2e失败: tmux step runE2EAgent cancelled


---
*由 Agent Control Center 自动生成*