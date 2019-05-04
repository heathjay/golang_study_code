package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pwd       float64
	Balance   float64
}

//方法
//1.存款
func (account *Account) Deposite(money float64, pwd float64) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("error")
		return
	}
	//检测存款金额是否正确
	if money <= 0 {
		fmt.Println("money不正确")
	}
	account.Balance += money
	fmt.Println("good job")
}

//取款功能
func (account *Account) WithDraw(money float64, pwd float64) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("error")
		return
	}
	//检测存款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("money不正确")
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

//查询余下的钱
func (account *Account) Query(pwd float64) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("error")
		return
	}
	fmt.Println("你的账号=", account.AccountNo, "你的余额=", account.Balance)
}
func main() {
	//测试
	var myaccount = Account{
		AccountNo: "sjsjsjs",
		Pwd:       6666,
		Balance:   1000,
	}
	myaccount.Query(6666)
}
