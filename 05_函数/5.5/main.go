package main

import (
	"fmt"
	"os"

	"5.5/test"
)

func main() {
	url := os.Args[1]
	fmt.Println(test.CountWordsAndImages(url))
}
