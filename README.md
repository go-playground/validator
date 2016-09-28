Package validator
================
<img align="right" src="https://raw.githubusercontent.com/go-playground/validator/v9/logo.png">
[![Join the chat at https://gitter.im/go-playground/validator](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-playground/validator?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
![Project status](https://img.shields.io/badge/version-9.2.1-green.svg)
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
###### Run on MacBook Pro (Retina, 15-inch, Late 2013) 2.6 GHz Intel Core i7 16 GB 1600 MHz DDR3 using Go version go1.7.1 darwin/amd64
```go
BenchmarkFieldSuccess-8                                       	20000000	       106 ns/op
BenchmarkFieldSuccessParallel-8                               	50000000	        33.7 ns/op
BenchmarkFieldFailure-8                                       	 5000000	       346 ns/op
BenchmarkFieldFailureParallel-8                               	20000000	       115 ns/op
BenchmarkFieldDiveSuccess-8                                   	 2000000	       739 ns/op
BenchmarkFieldDiveSuccessParallel-8                           	10000000	       246 ns/op
BenchmarkFieldDiveFailure-8                                   	 1000000	      1043 ns/op
BenchmarkFieldDiveFailureParallel-8                           	 5000000	       381 ns/op
BenchmarkFieldCustomTypeSuccess-8                             	 5000000	       270 ns/op
BenchmarkFieldCustomTypeSuccessParallel-8                     	20000000	        92.5 ns/op
BenchmarkFieldCustomTypeFailure-8                             	 5000000	       331 ns/op
BenchmarkFieldCustomTypeFailureParallel-8                     	20000000	       132 ns/op
BenchmarkFieldOrTagSuccess-8                                  	 2000000	       874 ns/op
BenchmarkFieldOrTagSuccessParallel-8                          	 5000000	       368 ns/op
BenchmarkFieldOrTagFailure-8                                  	 3000000	       566 ns/op
BenchmarkFieldOrTagFailureParallel-8                          	 5000000	       427 ns/op
BenchmarkStructLevelValidationSuccess-8                       	 5000000	       335 ns/op
BenchmarkStructLevelValidationSuccessParallel-8               	20000000	       124 ns/op
BenchmarkStructLevelValidationFailure-8                       	 2000000	       630 ns/op
BenchmarkStructLevelValidationFailureParallel-8               	10000000	       298 ns/op
BenchmarkStructSimpleCustomTypeSuccess-8                      	 3000000	       535 ns/op
BenchmarkStructSimpleCustomTypeSuccessParallel-8              	10000000	       170 ns/op
BenchmarkStructSimpleCustomTypeFailure-8                      	 2000000	       821 ns/op
BenchmarkStructSimpleCustomTypeFailureParallel-8              	 5000000	       379 ns/op
BenchmarkStructFilteredSuccess-8                              	 2000000	       769 ns/op
BenchmarkStructFilteredSuccessParallel-8                      	 5000000	       328 ns/op
BenchmarkStructFilteredFailure-8                              	 2000000	       594 ns/op
BenchmarkStructFilteredFailureParallel-8                      	10000000	       244 ns/op
BenchmarkStructPartialSuccess-8                               	 2000000	       682 ns/op
BenchmarkStructPartialSuccessParallel-8                       	 5000000	       291 ns/op
BenchmarkStructPartialFailure-8                               	 1000000	      1034 ns/op
BenchmarkStructPartialFailureParallel-8                       	 5000000	       392 ns/op
BenchmarkStructExceptSuccess-8                                	 1000000	      1014 ns/op
BenchmarkStructExceptSuccessParallel-8                        	10000000	       257 ns/op
BenchmarkStructExceptFailure-8                                	 2000000	       875 ns/op
BenchmarkStructExceptFailureParallel-8                        	 5000000	       405 ns/op
BenchmarkStructSimpleCrossFieldSuccess-8                      	 3000000	       545 ns/op
BenchmarkStructSimpleCrossFieldSuccessParallel-8              	10000000	       177 ns/op
BenchmarkStructSimpleCrossFieldFailure-8                      	 2000000	       787 ns/op
BenchmarkStructSimpleCrossFieldFailureParallel-8              	 5000000	       341 ns/op
BenchmarkStructSimpleCrossStructCrossFieldSuccess-8           	 2000000	       795 ns/op
BenchmarkStructSimpleCrossStructCrossFieldSuccessParallel-8   	10000000	       267 ns/op
BenchmarkStructSimpleCrossStructCrossFieldFailure-8           	 1000000	      1119 ns/op
BenchmarkStructSimpleCrossStructCrossFieldFailureParallel-8   	 3000000	       437 ns/op
BenchmarkStructSimpleSuccess-8                                	 5000000	       377 ns/op
BenchmarkStructSimpleSuccessParallel-8                        	20000000	       110 ns/op
BenchmarkStructSimpleFailure-8                                	 2000000	       785 ns/op
BenchmarkStructSimpleFailureParallel-8                        	 5000000	       302 ns/op
BenchmarkStructComplexSuccess-8                               	 1000000	      2159 ns/op
BenchmarkStructComplexSuccessParallel-8                       	 2000000	       723 ns/op
BenchmarkStructComplexFailure-8                               	  300000	      5237 ns/op
BenchmarkStructComplexFailureParallel-8                       	 1000000	      2378 ns/op
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
