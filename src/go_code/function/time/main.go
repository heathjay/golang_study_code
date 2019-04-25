package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
func main() {
	//日期和时间相关的函数的使用
	//1. 获取当前时间
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("type:%T\n", now)

	//返回年月日时分秒
	fmt.Printf("年:%v\n", now.Year())
	fmt.Printf("月:%v\n", now.Month())
	fmt.Printf("日:%v\n", now.Day())
	fmt.Printf("时:%v\n", now.Hour())
	fmt.Printf("分:%v\n", now.Minute())
	fmt.Printf("秒:%v\n", now.Second())

	//格式化日期时间
	fmt.Printf("%d-%d-%d-%d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dataStr := fmt.Sprintf("%d-%d-%d-%d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v", dataStr)

	//Format, 格式化的第二种方式：各个位置固定的
	fmt.Printf(now.Format("2006/01/02 15:04:06\n"))

	//每格一秒钟循环输出1，2，3，4，5。。。
	// var i int = 0
	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// 	if i >= 10 {
	// 		break
	// 	}
	// }
	// //每隔0.1秒就打印一个数字
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Millisecond * 100)
	// }

	//unix时间戳，unixnano时间戳——可以获取随机数字
	fmt.Printf("unix 纳秒时间戳 : %v,unixnano时间戳:%v\n", now.Unix(), now.UnixNano())

	//统计某一个函数需要的时间
	//执行test03前先获取当前的unix时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行时间：%v", start-end)

}
