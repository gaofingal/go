package main

import "fmt"

func main() {
	/*
		【Go 语言常量】
			const [type] = value
		【备注】可以省略类型说明，编译器会自动推断类型
		定义方式
			一：const WIDTH int = 20
			二：const HEIGHT  = 40
		多个相同类型的声明
			const c_name1,c_name2,c_name3 = c1,c2,c3
		常量还可以使用枚举
			const(
					Unknown = 0
					Female = 1
					Male = 2
				)
		iota：特殊的常量，可以认为是一个被编译器修改的常量
			在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
			const( a=iota;b=iota;c=iota) 相当于 const(a=iota;b;c)
	*/

	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
}
