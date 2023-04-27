package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

// 网址
const Base_url = "https://api.github.com/search/issues"

// 用户
type User struct {
	// 用户名
	Login   string `json:login`
	HTMLURL string `json:html_url`
}

type Issue struct {
	Number  int
	HTMLURL string `json:"html_url"`
	Title   string
	State   string
	User    *User
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// 模板
var temp = `
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
`

func WebServer() {
	http.HandleFunc("/", Handle)
	// 监听端口
	http.ListenAndServe("127.0.0.1:5000", nil)
}
func Handle(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	var report = template.New("report")
	report, err := report.Parse(temp)
	// 结构体
	var issue = &IssuesSearchResult{}
	if err != nil {
		fmt.Fprintln(w, err)
	}
	// 获取参数
	q := r.URL.Query().Get("key")
	if q == "" {
		fmt.Fprintln(w, "请输入要查询的内容")
		return
	}
	resp, err := http.Get(Base_url + "?q=" + q)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintln(w, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintln(w, resp.StatusCode)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(issue); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	resp.Body.Close()
	report.Execute(w, issue)
}
