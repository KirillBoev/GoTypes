package main

import "fmt"

func main() {
	counts := make(map[string]int)
	counts2 := make(map[string]int)
	counts = map[string]int{"apple": 1, "max": 3, "nix": 4, "Pony": 32, "banana": 67, "pop": 89}
	counts2 = map[string]int{"Marge": 13, "Max": 43, "Denis": 64, "Pony": 72, "Manik": 17, "Lion": 89}
	fmt.Println(mergeMaps(counts, counts2))
}

// Эта функция принимает два аргумента типа map[string]int, которые мы хотим объединить.
// Затем мы создаем новый map mergedMap для хранения объединенных значений.
// Затем мы проходимся по первой map m1 и добавляем все ее значения в mergedMap.
// Затем мы проходимся по второй map m2, и для каждого ключа k и значения v проверяем,
// есть ли уже такой ключ в mergedMap. Если ключ уже есть, то мы добавляем к его текущему значению значение из m2.
// Если ключа еще нет в mergedMap, то мы просто добавляем его со значением из m2. В конце мы возвращаем mergedMap.
func mergeMaps(m1 map[string]int, m2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for k, v := range m1 {
		mergedMap[k] = v
	}

	for k, v := range m2 {
		if val, ok := mergedMap[k]; ok {
			mergedMap[k] = val + v
		} else {
			mergedMap[k] = v
		}
	}

	return mergedMap
}
