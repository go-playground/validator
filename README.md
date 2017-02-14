Package validator
================
<img align="right" src="https://raw.githubusercontent.com/go-playground/validator/v9/logo.png">
[![Join the chat at https://gitter.im/go-playground/validator](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-playground/validator?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
![Project status](https://img.shields.io/badge/version-9.3.4-green.svg)
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
-   Customizable i18n aware error messages.
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

- [Simple](https://github.com/go-playground/validator/blob/v9/examples/simple/main.go)
- [Custom Field Types](https://github.com/go-playground/validator/blob/v9/examples/custom/main.go)
- [Struct Level](https://github.com/go-playground/validator/blob/v9/examples/struct-level/main.go)
- [Translations & Custom Errors](https://github.com/go-playground/validator/blob/v9/examples/translations/main.go)
- [Gin upgrade and/or override validator](https://github.com/go-playground/validator/tree/v9/examples/gin-upgrading-overriding)
- [wash - an example application putting it all together](https://github.com/bluesuncorp/wash)

Benchmarks
------
###### Run on i5-7600 16 GB 1600 MHz DDR4 using Go version go1.7.5 darwin/amd64
```go
BenchmarkFieldSuccess-4                                       	20000000	        84.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkFieldSuccessParallel-4                               	50000000	        31.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkFieldFailure-4                                       	 5000000	       299 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldFailureParallel-4                               	20000000	       104 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldDiveSuccess-4                                   	 2000000	       637 ns/op	     201 B/op	      11 allocs/op
BenchmarkFieldDiveSuccessParallel-4                           	10000000	       191 ns/op	     201 B/op	      11 allocs/op
BenchmarkFieldDiveFailure-4                                   	 2000000	       895 ns/op	     412 B/op	      16 allocs/op
BenchmarkFieldDiveFailureParallel-4                           	 5000000	       280 ns/op	     412 B/op	      16 allocs/op
BenchmarkFieldCustomTypeSuccess-4                             	10000000	       222 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldCustomTypeSuccessParallel-4                     	20000000	        70.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldCustomTypeFailure-4                             	 5000000	       313 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldCustomTypeFailureParallel-4                     	20000000	       103 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldOrTagSuccess-4                                  	 2000000	       743 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldOrTagSuccessParallel-4                          	 3000000	       485 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldOrTagFailure-4                                  	 3000000	       530 ns/op	     224 B/op	       5 allocs/op
BenchmarkFieldOrTagFailureParallel-4                          	 3000000	       402 ns/op	     224 B/op	       5 allocs/op
BenchmarkStructLevelValidationSuccess-4                       	10000000	       216 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructLevelValidationSuccessParallel-4               	20000000	        68.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructLevelValidationFailure-4                       	 3000000	       517 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructLevelValidationFailureParallel-4               	10000000	       169 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCustomTypeSuccess-4                      	 5000000	       385 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructSimpleCustomTypeSuccessParallel-4              	20000000	       108 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructSimpleCustomTypeFailure-4                      	 2000000	       700 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructSimpleCustomTypeFailureParallel-4              	 5000000	       241 ns/op	     440 B/op	      10 allocs/op
BenchmarkStructFilteredSuccess-4                              	 2000000	       606 ns/op	     288 B/op	       9 allocs/op
BenchmarkStructFilteredSuccessParallel-4                      	10000000	       198 ns/op	     288 B/op	       9 allocs/op
BenchmarkStructFilteredFailure-4                              	 3000000	       473 ns/op	     256 B/op	       7 allocs/op
BenchmarkStructFilteredFailureParallel-4                      	10000000	       158 ns/op	     256 B/op	       7 allocs/op
BenchmarkStructPartialSuccess-4                               	 2000000	       561 ns/op	     256 B/op	       6 allocs/op
BenchmarkStructPartialSuccessParallel-4                       	10000000	       176 ns/op	     256 B/op	       6 allocs/op
BenchmarkStructPartialFailure-4                               	 2000000	       803 ns/op	     480 B/op	      11 allocs/op
BenchmarkStructPartialFailureParallel-4                       	 5000000	       255 ns/op	     480 B/op	      11 allocs/op
BenchmarkStructExceptSuccess-4                                	 2000000	       868 ns/op	     496 B/op	      12 allocs/op
BenchmarkStructExceptSuccessParallel-4                        	10000000	       156 ns/op	     240 B/op	       5 allocs/op
BenchmarkStructExceptFailure-4                                	 2000000	       731 ns/op	     464 B/op	      10 allocs/op
BenchmarkStructExceptFailureParallel-4                        	10000000	       236 ns/op	     464 B/op	      10 allocs/op
BenchmarkStructSimpleCrossFieldSuccess-4                      	 3000000	       412 ns/op	      72 B/op	       3 allocs/op
BenchmarkStructSimpleCrossFieldSuccessParallel-4              	10000000	       121 ns/op	      72 B/op	       3 allocs/op
BenchmarkStructSimpleCrossFieldFailure-4                      	 2000000	       661 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCrossFieldFailureParallel-4              	10000000	       202 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldSuccess-4           	 3000000	       583 ns/op	      80 B/op	       4 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldSuccessParallel-4   	10000000	       167 ns/op	      80 B/op	       4 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldFailure-4           	 2000000	       852 ns/op	     320 B/op	       9 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldFailureParallel-4   	 5000000	       257 ns/op	     320 B/op	       9 allocs/op
BenchmarkStructSimpleSuccess-4                                	 5000000	       240 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSimpleSuccessParallel-4                        	20000000	        70.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSimpleFailure-4                                	 2000000	       657 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructSimpleFailureParallel-4                        	10000000	       210 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructComplexSuccess-4                               	 1000000	      1395 ns/op	     128 B/op	       8 allocs/op
BenchmarkStructComplexSuccessParallel-4                       	 3000000	       387 ns/op	     128 B/op	       8 allocs/op
BenchmarkStructComplexFailure-4                               	  300000	      4650 ns/op	    3040 B/op	      53 allocs/op
BenchmarkStructComplexFailureParallel-4                       	 1000000	      1372 ns/op	    3040 B/op	      53 allocs/op
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
