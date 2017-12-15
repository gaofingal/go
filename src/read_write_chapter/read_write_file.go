package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
