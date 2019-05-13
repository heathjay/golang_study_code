package main

import (
	"fmt"
	"runtime"
)

type Thread_t struct {
	s     *string
	t     *string
	t_len int
	s_len int
	count int
	//cpus  int
}

func Lev_dist01(thread_t Thread_t) int {

	// //initial
	// height := len(*str) + 1
	// width := len(*targe) + 1
	str := thread_t.s
	targe := thread_t.t
	height := thread_t.s_len + 1
	width := thread_t.t_len + 1
	count := thread_t.count
	//cpus := thread_t.cpus
	//slice 使用
	mtr := make([][]int, height)
	for i := 0; i < height; i++ {
		mtr[i] = make([]int, width)
		mtr[i][0] = i
	}
	for i := 0; i < width; i++ {
		mtr[0][i] = i
	}
	fmt.Println(mtr)
	//一方为0返回长度
	if height == 1 {
		return len(*targe)
	}
	if width == 1 {
		return len(*str)
	}
	fmt.Println("(******************)")
	//算法

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			dif := 1
			if (*str)[i-1] == (*targe)[thread_t.t_len*count+j-1] {
				dif = 0
			}
			mtr[i][j] = Min(Min(mtr[i-1][j]+1, mtr[i][j-1]+1), mtr[i-1][j-1]+dif)
		}

	}

	for _, v := range mtr {
		fmt.Println(v)
	}
	return mtr[height-1][width-1]

}
func Lev_dist(str *string, targe *string) int {

	//initial
	height := len(*str) + 1
	width := len(*targe) + 1

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
		return len(*targe)
	}
	if width == 1 {
		return len(*str)
	}
	fmt.Println("(******************)")
	//算法

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			dif := 1
			if (*str)[i-1] == (*targe)[j-1] {
				dif = 0
			}
			mtr[i][j] = Min(Min(mtr[i-1][j]+1, mtr[i][j-1]+1), mtr[i-1][j-1]+dif)
		}

	}

	for _, v := range mtr {
		fmt.Println(v)
	}
	return mtr[len(*str)][len(*targe)]

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

	//cpus
	cpus := runtime.NumCPU()
	fmt.Println("cpus:", cpus)
	par_len := Min(len(str), len(targe))

	var td []Thread_t
	td = make([]Thread_t, cpus)
	for i := 0; i < cpus; i++ {
		td[i].count = i
		//td[i].cpus = cpus
		td[i].s = &str
		td[i].t = &targe
		td[i].s_len = len(str)
		td[i].t_len = par_len / cpus
		res1 := Lev_dist01(td[i])
		fmt.Println(res1)
	}

	res := Lev_dist(&str, &targe)
	fmt.Println(res)

}
