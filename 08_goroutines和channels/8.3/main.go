package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}
	// 空结构体尺寸为0,通常当作信号来传播
	ch := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		// 发送信号，让等待阻塞的main函数继续执行
		ch <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	// 标准输入流关闭后，关闭conn的输入流
	conn.(*net.TCPConn).CloseWrite()
	// 阻塞等待服务端消息结束
	<-ch
}

func mustCopy(dst io.Writer, src io.Reader) {
	io.Copy(dst, src)
}
