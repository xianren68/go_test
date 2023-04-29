package main

import (
	"fmt"
	"os"

	"5.4/test"
	"golang.org/x/net/html"
)

func main() {
	// 接收输入流的html，并解析
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline:%v\n", err)
		os.Exit(1)
	}
	links := make([]string, 0)
	links = test.Visit(links, doc)
	for _, val := range links {
		fmt.Println(val)
	}
}
