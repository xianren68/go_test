// 编写一个rotate函数，通过一次循环完成旋转
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(rangeRotate(x, 2))

}

// 1. 循环,需要额外空间
func rangeRotate(slice []int, k int) []int {
	res := make([]int, len(slice))
	for i, val := range slice {
		res[(i+k)%len(slice)] = val
	}
	return res

}

// 2. 通过反转完成，不需额外空间
// 旋转切片
// slice 要旋转的切片
// k 要旋转的距离
func rotate(slice []int, k int) {
	k = k % len(slice)
	// 1. 全部反转
	reverse(slice)
	// 2. 局部反转
	reverse(slice[k:])
	reverse(slice[:k])
}

// 反转切片
func reverse(slice []int) {
	l := len(slice)
	for i := 0; i < l/2; i++ {
		slice[i], slice[l-i-1] = slice[l-i-1], slice[i]
	}
}
