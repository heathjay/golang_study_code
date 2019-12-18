package main

import (
	"fmt"
)

//selectSort完成排序
func SelectSort(arr *[5]int) {
	//对指针的处理能否影响外面
	//标准的访问方式: (*arr)[1] = 600
	//(*arr)[1] = 600
	//arr[1] = 600 等价自动把取地址加上去
	//1.先完成第一次大值和arr[0]先易后难

	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		//遍历后面的1---len(arr)-1

		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				//找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换一次需要一次判断
		if maxIndex != j {
			// temp := arr[0]
			// arr[0] = arr[maxIndex]
			// arr[maxIndex] = arr[0]
			//互换
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}

}
func main() {
	arr := [5]int{19, 32, 10, 100, 80}
	SelectSort(&arr)
	fmt.Println(arr)
}
