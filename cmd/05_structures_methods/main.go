package main

import "fmt"

func main() {

	emp := Employee{
		Person:   Person{Name: "Алиса"},
		Salary:   1000,
		Position: "manager",
	}

	// возвращает сумму бонуса от зарплаты.
	fmt.Println(emp.CalculateBonus(emp.Position))
}
