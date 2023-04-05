package main

import "fmt"

func main() {
	slice := []string{"Bob", "Danice", "Max", "Bob"}
	fmt.Println(removeDuplicates(slice))
}

// На каждой итерации цикла for, проверяем, была ли строка str уже обработана.
// Если строка не встречалась раньше, то добавляем ее в result и помечаем, что она уже была обработана,
// установив соответствующее значение в encountered в true.
// Таким образом, result будет содержать только уникальные строки из исходного слайса.
func removeDuplicates(strSlice []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, str := range strSlice {
		if encountered[str] == false {
			encountered[str] = true
			result = append(result, str)
		}
	}

	return result
}
