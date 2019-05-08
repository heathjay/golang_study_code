package main

import (
	"fmt"
	"go_code/project/project01_bill/utils"
)

func main() {
	fmt.Println("这是面向对象的方法")
	myAccount := utils.NewFamilyAccount()
	myAccount.DisplayMain()

}
