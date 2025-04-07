package main

import (
	"fmt"
	"learn-go/pkg/email"
)

func main() {
	user := email.User{}
	user.Name = "Marsel"
	user.Email = "abdullinmm@gmail.com"

	fmt.Println(email.SendEmail(&user))
}
