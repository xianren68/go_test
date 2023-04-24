// 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	charCount()
}
func charCount() {
	// 创建hash表
	hash := make(map[string]int, 4)
	input := bufio.NewReader(os.Stdin)
	for {
		r, _, err := input.ReadRune()
		if err == io.EOF { // 输入ctrl+z
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		t := whatType(r)
		hash[t]++
	}
	fmt.Printf("%#v", hash)
}

// 判断一个unicode码点为什么类型
func whatType(r rune) string {
	if unicode.IsNumber(r) {
		return "number"
	}
	if unicode.Is(unicode.Han, r) {
		return "中文"
	}
	if r >= 65 && r <= 122 {
		return "English"
	}

	return "any"
}
