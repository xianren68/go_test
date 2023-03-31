package main

import (
	"fmt"
	"os"
	"strconv"

	"2.2/conversion"
)

func main() {
	// 定义传入的值
	var target string
	// 单位
	var per byte
	if len(os.Args) > 1 {
		// 命令行参数
		target = os.Args[1]

	} else {
		// 输入
		fmt.Scanln(&target)

	}
	per = target[len(target)-1]
	target = target[:len(target)-1]
	// 转成浮点型
	fTarget, err := strconv.ParseFloat(target, 64)
	if err != nil {
		fmt.Println("输入有误")
	} else {
		if per == 'm' {
			// 米转英尺
			fmt.Println(conversion.Meters(fTarget).MToF())
		} else if per == 'f' {
			fmt.Println(conversion.Feet(fTarget).FTom())
		} else {
			fmt.Println("请输入正确的单位")
		}

	}

}
