package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    // 容量
	Top    int    //表示栈顶，因为栈顶固定
	arr    [5]int //数组模拟栈
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

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1, //栈为空

	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Show()
	val, _ := stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	fmt.Println("弹出", val)
}
