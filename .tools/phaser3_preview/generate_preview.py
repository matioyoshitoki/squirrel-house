#!/usr/bin/env python3
"""
Phaser3 设计资产预览生成器

从 agent 产出的设计代码中解析色板和组件定义，
自动生成国风风格的 HTML 预览页面。

用法:
    python3 generate_preview.py --input ./designs/issue-18/ --output ./preview/
"""

import argparse
import os
import re
import sys
from pathlib import Path


def parse_ui_colors(content: str) -> dict:
    """解析 ui-colors.ts，提取 UITheme 中的颜色常量。"""
    colors = {}
    # 匹配 UITheme = { ... } as const;
    theme_match = re.search(r'export\s+const\s+UITheme\s*=\s*\{(.*?)\}\s+as\s+const', content, re.DOTALL)
    if theme_match:
        body = theme_match.group(1)
        # 匹配每个字段: name: 0xRRGGBB
        for m in re.finditer(r'([a-zA-Z_]\w*)\s*:\s*0x([0-9a-fA-F]{6})', body):
            name = m.group(1)
            hex_val = f"#{m.group(2)}"
            colors[name] = hex_val
    return colors


def parse_component_exports(content: str) -> list[dict]:
    """解析 index.ts，提取导出的组件名和配置类型。跳过注释行。"""
    components = []
    # 匹配: export { CompName, type CompConfig } from './CompFile';
    pattern = r"export\s*\{\s*([a-zA-Z_]+)\s*,\s*type\s+([a-zA-Z_]+)Config\s*\}\s*from\s*['\"]\.\/([a-zA-Z_]+)['\"]"
    for m in re.finditer(pattern, content):
        # 跳过注释行（行首有 //）
        line_start = content.rfind('\n', 0, m.start()) + 1
        line = content[line_start:m.start()].strip()
        if line.startswith('//') or line.startswith('*'):
            continue
        comp_name = m.group(1)
        config_name = m.group(2) + "Config"
        file_name = m.group(3)
        if comp_name.startswith("UI"):
            components.append({
                "name": comp_name,
                "config": config_name,
                "file": file_name,
                "layer": None,
            })
    # 也匹配没有 type 的简单导出
    simple_pattern = r"export\s*\{\s*([a-zA-Z_]+)\s*\}\s*from\s*['\"]\.\/([a-zA-Z_]+)['\"]"
    existing_names = {c["name"] for c in components}
    for m in re.finditer(simple_pattern, content):
        line_start = content.rfind('\n', 0, m.start()) + 1
        line = content[line_start:m.start()].strip()
        if line.startswith('//') or line.startswith('*'):
            continue
        comp_name = m.group(1)
        file_name = m.group(2)
        if comp_name.startswith("UI") and comp_name not in existing_names:
            components.append({
                "name": comp_name,
                "config": None,
                "file": file_name,
                "layer": None,
            })
    return components


def parse_types(content: str) -> dict[str, list[dict]]:
    """解析 types.ts，提取每个 interface/ type 的字段。"""
    type_defs = {}
    # 匹配 export interface Name { ... }
    interface_pattern = r'export\s+interface\s+([a-zA-Z_]+)\s*\{(.*?)\}'
    for m in re.finditer(interface_pattern, content, re.DOTALL):
        name = m.group(1)
        body = m.group(2)
        fields = []
        for line in body.split('\n'):
            line = line.strip()
            if not line or line.startswith('/**') or line.startswith('*') or line.startswith('*/'):
                continue
            # 匹配 fieldName?: type;
            fm = re.match(r'([a-zA-Z_]+)\??:\s*([^;]+);', line)
            if fm:
                fields.append({"name": fm.group(1), "type": fm.group(2).strip()})
        type_defs[name] = fields
    return type_defs


# 预定义的国风色板映射（当无法从 UITheme 常量解析时使用）
_COLOR_PALETTE = {
    'inkBlack': '#1a1a1a',
    'paperYellow': '#f5f0e6',
    'cinnabarRed': '#c45c48',
    'goldBrown': '#c9a227',
    'indigoBlue': '#3d5a80',
    'stoneCyan': '#008080',
    'white': '#ffffff',
    'lightInk': '#666666',
    'creamWhite': '#faf8f3',
    'jadeGreen': '#5a9a6e',
    'boneWhite': '#e8e4d9',
    'amberYellow': '#e6b800',
}


def infer_colors_from_components(ts_files: list[Path]) -> dict:
    """从组件文件中推断色板：扫描 UITheme.xxx 引用，使用预定义映射。"""
    used = set()
    for f in ts_files:
        text = f.read_text(encoding='utf-8')
        for m in re.finditer(r'UITheme\.([a-zA-Z_]\w*)', text):
            used.add(m.group(1))
    colors = {}
    for name in sorted(used):
        if name in _COLOR_PALETTE:
            colors[name] = _COLOR_PALETTE[name]
    return colors


