package validator

import (
	. "github.com/go-playground/assert/v2"
	"testing"
)

func TestAddValidatorNilFunc(t *testing.T) {
	defer func() {
		r := recover()
		EqualSkip(t, 3, r, "validator Func can't be nil")
	}()
	AddValidator("nil_validator", nil)
}

func TestAddValidatorExistingName(t *testing.T) {
	defer func() {
		r := recover()
		EqualSkip(t, 3, r, "validator required already exists")
	}()
	AddValidator("required", func(fl FieldLevel) bool {
		return true
	})
}

func TestAddValidatorSuccess(t *testing.T) {
	validator := func(fl FieldLevel) bool {
		return true
	}
	AddValidator("testing_custom_validator", validator)
	_, ok := bakedInValidators["testing_custom_validator"]
	EqualSkip(t, 2, ok, true)

}
