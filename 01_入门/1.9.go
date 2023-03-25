// 1.9 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 获取命令行参数
	url := os.Args[1]
	// 判断url前缀是否为http/https
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		// 不是，直接添加前缀
		url = "http://" + url
	}
	// 发送请求
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// 打印请求状态码
	fmt.Println(res.Status)
}
