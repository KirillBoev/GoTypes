package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 13, 65}
	fmt.Println(findMinMax(&a))
}

func findMinMax(arr *[]int) (int, int) {
	// Получаем первый элемент массива
	min, max := (*arr)[0], (*arr)[0]

	// Проходим по всем элементам массива и обновляем значения min и max при необходимости
	for _, num := range *arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Возвращаем найденные значения
	return min, max
}
