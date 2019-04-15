package main

import (
	"fmt"
)

func main() {
	//var age int
	//接收
	// fmt.Println("pls input your age")
	// fmt.Scanln(&age)

	// if age := 20; age > 18 {
	// 	fmt.Println("warning you need to take responsibility of your action")
	// } else {
	// 	fmt.Println("do anything")
	// }

	//声明2个int32并赋值，如果大于50 helloword

	// var num1 int32 = 17
	// var num2 int32 = 35

	// if (num1 + num2) >= 50 {
	// 	fmt.Println("hello world")
	// }

	// //两个float64赋值，判断第一个数大于10.0
	// //切第二个小于20.0 打印两个值和

	// var n3 float64 = 11.0
	// var n4 float64 = 17.0
	// if n3 > 10.0 && n4 < 20.0 {
	// 	fmt.Println(n3 + n4)
	// }

	//定义两个int32 是否既能被3 又能被5整除

	// var n5 int32 = 3
	// var n6 int32 = 5

	// if (n5+n6)%3 == 0 && (n5+n6)%5 == 0 {
	// 	fmt.Println("hello wo")

	// }

	// //判断2020 年是否是闰年
	// var year int = 2020
	// if (year%4 == 0 && year%400 != 0) || (year%400 == 0) {
	// 	fmt.Println("瑞年")
	// }

	//二元一次方程求解

	// var b float64 = 100.0
	// var a float64 = 3.0
	// var c float64 = 6.0

	// m := b*b - 4*a*c
	// //多分枝判断
	// if m > 0 {
	// 	x1 := ((-b) + math.Sqrt(m)) / 2
	// 	x2 := ((-b) - math.Sqrt(m)) / 2
	// 	fmt.Printf("x1 = %v x2 = %v", x1, x2)
	// } else if m == 0 {
	// 	x1 := ((-b) + math.Sqrt(m)) / 2
	// 	fmt.Printf("x1 = %v ", x1)
	// } else if m < 0 {
	// 	fmt.Printf("无解。。。")
	// }

	// var height int32
	// var wealth float32
	// var outlook bool

	// fmt.Println("please input your height...")
	// fmt.Scanln(&height)

	// fmt.Println("please input your wealth")
	// fmt.Scanln(&wealth)

	// fmt.Println("please input your outlook")
	// fmt.Scanln(&outlook)

	// if height > 180 && wealth > 1.0 && outlook == true {
	// 	fmt.Println("you are lucky")
	// } else if height > 180 || wealth > 1.0 || outlook == true {
	// 	fmt.Println("all right")
	// } else {
	// 	fmt.Println("leave")
	// }

	// //跑步使用描述
	// //性别

	// var slot float64
	// var sex string

	// fmt.Println("input time")
	// fmt.Scanln(&slot)

	// if slot <= 8 {
	// 	fmt.Println("xingbie ")
	// 	fmt.Scanln(&sex)

	// 	if sex == "男" {
	// 		fmt.Println("男子组")
	// 	} else if sex == "女" {
	// 		fmt.Println("女子组")
	// 	}
	// }

	//旺季--bool
	//年龄 ---int

	// var month byte
	// var age byte

	// fmt.Println("几月份？")
	// fmt.Scanln(&month)
	// fmt.Println("年龄？")
	// fmt.Scanln(&age)

	// if month >= 4 && month <= 10 {
	// 	if age > 60 {
	// 		fmt.Println("老年票")
	// 	} else if age < 18 {
	// 		fmt.Println("少年票")
	// 	} else {
	// 		fmt.Println("正常票")
	// 	}
	// } else {
	// 	fmt.Println("单机票")
	// }

	//swith 的表达

	// var key byte
	// fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	// fmt.Scanf("%c", &key)

	// switch key {
	// case 'a':
	// 	fmt.Println("周一")
	// case 'b':
	// 	fmt.Println("周二")
	// case 'c':
	// 	fmt.Println("周三")
	// case 'd':
	// 	fmt.Println("周四")
	// case 'e':
	// 	fmt.Println("周五")
	// case 'f':
	// 	fmt.Println("周6")
	// case 'g':
	// 	fmt.Println("周7")
	// default:
	// 	fmt.Println("error")
	// }

	//switch 后面可以不带表达式，类似if --else 分支来使用
	// var age int = 10
	// switch {
	// case age == 10:
	// 	fmt.Println("age == 10")
	// case age == 20:
	// 	fmt.Println("age == 20")
	// default:
	// 	fmt.Println("gg")
	// }

	//switch 的穿透

	// var num int = 10
	// switch num {
	// case 10:
	// 	fmt.Println("ok1")
	// 	fallthrough //默认只能穿透一层，不带break 的感觉
	// case 20:
	// 	fmt.Println("ok2")
	// case 30:
	// 	fmt.Println("ok3")
	// }

	//输入字符进行switch进行判断
	// var char byte
	// fmt.Println("input a char")
	// fmt.Scanf("%c", &char)

	// switch char {
	// case 'a':
	// 	fmt.Println("A")
	// default:
	// 	fmt.Printf("%v", char)
	// }

	// var score float32
	// fmt.Println("score?")
	// fmt.Scanln(&score)

	// switch int(score / 60) {
	// case 1:
	// 	fmt.Println("几个")
	// case 0:
	// 	fmt.Println("不及格")
	// default:
	// 	fmt.Println("输入有误")
	// }

	//根据月份打印季节
	var month byte
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("spring")
	case 6, 7, 8:
		fmt.Println("summer")
	case 9, 10, 11:
		fmt.Println("fat")
	default:
		fmt.Println("winter")
	}
}
