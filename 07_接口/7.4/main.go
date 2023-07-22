package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// 定义一个html结构体
type HtmlReader struct {
	html []byte
	i    int
}

// 跟着strings.NewReader写的
func (h *HtmlReader) Read(p []byte) (n int, err error) {
	// 存储的索引大于byte数组，表示都读完了
	if h.i >= len(h.html) {
		return 0, io.EOF
	}
	// 复制p长度/剩余长度的字节到p
	n = copy(p, h.html[h.i:])
	// 变化索引
	h.i += n
	return
}

// NewReader，将html转化为reader,以供解析器解析
func NewReader(html string) *HtmlReader {
	return &HtmlReader{html: []byte(html), i: 0}
}

func main() {

	doc, err := html.Parse(NewReader(`
	<!DOCTYPE html>
	<html lang="en">
		<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <title>Document</title>
	</head>
	<body>
	   hahaha
	</body>
	</html>
		`))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v", doc)
}
