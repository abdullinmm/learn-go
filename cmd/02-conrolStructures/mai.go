package main

import "fmt"

func main() {

	var number int

	fmt.Println("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка: вы ввели не число")
	} else {
		fmt.Printf("Вы ввели число: %d\n", number)
	}

	x := chNum(number)

	for i := 1; i <= x; i++ {
		fmt.Println(i)
	}
}
