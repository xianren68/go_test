// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
package main

import "fmt"

func main() {
	str := "afdigneigu  \n\t   fjidfiefh  jsdifj  "
	bytes := []byte(str)
	onlySpace(bytes)
	fmt.Println(bytes)
	fmt.Println(string(bytes)) // afdigneigu fjidfiefh jsdifj
}

// 将多个空格转化为一个
func onlySpace(str []byte) {
	// 定义前后两个指针
	pre := 0
	cur := 1
	for cur < len(str) && str[cur] != 0 { // 遇到零值就停下
		if isSpace(str[pre]) { // 如果前面为空格
			for cur < len(str) && isSpace(str[cur]) { // 后面直接走到不是空格的位置
				cur++
			}
			// 通过copy往前挪
			copy(str[pre+1:], str[cur:])

			// 将最后几位转为0值
			// fmt.Println(pre, cur)
			for j := len(str) - 1; j > (len(str) - 1 - (cur - pre - 1)); j-- {
				str[j] = 0
			}
			pre += 2      // 跳过肯定不是空格的位置
			cur = pre + 1 // 调整cur的位置

		} else {
			pre++
			cur++
		}
	}
}

// 判断是否为空格
func isSpace(b byte) bool {
	spaces := []byte{'\t', '\n', '\v', '\f', '\r', ' '}
	for _, val := range spaces {
		if b == val {
			return true
		}
	}
	return false
}
