package main

import "fmt"

/*
	1. 定义一个链表,node
		链表增删该查
		假设加入比头节点还要小的节点，可能会少掉原来的头节点
	2.定义一个hashTable,链表数组
*/
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("%d-%s\n", this.Id, this.Name)
}

//第一个节点就直接存放东西
type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil

}
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil
	if cur == nil {
		this.Head = emp
		return
	} else {
		for {
			if cur != nil {
				if cur.Id >= emp.Id {
					break
				}
				pre = cur
				cur = cur.Next
			} else {
				break
			}
		}
		//退出时，我们看一下是否将emp添加到链表最后

		pre.Next = emp
		emp.Next = cur
		return

	}
}

func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Println(no, ":empty")
		return
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d,雇员id=%d name= %s->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//定义一个hashTable,含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}

//散列的方法
func (this *HashTable) HashFun(id int) int {
	return id % 7
}

func (this *HashTable) FindById(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

//显示所有元素
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}
func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {

		fmt.Println("========雇员系统菜单=======")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show显示所有雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 退出系统")
		fmt.Println("请输入你的选择:")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入id")
			fmt.Scanln(&id)
			fmt.Println("name:")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "exit":
		case "find":
			fmt.Println("请输入id:")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id = %d 不存在\n", id)
			} else {
				//编写一个方法显示雇员信息
				emp.ShowMe()
			}
		default:
			fmt.Println("input error")
		}
	}
}
