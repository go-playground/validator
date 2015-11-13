Package validator
================

[![Join the chat at https://gitter.im/bluesuncorp/validator](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-playground/validator?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://semaphoreci.com/api/v1/projects/ec20115f-ef1b-4c7d-9393-cc76aba74eb4/523019/badge.svg)](https://semaphoreci.com/joeybloggs/validator)
[![Coverage Status](https://coveralls.io/repos/go-playground/validator/badge.svg?branch=v8-development&service=github)](https://coveralls.io/github/go-playground/validator?branch=v8-development)
[![Go Report Card](http://goreportcard.com/badge/go-playground/validator)](http://goreportcard.com/report/go-playground/validator)
[![GoDoc](https://godoc.org/gopkg.in/go-playground/validator.v8?status.svg)](https://godoc.org/gopkg.in/go-playground/validator.v8)

Package validator implements value validations for structs and individual fields based on tags.

It has the following **unique** features:

-   Cross Field and Cross Struct validations by using validation tags or custom validators.  
-   Slice, Array and Map diving, which allows any or all levels of a multidimensional field to be validated.  
-   Handles type interface by determining it's underlying type prior to validation.
-   Handles custom field types such as sql driver Valuer see [Valuer](https://golang.org/src/database/sql/driver/types.go?s=1210:1293#L29)
-   Alias validation tags, which allows for mapping of several validations to a single tag for easier defining of validations on structs
-   Extraction of custom defined Field Name e.g. can specify to extract the JSON name while validating and have it available in the resulting FieldError

Installation
------------

Use go get.

	go get gopkg.in/go-playground/validator.v8

or to update

	go get -u gopkg.in/go-playground/validator.v8

Then import the validator package into your own code.

	import "gopkg.in/go-playground/validator.v8"

Error Return Value
-------

Validation functions return type error

They return type error to avoid the issue discussed in the following, where err is always != nil:

* http://stackoverflow.com/a/29138676/3158232
* https://github.com/go-playground/validator/issues/134

validator only returns nil or ValidationErrors as type error; so in you code all you need to do
is check if the error returned is not nil, and if it's not type cast it to type ValidationErrors
like so:

```go
err := validate.Struct(mystruct)
validationErrors := err.(validator.ValidationErrors)
 ```

Usage and documentation
------

Please see http://godoc.org/gopkg.in/go-playground/validator.v8 for detailed usage docs.

##### Examples:

Struct & Field validation
```go
package main

import (
	"fmt"

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
```

Custom Field Type
```go
package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// DbBackedUser User struct
type DbBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt64  `validate:"required"`
}

func main() {

	config := &validator.Config{TagName: "validate"}

	validate := validator.New(config)

	// register all sql.Null* types to use the ValidateValuer CustomTypeFunc
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}
	errs := validate.Struct(x)

	if len(errs.(validator.ValidationErrors)) > 0 {
		fmt.Printf("Errs:\n%+v\n", errs)
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

Benchmarks
------
###### Run on MacBook Pro (Retina, 15-inch, Late 2013) 2.6 GHz Intel Core i7 16 GB 1600 MHz DDR3 using Go 1.5.1
```go
$ go test -cpu=4 -bench=. -benchmem=true
PASS
BenchmarkFieldSuccess-4                            	 5000000	       291 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldFailure-4                            	 5000000	       294 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldDiveSuccess-4                        	  500000	      3498 ns/op	     528 B/op	      28 allocs/op
BenchmarkFieldDiveFailure-4                        	  300000	      4094 ns/op	     928 B/op	      32 allocs/op
BenchmarkFieldCustomTypeSuccess-4                  	 3000000	       460 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldCustomTypeFailure-4                  	 2000000	       758 ns/op	     400 B/op	       4 allocs/op
BenchmarkFieldOrTagSuccess-4                       	 1000000	      1393 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldOrTagFailure-4                       	 1000000	      1181 ns/op	     432 B/op	       6 allocs/op
BenchmarkStructSimpleCustomTypeSuccess-4           	 1000000	      1218 ns/op	      80 B/op	       5 allocs/op
BenchmarkStructSimpleCustomTypeFailure-4           	 1000000	      1748 ns/op	     624 B/op	      11 allocs/op
BenchmarkStructPartialSuccess-4                    	 1000000	      1392 ns/op	     400 B/op	      11 allocs/op
BenchmarkStructPartialFailure-4                    	 1000000	      1938 ns/op	     816 B/op	      16 allocs/op
BenchmarkStructExceptSuccess-4                     	 2000000	       903 ns/op	     368 B/op	       9 allocs/op
BenchmarkStructExceptFailure-4                     	 1000000	      1381 ns/op	     400 B/op	      11 allocs/op
BenchmarkStructSimpleCrossFieldSuccess-4           	 1000000	      1215 ns/op	     128 B/op	       6 allocs/op
BenchmarkStructSimpleCrossFieldFailure-4           	 1000000	      1781 ns/op	     560 B/op	      11 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldSuccess-4	 1000000	      1801 ns/op	     160 B/op	       8 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldFailure-4	 1000000	      2357 ns/op	     592 B/op	      13 allocs/op
BenchmarkStructSimpleSuccess-4                     	 1000000	      1161 ns/op	      48 B/op	       3 allocs/op
BenchmarkStructSimpleFailure-4                     	 1000000	      1818 ns/op	     624 B/op	      11 allocs/op
BenchmarkStructSimpleSuccessParallel-4             	 5000000	       375 ns/op	      48 B/op	       3 allocs/op
BenchmarkStructSimpleFailureParallel-4             	 2000000	       757 ns/op	     624 B/op	      11 allocs/op
BenchmarkStructComplexSuccess-4                    	  200000	      8053 ns/op	     432 B/op	      27 allocs/op
BenchmarkStructComplexFailure-4                    	  100000	     12634 ns/op	    3335 B/op	      69 allocs/op
BenchmarkStructComplexSuccessParallel-4            	 1000000	      2718 ns/op	     432 B/op	      27 allocs/op
BenchmarkStructComplexFailureParallel-4            	  300000	      5086 ns/op	    3336 B/op	      69 allocs/op
```

How to Contribute
------

There will always be a development branch for each version i.e. `v1-development`. In order to contribute, 
please make your pull requests against those branches.

If the changes being proposed or requested are breaking changes, please create an issue, for discussion
or create a pull request against the highest development branch for example this package has a
v1 and v1-development branch however, there will also be a v2-development branch even though v2 doesn't exist yet.

I strongly encourage everyone whom creates a custom validation function to contribute them and
help make this package even better.

License
------
Distributed under MIT License, please see license file in code for more details.
