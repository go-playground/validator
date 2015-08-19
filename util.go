package validator

import (
	"reflect"
	"strconv"
)

const (
	namespaceSeparator = "."
	leftBracket        = "["
	rightBracket       = "]"
)

func (v *Validate) extractType(current reflect.Value) (reflect.Value, reflect.Kind) {

	switch current.Kind() {
	case reflect.Ptr:

		if current.IsNil() {
			return current, reflect.Ptr
		}

		return v.extractType(current.Elem())

	case reflect.Interface:

		if current.IsNil() {
			return current, reflect.Interface
		}

		return v.extractType(current.Elem())

	case reflect.Invalid:
		return current, reflect.Invalid

	default:

		if v.config.hasCustomFuncs {
			if fn, ok := v.config.CustomTypeFuncs[current.Type()]; ok {
				return v.extractType(reflect.ValueOf(fn(current)))
			}
		}

		return current, current.Kind()
	}
}

// asInt retuns the parameter as a int64
// or panics if it can't convert
func asInt(param string) int64 {

	i, err := strconv.ParseInt(param, 0, 64)
	panicIf(err)

	return i
}

// asUint returns the parameter as a uint64
// or panics if it can't convert
func asUint(param string) uint64 {

	i, err := strconv.ParseUint(param, 0, 64)
	panicIf(err)

	return i
}

// asFloat returns the parameter as a float64
// or panics if it can't convert
func asFloat(param string) float64 {

	i, err := strconv.ParseFloat(param, 64)
	panicIf(err)

	return i
}

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}
