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
	utf8HexComma    = "0x2C"
	tagSeparator    = ","
	orSeparator     = "|"
	noValidationTag = "-"
	tagKeySeparator = "="
	structOnlyTag   = "structonly"
	omitempty       = "omitempty"
	required        = "required"
	fieldErrMsg     = "Field validation for \"%s\" failed on the \"%s\" tag"
	sliceErrMsg     = "Field validation for \"%s\" index \"%d\" failed on the \"%s\" tag"
	mapErrMsg       = "Field validation for \"%s\" key \"%s\" failed on the \"%s\" tag"
	structErrMsg    = "Struct:%s\n"
	diveTag         = "dive"
	diveSplit       = "," + diveTag
)

var structPool *pool

// Pool holds a channelStructErrors.
type pool struct {
	pool chan *StructErrors
}

// NewPool creates a new pool of Clients.
func newPool(max int) *pool {
	return &pool{
		pool: make(chan *StructErrors, max),
	}
}

// Borrow a StructErrors from the pool.
func (p *pool) Borrow() *StructErrors {
	var c *StructErrors

	select {
	case c = <-p.pool:
	default:
		c = &StructErrors{
			Errors:       map[string]*FieldError{},
			StructErrors: map[string]*StructErrors{},
		}
	}

	return c
}

// Return returns a StructErrors to the pool.
func (p *pool) Return(c *StructErrors) {

	select {
	case p.pool <- c:
	default:
		// let it go, let it go...
	}
}

type cachedTags struct {
	keyVals [][]string
	isOrVal bool
}

type cachedField struct {
	index          int
	name           string
	tags           []*cachedTags
	tag            string
	kind           reflect.Kind
	typ            reflect.Type
	isTime         bool
	isSliceOrArray bool
	isMap          bool
	isTimeSubtype  bool
	sliceSubtype   reflect.Type
	mapSubtype     reflect.Type
	sliceSubKind   reflect.Kind
	mapSubKind     reflect.Kind
	// DiveMaxDepth   uint64            // zero means no depth
	// DiveTags map[uint64]string // map of dive depth and associated tag as string]
	diveTag string
}

type cachedStruct struct {
	children int
	name     string
	kind     reflect.Kind
	fields   []*cachedField
}

type structsCacheMap struct {
	lock sync.RWMutex
	m    map[reflect.Type]*cachedStruct
}

func (s *structsCacheMap) Get(key reflect.Type) (*cachedStruct, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *structsCacheMap) Set(key reflect.Type, value *cachedStruct) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[key] = value
}

var structCache = &structsCacheMap{m: map[reflect.Type]*cachedStruct{}}

type fieldsCacheMap struct {
	lock sync.RWMutex
	m    map[string][]*cachedTags
}

func (s *fieldsCacheMap) Get(key string) ([]*cachedTags, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *fieldsCacheMap) Set(key string, value []*cachedTags) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[key] = value
}

var fieldsCache = &fieldsCacheMap{m: map[string][]*cachedTags{}}

// // SliceError contains a fields error for a single index within an array or slice
// // NOTE: library only checks the first dimension of the array so if you have a multidimensional
// // array [][]string that validations after the "dive" tag are applied to []string not each
// // string within it. It is not a dificulty with traversing the chain, but how to add validations
// // to what dimension of an array and even how to report on them in any meaningful fashion.
// type SliceError struct {
// 	Index uint64
// 	Field string
// 	Tag   string
// 	Kind  reflect.Kind
// 	Type  reflect.Type
// 	Param string
// 	Value interface{}
// }

// // This is intended for use in development + debugging and not intended to be a production error message.
// // it also allows SliceError to be used as an Error interface
// func (e *SliceError) Error() string {
// 	return fmt.Sprintf(sliceErrMsg, e.Field, e.Index, e.Tag)
// }

// // MapError contains a fields error for a single key within a map
// // NOTE: library only checks the first dimension of the array so if you have a multidimensional
// // array [][]string that validations after the "dive" tag are applied to []string not each
// // string within it. It is not a dificulty with traversing the chain, but how to add validations
// // to what dimension of an array and even how to report on them in any meaningful fashion.
// type MapError struct {
// 	Key   interface{}
// 	Field string
// 	Tag   string
// 	Kind  reflect.Kind
// 	Type  reflect.Type
// 	Param string
// 	Value interface{}
// }

