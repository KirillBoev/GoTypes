package main

import "fmt"

func main() {
	//Здесь мы создаем две переменные start и end, и передаем их адреса функции
	//countPrimesInRange с помощью операторов &. Функция возвращает количество найденных
	//простых чисел, которое мы выводим на экран с помощью fmt.Println.
	start := 1
	end := 100
	count := countPrimesInRange(&start, &end)
	fmt.Println(count)
}

// В этой функции мы проходимся по всем числам в диапазоне от start до end (включительно)
// и проверяем каждое число на простоту. Если число простое, увеличиваем счетчик на 1.
// В конце функция возвращает количество найденных простых чисел.
func countPrimesInRange(start *int, end *int) int {
	count := 0

	for i := *start; i <= *end; i++ {
		prime := true

		// Проверяем, является ли текущее число простым
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}

		if prime {
			count++
		}
	}

	return count
}
