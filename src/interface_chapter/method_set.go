package main

import (
	"fmt"
)
//指针方法可以通过指针调用
//值方法可以通过值调用
//接收者是值的方法可以通过指针调用，因为指针会首先被解引用
//接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

type Lener interface {
	Len() int
}
type Appender interface {
	Append(int)
}
type List []int

// receiver is bare
func (l List) Len() int {
	return len(l)
}

// receiver is pointer
func (l *List) Append(val int) {
	*l = append(*l, val)
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}
func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func main() {
	// A bare value
	var lst List

	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	//CountInto(lst, 1, 10)  // if you use this, you need read the above information

	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}

