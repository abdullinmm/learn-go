package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number int
	reader := bufio.NewReader(os.Stdin) // Для очистки буфера

	for {
		fmt.Print("Введите число: ")
		_, err := fmt.Scan(&number) // Пытаемся считать число

		if err != nil {
			// Читаем остаток строки из буфера, чтобы сбросить ошибку
			reader.ReadString('\n')
			fmt.Println("Ошибка: введите целое число.")
			continue
		}

		// Если ввод успешен, выходим из цикла
		break
	}

	x, msg := chNum(number)
	fmt.Println(msg)

	for i := 1; i <= x; i++ {
		fmt.Println(i)
	}
}
