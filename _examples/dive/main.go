package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// Test ...
type Test struct {
	Array []string          `validate:"required,gt=0,dive,required"`
	Map   map[string]string `validate:"required,gt=0,dive,keys,keymax,endkeys,required,max=1000"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	// registering alias so we can see the differences between
	// map key, value validation errors
	validate.RegisterAlias("keymax", "max=10")

	var test Test

	val(test)

	test.Array = []string{""}
	test.Map = map[string]string{"test > than 10": ""}
	val(test)
}

func val(test Test) {
	fmt.Println("testing")
	err := validate.Struct(test)
	fmt.Println(err)
}
