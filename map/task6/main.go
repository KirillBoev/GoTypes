package main

import (
	"fmt"
	"strings"
)

func main() {
	grades := map[string]string{
		"Alice":   "Dave",
		"Bob":     "Alice",
		"Charlie": "Bob",
		"Dave":    "Charlie",
	}
	fmt.Println(toLowerCase(grades))
}

// В этой функции мы используем функцию strings.
// ToLower() для преобразования ключей и значений исходного map в нижний регистр.
// Затем мы добавляем переведенные ключи и значения в новый map. Функция make() используется для создания нового map.
func toLowerCase(m map[string]string) map[string]string {
	result := make(map[string]string)
	for key, value := range m {
		lowerKey := strings.ToLower(key)
		lowerValue := strings.ToLower(value)
		result[lowerKey] = lowerValue
	}
	return result
}
