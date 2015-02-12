package validator_test

import (
	"fmt"
	"testing"

	"gopkg.in/joeybloggs/go-validate-yourself.v0"
)

type UserDetails struct {
	Address string `validate:"required"`
}

type User struct {
	FirstName string `validate:"required"`
	Details   *UserDetails
}

func TestValidateStruct(t *testing.T) {

	u := &User{
		FirstName: "Dean Karn",
		Details: &UserDetails{
			"26 Here Blvd.",
		},
	}

	errors := validator.ValidateStruct(u)

	fmt.Println(errors)
}
