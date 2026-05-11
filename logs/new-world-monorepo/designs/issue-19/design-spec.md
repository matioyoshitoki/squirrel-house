# Issue #19 设计规范：UI 容器组件重做（UIPanel / UIDialog / UIToast / UIAlert）

> **需求类型**: Layer 1 原子组件重做  
> **目标**: 对四个核心容器组件进行系统性重做，使其严格对齐 PRD-004 接口设计与 `docs/design-docs/ui-style-guide/` 视觉规范  
> **关联 PRD**: [PRD-004-ui-component-library-redesign.md](../../../prd/PRD-004-ui-component-library-redesign.md)  
> **依赖 Issue**: #18（UI 基础设施已完成 `types.ts` / `utils.ts` / `index.ts`）

---

## 1. 架构总览

```
designs/issue-19/src/ui/
├── UIPanel.ts    # 通用面板（可模态/非模态）
├── UIDialog.ts   # 模态对话框（继承 UIPanel）
├── UIToast.ts    # 轻提示（带类型色条）
└── UIAlert.ts    # 提示条（带类型色条，自动消失）
```

### 1.1 单向依赖规则

```
UIDialog → UIPanel → utils.ts / types.ts
UIToast  → utils.ts / types.ts
UIAlert  → utils.ts / types.ts
```

**禁止循环依赖。**

---

## 2. 现有实现问题分析

### 2.1 UIPanel（`apps/phaser3/src/ui/UIPanel.ts`）

| 问题编号 | 问题描述 | 违反规范 | 修复方案 |
|----------|----------|----------|----------|
| P-01 | 缺少 `show()` / `hide()` 动画接口 | PRD-004 5.5 | 增加标准入场/退场动画 |
| P-02 | 关闭按钮使用 `×` 字符渲染 | F-01（emoji/Unicode 图标禁令）| 改用 `UIIcon` + `cancel-btn-0.png` 纹理 |
| P-03 | 关闭按钮 `pointerover` 使用魔法数字色值 `'#ffffff'` | F-07 | 改用 `UITheme.white` |
| P-04 | 缺少模态遮罩支持 | PRD-004 5.5 | 增加 `showOverlay` / `overlayAlpha` / `isModal` / `onOverlayClick` |
| P-05 | 缺少 `setTitle()` / `setCloseVisible()` 方法 | PRD-004 5.5 | 补充动态更新接口 |
| P-06 | 使用 `export default` | F-11 | 改为命名导出 `export class UIPanel` |
| P-07 | 未设置 `depth` | 1.4 深度规范 | 默认 `UIDepthLayer.Panel`（2000）|

### 2.2 UIDialog（`apps/phaser3/src/ui/UIDialog.ts`）

| 问题编号 | 问题描述 | 违反规范 | 修复方案 |
|----------|----------|----------|----------|
| D-01 | **未继承 UIPanel**，自行实现背景/标题/按钮 | PRD-004 5.6 | 改为 `extends UIPanel`，复用遮罩和动画 |
| D-02 | 遮罩使用 `0x000000` 纯黑 | F-03 | 改用 `UITheme.inkBlack` |
| D-03 | `setDepth(5000)` 错误 | 1.4 深度规范 | 改为 `UIDepthLayer.Dialog`（3000）|
| D-04 | `show()` / `hide()` 未执行 `killTweensOf` | 1.3 动画铁律 | 动画前强制 `killTweensOf` |
| D-05 | 缺少 `closeOnOverlay` 选项 | PRD-004 5.6 | 增加配置项 |
| D-06 | 缺少静态工厂方法 `alert()` | PRD-004 5.6 | 增加快速单按钮弹窗 |
| D-07 | 使用 `export default` | F-11 | 改为命名导出 |
| D-08 | 确认回调在 hide 动画**之前**执行 | PRD-004 5.6 | 改为 hide 动画**完成后**执行 `onConfirm` |

