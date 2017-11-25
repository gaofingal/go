package main

import "hello"

func main() {

	/*
		变量声明初始化
		var a string = "huhu"  // 方式一 var 变量 变量类型 = 初始化值
		var b = "faddei"  // 方式二 var 变量 = 初始化值  （会根据初始化值自动识别变量类型）
		c := "xxx"  // 方式三 变量:= 初始化值 （变量类型自动识别）
		var a,b,c = "a","b","c"  // 批量声明变量

		常用于全局变量的声明
		var(
				a = "a"
				b = "b"
				c = "c"
			)

		变量赋值
		a,b,c = "a","b","c"  // 批量赋值
		a = "a"

		知识点1：局部变量声明之后必须使用，不使用会报错
		知识点2：变量声明之后不允许重复声明
	*/

	var (
		a = "a"
		b = "b"
		c = "c"
	)

	hello.Hello(a)
	hello.World(b)
	hello.Hi(c)
}
