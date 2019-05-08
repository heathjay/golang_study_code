package utils

import (
	"fmt"
)

//必要字段
//就是test
type FamilyAccount struct {
	name string
	pwd  string
	//声明一个变量接收用户输入的选项
	key string
	//声明一个变量来控制是否退出for循环
	loop bool
	//显示这个菜单

	//定义账户余额 balance string
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定一个变量，记录是否有收支的行为
	flag bool
	//收支的一个详情，使用一个字符串来记录
	details string
}

//工厂方法
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说	明",
	}
}

func (this *FamilyAccount) certificate() {
	for {
		fmt.Println("input your name:")
		fmt.Scanln(&this.name)
		fmt.Println("input pwd")
		fmt.Scanln(&this.pwd)
		if this.name == "jiang" && this.pwd == "888" {
			break
		}
	}
}

func (this *FamilyAccount) DisplayMain() {
	this.certificate()
	for {
		fmt.Println("------------家庭收支记账软件-----------")
		fmt.Println("			1.收支明细")
		fmt.Println("			2.登记收入")
		fmt.Println("			3.登记支出")
		fmt.Println("			4.transfor")
		fmt.Println("			5.退出软件")
		fmt.Println("请选择1-4")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.DisplayDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.translate()
		case "5":
			this.quite()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
}

func (this *FamilyAccount) translate() {
	fmt.Println("input money")
	fmt.Scanln(&this.money)
	if this.money < this.balance {
		this.balance -= this.money
		fmt.Println("本次支出的说明:")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
		this.flag = true
	} else {
		fmt.Println("bad job")
	}
}
func (this *FamilyAccount) DisplayDetails() {
	fmt.Println("------------当前支付明细记录-----------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("来一笔支出吧")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入的金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入的说明:")
	fmt.Scanln(&this.note)
	//将这笔收入情况，记录到details这个变量中去
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("登记支出..")
	fmt.Println("本次支出的金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("bad job")
		// return
	} else {
		this.balance -= this.money
		fmt.Println("本次支出的说明:")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
		this.flag = true
	}
}

func (this *FamilyAccount) quite() {
	//loop = false
	fmt.Println("确定推出吗yes/no")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，重新输入 y/n")
	}

	if choice == "y" {
		this.loop = false
	}

}
