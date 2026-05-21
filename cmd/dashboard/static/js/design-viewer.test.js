import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';

// Provide minimal DOM before module imports
beforeEach(() => {
  document.body.innerHTML = `
    <div id="design-overlay" class="hidden"></div>
    <div id="design-file-list"></div>
    <div id="design-filename"></div>
    <div id="design-preview-body"></div>
    <div id="close-design-btn"></div>
    <div id="run-e2e-btn"></div>
    <div id="build-preview-btn"></div>
    <div id="preview-status"></div>
    <div id="design-feedback-btn"></div>
  `;
});

afterEach(() => {
  vi.restoreAllMocks();
});

// Mock modules before importing design-viewer
const mockStore = {
  state: {
    designAssetsCache: {},
    currentProject: null,
  },
  get(key) {
    return key !== undefined ? this.state[key] : this.state;
  },
  set(patch) {
    Object.assign(this.state, patch);
  },
};

vi.mock('./state.js', () => ({
  store: mockStore,
}));

const mockApi = {
  designAssets: vi.fn(),
  buildDesignPreview: vi.fn(),
  submitDesignFeedback: vi.fn(),
};

vi.mock('./api.js', () => ({
  api: mockApi,
}));

vi.mock('./ui.js', () => ({
  showToast: vi.fn(),
}));

// Now import the module under test
const { loadDesignAssets, openDesignViewer } = await import('./design-viewer.js');

describe('loadDesignAssets', () => {
  beforeEach(() => {
    mockStore.state.designAssetsCache = {};
    mockApi.designAssets.mockReset();
    // Clear any leftover DOM cards
    document.body.innerHTML = `
      <div id="design-overlay" class="hidden"></div>
      <div id="design-file-list"></div>
      <div id="design-filename"></div>
      <div id="design-preview-body"></div>
      <div id="close-design-btn"></div>
      <div id="run-e2e-btn"></div>
      <div id="build-preview-btn"></div>
      <div id="preview-status"></div>
      <div id="design-feedback-btn"></div>
    `;
  });

  it('should cache assets and return them on success', async () => {
    const assets = [{ name: 'preview.html', url: '/preview.html', type: 'html' }];
    mockApi.designAssets.mockResolvedValue({ assets });

    const result = await loadDesignAssets(6);

    expect(mockApi.designAssets).toHaveBeenCalledWith(6);
    expect(result).toEqual(assets);
    expect(mockStore.state.designAssetsCache[6]).toEqual(assets);
  });

  it('should set cache to null when assets are empty', async () => {
    mockApi.designAssets.mockResolvedValue({ assets: [] });

    const result = await loadDesignAssets(6);

    expect(result).toBeNull();
    expect(mockStore.state.designAssetsCache[6]).toBeNull();
  });

  it('should set cache to null on API error', async () => {
    mockApi.designAssets.mockRejectedValue(new Error('network error'));

    const result = await loadDesignAssets(6);

    expect(result).toBeNull();
    expect(mockStore.state.designAssetsCache[6]).toBeNull();
  });

  it('should not make concurrent requests for the same issue', async () => {
    const assets = [{ name: 'a.html', url: '/a.html', type: 'html' }];
    mockApi.designAssets.mockImplementation(() =>
      new Promise((resolve) => setTimeout(() => resolve({ assets }), 50))
    );

    const p1 = loadDesignAssets(6);
    const p2 = loadDesignAssets(6);

    expect(mockApi.designAssets).toHaveBeenCalledTimes(1);

    const [r1, r2] = await Promise.all([p1, p2]);
    expect(r1).toEqual(assets);
    expect(r2).toEqual(assets);
  });

  it('should inject badge and button into existing issue card', async () => {
    document.body.innerHTML = `
      <div data-issue="6" class="card">
        <div class="flex items-start justify-between">
          <a href="#">#6: Test Issue</a>
        </div>
        <div class="flex flex-wrap gap-2">
          <button class="btn btn-sm btn-purple" onclick="startDesign()">开始设计</button>
          <button class="btn btn-sm btn-ghost" onclick="viewIssue(6)">查看</button>
        </div>
      </div>
    `;

    const assets = [{ name: 'preview.html', url: '/preview.html', type: 'html' }];
    mockApi.designAssets.mockResolvedValue({ assets });

    await loadDesignAssets(6);

    const card = document.querySelector('[data-issue="6"]');
    const badge = card.querySelector('.design-ready-badge');
    const btn = card.querySelector('.design-view-btn');

    expect(badge).not.toBeNull();
    expect(badge.textContent).toContain('设计资产已就绪');
    expect(btn).not.toBeNull();
    expect(btn.textContent).toContain('查看设计');
  });

  it('should not inject duplicate badge or button', async () => {
    document.body.innerHTML = `
      <div data-issue="6" class="card">
        <div class="flex items-start justify-between">
          <a href="#">#6: Test Issue</a>
          <span class="design-ready-badge">已就绪</span>
        </div>
        <div class="flex flex-wrap gap-2">
          <button class="btn btn-sm btn-primary design-view-btn">查看设计</button>
          <button class="btn btn-sm btn-purple">开始设计</button>
        </div>
      </div>
    `;

    const assets = [{ name: 'preview.html', url: '/preview.html', type: 'html' }];
    mockApi.designAssets.mockResolvedValue({ assets });

    await loadDesignAssets(6);

    const card = document.querySelector('[data-issue="6"]');
    expect(card.querySelectorAll('.design-ready-badge').length).toBe(1);
    expect(card.querySelectorAll('.design-view-btn').length).toBe(1);
  });

  it('should distinguish undefined (never loaded) from null (loaded but empty)', async () => {
    // First call returns empty
    mockApi.designAssets.mockResolvedValueOnce({ assets: [] });
    await loadDesignAssets(6);
    expect(mockStore.state.designAssetsCache[6]).toBeNull();

    // Second call should still hit the API because loadDesignAssets does not guard against null
    mockApi.designAssets.mockResolvedValueOnce({
      assets: [{ name: 'x.html', url: '/x.html', type: 'html' }],
    });
    await loadDesignAssets(6);
    expect(mockStore.state.designAssetsCache[6]).toHaveLength(1);
  });
});

describe('openDesignViewer', () => {
  beforeEach(() => {
    mockStore.state.designAssetsCache = {};
    mockStore.state.currentProject = null;
    document.body.innerHTML = `
      <div id="design-overlay" class="hidden"></div>
      <div id="design-file-list"></div>
      <div id="design-filename"></div>
      <div id="design-preview-body"></div>
      <div id="close-design-btn"></div>
      <div id="run-e2e-btn" class="hidden"></div>
      <div id="build-preview-btn" class="hidden"></div>
      <div id="preview-status" class="hidden"></div>
      <div id="design-feedback-btn" class="hidden"></div>
    `;
  });

  it('should show toast when no assets cached', async () => {
    const { showToast } = await import('./ui.js');
    mockStore.state.designAssetsCache[6] = null;

    await openDesignViewer(6);

    expect(showToast).toHaveBeenCalledWith('暂无设计资产', true);
  });

  it('should load assets if cache is undefined before opening', async () => {
    mockApi.designAssets.mockResolvedValue({
      assets: [{ name: 'preview.html', url: '/preview.html', type: 'html' }],
    });

    await openDesignViewer(6);

    expect(mockApi.designAssets).toHaveBeenCalledWith(6);
    const overlay = document.getElementById('design-overlay');
    expect(overlay.classList.contains('hidden')).toBe(false);
    expect(overlay.classList.contains('flex')).toBe(true);
  });
});
