package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// =============================================================================
// Platform 抽象层：统一封装 GitHub / GitLab 操作
// =============================================================================
// 通过 git remote URL 自动推断平台，支持配置覆盖。
// 所有 exec.Command("gh", ...) 的调用已迁移至此。
// =============================================================================

// MergeRequest 通用 MR/PR 模型（平台无关）
type MergeRequest struct {
	IID          int
	Title        string
	State        string
	SourceBranch string
	TargetBranch string
	Body         string
	URL          string
	CreatedAt    string
	Author       string
}

// Platform 定义 Git 平台操作接口
type Platform interface {
	Name() string      // "github" | "gitlab"
	CLIName() string   // "gh" | "glab"
	AuthStatus() error

	// MR / PR
	ListOpenMRs() ([]MergeRequest, error)
	ListMergedMRs() ([]MergeRequest, error)
	ViewMR(iid int) (*MergeRequest, error)
	FindMRByBranch(branch string) (int, error)
	GetMRDiffFiles(iid int) ([]string, error)
	CreateMRComment(iid int, body string) error
	UpdateLastMRComment(iid int, body string) error
	MergeMR(iid int) error
	RebaseMR(iid int) error
	CreateMR(sourceBranch, targetBranch, title, body string) (int, error)
	AddMRLabel(iid int, label string) error

	// Issue
	ListOpenIssues() ([]Issue, error)
	ViewIssue(iid int) (*Issue, error)
	CreateIssue(title, body string, labels []string) (int, error)
}

// =============================================================================
// 平台实例缓存与管理
// =============================================================================

var platformCache sync.Map // key=projectPath, value=Platform

// detectPlatform 通过 git remote URL 自动推断平台类型
func detectPlatform(projectPath string) string {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = projectPath
	out, err := cmd.Output()
	if err != nil {
		log.Printf("[platform] 无法获取 %s 的 git remote: %v, 默认使用 github", projectPath, err)
		return "github"
	}
	url := strings.TrimSpace(string(out))
	lower := strings.ToLower(url)

	if strings.Contains(lower, "gitlab") {
		return "gitlab"
	}

	// 从 host 推断自托管 GitLab
	host := extractGitHost(url)
	if strings.Contains(strings.ToLower(host), "gitlab") {
		return "gitlab"
	}

	return "github"
}

func extractGitHost(url string) string {
	// https://host.com/path
	if strings.HasPrefix(url, "https://") {
		rest := strings.TrimPrefix(url, "https://")
		parts := strings.SplitN(rest, "/", 2)
		return parts[0]
	}
	if strings.HasPrefix(url, "http://") {
		rest := strings.TrimPrefix(url, "http://")
		parts := strings.SplitN(rest, "/", 2)
		return parts[0]
	}
	// git@host.com:path
	if strings.HasPrefix(url, "git@") {
		rest := strings.TrimPrefix(url, "git@")
		parts := strings.SplitN(rest, ":", 2)
		return parts[0]
	}
	return ""
}

// getProjectPlatform 获取项目的平台类型（先检查配置覆盖，再自动推断）
func getProjectPlatform(projectPath string) string {
	absProject, _ := filepath.Abs(projectPath)

	projectMutex.RLock()
	for _, p := range config.Projects {
		absConfig, _ := filepath.Abs(p.Path)
		if absConfig == absProject && p.Platform != "" {
			projectMutex.RUnlock()
			return p.Platform
		}
	}
	projectMutex.RUnlock()

	return detectPlatform(projectPath)
}

// getPlatform 获取项目的 Platform 实例（带缓存）
func getPlatform(projectPath string) Platform {
	if p, ok := platformCache.Load(projectPath); ok {
		return p.(Platform)
	}

	platformType := getProjectPlatform(projectPath)
	var p Platform
	switch platformType {
	case "gitlab":
		p = &gitlabPlatform{projectPath: projectPath}
	default:
		p = &githubPlatform{projectPath: projectPath}
	}

	platformCache.Store(projectPath, p)
	log.Printf("[platform] 项目 %s 使用平台: %s", projectPath, platformType)
	return p
}

// invalidatePlatformCache 清除指定项目的平台缓存（切换项目后调用）
func invalidatePlatformCache(projectPath string) {
	platformCache.Delete(projectPath)
}

// =============================================================================
// GitHub 平台实现
// =============================================================================

type githubPlatform struct {
	projectPath string
}

func (g *githubPlatform) Name() string    { return "github" }
func (g *githubPlatform) CLIName() string { return "gh" }

