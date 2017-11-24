package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 2, 4, 2, 2, 7, 8, 9}

	f := filter(s, even)

	fmt.Println(f)

}

func filter(s []int, fn func(int) bool) []int {
	var res []int
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
