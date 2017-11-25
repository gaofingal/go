package main

import (
	"fmt"
)

func main() {
	var str string = "googles"
	res := string_reverse(str)
	fmt.Println(res)

}

func string_reverse(str string) (res string) {
	runes := []rune(str)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	res = string(runes)
	return
}
