package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int     // 容量
	Top    int     //表示栈顶，因为栈顶固定
	arr    [20]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	//先判断是否满
	if this.Top == this.MaxTop-1 {
		err = errors.New("full")
		fmt.Println("full")
		return err
	}

	this.Top++
	this.arr[this.Top] = val
	return
}

func (this *Stack) Pop() (val int, err error) {
	//先判断是否满
	if this.Top == -1 {
		err = errors.New("empty")
		fmt.Println("empty")
		return 0, err
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func (this *Stack) Show() {
	//遍历栈的时候需要从栈顶开始遍历
	if this.Top == -1 {
		fmt.Println("empty")
		return
	}
	curTop := this.Top
	for {
		if curTop == -1 {
			break
		}
		fmt.Println(this.arr[curTop])
		curTop--
	}

}

//is it */ 只做加减乘除
//用ascci码来做
//* 41
//+ 43
//- 45
// / 47
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算 方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("operation error:", oper)
	}
	return res
}

//优先定义 【* /】 = 1 ，【+-】 = 0
func (this *Stack) Priority(opt int) int {
	res := 0
	if opt == 42 || opt == 47 {
		res = 1
	} else if opt == 43 || opt == 45 {
		res = 0
	}
	return res
}
func main() {
	//1。 num stack
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//2. opt stack
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//3.index
	exp := "30+20*6-2"
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""
	for {
		ch := exp[index : index+1] //返回单个字符的"3
		temp := int([]byte(ch)[0]) //ascii
		if operStack.IsOper(temp) {
			//说明是符号
			//两个逻辑
			//需要判断优先级
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				//优先级和栈顶逻辑比较
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					//将计算结果入栈
					numStack.Push(result)
					//当前的符号押入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {
			//说明是数字
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}

		}
		//继续扫描
		if index == len(exp)-1 {
			break
		} else {
			index++
		}
	}
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		//将计算结果入栈
		numStack.Push(result)
	}
	//
	res, _ := numStack.Pop()
	fmt.Println(exp, "=", res)
}
