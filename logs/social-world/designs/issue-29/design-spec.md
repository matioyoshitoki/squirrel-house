# Issue #29 设计规范：开发模式悬浮抽屉（Dev Tools Drawer）

> 关联 PRD：[prd/v1-dev-tools-drawer.md](../../prd/v1-dev-tools-drawer.md)  
> 关联 BDD：[tests/bdd/dev_tools_drawer.feature](../../tests/bdd/dev_tools_drawer.feature)  
> 技术栈：Flutter (apps/mobile)

---

## 1. 页面/场景清单

| 场景 | 说明 | 对应 Widget | 状态覆盖 |
|------|------|-------------|----------|
| S1. Release 模式 | 入口完全不可见，Widget 不在渲染树中 | `SocialWorldApp` (条件编译) | — |
| S2. Debug 常态页面 | 任意页面右侧边缘显示半透明悬浮按钮 | `DevToolsToggleButton` | 正常态 |
| S3. 抽屉展开 | 点击悬浮按钮，右侧滑出工具面板 | `DevToolsDrawer` | 正常态 |
| S4. URL 编辑与保存 | 抽屉内修改 Server Base URL 并持久化 | `ServerUrlTool` (Drawer 内部) | 正常态、校验错误态、保存成功态 |
| S5. 抽屉关闭 | 点击遮罩或关闭按钮，抽屉收起，页面恢复交互 | `DevToolsDrawer` + Overlay | 正常态 |

---

## 2. 视觉系统

全部复用项目现有 `DesignTokens`，**禁止硬编码**。

### 2.1 色板

| Token | 色值 | 使用场景 |
|-------|------|----------|
| `DesignTokens.primary` | `#6750A4` | 悬浮按钮背景、抽屉标题、保存按钮、聚焦边框 |
| `DesignTokens.background` | `#FFFBFE` | 页面背景、抽屉背景、输入框填充色 |
| `DesignTokens.textPrimary` | `#1C1B1F` | 正文、标题、输入框文字 |
| `DesignTokens.error` | `#B3261E` | 校验失败边框、错误提示文字 |

**场景专用色（由 Token 组合）**：
- 悬浮按钮背景：`DesignTokens.primary.withOpacity(0.85)`（Debug 标识色，半透明不遮挡内容）
- 遮罩层：`Colors.black.withOpacity(0.15)`
- 禁用按钮背景：`Colors.grey`（Material 默认禁用色）

### 2.2 字体

| Token | 值 | 使用场景 |
|-------|-----|----------|
| `DesignTokens.fontSizeTitle` | `22.0` | 抽屉标题 "开发工具" |
| `DesignTokens.fontSizeBody` | `16.0` | 按钮文字、输入框文字、标签文字 |
| `DesignTokens.fontSizeCaption` | `12.0` | 错误提示、辅助说明 |

### 2.3 间距与尺寸

| Token | 值 | 使用场景 |
|-------|-----|----------|
| `DesignTokens.spacingMd` | `16.0` | 页面边距、抽屉内边距、卡片内边距 |
| `DesignTokens.spacingSm` | `12.0` | 标签与输入框间距、图标与文字间距 |
| `DesignTokens.spacingXs` | `8.0` | 紧凑元素间距 |
| `DesignTokens.radiusMd` | `12.0` | 卡片圆角、输入框圆角、按钮圆角 |
| `DesignTokens.radiusLg` | `16.0` | 抽屉顶部圆角（若采用底部 Sheet 备用方案） |

**固定尺寸**：
- 悬浮按钮直径：`56.0`
- 抽屉宽度：`屏幕宽度 × 0.8`，最大 `320.0`（适配 360dp ~ 414dp 屏幕）
- 输入框高度：`56.0`（同现有 `SwTextField` 视觉高度）
- 保存按钮高度：`48.0`
- 抽屉头部高度：`56.0`

---

## 3. 组件清单

### 3.1 复用现有原子/分子组件

| 组件 | 来源 | 使用位置 | 备注 |
|------|------|----------|------|
| `SwTextField` | `packages/design-system` | `ServerUrlTool` 内 URL 输入框 | 设置 `keyboardType: TextInputType.url` |
| `SwButton` (filled) | `packages/design-system` | `ServerUrlTool` 内 "保存配置" 按钮 | `isDisabled` 由校验状态驱动 |
| `SwCard` | `packages/design-system` | `ServerUrlTool` 外层容器 | 包裹输入框与按钮，形成独立工具卡片 |

### 3.2 新增组件（Feature 级，不进入 design-system atoms）

