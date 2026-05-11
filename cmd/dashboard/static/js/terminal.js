import { store } from './state.js';
import { escapeHtml } from './utils.js';

// Terminal-local state (high-frequency, not in global store)
let eventSource = null;
let lineCount = 0;
let isAutoScroll = true;
let currentLogFile = null;
let renderMode = 'rendered';
let rawLogLines = [];
let logBuffer = [];
let logFlushPending = false;
let parsedBlockPending = null;

const terminalOverlay = document.getElementById('terminal-overlay');
const terminalBody = document.getElementById('terminal-body');
const lineCountEl = document.getElementById('line-count');

// Scroll helpers
let userScrolledUp = false;

terminalBody?.addEventListener('scroll', () => {
    const threshold = 50;
    const atBottom = terminalBody.scrollHeight - terminalBody.scrollTop - terminalBody.clientHeight < threshold;
    userScrolledUp = !atBottom;
    const btn = document.getElementById('scroll-bottom-btn');
    if (btn) btn.classList.toggle('hidden', atBottom);
    isAutoScroll = atBottom;
});

document.getElementById('scroll-bottom-btn')?.addEventListener('click', () => {
    isAutoScroll = true;
    scrollToBottom();
});

function scrollToBottom() {
    if (terminalBody) {
        terminalBody.scrollTop = terminalBody.scrollHeight;
    }
}

// Main entry
export function openLogTerminalByFile(filename) {
    currentLogFile = filename;
    const filenameEl = document.getElementById('terminal-filename');
    if (filenameEl) filenameEl.textContent = 'tail -f ' + filename;
    terminalOverlay?.classList.remove('hidden');
    terminalOverlay?.classList.add('flex');

    if (terminalBody) terminalBody.innerHTML = '';
    lineCount = 0;
    isAutoScroll = true;
    rawLogLines = [];
    parsedBlockPending = null;
    logBuffer = [];

    if (eventSource) {
        eventSource.close();
        eventSource = null;
    }

    const es = new EventSource('/api/log-tail?file=' + encodeURIComponent(filename));
    eventSource = es;

    es.onmessage = (e) => {
        logBuffer.push(e.data);
        if (!logFlushPending) {
            logFlushPending = true;
            requestAnimationFrame(() => {
                flushLogBuffer();
                logFlushPending = false;
            });
        }
    };

    es.onerror = () => {
        console.log('SSE error, reconnecting...');
        es.close();
        eventSource = null;
        setTimeout(() => {
            if (terminalOverlay && terminalOverlay.classList.contains('flex')) {
                openLogTerminalByFile(filename);
            }
        }, 3000);
    };

    document.addEventListener('keydown', handleEscKey);
}

export function closeTerminal() {
    terminalOverlay?.classList.add('hidden');
    terminalOverlay?.classList.remove('flex');
    if (eventSource) {
        eventSource.close();
        eventSource = null;
    }
    logBuffer = [];
    logFlushPending = false;
    document.removeEventListener('keydown', handleEscKey);
}

export function clearTerminal() {
    if (terminalBody) terminalBody.innerHTML = '';
    lineCount = 0;
    logBuffer = [];
    rawLogLines = [];
    parsedBlockPending = null;
    if (lineCountEl) lineCountEl.textContent = '0 lines';
}

export function toggleRenderMode() {
    const newMode = renderMode === 'rendered' ? 'raw' : 'rendered';
    renderMode = newMode;
    const btn = document.getElementById('render-toggle-btn');
    if (btn) {
        btn.textContent = newMode === 'rendered' ? '渲染' : '原始';
        btn.classList.toggle('btn-primary', newMode === 'rendered');
        btn.classList.toggle('btn-outline', newMode !== 'rendered');
    }
    if (terminalBody) {
        terminalBody.innerHTML = '';
        lineCount = 0;
        parsedBlockPending = null;
        if (rawLogLines.length > 0) {
            const toRender = rawLogLines.slice();
            rawLogLines = [];
            logBuffer = toRender;
            flushLogBuffer();
        }
    }
}

export function handleEscKey(e) {
    if (e.key === 'Escape') {
        closeTerminal();
    }
}

