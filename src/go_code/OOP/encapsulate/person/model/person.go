package model

import "fmt"

type person struct {
	Name   string
	age    int //不可导出
	income float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了能访问age和income编写get set方法

func (person *person) SetAge(age int) {
	if age > 0 && age < 150 {
		person.age = age
	} else {
		fmt.Println("age error")
	}
}

func (person *person) GetAge() int {
	return person.age
}

func (person *person) SetIncome(income float64) {
	if income > 3000 && income < 30000 {
		person.income = income
	} else {
		fmt.Println("income error")
	}
}

func (person *person) GetIncome() float64 {
	return person.income
}
