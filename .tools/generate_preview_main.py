#!/usr/bin/env python3
import os
import re
import sys

def extract_type_defs(content):
    """提取文件中定义的所有类型信息。
    返回 {name: info}，其中 info 为:
      - 'class': 普通类（无字段信息）
      - ('enum', first_value): 枚举
      - {'kind': 'data_class', 'fields': {fname: ftype}}: 数据类（有 final 字段）
    """
    types = {}
    # 先提取所有 class 定义及其字段
    for m in re.finditer(r'class\s+(\w+)\s*(?:extends\s+\w+)?\s*\{', content):
        class_name = m.group(1)
        class_start = m.end()
        # 找到类体的结束位置
        brace_depth = 0
        in_class = False
        class_end = len(content)
        for i in range(class_start, len(content)):
            if content[i] == '{':
                brace_depth += 1
                in_class = True
            elif content[i] == '}':
                brace_depth -= 1
                if in_class and brace_depth == 0:
                    class_end = i + 1
                    break
        class_body = content[class_start:class_end]
        # 提取 final 字段
        fields = {}
        for fm in re.finditer(r'final\s+(\w+(?:<[^>]+>)?)\s+(\w+);', class_body):
            fields[fm.group(2)] = fm.group(1)
        # 如果有 final 字段，认为是数据类
        if fields:
            types[class_name] = {'kind': 'data_class', 'fields': fields}
        else:
            types[class_name] = 'class'
    # 提取 enum
    for m in re.finditer(r'enum\s+(\w+)\s*\{([^}]*)\}', content, re.DOTALL):
        enum_name = m.group(1)
        enum_body = m.group(2)
        vals = re.findall(r'^\s*(\w+)(?:\s*[,/]|\s*$)', enum_body, re.MULTILINE)
        skip = {'true', 'false', 'null', 'const', 'final', 'var', 'return', 'if', 'else'}
        vals = [v for v in vals if v not in skip]
        first_val = vals[0] if vals else 'values.first'
        types[enum_name] = ('enum', first_val)
    return types

def extract_field_types(content):
    types = {}
    for m in re.finditer(r'final\s+(\w+(?:<[^>]+>)?)\s+(\w+);', content):
        types[m.group(2)] = m.group(1)
    return types

def extract_widgets(dart_file, all_types):
    with open(dart_file, 'r', encoding='utf-8') as f:
        content = f.read()

    widgets = []
    pattern = re.compile(r'class\s+(\w+)\s+extends\s+(StatelessWidget|StatefulWidget|Widget)', re.MULTILINE)
    for m in pattern.finditer(content):
        class_name = m.group(1)
        if class_name.startswith('_'):
            continue
        class_start = m.end()
        class_end = len(content)
        brace_depth = 0
        in_class = False
        for i in range(class_start, len(content)):
            if content[i] == '{':
                brace_depth += 1
                in_class = True
            elif content[i] == '}':
                brace_depth -= 1
                if in_class and brace_depth == 0:
                    class_end = i + 1
                    break
        class_body = content[class_start:class_end]
        field_types = extract_field_types(class_body)
        
        ctor_re = re.compile(rf'(?:const|)\s+{re.escape(class_name)}\s*\((.*?)\)', re.DOTALL)
        ctor_m = None
        for cm in ctor_re.finditer(content):
            if cm.start() >= class_start and cm.end() <= class_end:
                # 排除 "return EditProfileSheet(...)" 这类方法调用，只匹配构造函数
                line_start = content.rfind('\n', 0, cm.start()) + 1
                line_prefix = content[line_start:cm.start()].strip()
                if line_prefix.startswith('return'):
                    continue
                ctor_m = cm
        
        params = []
        if ctor_m:
            raw = ctor_m.group(1).strip()
            named_m = re.search(r'\{(.*?)\}', raw, re.DOTALL)
            if named_m:
                named = named_m.group(1)
                for pm in re.finditer(r'required\s+this\.(\w+)', named):
                    pname = pm.group(1)
                    ptype = field_types.get(pname, 'dynamic')
                    params.append({'type': ptype, 'name': pname, 'required': True})
                for pm in re.finditer(r'(?<!required\s)this\.(\w+)', named):
                    pname = pm.group(1)
                    ptype = field_types.get(pname, 'dynamic')
                    params.append({'type': ptype, 'name': pname, 'required': False})
                for pm in re.finditer(r'(?:required\s+)?((?!this\b|required\b)\w+(?:<[^>]+>)?)\s+((?!this\b|required\b)\w+)', named):
                    ptype = pm.group(1)
                    pname = pm.group(2)
                    is_req = 'required' in named[pm.start():pm.end()]
                    params.append({'type': ptype, 'name': pname, 'required': is_req})
        # For StatefulWidget, also check the corresponding State class body
        check_body = class_body
        if m.group(2) == 'StatefulWidget':
            # Try common state class naming patterns
            state_patterns = [
                rf'class\s+_{re.escape(class_name)}State\s+extends\s+State<\s*{re.escape(class_name)}\s*>',
                rf'class\s+_{re.escape(class_name[0].lower() + class_name[1:])}State\s+extends\s+State<\s*{re.escape(class_name)}\s*>',
            ]
            for sp in state_patterns:
                sm = re.search(sp, content)
                if sm:
                    s_start = sm.end()
                    s_depth = 0
                    s_in = False
                    s_end = len(content)
                    for i in range(s_start, len(content)):
                        if content[i] == '{':
                            s_depth += 1
                            s_in = True
                        elif content[i] == '}':
                            s_depth -= 1
                            if s_in and s_depth == 0:
                                s_end = i + 1
                                break
                    check_body = check_body + content[s_start:s_end]
                    break
        has_scaffold = 'Scaffold' in check_body or 'Expanded' in check_body
        widgets.append({'name': class_name, 'params': params, 'has_scaffold': has_scaffold})
    return widgets

