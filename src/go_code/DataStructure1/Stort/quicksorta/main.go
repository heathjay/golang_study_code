package main

import "fmt"

//快速排序
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	pivot := array[(right+left)/2]
	temp := 0
	//将比pivote 小的数放在左边
	for l < r {
		//先从左边找到》=pivot的值
		for array[l] < pivot {
			l++
		}
		for array[r] < pivot {
			r--
		}
		if l >= r {
			break
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	//左右两边应完成
	//如果l==r，在移动一位
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
