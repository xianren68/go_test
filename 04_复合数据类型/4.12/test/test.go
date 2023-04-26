package test

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 定义用于接受的结构体
type Xkcd struct {
	// 标题
	Title string
	// 链接
	Link string
	// 封面
	Img string
}

// 地址
const Base_url = `https://xkcd.com/`

// 获取漫画信息并存到csv文件中作为离线索引
func GetSaveInfo() {
	// 创建文件
	f, err := os.Create("./xkcd.csv")
	if err != nil {
		fmt.Println("创建文件失败")
	}
	// 关闭文件
	defer f.Close()
	writer := csv.NewWriter(f)
	// 写入表头
	writer.Write([]string{"title", "link", "img"})
	// 循环获取每条数据
	for i := 1; ; i++ {
		xkcd := &Xkcd{}
		url := fmt.Sprintf("%s%d/info.0.json", Base_url, i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		// 判断状态码
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Println("没有更多了")
			return
		}
		if err = json.NewDecoder(resp.Body).Decode(xkcd); err != nil {
			resp.Body.Close()
			fmt.Println(err)
		}
		resp.Body.Close()
		writer.Write([]string{xkcd.Title, xkcd.Link, xkcd.Img})
		// 将缓存中的内容写入到文件里
		writer.Flush()

	}

}

// 通过关键字返回链接
func ResLink(key string) {
	// 寻找离线索引
	f, err := os.Open("./xkcd.csv")
	if err != nil {
		fmt.Println("打开文件失败，查看是否有离线索引")
	}
	defer f.Close()
	reader := csv.NewReader(f)
	// 循环对比
	for {
		result, err := reader.Read()
		// 到达文件尾部
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if key == result[0] {
			fmt.Println(result[1])
			return
		}
	}
	fmt.Println("暂时没有此漫画")
}
