// 用于记录一个int类型的数中1bit的个数
package test

// 定义一个byte数组，用来存储
var pc [256]byte

func init() {
	// 一个预处理数组的初始化操作
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 查表法
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 循环法
func RangeCount(x uint64) int {
	var res byte
	for i := 0; i <= 7; i++ {
		res += pc[byte(x>>(i*8))]
	}
	return int(res)
}
