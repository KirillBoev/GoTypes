package main

import "fmt"

func main() {
	slice1 := []int{3, 4, 6, 2, 6, 3, 7, 6, 9}
	slice2 := []int{5, 3, 56, 76, 23, 4, 56}
	fmt.Println(mergeAndRemoveDuplicates(slice1, slice2))
}

// Здесь мы используем функцию append для объединения двух слайсов, а затем проходим
// по объединенному слайсу и добавляем уникальные элементы в новый слайс с помощью
// вспомогательной функции contains. Обратите внимание, что мы создаем новый слайс
// с помощью make, чтобы избежать лишних реаллокаций памяти в процессе добавления элементов.
func mergeAndRemoveDuplicates(slice1 []int, slice2 []int) []int {
	merged := append(slice1, slice2...)   // объединяем слайсы
	unique := make([]int, 0, len(merged)) // создаем новый слайс для уникальных элементов

	// добавляем элементы в новый слайс, если они еще не встречались
	for _, val := range merged {
		if !contains(unique, val) {
			unique = append(unique, val)
		}
	}
	return unique
}

// вспомогательная функция для проверки, содержится ли элемент в слайсе
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
