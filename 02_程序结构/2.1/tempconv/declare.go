// 用于温度转换的包
package tempconv

import "fmt"

// 定义三种温度类型
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kelvin float64     // 绝对温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 0度
	BoilingC      Celsius = 100     // 100度
)

// string方法，通常用在输出的时候
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
