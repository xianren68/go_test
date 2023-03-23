// 1.1 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, res string
	for _, val := range os.Args {
		res += val + s
		s = " "
	}
	fmt.Println(res)
}
