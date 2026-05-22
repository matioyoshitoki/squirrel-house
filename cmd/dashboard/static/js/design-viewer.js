import { store } from './state.js';
import { escapeHtml } from './utils.js';
import { api } from './api.js';
import { showToast } from './ui.js';

let currentDesignIssue = null;
let _designEscHandler = null;
const _loadingPromises = new Map();

const DEFAULT_GROUPS = [
    { key: 'preview', title: '🔥 交互预览', types: ['flutter-web', 'phaser3-web', 'html'] },
    { key: 'docs', title: '🌐 原型与文档', types: ['markdown'] },
    { key: 'visual', title: '🖼️ 设计图', types: ['image'] },
    { key: 'code', title: '💻 组件代码', types: ['code'], collapsed: true },
    { key: 'other', title: '📎 其他', types: ['file', 'audio', 'json', 'font'] },
];

/*
 * Load design assets for an issue/PR, cache them, and inject
 * "查看设计" buttons / badges into already-rendered cards.
 */
export async function loadDesignAssets(issueNumber) {
    if (_loadingPromises.has(issueNumber)) {
        return _loadingPromises.get(issueNumber);
    }

    const promise = (async () => {
        try {
            const data = await api.designAssets(issueNumber);
            const assets = data.assets || [];
            const cache = store.get('designAssetsCache');
            cache[issueNumber] = assets.length > 0 ? assets : null;
            store.set({ designAssetsCache: { ...cache } });

            // Update issue cards
            const issueCard = document.querySelector(`[data-issue="${issueNumber}"]`);
            const prCard = document.querySelector(`[data-design-issue="${issueNumber}"]`);

            [issueCard, prCard].forEach(card => {
                if (card && assets.length > 0) {
                    // Insert badge if not present
                    const header = card.querySelector('.flex.items-start.justify-between');
                    if (header && !header.querySelector('.design-ready-badge')) {
                        const badge = document.createElement('span');
                        badge.className = 'inline-flex items-center gap-1.5 rounded-md px-2.5 py-1 text-xs font-semibold bg-primary/15 text-primary design-ready-badge';
                        badge.style.marginLeft = '0.5rem';
                        badge.innerHTML = '🎨 设计资产已就绪';
                        header.appendChild(badge);
                    }

                    // Insert "查看设计" button if not present
                    const actions = card.querySelector('.flex.flex-wrap.gap-2');
                    if (actions && !actions.querySelector('.design-view-btn')) {
                        const btn = document.createElement('button');
                        btn.className = 'btn btn-sm btn-primary design-view-btn';
                        btn.innerHTML = '🖼️ 查看设计';
                        btn.onclick = () => openDesignViewer(issueNumber);
                        const startDesignBtn = Array.from(actions.children).find(el => el.innerHTML && el.innerHTML.includes('开始设计'));
                        if (startDesignBtn) {
                            actions.insertBefore(btn, startDesignBtn);
                        } else {
                            actions.insertBefore(btn, actions.firstChild);
                        }
                    }
                }
            });

            return store.get('designAssetsCache')[issueNumber];
        } catch (e) {
            const cache = store.get('designAssetsCache');
            cache[issueNumber] = null;
            store.set({ designAssetsCache: { ...cache } });
            return null;
        }
    })();

    _loadingPromises.set(issueNumber, promise);
    try {
        return await promise;
    } finally {
        _loadingPromises.delete(issueNumber);
    }
}

/*
 * Open the design overlay for a given issue.
 * Loads assets first if they are not yet cached.
 */
