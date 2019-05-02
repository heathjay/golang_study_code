package main

import (
	"fmt"
)

func main() {
	//输出一组二维数组
	//声明一个二维数组
	var arr [4][6]int
	//赋值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println(arr)

	//遍历二维数组，按照要求输出徒刑
	for i := 0; i < 4; i++ {
		fmt.Println(arr[i])
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	//内存中的
	fmt.Printf("arr【1】地址%p：\n", &arr[1])
	fmt.Printf("arr[2]的地址：%p\n", &arr[2])
	fmt.Printf("arr【1】[0]地址%p：\n", &arr[1][0])

	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)

	//二维数组的遍历
	//traverse 双层for循环
	//for range方式遍历
	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v ", i, j, v2)
		}
		fmt.Println()
	}

}
