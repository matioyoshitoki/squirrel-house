/**
 * 国风水墨 UI 基础组件库统一入口
 *
 * 按以下顺序组织导出：
 *   1. 基础设施（types + utils）
 *   2. Layer 1: 原子组件（16 个）
 *   3. Layer 2: 通用交互组件（6 个）
 *
 * 所有组件在纹理资源缺失时均会自动回退到 Rectangle 渲染。
 * 禁止在此文件中引入任何副作用（CSS、polyfill、自动执行代码）。
 */

// ============================================================================
// 1. 基础设施
// ============================================================================

export {
  /** 国风主题色板（9 色） */
  UITheme,
  /** 检测纹理资源是否可用 */
  hasTexture,
  /** 创建圆角矩形纹理 */
  ensureRoundRectTexture,
  /** 国风文字样式生成器 */
  getInkTextStyle,
  /** 创建带国风边框的背景 */
  createInkBackground,
  /** 创建统一回退矩形（带 warn 输出） */
  createFallbackRect,
  /** 创建水墨装饰线 */
  createInkDecorationLine,
  /** 通用缓动配置 */
  UITweens,
  /** 获取下一个可用 depth */
  getNextDepth,
  /** 重置 Scene 的 depth 计数器 */
  resetDepthCounters,
  /** 数字色值转 HEX 字符串 */
  colorToHex,
  /** 数值裁剪 */
  clamp,
  /** 品质等级到回退颜色映射表 */
  qualityColorMap,
  /** 安全地停止目标对象上所有 tween */
  killTweensOf,
} from './utils';

export type {
  /** 位置配置 */
  UIPositionConfig,
  /** 尺寸配置 */
  UISizeConfig,
  /** 交互回调统一签名 */
  UIEventCallback,
  /** 品质等级 1~9 */
  UIQualityLevel,
  /** Toast / Alert / Badge 通知类型 */
  UINotificationType,
  /** 深度层级枚举 */
  UIDepthLayer,
  /** 交互组件状态机状态 */
  UIInteractiveState,
  /** 按钮专用状态 */
  UIButtonState,
  /** 统一缓动配置 */
  UITweenConfig,
  /** 资源回退配置 */
  UIFallbackConfig,
  /** 组件公共基础 Props */
  UIBaseConfig,
  /** 通用选项类型 */
  UIOptionItem,
  /** 品质颜色映射类型 */
  UIQualityColorMap,
} from './types';

// ============================================================================
// 2. Layer 1: 原子组件（16 个）
// ============================================================================

// export { UIButton, type UIButtonConfig } from './UIButton';
// export { UIIconButton, type UIIconButtonConfig } from './UIIconButton';
// export { UIPanel, type UIPanelConfig } from './UIPanel';
// export { UIDialog, type UIDialogConfig } from './UIDialog';
// export { UIList, type UIListConfig } from './UIList';
// export { UIListItem, type UIListItemConfig } from './UIListItem';
// export { UIText, type UITextConfig } from './UIText';
// export { UIAvatar, type UIAvatarConfig } from './UIAvatar';
// export { UIProgressBar, type UIProgressBarConfig } from './UIProgressBar';
// export { UIToast, ToastManager, type UIToastConfig, type ToastType } from './UIToast';
// export { UIResourceDisplay, type UIResourceDisplayConfig } from './UIResourceDisplay';
// export { UISegmentedControl, type UISegmentedControlConfig, type UISegmentedOption } from './UISegmentedControl';
// export { UIIcon, type UIIconConfig } from './UIIcon';
// export { UIAlert, type UIAlertConfig, type UIAlertType } from './UIAlert';
// export { UIBadge, type UIBadgeConfig, type UIBadgeType } from './UIBadge';
// export { UITooltip, type UITooltipConfig } from './UITooltip';

// ============================================================================
// 3. Layer 2: 通用交互组件（6 个）
// ============================================================================

// export { UISlider, type UISliderConfig } from './UISlider';
// export { UIToggle, type UIToggleConfig } from './UIToggle';
// export { UIStepper, type UIStepperConfig } from './UIStepper';
// export { UIDropdown, type UIDropdownConfig, type UIDropdownOption } from './UIDropdown';
// export { UIPagination, type UIPaginationConfig } from './UIPagination';
// export { UILoading, type UILoadingConfig } from './UILoading';

// ============================================================================
// 4. 向后兼容说明
// ============================================================================

/**
 * 迁移指南（供 Developer Agent 参考）：
 *
 * 1. 将本文件复制到 `apps/phaser3/src/ui/index.ts`，取消对应组件的注释。
 * 2. 将 `types.ts` 复制到 `apps/phaser3/src/ui/types.ts`。
 * 3. 将 `utils.ts` 复制到 `apps/phaser3/src/ui/utils.ts`（覆盖旧文件）。
 * 4. 将 `__tests__/setup.ts` 复制到 `apps/phaser3/src/ui/__tests__/setup.ts`。
 * 5. 各组件文件需逐步更新以导入 `types.ts` 中的公共基类：
 *    ```ts
 *    import { UIPositionConfig, UISizeConfig } from './types';
 *    export interface UIButtonConfig extends UIPositionConfig, Partial<UISizeConfig> { ... }
 *    ```
 * 6. 所有组件应使用 `createFallbackRect()` 替代直接 `add.rectangle()`，
 *    确保纹理缺失时统一输出 `[UI] fallback: ${key}` warn。
 * 7. 使用 `getNextDepth(UIDepthLayer.Dialog, scene)` 替代硬编码 depth 值。
 */
