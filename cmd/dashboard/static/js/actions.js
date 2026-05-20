import { api } from './api.js';
import { store } from './state.js';
import { showToast, showGlobalLoading, hideGlobalLoading, setButtonLoading, showConfirmDialog } from './ui.js';
import { escapeHtml, getRepoBaseUrl } from './utils.js';
import { openLogTerminalByFile } from './terminal.js';

// ─── Refresh helpers (lightweight data loaders) ───

export async function refreshRunningTasks() {
    try { store.set({ runningTasksData: await api.runningTasks() }); } catch {}
}

export async function refreshPullRequests() {
    try { store.set({ pullRequestsData: await api.pullRequests() }); } catch {}
}

export async function refreshMergedPRs() {
    try { store.set({ mergedPRsData: await api.mergedPRs() }); } catch {}
}

export async function refreshReviewTasks() {
    try { store.set({ reviewTasksData: await api.reviewTasks() }); } catch {}
}

export async function refreshDocGardenerTasks() {
    try { store.set({ docGardenerTasksData: await api.docGardenerTasks() }); } catch {}
}

export async function refreshRecentLogs() {
    try { store.set({ allLogsData: await api.logs() }); } catch {}
}

export async function refreshPromptEvolutionStats() {
    try { store.set({ promptEvolutionStatsData: await api.promptEvolutionStats() }); } catch {}
}

export async function refreshAll() {
    await Promise.all([
        refreshRunningTasks(),
        refreshPullRequests(),
        refreshMergedPRs(),
        refreshReviewTasks(),
        refreshDocGardenerTasks(),
        refreshRecentLogs(),
    ]);
}

// ─── Action handlers ───

