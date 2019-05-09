package main

import (
	"fmt"
	"os"
)

func main() {
	//命令行获取切片
	//os.Args 切片
	fmt.Println("命令行输入参数值,", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("arg[%v]:%v\n", i, v)
	}
}
