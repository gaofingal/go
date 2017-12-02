package main

import (
	"fmt"
	"math"
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

	switch t := areaInfo.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
