package main

import "fmt"

func main() {
	slice := []int{1, 0, 2, 0, 3, 0, 4}
	slice = removeZeros(slice)
	fmt.Println(slice)
}

// В этой функции мы создаем новый пустой слайс result.
// Затем мы перебираем каждый элемент из исходного слайса
// slice и добавляем его в result, если он не равен 0.
// В конце мы возвращаем result. Обратите внимание, что это
// т код не изменяет исходный слайс, а возвращает новый слайс
// без элементов, равных 0. Если вам нужно изменить исходный слайс,
// вы можете модифицировать функцию, чтобы она изменяла slice напрямую.
func removeZeros(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v != 0 {
			result = append(result, v)
		}
	}
	return result
}
