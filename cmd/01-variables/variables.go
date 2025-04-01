package main

import "fmt"

func ExampleVariables() {

	name := "Фрося"
	age := 25
	salary := 1500.50
	busy := true
	x := "gfgf"

	fmt.Println("Имя:", name,
		"\nВозраст:", age, "лет",
		"\nЗарплата:", salary, "руб.",
		"\nЗанят:", busy)

	fmt.Println(x)
}
