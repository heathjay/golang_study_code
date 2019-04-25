package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//统一utf-8
	//一个汉字占用3个字节
	str := "hello北"
	fmt.Println("string len:", len(str))

	//字符遍历，出现中文的问题 r := []rune(str2)
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符==%c\n", r[i])
	}

	//字符串转整数，: n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Println("转换结果是:", n)
	}

	//整数转字符串，str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("%T，str = %v", str, str)

	//字符串转[]byte: var bytes = []byte("hello world")
	var bytes = []byte("hello world")
	fmt.Printf("%v\n", bytes)

	//十进制转换2,8,16进制返回对应的字符串 func FormatInt
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制:%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制：%v\n", str)

	//统计一个字符串中有几个指定的字串Count()
	num := strings.Count("ceheese", "e")
	fmt.Printf("字符个数：%v\n", num)

	//不区分大小写的测试strings.EqualFold("abc", "ABC")
	fmt.Println(strings.EqualFold("abc", "ABC"))

	//返回字串在字符串第一次出现的index的值，如果没有返回-1
	index := strings.Index("nads_abc", "abc")
	fmt.Printf("index:%v\n", index)

	index = strings.LastIndex("hello go golang", "go")
	fmt.Printf("lastindex:%v\n", index)

	//str 替换,原先的str1不变化
	str1 := "go go hello"
	str = strings.Replace(str1, "go", "go语言", -1)
	fmt.Println(str)

	//拆分字符串
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)

	//去掉左右两边的指定字符串
	str = strings.Trim("!hello ! ", " !")
	fmt.Println(str)
}
