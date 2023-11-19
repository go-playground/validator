package validator

import (
	"errors"
	"fmt"
	"strings"
)

// default
func defaultValidator(fl FieldLevel) bool {
	if !hasValue(fl) {
		f := fl.Field()
		fmt.Printf("can %t", f.CanSet())
		fmt.Printf("can %t", f.CanAddr())
		if f.CanSet() && f.CanAddr() {
			err := mutate(&f, fl.Param())
			if err != nil {
				fmt.Printf("Mutation failed: %s", err)
			}
		}
	}

	return true
}

// default_with
func defaultWithValidator(fl FieldLevel) bool {
	// get required_with and default value
	params1 := strings.Split(fl.Param(), ":")

	if len(params1) != 2 {
		fmt.Printf("Bad parameter %s\n", fl.Param())
		return true
	}

	params := parseOneOfParam2(params1[0])
	for _, param := range params {
		if !requireCheckFieldKind(fl, param, true) {
			if !hasValue(fl) {
				f := fl.Field()
				if f.CanSet() && f.CanAddr() {
					err := mutate(&f, params1[1])
					if err != nil {
						fmt.Printf("Mutation failed: %s", err)
					}
				}
			}

			return true
		}
	}

	return true
}

// default_with_all
func defaultWithAllValidator(fl FieldLevel) bool {
	// get required_with and default value
	params1 := strings.Split(fl.Param(), ":")

	if len(params1) != 2 {
		fmt.Printf("Bad parameter %s\n", fl.Param())
		return true
	}

	params := parseOneOfParam2(params1[0])
	for _, param := range params {
		if requireCheckFieldKind(fl, param, true) {
			return true
		}
	}

	if !hasValue(fl) {
		f := fl.Field()
		if f.CanSet() && f.CanAddr() {
			err := mutate(&f, params1[1])
			if err != nil {
				fmt.Printf("Mutation failed: %s", err)
			}
		}
	}

	return true
}

// default_without
func defaultWithoutValidator(fl FieldLevel) bool {
	// get required_with and default value
	params1 := strings.Split(fl.Param(), ":")

	if len(params1) != 2 {
		fmt.Printf("Bad parameter %s\n", fl.Param())
		return true
	}

	if requireCheckFieldKind(fl, strings.TrimSpace(params1[0]), true) {
		if !hasValue(fl) {
			f := fl.Field()
			if f.CanSet() && f.CanAddr() {
				err := mutate(&f, params1[1])
				if err != nil {
					fmt.Printf("Mutation failed: %s", err)
				}
			}
		}

		return true
	}

	return true
}

// default_without_all
func defaultWithoutAllValidator(fl FieldLevel) bool {
	// get required_with and default value
	params1 := strings.Split(fl.Param(), ":")

	if len(params1) != 2 {
		fmt.Printf("Bad parameter %s\n", fl.Param())
		return true
	}

	params := parseOneOfParam2(params1[0])
	for _, param := range params {
		if !requireCheckFieldKind(fl, param, true) {
			return true
		}
	}

	if !hasValue(fl) {
		f := fl.Field()
		if f.CanSet() && f.CanAddr() {
			err := mutate(&f, params1[1])
			if err != nil {
				fmt.Printf("Mutation failed: %s", err)
			}
		}
	}

	return true
}

func RegisterCustomDefaultValidator(validate *Validate) error {
	if validate != nil {
		return validate.RegisterValidation("default", defaultValidator, true)
	}

	return errors.New("validate object is nil")
}

func RegisterCustomDefaultWithValidator(validate *Validate) error {
	if validate != nil {
		return validate.RegisterValidation("default_with", defaultWithValidator, true)
	}

	return errors.New("validate object is nil")
}

func RegisterCustomDefaultWithAllValidator(validate *Validate) error {
	if validate != nil {
		return validate.RegisterValidation("default_with_all", defaultWithAllValidator, true)
	}

	return errors.New("validate object is nil")
}

func RegisterCustomDefaultWithoutValidator(validate *Validate) error {
	if validate != nil {
		return validate.RegisterValidation("default_without", defaultWithoutValidator, true)
	}

	return errors.New("validate object is nil")
}

func RegisterCustomDefaultWithoutAllValidator(validate *Validate) error {
	if validate != nil {
		return validate.RegisterValidation("default_without_all", defaultWithoutAllValidator, true)
	}

	return errors.New("validate object is nil")
}