export async function openDesignViewer(issueNumber) {
    if (store.get('designAssetsCache')[issueNumber] === undefined) {
        await loadDesignAssets(issueNumber);
    }

    const assets = store.get('designAssetsCache')[issueNumber] || [];
    const overlay = document.getElementById('design-overlay');
    const fileList = document.getElementById('design-file-list');

    if (assets.length === 0) {
        showToast('暂无设计资产', true);
        return;
    }

    currentDesignIssue = issueNumber;
    window.currentDesignIssue = issueNumber; // 暴露给 inline onclick

    const projectName = store.get('currentProject')?.name || '';
    const groups = DEFAULT_GROUPS;

    let firstUrl = null;
    let firstExt = null;
    let html = '';

    groups.forEach(group => {
        const items = [];
        assets.forEach((asset, idx) => {
            if (group.types.includes(asset.type)) {
                items.push({ ...asset, _idx: idx });
            }
        });
        if (items.length === 0) return;
        if (firstUrl === null) {
            firstUrl = items[0].url;
            firstExt = items[0].type;
        }

        const isCollapsed = group.collapsed ? ' collapsed' : '';
        html += `<div class="design-file-group mb-1${isCollapsed}" data-group="${group.key}">`;
        html += `<div class="design-toggle flex cursor-pointer select-none items-center justify-between rounded-md px-3 py-2 text-[0.7rem] font-semibold uppercase tracking-wider text-muted-foreground transition-colors duration-150 hover:bg-secondary" onclick="toggleDesignGroup('${group.key}')">${group.title} <span class="text-[0.7rem] opacity-50">(${items.length})</span></div>`;
        html += `<div class="design-file-group-content${group.collapsed ? ' hidden' : ''}">`;
        items.forEach(asset => {
            let icon = '📝';
            if (asset.type === 'image') icon = '🖼️';
            else if (asset.type === 'html') icon = '🌐';
            else if (asset.type === 'flutter-web') icon = '🔥';
            else if (asset.type === 'phaser3-web') icon = '🎮';
            else if (asset.type === 'flutter-web') icon = '📱';
            else if (asset.type === 'code') icon = '💻';
            html += `<div class="design-thumbnail design-file-item flex cursor-pointer items-center gap-2 overflow-hidden text-ellipsis whitespace-nowrap rounded-lg border-l-transparent px-3 py-2.5 text-sm text-muted-foreground transition-all duration-150 hover:bg-secondary hover:text-foreground${asset._idx === 0 ? ' border-l-[3px] border-l-primary' : ''}" data-idx="${asset._idx}" title="${escapeHtml(asset.name)}" onclick="selectDesignAsset('${asset.url}', '${asset.type}')">` +
                `<span>${icon}</span>${escapeHtml(asset.name)}</div>`;
        });
        html += '</div></div>';
    });

    fileList.innerHTML = html;
    overlay.classList.remove('hidden');
    overlay.classList.add('flex');

    // Prevent body scroll
    document.body.style.overflow = 'hidden';

    // Keyboard handling
    if (!_designEscHandler) {
        _designEscHandler = (e) => {
            if (e.key === 'Escape') closeDesignViewer();
        };
    }
    document.addEventListener('keydown', _designEscHandler);

    // Backdrop click
    overlay.onclick = (e) => {
        if (e.target === overlay) closeDesignViewer();
    };

    // Close button
    const closeBtn = document.getElementById('close-design-btn');
    if (closeBtn) closeBtn.onclick = closeDesignViewer;

    // Build Preview button: show for phaser3 projects that have code but no preview
    const buildPreviewBtn = document.getElementById('build-preview-btn');
    const previewStatus = document.getElementById('preview-status');
    if (buildPreviewBtn) {
        const project = store.get('currentProject');
        const isPhaser3 = project?.techStack?.frontend === 'phaser3' || project?.techStack?.mobile === 'phaser3';
        const hasCode = assets.some(a => a.type === 'code' && a.name.startsWith('src/ui/'));
        const hasPreview = assets.some(a => a.type === 'phaser3-web' || a.type === 'flutter-web');
        if (isPhaser3 && hasCode && !hasPreview) {
            buildPreviewBtn.classList.remove('hidden');
            if (previewStatus) previewStatus.classList.add('hidden');
        } else {
            buildPreviewBtn.classList.add('hidden');
            if (previewStatus) previewStatus.classList.add('hidden');
        }
    }

    // Design feedback button: show when design assets exist
    const feedbackBtn = document.getElementById('design-feedback-btn');
    if (feedbackBtn) {
        if (assets.length > 0) {
            feedbackBtn.classList.remove('hidden');
        } else {
            feedbackBtn.classList.add('hidden');
        }
    }

    // Select first asset
    if (firstUrl !== null) {
        selectDesignAsset(firstUrl, firstExt);
    }
}