| 组件 | 路径建议 | 职责 |
|------|----------|------|
| `DevToolsToggleButton` | `apps/mobile/lib/presentation/widgets/dev_tools_toggle_button.dart` | 常驻半透明悬浮按钮，右侧居中，点击触发 `onToggle` |
| `DevToolsDrawer` | `apps/mobile/lib/presentation/widgets/dev_tools_drawer.dart` | 右侧滑出抽屉，含遮罩、头部、内容区；支持手势关闭 |
| `ServerUrlTool` | `apps/mobile/lib/presentation/widgets/server_url_tool.dart` | 抽屉内的 Server Base URL 编辑卡片，含输入框、校验、保存按钮 |
| `DevToolsWrapper` | `apps/mobile/lib/presentation/widgets/dev_tools_wrapper.dart` | 包裹在 `MaterialApp.router` 外层，负责在 Debug 模式下注入 Toggle + Drawer |

### 3.3 Widgetbook 注册要求

新增组件需在 `packages/design-system/widgetbook/main.dart` 的 **Features** 或 **Molecules** 目录下注册 use-case：

- `DevToolsDrawer`：默认态（含 `ServerUrlTool`）
- `ServerUrlTool`：正常态、错误态（URL 无效）、保存中态（`isLoading = true`）

> 注：`DevToolsToggleButton` 视觉简单，可与 `DevToolsDrawer` 合并在一个 use-case 中展示交互流程。

---

## 4. 图片/图标资源

本需求**无需新增图片资源**，全部使用 Material Icon 或内联 SVG。

| 图标 | 来源 | 使用位置 | 说明 |
|------|------|----------|------|
| `Icons.settings` / `Icons.construction` | Material | `DevToolsToggleButton` 中央 | 开发工具标识，建议 `Icons.construction` 更直观 |
| `Icons.public` | Material | `ServerUrlTool` 标签前缀 | URL/网络标识 |
| `Icons.close` | Material | `DevToolsDrawer` 头部右侧 | 关闭抽屉 |

---

## 5. 交互说明

### 5.1 悬浮按钮（DevToolsToggleButton）

- **位置**：`Alignment.centerRight`，右侧距边 `8.0`
- **尺寸**：`56.0 × 56.0`，圆形
- **背景**：`DesignTokens.primary.withOpacity(0.85)`，带 `BackdropFilter.blur(sigmaX: 4, sigmaY: 4)`
- **手势**：
  - 点击：展开抽屉
  - 长按/拖拽：**禁止拦截**，通过 `IgnorePointer` 的适当包裹或 `HitTestBehavior.translucent` 确保底层页面可正常滑动
- **可见性**：仅在 `kDebugMode == true` 时渲染；Release 模式下通过条件编译排除（建议用 `if (kDebugMode) ...` 或独立的 `DebugAppWrapper` Widget）

### 5.2 抽屉（DevToolsDrawer）

- **进入动画**：自右向左 `SlideTransition`，时长 `250ms`，曲线 `Curves.easeOutCubic`
- **退出动画**：自左向右 `SlideTransition`，时长 `200ms`，曲线 `Curves.easeInCubic`
- **关闭方式**：
  1. 点击左侧遮罩区域
  2. 点击头部 `✕` 关闭按钮
  3. （可选）向右滑动手势关闭
- **布局**：`Row` 包含左侧 `Expanded` 遮罩 + 右侧固定宽度抽屉面板

### 5.3 URL 编辑与保存（ServerUrlTool）

- **默认值**：显示当前生效的 `ServerBaseUrl`，初始为 `const String.fromEnvironment('API_BASE_URL', defaultValue: 'http://localhost:3000/api/v1')`
- **校验规则**：
  - 非空校验
  - URL 格式校验：必须以 `http://` 或 `https://` 开头，使用 `Uri.tryParse(url)?.hasAbsolutePath == true` 或正则 `^https?://.+`
- **错误反馈**：
  - 输入框边框变红（`DesignTokens.error`）
  - 输入框下方显示 `DesignTokens.fontSizeCaption` 红色错误文案
  - "保存" 按钮置灰（`isDisabled = true`）
- **保存成功**：
  - 显示底部 `SnackBar`："保存成功，新配置已生效"
  - （待产品确认）是否立即热重载 `DioClient` 的 `baseUrl`，或提示重启 App

### 5.4 BLoC / Cubit 设计

采用 **Cubit** 管理局部状态（状态简单，无复杂事件流）。

#### State

```dart
// apps/mobile/lib/presentation/blocs/dev_tools/dev_tools_state.dart
@freezed
class DevToolsState with _$DevToolsState {
  const factory DevToolsState({
    @Default('') String serverBaseUrl,
    @Default('') String urlError,
    @Default(false) bool isSaving,
    @Default(false) bool saveSuccess,
  }) = _DevToolsState;
}
```

#### Cubit

