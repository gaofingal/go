package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	haster := sha1.New()
	// Instructions one
	io.WriteString(haster, "test")
	b := []byte{}
	// Sum appends the current hash to b and returns the resulting slice.
	shaValue := haster.Sum(b)
	fmt.Printf("Result:%x \n", shaValue)
	fmt.Printf("Result:%d \n", shaValue)

	haster.Reset()
	data := []byte("Hello,Golang")
	// Instructions two
	n, err := haster.Write(data)
	if err != nil || n != len(data) {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checkSum := haster.Sum(b)
	fmt.Printf("Result: %x\n", checkSum)

}
