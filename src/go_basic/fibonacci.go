package main

import "fmt"

// 用闭包实现斐波那契
func main() {

	var i int
	f := fibonacci1()
	for i = 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

func fibonacci1() func(a int) int {
	var res int
	var last int
	var last_last int
	return func(a int) int {
		if a < 2 {
			res = a
		} else {
			res = last_last + last
		}
		last_last = last
		last = res
		return res
	}
}
