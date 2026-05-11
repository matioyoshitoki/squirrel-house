# PR #28 审查报告

**分支**: `feat/issue-18`  
**审查范围**: `apps/phaser3/src/ui/` 基础设施重建 + Layer 1 组件适配  
**关联 PRD**: PRD-004-ui-component-library-redesign.md  
**设计资产**: `designs/issue-18/design-spec.md`

---

## 审查结论

**审查结论: NEEDS_MAJOR_FIX**

本 PR 完成了 Issue #18 要求的 UI 基础设施重建（`types.ts`、`utils.ts`、`__tests__/setup.ts`、`index.ts`），并对现有 Layer 1 组件进行了命名规范统一和资源回退机制适配，整体方向正确。但**新增的 `UIListItem` 组件实现与 PRD-004 接口设计严重不符**，且存在多项 Foundation 规范违规（硬编码色值、荧光色、emoji/文本图标风险）。在修复以下 Blocking + Major 问题前，不建议合并。

---

## 阻塞问题（Blocking）

### 1. `UIListItem.icon` 实现与 PRD-004 及 Foundation 规范严重冲突

- **位置**: `apps/phaser3/src/ui/UIListItem.ts:52-57`
- **问题**: `icon` 参数被直接当作文本渲染：
  ```typescript
  this.iconText = this.scene.add.text(-width / 2 + 20, 0, icon, {
    fontSize: '18px',
  }).setOrigin(0.5);
  ```
  但 PRD-004 第 5.7 节明确定义 `icon?: string` 为**纹理键名**（配合 `iconSize?: number`），应使用 `scene.add.image()` 加载纹理。当前实现会导致业务层传入纹理键名时显示为乱码文字而非图标。同时，该实现违反了 Foundation 规范 **F-01**（禁止 emoji 作为图标/按钮），因为 `add.text` 的语义允许传入任意字符内容。
- **影响**: `UIListItem` 无法正常显示图标；若旧业务代码将 emoji 传入 `icon`，会在新组件体系下继续存在 F-01 违规。
- **修复**: 
  ```typescript
  // UIListItem.ts:52-57 替换为纹理渲染
  if (icon && hasTexture(this.scene, icon)) {
    this.iconImage = this.scene.add.image(-width / 2 + 20, 0, icon);
    this.iconImage.setDisplaySize(iconSize, iconSize);
  } else if (icon) {
    this.iconImage = createFallbackRect(this.scene, iconSize, iconSize, icon);
  }
  ```
- **依据**: PRD-004 第 5.7 节 `UIListItemConfig` 接口设计（`icon?: string; iconSize?: number`）；`docs/design-docs/ui-style-guide/01-foundation.md` 第 115 行 F-01。

---

## 严重问题（Major）

### 2. `UIAvatar.ts` 滥用 `createFallbackRect` 导致无意义告警

- **位置**: `apps/phaser3/src/ui/UIAvatar.ts:53`
- **问题**: 
  ```typescript
  this.bg = createFallbackRect(this.scene, size, size, 'avatar-bg', bgColor, strokeColor);
  ```
  `createFallbackRect` 的 JSDoc 明确说明其用途是"**当组件依赖的纹理缺失时**，使用此函数创建视觉可辨的回退矩形，并打印统一的 `console.warn` 提示"。但 `UIAvatar` 的背景矩形是**正常无纹理时的默认渲染**，并非纹理缺失场景。这会导致每次创建 `UIAvatar` 都输出 `[UI] fallback: avatar-bg`，淹没真正的纹理缺失告警。
- **影响**: 开发调试时无法通过 `console.warn` 有效识别真实的资源缺失问题。
- **修复**: 
  ```typescript
  this.bg = createInkBackground(this.scene, size, size, bgColor, strokeColor);
  ```
- **依据**: `apps/phaser3/src/ui/utils.ts` 第 180-206 行 `createFallbackRect` 的 JSDoc；对比 `UIButton.ts` 第 57-61 行正确区分了 "texture 缺失" 和 "无 texture" 两种分支。

### 3. `UITheme` 引入高饱和度荧光色，违反 Foundation 规范 F-04

- **位置**: `apps/phaser3/src/ui/utils.ts:34,36`
- **问题**: 
  ```typescript
  emeraldGreen: 0x4aff4a,  // 高饱和度荧光绿
  coralRed: 0xff6b6b,      // 高饱和度荧光红
  ```
  Foundation 规范 **F-04** 明确禁止"使用高饱和度荧光色（如 `#00ff00`、`#ff00ff`）"。旧代码中这些值以魔法数字形式存在于 `UIToast.ts`，但本 PR 将其提取为 `UITheme` 官方常量，相当于将违规色值**正式纳入色板体系**，扩大了影响面（任何组件都可引用）。
