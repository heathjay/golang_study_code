package main

func main() {
	//测试一下int8 的范围
	// var j int8 = -128
	// fmt.Println("j=", j)

	// //一个变量占用的字节大小和数据类型,引入的unsafe包中的sizeof（）
	// var n2 int64 = 10
	// fmt.Printf("n2 的数据类型：%T	n2的字节数 %d\n", n2, unsafe.Sizeof(n2))

	//尽量使用空间小的数据类型
	//var age int64 = 100
	//var age byte = 100

	//浮点类型
	// var price float32 = 89.12
	// fmt.Println("price =", price)

	// var num1 float32 = -123.0000001
	// var num2 float64 = -123.0000001
	// fmt.Println("num1=", num1, "num2", num2)

	//浮点数的默认类型
	// var num5 = 1.1
	// fmt.Println("num5=", num5)

	// //十进制
	// num6 := 5.12
	// num7 := .123 //0.123
	// fmt.Println("num6", num6, "num7", num7)

	// //科学计数法
	// num8 := 5.1234e2 //10的2次方
	// fmt.Println("num8=", num8)
	// num9 := 5.1234e-2 //10的2次方
	// fmt.Println("num9=", num9)

	// //字符类型
	// var c1 byte = 'a'
	// var c2 byte = '0' //字符0
	// var c byte = 'a'

	// fmt.Println("c1=", c1, "c2=", c2, "c=", c)

	// //若果想要输出对应字符，需要格式化输出
	// //fmt.Printf("c1=%c c2=%c，c3=%c", c1, c2, c3)

	// var c4 int = 22269
	// fmt.Printf("c4=%c", c4)

	// c5 := 10 + 'a'
	// fmt.Println("c5=", c5)

	// //bool类型
	// var b = false
	// fmt.Println("b=", b)
	// //占用空间
	// fmt.Println("b 占用的空间=", unsafe.Sizeof(b))

	//string 的基本使用
	// var address string = "北京您好"

	// fmt.Println("address=", address)

	// //go中的字符串一旦赋值，他的字符串就是不可变的

	// var str = "hello"
	// str[0] = 'a'

	// //输出源代码效果
	// str2 := "abc\ngao"
	// fmt.Println(str2)

	// //源代码输出,使用的是反引号
	// str3 := `package main

	// import "fmt"

	// func main() {
	// 	//定义变量，申明变量
	// 	var i int
	// 	//给i赋值
	// 	i = 10
	// 	fmt.Println("i =", i)
	// }
	// `
	// fmt.Println(str3)

	// //字符串的拼接
	// var str = "hello" + "world"
	// str += "haha"
	// fmt.Println(str)

	// //当一个拼接操作很长时，可以分行,把+留在上面
	// var str1 = "hello" + "world" + "world" + "world" + "world" +
	// 	"world" + "world" +
	// 	"world" + "world"
	// fmt.Println(str1)

	// //数据类型转换
	// var i int32 = 100
	// //var j float32 = i
	// var n1 float32 = float32(i)
	// fmt.Println("n1=", n1)

	// var num1 int64 = 999999
	// var num2 int8 = int8(num1)
	// fmt.Println("num2=", num2)

	// var n1 int32 = 12
	// var n2 int64
	// var n3 int8

	// n2 = int64(n1) + 20 //int32--->int64
	// n3 = int8(n1) + 20  //int32 ---->int8

	// fmt.Println("n2=", n2, "n3=", n3)

	//基础数据类型转换成string
	// var num1 int = 99
	//var num2 float64 = 23.456
	// var b bool = true
	// var mychar byte = 'h'
	// var str string //空的str

	// str = fmt.Sprintf("%d", num1)
	// fmt.Printf("string type %T str=%p", str, str)
	// str = fmt.Sprintf("%f", num2)
	// fmt.Printf("string type %T str=%p\n", str, str)

	// str = fmt.Sprintf("%t", b)
	// fmt.Printf("string type %T str=%p\n", str, str)
	// str = fmt.Sprintf("%c", mychar)
	// fmt.Printf("string type %T str=%q\n", str, str)

	//string strconv
	//formatbool
	//func formatint
	//
	// var num1 int = 99
	// var num2 float64 = 23.456
	// var b bool = true
	// // var mychar byte = 'h'
	// var str string //空的str

	// str = strconv.FormatInt(int64(num1), 10)
	// fmt.Printf("string type=%T str=%q\n", str, str)

	// //'f' 表示一种格式，10小数点后保留，64-float64
	// str = strconv.FormatFloat(num2, 'f', 10, 64)
	// fmt.Printf("string type=%T str=%q\n", str, str)

	// str = strconv.FormatBool(b)
	// fmt.Printf("string type=%T str=%q\n", str, str)

	// var num5 int = 4567
	// str := strconv.Itoa(num5)
	// fmt.Printf("string type %T string = %q", str, str)

	// //string to base type
	// var str string = "true"
	// var b bool
	// //函数本身会返回两个值，
	// //第一个值是一个bool， 第二个是err error
	// //关心的是bool， 用_忽略err
	// b, _ = strconv.ParseBool(str)
	// fmt.Printf("b type = %T, b=%v\n", b, b)

	// var str2 string = "12345"
	// var num int64
	// //str,base(进制),bitSize--0-int 8-int8
	// num, _ = strconv.ParseInt(str2, 10, 64)
	// fmt.Printf("num type = %T, num=%v\n", num, num)

	// var str3 string = "123.456"
	// var num2 float64
	// num2, _ = strconv.ParseFloat(str3, 64)
	// fmt.Printf("num type = %T, num=%v\n", num2, num2)

	// //pointer
	// //基本数据类型在内存中的布局
	// var i int = 10
	// //i 的地址
	// fmt.Println("i的地址是：", &i)

	// //下面的 var ptr *int = &i
	// //ptr 是指针变量
	// //ptr的指针类型是*int
	// //ptr本身的值&i

	// var ptr *int = &i
	// fmt.Printf("ptr=%v\n", ptr)
	// fmt.Printf("ptr 的地址：%v", &ptr)
	// fmt.Printf("ptr 指向的值：%v", *ptr)

	// var num int = 9
	// fmt.Printf("num 指针地址：%v\n", &num)

	// //这里进行修改会导致num的变化
	// var ptr *int = &num
	// *ptr = 10
	// fmt.Printf("num 的值是：%v\n", num)
}
