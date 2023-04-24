package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq()
}

// 统计词频
func wordfreq() {
	hash := make(map[string]int, 20)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		hash[input.Text()]++
	}
	fmt.Printf("%#v", hash)
}
