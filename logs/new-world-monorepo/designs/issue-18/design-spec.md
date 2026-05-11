# Issue #18 设计规范：UI 基础设施重建（utils / types / setup / index.ts）

> **需求类型**: UI 基础设施设计  
> **目标**: 重建 `apps/phaser3/src/ui/` 的基础设施层，为 PRD-004 的 Layer 1 + Layer 2 组件重做提供统一、类型安全、可测试的基座  
> **关联 PRD**: [PRD-004-ui-component-library-redesign.md](../../../prd/PRD-004-ui-component-library-redesign.md)

---

## ⚠️ 重要冲突报告

`origin/feature/issue-18-ui-infrastructure` 分支中的设计产物与本项目实际技术栈**严重不符**。该分支产物面向 **React + Tailwind CSS** 的 Web UI 库（包含 `classnames.ts`、`dom.ts`、`useId` Hook 等），而 `new-world-monorepo` 的前端是 **Phaser3 游戏引擎**（`apps/phaser3/src/ui/` 中为 Phaser3 GameObjects 组件）。

**冲突细节：**

| 现有分支产物 | 项目实际需求 |
|-------------|-------------|
| React `AsProp`、`PolymorphicRef` 类型 | Phaser3 `Phaser.Scene` + `Phaser.GameObjects.Container` |
| Tailwind `cx` + `tailwind-merge` | Phaser3 `Graphics` / `Rectangle` + `UITheme` 色板常量 |
| React `ConfigProvider` + `useConfig` Hook | Phaser3 `Scene` 上下文 + 纯对象配置 |
| DOM 工具 `canUseDOM`、`ownerDocument` | Phaser3 纹理加载、GameObjects 生命周期管理 |
| CSS 变量主题切换 (`data-theme`) | Phaser3 `UITheme` 数字色值 + 运行时渲染 |

**决策**: 本设计资产**完全基于 Phaser3 技术栈**重新编写，覆盖 PRD-004 Phase 1 的四个基础设施文件。开发 Agent 应以本设计资产为准，忽略 `feature/issue-18-ui-infrastructure` 分支中的 React 方向产物。

---

## 1. 架构总览

```
src/ui/
├── types.ts              # 全局 UI 类型定义（零运行时开销）
├── utils.ts              # 运行时工具函数与色板常量
├── __tests__/
│   └── setup.ts          # UI 组件测试基础设施（MockScene + helper）
└── index.ts              # 统一入口（按组件粒度暴露）
```

### 1.1 单向依赖规则

```
index.ts → types.ts / utils.ts / 各组件
各组件 → types.ts / utils.ts
types.ts → 无依赖（纯类型）
utils.ts → types.ts（可选，类型引用）
__tests__/setup.ts → utils.ts / types.ts
```

**禁止循环依赖。**

---

## 2. 类型层设计 (`types.ts`)

### 2.1 设计决策

| 决策项 | 选择 | 原因 |
|--------|------|------|
| 位置/尺寸基类 | `UIPositionConfig` + `UISizeConfig` | PRD-004 统一命名，所有组件复用 |
| 回调签名 | `UIEventCallback<T>` | 统一事件回调类型，避免各组件自定义 |
| 品质等级 | `UIQualityLevel = 1 \| 2 \| ... \| 9` | 与 `quality-icon-{1~9}.png` 纹理体系对齐 |
| 通知类型 | `UINotificationType` | Toast / Alert / Badge 共享，避免重复定义 |
| 深度层级 | `UIDepthLayer` enum | 规范 Scene / UI / Panel / Dialog / Toast 的深度值 |

### 2.2 核心类型清单

- `UIPositionConfig` — `x?: number; y?: number`
- `UISizeConfig` — `width?: number; height?: number`
- `UIEventCallback<T>` — `(value: T) => void`
- `UIQualityLevel` — `1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9`
- `UINotificationType` — `'info' | 'success' | 'warning' | 'error'`
- `UIDepthLayer` — 深度层级枚举（Scene=0, UI=1000, Panel=2000, Dialog=3000, Overlay=4000, Toast=5000）
- `UIButtonState` / `UIInteractiveState` — 交互状态机类型
- `UITweenConfig` — 缓动配置类型（duration, ease, delay 等）

### 2.3 约束规则

1. **类型文件仅含 `export interface` / `export type` / `export enum`，禁止混入运行时逻辑或副作用**
2. **所有类型以 `export` 暴露，不使用 `export default`** — 强制命名导入
3. **复杂类型必须附 JSDoc**，说明典型使用场景
4. **与现有组件向后兼容** — 现有组件的 Config 接口通过 `extends` 复用新基类，不破坏已有代码

---

## 3. 工具层设计 (`utils.ts`)

### 3.1 设计决策

