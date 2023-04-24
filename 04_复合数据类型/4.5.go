// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
package main

import "fmt"

func main() {
	str := []string{"11", "11", "22", "23", "24", "24", "37"}
	str = deduplication(str)
	fmt.Println(str) // [11 22 23 24 37]
}

// 原地消除重复值函数
func deduplication(str []string) []string {
	for i := 1; i < len(str); {
		if str[i] == str[i-1] {
			copy(str[i:], str[i+1:])
			// 删除最后一位
			str = str[:len(str)-1]
		}
		i++
	}
	return str
}
