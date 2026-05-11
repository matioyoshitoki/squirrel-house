/**
 * UI 基础工具函数
 *
 * 提供国风主题色、资源检查、回退渲染、深度分配等通用能力。
 * 所有依赖纹理的组件必须通过本文件的工具执行统一回退契约。
 */

import { Scene } from 'phaser';
import { UIDepthLayer } from './types';

// ============================================================================
// 1. 色板常量
// ============================================================================

/** 国风水墨主题色板 */
export const UITheme = {
  /** 墨黑 - 主背景、弹窗底色、面板底色 */
  inkBlack: 0x1a1a1a,
  /** 宣纸黄 - 次级背景、输入框底色 */
  paperYellow: 0xf5f0e6,
  /** 朱砂红 - 危险操作、删除按钮、警告提示 */
  cinnabarRed: 0xc45c48,
  /** 石青 - 主按钮（确认、升级）、信息提示 */
  stoneCyan: 0x4a9eff,
  /** 黛蓝 - 选中高亮、Hover 态、悬浮面板 */
  indigoBlue: 0x3d5a80,
  /** 金棕 - 边框、装饰线、标题强调 */
  goldBrown: 0xc9a227,
  /** 淡墨 - 禁用状态、占位文字、次要信息 */
  lightInk: 0x666666,
  /** 纯白 - 深色背景上的正文文字 */
  white: 0xffffff,
  /** 墨色文字 - 用于浅色背景 */
  inkText: 0x2c2c2c,
} as const;

// ============================================================================
// 2. 缓动配置
// ============================================================================

/** 通用缓动配置 —— 所有 UI 动效必须统一使用 */
export const UITweens = {
  /** 弹窗/面板入场：scale 0.8→1.0, alpha 0→1 */
  popup: {
    duration: 200,
    ease: 'Back.out',
  },
  /** 弹窗/面板退场：scale 1.0→0.8, alpha 1→0 */
  fade: {
    duration: 150,
    ease: 'Power2',
  },
  /** Hover 放大：scale 1.0→1.05 */
  hover: {
    duration: 100,
    ease: 'Sine.easeInOut',
  },
  /** Pressed 缩小：scale 1.0→0.95，yoyo 恢复 */
  press: {
    duration: 50,
    ease: 'Power2',
    yoyo: true,
  },
  /** 脉冲呼吸：scale 1.0→1.1，无限循环 */
  pulse: {
    duration: 500,
    ease: 'Sine.easeInOut',
    yoyo: true,
    repeat: -1,
  },
  /** 数值变化 tween：进度条/数值平滑过渡 */
  tweenValue: {
    duration: 500,
    ease: 'Power2',
  },
  /** Toast 入场：y +20→0, alpha 0→1 */
  toastIn: {
    duration: 300,
    ease: 'Back.out',
  },
  /** Toast 退场：y 0→-10, alpha 1→0 */
  toastOut: {
    duration: 200,
    ease: 'Power2',
  },
} as const;

// ============================================================================
// 3. 纹理与资源检查
// ============================================================================

/**
 * 检测纹理资源是否可用
 */
export function hasTexture(scene: Scene, key: string): boolean {
  return scene.textures.exists(key);
}

// ============================================================================
// 4. 回退渲染工具
// ============================================================================

/**
 * 创建带国风边框的背景（Rectangle 回退）
 *
 * @param scene Phaser 场景
 * @param width 宽度
 * @param height 高度
 * @param fillColor 填充色，默认 UITheme.inkBlack
 * @param strokeColor 边框色，默认 UITheme.goldBrown
 * @returns Rectangle 对象
 */
export function createInkBackground(
  scene: Scene,
  width: number,
  height: number,
  fillColor: number = UITheme.inkBlack,
  strokeColor: number = UITheme.goldBrown
): Phaser.GameObjects.Rectangle {
  const bg = scene.add.rectangle(0, 0, width, height, fillColor, 0.95);
  bg.setStrokeStyle(2, strokeColor);
  return bg;
}

/**
 * 创建统一回退矩形，纹理缺失时调用。
 *
 * 遵循 PRD-004 资源回退统一标准：
 * 1. 检查纹理可用性（调用方已检查）
 * 2. 使用 Rectangle 回退渲染
 * 3. 打印 `[UI] fallback: ${textureKey}` warn
 * 4. 回退元素尺寸与预期纹理一致，边框颜色协调，alpha ≥ 0.8
 * 5. 回退状态下交互功能保持可用
 *
 * @param scene Phaser 场景
 * @param width 宽度
 * @param height 高度
 * @param fillColor 填充色
 * @param strokeColor 边框色
 * @param textureKey 关联的纹理键（用于打印 warn）
 * @returns Rectangle 对象
 */
export function createFallbackRect(
  scene: Scene,
  width: number,
  height: number,
  fillColor: number = UITheme.inkBlack,
  strokeColor: number = UITheme.goldBrown,
  textureKey?: string
): Phaser.GameObjects.Rectangle {
  const rect = scene.add.rectangle(0, 0, width, height, fillColor, 0.9);
  rect.setStrokeStyle(2, strokeColor);

  if (textureKey) {
    // eslint-disable-next-line no-console
    console.warn(`[UI] fallback: ${textureKey}`);
  }

  return rect;
}

// ============================================================================
// 5. 文字样式工具
// ============================================================================

