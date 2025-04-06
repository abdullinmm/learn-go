package main

import (
	"fmt"
	"math/rand"
)

func main() {

	s := make([]int, 10)

	populateSlice(&s)
	s1, err := doubleSlice(s)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("s: %v, длина = %d, емкость = %d", s, len(s), cap(s))
	fmt.Printf("\ns: %v, длина = %d, емкость = %d", s1, len(s1), cap(s1))
}

func populateSlice(s *[]int) error {

	if s == nil {
		return fmt.Errorf("slice is nil")
	}

	for i, _ := range *s {

		(*s)[i] = rand.Intn(100)
		// *s = append(*s, 100)
	}

	if len(*s) == 0 {
		return fmt.Errorf("the function returned an empty slice")
	}

	return nil
}

func doubleSlice(s []int) ([]int, error) {

	doubled := make([]int, len(s))
	for i, num := range s {
		doubled[i] = num * 2
	}

	if len(doubled) == 0 {
		return nil, fmt.Errorf("the function returned an empty slice")
	}

	return doubled, nil
}
