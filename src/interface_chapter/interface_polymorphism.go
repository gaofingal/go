package main

import "fmt"

type valuable interface {
	getValue() float32
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

func showValue(v valuable) {
	fmt.Printf("value of the asset is %f\n", v.getValue())
}

func main() {
	var o valuable = stockPosition{"Golang", 99.9, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}
