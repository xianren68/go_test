package main

import (
	"os"
	"strings"

	"4.12/test"
)

func main() {
	// 获取离线列表
	// test.GetSaveInfo()
	key := os.Args[1:]
	test.ResLink(strings.Join(key, " "))

}
