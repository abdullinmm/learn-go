package main

func chNum(num int) (int, string) {
	if num < 10 {
		return num, "Число меньше или равно 10"
	} else {
		return 10, ""
	}
}
