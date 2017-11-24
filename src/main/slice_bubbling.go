package main

import "fmt"

func main() {
	number := []int{4, 5, 6, 3, 2, 9, 8, 7, 1}

	res := sortBubble(number)
	fmt.Println(res)
}

func sortBubble(ns []int) []int {
	var i int
	for i = 0; i < len(ns); i++ {
		if ns[i] > ns[i+1] {
			swap(&ns[i], &ns[i+1])
		}
	}
	return ns
}
