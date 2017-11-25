package main

import (
	"fmt"
)

func main() {
	number := []int{4, 5, 6, 3, 2, 9, 8, 7, 1}

	res := sortBubble(number)
	fmt.Println(res)
}

func sortBubble(ns []int) []int {
	var i, j, tmp int
	for i = 0; i < len(ns); i++ {
		for j = 0; j < len(ns)-1; j++ {
			if ns[j] > ns[j+1] {
				tmp = ns[j]
				ns[j] = ns[j+1]
				ns[j+1] = tmp
			}
		}
	}
	return ns
}
