package main

import (
	"fmt"
	"math/rand"
)

func main() {

	s := make([]int, 10)

	populate_slice(&s)

	fmt.Printf("s: %v, длина = %d, емкость = %d", s, len(s), cap(s))
}

func populate_slice(s *[]int) {
	for i := range *s {
		(*s)[i] = rand.Intn(100)
	}
}
