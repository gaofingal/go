package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 7: "Sunday"}
)

func main() {

	//  顺序输出map中的值
	var i int = 0
	keys := make([]int, len(barVal))

	for key := range barVal {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	for _, value := range keys {
		fmt.Printf("%d = %s\n", value, barVal[value])
	}
}
