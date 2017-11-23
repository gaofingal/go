package main

import "fmt"

func main() {
	var a = []int{5, 3, 4, 2, 8}
	fmt.Println(min(&a))
}

func min(a *[]int) int {
	b := *a
	var min int = b[0]
	var i int
	for i = 0; i < len(b); i++ {
		if b[i] < min {
			min = b[i]
		}
	}
	return min
}
