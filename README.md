Package validator
================
<img align="right" src="https://raw.githubusercontent.com/go-playground/validator/v9/logo.png">
[![Join the chat at https://gitter.im/go-playground/validator](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-playground/validator?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
![Project status](https://img.shields.io/badge/version-9.1.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/validator/branches/v9/badge.svg)](https://semaphoreci.com/joeybloggs/validator)
[![Coverage Status](https://coveralls.io/repos/go-playground/validator/badge.svg?branch=v9&service=github)](https://coveralls.io/github/go-playground/validator?branch=v9)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/validator)](https://goreportcard.com/report/github.com/go-playground/validator)
[![GoDoc](https://godoc.org/gopkg.in/go-playground/validator.v9?status.svg)](https://godoc.org/gopkg.in/go-playground/validator.v9)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package validator implements value validations for structs and individual fields based on tags.

It has the following **unique** features:

-   Cross Field and Cross Struct validations by using validation tags or custom validators.  
-   Slice, Array and Map diving, which allows any or all levels of a multidimensional field to be validated.  
-   Handles type interface by determining it's underlying type prior to validation.
-   Handles custom field types such as sql driver Valuer see [Valuer](https://golang.org/src/database/sql/driver/types.go?s=1210:1293#L29)
-   Alias validation tags, which allows for mapping of several validations to a single tag for easier defining of validations on structs
-   Extraction of custom defined Field Name e.g. can specify to extract the JSON name while validating and have it available in the resulting FieldError
-   Default validator for the [gin](https://github.com/gin-gonic/gin) web framework; upgrading from v8 to v9 in gin see [here](https://github.com/go-playground/validator/tree/v9/examples/gin-upgrading-overriding)

Installation
------------

Use go get.

	go get gopkg.in/go-playground/validator.v9

Then import the validator package into your own code.

	import "gopkg.in/go-playground/validator.v9"

Error Return Value
-------

Validation functions return type error

They return type error to avoid the issue discussed in the following, where err is always != nil:

* http://stackoverflow.com/a/29138676/3158232
* https://github.com/go-playground/validator/issues/134

Validator only InvalidValidationError for bad validation input, nil or ValidationErrors as type error; so, in your code all you need to do is check if the error returned is not nil, and if it's not check if error is InvalidValidationError ( if necessary, most of the time it isn't ) type cast it to type ValidationErrors like so:

```go
err := validate.Struct(mystruct)
validationErrors := err.(validator.ValidationErrors)
 ```

Usage and documentation
------

Please see http://godoc.org/gopkg.in/go-playground/validator.v9 for detailed usage docs.

##### Examples:

Struct & Field validation
```go
package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	validateStruct()
	validateVariable()
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
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( map[string]*FieldError )
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
```

Custom Field Type
```go
package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"gopkg.in/go-playground/validator.v9"
)

// DbBackedUser User struct
type DbBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt64  `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	// register all sql.Null* types to use the ValidateValuer CustomTypeFunc
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	// build object for validation
	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}

	err := validate.Struct(x)

	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateValuer implements validator.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}
```

Struct Level Validation
```go
package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
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

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	// register validation for 'User'
	// NOTE: only have to register a non-pointer type for 'User', validator
	// interanlly dereferences during it's type checks.
	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	// build 'User' info, normally posted data etc...
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

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
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
func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}

	// plus can to more, even with different tag than "fnameorlname"
}
```

Benchmarks
------
###### Run on MacBook Pro (Retina, 15-inch, Late 2013) 2.6 GHz Intel Core i7 16 GB 1600 MHz DDR3 using Go version go1.7 darwin/amd64
```go
BenchmarkFieldSuccess-8                                       	20000000	       105 ns/op	       0 B/op	       0 allocs/op
BenchmarkFieldSuccessParallel-8                               	50000000	        35.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkFieldFailure-8                                       	 5000000	       337 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldFailureParallel-8                               	20000000	       120 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldDiveSuccess-8                                   	 2000000	       716 ns/op	     201 B/op	      11 allocs/op
BenchmarkFieldDiveSuccessParallel-8                           	10000000	       253 ns/op	     201 B/op	      11 allocs/op
BenchmarkFieldDiveFailure-8                                   	 1000000	      1060 ns/op	     412 B/op	      16 allocs/op
BenchmarkFieldDiveFailureParallel-8                           	 5000000	       360 ns/op	     413 B/op	      16 allocs/op
BenchmarkFieldCustomTypeSuccess-8                             	 5000000	       299 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldCustomTypeSuccessParallel-8                     	20000000	        86.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldCustomTypeFailure-8                             	 5000000	       341 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldCustomTypeFailureParallel-8                     	20000000	       140 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldOrTagSuccess-8                                  	 2000000	       893 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldOrTagSuccessParallel-8                          	 5000000	       431 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldOrTagFailure-8                                  	 2000000	       563 ns/op	     224 B/op	       5 allocs/op
BenchmarkFieldOrTagFailureParallel-8                          	 5000000	       417 ns/op	     224 B/op	       5 allocs/op
BenchmarkStructLevelValidationSuccess-8                       	 5000000	       339 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructLevelValidationSuccessParallel-8               	20000000	       114 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructLevelValidationFailure-8                       	 2000000	       630 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructLevelValidationFailureParallel-8               	 5000000	       291 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCustomTypeSuccess-8                      	 3000000	       540 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructSimpleCustomTypeSuccessParallel-8              	10000000	       176 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructSimpleCustomTypeFailure-8                      	 2000000	       821 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructSimpleCustomTypeFailureParallel-8              	 5000000	       336 ns/op	     440 B/op	      10 allocs/op
BenchmarkStructPartialSuccess-8                               	 2000000	       686 ns/op	     256 B/op	       6 allocs/op
BenchmarkStructPartialSuccessParallel-8                       	 5000000	       282 ns/op	     256 B/op	       6 allocs/op
BenchmarkStructPartialFailure-8                               	 2000000	       931 ns/op	     480 B/op	      11 allocs/op
BenchmarkStructPartialFailureParallel-8                       	 5000000	       394 ns/op	     480 B/op	      11 allocs/op
BenchmarkStructExceptSuccess-8                                	 1000000	      1017 ns/op	     496 B/op	      12 allocs/op
BenchmarkStructExceptSuccessParallel-8                        	10000000	       233 ns/op	     240 B/op	       5 allocs/op
BenchmarkStructExceptFailure-8                                	 2000000	       864 ns/op	     464 B/op	      10 allocs/op
BenchmarkStructExceptFailureParallel-8                        	 5000000	       393 ns/op	     464 B/op	      10 allocs/op
BenchmarkStructSimpleCrossFieldSuccess-8                      	 3000000	       552 ns/op	      72 B/op	       3 allocs/op
BenchmarkStructSimpleCrossFieldSuccessParallel-8              	10000000	       202 ns/op	      72 B/op	       3 allocs/op
BenchmarkStructSimpleCrossFieldFailure-8                      	 2000000	       798 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCrossFieldFailureParallel-8              	 5000000	       356 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldSuccess-8           	 2000000	       825 ns/op	      80 B/op	       4 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldSuccessParallel-8   	 5000000	       300 ns/op	      80 B/op	       4 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldFailure-8           	 2000000	      1103 ns/op	     320 B/op	       9 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldFailureParallel-8   	 3000000	       433 ns/op	     320 B/op	       9 allocs/op
BenchmarkStructSimpleSuccess-8                                	 5000000	       360 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSimpleSuccessParallel-8                        	20000000	       110 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSimpleFailure-8                                	 2000000	       783 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructSimpleFailureParallel-8                        	 5000000	       358 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructComplexSuccess-8                               	 1000000	      2120 ns/op	     128 B/op	       8 allocs/op
BenchmarkStructComplexSuccessParallel-8                       	 2000000	       659 ns/op	     128 B/op	       8 allocs/op
BenchmarkStructComplexFailure-8                               	  300000	      5126 ns/op	    3041 B/op	      53 allocs/op
BenchmarkStructComplexFailureParallel-8                       	 1000000	      2261 ns/op	    3041 B/op	      53 allocs/op
```

Complimentary Software
----------------------

Here is a list of software that compliments using this library either pre or post validation.

* [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.
* [Conform](https://github.com/leebenson/conform) - Trims, sanitizes & scrubs data based on struct tags.

How to Contribute
------

Make a pull request...

License
------
Distributed under MIT License, please see license file in code for more details.
