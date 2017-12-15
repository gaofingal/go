package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("gao.txt")
	if err != nil {
		fmt.Printf("openning the file fail!!!")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, err := inputReader.ReadString('\n')
		fmt.Printf("the input was: %s", inputString)
		if err == io.EOF {
			return
		}
	}
}
