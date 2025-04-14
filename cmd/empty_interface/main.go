package main

import "fmt"

func main() {

	var str string
	var num int
	printType(str)
	printTypeSlice(str, num)
}

// func printType(i interface{}) {
// 	fmt.Printf("%T", i)
// }

// то же самое печатет тип переменной
func printType(a any) {
	fmt.Printf("var type: %T", a)
}

// распечатает тип []interface{}
func printTypeSlice(a ...any) {
	fmt.Printf("\nvar type: %T", a)
}
