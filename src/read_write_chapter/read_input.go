package main

import "fmt"

func main() {
	var (
		firstName string
		lastName  string
		s         string
		i         int
		f         float32
		input     = "56.12 / 5212 / Go"
		format    = "%f  %d  %s"
	)
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi,%s %s", firstName, lastName)
	fmt.Scanf(format, &f, &i, &s)
	fmt.Printf("f=%f,i=%d,s=%s\n", f, i, s)
	fmt.Sscanf(input, format, &f, &i, &s)
}
