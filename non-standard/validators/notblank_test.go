package validators

import (
	"testing"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/assert.v1"
)

type test struct {
	String    string      `validate:"notblank"`
	Array     []int       `validate:"notblank"`
	Pointer   *int        `validate:"notblank"`
	Number    int         `validate:"notblank"`
	Interface interface{} `validate:"notblank"`
	Func      func()      `validate:"notblank"`
}

func TestNotBlank(t *testing.T) {
	v := validator.New()
	err := v.RegisterValidation("notblank", NotBlank)
	assert.Equal(t, nil, err)

	// Errors
	var x *int
	invalid := test{
		String:    " ",
		Array:     []int{},
		Pointer:   x,
		Number:    0,
		Interface: nil,
		Func:      nil,
	}
	fieldsWithError := []string{
		"String",
		"Array",
		"Pointer",
		"Number",
		"Interface",
		"Func",
	}

	errors := v.Struct(invalid).(validator.ValidationErrors)
	var fields []string
	for _, err := range errors {
		fields = append(fields, err.Field())
	}

	assert.Equal(t, fieldsWithError, fields)

	// No errors
	y := 1
	x = &y
	valid := test{
		String:    "str",
		Array:     []int{1},
		Pointer:   x,
		Number:    1,
		Interface: "value",
		Func:      func() {},
	}

	err = v.Struct(valid)
	assert.Equal(t, nil, err)
}
