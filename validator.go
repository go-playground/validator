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
	"sync"
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
	existsTag           = "exists"
	fieldErrMsg         = "Key: \"%s\" Error:Field validation for \"%s\" failed on the \"%s\" tag"
	arrayIndexFieldName = "%s[%d]"
	mapIndexFieldName   = "%s[%v]"
	invalidValidation   = "Invalid validation tag on field %s"
	undefinedValidation = "Undefined validation function on field %s"
)

var (
	timeType    = reflect.TypeOf(time.Time{})
	timePtrType = reflect.TypeOf(&time.Time{})
	errsPool    = &sync.Pool{New: newValidationErrors}
	tagsCache   = &tagCacheMap{m: map[string][]*tagCache{}}
)

// returns new ValidationErrors to the pool
func newValidationErrors() interface{} {
	return ValidationErrors{}
}

type tagCache struct {
	tagVals [][]string
	isOrVal bool
}

type tagCacheMap struct {
	lock sync.RWMutex
	m    map[string][]*tagCache
}

func (s *tagCacheMap) Get(key string) ([]*tagCache, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *tagCacheMap) Set(key string, value []*tagCache) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[key] = value
}

// Validate contains the validator settings passed in using the Config struct
type Validate struct {
	config Config
}

// Config contains the options that a Validator instance will use.
// It is passed to the New() function
type Config struct {
	TagName         string
	ValidationFuncs map[string]Func
	CustomTypeFuncs map[reflect.Type]CustomTypeFunc
	hasCustomFuncs  bool
}

// CustomTypeFunc allows for overriding or adding custom field type handler functions
// field = field value of the type to return a value to be validated
// example Valuer from sql drive see https://golang.org/src/database/sql/driver/types.go?s=1210:1293#L29
type CustomTypeFunc func(field reflect.Value) interface{}

