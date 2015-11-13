package main

import (
	"fmt"
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// User contains user information
type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func main() {

	config := &validator.Config{TagName: "validate"}

	validate = validator.New(config)
	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	validateStruct()
}

// UserStructLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For Example this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//
// NOTE: you may ask why wouldn't I just do this outside of validator, because doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.
func UserStructLevelValidation(v *validator.Validate, structLevel *validator.StructLevel) {

	user := structLevel.CurrentStruct.Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		structLevel.ReportError(reflect.ValueOf(user.FirstName), "FirstName", "fname", "fnameorlname")
		structLevel.ReportError(reflect.ValueOf(user.LastName), "LastName", "lname", "fnameorlname")
	}

	// plus can to more, even with different tag than "fnameorlname"
}

func validateStruct() {

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "Unknown",
	}

	user := &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( map[string]*FieldError )
	errs := validate.Struct(user)

	if errs != nil {

		fmt.Println(errs) // output: Key: 'User.LastName' Error:Field validation for 'LastName' failed on the 'fnameorlname' tag
		//	                         Key: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'fnameorlname' tag
		err := errs.(validator.ValidationErrors)["User.FirstName"]
		fmt.Println(err.Field) // output: FirstName
		fmt.Println(err.Tag)   // output: fnameorlname
		fmt.Println(err.Kind)  // output: string
		fmt.Println(err.Type)  // output: string
		fmt.Println(err.Param) // output:
		fmt.Println(err.Value) // output:

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}
