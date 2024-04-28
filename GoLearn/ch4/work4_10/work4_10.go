package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// NOTE: 定义合适的类型和常量

const IssuesURL = "https://api.github.com/search/issues"

// NOTE: 和前面一样，即使对应的JSON对象名是小写字母，每个结构体的成员名也是声明为大写字母开头的。
// NOTE: 有些JSON成员名字和Go结构体成员名字并不相同，因此需要Go语言结构体成员Tag来指定对应的JSON名字。

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues 函数发出一个HTTP请求，然后解码返回的JSON格式的结果。
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// 用户提供的查询条件可能包含类似?和&之类的特殊字符，为了避免对URL造成冲突，用url.QueryEscape来对查询中的特殊字符进行转义操作
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	// 使用了基于流式的解码器json.Decoder，它可以从一个输入流解码JSON数据
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total count %d issues: \n", result.TotalCount)
	now := time.Now()
	aMonthAgo := now.AddDate(0, -1, 0)
	aYearAgo := now.AddDate(-1, 0, 0)

	var lessMonthIssues []*Issue
	var lessYearMoreThanMonthIssues []*Issue
	var moreThanYearIssues []*Issue

	for _, item := range result.Items {
		if item.CreatedAt.After(aMonthAgo) {
			lessMonthIssues = append(lessMonthIssues, item)
		} else if item.CreatedAt.After(aYearAgo) {
			lessYearMoreThanMonthIssues = append(lessYearMoreThanMonthIssues, item)
		} else {
			moreThanYearIssues = append(moreThanYearIssues, item)
		}
	}
	// NOTE: 同一个issue对应多个参数，例如Title HTMLURL CreatedAt User.Login
	for _, issue := range lessMonthIssues {
		fmt.Printf("最近一月:\t%s\t%s\t%s\t%s\n", issue.Title, issue.HTMLURL, issue.CreatedAt, issue.User.Login)
	}

	for _, issue := range lessYearMoreThanMonthIssues {
		fmt.Printf("最近一年:\t%s\t%s\t%s\t%s\n", issue.Title, issue.HTMLURL, issue.CreatedAt, issue.User.Login)
	}

	for _, issue := range moreThanYearIssues {
		fmt.Printf("一年以前:\t%s\t%s\t%s\t%s\n", issue.Title, issue.HTMLURL, issue.CreatedAt, issue.User.Login)
	}
}