func (g *githubPlatform) run(args ...string) ([]byte, error) {
	cmd := exec.Command("gh", args...)
	cmd.Dir = g.projectPath
	return cmd.CombinedOutput()
}

func (g *githubPlatform) AuthStatus() error {
	_, err := g.run("auth", "status")
	return err
}

func (g *githubPlatform) listPRs(state string) ([]MergeRequest, error) {
	output, err := g.run("pr", "list", "--state", state, "--json", "number,title,state,baseRefName,headRefName,url,createdAt,author")
	if err != nil {
		return nil, fmt.Errorf("list PRs: %w, output: %s", err, string(output))
	}
	var raw []struct {
		Number      int    `json:"number"`
		Title       string `json:"title"`
		State       string `json:"state"`
		BaseRefName string `json:"baseRefName"`
		HeadRefName string `json:"headRefName"`
		URL         string `json:"url"`
		CreatedAt   string `json:"createdAt"`
		Author      struct {
			Login string `json:"login"`
		} `json:"author"`
	}
	if err := json.Unmarshal(output, &raw); err != nil {
		return nil, err
	}
	result := make([]MergeRequest, len(raw))
	for i, r := range raw {
		result[i] = MergeRequest{
			IID: r.Number, Title: r.Title, State: r.State,
			SourceBranch: r.HeadRefName, TargetBranch: r.BaseRefName,
			Body: "", URL: r.URL, CreatedAt: r.CreatedAt, Author: r.Author.Login,
		}
	}
	return result, nil
}

func (g *githubPlatform) ListOpenMRs() ([]MergeRequest, error) {
	return g.listPRs("open")
}

func (g *githubPlatform) ListMergedMRs() ([]MergeRequest, error) {
	return g.listPRs("merged")
}

func (g *githubPlatform) ViewMR(iid int) (*MergeRequest, error) {
	output, err := g.run("pr", "view", strconv.Itoa(iid), "--json", "number,title,state,baseRefName,headRefName,body,url,createdAt,author")
	if err != nil {
		return nil, fmt.Errorf("view PR %d: %w, output: %s", iid, err, string(output))
	}
	var r struct {
		Number      int    `json:"number"`
		Title       string `json:"title"`
		State       string `json:"state"`
		BaseRefName string `json:"baseRefName"`
		HeadRefName string `json:"headRefName"`
		Body        string `json:"body"`
		URL         string `json:"url"`
		CreatedAt   string `json:"createdAt"`
		Author      struct {
			Login string `json:"login"`
		} `json:"author"`
	}
	if err := json.Unmarshal(output, &r); err != nil {
		return nil, err
	}
	return &MergeRequest{
		IID: r.Number, Title: r.Title, State: r.State,
		SourceBranch: r.HeadRefName, TargetBranch: r.BaseRefName,
		Body: r.Body, URL: r.URL, CreatedAt: r.CreatedAt, Author: r.Author.Login,
	}, nil
}

func (g *githubPlatform) FindMRByBranch(branch string) (int, error) {
	output, err := g.run("pr", "list", "--head", branch, "--json", "number", "--state", "open")
	if err != nil {
		return 0, fmt.Errorf("find PR by branch %s: %w, output: %s", branch, err, string(output))
	}
	var prs []struct {
		Number int `json:"number"`
	}
	if err := json.Unmarshal(output, &prs); err != nil {
		return 0, err
	}
	if len(prs) == 0 {
		return 0, nil
	}
	return prs[0].Number, nil
}

func (g *githubPlatform) GetMRDiffFiles(iid int) ([]string, error) {
	output, err := g.run("pr", "diff", strconv.Itoa(iid), "--name-only")
	if err != nil {
		return nil, fmt.Errorf("PR %d diff: %w, output: %s", iid, err, string(output))
	}
	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(files) == 1 && files[0] == "" {
		return []string{}, nil
	}
	return files, nil
}

