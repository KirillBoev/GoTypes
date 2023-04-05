package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(take(numbers, 2))  // [1 2]
	fmt.Println(take(numbers, 5))  // [1 2 3 4 5]
	fmt.Println(take(numbers, 10)) // [1 2 3 4 5]
}

// Здесь мы передаем исходный слайс и количество элементов,
// которые мы хотим извлечь. Если n больше, чем длина исходного слайса,
// мы присваиваем n значение длины исходного слайса, чтобы избежать ошибок.
// Затем мы возвращаем новый слайс, который содержит первые n элементов исходного слайса.
func take(slice []int, n int) []int {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n]
}
