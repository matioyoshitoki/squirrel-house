## 审查结论

审查结论: **NEEDS_MAJOR_FIX**

> 本 PR 对 Layer 1 原子组件（UIPanel、UIDialog、UIToast、UIAlert 等）进行了系统性重做，核心容器组件的实现质量较高，接口与 PRD-004 基本对齐，`export default` 已全部清理，动画铁律（`killTweensOf`）和 `show()` 幂等性均得到落实。但存在多项影响规范符合性与维护性的问题，需在合并前修复。

---

## 阻塞问题（Blocking）

**无**

> 本 PR 未引入功能不可用、数据丢失或安全漏洞类问题。以下问题虽严重，但不属于 Blocking 定义范畴。

---

## 严重问题（Major）

### 1. UIList.ts 新增空状态占位使用硬编码色值 `#888888`
- **位置**: `apps/phaser3/src/ui/UIList.ts:115`
- **问题**: 本次 PR 新增的 `showEmptyState()` 方法中直接写入 `color: '#888888'`，违反 Foundation F-07「禁止在业务代码中直接使用魔法数字色值」。这是本 PR **新引入**的违规，旧代码中不存在此行。
- **影响**: 代码无法通过 `01-foundation.md` 定义的扫描脚本；风格维护困难。
- **修复**: 将 `color: '#888888'` 改为 `color: colorToHex(UITheme.mediumInk)`，并确认 `getInkTextStyle()` 是否更适用。
- **依据**: `docs/design-docs/ui-style-guide/01-foundation.md` 第 151 行（F-07）、`docs/frontend/agent-design-guide.md` Step 3.1 检查清单

### 2. ToastManager.getInstance 签名与 PRD-004 不一致，且存在单例行为缺陷
- **位置**: `apps/phaser3/src/ui/UIToast.ts:315`、`apps/phaser3/src/city-builder/ui/ProgressToast.ts:103`
- **问题**: PRD-004 5.11 明确定义 `static getInstance(scene: Phaser.Scene): ToastManager`，但代码实现为 `static getInstance(scene: Scene, startY?: number): ToastManager`。`ProgressToast.ts` 调用 `ToastManager.getInstance(scene, this.startY)` 传入第二个参数。由于 `ToastManager` 是单例（`WeakMap<Scene, ToastManager>`），当同一 Scene 先以 `startY=A` 获取实例后，再次以 `startY=B` 调用会返回旧实例且 `startY` 保持为 `A`，导致布局行为不符合预期。
- **影响**: 单例语义被破坏；不同调用方传入不同的 `startY` 时行为不可预期。
- **修复**: 
  1. 移除 `startY` 参数，严格对齐 PRD-004 接口：`static getInstance(scene: Phaser.Scene): ToastManager`
  2. 若 `ProgressToast` 需要自定义 `startY`，应在构造后通过实例方法（如 `setStartY()`）设置，或单独维护 `ProgressToastManager`
  3. 同步更新 `docs/modules/client/ui.md` 中 `ToastManager` 的接口声明
- **依据**: `prd/PRD-004-ui-component-library-redesign.md` 第 693 行

### 3. UIListItem 组件重做后缺少单元测试
- **位置**: `apps/phaser3/src/ui/UIListItem.ts`（无对应测试文件）
- **问题**: `UIListItem.ts` 在本次 PR 中进行了实质性修改：新增 `iconSize`、`selectedBorderTexture`、`bgTexture` 配置，将 `icon` 从 Unicode 文本改为纹理/回退 Rectangle，新增 `createFallbackRect` 回退逻辑。但 `apps/phaser3/src/ui/__tests__/` 目录下不存在 `UIListItem.test.ts`。
- **影响**: PRD-004 要求 Layer 1 原子组件「每个组件独立测试文件，覆盖初始化、交互状态、资源回退、内存清理」。缺少测试意味着后续重构缺乏回归保护。
- **修复**: 新增 `apps/phaser3/src/ui/__tests__/UIListItem.test.ts`，至少覆盖：
  - 默认初始化
  - `icon` 纹理存在 vs 缺失时的回退行为（验证 `console.warn`）
  - `setSelected(true/false)` 视觉变化
  - `setItemData` / `getItemData`
- **依据**: `prd/PRD-004-ui-component-library-redesign.md` 第 1172 行、第 1209 行

