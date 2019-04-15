package main

import (
	"fmt"
)

func main() {

	// //给rand一个随机的种子
	// //time.Now().Unix() 会返回一个从1970:01:01到现在的秒数

	// fmt.Println(time.Now().Unix())
	// //随机生成数，1-100，直到有99跳出

	// count := 0

	// for {

	// 	rand.Seed((time.Now().Unix()))
	// 	n := rand.Intn(100) + 1 //0-100的数，【0-100）
	// 	count++
	// 	if n == 99 {
	// 		break

	// 	}
	// }
	// fmt.Println(count)
	// label2:
	// 	//指定标签的形式来使用break
	// 	//配合标签非常有用
	// 	for i := 0; i < 4; i++ {

	// 		//label1:
	// 		for j := 0; j < 10; j++ {
	// 			if j == 2 {
	// 				break label2
	// 			}
	// 			fmt.Println("j=", j)
	// 		}
	// 	}

	//100以内的数求和，求出当和第一次大于20的当前数
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= 20 {
			fmt.Println("now sum = ", sum, "i =", i)
			break
		}
	}

	//三次机会提示还剩多少次，张无忌，888
	var name string = "张无忌"
	var password string = "888"

	var str string
	var psw string

	for i := 1; i < 3; i++ {
		fmt.Println("input name:")
		fmt.Scanln(&str)
		fmt.Println("input password")
		fmt.Scanln(&psw)
		if str == name && psw == password {
			fmt.Println("good job")
			break
		} else {
			fmt.Println("you only have ", 3-i, "chance")
		}
	}

}
