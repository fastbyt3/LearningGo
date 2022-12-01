package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	var res float64

	for pt := range path {
		if pt > 0 {
			res += path[pt-1].Distance(path[pt])
		}
	}

	return res
}

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Printf("\nPerim: %v\n", perim)

	dist := perim.Distance()
	fmt.Printf("Distance: %v\n", dist)
}
