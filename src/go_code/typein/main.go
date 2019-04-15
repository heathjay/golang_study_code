package main

import (
	"fmt"
)

//键盘输入语句

func main() {
	var name string
	var age int
	var income float32
	var quality bool

	//当程序执行到scanln 的时候会停在那里等待用户输入
	// fmt.Println("input name")
	// fmt.Scanln(&name)

	// fmt.Println("input age")
	// fmt.Scanln(&age)

	// fmt.Println("input income")
	// fmt.Scanln(&income)

	// fmt.Println("input is pass")
	// fmt.Scanln(&quality)

	//fmt.Scanf()

	fmt.Println("input name age income quality, split using block")

	fmt.Scanf("%s %d %f %t", &name, &age, &income, &quality)

	fmt.Printf("name = %v age =%v income=%v quality=%v\n", name, age, income, quality)
}
