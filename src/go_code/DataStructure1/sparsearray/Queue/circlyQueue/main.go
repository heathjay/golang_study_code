package main

import (
	"errors"
	"fmt"
	"os"
)

/*
	无法利用数组空间，一次性
	使用数组实现一个环形队列
	通过取模的方式去实现
	1. 队列容量需要空一个作为约定，在队列堆满的时候
	2. head == tail为空
	3. (tail+1)%maxSize == head 为满
	4. 初始化，tail = head =0
	5. 统计该队列有多少个元素 (tail + maxSize - head ) % maxSize
*/
type CircleQueue struct {
	Head    int
	Tail    int
	MaxSize int
	Array   [5]int
}

func (this *CircleQueue) Push(data int) (err error) {
	if this.IsFull() {
		err = errors.New("Queue is full")
		return
	}

	this.Array[this.Tail] = data //直接把值放给尾部,tail 并没有包含最后一个值
	this.Tail = (this.Tail + 1) % this.MaxSize
	return
}

func (this *CircleQueue) Pop() (data int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("queue is empty")
	}
	data = this.Array[this.Head]
	this.Head = (this.Head + 1) % this.MaxSize
	return
}

func (this *CircleQueue) IsFull() bool {
	return (this.Tail+1)%this.MaxSize == this.Head
}

func (this *CircleQueue) IsEmpty() bool {
	return this.Tail == this.Head
}

func (this *CircleQueue) Size() int {
	return (this.Tail + this.MaxSize - this.Head) % this.MaxSize
}
func (this *CircleQueue) ShowCirleQueue() {
	//取出当前队有多少缘故
	size := this.Size()
	if size == 0 {
		fmt.Println("queue is empty")
	}
	fmt.Println("queue is ")
	//设计一个辅助变量
	tempHead := this.Head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.Array[tempHead])
		tempHead = (tempHead + 1) % this.MaxSize
	}
	fmt.Println()
}
func main() {
	queue := CircleQueue{
		MaxSize: 5,
		Head:    0,
		Tail:    0,
	}
	var key string
	var data int
	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2. 输入get表示从队列中获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入Eixt退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要如队列的数:")
			fmt.Scanln(&data)
			err := queue.Push(data)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			data, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("get:", data)
		case "show":
			queue.ShowCirleQueue()
		case "exit":
			os.Exit(-1)
		}

	}
}
