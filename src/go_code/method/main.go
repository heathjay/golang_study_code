package main

import (
	"fmt"
	"math"
)

//6.给person实现string方法
type Person struct {
	Name string
	Age  int
}

// func (person *Person) String() string {
// 	str := fmt.Sprintf("name==[%v] age=[%v]", person.Name, person.Age)
// 	return str
// }

//1.1
func (person *Person) speak() {
	(*person).Name = "jack"
	fmt.Println("我是好人", (*person).Name)
}

//1.2
func (person Person) jisuan() {
	res := 0
	for i := 0; i <= 10000; i++ {
		res += i
	}
	fmt.Println(person.Name, "res=", res)
}

//1.3
func (person Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(person.Name, "res=", res)
}

//1.4
func (person Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

type Circle struct {
	Radius float64
}

func (circle Circle) area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}

//5.int，float32数据类型也可以有自己的绑定的方法
type integer int

func (num integer) print() {
	fmt.Println("num=", num)
}

//p代表一个副本，默认是值拷贝，不能影响到main函数内的东西，
//p这个名字可以被指定，不是固定的，后面的不能改，后面的是类型，不能被改变
func (p Person) test() {
	p.Name = "Jack"
	fmt.Println("test()", p.Name) //调用的时候，可以直接变成jack，
}
func main() {

	var person Person
	person.Name = "tom"
	person.test()       //调用方法，类似于函数的传参数。
	fmt.Println(person) //输出tom

	//1.方法快速入门
	//1.给person结构体添加一个speak方法，输出xxx是好人
	person.speak()
	//2.给person添加一个jisuan方法，可以计算1+...+1000的结果
	//方法体内可以像函数一样，进行各种运算。
	person.jisuan()
	//3.给person添加一个jisuan2方法，可以计算累加到n的结果
	person.jisuan2(30)
	//4.有返回值
	n1 := 10
	n2 := 20
	res := person.getSum(n1, n2)
	fmt.Println(res)

	//2.方法的调用和传参机制
	//2.1方法调用时，基本和函数相同，不同的是方法调用时，会将调用方法的变量，当作实参也传递给方法，下面举例说明
	//画出getSum方法的执行过程 +
	//main栈 n1 --> 10 n2 ---> 20 p--->[tom]Name///res := 30
	//getSum栈 n1 --> 10 n2 ---> 20 p已经存在  p--->[tom]Name///return n1 + n2
	//p当作形参传给getsum空间， 如果变量是值类型，就进行值拷贝，是引用类型，就是地址拷贝
	//2.2先声明一个结构体Circle,字段radius,声明一个方法area 和 Circle绑定，可以返回面积，
	var circle Circle
	circle.Radius = 6
	res2 := circle.area()
	fmt.Println(res2)

	//3.方法的声明
	//func(recevier type) methodName(参数列表)(返回值列表){
	//方法体，
	//return 返回值---并不是必须的
	//}

	//4.注意事项
	//4.1希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理，用的最多的是结构体指针类型，
	/*
		func (person *Person) speak() {
		(*person).Name = "jack" //可以person.Name ---被优化了
		fmt.Println("我是好人", (*person).Name)
		}
		person.speak()		//标准写法：（&person）.speak() 但是编译器会识别帮加上。
	*/
	//speak栈内，person是地址值指向main栈内的地址，
	person.speak()
	fmt.Println(person)

	//5.golang中的方法作用在指定的数据类型上，因此自定义类型，都可以有方法，不仅仅是struct，比如int，float32 都可以有方法
	//5.int，float32数据类型也可以有自己的绑定的方法
	// func (num int) print(){
	// 	fmt.Println("num=",i)
	// }
	var i integer = 10
	i.print()

	//6.如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的string（）进行输出
	// func(person *Person) String()string{
	// 	str := fmt.Sprintf("name==[%v] age=[%v]",person.Name,person.Age)
	// }
	stu := Person{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(stu)
	fmt.Println(&stu) //传一个地址过去name==[tom] age=[20]

	(&stu).test()
	fmt.Println(stu) //只关注方法类型，本质仍然是值拷贝，

}
