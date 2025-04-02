package utils

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Привет, %s! Добро пожаловать в Go.", name)
}
