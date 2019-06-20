package main

import (
	"fmt"
	"math"
	"runtime"
)

type work_unix struct {
	start uint32
	end   uint32
	v     uint32
	dist  *[]uint32
	k     uint32
	exit  chan bool
	//wg *sync.WaitGroup
}

func FindPath(tw *work_unix) {
	var i, j uint32
	k := tw.k
	v := tw.v
	//wg := tw.wg

	//defer wg.Done()

	go func() {
		for i = tw.start; i < tw.end; i++ {
			if k != i {
				tempk := (*tw.dist)[int(i*v+k)]
				if tempk != math.MaxUint32 {
					for j = 0; j < tw.v; j++ {
						if j != i && j != k {
							tempj := (*tw.dist)[int(k*v+j)]
							if tempj != math.MaxUint32 {
								inter := tempj + tempk
								now := (*tw.dist)[int(i*v+j)]
								if inter < now {
									(*tw.dist)[int(i*v+j)] = inter
								}
							}
						}
					}
				}

			}
		}
		tw.exit <- true
	}()

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
	fmt.Println(sum)
	fmt.Println("path=", path)
	solution := sum / path
	fmt.Println(solution)
}
func main() {

	var v, e uint32
	fmt.Scanf("%d %d", &v, &e)
	var i, k uint32
	dist := make([]uint32, v*v)
	for i = 0; i < v*v; i++ {
		dist[i] = math.MaxUint32
	}

	for i = 0; i < v; i++ {
		dist[i*v+i] = 0
	}

	var source, dest, cost uint32
	for i = 0; i < e; i++ {
		fmt.Scanf("%d %d %d", &source, &dest, &cost)

		if cost < dist[source*v+dest] {
			dist[source*v+dest] = cost
		}
	}

	cpus := runtime.NumCPU() * 50

	work := make([]work_unix, cpus+1)
	size := int(v) / cpus
	//ar wg sync.WaitGroup
	//exitChan := make([]chan bool, v)

	for k = 0; k < v; k++ {

		exitChan := make(chan bool, cpus+1)
		for j := 0; j < cpus+1; j++ {

			work[j].dist = &dist
			work[j].v = v
			work[j].start = uint32(size * j)
			//work[j].wg = &wg
			work[j].exit = exitChan
			if j == cpus {
				work[cpus].end = v
			} else {
				work[j].end = uint32(size * (j + 1))
			}
			work[j].k = k

			//wg.Add(1)
			go FindPath(&work[j])

		}
		//wg.Wait()
		for n := 0; n < cpus+1; n++ {
			<-exitChan
		}
		close(exitChan)

	}

	amd(&dist, int(v))

	//fmt.Println("cpus=", cpus, "time=", endTime-startTime)
}
