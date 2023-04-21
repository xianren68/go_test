package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(differentCount(c1, c2))
}

func differentCount(c1, c2 [32]uint8) (count int) {
	count = 0
	for i, _ := range c1 {
		count += popCount(c1[i], c2[i])
	}
	return
}

// 比较每一个bit位是否相同
func popCount(x, y byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		// 判断是否是偶数来判断
		// xb := x % 2
		// yb := y % 2
		// if xb != yb {
		// 	count++
		// }
		// x = x >> 1
		// y = y >> 1
		// 直接获取最后一位,因为1只有最后一位是1，其余全为0
		xb, yb := (x>>i)&1, (y>>i)&1
		// 异或运算，相同为0
		if xb^yb == 1 {
			count++
		}

	}
	return count
}
