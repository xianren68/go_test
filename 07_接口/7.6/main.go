package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// 华氏转摄氏
func FtOC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 绝对转摄氏
func KtoC(k Kelvin) Celsius {
	return Celsius(float64(k) + 273.15)
}
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type celsiusFlag struct {
	Celsius
}

// set方法，用于满足flag-value接口，获取命令行参数
func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		c.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = FtOC(Fahrenheit(value))
		return nil
	case "K", "°K":
		c.Celsius = KtoC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	// 将命令行参数传入f的set方法
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var c = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*c)
}
