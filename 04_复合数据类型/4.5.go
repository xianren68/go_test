// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
package main

import "fmt"

func main() {
	str := []string{"11", "11", "22", "23", "24", "24", "37"}
	deduplication(str)
	fmt.Println(str) // [11 22 23 24 37  ]
}

// 原地消除重复值函数
func deduplication(str []string) {
	for i := 1; i < len(str); {
		if str[i] == str[i-1] {
			copy(str[i:], str[i+1:])
			// 最后一位设为空串
			str[len(str)-1] = ""
		}
		i++
	}
}
