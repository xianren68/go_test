package main

import (
	"fmt"
	"time"

	"2.5/test"
)

func main() {
	// 查表法
	start := time.Now()
	x := test.PopCount(184467440737095516)
	fmt.Printf("查表法耗时%fs,%d\n", time.Since(start).Seconds(), x)
	start = time.Now()
	x = test.RangeCount(184467440737095516)
	fmt.Printf("清零法耗时%fs,%d\n", time.Since(start).Seconds(), x)
}
