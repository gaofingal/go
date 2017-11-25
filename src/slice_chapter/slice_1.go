package main

import "fmt"

func main() {
	str := "google"
	str1, str2 := splitByIndex(str, 2)
	fmt.Printf("first %s   second %s", str1, str2)

	fmt.Println()

	str3 := splitAndMerge(str)
	fmt.Println(str3)
}

func splitByIndex(str string, index int) (string, string) {
	return str[:index], str[index:]
}

func splitAndMerge(str string) string {
	return str[len(str)/2:] + str[:len(str)/2]
}
