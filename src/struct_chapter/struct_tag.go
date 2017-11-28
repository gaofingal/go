package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 int    "How much there are"
	field3 string "The name of the thing"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {
	tt := TagType{true, 12, "Gao"}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
