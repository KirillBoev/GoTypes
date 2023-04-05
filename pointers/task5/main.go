package main

import "fmt"

func main() {
	//В этом примере мы создаем новый экземпляр структуры Person и передаем его адрес функции ChangeName.
	//Функция изменяет поле Name на "Bob", и после вызова значения этого поля изменились на новое значение.
	person := &Person{Name: "Alice", Age: 30}
	fmt.Println(person.Name) // Alice

	ChangeName(person, "Bob")

	fmt.Println(person.Name)
}

type Person struct {
	Name string
	Age  int
}

// Функция ChangeName принимает указатель на структуру Person и новое имя,
// которое нужно присвоить полю Name. Затем она изменяет поле Name на указанное значение, используя переданный указатель.
func ChangeName(p *Person, newName string) {
	p.Name = newName
}
