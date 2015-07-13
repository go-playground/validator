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
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	utf8HexComma      = "0x2C"
	tagSeparator      = ","
	orSeparator       = "|"
	tagKeySeparator   = "="
	structOnlyTag     = "structonly"
	omitempty         = "omitempty"
	skipValidationTag = "-"
	fieldErrMsg       = "Key: \"%s\" Error:Field validation for \"%s\" failed on the \"%s\" tag"
	invaldField       = "Invalid field passed to traverseField"
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
// top     = top level struct when validating by struct otherwise nil
// current = current level struct when validating by struct otherwise optional comparison value
// f       = field value for validation
// param   = parameter used in validation i.e. gt=0 param would be 0
type Func func(top interface{}, current interface{}, f interface{}, param string) bool

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

// Struct validates a struct, even it's nested structs, and returns a struct containing the errors
// NOTE: Nested Arrays, or Maps of structs do not get validated only the Array or Map itself; the reason is that there is no good
// way to represent or report which struct within the array has the error, besides can validate the struct prior to adding it to
// the Array or Map.
func (v *Validate) Struct(current interface{}) ValidationErrors {

	errs := map[string]*FieldError{}
	sv := reflect.ValueOf(current)

	v.tranverseStruct(sv, sv, sv, "", errs)

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (v *Validate) tranverseStruct(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors) {

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	if current.Kind() != reflect.Struct && current.Kind() != reflect.Interface {
		panic("value passed for validation is not a struct")
	}

	typ := current.Type()
	errPrefix += typ.Name() + "."
	numFields := current.NumField()

	for i := 0; i < numFields; i++ {
		v.traverseField(topStruct, currentStruct, current.Field(i), errPrefix, errs, true, typ.Field(i).Tag.Get(v.config.TagName))
	}
}

func (v *Validate) traverseField(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, isStructField bool, tag string) {

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

			errs[errPrefix+typ.Name()] = &FieldError{
				Field: typ.Name(),
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

	case reflect.Invalid:
		panic(invaldField)
	case reflect.Struct, reflect.Interface:

		if kind == reflect.Interface {

			current = current.Elem()
			kind = current.Kind()

			if kind == reflect.Ptr && !current.IsNil() {
				current = current.Elem()
				kind = current.Kind()
			}

			if kind != reflect.Struct {
				goto FALLTHROUGH
			}
		}

		if typ != timeType && typ != timePtrType {

			if isStructField {

				// required passed validationa above so stop here
				// if only validating the structs existance.
				if strings.Contains(tag, structOnlyTag) {
					return
				}

				v.tranverseStruct(topStruct, current, current, errPrefix, errs)
				return
			}

			panic(invaldField)
		}
	FALLTHROUGH:
		fallthrough
	default:
		if len(tag) == 0 {
			return
		}
	}

	// for _, t := range strings.Split(tag, tagSeparator) {

	// 			if t == diveTag {

	// 				cField.dive = true
	// 				cField.diveTag = strings.TrimLeft(strings.SplitN(tag, diveTag, 2)[1], ",")
	// 				break
	// 			}

	// 			orVals := strings.Split(t, orSeparator)
	// 			cTag := &cachedTags{isOrVal: len(orVals) > 1, keyVals: make([][]string, len(orVals))}
	// 			cField.tags = append(cField.tags, cTag)

	// 			for i, val := range orVals {
	// 				vals := strings.SplitN(val, tagKeySeparator, 2)

	// 				key := strings.TrimSpace(vals[0])

	// 				if len(key) == 0 {
	// 					panic(fmt.Sprintf("Invalid validation tag on field %s", name))
	// 				}

	// 				param := ""
	// 				if len(vals) > 1 {
	// 					param = strings.Replace(vals[1], utf8HexComma, ",", -1)
	// 				}

	// 				cTag.keyVals[i] = []string{key, param}
	// 			}
	// 		}
}
