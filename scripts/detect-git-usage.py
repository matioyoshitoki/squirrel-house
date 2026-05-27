#!/usr/bin/env python3
"""
精确检测 agent 日志中的本地 git 命令调用。
只统计 FunctionBody(name='Shell'...) 的 arguments 中的真实 git 命令，
避免误报 reasoning/prompt/result 中的声明文本。
"""
import re
import sys

def main():
    if len(sys.argv) < 2:
        print(0)
        return
    log_file = sys.argv[1]
    git_cmds = sys.argv[2] if len(sys.argv) > 2 else "status|diff|log|show|blame|checkout|add|commit|push"
    try:
        with open(log_file) as f:
            content = f.read()
        # 精确提取 Shell 工具调用的 arguments
        matches = re.findall(r"name='Shell'.*?arguments='(.*?)'\s*\)", content, re.DOTALL)
        count = 0
        pattern = re.compile(rf'\bgit\s+({git_cmds})\b', re.IGNORECASE)
        for args in matches:
            count += len(pattern.findall(args))
        print(count)
    except Exception:
        print(0)

if __name__ == "__main__":
    main()
