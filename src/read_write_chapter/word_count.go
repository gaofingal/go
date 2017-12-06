package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	charsNumber = 0
	wordsNumber = 0
	linesNumber = 0
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input,type S to stop:")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occured:%s\n", err)
		}
		if input == "S\r\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters:%d\n", charsNumber)
			fmt.Printf("Number of words:%d\n", wordsNumber)
			fmt.Printf("Number of lines:%d\n", linesNumber)
		}
		counterNumber(input)
	}
}

func counterNumber(input string) {
	charsNumber += len(input) - 2
	wordsNumber += len(strings.Fields(input))
	linesNumber++
}
