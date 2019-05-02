package main

import (
	"fmt"
)

func BubbleSort(arr *[5]int) {
	fmt.Println("冒泡排序", *arr)
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println(*arr)
}
func main() {
	//arr = [23, 65, 80, 57, 13]
	//外层排序_第一次比较--让前面的数和后面的数进行比较。
	//如果前面的数字大，则交换
	//第一轮找到最大值
	//第二轮比较找到次大值
	//第三轮找到第三大的值
	//1.经过len(arr) - 1 外层循环比较，每一轮会确定一个数的位置

	var arr = [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	//冒泡排序

}
