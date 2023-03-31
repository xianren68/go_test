package main

import (
	"fmt"
	"time"

	"2.3/test"
)

func main() {
	// 查表法
	start := time.Now()
	x := test.PopCount(18446744073709551615)
	fmt.Printf("查表法耗时%fs,%d\n", time.Since(start).Seconds(), x)
	start = time.Now()
	x = test.RangeCount(18446744073709551615)
	fmt.Printf("循环法耗时%fs,%d\n", time.Since(start).Seconds(), x)
}