/**
 * 国风文字样式
 *
 * 自动注入字体栈 `"Microsoft YaHei", "SimHei", sans-serif`。
 * 所有 UI 文字必须通过本函数生成样式，禁止手动写字体栈。
 *
 * @param fontSize 字号，默认 '16px'
 * @param color 文字颜色，默认 '#ffffff'
 * @param options 额外 TextStyle 选项
 * @returns Phaser 文字样式对象
 */
export function getInkTextStyle(
  fontSize: string | number = '16px',
  color: string = '#ffffff',
  options: Partial<Phaser.Types.GameObjects.Text.TextStyle> = {}
): Phaser.Types.GameObjects.Text.TextStyle {
  const size = typeof fontSize === 'number' ? `${fontSize}px` : fontSize;
  return {
    fontFamily: '"Microsoft YaHei", "SimHei", sans-serif',
    fontSize: size,
    color,
    ...options,
  };
}

// ============================================================================
// 6. 装饰工具
// ============================================================================

/**
 * 创建水墨装饰线
 *
 * @param scene Phaser 场景
 * @param width 线宽
 * @param color 颜色，默认 UITheme.goldBrown
 * @returns Rectangle 对象（高 2px）
 */
export function createInkDecorationLine(
  scene: Scene,
  width: number,
  color: number = UITheme.goldBrown
): Phaser.GameObjects.Rectangle {
  return scene.add.rectangle(0, 0, width, 2, color);
}

// ============================================================================
// 7. 纹理生成工具
// ============================================================================

/**
 * 创建圆角矩形纹理（如果不存在）。
 * 用于需要圆角视觉的组件回退渲染。
 *
 * @param scene Phaser 场景
 * @param key 纹理键
 * @param radius 圆角半径
 */
export function ensureRoundRectTexture(scene: Scene, key: string, radius: number): void {
  if (scene.textures.exists(key)) {
    return;
  }

  const size = radius * 2;
  const graphics = scene.add.graphics();
  graphics.fillStyle(0xffffff, 1);
  graphics.fillRoundedRect(0, 0, size, size, radius);
  graphics.generateTexture(key, size, size);
  graphics.destroy();
}

// ============================================================================
// 8. 深度分配工具
// ============================================================================

/**
 * 同层级 depth 计数器（按 Scene 隔离）
 * 弱引用：Scene 销毁后自动回收
 */
const depthCounters = new WeakMap<Scene, Map<UIDepthLayer, number>>();

/**
 * 获取下一个可用的 depth 值。
 *
 * 同一 UIDepthLayer 的多个组件同时显示时，后创建的组件自动递增 depth（+1），
 * 确保最上层可交互。
 *
 * @param layer 深度层级
 * @param scene Phaser 场景（用于隔离计数器）
 * @returns 分配的深度值
 *
 * @example
 *   const depth = getNextDepth(UIDepthLayer.Dialog, scene); // 首次返回 3000，再次返回 3001
 */
export function getNextDepth(layer: UIDepthLayer, scene?: Scene): number {
  if (!scene) {
    return layer;
  }

  let sceneMap = depthCounters.get(scene);
  if (!sceneMap) {
    sceneMap = new Map();
    depthCounters.set(scene, sceneMap);
  }

  const current = sceneMap.get(layer) ?? layer;
  sceneMap.set(layer, current + 1);
  return current;
}

/**
 * 重置指定 Scene 的深度计数器。
 * 通常在 Scene shutdown 时调用，避免内存泄漏。
 *
 * @param scene Phaser 场景
 */
export function resetDepthCounters(scene: Scene): void {
  depthCounters.delete(scene);
}

// ============================================================================
// 9. 数值与颜色工具
// ============================================================================

/**
 * 数字色值转 HEX 字符串
 *
 * @param color 数字色值，如 0x1a1a1a
 * @returns HEX 字符串，如 '#1a1a1a'
 *
 * @example
 *   colorToHex(UITheme.inkBlack); // '#1a1a1a'
 */
export function colorToHex(color: number): string {
  return '#' + color.toString(16).padStart(6, '0');
}

/**
 * 数值裁剪
 * 与 Phaser.Math.Clamp 行为一致，作为 utils 层独立工具提供，
 * 便于非 Phaser 环境（如测试）中使用。
 *
 * @param value 输入值
 * @param min 最小值
 * @param max 最大值
 * @returns 裁剪后的值
 */
export function clamp(value: number, min: number, max: number): number {
  return Math.max(min, Math.min(max, value));
}

/**
 * 品质等级到回退颜色的映射表
 * 当 quality-icon 纹理缺失时，使用对应颜色作为品质标识。
 */
export const qualityColorMap: Record<number, number> = {
  1: 0xcccccc, // 白
  2: 0x9ccc65, // 浅绿
  3: 0x4caf50, // 绿
  4: 0x4dd0e1, // 青
  5: 0x2196f3, // 蓝
  6: 0x7c4dff, // 紫
  7: 0xff9800, // 橙
  8: 0xff5722, // 深橙
  9: 0xf44336, // 红
};

// ============================================================================
// 10. 动画辅助
// ============================================================================

/**
 * 安全地停止目标对象上所有 tween，然后执行回调。
 * 用于防止连续触发时动画叠加。
 *
 * @param scene Phaser 场景
 * @param target tween 目标对象
 * @param callback 清理后的回调
 */
export function killTweensOf(
  scene: Scene,
  target: Phaser.Tweens.TweenTarget,
  callback?: () => void
): void {
  scene.tweens.killTweensOf(target);
  callback?.();
}
