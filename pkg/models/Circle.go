package models

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0.0, fmt.Errorf("invalid Radius value %.2f", c.Radius)
	}
	return math.Pi * math.Pow(c.Radius, 2), nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.Radius <= 0 {
		return 0.0, fmt.Errorf("invalid Radius value %.2f", c.Radius)
	}
	return 2 * math.Pi * c.Radius, nil
}
