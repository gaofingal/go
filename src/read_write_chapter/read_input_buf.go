package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		inputReader *bufio.Reader
		inputV2     string
		err         error
	)
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	inputV2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", inputV2)
	}
}
