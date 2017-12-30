package sortpipline

import (
	"bufio"
	"encoding/binary"
	"io"
	"math/rand"
	"net"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(ch <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		var a []int

		for val := range ch {
			a = append(a, val)
		}

		sort.Ints(a)

		for _, val := range a {
			out <- val
		}
		close(out)
	}()

	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		res1, ok1 := <-in1
		res2, ok2 := <-in2

		for ok1 || ok2 {
			if !ok2 || (ok1 && res1 <= res2) {
				out <- res1
				res1, ok1 = <-in1
			} else {
				out <- res2
				res2, ok2 = <-in2
			}
		}

		close(out)
	}()
	return out
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				val := int(binary.BigEndian.Uint64(buffer))
				out <- val
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

func WriterSink(writer io.Writer, ch <-chan int) {
	for val := range ch {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(val))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}

func ReaderNetSource(addr string) <-chan int {
	out := make(chan int)

	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		reader := bufio.NewReader(conn)
		chVal := ReaderSource(reader, -1)
		for v := range chVal {
			out <- v
		}
		close(out)
	}()

	return out
}

func WriteNetSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listener.Close()
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSink(writer, in)
	}()
}
