package main

import (
	"ch4/github"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	//启动一个web服务器
	http.HandleFunc("/", handle)
	fmt.Printf("visit http://localhost:8888/?q=xxx\n")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

// 使用q参数查询
func handle(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	if query == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<html><body>You need a param 'q'</body></html>")
		return
	}
	var result *github.IssuesSearchResult
	var keywords = strings.Split(query, " ")
	result, _ = github.SearchIssues(keywords)

	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
	err := issueList.Execute(w, result)
	if err != nil {
		return
	}
}
