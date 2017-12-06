package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cal := make([]int, 0)
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an number or quit to stop:")
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error,please check")
			os.Exit(1)
		}
		token = strings.TrimRight(token, "\r\n")
		token = string(token)
		switch {
		case token == "quit":
			fmt.Println("Stopped")
			return
		case token >= "0" && token <= "9999":
			i, _ := strconv.Atoi(token)
			cal = append(cal, i)
		case token == "+":
			a := cal[0]
			b := cal[1]
			fmt.Printf("Result: [%d]%s[%d]=%d", a, token, b, a+b)
		case token == "-":
			a := cal[0]
			b := cal[1]
			fmt.Printf("Result: [%d]%s[%d]=%d", a, token, b, a-b)
		case token == "*":
			a := cal[0]
			b := cal[1]
			fmt.Printf("Result: [%d]%s[%d]=%d", a, token, b, a*b)
		case token == "/":
			a := cal[0]
			b := cal[1]
			fmt.Printf("Result: [%d]%s[%d]=%d", a, token, b, a/b)
		default:
			fmt.Println("No valid input")
		}

	}
}
