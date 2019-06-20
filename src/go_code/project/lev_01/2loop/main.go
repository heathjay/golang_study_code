package main

import (
	"fmt"
	"sync"
)

func work(dist [][]int, startX int, startY int, endX int, endY int, s1 string, s2 string, wgIn *sync.WaitGroup, s string) {

	defer wgIn.Done()
	//defer wg.Done()
	s1 = "a" + s1
	s2 = "a" + s2
	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			dif := 1
			// if i-1 < 0 || j-1 < 0 || i-1 > len(s1) || j-1 > len(s2) {
			// 	//fmt.Println("startX= ", i, "startY=", j, "endX", endX, "endY", endY)
			// 	break
			// }

			if s2[i-1] == s1[j-1] {
				dif = 0
			}

			dist[i][j] = Min(Min(dist[i-1][j]+1, dist[i][j-1]+1), dist[i-1][j-1]+dif)
		}
	}

	//defer wgIn.Done()
	//fmt.Println("s", s, "startX= ", startX, "startY=", startY, "endX", endX, "endY", endY)
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func MoveGo(cpus int, AR int, AC int, BR int, BC int, dp [][]int, xT int, yT int, s1 string, s2 string, wg *sync.WaitGroup) {
	defer wg.Done()
	var wgIn sync.WaitGroup
	//fmt.Println("AR=", AR, "AC=", AC, "BR=", BR, "BC=", BC)
	for {

		if AR <= 0 || AC <= 0 {
			break
		}
		if BR <= 0 || BC <= 0 {
			break
		}
		ACend := AC + xT
		if AC == cpus*xT+1 {
			ACend = len(s1)
		}

		ARend := AR + yT
		if AR == cpus*yT+1 {
			ARend = len(s2)
		}
		wgIn.Add(1)
		//fmt.Println("AR=", AR, "AC=", AC)
		go work(dp, AR, AC, ARend, ACend, s1, s2, &wgIn, "a")
		if AR == BR && AC == BC {
			break
		}

		AC -= xT
		AR += yT

	}
	wgIn.Wait()
}

func Distribut(s1, s2 string) int {

	m := len(s1) + 1
	n := len(s2) + 1
	cpus := 4
	if n == 1 {
		return len(s1)
	}
	if m == 1 {
		return len(s2)
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = i
	}
	for i := 0; i < m; i++ {
		dp[0][i] = i
	}

	AR := 1
	AC := 1
	BR := 1
	BC := 1
	xT := (m - 1) / cpus
	yT := (n - 1) / cpus
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		fmt.Println("AR=", AR, "AC=", AC, "BR=", BR, "BC=", BC)
		go MoveGo(cpus, AR, AC, BR, BC, dp, xT, yT, s1, s2, &wg)

		if BR == AR && BC == AC && AR != 1 {
			break
		}
		wg.Wait()
		if AC == cpus*xT+1 {
			AR += yT
		} else {
			AC += xT
		}

		if BR == cpus*yT+1 {
			BC += xT
		} else {
			BR += yT
		}

	}

	//wg.Wait()
	for i := 0; i < n; i++ {
		fmt.Println(dp[i])
	}

	return dp[len(s2)][len(s1)]
}

func main() {

	var s1 string
	var s2 string
	// fmt.Scanln(&s1)
	// fmt.Scanln(&s2)
	s1 = "Loremipsumdolorsitametconsectetur"
	s2 = "adipiscin"

	// x := s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1 + s1
	// x = x + x + x + x + x + x + x + x + x + x + x + x + x + x + x + x
	// y := s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2 + s2
	// y = y + y + y + y + y + y + y + y + y + y + y + y + y + y + y + y
	//res := Distribut(x, y)
	res := Distribut(s1, s2)
	fmt.Println(res)
}
