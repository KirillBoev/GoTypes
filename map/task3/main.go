package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[string]int{
		"apple":  1,
		"orange": 2,
		"banana": 3,
	}

	// Функция, которая возвращает true для ключей, начинающихся с буквы "a"
	condition := func(k string) bool {
		return strings.HasPrefix(k, "a")
	}

	filtered := filterMap(m, condition)
	fmt.Println(filtered) // вывод: map[apple:1]
}

// Конечный результат зависит от того, какое условие должны удовлетворять ключи и какие
// значения должны соответствовать этим ключам. Ниже приведен пример функции,
// которая принимает карту m и функцию condition, которая принимает ключ и возвращает логическое значение,
// указывающее, должен ли ключ быть включен в результирующую карту.

// Здесь создается новая пустая карта result, и каждый ключ из m проверяется с помощью функции condition.
// Если функция возвращает true для ключа, то он добавляется в result вместе со значением, связанным с этим ключом в m.
// Функция возвращает новую карту, содержащую только те элементы,
// для которых функция condition возвращает true при вызове с ключом.
func filterMap(m map[string]int, condition func(string) bool) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		if condition(k) {
			result[k] = v
		}
	}
	return result
}
