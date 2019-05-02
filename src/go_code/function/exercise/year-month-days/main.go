package main

import (
	"errors"
	"fmt"
)

//题目1
//循环打印输入的月份天数【用continue实现】
//要有判断输入的月份是否错误语句

func checkError(month int) (err error) {
	if month <= 12 && month >= 1 {
		return nil
	} else {
		return errors.New("月份输入错误...")
	}
}

func showDays(month int, year int) (days int) {

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if year%4 == 0 && year%400 != 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}
func main() {
	var month int
	var year int
	for {
		fmt.Println("please input the year")
		fmt.Scanln(&year)
		fmt.Println("please input the month")
		fmt.Scanln(&month)
		err := checkError(month)
		if err != nil {
			fmt.Printf("err:%v", err)
			continue
		} else {
			days := showDays(month, year)
			fmt.Printf("input time:%v-%v-%v", month, year, days)
		}

	}
}
