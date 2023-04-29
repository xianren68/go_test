package main

import (
	"fmt"
	"os"

	"5.3/test"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	str := test.Visit(doc)
	fmt.Printf("%s", str)
}
