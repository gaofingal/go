package main

import (
	"bufio"
	"fmt"
	"os"
	"pipline/sortpipline"
	"strconv"
)

func main() {
	p := CreateNetPipeline("small.int", 512, 4)
	for v := range p {
		fmt.Println(v)
	}
	writeToFile(p, "small.out")
	printFile("small.out")
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	p := sortpipline.ReaderSource(reader, -1)
	for v := range p {
		fmt.Println(v)
	}
}

func createPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	sortResult := make([]<-chan int, 0)

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)

		sortResult = append(sortResult, sortpipline.InMemSort(sortpipline.ReaderSource(bufio.NewReader(file), chunkSize)))
	}
	return sortpipline.MergeN(sortResult...)
}

func writeToFile(ch <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	sortpipline.WriterSink(writer, ch)
}

func CreateNetPipeline(fileName string, fileSize int, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount

	sortedMem := make([]<-chan int, 0)
	sortAddr := make([]string, 0)

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		chData := sortpipline.ReaderSource(bufio.NewReader(file), chunkSize)
		chSend := sortpipline.InMemSort(chData)

		addr := ":" + strconv.Itoa(7000+i)
		sortAddr = append(sortAddr, addr)
		sortpipline.WriteNetSink(addr, chSend)
	}
	for _, addr := range sortAddr {
		sortedMem = append(sortedMem, sortpipline.ReaderNetSource(addr))
	}

	return sortpipline.MergeN(sortedMem...)
}
