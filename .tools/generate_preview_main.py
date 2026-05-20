#!/usr/bin/env python3
import os
import re
import sys

def extract_types(content):
    """提取文件中定义的所有类型 {name: kind}，kind 为 class/enum"""
    types = {}
    for m in re.finditer(r'class\s+(\w+)', content):
        types[m.group(1)] = 'class'
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

def mock_value(param_type, param_name, known_types, all_type_defs):
    pt = param_type.strip()
    if pt in all_type_defs:
        td = all_type_defs[pt]
        if isinstance(td, tuple) and td[0] == 'enum':
            return f'{pt}.{td[1]}'
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
        return f'{pt}()'
    return f'{pt}()'

def build_instance(class_name, params, known_types, all_type_defs, has_scaffold=False):
    req_params = [p for p in params if p['required']]
    if not req_params:
        code = f'{class_name}()'
    else:
        args = []
        has_unknown = False
        for p in req_params:
            val = mock_value(p['type'], p['name'], known_types, all_type_defs)
            if p['type'] not in {'String', 'int', 'double', 'bool', 'Key'} and not any(x in p['type'] for x in ['Callback', 'VoidCallback', 'Gesture', 'ValueChanged', 'List', 'Map']):
                td = all_type_defs.get(p['type'])
                if not (isinstance(td, tuple) and td[0] == 'enum'):
                    has_unknown = True
            args.append(f'{p["name"]}: {val}')
        code = f'{class_name}({", ".join(args)})'
    if has_scaffold:
        code = f'Container(height: 400, decoration: BoxDecoration(borderRadius: BorderRadius.circular(8), border: Border.all(color: Colors.grey)), clipBehavior: Clip.antiAlias, child: {code})'
    return code, True

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
            types = extract_types(content)
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

    for name, filepath in widget_files.items():
        rel = os.path.relpath(filepath, preview_dir)
        import_path = 'preview/' + rel.replace('\\', '/')
        imports.add(f"import '{import_path}';")

    for name, filepath in widget_files.items():
        widgets = extract_widgets(filepath, all_type_defs)
        for w in widgets:
            if w['name'] != name:
                continue
            code, is_safe = build_instance(w['name'], w['params'], known_types, all_type_defs, w.get('has_scaffold', False))
            instances.append(f"            _buildCard('{w['name']}', {code}),")

    if not instances:
        instances.append("            const Center(child: Text('未找到可预览的 Widget')),")

    main = f"""import 'package:flutter/material.dart';
{chr(10).join(sorted(imports))}

void main() {{
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
              padding: const EdgeInsets.all(12),
              decoration: BoxDecoration(
                color: const Color(0xFF1a1a2e),
                borderRadius: BorderRadius.circular(8),
              ),
              child: child is Container ? child : Center(child: child),
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
