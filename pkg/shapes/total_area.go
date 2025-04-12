package shape

func TotalArea(shapes ...Shape) (total float64) {
	for _, s := range shapes {
		area, _ := s.Area()
		total = total + area
	}
	return
}
