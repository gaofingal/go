package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func main() {
	/*
			【函数定义】
		func function_name([parameter list])[return_type]{
			//TODO
		}
		func : 函数有func开始声明
		function_name : 函数名称，函数名和参数列表一起构成函数签名
		parameter_list : 参数列表，参数就像是一个占位符，当函数被调用时，你可以将值船体给参数，这个值被称为实际参数。
						参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
		return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
		函数体：函数定义的代码集合
	*/

	// 普通用法
	var x, y int = 4, 6
	ret := max(x, y)
	fmt.Printf("最大值是: %d", ret)

	// 传地址
	swap(&x, &y)

	// 函数作为值
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	//匿名函数的闭包效果
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

// 求最大值
func max(x, y int) int {
	var ret int

	if x > y {
		ret = x
	} else {
		ret = y
	}

	return ret
}

// 交换地址指针
func swap(x *int, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp
}

func getSquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

// 匿名函数 作为闭包
func getSequence() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

// 该method 属于Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return c.radius * c.radius
}
