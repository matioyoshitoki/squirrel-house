import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import {
  renderRunningTasks,
  renderIssues,
  renderPullRequests,
} from './renderers.js';

function createContainer(id) {
  const el = document.createElement('div');
  el.id = id;
  document.body.appendChild(el);
  return el;
}

function cleanupContainers() {
  document.body.innerHTML = '';
}

describe('renderRunningTasks', () => {
  beforeEach(() => {
    cleanupContainers();
    createContainer('running-tasks');
  });

  afterEach(() => {
    cleanupContainers();
  });

  it('should show "查看设计" button when design task has assets', () => {
    const tasks = [{
      id: 'design-6',
      type: 'design',
      status: 'running',
      targetId: 6,
      targetTitle: 'Design Issue',
      logFile: 'logs/design-issue-6.log',
      startTime: new Date().toISOString(),
    }];
    const designAssetsCache = {
      6: [{ name: 'preview.html', url: '/preview.html', type: 'html' }],
    };

    renderRunningTasks(tasks, { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('running-tasks');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).not.toBeNull();
    expect(btn.textContent).toBe('查看设计');
  });

  it('should not show "查看设计" button when design task has no assets', () => {
    const tasks = [{
      id: 'design-6',
      type: 'design',
      status: 'running',
      targetId: 6,
      targetTitle: 'Design Issue',
      logFile: 'logs/design-issue-6.log',
      startTime: new Date().toISOString(),
    }];
    const designAssetsCache = {};

    renderRunningTasks(tasks, { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('running-tasks');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).toBeNull();
  });

  it('should not show completed design tasks', () => {
    const tasks = [{
      id: 'design-6',
      type: 'design',
      status: 'success',
      targetId: 6,
      targetTitle: 'Design Issue',
      logFile: 'logs/design-issue-6.log',
      startTime: new Date().toISOString(),
    }];
    const designAssetsCache = {
      6: [{ name: 'preview.html', url: '/preview.html', type: 'html' }],
    };

    renderRunningTasks(tasks, { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('running-tasks');
    expect(container.textContent).toContain('暂无运行中的任务');
  });
});

describe('renderIssues', () => {
  beforeEach(() => {
    cleanupContainers();
    createContainer('plan-issues');
    createContainer('prd-issues');
    createContainer('triage-issues');
  });

  afterEach(() => {
    cleanupContainers();
  });

  it('should show "查看设计" button when issue has design assets', () => {
    const issues = [{
      number: 6,
      title: 'Test Issue',
      labels: [{ name: 'plan-ready' }],
      classification: 'full-stack',
      updatedAt: new Date().toISOString(),
      body: '',
    }];
    const designAssetsCache = {
      6: [{ name: 'preview.html', url: '/preview.html', type: 'html' }],
    };

    renderIssues(issues, [], { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('plan-issues');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).not.toBeNull();
    expect(btn.textContent).toBe('查看设计');
  });

  it('should not show "查看设计" button when designAssetsCache is null', () => {
    const issues = [{
      number: 6,
      title: 'Test Issue',
      labels: [{ name: 'plan-ready' }],
      classification: 'full-stack',
      updatedAt: new Date().toISOString(),
      body: '',
    }];
    const designAssetsCache = { 6: null };

    renderIssues(issues, [], { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('plan-issues');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).toBeNull();
  });

  it('should not show "查看设计" button when designAssetsCache is undefined', () => {
    const issues = [{
      number: 6,
      title: 'Test Issue',
      labels: [{ name: 'plan-ready' }],
      classification: 'full-stack',
      updatedAt: new Date().toISOString(),
      body: '',
    }];
    const designAssetsCache = {};

    renderIssues(issues, [], { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('plan-issues');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).toBeNull();
  });

  it('should show "修改设计" button when issue already has assets', () => {
    const issues = [{
      number: 6,
      title: 'Test Issue',
      labels: [{ name: 'plan-ready' }],
      classification: 'full-stack',
      updatedAt: new Date().toISOString(),
      body: '',
    }];
    const designAssetsCache = {
      6: [{ name: 'preview.html', url: '/preview.html', type: 'html' }],
    };

    renderIssues(issues, [], { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('plan-issues');
    const btn = Array.from(container.querySelectorAll('button')).find(
      b => b.textContent.includes('修改设计')
    );
    expect(btn).not.toBeNull();
  });

  it('should show "开始设计" button when issue has no assets', () => {
    const issues = [{
      number: 6,
      title: 'Test Issue',
      labels: [{ name: 'plan-ready' }],
      classification: 'full-stack',
      updatedAt: new Date().toISOString(),
      body: '',
    }];
    const designAssetsCache = {};

    renderIssues(issues, [], { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('plan-issues');
    const btn = Array.from(container.querySelectorAll('button')).find(
      b => b.textContent.includes('开始设计')
    );
    expect(btn).not.toBeNull();
  });

  it('should not show design buttons for backend-only issues', () => {
    const issues = [{
      number: 6,
      title: 'Backend API',
      labels: [{ name: 'plan-ready' }],
      classification: 'backend-only',
      updatedAt: new Date().toISOString(),
      body: '',
    }];
    const designAssetsCache = {};

    renderIssues(issues, [], { repo: 'test/repo' }, designAssetsCache);

    const container = document.getElementById('plan-issues');
    const startBtn = Array.from(container.querySelectorAll('button')).find(
      b => b.textContent.includes('开始设计')
    );
    const viewBtn = container.querySelector('.design-view-btn');
    expect(startBtn).toBeUndefined();
    expect(viewBtn).toBeNull();
  });

  it('should show "运行中" button when any task for the issue is running', () => {
    const issues = [{
      number: 6,
      title: 'Test Issue',
      labels: [{ name: 'plan-ready' }],
      classification: 'full-stack',
      updatedAt: new Date().toISOString(),
      body: '',
    }];
    const runningTasks = [{
      type: 'design',
      status: 'running',
      targetId: 6,
    }];

    renderIssues(issues, runningTasks, { repo: 'test/repo' }, {});

    const container = document.getElementById('plan-issues');
    const runningBtn = Array.from(container.querySelectorAll('button')).find(
      b => b.disabled && b.textContent === '运行中'
    );
    expect(runningBtn).toBeDefined();
  });
});

describe('renderPullRequests', () => {
  beforeEach(() => {
    cleanupContainers();
    createContainer('pending-prs');
  });

  afterEach(() => {
    cleanupContainers();
  });

  it('should show "查看设计" button for design PRs with cached assets', () => {
    const prs = [{
      number: 42,
      title: 'design: assets for issue #6',
      createdAt: new Date().toISOString(),
      branch: 'feature/design-6',
      url: 'https://github.com/test/repo/pull/42',
    }];
    const designAssetsCache = {
      6: [{ name: 'preview.html', url: '/preview.html', type: 'html' }],
    };

    renderPullRequests(prs, [], [], designAssetsCache);

    const container = document.getElementById('pending-prs');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).not.toBeNull();
    expect(btn.textContent).toBe('查看设计');
  });

  it('should not show "查看设计" button when design assets are not cached', () => {
    const prs = [{
      number: 42,
      title: 'design: assets for issue #6',
      createdAt: new Date().toISOString(),
      branch: 'feature/design-6',
      url: 'https://github.com/test/repo/pull/42',
    }];
    const designAssetsCache = {};

    renderPullRequests(prs, [], [], designAssetsCache);

    const container = document.getElementById('pending-prs');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).toBeNull();
  });

  it('should not show "查看设计" button for non-design PRs', () => {
    const prs = [{
      number: 42,
      title: 'feat: add login page',
      createdAt: new Date().toISOString(),
      branch: 'feature/login',
      url: 'https://github.com/test/repo/pull/42',
    }];
    const designAssetsCache = {};

    renderPullRequests(prs, [], [], designAssetsCache);

    const container = document.getElementById('pending-prs');
    const btn = container.querySelector('.design-view-btn');
    expect(btn).toBeNull();
  });
});
