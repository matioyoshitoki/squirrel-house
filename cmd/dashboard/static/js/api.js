async function post(path, body) {
    const res = await fetch(path, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
    });
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

async function get(path) {
    const res = await fetch(path);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export const api = {
    projects: () => get('/api/projects'),
    switchProject: (name) => post('/api/switch-project', { projectName: name }),

    issues: () => get('/api/issues'),
    runningTasks: () => get('/api/running-tasks'),
    logs: () => get('/api/logs'),
    pullRequests: () => get('/api/pull-requests'),
    mergedPRs: () => get('/api/pull-requests?state=merged'),
    reviewTasks: () => get('/api/review-tasks'),
    docGardenerTasks: () => get('/api/doc-gardener-tasks'),
    designAssets: (issueNumber) => get(`/api/design-assets?issueNumber=${issueNumber}`),
    buildDesignPreview: (issueNumber) => post('/api/build-design-preview', { issueNumber }),
    submitDesignFeedback: (issueNumber, feedback) => post('/api/design-feedback', { issueNumber, feedback }),
    reviewReport: (prNumber) => get(`/api/review-report?pr=${prNumber}`),

    reviewPR: (prNumber, branch) => post('/api/review-pr', { prNumber, branch }),
    docGardener: (prNumber, baseBranch) => post('/api/doc-gardener', { prNumber, baseBranch }),
    mergedDocGardener: (prNumber) => post('/api/merged-doc-gardener', { prNumber }),
    mergePR: (prNumber) => post('/api/merge-pr', { prNumber }),
    startDev: (issueNumber, issueTitle, issueType) => post('/api/start-dev', { issueNumber, issueTitle, issueType }),
    runE2E: (issueNumber, branch, taskKind) => post('/api/run-e2e', { issueNumber, branch, taskKind }),
    startDesign: (issueNumber, issueTitle) => post('/api/start-design', { issueNumber, issueTitle }),
    startPM: (title, description) => post('/api/start-pm', { title, description }),
    resumeDev: (issueNumber, issueTitle) => post('/api/resume-dev', { issueNumber, issueTitle }),
    stopTask: (issueNumber, taskType) => post('/api/stop-task', { issueNumber, taskType }),
    stopReviewTask: (prNumber) => post('/api/stop-task', { prNumber, taskType: 'review' }),
    reworkPR: (prNumber) => post('/api/rework-pr', { prNumber }),
    promptEvolutionStats: () => get('/api/prompt-evolution-stats'),
};
