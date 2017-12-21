package main

import (
	"fmt"
	"time"
)

func main() {
	// 有缓冲和没有缓冲打印结果不同
	/* output：原因，main()已经停止
	Sending 10
	Sent: 10
	*/
	//c := make(chan int, 10)
	/* Output:
	Sending 10
	Received: 10
	Sent: 10
	*/
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("Received:", x)
	}()
	fmt.Println("Sending", 10)
	c <- 10
	fmt.Println("Sent:", 10)
}
