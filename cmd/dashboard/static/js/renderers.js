import { escapeHtml, formatDuration, formatBytes, getIssueType, getIssueTypeDisplay } from './utils.js';

const CARD = 'card card-hover p-4';
const EMPTY = 'card p-8 text-center text-sm text-muted-foreground';

/* ─── Running Tasks ─── */
export function renderRunningTasks(tasks, currentProject, designAssetsCache) {
    const container = document.getElementById('running-tasks');
    if (!tasks) tasks = [];
    const running = tasks.filter(t => t.status === 'running');

    if (running.length === 0) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无运行中的任务</div>';
        return;
    }

    container.innerHTML = running.map(task => {
        const startTime = new Date(task.startTime);
        const duration = Date.now() - startTime.getTime();
        const isRework = task.type === 'rework';
        let taskLabel = '开发中';
        if (task.type === 'rework') taskLabel = '返工修复中';
        else if (task.type === 'design') taskLabel = '设计中';
        else if (task.type === 'pm') taskLabel = 'PRD 产出中';
        else if (task.type === 'doc-gardener') taskLabel = 'Doc Gardener';

        let designBtn = '';
        const assets = designAssetsCache && designAssetsCache[task.targetId];
        if (assets && assets.length > 0) {
            designBtn = '<button class="btn btn-sm btn-primary design-view-btn" onclick="openDesignViewer(' + task.targetId + ')">查看设计</button>';
        }

        return '<div class="' + CARD + ' pulse-running" data-issue="' + task.targetId + '">' +
            '<div class="flex items-start justify-between gap-3 mb-2">' +
                '<a href="' + (currentProject && currentProject.repo ? 'https://github.com/' + currentProject.repo + '/issues/' : '#') + '' + task.targetId + '" ' +
                   'class="text-sm font-semibold text-foreground hover:text-primary transition-colors" target="_blank">' +
                    '#' + task.targetId + ': ' + escapeHtml(task.targetTitle) +
                '</a>' +
                '<span class="running-time shrink-0 text-xs font-medium text-success" data-start="' + startTime.getTime() + '" data-label="' + taskLabel + '">' +
                    taskLabel + ' · ' + formatDuration(duration) +
                '</span>' +
            '</div>' +
            '<div class="mb-3 text-xs text-muted-foreground">' + startTime.toLocaleString() + ' 开始</div>' +
            '<div class="flex flex-wrap gap-2">' +
                '<button class="btn btn-sm btn-ghost" onclick="copyTaskId(\'' + escapeHtml(task.id) + '\')" title="复制任务ID">复制ID</button>' +
                designBtn +
                '<button class="btn btn-sm btn-primary" onclick="openLogTerminalByFile(\'' + task.logFile + '\')">监控</button>' +
                '<button class="btn btn-sm btn-destructive" onclick="stopTask(event, ' + task.targetId + ', \'' + task.type + '\')">停止</button>' +
            '</div>' +
        '</div>';
    }).join('');
}

