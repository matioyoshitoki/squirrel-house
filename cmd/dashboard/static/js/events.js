// events.js — SSE client for real-time state updates
// Replaces aggressive polling with server-push + fallback polling.

let source = null;
let reconnectDelay = 1000;
const MAX_RECONNECT_DELAY = 30000;
let reconnectTimer = null;
let connected = false;
const connectionCbs = [];
const eventCbs = {};

function setConnected(val) {
    if (connected === val) return;
    connected = val;
    connectionCbs.forEach(fn => fn(val));
}

function scheduleReconnect() {
    if (reconnectTimer) return;
    reconnectTimer = setTimeout(() => {
        reconnectTimer = null;
        startEventStream();
    }, reconnectDelay);
    reconnectDelay = Math.min(reconnectDelay * 2, MAX_RECONNECT_DELAY);
}

export function isConnected() {
    return connected;
}

export function onConnectionChange(cb) {
    connectionCbs.push(cb);
    return () => {
        const idx = connectionCbs.indexOf(cb);
        if (idx !== -1) connectionCbs.splice(idx, 1);
    };
}

export function onEvent(type, cb) {
    if (!eventCbs[type]) eventCbs[type] = [];
    eventCbs[type].push(cb);
    return () => {
        const arr = eventCbs[type];
        if (!arr) return;
        const idx = arr.indexOf(cb);
        if (idx !== -1) arr.splice(idx, 1);
    };
}

function dispatchEvent(type, data) {
    const arr = eventCbs[type];
    if (arr) arr.forEach(fn => {
        try { fn(data); } catch (e) { console.error('[SSE] handler error:', e); }
    });
    // Also dispatch to wildcard listeners
    const wild = eventCbs['*'];
    if (wild) wild.forEach(fn => {
        try { fn(type, data); } catch (e) { console.error('[SSE] wildcard handler error:', e); }
    });
}

export function startEventStream() {
    if (source) return;

    try {
        source = new EventSource('/api/events');
    } catch (e) {
        console.error('[SSE] failed to create EventSource:', e);
        scheduleReconnect();
        return;
    }

    source.onopen = () => {
        reconnectDelay = 1000;
        setConnected(true);
        console.log('[SSE] connected');
    };

    source.onerror = () => {
        if (source) {
            source.close();
            source = null;
        }
        setConnected(false);
        console.log('[SSE] disconnected, will retry...');
        scheduleReconnect();
    };

    source.addEventListener('connected', (e) => {
        try {
            const data = JSON.parse(e.data);
            dispatchEvent('connected', data);
        } catch (_) {}
    });

    source.addEventListener('task:changed', (e) => {
        try {
            const data = JSON.parse(e.data);
            dispatchEvent('task:changed', data);
        } catch (_) {}
    });

    source.addEventListener('pr:merged', (e) => {
        try {
            const data = JSON.parse(e.data);
            dispatchEvent('pr:merged', data);
        } catch (_) {}
    });

    source.addEventListener('project:switched', (e) => {
        try {
            const data = JSON.parse(e.data);
            dispatchEvent('project:switched', data);
        } catch (_) {}
    });
}

export function closeEventStream() {
    if (reconnectTimer) {
        clearTimeout(reconnectTimer);
        reconnectTimer = null;
    }
    if (source) {
        source.close();
        source = null;
    }
    setConnected(false);
}
