package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data := []byte("Hello, world!")
	buf := writeToBuffer(data)
	fmt.Println(buf.String())
}

// В данном примере мы создаем новый буфер, используя тип bytes.Buffer.
// Затем мы используем функцию io.Copy() для записи данных в буфер из источника данных типа bytes.Reader.
// Функция io.Copy() автоматически обрабатывает передачу данных из источника в буфер,
// пока не будет достигнут конец источника или пока не произойдет ошибка.
//
// Данная функция принимает входные данные в виде среза байтов data и возвращает
// указатель на буфер bytes.Buffer, который содержит записанные данные.
// Если произошла ошибка в процессе записи, необходимо обработать ее в соответствии с логикой вашего приложения.
func writeToBuffer(data []byte) *bytes.Buffer {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, bytes.NewReader(data))
	if err != nil {
		// Обработка ошибки
	}
	return buf
}