// // This is intended for use in development + debugging and not intended to be a production error message.
// // it also allows MapError to be used as an Error interface
// func (e *MapError) Error() string {
// 	return fmt.Sprintf(mapErrMsg, e.Field, e.Key, e.Tag)
// }

// FieldError contains a single field's validation error along
// with other properties that may be needed for error message creation
type FieldError struct {
	Field            string
	Tag              string
	Kind             reflect.Kind
	Type             reflect.Type
	Param            string
	Value            interface{}
	isPlaceholderErr bool
	IsSliceOrArray   bool
	IsMap            bool
	// Key              interface{}
	// Index            int
	SliceOrArrayErrs []error               // counld be FieldError, StructErrors
	MapErrs          map[interface{}]error // counld be FieldError, StructErrors
}

// This is intended for use in development + debugging and not intended to be a production error message.
// it also allows FieldError to be used as an Error interface
func (e *FieldError) Error() string {
	return fmt.Sprintf(fieldErrMsg, e.Field, e.Tag)
}

// StructErrors is hierarchical list of field and struct validation errors
// for a non hierarchical representation please see the Flatten method for StructErrors
type StructErrors struct {
	// Name of the Struct
	Struct string
	// Struct Field Errors
	Errors map[string]*FieldError
	// Struct Fields of type struct and their errors
	// key = Field Name of current struct, but internally Struct will be the actual struct name unless anonymous struct, it will be blank
	StructErrors map[string]*StructErrors

	// Index int
	// Key   interface{}
	// IsSliceOrArrayError bool
	// IsMapError          bool
	// Key                 interface{}
	// Index               uint64
}

// This is intended for use in development + debugging and not intended to be a production error message.
// it also allows StructErrors to be used as an Error interface
func (e *StructErrors) Error() string {
	buff := bytes.NewBufferString(fmt.Sprintf(structErrMsg, e.Struct))

	for _, err := range e.Errors {
		buff.WriteString(err.Error())
		buff.WriteString("\n")
	}

	for _, err := range e.StructErrors {
		buff.WriteString(err.Error())
	}

	return buff.String()
}

