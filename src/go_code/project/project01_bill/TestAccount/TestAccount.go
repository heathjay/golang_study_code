package main

import (
	"fmt"
)

func main() {

	//声明一个变量接收用户输入的选项
	key := ""
	//声明一个变量来控制是否退出for循环
	// loop := true
	//显示这个菜单

	//定义账户余额 balance string
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定一个变量，记录是否有收支的行为
	flag := false
	//收支的一个详情，使用一个字符串来记录
	details := "收支\t账户金额\t收支金额\t说	明"
label:
	for {
		fmt.Println("------------家庭收支记账软件-----------")
		fmt.Println("			1.收支明细")
		fmt.Println("			2.登记收入")
		fmt.Println("			3.登记支出")
		fmt.Println("			4.退出软件")
		fmt.Println("请选择1-4")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("------------当前支付明细记录-----------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("来一笔支出吧")
			}
		case "2":
			fmt.Println("本次收入的金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入的说明:")
			fmt.Scanln(&note)
			//将这笔收入情况，记录到details这个变量中去
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("登记支出..")
			fmt.Println("本次支出的金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("bad job")
				break
			}
			balance -= money
			fmt.Println("本次支出的说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "4":
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
				break label
			} else {
				continue
			}

		default:
			fmt.Println("请输入正确的选项")
		}
		// if !loop {
		// 	break
		// }
	}

	fmt.Println("你已经退出该菜单")
}
