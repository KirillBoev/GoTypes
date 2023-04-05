package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(reverseSlice(numbers))
}

// В этом коде мы используем цикл for, чтобы пройти по исходному слайсу
// и в обратном порядке поместить каждый элемент в новый слайс.
// Мы также используем функцию make() для создания нового слайса с той же длиной,
// что и исходный, и затем возвращаем его в качестве результата.
// Эта функция также проверяет, если исходный слайс пуст, то он возвращает исходный слайс без изменений.
func reverseSlice(input []int) []int {
	if len(input) == 0 {
		return input
	}
	output := make([]int, len(input))
	for i, j := 0, len(input)-1; i < len(input); i, j = i+1, j-1 {
		output[i] = input[j]
	}
	return output
}
