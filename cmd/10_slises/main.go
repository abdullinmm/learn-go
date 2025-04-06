package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	// Создаем slice
	s := make([]int, 10)
	// Инициализация локального генератора
	r := rand.New((rand.NewPCG(0, uint64(time.Now().UnixNano()))))

	// Заполняем случайными числами
	if err := populateSlice(&s, r); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Исходный slice: %v\n", s)

	// Удваиваем элементы
	doubleSlice(&s)
	fmt.Printf("После удвоения: %v\n", s)
}

func populateSlice(s *[]int, r *rand.Rand) error {
	if s == nil {
		return fmt.Errorf("slice is nil")
	}
	for i := range *s {
		(*s)[i] = r.IntN(100)
	}
	return nil
}

func doubleSlice(s *[]int) {
	for i := range *s {
		(*s)[i] *= 2
	}
}
