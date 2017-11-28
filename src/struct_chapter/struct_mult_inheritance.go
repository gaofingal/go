package main

import "fmt"

//  多继承的方式和实现

type Musician struct {
}

type Soldier struct {
}

type Doctor struct {
}

type SupperMan struct {
	Soldier
	Doctor
	Musician
}

func (m *Musician) Sing() string {
	return "i can singing"
}
func (s *Soldier) War() string {
	return "i can use gun and kill enemy"
}

func (d *Doctor) Heal() string {
	return "i can heal your injure"
}

func main() {
	man := new(SupperMan)
	fmt.Println(man.Sing())
	fmt.Println(man.War())
	fmt.Println(man.Heal())
}