/* ─── Interrupted Tasks ─── */
export function renderInterruptedTasks(tasks) {
    const container = document.getElementById('interrupted-tasks');
    if (!tasks) tasks = [];
    const interrupted = tasks.filter(t => t.status === 'interrupted');

    if (interrupted.length === 0) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无中断的任务</div>';
        return;
    }

    container.innerHTML = interrupted.map(task => {
        const logFile = task.logFile.replace(/^logs\//, '');
        const typeLabel = task.type === 'rework' ? '返工修复' :
                          task.type === 'doc-gardener' ? '文档检查' : '开发';
        const btnText = task.type === 'rework' ? '恢复返工' :
                        task.type === 'doc-gardener' ? '恢复文档检查' : '恢复开发';
        const resumeFn = task.type === 'rework' ? 'resumeRework' :
                         task.type === 'doc-gardener' ? 'resumeDocGardener' : 'resumeDev';
        const targetLabel = task.type === 'rework' || task.type === 'doc-gardener' ?
            'PR #' + task.targetId : 'Issue #' + task.targetId;
        return '<div class="' + CARD + ' border-l-[3px] border-l-warning" data-issue="' + task.targetId + '" data-type="' + task.type + '">' +
            '<div class="flex items-start justify-between gap-3 mb-2">' +
                '<span class="text-sm font-semibold text-foreground">' + targetLabel + '</span>' +
                '<span class="badge bg-warning/10 border-warning/20 text-warning">已中断 (' + typeLabel + ')</span>' +
            '</div>' +
            '<div class="mb-2 text-xs text-warning">任务因步骤限制中断，点击按钮继续完成</div>' +
            '<div class="mb-3 text-xs text-muted-foreground">' + new Date(task.startTime).toLocaleString() + '</div>' +
            '<div class="flex flex-wrap gap-2">' +
                '<button class="btn btn-sm btn-ghost" onclick="copyTaskId(\'' + escapeHtml(task.id) + '\')" title="复制任务ID">复制ID</button>' +
                '<button class="btn btn-sm btn-primary" onclick="' + resumeFn + '(event, ' + task.targetId + ', \'' + logFile + '\')">' + btnText + '</button>' +
                '<button class="btn btn-sm btn-outline" onclick="openLogTerminalByFile(\'' + logFile + '\')">查看日志</button>' +
            '</div>' +
        '</div>';
    }).join('');
}

/* ─── Doc Gardener Tasks ─── */
export function renderDocGardenerTasks(tasks) {
    const container = document.getElementById('doc-gardener-tasks');
    if (!tasks) tasks = [];
    const running = tasks.filter(t => t.status === 'running');

    if (running.length === 0) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无文档检查任务</div>';
        return;
    }

    container.innerHTML = running.map(task => {
        const startTime = new Date(task.startTime);
        const duration = Date.now() - startTime.getTime();
        const label = task.merged ? '已合并 PR 文档补全' : 'PR 文档检查';
        return '<div class="' + CARD + ' border-l-[3px] border-l-success" data-pr="' + task.prNumber + '">' +
            '<div class="flex items-start justify-between gap-3 mb-2">' +
                '<span class="text-sm font-semibold text-foreground">' + label + ' · PR #' + task.prNumber + '</span>' +
                '<span class="badge bg-success/10 border-success/20 text-success">检查中</span>' +
            '</div>' +
            '<div class="mb-3 text-xs text-muted-foreground">' + formatDuration(duration) + ' · ' + startTime.toLocaleString() + ' 开始</div>' +
            '<div class="flex flex-wrap gap-2">' +
                '<button class="btn btn-sm btn-primary" onclick="openLogTerminalByFile(\'' + task.logFile + '\')">监控</button>' +
            '</div>' +
        '</div>';
    }).join('');
}

/* ─── Issues ─── */
export function renderIssues(issues, runningTasks, currentProject, designAssetsCache) {
    if (!issues) issues = [];
    if (!runningTasks) runningTasks = [];

    const runningIds = new Set(runningTasks.filter(t => t.status === 'running').map(t => t.targetId));

    const plans = issues.filter(i => i.labels.some(l => l.name === 'plan-ready')).sort((a, b) => b.number - a.number);
    const prds = issues.filter(i => i.labels.some(l => l.name === 'prd-ready')).sort((a, b) => b.number - a.number);
    const triages = issues.filter(i => !i.labels.some(l => l.name === 'plan-ready' || l.name === 'prd-ready')).sort((a, b) => b.number - a.number);

    renderIssueGroup('plan-issues', plans, runningIds, currentProject, designAssetsCache, runningTasks);
    renderIssueGroup('prd-issues', prds, runningIds, currentProject, designAssetsCache, runningTasks);
    renderIssueGroup('triage-issues', triages, runningIds, currentProject, designAssetsCache, runningTasks);
}

