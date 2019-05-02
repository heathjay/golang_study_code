package main

import (
	"fmt"
)

func main() {
	//定义二维数组，用来存储三个班级，每个班级5个同学的成绩
	//并求出每个班级平均分，以及所有班级的平均分
	var score [3][5]float64
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Printf("input score of %v student in %vth class\n ", j, i)
			fmt.Scanln(&score[i][j])
		}
	}

	//	fmt.Println(score)
	totalSum := 0.0
	for i, v := range score {
		sum := 0.0
		for _, v2 := range v {
			sum += v2
		}
		totalSum += sum
		fmt.Println(i, " class sum:=", sum,
			"average := ", sum/float64(len(score[i])))
	}
	fmt.Println("所有班级：sum :=", totalSum, "average :=", totalSum/15)
}
