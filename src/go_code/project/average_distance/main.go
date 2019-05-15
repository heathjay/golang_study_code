package main

import (
	"fmt"
	"math"
)

func Cal_dist(dist_slice *[]float64, vert int) {
	fmt.Println("*******cal_dis******")
	//cal_dist of dist[i][j]
	for i := 0; i < vert; i++ {
		for k := 0; k < vert; k++ {
			for j := 0; j < vert; j++ {
				inter := (*dist_slice)[i*vert+k] + (*dist_slice)[k*vert+j]
				if inter >= (*dist_slice)[i*vert+k] && inter >= (*dist_slice)[k*vert+j] && inter < (*dist_slice)[i*vert+j] {
					(*dist_slice)[i*vert+j] = inter
				}

			}
		}
	}

}

func main() {
	var vert, edge int
	fmt.Println("please input the vertices:")
	fmt.Scanln(&vert)
	fmt.Println("please input the edges:")
	fmt.Scanln(&edge)

	var dist_slice []float64
	dist_slice = make([]float64, vert*vert)
	for i := 0; i < vert*vert; i++ {
		dist_slice[i] = math.Inf(1)
	}

	for i := 0; i < vert; i++ {
		dist_slice[i*vert+i] = 0
	}

	fmt.Println("please input your nodes diagram + cost")
	for i := 0; i < edge; i++ {
		var source, dest int
		var cost float64
		fmt.Println("source:  destination: cost: (source[0:vert)) destination[0:vert)")
		fmt.Scanf("%d %d %f", &source, &dest, &cost)
		// fmt.Println("source:", source)
		if dist_slice[source*vert+dest] > cost {
			dist_slice[source*vert+dest] = cost
		}
	}
	fmt.Println(dist_slice)
	Cal_dist(&dist_slice, vert)
	fmt.Println(dist_slice)
}
