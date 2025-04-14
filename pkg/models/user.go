package models

import (
	"errors"
	"regexp"
	"strings"
)

// Предопределенные ошибки валидации
var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrUnderage     = errors.New("user must be at least 18 years old")
	ErrInvalidName  = errors.New("name must be at least 2 characters long")

	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Age   int    `json:"age" xml:"age"`
	Email string `json:"email" xml:"email"`
}

// Метод валидации пользователя (использует указатель для избегания копирования)
func (u *User) Validate() error {
	// Валидация email
	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}
	// Валидация возраста
	if u.Age < 18 || u.Age < 0 {
		return ErrUnderage
	}
	// Дополнительная валидация имени
	if len(u.Name) < 2 || strings.TrimSpace(u.Name) == "" {
		return ErrInvalidName
	}

	return nil
}

// Вспомогательная функция для проверки email
func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
