package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type reader struct {
	i    int
	read io.Reader
	n    int
}

func (r *reader) Read(p []byte) (n int, err error) {
	// 到达n,直接结束
	if r.i >= r.n {
		return 0, io.EOF
	}
	// 剩余的长度不够p
	if r.n-r.i < len(p) {
		n, _ = r.read.Read(p[:r.n-r.i])
		err = io.EOF
		return

	}
	// 通过嵌套读取另一个流
	n, err = r.read.Read(p)
	// 另一个流读到尾部了
	if err != nil {
		return
	}
	// 修改索引
	r.i += n
	return
}
func LimitReader(r io.Reader, n int) io.Reader {
	return &reader{i: 0, read: r, n: n}
}

func main() {
	// 创建一个读入流
	r := strings.NewReader("1234567890")
	r1 := LimitReader(r, 4)
	s := bufio.NewScanner(r1)
	s.Split(bufio.ScanBytes)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