// Flatten flattens the StructErrors hierarchical structure into a flat namespace style field name
// for those that want/need it
func (e *StructErrors) Flatten() map[string]*FieldError {

	if e == nil {
		return nil
	}

	errs := map[string]*FieldError{}

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

// Func accepts all values needed for file and cross field validation
// top     = top level struct when validating by struct otherwise nil
// current = current level struct when validating by struct otherwise optional comparison value
// f       = field value for validation
// param   = parameter used in validation i.e. gt=0 param would be 0
type Func func(top interface{}, current interface{}, f interface{}, param string) bool

// Validate implements the Validate Struct
// NOTE: Fields within are not thread safe and that is on purpose
// Functions and Tags should all be predifined before use, so subscribe to the philosiphy
// or make it thread safe on your end
type Validate struct {
	// tagName being used.
	tagName string
	// validateFuncs is a map of validation functions and the tag keys
	validationFuncs map[string]Func
}

// New creates a new Validate instance for use.
func New(tagName string, funcs map[string]Func) *Validate {

	structPool = newPool(10)

	return &Validate{
		tagName:         tagName,
		validationFuncs: funcs,
	}
}

// SetTag sets tagName of the Validator to one of your choosing after creation
// perhaps to dodge a tag name conflict in a specific section of code
// NOTE: this method is not thread-safe
func (v *Validate) SetTag(tagName string) {
	v.tagName = tagName
}

// SetMaxStructPoolSize sets the  struct pools max size. this may be usefull for fine grained
// performance tuning towards your application, however, the default should be fine for
// nearly all cases. only increase if you have a deeply nested struct structure.
// NOTE: this method is not thread-safe
// NOTE: this is only here to keep compatibility with v5, in v6 the method will be removed
// and the max pool size will be passed into the New function
func (v *Validate) SetMaxStructPoolSize(max int) {
	structPool = newPool(max)
}

// AddFunction adds a validation Func to a Validate's map of validators denoted by the key
// NOTE: if the key already exists, it will get replaced.
// NOTE: this method is not thread-safe
func (v *Validate) AddFunction(key string, f Func) error {

	if len(key) == 0 {
		return errors.New("Function Key cannot be empty")
	}

	if f == nil {
		return errors.New("Function cannot be empty")
	}

	v.validationFuncs[key] = f

	return nil
}

// Struct validates a struct, even it's nested structs, and returns a struct containing the errors
// NOTE: Nested Arrays, or Maps of structs do not get validated only the Array or Map itself; the reason is that there is no good
// way to represent or report which struct within the array has the error, besides can validate the struct prior to adding it to
// the Array or Map.
func (v *Validate) Struct(s interface{}) *StructErrors {

	return v.structRecursive(s, s, s)
}

// structRecursive validates a struct recursivly and passes the top level and current struct around for use in validator functions and returns a struct containing the errors
func (v *Validate) structRecursive(top interface{}, current interface{}, s interface{}) *StructErrors {

	structValue := reflect.ValueOf(s)

	if structValue.Kind() == reflect.Ptr && !structValue.IsNil() {
		return v.structRecursive(top, current, structValue.Elem().Interface())
	}

	if structValue.Kind() != reflect.Struct && structValue.Kind() != reflect.Interface {
		panic("interface passed for validation is not a struct")
	}

	structType := reflect.TypeOf(s)

	var structName string
	var numFields int
	var cs *cachedStruct
	var isCached bool

	cs, isCached = structCache.Get(structType)

	if isCached {
		structName = cs.name
		numFields = cs.children
	} else {
		structName = structType.Name()
		numFields = structValue.NumField()
		cs = &cachedStruct{name: structName, children: numFields}
		structCache.Set(structType, cs)
	}

	validationErrors := structPool.Borrow()
	validationErrors.Struct = structName

	for i := 0; i < numFields; i++ {

		var valueField reflect.Value
		var cField *cachedField
		var typeField reflect.StructField

		if isCached {
			cField = cs.fields[i]
			valueField = structValue.Field(cField.index)

			if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
				valueField = valueField.Elem()
			}
		} else {
			valueField = structValue.Field(i)

			if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
				valueField = valueField.Elem()
			}

			typeField = structType.Field(i)

			cField = &cachedField{index: i, tag: typeField.Tag.Get(v.tagName), isTime: valueField.Type() == reflect.TypeOf(time.Time{})}

			if cField.tag == noValidationTag {
				cs.children--
				continue
			}

			// if no validation and not a struct (which may containt fields for validation)
			if cField.tag == "" && ((valueField.Kind() != reflect.Struct && valueField.Kind() != reflect.Interface) || valueField.Type() == reflect.TypeOf(time.Time{})) {
				cs.children--
				continue
			}

			cField.name = typeField.Name
			cField.kind = valueField.Kind()
			cField.typ = valueField.Type()
		}

		// this can happen if the first cache value was nil
		// but the second actually has a value
		if cField.kind == reflect.Ptr {
			cField.kind = valueField.Kind()
		}

		switch cField.kind {

		case reflect.Struct, reflect.Interface:

			if !unicode.IsUpper(rune(cField.name[0])) {
				cs.children--
				continue
			}

			if cField.isTime {

				// cField.isTime = true

				if fieldError := v.fieldWithNameAndValue(top, current, valueField.Interface(), cField.tag, cField.name, false, cField); fieldError != nil {
					validationErrors.Errors[fieldError.Field] = fieldError
					// free up memory reference
					fieldError = nil
				}

			} else {

				if strings.Contains(cField.tag, structOnlyTag) {
					cs.children--
					continue
				}

				if valueField.Kind() == reflect.Ptr && valueField.IsNil() {

					if strings.Contains(cField.tag, omitempty) {
						continue
					}

					if strings.Contains(cField.tag, required) {

						validationErrors.Errors[cField.name] = &FieldError{
							Field: cField.name,
							Tag:   required,
							Value: valueField.Interface(),
						}

						continue
					}
				}

				if structErrors := v.structRecursive(top, valueField.Interface(), valueField.Interface()); structErrors != nil {
					validationErrors.StructErrors[cField.name] = structErrors
					// free up memory map no longer needed
					structErrors = nil
				}
			}

		case reflect.Slice, reflect.Array:
			cField.isSliceOrArray = true
			cField.sliceSubtype = cField.typ.Elem()
			cField.isTimeSubtype = cField.sliceSubtype == reflect.TypeOf(time.Time{})
			cField.sliceSubKind = cField.sliceSubtype.Kind()

			if fieldError := v.fieldWithNameAndValue(top, current, valueField.Interface(), cField.tag, cField.name, false, cField); fieldError != nil {
				validationErrors.Errors[fieldError.Field] = fieldError
				// free up memory reference
				fieldError = nil
			}

		case reflect.Map:
			cField.isMap = true
			cField.mapSubtype = cField.typ.Elem()
			cField.isTimeSubtype = cField.mapSubtype == reflect.TypeOf(time.Time{})
			cField.mapSubKind = cField.mapSubtype.Kind()

			if fieldError := v.fieldWithNameAndValue(top, current, valueField.Interface(), cField.tag, cField.name, false, cField); fieldError != nil {
				validationErrors.Errors[fieldError.Field] = fieldError
				// free up memory reference
				fieldError = nil
			}

		default:
			if fieldError := v.fieldWithNameAndValue(top, current, valueField.Interface(), cField.tag, cField.name, false, cField); fieldError != nil {
				validationErrors.Errors[fieldError.Field] = fieldError
				// free up memory reference
				fieldError = nil
			}
		}

		if !isCached {
			cs.fields = append(cs.fields, cField)
		}
	}

	if len(validationErrors.Errors) == 0 && len(validationErrors.StructErrors) == 0 {
		structPool.Return(validationErrors)
		return nil
	}

	return validationErrors
}

