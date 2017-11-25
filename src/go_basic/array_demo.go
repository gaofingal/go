package main

import "fmt"

func main() {
	/*
		【Go语言数组语法格式】
		var variable_name [size] variable_type
		数组长度必须是整数且大于0,
		初始化数组 var balance = [5]float32{1.3,1.3,1.3,1.3}  {}中的个数不能大于[]中的数字
		动态大小 var balance = [...]float32{1.3,1.3,1.3,1.3,1.3}
	*/

	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; i < 10; j++ {
		fmt.Printf("Element[%d] = %d", j, n[j])
	}
}
