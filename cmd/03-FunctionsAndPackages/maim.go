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
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println(utils.Greet(name))
}
