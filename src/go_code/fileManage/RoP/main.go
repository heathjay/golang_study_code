package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//1.- 读取文件的内容并显示在终端（带缓冲的方式），使用`os.Open file.Close(),bufio.NewRead(),reader.ReadString`函数和方法

func main() {

	//file的叫法：
	//1.file对象
	//2.file句柄
	//3.file指针
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//及时关闭，需要及时关闭file
	//要及时关闭file句柄，否则要有内存泄漏
	defer file.Close()

	//输出文件，看文件是什么
	fmt.Printf("file = %v", file)

	//创建一个*Reader,默认缓冲区是defaultBufSize = 4096
	//好处是不是一次性读入内存，读一部分处理一部分，
	//循环读取文件内容
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //他会把\n也读出来
		if err == io.EOF {                  //表示文件的末尾，结束读取
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

	//ioutil包内的readfile,只适用于文件不太大的方式
	//返回的是[]byte 切片,
	//
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("bad job")
	}
	//读取到的内容显示到终端，
	//输出是切片格式，会是一堆数字，
	//没有打开也不用关闭
	//都被隐藏在readFile里面了
	fmt.Printf("%v", string(content))

	//关闭文件
	// err = file.Close()
	// if err != nil {
	// 	fmt.Println("file close err")
	// }

}
