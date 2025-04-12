package shape

import (
	"learn-go/pkg/models"
	"math"
	"testing"
)

func TestShapes(t *testing.T) {
	test := []struct {
		shape  Shape
		name   string
		answer float64
	}{
		// Обычные случаи
		{
			shape:  models.Circle{Radius: 2.2},
			name:   "Area of circul",
			answer: 15.21},
		{
			shape:  models.Triangle{Height: 2.1, Base: 0.5},
			name:   "Area of Triangle",
			answer: 0.53},
		{
			shape:  models.Rectangle{Height: 4.0, Width: 2.5},
			name:   "Area of Rectangle",
			answer: 10.00},
		// todo: Граничные значения
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if result, _ := tt.shape.Area(); math.Round(result*100)/100 != math.Round(tt.answer*100)/100 {
				t.Errorf("Expected %.2f, got %.2f", tt.answer, result)
				t.Fail()
			}
		})
	}
}
