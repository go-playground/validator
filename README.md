Package validator
================
<img align="right" src="https://raw.githubusercontent.com/go-playground/validator/v9/logo.png">[![Join the chat at https://gitter.im/go-playground/validator](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-playground/validator?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
![Project status](https://img.shields.io/badge/version-9.15.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/validator/branches/v9/badge.svg)](https://semaphoreci.com/joeybloggs/validator)
[![Coverage Status](https://coveralls.io/repos/go-playground/validator/badge.svg?branch=v9&service=github)](https://coveralls.io/github/go-playground/validator?branch=v9)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/validator)](https://goreportcard.com/report/github.com/go-playground/validator)
[![GoDoc](https://godoc.org/gopkg.in/go-playground/validator.v9?status.svg)](https://godoc.org/gopkg.in/go-playground/validator.v9)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package validator implements value validations for structs and individual fields based on tags.

It has the following **unique** features:

-   Cross Field and Cross Struct validations by using validation tags or custom validators.  
-   Slice, Array and Map diving, which allows any or all levels of a multidimensional field to be validated.
-   Ability to dive into both map keys and values for validation  
-   Handles type interface by determining it's underlying type prior to validation.
-   Handles custom field types such as sql driver Valuer see [Valuer](https://golang.org/src/database/sql/driver/types.go?s=1210:1293#L29)
-   Alias validation tags, which allows for mapping of several validations to a single tag for easier defining of validations on structs
-   Extraction of custom defined Field Name e.g. can specify to extract the JSON name while validating and have it available in the resulting FieldError
-   Customizable i18n aware error messages.
-   Default validator for the [gin](https://github.com/gin-gonic/gin) web framework; upgrading from v8 to v9 in gin see [here](https://github.com/go-playground/validator/tree/v9/_examples/gin-upgrading-overriding)

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

- [Simple](https://github.com/go-playground/validator/blob/v9/_examples/simple/main.go)
- [Custom Field Types](https://github.com/go-playground/validator/blob/v9/_examples/custom/main.go)
- [Struct Level](https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go)
- [Translations & Custom Errors](https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go)
- [Gin upgrade and/or override validator](https://github.com/go-playground/validator/tree/v9/_examples/gin-upgrading-overriding)
- [wash - an example application putting it all together](https://github.com/bluesuncorp/wash)

Benchmarks
------
###### Run on MacBook Pro (15-inch, 2017) Go version go1.9.4 darwin/amd64
```go
goos: darwin
goarch: amd64
pkg: github.com/go-playground/validator
BenchmarkFieldSuccess-8                                       	20000000	        86.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkFieldSuccessParallel-8                               	50000000	        27.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkFieldFailure-8                                       	 5000000	       297 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldFailureParallel-8                               	20000000	       107 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldArrayDiveSuccess-8                              	 2000000	       618 ns/op	     201 B/op	      11 allocs/op
BenchmarkFieldArrayDiveSuccessParallel-8                      	10000000	       225 ns/op	     201 B/op	      11 allocs/op
BenchmarkFieldArrayDiveFailure-8                              	 2000000	       863 ns/op	     412 B/op	      16 allocs/op
BenchmarkFieldArrayDiveFailureParallel-8                      	 5000000	       322 ns/op	     413 B/op	      16 allocs/op
BenchmarkFieldMapDiveSuccess-8                                	 1000000	      1336 ns/op	     432 B/op	      18 allocs/op
BenchmarkFieldMapDiveSuccessParallel-8                        	 3000000	       474 ns/op	     432 B/op	      18 allocs/op
BenchmarkFieldMapDiveFailure-8                                	 1000000	      1103 ns/op	     512 B/op	      16 allocs/op
BenchmarkFieldMapDiveFailureParallel-8                        	 5000000	       412 ns/op	     512 B/op	      16 allocs/op
BenchmarkFieldMapDiveWithKeysSuccess-8                        	 1000000	      1572 ns/op	     480 B/op	      21 allocs/op
BenchmarkFieldMapDiveWithKeysSuccessParallel-8                	 3000000	       615 ns/op	     480 B/op	      21 allocs/op
BenchmarkFieldMapDiveWithKeysFailure-8                        	 1000000	      1438 ns/op	     721 B/op	      21 allocs/op
BenchmarkFieldMapDiveWithKeysFailureParallel-8                	 3000000	       543 ns/op	     721 B/op	      21 allocs/op
BenchmarkFieldCustomTypeSuccess-8                             	10000000	       230 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldCustomTypeSuccessParallel-8                     	20000000	        82.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkFieldCustomTypeFailure-8                             	 5000000	       284 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldCustomTypeFailureParallel-8                     	20000000	       118 ns/op	     208 B/op	       4 allocs/op
BenchmarkFieldOrTagSuccess-8                                  	 2000000	       824 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldOrTagSuccessParallel-8                          	 3000000	       472 ns/op	      16 B/op	       1 allocs/op
BenchmarkFieldOrTagFailure-8                                  	 3000000	       487 ns/op	     224 B/op	       5 allocs/op
BenchmarkFieldOrTagFailureParallel-8                          	 5000000	       405 ns/op	     224 B/op	       5 allocs/op
BenchmarkStructLevelValidationSuccess-8                       	10000000	       214 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructLevelValidationSuccessParallel-8               	20000000	        78.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructLevelValidationFailure-8                       	 3000000	       475 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructLevelValidationFailureParallel-8               	10000000	       200 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCustomTypeSuccess-8                      	 3000000	       403 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructSimpleCustomTypeSuccessParallel-8              	10000000	       143 ns/op	      32 B/op	       2 allocs/op
BenchmarkStructSimpleCustomTypeFailure-8                      	 2000000	       655 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructSimpleCustomTypeFailureParallel-8              	 5000000	       286 ns/op	     440 B/op	      10 allocs/op
BenchmarkStructFilteredSuccess-8                              	 2000000	       598 ns/op	     288 B/op	       9 allocs/op
BenchmarkStructFilteredSuccessParallel-8                      	10000000	       231 ns/op	     288 B/op	       9 allocs/op
BenchmarkStructFilteredFailure-8                              	 3000000	       455 ns/op	     256 B/op	       7 allocs/op
BenchmarkStructFilteredFailureParallel-8                      	10000000	       197 ns/op	     256 B/op	       7 allocs/op
BenchmarkStructPartialSuccess-8                               	 3000000	       552 ns/op	     256 B/op	       6 allocs/op
BenchmarkStructPartialSuccessParallel-8                       	10000000	       206 ns/op	     256 B/op	       6 allocs/op
BenchmarkStructPartialFailure-8                               	 2000000	       750 ns/op	     480 B/op	      11 allocs/op
BenchmarkStructPartialFailureParallel-8                       	 5000000	       317 ns/op	     480 B/op	      11 allocs/op
BenchmarkStructExceptSuccess-8                                	 2000000	       853 ns/op	     496 B/op	      12 allocs/op
BenchmarkStructExceptSuccessParallel-8                        	10000000	       179 ns/op	     240 B/op	       5 allocs/op
BenchmarkStructExceptFailure-8                                	 2000000	       698 ns/op	     464 B/op	      10 allocs/op
BenchmarkStructExceptFailureParallel-8                        	 5000000	       276 ns/op	     464 B/op	      10 allocs/op
BenchmarkStructSimpleCrossFieldSuccess-8                      	 3000000	       412 ns/op	      72 B/op	       3 allocs/op
BenchmarkStructSimpleCrossFieldSuccessParallel-8              	10000000	       148 ns/op	      72 B/op	       3 allocs/op
BenchmarkStructSimpleCrossFieldFailure-8                      	 2000000	       630 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCrossFieldFailureParallel-8              	10000000	       244 ns/op	     304 B/op	       8 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldSuccess-8           	 2000000	       610 ns/op	      80 B/op	       4 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldSuccessParallel-8   	10000000	       205 ns/op	      80 B/op	       4 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldFailure-8           	 2000000	       861 ns/op	     320 B/op	       9 allocs/op
BenchmarkStructSimpleCrossStructCrossFieldFailureParallel-8   	 5000000	       315 ns/op	     320 B/op	       9 allocs/op
BenchmarkStructSimpleSuccess-8                                	 5000000	       279 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSimpleSuccessParallel-8                        	20000000	        86.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSimpleFailure-8                                	 2000000	       636 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructSimpleFailureParallel-8                        	10000000	       264 ns/op	     424 B/op	       9 allocs/op
BenchmarkStructComplexSuccess-8                               	 1000000	      1539 ns/op	     128 B/op	       8 allocs/op
BenchmarkStructComplexSuccessParallel-8                       	 3000000	       557 ns/op	     128 B/op	       8 allocs/op
BenchmarkStructComplexFailure-8                               	  300000	      4136 ns/op	    3041 B/op	      53 allocs/op
BenchmarkStructComplexFailureParallel-8                       	 1000000	      1855 ns/op	    3041 B/op	      53 allocs/op
```

Complementary Software
----------------------

Here is a list of software that complements using this library either pre or post validation.

* [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.
* [mold](https://github.com/go-playground/mold) - A general library to help modify or set data within data structures and other objects

How to Contribute
------

Make a pull request...

License
------
Distributed under MIT License, please see license file within the code for more details.
