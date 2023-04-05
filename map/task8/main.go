package main

import "fmt"

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	deleteByKeyCondition(m, func(k string) bool {
		return len(k) == 3
	})
	fmt.Println(m)
}

func deleteByKeyCondition(m map[string]int, condition func(string) bool) {
	for k := range m {
		if condition(k) {
			delete(m, k)
		}
	}
}
