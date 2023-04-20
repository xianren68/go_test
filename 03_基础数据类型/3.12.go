// 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isEqual("abced", "baacde"))
}

// 接入接口用于排序
type runeArr []rune

func (r runeArr) Len() int {
	return len(r)
}
func (r runeArr) Less(i, j int) bool {
	return r[i] > r[j]
}
func (r runeArr) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// 判断组成两个字符串的字符相等
func isEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// 排序
	s1Arr := []rune(s1)
	s2Arr := []rune(s2)
	sort.Sort(runeArr(s1Arr))
	sort.Sort(runeArr(s2Arr))
	for i, _ := range s1Arr {
		if s1Arr[i] != s2Arr[i] {
			return false
		}
	}
	return true

}

// 哈希表记录次数也可以做
