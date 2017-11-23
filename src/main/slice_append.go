package main

import (
	"fmt"
)

func main() {
	var s = []int{1, 2, 3, 4}
	var factory int = 6

	s1 := make([]int, factory)
	n := copy(s, s1)
	for _, item := range s1 {
		fmt.Println(item)
	}
	fmt.Printf("length is %d", n)
}
