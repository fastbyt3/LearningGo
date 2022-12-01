package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(x float64) {
	p.X *= x
	p.Y *= x
}

func main() {
	p := Point{X: 10, Y: 15}
	fmt.Printf("Point: %v\n", p)

	// (&p).ScaleBy(2) <- below ins is translated to this
	p.ScaleBy(2)
	fmt.Printf("Point after scale by 2: %v\n", p)
}
