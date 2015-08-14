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

	switch kind {
	case reflect.Ptr, reflect.Interface, reflect.Invalid:

		if len(namespace) == 0 {
			return current, kind, true
		}

		return current, kind, false

	case reflect.Struct:

		typ := current.Type()
		fld := namespace

		if typ != timeType && typ != timePtrType {

			idx := strings.Index(namespace, namespaceSeparator)

			if idx != -1 {
				fld = namespace[:idx]
			}

			ns := namespace[idx+1:]

			bracketIdx := strings.Index(fld, leftBracket)
			if bracketIdx != -1 {
				fld = fld[:bracketIdx]

				if idx == -1 {
					ns = namespace[bracketIdx:]
				} else {
					ns = namespace[bracketIdx:]
				}
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

		return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(namespace[idx:idx2])), namespace[endIdx+1:])
	}

	return current, kind, true
}
