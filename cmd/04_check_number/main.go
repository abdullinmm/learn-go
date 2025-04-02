package main

import (
	"bufio"
	"fmt"
	"learn-go/pkg/utils"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите число: ")
		number, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Ошибка чтения ввода: ", err)
			continue
		}

		number = strings.TrimSpace(number)

		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Ошибка ввода. Ожидается число:", err)
			continue
		}

		fmt.Printf("Число %d четное: %t\n", num, utils.IsEven(num))
		fmt.Printf("Число %d положитеьное: %t", num, utils.IsPositive(num))
		// Если ввод успешен, выходим из цикла
		break
	}
}
