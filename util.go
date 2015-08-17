package validator

import (
	"reflect"
	"strconv"
	"strings"
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

func (v *Validate) getStructFieldOK(current reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool) {

	current, kind := v.extractType(current)

	if kind == reflect.Invalid {
		return current, kind, false
	}

	if len(namespace) == 0 {
		return current, kind, true
	}

	switch kind {

	case reflect.Ptr, reflect.Interface:

		return current, kind, false

	case reflect.Struct:

		typ := current.Type()
		fld := namespace
		ns := namespace

		if typ != timeType && typ != timePtrType {

			idx := strings.Index(namespace, namespaceSeparator)

			if idx != -1 {
				fld = namespace[:idx]
				ns = namespace[idx+1:]
			} else {
				ns = ""
				idx = len(namespace)
			}

			bracketIdx := strings.Index(fld, leftBracket)
			if bracketIdx != -1 {
				fld = fld[:bracketIdx]

				ns = namespace[bracketIdx:]
			}

			current = current.FieldByName(fld)

			return v.getStructFieldOK(current, ns)
		}

	case reflect.Array, reflect.Slice:
		idx := strings.Index(namespace, leftBracket)
		idx2 := strings.Index(namespace, rightBracket)

		arrIdx, _ := strconv.Atoi(namespace[idx+1 : idx2])

		if arrIdx >= current.Len() {
			return current, kind, false
		}

		startIdx := idx2 + 1

		if startIdx < len(namespace) {
			if namespace[startIdx:startIdx+1] == namespaceSeparator {
				startIdx++
			}
		}

		return v.getStructFieldOK(current.Index(arrIdx), namespace[startIdx:])

	case reflect.Map:
		idx := strings.Index(namespace, leftBracket) + 1
		idx2 := strings.Index(namespace, rightBracket)

		endIdx := idx2

		if endIdx+1 < len(namespace) {
			if namespace[endIdx+1:endIdx+2] == namespaceSeparator {
				endIdx++
			}
		}

		key := namespace[idx:idx2]

		switch current.Type().Key().Kind() {
		case reflect.Int:
			i, _ := strconv.Atoi(key)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(i)), namespace[endIdx+1:])
		case reflect.Int8:
			i, _ := strconv.ParseInt(key, 10, 8)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(int8(i))), namespace[endIdx+1:])
		case reflect.Int16:
			i, _ := strconv.ParseInt(key, 10, 16)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(int16(i))), namespace[endIdx+1:])
		case reflect.Int32:
			i, _ := strconv.ParseInt(key, 10, 32)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(int32(i))), namespace[endIdx+1:])
		case reflect.Int64:
			i, _ := strconv.ParseInt(key, 10, 64)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(i)), namespace[endIdx+1:])
		case reflect.Uint:
			i, _ := strconv.ParseUint(key, 10, 0)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(uint(i))), namespace[endIdx+1:])
		case reflect.Uint8:
			i, _ := strconv.ParseUint(key, 10, 8)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(uint8(i))), namespace[endIdx+1:])
		case reflect.Uint16:
			i, _ := strconv.ParseUint(key, 10, 16)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(uint16(i))), namespace[endIdx+1:])
		case reflect.Uint32:
			i, _ := strconv.ParseUint(key, 10, 32)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(uint32(i))), namespace[endIdx+1:])
		case reflect.Uint64:
			i, _ := strconv.ParseUint(key, 10, 64)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(i)), namespace[endIdx+1:])
		case reflect.Float32:
			f, _ := strconv.ParseFloat(key, 32)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(float32(f))), namespace[endIdx+1:])
		case reflect.Float64:
			f, _ := strconv.ParseFloat(key, 64)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(f)), namespace[endIdx+1:])
		case reflect.Bool:
			b, _ := strconv.ParseBool(key)
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(b)), namespace[endIdx+1:])

		// reflect.Type = string
		default:
			return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(key)), namespace[endIdx+1:])
		}
	}

	// if got here there was more namespace, cannot go any deeper
	panic("Invalid field namespace")
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
