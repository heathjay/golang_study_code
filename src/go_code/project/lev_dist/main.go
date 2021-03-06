package main

import "fmt"

func Lev_dist(str string, targe string) int {

	//initial
	height := len(str) + 1
	width := len(targe) + 1

	//slice 使用
	mtr := make([][]int, height)
	for i := 0; i < height; i++ {
		mtr[i] = make([]int, width)
		mtr[i][0] = i
	}
	for i := 0; i < width; i++ {
		mtr[0][i] = i
	}

	//一方为0返回长度
	if height == 1 {
		return len(targe)
	}
	if width == 1 {
		return len(str)
	}
	//	fmt.Println("(******************)")
	//算法

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			dif := 1
			if str[i-1] == targe[j-1] {
				dif = 0
			}
			mtr[i][j] = Min(Min(mtr[i-1][j]+1, mtr[i][j-1]+1), mtr[i-1][j-1]+dif)
		}

	}

	// for _, v := range mtr {
	// 	fmt.Println(v)
	// }
	return mtr[len(str)][len(targe)]

}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	var str string
	var targe string

	fmt.Scanln(&str)
	fmt.Scanln(&targe)

	res := Lev_dist(str, targe)
	fmt.Println(res)
}
