package model

import "fmt"

type account struct {
	name    string
	pwd     string
	balance float64
}

func NewAccount(name string, pwd string, balance float64) *account {
	if len(name) < 6 || len(name) > 10 {
		fmt.Println("账号长度不对")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度不对")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额不对")
		return nil
	}
	return &account{
		name:    name,
		pwd:     pwd,
		balance: balance,
	}
}

func (account *account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("密码长度不对")
	} else {
		account.pwd = pwd
	}
}

func (account *account) GetPwd() string {
	return account.pwd
}

func (account *account) SetName(name string) {
	if len(name) < 6 || len(name) > 10 {
		fmt.Println("账号长度不对")
	} else {
		account.name = name
	}
}

func (account *account) GetName() string {
	return account.name
}
