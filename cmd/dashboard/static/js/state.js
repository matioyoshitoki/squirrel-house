/**
 * Tiny reactive store with selective subscriptions.
 */
export function createStore(initialState) {
    let state = { ...initialState };
    const listeners = new Map(); // key -> Set<fn>

    function notify(changedKeys) {
        const toCall = new Set();
        for (const key of changedKeys) {
            listeners.get(key)?.forEach(fn => toCall.add(fn));
        }
        // wildcard listeners
        listeners.get('*')?.forEach(fn => toCall.add(fn));
        toCall.forEach(fn => fn(state));
    }

    return {
        get(key) {
            return key !== undefined ? state[key] : state;
        },

        set(patch) {
            const changed = [];
            for (const [key, value] of Object.entries(patch)) {
                if (state[key] !== value) {
                    state[key] = value;
                    changed.push(key);
                }
            }
            if (changed.length) notify(changed);
        },

        subscribe(keys, fn) {
            const list = Array.isArray(keys) ? keys : [keys];
            for (const key of list) {
                if (!listeners.has(key)) listeners.set(key, new Set());
                listeners.get(key).add(fn);
            }
            // Return unsubscribe function
            return () => {
                for (const key of list) {
                    listeners.get(key)?.delete(fn);
                }
            };
        },
    };
}

export const store = createStore({
    // Projects
    currentProject: null,

    // Stage filter
    currentStage: 'all',

    // Tasks
    runningTasksData: [],
    reviewTasksData: [],
    docGardenerTasksData: [],

    // PRs
    pullRequestsData: [],
    mergedPRsData: [],

    // Issues
    issuesData: [],

    // Logs
    allLogsData: [],

    // Prompt Evolution
    promptEvolutionStatsData: null,

    // Design assets cache: { issueNumber: assets[] | null }
    designAssetsCache: {},

    // Global loading overlay ref count
    globalLoadingCount: 0,

    // Terminal state
    eventSource: null,
    lineCount: 0,
    isAutoScroll: true,
    currentLogFile: null,
    renderMode: 'rendered',
    rawLogLines: [],
    logBuffer: [],
    logFlushPending: false,
    parsedBlockPending: null,
});

// Convenience accessors
export const $ = (sel) => document.querySelector(sel);
export const $$ = (sel) => document.querySelectorAll(sel);
