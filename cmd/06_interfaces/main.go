package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct {
	Voice string
}

// Sound implements Animal.
func (d Dog) Sound() string {
	return d.Voice
}

type Cat struct {
	Voice string
}

// Sound implements Animal.
func (c Cat) Sound() string {
	return c.Voice
}

func main() {
	dog := Dog{Voice: "Гав"} // задаем значение
	cat := Cat{Voice: "Мяу"} // задаем значение

	MakeSound(dog)
	MakeSound(cat)
}

func MakeSound(a Animal) {
	fmt.Println(a.Sound()) // выводим результат
}
