/**
 * UIToast — 国风轻提示组件（重做）
 *
 * 仅承载消息提示职责，通过左侧 4px 类型色条区分 info/success/warning/error。
 * 进度/操作倒计时等复杂逻辑应迁移到业务层或专用组件。
 */

import { Scene } from 'phaser';
import { UITheme, UITweens, hasTexture, createInkBackground, createFallbackRect, getInkTextStyle, killTweensOf, getNextDepth } from './utils';
import { UIDepthLayer } from './types';

import type { UIPositionConfig, UISizeConfig, UINotificationType } from './types';

export interface UIToastConfig extends UIPositionConfig, Partial<UISizeConfig> {
  /** 提示消息 */
  message: string;
  /** 提示类型，决定左侧色条颜色 */
  type?: UINotificationType;
  /** 自动消失延迟（毫秒），默认 3000 */
  duration?: number;
  /** 左侧图标纹理（可选） */
  icon?: string;
  /** 背景纹理（可选） */
  texture?: string;
  /** 显式覆盖 depth */
  depth?: number;
}

export class UIToast extends Phaser.GameObjects.Container {
  private config: Required<UIToastConfig>;
  private background!: Phaser.GameObjects.Rectangle | Phaser.GameObjects.Image;
  private leftBar!: Phaser.GameObjects.Rectangle;
  private messageText!: Phaser.GameObjects.Text;
  private iconImage?: Phaser.GameObjects.Image;
  private autoHideEvent?: Phaser.Time.TimerEvent;

  constructor(scene: Scene, config: UIToastConfig) {
    const camera = scene.cameras.main;
    super(scene, config.x ?? camera.width / 2, config.y ?? camera.height - 100);

    this.config = {
      x: config.x ?? camera.width / 2,
      y: config.y ?? camera.height - 100,
      width: config.width ?? 350,
      height: config.height ?? 56,
      message: config.message,
      type: config.type ?? 'info',
      duration: config.duration ?? 3000,
      icon: config.icon ?? '',
      texture: config.texture ?? '',
      depth: config.depth ?? getNextDepth(UIDepthLayer.Toast, scene),
    };

    this.setDepth(this.config.depth);
    this.setScrollFactor(0);

    this.createBackground();
    this.createLeftBar();
    this.createIcon();
    this.createMessage();

    // 初始态：用于入场动画
    this.setVisible(true);
    this.setAlpha(0);
    this.y += 20;

    scene.add.existing(this);
  }

  /** 获取类型色 */
  private getTypeColor(): number {
    switch (this.config.type) {
      case 'success': return UITheme.stoneCyan;
      case 'warning': return UITheme.goldBrown;
      case 'error': return UITheme.cinnabarRed;
      default: return UITheme.indigoBlue;
    }
  }

  /** 创建背景 */
  private createBackground(): void {
    const { texture, width, height } = this.config;
    const strokeColor = this.getTypeColor();

    if (texture && hasTexture(this.scene, texture)) {
      this.background = this.scene.add.image(0, 0, texture);
      this.background.setDisplaySize(width, height);
    } else if (texture) {
      this.background = createFallbackRect(this.scene, width, height, texture, UITheme.inkBlack, strokeColor);
    } else {
      this.background = createInkBackground(this.scene, width, height, UITheme.inkBlack, strokeColor);
    }

    if (this.background instanceof Phaser.GameObjects.Rectangle) {
      this.background.setFillStyle(UITheme.inkBlack, 0.95);
    }

    this.add(this.background);
  }

  /** 创建左侧 4px 类型色条 */
  private createLeftBar(): void {
    const { height } = this.config;
    const color = this.getTypeColor();

    this.leftBar = this.scene.add.rectangle(
      -this.config.width / 2 + 2,
      0,
      4,
      height - 8,
      color,
      1
    );

    this.add(this.leftBar);
  }

  /** 创建左侧图标（可选） */
  private createIcon(): void {
    const { icon } = this.config;
    if (!icon) return;

    const iconX = -this.config.width / 2 + 24;
    const size = 24;

    if (hasTexture(this.scene, icon)) {
      this.iconImage = this.scene.add.image(iconX, 0, icon);
      this.iconImage.setDisplaySize(size, size);
    } else {
      const fallback = createFallbackRect(this.scene, size, size, icon, UITheme.lightInk, UITheme.goldBrown);
      fallback.setPosition(iconX, 0);
      this.add(fallback);
      return;
    }

    this.add(this.iconImage);
  }