def mock_value(param_type, param_name, known_types, all_type_defs, visited=None):
    """为给定类型生成 mock 值。支持递归构造数据类。"""
    if visited is None:
        visited = set()
    pt = param_type.strip()
    # 防止循环依赖
    if pt in visited:
        return f'{pt}()'
    visited.add(pt)
    try:
        if pt in all_type_defs:
            td = all_type_defs[pt]
            if isinstance(td, tuple) and td[0] == 'enum':
                return f'{pt}.{td[1]}'
            if isinstance(td, dict) and td.get('kind') == 'data_class':
                # 递归构造数据类实例
                fields = td.get('fields', {})
                args = []
                for fname, ftype in fields.items():
                    fval = mock_value(ftype, fname, known_types, all_type_defs, visited.copy())
                    args.append(f'{fname}: {fval}')
                if args:
                    return f'{pt}({", ".join(args)})'
                return f'{pt}()'
        if pt == 'String':
            return f'"{param_name}"'
        if pt == 'int':
            return '0'
        if pt == 'double':
            return '0.0'
        if pt == 'bool':
            return 'false'
        if 'Callback' in pt or 'VoidCallback' in pt or 'GestureTapCallback' in pt:
            return '() {}'
        if 'ValueChanged' in pt or 'ValueSetter' in pt:
            return '(v) {}'
        if pt.startswith('List'):
            return '[]'
        if pt.startswith('Map'):
            return '{}'
        if pt == 'Key':
            return 'const ValueKey("preview")'
        if pt in known_types:
            # 检查是否是可递归构造的数据类
            td = all_type_defs.get(pt)
            if isinstance(td, dict) and td.get('kind') == 'data_class':
                fields = td.get('fields', {})
                args = []
                for fname, ftype in fields.items():
                    fval = mock_value(ftype, fname, known_types, all_type_defs, visited.copy())
                    args.append(f'{fname}: {fval}')
                if args:
                    return f'{pt}({", ".join(args)})'
            return f'{pt}()'
        return f'{pt}()'
    finally:
        visited.discard(pt)

def build_instance(class_name, params, known_types, all_type_defs, has_scaffold=False):
    req_params = [p for p in params if p['required']]
    has_unknown = False
    if not req_params:
        code = f'{class_name}()'
    else:
        args = []
        for p in req_params:
            val = mock_value(p['type'], p['name'], known_types, all_type_defs)
            # 判断是否为无法安全 mock 的类型（非基础类型、非 enum、非 callback、非 List/Map、非已知数据类）
            if p['type'] not in {'String', 'int', 'double', 'bool', 'Key'} and not any(x in p['type'] for x in ['Callback', 'VoidCallback', 'Gesture', 'ValueChanged', 'List', 'Map']):
                td = all_type_defs.get(p['type'])
                if not (isinstance(td, tuple) and td[0] == 'enum') and not (isinstance(td, dict) and td.get('kind') == 'data_class'):
                    has_unknown = True
            args.append(f'{p["name"]}: {val}')
        code = f'{class_name}({", ".join(args)})'
    # 包含 Scaffold 的 widget 在嵌套 MaterialApp 中可能表现异常，用 SizedBox 固定尺寸
    if has_scaffold:
        code = f'SizedBox(height: 350, width: double.infinity, child: {code})'
    return code, not has_unknown

