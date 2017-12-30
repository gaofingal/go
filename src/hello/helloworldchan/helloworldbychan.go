package main

import "fmt"

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go printHelloWorld(ch, i)
	}
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
func printHelloWorld(ch chan string, i int) {
	for {
		ch <- fmt.Sprintf("Hello world from goroutine %d!\n", i)
	}
}
