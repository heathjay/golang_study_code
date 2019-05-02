package main

import (
	"fmt"
)

func test(slice []int) {
	slice[0] = 1000
}
func main() {
	//使用常规for遍历
	var arr = [...]int{10, 20, 30, 40, 50}
	//slice := arr[1:4] // 20, 30, 40
	slice := arr[0:len(arr)]

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\n", i, slice[i])
	}

	for index, value := range slice {
		fmt.Printf("slice[%v]=%v\n", index, value)
	}
	slice2 := slice[2:3]
	slice2[0] = 100
	fmt.Println(slice2)
	fmt.Println(slice)
	fmt.Println(arr)

	//append 内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	fmt.Println(slice3)

	//通过append给slice3追加具体的元素，要跟数据类型相同
	slice4 := append(slice3, 400, 500, 600)
	fmt.Println(slice4)
	fmt.Println(slice3)
	//切片再追加切片
	slice3 = append(slice3, slice3...)
	fmt.Println(slice3)

	var a []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)

	fmt.Println(slice5)
	copy(slice5, a)
	fmt.Println(slice5)

	test(slice)
	fmt.Println(slice)
}