def generate_main_dart(preview_dir, output_file):
    imports = set()
    instances = []
    known_types = set()
    all_type_defs = {}

    # Skip widgetbook/ directory — it contains a standalone Widgetbook app
    # that should not be nested inside another MaterialApp
    skip_dirs = {'widgetbook'}

    for root, _, files in os.walk(preview_dir):
        if any(skip in root.split(os.sep) for skip in skip_dirs):
            continue
        for filename in files:
            if not filename.endswith('.dart'):
                continue
            filepath = os.path.join(root, filename)
            with open(filepath, 'r', encoding='utf-8') as f:
                content = f.read()
            types = extract_type_defs(content)
            all_type_defs.update(types)

    # 收集所有 widget，按名称去重（同名 widget 只保留第一次出现的）
    widget_files = {}  # name -> filepath
    for root, _, files in os.walk(preview_dir):
        if any(skip in root.split(os.sep) for skip in skip_dirs):
            continue
        for filename in files:
            if not filename.endswith('.dart'):
                continue
            filepath = os.path.join(root, filename)
            widgets = extract_widgets(filepath, all_type_defs)
            for w in widgets:
                known_types.add(w['name'])
                if w['name'] not in widget_files:
                    widget_files[w['name']] = filepath

    # 按文件分组，直接导入整个文件（包含 widget 和数据类）
    file_widgets = {}
    for name, filepath in widget_files.items():
        file_widgets.setdefault(filepath, []).append(name)
    for filepath, names in file_widgets.items():
        rel = os.path.relpath(filepath, preview_dir)
        import_path = 'preview/' + rel.replace('\\', '/')
        imports.add(f"import '{import_path}';")

    for name, filepath in widget_files.items():
        widgets = extract_widgets(filepath, all_type_defs)
        for w in widgets:
            if w['name'] != name:
                continue
            code, is_safe = build_instance(w['name'], w['params'], known_types, all_type_defs, w.get('has_scaffold', False))
            if is_safe:
                instances.append(f"            _buildCard('{w['name']}', {code}),")
            else:
                # 跳过不安全的 widget，显示占位提示
                instances.append(f"            _buildCard('{w['name']} (参数复杂，暂不支持预览)', const Center(child: Text('Complex params'))),")

    if not instances:
        instances.append("            const Center(child: Text('未找到可预览的 Widget')),")

    main = f"""import 'package:flutter/material.dart';
{chr(10).join(sorted(imports))}
import 'dart:ui';

void main() {{
  // 捕获构建错误，避免 release 模式下显示空白
  ErrorWidget.builder = (FlutterErrorDetails details) {{
    return Container(
      height: 60,
      width: double.infinity,
      color: const Color(0xFF2d1b1b),
      padding: const EdgeInsets.all(12),
      alignment: Alignment.centerLeft,
      child: Text(
        'Render Error: ' + details.exception.toString(),
        style: const TextStyle(color: Colors.redAccent, fontSize: 12),
      ),
    );
  }};
  // 捕获全局异步错误，防止未捕获异常导致整个应用白屏
  PlatformDispatcher.instance.onError = (Object error, StackTrace stack) {{
    return true;
  }};
  runApp(const DesignPreviewApp());
}}

class DesignPreviewApp extends StatelessWidget {{
  const DesignPreviewApp({{super.key}});

  @override
  Widget build(BuildContext context) {{
    return MaterialApp(
      title: 'Design Preview',
      debugShowCheckedModeBanner: false,
      theme: ThemeData.dark(),
      home: Container(
        color: const Color(0xFF0f172a),
        padding: const EdgeInsets.all(16),
        child: ListView(
          children: [
            const Center(
              child: Padding(
                padding: EdgeInsets.symmetric(vertical: 16),
                child: Text('Design Assets Preview', style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
              ),
            ),
{chr(10).join(instances)}
          ],
        ),
      ),
    );
  }}

  Widget _buildCard(String title, Widget child) {{
    return Card(
      margin: const EdgeInsets.only(bottom: 16),
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(title, style: const TextStyle(fontSize: 16, fontWeight: FontWeight.bold)),
            const SizedBox(height: 12),
            Container(
              height: 350,
              width: double.infinity,
              decoration: BoxDecoration(
                color: const Color(0xFF1a1a2e),
                borderRadius: BorderRadius.circular(8),
                border: Border.all(color: Colors.grey),
              ),
              alignment: Alignment.center,
              clipBehavior: Clip.antiAlias,
              child: child,
            ),
          ],
        ),
      ),
    );
  }}
}}
"""
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(main)

if __name__ == '__main__':
    generate_main_dart(sys.argv[1], sys.argv[2])
