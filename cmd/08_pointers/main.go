package main

import "fmt"

func main() {
	x := 42
	var p *int = &x                      // p хранит адрес x
	fmt.Println("Значение x:", x)        // 42
	fmt.Println("Адрес x:", &x)          // Например, 0xc00001a0b0
	fmt.Println("Указатель p:", p)       // Тот же адрес, что &x
	fmt.Println("Значение через p:", *p) // 42
}
