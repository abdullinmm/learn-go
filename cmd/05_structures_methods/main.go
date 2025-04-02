package main

import (
	"fmt"
	"learn-go/pkg/utils"
)

func main() {

	emp := Employee{
		Person:   Person{Name: "Алиса"},
		Salary:   -1000.00,
		Position: "manager",
	}

	if !utils.IsPositive(int(emp.Salary)) {
		fmt.Printf("У сотрудника %s отрицательная зарплата %.2f", emp.Name, emp.Salary)
		return
	}

	// возвращает сумму бонуса от зарплаты.
	fmt.Println(emp.CalculateBonus())
}
