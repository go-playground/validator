package validator

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

const (
	fieldErrMsg = "Key: '%s' Error:Field validation for '%s' failed on the '%s' tag"
)

// InvalidValidationError describes an invalid argument passed to
// `Struct`, `StructExcept`, StructPartial` or `Field`
type InvalidValidationError struct {
	Type reflect.Type
}

// Error returns InvalidValidationError message
func (e *InvalidValidationError) Error() string {

	if e.Type == nil {
		return "validator: (nil)"
	}

	return "validator: (nil " + e.Type.String() + ")"
}

// ValidationErrors is an array of FieldError's
// for use in custom error messages post validation.
type ValidationErrors []FieldError

// Error is intended for use in development + debugging and not intended to be a production error message.
// It allows ValidationErrors to subscribe to the Error interface.
// All information to create an error message specific to your application is contained within
// the FieldError found within the ValidationErrors array
func (ve ValidationErrors) Error() string {

	buff := bytes.NewBufferString("")

	var err *fieldError

	for i := 0; i < len(ve); i++ {

		err = ve[i].(*fieldError)
		buff.WriteString(err.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// FieldError contains all functions to get error details
type FieldError interface {

	// returns the validation tag that failed. if the
	// validation was an alias, this will return the
	// alias name and not the underlying tag that failed.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "iscolor"
	Tag() string

	// returns the validation tag that failed, even if an
	// alias the actual tag within the alias will be returned.
	// If an 'or' validation fails the entire or will be returned.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "hexcolor|rgb|rgba|hsl|hsla"
	ActualTag() string

	// returns the namespace for the field error, with the tag
	// name taking precedence over the fields actual name.
	//
	// eq. JSON name "User.fname" see ActualNamespace for comparison
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract it's name
	Namespace() string

	// returns the namespace for the field error, with the fields
	// actual name.
	//
	// eq. "User.FirstName" see Namespace for comparison
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract it's name
	ActualNamespace() string

	// returns the fields name with the tag name taking precedence over the
	// fields actual name.
	//
	// eq. JSON name "fname"
	// see ActualField for comparison
	Field() string

	// returns the fields actual name.
	//
	// eq.  "FirstName"
	// see Field for comparison
	ActualField() string

	// returns the actual fields value in case needed for creating the error
	// message
	Value() interface{}

	// returns the param value, already converted into the fields type for
	// comparison; this will also help with generating an error message
	Param() interface{}

	// Kind returns the Field's reflect Kind
	//
	// eg. time.Time's kind is a struct
	Kind() reflect.Kind

	// Type returns the Field's reflect Type
	//
	// // eg. time.Time's type is time.Time
	Type() reflect.Type
}

// compile time interface checks
var _ FieldError = new(fieldError)
var _ error = new(fieldError)

// fieldError contains a single field's validation error along
// with other properties that may be needed for error message creation
// it complies with the FieldError interface
type fieldError struct {
	tag         string
	actualTag   string
	ns          string
	actualNs    string
	field       string
	actualField string
	value       interface{}
	param       interface{}
	kind        reflect.Kind
	typ         reflect.Type
}

// Tag returns the validation tag that failed.
func (fe *fieldError) Tag() string {
	return fe.tag
}

// ActualTag returns the validation tag that failed, even if an
// alias the actual tag within the alias will be returned.
func (fe *fieldError) ActualTag() string {
	return fe.actualTag
}

// Namespace returns the namespace for the field error, with the tag
// name taking precedence over the fields actual name.
func (fe *fieldError) Namespace() string {
	return fe.ns
}

// ActualNamespace returns the namespace for the field error, with the fields
// actual name.
func (fe *fieldError) ActualNamespace() string {
	return fe.actualNs
}

// Field returns the fields name with the tag name taking precedence over the
// fields actual name.
func (fe *fieldError) Field() string {
	return fe.field
}

// ActualField returns the fields actual name.
func (fe *fieldError) ActualField() string {
	return fe.actualField
}

// Value returns the actual fields value in case needed for creating the error
// message
func (fe *fieldError) Value() interface{} {
	return fe.value
}

// Param returns the param value, already converted into the fields type for
// comparison; this will also help with generating an error message
func (fe *fieldError) Param() interface{} {
	return fe.param
}

// Kind returns the Field's reflect Kind
func (fe *fieldError) Kind() reflect.Kind {
	return fe.kind
}

// Type returns the Field's reflect Type
func (fe *fieldError) Type() reflect.Type {
	return fe.typ
}

// Error returns the fieldError's error message
func (fe *fieldError) Error() string {
	return fmt.Sprintf(fieldErrMsg, fe.ns, fe.field, fe.tag)
}
