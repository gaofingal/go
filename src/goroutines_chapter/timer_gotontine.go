package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	boo := time.After(5e8)

	for {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-boo:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(".")
			time.Sleep(5e7)
		}
	}
}
