package main

import "fmt"

type Employee struct {
	name   string
	salary *Salary
}

type Salary struct {
}

func main() {
	em1 := &Employee{"gao", &Salary{}}
	salary := em1.salary.getSalarySize()
	fmt.Printf("%s 的薪水是 %f", em1.name, salary)
}

func (sa *Salary) getSalarySize() float64 {
	var salary float64
	salary = 4000.00
	return salary
}
