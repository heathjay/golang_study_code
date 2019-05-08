package main

import (
	"fmt"
)

func main() {
	name := ""
	fmt.Scanf("%s", &name)
	fmt.Printf("%T %v---", name, name)
	if name != "" {
		fmt.Println("erro")
	} else {
		fmt.Println("good")
	}
}
