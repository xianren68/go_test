package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

// 递归将数组中的数据添加到树中
func add(value int, t *tree) *tree {
	if t == nil {
		t = &tree{}
		t.value = value
		return t
	}
	if t.value > value {
		t.left = add(value, t.left)
	} else {
		t.right = add(value, t.right)
	}
	return t
}
func sort(values []int) *tree {
	var root *tree
	for _, val := range values {
		root = add(val, root)
	}
	// 通过这个切片来改变原切片底层数组的值（超过容量时才改变地址）
	appendvalues(values[:0], root)
	// 返回个值方便测试
	return root

}

// 递归把树上的数据添加到数组中
func appendvalues(values []int, t *tree) []int {
	if t != nil {
		values = appendvalues(values, t.left)
		values = append(values, t.value)
		values = appendvalues(values, t.right)
	}
	return values
}

// string 方法
func (t *tree) String() string {
	values := appendvalues(make([]int, 0), t)
	return fmt.Sprintf("%v", values)
}
func main() {
	var m = []int{1, 4, 7, 9, 3, 5, 7, 2, 4}
	t := sort(m)
	fmt.Println(m)
	fmt.Println(t)

}
