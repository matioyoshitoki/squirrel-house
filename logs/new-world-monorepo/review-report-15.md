## 审查结果

| 检查项 | 状态 | 说明 |
|--------|------|------|
| 全链路完整性 | ✅ | 本 PR 为纯前端 UI 交互改进，仅涉及 Phaser3 客户端，无需后端/DB 变更 |
| 规范符合性 | ❌ | 新增代码使用魔法数字色值，违反 `docs/design-docs/ui-style-guide/01-foundation.md` F-07 |
| 契约一致性 | N/A | 未涉及共享类型包或前后端 DTO 变更 |
| 测试覆盖 | ❌ | 新增 `showMoveConfirm` 功能未补充单元测试；`BuildingActionBar` 无测试覆盖 |
| 文档同步 | ❌ | 删除了 `ui-demo-screenshot.png` 无说明；`city-building.md` 中的交互架构描述与代码不一致 |
| UI/设计一致性 | ❌ | `designs/issue-14/` 缺失，且 PR 非纯 bug 修复，引入新 UI 按钮与交互流程，无法验证设计合规性 |

## 问题详情

1. **问题：前端/UI 变更缺少设计资产，无法验证实现是否符合设计规范**
   **建议**：`designs/issue-14/` 目录不存在，`design-spec.md` 缺失。根据审查规则，PR 引入了新的「确认移动/取消移动」按钮及交互流程（从 Building 内部 OK 按钮重构为 BuildingActionBar 统一管理的双按钮），不满足纯 bug 修复例外条件。请在 `designs/issue-14/` 下补充设计资产（至少包含 `design-spec.md`），或在 PR 描述中提供视觉验收截图证据。

2. **问题：新增代码使用魔法数字色值，违反 UI Foundation 规范**
   **建议**：修改 `apps/phaser3/src/city-builder/ui/BuildingActionBar.ts`，将 `createMoveConfirmButtons()` 中的魔法数字替换为 `UITheme` 常量：
   - `0x00aa00` → `UITheme.stoneCyan` 或相近语义颜色
   - `0x00cc00` → 使用 `UITheme` Hover 变体或计算得出
   - `0xff4444` → `UITheme.cinnabarRed`
   - `0xff6666` → 使用 `UITheme.cinnabarRed` 的 Hover 变体
   - `0xffffff` → `UITheme.white`
   具体规范参见 `docs/design-docs/ui-style-guide/01-foundation.md` 第 1.1 节。

3. **问题：新增功能缺少单元测试**
   **建议**：在 `apps/phaser3/src/city-builder/ui/__tests__/` 下新建 `BuildingActionBar.test.ts`，至少覆盖以下场景：
   - `showMoveConfirm()` 被调用后操作栏可见且包含确认/取消按钮
   - 点击确认按钮触发 `onConfirm` 回调
   - 点击取消按钮触发 `onCancel` 回调
   - `hide()` 后操作栏不可见
   可参考同目录下 `RemoveConfirmDialog.test.ts` 的测试写法。

4. **问题：模块文档与代码实现不一致**
   **建议**：更新 `docs/modules/client/city-building.md`：
   - 状态转换图中 `edit` 状态的说明"显示操作按钮(完成/取消)"应更新为"显示确认/取消按钮"
   - 组件架构图中 Building 内部包含"操作按钮(Buttons)"的描述已过时，需说明操作栏现由 `UIManager` / `BuildingActionBar` 统一管理
   同时，PR 删除了 `docs/design-docs/ui-style-guide/ui-demo-screenshot.png`，请在相关文档或 PR 描述中说明删除原因。

5. **问题：`make lint` 在项目中存在历史失败（非本 PR 引入）**
   **说明**：`make lint` 因 Go 代码类型检查错误而失败，与本次前端变更无关。但建议维护者关注并清理 Go 侧 lint 问题，避免影响后续 PR 的机械检查。
