package main

import (
	"fmt"
	"go_code/OOP/encapsulate/account/model"
)

func main() {
	account := model.NewAccount("jjaaajjj", "666666", 40)
	if account != nil {
		fmt.Println("good job")
	} else {
		fmt.Println("error")
	}
	account.SetPwd("joa898")
	fmt.Println("pwd set：=", account.GetPwd())
	account.SetName("goodgoo")
	fmt.Println("name?=", account.GetName())
}
