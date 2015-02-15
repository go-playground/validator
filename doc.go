/*
Package validator implements value validations for structs and individual fields based on tags.

Built In Validator

The package contains a built in Validator instance for use,
but you may also create a new instance if needed.

	// built in
	errs := validator.ValidateStruct(//your struct)
	errs := validator.ValidateFieldByTag(field, "omitempty,min=1,max=10")

	// new
	newValidator = validator.New("struct tag name", validator.BakedInFunctions)

A simple example usage:

	type UserDetail {
		Details string `validate:"-"`
	}

	type User struct {
		Name         string     `validate:"required,max=60"`
		PreferedName string     `validate:"omitempty,max=60"`
		Sub          UserDetail
	}

	user := &User {
		Name: "",
	}

	// errs will contain a hierarchical list of errors
	// using the StructValidationErrors struct
	// or nil if no errors exist
	errs := validator.ValidateStruct(user)

	// in this case 1 error Name is required
	errs.Struct will be "User"
	errs.StructErrors will be empty <-- fields that were structs
	errs.Errors will have 1 error of type FieldValidationError

Error Handling

The error can be used like so

	fieldErr, _ := errs["Name"]
	fieldErr.Field    // "Name"
	fieldErr.ErrorTag // "required"

Both StructValidationErrors and FieldValidationError implement the Error interface but it's
intended use is for development + debugging, not a production error message.

	fieldErr.Error() // Field validation for "Name" failed on the "required" tag
	errs.Error()
	// Struct: User
	// Field validation for "Name" failed on the "required" tag

Why not a better error message? because this library intends for you to handle your own error messages

Why should I handle my own errors? Many reasons, for me building and internationalized application
I needed to know the field and what validation failed so that I could provide an error in the users specific language.

	if fieldErr.Field == "Name" {
		switch fieldErr.ErrorTag
		case "required":
			return "Translated string based on field + error"
		default:
		return "Translated string based on field"
	}

The hierarchical structure is hard to work with sometimes.. Agreed Flatten function to the rescue!
Flatten will return a map of FieldValidationError's but the field name will be namespaced.

	// if UserDetail Details field failed validation
	Field will be "Sub.Details"

	// for Name
	Field will be "Name"

Custom Functions

Custom functions can be added

	//Structure
	func customFunc(field interface{}, param string) bool {

		if whatever {
			return false
		}

		return true
	}

	validator.AddFunction("custom tag name", customFunc)
	// NOTE: using the same tag name as an existing function
	//       will overwrite the existing one

Custom Tag Name

A custom tag name can be set to avoid conficts, or just have a shorter name

	validator.SetTag("valid")

Multiple Validators

Multiple validators on a field will process in the order defined

	type Test struct {
		Field `validate:"max=10,min=1"`
	}

	// max will be checked then min

Bad Validator definitions are not handled by the library

	type Test struct {
		Field `validate:"min=10,max=0"`
	}

	// this definition of min max will never validate

Baked In Validators and Tags

Here is a list of the current built in validators:

	-
		Tells the validation to skip this struct field; this is particularily
		handy in ignoring embedded structs from being validated. (Usage: -)

	|
		This is the 'or' operator allowing multiple validators to be used and
		accepted. (Usage: rbg|rgba) <-- this would allow either rgb or rgba
		colors to be accepted. This can also be combined with 'and' for example
		( Usage: omitempty,rgb|rgba)

	omitempty
		Allows conitional validation, for example if a field is not set with
		a value (Determined by the required validator) then other validation
		such as min or max won't run, but if a value is set validation will run.
		(Usage: omitempty)

	required
		This validates that the value is not the data types default value.
		For numbers ensures value is not zero. For strings ensures value is
		not "". For slices, arrays, and maps, ensures the length is not zero.
		(Usage: required)

	len
		For numbers, max will ensure that the value is
		equal to the parameter given. For strings, it checks that
		the string length is exactly that number of characters. For slices,
		arrays, and maps, validates the number of items. (Usage: len=10)

	max
		For numbers, max will ensure that the value is
		less than or equal to the parameter given. For strings, it checks
		that the string length is at most that number of characters. For
		slices, arrays, and maps, validates the number of items. (Usage: max=10)

	min
		For numbers, min will ensure that the value is
		greater or equal to the parameter given. For strings, it checks that
		the string length is at least that number of characters. For slices,
		arrays, and maps, validates the number of items. (Usage: min=10)

	alpha
		This validates that a strings value contains alpha characters only
		(Usage: alpha)

	alphanum
		This validates that a strings value contains alphanumeric characters only
		(Usage: alphanum)

	numeric
		This validates that a strings value contains a basic numeric value.
		basic excludes exponents etc...
		(Usage: numeric)

	hexadecimal
		This validates that a strings value contains a valid hexadecimal.
		(Usage: hexadecimal)

	hexcolor
		This validates that a strings value contains a valid hex color including
		hashtag (#)
		(Usage: hexcolor)

	rgb
		This validates that a strings value contains a valid rgb color
		(Usage: rgb)

	rgba
		This validates that a strings value contains a valid rgba color
		(Usage: rgba)

	hsl
		This validates that a strings value contains a valid hsl color
		(Usage: hsl)

	hsla
		This validates that a strings value contains a valid hsla color
		(Usage: hsla)

	email
		This validates that a strings value contains a valid email
		This may not conform to all possibilities of any rfc standard, but neither
		does any email provider accept all posibilities...
		(Usage: email)

Validator notes:

	regex
		a regex validator won't be added because commas and = signs can be part of
		a regex which conflict with the validation definitions, although workarounds
		can be made, they take away from using pure regex's. Furthermore it's quick
		and dirty but the regex's become harder to maintain and are not reusable, so
		it's as much as a programming philosiphy as anything.

		In place of this new validator functions should be created; a regex can be
		used within the validator function and even be precompiled for better efficiency.

		And the best reason, you can sumit a pull request and we can keep on adding to the
		validation library of this package!

Panics

This package panics when bad input is provided, this is by design, bad code like that should not make it to production.

	type Test struct {
		TestField string `validate:"nonexistantfunction=1"`
	}

	t := &Test{
		TestField: "Test"
	}

	validator.ValidateStruct(t) // this will panic
*/
package validator