/*
 * Close the design overlay.
 */
export function closeDesignViewer() {
    const overlay = document.getElementById('design-overlay');
    if (!overlay) return;
    overlay.classList.add('hidden');
    overlay.classList.remove('flex');
    document.body.style.overflow = '';
    if (_designEscHandler) {
        document.removeEventListener('keydown', _designEscHandler);
    }
    currentDesignIssue = null;
}

/*
 * Toggle visibility of a design group (accordion style).
 */
export function toggleDesignGroup(groupName) {
    const group = document.querySelector(`.design-file-group[data-group="${groupName}"]`);
    if (!group) return;
    group.classList.toggle('collapsed');
    const content = group.querySelector('.design-file-group-content');
    if (content) content.classList.toggle('hidden');
    const toggle = group.querySelector('.design-toggle');
    if (toggle) {
        const isCollapsed = group.classList.contains('collapsed');
        toggle.classList.toggle('after:content-[\'_▶\']', isCollapsed);
        toggle.classList.toggle('after:content-[\'_▼\']', !isCollapsed);
    }
}

/*
 * Show an asset in the preview area.
 * `ext` is the asset type: image | html | flutter-web | code | markdown | file
 */
export function selectDesignAsset(url, ext) {
    const filenameEl = document.getElementById('design-filename');
    const body = document.getElementById('design-preview-body');
    if (!body) return;

    // 重置可能影响布局的类
    body.classList.remove('relative');

    // Update active thumbnail state
    document.querySelectorAll('.design-thumbnail').forEach(el => {
        const active = el.getAttribute('onclick') && el.getAttribute('onclick').includes(`'${url}'`);
        el.classList.toggle('border-l-[3px]', active);
        el.classList.toggle('border-l-primary', active);
        el.classList.toggle('border-l-transparent', !active);
    });

    // Try to derive a filename for the header from the URL
    const urlObj = new URL(url, window.location.href);
    const pathParts = urlObj.pathname.split('/');
    const name = decodeURIComponent(pathParts[pathParts.length - 1]) || '未命名';
    if (filenameEl) filenameEl.textContent = name;

    if (ext === 'image') {
        body.innerHTML = `<img class="max-h-full max-w-full rounded-lg border border-border object-contain" src="${url}" alt="${escapeHtml(name)}">`;
    } else if (ext === 'html' || ext === 'flutter-web' || ext === 'phaser3-web') {
        // 加时间戳防止 iframe 缓存旧内容
        const cacheBustUrl = url + (url.includes('?') ? '&' : '?') + '_t=' + Date.now();
        // 使用 absolute 定位确保 iframe 填满 flex 容器，避免 h-full 在 flex item 中失效
        body.classList.add('relative');
        body.innerHTML = `<iframe class="absolute inset-0 h-full w-full rounded-lg border-none bg-white" src="${cacheBustUrl}"></iframe>`;
    } else if (ext === 'audio') {
        body.innerHTML = `<audio class="w-full rounded-lg" controls src="${url}"></audio>`;
    } else if (ext === 'json') {
        body.innerHTML = '<div class="px-4 py-12 text-center text-sm text-muted-foreground">加载中...</div>';
        fetch(url)
            .then(r => r.text())
            .then(text => {
                try {
                    const obj = JSON.parse(text);
                    const formatted = JSON.stringify(obj, null, 2);
                    body.innerHTML = `<pre class="h-full w-full overflow-auto whitespace-pre-wrap rounded-lg bg-[#0d1117] p-4 font-mono text-sm leading-relaxed text-foreground"><code>${escapeHtml(formatted)}</code></pre>`;
                } catch (e) {
                    body.innerHTML = `<pre class="h-full w-full overflow-auto whitespace-pre-wrap rounded-lg bg-[#0d1117] p-4 font-mono text-sm leading-relaxed text-foreground"><code>${escapeHtml(text)}</code></pre>`;
                }
            })
            .catch(() => {
                body.innerHTML = '<div class="px-4 py-10 text-center text-[0.9375rem] text-muted-foreground">加载失败</div>';
            });
    } else if (ext === 'code') {
        body.innerHTML = '<div class="px-4 py-12 text-center text-sm text-muted-foreground">加载中...</div>';
        fetch(url)
            .then(r => r.text())
            .then(text => {
                const fileExt = name.split('.').pop().toLowerCase();
                body.innerHTML = `<pre class="h-full w-full overflow-auto whitespace-pre-wrap rounded-lg bg-[#0d1117] p-4 font-mono text-sm leading-relaxed text-foreground"><code>${highlightCode(text, fileExt)}</code></pre>`;
            })
            .catch(() => {
                body.innerHTML = '<div class="px-4 py-10 text-center text-[0.9375rem] text-muted-foreground">加载失败</div>';
            });
    } else if (ext === 'markdown') {
        body.innerHTML = '<div class="px-4 py-12 text-center text-sm text-muted-foreground">加载中...</div>';
        fetch(url)
            .then(r => r.text())
            .then(text => {
                const html = typeof marked !== 'undefined' ? marked.parse(text) : escapeHtml(text);
                body.innerHTML = `<div class="markdown-body h-full w-full overflow-auto rounded-lg bg-[#0d1117] p-6 text-sm leading-relaxed text-foreground">${html}</div>`;
            })
            .catch(() => {
                body.innerHTML = '<div class="px-4 py-10 text-center text-[0.9375rem] text-muted-foreground">加载失败</div>';
            });
    } else {
        body.innerHTML = '<div class="px-4 py-10 text-center text-[0.9375rem] text-muted-foreground">该文件类型暂不支持预览</div>';
    }
}

