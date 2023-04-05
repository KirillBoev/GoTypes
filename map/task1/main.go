package main

import "fmt"

func main() {
	s := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	counts := countElements(s)
	fmt.Println(counts)
}

// Эта функция принимает в качестве входных данных слайс строк s,
// а затем проходится по каждому элементу и увеличивает счетчик вхождений
// этого элемента в counts. Наконец, функция возвращает counts,
// который является map[string]int, где ключ - это элемент, а значение - количество его вхождений.
func countElements(s []string) map[string]int {
	counts := make(map[string]int)
	for _, elem := range s {
		counts[elem]++
	}
	return counts
}
