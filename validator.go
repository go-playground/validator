package validator

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"unicode"
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
	Struct string
	Errors []*FieldValidationError
}

// // ArrayStructValidationErrors is a struct that contains a 2D flattened list of struct specific StructValidationErrors
// type ArrayStructValidationErrors struct {
// 	// Key = Struct Name
// 	Errors map[string][]StructValidationErrors
// }

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

// ValidateStruct validates a struct and returns a struct containing the errors
func ValidateStruct(s interface{}) []*StructValidationErrors {

	return internalValidator.ValidateStruct(s)
}

// ValidateStruct validates a struct and returns a struct containing the errors
func (v *Validator) ValidateStruct(s interface{}) []*StructValidationErrors {

	errorArray := []*StructValidationErrors{}
	structValue := reflect.ValueOf(s)
	structType := reflect.TypeOf(s)
	structName := structType.Name()

	var currentStructError = &StructValidationErrors{
		Struct: structName,
	}

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

		tag := typeField.Tag.Get(v.tagName)

		if tag == "-" {
			continue
		}

		// if no validation and not a struct (which may containt fields for validation)
		if tag == "" && valueField.Kind() != reflect.Struct {
			continue
		}

		switch valueField.Kind() {

		case reflect.Struct:

			if !unicode.IsUpper(rune(typeField.Name[0])) {
				continue
			}

			if structErrors := v.ValidateStruct(valueField.Interface()); structErrors != nil {
				errorArray = append(errorArray, structErrors...)
			}

		default:

			if fieldError := v.validateStructFieldByTag(valueField.Interface(), typeField.Name, tag); fieldError != nil {
				currentStructError.Errors = append(currentStructError.Errors, fieldError)
			}
		}
	}

	if currentStructError.Errors != nil {
		errorArray = append(errorArray, currentStructError)
	}

	if len(errorArray) == 0 {
		return nil
	}

	return errorArray
}

// ValidateFieldWithTag validates the given field by the given tag arguments
func (v *Validator) validateStructFieldByTag(f interface{}, name string, tag string) *FieldValidationError {

	// // This is a double check if coming from ValidateStruct but need to be here in case function is called directly
	// if tag == "-" {
	// 	return nil
	// }
	//
	// valueField := reflect.ValueOf(f)
	//
	// if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
	// 	return v.validateStructFieldByTag(valueField.Elem().Interface(), name, tag)
	// }
	//
	// // fmt.Println(typeField.Name)
	//
	// switch valueField.Kind() {
	//
	// case reflect.Struct, reflect.Invalid:
	// 	log.Fatal("Invalid field passed to ValidateFieldWithTag")
	// }
	//
	// // typeField := reflect.TypeOf(f)
	// // name := ""
	// valTags := strings.Split(tag, ",")
	//
	// for _, valTag := range valTags {
	//
	// 	vals := strings.Split(valTag, "=")
	// 	key := strings.Trim(vals[0], " ")
	//
	// 	if len(key) == 0 {
	// 		log.Fatalf("Invalid validation tag on field %s", name)
	// 	}
	//
	// 	valFunc := v.validationFuncs[key]
	// 	if valFunc == nil {
	// 		log.Fatalf("Undefined validation function on field %s", name)
	// 	}
	//
	// 	param := ""
	// 	if len(vals) > 1 {
	// 		param = strings.Trim(vals[1], " ")
	// 	}
	//
	// 	if err := valFunc(f, param); !err {
	//
	// 		return &FieldValidationError{
	// 			Field:    name,
	// 			ErrorTag: key,
	// 		}
	// 	}
	//
	// }
	//
	// return nil
	//
	//

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

	valueField := reflect.ValueOf(f)

	if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
		return v.ValidateFieldByTag(valueField.Elem().Interface(), tag)
	}

	// fmt.Println(typeField.Name)

	switch valueField.Kind() {

	case reflect.Struct, reflect.Invalid:
		log.Fatal("Invalid field passed to ValidateFieldWithTag")
	}

	// typeField := reflect.TypeOf(f)
	// name := ""
	valTags := strings.Split(tag, ",")

	for _, valTag := range valTags {

		vals := strings.Split(valTag, "=")
		key := strings.Trim(vals[0], " ")

		if len(key) == 0 {
			log.Fatalf("Invalid validation tag on field %s", name)
		}

		valFunc := v.validationFuncs[key]
		if valFunc == nil {
			log.Fatalf("Undefined validation function on field %s", name)
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
