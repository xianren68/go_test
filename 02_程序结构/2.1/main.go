package main

import (
	"fmt"

	"2.1/tempconv"
)

func main() {
	var x tempconv.Celsius = -40
	var y tempconv.Fahrenheit = 222
	var z tempconv.Kelvin = 0
	fmt.Println(x.CToF(), y.FToK(), z.KToC())
}
