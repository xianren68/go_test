package main

import (
	"bufio"
	"fmt"

	"bytes"
)

type wordCounts int
type lineCounts int

func (w *wordCounts) Write(p []byte) (int, error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*w++
	}
	return int(*w), nil

}
func (l *lineCounts) Write(p []byte) (int, error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*l++
	}
	return int(*l), nil
}
func main() {
	var w wordCounts
	var l lineCounts
	fmt.Fprintf(&w, "hello li, this is %s", "go")
	fmt.Fprintf(&l, `hello
	li
	hao
	xxx
	xx
	%s`, 12)
	fmt.Println(w, l)
}
