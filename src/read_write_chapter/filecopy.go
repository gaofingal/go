package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	rs, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Println("Copy error")
		return
	}
	fmt.Println("Copy done:", rs)
}