// Field allows validation of a single field, still using tag style validation to check multiple errors
func (v *Validate) Field(f interface{}, tag string) *FieldError {

	return v.FieldWithValue(nil, f, tag)
}

// FieldWithValue allows validation of a single field, possibly even against another fields value, still using tag style validation to check multiple errors
func (v *Validate) FieldWithValue(val interface{}, f interface{}, tag string) *FieldError {

	return v.fieldWithNameAndValue(nil, val, f, tag, "", true, nil)
}

func (v *Validate) fieldWithNameAndValue(val interface{}, current interface{}, f interface{}, tag string, name string, isSingleField bool, cacheField *cachedField) *FieldError {

	var cField *cachedField
	var isCached bool
	// var isInDive bool
	var valueField reflect.Value

	// This is a double check if coming from validate.Struct but need to be here in case function is called directly
	if tag == noValidationTag {
		return nil
	}

	if strings.Contains(tag, omitempty) && !hasValue(val, current, f, "") {
		return nil
	}

	valueField = reflect.ValueOf(f)

	if cacheField == nil {
		// valueField = reflect.ValueOf(f)

		if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
			valueField = valueField.Elem()
			f = valueField.Interface()
		}

		cField = &cachedField{name: name, kind: valueField.Kind(), tag: tag, typ: valueField.Type()}

		switch cField.kind {
		case reflect.Slice, reflect.Array:
			cField.isSliceOrArray = true
			cField.sliceSubtype = cField.typ.Elem()
			cField.isTimeSubtype = cField.sliceSubtype == reflect.TypeOf(time.Time{})
			cField.sliceSubKind = cField.sliceSubtype.Kind()
		case reflect.Map:
			cField.isMap = true
			cField.mapSubtype = cField.typ.Elem()
			cField.isTimeSubtype = cField.mapSubtype == reflect.TypeOf(time.Time{})
			cField.mapSubKind = cField.mapSubtype.Kind()
		}
	} else {
		cField = cacheField
	}

	switch cField.kind {

	case reflect.Struct, reflect.Interface, reflect.Invalid:

		if cField.typ != reflect.TypeOf(time.Time{}) {
			panic("Invalid field passed to ValidateFieldWithTag")
		}
	}

	if len(cField.tags) == 0 {

		if isSingleField {
			cField.tags, isCached = fieldsCache.Get(tag)
		}

		if !isCached {

			for k, t := range strings.Split(tag, tagSeparator) {

				if t == diveTag {

					if k == 0 {
						cField.diveTag = tag[4:]
					} else {
						cField.diveTag = strings.SplitN(tag, diveSplit, 2)[1][1:]
					}

					break
				}

				orVals := strings.Split(t, orSeparator)
				cTag := &cachedTags{isOrVal: len(orVals) > 1, keyVals: make([][]string, len(orVals))}

				// if isInDive {

				// s, ok := cField.DiveTags[cField.DiveMaxDepth]

				// if ok {
				// 	cField.DiveTags[cField.DiveMaxDepth] = cField.DiveTags[cField.DiveMaxDepth] + tagSeparator + tag
				// } else {
				// cField.DiveTags[cField.DiveMaxDepth] = tag
				// }

				// continue

				// } else {
				cField.tags = append(cField.tags, cTag)
				// }

				for i, val := range orVals {
					vals := strings.SplitN(val, tagKeySeparator, 2)

					key := strings.TrimSpace(vals[0])

					if len(key) == 0 {
						panic(fmt.Sprintf("Invalid validation tag on field %s", name))
					}

					param := ""
					if len(vals) > 1 {
						param = strings.Replace(vals[1], utf8HexComma, ",", -1)
					}

					cTag.keyVals[i] = []string{key, param}
				}
			}

			if isSingleField {
				fieldsCache.Set(cField.tag, cField.tags)
			}
		}
	}

	var fieldErr *FieldError
	var err error

	for _, cTag := range cField.tags {

		if cTag.isOrVal {

			errTag := ""

			for _, val := range cTag.keyVals {

				fieldErr, err = v.fieldWithNameAndSingleTag(val, current, f, val[0], val[1], name)

				if err == nil {
					return nil
				}

				errTag += orSeparator + fieldErr.Tag
			}

			errTag = strings.TrimLeft(errTag, orSeparator)

			fieldErr.Tag = errTag
			fieldErr.Kind = cField.kind
			fieldErr.Type = cField.typ

			return fieldErr
		}

		if fieldErr, err = v.fieldWithNameAndSingleTag(val, current, f, cTag.keyVals[0][0], cTag.keyVals[0][1], name); err != nil {

			fieldErr.Kind = cField.kind
			fieldErr.Type = cField.typ

			return fieldErr
		}
	}

	if len(cField.diveTag) > 0 {

		if cField.isSliceOrArray {

			if errs := v.traverseSliceOrArray(val, current, valueField, cField); errs != nil && len(errs) > 0 {

				return &FieldError{
					Field:            cField.name,
					Kind:             cField.kind,
					Type:             cField.typ,
					Value:            f,
					isPlaceholderErr: true,
					IsSliceOrArray:   true,
					// Index:            i,
					SliceOrArrayErrs: errs,
				}
			}
			// return if error here
		} else if cField.isMap {
			// return if error here
		} else {
			// throw error, if not a slice or map then should not have gotten here
		}

		// dive tags need to be passed to traverse
		// traverse needs to call a SliceOrArray recursive function to meet depth requirements

		// for depth, diveTag := range cField.DiveTags {

		// 	// error returned should be added to SliceOrArrayErrs
		// 	if errs := v.traverseSliceOrArrayField(val, current, depth, currentDepth+1, diveTag, cField, valueField); len(errs) > 0 {
		// 		// result := &FieldError{
		// 		// 	Field:            cField.name,
		// 		// 	Kind:             cField.kind,
		// 		// 	Type:             cField.typ,
		// 		// 	Value:            valueField.Index(i).Interface(),
		// 		// 	isPlaceholderErr: true,
		// 		// 	IsSliceOrArray:true,
		// 		// 	Index:i,
		// 		// 	SliceOrArrayErrs:
		// 		// }
		// 	}

		// 	for _, tag := range diveTag {

		// 		fmt.Println("Depth:", depth, " Tag:", tag, " SliceType:", cField.SliceSubtype, " MapType:", cField.MapSubtype, " Kind:", cField.kind)

		// 	}
		// }
	}

	return nil
}

