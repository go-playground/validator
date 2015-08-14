package validator

import (
	"reflect"
	"strconv"
	"strings"
)

func (v *Validate) determineType(current reflect.Value) (reflect.Value, reflect.Kind) {

	switch current.Kind() {
	case reflect.Ptr:

		if current.IsNil() {
			return current, reflect.Ptr
		}

		return v.determineType(current.Elem())

	case reflect.Interface:

		if current.IsNil() {
			return current, reflect.Interface
		}

		return v.determineType(current.Elem())

	case reflect.Invalid:

		return current, reflect.Invalid

	default:

		// fmt.Println(current.Kind())
		if v.config.hasCustomFuncs {
			if fn, ok := v.config.CustomTypeFuncs[current.Type()]; ok {
				return v.determineType(reflect.ValueOf(fn(current)))
			}
		}

		return current, current.Kind()
	}
}

func (v *Validate) getStructFieldOK(current reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool) {

	// fmt.Println("NS:", namespace)

	current, kind := v.determineType(current)

	// fmt.Println("getStructFieldOK - ", current, kind)

	switch kind {
	case reflect.Ptr, reflect.Interface, reflect.Invalid:
		return current, kind, false

	case reflect.Struct:

		typ := current.Type()
		fld := namespace

		if typ != timeType && typ != timePtrType {

			idx := strings.Index(namespace, namespaceSeparator)

			// fmt.Println("IDX:", namespace, idx)
			if idx != -1 {
				fld = namespace[:idx]
			}

			ns := namespace[idx+1:]

			bracketIdx := strings.Index(fld, "[")
			if bracketIdx != -1 {
				fld = fld[:bracketIdx]
				// fmt.Println("NSS:", ns)

				if idx == -1 {
					ns = namespace[bracketIdx:]
				} else {
					ns = namespace[bracketIdx:]
				}
				// fmt.Println("NSS2:", ns)
			}

			// fmt.Println("Looking for field:", fld)
			current = current.FieldByName(fld)

			// fmt.Println("Current:", current)

			return v.getStructFieldOK(current, ns)
		}

	case reflect.Array, reflect.Slice:
		idx := strings.Index(namespace, "[")
		idx2 := strings.Index(namespace, "]")
		// idx3 := strings.Index(namespace, namespaceSeparator)

		// if idx3 == -1 {
		// 	idx3 = 0
		// } else {
		// 	idx3 = 1
		// }
		//

		arrIdx, _ := strconv.Atoi(namespace[idx+1 : idx2])

		// fmt.Println("ArrayIndex:", arrIdx)
		// fmt.Println("LEN:", current.Len())
		if arrIdx >= current.Len() {
			return current, kind, false
		}

		startIdx := idx2 + 1

		if startIdx < len(namespace) {
			if namespace[startIdx:startIdx+1] == "." {
				startIdx++
			}
		}

		return v.getStructFieldOK(current.Index(arrIdx), namespace[startIdx:])

	case reflect.Map:
		idx := strings.Index(namespace, "[") + 1
		idx2 := strings.Index(namespace, "]")

		endIdx := idx2

		// fmt.Println("END IDX:", endIdx)
		// fmt.Println("L NS:", len(namespace))
		// fmt.Println("NS:", namespace)

		if endIdx+1 < len(namespace) {
			if namespace[endIdx+1:endIdx+2] == "." {
				endIdx++
			}
		}

		// fmt.Println("KEY:", namespace[idx:idx2])
		// fmt.Println("KEY NS:", namespace[endIdx+1:])

		return v.getStructFieldOK(current.MapIndex(reflect.ValueOf(namespace[idx:idx2])), namespace[endIdx+1:])
	}

	// fmt.Println("Returning field")
	return current, kind, true
}
