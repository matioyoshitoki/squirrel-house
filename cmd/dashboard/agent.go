package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// agentTools 返回指定 agent 类型所需的工具列表
func agentTools(agentType string) []string {
	switch agentType {
	case "developer", "rework":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:WriteFile",
			"kimi_cli.tools.file:StrReplaceFile",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.web:SearchWeb",
			"kimi_cli.tools.web:FetchURL",
			"kimi_cli.tools.todo:SetTodoList",
			"kimi_cli.tools.ask_user:AskUserQuestion",
			"kimi_cli.tools.agent:Agent",
			"kimi_cli.tools.background:TaskList",
			"kimi_cli.tools.background:TaskOutput",
			"kimi_cli.tools.background:TaskStop",
		}
	case "reviewer":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:WriteFile",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.web:FetchURL",
			"kimi_cli.tools.file:ReadMediaFile",
			"kimi_cli.tools.todo:SetTodoList",
		}
	case "doc-gardener":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:WriteFile",
			"kimi_cli.tools.file:StrReplaceFile",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.web:SearchWeb",
			"kimi_cli.tools.web:FetchURL",
			"kimi_cli.tools.todo:SetTodoList",
			"kimi_cli.tools.ask_user:AskUserQuestion",
			"kimi_cli.tools.background:TaskList",
			"kimi_cli.tools.background:TaskOutput",
			"kimi_cli.tools.background:TaskStop",
		}
	case "pr-doc-gardener":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.web:FetchURL",
			"kimi_cli.tools.todo:SetTodoList",
		}
	case "ui-designer":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:WriteFile",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.web:FetchURL",
			"kimi_cli.tools.file:ReadMediaFile",
			"kimi_cli.tools.todo:SetTodoList",
		}
	case "product-manager":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:WriteFile",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.web:FetchURL",
			"kimi_cli.tools.todo:SetTodoList",
			"kimi_cli.tools.ask_user:AskUserQuestion",
		}
	case "e2e-tester":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:WriteFile",
			"kimi_cli.tools.file:StrReplaceFile",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.file:ReadMediaFile",
			"kimi_cli.tools.todo:SetTodoList",
			"kimi_cli.tools.ask_user:AskUserQuestion",
		}
	case "prompt-evolution":
		return []string{
			"kimi_cli.tools.file:ReadFile",
			"kimi_cli.tools.file:WriteFile",
			"kimi_cli.tools.file:StrReplaceFile",
			"kimi_cli.tools.file:Glob",
			"kimi_cli.tools.file:Grep",
			"kimi_cli.tools.shell:Shell",
			"kimi_cli.tools.todo:SetTodoList",
		}
	default:
		return agentTools("developer")
	}
}

// promptTemplateData 为 prompt 模板提供的数据
type promptTemplateData struct {
	Project
	HasEnvVars   bool
	HasDesignCheck bool
	HasUI        bool
	DeliveryChain string
}

// buildPromptData 根据项目配置构建模板数据
func buildPromptData(p Project) promptTemplateData {
	d := promptTemplateData{Project: p, HasEnvVars: true}

	parts := []string{"DB", "API"}
	if p.TechStack.Mobile != "" {
		parts = append(parts, "Mobile")
	}
	if p.TechStack.Frontend != "" {
		parts = append(parts, "Frontend")
	}
	parts = append(parts, "Docs")
	d.DeliveryChain = strings.Join(parts, " + ")

	d.HasUI = p.TechStack.Mobile != "" || p.TechStack.Frontend != ""
	d.HasDesignCheck = d.HasUI

	if d.DefaultBranch == "" {
		d.DefaultBranch = "main"
	}
	return d
}

// renderPrompt 渲染指定 agent 类型的 prompt 模板
func renderPrompt(agentType string, p Project) (string, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var templateFile string
	switch agentType {
	case "developer", "rework":
		templateFile = filepath.Join(baseDir, "prompts", "developer.md")
	case "e2e-tester":
		templateFile = filepath.Join(baseDir, "prompts", "e2e-tester.md")
	case "reviewer":
		templateFile = filepath.Join(baseDir, "prompts", "reviewer.md")
	case "ui-designer":
		templateFile = filepath.Join(baseDir, "prompts", "ui-designer.md")
	case "doc-gardener":
		templateFile = filepath.Join(baseDir, "prompts", "doc-gardener.md")
	case "pr-doc-gardener":
		templateFile = filepath.Join(baseDir, "prompts", "pr-doc-gardener.md")
	case "product-manager":
		templateFile = filepath.Join(baseDir, "prompts", "product-manager.md")
	case "prompt-evolution":
		templateFile = filepath.Join(baseDir, "prompts", "prompt-evolution.md")
	default:
		templateFile = filepath.Join(baseDir, "prompts", "developer.md")
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", fmt.Errorf("parse prompt template %s: %w", templateFile, err)
	}

	data := buildPromptData(p)
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute prompt template: %w", err)
	}
	prompt := buf.String()
	
	// 根据平台自动适配 prompt 中的 CLI 命令
	platform := getPlatform(p.Path)
	prompt = adaptPromptForPlatform(prompt, platform)
	
	return prompt, nil
}

