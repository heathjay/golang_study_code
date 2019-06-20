package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

type work_unix struct {
	start uint32
	end   uint32
	v     uint32
	dist  *[]uint32
	//exit  *[]chan bool
	wg *[]sync.WaitGroup
	//cpus int
	exit chan bool
}

func FindPath(tw work_unix) {
	var i, j, k uint32
	v := tw.v
	start := tw.start
	end := tw.end
	dist := tw.dist
	//exit := tw.exit
	wg := tw.wg

	// cpus := tw.cpus

	for k = 0; k < v; k++ {
		if k != 0 {
			(*wg)[k-1].Wait()
		}
		go func() {
			for i = start; i < end; i++ {
				if k != i {
					tempk := (*dist)[int(i*v+k)]
					if tempk != math.MaxUint32 {
						for j = 0; j < v; j++ {
							if j != i && j != k {
								tempj := (*dist)[int(k*v+j)]
								if tempj != math.MaxUint32 {
									inter := tempj + tempk
									now := (*dist)[int(i*v+j)]
									if inter < now {
										(*dist)[int(i*v+j)] = inter
									}
								}
							}
						}
					}

				}
			}
			(*wg)[k].Done()
		}()

	}
	tw.exit <- true

}

func amd(dist *[]uint32, v int) {
	var sum, path uint32 = 0, 0
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			t := (*dist)[i*v+j]
			if t != math.MaxUint32 && i != j {
				sum += t
				path++
			}
		}
	}

	fmt.Println(sum / path)
}
func main() {
	//startTime := time.Now().UnixNano()

	cpus := runtime.NumCPU() * 20
	//cpus := 2

	//fmt.Println("cpus=", cpus)
	runtime.GOMAXPROCS(cpus)
	var v, e uint32
	fmt.Scanf("%d %d", &v, &e)
	var i uint32

	wg := make([]sync.WaitGroup, v)
	//var wg sync.WaitGroup
	exit := make(chan bool, cpus+1)
	dist := make([]uint32, v*v)
	for i = 0; i < v*v; i++ {
		dist[i] = math.MaxUint32
	}

	for i = 0; i < v; i++ {
		dist[i*v+i] = 0
		//exitChan[i] = make(chan bool, cpus+1)
		wg[i].Add(cpus + 1)

	}

	var source, dest, cost uint32

	for i = 0; i < e; i++ {
		fmt.Scanf("%d %d %d", &source, &dest, &cost)
		if cost < dist[source*v+dest] {
			dist[source*v+dest] = cost
		}
	}

	work := make([]work_unix, cpus+1)
	size := int(v) / cpus

	for j := 0; j < cpus+1; j++ {
		work[j].dist = &dist
		work[j].v = v
		work[j].start = uint32(size * j)
		work[j].wg = &wg
		work[j].exit = exit
		if j == cpus {
			work[j].end = v
		} else {
			work[j].end = uint32(size * (j + 1))
		}
		go FindPath(work[j])
		// go func(work work_unix) {
		// 	FindPath(work[j])
		// }()
	}
	for j := 0; j < cpus+1; j++ {
		<-exit
	}

	// FindPath(temp)
	amd(&dist, int(v))
	//endTime := time.Now().UnixNano()
	//fmt.Println("cpus=", cpus, "time=", endTime-startTime)
}
