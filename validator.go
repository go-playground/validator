/**
 * Package validator
 *
 * MISC:
 * - anonymous structs - they don't have names so expect the Struct name within StructValidationErrors to be blank
 *
 */

package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

const (
	defaultTagName         = "validate"
	omitempty              = "omitempty"
	validationFieldErrMsg  = "Field validation for \"%s\" failed on the \"%s\" tag\n"
	validationStructErrMsg = "Struct:%s\n"
)

// FieldValidationError contains a single fields validation error
type FieldValidationError struct {
	Field    string
	ErrorTag string
}

// This is intended for use in development + debugging and not intended to be a production error message.
// it also allows FieldValidationError to be used as an Error interface
func (e *FieldValidationError) Error() string {
	return fmt.Sprintf(validationFieldErrMsg, e.Field, e.ErrorTag)
}

// StructValidationErrors is hierarchical list of field and struct errors
type StructValidationErrors struct {
	// Name of the Struct
	Struct string
	// Struct Field Errors
	Errors map[string]*FieldValidationError
	// Struct Fields of type struct and their errors
	// key = Field Name of current struct, but internally Struct will be the actual struct name unless anonymous struct, it will be blank
	StructErrors map[string]*StructValidationErrors
}

// This is intended for use in development + debugging and not intended to be a production error message.
// it also allows StructValidationErrors to be used as an Error interface
func (e *StructValidationErrors) Error() string {

	s := fmt.Sprintf(validationStructErrMsg, e.Struct)

	for _, err := range e.Errors {
		s += err.Error()
	}

	for _, sErr := range e.StructErrors {
		s += sErr.Error()
	}

	return fmt.Sprintf("%s\n\n", s)
}

// Flatten flattens the StructValidationErrors hierarchical sctructure into a flat namespace style field name
// for those that want/need it
func (e *StructValidationErrors) Flatten() map[string]*FieldValidationError {

	if e == nil {
		return nil
	}

	errs := map[string]*FieldValidationError{}

	for _, f := range e.Errors {

		errs[f.Field] = f
	}

	for key, val := range e.StructErrors {

		otherErrs := val.Flatten()

		for _, f2 := range otherErrs {

			f2.Field = fmt.Sprintf("%s.%s", key, f2.Field)
			errs[f2.Field] = f2
		}
	}

	return errs
}

// ValidationFunc that accepts the value of a field and parameter for use in validation (parameter not always used or needed)
type ValidationFunc func(v interface{}, param string) bool

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

// var bakedInValidators = map[string]ValidationFunc{}

var internalValidator = NewValidator(defaultTagName, BakedInValidators)

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

// ValidateStruct validates a struct and returns a struct containing the errors
func ValidateStruct(s interface{}) *StructValidationErrors {

	return internalValidator.ValidateStruct(s)
}

// ValidateStruct validates a struct and returns a struct containing the errors
func (v *Validator) ValidateStruct(s interface{}) *StructValidationErrors {

	structValue := reflect.ValueOf(s)
	structType := reflect.TypeOf(s)
	structName := structType.Name()

	validationErrors := &StructValidationErrors{
		Struct:       structName,
		Errors:       map[string]*FieldValidationError{},
		StructErrors: map[string]*StructValidationErrors{},
	}

	if structValue.Kind() == reflect.Ptr && !structValue.IsNil() {
		return v.ValidateStruct(structValue.Elem().Interface())
	}

	if structValue.Kind() != reflect.Struct && structValue.Kind() != reflect.Interface {
		panic("interface passed for validation is not a struct")
	}

	var numFields = structValue.NumField()

	for i := 0; i < numFields; i++ {

		valueField := structValue.Field(i)
		typeField := structType.Field(i)

		if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
			valueField = valueField.Elem()
		}

		tag := typeField.Tag.Get(v.tagName)

		if tag == "-" {
			continue
		}

		// if no validation and not a struct (which may containt fields for validation)
		if tag == "" && valueField.Kind() != reflect.Struct && valueField.Kind() != reflect.Interface {
			continue
		}

		switch valueField.Kind() {

		case reflect.Struct, reflect.Interface:

			if !unicode.IsUpper(rune(typeField.Name[0])) {
				continue
			}

			if structErrors := v.ValidateStruct(valueField.Interface()); structErrors != nil {
				validationErrors.StructErrors[typeField.Name] = structErrors
				// free up memory map no longer needed
				structErrors = nil
			}

		default:

			if fieldError := v.validateStructFieldByTag(valueField.Interface(), typeField.Name, tag); fieldError != nil {
				validationErrors.Errors[fieldError.Field] = fieldError
				// free up memory reference
				fieldError = nil
			}
		}
	}

	if len(validationErrors.Errors) == 0 && len(validationErrors.StructErrors) == 0 {
		return nil
	}

	return validationErrors
}

// ValidateFieldWithTag validates the given field by the given tag arguments
func (v *Validator) validateStructFieldByTag(f interface{}, name string, tag string) *FieldValidationError {

	if err := v.validateFieldByNameAndTag(f, name, tag); err != nil {
		return &FieldValidationError{
			Field:    name,
			ErrorTag: err.Error(),
		}
	}

	return nil
}

// ValidateFieldByTag allows validation of a single field with the internal validator, still using tag style validation to check multiple errors
func ValidateFieldByTag(f interface{}, tag string) error {

	return internalValidator.validateFieldByNameAndTag(f, "", tag)
}

// ValidateFieldByTag allows validation of a single field, still using tag style validation to check multiple errors
func (v *Validator) ValidateFieldByTag(f interface{}, tag string) error {

	return v.validateFieldByNameAndTag(f, "", tag)
}

func (v *Validator) validateFieldByNameAndTag(f interface{}, name string, tag string) error {

	// This is a double check if coming from ValidateStruct but need to be here in case function is called directly
	if tag == "-" {
		return nil
	}

	if strings.Contains(tag, omitempty) && !required(f, "") {
		return nil
	}

	valueField := reflect.ValueOf(f)

	if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
		return v.ValidateFieldByTag(valueField.Elem().Interface(), tag)
	}

	switch valueField.Kind() {

	case reflect.Struct, reflect.Invalid:
		panic("Invalid field passed to ValidateFieldWithTag")
	}

	// TODO: validate commas in regex's
	valTags := strings.Split(tag, ",")

	for _, valTag := range valTags {

		// TODO: validate = in regex's
		vals := strings.Split(valTag, "=")
		key := strings.Trim(vals[0], " ")

		if len(key) == 0 {
			panic(fmt.Sprintf("Invalid validation tag on field %s", name))
		}

		// OK to continue because we checked it's existance before getting into this loop
		if key == omitempty {
			continue
		}

		valFunc, ok := v.validationFuncs[key]
		if !ok {
			panic(fmt.Sprintf("Undefined validation function on field %s", name))
		}

		param := ""
		if len(vals) > 1 {
			param = strings.Trim(vals[1], " ")
		}

		if err := valFunc(f, param); !err {

			return errors.New(key)
		}

	}

	return nil
}
