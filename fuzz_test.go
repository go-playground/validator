//go:build go1.18
// +build go1.18

// Copyright (c) 2015 Dean Karn
// SPDX-License-Identifier: MIT

package validator_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

// FuzzValidatorStruct validates arbitrary struct field values
// against validation tags. The validator processes user input
// through struct tags — a bug here can bypass validation.
//
// Validator has 20K+ GitHub stars and is the standard Go
// struct validation library.
func FuzzValidatorStruct(f *testing.F) {
	validate := validator.New()

	f.Add("test@example.com", "email")
	f.Add("123", "gte=0,lte=100")
	f.Add("", "required")
	f.Add("not-a-url", "url")

	f.Fuzz(func(t *testing.T, value, tag string) {
		if len(value) > 10000 || len(tag) > 10000 {
			return
		}

		type testStruct struct {
			Field string `validate:"%s"`
		}

		// Validate must never panic
		s := testStruct{Field: value}
		_ = validate.Struct(s)
	})
}
