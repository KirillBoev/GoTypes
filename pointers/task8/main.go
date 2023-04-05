package main

import "fmt"

func main() {
	//Вы можете вызвать эту функцию, передав указатель на экземпляр структуры
	//Point в качестве первого аргумента,
	//а затем значения, на которые вы хотите установить поля X и Y,
	//в качестве второго и третьего аргументов соответственно. Вот как это может выглядеть:
	p := &Point{1, 2}
	fmt.Println(p) // Выведет: &{1 2}

	setPoint(p, 3, 4)
	fmt.Println(p) // Выведет: &{3 4}
	//В этом примере мы создаем экземпляр структуры Point с полями X и Y,
	//установленными на 1 и 2 соответственно. Затем мы выводим этот экземпляр на экран,
	//чтобы убедиться, что его поля правильно установлены.
	//
	//Затем мы вызываем функцию setPoint и передаем ей указатель на этот экземпляр,
	//а также значения 3 и 4 для полей X и Y. Функция изменяет поля X и Y структуры на заданные значения.
	//
	//Наконец, мы снова выводим экземпляр структуры на экран, чтобы убедиться,
	//что его поля были изменены на новые значения 3 и 4.
}

type Point struct {
	X int
	Y int
}

func setPoint(p *Point, x int, y int) {
	p.X = x
	p.Y = y
}