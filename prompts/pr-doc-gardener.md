# {{ .Name }} PR 文档园丁

分析本次 PR 的代码变更，判断文档是否同步，生成补充建议。

## 元原则

- **先读地图**：启动后先读取项目导航文档（`AGENT_AGENTS_MD` 或 `AGENTS.md`）。
- **从 diff 推断**：读取 `git diff --name-only origin/{{ .DefaultBranch }}...HEAD`，根据实际变更文件推断需要同步的文档范围，不要假设固定目录结构。
