package main

func main() {

	//判断一个数与0的关系
	// var num int

	// fmt.Println("plz input a number.....")
	// fmt.Scanln(&num)

	// if num > 0 {
	// 	fmt.Printf("%v >0\n", num)
	// } else if num < 0 {
	// 	fmt.Printf("%v <0\n", num)
	// } else {
	// 	fmt.Printf("%v ==0\n", num)
	// }

	//判断是否是闰年

	// var year int
	// fmt.Println("plz input a year.....")
	// fmt.Scanln(&year)

	// if year%4 == 0 && year%400 != 0 {
	// 	fmt.Println("it is a run nian")
	// } else {
	// 	fmt.Println("it is not a run nian")
	// }

	//判断一个整数是否是水仙花数，153 = 1*1*1 + 3*3*3 + 5*5*5

	// var num int

	// fmt.Println("plz input a number(3 b)...")
	// fmt.Scanln(&num)

	// h_num := num / 100
	// m_num := num/10 - h_num*10
	// l_num := num % 10
	// if (h_num*h_num*h_num + m_num*m_num*m_num + l_num*l_num*l_num) == num {
	// 	fmt.Println("good job ")
	// } else {
	// 	fmt.Println("terrible")
	// }

	//保存用户名和密码， 判断用户名是否为“张三”， 密码是否是1234。如果是，提示登录成功，否则提示登录失败
	// var name string
	// var password string

	// var str string = "张三"
	// var num string = "1234"

	// fmt.Println("plz input your name....")
	// fmt.Scanln(&name)
	// fmt.Println("plz input your password..")
	// fmt.Scanln(&password)

	// if str == name && num == password {
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("shit")
	// }

	//根据输入的年份和月份，求天数

	// var flag int = 0
	// var year, month int

	// fmt.Println("input year and month splitting by block")
	// fmt.Scanf("%d %d", &year, &month)
	// //fmt.Println(year, month)

	// //判断闰年
	// if year%4 == 0 && year%400 != 0 {
	// 	flag = 1
	// }

	// switch month {
	// case 1, 3, 5, 7, 8, 10, 12:
	// 	fmt.Println("31days")
	// case 2:
	// 	if flag == 1 {
	// 		fmt.Println("29days")
	// 	} else {
	// 		fmt.Println("28days")
	// 	}
	// default:
	// 	fmt.Println("30days")
	// }

	//根据公式(身高-108)*2=体重可以有10公斤的浮动

	// var standard float32
	// var height float32
	// var weight float32

	// fmt.Println("plz input height, weight....")
	// fmt.Scanf("%f %f", &height, &weight)

	// standard = (height - 108) * 2

	// if (standard-10) <= weight && weight <= (standard+10) {
	// 	fmt.Println("good job")
	// } else {
	// 	fmt.Println("not a good ")
	// }

	//判断考试成绩的等级
	// var score float32

	// fmt.Println("input your score")
	// fmt.Scanln(&score)

	// switch score/10{
	// case 9, 100:
	// 	fmt.Println("a")
	// case 8:
	// 	fmt.Println("a")
	// }

	// var a int
	// var b int
	// fmt.Println("input a")
	// fmt.Scanln(&a)
	// fmt.Println("input b")
	// fmt.Scanln(&b)

	// if a%b == 0 || a+b >= 1000 {
	// 	fmt.Println("a=", a)

	// } else {
	// 	fmt.Println("b=", b)
	// }

}
