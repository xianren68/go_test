package test

import (
	"golang.org/x/net/html"
)

func Visit(count map[string]int, n *html.Node) {
	// 如果当前节点为空，递归结束
	if n == nil {
		return
	}
	// 遍历，并记录至hash表中
	if n.Type == html.ElementNode {
		// 对应次数++
		count[n.Data]++
	}
	// 进入子节点
	Visit(count, n.FirstChild)
	// 进入兄弟节点
	Visit(count, n.NextSibling)
}
