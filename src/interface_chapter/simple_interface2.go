package main

import (
	"fmt"
)

func main() {
	classifier(12, 2.34, "gao", complex(1, 4), nil, false)

	// this is not in switch-type, it can use fallthrough
	var flag int = 1
	switch flag {
	case 1:
		fmt.Printf("this is 1\n")
		fallthrough
	case 2:
		fmt.Printf("this is 2\n")
	case 3:
		fmt.Printf("this is 3\n")
	default:
		fmt.Printf("this is default")
	}
}

func classifier(items ...interface{}) {
	for k, v := range items {
		//  cannot use fallthrough,if you do,it would make error message
		switch v.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", k)
		case float64:
			fmt.Printf("Param #%d is a float\n", k)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", k)
		case string:
			fmt.Printf("Param #%d is a string\n", k)
		default:
			fmt.Printf("Param #%d is a unknow\n", k)
		}
	}
}