function renderIssueGroup(containerId, issues, runningIds, currentProject, designAssetsCache, runningTasks) {
    const container = document.getElementById(containerId);
    if (!container) return;

    if (issues.length === 0) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无数据</div>';
        return;
    }

    container.innerHTML = issues.map(issue => {
        const isRunning = runningIds.has(issue.number);
        const labelsHtml = issue.labels.map(l => '<span class="badge bg-secondary border-secondary text-muted-foreground">' + l.name + '</span>').join('');
        const classification = issue.classification || 'full-stack';
        const designTask = runningTasks.find(t => t.type === 'design' && t.targetId === issue.number);
        const isDesigning = designTask && designTask.status === 'running';
        const issueType = getIssueType(issue.labels, issue.title, issue.body);
        const display = getIssueTypeDisplay(issueType);

        let actionsHtml = '<button class="btn btn-sm btn-ghost" onclick="viewIssue(' + issue.number + ')">查看</button>';

        const e2eRunning = runningTasks.find(t => t.type === 'e2e' && t.targetId === issue.number);
        const isE2ERunning = e2eRunning && e2eRunning.status === 'running';

        if (isRunning) {
            actionsHtml = '<button class="btn btn-sm btn-outline" disabled>运行中</button>' + actionsHtml;
        } else {
            const designAssets = designAssetsCache[issue.number];
            const hasDesignAssets = designAssets && designAssets.length > 0;
            let viewDesignBtn = '';
            if (hasDesignAssets) {
                viewDesignBtn = '<button class="btn btn-sm btn-primary design-view-btn" onclick="openDesignViewer(' + issue.number + ')">查看设计</button>';
            }
            let designBtn = '';
            if (classification !== 'backend-only') {
                if (hasDesignAssets) {
                    designBtn = '<button class="btn btn-sm btn-purple" onclick="startDesign(event, ' + issue.number + ', \'' +
                        escapeHtml(issue.title) + '\')">🔄 修改设计</button>';
                } else {
                    designBtn = '<button class="btn btn-sm btn-purple" onclick="startDesign(event, ' + issue.number + ', \'' +
                        escapeHtml(issue.title) + '\')">开始设计</button>';
                }
            }

            if (issueType === 'ui-feedback') {
                actionsHtml = designBtn + viewDesignBtn + actionsHtml;
            } else {
                let devBtn = '<button class="btn btn-sm ' + display.btnClass + '" onclick="startDev(event, ' + issue.number + ', \'' +
                    escapeHtml(issue.title) + '\', \'' + issueType + '\')">' + display.text + '</button>';
                if (isDesigning) {
                    devBtn = '<button class="btn btn-sm btn-outline" disabled title="设计任务进行中，建议等待完成后再开发">' + display.text + '（等待设计）</button>';
                }
                let e2eBtn = '';
                if (classification !== 'backend-only') {
                    if (isE2ERunning) {
                        e2eBtn = '<button class="btn btn-sm btn-outline" disabled>🧪 E2E 中</button>';
                    } else {
                        e2eBtn = '<button class="btn btn-sm btn-outline" onclick="runE2E(event, ' + issue.number + ', \'\', \'' + classification + '\')">🧪 运行 E2E</button>';
                    }
                }
                if (classification === 'backend-only') {
                    actionsHtml = devBtn + actionsHtml;
                } else {
                    actionsHtml = devBtn + e2eBtn + viewDesignBtn + designBtn + actionsHtml;
                }
            }
        }

        return '<div class="' + CARD + '" data-issue="' + issue.number + '">' +
            '<div class="flex items-start justify-between gap-3 mb-2">' +
                '<a href="' + (currentProject && currentProject.repo ? 'https://github.com/' + currentProject.repo + '/issues/' : '#') + '' + issue.number + '" ' +
                   'class="text-sm font-semibold text-foreground hover:text-primary transition-colors" target="_blank">' +
                    '#' + issue.number + ': ' + escapeHtml(issue.title) +
                '</a>' +
            '</div>' +
            '<div class="mb-2 text-xs text-muted-foreground">' + new Date(issue.updatedAt).toLocaleString() + '</div>' +
            '<div class="flex flex-wrap gap-1.5 mb-3">' + labelsHtml + '</div>' +
            '<div class="flex flex-wrap gap-2">' + actionsHtml + '</div>' +
        '</div>';
    }).join('');


}

