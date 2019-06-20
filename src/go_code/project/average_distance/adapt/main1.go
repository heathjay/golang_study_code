package main

import (
	"fmt"
	"math"
	"time"
)

type work_unix struct {
	start uint32
	end   uint32
	v     uint32
	dist  *[]uint32
	k     int
	exit  chan bool
}

func FindPath(tw *work_unix) {
	var i, j uint32
	k := tw.k
	v := tw.v

	//for k = 0; k < tw.v; k++ {
	for i = tw.start; i < tw.end; i++ {
		if k != int(i) {
			tempk := (*tw.dist)[int(i*v)+k]
			if tempk != math.MaxUint32 {
				for j = 0; j < tw.v; j++ {
					if j != i && int(j) != k {
						tempj := (*tw.dist)[k*int(v)+int(j)]
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
	//}

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
	startTime := time.Now().UnixNano()
	var v uint32
	//	var e uint32
	// fmt.Scanln(&v)
	// fmt.Scanln(&e)
	v = 10
	//	e = 7
	var i uint32
	dist := make([]uint32, v*v)
	for i = 0; i < v*v; i++ {
		dist[i] = math.MaxUint32
	}

	for i = 0; i < v; i++ {
		dist[i*v+i] = 0
	}

	//var source, dest, cost uint32
	input := [7][3]uint32{{0, 4, 3}, {4, 7, 4}, {4, 8, 6}, {7, 0, 7}, {7, 4, 1}, {8, 3, 8}, {9, 3, 2}}
	for _, y := range input {
		if y[2] < math.MaxUint32 {
			dist[y[0]*v+y[1]] = y[2]
		}

	}
	// for i = 0; i < e; i++ {
	// 	fmt.Scanln(&source)
	// 	fmt.Scanln(&dest)
	// 	fmt.Scanln(&cost)
	// 	if cost < dist[source*v+dest] {
	// 		dist[source*v+dest] = cost
	// 	}
	// }

	//cpus := runtime.NumCPU()
	// temp := &work_unix{
	// 	0,
	// 	v,
	// 	v,
	// 	&dist,
	// }

	//cpus := runtime.NumCPU()
	//cpus := 1
	cpus := 2
	//cpus := 10

	work := make([]work_unix, cpus+1)
	size := int(v) / cpus
	for j := 0; j < cpus; j++ {
		work[j].dist = &dist
		work[j].v = v
		work[j].start = uint32(size * j)
		work[j].end = uint32(size * (j + 1))
	}
	work[cpus].dist = &dist
	work[cpus].start = uint32(size * cpus)
	work[cpus].end = v
	work[cpus].v = v
	exitChan := make([]chan bool, v)

	for k := 0; k < int(v); k++ {
		exitChan[k] = make(chan bool, cpus+1)
		for j := 0; j < cpus+1; j++ {
			work[j].k = k
			work[j].exit = exitChan[k]
			go FindPath(&work[j])

		}
		for n := 0; n < cpus+1; n++ {
			<-exitChan[k]
		}
		close(exitChan[k])

	}
	// FindPath(temp)
	amd(&dist, int(v))
	endTime := time.Now().UnixNano()
	fmt.Println("cpus=", cpus, "time=", endTime-startTime)
}
