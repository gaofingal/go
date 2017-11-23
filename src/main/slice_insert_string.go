package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{7, 8}
	n := 0

	s3 := insertStringSlice(n, s1, s2)
	for _, v := range s3 {
		fmt.Println(v)
	}
}

func insertStringSlice(n int, dst []int, sou []int) []int {
	if n == 0 {
		return sou
	} else if n == len(dst) {
		copy(dst, sou)
		return dst
	} else {
		c1 := dst[0:n]
		c2 := dst[n:]
		copy(c1, sou)
		copy(c1, c2)
		return c1

	}

}
