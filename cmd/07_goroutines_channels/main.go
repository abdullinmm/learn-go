package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 3; i++ {
		ch := make(chan int)
		go func() {
			for j := 1; j <= 5; j++ {
				ch <- j
			}
			close(ch) // Закрываем канал
			time.Sleep(100 * time.Second)
		}()

		for num := range ch { // Читаем до закрытия
			fmt.Println(num)
		}
	}
}
