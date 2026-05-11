/**
 * UI 组件公共类型定义
 *
 * 本文件仅包含类型定义（零运行时开销），供所有 Layer 1 / Layer 2 组件共享。
 * 禁止混入任何运行时逻辑或副作用。
 */

/**
 * UI 组件公共位置配置
 */
export interface UIPositionConfig {
  x?: number;
  y?: number;
}

/**
 * UI 组件公共尺寸配置
 */
export interface UISizeConfig {
  width?: number;
  height?: number;
}

/**
 * 交互回调统一签名
 * @example
 *   onClick?: UIEventCallback<void>;
 *   onChange?: UIEventCallback<number>;
 */
export type UIEventCallback<T = void> = (value: T) => void;

/**
 * 品质等级 1~9，对应 quality-icon 和 quality-bg 纹理体系
 */
export type UIQualityLevel = 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9;

/**
 * Toast / Alert / Badge 通知类型
 */
export type UINotificationType = 'info' | 'success' | 'warning' | 'error';

/**
 * UI 深度层级枚举
 *
 * 与 Foundation 规范对齐：
 * - Scene: 0 ~ 999
 * - UI: 1000
 * - Panel: 2000
 * - Dialog: 3000
 * - Overlay: 4000
 * - Toast / Alert / Tooltip: 5000+
 *
 * 同一层级多个组件同时显示时，后创建的组件应自动递增 depth（+1）。
 */
export enum UIDepthLayer {
  Scene = 0,
  UI = 1000,
  Panel = 2000,
  Dialog = 3000,
  Overlay = 4000,
  Toast = 5000,
}

/**
 * 交互组件状态机状态
 *
 * @see docs/design-docs/ui-style-guide/03-component.md 组件状态机
 */
export type UIInteractiveState = 'normal' | 'hover' | 'pressed' | 'disabled';

/**
 * 按钮专用状态（与 UIInteractiveState 语义一致，用于需要更严格约束的场景）
 */
export type UIButtonState = UIInteractiveState;

/**
 * 统一缓动配置类型
 */
export interface UITweenConfig {
  /** 动画时长（毫秒） */
  duration: number;
  /** 缓动函数名称 */
  ease: string;
  /** 延迟时间（毫秒），可选 */
  delay?: number;
  /** 是否往返播放（yoyo），可选 */
  yoyo?: boolean;
  /** 重复次数，-1 为无限循环，可选 */
  repeat?: number;
}

/**
 * 资源回退配置
 */
export interface UIFallbackConfig {
  /** 回退填充色 */
  fillColor: number;
  /** 回退边框色 */
  strokeColor: number;
  /** 回退透明度 */
  alpha?: number;
  /** 关联的纹理键（用于打印 warn） */
  textureKey?: string;
}

/**
 * 组件公共基础 Props（所有 UI 组件 Config 的最小交集）
 *
 * @example
 *   interface UIButtonConfig extends UIPositionConfig, Partial<UISizeConfig> {
 *     // 组件特有配置...
 *   }
 */
export interface UIBaseConfig extends UIPositionConfig, Partial<UISizeConfig> {
  /** 自定义深度值，覆盖默认层级 */
  depth?: number;
}

/**
 * 分段控制 / 下拉选项的通用选项类型
 */
export interface UIOptionItem {
  key: string;
  label: string;
  icon?: string;
  isDisabled?: boolean;
}

/**
 * 品质到颜色的映射回退配置
 * 当 quality-icon 纹理缺失时使用
 */
export type UIQualityColorMap = Record<UIQualityLevel, number>;
