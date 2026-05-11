/**
 * UIDialog — 国风模态对话框组件（重做）
 *
 * 继承 UIPanel，复用其遮罩、背景、标题、关闭按钮及 show/hide 动画。
 * 提供确认/取消双按钮布局，支持快速单按钮 alert 工厂方法。
 */

import { Scene } from 'phaser';
import { UIPanel } from './UIPanel';
import { UITheme, UITweens, killTweensOf, getNextDepth } from './utils';
import { UIDepthLayer } from './types';
import { UIButton } from './UIButton';

import type { UIPositionConfig, UISizeConfig } from './types';

export interface UIDialogConfig extends UIPositionConfig, Partial<UISizeConfig> {
  /** 对话框标题（必填） */
  title: string;
  /** 消息正文（必填） */
  message: string;
  /** 确认按钮文字 */
  confirmText?: string;
  /** 取消按钮文字 */
  cancelText?: string;
  /** 是否显示取消按钮 */
  showCancel?: boolean;
  /** 确认按钮纹理 */
  confirmTexture?: string;
  /** 取消按钮纹理 */
  cancelTexture?: string;
  /** 确认回调 */
  onConfirm?: () => void;
  /** 取消回调 */
  onCancel?: () => void;
  /** 点击遮罩是否关闭（默认 false） */
  closeOnOverlay?: boolean;
  /** 显式覆盖 depth */
  depth?: number;
}

export class UIDialog extends UIPanel {
  private dialogConfig: Required<UIDialogConfig>;
  private messageText!: Phaser.GameObjects.Text;
  private confirmButton!: UIButton;
  private cancelButton?: UIButton;

  constructor(scene: Scene, config: UIDialogConfig) {
    const camera = scene.cameras.main;

    // 构建 UIPanel 基类配置
    const panelConfig = {
      x: config.x ?? camera.width / 2,
      y: config.y ?? camera.height / 2,
      width: config.width ?? 360,
      height: config.height ?? 220,
      title: config.title,
      showClose: false,
      showOverlay: true,
      overlayAlpha: 0.6,
      isModal: true,
      depth: config.depth ?? getNextDepth(UIDepthLayer.Dialog, scene),
      onOverlayClick: () => {
        if (config.closeOnOverlay) {
          this.cancel();
        }
      },
    };

    super(scene, panelConfig);

    this.dialogConfig = {
      x: config.x ?? camera.width / 2,
      y: config.y ?? camera.height / 2,
      width: config.width ?? 360,
      height: config.height ?? 220,
      title: config.title,
      message: config.message,
      confirmText: config.confirmText ?? '确认',
      cancelText: config.cancelText ?? '取消',
      showCancel: config.showCancel ?? true,
      confirmTexture: config.confirmTexture ?? '',
      cancelTexture: config.cancelTexture ?? '',
      onConfirm: config.onConfirm ?? (() => {}),
      onCancel: config.onCancel ?? (() => {}),
      closeOnOverlay: config.closeOnOverlay ?? false,
      depth: config.depth ?? getNextDepth(UIDepthLayer.Dialog, scene),
    };

    this.createMessage();
    this.createButtons();
  }

  /** 创建消息正文 */
  private createMessage(): void {
    const { message, width } = this.dialogConfig;

    this.messageText = this.scene.add.text(
      0,
      -10,
      message,
      {
        fontFamily: '"Microsoft YaHei", "SimHei", sans-serif',
        fontSize: '14px',
        color: '#cccccc',
        align: 'center',
        wordWrap: { width: width - 40 },
      }
    ).setOrigin(0.5);

    this.addContent(this.messageText);
  }

  /** 创建确认/取消按钮 */
  private createButtons(): void {
    const { width, height, showCancel, confirmText, cancelText } = this.dialogConfig;
    const btnY = height / 2 - 45;

    this.confirmButton = new UIButton(this.scene, {
      x: showCancel ? 80 : 0,
      y: btnY,
      width: 120,
      height: 40,
      text: confirmText,
      bgColor: UITheme.cinnabarRed,
      hoverColor: UITheme.cinnabarRedLight,
      onClick: () => this.confirm(),
    });

    this.addContent(this.confirmButton);

    if (showCancel) {
      this.cancelButton = new UIButton(this.scene, {
        x: -80,
        y: btnY,
        width: 120,
        height: 40,
        text: cancelText,
        bgColor: UITheme.lightInk,
        hoverColor: UITheme.mediumInk,
        onClick: () => this.cancel(),
      });

      this.addContent(this.cancelButton);
    }
  }

  /** 确认操作：先 hide 动画，完成后执行 onConfirm */
  private confirm(): void {
    killTweensOf(this.scene, this);
    this.hide(() => {
      this.dialogConfig.onConfirm?.();
    });
  }

  /** 取消操作：先 hide 动画，完成后执行 onCancel */
  private cancel(): void {
    killTweensOf(this.scene, this);
    this.hide(() => {
      this.dialogConfig.onCancel?.();
    });
  }

  /** 创建并显示对话框（双按钮） */
  static show(scene: Scene, config: UIDialogConfig): UIDialog {
    const dialog = new UIDialog(scene, config);
    dialog.show();
    return dialog;
  }

  /** 创建并显示单按钮提示框（alert 模式） */
  static alert(scene: Scene, title: string, message: string): UIDialog {
    const dialog = new UIDialog(scene, {
      title,
      message,
      showCancel: false,
      confirmText: '确定',
      onConfirm: () => {},
    });
    dialog.show();
    return dialog;
  }
}
