package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filterEvenNumbers(numbers)
	fmt.Println(evens)
}

// Здесь мы создаем новый слайс evens, который будет содержать только
// четные числа из исходного слайса. Затем мы проходим по исходному слайсу
// numbers с помощью цикла for range, проверяем каждое число на четность с помощью
// оператора %, и если число четное, добавляем его в evens с помощью функции append.
// В конце функции мы возвращаем новый слайс evens.
func filterEvenNumbers(numbers []int) []int {
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}
