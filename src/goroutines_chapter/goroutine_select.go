package main

import (
	"fmt"
	"time"
)

func chan1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func chan2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func getChan(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1:%d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2:%d\n", v)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go chan1(ch1)
	go chan2(ch2)
	go getChan(ch1, ch2)

	time.Sleep(1e9)
}
