package main

import (
	"fmt"
	"sort"
)

func main() {
	// 测试
	t1 := []int{1, 2, 3, 2, 1}
	fmt.Println(isPalindrome(sort.IntSlice(t1)))
	t2 := []string{"hello", "world", "hello"}
	fmt.Println(isPalindrome(sort.StringSlice(t2)))
	t3 := []string{"hello", "world", "hell"}
	fmt.Println(isPalindrome(sort.StringSlice(t3)))
}
func isPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		// 两边依次比较，有不同则跳出
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