// Raw line coloring
function createLogLine(text) {
    const lineDiv = document.createElement('div');
    lineDiv.className = 'log-line py-0.5 whitespace-pre-wrap break-all hover:bg-secondary/40';

    let contentClass = '';
    const lower = text.toLowerCase();
    if (lower.includes('error') || lower.includes('❌') || lower.includes('fail')) {
        contentClass = 'text-destructive';
    } else if (lower.includes('success') || lower.includes('✅') || lower.includes('done')) {
        contentClass = 'text-success';
    } else if (lower.includes('warn') || lower.includes('⚠️')) {
        contentClass = 'text-warning';
    } else if (lower.includes('toolcall') || lower.includes('toolresult')) {
        contentClass = 'text-[#a78bfa]';
    } else if (lower.includes('thinkpart')) {
        contentClass = 'text-muted-foreground italic';
    } else if (text.startsWith('>') || text.startsWith('$')) {
        contentClass = 'text-primary font-medium';
    }

    const escaped = text.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    lineCount++;
    lineDiv.innerHTML =
        '<span class="inline-block w-14 select-none pr-3 text-right text-muted-foreground/70">' + lineCount + '</span>' +
        '<span class="text-foreground ' + contentClass + '">' + escaped + '</span>';

    return lineDiv;
}

