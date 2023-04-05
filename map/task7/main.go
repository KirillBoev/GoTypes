package main

import "fmt"

func main() {
	grades := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 76,
		"Dave":    89,
	}
	fmt.Println(sumMapValues(grades))
}

// Эта функция объявляет переменную sum, которая инициализируется нулем.
// Затем она использует цикл for для итерации по всем значениям
// в карте m с помощью оператора _ (игнорирования ключа) и переменной value,
// которая содержит значение для текущего ключа. Каждое значение добавляется к переменной sum.
// В конце функция возвращает sum.
func sumMapValues(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}
