package main

import "fmt"

func main() {
	//Здесь мы вызываем функцию containsValue и передаем ей нашу map grades и функцию-условие,
	//которая проверяет, больше ли оценка 90. Результат выполнения программы будет true,
	//потому что у нас есть студент Bob, чья оценка выше 90. Если бы мы передали условие,
	//которое проверяет, больше ли оценка 95, то результат был бы false, потому что ни у кого из студентов нет оценки выше 95.
	grades := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 76,
		"Dave":    89,
	}

	hasHighGrade := containsValue(grades, func(grade int) bool {
		return grade > 90
	})

	fmt.Println(hasHighGrade)
}

// Здесь функция containsValue принимает map m типа map[string]int и функцию condition,
// которая принимает значение типа int и возвращает булево значение.
//
// Внутри функции containsValue мы перебираем все значения map'а,
// используя цикл for и ключевое слово range. Для каждого значения мы вызываем функцию condition и,
// если она возвращает true, возвращаем true. Если мы перебрали все значения map'а
// и не нашли ни одного значения, удовлетворяющего условию, то возвращаем false.
func containsValue(m map[string]int, condition func(int) bool) bool {
	for _, v := range m {
		if condition(v) {
			return true
		}
	}
	return false
}
