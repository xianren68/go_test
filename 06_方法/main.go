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
	x.AddAll(66, 88)
	fmt.Println(c)
	c.Add(77)
	c.AddAll(5, 6, 7)
	fmt.Println(c)
	c.SymmetricDifference(x)
	fmt.Println(c)
	fmt.Println(c.Elems())
}
