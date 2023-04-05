package main

import "fmt"

func main() {
	inputMap := map[string]int{"a": 1, "b": 2, "c": 3}
	multiplier := 2
	outputMap := multiplyMapValues(inputMap, multiplier)
	fmt.Println(outputMap) // Output: map[a:2 b:4 c:6]
}

// Эта функция принимает в качестве аргументов исходный map inputMap и множитель multiplier.
// Она создает новый map outputMap и перебирает элементы inputMap, умножая значения на multiplier
// и сохраняя их в outputMap. После завершения цикла функция возвращает outputMap.
func multiplyMapValues(inputMap map[string]int, multiplier int) map[string]int {
	outputMap := make(map[string]int)
	for key, value := range inputMap {
		outputMap[key] = value * multiplier
	}
	return outputMap
}
