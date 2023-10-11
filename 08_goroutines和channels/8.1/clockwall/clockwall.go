package main

import (
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

/*
时区id
China Standard Time
Eastern Standard Time
Hawaiian Standard Time
*/
func main() {
	// 设置等待组
	waitGroup := sync.WaitGroup{}
	// 从命令行获取要访问的端口及其对应的时区
	for _, val := range os.Args[1:] {
		// 以等号分割
		line := strings.Split(val, "=")
		// 获取对应的时区及端口
		tz, address := line[0], line[1]
		// 等待组+1
		waitGroup.Add(1)
		go getTime(address, tz, &waitGroup)
	}
	waitGroup.Wait()
}

// 从指定服务器获取其时间
func getTime(address string, tz string, wait *sync.WaitGroup) {
	// 等待组-1
	defer wait.Done()
	// 创建tcp client
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	// 读取服务端发送过来的时间
	time := make([]byte, 100)
	conn.Read(time[:])
	// 输出到标准输出流(时区=时间)
	io.WriteString(os.Stdout, tz+"="+string(time)+"\n")
}
