package main

import (
	"fmt"
	"learn-go/pkg/models"
	"learn-go/pkg/serializers"
)

func main() {
	user := models.User{}
	user.Name = "Marsel"
	user.Age = 20

	jsonSerializer := serializers.JSONSerializer{}
	SaveData(jsonSerializer, user)
}

func SaveData(serializer serializers.Serializer, data interface{}) {
	bytes, err := serializer.Serializer(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Serialized data:\n" + string(bytes))
}