async function resumeTask(event, targetId, title, apiPath, logFile) {
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('恢复任务中...');
    try {
        const response = await fetch(apiPath, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ issueNumber: targetId, issueTitle: title })
        });
        const result = await response.json();
        if (!response.ok) {
            showToast(result.error || '恢复失败', true);
            return;
        }
        showToast(result.message);
        await refreshRunningTasks();
        setTimeout(() => {
            const tasks = store.get('runningTasksData');
            const task = tasks.find(t => t.targetId === targetId && t.status === 'running');
            if (task) {
                openLogTerminalByFile(task.logFile);
            }
        }, 2000);
    } catch (error) {
        showToast('恢复失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function resumeDev(event, issueNumber, logFile) {
    if (!await showConfirmDialog('恢复开发', '确认恢复 Issue #' + issueNumber + ' 的开发?\n\nAgent 将基于之前的进度继续。')) return;
    await resumeTask(event, issueNumber, 'Issue #' + issueNumber, '/api/resume-dev', logFile);
}

export async function resumeRework(event, prNumber, logFile) {
    if (!await showConfirmDialog('恢复返工修复', '确认恢复 PR #' + prNumber + ' 的返工修复?\n\nAgent 将基于之前的进度继续。')) return;
    await resumeTask(event, prNumber, 'PR #' + prNumber, '/api/resume-dev', logFile);
}

export async function resumeDocGardener(event, prNumber, logFile) {
    if (!await showConfirmDialog('恢复文档检查', '确认恢复 PR #' + prNumber + ' 的文档检查?\n\nAgent 将基于之前的进度继续。')) return;
    await resumeTask(event, prNumber, 'PR #' + prNumber, '/api/resume-dev', logFile);
}

export async function stopTask(event, number, taskType) {
    const typeLabel = taskType === 'design' ? '设计' : taskType === 'pm' ? 'PRD 产出' : '开发';
    if (!await showConfirmDialog('确认停止', '确认停止 Issue #' + number + ' 的' + typeLabel + '?')) return;
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('停止任务中...');
    try {
        const result = await api.stopTask(number, taskType || 'dev');
        showToast(result.message);
        await refreshRunningTasks();
    } catch (error) {
        showToast('停止失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function stopReviewTask(event, prNumber) {
    if (!await showConfirmDialog('确认停止 Review', '确认停止 PR #' + prNumber + ' 的 Review?')) return;
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('停止 Review 中...');
    try {
        const result = await api.stopReviewTask(prNumber);
        showToast(result.message);
        await refreshReviewTasks();
        await refreshPullRequests();
    } catch (error) {
        showToast('停止失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function agentReviewPR(event, prNumber, branch) {
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('Agent Review 启动中...');
    try {
        const result = await api.reviewPR(prNumber, branch);
        showToast(result.message);
        // 乐观更新
        const current = store.get('reviewTasksData');
        store.set({ reviewTasksData: [...current, {
            id: 'review-' + prNumber,
            type: 'review',
            status: 'running',
            targetId: prNumber,
            targetTitle: 'PR #' + prNumber,
            branch: branch,
            logFile: result.logFile,
            startTime: new Date().toISOString()
        }]});
        await refreshReviewTasks();
        await refreshPullRequests();
        setTimeout(() => {
            const tasks = store.get('reviewTasksData');
            const task = tasks.find(t => t.targetId === prNumber && t.status === 'running');
            if (task && task.logFile) {
                openLogTerminalByFile(task.logFile.replace(/^logs\//, ''));
            }
        }, 2000);
    } catch (error) {
        showToast('Review 启动失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function mergePR(event, prNumber) {
    if (!await showConfirmDialog('确认合并', '确认合并 PR #' + prNumber + '?')) return;
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('合并 PR 中...');
    try {
        const result = await api.mergePR(prNumber);
        showToast(result.message);
        await refreshPullRequests();
        await refreshMergedPRs();
    } catch (error) {
        showToast('合并失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function runDocGardener(event, prNumber, baseRefName) {
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('Doc Gardener 启动中...');
    try {
        const result = await api.docGardener(prNumber, baseRefName);
        showToast(result.message);
        // 乐观更新
        const current = store.get('docGardenerTasksData');
        store.set({ docGardenerTasksData: [...current, {
            id: 'doc-gardener-' + prNumber,
            type: 'doc-gardener',
            status: 'running',
            targetId: prNumber,
            prNumber: prNumber,
            targetTitle: 'PR #' + prNumber,
            logFile: result.logFile,
            startTime: new Date().toISOString()
        }]});
        await refreshDocGardenerTasks();
        setTimeout(() => {
            const tasks = store.get('docGardenerTasksData');
            const task = tasks.find(t => t.prNumber === prNumber && t.status === 'running');
            if (task && task.logFile) {
                openLogTerminalByFile(task.logFile.replace(/^logs\//, ''));
            }
        }, 2000);
    } catch (error) {
        showToast('文档检查启动失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function runMergedDocGardener(event, number) {
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('Merged Doc Gardener 启动中...');
    try {
        const result = await api.mergedDocGardener(number);
        showToast(result.message);
        // 乐观更新
        const current = store.get('docGardenerTasksData');
        store.set({ docGardenerTasksData: [...current, {
            id: 'doc-gardener-' + number,
            type: 'doc-gardener',
            status: 'running',
            targetId: number,
            prNumber: number,
            targetTitle: 'PR #' + number,
            merged: true,
            logFile: result.logFile,
            startTime: new Date().toISOString()
        }]});
        await refreshDocGardenerTasks();
        if (result.logFile) openLogTerminalByFile(result.logFile);
    } catch (error) {
        showToast('文档检查启动失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function reworkPR(event, number) {
    if (!await showConfirmDialog('确认标记返工', '确认标记 PR #' + number + ' 需要返工?\n\n这将在飞书通知开发者进行修改。')) return;
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('标记返工中...');
    try {
        const result = await api.reworkPR(number);
        showToast(result.message);
        await refreshReviewTasks();
        await refreshPullRequests();
        await refreshRunningTasks();
    } catch (error) {
        showToast('标记返工失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function startDev(event, issueNumber, issueTitle, issueType) {
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('启动开发任务中...');
    try {
        const result = await api.startDev(issueNumber, issueTitle, issueType);
        showToast(result.message);
        // 乐观更新：立即在列表中渲染临时任务
        const current = store.get('runningTasksData');
        store.set({ runningTasksData: [...current, {
            id: 'dev-' + issueNumber,
            type: 'dev',
            status: 'running',
            targetId: issueNumber,
            targetTitle: issueTitle,
            logFile: result.logFile,
            startTime: new Date().toISOString()
        }]});
        await refreshRunningTasks();
        setTimeout(() => {
            const tasks = store.get('runningTasksData');
            const task = tasks.find(t => t.targetId === issueNumber && t.status === 'running');
            if (task) openLogTerminalByFile(task.logFile);
        }, 2000);
    } catch (error) {
        showToast('启动失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function runE2E(event, issueNumber, branch, taskKind) {
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('启动 E2E 测试中...');
    try {
        const result = await api.runE2E(issueNumber, branch, taskKind);
        showToast(result.message);
        // 乐观更新
        const current = store.get('runningTasksData');
        store.set({ runningTasksData: [...current, {
            id: 'e2e-' + issueNumber,
            type: 'e2e',
            status: 'running',
            targetId: issueNumber,
            targetTitle: 'E2E for Issue #' + issueNumber,
            startTime: new Date().toISOString()
        }]});
        await refreshRunningTasks();
    } catch (error) {
        showToast('启动 E2E 失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export async function startDesign(event, issueNumber, issueTitle) {
    const btn = event && event.target;
    setButtonLoading(btn, true);
    showGlobalLoading('启动设计任务中...');
    try {
        const result = await api.startDesign(issueNumber, issueTitle);
        showToast(result.message);
        // 乐观更新
        const current = store.get('runningTasksData');
        store.set({ runningTasksData: [...current, {
            id: 'design-' + issueNumber,
            type: 'design',
            status: 'running',
            targetId: issueNumber,
            targetTitle: issueTitle,
            logFile: result.logFile,
            startTime: new Date().toISOString()
        }]});
        await refreshRunningTasks();
        setTimeout(() => {
            const tasks = store.get('runningTasksData');
            const task = tasks.find(t => t.targetId === issueNumber && t.status === 'running');
            if (task) openLogTerminalByFile(task.logFile);
        }, 2000);
    } catch (error) {
        showToast('启动失败: ' + error.message, true);
    } finally {
        setButtonLoading(btn, false);
        hideGlobalLoading();
    }
}

export function viewIssue(number) {
    const proj = store.get('currentProject');
    const base = getRepoBaseUrl(proj);
    window.open(base ? base + '/issues/' + number : '#', '_blank');
}

export async function viewReviewReport(number) {
    try {
        const data = await api.reviewReport(number);
        if (data.status === 'not_found') {
            showToast('未找到该 PR 的 Review 记录', true);
            return;
        }

        const modal = document.getElementById('report-modal');
        const titleEl = document.getElementById('report-title');
        const metaEl = document.getElementById('report-meta');
        const bodyEl = document.getElementById('report-body');
        const logSection = document.getElementById('report-log-section');
        const logEl = document.getElementById('report-log');

        titleEl.textContent = 'Review 报告 - PR #' + number;

        const statusText = data.status === 'success' ? '完成' : '失败';
        const resultText = data.result === 'passed' ? '通过' : '需修改';

        metaEl.innerHTML =
            '<div><span class="text-muted-foreground">状态:</span> ' + statusText + '</div>' +
            (data.result ? '<div><span class="text-muted-foreground">结果:</span> ' + resultText + '</div>' : '') +
            '<div><span class="text-muted-foreground">开始:</span> ' + new Date(data.startTime).toLocaleString() + '</div>' +
            (data.endTime ? '<div><span class="text-muted-foreground">结束:</span> ' + new Date(data.endTime).toLocaleString() + '</div>' : '');

        const reportText = (data.report || '暂无报告内容').trim().replace(/\n{3,}/g, '\n\n');
        bodyEl.innerHTML = typeof marked !== 'undefined' ? marked.parse(reportText) : escapeHtml(reportText);

        if (data.log) {
            logSection.classList.remove('hidden');
            logEl.textContent = data.log.substring(0, 5000) + (data.log.length > 5000 ? '\n\n... (日志已截断)' : '');
        } else {
            logSection.classList.add('hidden');
        }

        modal.classList.remove('hidden');
        modal.classList.add('flex');
    } catch (error) {
        showToast('查看报告失败: ' + error.message, true);
    }
}

export function closeReportModal() {
    const modal = document.getElementById('report-modal');
    modal?.classList.add('hidden');
    modal?.classList.remove('flex');
}
