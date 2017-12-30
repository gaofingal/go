package main

import (
	"bufio"
	"fmt"
	"os"
	"pipline/sortpipline"
)

func main() {
	const (
		filename   = "small.int"
		chunkCount = 64
	)

	file, err := os.Create(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	res := sortpipline.RandomSource(chunkCount)
	write := bufio.NewWriter(file)
	sortpipline.WriterSink(write, res)
	write.Flush()

	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	chData := sortpipline.InMemSort(sortpipline.ReaderSource(bufio.NewReader(readFile), -1))
	count := 0
	for v := range chData {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}

}

func mergeDemo() {
	res := sortpipline.Merge(sortpipline.InMemSort(sortpipline.ArraySource(2, 5, 4, 3)),
		sortpipline.InMemSort(sortpipline.ArraySource(1, 7, 8, 6)))
	for val := range res {
		fmt.Println(val)
	}
}
