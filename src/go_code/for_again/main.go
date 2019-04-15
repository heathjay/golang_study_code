package main

func main() {
	//输出10 hello

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("hello")
	// }

	//循环条件时返回一个bool值的，将变量初始化和变量的迭代写在其他地方
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("hello")
	// 	i++
	// }

	//死循环配合break
	// key := 1
	// for {
	// 	if key <= 10 {
	// 		fmt.Println("ok")
	// 	} else {
	// 		break
	// 	}
	// 	key++
	// }

	//遍历字符串
	//方式1
	//传统方式

	// var str string = "hello world !"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c\n", str[i])
	// }

	//方式2
	//for-range 遍历字符串
	//index 按顺序编号
	//val 值
	// str := "abc~ok 上海"
	// for index, val := range str {
	// 	fmt.Printf("index=%d, val=%c\n", index, val)
	// }

	// //切片
	// var str string = "hello world ! 北京"
	// str2 := []rune(str)
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("%c\n", str2[i])
	// }

	//练习
	// //1-100，9的倍数
	// var max int = 100
	// var count uint64 = 0
	// var sum uint64 = 0
	// for i := 1; i <= max; i++ {
	// 	if i%9 == 0 {
	// 		fmt.Printf("ok %v\n", i)
	// 		count++
	// 		sum += uint64(i)
	// 	}
	// }

	// fmt.Printf("sum %v\n", sum)

	// //完成下面输出， 6可变
	// var n int = 6
	// for i := 0; i <= n; i++ {
	// 	fmt.Printf("%v + %v = %v \n", i, n-i, n)
	// }

	//多重循环
	//3个班，每个班5名学员，求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]

	// var class_amount int = 3
	// var persons int = 5
	// var temp float64 = 0
	// var bigSum float64 = 0.0
	// var bigCount int = 0
	// for i := 1; i <= class_amount; i++ {
	// 	sum := 0.0
	// 	count := 0
	// 	for j := 1; j <= persons; j++ {
	// 		fmt.Printf("input the %vth score\n", j)
	// 		fmt.Scanln(&temp)
	// 		if temp >= 60 {
	// 			count++
	// 			bigCount++
	// 		}
	// 		sum += temp
	// 	}
	// 	avg := sum / float64(persons)
	// 	bigSum += sum
	// 	fmt.Printf("%v avarage is  %v 及格人数%v\n", i, avg, count)
	// }
	// fmt.Printf("big sum is %v\n 及格人数 %v", bigSum, bigCount)

	//打印金字塔

	// //i表示层数
	// for i := 1; i <= 3; i++ {
	// 	//j表示每层多少星星
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Printf("\n")
	// }

	//打印半个金字塔
	//i表示层数
	// var num int = 9

	// for i := 1; i <= num; i++ {

	// 	//打印空格
	// 	for k := 1; k <= num-i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	//j表示每层多少星星
	// 	for j := 1; j <= 2*i-1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Printf("\n")
	// }

	//打印空心金字塔
	// var num int = 9

	// for i := 1; i <= num; i++ {

	// 	//打印空格
	// 	for k := 1; k <= num-i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	//j表示每层多少星星

	// 	for j := 1; j <= 2*i-1; j++ {
	// 		if j == 1 || j == 2*i-1 || i == num {
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}

	// 	}
	// 	fmt.Printf("\n")
	// }

	//打印99乘法表
	// var num int = 9
	// for i := 1; i <= num; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%v * %v = %v  ", i, j, i*j)
	// 	}
	// 	fmt.Println()
	// }

}
