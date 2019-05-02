package main

import "fmt"

func main() {
	//演示切片的基本操作
	// var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// //声明/定义一个切片
	// //slice := intArr[1:3]
	// //1. slice 是切片名
	// //2. intArr 表示slice 引用到intArr数组的第二个元素到下标为3,但是不包含3【1,3）
	// slice := intArr[1:3]

	// fmt.Println("intArr:=", intArr)
	// fmt.Println("slice elements:= ", slice)
	// fmt.Println("length of slice:", len(slice))
	// fmt.Println("capability of slice:=", cap(slice)) //切片的容量可以动态变化的会进行自动增长 稍微多一些

	// fmt.Println("intarr[1] address：=", &intArr[1])
	// fmt.Printf("slice[0]的地址=%p\n", &slice[0])
	// fmt.Println("slice 的值是：", slice[0])

	// slice[1] = 34
	// fmt.Println(intArr)

	//slice usage
	//slice make usage
	//对于切片，必须make后才能使用，或者对数组进行引用——slice := intArr[1:3] 1-2
	//make([]type, len, cap)
	var slice = make([]float64, 4, 10)
	slice[2] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice elements:= ", slice)
	fmt.Println("length of slice:", len(slice))
	fmt.Println("capability of slice:=", cap(slice)) //切片的容量可以动态变化的会进行自动增长 稍微多一些

	var slice2 = []string{"tom", "joy", "mary"}
	fmt.Println(slice2)
	fmt.Println("slice elements:= ", slice2)
	fmt.Println("length of slice:", len(slice2))
	fmt.Println("capability of slice:=", cap(slice2))

}
