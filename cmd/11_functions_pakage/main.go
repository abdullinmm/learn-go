package main

import (
	"fmt"
	"learn-go/pkg/calculator"
	"learn-go/pkg/strings"
)

func main() {
	fmt.Println("Sum a + b:", calculator.Add(1, 2))

	d, err := calculator.Divide(4, 6)
	if err != nil {
		fmt.Printf("Return error: %v", err)
	}
	fmt.Printf("Divide a on b: %.2f", d)

	str, err := strings.Revers("ABC")
	if err != nil {
		fmt.Printf("str: %v", err)
	}
	fmt.Printf("\n%s", str)
}
