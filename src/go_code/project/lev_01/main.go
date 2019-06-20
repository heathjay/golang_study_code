package main

import "fmt"

func Lev_dist(str string, targe string) int {

	//initial
	width := len(str) + 2
	height := len(targe) + 2
	var maxLen int
	if len(str) > len(targe) {
		maxLen = len(str)
	} else {
		maxLen = len(targe)
	}
	//slice 使用
	mtr := make([][]int, height)
	for i := 0; i < height; i++ {
		mtr[i] = make([]int, width)
		mtr[i][0] = maxLen
		if i > 0 {
			mtr[i][1] = i - 1
		}
	}
	for i := 0; i < width; i++ {
		mtr[0][i] = maxLen
		if i > 0 {
			mtr[1][i] = i - 1
		}
	}

	//一方为0返回长度
	if height == 1 {
		return len(str)
	}
	if width == 1 {
		return len(targe)
	}
	//fmt.Println("(******************)")
	//算法
	for _, v := range mtr {
		fmt.Println(v)
	}

	for i := 2; i < width; i++ {
		for j := 2; j < height; j++ {
			if str[i-2] == targe[j-2] {
				mtr[i][j] = mtr[i-1][j-1]
			} else if str[i-3] == targe[j-2] && i >= 3 {
				mtr[i][j] = 1 + Min(Min(mtr[i][j-1], mtr[i-1][j]), mtr[i-2][j-2])
			} else if str[i-2] == targe[j-3] && j >= 3 {
				mtr[i][j] = 1 + Min(Min(mtr[i][j-1], mtr[i-1][j]), mtr[i-2][j-2])
			} else {
				mtr[i][j] = 1 + Min(Min(mtr[i-1][j-1], mtr[i][j-1]), mtr[i-1][j])
			}

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