### 2.3 UIToast（`apps/phaser3/src/ui/UIToast.ts`）

| 问题编号 | 问题描述 | 违反规范 | 修复方案 |
|----------|----------|----------|----------|
| T-01 | 组件职责混杂：同时承载「进度通知」+「轻提示」| PRD-004 5.11 | 按 PRD-004 简化为纯消息 Toast，进度逻辑迁移到业务层 |
| T-02 | 没有左侧类型色条 | PRD-004 5.11 | 增加 4px 左侧色条，与 UIAlert 规范对齐 |
| T-03 | `hide()` 退场动画参数错误 | 3.4 动画标准 | 统一为 `y: 0 → -10`, `duration=200ms` |
| T-04 | 入场/退场未执行 `killTweensOf` | 1.3 动画铁律 | 动画前强制 `killTweensOf` |
| T-05 | `ToastManager` 非单例，无 `getInstance` | PRD-004 5.11 | 增加单例获取接口 |
| T-06 | `depth = 4000` 错误 | 1.4 深度规范 | 改为 `UIDepthLayer.Toast`（5000）|
| T-07 | 使用 `export default` | F-11 | 改为命名导出 |
| T-08 | `ToastType` 包含业务语义（`build`/`upgrade`/`remove`）| PRD-004 5.11 | 统一为 `UINotificationType`（info/success/warning/error）|

### 2.4 UIAlert（`apps/phaser3/src/ui/UIAlert.ts`）

| 问题编号 | 问题描述 | 违反规范 | 修复方案 |
|----------|----------|----------|----------|
| A-01 | 使用 Unicode 符号 `✓` `!` `×` `i` 作为类型图标 | F-01 | **删除图标文字**，仅用左侧 4px 色条标识类型 |
| A-02 | 缺少 `show()` 方法 | PRD-004 5.15 | 增加显式 show 接口（含入场动画）|
| A-03 | `hide()` 未执行 `killTweensOf` | 1.3 动画铁律 | 动画前强制 `killTweensOf` |
| A-04 | 未传入 `isAutoHide` 但写入了类型定义 | PRD-004 5.15 | 保留 `isAutoHide`，默认 `true` |
| A-05 | 使用 `export default` | F-11 | 改为命名导出 |
| A-06 | 缺少静态 `getTypeColor()` 方法 | PRD-004 5.15 | 增加公共静态方法 |

---

## 3. 组件详细设计

### 3.1 UIPanel — 通用面板

#### 视觉标准

| 属性 | 默认值 | 说明 |
|------|--------|------|
| 宽度 | `400px` | 可通过 Config 覆盖 |
| 高度 | `300px` | 可通过 Config 覆盖 |
| 背景 | `UITheme.inkBlack`（透明度 0.98）| 回退时使用 `createInkBackground()` |
| 边框 | `2px solid UITheme.goldBrown` | 回退时必须有边框 |
| 标题栏高度 | `60px` | 标题距顶部 `30px` |
| 标题字号 | `18px` | 加粗，居中，颜色 `UITheme.white` |
| 内容区内边距 | `20px` | 四边统一 |
| 关闭按钮 | `24px` | `UIIcon` 加载 `cancel-btn-0.png`，距边缘 `(width/2 - 20, -height/2 + 20)` |
| 关闭按钮 Hover | 边框颜色变为 `UITheme.cinnabarRed` | 或图标缩放 `1.1` |

#### 动画标准

| 方法 | 参数 | 说明 |
|------|------|------|
| `show()` | `scale: 0.8 → 1.0`, `alpha: 0 → 1`, `duration=200ms`, `ease='Back.out'` | 任何 show 前必须先 `killTweensOf(scene, this)` |
| `hide()` | `scale: 1.0 → 0.8`, `alpha: 1 → 0`, `duration=150ms`, `ease='Power2'` | 任何 hide 前必须先 `killTweensOf(scene, this)` |

#### 深度标准

