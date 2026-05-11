import { store } from './state.js';

const toastEl = document.getElementById('toast');
const globalLoadingEl = document.getElementById('global-loading');
const loadingTextEl = document.getElementById('loading-text');

export function showToast(message, isError = false) {
    toastEl.textContent = message;
    toastEl.classList.remove('hidden', 'bg-destructive');
    toastEl.classList.add('block');
    toastEl.style.animation = 'slideIn 0.2s ease-out';
    if (isError) {
        toastEl.classList.remove('bg-success');
        toastEl.classList.add('bg-destructive');
    } else {
        toastEl.classList.remove('bg-destructive');
        toastEl.classList.add('bg-success');
    }
    setTimeout(() => {
        toastEl.classList.remove('block');
        toastEl.style.animation = '';
        toastEl.classList.add('hidden');
    }, 3000);
}

export function showGlobalLoading(text) {
    const count = store.get('globalLoadingCount') + 1;
    store.set({ globalLoadingCount: count });
    if (count === 1) {
        globalLoadingEl.classList.remove('hidden');
        globalLoadingEl.classList.add('flex');
    }
    loadingTextEl.textContent = text || '执行中...';
}

export function hideGlobalLoading() {
    const count = Math.max(0, store.get('globalLoadingCount') - 1);
    store.set({ globalLoadingCount: count });
    if (count === 0) {
        globalLoadingEl.classList.remove('flex');
        globalLoadingEl.classList.add('hidden');
    }
}

export function setButtonLoading(btn, loading) {
    if (!btn) return;
    if (loading) {
        btn.disabled = true;
        btn.classList.add('btn-loading');
        btn.dataset.originalText = btn.textContent;
    } else {
        btn.disabled = false;
        btn.classList.remove('btn-loading');
    }
}

export function switchStage(stage) {
    store.set({ currentStage: stage });

    // Update nav buttons
    document.querySelectorAll('.stage-btn').forEach(btn => {
        btn.classList.toggle('active', btn.dataset.stage === stage);
    });

    // Show/hide sections
    document.querySelectorAll('.stage-section').forEach(sec => {
        const secStage = sec.dataset.stage;
        let show = false;
        if (stage === 'all') {
            show = true;
        } else if (stage === 'dev') {
            show = secStage === 'running' || secStage === 'doc' || secStage === 'interrupted';
        } else {
            show = secStage === stage;
        }
        sec.style.display = show ? '' : 'none';
    });
}

export function showConfirmDialog(title, message) {
    return new Promise((resolve) => {
        const dialog = document.getElementById('confirm-dialog');
        const titleEl = document.getElementById('confirm-title');
        const messageEl = document.getElementById('confirm-message');
        const okBtn = document.getElementById('confirm-ok');
        const cancelBtn = document.getElementById('confirm-cancel');
        if (!dialog || !okBtn || !cancelBtn) { resolve(false); return; }

        titleEl.textContent = title || '确认';
        messageEl.textContent = message || '';
        dialog.classList.remove('hidden');
        dialog.classList.add('flex');

        const cleanup = () => {
            dialog.classList.add('hidden');
            dialog.classList.remove('flex');
            okBtn.onclick = null;
            cancelBtn.onclick = null;
            dialog.onclick = null;
        };

        okBtn.onclick = () => { cleanup(); resolve(true); };
        cancelBtn.onclick = () => { cleanup(); resolve(false); };
        dialog.onclick = (e) => { if (e.target === dialog) { cleanup(); resolve(false); } };
    });
}

export function updateStageBadges({ runningTasksData, pullRequestsData, issuesData, mergedPRsData }) {
    const running = (runningTasksData || []).filter(t => t.status === 'running').length;
    const interrupted = (runningTasksData || []).filter(t => t.status === 'interrupted').length;
    const prs = pullRequestsData || [];
    const allIssues = issuesData || [];
    const merged = (mergedPRsData || []).length;

    const plans = allIssues.filter(i => i.labels.some(l => l.name === 'plan-ready'));
    const prds = allIssues.filter(i => i.labels.some(l => l.name === 'prd-ready'));
    const triages = allIssues.filter(i => !i.labels.some(l => l.name === 'plan-ready' || l.name === 'prd-ready'));

    const setBadge = (id, count) => {
        const el = document.getElementById('stage-badge-' + id);
        if (el) {
            el.textContent = count;
            el.style.display = count > 0 ? '' : 'none';
        }
    };

    setBadge('all', running + interrupted + triages.length + plans.length + prds.length + prs.filter(p => !p.merged).length + merged);
    setBadge('triage', triages.length);
    setBadge('plan', plans.length);
    setBadge('prd', prds.length);
    setBadge('dev', running + interrupted);
    setBadge('review', prs.filter(p => !p.merged).length);
    setBadge('merged', merged);
}
