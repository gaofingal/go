package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//v.SetFloat(3.44)
	fmt.Println("settablility of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v：", v.CanSet())
	v = v.Elem()
	fmt.Println("the Elem of v is :", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(4.333)
	fmt.Println(v.Interface())
}