- 非模态：`UIDepthLayer.Panel`（2000）
- 模态（`isModal=true`）：`UIDepthLayer.Dialog`（3000）
- 遮罩层：始终 `UIDepthLayer.Overlay`（4000）

#### 接口设计

```typescript
export interface UIPanelConfig extends UIPositionConfig, Partial<UISizeConfig> {
  title?: string;
  texture?: string;
  bgColor?: number;
  strokeColor?: number;
  showClose?: boolean;
  showOverlay?: boolean;
  overlayAlpha?: number;
  isModal?: boolean;
  onClose?: () => void;
  onOverlayClick?: () => void;
  depth?: number;
}

export class UIPanel extends Phaser.GameObjects.Container {
  constructor(scene: Phaser.Scene, config: UIPanelConfig);
  show(): void;
  hide(): void;
  setTitle(title: string): void;
  getContentBounds(): Phaser.Geom.Rectangle;
  addContent(child: Phaser.GameObjects.GameObject): void;
  setCloseVisible(isVisible: boolean): void;
}
```

#### 禁止事项

- 禁止标题文字截断（标题宽度应 ≤ 面板宽度 - 80px）
- 禁止内容区直接贴边（必须有 20px 内边距）
- 禁止缺少 `show()` / `hide()` 标准动画接口
- 禁止使用 `UIText` 渲染 `×` 字符作为关闭按钮（F-01）
- 禁止遮罩使用纯黑 `0x000000`（F-03）

---

### 3.2 UIDialog — 模态对话框

#### 视觉标准

**必须继承 `UIPanel`**，复用其所有标准，并追加：

| 属性 | 标准 |
|------|------|
| 宽度 | `360px` |
| 高度 | `220px` |
| 模态遮罩 | 全屏 `Rectangle`，颜色 `UITheme.inkBlack`，透明度 `0.6` |
| 确认按钮 | 底部偏右，`bgColor: UITheme.cinnabarRed` |
| 取消按钮 | 底部偏左，`bgColor: UITheme.lightInk` |
| 按钮间距 | 确认与取消间距 `160px`（即确认 `x=80`，取消 `x=-80`）|
| 消息文字 | `14px`，颜色 `#cccccc`（即 `0xcccccc`），居中，自动换行 |

#### 动画标准

- 继承 `UIPanel` 的 show/hide 动画，禁止自行实现不同参数的动画

#### 行为标准

| 动作 | 行为 |
|------|------|
| 点击确认 | 先播放 `hide()` 动画，动画**完成后**再执行 `onConfirm` |
| 点击取消 | 先播放 `hide()` 动画，动画**完成后**再执行 `onCancel` |
| 点击遮罩 | 默认不关闭；当 `closeOnOverlay=true` 时，执行 `cancel()` 逻辑 |
| `depth` | `UIDepthLayer.Dialog`（3000），禁止 5000 |

#### 接口设计

```typescript
export interface UIDialogConfig extends Omit<UIPanelConfig, 'title' | 'showClose'> {
  title: string;
  message: string;
  confirmText?: string;
  cancelText?: string;
  showCancel?: boolean;
  confirmTexture?: string;
  cancelTexture?: string;
  onConfirm?: () => void;
  onCancel?: () => void;
  closeOnOverlay?: boolean;
  depth?: number;
}

export class UIDialog extends UIPanel {
  constructor(scene: Phaser.Scene, config: UIDialogConfig);
  static show(scene: Phaser.Scene, config: UIDialogConfig): UIDialog;
  static alert(scene: Phaser.Scene, title: string, message: string): UIDialog;
}
```

---

### 3.3 UIToast — 轻提示

#### 视觉标准

