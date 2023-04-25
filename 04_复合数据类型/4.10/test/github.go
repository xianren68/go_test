package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 网址
const Base_url = "https://api.github.com/search/issues"

// 用户
type User struct {
	// 用户名
	Login string `json:login`
	// 用户id
	Id int `json:id`
	// 用户主页
	HTMLURL string `json:html_url`
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

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// 获取数据
func SearchIssure(term []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(term, " "))
	resp, err := http.Get(Base_url + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("状态码错误：%d", resp.StatusCode)
	}
	var result = &IssuesSearchResult{}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil

}

// 处理数据
func Process(data *IssuesSearchResult) {
	// 创建一个hash表
	hash := make(map[string][]*Issue, 3)
	hash["不到一月"] = make([]*Issue, 0)
	hash["不到一年"] = make([]*Issue, 0)
	hash["超过一年"] = make([]*Issue, 0)
	// 遍历数据
	for _, val := range data.Items {
		// 获取现在的时间
		now := time.Now()
		year := now.Year()
		month := now.Month()
		// 创建时间
		create_year := val.CreatedAt.Year()
		create_month := val.CreatedAt.Month()
		if year == create_year && month == create_month {
			hash["不到一月"] = append(hash["不到一月"], val)
		} else if year == create_year {
			hash["不到一年"] = append(hash["不到一年"], val)
		} else {
			hash["超过一年"] = append(hash["超过一年"], val)
		}
	}
	fmt.Printf("数据总条数，%v\n", data.TotalCount)
	for key, val := range hash {
		fmt.Printf("%s:\n", key)
		for _, v := range val {
			fmt.Printf("#%-5d %9.9s %.55s %9.9v\n",
				v.Number, v.User.Login, v.Title, v.CreatedAt)
		}
	}
}
