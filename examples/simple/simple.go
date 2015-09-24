package main

import (
	"errors"
	"fmt"
	"reflect"

	sql "database/sql/driver"

	"gopkg.in/go-playground/validator.v8"
)

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
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

	validateStruct()
	validateField()
}

func validateStruct() {

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( map[string]*FieldError )
	errs := validate.Struct(user)

	if errs != nil {

		fmt.Println(errs) // output: Key: "User.Age" Error:Field validation for "Age" failed on the "lte" tag
		//	                         Key: "User.Addresses[0].City" Error:Field validation for "City" failed on the "required" tag
		err := errs.(validator.ValidationErrors)["User.Addresses[0].City"]
		fmt.Println(err.Field) // output: City
		fmt.Println(err.Tag)   // output: required
		fmt.Println(err.Kind)  // output: string
		fmt.Println(err.Type)  // output: string
		fmt.Println(err.Param) // output:
		fmt.Println(err.Value) // output:

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateField() {
	myEmail := "joeybloggs.gmail.com"

	errs := validate.Field(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}

var validate2 *validator.Validate

type valuer struct {
	Name string
}

func (v valuer) Value() (sql.Value, error) {

	if v.Name == "errorme" {
		return nil, errors.New("some kind of error")
	}

	if v.Name == "blankme" {
		return "", nil
	}

	if len(v.Name) == 0 {
		return nil, nil
	}

	return v.Name, nil
}

// ValidateValuerType implements validator.CustomTypeFunc
func ValidateValuerType(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(sql.Valuer); ok {
		val, err := valuer.Value()
		if err != nil {
			// handle the error how you want
			return nil
		}

		return val
	}

	return nil
}

func main2() {

	config := &validator.Config{TagName: "validate"}

	validate2 = validator.New(config)
	validate2.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})

	validateCustomFieldType()
}

func validateCustomFieldType() {
	val := valuer{
		Name: "blankme",
	}

	errs := validate2.Field(val, "required")
	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "required" tag
		return
	}

	// all ok
}
