package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

//go:generate enumer -type=Enum
type Enum uint8

const (
	Zero Enum = iota
	One
	Two
	Three
)

func (e *Enum) Validate() error {
	if e == nil {
		return errors.New("can't be nil")
	}

	return nil
}

type Struct struct {
	Foo *Enum `validate:"validateFn"`         // uses Validate() error by default
	Bar Enum  `validate:"validateFn=IsAEnum"` // uses IsAEnum() bool provided by enumer
}

func main() {
	validate := validator.New()

	var x Struct

	x.Bar = Enum(64)

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
