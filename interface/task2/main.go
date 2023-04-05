package main

import "fmt"

func main() {
	s := "string"
	fmt.Println(isString(s))
}

// Эта функция использует интерфейс пустого интерфейса interface{}
// для принятия любого типа данных в качестве аргумента.
// Она затем использует типовое утверждение data.(string) для проверки,
// является ли переданный аргумент строкой. Если data является строкой,
// функция вернет true, иначе - false.
func isString(data interface{}) bool {
	_, ok := data.(string)
	return ok
}
