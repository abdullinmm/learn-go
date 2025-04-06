package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // Для ожидания завершения горутин

	// Создаем один общий канал для всех горутин
	ch := make(chan int)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Добавляем задачу в WaitGroup

		go func() {
			defer wg.Done() // Уменьшаем счетчик при завершении

			for j := 1; j <= 5; j++ {
				ch <- j
				time.Sleep(100 * time.Millisecond) // Имитация задержки
			}
		}()
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Читаем все данные из канала
	for num := range ch {
		fmt.Println(num)
	}
}
