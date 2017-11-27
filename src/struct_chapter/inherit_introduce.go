package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

type NamePoint struct {
	name string
	point
}

func (p *point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (np *NamePoint) Abs() float64 {
	return np.point.Abs() * 100
}

func main() {
	n := &NamePoint{"gao", point{3, 4}}
	fmt.Println(n.Abs())
}
