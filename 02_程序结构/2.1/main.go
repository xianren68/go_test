package main

import (
	"fmt"

	"2.1/tempconv"
)

func main() {
	var x tempconv.Celsius = 12.12
	var y tempconv.Fahrenheit = 222
	var z tempconv.Kelvin = 0
	fmt.Println(x.CToK(), y.FToK(), z.KToC())
}
