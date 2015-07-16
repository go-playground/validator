/**
 * Package validator
 *
 * MISC:
 * - anonymous structs - they don't have names so expect the Struct name within StructErrors to be blank
 *
 */

package validator

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
)

const (
	utf8HexComma        = "0x2C"
	utf8Pipe            = "0x7C"
	tagSeparator        = ","
	orSeparator         = "|"
	tagKeySeparator     = "="
	structOnlyTag       = "structonly"
	omitempty           = "omitempty"
	skipValidationTag   = "-"
	diveTag             = "dive"
	fieldErrMsg         = "Key: \"%s\" Error:Field validation for \"%s\" failed on the \"%s\" tag"
	arrayIndexFieldName = "%s[%d]"
	mapIndexFieldName   = "%s[%v]"
	invalidValidation   = "Invalid validation tag on field %s"
	undefinedValidation = "Undefined validation function on field %s"
)

var (
	timeType    = reflect.TypeOf(time.Time{})
	timePtrType = reflect.TypeOf(&time.Time{})
)

// Validate implements the Validate Struct
// NOTE: Fields within are not thread safe and that is on purpose
// Functions and Tags should all be predifined before use, so subscribe to the philosiphy
// or make it thread safe on your end
type Validate struct {
	config Config
}

// Config contains the options that Validator with use
// passed to the New function
type Config struct {
	TagName         string
	ValidationFuncs map[string]Func
}

