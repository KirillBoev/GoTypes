package main

import "math"

func main() {

}

// Теперь Circle реализует интерфейс Shape,
// поскольку он имеет методы Area() и Perimeter() с нужной сигнатурой.
// Точно так же можно реализовать этот интерфейс для других фигур, таких как квадрат или прямоугольник.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
