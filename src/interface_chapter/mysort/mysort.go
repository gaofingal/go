package mysort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}
func (ia IntArray) Less(i, j int) bool {
	return ia[i] < ia[j]
}
func (ia IntArray)Swap(i,j int){
	ia[i],ia[j] = ia[j],ia[i]
}

type StringArray []string
func (sa StringArray) Len() int {
	return len(sa)
}
func (sa StringArray) Less(i, j int) bool {
	return sa[i] < sa[j]
}
func (sa StringArray)Swap(i,j int){
	sa[i],sa[j] = sa[j],sa[i]
}