| 属性 | 默认值 | 说明 |
|------|--------|------|
| 宽度 | `350px` | 可通过 Config 覆盖 |
| 高度 | `56px` | 可通过 Config 覆盖 |
| 背景 | `UITheme.inkBlack`（透明度 0.95）| 回退时使用 `createInkBackground()` |
| 边框 | `2px solid UITheme.goldBrown` | 回退时必须有边框 |
| 左侧色条 | `4px` 宽，高度 `height - 8px` | 类型色，与 UIAlert 完全对齐 |
| 圆角 | 无 | Rectangle 直角 |
| 图标区 | 预留 `32px` | 如传入 `icon` 纹理则显示，否则留空 |
| 消息文字 | `14px`，颜色 `UITheme.white` | 左对齐 |

#### 类型色标准

| 类型 | 左侧色条 | 使用场景 |
|------|----------|----------|
| `info` | `UITheme.indigoBlue` | 普通提示 |
| `success` | `UITheme.stoneCyan` | 操作成功 |
| `warning` | `UITheme.goldBrown` | 警告提示 |
| `error` | `UITheme.cinnabarRed` | 错误提示 |

> 色条宽度**必须为 4px**（PRD-004 5.15 强制要求）。

#### 动画标准

| 方法 | 参数 | 说明 |
|------|------|------|
| `show()` | `y: +20 → 0`, `alpha: 0 → 1`, `duration=300ms`, `ease='Back.out'` | 必须先 `killTweensOf(scene, this)` |
| `dismiss()` | `y: 0 → -10`, `alpha: 1 → 0`, `duration=200ms`, `ease='Power2'` | 必须先 `killTweensOf(scene, this)` |

#### ToastManager 标准

| 行为 | 标准 |
|------|------|
| 单例获取 | `ToastManager.getInstance(scene)` |
| 队列上限 | 最多同时显示 3 个 Toast |
| 队列满策略 | 最早的 Toast 提前 `dismiss()` |
| 垂直间距 | 相邻 Toast 间距 `64px` |
| 自动消失 | 默认 `duration=3000ms` |
| 重新排列 | Toast 消失后，剩余 Toast 以 `duration=200ms` `Power2` 补位 |

#### 接口设计

```typescript
export interface UIToastConfig extends UIPositionConfig, Partial<UISizeConfig> {
  message: string;
  type?: UINotificationType;
  duration?: number;
  icon?: string;
  texture?: string;
  depth?: number;
}

export class UIToast extends Phaser.GameObjects.Container {
  constructor(scene: Phaser.Scene, config: UIToastConfig);
  show(): void;
  dismiss(): void;
  static show(scene: Phaser.Scene, config: UIToastConfig): UIToast;
}

export class ToastManager {
  static getInstance(scene: Phaser.Scene): ToastManager;
  show(config: UIToastConfig): UIToast;
  clear(): void;
  getCount(): number;
}
```

#### 禁止事项

- 禁止返回 `✓`、`!`、`×`、`i` 等 Unicode 符号作为类型图标（F-01）
- 禁止缺少 `depth` 设置，必须 `UIDepthLayer.Toast`（5000+）
- 禁止在 Toast 内部处理业务进度逻辑（如建造倒计时）

---

### 3.4 UIAlert — 提示条

#### 视觉标准

| 属性 | 默认值 | 说明 |
|------|--------|------|
| 宽度 | `280px` | 可通过 Config 覆盖 |
| 高度 | `44px` | 可通过 Config 覆盖 |
| 背景 | `UITheme.inkBlack`（透明度 0.95）| 回退时使用 `createInkBackground()` |
| 圆角 | 无 | Rectangle 直角 |
| 左侧色条 | `4px` 宽，高度 `height - 4px` | 类型色 |
| 消息文字 | `14px`，颜色 `UITheme.white` | 居中 |

#### 类型色标准

与 UIToast 完全对齐：

| 类型 | 左侧色条 | 使用场景 |
|------|----------|----------|
| `info` | `UITheme.indigoBlue` | 普通提示 |
| `success` | `UITheme.stoneCyan` | 操作成功 |
| `warning` | `UITheme.goldBrown` | 警告提示 |
| `error` | `UITheme.cinnabarRed` | 错误提示 |

