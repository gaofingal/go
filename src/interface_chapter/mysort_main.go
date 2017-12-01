package main

import (
	"./mysort"
	"fmt"
)

type dayArray struct {
	data []*day
}
type day struct {
	num       int
	shortName string
	longName  string
}

func (da *dayArray) Len() int {
	return len(da.data)
}
func (da *dayArray) Less(i, j int) bool {
	return da.data[i].num < da.data[j].num
}
func (da *dayArray) Swap(i, j int) {
	da.data[i], da.data[j] = da.data[j], da.data[i]
}
func main() {
	ms := mysort.IntArray{2, 43, 56, 33, 98, 11, 12, 3}
	mysort.Sort(ms)
	fmt.Println(ms)

	ss := mysort.StringArray{"gao", "hu", "kai", "fu", "lian", "xie"}
	mysort.Sort(ss)
	fmt.Println(ss)

	Sunday := day{1, "SUN", "Sunday"}
	Monday := day{2, "MON", "Monday"}
	Tuesday := day{3, "TUE", "Tuesday"}
	Wednesday := day{4, "WED", "Wednesday"}
	Thursday := day{5, "THU", "Thursday"}
	Friday := day{6, "FRI", "Friday"}
	Saturday := day{7, "SAT", "Saturday"}

	var da dayArray
	da.data = []*day{&Sunday, &Monday, &Tuesday, &Wednesday, &Thursday, &Friday, &Saturday}
	mysort.Sort(&da)
	for _, d := range da.data {
		fmt.Printf("%s,", d.longName)
	}

}