  /** 创建消息文字 */
  private createMessage(): void {
    const { message, width, icon } = this.config;
    const paddingLeft = icon ? 48 : 16;

    this.messageText = this.scene.add.text(
      -width / 2 + paddingLeft + 8,
      0,
      message,
      getInkTextStyle('14px', colorToHex(UITheme.white))
    ).setOrigin(0, 0.5);

    this.add(this.messageText);
  }

  /** 入场动画 */
  show(): void {
    killTweensOf(this.scene, this);
    this.scene.tweens.add({
      targets: this,
      alpha: 1,
      y: this.y - 20,
      duration: 300,
      ease: 'Back.out',
    });

    // 自动消失
    if (this.config.duration > 0) {
      this.autoHideEvent = this.scene.time.delayedCall(this.config.duration, () => {
        this.dismiss();
      });
    }
  }

  /** 退场动画并销毁 */
  dismiss(): void {
    if (this.autoHideEvent) {
      this.autoHideEvent.remove();
      this.autoHideEvent = undefined;
    }

    killTweensOf(this.scene, this);
    this.scene.tweens.add({
      targets: this,
      alpha: 0,
      y: this.y - 10,
      duration: 200,
      ease: 'Power2',
      onComplete: () => {
        this.destroy();
      },
    });
  }

  /** 创建并显示 Toast */
  static show(scene: Scene, config: UIToastConfig): UIToast {
    const toast = new UIToast(scene, config);
    toast.show();
    return toast;
  }

  destroy(fromScene?: boolean): void {
    if (this.autoHideEvent) {
      this.autoHideEvent.remove();
      this.autoHideEvent = undefined;
    }
    super.destroy(fromScene);
  }
}

/** 数字色值转 HEX 字符串 */
function colorToHex(color: number): string {
  return `#${color.toString(16).padStart(6, '0')}`;
}

// ============================================================================
// ToastManager — Toast 队列管理器（单例）
// ============================================================================

export class ToastManager {
  private static instances = new WeakMap<Scene, ToastManager>();

  private scene: Scene;
  private toasts: UIToast[] = [];
  private startY: number;
  private spacing: number = 64;

  private constructor(scene: Scene, startY?: number) {
    this.scene = scene;
    this.startY = startY ?? scene.cameras.main.height - 100;
  }

  /** 获取指定 Scene 的单例管理器 */
  static getInstance(scene: Scene): ToastManager {
    if (!ToastManager.instances.has(scene)) {
      ToastManager.instances.set(scene, new ToastManager(scene));
    }
    return ToastManager.instances.get(scene)!;
  }

  /** 显示一个 Toast，自动处理队列和堆叠 */
  show(config: UIToastConfig): UIToast {
    // 队列满时（最多 3 个），最早的 Toast 提前 dismiss
    if (this.toasts.length >= 3) {
      const earliest = this.toasts[0];
      if (earliest) {
        earliest.dismiss();
      }
    }

    const y = this.startY - this.toasts.length * this.spacing;

    const toast = new UIToast(this.scene, {
      ...config,
      y,
    });

    // 拦截原始 dismiss，加入队列管理
    const originalDismiss = toast['dismiss'].bind(toast);
    toast['dismiss'] = () => {
      originalDismiss();
      this.removeToast(toast);
    };

    toast.show();
    this.toasts.push(toast);
    return toast;
  }

  /** 清空所有 Toast */
  clear(): void {
    this.toasts.forEach(toast => toast.destroy());
    this.toasts = [];
  }

  /** 获取当前 Toast 数量 */
  getCount(): number {
    return this.toasts.length;
  }

  private removeToast(toast: UIToast): void {
    const index = this.toasts.indexOf(toast);
    if (index !== -1) {
      this.toasts.splice(index, 1);
      this.repositionToasts();
    }
  }

  private repositionToasts(): void {
    this.toasts.forEach((toast, index) => {
      const targetY = this.startY - index * this.spacing;
      killTweensOf(this.scene, toast);
      this.scene.tweens.add({
        targets: toast,
        y: targetY,
        duration: 200,
        ease: 'Power2',
      });
    });
  }
}
