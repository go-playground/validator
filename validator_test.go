package validator_test

import (
	"fmt"
	"testing"

	"github.com/joeybloggs/go-validate-yourself"
)

type UserDetails struct {
	Address string `validate:"omitempty,length=6"`
}

type User struct {
	FirstName string `validate:"required"`
	Details   *UserDetails
}

func TestValidateStruct(t *testing.T) {

	u := &User{
		FirstName: "",
		Details: &UserDetails{
			"",
		},
	}

	errors := validator.ValidateStruct(u)

	fmt.Println(errors == nil)

	for _, i := range errors {
		fmt.Printf("Error Struct:%s\n", i.Struct)

		for _, j := range i.Errors {

			fmt.Printf("Error Field:%s Error Tag:%s\n", j.Field, j.ErrorTag)
			fmt.Println(j.Error())
		}
	}

}

// func TestValidateField(t *testing.T) {
//
// 	u := &User{
// 		FirstName: "Dean Karn",
// 		Details: &UserDetails{
// 			"26 Here Blvd.",
// 		},
// 	}
//
// 	err := validator.ValidateFieldByTag(u.FirstName, "required")
//
// 	fmt.Println(err == nil)
// 	fmt.Println(err)
// }