// adaptPromptForPlatform 根据平台类型自动替换 prompt 中的 CLI 命令和文案
func adaptPromptForPlatform(prompt string, platform Platform) string {
	if platform.Name() == "github" {
		return prompt
	}
	// GitLab 适配映射
	replacements := []struct{ old, new string }{
		{"gh auth status", "glab auth status"},
		{"gh pr diff", "glab mr diff"},
		{"gh pr view", "glab mr view"},
		{"gh pr comment", "glab mr note"},
		{"gh pr create", "glab mr create"},
		{"gh pr edit", "glab mr update"},
		{"gh pr merge", "glab mr merge"},
		{"gh pr update-branch", "glab mr rebase"},
		{"gh issue create", "glab issue create"},
		{"gh issue view", "glab issue view"},
		{"gh issue list", "glab issue list"},
		{"GitHub CLI", "GitLab CLI"},
		{"GitHub Issue", "GitLab Issue"},
		{"GitHub PR", "GitLab MR"},
		{"GitHub", "GitLab"},
		{" PullRequest", " MergeRequest"},
	}
	for _, r := range replacements {
		prompt = strings.ReplaceAll(prompt, r.old, r.new)
	}
	return prompt
}

// prepareAgentConfig 生成临时 agent 配置文件和 prompt 文件，返回 agent YAML 路径
func prepareAgentConfig(agentType string) string {
	projectMutex.RLock()
	idx := currentProjectIndex
	projects := config.Projects
	projectMutex.RUnlock()

	if idx < 0 || idx >= len(projects) {
		idx = 0
	}
	if len(projects) == 0 {
		return ""
	}
	project := projects[idx]

	baseDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	// 渲染 prompt
	promptContent, err := renderPrompt(agentType, project)
	if err != nil {
		return ""
	}

	timestamp := fmt.Sprintf("%d", os.Getpid())
	promptFile := filepath.Join(os.TempDir(), fmt.Sprintf("dashboard-prompt-%s-%s-%s.md", project.Name, agentType, timestamp))
	if err := os.WriteFile(promptFile, []byte(promptContent), 0644); err != nil {
		return ""
	}

	// 生成临时 agent YAML
	agentFile := filepath.Join(os.TempDir(), fmt.Sprintf("dashboard-agent-%s-%s-%s.yaml", project.Name, agentType, timestamp))

	tools := agentTools(agentType)
	toolLines := make([]string, 0, len(tools))
	for _, t := range tools {
		toolLines = append(toolLines, fmt.Sprintf("    - \"%s\"", t))
	}

	subagentDir := filepath.Join(baseDir, "agents", "subagents")
	yamlContent := fmt.Sprintf(`version: 1
agent:
  name: dashboard-generated-%s
  extend: default

  loop_control:
    max_steps_per_turn: 10000

  system_prompt_path: %s

  tools:
%s

  subagents:
    coder:
      path: %s
      description: "通用软件工程任务：读写文件、运行命令、搜索代码"
    explore:
      path: %s
      description: "快速只读代码探索：搜索、阅读、总结，不修改代码"
    plan:
      path: %s
      description: "实现规划与架构设计：分析文件、制定方案"
    reviewer:
      path: %s
      description: "A2A 代码审查：检查代码规范、架构合规、测试覆盖"
`, agentType, promptFile, strings.Join(toolLines, "\n"),
		filepath.Join(subagentDir, "coder.yaml"),
		filepath.Join(subagentDir, "explore.yaml"),
		filepath.Join(subagentDir, "plan.yaml"),
		filepath.Join(subagentDir, "reviewer.yaml"))

	if err := os.WriteFile(agentFile, []byte(yamlContent), 0644); err != nil {
		return ""
	}
	return agentFile
}
