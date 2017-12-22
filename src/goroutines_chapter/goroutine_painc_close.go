package main

import "fmt"

func tel(ch chan int) {
	for i := 1; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

//func main() {
//	ch := make(chan int)
//	go tel(ch)
//	for {
//		if i, ok := <-ch; ok {
//			fmt.Printf("channel result:%d\n", i)
//		}
//	}
//}

func main() {
	ch := make(chan int)
	go tel(ch)
	for res := range ch {
		fmt.Printf("channle result:%d\n", res)
	}
}
