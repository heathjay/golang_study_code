package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//1.创建一个新文件写一个文件输入5如hello godarn
func WriteOp() {
	file, err := os.OpenFile("./abc.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("open file worse, err := ", err)
	}
	fmt.Println("open file successful, ready to write hello garden")

	//准备写入内容
	//带缓存的writer bufio.NewWriter
	str := "hello garden\r\n" //有些编辑器认识\r有些认识\n
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str) //还没有真正写入到磁盘里面
		//在调用writestring方法时，
		//其实内容时先写入到缓存中去
		//需要把缓存内容存到磁盘里面

	}
	writer.Flush()
	//及时关闭file句柄防止内存泄漏
}

//2.打开一个已经存在的文件，把里面的内容覆盖掉覆盖成“你好尚硅谷”
func WriteOp2() {
	file, err := os.OpenFile("./abc.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "你好尚硅谷\r\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

//3.打开一个存在的文件，在原来的文件内容中国呢追加‘ABC|english’
func WriteOp3() {
	file, err := os.OpenFile("./abc.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "abc!english\r\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

//4.打开一个存在的文件，将原来的内容读出来显示在终端，并且追加5句hello,beijign
func WROper() {
	file, err := os.OpenFile("./abc.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(string(content))
	}

	str := "hellobeijing\r\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

//5.编写一个程序，把一个文件的内容写到另一个文件里面去，这两个文件已经存在了

func CopyWrite() {
	//首先读取需要拷贝的内容
	file1Path := "../RoP/test.txt"
	file2Path := "./abc.txt"

	//ioutil包 readfile
	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("bad")
		return
	}

	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Println("writer err:", err)
	}
	//讲读取到的内容写入到dest

}

func main() {
	WriteOp()
	WriteOp2()
	WriteOp3()
	WROper()
	CopyWrite()
}
