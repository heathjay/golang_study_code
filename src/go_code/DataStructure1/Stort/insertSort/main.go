package main

import "fmt"

func InsertSort(arr *[7]int) {
	//完成第一次给第二个元素找到合适的值
	for j := 1; j < len(arr); j++ {
		insertVal := arr[j]
		inserIndex := j - 1
		//从大到小
		for inserIndex >= 0 && arr[inserIndex] < insertVal {
			arr[inserIndex+1] = arr[inserIndex] //数据后移
			inserIndex--
		}
		//插入
		if inserIndex+1 != j {
			arr[inserIndex+1] = insertVal
		}
	}

}

func main() {
	arr := [7]int{23, 0, 12, 56, 34, -1, 55}
	InsertSort(&arr)
	fmt.Println(arr)
}