def parse_components_from_files(ts_files: list[Path]) -> tuple[list[dict], dict[str, list[dict]]]:
    """直接从 .ts 组件文件中提取组件列表和类型定义（无 index.ts 时的 fallback）。"""
    components = []
    type_defs = {}
    seen_comps = set()

    for f in ts_files:
        text = f.read_text(encoding='utf-8')

        # 1) 提取 export class XXX extends ... (Phaser.GameObjects.* 或 UI* 基类)
        for m in re.finditer(
            r'export\s+class\s+([a-zA-Z_]\w+)\s+extends\s+([a-zA-Z_][\w.]*)',
            text
        ):
            comp_name = m.group(1)
            base_name = m.group(2)
            if comp_name.startswith('_') or comp_name in seen_comps:
                continue
            # 只接受 Phaser.GameObjects.* 或 UI 开头的基类（排除普通工具类）
            if not (base_name.startswith('Phaser.GameObjects.') or base_name.startswith('UI')):
                continue
            seen_comps.add(comp_name)
            config_name = f"{comp_name}Config"
            components.append({
                "name": comp_name,
                "config": config_name,
                "file": f.stem,
                "layer": None,
            })

        # 2) 提取 export interface XXXConfig { ... }（支持 extends）
        for m in re.finditer(r'export\s+interface\s+([a-zA-Z_]\w+)(?:\s+extends\s+[^{]+)?\s*\{(.*?)\}', text, re.DOTALL):
            iface_name = m.group(1)
            body = m.group(2)
            fields = []
            for line in body.split('\n'):
                line = line.strip()
                if not line or line.startswith('/**') or line.startswith('*') or line.startswith('*/'):
                    continue
                fm = re.match(r'([a-zA-Z_]+)\??:\s*([^;]+);', line)
                if fm:
                    fields.append({"name": fm.group(1), "type": fm.group(2).strip()})
            type_defs[iface_name] = fields

    return components, type_defs


def _clean_jsdoc(raw: str) -> str:
    """清理 JSDoc 原始文本，提取第一句简短描述。"""
    lines = []
    for line in raw.split('\n'):
        line = re.sub(r'^[ \t]*\*\s?', '', line).strip()
        if line and not line.startswith('@'):
            lines.append(line)
    short = lines[0] if lines else ''
    if len(short) > 80:
        short = short[:77] + '...'
    return short


def parse_infra_exports(content: str) -> list[dict]:
    """解析基础设施文件（types.ts / utils.ts），提取 export 的符号名和 JSDoc 简短描述。

    策略：先找到所有 JSDoc 块，再用 re.match 严格检查 JSDoc 结束后紧接着是否是 export，
    防止跨 export / 跨字段注释匹配。
    """
    exports = []
    export_kw = r'(?:const|function|type|interface|class|enum)'
    for m in re.finditer(r'/\*\*(.*?)\*/', content, re.DOTALL):
        after = content[m.end():m.end() + 200]
        em = re.match(r'(?:[ \t]*\n){0,2}[ \t]*export\s+(' + export_kw + r')\s+([a-zA-Z_][a-zA-Z0-9_]*)', after)
        if not em:
            continue
        name = em.group(2)
        kind = em.group(1)
        desc = _clean_jsdoc(m.group(1))
        exports.append({"name": name, "desc": desc, "kind": kind})

    # 也匹配 export type { Name1, Name2 } from './file' 中的具名导出
    for m in re.finditer(r'^[ \t]*export\s+type\s*\{\s*([^}]+)\}', content, re.MULTILINE):
        for item in m.group(1).split(','):
            item = item.strip()
            jsdoc_match = re.search(r'/\*\*(.*?)\*/', item, re.DOTALL)
            desc = ""
            if jsdoc_match:
                desc = _clean_jsdoc(jsdoc_match.group(1))
                item = re.sub(r'/\*.*?\*/', '', item, flags=re.DOTALL)
            clean = item.strip()
            if clean.startswith('type '):
                clean = clean[5:].strip()
            if clean and not clean.startswith('*') and not any(e["name"] == clean for e in exports):
                exports.append({"name": clean, "desc": desc or "Type export", "kind": "type"})

    # 匹配 export { Name1, Name2 } from './file' 中的具名值导出
    for m in re.finditer(r'^[ \t]*export\s*\{\s*([^}]+)\}\s*from', content, re.MULTILINE):
        for item in m.group(1).split(','):
            item = item.strip()
            jsdoc_match = re.search(r'/\*\*(.*?)\*/', item, re.DOTALL)
            desc = ""
            if jsdoc_match:
                desc = _clean_jsdoc(jsdoc_match.group(1))
                item = re.sub(r'/\*.*?\*/', '', item, flags=re.DOTALL)
            clean = item.strip()
            if clean.startswith('type '):
                clean = clean[5:].strip()
                kind = "type"
            else:
                kind = "value"
            if clean and not clean.startswith('*') and not any(e["name"] == clean for e in exports):
                exports.append({"name": clean, "desc": desc or "Re-export", "kind": kind})

    return exports


