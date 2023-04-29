package main

import (
	"fmt"
	"os"

	"5.2/test"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	count := make(map[string]int)
	test.Visit(count, doc)
	fmt.Println("%#v", count)

}
