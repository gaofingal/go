package main

import (
	"fmt"
	"test_chapter/even"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
