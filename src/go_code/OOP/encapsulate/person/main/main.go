package main

import (
	"fmt"
	"go_code/OOP/encapsulate/model"
)

func main() {
	p := model.NewPerson("smith")
	fmt.Println(*p)
	p.SetAge(50)
	p.SetIncome(6666)
	fmt.Println(p.Name, "age=", p.GetAge(), "income=", p.GetIncome())
}
