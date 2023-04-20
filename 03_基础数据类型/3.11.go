// 完善comma函数，以支持浮点数处理和一个可选的正负号的处理,这题题意依然是只处理整数部分，对于正负号或
// 浮点数小数部分都按原样即可
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("+123458974632.79238594"))
}

func comma(s string) string {
	// 字符串前缀，整数部分，小数部分
	var prefix, intstr, floatstr string
	// 判断是否拥有正负号
	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		prefix = s[:1]
		intstr = s[1:]
	}
	// 切分整数和小数部分
	if strings.Contains(intstr, ".") {
		arr := strings.Split(intstr, ".")
		intstr = arr[0]
		floatstr = arr[1]
	}
	// 拼接并添加逗号
	var res bytes.Buffer
	res.WriteString(prefix)
	for i, val := range []byte(intstr) {
		res.WriteByte(val)
		if (i+1)%3 == 0 && i != len(intstr)-1 {
			fmt.Fprint(&res, ",")
		}
	}
	res.WriteByte('.')
	res.WriteString(floatstr)
	return res.String()
}
