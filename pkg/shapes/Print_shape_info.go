package shape

import "fmt"

func PrintShapeInfo(s Shape) {
	result, err := s.Area()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Area: %.2f\n", result)
}