def parse_design_spec_components(content: str, known_components: list[str]) -> dict[str, str]:
    """解析 design-spec.md，提取每个组件的描述。"""
    descriptions = {}
    for comp in known_components:
        # 匹配 ## 3.X CompName 或 ### CompName
        pattern = rf'(?:##|###)\s+\d*\.?\s*{re.escape(comp)}\b(.*?)(?=(?:##|###)\s+\d|\Z)'
        m = re.search(pattern, content, re.DOTALL | re.IGNORECASE)
        if m:
            desc = m.group(1).strip()
            # 只取第一行非空行作为简短描述
            for line in desc.split('\n'):
                line = line.strip()
                if line and not line.startswith('#') and not line.startswith('|') and not line.startswith('-'):
                    descriptions[comp] = line[:120]
                    break
    return descriptions


def get_component_visual(name: str, colors: dict) -> str:
    """根据组件名称和色板生成对应的 CSS 渲染 HTML。"""
    cinnabar = colors.get('cinnabarRed', '#c45c48')
    gold = colors.get('goldBrown', '#c9a227')
    indigo = colors.get('indigoBlue', '#3d5a80')
    ink = colors.get('inkBlack', '#1a1a1a')
    paper = colors.get('paperYellow', '#f5f0e6')
    stone = colors.get('stoneCyan', '#008080')
    white = colors.get('white', '#ffffff')
    light_ink = colors.get('lightInk', '#666666')

    # 通用容器样式
    box = 'display:flex;align-items:center;justify-content:center;height:100%;width:100%;'

    visuals = {
        'UIButton': (
            f'<div style="{box}">'
            f'<div style="background:{cinnabar};color:#fff;padding:8px 24px;border-radius:4px;'
            f'font-size:13px;font-weight:600;border:1px solid {gold};box-shadow:0 2px 6px rgba(0,0,0,0.18);'
            f'letter-spacing:2px;">确定</div></div>'
        ),
        'UIIconButton': (
            f'<div style="{box}">'
            f'<div style="width:38px;height:38px;border-radius:50%;background:{paper};'
            f'border:2px solid {gold};display:flex;align-items:center;justify-content:center;'
            f'color:{cinnabar};font-size:18px;box-shadow:0 1px 4px rgba(0,0,0,0.1);">✦</div></div>'
        ),
        'UIPanel': (
            f'<div style="{box}">'
            f'<div style="background:{paper};border:1px solid {gold};border-radius:6px;'
            f'padding:12px 14px;width:90%;box-shadow:0 2px 8px rgba(0,0,0,0.08);">'
            f'<div style="height:3px;background:{gold};border-radius:2px;margin-bottom:8px;"></div>'
            f'<div style="height:6px;background:#e0dcd4;border-radius:2px;width:70%;margin-bottom:5px;"></div>'
            f'<div style="height:6px;background:#e0dcd4;border-radius:2px;width:50%;"></div></div></div>'
        ),
        'UIDialog': (
            f'<div style="{box}">'
            f'<div style="background:#fff;border:2px solid {gold};border-radius:8px;'
            f'padding:14px 16px;width:85%;box-shadow:0 6px 20px rgba(0,0,0,0.15);text-align:center;">'
            f'<div style="font-size:13px;color:{ink};margin-bottom:10px;font-weight:600;">确认操作</div>'
            f'<div style="display:flex;gap:10px;justify-content:center;">'
            f'<div style="background:{cinnabar};color:#fff;padding:5px 16px;border-radius:3px;font-size:11px;letter-spacing:1px;">确定</div>'
            f'<div style="background:{paper};color:{ink};padding:5px 16px;border-radius:3px;font-size:11px;border:1px solid {gold};">取消</div>'
            f'</div></div></div>'
        ),
        'UIList': (
            f'<div style="{box}">'
            f'<div style="background:#fff;border:1px solid {gold};border-radius:4px;overflow:hidden;width:90%;">'
            f'<div style="padding:6px 10px;border-bottom:1px solid #eee;font-size:11px;color:{ink};display:flex;align-items:center;gap:6px;">'
            f'<span style="width:6px;height:6px;background:{stone};border-radius:50%;"></span>列表项一</div>'
            f'<div style="padding:6px 10px;border-bottom:1px solid #eee;font-size:11px;color:{ink};display:flex;align-items:center;gap:6px;">'
            f'<span style="width:6px;height:6px;background:{stone};border-radius:50%;"></span>列表项二</div>'
            f'<div style="padding:6px 10px;font-size:11px;color:{ink};display:flex;align-items:center;gap:6px;">'
            f'<span style="width:6px;height:6px;background:{stone};border-radius:50%;"></span>列表项三</div>'
            f'</div></div>'
        ),
        'UIListItem': (
            f'<div style="{box}">'
            f'<div style="display:flex;align-items:center;gap:8px;padding:7px 12px;background:#fff;'
            f'border:1px solid {gold};border-radius:4px;width:85%;">'
            f'<div style="width:22px;height:22px;border-radius:50%;background:{stone};"></div>'
            f'<div><div style="font-size:11px;color:{ink};font-weight:600;">列表项标题</div>'
            f'<div style="font-size:10px;color:{light_ink};">辅助描述文本</div></div></div></div>'
        ),
        'UIText': (
            f'<div style="{box}flex-direction:column;gap:4px;">'
            f'<div style="font-size:14px;color:{ink};font-weight:600;letter-spacing:1px;">国风标题文本</div>'
            f'<div style="font-size:11px;color:{light_ink};">辅助说明文字内容</div></div>'
        ),
        'UIAvatar': (
            f'<div style="{box}">'
            f'<div style="width:44px;height:44px;border-radius:50%;background:{paper};'
            f'border:2px solid {gold};display:flex;align-items:center;justify-content:center;'
            f'color:{cinnabar};font-size:20px;font-weight:bold;box-shadow:0 2px 6px rgba(0,0,0,0.1);">甲</div></div>'
        ),
        'UIProgressBar': (
            f'<div style="{box}flex-direction:column;gap:3px;width:85%;">'
            f'<div style="background:#e0dcd4;border-radius:10px;height:10px;overflow:hidden;width:100%;">'
            f'<div style="background:{cinnabar};width:65%;height:100%;border-radius:10px;"></div></div>'
            f'<div style="font-size:10px;color:{light_ink};text-align:right;width:100%;">65%</div></div>'
        ),
        'UIResourceDisplay': (
            f'<div style="{box}">'
            f'<div style="display:flex;align-items:center;gap:8px;background:{paper};'
            f'padding:7px 12px;border-radius:4px;border:1px solid {gold};">'
            f'<span style="color:{gold};font-size:16px;">◆</span>'
            f'<span style="font-size:13px;color:{ink};font-weight:bold;letter-spacing:1px;">1,234</span>'
            f'<span style="font-size:10px;color:{light_ink};">元宝</span></div></div>'
        ),
        'UISegmentedControl': (
            f'<div style="{box}">'
            f'<div style="display:flex;border:1px solid {gold};border-radius:4px;overflow:hidden;width:90%;">'
            f'<div style="flex:1;background:{cinnabar};color:#fff;padding:5px 8px;font-size:11px;text-align:center;font-weight:600;">全部</div>'
            f'<div style="flex:1;background:#fff;color:{ink};padding:5px 8px;font-size:11px;text-align:center;">未完成</div>'
            f'<div style="flex:1;background:#fff;color:{ink};padding:5px 8px;font-size:11px;text-align:center;">已完成</div>'
            f'</div></div>'
        ),
        'UIIcon': (
            f'<div style="{box}gap:10px;">'
            f'<span style="color:{cinnabar};font-size:22px;">★</span>'
            f'<span style="color:{gold};font-size:22px;">◆</span>'
            f'<span style="color:{stone};font-size:22px;">●</span></div>'
        ),
        'UIAlert': (
            f'<div style="{box}">'
            f'<div style="background:#fff8f0;border:1px solid {cinnabar};border-radius:4px;'
            f'padding:8px 12px;display:flex;align-items:center;gap:8px;width:90%;">'
            f'<span style="color:{cinnabar};font-size:16px;font-weight:bold;">!</span>'
            f'<span style="font-size:11px;color:{ink};">这是一条警告提示信息</span></div></div>'
        ),
        'UIBadge': (
            f'<div style="{box}gap:8px;">'
            f'<span style="background:{cinnabar};color:#fff;padding:3px 10px;border-radius:10px;font-size:10px;font-weight:600;">热门</span>'
            f'<span style="background:{gold};color:#fff;padding:3px 10px;border-radius:10px;font-size:10px;font-weight:600;">新品</span>'
            f'<span style="background:{stone};color:#fff;padding:3px 10px;border-radius:10px;font-size:10px;font-weight:600;">限时</span></div>'
        ),
        'UITooltip': (
            f'<div style="{box}">'
            f'<div style="position:relative;display:inline-block;">'
            f'<div style="background:{ink};color:#fff;padding:5px 10px;border-radius:4px;font-size:11px;white-space:nowrap;">提示文本</div>'
            f'<div style="width:0;height:0;border-left:5px solid transparent;border-right:5px solid transparent;'
            f'border-top:5px solid {ink};margin:0 auto;"></div></div></div>'
        ),
        'UISlider': (
            f'<div style="{box}width:85%;">'
            f'<div style="position:relative;height:20px;display:flex;align-items:center;width:100%;">'
            f'<div style="width:100%;height:4px;background:#e0dcd4;border-radius:2px;"></div>'
            f'<div style="position:absolute;left:40%;width:18px;height:18px;background:{cinnabar};border-radius:50%;'
            f'border:2px solid #fff;box-shadow:0 1px 4px rgba(0,0,0,0.25);cursor:pointer;"></div></div></div>'
        ),
        'UIToggle': (
            f'<div style="{box}gap:10px;">'
            f'<div style="width:38px;height:22px;background:{cinnabar};border-radius:11px;position:relative;">'
            f'<div style="position:absolute;right:2px;top:2px;width:18px;height:18px;background:#fff;border-radius:50%;box-shadow:0 1px 2px rgba(0,0,0,0.2);"></div></div>'
            f'<span style="font-size:12px;color:{ink};font-weight:600;">开启</span></div>'
        ),
        'UIStepper': (
            f'<div style="{box}">'
            f'<div style="display:flex;align-items:center;border:1px solid {gold};border-radius:4px;overflow:hidden;">'
            f'<div style="width:30px;height:30px;background:{paper};display:flex;align-items:center;justify-content:center;'
            f'color:{ink};font-size:18px;cursor:pointer;font-weight:bold;">-</div>'
            f'<div style="width:44px;height:30px;background:#fff;display:flex;align-items:center;justify-content:center;'
            f'color:{ink};font-size:14px;font-weight:bold;border-left:1px solid {gold};border-right:1px solid {gold};">3</div>'
            f'<div style="width:30px;height:30px;background:{paper};display:flex;align-items:center;justify-content:center;'
            f'color:{ink};font-size:18px;cursor:pointer;font-weight:bold;">+</div></div></div>'
        ),
        'UIDropdown': (
            f'<div style="{box}width:85%;">'
            f'<div style="background:#fff;border:1px solid {gold};border-radius:4px;padding:7px 10px;'
            f'font-size:12px;color:{ink};display:flex;justify-content:space-between;align-items:center;">'
            f'<span>请选择选项</span><span style="color:{gold};font-size:10px;">▼</span></div></div>'
        ),
        'UIPagination': (
            f'<div style="{box}gap:5px;">'
            f'<div style="width:26px;height:26px;background:{paper};border:1px solid {gold};border-radius:3px;'
            f'display:flex;align-items:center;justify-content:center;font-size:11px;color:{ink};cursor:pointer;">1</div>'
            f'<div style="width:26px;height:26px;background:{cinnabar};border:1px solid {cinnabar};border-radius:3px;'
            f'display:flex;align-items:center;justify-content:center;font-size:11px;color:#fff;font-weight:600;cursor:pointer;">2</div>'
            f'<div style="width:26px;height:26px;background:{paper};border:1px solid {gold};border-radius:3px;'
            f'display:flex;align-items:center;justify-content:center;font-size:11px;color:{ink};cursor:pointer;">3</div></div>'
        ),
        'UILoading': (
            f'<div style="{box}gap:5px;">'
            f'<div class="load-dot" style="width:9px;height:9px;background:{cinnabar};border-radius:50%;animation:loadPulse 1s infinite;"></div>'
            f'<div class="load-dot" style="width:9px;height:9px;background:{gold};border-radius:50%;animation:loadPulse 1s infinite 0.2s;"></div>'
            f'<div class="load-dot" style="width:9px;height:9px;background:{stone};border-radius:50%;animation:loadPulse 1s infinite 0.4s;"></div></div>'
        ),
        'UIAlert': (
            f'<div style="{box}">'
            f'<div style="display:flex;align-items:center;background:{paper};border:1px solid {gold};border-radius:4px;'
            f'padding:8px 14px;width:85%;box-shadow:0 2px 6px rgba(0,0,0,0.08);position:relative;overflow:hidden;">'
            f'<div style="position:absolute;left:0;top:2px;bottom:2px;width:3px;background:{indigo};border-radius:2px;"></div>'
            f'<span style="font-size:12px;color:{ink};margin-left:6px;">提示消息内容</span></div></div>'
        ),
        'UIToast': (
            f'<div style="{box}">'
            f'<div style="display:flex;align-items:center;background:{paper};border:1px solid {gold};border-radius:4px;'
            f'padding:10px 14px;width:85%;box-shadow:0 2px 6px rgba(0,0,0,0.08);position:relative;overflow:hidden;">'
            f'<div style="position:absolute;left:0;top:2px;bottom:2px;width:3px;background:{stone};border-radius:2px;"></div>'
            f'<span style="font-size:12px;color:{ink};margin-left:6px;">轻提示消息</span></div></div>'
        ),
        'UIDialog': (
            f'<div style="{box}">'
            f'<div style="background:#fff;border:2px solid {gold};border-radius:8px;'
            f'padding:14px 16px;width:80%;box-shadow:0 6px 20px rgba(0,0,0,0.15);text-align:center;">'
            f'<div style="font-size:13px;color:{ink};margin-bottom:10px;font-weight:600;">确认操作</div>'
            f'<div style="display:flex;gap:10px;justify-content:center;">'
            f'<div style="background:{cinnabar};color:#fff;padding:5px 16px;border-radius:3px;font-size:11px;letter-spacing:1px;">确定</div>'
            f'<div style="background:{paper};color:{ink};padding:5px 16px;border-radius:3px;font-size:11px;border:1px solid {gold};">取消</div>'
            f'</div></div></div>'
        ),
    }
    return visuals.get(name, f'<div style="{box}color:{light_ink};font-size:12px;letter-spacing:1px;">{name[2:]}</div>')


