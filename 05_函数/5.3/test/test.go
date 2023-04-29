package test

import (
	"golang.org/x/net/html"
)

func Visit(n *html.Node) string {
	str := ""
	// 如果当前节点为空，递归结束
	if n == nil {
		return str
	}
	// 文本节点
	if n.Type == html.TextNode {
		str += n.Data
	}
	// 不是style和script才递归它们的子节点
	if n.Data != "style" && n.Data != "script" {
		// 进入子节点
		str += Visit(n.FirstChild)
	}
	// 进入兄弟节点
	str += Visit(n.NextSibling)
	return str
}
