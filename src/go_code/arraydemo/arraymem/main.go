package main

import (
	"fmt"
)

//go 认为[3]int 和[4]int 不是一个数据类型
//形参声明的时候带上长度
func test(arr [3]int) {
	arr[0] = 88
}

//指针传递
func test01(arr *[3]int) {
	(*arr)[0] = 88 //数组指针访问元素
}

func main() {
	// var inputArr [5]float64
	// for i := 0; i < len(inputArr); i++ {
	// 	fmt.Println("请输入第,", i, "数据值")
	// 	fmt.Scanln(&inputArr[i])
	// }

	// for i := 0; i < len(inputArr); i++ {

	// 	fmt.Println("第", i, "元素:", inputArr[i])
	// }
	// //for range
	// for index, value := range inputArr {

	// 	fmt.Println("第", index, "元素:", value)
	// }

	//四种初始化数组
	//1.定义就初始化，
	// var numArr [3]int = [3]int{1, 2, 3}
	// fmt.Println("numArr:", numArr)

	// //2.类型推导
	// var numArr1 = [3]int{1, 2, 3}
	// fmt.Println("numArr1:", numArr1)

	// //3....数量不定[...]是固定写法
	// var numArr02 = [...]int{1, 2, 3}
	// fmt.Println("numArr02:", numArr02)

	// //4.
	// var numArr03 = [...]int{1: 800, 0: 900, 2: 1000}
	// fmt.Println("numArr03:", numArr03)

	//数组的遍历方式
	//for range 数组名{}

	//数组是值传递，因此先拷贝后使用，传入test里面改变第二个88
	arr := [3]int{11, 22, 33}
	test(arr)
	fmt.Println(arr) //11,22,33

	//在其他函数中改变该值
	test01(&arr)
	fmt.Println(arr) //88，22，33
}