- **影响**: 与国风水墨的低饱和度视觉风格冲突；`success` 类型 Toast 使用了荧光绿而非规范已有的 `stoneCyan`。
- **修复**: 
  - 删除 `emeraldGreen`，`UIToast` 的 `success` 类型直接使用 `UITheme.stoneCyan`（Foundation 已定义的 success 色）。
  - 删除 `coralRed`，`UIToast` 的 `remove` 类型直接使用 `UITheme.cinnabarRed`。
- **依据**: `docs/design-docs/ui-style-guide/01-foundation.md` 第 119 行 F-04。

### 4. `UIListItem.ts` 使用硬编码颜色值

- **位置**: `apps/phaser3/src/ui/UIListItem.ts:68,74`
- **问题**: 
  ```typescript
  color: '#dddddd',  // title
  color: '#888888',  // subtitle
  ```
  违反 Foundation 规范 **F-07**（禁止在业务代码中直接使用魔法数字色值）。
- **修复**: 使用 `UITheme.white` 替代 `#dddddd`；使用 `UITheme.lightInk` 替代 `#888888`（或新增 `UITheme.secondaryText` 语义色）。
- **依据**: `docs/design-docs/ui-style-guide/01-foundation.md` 第 126 行 F-07。

### 5. `UIListItem.ts` subtitle 字号不在 Foundation 规范层级中

- **位置**: `apps/phaser3/src/ui/UIListItem.ts:75`
- **问题**: `fontSize: '11px'` 不在 Foundation 定义的字号层级中。规范层级为：标签 `10px`、辅助 `12px`、正文 `14px`、标题 `16/18/20px`。
- **修复**: 改为 `'12px'`（辅助层级）或 `'10px'`（标签层级）。
- **依据**: `docs/design-docs/ui-style-guide/01-foundation.md` 第 49-58 行字号层级表。

### 6. `UIListItem.ts` 作为新组件缺少独立单元测试

- **位置**: 缺失 `apps/phaser3/src/ui/__tests__/UIListItem.test.ts`
- **问题**: PRD-004 第 5.1 节目录结构和第 7.3 节测试要求均明确：Layer 1 原子组件"每个组件独立测试文件，覆盖初始化、交互状态、资源回退、内存清理"。`UIListItem` 是本 PR 新增的 Layer 1 组件，却无对应测试。
- **修复**: 新增 `__tests__/UIListItem.test.ts`，至少覆盖：
  - 初始化渲染（title、subtitle、icon 纹理/回退）
  - `setSelected(true/false)` 视觉变化
  - `setHover(true/false)` 视觉变化
  - 纹理缺失时触发 `console.warn('[UI] fallback: ...')`
- **依据**: PRD-004 第 5.1 节目录结构；第 7.3 节测试要求。

### 7. `UIListItem.ts` 接口与 PRD-004 不一致

- **位置**: `apps/phaser3/src/ui/UIListItem.ts`
- **问题**: 
  1. 方法命名：PRD-004 要求 `setData(data)` / `getData()`，实现为 `setItemData` / `getItemData`。
  2. 缺少参数：PRD-004 要求的 `selectedBorderTexture?: string`、`bgTexture?: string`、`iconSize?: number` 均未实现。
  3. 多余暴露：存在未在 PRD-004 中定义的 `getBackground()` 公共方法，破坏了封装。`UIList.ts` 通过此方法获取 Rectangle 绑定事件，属于实现层耦合。
- **修复**: 
  - 重命名 `setItemData` → `setData`，`getItemData` → `getData`。
  - 补充 `selectedBorderTexture`、`bgTexture`、`iconSize` 支持。
  - 移除 `getBackground()`；`UIList` 应通过委托事件（如 `listItem.on('pointerover', ...)`）或内部封装来处理交互。
- **依据**: PRD-004 第 5.7 节 `UIListItemConfig` / `UIListItem` 接口设计。

### 8. `UIList.ts` 使用魔法数字设置 depth

- **位置**: `apps/phaser3/src/ui/UIList.ts:50`
- **问题**: `this.setDepth(1000)` 使用了魔法数字，而项目中已定义 `UIDepthLayer.UI = 1000`。`01-foundation.md` 要求"所有 UI 组件必须显式设置 `depth`"。
- **修复**: `this.setDepth(UIDepthLayer.UI);`
- **依据**: `docs/design-docs/ui-style-guide/01-foundation.md` 第 89-98 行深度规范；`apps/phaser3/src/ui/types.ts` 第 48-61 行 `UIDepthLayer` 枚举。

