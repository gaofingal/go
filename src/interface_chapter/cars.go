package main

import "fmt"

type Any interface {
}

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type Cars []*Car

func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	ix := 0
	cs.Process(func(car *Car) {
		result[ix] = f(car)
		ix++
	})
	return result
}

func main() {
	//  make some cars
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	allCars := Cars{ford, bmw, merc, bmw2}

	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return car.Manufacturer == "BMW"
	})
	fmt.Printf("allNewBMWs:%t", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	fmt.Println()
	allCars.Process(sortedAppender)
	fmt.Println(sortedCars)
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	appender := func(car *Car) {
		if _, ok := sortedCars[car.Manufacturer]; ok {
			sortedCars[car.Manufacturer] = append(sortedCars[car.Manufacturer], car)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], car)
		}
	}

	return appender, sortedCars

}
