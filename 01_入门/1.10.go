// 1.10  找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出到文件，以便于进行对比。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// 程序开始的时间
	start := time.Now()
	// 创建一个管道
	ch := make(chan string)
	// 获取命令行参数
	urls := os.Args[1:]
	for _, url := range urls {
		// 并发访问网址
		go fetch(url, ch)
		go fetch(url, ch)
	}
	for i := 0; i < 2*len(urls); i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("总时间", time.Since(start).Seconds())

}
func fetch(url string, ch chan<- string) {
	// 记录每个请求的开始时间
	start := time.Now()
	// 发送请求
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// 创建一个文件
	f, err := os.Create(strings.TrimPrefix(url, "http://") + ".txt")
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// 将响应体返回到文件中
	nbytes, err := io.Copy(f, res.Body)
	res.Body.Close()
	// 关闭文件
	f.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// 用时
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
