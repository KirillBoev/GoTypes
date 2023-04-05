package main

import "fmt"

func main() {
	slice := []int{1, 5, 3, 8, 2}
	max := max(slice)
	fmt.Println(max)
}

// Эта функция принимает в качестве аргумента
// слайс целых чисел и возвращает максимальное значение.
// Если слайс пустой, функция генерирует панику.
func max(slice []int) int {
	if len(slice) == 0 {
		panic("Cannot find max of an empty slice")
	}
	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}
