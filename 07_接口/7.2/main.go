package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	w     io.Writer
	count *int64
}

func (c *CountWriter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanBytes)
	var n int
	for sc.Scan() {
		n++
		*(*c).count++
	}
	return n, nil
}
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	return &CountWriter{w, &c}, &c
}
func main() {
	nW, _ := CountingWriter(os.Stdin)
	fmt.Fprintf(nW, "xyzd,fi%s", "dfesd")
}
