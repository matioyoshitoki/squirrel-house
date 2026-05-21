import { store } from './state.js';
import { api } from './api.js';
import { showToast, showGlobalLoading, hideGlobalLoading, switchStage, updateStageBadges } from './ui.js';
import { updateRunningTimes } from './utils.js';
import * as renderers from './renderers.js';
import {
    openLogTerminalByFile, closeTerminal, clearTerminal, toggleRenderMode
} from './terminal.js';
import {
    openDesignViewer, closeDesignViewer, loadDesignAssets
} from './design-viewer.js';
import { startEventStream, onEvent, onConnectionChange } from './events.js';
import * as actions from './actions.js';

// ─── Stage switching ───
document.querySelectorAll('.stage-btn').forEach(btn => {
    btn.addEventListener('click', () => switchStage(btn.dataset.stage));
});

// ─── Refresh button ───
document.getElementById('refresh-btn')?.addEventListener('click', loadAll);

// ─── Project selector ───
document.getElementById('project-selector')?.addEventListener('change', async (e) => {
    try {
        await api.switchProject(e.target.value);
        showToast('已切换项目: ' + e.target.value);
        loadAll();
    } catch (err) {
        showToast('切换项目失败: ' + err.message, true);
    }
});

// ─── PM Modal ───
document.getElementById('open-pm-btn')?.addEventListener('click', () => {
    const modal = document.getElementById('pm-modal');
    modal?.classList.remove('hidden');
    modal?.classList.add('flex');
});
document.getElementById('close-pm-btn')?.addEventListener('click', closePMModal);
document.getElementById('start-pm-btn')?.addEventListener('click', startPM);
document.getElementById('pm-modal')?.addEventListener('click', (e) => {
    if (e.target === e.currentTarget) closePMModal();
});

function closePMModal() {
    const modal = document.getElementById('pm-modal');
    modal?.classList.add('hidden');
    modal?.classList.remove('flex');
    const title = document.getElementById('pm-title');
    const desc = document.getElementById('pm-desc');
    if (title) title.value = '';
    if (desc) desc.value = '';
}

async function startPM() {
    const title = document.getElementById('pm-title')?.value.trim();
    const desc = document.getElementById('pm-desc')?.value.trim();
    if (!title) {
        showToast('请输入标题', true);
        return;
    }
    try {
        const result = await api.startPM(title, desc);
        showToast('PM 任务已启动');
        closePMModal();
        // 乐观更新
        if (result.issueNumber) {
            const current = store.get('runningTasksData');
            store.set({ runningTasksData: [...current, {
                id: 'pm-' + result.issueNumber,
                type: 'pm',
                status: 'running',
                targetId: result.issueNumber,
                targetTitle: title,
                logFile: result.logFile,
                startTime: new Date().toISOString()
            }]});
        }
        loadAll();
    } catch (err) {
        showToast('启动失败: ' + err.message, true);
    }
}

// ─── Report Modal ───
document.getElementById('close-report-btn')?.addEventListener('click', actions.closeReportModal);
document.getElementById('report-modal')?.addEventListener('click', (e) => {
    if (e.target === e.currentTarget) actions.closeReportModal();
});

// ─── Data loaders ───
// All loaders write to the store. Rendering is handled by subscriptions below.

export async function loadProjects() {
    try {
        const data = await api.projects();
        store.set({ currentProject: data.currentProject });
        const sel = document.getElementById('project-selector');
        if (sel) {
            sel.innerHTML = '';
            data.projects.forEach((p, idx) => {
                const opt = document.createElement('option');
                opt.value = p.name;
                opt.textContent = p.name;
                if (idx === data.currentIndex) opt.selected = true;
                sel.appendChild(opt);
            });
        }
    } catch (err) {
        showToast('加载项目失败', true);
    }
}

export async function loadAll() {
    const btn = document.getElementById('refresh-btn');
    btn?.classList.add('rotate-180');
    setTimeout(() => btn?.classList.remove('rotate-180'), 500);

    showGlobalLoading();
    await Promise.all([
        loadProjects(),
        loadIssues(),
        actions.refreshRunningTasks(),
        actions.refreshPullRequests(),
        actions.refreshMergedPRs(),
        actions.refreshReviewTasks(),
        actions.refreshDocGardenerTasks(),
        actions.refreshRecentLogs(),
        actions.refreshPromptEvolutionStats(),
    ]);
    hideGlobalLoading();
}

export async function loadTasks() {
    showGlobalLoading();
    await Promise.all([
        actions.refreshRunningTasks(),
        actions.refreshPullRequests(),
        actions.refreshReviewTasks(),
        actions.refreshDocGardenerTasks(),
        actions.refreshRecentLogs(),
    ]);
    hideGlobalLoading();
}

async function loadIssues() {
    try {
        const data = await api.issues();
        store.set({ issuesData: data });
    } catch (err) {
        showToast('加载 Issues 失败', true);
    }
}

// ─── Store subscriptions ───
// Each view subscribes to the state slices it needs. Since store.set() batches
// key changes into a single notification per call, a listener is invoked at most
// once even if several of its subscribed keys change together.

store.subscribe(['runningTasksData', 'currentProject', 'designAssetsCache'], (s) => {
    renderers.renderRunningTasks(s.runningTasksData, s.currentProject, s.designAssetsCache);
});

