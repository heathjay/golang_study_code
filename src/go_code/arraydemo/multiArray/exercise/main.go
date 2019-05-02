package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func OneRnad() {
	//随机生成10个整数
	rand.Seed(time.Now().Unix())
	var arrRand [10]int
	for i := 0; i < len(arrRand); i++ {
		arrRand[i] = rand.Intn(100) + 1
	}
	fmt.Println(arrRand)
	//倒序打印
	for i := 0; i < len(arrRand)/2; i++ {
		tmp := arrRand[i]
		arrRand[i] = arrRand[len(arrRand)-1-i]
		arrRand[len(arrRand)-1-i] = tmp
	}
	fmt.Println(arrRand)
	//求出平均值
	sum := 0.0
	max := arrRand[0]
	maxIndex := 0
	min := arrRand[0]
	minIndex := 0
	for i, v := range arrRand {
		sum += float64(v)
		if v > max {
			max = v
			maxIndex = i
		}
		if v < min {
			min = v
			minIndex = i
		}
	}
	fmt.Println("sum:=", sum, "average:=", sum/float64(len(arrRand)))
	fmt.Println("max:=", max, "maxIndex:=", maxIndex, "min:=", min, "minIndex:=", minIndex)
}

func SecondInsert(arr *[5]int, num int) (arr1 [6]int) {
	var index int
	for i := 0; i < len(arr); i++ {
		if arr[i] > num {
			index = i
			break
		}
		arr1[i] = arr[i]
		fmt.Println(i)
	}
	arr1[index] = num
	for i := index + 1; i < len(arr1); i++ {
		arr1[i] = arr[i-1]
	}
	return arr1
}

func ThirdInput() {
	fmt.Println("please input [3][4]value")
	var arr [3][4]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("pls input [%v][%v] number ", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == 2 || j == 0 || j == 3 {
				arr[i][j] = 0
			}
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}

func FourthChange(arr *[4][4]int) {
	temp := (*arr)[0]
	(*arr)[0] = (*arr)[3]
	(*arr)[3] = temp

	temp = (*arr)[1]
	(*arr)[1] = (*arr)[2]
	(*arr)[2] = temp

}
func FifthReverse(arr *[5]int) {
	for i := 0; i < len(arr)/2; i++ {
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[len(arr)-i-1]
		(*arr)[len(arr)-i-1] = temp
	}
}

func SixSearch(arr *[10]string, str string) {
	slice := make([]int, 0)
	for i, v := range arr {
		if v == str {
			slice = append(slice, i)
		}
	}
	fmt.Println(slice)
}

func MaoSort(arr *[10]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func BinarySearch(arr *[10]int, num int, startPoint int, endPoint int) {

	if startPoint > endPoint {
		fmt.Println("找不到")
		return
	}
	middle := (startPoint + endPoint) / 2
	if num > (*arr)[middle] {
		BinarySearch(arr, num, middle+1, endPoint)
	} else if num < (*arr)[middle] {
		BinarySearch(arr, num, startPoint, middle-1)
	} else {
		fmt.Println("good job")
	}
}
func SevenSortAndSearch() {
	var num int
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}
	fmt.Println(arr)
	MaoSort(&arr)
	fmt.Println(arr)
	fmt.Scanln(&num)
	BinarySearch(&arr, num, 0, 9)
}

func EightMaxMin() {
	fmt.Println("8.")
	var arr [5]int
	var Max int
	var MaxIndex int
	var Min int
	var MinIndex int
	for i := 0; i < len(arr); i++ {
		fmt.Printf("input num into arr[%v]", i)
		fmt.Scanln(&arr[i])
		if i == 0 {
			Max = arr[0]
			Min = arr[0]
		}
		if arr[i] > Max {
			Max = arr[i]
			MaxIndex = i
		}
		if arr[i] <= Min {
			Min = arr[i]
			MinIndex = i
		}
	}
	fmt.Println(arr)
	fmt.Println("MaxIndex:=", MaxIndex, "max:=", Max, "MinIndex:=", MinIndex, "min:=", Min)
}

func NightFindAverage(arr *[8]int) {
	sum := 0.0
	var count = 0
	for _, v := range arr {
		sum += float64(v)
	}
	average := sum / float64(len(*arr))
	fmt.Println(average)
	for _, v := range arr {
		if float64(v) > average {
			count++
		}
	}
	fmt.Println("greater:", count, "lower:", len(*arr)-count)
}

func Ten() {
	var arr [8]float64
	var sum = 0.0
	var average = 0.0
	var Max float64
	var MaxIndex int
	var Min float64
	var MinIndex int
	fmt.Println("10.")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("plz input a[%v]", i)
		fmt.Scan(&arr[i])
		if i == 0 {
			Max = arr[i]
			Min = arr[i]
		}
		if Max < arr[i] {
			Max = arr[i]
			MaxIndex = i
		}
		if Min > arr[i] {
			Min = arr[i]
			MinIndex = i
		}
		sum += arr[i]
	}
	sum = sum - Max - Min
	average = sum / float64(len(arr))

	var diff float64
	var best = make([]int, 0)
	fmt.Println(best)
	for i, v := range arr {
		if i == 0 {
			diff = math.Abs(v - average)
		}
		if diff > math.Abs(v-average) {
			diff = math.Abs(v - average)

		} else if diff == math.Abs(v-average) {
			best = append(best, i)
		}
	}

	fmt.Printf("sum:= %v average := %v MaxIndex:= %v MinIndex:=%v Best:=%v diff=%v\n", sum, average, MaxIndex, MinIndex, best, diff)
}
func main() {
	//1.随机生成10个整数保存到数组，并且倒序打印，以及求平均值，最大值和最小值的下标，并查找里面是否有55
	OneRnad()
	//2.已经有个排序好（升序）的数组，要求插入一个元素，最后打印一个该数组，顺序仍旧是升序
	var arr = [5]int{1, 2, 3, 8, 9}
	num := 5
	res := SecondInsert(&arr, num)
	fmt.Println(res)
	//3.定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清零
	//ThirdInput()
	//4.定义一个[4][4]的二维数组，逐个从键盘输入值，然后将第一行和第四行的数据进行交换，将第二行和第三行的数据进行交换
	var arr4 [4][4]int = [4][4]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	fmt.Println(arr4)
	FourthChange(&arr4)
	fmt.Println(arr4)
	//5.试保存1 3 5 7 9  存到数组中，并倒序打印
	var arr5 = [5]int{1, 3, 5, 7, 9}
	FifthReverse(&arr5)
	fmt.Println(arr5)

	//6.写出查找的核心代码，比如已知数组arr [10]string里面保存了10个元素，现要查找""AA"在其中是否存在，并打印提示，“如果有多个AA”也要找到对应的下标
	var arr6 = [10]string{"AA", "BB", "CC", "DD", "AA", "CC", "DD", "GG", "AA"}
	SixSearch(&arr6, "AA")

	//7.随机生成10个使用冒泡法排序，排序后使用二分法进行查找
	SevenSortAndSearch()

	//8.编写一个函数可以接受一个数组，该数组有五个数，请找出最大的数和最小的数对应的下表是多少
	//EightMaxMin()

	//9.定义一个数组，并给出8个整数，酋该数组中大于平均值的个数，和小于平均值的个数
	var arr9 [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	NightFindAverage(&arr9)

	//10.跳水比赛。8个评委打分，运动员的成绩是8个成绩去掉一个最高分，一个最低分，剩下的分数平均分就是该运动员的得分
	//1.找到最高分的评委和最低分的评委
	//2.找到最佳的评委和最差的评委就是两个极端相近的
	Ten()
	//
}
