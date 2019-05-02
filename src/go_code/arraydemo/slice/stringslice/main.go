package main

import (
	"fmt"
)

func main() {
	//string 本身是个byte类型，可以进行切片处理
	str := "hello@atguigu"
	//用切片获取atguigu
	slice := str[6:]

	fmt.Println(slice)
	//修改string -> 【】byte /或者【】rane -> 修改 ->重写
	//“hello@atguigu”=>"zello@atguigu"
	//但是我们要改成一个中文不成功
	// arr1 := []byte(str)
	// arr1[0] = 'z'
	// str = string(arr1)
	// fmt.Println(str)

	//[]byte字节来处理，汉字是三个字节，会出现乱码
	//解决方法将string转成 【】rune即可
	arr1 := []rune(str)
	arr1[0] = '北'
	str = string(arr1)
	fmt.Println(str)
}
