// 用于转化长度的包
package conversion

import "fmt"

// 定义类型
type Meters float64 // 米
type Feet float64   // 英尺

// String方法
func (m Meters) String() string {
	return fmt.Sprintf("%.3gm", m)
}
func (f Feet) String() string {
	return fmt.Sprintf("%.3gf", f)
}

// 转化方法
func (m Meters) MToF() Feet {
	return Feet(m * 3.28)
}
func (f Feet) FTom() Meters {
	return Meters(f / 3.28)
}
