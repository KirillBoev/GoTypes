package main

//В этом примере определены две структуры, Rectangle и Circle,
//которые реализуют интерфейс Shape с помощью методов Area и Perimeter.
//Функция printShapeInfo принимает на вход интерфейс Shape и выводит его площадь
//и периметр с помощью методов Area и Perimeter. В функции main создаются экземпляры
//структур Rectangle и Circle, и их информация выводится на экран с помощью функции printShapeInfo.
import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
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

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %v\nPerimeter: %v\n", s.Area(), s.Perimeter())
}

func main() {
	r := Rectangle{width: 5, height: 10}
	c := Circle{radius: 7}

	printShapeInfo(r)
	printShapeInfo(c)
}
