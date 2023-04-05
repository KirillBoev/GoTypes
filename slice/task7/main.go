package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = removeElements(s, 3) // удаляем элементы, равные 3
	fmt.Println(s)           // [1 2 4 5]

	s = []int{1, 2, 3, 4, 5}
	s = removeElements2(s, func(element int) bool { return element%2 == 0 }) // удаляем четные элементы
	fmt.Println(s)                                                           // [1 3 5]
}

// Пример 1: Удаление элементов, равных определенному значению
func removeElements(s []int, value int) []int {
	var result []int
	for _, element := range s {
		if element != value {
			result = append(result, element)
		}
	}
	return result
}

// Пример 2: Удаление элементов, удовлетворяющих определенному условию
func removeElements2(s []int, condition func(int) bool) []int {
	var result []int
	for _, element := range s {
		if !condition(element) {
			result = append(result, element)
		}
	}
	return result
}
