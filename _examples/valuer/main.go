package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Nullable wraps a generic value.
type Nullable[T any] struct {
	Data T
}

// ValidatorValue returns the inner value that should be validated.
func (n Nullable[T]) ValidatorValue() any {
	return n.Data
}

type Config struct {
	Name string `validate:"required"`
}

type Record struct {
	Config Nullable[Config] `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	validate = validator.New()

	// valid case: ValidatorValue unwraps Config and Name passes required.
	valid := Record{
		Config: Nullable[Config]{
			Data: Config{Name: "validator"},
		},
	}
	err := validate.Struct(valid)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	// invalid case: Name is empty after unwrapping, so required fails on Config.Name.
	invalid := Record{
		Config: Nullable[Config]{},
	}
	err = validate.Struct(invalid)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) && len(validationErrs) > 0 {
			fmt.Printf("First error namespace: %s\n", validationErrs[0].Namespace())
		}
	}
}
