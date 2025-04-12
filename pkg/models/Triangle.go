package models

type Triangle struct {
	Base, Height, Side float64
}

func (t Triangle) Area() (float64, error) {
	return (t.Base * t.Height) / 2, nil
}

func (t Triangle) Perimeter() (float64, error) {
	return (t.Base + t.Side) * 2, nil
}
