package main

import (
	"fmt"
	utils "go_code/function/untils"
)

type myFunType func(int, int) int

//通过函数直接改变其他函数中的变量值，直接进行地址传递。
//调用函数内的变量中的值是目标变量的地址，通过取内容来改变该值
//目标变量通过取地址传值
// func test(n1 *int) {
// 	*n1 = *n1 + 10
// 	fmt.Println("n1 =", *n1)
// 	fmt.Println("n1 var = ", n1)
// 	fmt.Println("n1 address", &n1)
// }

//函数本身也是一个数据类型
//支持对函数返回值命名
func getSum(n1 int, n2 int) (sum int) {
	sum = n1 + n2
	return
}

//函数既然是一种数据类型，函数就可以当作一个形参进行数据调用。
func myFun(funvar myFunType, n1 int, n2 int) int {
	return funvar(n1, n2)
}

//自定义数据类型的使用，
// Type 自定义数据类型名 数据类型
// Type myint int
// Type mySum func(int, int) int

//go支持可变参数
//支持0到多个参数
// func sum(args... int) sum int{

// }
//支持1到多个参数
//写一个sum函数，最少有一个值进来
func sum2(n1 int, args ...int) (sum int) {
	sum = n1
	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[0表示取出切片的第一个元素的值，其他以此类推]
	}
	return
}

//交换n1 和 n2 的值
func change(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

// 14.init函数
// 每一个源文件中，都可以包含一个init函数，该函数会在main函数执行之前，被go运行框架调用，
//变量定义->init函数->main函数

var age = test()

func test() int {
	fmt.Println("test()")
	return 90
}
func init() {
	fmt.Println("init().....")
}

func main() {
	// n := 20
	// test(&n)
	// fmt.Println("main n1 =", n)
	// fmt.Println("mian n address:", &n)

	fmt.Println("main()...", age)
	fmt.Println("Age=", utils.Age, "Name", utils.Name)
	a := getSum
	res := a(10, 40)
	res2 := myFun(getSum, 50, 60)
	fmt.Printf("a 的类型是 %T , getSum 类型是 %T\n", a, getSum)
	fmt.Println("res:", res, "res2:", res2)

	type myInt int //给int取一个别名，但还是不同的数据类型

	var num1 myInt

	num1 = 40
	fmt.Printf("num1 %T \n", num1)
	fmt.Println("num is", num1)

	//函数类型来一个别名，作为函数形参就可以简化：

	res3 := myFun(getSum, 500, 600)
	fmt.Println("res3=", res3)

	res4 := sum2(500, 20, 33, 44, 505, 77)
	fmt.Println("res4=", res4)

	n1 := 10
	n2 := 20
	change(&n1, &n2)
	fmt.Println("change : ", n1, n2)
}
