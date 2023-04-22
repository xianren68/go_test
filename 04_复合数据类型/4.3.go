// 重写reverse函数，使用数组指针代替slice
package main

import "fmt"

const length = 10

func main() {
	arr := [length]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(*(reverse(&arr)))
}

// 用来反转数组
func reverse(arr *[length]int) *[length]int {
	for i := 0; i < length/2; i++ {
		// 换位
		(*arr)[i], (*arr)[length-1-i] = (*arr)[length-1-i], (*arr)[i]
	}
	return arr
}
