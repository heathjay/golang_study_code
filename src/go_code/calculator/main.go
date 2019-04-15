package main

import "fmt"

// func test() bool {
// 	fmt.Println("test...")
// 	return true
// }

func main() {

	// //重点讲解 /%
	// //除法取整，两边都是整数，保留整数部分
	// //
	// fmt.Println(10 / 4)
	// var f1 float32 = 10 / 4
	// fmt.Println(f1)

	// //如果希望保留小数部分
	// //需要浮点数参与运算

	// var f2 float32 = 10.0 / 4
	// fmt.Println(f2)

	// //取模%
	// //公式 a%b = a -a/b*b
	// fmt.Println(10 % 3)
	// fmt.Println(-10 % 3) //-10 - (-10)/3*3 = -1
	// fmt.Println(10 % -3)
	// fmt.Println(-10 % -3)

	// //++--
	// var i int = 10
	// i++
	// fmt.Println("i=", i)
	// i--
	// fmt.Println("i=", i)

	// var week int = 97 / 7
	// var days int = 97 % 7
	// fmt.Println("week=", week, "days=", days)

	// //定义一个变量保存华氏温度，
	// var huashi float32 = 134.2
	// var sheshi float32 = 5.0 / 9 * (huashi - 100)
	// fmt.Println("摄氏温度", sheshi)

	// test() ---- >不执行， 短路与
	// var i int = 10
	// if i < 9 && test() {
	// 	fmt.Println("ok........")
	// }

	//求出两个数的最大
	var num1 int = 10
	var num2 int = 40
	var num3 int = 45
	var max int
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}

	if num3 > max {
		max = num3
	}

	fmt.Println("max = ", max)

}
