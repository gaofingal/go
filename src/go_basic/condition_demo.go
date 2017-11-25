package main

import (
	"fmt"
)

func main() {
	/*
		 【if 语句】
		 if 布尔表达式{
				 conditional code
			}else{
				conditional code
			}
		【switch 语句】
		switch var1{
				case val1:
					...
				case val2:
					...
				default:
					...
			}
	*/
	var a int = 10
	var grade string = "A"

	if a > 0 {
		fmt.Println("a 大于0")
	} else {
		fmt.Println("a 小于0")
	}

	switch a {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀!\n")
	case grade == "B":
		fmt.Println("良好!\n")
	case grade == "C":
		fmt.Println("及格!\n")
	case grade == "D":
		fmt.Println("不及格!\n")
	default:
		fmt.Println("差")
	}

	fmt.Printf("你的等级是%s\n", grade)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 ：%T", i)
	case int:
		fmt.Printf("x 的类型 ：%T", i)
	case float32:
		fmt.Printf("x 的类型 ：%T", i)
	case float64:
		fmt.Printf("x 的类型 ：%T", i)
	case bool, string:
		fmt.Printf("x 的类型 ：%T", i)
	}

}
