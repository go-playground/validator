package validator_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/joeybloggs/go-validate-yourself"
)

// type UserDetails struct {
// 	Address string `validate:"omitempty,length=6"`
// 	Sub     struct {
// 		A string `validate:"required"`
// 	}
// }

type UserDetails struct {
	Address string `validate:"omitempty,length=6"`
	Sub     struct {
		A string `validate:"required"`
	}
}

type User struct {
	FirstName string `validate:"required"`
	Details   *UserDetails
}

// func Test(t *testing.T) { TestingT(t) }
//
// type MySuite struct{}
//
// var _ = Suite(&MySuite{})
//
// func (s *MySuite) SetUpTest(c *C) {
// 	s.dir = c.MkDir()
// 	// Use s.dir to prepare some data.
// }

// func RecursiveErrorReporter(e *validator.StructValidationErrors) {
//
// 	// log.Printf("Error within Struct:%s\n", e.Struct)
//
// 	// for _, f := range e.Errors {
// 	// 	log.Println(f.Error())
// 	// }
// }

func TestValidateStruct(t *testing.T) {

	u := &User{
		FirstName: "",
		Details: &UserDetails{
			Address: "",
			Sub: struct {
				A string `validate:"required"`
			}{
				A: "",
			},
		},
	}

	errors := validator.ValidateStruct(u)

	fmt.Println(errors == nil)
	log.Println(errors.Error())

	// for _, i := range errors {
	// 	fmt.Printf("Error Struct:%s\n", i.Struct)
	//
	// 	for _, j := range i.Errors {
	//
	// 		fmt.Printf("Error Field:%s Error Tag:%s\n", j.Field, j.ErrorTag)
	// 		fmt.Println(j.Error())
	// 	}
	// }

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
