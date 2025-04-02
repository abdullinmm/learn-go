package main

import "fmt"

func main() {

	emp := Employee{
		Person:   Person{Name: "Алиса"},
		Salary:   1000,
		Position: "manager",
	}

	fmt.Println(emp.CalculateBonus(emp.Position))
}
