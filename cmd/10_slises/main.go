package main

import (
	"fmt"
	"math/rand"
)

func main() {

	s := make([]int, 10)

	populateSlice(&s)
	s1 := doubleSlice(s)

	fmt.Printf("s: %v, длина = %d, емкость = %d", s, len(s), cap(s))
	fmt.Printf("\ns: %v, длина = %d, емкость = %d", s1, len(s1), cap(s1))
}

func populateSlice(s *[]int) {

	if s == nil {
		fmt.Println("Срез является nil")
		return
	}

	for i, _ := range *s {

		(*s)[i] = rand.Intn(100)
		// *s = append(*s, 100)
	}
}

func doubleSlice(s []int) []int {

	doubled := make([]int, len(s))
	for i, num := range s {
		doubled[i] = num * 2
	}

	if len(doubled) == 0 {
		fmt.Println("Функция вернула пустой срез")
		return nil
	}

	return doubled
}
