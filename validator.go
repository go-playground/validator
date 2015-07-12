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
)

const (
	fieldErrMsg = "Key: \"%s\" Error:Field validation for \"%s\" failed on the \"%s\" tag"
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
	// Kind             reflect.Kind
	// Type             reflect.Type
	// Param            string
	// Value            interface{}
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

	v.structRecursive(sv, sv, "", errs)

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// func (v *Validate) structRecursive(top interface{}, current interface{}, s interface{}, errs map[string]*FieldError) {

func (v *Validate) structRecursive(top reflect.Value, current reflect.Value, errPrefix string, errs ValidationErrors) {

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	if current.Kind() != reflect.Struct && current.Kind() != reflect.Interface {
		panic("value passed for validation is not a struct")
	}

	// errs[errPrefix+"Name"] = &FieldError{Field: "Name", Tag: "required"}

	// if depth < 3 {
	// 	v.structRecursive(top, current, errs)
	// }
}

// // Struct validates a struct, even it's nested structs, and returns a struct containing the errors
// // NOTE: Nested Arrays, or Maps of structs do not get validated only the Array or Map itself; the reason is that there is no good
// // way to represent or report which struct within the array has the error, besides can validate the struct prior to adding it to
// // the Array or Map.
// func (v *Validate) Struct(s interface{}) map[string]*FieldError {

// 	// var err *FieldError
// 	errs := map[string]*FieldError{}
// 	errchan := make(chan *FieldError)
// 	done := make(chan bool)
// 	// wg := &sync.WaitGroup{}

// 	go v.structRecursive(s, s, s, 0, errchan, done)

// LOOP:
// 	for {
// 		select {
// 		case err := <-errchan:
// 			errs[err.Field] = err
// 			// fmt.Println(err)
// 		case <-done:
// 			// fmt.Println("All Done")
// 			break LOOP
// 		}
// 	}

// 	return errs
// }

// func (v *Validate) structRecursive(top interface{}, current interface{}, s interface{}, depth int, errs chan *FieldError, done chan bool) {

// 	errs <- &FieldError{Field: "Name"}

// 	if depth < 1 {
// 		// wg.Add(1)
// 		v.structRecursive(s, s, s, depth+1, errs, done)
// 	}

// 	// wg.Wait()

// 	if depth == 0 {
// 		// wg.Wait()
// 		done <- true
// 		// return
// 	} else {
// 		// wg.Done()
// 	}
// }
