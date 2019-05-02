package main

import (
	"fmt"
)

func outputAZ() {
	var str1 = 'A'
	var str2 = 'a'
	for {
		fmt.Printf("%c ", str1)
		str1++
		if str1 == 'Z'+1 {
			break
		}
	}
	for {
		fmt.Printf("%c ", str2)
		str2++
		if str2 == 'z'+1 {
			break
		}
	}
}
func main() {
	outputAZ()
}