| 决策项 | 选择 | 原因 |
|--------|------|------|
| 色板常量 | `UITheme` 对象（数字色值） | Phaser3 原生接受 `number` 颜色，比字符串更高效 |
| 缓动配置 | `UITweens` 对象 | 统一所有组件动画参数，避免魔法数字 |
| 纹理检查 | `hasTexture(scene, key)` | 封装 `scene.textures.exists()` |
| 回退渲染 | `createInkBackground()` + `createFallbackRect()` | 国风边框 + 统一告警输出 |
| 文字样式 | `getInkTextStyle()` | 自动注入字体栈 `"Microsoft YaHei", "SimHei", sans-serif` |
| 深度分配 | `getNextDepth(layer)` | 同层级组件后创建自动递增 depth |

### 3.2 核心工具清单

| 工具名 | 职责 | 签名 |
|--------|------|------|
| `UITheme` | 9 色国风色板常量 | `const UITheme = { inkBlack: 0x1a1a1a, ... }` |
| `UITweens` | 标准缓动参数 | `popup / fade / hover / press / pulse / tweenValue` |
| `hasTexture(scene, key)` | 检测纹理是否可用 | `(scene: Scene, key: string) => boolean` |
| `createInkBackground(scene, w, h, fill, stroke)` | 国风边框矩形回退 | 返回 `Phaser.GameObjects.Rectangle` |
| `createFallbackRect(scene, w, h, color, stroke, key?)` | 统一回退矩形（带 warn） | 返回 `Rectangle`，纹理缺失时打印 `[UI] fallback: ${key}` |
| `getInkTextStyle(fontSize, color, options)` | 国风文字样式 | 返回 `Phaser.Types.GameObjects.Text.TextStyle` |
| `createInkDecorationLine(scene, width, color)` | 水墨装饰线 | 返回 `Rectangle`（高 2px） |
| `ensureRoundRectTexture(scene, key, radius)` | 生成圆角矩形纹理 | 用于需要圆角的组件 |
| `getNextDepth(layer, scene?)` | 获取下一个可用 depth | `(layer: UIDepthLayer, scene?) => number` |
| `colorToHex(color)` | 数字色值转 HEX 字符串 | `(color: number) => string` |
| `clamp(value, min, max)` | 数值裁剪 | 与 `Phaser.Math.Clamp` 行为一致 |

### 3.3 约束规则

1. **纹理缺失时必须打印 `console.warn('[UI] fallback: ${textureKey}')`** — PRD-004 强制要求
2. **回退元素尺寸与预期纹理一致，边框颜色与 `UITheme` 协调，透明度 `alpha ≥ 0.8`**
3. **回退状态下组件的交互功能必须保持可用**
4. **禁止在 utils 中引入任何 UI 组件或业务逻辑** — 保持纯工具层
5. **所有新增工具函数必须有 JSDoc**

---

## 4. 测试基础设施设计 (`__tests__/setup.ts`)

### 4.1 设计决策

| 决策项 | 选择 | 原因 |
|--------|------|------|
| MockScene | 扩展现有 `MockScene` | `src/test-setup.ts` 已提供基础 Mock，UI 层需要补充 Phaser3 GameObjects 专用 mock |
| 通用 helper | `simulateClick()` / `simulateHover()` / `simulatePointerOut()` | 封装事件触发，测试代码更简洁 |
| 纹理模拟 | `mockTexture(scene, key)` | 让 `hasTexture()` 在测试中返回 true |
| 回退断言 | `expectFallbackWarn(key)` | 验证纹理缺失时正确触发 `console.warn` |
| tween 断言 | `expectTweenCreated(scene, targets)` | 验证 tween 被正确创建 |

### 4.2 核心模块清单

- `UIMockScene` — 扩展 `MockScene`，增加 `textures.add(key)` 模拟纹理注册
- `simulateClick(target)` — 触发 `pointerdown` + `pointerup` 事件
- `simulateHover(target)` — 触发 `pointerover`
- `simulatePointerOut(target)` — 触发 `pointerout`
- `mockTexture(scene, key)` — 注册 mock 纹理，使 `hasTexture()` 返回 true
- `clearMockTextures(scene)` — 清空 mock 纹理注册表
- `expectFallbackWarn(key)` — Vitest spy 辅助，验证 `console.warn` 调用
- `createMockContainer()` — 快速创建可交互的 mock container

### 4.3 约束规则

1. **测试 helper 不依赖真实 Phaser3 引擎** — 完全基于 MockScene
2. **每个 helper 必须处理事件监听器清理** — 防止测试间状态泄漏
3. **`setup.ts` 必须导出 `cleanup()` 函数** — 供 `afterEach` 调用

---

## 5. 入口文件设计 (`index.ts`)

### 5.1 设计决策

| 决策项 | 选择 | 原因 |
|--------|------|------|
| 导出策略 | 子路径分层 + 统一入口 | 消费者可按需导入 |
| 类型导出 | `export { type XxxConfig }` | TypeScript 3.8+ 显式类型导出，避免值污染 |
| 顺序 | 先 utils/types，再 Layer 1，再 Layer 2 | 逻辑清晰，便于查阅 |

