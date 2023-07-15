package main

import (
	"6/practice"
	"fmt"
)

func main() {
	x := &practice.IntSet{}
	x.Add(1)
	x.Add(33)
	x.Add(55)
	fmt.Println(x.Has(1))
	fmt.Println(x.Len())
	c := x.Copy()
	fmt.Println(c)
	c.Add(77)
	fmt.Println(c, x)
}
