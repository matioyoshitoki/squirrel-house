## 审查结果

| 检查项 | 状态 | 说明 |
|--------|------|------|
| 全链路完整性 | ✅ | Bug 修复范围合理，仅涉及 Mobile Widget + 测试，无需 DB/API 变更 |
| 规范符合性 | ✅ | 使用 `SwCard`/`SwTextField`/`SwButton` 等 design-system 组件，无硬编码 Token；`make check-docs` 通过 |
| 契约一致性 | N/A | 无共享类型或 API 契约变更 |
| 测试覆盖 | ⚠️ | 现有 Widget 测试已更新，新增回归测试 `entered text is not reset by state changes`；但未覆盖 `load()` 异步晚于首帧的边界场景 |
| 文档同步 | ✅ | 纯 Bug 修复，无新模块/新 API/架构变更，无需更新模块文档 |
| UI/设计一致性 | N/A | 纯 Bug 修复（修复输入框重置与复制粘贴），无新增 UI 设计或页面结构变更；设计资产存在性按 Bug Fix 例外标记为 N/A |

**综合结论：LGTM（建议合入，附一条改进建议）**

---

## 问题详情

### 1. 异步加载竞态边界（建议改进）

**问题**：`ServerUrlTool` 移除了 `BlocListener` 后，仅通过首帧 `addPostFrameCallback` 同步 `load()` 结果：

```dart
// apps/mobile/lib/presentation/widgets/server_url_tool.dart
WidgetsBinding.instance.addPostFrameCallback((_) {
  if (!mounted) return;
  final state = context.read<DevToolsCubit>().state;
  if (_controller.text.isEmpty && state.serverBaseUrl.isNotEmpty) {
    _controller.text = state.serverBaseUrl;
  }
});
```

若 `DevToolsRepository.getServerBaseUrl()`（SharedPreferences 读取）耗时超过一帧，`load()` 在 `addPostFrameCallback` 之后才 emit 新 state，此时 `_controller` 将永远不会被填充已持久化的 URL。

**影响评估**：实际使用场景中 `DevToolsWrapper` 在 App 启动时即创建 Cubit 并调用 `load()`，而 `ServerUrlTool` 位于默认关闭的抽屉内，用户需经过「看到悬浮按钮 → 点击」后才进入树，因此 `load()` 通常已提前完成。该竞态在生产使用中触发概率极低，但代码层面存在 correctness 缺陷。

**建议**：在保留 `BlocListener` 的前提下，增加一个 `_initialSyncDone` 标志，仅处理首次异步加载完成时的同步，后续用户编辑不再被重置：

```dart
bool _initialSyncDone = false;

@override
Widget build(BuildContext context) {
  return BlocListener<DevToolsCubit, DevToolsState>(
    listenWhen: (previous, current) =>
        !_initialSyncDone &&
        previous.serverBaseUrl != current.serverBaseUrl &&
        _controller.text != current.serverBaseUrl,
    listener: (context, state) {
      _controller.text = state.serverBaseUrl;
      _controller.selection = TextSelection.collapsed(
        offset: state.serverBaseUrl.length,
      );
      _initialSyncDone = true;
    },
    child: BlocBuilder<DevToolsCubit, DevToolsState>(...),
  );
}
```

---

## 补充说明

- **机械检查**：`make test` 与 `make lint` 因当前环境缺少 `jest`/`eslint`/`flutter` 命令而失败，属于环境依赖缺失，与本次 PR 变更无关。`make check-docs` 已通过。
- **设计资产**：`.design-assets.txt` 指向 `designs/issue-29/`（DevTools Drawer 原始设计），本次 PR 为 Issue #32 的纯行为 Bug 修复，未引入新视觉设计，故不强制要求 `designs/issue-32/design-spec.md`。
- **Golden 截图**：`apps/mobile/test/dev_tools_golden_test.dart` 与 `test/goldens/` 未被修改；由于仅移除了不渲染的 `BlocListener`，视觉输出应无变化，Golden 基准图无需更新。
