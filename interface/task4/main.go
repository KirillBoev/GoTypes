package main

import (
	"fmt"
	"strconv"
)

func toString(v interface{}) string {
	switch v := v.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func main() {
	fmt.Println(toString(42))             // "42"
	fmt.Println(toString(3.14))           // "3.14"
	fmt.Println(toString(true))           // "true"
	fmt.Println(toString("hello"))        // "hello"
	fmt.Println(toString([]int{1, 2, 3})) // "[1 2 3]"
}
