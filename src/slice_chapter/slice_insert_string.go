package main

import "fmt"

func main() {
	s := []string{"a", "a", "a", "a", "a", "a"}
	in := []string{"b", "b"}

	res := insertStringSlice(0, s, in)
	fmt.Println(res)

}

func insertStringSlice(index int, dst []string, sou []string) []string {
	res := make([]string, len(dst)+len(sou))
	at := copy(res, dst[:index])
	at += copy(res, sou)
	copy(res[at:], dst[index:])
	return res
}
