package main

import "fmt"

//斐波纳切数列
// func cal(n int) int {
// 	var res int
// 	if n <= 2 {
// 		res = 1
// 	} else {
// 		res = cal(n-1) + cal(n-2)

// 	}
// 	return res
// }

//猴子吃桃问题;一天吃一半多一个，每天都吃一半多一个吃到第十天只剩一个桃子

func num(i int) int {

	if i == 10 {
		return 1
	} else {
		return (num(i+1) + 1) * 2
	}
}
func main() {

	//斐波纳切数列
	// var n int
	// fmt.Println("input the number ")
	// fmt.Scanln(&n)fmt.Println("pho = ",1)

	// res := cal(n)
	// fmt.Println("res :=", res)
	fmt.Println("第一天有桃子", num(1))
	fmt.Println("第9天有桃子", num(9))
	fmt.Println("第十天有桃子:", num(10))
}
