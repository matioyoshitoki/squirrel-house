/**
 * UIPanel — 国风通用面板组件（重做）
 *
 * 带标题栏、关闭按钮、模态遮罩，支持标准入场/退场动画。
 * 所有依赖纹理的渲染在缺失时自动回退到 Rectangle，并打印统一 warn。
 */

import { Scene } from 'phaser';
import {
  UITheme,
  UITweens,
  hasTexture,
  createInkBackground,
  createFallbackRect,
  getInkTextStyle,
  createInkDecorationLine,
  killTweensOf,
  getNextDepth,
} from './utils';
import { UIDepthLayer } from './types';

import type { UIPositionConfig, UISizeConfig } from './types';

export interface UIPanelConfig extends UIPositionConfig, Partial<UISizeConfig> {
  /** 面板标题，空字符串时不显示标题栏 */
  title?: string;
  /** 背景纹理键名 */
  texture?: string;
  /** 背景填充色（回退时使用） */
  bgColor?: number;
  /** 边框颜色（回退时使用） */
  strokeColor?: number;
  /** 是否显示关闭按钮 */
  showClose?: boolean;
  /** 是否显示模态遮罩 */
  showOverlay?: boolean;
  /** 遮罩透明度 */
  overlayAlpha?: number;
  /** 是否为模态面板（禁止点击遮罩关闭时设为 true） */
  isModal?: boolean;
  /** 关闭按钮点击回调 */
  onClose?: () => void;
  /** 遮罩点击回调 */
  onOverlayClick?: () => void;
  /** 显式覆盖 depth，默认自动分配 */
  depth?: number;
}

export class UIPanel extends Phaser.GameObjects.Container {
  protected config: Required<UIPanelConfig>;
  protected background!: Phaser.GameObjects.Rectangle | Phaser.GameObjects.Image;
  protected overlay?: Phaser.GameObjects.Rectangle;
  protected titleText?: Phaser.GameObjects.Text;
  protected titleLine?: Phaser.GameObjects.Rectangle;
  protected closeButton?: Phaser.GameObjects.Container;

  constructor(scene: Scene, config: UIPanelConfig) {
    super(scene, config.x ?? 0, config.y ?? 0);

    this.config = {
      x: config.x ?? 0,
      y: config.y ?? 0,
      width: config.width ?? 400,
      height: config.height ?? 300,
      title: config.title ?? '',
      texture: config.texture ?? '',
      bgColor: config.bgColor ?? UITheme.inkBlack,
      strokeColor: config.strokeColor ?? UITheme.goldBrown,
      showClose: config.showClose ?? false,
      showOverlay: config.showOverlay ?? false,
      overlayAlpha: config.overlayAlpha ?? 0.6,
      isModal: config.isModal ?? false,
      onClose: config.onClose ?? (() => {}),
      onOverlayClick: config.onOverlayClick ?? (() => {}),
      depth: config.depth ?? getNextDepth(config.isModal ? UIDepthLayer.Dialog : UIDepthLayer.Panel, scene),
    };

    this.setDepth(this.config.depth);

    if (this.config.showOverlay) {
      this.createOverlay();
    }

    this.createBackground();
    this.createTitle();
    this.createCloseButton();

    // 初始隐藏，等待 show() 调用
    this.setVisible(false);
    this.setAlpha(0);
    this.setScale(0.8);

    scene.add.existing(this);
  }

  /** 创建模态遮罩层 */
  protected createOverlay(): void {
    const camera = this.scene.cameras.main;
    this.overlay = this.scene.add.rectangle(
      0,
      0,
      camera.width * 2,
      camera.height * 2,
      UITheme.inkBlack,
      this.config.overlayAlpha
    );
    this.overlay.setInteractive();
    this.overlay.setDepth(UIDepthLayer.Overlay);

    this.overlay.on('pointerdown', () => {
      this.config.onOverlayClick?.();
    });

    this.add(this.overlay);
  }

  /** 创建面板背景 */
  protected createBackground(): void {
    const { texture, width, height, bgColor, strokeColor } = this.config;

    if (texture && hasTexture(this.scene, texture)) {
      this.background = this.scene.add.image(0, 0, texture);
      this.background.setDisplaySize(width, height);
    } else if (texture) {
      this.background = createFallbackRect(this.scene, width, height, texture, bgColor, strokeColor);
    } else {
      this.background = createInkBackground(this.scene, width, height, bgColor, strokeColor);
    }

    // 背景透明度与规范对齐：0.98
    if (this.background instanceof Phaser.GameObjects.Rectangle) {
      this.background.setFillStyle(bgColor, 0.98);
    }

    this.add(this.background);
  }

