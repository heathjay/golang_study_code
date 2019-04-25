package main

import (
	"fmt"
)

// func test(n1 int) {
// 	n1 = n1 + 1
// 	fmt.Println("test number is ", n1)

// }
//		编写一个函数，返回和和差
// 返回值 可以多个
// 返回值可以被忽略 _ 占位置

// func getSumAndDiff(n1 int, n2 int) (int, int) {
// 	sum := n1 + n2
// 	dif := n1 - n2
// 	fmt.Println("sum = ", sum, "dif = ", dif)
// 	return dif, sum
// }

//迭代分析
func test(n1 int) {
	if n1 > 2 {
		n1--
		test(n1)
	} else {
		fmt.Println("n1 = ", n1)
	}
}

func main() {
	// n1 := 10
	// //调用
	// test(n1)
	// fmt.Println("main number is ", n1)

	//dif, sum := getSumAndDiff(11, 15)
	// _, sum := getSumAndDiff(11, 15)
	// //fmt.Println("main sum", sum, "main dif", dif)
	// fmt.Println("main sum", sum)

	test(4)
}
