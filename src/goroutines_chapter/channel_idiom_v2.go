package main

import (
	"fmt"
	"time"
)

func pumpV5() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suckV5(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	suckV5(pumpV5())
	time.Sleep(1e9)
}