func (v *Validate) traverseSliceOrArray(val interface{}, current interface{}, valueField reflect.Value, cField *cachedField) []error {

	errs := make([]error, 0)

	for i := 0; i < valueField.Len(); i++ {

		idxField := valueField.Index(i)

		switch cField.sliceSubKind {
		case reflect.Struct, reflect.Interface:

			if cField.isTimeSubtype || idxField.Type() == reflect.TypeOf(time.Time{}) {

				if fieldError := v.fieldWithNameAndValue(val, current, idxField.Interface(), cField.diveTag, cField.name, true, nil); fieldError != nil {
					errs = append(errs, fieldError)
				}

				continue
			}

			if idxField.Kind() == reflect.Ptr && idxField.IsNil() {

				if strings.Contains(cField.tag, omitempty) {
					continue
				}

				if strings.Contains(cField.tag, required) {

					errs = append(errs, &FieldError{
						Field: cField.name,
						Tag:   required,
						Value: idxField.Interface(),
						Kind:  reflect.Ptr,
						Type:  cField.sliceSubtype,
					})

					continue
				}
			}

			if structErrors := v.structRecursive(val, current, idxField.Interface()); structErrors != nil {
				errs = append(errs, structErrors)
			}

		default:
			if fieldError := v.fieldWithNameAndValue(val, current, idxField.Interface(), cField.diveTag, cField.name, true, nil); fieldError != nil {
				errs = append(errs, fieldError)
			}
		}
	}

	return errs
}

