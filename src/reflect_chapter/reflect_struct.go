package main

import (
	"fmt"
	"reflect"
)

type NotknowType struct {
	s1, s2, s3 string
}

func (n NotknowType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = NotknowType{"Go", "PHP", "Java"}

func main() {
	value := reflect.ValueOf(secret) // Go - PHP - Java
	typ := reflect.TypeOf(secret)    // main.NotknowType
	kind := value.Kind()             // struct
	fmt.Println(value)
	fmt.Println(typ)
	fmt.Println(kind)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d:%v \n", i, value.Field(i))
	}
	result := value.Method(0).Call(nil)
	fmt.Println(result)
}
