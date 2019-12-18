package main

import (
	"fmt"
)

/*
	先写一个函数，使用单向列表
	n 小孩1---n
	编号
	数到m出列
	下一位又开始从1开始

*/

type Boy struct {
	No   int
	Next *Boy //指向下一个小孩的指针
}

//编写一个函数，构成单向环形列表，
//num：表示小孩的个数
//*Boy 返回该环形列表的第一个小孩的指针

func AddBody(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	//判断
	if num < 1 {
		//说明这个不能构建
		fmt.Println("num的值不对")
		return first
	}

	//循环构建这个环形列表
	for i := 1; i <= num; i++ {
		//第一个小孩比较特殊

		boy := &Boy{
			No: i,
		}
		//需要一个step指针，

		if i == 1 {
			//先创建一个小孩
			first = boy //一旦完成，头指针形成
			curBoy = boy
			curBoy.Next = first //形成一个圈
		} else {
			//保留first作为头指针
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

//显示单向列表
//进行链表的遍历
func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("empty")
		return
	}

	//创建一个指针帮助遍历
	curBoy := first
	for {
		//说明至少有一个小孩
		fmt.Printf("小孩编号 = %d ->", curBoy.No)
		//退出条件？
		if curBoy.Next == first {
			break
		}
		//下移动
		curBoy = curBoy.Next
	}
	fmt.Println()
}

//分析思路
//1.编写一个函数，playGamge
func PlayGame(first *Boy, startNo int, countNum int) {
	//2.使用算法，在环形列表中留下最后一个人
	//1. 空链表先处理
	if first.Next == nil {
		fmt.Println("空的列表.")
		return
	}

	//2.定义辅助指针，帮助我们删除小孩
	//3.判断startNo <= 小孩总数
	tail := first
	//4.让tail指向环形列表的最后一个小孩，他会帮助我们删掉
	for {
		if tail.Next == first {
			//说明已经到了最后的小孩了
			break
		}
		tail = tail.Next
	}
	fmt.Println(tail.No)
	//让first移动到startNo
	for i := 0; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next

	}
	//开始数countNum,然后就开始删除first 指向的小孩
	for {
		//开始数countNum - 1
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号：%d\n", first.No)
		//删除first执行的小数
		first = first.Next
		tail.Next = first
		//判断如果tail == first ,说明圈中只有一个小孩，留下这个小孩,
		if tail == first {
			break
		}

	}
	fmt.Printf("小孩编号为%d ", first.No)
}
func main() {
	first := AddBody(5)
	//显示
	ShowBoy(first)
	PlayGame(first, 2, 3)
}
