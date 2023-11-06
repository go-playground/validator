package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email string `validate:"required,email"`
	Phone string `validate:"required,regex=^__comma__[0-9]{2}$"`
}

func main() {
	user := User{
		Email: "test@example.com",
		Phone: ",99",
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("Validation passed")
}
