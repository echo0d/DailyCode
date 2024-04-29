package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/go-github/v35/github"
	"github.com/manifoldco/promptui"
	"golang.org/x/oauth2"
)

func main() {
	// 设置GitHub认证信息
	token := "YOUR_GITHUB_TOKEN"
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	// 包"github.com/google/go-github/v35/github"是一个用于与GitHub API进行交互的Go语言库。用于管理GitHub上的仓库、问题、拉取请求、用户等。
	client := github.NewClient(tc)

	// 提示用户选择操作
	action := promptForAction()
	switch action {
	case "Create":
		createIssue(client)
	case "Read":
		readIssue(client)
	case "Update":
		updateIssue(client)
	case "Close":
		closeIssue(client)
	default:
		fmt.Println("Invalid action")
	}
}

// 提示用户选择操作
func promptForAction() string {
	// 包"github.com/manifoldco/promptui"是一个用于在命令行应用程序中创建交互式提示的Go语言库。 该库提供了两种主要的输入模式：
	// Prompt：提供单行输入框，支持可选的实时验证、确认和输入掩码。
	// Select：提供选项列表供用户选择，支持分页、搜索、详细视图和自定义模板。

	prompt := promptui.Select{
		Label: "Select an action",
		Items: []string{"Create", "Read", "Update", "Close"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// 创建GitHub issue
func createIssue(client *github.Client) {
	owner := promptForInput("Owner")
	repo := promptForInput("Repository")
	title := promptForInput("Title")
	body := promptForEditorInput()

	issueRequest := &github.IssueRequest{
		Title: &title,
		Body:  &body,
	}

	issue, _, err := client.Issues.Create(context.Background(), owner, repo, issueRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue created: %s\n", *issue.HTMLURL)
}

// 读取GitHub issue
func readIssue(client *github.Client) {
	owner := promptForInput("Owner")
	repo := promptForInput("Repository")
	numberStr := promptForInput("Issue number")

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Fatal(err)
	}

	issue, _, err := client.Issues.Get(context.Background(), owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue details:\nTitle: %s\nBody: %s\n", *issue.Title, *issue.Body)
}

// 更新GitHub issue
func updateIssue(client *github.Client) {
	owner := promptForInput("Owner")
	repo := promptForInput("Repository")
	numberStr := promptForInput("Issue number")

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Fatal(err)
	}

	body := promptForEditorInput()

	issueRequest := &github.IssueRequest{
		Body: &body,
	}

	_, _, err = client.Issues.Edit(context.Background(), owner, repo, number, issueRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Issue updated")
}

// 关闭GitHub issue
func closeIssue(client *github.Client) {
	owner := promptForInput("Owner")
	repo := promptForInput("Repository")
	numberStr := promptForInput("Issue number")

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = client.Issues.Edit(context.Background(), owner, repo, number, &github.IssueRequest{State: github.String("closed")})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Issue closed")
}

// 提示用户输入
func promptForInput(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// 打开编辑器获取输入
func promptForEditorInput() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "notepad" // 默认使用nano编辑器
	}

	filePath := "issue.txt"
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(content))
}