def generate_html(colors: dict, components: list[dict], type_defs: dict, descriptions: dict, infra_exports: dict, issue_number: str) -> str:
    """生成国风预览 HTML 页面。"""
    # 颜色按名称排序
    sorted_colors = sorted(colors.items(), key=lambda x: x[0])

    # 判断本次产出性质
    is_infra_rebuild = len(components) == 0 and (len(infra_exports.get("types", [])) > 0 or len(infra_exports.get("utils", [])) > 0)
    page_title = f"Issue #{issue_number} — {'UI基础设施重建' if is_infra_rebuild else '国风 UI 组件预览'}"
    page_subtitle = "基础设施: types.ts + utils.ts + index.ts 统一入口" if is_infra_rebuild else "由设计代码自动生成 · 色板与组件定义源自 src/ui/"

    # 色板卡片 HTML
    color_cards = ""
    for name, hex_val in sorted_colors:
        # 计算文字颜色（深色背景用白字，浅色用黑字）
        r = int(hex_val[1:3], 16)
        g = int(hex_val[3:5], 16)
        b = int(hex_val[5:7], 16)
        brightness = (r * 299 + g * 587 + b * 114) / 1000
        text_color = "#ffffff" if brightness < 128 else "#2c2c2c"
        color_cards += f"""
        <div class="color-card">
            <div class="color-swatch" style="background:{hex_val};color:{text_color}">{hex_val}</div>
            <div class="color-name">{name}</div>
        </div>"""

    # 基础设施卡片 HTML — 紧凑网格
    infra_cards = ""
    for category in ["types", "utils"]:
        items = infra_exports.get(category, [])
        if items:
            infra_cards += f'<div class="infra-category">📁 {category}.ts</div>'
            infra_cards += '<div class="infra-grid">'
            for item in items:
                kind = item.get("kind", "value")
                infra_cards += f"""
                <div class="infra-card">
                    <div class="infra-kind">{kind}</div>
                    <div class="infra-name">{item['name']}</div>
                    <div class="infra-desc">{item.get('desc', '')}</div>
                </div>"""
            infra_cards += '</div>'

    # 组件卡片 HTML
    comp_cards = ""
    for comp in components:
        name = comp["name"]
        config = comp["config"]
        desc = descriptions.get(name, "")
        fields_html = ""
        if config and config in type_defs:
            fields = type_defs[config][:6]  # 最多显示6个字段
            for f in fields:
                fields_html += f'<span class="field"><code>{f["name"]}?</code>: {f["type"]}</span>'
        # 组件占位视觉：用 CSS 画一个简化轮廓
        comp_cards += f"""
        <div class="comp-card">
            <div class="comp-header">
                <span class="comp-name">{name}</span>
                <span class="comp-badge">{comp.get('layer', 'Layer 1')}</span>
            </div>
            <div class="comp-visual">
                {get_component_visual(name, colors)}
            </div>
            <div class="comp-fields">{fields_html}</div>
            {f'<div class="comp-desc">{desc}</div>' if desc else ''}
        </div>"""

    # 完整的 HTML 页面
    html = f"""<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>{page_title}</title>
<style>
  :root {{
    --bg: #f7f5f0;
    --card-bg: #ffffff;
    --border: #e0dcd4;
    --ink-black: #1a1a1a;
    --paper-yellow: #f5f0e6;
    --cinnabar-red: #c45c48;
    --stone-cyan: #4a9eff;
    --indigo-blue: #3d5a80;
    --gold-brown: #c9a227;
    --light-ink: #666666;
    --white: #ffffff;
    --ink-text: #2c2c2c;
  }}
  * {{ box-sizing: border-box; margin: 0; padding: 0; }}
  body {{
    font-family: "Microsoft YaHei", "PingFang SC", "Noto Sans SC", sans-serif;
    background: var(--bg);
    color: var(--ink-text);
    line-height: 1.6;
    padding-bottom: 60px;
  }}
  header {{
    background: var(--ink-black);
    color: var(--white);
    padding: 28px 40px;
    border-bottom: 3px solid var(--gold-brown);
  }}
  header h1 {{ font-size: 22px; margin-bottom: 6px; }}
  header p {{ color: var(--light-ink); font-size: 13px; }}
  .container {{ max-width: 1200px; margin: 0 auto; padding: 28px 40px; }}

  .section {{
    background: var(--card-bg);
    border: 1px solid var(--border);
    border-radius: 6px;
    margin-bottom: 24px;
    overflow: hidden;
  }}
  .section-header {{
    background: var(--paper-yellow);
    padding: 14px 20px;
    border-bottom: 2px solid var(--gold-brown);
    font-size: 15px;
    font-weight: bold;
    color: var(--ink-black);
    display: flex;
    align-items: center;
    gap: 8px;
  }}
  .section-body {{ padding: 20px; }}

  /* Color Palette */
  .color-grid {{
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
    gap: 16px;
  }}
  .color-card {{
    border: 1px solid var(--border);
    border-radius: 6px;
    overflow: hidden;
    text-align: center;
  }}
  .color-swatch {{
    height: 64px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 13px;
    font-weight: bold;
    font-family: "SF Mono", Consolas, monospace;
  }}
  .color-name {{
    padding: 8px;
    font-size: 12px;
    color: var(--light-ink);
    font-family: "SF Mono", Consolas, monospace;
    background: #faf8f4;
  }}

  /* Component Cards */
  .comp-grid {{
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 16px;
  }}
  .comp-card {{
    border: 1px solid var(--border);
    border-radius: 6px;
    padding: 16px;
    background: #faf8f4;
    transition: box-shadow 0.2s;
  }}
  .comp-card:hover {{
    box-shadow: 0 4px 16px rgba(0,0,0,0.08);
    border-color: var(--gold-brown);
  }}
  .comp-header {{
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
  }}
  .comp-name {{
    font-size: 15px;
    font-weight: bold;
    color: var(--ink-black);
    font-family: "SF Mono", Consolas, monospace;
  }}
  .comp-badge {{
    font-size: 10px;
    padding: 2px 8px;
    border-radius: 10px;
    background: var(--paper-yellow);
    color: var(--indigo-blue);
    border: 1px solid var(--gold-brown);
  }}
  .comp-visual {{
    background: var(--white);
    border: 1px dashed var(--border);
    border-radius: 4px;
    height: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 12px;
  }}
  .comp-fields {{
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
    margin-bottom: 8px;
  }}
  .comp-fields .field {{
    font-size: 11px;
    padding: 2px 8px;
    background: var(--white);
    border: 1px solid var(--border);
    border-radius: 3px;
    color: var(--indigo-blue);
  }}
  .comp-fields code {{
    font-family: "SF Mono", Consolas, monospace;
    color: var(--cinnabar-red);
  }}
  .comp-desc {{
    font-size: 12px;
    color: var(--light-ink);
    border-top: 1px solid var(--border);
    padding-top: 8px;
    margin-top: 4px;
  }}

  /* Infrastructure exports */
  .infra-category {{
    font-size: 13px;
    font-weight: bold;
    color: var(--indigo-blue);
    margin: 16px 0 10px 0;
    padding-bottom: 4px;
    border-bottom: 1px solid var(--border);
  }}
  .infra-grid {{
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 10px;
    margin-bottom: 8px;
  }}
  .infra-card {{
    border: 1px solid var(--border);
    border-radius: 6px;
    padding: 10px 12px;
    background: #faf8f4;
    transition: box-shadow 0.15s, border-color 0.15s;
  }}
  .infra-card:hover {{
    box-shadow: 0 2px 10px rgba(0,0,0,0.06);
    border-color: var(--gold-brown);
  }}
  .infra-kind {{
    font-size: 9px;
    padding: 1px 5px;
    border-radius: 3px;
    background: var(--indigo-blue);
    color: #fff;
    display: inline-block;
    margin-bottom: 6px;
    font-family: "SF Mono", Consolas, monospace;
    text-transform: uppercase;
  }}
  .infra-name {{
    font-family: "SF Mono", Consolas, monospace;
    font-size: 12px;
    color: var(--cinnabar-red);
    font-weight: 600;
    margin-bottom: 4px;
  }}
  .infra-desc {{
    font-size: 11px;
    color: var(--light-ink);
    line-height: 1.4;
    max-height: 2.8em;
    overflow: hidden;
  }}

  .status-bar {{
    text-align: center;
    padding: 16px;
    color: var(--light-ink);
    font-size: 12px;
  }}
  @keyframes loadPulse {{
    0%, 100% {{ opacity: 1; transform: scale(1); }}
    50% {{ opacity: 0.4; transform: scale(0.8); }}
  }}
</style>
</head>
<body>
<header>
  <h1>🎨 {page_title}</h1>
  <p>{page_subtitle}</p>
</header>

<div class="container">
  <div class="section">
    <div class="section-header">🎨 国风水墨色板（{len(colors)} 色）</div>
    <div class="section-body">
      <div class="color-grid">
        {color_cards}
      </div>
    </div>
  </div>

  {f'<div class="section"><div class="section-header">🏗️ 基础设施导出</div><div class="section-body">{infra_cards}</div></div>' if is_infra_rebuild else ''}

  <div class="section">
    <div class="section-header">🧩 组件清单（共 {len(components)} 个{'' if not is_infra_rebuild else ' · Layer 1/2 组件注释中，见迁移指南'})</div>
    <div class="section-body">
      <div class="comp-grid">
        {comp_cards}
      </div>
    </div>
  </div>

  <div class="status-bar">
    ⚠️ 本页面由系统自动生成，仅展示组件接口定义与色板。实际视觉效果以 Phaser3 渲染为准。
  </div>
</div>
</body>
</html>"""
    return html


