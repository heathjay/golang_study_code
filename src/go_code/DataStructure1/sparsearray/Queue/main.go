package main

import (
	"errors"
	"fmt"
	"os"
)

/*
	创建一个数组 array，作为Queue的一个字段
	front = rear = -1
	基本操作:
		AddQueue:
		GetQueue
		ShowQueue
*/

type Queue struct {
	Array   [5]int
	Rear    int
	Front   int
	MaxSize int
}

func (this *Queue) AddQueue(data int) (err error) {
	if this.Rear == this.MaxSize-1 {

		return errors.New("Queue full")
	}
	this.Array[this.Rear+1] = data
	this.Rear += 1
	return
}

func (this *Queue) ShowQueue() {
	//front指向队首但不含队首
	fmt.Println("当前队列")
	for i := this.Front + 1; i <= this.Rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.Array[i])
	}
	fmt.Println()
}

func (this *Queue) GetQueue() (data int, err error) {
	if this.Front == this.Rear {
		fmt.Println("为空")
		err = errors.New("queue is empty")
		return -1, err
	}
	this.Front++
	data = this.Array[this.Front]
	return data, err

}

func main() {
	//先创建一个队列
	queue := Queue{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
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
			err := queue.AddQueue(data)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			data, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("get:", data)
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(-1)
		}

	}

}
