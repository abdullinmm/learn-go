package main

import (
	"fmt"
	"learn-go/pkg/models"
	shape "learn-go/pkg/shapes"
)

func main() {
	circle := models.Circle{Radius: 2.2}
	rectangle := models.Rectangle{Height: 4.0, Width: 2.5}
	triangle := models.Triangle{Height: 2.1, Base: 0.5}

	shape.PrintShapeInfo(circle)
	shape.PrintShapeInfo(rectangle)
	shape.PrintShapeInfo(triangle)

	shape.TotalArea(circle, rectangle, triangle)

	fmt.Printf("Total area: %.2f", shape.TotalArea(circle, rectangle, triangle))
}
