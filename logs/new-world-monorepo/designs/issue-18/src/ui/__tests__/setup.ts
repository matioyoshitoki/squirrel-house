/**
 * UI 组件测试基础设施
 *
 * 提供 MockScene 扩展、事件模拟辅助、纹理模拟、回退断言等能力。
 * 所有 UI 组件单元测试应通过 `import { UIMockScene, simulateClick } from './setup'` 使用。
 *
 * @migration 当本文件复制到 `apps/phaser3/src/ui/__tests__/setup.ts` 时，
 *            请将路径 `../../../../../../apps/phaser3/src/test-setup` 替换为 `../../test-setup`。
 */

import { vi, type SpyInstance } from 'vitest';

// Re-export existing test utilities for convenience
// 迁移时改为: export { MockScene, emitEvent } from '../../test-setup';
export { MockScene, emitEvent } from '../../../../../../apps/phaser3/src/test-setup';

// ============================================================================
// 1. 扩展 MockScene（纹理注册能力）
// ============================================================================

/**
 * UI 专用 MockScene，在基础 MockScene 上增加纹理注册能力。
 *
 * 使用方式：
 * ```ts
 * let scene: UIMockScene;
 * beforeEach(() => {
 *   scene = new UIMockScene();
 * });
 * ```
 */
export class UIMockScene {
  private registeredTextures = new Set<string>();
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  private _mockScene: any;

  constructor() {
    // Dynamically create a MockScene-like object with texture registration
    // 迁移时改为: const { MockScene } = require('../../test-setup');
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const { MockScene } = require('../../../../../../apps/phaser3/src/test-setup');
    this._mockScene = new MockScene();

    // Override textures.exists to check registered textures
    const originalExists = this._mockScene.textures.exists;
    this._mockScene.textures.exists = (key: string) => {
      return this.registeredTextures.has(key) || originalExists(key);
    };
  }

  /** 访问底层 MockScene 实例（用于类型转换） */
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  get native(): any {
    return this._mockScene;
  }

  /** 注册 mock 纹理 */
  registerTexture(key: string): void {
    this.registeredTextures.add(key);
  }

  /** 清除所有 mock 纹理 */
  clearTextures(): void {
    this.registeredTextures.clear();
  }

  // Proxy common properties for convenience
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  get cameras(): any { return this._mockScene.cameras; }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  get tweens(): any { return this._mockScene.tweens; }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  get time(): any { return this._mockScene.time; }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  get input(): any { return this._mockScene.input; }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  get textures(): any { return this._mockScene.textures; }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  get add(): any { return this._mockScene.add; }
}

// ============================================================================
// 2. 事件模拟辅助
// ============================================================================

/**
 * 模拟点击事件（pointerdown + pointerup）
 *
 * @param target 目标对象（需支持 on/emit 事件接口）
 */
export function simulateClick(target: { emit: (event: string, ...args: unknown[]) => unknown }): void {
  target.emit('pointerdown');
  target.emit('pointerup');
}

/**
 * 模拟 Hover 事件（pointerover）
 *
 * @param target 目标对象
 */
export function simulateHover(target: { emit: (event: string, ...args: unknown[]) => unknown }): void {
  target.emit('pointerover');
}

/**
 * 模拟 PointerOut 事件
 *
 * @param target 目标对象
 */
export function simulatePointerOut(target: { emit: (event: string, ...args: unknown[]) => unknown }): void {
  target.emit('pointerout');
}

/**
 * 模拟 PointerDown 事件
 *
 * @param target 目标对象
 */
export function simulatePointerDown(target: { emit: (event: string, ...args: unknown[]) => unknown }): void {
  target.emit('pointerdown');
}

/**
 * 模拟 PointerUp 事件
 *
 * @param target 目标对象
 */
export function simulatePointerUp(target: { emit: (event: string, ...args: unknown[]) => unknown }): void {
  target.emit('pointerup');
}

// ============================================================================
// 3. 纹理模拟
// ============================================================================

/**
 * 为指定 Scene 注册 mock 纹理，使 `hasTexture()` 返回 true。
 *
 * @param scene MockScene 实例
 * @param key 纹理键
 */
