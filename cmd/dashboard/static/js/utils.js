export function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

export function formatDuration(ms) {
    const seconds = Math.floor(ms / 1000);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);
    if (hours > 0) return `${hours}h ${minutes % 60}m`;
    if (minutes > 0) return `${minutes}m ${seconds % 60}s`;
    return `${seconds}s`;
}

export function formatBytes(bytes) {
    if (bytes < 1024) return bytes + ' B';
    if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB';
    return (bytes / (1024 * 1024)).toFixed(1) + ' MB';
}

export function getIssueType(labels, title, body) {
    if (labels && labels.some(l => l.name === 'bug')) return 'bug';
    if (labels && labels.some(l => l.name === 'feature')) return 'feature';
    if (labels && labels.some(l => l.name === 'performance')) return 'performance';
    if (labels && labels.some(l => l.name === 'tech-debt')) return 'tech-debt';
    if (labels && labels.some(l => l.name === 'ui-feedback')) return 'ui-feedback';
    const text = (title + ' ' + (body || '')).toLowerCase();
    if (text.includes('bug') || text.includes('fix')) return 'bug';
    if (text.includes('performance') || text.includes('slow')) return 'performance';
    if (text.includes('tech debt') || text.includes('refactor')) return 'tech-debt';
    return 'feature';
}

export function getIssueTypeDisplay(type) {
    const map = {
        bug:         { text: '修复 Bug',     btnClass: 'btn-destructive' },
        feature:     { text: '开发 Feature', btnClass: 'btn-info' },
        performance: { text: '性能优化',     btnClass: 'btn-warning' },
        'tech-debt': { text: '技术债',       btnClass: 'btn-amber' },
        'ui-feedback': { text: 'UI 反馈',    btnClass: 'btn-cyan' },
    };
    return map[type] || map.feature;
}

export function updateRunningTimes() {
    document.querySelectorAll('.running-time').forEach(el => {
        const start = Number(el.dataset.start);
        if (!start || isNaN(start)) return;
        const label = el.dataset.label || '';
        el.textContent = (label ? label + ' · ' : '') + formatDuration(Date.now() - start);
    });
}
