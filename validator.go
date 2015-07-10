/**
 * Package validator
 *
 * MISC:
 * - anonymous structs - they don't have names so expect the Struct name within StructErrors to be blank
 *
 */

package validator

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

// FieldError contains a single field's validation error along
// with other properties that may be needed for error message creation
type FieldError struct {
	Field string
	// Tag              string
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
func (v *Validate) Struct(s interface{}) map[string]*FieldError {

	// var err *FieldError
	errs := map[string]*FieldError{}
	// errchan := make(chan *FieldError)
	// done := make(chan bool)
	// wg := &sync.WaitGroup{}

	v.structRecursive(s, s, s, 0, errs)

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

	return errs
}

func (v *Validate) structRecursive(top interface{}, current interface{}, s interface{}, depth int, errs map[string]*FieldError) {

	errs["Name"] = &FieldError{Field: "Name"}

	if depth < 3 {
		// wg.Add(1)
		v.structRecursive(s, s, s, depth+1, errs)
	}

	// wg.Wait()

	// if depth == 0 {
	// 	// wg.Wait()
	// 	done <- true
	// 	// return
	// } else {
	// 	// wg.Done()
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
