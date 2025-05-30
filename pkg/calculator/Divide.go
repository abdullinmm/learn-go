package calculator

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
