package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	if containsMultipleOfThree(numbers) {
		fmt.Println("At least one element is a multiple of 3.")
	} else {
		fmt.Println("No elements are multiples of 3.")
	}
}

// Конечное условие, которое мы проверяем, может быть любым.
// Предположим, что нам нужно проверить, содержит ли слайс numbers
// хотя бы одно значение, которое делится на 3
// Эта функция перебирает каждый элемент слайса numbers и проверяет, делится ли он на 3.
// Если значение делится на 3, то функция возвращает true, иначе она продолжает перебирать элементы.
// Если ни один элемент не удовлетворяет условию, функция вернет false.
func containsMultipleOfThree(numbers []int) bool {
	for _, num := range numbers {
		if num%3 == 0 {
			return true
		}
	}
	return false
}
