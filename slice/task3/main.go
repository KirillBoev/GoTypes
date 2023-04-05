package main

import (
	"fmt"
	"sort"
)

func main() {
	mySlice := []string{"xanaxin", "banana", "orange", "apple", "pear"}
	sortedSlice := sortStringsAsc(mySlice)
	fmt.Println(sortedSlice)
}

// Эта функция использует встроенную функцию sort.Strings(),
// которая сортирует переданный слайс строк в порядке возрастания.
// Затем функция просто возвращает отсортированный слайс.
func sortStringsAsc(slice []string) []string {
	sort.Strings(slice)
	return slice
}
