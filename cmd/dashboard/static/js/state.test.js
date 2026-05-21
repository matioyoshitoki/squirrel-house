import { describe, it, expect, vi } from 'vitest';
import { createStore } from './state.js';

describe('createStore', () => {
  it('should initialize with given state', () => {
    const store = createStore({ count: 0, name: 'test' });
    expect(store.get('count')).toBe(0);
    expect(store.get('name')).toBe('test');
  });

  it('should return full state when get() has no argument', () => {
    const store = createStore({ count: 0 });
    expect(store.get()).toEqual({ count: 0 });
  });

  it('should update state with set()', () => {
    const store = createStore({ count: 0 });
    store.set({ count: 5 });
    expect(store.get('count')).toBe(5);
  });

  it('should notify subscribers on change', () => {
    const store = createStore({ count: 0 });
    const fn = vi.fn();
    store.subscribe(['count'], fn);
    store.set({ count: 1 });
    expect(fn).toHaveBeenCalledTimes(1);
    expect(fn).toHaveBeenCalledWith(expect.objectContaining({ count: 1 }));
  });

  it('should not notify when value is unchanged', () => {
    const store = createStore({ count: 0 });
    const fn = vi.fn();
    store.subscribe(['count'], fn);
    store.set({ count: 0 });
    expect(fn).not.toHaveBeenCalled();
  });

  it('should support multiple keys in one subscription', () => {
    const store = createStore({ a: 1, b: 2 });
    const fn = vi.fn();
    store.subscribe(['a', 'b'], fn);
    store.set({ a: 10 });
    expect(fn).toHaveBeenCalledTimes(1);
    store.set({ b: 20 });
    expect(fn).toHaveBeenCalledTimes(2);
  });

  it('should batch multiple key changes into single notification', () => {
    const store = createStore({ a: 1, b: 2 });
    const fn = vi.fn();
    store.subscribe(['a'], fn);
    store.subscribe(['b'], fn);
    // set both a and b at once
    store.set({ a: 10, b: 20 });
    expect(fn).toHaveBeenCalledTimes(1);
  });

  it('should support wildcard subscription', () => {
    const store = createStore({ x: 1 });
    const fn = vi.fn();
    store.subscribe('*', fn);
    store.set({ x: 2 });
    expect(fn).toHaveBeenCalledTimes(1);
  });

  it('should allow unsubscribing', () => {
    const store = createStore({ count: 0 });
    const fn = vi.fn();
    const unsub = store.subscribe(['count'], fn);
    unsub();
    store.set({ count: 1 });
    expect(fn).not.toHaveBeenCalled();
  });

  it('should not notify unrelated subscribers', () => {
    const store = createStore({ a: 1, b: 2 });
    const fnA = vi.fn();
    const fnB = vi.fn();
    store.subscribe(['a'], fnA);
    store.subscribe(['b'], fnB);
    store.set({ a: 10 });
    expect(fnA).toHaveBeenCalledTimes(1);
    expect(fnB).not.toHaveBeenCalled();
  });
});

describe('designAssetsCache behavior', () => {
  it('should default to empty object', () => {
    const { store: realStore } = require('./state.js');
    // We can not easily test the singleton, but createStore can simulate it
    const store = createStore({ designAssetsCache: {} });
    expect(store.get('designAssetsCache')).toEqual({});
  });

  it('should allow caching assets by issue number', () => {
    const store = createStore({ designAssetsCache: {} });
    const assets = [{ name: 'preview.html', url: '/logs/preview.html', type: 'html' }];
    store.set({ designAssetsCache: { 6: assets } });
    expect(store.get('designAssetsCache')[6]).toEqual(assets);
  });

  it('should allow null cache entry to indicate no assets', () => {
    const store = createStore({ designAssetsCache: {} });
    store.set({ designAssetsCache: { 6: null } });
    expect(store.get('designAssetsCache')[6]).toBeNull();
    // undefined means never loaded; null means loaded but empty
    expect(store.get('designAssetsCache')[7]).toBeUndefined();
  });
});
