package utils

import "fmt"

/*
принимает имя (строка) и возвращает приветствие в формате:
"Привет, {Имя}! Добро пожаловать в Go."
*/
func Greet(name string) string {
	return fmt.Sprintf("Привет, %s! Добро пожаловать в Go.", name)
}
