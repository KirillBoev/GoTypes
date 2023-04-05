package main

import (
	"fmt"
	"reflect"
)

// В данном примере используется пакет reflect,
// который позволяет определить тип данных во время выполнения программы.
// Функция printType принимает аргумент типа interface{}, что позволяет ей работать с любыми типами данных,
// а не только с определенными заранее.
func printType(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func main() {
	printType(42)             // int
	printType("Hello")        // string
	printType(true)           // bool
	printType(3.14)           // float64
	printType([]int{1, 2, 3}) // []int
}