// func (v *Validate) traverseSliceOrArrayField(val interface{}, current interface{}, depth uint64, currentDepth uint64, diveTags []*cachedTags, cField *cachedField, valueField reflect.Value) []error {

// 	for i := 0; i < valueField.Len(); i++ {

// 		if depth != currentDepth {

// 			switch cField.SliceSubKind {
// 			case reflect.Slice, reflect.Array:
// 				return v.fieldWithNameAndValue(val, current, valueField.Index(i).Interface(), cField.tag, cField.name, false, cField, currentDepth)

// 				// type FieldError struct {
// 				// 	Field            string
// 				// 	Tag              string
// 				// 	Kind             reflect.Kind
// 				// 	Type             reflect.Type
// 				// 	Param            string
// 				// 	Value            interface{}
// 				// 	HasErr           bool
// 				// 	IsSliceOrArray   bool
// 				// 	IsMap            bool
// 				// 	Key              interface{}
// 				// 	Index            uint64
// 				// 	SliceOrArrayErrs []*error               // counld be FieldError, StructErrors
// 				// 	MapErrs          map[interface{}]*error // counld be FieldError, StructErrors
// 				// }

// 				// result := &FieldError{
// 				// 	Field:            cField.name,
// 				// 	Kind:             cField.kind,
// 				// 	Type:             cField.typ,
// 				// 	Value:            valueField.Index(i).Interface(),
// 				// 	isPlaceholderErr: true,
// 				// 	IsSliceOrArray:true,
// 				// 	Index:i,
// 				// 	SliceOrArrayErrs:
// 				// }
// 				// validationErrors.Errors[fieldError.Field] = fieldError
// 				// // free up memory reference
// 				// fieldError = nil
// 				// }
// 			default:
// 				panic("attempting to dive deeper, but Kind is not a Slice nor Array")
// 			}
// 		}

// 		// switch cField.SliceSubKind {
// 		// case reflect.Struct, reflect.Interface:
// 		//  // need to check if required tag and or omitempty just like in struct recirsive
// 		// 	if cField.isTimeSubtype || valueField.Type() == reflect.TypeOf(time.Time{}) {

// 		// 		if fieldError := v.fieldWithNameAndValue(top, current, valueField.Index(i).Interface(), cField.tag, cField.name, false, cField); fieldError != nil {
// 		// 			validationErrors.Errors[fieldError.Field] = fieldError
// 		// 			// free up memory reference
// 		// 			fieldError = nil
// 		// 		}
// 		// 	}
// 		// }
// 		fmt.Println(valueField.Index(i))
// 	}
// 	// fmt.Println(v)
// 	// for _, item := range arr {

// 	// }
// 	return nil
// }

func (v *Validate) fieldWithNameAndSingleTag(val interface{}, current interface{}, f interface{}, key string, param string, name string) (*FieldError, error) {

	// OK to continue because we checked it's existance before getting into this loop
	if key == omitempty {
		return nil, nil
	}

	valFunc, ok := v.validationFuncs[key]
	if !ok {
		panic(fmt.Sprintf("Undefined validation function on field %s", name))
	}

	if err := valFunc(val, current, f, param); err {
		return nil, nil
	}

	return &FieldError{
		Field: name,
		Tag:   key,
		Value: f,
		Param: param,
	}, errors.New(key)
}