#### 动画标准

| 方法 | 参数 | 说明 |
|------|------|------|
| `show()` | `y: +20 → 0`, `alpha: 0 → 1`, `duration=200ms`, `ease='Back.out'` | 必须先 `killTweensOf(scene, this)` |
| `hide()` | `alpha: 1 → 0`, `y: 0 → -10`, `duration=200ms`, `ease='Power2'` | 必须先 `killTweensOf(scene, this)` |
| 自动消失 | 延迟 `duration`（默认 3000ms）后触发 `hide()` | — |

#### 接口设计

```typescript
export interface UIAlertConfig extends UIPositionConfig, Partial<UISizeConfig> {
  message: string;
  type?: UINotificationType;
  duration?: number;
  isAutoHide?: boolean;
  depth?: number;
}

export class UIAlert extends Phaser.GameObjects.Container {
  constructor(scene: Phaser.Scene, config: UIAlertConfig);
  show(): void;
  hide(): void;
  static show(scene: Phaser.Scene, config: UIAlertConfig): UIAlert;
  static getTypeColor(type: UINotificationType): number;
}
```

#### 禁止事项

- 禁止使用 Unicode 符号作为类型图标（F-01）
- 禁止缺少 `depth` 设置，必须 `UIDepthLayer.Toast`（5000+）
- 禁止色条宽度不等于 4px

---

## 4. 状态定义

### 4.1 UIPanel 状态机

```
[HIDDEN] ──show()──▶ [SHOWING] ──动画完成──▶ [VISIBLE]
   ▲                                              │
   │                                              │
   └────────hide()────────动画完成───────────────┘
```

- `SHOWING` / `HIDING` 态中再次调用 `show()` / `hide()` 时，先 `killTweensOf` 再播放新动画
- 模态模式下，遮罩层的显隐与面板同步

### 4.2 UIDialog 状态机

继承 UIPanel 状态机，追加：

```
[VISIBLE] ──点击确认──▶ [CONFIRMING] ──hide完成──▶ [HIDDEN] → onConfirm()
[VISIBLE] ──点击取消──▶ [CANCELING]  ──hide完成──▶ [HIDDEN] → onCancel()
```

### 4.3 UIToast 状态机

```
[HIDDEN] ──show()──▶ [SHOWING] ──动画完成──▶ [VISIBLE] ──延迟──▶ [DISMISSING] ──动画完成──▶ [DESTROYED]
   │                                                                                          ▲
   │                                                                                          │
   └──────────────────────dismiss() 提前触发──────────────────────────────────────────────────┘
```

### 4.4 UIAlert 状态机

与 UIToast 类似，但 `isAutoHide=false` 时停留在 `[VISIBLE]` 直到显式调用 `hide()`。

---

## 5. Accessibility 设计