/* ─── Pull Requests ─── */
export function renderPullRequests(prs, reviewTasks, runningTasks, designAssetsCache) {
    const container = document.getElementById('pending-prs');
    if (!reviewTasks) reviewTasks = [];
    if (!runningTasks) runningTasks = [];

    if (prs.length === 0) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无等待 Review 的 PR</div>';
        return;
    }

    const reviewingIds = new Set(reviewTasks.filter(t => t.status === 'running').map(t => t.targetId));
    const knownReviews = new Map(reviewTasks.filter(t => ['success','failed','stopped','interrupted'].includes(t.status)).map(t => [t.targetId, t]));

    container.innerHTML = prs.map(pr => {
        const time = pr.createdAt ? new Date(pr.createdAt).toLocaleString() : '未知';
        const isReviewing = reviewingIds.has(pr.number);
        const reviewTask = knownReviews.get(pr.number);

        const relatedReworkTask = runningTasks.find(t =>
            ['running', 'failed', 'interrupted'].includes(t.status) &&
            t.type === 'rework' &&
            t.prNumber === pr.number
        );
        const isReworkRunning = relatedReworkTask && relatedReworkTask.status === 'running';
        const isReworkFailed = relatedReworkTask && (relatedReworkTask.status === 'failed' || relatedReworkTask.status === 'interrupted');

        let statusHtml = '';
        let buttonsHtml = '';
        let extraClass = '';

        let designIssueNumber = null;
        const designMatch = pr.title.match(/design:\s+assets\s+for\s+issue\s+#(\d+)/i);
        if (designMatch) {
            designIssueNumber = parseInt(designMatch[1], 10);
        }

        if (isReworkRunning) {
            statusHtml = '<span class="badge bg-warning/10 border-warning/20 text-warning ml-2">返工修复中</span>';
            buttonsHtml = '<button class="btn btn-sm btn-outline" disabled>返工修复中</button>' +
                '<button class="btn btn-sm btn-primary" onclick="openLogTerminalByFile(\'' + relatedReworkTask.logFile + '\')">查看日志</button>' +
                '<button class="btn btn-sm btn-ghost" onclick="window.open(\'' + pr.url + '\', \'_blank\')">查看 PR</button>';
        } else if (isReworkFailed) {
            statusHtml = '<span class="badge bg-destructive/10 border-destructive/20 text-destructive ml-2">返工失败</span>';
            extraClass = ' border-l-[3px] border-l-warning';
            const logFile = relatedReworkTask.logFile.replace(/^logs\//, '');
            buttonsHtml = '<button class="btn btn-sm btn-warning" onclick="resumeRework(event, ' + pr.number + ', \'' + logFile + '\')">恢复返工</button>' +
                '<button class="btn btn-sm btn-primary" onclick="viewReviewReport(' + pr.number + ')">查看报告</button>' +
                '<button class="btn btn-sm btn-ghost" onclick="window.open(\'' + pr.url + '\', \'_blank\')">查看 PR</button>';
        } else if (isReviewing) {
            statusHtml = '<span class="badge bg-success/10 border-success/20 text-success ml-2">Review 中</span>';
            const logFile = reviewTasks.find(t => t.targetId === pr.number)?.logFile?.replace(/^logs\//, '') || '';
            buttonsHtml = '<button class="btn btn-sm btn-destructive" onclick="stopReviewTask(event, ' + pr.number + ')">停止 Review</button>' +
                '<button class="btn btn-sm btn-primary" onclick="openLogTerminalByFile(\'' + logFile + '\')">查看日志</button>' +
                '<button class="btn btn-sm btn-ghost" onclick="window.open(\'' + pr.url + '\', \'_blank\')">查看 PR</button>';
        } else if (reviewTask) {
            if (reviewTask.status === 'stopped' || reviewTask.status === 'interrupted') {
                const statusText = reviewTask.status === 'stopped' ? '已停止' : '已中断';
                statusHtml = '<span class="badge bg-warning/10 border-warning/20 text-warning ml-2">' + statusText + '</span>';
                extraClass = ' border-l-[3px] border-l-warning';
                buttonsHtml = '<button class="btn btn-sm btn-primary" onclick="agentReviewPR(event, ' + pr.number + ', \'' + pr.branch + '\')">重新 Review</button>' +
                    '<button class="btn btn-sm btn-ghost" onclick="window.open(\'' + pr.url + '\', \'_blank\')">查看 PR</button>';
            } else {
                const isPassed = reviewTask.result === 'passed';
                const statusClass = isPassed ? 'bg-primary/10 border-primary/20 text-primary' : 'bg-destructive/10 border-destructive/20 text-destructive';
                const statusText = isPassed ? '通过' : '需修改';
                let docSyncHint = '';
                const report = reviewTask.report || '';
                if (report.includes('文档同步') && (report.includes('未通过') || report.includes('文档缺失') || report.includes('缺少文档'))) {
                    docSyncHint = ' <span title="文档同步未通过" class="cursor-help text-warning">!</span>';
                }
                statusHtml = '<span class="badge ' + statusClass + ' ml-2">' + statusText + '</span>' + docSyncHint;

                if (isPassed) {
                    buttonsHtml = '<button class="btn btn-sm btn-primary" onclick="viewReviewReport(' + pr.number + ')">查看报告</button>' +
                        '<button class="btn btn-sm btn-primary" onclick="runDocGardener(event, ' + pr.number + ', \'' + (pr.baseRefName || '') + '\')">文档检查</button>' +
                        '<button class="btn btn-sm btn-success" onclick="mergePR(event, ' + pr.number + ')">合并 PR</button>' +
                        '<button class="btn btn-sm btn-ghost" onclick="window.open(\'' + pr.url + '\', \'_blank\')">查看 PR</button>';
                } else {
                    buttonsHtml = '<button class="btn btn-sm btn-primary" onclick="viewReviewReport(' + pr.number + ')">查看报告</button>' +
                        '<button class="btn btn-sm btn-warning" onclick="reworkPR(event, ' + pr.number + ')">标记返工</button>' +
                        '<button class="btn btn-sm btn-primary" onclick="agentReviewPR(event, ' + pr.number + ', \'' + pr.branch + '\')">重新 Review</button>';
                }
            }
        } else {
            buttonsHtml = '<button class="btn btn-sm btn-success" onclick="agentReviewPR(event, ' + pr.number + ', \'' + pr.branch + '\')">Agent Review</button>' +
                '<button class="btn btn-sm btn-primary" onclick="mergePR(event, ' + pr.number + ')">直接通过</button>' +
                '<button class="btn btn-sm btn-ghost" onclick="window.open(\'' + pr.url + '\', \'_blank\')">查看 PR</button>';
        }

        let designBtnHtml = '';
        if (designIssueNumber && designAssetsCache[designIssueNumber]) {
            designBtnHtml = '<button class="btn btn-sm btn-primary design-view-btn" onclick="openDesignViewer(' + designIssueNumber + ')">查看设计</button>';
        }

        return '<div class="' + CARD + extraClass + '" data-pr="' + pr.number + '"' + (designIssueNumber ? ' data-design-issue="' + designIssueNumber + '"' : '') + '>' +
            '<div class="flex items-start justify-between gap-3 mb-2">' +
                '<a href="' + pr.url + '" class="text-sm font-semibold text-foreground hover:text-primary transition-colors" target="_blank">' +
                    '#' + pr.number + ': ' + escapeHtml(pr.title) +
                '</a>' +
                statusHtml +
            '</div>' +
            '<div class="mb-3 text-xs text-muted-foreground">' + time + ' · 分支: ' + pr.branch + '</div>' +
            '<div class="flex flex-wrap gap-2">' + designBtnHtml + buttonsHtml + '</div>' +
        '</div>';
    }).join('');
}

/* ─── Merged PRs ─── */
export function renderMergedPRs(prs) {
    const container = document.getElementById('merged-prs');

    if (prs.length === 0) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无最近合并的 PR</div>';
        return;
    }

    container.innerHTML = prs.map(pr => {
        const time = pr.createdAt ? new Date(pr.createdAt).toLocaleString() : '未知';
        return '<div class="' + CARD + ' opacity-75" data-pr="' + pr.number + '">' +
            '<div class="flex items-start justify-between gap-3 mb-2">' +
                '<a href="' + pr.url + '" class="text-sm font-semibold text-foreground hover:text-primary transition-colors" target="_blank">' +
                    '#' + pr.number + ': ' + escapeHtml(pr.title) +
                '</a>' +
                '<span class="badge bg-primary/10 border-primary/20 text-primary">已合并</span>' +
            '</div>' +
            '<div class="mb-3 text-xs text-muted-foreground">' + time + ' · 分支: ' + pr.branch + '</div>' +
            '<div class="flex flex-wrap gap-2">' +
                '<button class="btn btn-sm btn-primary" onclick="runMergedDocGardener(event, ' + pr.number + ')">文档检查</button>' +
                '<button class="btn btn-sm btn-ghost" onclick="window.open(\'' + pr.url + '\', \'_blank\')">查看 PR</button>' +
            '</div>' +
        '</div>';
    }).join('');
}

/* ─── Logs ─── */
export function renderLogs(logs, containerId) {
    const container = document.getElementById(containerId);
    if (!logs) logs = [];

    if (logs.length === 0) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无日志</div>';
        return;
    }

    container.innerHTML = logs.map(log => {
        const time = log.modifiedAt ? new Date(log.modifiedAt).toLocaleString() : '未知';
        const size = formatBytes(log.size);
        let title = 'File';
        if (log.filename.startsWith('dev-issue-')) {
            title = 'Issue #' + log.issueNumber;
        } else if (log.filename.startsWith('design-issue-')) {
            title = 'Issue #' + log.issueNumber + ' Design';
        } else if (log.filename.startsWith('review-pr-')) {
            title = 'PR #' + log.issueNumber + ' Review';
        } else if (log.filename.startsWith('review-report-')) {
            title = 'PR #' + log.issueNumber + ' Report';
        } else if (log.filename.startsWith('review-state-')) {
            title = 'PR #' + log.issueNumber + ' State';
        } else if (log.filename.startsWith('pm-issue-')) {
            title = 'Issue #' + log.issueNumber + ' PRD';
        } else if (log.filename.startsWith('doc-gardener-pr-')) {
            title = 'PR #' + log.issueNumber + ' Doc Gardener';
        } else if (log.filename.startsWith('e2e-issue-')) {
            title = 'Issue #' + log.issueNumber + ' E2E';
        } else if (log.filename.startsWith('rework-pr-')) {
            title = 'PR #' + log.issueNumber + ' Rework';
        }
        return '<div class="' + CARD + ' p-3">' +
            '<div class="flex items-center justify-between gap-3">' +
                '<div>' +
                    '<div class="text-sm font-medium text-foreground">' + title + '</div>' +
                    '<div class="text-xs text-muted-foreground">' + time + ' · ' + size + '</div>' +
                '</div>' +
                '<button class="btn btn-sm btn-primary" onclick="openLogTerminalByFile(\'' + log.filename + '\')">查看</button>' +
            '</div>' +
        '</div>';
    }).join('');
}


