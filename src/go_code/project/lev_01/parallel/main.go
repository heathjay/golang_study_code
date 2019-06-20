package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func editDist(s1 string, s2 string, m int, n int, dp [][]int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if dp[m-1][n-1] != -1 {
		return dp[m-1][n-1]
	}

	if s1[m-1] == s2[n-1] {
		dp[m-1][n-1] = editDist(s1, s2, m-1, n-1, dp)
		return dp[m-1][n-1]
	}

	dp[m-1][n-1] = 1 + min(editDist(s1, s2, m, n-1, dp), min(editDist(s1, s2, m-1, n, dp), editDist(s1, s2, m-1, n-1, dp)))
	return dp[m-1][n-1]
}

func main() {

	var s1 string
	var s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	m := len(s1)
	n := len(s2)

	// Declare a dp array which stores
	// the answer to recursive calls
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	res := editDist(s1, s2, m, n, dp)
	fmt.Println(res)

}
