package validator_test

import (
	"fmt"

	"../validator"
)

func ExampleValidate_new() {
	config := validator.Config{
		TagName:         "validate",
		ValidationFuncs: validator.BakedInValidators,
	}

	validator.New(config)
}

func ExampleValidate_field() {
	// This should be stored somewhere globally
	var validate *validator.Validate

	config := validator.Config{
		TagName:         "validate",
		ValidationFuncs: validator.BakedInValidators,
	}

	validate = validator.New(config)

	i := 0
	errs := validate.Field(i, "gt=1,lte=10")
	err := errs[""]
	fmt.Println(err.Field)
	fmt.Println(err.Tag)
	fmt.Println(err.Kind) // NOTE: Kind and Type can be different i.e. time Kind=struct and Type=time.Time
	fmt.Println(err.Type)
	fmt.Println(err.Param)
	fmt.Println(err.Value)
	//Output:
	//
	//gt
	//int
	//int
	//1
	//0
}

func ExampleValidate_struct() {
	// This should be stored somewhere globally
	var validate *validator.Validate

	config := validator.Config{
		TagName:         "validate",
		ValidationFuncs: validator.BakedInValidators,
	}

	validate = validator.New(config)

	type ContactInformation struct {
		Phone  string `validate:"required"`
		Street string `validate:"required"`
		City   string `validate:"required"`
	}

	type User struct {
		Name               string `validate:"required,excludesall=!@#$%^&*()_+-=:;?/0x2C"` // 0x2C = comma (,)
		Age                int8   `validate:"required,gt=0,lt=150"`
		Email              string `validate:"email"`
		ContactInformation []*ContactInformation
	}

	contactInfo := &ContactInformation{
		Street: "26 Here Blvd.",
		City:   "Paradeso",
	}

	user := &User{
		Name:               "Joey Bloggs",
		Age:                31,
		Email:              "joeybloggs@gmail.com",
		ContactInformation: []*ContactInformation{contactInfo},
	}

	errs := validate.Struct(user)
	for _, v := range errs {
		fmt.Println(v.Field) // Phone
		fmt.Println(v.Tag)   // required
		//... and so forth
		//Output:
		//Phone
		//required
	}
}
