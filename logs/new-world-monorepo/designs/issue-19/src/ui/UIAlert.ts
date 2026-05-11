/**
 * UIAlert — 国风提示条组件（重做）
 *
 * 通过左侧 4px 类型色条标识 info/success/warning/error，
 * 禁止使用 Unicode 符号作为类型图标。
 */

import { Scene } from 'phaser';
import { UITheme, UITweens, killTweensOf, getNextDepth } from './utils';
import { UIDepthLayer } from './types';

import type { UIPositionConfig, UISizeConfig, UINotificationType } from './types';

export interface UIAlertConfig extends UIPositionConfig, Partial<UISizeConfig> {
  /** 提示消息 */
  message: string;
  /** 提示类型，决定左侧色条颜色 */
  type?: UINotificationType;
  /** 自动消失延迟（毫秒），默认 3000 */
  duration?: number;
  /** 是否自动消失 */
  isAutoHide?: boolean;
  /** 显式覆盖 depth */
  depth?: number;
}

export class UIAlert extends Phaser.GameObjects.Container {
  private config: Required<UIAlertConfig>;
  private background!: Phaser.GameObjects.Rectangle;
  private leftBar!: Phaser.GameObjects.Rectangle;
  private messageText!: Phaser.GameObjects.Text;
  private autoHideEvent?: Phaser.Time.TimerEvent;

  constructor(scene: Scene, config: UIAlertConfig) {
    super(scene, config.x ?? 0, config.y ?? 0);

    this.config = {
      x: config.x ?? 0,
      y: config.y ?? 0,
      width: config.width ?? 280,
      height: config.height ?? 44,
      message: config.message,
      type: config.type ?? 'info',
      duration: config.duration ?? 3000,
      isAutoHide: config.isAutoHide ?? true,
      depth: config.depth ?? getNextDepth(UIDepthLayer.Toast, scene),
    };

    this.setDepth(this.config.depth);

    this.createBackground();
    this.createLeftBar();
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
    const { width, height } = this.config;
    const color = this.getTypeColor();

    this.background = this.scene.add.rectangle(0, 0, width, height, UITheme.inkBlack, 0.95);
    this.background.setStrokeStyle(2, color);
    this.background.setOrigin(0.5);

    this.add(this.background);
  }

  /** 创建左侧 4px 类型色条 */
  private createLeftBar(): void {
    const { width, height } = this.config;
    const color = this.getTypeColor();

    this.leftBar = this.scene.add.rectangle(
      -width / 2 + 2,
      0,
      4,
      height - 4,
      color,
      1
    );

    this.add(this.leftBar);
  }

  /** 创建消息文字 */
  private createMessage(): void {
    const { message } = this.config;

    this.messageText = this.scene.add.text(
      0,
      0,
      message,
      {
        fontFamily: '"Microsoft YaHei", "SimHei", sans-serif',
        fontSize: '14px',
        color: '#ffffff',
      }
    ).setOrigin(0.5);

    this.add(this.messageText);
  }

  /** 入场动画 */
  show(): void {
    killTweensOf(this.scene, this);
    this.scene.tweens.add({
      targets: this,
      alpha: 1,
      y: this.y - 20,
      duration: 200,
      ease: 'Back.out',
    });

    if (this.config.isAutoHide) {
      this.autoHideEvent = this.scene.time.delayedCall(this.config.duration, () => {
        this.hide();
      });
    }
  }

  /** 退场动画并销毁 */
  hide(): void {
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

  /** 创建并显示 Alert */
  static show(scene: Scene, config: UIAlertConfig): UIAlert {
    const alert = new UIAlert(scene, config);
    alert.show();
    return alert;
  }

  /** 获取指定类型的主题色 */
  static getTypeColor(type: UINotificationType): number {
    switch (type) {
      case 'success': return UITheme.stoneCyan;
      case 'warning': return UITheme.goldBrown;
      case 'error': return UITheme.cinnabarRed;
      default: return UITheme.indigoBlue;
    }
  }

  destroy(fromScene?: boolean): void {
    if (this.autoHideEvent) {
      this.autoHideEvent.remove();
      this.autoHideEvent = undefined;
    }
    super.destroy(fromScene);
  }
}