### 5.2 导出结构

```ts
// === 基础设施 ===
export {
  UITheme, hasTexture, ensureRoundRectTexture,
  getInkTextStyle, createInkBackground, createFallbackRect,
  createInkDecorationLine, UITweens,
  getNextDepth, colorToHex, clamp,
} from './utils';

export type {
  UIPositionConfig, UISizeConfig, UIEventCallback,
  UIQualityLevel, UINotificationType, UIDepthLayer,
  UIButtonState, UIInteractiveState, UITweenConfig,
} from './types';

// === Layer 1: 原子组件 ===
export { UIButton, type UIButtonConfig } from './UIButton';
// ... (16 个组件)

// === Layer 2: 通用交互组件 ===
export { UISlider, type UISliderConfig } from './UISlider';
// ... (6 个组件)
```

### 5.3 约束规则

1. **统一入口文件禁止引入任何副作用** — 避免无意的 bundle 膨胀
2. **所有组件必须显式导出其 Config 类型** — 便于业务层类型安全使用
3. **utils 中的内部辅助函数（如非公共 API）不导出** — 保持接口精简

---

## 6. 状态定义

### 6.1 交互状态机

所有交互组件必须支持的状态机（已由 `types.ts` 中的 `UIInteractiveState` 建模）：

```
[NORMAL] ──pointerover──▶ [HOVER] ──pointerout──▶ [NORMAL]
   │                           │
   │                    ──pointerdown──▶ [PRESSED]
   │                                        │
   │                                 ──pointerup──▶ [HOVER]（仍在区域内）
   │                                        │
   │                                 ──pointerup outside──▶ [NORMAL]
   │                                        │
   ▼                              ──pointerout during press──▶ [NORMAL]
[DISABLED] ◀────setDisabled(true)────┘
```

### 6.2 深度分配状态

| 状态 | 触发条件 | 行为 |
|------|----------|------|
| 首次创建 | `new UIPanel(...)` | `depth = UIDepthLayer.Dialog`（3000） |
| 同层叠加 | 第二个 Dialog 打开 | `depth = 3001`（自动递增） |
| 显式覆盖 | 传入 `depth: 9999` | 使用传入值，跳过自动分配 |

---

## 7. Accessibility 设计

1. **色板对比度**：`UITheme.inkBlack` (#1a1a1a) 背景 + `UITheme.white` (#ffffff) 文字 = 对比度 14.5:1，远超 WCAG AAA 标准（7:1）
2. **禁用状态**：`alpha ≤ 0.5` + `disableInteractive()`，视觉上与可用态有明显区分
3. **焦点管理**：`UIInput` 等组件提供 focus/blur 视觉反馈（边框颜色变化）
4. **回退可见性**：纹理缺失时回退色块带边框，在复杂背景下可辨

---

## 8. 响应式 / 适配设计

1. **组件尺寸使用配置参数** — 所有 `width` / `height` 均可通过 Config 传入，不硬编码
2. **Scene 尺寸读取** — `scene.cameras.main.width/height` 用于居中对齐等计算
3. **1280×720 与 1920×1080 兼容** — 组件位置通过相对值或配置传入，无绝对坐标

---

## 9. 命名规范

| 层级 | 命名规则 | 示例 |
|------|----------|------|
| 类型定义 | PascalCase，接口后缀 `Config` | `UIButtonConfig` |
| 类型别名 | PascalCase | `UIEventCallback` |
| 工具常量 | PascalCase | `UITheme`、`UITweens` |
| 工具函数 | camelCase | `hasTexture`、`createInkBackground` |
| 枚举 | PascalCase + 语义名 | `UIDepthLayer` |
| 文件命名 | PascalCase（组件）/ camelCase（工具） | `UIButton.ts`、`utils.ts` |

---

## 10. 验收标准（供开发 Agent 参考）

- [ ] `types.ts` 包含 PRD-004 5.2 节全部公共类型 + `UIDepthLayer` / `UIInteractiveState` / `UITweenConfig`
- [ ] `utils.ts` 重写后包含：色板、缓动、纹理检查、回退渲染（带 warn）、文字样式、装饰线、圆角纹理生成、深度分配、颜色转换
- [ ] 所有纹理缺失回退统一使用 `createFallbackRect()` 并打印 `[UI] fallback: ${key}` warn
- [ ] `__tests__/setup.ts` 提供 `UIMockScene`、`simulateClick`、`simulateHover`、`mockTexture`、`expectFallbackWarn` 等 helper
- [ ] `index.ts` 按「基础设施 → Layer 1 → Layer 2」顺序导出，无额外逻辑
- [ ] `pnpm --filter phaser3 lint` 通过
- [ ] `pnpm --filter phaser3 type-check` 通过
- [ ] 无循环依赖：`setup.ts` → `utils.ts` → `types.ts` 单向
