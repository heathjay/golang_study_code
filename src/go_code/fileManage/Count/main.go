package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定一个结构体用于保存统计的结果
type CountChar struct {
	ChCount    int
	NumCount   int
	SpCount    int
	OtherCount int
}

func main() {
	//打开一个文件，创建一个Reader
	//每读出一行，就去统计该行有多少个英文数字空格和其他字符，
	//然后将结果保存到一个结构体中

	file, err := os.Open("./abc.txt")
	if err != nil {
		fmt.Println("bad job:", err)
		return
	}
	defer file.Close()
	var count CountChar
	//创建一个reader
	//reader是带buffer

	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历string进行统计，
		for _, v := range content {
			//fmt.Print(string(v))
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++

			}
		}

	}
	fmt.Printf("字数个数：%v,数字个数：%v,空格个数:%v,其他个数：%v", count.ChCount, count.NumCount, count.SpCount, count.OtherCount)
}
