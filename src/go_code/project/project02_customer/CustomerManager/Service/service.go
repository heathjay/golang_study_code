package service

import (
	"fmt"
	model "go_code/project/project02_customer/CustomerManager/Model"
	"strconv"
)

//该结构体完成对customer的操作，包括增删改查，
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段还可以作为新客户的id号+1
	customerNum int
}

//2.
func NewCustomerService() *CustomerService {
	//为了能让客户在切片中看到，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@163.com")
	customerService.customers = append(customerService.customers, *customer)
	return customerService
}

//2.返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//3.添加客户
//一定要用引用类型，不然只有一个客户
func (this *CustomerService) Add(customer model.Customer) bool {
	//添加的顺序就是id
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//4.删除
//根据id号进行删除，如果没有该客户返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers 切片
	for i, v := range this.customers {
		if v.Id == id {
			//找到
			index = i
		}
	}
	return index
}

//4.根据id删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	} else {
		//如何在切片里面删掉
		//从切片索引里面从0开始到index但是不包含index对应下标
		this.customers = append(this.customers[:index], this.customers[index+1:]...)
		return true
	}

}

//5.update
func (this *CustomerService) Update(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("id不存在")
		return false
	} else {

		fmt.Printf("姓名(%v)", this.customers[index].Name)
		name := ""
		fmt.Scanln(&name)
		fmt.Printf("%T---", name)
		if name != "" {
			this.customers[index].Name = name

		}

		fmt.Printf("性别(%v)", this.customers[index].Gender)
		gender := ""
		fmt.Scanln(&gender)
		if gender != "" {
			this.customers[index].Gender = gender
			fmt.Printf("good%v", gender)
		}

		fmt.Printf("年龄(%v)", this.customers[index].Age)
		age := ""
		fmt.Scanln(&age)
		if age != "" {
			this.customers[index].Age, _ = strconv.Atoi(age)
		}

		fmt.Printf("电话(%v)", this.customers[index].Phone)
		phone := ""
		fmt.Scanln(&phone)
		if phone != "" {
			this.customers[index].Phone = phone
		}
		fmt.Printf("邮箱(%v)", this.customers[index].Email)
		email := ""
		fmt.Scanln(&email)
		if email != "" {
			this.customers[index].Email = email
		}
		return true
	}
}