// Func accepts all values needed for file and cross field validation
// topStruct     = top level struct when validating by struct otherwise nil
// currentStruct = current level struct when validating by struct otherwise optional comparison value
// field         = field value for validation
// param         = parameter used in validation i.e. gt=0 param would be 0
type Func func(topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool

// ValidationErrors is a type of map[string]*FieldError
// it exists to allow for multiple errors to be passed from this library
// and yet still subscribe to the error interface
type ValidationErrors map[string]*FieldError

// Error is intended for use in development + debugging and not intended to be a production error message.
// It allows ValidationErrors to subscribe to the Error interface.
// All information to create an error message specific to your application is contained within
// the FieldError found within the ValidationErrors map
func (ve ValidationErrors) Error() string {

	buff := bytes.NewBufferString("")

	for key, err := range ve {
		buff.WriteString(fmt.Sprintf(fieldErrMsg, key, err.Field, err.Tag))
		buff.WriteString("\n")
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
}

// New creates a new Validate instance for use.
func New(config Config) *Validate {

	if config.CustomTypeFuncs != nil && len(config.CustomTypeFuncs) > 0 {
		config.hasCustomFuncs = true
	}

	return &Validate{config: config}
}

// RegisterValidation adds a validation Func to a Validate's map of validators denoted by the key
// NOTE: if the key already exists, the previous validation function will be replaced.
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

// RegisterCustomTypeFunc registers a CustomTypeFunc against a number of types
func (v *Validate) RegisterCustomTypeFunc(fn CustomTypeFunc, types ...interface{}) {

	if v.config.CustomTypeFuncs == nil {
		v.config.CustomTypeFuncs = map[reflect.Type]CustomTypeFunc{}
	}

	for _, t := range types {
		v.config.CustomTypeFuncs[reflect.TypeOf(t)] = fn
	}

	v.config.hasCustomFuncs = true
}

// Field validates a single field using tag style validation and returns ValidationErrors
// NOTE: it returns ValidationErrors instead of a single FieldError because this can also
// validate Array, Slice and maps fields which may contain more than one error
func (v *Validate) Field(field interface{}, tag string) ValidationErrors {

	errs := errsPool.Get().(ValidationErrors)
	fieldVal := reflect.ValueOf(field)

	v.traverseField(fieldVal, fieldVal, fieldVal, "", errs, false, tag, "")

	if len(errs) == 0 {
		errsPool.Put(errs)
		return nil
	}

	return errs
}

// FieldWithValue validates a single field, against another fields value using tag style validation and returns ValidationErrors
// NOTE: it returns ValidationErrors instead of a single FieldError because this can also
// validate Array, Slice and maps fields which may contain more than one error
func (v *Validate) FieldWithValue(val interface{}, field interface{}, tag string) ValidationErrors {

	errs := errsPool.Get().(ValidationErrors)
	topVal := reflect.ValueOf(val)

	v.traverseField(topVal, topVal, reflect.ValueOf(field), "", errs, false, tag, "")

	if len(errs) == 0 {
		errsPool.Put(errs)
		return nil
	}

	return errs
}

// Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
func (v *Validate) Struct(current interface{}) ValidationErrors {

	errs := errsPool.Get().(ValidationErrors)
	sv := reflect.ValueOf(current)

	v.tranverseStruct(sv, sv, sv, "", errs, true)

	if len(errs) == 0 {
		errsPool.Put(errs)
		return nil
	}

	return errs
}

// tranverseStruct traverses a structs fields and then passes them to be validated by traverseField
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

// traverseField validates any field, be it a struct or single field, ensures it's validity and passes it along to be validated via it's tag options
func (v *Validate) traverseField(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, isStructField bool, tag string, name string) {

	if tag == skipValidationTag {
		return
	}

	kind := current.Kind()

	if kind == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
		kind = current.Kind()
	}

	// this also allows for tags 'required' and 'omitempty' to be used on
	// nested struct fields because when len(tags) > 0 below and the value is nil
	// then required failes and we check for omitempty just before that
	if ((kind == reflect.Ptr || kind == reflect.Interface) && current.IsNil()) || kind == reflect.Invalid {

		if strings.Contains(tag, omitempty) {
			return
		}

		if len(tag) > 0 {

			tags := strings.Split(tag, tagSeparator)
			var param string
			vals := strings.SplitN(tags[0], tagKeySeparator, 2)

			if len(vals) > 1 {
				param = vals[1]
			}

			if kind == reflect.Invalid {
				errs[errPrefix+name] = &FieldError{
					Field: name,
					Tag:   vals[0],
					Param: param,
					Kind:  kind,
				}
				return
			}

			errs[errPrefix+name] = &FieldError{
				Field: name,
				Tag:   vals[0],
				Param: param,
				Value: current.Interface(),
				Kind:  kind,
				Type:  current.Type(),
			}

			return
		}
		// if we get here tag length is zero and we can leave
		if kind == reflect.Invalid {
			return
		}
	}

	typ := current.Type()

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

				if v.config.hasCustomFuncs {
					if fn, ok := v.config.CustomTypeFuncs[typ]; ok {
						v.traverseField(topStruct, currentStruct, reflect.ValueOf(fn(current)), errPrefix, errs, isStructField, tag, name)
						return
					}
				}

				// required passed validation above so stop here
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

	if v.config.hasCustomFuncs {
		if fn, ok := v.config.CustomTypeFuncs[typ]; ok {
			v.traverseField(topStruct, currentStruct, reflect.ValueOf(fn(current)), errPrefix, errs, isStructField, tag, name)
			return
		}
	}

	tags, isCached := tagsCache.Get(tag)

	if !isCached {

		tags = []*tagCache{}

		for _, t := range strings.Split(tag, tagSeparator) {

			if t == diveTag {
				tags = append(tags, &tagCache{tagVals: [][]string{{t}}})
				break
			}

			// if a pipe character is needed within the param you must use the utf8Pipe representation "0x7C"
			orVals := strings.Split(t, orSeparator)
			cTag := &tagCache{isOrVal: len(orVals) > 1, tagVals: make([][]string, len(orVals))}
			tags = append(tags, cTag)

			var key string
			var param string

			for i, val := range orVals {
				vals := strings.SplitN(val, tagKeySeparator, 2)
				key = vals[0]

				if len(key) == 0 {
					panic(strings.TrimSpace(fmt.Sprintf(invalidValidation, name)))
				}

				if len(vals) > 1 {
					param = strings.Replace(strings.Replace(vals[1], utf8HexComma, ",", -1), utf8Pipe, "|", -1)
				}

				cTag.tagVals[i] = []string{key, param}
			}
		}
		tagsCache.Set(tag, tags)
	}

	var dive bool
	var diveSubTag string

	for _, cTag := range tags {

		if cTag.tagVals[0][0] == existsTag {
			continue
		}

		if cTag.tagVals[0][0] == diveTag {
			dive = true
			diveSubTag = strings.TrimLeft(strings.SplitN(tag, diveTag, 2)[1], ",")
			break
		}

		if cTag.tagVals[0][0] == omitempty {

			if !hasValue(topStruct, currentStruct, current, typ, kind, "") {
				return
			}
			continue
		}

		if v.validateField(topStruct, currentStruct, current, typ, kind, errPrefix, errs, cTag, name) {
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

// traverseSlice traverses a Slice or Array's elements and passes them to traverseField for validation
func (v *Validate) traverseSlice(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, tag string, name string) {

	for i := 0; i < current.Len(); i++ {
		v.traverseField(topStruct, currentStruct, current.Index(i), errPrefix, errs, false, tag, fmt.Sprintf(arrayIndexFieldName, name, i))
	}
}

// traverseMap traverses a map's elements and passes them to traverseField for validation
func (v *Validate) traverseMap(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors, tag string, name string) {

	for _, key := range current.MapKeys() {
		v.traverseField(topStruct, currentStruct, current.MapIndex(key), errPrefix, errs, false, tag, fmt.Sprintf(mapIndexFieldName, name, key.Interface()))
	}
}

// validateField validates a field based on the provided tag's key and param values and returns true if there is an error or false if all ok
func (v *Validate) validateField(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, currentType reflect.Type, currentKind reflect.Kind, errPrefix string, errs ValidationErrors, cTag *tagCache, name string) bool {

	var valFunc Func
	var ok bool

	if cTag.isOrVal {

		errTag := ""

		for _, val := range cTag.tagVals {

			valFunc, ok = v.config.ValidationFuncs[val[0]]
			if !ok {
				panic(strings.TrimSpace(fmt.Sprintf(undefinedValidation, name)))
			}

			if valFunc(topStruct, currentStruct, current, currentType, currentKind, val[1]) {
				return false
			}

			errTag += orSeparator + val[0]
		}

		errs[errPrefix+name] = &FieldError{
			Field: name,
			Tag:   errTag[1:],
			Value: current.Interface(),
			Type:  currentType,
			Kind:  currentKind,
		}

		return true
	}

	valFunc, ok = v.config.ValidationFuncs[cTag.tagVals[0][0]]
	if !ok {
		panic(strings.TrimSpace(fmt.Sprintf(undefinedValidation, name)))
	}

	if valFunc(topStruct, currentStruct, current, currentType, currentKind, cTag.tagVals[0][1]) {
		return false
	}

	errs[errPrefix+name] = &FieldError{
		Field: name,
		Tag:   cTag.tagVals[0][0],
		Value: current.Interface(),
		Param: cTag.tagVals[0][1],
		Type:  currentType,
		Kind:  currentKind,
	}

	return true
}
