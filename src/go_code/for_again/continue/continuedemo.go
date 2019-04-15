package main

import (
	"fmt"
)

func main() {
	//label2:
	//指定标签的形式来使用break
	//配合标签非常有用
	// for i := 0; i < 4; i++ {

	// 	//label1:
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			continue
	// 		}
	// 		fmt.Println("j=", j)
	// 	}
	// }

	//

	// here:
	// 	for i := 0; i < 2; i++ {
	// 		for j := 0; j < 4; j++ {
	// 			if j == 2 {
	// 				continue here
	// 			}
	// 			fmt.Println("i=", i, "j", j)
	// 		}
	// 	}

	//continue 实现打印1-100之内的奇数

	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// //从键盘读入的个数不确定的整数，并判断读入的整数和负数的个数， 当输入为0时结束程序

	// var z_count int = 0
	// var f_count int = 0
	// for {
	// 	var num int64
	// 	fmt.Println("please input number")
	// 	fmt.Scanln(&num)
	// 	if num > 0 {
	// 		z_count++
	// 		continue
	// 	} else if num < 0 {
	// 		f_count++
	// 		continue
	// 	} else if num == 0 {
	// 		break
	// 	}
	// }
	// fmt.Println("z_count=", z_count, "f_count", f_count)

	//100000
	//>50000 5%
	//<50000 1000
	var count int = 0
	var money float64 = 100000
	for {
		if money <= 1000 {
			break
		}

		if money > 50000 {
			money = money - money*0.05
		} else {
			money -= 1000
		}
		count++
	}

	fmt.Println("money", money, "count", count)
}
