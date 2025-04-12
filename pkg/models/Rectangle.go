package models

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() (float64, error) {
	return r.Width * r.Height, nil
}

func (r Rectangle) Perimeter() (float64, error) {
	return 2 * (r.Width + r.Height), nil
}
