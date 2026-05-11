#!/usr/bin/env python3
"""
E2E 测试报告生成器
解析 Maestro 的 debug output，生成详细的 Markdown 报告

Usage: python3 e2e-report-gen.py <debug_dir> <output_md>
"""

import json
import os
import sys
from glob import glob
from datetime import datetime


def describe_command(cmd_type, detail):
    """将 Maestro 命令转换为人类可读描述"""
    detail_str = str(detail)

    if 'visible' in detail_str:
        m = detail.get('condition', {}).get('visible', {})
        t = m.get('text') or m.get('textRegex', '')
        return f'assertVisible "{t}"'

    if 'tapOnElement' in detail_str:
        m = detail.get('element', {})
        t = m.get('text') or m.get('point', '') or m.get('id', '')
        return f'tapOn "{t}"'

    if 'tapOnPoint' in detail_str:
        return 'tapOn point'

    if 'inputTextCommand' in detail_str:
        t = detail.get('text', '')
        return f'inputText "{t}"'

    if 'launchAppCommand' in detail_str:
        return 'launchApp'

    if 'runFlowCommand' in detail_str:
        return f'runFlow {detail.get("path", "")}'

    if 'pressKeyCommand' in detail_str:
        return f'pressKey {detail.get("code", "")}'

    if 'waitForAnimationToEndCommand' in detail_str:
        return 'waitForAnimationToEnd'

    if 'hideKeyboardCommand' in detail_str:
        return 'hideKeyboard'

    if 'eraseTextCommand' in detail_str:
        return 'eraseText'

    if 'defineVariablesCommand' in detail_str:
        return 'defineVariables'

    if 'applyConfigurationCommand' in detail_str:
        return 'applyConfiguration'

    return cmd_type


def parse_flow_commands(json_path):
    """解析单个 flow 的 commands JSON"""
    with open(json_path) as f:
        data = json.load(f)

    steps = []
    first_failure = None

    for item in data:
        meta = item.get('metadata', {})
        status = meta.get('status', '?')
        cmd = item.get('command', {})
        cmd_type = list(cmd.keys())[0] if cmd else '?'
        cmd_detail = cmd.get(cmd_type, {})

        desc = describe_command(cmd_type, cmd_detail)
        error_msg = ''

        if status == 'FAILED':
            error_msg = meta.get('error', {}).get('message', '')
            if not first_failure:
                first_failure = {
                    'step': len(steps) + 1,
                    'desc': desc,
                    'reason': error_msg
                }

        steps.append({
            'index': len(steps) + 1,
            'status': status,
            'desc': desc,
            'error': error_msg,
            'duration': meta.get('duration', 0)
        })

    passed = sum(1 for s in steps if s['status'] == 'COMPLETED')
    failed = sum(1 for s in steps if s['status'] == 'FAILED')

    return {
        'steps': steps,
        'passed_steps': passed,
        'failed_steps': failed,
        'first_failure': first_failure
    }


def find_screenshots(debug_dir, flow_name):
    """查找 flow 对应的失败截图"""
    pattern = os.path.join(debug_dir, f'screenshot-*-({flow_name}).png')
    return [os.path.basename(f) for f in glob(pattern)]


