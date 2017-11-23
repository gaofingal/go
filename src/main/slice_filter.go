package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 2, 4, 2, 2, 7, 8, 9}

	f := filter(s)

	fmt.Println(f(2))

}

func filter(s []int) func(int) bool {
	a := make([]int, 1)
	flag := false
	return func(i int) bool {
		for _, v := range s {
			if v == i {
				a = append(a, v)
				flag = true
			}
		}
		a = a[1:]
		for _, iv := range a {
			fmt.Println(iv)
		}
		return flag
	}
}
