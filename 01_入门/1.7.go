// 1.7 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 获取命令行参数输入的网址
	urls := os.Args[1:]
	for _, url := range urls {
		// 向对应网站发送请求
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		// 获取响应体,并将其直接拷贝到标准输出流
		_, err = io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}
}
