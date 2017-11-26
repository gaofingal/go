package main

import "fmt"

type Rectangle struct {
	width int
	heigh int
}

func (r Rectangle) area() int {
	var area int
	w := r.width
	h := r.heigh

	area = w * h
	return area
}

func (r Rectangle) primeter() int {
	var primeter int
	w := r.width
	h := r.heigh
	primeter = (w + h) * 2
	return primeter
}

// 定义矩形结构体，其中包括求周长面积的方法
func main() {
	var r1 Rectangle
	r1.heigh = 7
	r1.width = 9

	area := r1.area()
	primete := r1.primeter()
	fmt.Printf("面积为%d  周长为%d",area,primete)
}
