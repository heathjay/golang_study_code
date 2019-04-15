package main

import (
	"fmt"
	"go_code/function/untils"
)

//将计算的功能放到一个函数中
//然后在需要使用，调用即可

func main() {
	//输入两个数，一个运算符，得出结果

	result := untils.Cal(1.2, 2.3, '+')
	fmt.Println("res= ", result)

}
