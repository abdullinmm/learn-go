package main

import (
	"fmt"
	"learn-go/pkg/models"
	"learn-go/pkg/validator"
)

func main() {
	user := models.User{}
	user.Email = "abdullinmm"
	user.Age = 17
	if v := validator.ValidatorUser(&user); v != nil {
		fmt.Printf("%s", v)
	}
}
