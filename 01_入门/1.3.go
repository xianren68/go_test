// 1.3  做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 获取纳秒类型的时间戳
	start := time.Now().UnixNano()
	var s, res string
	for _, val := range os.Args[1:] {
		res += val + s
		s = " "
	}
	fmt.Println(res)
	end := time.Now().UnixNano()
	fmt.Println("普通时间", end-start)
	start = time.Now().UnixNano()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end = time.Now().UnixNano()
	fmt.Println("Join时间", end-start)
}