def generate_report(debug_dir, output_md):
    # 读取 Maestro 整体输出
    maestro_output = ''
    maestro_log = os.path.join(debug_dir, 'maestro-output.log')
    if os.path.exists(maestro_log):
        with open(maestro_log) as f:
            maestro_output = f.read()

    # 解析整体结果
    passed_lines = [l for l in maestro_output.split('\n') if l.startswith('[Passed]')]
    failed_lines = [l for l in maestro_output.split('\n') if l.startswith('[Failed]')]
    passed_count = len(passed_lines)
    failed_count = len(failed_lines)
    total_count = passed_count + failed_count

    # 读取 exit code
    exit_code = 'N/A'
    exit_file = os.path.join(debug_dir, 'exit-code.txt')
    if os.path.exists(exit_file):
        with open(exit_file) as f:
            exit_code = f.read().strip()

    # 收集所有截图
    all_screenshots = sorted(glob(os.path.join(debug_dir, 'screenshot-*.png')))

    # 解析 flow 结果
    flow_results = []
    for line in maestro_output.split('\n'):
        if line.startswith('[Passed]') or line.startswith('[Failed]'):
            parts = line.split()
            if len(parts) >= 2:
                fname = parts[1]
                status = 'PASSED' if line.startswith('[Passed]') else 'FAILED'
                duration = parts[2] if len(parts) >= 3 else ''
                reason = ' '.join(parts[3:]) if len(parts) > 3 else ''
                flow_results.append({
                    'name': fname,
                    'status': status,
                    'duration': duration,
                    'reason': reason,
                    'line': line
                })

    lines = []
    lines.append('# E2E 测试详细报告')
    lines.append('')
    lines.append(f'- **测试时间**: {datetime.now().strftime("%Y-%m-%d %H:%M:%S")}')
    lines.append(f'- **总体结果**: {"✅ 通过" if exit_code == "0" else "❌ 失败"}')
    lines.append(f'- **Exit Code**: {exit_code}')
    lines.append('')
    lines.append('## 汇总')
    lines.append('')
    lines.append('| 指标 | 数量 |')
    lines.append('|------|------|')
    lines.append(f'| Flow 总数 | {total_count} |')
    lines.append(f'| ✅ 通过 | {passed_count} |')
    lines.append(f'| ❌ 失败 | {failed_count} |')
    lines.append(f'| 📸 失败截图 | {len(all_screenshots)} |')
    lines.append('')

    # 逐个 Flow 详细报告
    for flow in flow_results:
        fname = flow['name']
        status_icon = '✅' if flow['status'] == 'PASSED' else '❌'
        lines.append(f'## {status_icon} Flow: {fname}')
        lines.append('')
        lines.append(f'- **结果**: {flow["status"]} {flow["duration"]}')
        if flow['reason']:
            lines.append(f'- **失败原因**: {flow["reason"]}')

        cmd_json = os.path.join(debug_dir, f'commands-({fname}).json')
        if os.path.exists(cmd_json):
            detail = parse_flow_commands(cmd_json)
            lines.append(f'- **步骤统计**: {detail["passed_steps"]} 通过 / {detail["failed_steps"]} 失败 / {len(detail["steps"])} 总计')
            lines.append('')
            lines.append('### 步骤详情')
            lines.append('')
            lines.append('| # | 状态 | 步骤 | 错误 |')
            lines.append('|---|------|------|------|')
            for s in detail['steps']:
                icon = '✅' if s['status'] == 'COMPLETED' else '❌' if s['status'] == 'FAILED' else '⏳'
                err = s['error'] if s['error'] else '-'
                lines.append(f'| {s["index"]} | {icon} | {s["desc"]} | {err} |')

            if detail['first_failure']:
                ff = detail['first_failure']
                lines.append('')
                lines.append(f'**首个失败步骤**: 第 {ff["step"]} 步 `{ff["desc"]}`')
                lines.append(f'**错误信息**: {ff["reason"]}')
        else:
            lines.append('- **步骤详情**: 未找到 commands JSON')

        screenshots = find_screenshots(debug_dir, fname)
        if screenshots:
            lines.append('')
            lines.append('### 失败截图')
            lines.append('')
            for s in screenshots:
                lines.append(f'- `{s}`')

        lines.append('')

    if all_screenshots:
        lines.append('## 全部截图')
        lines.append('')
        for s in all_screenshots:
            lines.append(f'- `{os.path.basename(s)}`')
        lines.append('')

    lines.append('## 相关路径')
    lines.append('')
    lines.append(f'- 📸 Debug 目录: `{debug_dir}`')
    lines.append('')

    with open(output_md, 'w') as f:
        f.write('\n'.join(lines))

    # 打印摘要到 stdout
    print(f'✅ 报告已生成: {output_md}')
    print(f'   Flow: {total_count} ({passed_count} 通过, {failed_count} 失败)')
    print(f'   截图: {len(all_screenshots)}')
    if failed_count > 0:
        print(f'   失败 Flow:')
        for fl in failed_lines:
            print(f'     - {fl}')


if __name__ == '__main__':
    if len(sys.argv) < 3:
        print('Usage: e2e-report-gen.py <debug_dir> <output_md>')
        sys.exit(1)
    generate_report(sys.argv[1], sys.argv[2])
