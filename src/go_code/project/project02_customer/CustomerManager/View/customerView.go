package main

import (
	"fmt"
	model "go_code/project/project02_customer/CustomerManager/Model"
	service "go_code/project/project02_customer/CustomerManager/Service"
)

type customerView struct {
	//定义必要的字段
	key string
	//退出的loop,是否循环显示主菜单
	loop bool
	//2.增加了该字段
	customerService *service.CustomerService
}

func (this *customerView) displayMain() {
	for {
		fmt.Println("---------客户信息管理系统--------")
		fmt.Println("---------1.添加用户--------")
		fmt.Println("---------2.修改用户-------")
		fmt.Println("---------3.删除客户--------")
		fmt.Println("---------4.显示列表--------")
		fmt.Println("---------5.退出--------")
		fmt.Println("---------6.更新--------")
		fmt.Print("---请选择1-6")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			this.delete()
		case "4":
			fmt.Println("列 表 显 示")
			this.list()
		case "5":
			this.quite()
		case "6":
			this.update()
		default:
			fmt.Println("again ")
		}
		if !this.loop {
			fmt.Println("你退出了客户关系管理系统")
			break
		}
	}
}

func NewCustomerView() *customerView {
	return &customerView{
		key:  "",
		loop: true,
	}
}

//2.list
func (this *customerView) list() {
	//获取当前所有用户的客户信息
	customer := this.customerService.List()
	fmt.Println("-----------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, v := range customer {
		fmt.Println(v.GetInfo())
	}
	fmt.Println("-----------------列表完成-------------------")
}

//3.add
func (this *customerView) add() {
	fmt.Println("----------------添加客户--------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("email:")
	email := ""
	fmt.Scanln(&email)
	//ID没有让用户直接输入，因为Id是唯一的需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//需要给客户分配一个id
	if this.customerService.Add(*customer) {
		fmt.Println("good job")
	} else {
		fmt.Println("error")
	}
}

//4.delete
func (this *customerView) delete() {
	id := -1
	fmt.Println("-----------删除用户--------------")
	fmt.Println("输入客户id(放弃请输入-1):")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认删除请输入y：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，重新输入 y/n")
	}
	if choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("-----------删除成功--------------")
		} else {
			fmt.Println("id不存在")
		}
	} else {
		return
	}

}

//5.update
func (this *customerView) update() {
	fmt.Println("请输入update：id")
	id := -1
	fmt.Scanln(&id)
	if this.customerService.Update(id) {
		fmt.Println("good ")
	}
}
func (this *customerView) quite() {
	fmt.Println("确认删除请输入y：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，重新输入 y/n")
	}
	if choice == "y" || choice == "Y" {
		this.loop = false
	} else {
		return
	}
}

func main() {
	customerV := NewCustomerView()

	//2.这里完成对CustomerView结构体的customerService
	customerV.customerService = service.NewCustomerService()

	customerV.displayMain()
}
