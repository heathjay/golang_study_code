package main

import (
	"fmt"
)

func main() {

	num1 := 100
	fmt.Printf("num1的类型%T，num1的值：%v，num1的地址：%v\n", num1, num1, &num1)

	num2 := new(int) //*int

	//num2的类型--》*int
	//num2的值 = 地址
	//num2的地址--地址
	//num2指向的值是0
	fmt.Printf("num2的类型%T，num2的值：%v，num2的地址：%v，num2指针指向的值是：%v\n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2的类型%T，num2的值：%v，num2的地址：%v，num2指针指向的值是：%v\n", num2, num2, &num2, *num2)
}
