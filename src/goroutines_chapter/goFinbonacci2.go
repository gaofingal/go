package main

import "fmt"

func fibonacci2(n int, ch chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