/* ─── Prompt Evolution Stats ─── */
export function renderPromptEvolutionStats(data) {
    const container = document.getElementById('prompt-evolution-stats');
    if (!container) return;
    if (!data) {
        container.innerHTML = '<div class="' + EMPTY + '">暂无统计数据</div>';
        return;
    }

    const state = data.state || {};
    const report = data.report || {};
    const agentStats = report.agentStats || {};
    const issues = report.topIssues || [];
    const runningTask = data.runningTask;

    let html = '';

    // 状态概览
    const lastTime = state.lastEvolutionTime ? new Date(state.lastEvolutionTime).toLocaleString() : '从未';
    const nextTrigger = state.taskCountSinceLast >= 5 ? '满足条件，下次任务完成后触发' : `还需 ${5 - (state.taskCountSinceLast || 0)} 个任务`;
    html += '<div class="mb-3 rounded-lg bg-secondary/50 p-3 space-y-1.5">' +
        '<div class="flex items-center justify-between"><span class="text-xs text-muted-foreground">上次分析</span><span class="text-xs font-medium">' + lastTime + '</span></div>' +
        '<div class="flex items-center justify-between"><span class="text-xs text-muted-foreground">任务计数</span><span class="text-xs font-medium">' + (state.taskCountSinceLast || 0) + ' / 5</span></div>' +
        '<div class="flex items-center justify-between"><span class="text-xs text-muted-foreground">触发条件</span><span class="text-xs font-medium text-primary">' + nextTrigger + '</span></div>' +
        (runningTask ? '<div class="flex items-center justify-between"><span class="text-xs text-muted-foreground">演进任务</span><span class="text-xs font-medium text-success">运行中 #' + runningTask.targetID + '</span></div>' : '') +
        '</div>';

    // Agent 统计表格
    if (Object.keys(agentStats).length > 0) {
        html += '<div class="mb-3 overflow-hidden rounded-lg border border-border">' +
            '<table class="w-full text-xs">' +
            '<thead><tr class="bg-secondary/70"><th class="px-2 py-1.5 text-left font-medium">Agent</th><th class="px-2 py-1.5 text-right font-medium">任务</th><th class="px-2 py-1.5 text-right font-medium">步数</th><th class="px-2 py-1.5 text-right font-medium">成功率</th></tr></thead>' +
            '<tbody>';
        for (const [agent, stat] of Object.entries(agentStats)) {
            const successRate = (stat.successRate * 100).toFixed(0);
            const successColor = stat.successRate >= 0.7 ? 'text-success' : stat.successRate >= 0.4 ? 'text-warning' : 'text-destructive';
            html += '<tr class="border-t border-border/50 hover:bg-secondary/30">' +
                '<td class="px-2 py-1.5 font-medium">' + escapeHtml(agent) + '</td>' +
                '<td class="px-2 py-1.5 text-right">' + stat.count + '</td>' +
                '<td class="px-2 py-1.5 text-right">' + (stat.avgSteps ? stat.avgSteps.toFixed(1) : '-') + '</td>' +
                '<td class="px-2 py-1.5 text-right ' + successColor + '">' + successRate + '%</td>' +
                '</tr>';
        }
        html += '</tbody></table></div>';
    }

    // Top Issues（只显示前 5 个）
    if (issues.length > 0) {
        html += '<div class="space-y-1.5">' +
            '<div class="text-xs font-semibold text-muted-foreground uppercase tracking-wider">热点问题</div>';
        const top5 = issues.slice(0, 5);
        for (const issue of top5) {
            html += '<div class="rounded-md bg-destructive/5 border border-destructive/10 px-2.5 py-1.5">' +
                '<div class="text-xs font-medium text-destructive">' + escapeHtml(issue.pattern) + '</div>' +
                '<div class="text-[0.7rem] text-muted-foreground mt-0.5">' + escapeHtml(issue.description) + '</div>' +
                '</div>';
        }
        if (issues.length > 5) {
            html += '<div class="text-center text-[0.7rem] text-muted-foreground">还有 ' + (issues.length - 5) + ' 个问题...</div>';
        }
        html += '</div>';
    }

    // 历史报告
    const history = data.historyReports || [];
    if (history.length > 0) {
        html += '<div class="mt-3 space-y-1.5">' +
            '<div class="text-xs font-semibold text-muted-foreground uppercase tracking-wider">历史报告</div>';
        for (const hr of history.slice(0, 5)) {
            const reportUrl = hr.url || '#';
            html += '<div class="flex items-center justify-between text-xs">' +
                '<a href="' + reportUrl + '" target="_blank" class="text-primary hover:underline truncate max-w-[180px]" title="查看 JSON 报告">' + escapeHtml(hr.filename) + '</a>' +
                '<span class="text-muted-foreground">' + new Date(hr.createdAt).toLocaleDateString() + '</span>' +
                '</div>';
        }
        html += '</div>';
    }

    container.innerHTML = html;
}