export function mockTexture(
  scene: { textures: { exists: (key: string) => boolean } },
  key: string
): void {
  const originalExists = scene.textures.exists;
  const registered = new Set<string>();

  // If already patched, just register
  if ((scene as unknown as { __mockTextures?: Set<string> }).__mockTextures) {
    (scene as unknown as { __mockTextures: Set<string> }).__mockTextures.add(key);
    return;
  }

  registered.add(key);
  (scene as unknown as { __mockTextures: Set<string> }).__mockTextures = registered;

  scene.textures.exists = (k: string) => registered.has(k) || originalExists(k);
}

/**
 * 清除指定 Scene 的所有 mock 纹理注册
 *
 * @param scene MockScene 实例
 */
export function clearMockTextures(
  scene: { textures: { exists: (key: string) => boolean } }
): void {
  const registered = (scene as unknown as { __mockTextures?: Set<string> }).__mockTextures;
  if (registered) {
    registered.clear();
  }
}

// ============================================================================
// 4. 断言辅助
// ============================================================================

/**
 * 验证是否打印了纹理回退 warn。
 *
 * 使用方式：
 * ```ts
 * const warnSpy = vi.spyOn(console, 'warn').mockImplementation(() => {});
 * new UIButton(scene, { text: 'test', texture: 'missing' });
 * expectFallbackWarn(warnSpy, 'missing');
 * warnSpy.mockRestore();
 * ```
 *
 * @param warnSpy console.warn 的 spy
 * @param textureKey 期望的纹理键
 */
export function expectFallbackWarn(
  warnSpy: SpyInstance<unknown[], unknown>,
  textureKey: string
): void {
  expect(warnSpy).toHaveBeenCalledWith(`[UI] fallback: ${textureKey}`);
}

/**
 * 验证 tween 是否被创建。
 *
 * @param scene MockScene 实例
 * @param target 期望的 tween 目标
 * @returns 是否找到匹配的 tween
 */
export function expectTweenCreated(
  scene: { tweens: { tweens: Array<{ targets: unknown }> } },
  target: unknown
): boolean {
  return scene.tweens.tweens.some((t) => t.targets === target);
}

// ============================================================================
// 5. 通用辅助
// ============================================================================

/**
 * 创建可交互的 mock GameObject（用于测试事件绑定）
 *
 * @returns 模拟的 GameObject
 */
export function createMockGameObject(): {
  id: string;
  interactive: boolean;
  events: Record<string, ((...args: unknown[]) => void)[]>;
  setInteractive: (opts?: { useHandCursor?: boolean }) => unknown;
  disableInteractive: () => unknown;
  on: (event: string, handler: (...args: unknown[]) => void) => unknown;
  off: (event: string, handler: (...args: unknown[]) => void) => unknown;
  emit: (event: string, ...args: unknown[]) => unknown;
} {
  const obj = {
    id: Math.random().toString(36).slice(2),
    interactive: false,
    events: {} as Record<string, ((...args: unknown[]) => void)[]>,
    setInteractive(opts?: { useHandCursor?: boolean }) {
      this.interactive = true;
      return this;
    },
    disableInteractive() {
      this.interactive = false;
      return this;
    },
    on(event: string, handler: (...args: unknown[]) => void) {
      if (!this.events[event]) this.events[event] = [];
      this.events[event].push(handler);
      return this;
    },
    off(event: string, handler: (...args: unknown[]) => void) {
      if (!this.events[event]) return this;
      this.events[event] = this.events[event].filter((h) => h !== handler);
      return this;
    },
    emit(event: string, ...args: unknown[]) {
      (this.events[event] ?? []).forEach((h) => h(...args));
      return this;
    },
  };
  return obj;
}

/**
 * 清理函数 —— 每个测试后调用，重置全局状态。
 *
 * 使用方式：
 * ```ts
 * afterEach(() => {
 *   cleanup();
 * });
 * ```
 */
export function cleanup(): void {
  // 迁移时改为: const { cleanup: baseCleanup } = require('../../test-setup');
  // eslint-disable-next-line @typescript-eslint/no-var-requires
  const { cleanup: baseCleanup } = require('../../../../../../apps/phaser3/src/test-setup');
  baseCleanup();
}

// ============================================================================
// 6. Scene 生命周期辅助
// ============================================================================

/**
 * 模拟 Scene shutdown 事件，验证组件是否正确清理。
 *
 * @param scene MockScene 实例
 */
export function simulateSceneShutdown(
  scene: { emit: (event: string) => unknown }
): void {
  scene.emit('shutdown');
}
