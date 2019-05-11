package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpus := runtime.NumCPU()
	fmt.Println("现在cpus ：= ", cpus)

	//可以自己设置使用多少个cpu
	runtime.GOMAXPROCS(cpus - 1)
	fmt.Println("ok")
}