### 9. `UISegmentedControl.ts` 仍存在硬编码颜色

- **位置**: `apps/phaser3/src/ui/UISegmentedControl.ts:169`
- **问题**: `'#aaaaaa'` 未使用 `UITheme` 常量。
- **修复**: `label.setColor(colorToHex(UITheme.lightInk));`
- **依据**: `docs/design-docs/ui-style-guide/01-foundation.md` 第 126 行 F-07。

---

## 轻微问题（Minor）

### 10. `UIBadge.ts` `primary` 类型 stroke 色变更无说明

- **位置**: `apps/phaser3/src/ui/UIBadge.ts:43`
- **问题**: `primary` 类型的 `stroke` 从 `0x3d8edb` 改为与 `bg` 相同的 `UITheme.stoneCyan`，导致选中态无边框区分。若是设计调整，应在 PR 描述或设计资产中说明意图。
- **修复**: 恢复独立的 stroke 色（如 `0x3d8edb` 的合规替代），或更新 `03-component.md` 说明此变更。

### 11. 现有测试深度不足

- **位置**: `apps/phaser3/src/ui/__tests__/UIButton.test.ts:39-42`
- **问题**: `disabled` 状态测试仅验证 `expect(button).toBeDefined()`，未验证 `alpha ≤ 0.5` 或 `disableInteractive()` 是否生效。本 PR 重命名了 `disabled` → `isDisabled`，是补充此类断言的好时机。
- **修复**: 增加断言验证 disabled 状态下 `target.alpha` 和交互状态。

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 功能完整性 | ⚠️ 部分完成 | Issue #18 基础设施（types/utils/setup/index）已完成。但 `UIListItem` 的 `icon` 实现与 PRD-004 接口严重不符，若业务传入纹理键名将无法渲染图标。 |
| 代码规范性 | ❌ 未通过 | 命名规范化（`isDisabled`/`showClose` 等）完成较好，但存在硬编码色值（`#dddddd`、`#aaaaaa`）、魔法数字 depth、荧光色常量等问题。 |
| 全链路完整性 | ✅ 通过 | 纯前端组件 PR，不涉及 DB/API/Mobile。`BuildingSelector` 已适配新类型导出。 |
| 规范符合性 | ❌ 未通过 | Foundation F-01（icon 文本渲染）、F-04（荧光色 `emeraldGreen`/`coralRed`）、F-07（硬编码色值）均有违反；PRD-004 接口一致性不足。 |
| 测试覆盖 | ❌ 未通过 | 新增 `UIListItem` 无独立测试；`utils.ts` 新增函数有基础测试但边界覆盖不足（如 `killTweensOf` 未验证 tween 是否真被 kill）。 |
| 文档同步 | ✅ 通过 | 本 PR 属于 PRD-004 Phase 1，设计资产 `design-spec.md` 已存在；项目级文档（`docs/`）按计划在 Phase 4 统一更新，当前无需同步。 |

---

## 风险与建议

1. **`UIListItem` 是当前最大风险点**：其 `icon` 实现方式决定了下游业务模块（Fish、Weapon、Market 等）是否能正确使用图标。建议**优先修复为纹理渲染模式**，并补充 `iconSize` 支持，否则 Phase 2/3 的组件开发将建立在错误基座上。

2. **`createFallbackRect` 的滥用会污染日志**：除 `UIAvatar` 外，建议全量扫描本 PR 中所有调用 `createFallbackRect` 的位置，确保只有"传入 texture 但纹理缺失"的场景才触发 warn。正常默认回退应使用 `createInkBackground`。

3. **荧光色债务不要固化到 `UITheme`**：`emeraldGreen` 和 `coralRed` 在旧代码中是魔法数字，尚可视为遗留问题；一旦进入 `UITheme` 官方常量，后续全项目都可能引用，清理成本更高。建议在合并前修正。

4. **合并前执行检查清单**：
   ```bash
   cd apps/phaser3
   pnpm lint
   pnpm type-check
   pnpm test        # 确认 utils.test.ts 全绿
   ```

5. **建议修复后复审的重点**：`UIListItem.test.ts` 是否覆盖了四态交互、资源回退、以及 `UIList` 与 `UIListItem` 的联动逻辑。