def main():
    parser = argparse.ArgumentParser(description="Phaser3 设计资产预览生成器")
    parser.add_argument("--input", required=True, help="设计资产输入目录 (designs/issue-N/)")
    parser.add_argument("--output", required=True, help="预览输出目录")
    parser.add_argument("--issue", default="", help="Issue 编号")
    args = parser.parse_args()

    input_dir = Path(args.input)
    output_dir = Path(args.output)

    if not input_dir.exists():
        print(f"❌ 输入目录不存在: {input_dir}", file=sys.stderr)
        sys.exit(1)

    # 收集所有 ts 文件（供多处 fallback 使用）
    ui_dir = input_dir / "src" / "ui"
    ts_files = list(ui_dir.glob("*.ts")) if ui_dir.exists() else []

    # 1. 解析色板（优先 ui-colors.ts，回退到 utils.ts，再回退到组件文件推断）
    colors = {}
    colors_file = input_dir / "src" / "ui" / "styles" / "ui-colors.ts"
    utils_file = input_dir / "src" / "ui" / "utils.ts"
    if colors_file.exists():
        colors = parse_ui_colors(colors_file.read_text(encoding="utf-8"))
        print(f"✅ 解析色板: {len(colors)} 个颜色 (from styles/ui-colors.ts)")
    elif utils_file.exists():
        colors = parse_ui_colors(utils_file.read_text(encoding="utf-8"))
        print(f"✅ 解析色板: {len(colors)} 个颜色 (from utils.ts)")
    elif ts_files:
        colors = infer_colors_from_components(ts_files)
        print(f"✅ 推断色板: {len(colors)} 个颜色 (from 组件文件)")
    else:
        print(f"⚠️ 未找到色板文件: {colors_file} 或 {utils_file}")

    # 2. 解析组件导出
    components = []
    type_defs = {}
    index_file = input_dir / "src" / "ui" / "index.ts"
    if index_file.exists():
        components = parse_component_exports(index_file.read_text(encoding="utf-8"))
        print(f"✅ 解析组件: {len(components)} 个 (from index.ts)")
    elif ts_files:
        components, type_defs = parse_components_from_files(ts_files)
        print(f"✅ 解析组件: {len(components)} 个 (from 组件文件)")
    else:
        print(f"⚠️ 未找到入口文件: {index_file}")

    # 3. 解析类型定义
    types_file = input_dir / "src" / "ui" / "types.ts"
    if types_file.exists():
        type_defs = parse_types(types_file.read_text(encoding="utf-8"))
        print(f"✅ 解析类型: {len(type_defs)} 个 (from types.ts)")
    elif not type_defs and ts_files:
        _, type_defs = parse_components_from_files(ts_files)
        print(f"✅ 解析类型: {len(type_defs)} 个 (from 组件文件)")

    # 4. 解析基础设施导出（types + utils）
    infra_exports = {"types": [], "utils": []}
    if types_file.exists():
        infra_exports["types"] = parse_infra_exports(types_file.read_text(encoding="utf-8"))
        print(f"✅ 解析 types 导出: {len(infra_exports['types'])} 个")
    if utils_file.exists():
        infra_exports["utils"] = parse_infra_exports(utils_file.read_text(encoding="utf-8"))
        print(f"✅ 解析 utils 导出: {len(infra_exports['utils'])} 个")

    # 5. 解析 design-spec.md 获取描述
    descriptions = {}
    spec_file = input_dir / "design-spec.md"
    if spec_file.exists():
        descriptions = parse_design_spec_components(
            spec_file.read_text(encoding="utf-8"),
            [c["name"] for c in components]
        )
        print(f"✅ 解析设计规范描述: {len(descriptions)} 个")

    # 6. 推断 Layer
    for comp in components:
        name = comp["name"]
        if name in ["UISlider", "UIToggle", "UIStepper", "UIDropdown", "UIPagination", "UILoading"]:
            comp["layer"] = "Layer 2"
        else:
            comp["layer"] = "Layer 1"

    # 7. 生成 HTML
    issue_number = args.issue or input_dir.name.replace("issue-", "")
    html = generate_html(colors, components, type_defs, descriptions, infra_exports, issue_number)

    output_dir.mkdir(parents=True, exist_ok=True)
    output_file = output_dir / "index.html"
    output_file.write_text(html, encoding="utf-8")
    print(f"✅ 预览页面已生成: {output_file}")


if __name__ == "__main__":
    main()
