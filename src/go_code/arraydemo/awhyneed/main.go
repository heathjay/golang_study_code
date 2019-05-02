package main

import (
	"fmt"
)

func main() {
	/*?l六只鸡的体重，求出他们呢的和以及平均体重*/

	hen1 := 3
	hen2 := 5
	hen3 := 11

	totoalweight := hen1 + hen2 + hen3
	averageweight := totoalweight / 3

	//avagw := fmt.Sprintf("%.2f", totalweight/3)四舍五入保留小数点2位。
	fmt.Printf("total:%v , averageweight%v", totoalweight, averageweight)

	//使用数组的方式来解决问题，定一个数组
	var hens [3]float64
	//给数组的每个元素赋值;
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0 //hens从第一个元素开始0

	totalWeight := 0.0
	//数组可以遍历，求出总体中，
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	fmt.Printf("total:%v\n", totalWeight)

}