// Func accepts all values needed for file and cross field validation
// topStruct     = top level struct when validating by struct otherwise nil
// currentStruct = current level struct when validating by struct otherwise optional comparison value
// field         = field value for validation
// param         = parameter used in validation i.e. gt=0 param would be 0
type Func func(topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool

// ValidationErrors is a type of map[string]*FieldError
// it exists to allow for multiple errors passed from this library
// and yet still comply to the error interface
type ValidationErrors map[string]*FieldError

// This is intended for use in development + debugging and not intended to be a production error message.
// It allows ValidationErrors to subscribe to the Error interface.
// All information to create an error message specific to your application is contained within
// the FieldError found in the ValidationErrors
func (ve ValidationErrors) Error() string {

	buff := bytes.NewBufferString("")

	for key, err := range ve {
		buff.WriteString(fmt.Sprintf(fieldErrMsg, key, err.Field, err.Tag))
	}

	return strings.TrimSpace(buff.String())
}

// FieldError contains a single field's validation error along
// with other properties that may be needed for error message creation
type FieldError struct {
	Field string
	Tag   string
	Kind  reflect.Kind
	Type  reflect.Type
	Param string
	Value interface{}
	// IsPlaceholderErr bool
	// IsSliceOrArray   bool
	// IsMap            bool
	// SliceOrArrayErrs map[int]error         // counld be FieldError, StructErrors
	// MapErrs          map[interface{}]error // counld be FieldError, StructErrors
}

// New creates a new Validate instance for use.
func New(config Config) *Validate {

	// structPool = &sync.Pool{New: newStructErrors}

	return &Validate{config: config}
}

// RegisterValidation adds a validation Func to a Validate's map of validators denoted by the key
// NOTE: if the key already exists, it will get replaced.
// NOTE: this method is not thread-safe
func (v *Validate) RegisterValidation(key string, f Func) error {

	if len(key) == 0 {
		return errors.New("Function Key cannot be empty")
	}

	if f == nil {
		return errors.New("Function cannot be empty")
	}

	v.config.ValidationFuncs[key] = f

	return nil
}

// Field allows validation of a single field, still using tag style validation to check multiple errors
func (v *Validate) Field(field interface{}, tag string) ValidationErrors {

	errs := map[string]*FieldError{}
	fieldVal := reflect.ValueOf(field)

	v.traverseField(fieldVal, fieldVal, fieldVal, "", errs, false, tag, "")

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// FieldWithValue allows validation of a single field, possibly even against another fields value, still using tag style validation to check multiple errors
func (v *Validate) FieldWithValue(val interface{}, field interface{}, tag string) ValidationErrors {

	errs := map[string]*FieldError{}
	topVal := reflect.ValueOf(val)

	v.traverseField(topVal, topVal, reflect.ValueOf(field), "", errs, false, tag, "")

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Struct validates a struct, even it's nested structs, and returns a struct containing the errors
// NOTE: Nested Arrays, or Maps of structs do not get validated only the Array or Map itself; the reason is that there is no good
// way to represent or report which struct within the array has the error, besides can validate the struct prior to adding it to
// the Array or Map.
func (v *Validate) Struct(current interface{}) ValidationErrors {

	errs := map[string]*FieldError{}
	sv := reflect.ValueOf(current)

	v.tranverseStruct(sv, sv, sv, "", errs, true)

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (v *Validate) tranverseStruct(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, useStructName bool) {

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	if current.Kind() != reflect.Struct && current.Kind() != reflect.Interface {
		panic("value passed for validation is not a struct")
	}

	typ := current.Type()

	if useStructName {
		errPrefix += typ.Name() + "."
	}

	numFields := current.NumField()

	var fld reflect.StructField

	for i := 0; i < numFields; i++ {
		fld = typ.Field(i)

		if !unicode.IsUpper(rune(fld.Name[0])) {
			continue
		}

		v.traverseField(topStruct, currentStruct, current.Field(i), errPrefix, errs, true, fld.Tag.Get(v.config.TagName), fld.Name)
	}
}

func (v *Validate) traverseField(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, isStructField bool, tag string, name string) {

	if tag == skipValidationTag {
		return
	}

	kind := current.Kind()

	if kind == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
		kind = current.Kind()
	}

	typ := current.Type()

	// this also allows for tags 'required' and 'omitempty' to be used on
	// nested struct fields because when len(tags) > 0 below and the value is nil
	// then required failes and we check for omitempty just before that
	if (kind == reflect.Ptr || kind == reflect.Interface) && current.IsNil() {

		if strings.Contains(tag, omitempty) {
			return
		}

		tags := strings.Split(tag, tagSeparator)

		if len(tags) > 0 {

			var param string
			vals := strings.SplitN(tags[0], tagKeySeparator, 2)

			if len(vals) > 1 {
				param = vals[1]
			}

			errs[errPrefix+name] = &FieldError{
				Field: name,
				Tag:   vals[0],
				Param: param,
				Value: current.Interface(),
				Kind:  kind,
				Type:  typ,
			}

			return
		}
	}

	switch kind {
	case reflect.Struct, reflect.Interface:

		if kind == reflect.Interface {

			current = current.Elem()
			kind = current.Kind()

			if kind == reflect.Ptr && !current.IsNil() {
				current = current.Elem()
				kind = current.Kind()
			}

			// changed current, so have to get inner type again
			typ = current.Type()

			if kind != reflect.Struct {
				goto FALLTHROUGH
			}
		}

		if typ != timeType && typ != timePtrType {

			if kind == reflect.Struct {

				// required passed validationa above so stop here
				// if only validating the structs existance.
				if strings.Contains(tag, structOnlyTag) {
					return
				}

				v.tranverseStruct(topStruct, current, current, errPrefix+name+".", errs, false)
				return
			}
		}
	FALLTHROUGH:
		fallthrough
	default:
		if len(tag) == 0 {
			return
		}
	}

	var dive bool
	var diveSubTag string

	for _, t := range strings.Split(tag, tagSeparator) {

		if t == diveTag {

			dive = true
			diveSubTag = strings.TrimLeft(strings.SplitN(tag, diveTag, 2)[1], ",")
			break
		}

		// no use in checking tags if it's empty and is ok to be
		// omitempty needs to be the first tag if you wish to use it
		if t == omitempty {

			if !hasValue(topStruct, currentStruct, current, typ, kind, "") {
				return
			}
			continue
		}

		var key string
		var param string

		// if a pipe character is needed within the param you must use the utf8Pipe representation "0x7C"
		if strings.Index(t, orSeparator) == -1 {
			vals := strings.SplitN(t, tagKeySeparator, 2)
			key = vals[0]

			if len(key) == 0 {
				panic(strings.TrimSpace(fmt.Sprintf(invalidValidation, name)))
			}

			if len(vals) > 1 {
				param = strings.Replace(strings.Replace(vals[1], utf8HexComma, ",", -1), utf8Pipe, "|", -1)
			}
		} else {
			key = t
		}

		if v.validateField(topStruct, currentStruct, current, typ, kind, errPrefix, errs, key, param, name) {
			return
		}
	}

	if dive {
		// traverse slice or map here
		// or panic ;)
		switch kind {
		case reflect.Slice, reflect.Array:
			v.traverseSlice(topStruct, currentStruct, current, errPrefix, errs, diveSubTag, name)
		case reflect.Map:
			v.traverseMap(topStruct, currentStruct, current, errPrefix, errs, diveSubTag, name)
		default:
			// throw error, if not a slice or map then should not have gotten here
			// bad dive tag
			panic("dive error! can't dive on a non slice or map")
		}
	}
}

func (v *Validate) traverseSlice(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, tag string, name string) {

	for i := 0; i < current.Len(); i++ {

		idxField := current.Index(i)

		if idxField.Kind() == reflect.Ptr && !idxField.IsNil() {
			idxField = idxField.Elem()
		}

		v.traverseField(topStruct, currentStruct, idxField, errPrefix, errs, false, tag, fmt.Sprintf(arrayIndexFieldName, name, i))
	}
}

func (v *Validate) traverseMap(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, tag string, name string) {

	for _, key := range current.MapKeys() {

		idxField := current.MapIndex(key)

		if idxField.Kind() == reflect.Ptr && !idxField.IsNil() {
			idxField = idxField.Elem()
		}

		v.traverseField(topStruct, currentStruct, idxField, errPrefix, errs, false, tag, fmt.Sprintf(mapIndexFieldName, name, key.Interface()))
	}
}

// validateField validates a field based on the provided key tag and param and return true if there is an error false if all ok
func (v *Validate) validateField(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, currentType reflect.Type, currentKind reflect.Kind, errPrefix string, errs ValidationErrors, key string, param string, name string) bool {

	// check if key is orVals, it could be!
	orVals := strings.Split(key, orSeparator)

	if len(orVals) > 1 {

		var errTag string

		for _, val := range orVals {
			vals := strings.SplitN(val, tagKeySeparator, 2)

			if len(vals[0]) == 0 {
				panic(strings.TrimSpace(fmt.Sprintf(invalidValidation, name)))
			}

			param := ""
			if len(vals) > 1 {
				param = strings.Replace(strings.Replace(vals[1], utf8HexComma, ",", -1), utf8Pipe, "|", -1)
			}

			// validate and keep track!
			valFunc, ok := v.config.ValidationFuncs[vals[0]]
			if !ok {
				panic(strings.TrimSpace(fmt.Sprintf(undefinedValidation, name)))
			}

			if valFunc(topStruct, currentStruct, current, currentType, currentKind, param) {
				return false
			}

			errTag += orSeparator + vals[0]
		}

		errs[errPrefix+name] = &FieldError{
			Field: name,
			Tag:   errTag[1:],
			Value: current.Interface(),
			Param: param,
			Type:  currentType,
			Kind:  currentKind,
		}

		return true
	}

	valFunc, ok := v.config.ValidationFuncs[key]
	if !ok {
		panic(strings.TrimSpace(fmt.Sprintf(undefinedValidation, name)))
	}

	if valFunc(topStruct, currentStruct, current, currentType, currentKind, param) {
		return false
	}

	errs[errPrefix+name] = &FieldError{
		Field: name,
		Tag:   key,
		Value: current.Interface(),
		Param: param,
		Type:  currentType,
		Kind:  currentKind,
	}

	return true
}
