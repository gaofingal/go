package main

import "fmt"

type shape interface {
	Area() float32
}

type square struct {
	side float32
}

type rectangle struct {
	width  float32
	height float32
}

func (re *rectangle) Area() float32 {
	return re.height * re.width
}

func (sq *square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(square)
	sq1.side = 5

	var areaInterface shape
	areaInterface = sq1
	fmt.Printf("The square has area: %f\n", areaInterface.Area())
	fmt.Println()
	r := &rectangle{5, 3}
	q := &square{5}
	shapes := []shape{q, r}
	for n, _ := range shapes {
		fmt.Println("Shape details:", shapes[n])
		fmt.Println("Area of this shape is:", shapes[n].Area())
	}
}
