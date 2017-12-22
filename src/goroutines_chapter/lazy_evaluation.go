package main

import "fmt"

var resume chan int

func integers() chan int {
	yeild := make(chan int)
	cont := 0
	go func() {
		for {
			yeild <- cont
			cont++
		}
	}()
	return yeild
}

func generateInteger() int {
	return <-resume
}

func main() {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}
