package main

import (
	"math"
	"fmt"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
func main() {
	var areaInfo Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaInfo = sq1
	if t, ok := areaInfo.(*Square); ok {
		fmt.Printf("The type of areaInfo is :%T\n", t)
	}
	if u, ok := areaInfo.(*Circle); ok {
		fmt.Printf("The type of areaInfo is :%T\n", u)
	} else {
		fmt.Println("areainfo does not contain a variable of type Circle")
	}
}
