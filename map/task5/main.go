package main

import "fmt"

func main() {
	grades := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 76,
		"Dave":    89,
	}
	fmt.Println(reverseMap(grades))
}

// Эта функция принимает исходный map, который имеет ключи типа string и значения типа int,
// и возвращает новый map, в котором ключи и значения поменяны местами. Новый map имеет тип map[int]string.
//
// Мы создаем новый пустой map reversedMap с помощью функции make,
// а затем проходим по всем элементам исходного map с помощью цикла for range.
// Внутри цикла мы используем ключ и значение текущего элемента, чтобы добавить новый элемент в reversedMap.
// Мы помещаем значение в качестве нового ключа и старый ключ в качестве нового значения.
//
// В конце функция возвращает reversedMap, который теперь содержит поменянные ключи и значения.
func reverseMap(originalMap map[string]int) map[int]string {
	reversedMap := make(map[int]string)
	for key, value := range originalMap {
		reversedMap[value] = key
	}
	return reversedMap
}