/*
 * Simple syntax highlighting for code preview.
 * Supports Dart and JS-like languages.
 */
export function highlightCode(code, ext) {
    const isDart = ext === 'dart';
    const isJsLike = ['js', 'ts', 'tsx', 'jsx', 'css', 'scss', 'less'].includes(ext);
    if (!isDart && !isJsLike) return escapeHtml(code);

    const keywords = new Set([
        'class', 'extends', 'implements', 'with', 'mixin', 'abstract', 'interface',
        'factory', 'static', 'final', 'const', 'late', 'required', 'void', 'return',
        'if', 'else', 'for', 'while', 'do', 'switch', 'case', 'break', 'continue',
        'default', 'try', 'catch', 'finally', 'throw', 'assert', 'async', 'await',
        'yield', 'sync', 'new', 'this', 'super', 'get', 'set', 'operator', 'external',
        'native', 'covariant', 'var', 'dynamic', 'typedef', 'enum', 'part', 'library',
        'export', 'import', 'show', 'hide', 'as', 'is', 'in', 'true', 'false', 'null',
        'function', 'let', 'typeof', 'instanceof'
    ]);
    const typeKeywords = new Set([
        'int', 'double', 'String', 'bool', 'List', 'Map', 'Set', 'Future', 'Stream',
        'Widget', 'BuildContext', 'State', 'StatefulWidget', 'StatelessWidget',
        'Element', 'Key', 'Object', 'dynamic', 'num'
    ]);

    let html = '';
    let i = 0;
    while (i < code.length) {
        // Multi-line comment /* */
        if (code.substr(i, 2) === '/*') {
            let end = code.indexOf('*/', i + 2);
            if (end === -1) { end = code.length; } else { end += 2; }
            html += `<span style="color:#6a9955;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Single-line comment //
        if (code.substr(i, 2) === '//') {
            let end = code.indexOf('\n', i);
            if (end === -1) end = code.length;
            html += `<span style="color:#6a9955;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Single-quoted string
        if (code[i] === "'") {
            let end = i + 1;
            while (end < code.length && code[end] !== "'") {
                if (code[end] === '\\' && end + 1 < code.length) end++;
                end++;
            }
            if (end < code.length) end++;
            html += `<span style="color:#ce9178;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Double-quoted string
        if (code[i] === '"') {
            let end = i + 1;
            while (end < code.length && code[end] !== '"') {
                if (code[end] === '\\' && end + 1 < code.length) end++;
                end++;
            }
            if (end < code.length) end++;
            html += `<span style="color:#ce9178;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Triple-quoted string '''
        if (code.substr(i, 3) === "'''") {
            let end = code.indexOf("'''", i + 3);
            if (end === -1) end = code.length; else end += 3;
            html += `<span style="color:#ce9178;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Triple-quoted string """
        if (code.substr(i, 3) === '"""') {
            let end = code.indexOf('"""', i + 3);
            if (end === -1) end = code.length; else end += 3;
            html += `<span style="color:#ce9178;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Annotation @
        if (code[i] === '@') {
            let end = i + 1;
            while (end < code.length && /[a-zA-Z0-9_]/.test(code[end])) end++;
            html += `<span style="color:#4ec9b0;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Identifier / keyword
        if (/[a-zA-Z_]/.test(code[i])) {
            let end = i + 1;
            while (end < code.length && /[a-zA-Z0-9_]/.test(code[end])) end++;
            const word = code.slice(i, end);
            const esc = escapeHtml(word);
            if (keywords.has(word)) {
                html += `<span style="color:#569cd6;">${esc}</span>`;
            } else if (typeKeywords.has(word) || /^[A-Z]/.test(word)) {
                html += `<span style="color:#4ec9b0;">${esc}</span>`;
            } else {
                html += esc;
            }
            i = end; continue;
        }
        // Number
        if (/\d/.test(code[i])) {
            let end = i + 1;
            while (end < code.length && /[\d.]/.test(code[end])) end++;
            html += `<span style="color:#b5cea8;">${escapeHtml(code.slice(i, end))}</span>`;
            i = end; continue;
        }
        // Other characters
        html += escapeHtml(code[i]);
        i++;
    }
    return html;
}

/**
 * Build design preview for Phaser3 projects via dashboard API.
 */
let _buildingPreview = false;
export async function buildDesignPreviewHandler(issueNumber) {
    if (_buildingPreview) return;
    _buildingPreview = true;

    const btn = document.getElementById('build-preview-btn');
    const statusEl = document.getElementById('preview-status');

    if (btn) {
        btn.disabled = true;
        btn.textContent = '🎨 生成中...';
    }
    if (statusEl) {
        statusEl.classList.remove('hidden');
        statusEl.textContent = '请求中...';
    }

    try {
        const res = await api.buildDesignPreview(issueNumber);
        if (res.success) {
            showToast('预览生成成功');
            if (statusEl) statusEl.textContent = '✅ 完成';
            // 先清空预览区域，防止显示旧内容
            const body = document.getElementById('design-preview-body');
            if (body) body.innerHTML = '<div class="text-sm text-muted-foreground">加载新预览中...</div>';
            // 刷新资产列表
            await loadDesignAssets(issueNumber);
            // 直接选中新生成的 preview，避免重新渲染整个 overlay 导致的闪烁
            const assets = store.get('designAssetsCache')[issueNumber] || [];
            const previewAsset = assets.find(a => a.type === 'phaser3-web' || a.type === 'flutter-web' || a.type === 'image');
            if (previewAsset) {
                selectDesignAsset(previewAsset.url, previewAsset.type);
            }
            // 刷新左侧文件列表（只更新 fileList，不关闭/重新打开 overlay）
            const fileList = document.getElementById('design-file-list');
            if (fileList) {
                const projectName = store.get('currentProject')?.name || '';
                const groups = DEFAULT_GROUPS;
                let firstUrl = null, firstExt = null, html = '';
                groups.forEach(group => {
                    const items = [];
                    assets.forEach((asset, idx) => {
                        if (group.types.includes(asset.type)) {
                            items.push({ ...asset, _idx: idx });
                        }
                    });
                    if (items.length === 0) return;
                    if (firstUrl === null) {
                        firstUrl = items[0].url;
                        firstExt = items[0].type;
                    }
                    const isCollapsed = group.collapsed ? ' collapsed' : '';
                    html += `<div class="design-file-group mb-1${isCollapsed}" data-group="${group.key}">`;
                    html += `<div class="design-toggle flex cursor-pointer select-none items-center justify-between rounded-md px-3 py-2 text-[0.7rem] font-semibold uppercase tracking-wider text-muted-foreground transition-colors duration-150 hover:bg-secondary" onclick="toggleDesignGroup('${group.key}')">${group.title} <span class="text-[0.7rem] opacity-50">(${items.length})</span></div>`;
                    html += `<div class="design-file-group-content${group.collapsed ? ' hidden' : ''}">`;
                    items.forEach(asset => {
                        let icon = '📝';
                        if (asset.type === 'image') icon = '🖼️';
                        else if (asset.type === 'html') icon = '🌐';
                        else if (asset.type === 'flutter-web') icon = '📱';
                        else if (asset.type === 'phaser3-web') icon = '🎮';
                        else if (asset.type === 'code') icon = '💻';
                        html += `<div class="design-thumbnail design-file-item flex cursor-pointer items-center gap-2 overflow-hidden text-ellipsis whitespace-nowrap rounded-lg border-l-transparent px-3 py-2.5 text-sm text-muted-foreground transition-all duration-150 hover:bg-secondary hover:text-foreground${asset._idx === 0 ? ' border-l-[3px] border-l-primary' : ''}" data-idx="${asset._idx}" title="${escapeHtml(asset.name)}" onclick="selectDesignAsset('${asset.url}', '${asset.type}')">` +
                            `<span>${icon}</span>${escapeHtml(asset.name)}</div>`;
                    });
                    html += '</div></div>';
                });
                fileList.innerHTML = html;
            }
            // 隐藏生成预览按钮（因为 preview 已存在）
            if (btn) btn.classList.add('hidden');
        } else {
            showToast(res.error || '预览生成失败', true);
            if (statusEl) statusEl.textContent = '❌ 失败';
        }
    } catch (e) {
        showToast('预览请求失败: ' + e.message, true);
        if (statusEl) statusEl.textContent = '❌ 失败';
    } finally {
        _buildingPreview = false;
        if (btn) {
            btn.disabled = false;
            btn.textContent = '🎨 生成预览';
        }
    }
}

/**
 * Show design feedback form
 */
export async function showDesignFeedbackForm(issueNumber) {
    const feedback = prompt('请输入设计反馈（如：SplashPage 需要添加加载动画、LoginPage 按钮颜色需要调整等）：');
    if (!feedback || !feedback.trim()) return;

    try {
        const res = await api.submitDesignFeedback(issueNumber, feedback.trim());
        if (res.success) {
            showToast('反馈已保存，下次点击「修改设计」会带上此反馈');
        } else {
            showToast(res.error || '保存反馈失败', true);
        }
    } catch (e) {
        showToast('提交反馈失败: ' + e.message, true);
    }
}

// Expose to window for inline onclick handlers
window.loadDesignAssets = loadDesignAssets;
window.openDesignViewer = openDesignViewer;
window.closeDesignViewer = closeDesignViewer;
window.toggleDesignGroup = toggleDesignGroup;
window.selectDesignAsset = selectDesignAsset;
window.highlightCode = highlightCode;
window.buildDesignPreviewHandler = buildDesignPreviewHandler;
window.showDesignFeedbackForm = showDesignFeedbackForm;
