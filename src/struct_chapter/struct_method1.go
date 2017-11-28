package main

import "fmt"

//  Invector 是[]int的alias
type Invector []int // alias type

func (v Invector) Sum() int {
	var s int
	for _, x := range v {
		s += x
	}
	return s
}

func main() {
	ix := Invector{1, 2, 3, 45, 65}
	fmt.Println(ix.Sum())

}
