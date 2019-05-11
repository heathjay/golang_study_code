package main

import "fmt"

func main() {
	//map 管道
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"

	m2 := make(map[string]string, 5)
	m2["bb"] = "baba"
	m2["mm"] = "mother"

	mapChan <- m1
	mapChan <- m2

	//m3 := make(chan map[string]string, 5)
	m3 := <-mapChan
	fmt.Println(m3)

}
