package main

import (
	"bufio"
	"fmt"
	"learn-go/pkg/utils"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя:")
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	name = strings.TrimSpace(name)

	fmt.Println(utils.Greet(name))
}
