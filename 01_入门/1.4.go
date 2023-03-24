// 1.4 修改 dup2，出现重复的行时打印文件名称。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 用于记录每行出现次数的hash表
	hash := make(map[string]int)
	// 用于记录文件名的hash表,每一行对应的文件名存到一个数组中
	fileHash := make(map[string][]string)
	// 获取命令行参数
	files := os.Args[1:]
	// 如果为空，则在控制台上输入
	if len(files) == 0 {
		countlines(os.Stdin, hash, fileHash)
	} else {
		// 不为空，遍历文件列表
		for _, file := range files {
			// 打开文件
			f, err := os.Open(file)
			// 判断文件路径等是否出错
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 传入contlines进行处理
			countlines(f, hash, fileHash)
			f.Close()
		}
	}

	for i, val := range hash {
		fmt.Println(i, val)
		// 打印出现的文件
		if val > 1 {
			for _, v := range fileHash[i] {
				fmt.Printf("%s ", v)
			}
			fmt.Println()
		}
	}

}

// 用于打印输入流中每一行出现的次数
func countlines(f *os.File, hash map[string]int, fileHash map[string][]string) {
	// 创建读入流
	input := bufio.NewScanner(f)
	// 一行一行读取
	for input.Scan() {
		hash[input.Text()]++
		// 将对应的文件名加入数组
		fileHash[input.Text()] = append(fileHash[input.Text()], f.Name())
	}
}
