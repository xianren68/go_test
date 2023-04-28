package test

import (
	"golang.org/x/net/html"
)

// 遍历节点树
func Visit(links []string, n *html.Node) []string {
	// 如果当前节点为空，递归结束
	if n == nil {
		return links
	}
	// 寻找元素节点，并且是a标签的元素节点
	if n.Type == html.ElementNode && n.Data == "a" {
		// 寻找a标签的链接属性
		for _, val := range n.Attr {
			if val.Key == "href" {
				links = append(links, val.Val)
			}
		}
	}
	// 递归
	if n.FirstChild != nil {
		// 进入子节点
		links = Visit(links, n.FirstChild)
	}
	// 继续下一个兄弟节点
	links = Visit(links, n.NextSibling)
	return links
}
