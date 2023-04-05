package main

import "sort"

// Эти функции используют методы Len(), Swap() и Less() из интерфейса Sortable
// для сортировки срезов чисел и строк в определенном порядке.
// В обоих случаях мы определяем новый тип (ascendingSorter и descendingSorter),
// который реализует интерфейс Sortable, чтобы сортировка могла быть выполнена.
func main() {

}

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

// Функция, которая сортирует срез чисел в порядке возрастания.
func SortAscending(numbers []int) {
	sort.Sort(ascendingSorter(numbers))
}

// Сортировщик для сортировки в порядке возрастания.
type ascendingSorter []int

func (s ascendingSorter) Len() int {
	return len(s)
}

func (s ascendingSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ascendingSorter) Less(i, j int) bool {
	return s[i] < s[j]
}

// Функция, которая сортирует срез строк в порядке убывания.
func SortDescending(strings []string) {
	sort.Sort(descendingSorter(strings))
}

// Сортировщик для сортировки в порядке убывания.
type descendingSorter []string

func (s descendingSorter) Len() int {
	return len(s)
}

func (s descendingSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s descendingSorter) Less(i, j int) bool {
	return s[i] > s[j]
}
