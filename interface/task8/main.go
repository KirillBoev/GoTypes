package main

import "fmt"

func main() {
	s := StringList{"foo", "bar", "baz"}
	Sort(s)
	fmt.Println(s) // [bar baz foo]
}

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(s Sortable) {
	for i := 0; i < s.Len()-1; i++ {
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(j, i) {
				s.Swap(i, j)
			}
		}
	}
}

// Для использования этой функции необходимо реализовать интерфейс Sortable
// для конкретного типа данных. Например, для сортировки списка строк можно реализовать такой интерфейс:
type StringList []string

func (s StringList) Len() int {
	return len(s)
}

func (s StringList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