// Kimi block parsing
function parseKimiBlocks(rawTexts) {
    const allLines = [];
    for (const text of rawTexts) {
        if (text) allLines.push(...text.split('\n'));
    }
    const blocks = [];
    let current = parsedBlockPending;
    const topLevelPattern = /^(TurnBegin|TurnEnd|StepBegin|ThinkPart|ToolCall|ToolResult|ToolCallPart|StatusUpdate|TextPart|ToolResultPart)\(/;

    for (const line of allLines) {
        const trimmed = line.trimStart();
        const match = trimmed.match(topLevelPattern);
        if (match) {
            if (current) blocks.push(current);
            current = { type: match[1], lines: [line] };
        } else {
            if (!current) current = { type: 'plain', lines: [] };
            current.lines.push(line);
        }
    }
    parsedBlockPending = current;
    return blocks;
}

function extractValue(text, key) {
    const regex = new RegExp(key + "='([^']*?)'(?:,|\\))");
    const m = text.match(regex);
    if (m) return m[1];
    const regex2 = new RegExp(key + '="([^"]*?)"(?:,|\\))');
    const m2 = text.match(regex2);
    if (m2) return m2[1];
    return '';
}

function extractField(lines, fieldName) {
    const text = lines.join('\n');
    const regex = new RegExp(fieldName + "=['\"]([\\s\\S]*?)['\"](?:,\\s*\\n|\\))");
    const m = text.match(regex);
    if (m) {
        return m[1]
            .replace(/\\n/g, '\n')
            .replace(/\\t/g, '\t')
            .replace(/\\r/g, '\r');
    }
    const regex2 = new RegExp(fieldName + '=([^,\\n\\)]+)');
    const m2 = text.match(regex2);
    if (m2) return m2[1].trim();
    return '';
}

function unwrapHardBreaks(text) {
    return text.replace(/(\w)\n(?=\w|-\w)/g, '$1');
}

function createRenderedBlock(block, isPending) {
    const div = document.createElement('div');
    div.className = 'rendered-block my-1.5';
    if (isPending) div.setAttribute('data-pending', 'true');
    const rawText = block.lines.join('\n');

    switch (block.type.toLowerCase()) {
        case 'turnbegin': {
            div.className += ' my-2 border-l-[3px] border-l-primary pl-3';
            const hasPrompt = rawText.includes('user_input=');
            div.innerHTML = '<div class="font-semibold text-primary cursor-pointer select-none hover:opacity-80" onclick="this.nextElementSibling.classList.toggle(\'hidden\')">Turn 开始 ' + (hasPrompt ? '(点击展开 prompt)' : '') + '</div>' +
                '<div class="hidden"><pre class="mt-2 text-xs whitespace-pre-wrap text-muted-foreground">' + escapeHtml(rawText.substring(0, 500)) + (rawText.length > 500 ? '... (' + rawText.length + ' chars)' : '') + '</pre></div>';
            break;
        }
        case 'turnend': {
            div.className += ' text-muted-foreground/70 text-xs py-1 border-t border-dashed border-border mt-2';
            div.innerHTML = 'Turn 结束';
            break;
        }
        case 'stepbegin': {
            div.className += ' my-2 rounded-md border border-primary/25 bg-primary/10 px-3 py-1.5 font-semibold text-primary';
            const n = rawText.match(/n=(\d+)/);
            div.innerHTML = 'Step ' + (n ? n[1] : '?');
            break;
        }
        case 'thinkpart': {
            div.className += ' my-1.5 rounded-r-md border-l-[3px] border-l-muted-foreground bg-muted-foreground/5 px-3 py-2';
            let think = extractField(block.lines, 'think');
            think = unwrapHardBreaks(think);
            const thinkHtml = typeof marked !== 'undefined' ? marked.parse(think) : escapeHtml(think);
            div.innerHTML = '<div class="text-xs font-semibold text-muted-foreground mb-1">思考</div>' +
                '<div class="kimi-markdown text-foreground/80 italic">' + thinkHtml + '</div>';
            break;
        }
        case 'toolcall': {
            div.className += ' my-1.5 rounded-md border border-[#a78bfa]/20 bg-[#a78bfa]/5 px-3 py-1.5';
            const name = extractField(block.lines, 'name');
            const args = extractField(block.lines, 'arguments').replace(/[\n\r]/g, '');
            let summary = name || 'Tool';
            let body = rawText;
            if (args && args.length > 0) {
                try {
                    const parsed = JSON.parse(args);
                    if (parsed.command) {
                        summary += ': ' + parsed.command.substring(0, 120);
                    } else if (parsed.path) {
                        summary += ': ' + parsed.path;
                    } else if (parsed.query) {
                        summary += ': ' + parsed.query.substring(0, 120);
                    } else {
                        summary += ': ' + JSON.stringify(parsed).substring(0, 120);
                    }
                    body = JSON.stringify(parsed, null, 2);
                } catch (e) {
                    summary += ': ' + args.substring(0, 120);
                }
            }
            div.innerHTML = '<div class="text-xs font-semibold text-[#a78bfa] cursor-pointer select-none" onclick="const b=this.nextElementSibling;b.classList.toggle(\'hidden\');b.classList.toggle(\'block\')">' + escapeHtml(summary) + '</div>' +
                '<div class="hidden text-foreground/80 mt-1"><pre class="m-0 text-xs leading-relaxed overflow-x-auto whitespace-pre">' + escapeHtml(body) + '</pre></div>';
            break;
        }
        case 'toolresult': {
            div.className += ' my-1.5 rounded-r-md border-l-[3px] border-l-success bg-success/5 px-3 py-1.5';
            let result = extractField(block.lines, 'result') || extractField(block.lines, 'output') || extractField(block.lines, 'content') || rawText;
            // kimi 输出可能自带行号前缀（如 "     1\t内容"），去掉以正确渲染 Markdown
            result = result.replace(/^\s+\d+\t/gm, '');
            const lines = result.split('\n');
            const summary = (lines[0] || 'Result').substring(0, 100);
            const isMd = /^(#{1,6}\s|^\|.+\||^[-*+]\s|^\d+\.\s|\*\*[^*]+\*\*|```|^\[.+\]\(.+\)|^>\s)/m.test(result);
            let bodyHtml;
            if (isMd && typeof marked !== 'undefined') {
                bodyHtml = marked.parse(result);
            } else {
                bodyHtml = '<pre class="m-0 overflow-x-auto whitespace-pre">' + escapeHtml(result) + '</pre>';
            }
            div.innerHTML = '<div class="text-xs font-semibold text-success cursor-pointer select-none" onclick="const b=this.nextElementSibling;b.classList.toggle(\'hidden\');b.classList.toggle(\'block\')">' + escapeHtml(summary) + '</div>' +
                '<div class="hidden text-foreground/80 text-xs font-mono leading-relaxed mt-1">' + bodyHtml + '</div>';
            break;
        }
        case 'toolcallpart': {
            div.className += ' my-1.5 rounded-md border border-[#a78bfa]/20 bg-[#a78bfa]/5 px-3 py-1.5';
            // kimi 输出可能把长参数截断为多行，先合并再解析
            const args = extractField(block.lines, 'arguments_part').replace(/[\n\r]/g, '');
            let summary = 'Tool片段';
            let body = rawText;
            if (args && args.length > 0) {
                try {
                    const parsed = JSON.parse(args);
                    if (parsed.command) summary += ': ' + parsed.command.substring(0, 120);
                    else if (parsed.path) summary += ': ' + parsed.path;
                    else if (parsed.query) summary += ': ' + parsed.query.substring(0, 120);
                    else summary += ': ' + JSON.stringify(parsed).substring(0, 120);
                    body = JSON.stringify(parsed, null, 2);
                } catch (e) {
                    summary += ': ' + args.substring(0, 120);
                }
            }
            div.innerHTML = '<div class="text-xs font-semibold text-[#a78bfa] cursor-pointer select-none" onclick="this.parentElement.classList.toggle(\'expanded\')">' + escapeHtml(summary) + '</div>' +
                '<div class=""><pre class="m-0 text-xs leading-normal overflow-x-auto whitespace-pre">' + escapeHtml(body) + '</pre></div>';
            break;
        }
        case 'statusupdate': {
            div.className += ' text-muted-foreground/70 text-xs py-0.5 px-2';
            const usage = extractField(block.lines, 'token_usage');
            const ctx = extractField(block.lines, 'context_tokens');
            const maxCtx = extractField(block.lines, 'max_context_tokens');
            let status = '';
            if (ctx && maxCtx) {
                const pct = Math.round((parseInt(ctx) / parseInt(maxCtx)) * 100);
                status = 'Context: ' + ctx + ' / ' + maxCtx + ' (' + pct + '%)';
            } else if (usage) {
                status = 'Tokens: ' + usage.substring(0, 100);
            } else {
                status = 'Status Update';
            }
            div.textContent = status;
            break;
        }
        case 'textpart': {
            div.className += ' my-1.5 rounded-r-md border-l-[3px] border-l-muted-foreground bg-muted-foreground/5 px-3 py-2';
            let text = extractField(block.lines, 'text');
            text = unwrapHardBreaks(text);
            const textHtml = typeof marked !== 'undefined' ? marked.parse(text) : escapeHtml(text);
            div.innerHTML = '<div class="text-xs font-semibold text-muted-foreground mb-1">文本</div>' +
                '<div class="kimi-markdown text-foreground/80 italic">' + textHtml + '</div>';
            break;
        }
        default: {
            div.className += ' text-foreground whitespace-pre-wrap py-0.5';
            div.textContent = rawText;
            break;
        }
    }
    return div;
}

export function flushLogBuffer() {
    if (logBuffer.length === 0) return;
    const lines = logBuffer.splice(0, logBuffer.length);
    rawLogLines.push(...lines);
    const frag = document.createDocumentFragment();

    if (renderMode === 'rendered') {
        const blocks = parseKimiBlocks(lines);
        const oldPending = terminalBody?.querySelector('[data-pending="true"]');
        if (oldPending) oldPending.remove();
        for (const block of blocks) {
            frag.appendChild(createRenderedBlock(block));
            lineCount += block.lines.length || 1;
        }
        if (parsedBlockPending) {
            frag.appendChild(createRenderedBlock(parsedBlockPending, true));
            lineCount += parsedBlockPending.lines.length || 1;
        }
    } else {
        for (const text of lines) {
            frag.appendChild(createLogLine(text));
        }
    }

    terminalBody?.appendChild(frag);

    while (terminalBody && terminalBody.childElementCount > 5000) {
        terminalBody.removeChild(terminalBody.firstElementChild);
    }

    if (lineCountEl) lineCountEl.textContent = lineCount + ' lines';

    if (isAutoScroll) {
        scrollToBottom();
    }
}

// Wire up buttons
document.getElementById('close-terminal-btn')?.addEventListener('click', closeTerminal);
document.getElementById('clear-terminal-btn')?.addEventListener('click', clearTerminal);
document.getElementById('render-toggle-btn')?.addEventListener('click', toggleRenderMode);
