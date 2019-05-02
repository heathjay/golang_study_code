package main

import (
	"fmt"
)

//判断打鱼还是晒网
//中国有句古话“三天打鱼，两天晒网”，如果从1990年1月1号开始执行“三天打鱼两天晒网”如何判断接下来那一天是打鱼还是晒网

func findOut(year int, month int, day int) {
	var standardYear int = 1990
	//var standardMonth int = 1
	//var standardDay int = 1
	var YearDay int = 0
	var MonthDay int = 0
	difYear := year - standardYear
	switch difYear % 4 {
	case 0:
		YearDay = difYear * (365*3 + 366)
	case 1:
		YearDay = (difYear-1)*(365*3+366) + 365
	case 2:
		YearDay = (difYear-1)*(365*3+366) + 365 + 366
	case 3:
		YearDay = (difYear-1)*(365*3+366) + 365 + 366 + 365
	}
	difMonth := month - 1
	switch difMonth {
	case 0:
		MonthDay = 0
	case 1:
		MonthDay = 31
	case 2:
		MonthDay = 61
	case 3:
		MonthDay = 92
	case 4:
		MonthDay = 122
	case 5:
		MonthDay = 153
	case 6:
		MonthDay = 183
	case 7:
		MonthDay = 214
	case 8:
		MonthDay = 245
	case 9:
		MonthDay = 275
	case 10:
		MonthDay = 305
	case 11:
		MonthDay = 336
	}

	days := YearDay + MonthDay + day - 1
	key := days % 5
	var flag string
	switch key {
	case 1, 2, 3:
		flag = "fishing"
	case 0, 4:
		flag = "neting"
	}
	fmt.Println(flag)
}
func main() {
	var year int
	var month int
	var day int
	fmt.Println("please input year month day:")
	fmt.Scanf("%d %d %d", &year, &month, &day)
	findOut(year, month, day)
}