### 4. city-builder/ui/ 旧组件迁移不完整，大量硬编码色值残留
- **位置**: `apps/phaser3/src/city-builder/ui/*.ts`（多文件）
- **问题**: PRD-004 Phase 4 要求「`city-builder/ui/` 中的旧组件全部迁移使用新版 `src/ui/` 组件」。本 PR 仅：
  - 移除了全部 `export default`
  - 对 `TopInfoBar.ts` 做了部分 `UITheme` 迁移
  - 对 `ProgressToast.ts` 调整了 `ToastManager` 调用
  其余文件（`BuildingActionBar.ts`、`BuildingDetailPanel.ts`、`FishListBar.ts`、`RemoveConfirmDialog.ts`、`ReturnButton.ts`、`RightControlButtons.ts`、`UpgradeConfirmDialog.ts` 等）仍保留大量硬编码色值，例如：
  - `0x00aa00`、`0xff4444`、`0xffd700`
  - `#ff6b6b`、`#d4af37`、`#4aff4a`、`#888888`、`#aaaaaa`、`#cccccc`
- **影响**: 代码库处于半迁移状态，新旧风格混杂；Foundation 扫描脚本会持续报警。
- **修复**: 若本 PR 不打算完成 Phase 4，应在 PR 描述中明确说明 city-builder 迁移不在本次范围；若打算包含，需完成剩余文件的 `UITheme` 替换与旧模式清理。
- **依据**: `prd/PRD-004-ui-component-library-redesign.md` 第 1277 行、第 1209 行；`docs/design-docs/ui-style-guide/01-foundation.md` 第 216–228 行（Foundation 扫描脚本）

### 5. docs/modules/client/INDEX.md 未按 PRD-004 要求更新
- **位置**: `docs/modules/client/INDEX.md`
- **问题**: PRD-004 9.3 文档验收清单明确要求「`docs/modules/client/INDEX.md` 中 'ui' 公共模块说明已更新」。经 `git diff` 验证，该文件在本 PR 中无任何变更。
- **影响**: 开发者通过 INDEX.md 导航时无法获取新版组件库的信息入口，文档索引与实际内容脱节。
- **修复**: 在 `docs/modules/client/INDEX.md` 的「ui」模块说明中，更新为指向新版 `docs/modules/client/ui.md` 的链接，并简要说明 PRD-004 重做后的组件范围（Layer 1 + Layer 2）。
- **依据**: `prd/PRD-004-ui-component-library-redesign.md` 第 1218 行

---

## 轻微问题（Minor）

### 1. 多个已修改的 UI 组件遗留硬编码色值未修复
- **位置**: 
  - `apps/phaser3/src/ui/UIAvatar.ts:47` `textColor: config.textColor ?? '#ffffff'`
  - `apps/phaser3/src/ui/UIIconButton.ts:121` `color: '#ffffff'`
  - `apps/phaser3/src/ui/UIResourceDisplay.ts:37,39` `iconColor: config.iconColor ?? '#ffffff'`、`valueColor: config.valueColor ?? '#ffffff'`
  - `apps/phaser3/src/ui/UIProgressBar.ts:78` `color: '#ffffff'`
  - `apps/phaser3/src/ui/UIText.ts:25` `color: '#ffffff'`
  - `apps/phaser3/src/ui/UITooltip.ts:72,86` `color: '#ffffff'`、`color: '#cccccc'`
  - `apps/phaser3/src/ui/UIScrollView.ts:59` `this.maskGraphics.fillStyle(0xffffff, 1)`（mask  fill 可酌情保留）
  - `apps/phaser3/src/ui/UIInput.ts:68,148` `color: '#2c2c2c'`、`color: '#999999'`
- **问题**: 上述文件在本次 PR 的变更文件列表中（`git diff` 有输出），但 PR 未修复其中预存的硬编码色值。虽然这些色值在旧代码中已存在，但鉴于本 PR 的目标是「系统性重做」并「对齐视觉规范」，遗留这些问题会降低重构的完整性。
- **影响**: 代码规范性不足；Foundation 扫描脚本持续报警。
- **修复**: 在本 PR 或后续跟进 PR 中，统一将上述色值替换为 `colorToHex(UITheme.xxx)` 或 `getInkTextStyle()` 注入。
- **依据**: `docs/design-docs/ui-style-guide/01-foundation.md` 第 151 行（F-07）

