package main

// Определение структуры Person
type Person struct {
	Name string
}

// Определение структуры Employee, которая включает Person
type Employee struct {
	Person
	Salary   float64
	Position string
}

// Метод CalculateBonus для структуры Employee
func (e Employee) CalculateBonus() float64 {
	// Проверяем, совпадает ли позиция сотрудника с переданной позицией
	if e.Position == "manager" {
		// Возвращаем бонус, который составляет 10% от зарплаты
		return (e.Salary / 100) * 10
	}
	// Если позиции не совпадают, возвращаем 5% от зарплаты
	return (e.Salary / 100) * 5
}
