package main

import (
	"fmt"
	"os"
)

// Этот интерфейс определяет один метод Write(),
// который принимает срез байтов p и возвращает количество записанных байтов n и ошибку err.
// Метод Write() должен записывать данные из среза p в буфер и возвращать количество записанных байтов n и ошибку err.
// Если в процессе записи возникла ошибка, она должна быть возвращена в качестве значения err. Если запись прошла успешно,
// err должна быть равна nil.
type Writer interface {
	Write(p []byte) (n int, err error)
}

// В этом примере мы создаем файл "output.txt" и получаем объект *os.File,
// который реализует интерфейс Writer. Затем мы записываем строку "Hello, world!\n" в файл,
// используя метод Write() объекта *os.File. В конце мы выводим количество записанных байтов.
func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := file
	data := []byte("Hello, world!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wrote %d bytes to file\n", n)
}
