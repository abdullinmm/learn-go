package strings

import (
	"fmt"
)

// функция переворачивает строку
func Revers(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("String is empty")
	}

	r := []rune(s)
	rev := make([]rune, len(r))
	for i, v := range r {
		rev[len(r)-1-i] = v
	}
	return string(rev), nil
}