```dart
// apps/mobile/lib/presentation/blocs/dev_tools/dev_tools_cubit.dart
class DevToolsCubit extends Cubit<DevToolsState> {
  DevToolsCubit(this._repository)
      : super(const DevToolsState());

  final DevToolsRepository _repository;

  Future<void> load() async {
    final url = await _repository.getServerBaseUrl();
    emit(state.copyWith(serverBaseUrl: url));
  }

  void onUrlChanged(String url) {
    emit(state.copyWith(
      serverBaseUrl: url,
      urlError: _validateUrl(url),
      saveSuccess: false,
    ));
  }

  String _validateUrl(String url) {
    if (url.trim().isEmpty) return '请输入服务器地址';
    if (!url.trim().startsWith(RegExp(r'^https?://'))) {
      return '请输入有效的 HTTP/HTTPS URL';
    }
    return '';
  }

  Future<void> save() async {
    final error = _validateUrl(state.serverBaseUrl);
    if (error.isNotEmpty) return;
    emit(state.copyWith(isSaving: true));
    await _repository.saveServerBaseUrl(state.serverBaseUrl);
    emit(state.copyWith(isSaving: false, saveSuccess: true));
  }
}
```

#### Repository 接口（Domain 层）

```dart
// apps/mobile/lib/domain/repositories/dev_tools_repository.dart
abstract class DevToolsRepository {
  Future<String> getServerBaseUrl();
  Future<void> saveServerBaseUrl(String url);
}
```

#### Repository 实现（Data 层）

```dart
// apps/mobile/lib/data/repositories/dev_tools_repository_impl.dart
class DevToolsRepositoryImpl implements DevToolsRepository {
  DevToolsRepositoryImpl(this._prefs);
  final SharedPreferences _prefs;

  static const _key = 'dev_tools_server_base_url';
  static const _default = 'http://localhost:3000/api/v1';

  @override
  Future<String> getServerBaseUrl() async {
    return _prefs.getString(_key) ?? _default;
  }

  @override
  Future<void> saveServerBaseUrl(String url) async {
    await _prefs.setString(_key, url);
  }
}
```

---

## 6. 待产品确认

1. **保存后生效时机**：URL 保存后是否要求立即重启 App？若立即生效，需设计 `DioClient` 热重载 `baseUrl` 的机制；若需重启，需在保存成功后追加提示文案 "重启后生效"。
2. **抽屉关闭后输入态保留**：关闭抽屉后再次打开，是否保留上一次的输入内容（未保存）？建议保留，避免误操作丢失编辑内容。
3. **是否支持 URL 历史记录**：是否需要保存最近使用的 3~5 个 URL 供快速切换？本期 PRD 未要求，若后续有需求可扩展 `ServerUrlTool` 为下拉选择框。

> **当前假设**：保存后立即生效（通过更新 `DioClient` 单例的 `baseUrl`）；关闭抽屉保留输入态；不支持历史记录。

---

## 7. 需更新文档

### 7.1 模块注册

- **`docs/modules/INDEX.md`**：新增 `dev-tools.md` 模块条目，说明为 Debug 专用工具模块，关联 PRD `v1-dev-tools-drawer.md`。

### 7.2 新增模块实现文档

- **`docs/modules/dev-tools.md`**（新建）：按模块文档规范编写，包含：
  - 时序图：点击悬浮按钮 → 打开抽屉 → 修改 URL → 校验 → 持久化到 SharedPreferences → 更新 DioClient
  - 关键文件索引：`DevToolsCubit`、`DevToolsRepository`、`DevToolsToggleButton`、`DevToolsDrawer`、`ServerUrlTool`
  - 条件编译说明：如何在 `app.dart` 中通过 `kDebugMode` 注入 `DevToolsWrapper`
  - 边界情况：Release 模式完全排除、URL 格式校验失败、SharedPreferences 读写异常

### 7.3 Flutter 编码规范补充

- **`docs/design-docs/flutter.md`**：新增 "Debug-Only Widget 条件编译模式" 小节，明确：
  - 所有调试工具 Widget 必须包裹在 `if (kDebugMode) ...` 中
  - 禁止在 Release 模式下将调试 Widget 编译进产物
  - 调试工具的状态管理允许不遵循 `BLoC` 而使用 `StatefulWidget` 或 `ValueNotifier`（因仅开发期使用），但建议统一用 Cubit 保持代码可读性

### 7.4 UI 规格文档归档

- **`docs/design-docs/ui-specs/dev-tools/dev-tools-drawer-wireframe.md`**（新建）：归档本设计资产（引用 `designs/issue-29/wireframe.svg` 和 `mockup.html`），记录设计评审结论和变更历史，供后续迭代参考。

---

## 8. Golden File / 测试建议

- `DevToolsDrawer` 默认态截图（`test/goldens/dev_tools_drawer_default.png`）
- `ServerUrlTool` 错误态截图（`test/goldens/server_url_tool_error.png`）
- Widget 测试：`DevToolsToggleButton` 点击事件、`ServerUrlTool` 校验逻辑、`DevToolsDrawer` 开闭动画状态
