package main

import (
	"fmt"
	"os"

	"4.10/test"
)

func main() {
	query := os.Args[1:]
	data, err := test.SearchIssure(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	test.Process(data)
}
