package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outFile, err := os.OpenFile("gao.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Writting file fail!!!!")
		return
	}
	defer outFile.Close()

	outputWriter := bufio.NewWriter(outFile)
	outputString := "Hello,Golang!!!"
	outputWriter.WriteString(outputString)
	outputWriter.Flush()
}
