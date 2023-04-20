// 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("ytsfssentxfesfegad"))
}

func comma(s string) string {
	if len(s) < 3 {
		return s
	}
	var res bytes.Buffer
	for i, val := range []byte(s) {
		res.WriteByte(val)
		if (i+1)%3 == 0 && i != len(s)-1 {
			fmt.Fprint(&res, ",")
		}
	}
	return res.String()
}
