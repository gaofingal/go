package main

import "fmt"

func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

func numEchoRange(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func main() {
	numChan := make(chan int)
	done := make(chan bool)

	// 提供数字
	go numGen(0, 10, numChan)
	// 打印数字
	go numEchoRange(numChan, done)
	// main() 等待两个协程结束
	<-done
}
