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
	"time"
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
	Kind     reflect.Kind
	Param    string
	Value    interface{}
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

// ValidationFunc that accepts a value(optional usage), a field and parameter(optional usage) for use in validation
type ValidationFunc func(val interface{}, v interface{}, param string) bool

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

	if structValue.Kind() == reflect.Ptr && !structValue.IsNil() {
		return v.ValidateStruct(structValue.Elem().Interface())
	}

	return v.validateStructRecursive(s, s)
}

// validateStructRecursive validates a struct recursivly and passes the top level struct around for use in validator functions and returns a struct containing the errors
func (v *Validator) validateStructRecursive(top interface{}, s interface{}) *StructValidationErrors {

	structValue := reflect.ValueOf(s)
	structType := reflect.TypeOf(s)
	structName := structType.Name()

	validationErrors := &StructValidationErrors{
		Struct:       structName,
		Errors:       map[string]*FieldValidationError{},
		StructErrors: map[string]*StructValidationErrors{},
	}

	if structValue.Kind() == reflect.Ptr && !structValue.IsNil() {
		return v.validateStructRecursive(top, structValue.Elem().Interface())
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

			if valueField.Type() == reflect.TypeOf(time.Time{}) {

				if fieldError := v.validateFieldByNameAndTagAndValue(top, valueField.Interface(), typeField.Name, tag); fieldError != nil {
					validationErrors.Errors[fieldError.Field] = fieldError
					// free up memory reference
					fieldError = nil
				}

			} else {

				if structErrors := v.ValidateStruct(valueField.Interface()); structErrors != nil {
					validationErrors.StructErrors[typeField.Name] = structErrors
					// free up memory map no longer needed
					structErrors = nil
				}
			}

		default:

			if fieldError := v.validateFieldByNameAndTagAndValue(top, valueField.Interface(), typeField.Name, tag); fieldError != nil {
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

// ValidateFieldByTag allows validation of a single field with the internal validator, still using tag style validation to check multiple errors
func ValidateFieldByTag(f interface{}, tag string) *FieldValidationError {

	return internalValidator.ValidateFieldByTag(f, tag)
}

// ValidateFieldByTag allows validation of a single field, still using tag style validation to check multiple errors
func (v *Validator) ValidateFieldByTag(f interface{}, tag string) *FieldValidationError {

	return v.ValidateFieldByTagAndValue(nil, f, tag)
}

// ValidateFieldByTagAndValue allows validation of a single field with the internal validator, still using tag style validation to check multiple errors
func ValidateFieldByTagAndValue(val interface{}, f interface{}, tag string) *FieldValidationError {

	return internalValidator.ValidateFieldByTagAndValue(val, f, tag)
}

// ValidateFieldByTagAndValue allows validation of a single field, still using tag style validation to check multiple errors
func (v *Validator) ValidateFieldByTagAndValue(val interface{}, f interface{}, tag string) *FieldValidationError {

	return v.validateFieldByNameAndTagAndValue(val, f, "", tag)
}

func (v *Validator) validateFieldByNameAndTagAndValue(val interface{}, f interface{}, name string, tag string) *FieldValidationError {

	// This is a double check if coming from ValidateStruct but need to be here in case function is called directly
	if tag == "-" {
		return nil
	}

	if strings.Contains(tag, omitempty) && !hasValue(val, f, "") {
		return nil
	}

	valueField := reflect.ValueOf(f)

	if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
		return v.validateFieldByNameAndTagAndValue(val, valueField.Elem().Interface(), name, tag)
	}

	switch valueField.Kind() {

	case reflect.Struct, reflect.Interface, reflect.Invalid:

		if valueField.Type() != reflect.TypeOf(time.Time{}) {
			panic("Invalid field passed to ValidateFieldWithTag")
		}
	}

	var valErr *FieldValidationError
	var err error
	valTags := strings.Split(tag, ",")

	for _, valTag := range valTags {

		orVals := strings.Split(valTag, "|")

		if len(orVals) > 1 {

			errTag := ""

			for _, val := range orVals {

				valErr, err = v.validateFieldByNameAndSingleTag(val, f, name, val)

				if err == nil {
					return nil
				}

				errTag += "|" + valErr.ErrorTag

			}

			errTag = strings.TrimLeft(errTag, "|")

			valErr.ErrorTag = errTag
			valErr.Kind = valueField.Kind()

			return valErr
		}

		if valErr, err = v.validateFieldByNameAndSingleTag(val, f, name, valTag); err != nil {
			valErr.Kind = valueField.Kind()

			return valErr
		}
	}

	return nil
}

func (v *Validator) validateFieldByNameAndSingleTag(val interface{}, f interface{}, name string, valTag string) (*FieldValidationError, error) {

	vals := strings.Split(valTag, "=")
	key := strings.Trim(vals[0], " ")

	if len(key) == 0 {
		panic(fmt.Sprintf("Invalid validation tag on field %s", name))
	}

	valErr := &FieldValidationError{
		Field:    name,
		ErrorTag: key,
		Value:    f,
		Param:    "",
	}

	// OK to continue because we checked it's existance before getting into this loop
	if key == omitempty {
		return valErr, nil
	}

	valFunc, ok := v.validationFuncs[key]
	if !ok {
		panic(fmt.Sprintf("Undefined validation function on field %s", name))
	}

	param := ""
	if len(vals) > 1 {
		param = strings.Trim(vals[1], " ")
	}

	if err := valFunc(val, f, param); !err {
		valErr.Param = param
		return valErr, errors.New(key)
	}

	return valErr, nil
}