  /** 创建标题栏 */
  protected createTitle(): void {
    if (!this.config.title) return;

    const { width } = this.config;
    const titleY = -this.config.height / 2 + 30;

    this.titleText = this.scene.add.text(
      0,
      titleY,
      this.config.title,
      getInkTextStyle('18px', colorToHex(UITheme.white), { fontStyle: 'bold' })
    ).setOrigin(0.5);

    this.titleLine = createInkDecorationLine(this.scene, width - 40, UITheme.goldBrown);
    this.titleLine.setPosition(0, titleY + 18);

    this.add([this.titleText, this.titleLine]);
  }

  /** 创建关闭按钮（使用 UIIcon 规范，禁止 Unicode ×） */
  protected createCloseButton(): void {
    if (!this.config.showClose) return;

    const closeX = this.config.width / 2 - 20;
    const closeY = -this.config.height / 2 + 20;
    const size = 24;
    const textureKey = 'cancel-btn-0';

    this.closeButton = this.scene.add.container(closeX, closeY);

    let icon: Phaser.GameObjects.Rectangle | Phaser.GameObjects.Image;

    if (hasTexture(this.scene, textureKey)) {
      icon = this.scene.add.image(0, 0, textureKey);
      icon.setDisplaySize(size, size);
    } else {
      icon = createFallbackRect(this.scene, size, size, textureKey, UITheme.inkBlack, UITheme.goldBrown);
    }

    icon.setInteractive({ useHandCursor: true });

    icon.on('pointerover', () => {
      if (icon instanceof Phaser.GameObjects.Rectangle) {
        icon.setStrokeStyle(2, UITheme.cinnabarRed);
      }
      this.scene.tweens.add({
        targets: this.closeButton,
        scale: 1.1,
        duration: UITweens.hover.duration,
        ease: UITweens.hover.ease,
      });
    });

    icon.on('pointerout', () => {
      if (icon instanceof Phaser.GameObjects.Rectangle) {
        icon.setStrokeStyle(2, UITheme.goldBrown);
      }
      this.scene.tweens.add({
        targets: this.closeButton,
        scale: 1,
        duration: UITweens.hover.duration,
        ease: UITweens.hover.ease,
      });
    });

    icon.on('pointerdown', () => {
      killTweensOf(this.scene, this.closeButton);
      this.scene.tweens.add({
        targets: this.closeButton,
        scale: 0.95,
        duration: UITweens.press.duration,
        ease: UITweens.press.ease,
        yoyo: true,
        onComplete: () => {
          this.config.onClose?.();
        },
      });
    });

    this.closeButton.add(icon);
    this.add(this.closeButton);
  }

  /** 显示面板（入场动画） */
  show(): void {
    this.setVisible(true);
    if (this.overlay) {
      this.overlay.setVisible(true);
    }

    killTweensOf(this.scene, this);
    this.scene.tweens.add({
      targets: this,
      alpha: 1,
      scale: 1,
      duration: UITweens.popup.duration,
      ease: UITweens.popup.ease,
    });
  }

  /** 隐藏面板（退场动画） */
  hide(onComplete?: () => void): void {
    killTweensOf(this.scene, this);
    this.scene.tweens.add({
      targets: this,
      alpha: 0,
      scale: 0.8,
      duration: UITweens.fade.duration,
      ease: UITweens.fade.ease,
      onComplete: () => {
        this.setVisible(false);
        if (this.overlay) {
          this.overlay.setVisible(false);
        }
        onComplete?.();
      },
    });
  }

  /** 动态更新标题 */
  setTitle(title: string): void {
    this.config.title = title;

    if (this.titleText) {
      this.titleText.setText(title);
      this.titleText.setVisible(title.length > 0);
    }

    if (this.titleLine) {
      this.titleLine.setVisible(title.length > 0);
    }
  }

  /** 获取内容区可用矩形（供业务组件安全放置子元素） */
  getContentBounds(): Phaser.Geom.Rectangle {
    const { width, height } = this.config;
    const topPadding = this.config.title ? 60 : 20;
    return new Phaser.Geom.Rectangle(
      -width / 2 + 20,
      -height / 2 + topPadding,
      width - 40,
      height - topPadding - 20
    );
  }

  /** 向面板内容区添加子元素 */
  addContent(child: Phaser.GameObjects.GameObject): void {
    this.add(child);
  }

  /** 控制关闭按钮显隐 */
  setCloseVisible(isVisible: boolean): void {
    this.config.showClose = isVisible;
    if (this.closeButton) {
      this.closeButton.setVisible(isVisible);
    }
  }
}

/** 数字色值转 HEX 字符串（局部辅助，避免循环导入） */
function colorToHex(color: number): string {
  return `#${color.toString(16).padStart(6, '0')}`;
}
