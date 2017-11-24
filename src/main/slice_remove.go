package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d", "e", "f"}

	res := remove(s, 2, 4)
	fmt.Println(res)
}

func remove(s []string, start int, end int) (res []string) {
	remove_amount := end - start
	new_slice_len := len(s) - remove_amount

	res = make([]string, new_slice_len)
	at := copy(res, s[:start])
	copy(res[at:], s[end:])
	return
}