### 2. designs/issue-19 与 apps/phaser3/src/ui/ 测试文件存在漂移
- **位置**: `designs/issue-19/src/ui/__tests__/` vs `apps/phaser3/src/ui/__tests__/`
- **问题**: `designs/` 目录缺少 `UIAvatar.test.ts`、`UIBadge.test.ts`、`UIButton.test.ts`、`UIIconButton.test.ts`、`UIList.test.ts`、`UIResourceDisplay.test.ts`、`UISegmentedControl.test.ts`、`utils.test.ts`，且同名的 `UIPanel.test.ts`、`UIDialog.test.ts`、`UIToast.test.ts`、`UIAlert.test.ts` 内容也与生产测试不同。
- **影响**: 较低；`designs/issue-19/design-spec.md` 已明确标注「实现以 `apps/phaser3/src/ui/` 为准」，后续 Agent 应不会误引用。
- **修复**（可选）: 在 `design-spec.md` 中追加一句说明「`designs/issue-19/src/ui/__tests__/` 为早期草稿，生产测试以 `apps/phaser3/src/ui/__tests__/` 为准」。
- **依据**: `designs/issue-19/design-spec.md` 第 12 行

### 3. .gitignore 新增 `package-lock.json` 与项目包管理器不一致
- **位置**: `.gitignore`
- **问题**: 项目使用 `pnpm`（存在 `pnpm-lock.yaml`、`pnpm-workspace.yaml`），新增 `package-lock.json` 到 `.gitignore` 虽无害，但暗示可能存在混用 npm 的风险。
- **影响**: 极低。
- **修复**: 无需修复，或改为统一忽略 `package-lock.json`、`yarn.lock`、`pnpm-lock.yaml`（若希望强制统一包管理器）。

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 功能完整性 | ✅ | Issue #19 范围内的 UIPanel / UIDialog / UIToast / UIAlert 重做完成，接口与 PRD-004 基本一致；UIButton / UIList / UIListItem 的改进也符合 PRD-004 要求 |
| 代码规范性 | ❌ | 新增硬编码色值（`UIList.ts` `#888888`）；`ToastManager.getInstance` 签名偏离 PRD；多个已修改组件遗留旧色值 |
| 全链路完整性 | ✅ | 本次变更仅涉及前端 UI 组件层，不涉及 DB / API / Mobile，组件内部数据流自洽 |
| 规范符合性 | ⚠️ | 核心容器组件对齐 `docs/design-docs/ui-style-guide/`，`export default` 已全部清理，`killTweensOf` 已落实，深度规范已修正；但 city-builder 迁移不完整，遮罩层已正确加入 Scene |
| 测试覆盖 | ❌ | UIListItem 无测试文件；部分组件（UIIcon、UIInput、UIScrollView、UIText、UITooltip、UIProgressBar）缺少测试，未达到 PRD-004 要求的 ≥80% 覆盖 |
| 文档同步 | ❌ | `docs/frontend/ui-components.md`、`docs/modules/client/ui.md`、`docs/design-docs/ui-style-guide/01-foundation.md`、`03-component.md` 已更新；但 `docs/modules/client/INDEX.md`、`docs/design-docs/ui-style-guide/INDEX.md` 未更新 |

---

## 风险与建议

1. **单例行为风险（ToastManager）**: `ToastManager.getInstance(scene, startY)` 的第二参数在单例模式下会被静默忽略，可能导致 `ProgressToast` 与其他 Toast 使用不同的 `startY` 时期望落空。建议立即对齐 PRD-004 单参数签名，避免后续业务代码依赖错误行为。

2. **city-builder 半迁移风险**: 当前 `city-builder/ui/` 中大量文件仍使用硬编码颜色，与新版 `src/ui/` 规范脱节。建议在下一个 PR 中专项完成 city-builder 迁移，或在 `docs/PLANS.md` / `exec-plans/tech-debt-tracker.md` 中记录此项技术债务，防止遗忘。

3. **测试债务**: PRD-004 要求全部 Layer 1 组件有独立测试。当前缺失 `UIListItem.test.ts` 及若干其他组件测试。建议在合并本 PR 后优先补齐 UIListItem 测试，再逐步覆盖剩余组件。

4. **设计资产漂移**: `designs/issue-19/` 与 `apps/phaser3/src/ui/` 的测试文件存在差异。虽然 `design-spec.md` 已声明以 `apps/` 为准，但建议在 Issue 关闭前将 `designs/issue-19/design-spec.md` 中的「差异说明」章节补充具体差异列表，降低后续 Agent 的困惑。
