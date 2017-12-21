package main

import (
	"fmt"
	"time"
)

func pumpV2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch1 := make(chan int)
	go pumpV2(ch1)
	go suck(ch1)
	time.Sleep(1e9)
}
