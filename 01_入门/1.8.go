// 1.8  修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
package main

import (
	"fmt"
	"io"
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
	// 将返回体复制到返回体
	_, err = io.Copy(os.Stdout, res.Body)
	// 关闭读入流
	res.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
