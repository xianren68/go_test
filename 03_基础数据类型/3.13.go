package main

import (
	"fmt"
)

func main() {
	const (
		KB = 1000
		MB = KB * KB
		GB = MB * KB
		TB = GB * KB
		PB = TB * KB
		EB = PB * KB
		ZB = EB * KB
		YB = ZB * KB
	)
	fmt.Println(KB, MB)
}
