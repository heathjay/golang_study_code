package main

import (
	"fmt"
)

type MethodUtils struct {
}

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (calculator *Calculator) getSum() float64 {
	return calculator.Num1 + calculator.Num2
}
func (calculator *Calculator) getDiff() float64 {
	return calculator.Num1 - calculator.Num2
}

func (calculator *Calculator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calculator.Num1 + calculator.Num2
	case '-':
		res = calculator.Num1 - calculator.Num2
	case '*':
		res = calculator.Num1 * calculator.Num2
	case '/':
		res = calculator.Num1 / calculator.Num2
	default:
		fmt.Println("operator error")

	}
	return res
}
func (methodUtils MethodUtils) RectPrint() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (methodUtils MethodUtils) RectPrintf(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (methodUtils MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println("even num")
	} else {
		fmt.Println("oven")
	}
}

func (mu *MethodUtils) MutipleTable(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v ", j, i, i*j)
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) Change(arr *[3][3]int) {
	for i := 0; i < len(*arr); i++ {
		for j := i; j < len((*arr)[i]); j++ {
			tmp := (*arr)[i][j]
			(*arr)[i][j] = (*arr)[j][i]
			(*arr)[j][i] = tmp
		}
	}
}
func main() {
	//1.编写一个结构体MethodUtils，编写一个方法，方法不需要函数，在方法中打印一个10*8的举行，在main方法中调用该方法
	var methodUtils MethodUtils
	methodUtils.RectPrint()
	//2.方法中传入m, n 打印矩形
	fmt.Println()
	methodUtils.RectPrintf(5, 18)
	//3.返回一个面积值
	res := methodUtils.area(8, 7)
	fmt.Println(res)
	//4.编写一个方法，判断是奇数还是偶数
	methodUtils.JudgeNum(10)

	//5.定义一个计算器的结构体Calculator
	//实现加减乘除功能
	//实现形式1:分四个方法完成
	//实现方式2：用一个方法搞定，需要接收两个数，还有一个运算符
	var calculator Calculator
	calculator.Num1 = 1.1
	calculator.Num2 = 2.2
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", calculator.getSum()))
	fmt.Println("diff:=", calculator.getDiff())
	res1 := calculator.getRes('+')
	fmt.Println("res1:=", res1)

	//6.在methodUtils结构体编写一个方法，从键盘接收整数(1-9,)打印对应的乘法表
	methodUtils.MutipleTable(4)

	//编写方法，是给定的一个二维数组（3*3）转置
	var arr [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr)
	methodUtils.Change(&arr)
	fmt.Println(arr)

}