1. **色板对比度**：`UITheme.inkBlack` (#1a1a1a) 背景 + `UITheme.white` (#ffffff) 文字 = 对比度 14.5:1，远超 WCAG AAA（7:1）
2. **禁用状态**：`alpha ≤ 0.5` + `disableInteractive()`，视觉上与可用态有明显区分
3. **回退可见性**：纹理缺失时回退色块带 `2px` 边框，在复杂背景下可辨
4. **动画可感知**：入场/退场动画时长 ≤ 300ms，不会引起前庭功能障碍

---

## 6. 响应式 / 适配设计

1. **组件尺寸全部通过 Config 传入**，无硬编码绝对值
2. **居中对齐**：`UIDialog` 默认以 `scene.cameras.main.width/2` 和 `height/2` 居中
3. **Toast 锚点**：默认底部居中（`y = camera.height - 100`）
4. **Alert 锚点**：默认顶部居中（`y = 60`）
5. **1280×720 与 1920×1080 兼容**：所有位置通过相对值或配置传入

---

## 7. 素材决策日志

| UI 元素 | 搜索关键词 | 匹配结果 | 决策 | 加载 key |
|---------|-----------|----------|------|---------|
| 面板背景 | `ui-base` + `panel` | `login-plane.png`, `pop-ui.png`, `ui-small-panel.png` | 复用 | 由业务层传入 `texture` |
| 关闭按钮图标 | `ui-icon` + `cancel` | `cancel-btn-0.png` | 复用 | `cancel-btn-0` |
| Toast 背景 | `ui-base` + `panel` | `ui-small-panel.png` | 复用 | 由业务层传入 `texture` |
| Toast/Alert 类型图标 | — | — | **不使用图标** | 仅用左侧 4px 色条标识类型 |

---

## 8. 命名规范

| 层级 | 命名规则 | 示例 |
|------|----------|------|
| 布尔参数 | `isXxx` 或 `showXxx` | `isModal`, `showClose`, `isAutoHide` |
| 回调函数 | `onXxx` | `onClose`, `onConfirm`, `onOverlayClick` |
| 颜色参数 | `xxxColor` | `bgColor`, `strokeColor` |
| 组件文件 | PascalCase | `UIPanel.ts`, `UIDialog.ts` |
| 导出方式 | 命名导出 | `export class UIPanel` |

---

## 9. 验收标准（供开发 Agent 参考）

### 9.1 UIPanel

- [ ] `show()` / `hide()` 标准动画正确执行，先 `killTweensOf`
- [ ] 关闭按钮使用 `UIIcon` + `cancel-btn-0.png`，Hover 时颜色变为 `UITheme.cinnabarRed`
- [ ] `showOverlay=true` 时显示全屏遮罩，颜色 `UITheme.inkBlack`，透明度 `overlayAlpha`
- [ ] `isModal=true` 时 depth 为 3000，否则为 2000
- [ ] `getContentBounds()` 返回正确内容区矩形（含标题栏 60px 偏移）
- [ ] 无 `export default`

### 9.2 UIDialog

- [ ] 继承 `UIPanel`，复用 `show()` / `hide()` / 遮罩逻辑
- [ ] 遮罩颜色为 `UITheme.inkBlack`，非 `0x000000`
- [ ] `depth = UIDepthLayer.Dialog`（3000）
- [ ] 点击确认后，先 `hide()` 动画完成，再执行 `onConfirm`
- [ ] `closeOnOverlay=true` 时点击遮罩触发 `cancel()` 逻辑
- [ ] 提供静态工厂方法 `alert()`（单按钮快速提示）
- [ ] 无 `export default`

### 9.3 UIToast

- [ ] 左侧 4px 类型色条正确显示
- [ ] `show()` 入场动画：`y +20 → 0`, `alpha 0 → 1`, `300ms`, `Back.out`
- [ ] `dismiss()` 退场动画：`y 0 → -10`, `alpha 1 → 0`, `200ms`, `Power2`
- [ ] 动画前均执行 `killTweensOf`
- [ ] `ToastManager.getInstance(scene)` 返回单例
- [ ] 队列满（3 个）时最早 Toast 提前 dismiss
- [ ] `depth = UIDepthLayer.Toast`（5000）
- [ ] 无 `export default`

### 9.4 UIAlert

- [ ] 无 Unicode 类型图标，仅用左侧 4px 色条
- [ ] `show()` / `hide()` 动画正确，先 `killTweensOf`
- [ ] `isAutoHide=true`（默认）时，延迟 3000ms 自动触发 `hide()`
- [ ] `getTypeColor('info') === UITheme.indigoBlue`
- [ ] `depth = UIDepthLayer.Toast`（5000）
- [ ] 无 `export default`

### 9.5 通用

- [ ] `pnpm --filter phaser3 lint` 通过
- [ ] `pnpm --filter phaser3 type-check` 通过
- [ ] 无魔法数字色值（`grep -rn "0x[0-9a-fA-F]\{6\}"` 仅命中 `utils.ts`）
- [ ] 无 `export default`
