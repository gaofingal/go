package main

import "fmt"

type Stringer interface {
	makeString() string
}

type struct1 struct {
}

func (s1 *struct1) makeString() string {
	return "this is struct1"
}

func main() {
	var v Stringer
	v = new(struct1)
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements Stringer:%s\n",sv.makeString())
	}
}
