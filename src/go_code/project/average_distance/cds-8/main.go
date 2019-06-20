package main

import (
	"fmt"
	"math"
	"sync"
)

/*
func FindPath(dist []uint32, v uint32) {
	var k, m uint32
	var cpus uint32 = 150
	size := v / cpus
	var wg sync.WaitGroup

	for k = 0; k < v; k++ {
		for m = 0; m < cpus; m++ {
			wg.Add(1)
			go loop(k, m*size, (m+1)*size, &wg, dist, v)
		}
		wg.Add(1)
		go loop(k, m*size, v, &wg, dist, v)
		wg.Wait()
	}

}
*/
func loop(k uint32, start uint32, end uint32, wg *sync.WaitGroup, dist []uint32, v uint32) {
	go func() {
		var i, j uint32
		(*wg).Done()
		for i = start; i < end; i++ {
			// if k != i {
			// 	tempk := dist[i*v+k]
			// 	if tempk != math.MaxUint32 {
			// 		for j = 0; j < v; j++ {
			// 			if j != i && j != k {
			// 				tempj := dist[k*v+j]
			// 				if tempj != math.MaxUint32 {
			// 					inter := tempj + tempk
			// 					now := dist[i*v+j]
			// 					if inter < now {
			// 						dist[i*v+j] = inter
			// 					}
			// 				}
			// 			}
			// 		}
			// 	}

			// }
			for j = 0; j < v; j++ {
				intermediary := dist[i*v+k] + dist[k*v+j]
				//check for overflows
				if (intermediary >= dist[i*v+k]) && (intermediary >= dist[k*v+j]) && (intermediary < dist[i*v+j]) {
					dist[i*v+j] = dist[i*v+k] + dist[k*v+j]
				}
			}
		}
	}()
}

func amd(dist []uint32, v int) {
	var sum, path uint32 = 0, 0
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			//t := (*dist)[i*v+j]
			if dist[i*v+j] != math.MaxUint32 && i != j {
				sum += dist[i*v+j]
				path++
			}
		}
	}

	solution := sum / path
	fmt.Println(solution)
}

func main() {

	var v, e uint32
	fmt.Scanf("%d %d", &v, &e)

	var i uint32
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

	var k, m uint32
	var cpus uint32 = 150
	size := v / cpus
	var wg sync.WaitGroup

	for k = 0; k < v; k++ {
		for m = 0; m < cpus; m++ {
			wg.Add(1)
			go loop(k, m*size, (m+1)*size, &wg, dist, v)
		}
		wg.Add(1)
		go loop(k, m*size, v, &wg, dist, v)
		wg.Wait()
	}
	//FindPath(dist, v)
	amd(dist, int(v))

}
