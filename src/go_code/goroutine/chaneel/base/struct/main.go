package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//struct 管道
	cat1 := Cat{Name: "goo", Age: 19}
	cat2 := Cat{Name: "jj", Age: 20}

	var structChan chan Cat
	structChan = make(chan Cat, 10)
	structChan <- cat1
	structChan <- cat2

	cat3 := <-structChan
	fmt.Println(cat3)
}
