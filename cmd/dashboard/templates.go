package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

// ctxTemplates 缓存加载的 prompt 模板
var ctxTemplates map[string]*template.Template

type ctxTemplatesConfig struct {
	Templates map[string]string `yaml:"templates"`
}

// loadCtxTemplates 从 prompts/ctx-templates.yaml 加载所有模板
func loadCtxTemplates() error {
	data, err := os.ReadFile("prompts/ctx-templates.yaml")
	if err != nil {
		return fmt.Errorf("read ctx-templates.yaml: %w", err)
	}

	var cfg ctxTemplatesConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("parse ctx-templates.yaml: %w", err)
	}

	ctxTemplates = make(map[string]*template.Template)
	for name, tmplStr := range cfg.Templates {
		tmpl, err := template.New(name).Parse(tmplStr)
		if err != nil {
			return fmt.Errorf("parse template %s: %w", name, err)
		}
		ctxTemplates[name] = tmpl
	}

	return nil
}

// renderCtxTemplate 渲染指定名称的模板
func renderCtxTemplate(name string, data map[string]interface{}) (string, error) {
	tmpl, ok := ctxTemplates[name]
	if !ok {
		return "", fmt.Errorf("template not found: %s", name)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("render template %s: %w", name, err)
	}
	return buf.String(), nil
}

