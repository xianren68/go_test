package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var flag string
	if len(os.Args) > 1 {
		flag = os.Args[1]
	}
	// 输入的字符串
	var input string
	fmt.Scanln(&input)
	if flag == "" {
		fmt.Println(sha256.Sum256([]byte(input)))
	} else if flag == "SHA384" {
		fmt.Println(sha512.Sum384([]byte(input)))
	} else if flag == "SHA512" {
		fmt.Println(sha512.Sum512([]byte(input)))
	} else {
		fmt.Println("请输入正确的加密方式")
	}

}
