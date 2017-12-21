package main

import (
	"fmt"
	"time"
)

func pumpV4() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suckV4(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	stream := pumpV4()
	go suckV4(stream)
	time.Sleep(1e9)
}
