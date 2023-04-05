package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"a": 3, "b": 1, "c": 2}
	sorted := sortMapByValues(m)
	fmt.Println(sorted)
}

func sortMapByValues(m map[string]int) map[string]int {
	// Создаем слайс пар ключ-значение из исходной map.
	pairs := make([]Pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair{k, v})
	}

	// Определяем функцию сортировки, которая будет сравнивать значения в парах.
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	// Создаем новую map и заполняем ее отсортированными парами.
	sortedMap := make(map[string]int)
	for _, pair := range pairs {
		sortedMap[pair.Key] = pair.Value
	}

	return sortedMap
}

// Структура Pair представляет пару ключ-значение для сортировки.
type Pair struct {
	Key   string
	Value int
}
