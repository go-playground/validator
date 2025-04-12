package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Enum uint8

const (
	Zero Enum = iota
	One
	Two
)

func (e Enum) NotZero() bool {
	return e != Zero
}

func (e *Enum) Validate() error {
	if e == nil {
		return errors.New("can't be nil")
	}

	return nil
}

type Struct struct {
	Foo *Enum `validate:"validateFn"`         //uses Validate() error by default
	Bar Enum  `validate:"validateFn=NotZero"` // uses NotZero() bool
}

func main() {
	validate := validator.New()

	var x Struct

	if err := validate.Struct(x); err != nil {
		fmt.Printf("Expected Err(s):\n%+v\n", err)
	}

	x = Struct{
		Foo: new(Enum),
		Bar: One,
	}

	if err := validate.Struct(x); err != nil {
		fmt.Printf("Unexpected Err(s):\n%+v\n", err)
	}
}
