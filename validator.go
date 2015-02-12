package validator

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// FieldValidationError contains a single fields validation error
type FieldValidationError struct {
	Field    string
	ErrorTag string
}

// StructValidationErrors is a slice of errors for struct fields ( Excluding struct fields)
// NOTE: if a field within a struct is a struct it's errors will not be contained within the current
// StructValidationErrors but rather a new ArrayValidationErrors is created for each struct
type StructValidationErrors struct {
	Errors []FieldValidationError
}

// ArrayStructValidationErrors is a struct that contains a 2D flattened list of struct specific StructValidationErrors
type ArrayStructValidationErrors struct {
	// Key = Struct Name
	Errors map[string][]StructValidationErrors
}

// ValidationFunc that accepts the value of a field and parameter for use in validation (parameter not always used or needed)
type ValidationFunc func(v interface{}, param string) error

// Validator implements the Validator Struct
// NOTE: Fields within are not thread safe and that is on purpose
// Functions Tags etc. should all be predifined before use, so subscribe to the philosiphy
// or make it thread safe on your end
type Validator struct {
	// TagName being used.
	tagName string
	// validationFuncs is a map of validation functions and the tag keys
	validationFuncs map[string]ValidationFunc
}

var bakedInValidators = map[string]ValidationFunc{}

var internalValidator = NewValidator("validate", bakedInValidators)

// NewValidator creates a new Validator instance
// NOTE: it is not necessary to create a new validator as the internal one will do in 99.9% of cases, but the option is there.
func NewValidator(tagName string, funcs map[string]ValidationFunc) *Validator {
	return &Validator{
		tagName:         tagName,
		validationFuncs: funcs,
	}
}

// SetTag sets the baked in Validator's tagName to one of your choosing
func SetTag(tagName string) {
	internalValidator.SetTag(tagName)
}

// SetTag sets tagName of the Validator to one of your choosing
func (v *Validator) SetTag(tagName string) {
	v.tagName = tagName
}

// AddFunction adds a ValidationFunc to the baked in Validator's map of validators denoted by the key
func AddFunction(key string, f ValidationFunc) error {
	return internalValidator.AddFunction(key, f)
}

// AddFunction adds a ValidationFunc to a Validator's map of validators denoted by the key
func (v *Validator) AddFunction(key string, f ValidationFunc) error {

	if len(key) == 0 {
		return errors.New("Function Key cannot be empty")
	}

	if f == nil {
		return errors.New("Function Key cannot be empty")
	}

	// Commented out to allow overwritting of Baked In Function if so desired.
	// if v.ValidationFuncs[key] != nil {
	// 	return errors.New(fmt.Sprintf("Validation Function with key: %s already exists.", key))
	// }

	v.validationFuncs[key] = f

	return nil
}

func ValidateStruct(s interface{}) ArrayStructValidationErrors {

	return internalValidator.ValidateStruct(s)
}

func (v *Validator) ValidateStruct(s interface{}) ArrayStructValidationErrors {

	var errorStruct = ArrayStructValidationErrors{}

	structValue := reflect.ValueOf(s)
	structType := reflect.TypeOf(s)

	if structValue.Kind() == reflect.Ptr && !structValue.IsNil() {
		return v.ValidateStruct(structValue.Elem().Interface())
	}

	if structValue.Kind() != reflect.Struct {
		log.Fatal("interface passed for validation is not a struct")
	}

	var numFields = structValue.NumField()

	for i := 0; i < numFields; i++ {

		valueField := structValue.Field(i)
		typeField := structType.Field(i)

		if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
			valueField = valueField.Elem()
		}

		fmt.Println(typeField.Name)
	}

	return errorStruct
}
