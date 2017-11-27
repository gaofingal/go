package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p Person) (reP Person) {
	reP.firstName = strings.ToUpper(p.firstName)
	reP.lastName = strings.ToUpper(p.lastName)
	return
}
func upPersonV2(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 方式1
	var p1 Person
	p1.firstName = "Jon"
	p1.lastName = "noe"
	pp1 := upPerson(p1)
	fmt.Printf("struct 1: %v", pp1)
	fmt.Println()

	// 方式2
	p2 := new(Person)
	p2.firstName = "Paul"
	p2.lastName = "harden"
	pp2 := upPerson(*p2)
	fmt.Printf("struct 1: %v", pp2)
	fmt.Println()

	// 方式3
	p3 := &Person{firstName: "gao", lastName: "faddei"}
	pp3 := upPerson(*p3)
	fmt.Printf("struct 1: %v", pp3)

}
