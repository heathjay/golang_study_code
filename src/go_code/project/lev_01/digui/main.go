package main

import "fmt"

func edit(s1 string, s2 string, m int, n int) int {
	if m == 0 {
		return n
	} else if n == 0 {
		return m
	} else if s1[m-1] == s2[n-1] {
		return edit(s1, s2, m-1, n-1)
	} else {
		d1 := edit(s1, s2, m-1, n) + 1
		d2 := edit(s1, s2, m, n-1) + 1
		d3 := edit(s1, s2, m-1, n-1) + 1
		return min(d1, min(d2, d3))
	}

}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func main() {
	var s1 string
	var s2 string
	s1 = "Loremipsumdolorsitametconsectetur"
	s2 = "adipiscin"

	res := edit(s1, s2, len(s1), len(s2))
	fmt.Println(res)
}
