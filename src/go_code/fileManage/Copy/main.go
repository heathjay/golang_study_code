package main

import (
	"bufio"
	"fmt"
	"io" //提供copy
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (writter int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file error:", err)

	}
	defer srcFile.Close()
	//通过file获取到reader
	reader := bufio.NewReader(srcFile)

	//打开文件
	//可能不存在
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("opehn file err :", err)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)

}

func main() {
	//将RoP文件下面的1.png拷贝到Copy下面
	//编写一个函数，接收两个文件路径，一个srcName，dstName
	srcFile := "../RoP/1.png"
	dstFile := "./1.png"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("copy good")
	} else {
		fmt.Println("bad:", err)
	}
}
