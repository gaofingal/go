package main

import (
	"../hello"
	"fmt"
)

//  使用外部调用类型

func main() {
	st := new(hello.OutsideType)
	st.C = 4.4

	str := st.Get()
	fmt.Println(str)
}
