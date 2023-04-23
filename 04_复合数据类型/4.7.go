// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
package main

import "fmt"

func main() {
	str := "hfigbiredguivddkdk"
	bytes := []byte(str)
	reverse(bytes)
	fmt.Println(string(bytes))
}

// 反转函数
func reverse(bytes []byte) {
	l := len(bytes)
	for i := 0; i < l/2; i++ {
		bytes[i], bytes[l-1-i] = bytes[l-1-i], bytes[i]
	}
}