store.subscribe(['runningTasksData'], (s) => {
    renderers.renderInterruptedTasks(s.runningTasksData);
});

store.subscribe(['docGardenerTasksData'], (s) => {
    renderers.renderDocGardenerTasks(s.docGardenerTasksData);
});

store.subscribe(['issuesData', 'runningTasksData', 'currentProject', 'designAssetsCache'], (s) => {
    renderers.renderIssues(s.issuesData, s.runningTasksData, s.currentProject, s.designAssetsCache);
    // Preload design assets for issues on demand
    s.issuesData.forEach(issue => {
        if (s.designAssetsCache[issue.number] === undefined) {
            loadDesignAssets(issue.number);
        }
    });
});

store.subscribe(['pullRequestsData', 'reviewTasksData', 'runningTasksData', 'designAssetsCache'], (s) => {
    renderers.renderPullRequests(s.pullRequestsData, s.reviewTasksData, s.runningTasksData, s.designAssetsCache);
    // Preload design assets referenced by PR titles on demand
    s.pullRequestsData.forEach(pr => {
        const match = pr.title.match(/design:\s+assets\s+for\s+issue\s+#(\d+)/i);
        if (match) {
            const num = parseInt(match[1], 10);
            if (s.designAssetsCache[num] === undefined) {
                loadDesignAssets(num);
            }
        }
    });
});

store.subscribe(['mergedPRsData'], (s) => {
    renderers.renderMergedPRs(s.mergedPRsData);
});

store.subscribe(['allLogsData'], (s) => {
    renderers.renderLogs(s.allLogsData.slice(0, 5), 'recent-logs');
});

store.subscribe(['promptEvolutionStatsData'], (s) => {
    renderers.renderPromptEvolutionStats(s.promptEvolutionStatsData);
});

store.subscribe(['runningTasksData', 'pullRequestsData', 'issuesData', 'mergedPRsData'], (s) => {
    updateStageBadges(s);
});

// ─── SSE + fallback polling ───
// No polling while SSE is connected. When disconnected, fall back
// to aggressive polling so the UI never goes stale.
let tasksInterval = null;
let allInterval = null;

function startFallbackPolling() {
    if (tasksInterval || allInterval) return;
    tasksInterval = setInterval(loadTasks, 3000);
    allInterval = setInterval(loadAll, 30000);
}

function stopPolling() {
    if (tasksInterval) { clearInterval(tasksInterval); tasksInterval = null; }
    if (allInterval) { clearInterval(allInterval); allInterval = null; }
}

onConnectionChange((connected) => {
    if (connected) {
        stopPolling();
    } else {
        startFallbackPolling();
    }
});

onEvent('connected', () => {
    // Fresh full sync on (re)connect to catch anything missed
    loadAll();
});

onEvent('task:changed', (data) => {
    loadTasks();
    // 设计任务成功完成后，强制刷新对应 issue 的设计资产缓存
    // 防止资产目录在首次加载后才生成导致按钮不显示
    if (data && data.type === 'design' && data.status === 'success' && data.targetId) {
        loadDesignAssets(data.targetId);
    }
});

onEvent('pr:merged', () => {
    loadAll();
});

onEvent('project:switched', () => {
    window.location.reload();
});

// ─── Boot ───
loadAll();
startEventStream();
startFallbackPolling(); // poll aggressively until SSE connects
setInterval(updateRunningTimes, 1000);
switchStage('all');

window.addEventListener('beforeunload', () => {
    import('./events.js').then(m => m.closeEventStream());
});

function copyTaskId(id) {
    navigator.clipboard.writeText(id).then(() => {
        showToast('已复制任务ID: ' + id);
    }).catch(() => {
        showToast('复制失败', true);
    });
}

// ─── Expose to global for inline onclick ───
window.loadAll = loadAll;
window.switchStage = switchStage;
window.closePMModal = closePMModal;
window.startPM = startPM;
window.closeReportModal = actions.closeReportModal;
window.copyTaskId = copyTaskId;

window.openLogTerminalByFile = openLogTerminalByFile;
window.closeTerminal = closeTerminal;
window.clearTerminal = clearTerminal;
window.toggleRenderMode = toggleRenderMode;

window.openDesignViewer = openDesignViewer;
window.closeDesignViewer = closeDesignViewer;
window.toggleDesignGroup = window.toggleDesignGroup;
window.selectDesignAsset = window.selectDesignAsset;

window.stopTask = actions.stopTask;
window.stopReviewTask = actions.stopReviewTask;
window.resumeDev = actions.resumeDev;
window.resumeRework = actions.resumeRework;
window.resumeDocGardener = actions.resumeDocGardener;
window.viewIssue = actions.viewIssue;
window.viewReviewReport = actions.viewReviewReport;
window.startDev = actions.startDev;
window.runE2E = actions.runE2E;
window.startDesign = actions.startDesign;
window.agentReviewPR = actions.agentReviewPR;
window.mergePR = actions.mergePR;
window.runDocGardener = actions.runDocGardener;
window.runMergedDocGardener = actions.runMergedDocGardener;
window.reworkPR = actions.reworkPR;
