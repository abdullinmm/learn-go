package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

// Sound implements Animal.
func (d Dog) Sound() string {
	return "Гав!"
}

type Cat struct{}

// Sound implements Animal.
func (c Cat) Sound() string {
	return "Мяу!"
}

type Cow struct{}

func (cow Cow) Sound() string {
	return "Муу!"
}

func main() {
	dog := Dog{} // задаем значение
	cat := Cat{} // задаем значение
	cow := Cow{} // задаем значение

	MakeSound(dog) // Выведет: Гав!
	MakeSound(cat) // Выведет: Мяу!
	MakeSound(cow) // Выведет: Муу!
}

func MakeSound(a Animal) {
	fmt.Println(a.Sound()) // выводим результат
}
