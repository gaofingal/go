package main

import "fmt"

type Simpler interface {
	Get() int
	Set(n int)
}

type simple struct {
	n int
}

func (s *simple) Get() int {
	return s.n
}

func (s *simple) Set(n int) {
	s.n = n
}

func getSimper(s Simpler, n int) int {
	s.Set(n)
	a := s.Get()
	return a
}

func main() {
	var si Simpler = &simple{}
	a := getSimper(si, 9)
	fmt.Println(a)
}