func (g *githubPlatform) CreateMRComment(iid int, body string) error {
	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("gh-comment-%d-%d.md", iid, time.Now().Unix()))
	if err := os.WriteFile(tmpFile, []byte(body), 0644); err != nil {
		return err
	}
	defer os.Remove(tmpFile)
	output, err := g.run("pr", "comment", strconv.Itoa(iid), "--body-file", tmpFile)
	if err != nil {
		return fmt.Errorf("comment PR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (g *githubPlatform) UpdateLastMRComment(iid int, body string) error {
	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("gh-comment-%d-%d.md", iid, time.Now().Unix()))
	if err := os.WriteFile(tmpFile, []byte(body), 0644); err != nil {
		return err
	}
	defer os.Remove(tmpFile)
	output, err := g.run("pr", "comment", strconv.Itoa(iid), "--edit-last", "--body-file", tmpFile)
	if err != nil {
		return fmt.Errorf("edit-last comment PR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (g *githubPlatform) MergeMR(iid int) error {
	output, err := g.run("pr", "merge", strconv.Itoa(iid), "--squash", "--delete-branch")
	if err != nil {
		return fmt.Errorf("merge PR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (g *githubPlatform) RebaseMR(iid int) error {
	output, err := g.run("pr", "update-branch", strconv.Itoa(iid))
	if err != nil {
		return fmt.Errorf("update-branch PR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (g *githubPlatform) CreateMR(sourceBranch, targetBranch, title, body string) (int, error) {
	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("gh-pr-body-%d.md", time.Now().Unix()))
	if err := os.WriteFile(tmpFile, []byte(body), 0644); err != nil {
		return 0, err
	}
	defer os.Remove(tmpFile)
	output, err := g.run("pr", "create", "--base", targetBranch, "--head", sourceBranch, "--title", title, "--body-file", tmpFile, "--json", "number", "-q", ".number")
	if err != nil {
		return 0, fmt.Errorf("create PR: %w, output: %s", err, string(output))
	}
	numStr := strings.TrimSpace(string(output))
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("parse PR number: %w (output: %s)", err, numStr)
	}
	return num, nil
}

func (g *githubPlatform) AddMRLabel(iid int, label string) error {
	output, err := g.run("pr", "edit", strconv.Itoa(iid), "--add-label", label)
	if err != nil {
		return fmt.Errorf("add label to PR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (g *githubPlatform) ListOpenIssues() ([]Issue, error) {
	output, err := g.run("issue", "list", "--state", "open", "--json", "number,title,state,labels,body,updatedAt")
	if err != nil {
		return nil, fmt.Errorf("list issues: %w, output: %s", err, string(output))
	}
	var issues []Issue
	if err := json.Unmarshal(output, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (g *githubPlatform) ViewIssue(iid int) (*Issue, error) {
	output, err := g.run("issue", "view", strconv.Itoa(iid), "--json", "number,title,state,labels,body")
	if err != nil {
		return nil, fmt.Errorf("view issue %d: %w, output: %s", iid, err, string(output))
	}
	var issue Issue
	if err := json.Unmarshal(output, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (g *githubPlatform) CreateIssue(title, body string, labels []string) (int, error) {
	args := []string{"issue", "create", "--title", title, "--body", body}
	if len(labels) > 0 {
		args = append(args, "--label", strings.Join(labels, ","))
	}
	args = append(args, "--json", "number", "-q", ".number")
	output, err := g.run(args...)
	if err != nil {
		return 0, fmt.Errorf("create issue: %w, output: %s", err, string(output))
	}
	numStr := strings.TrimSpace(string(output))
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("parse issue number: %w (output: %s)", err, numStr)
	}
	return num, nil
}

// =============================================================================
// GitLab 平台实现
// =============================================================================

type gitlabPlatform struct {
	projectPath string
}

func (gl *gitlabPlatform) Name() string    { return "gitlab" }
func (gl *gitlabPlatform) CLIName() string { return "glab" }

func (gl *gitlabPlatform) run(args ...string) ([]byte, error) {
	cmd := exec.Command("glab", args...)
	cmd.Dir = gl.projectPath
	return cmd.CombinedOutput()
}

func (gl *gitlabPlatform) AuthStatus() error {
	// glab auth status 在有多个 host 配置或 token 被标记为 invalid 时也会返回非零退出码
	// 使用轻量 API 调用验证实际认证状态
	output, err := gl.run("mr", "list", "--per-page", "1", "-F", "json")
	if err != nil {
		return fmt.Errorf("glab API call failed: %w, output: %s", err, string(output))
	}
	return nil
}

func (gl *gitlabPlatform) listMRs(state string) ([]MergeRequest, error) {
	args := []string{"mr", "list", "-F", "json"}
	if state == "merged" {
		args = append(args, "--merged")
	}
	output, err := gl.run(args...)
	if err != nil {
		return nil, fmt.Errorf("list MRs: %w, output: %s", err, string(output))
	}
	var raw []struct {
		IID          int    `json:"iid"`
		Title        string `json:"title"`
		State        string `json:"state"`
		SourceBranch string `json:"source_branch"`
		TargetBranch string `json:"target_branch"`
		Description  string `json:"description"`
		WebURL       string `json:"web_url"`
		CreatedAt    string `json:"created_at"`
		Author       struct {
			Username string `json:"username"`
		} `json:"author"`
	}
	if err := json.Unmarshal(output, &raw); err != nil {
		return nil, err
	}
	result := make([]MergeRequest, len(raw))
	for i, r := range raw {
		result[i] = MergeRequest{
			IID: r.IID, Title: r.Title, State: r.State,
			SourceBranch: r.SourceBranch, TargetBranch: r.TargetBranch,
			Body: r.Description, URL: r.WebURL, CreatedAt: r.CreatedAt, Author: r.Author.Username,
		}
	}
	return result, nil
}

func (gl *gitlabPlatform) ListOpenMRs() ([]MergeRequest, error) {
	return gl.listMRs("")
}

func (gl *gitlabPlatform) ListMergedMRs() ([]MergeRequest, error) {
	return gl.listMRs("merged")
}

func (gl *gitlabPlatform) ViewMR(iid int) (*MergeRequest, error) {
	output, err := gl.run("mr", "view", strconv.Itoa(iid), "-F", "json")
	if err != nil {
		return nil, fmt.Errorf("view MR %d: %w, output: %s", iid, err, string(output))
	}
	var r struct {
		IID          int    `json:"iid"`
		Title        string `json:"title"`
		State        string `json:"state"`
		SourceBranch string `json:"source_branch"`
		TargetBranch string `json:"target_branch"`
		Description  string `json:"description"`
		WebURL       string `json:"web_url"`
		CreatedAt    string `json:"created_at"`
		Author       struct {
			Username string `json:"username"`
		} `json:"author"`
	}
	if err := json.Unmarshal(output, &r); err != nil {
		return nil, err
	}
	return &MergeRequest{
		IID: r.IID, Title: r.Title, State: r.State,
		SourceBranch: r.SourceBranch, TargetBranch: r.TargetBranch,
		Body: r.Description, URL: r.WebURL, CreatedAt: r.CreatedAt, Author: r.Author.Username,
	}, nil
}

func (gl *gitlabPlatform) FindMRByBranch(branch string) (int, error) {
	output, err := gl.run("mr", "list", "--source-branch", branch, "-F", "json")
	if err != nil {
		return 0, fmt.Errorf("find MR by branch %s: %w, output: %s", branch, err, string(output))
	}
	var mrs []struct {
		IID int `json:"iid"`
	}
	if err := json.Unmarshal(output, &mrs); err != nil {
		return 0, err
	}
	if len(mrs) == 0 {
		return 0, nil
	}
	return mrs[0].IID, nil
}

func (gl *gitlabPlatform) GetMRDiffFiles(iid int) ([]string, error) {
	// glab mr diff 输出 patch 格式，解析较复杂。
	// 更可靠的方式：在 worktree 中用 git diff
	output, err := gl.run("mr", "diff", strconv.Itoa(iid))
	if err != nil {
		return nil, fmt.Errorf("MR %d diff: %w, output: %s", iid, err, string(output))
	}
	// 解析 patch 格式：diff --git a/file b/file
	var files []string
	re := regexp.MustCompile(`^diff --git a/(.+) b/`)
	for _, line := range strings.Split(string(output), "\n") {
		if m := re.FindStringSubmatch(line); len(m) > 1 {
			files = append(files, m[1])
		}
	}
	return files, nil
}

func (gl *gitlabPlatform) CreateMRComment(iid int, body string) error {
	output, err := gl.run("mr", "note", strconv.Itoa(iid), "-m", body)
	if err != nil {
		return fmt.Errorf("comment MR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (gl *gitlabPlatform) UpdateLastMRComment(iid int, body string) error {
	// glab 没有 edit-last，需要列出 notes 找到最新的再更新
	listOut, err := gl.run("mr", "note", "list", strconv.Itoa(iid), "-F", "json")
	if err != nil {
		// 没有 notes 时直接新建
		return gl.CreateMRComment(iid, body)
	}
	var notes []struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(listOut, &notes); err != nil || len(notes) == 0 {
		return gl.CreateMRComment(iid, body)
	}
	lastID := notes[len(notes)-1].ID
	output, err := gl.run("mr", "note", "update", strconv.Itoa(iid), strconv.Itoa(lastID), "-m", body)
	if err != nil {
		return fmt.Errorf("update note %d on MR %d: %w, output: %s", lastID, iid, err, string(output))
	}
	return nil
}

func (gl *gitlabPlatform) MergeMR(iid int) error {
	output, err := gl.run("mr", "merge", strconv.Itoa(iid), "-s")
	if err != nil {
		return fmt.Errorf("merge MR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (gl *gitlabPlatform) RebaseMR(iid int) error {
	output, err := gl.run("mr", "rebase", strconv.Itoa(iid))
	if err != nil {
		return fmt.Errorf("rebase MR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (gl *gitlabPlatform) CreateMR(sourceBranch, targetBranch, title, body string) (int, error) {
	output, err := gl.run("mr", "create", "-b", targetBranch, "-s", sourceBranch, "-t", title, "-d", body, "-F", "json")
	if err != nil {
		return 0, fmt.Errorf("create MR: %w, output: %s", err, string(output))
	}
	var result struct {
		IID int `json:"iid"`
	}
	if err := json.Unmarshal(output, &result); err != nil {
		// fallback: 尝试从输出中解析 URL 或其他信息
		return 0, fmt.Errorf("parse create MR response: %w, output: %s", err, string(output))
	}
	return result.IID, nil
}

func (gl *gitlabPlatform) AddMRLabel(iid int, label string) error {
	output, err := gl.run("mr", "update", strconv.Itoa(iid), "-l", label)
	if err != nil {
		return fmt.Errorf("add label to MR %d: %w, output: %s", iid, err, string(output))
	}
	return nil
}

func (gl *gitlabPlatform) ListOpenIssues() ([]Issue, error) {
	output, err := gl.run("issue", "list", "-O", "json")
	if err != nil {
		return nil, fmt.Errorf("list issues: %w, output: %s", err, string(output))
	}
	var raw []struct {
		IID       int      `json:"iid"`
		Title     string   `json:"title"`
		State     string   `json:"state"`
		Labels    []string `json:"labels"`
		Description string `json:"description"`
		UpdatedAt string   `json:"updated_at"`
	}
	if err := json.Unmarshal(output, &raw); err != nil {
		// glab 在没有 issues 时输出文本提示而非 JSON（如 "No open issues match..."），此时返回空数组
		trimmed := strings.TrimSpace(string(output))
		if len(trimmed) > 0 && !strings.HasPrefix(trimmed, "[") && !strings.HasPrefix(trimmed, "{") {
			return []Issue{}, nil
		}
		return nil, err
	}
	result := make([]Issue, len(raw))
	for i, r := range raw {
		labels := make([]Label, len(r.Labels))
		for j, l := range r.Labels {
			labels[j] = Label{Name: l}
		}
		result[i] = Issue{
			Number: r.IID, Title: r.Title, State: r.State,
			Labels: labels, Body: r.Description, UpdatedAt: r.UpdatedAt,
		}
	}
	return result, nil
}

func (gl *gitlabPlatform) ViewIssue(iid int) (*Issue, error) {
	output, err := gl.run("issue", "view", strconv.Itoa(iid), "-F", "json")
	if err != nil {
		return nil, fmt.Errorf("view issue %d: %w, output: %s", iid, err, string(output))
	}
	var r struct {
		IID         int      `json:"iid"`
		Title       string   `json:"title"`
		State       string   `json:"state"`
		Labels      []string `json:"labels"`
		Description string   `json:"description"`
	}
	if err := json.Unmarshal(output, &r); err != nil {
		return nil, err
	}
	labels := make([]Label, len(r.Labels))
	for j, l := range r.Labels {
		labels[j] = Label{Name: l}
	}
	return &Issue{
		Number: r.IID, Title: r.Title, State: r.State,
		Labels: labels, Body: r.Description,
	}, nil
}

func (gl *gitlabPlatform) CreateIssue(title, body string, labels []string) (int, error) {
	args := []string{"issue", "create", "-t", title, "-d", body}
	if len(labels) > 0 {
		args = append(args, "-l", strings.Join(labels, ","))
	}
	output, err := gl.run(args...)
	if err != nil {
		return 0, fmt.Errorf("create issue: %w, output: %s", err, string(output))
	}
	// glab issue create 输出形如：
	// https://gitlab.com/group/project/-/issues/42
	// 或包含 IID
	outStr := string(output)
	re := regexp.MustCompile(`/issues/(\d+)`)
	if m := re.FindStringSubmatch(outStr); len(m) > 1 {
		num, _ := strconv.Atoi(m[1])
		return num, nil
	}
	return 0, fmt.Errorf("无法从输出解析 issue 编号: %s", outStr)
}
