package main

import (
	"fmt"
)

//一个小计算器，10 5固定，选择输入算子
func operatorSelect(n int) {
	var num1 = 10
	var num2 = 5
	var str string
	var res int

	switch n {
	case 1:
		res = num1 + num2
		str = "+"
	case 2:
		res = num1 - num2
		str = "-"
	case 3:
		res = num1 * num2
		str = "*"
	case 4:
		res = num1 / num2
		str = "/"
	}
	fmt.Println(num1, str, num2, "=", res)
}
func main() {
	var n int

	fmt.Println("-----小小计算器-----")
	fmt.Printf("1.加\n2.-\n3.*\n4./\n0.exit\n")
	fmt.Scanln(&n)
	operatorSelect(n)
}
