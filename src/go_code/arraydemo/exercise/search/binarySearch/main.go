package main

import (
	"fmt"
)

//二分查找,
/*
1.中间比较
2.逐步缩小左右两边
*/
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	//判断leftIndex 是否 > rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//从小到大leftIndex
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("good job", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, -6)
}
