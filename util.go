package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	blank              = ""
	namespaceSeparator = "."
	leftBracket        = "["
	rightBracket       = "]"
	restrictedTagChars = ".[],|=+()`~!@#$%^&*\\\"/?<>{}"
	restrictedAliasErr = "Alias \"%s\" either contains restricted characters or is the same as a restricted tag needed for normal operation"
	restrictedTagErr   = "Tag \"%s\" either contains restricted characters or is the same as a restricted tag needed for normal operation"
)

var (
	restrictedTags = map[string]*struct{}{
		diveTag:           emptyStructPtr,
		existsTag:         emptyStructPtr,
		structOnlyTag:     emptyStructPtr,
		omitempty:         emptyStructPtr,
		skipValidationTag: emptyStructPtr,
		utf8HexComma:      emptyStructPtr,
		utf8Pipe:          emptyStructPtr,
	}
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

		if v.hasCustomFuncs {
			if fn, ok := v.customTypeFuncs[current.Type()]; ok {
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
				ns = blank
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

func (v *Validate) parseTags(tag, fieldName string) *cachedTag {

	cTag := &cachedTag{}

	v.parseTagsRecursive(cTag, tag, fieldName, blank, false)
	return cTag
}

func (v *Validate) parseTagsRecursive(cTag *cachedTag, tag, fieldName, alias string, isAlias bool) bool {

	if len(tag) == 0 {
		return true
	}

	for _, t := range strings.Split(tag, tagSeparator) {

		if v.hasAliasValidators {
			// check map for alias and process new tags, otherwise process as usual
			if tagsVal, ok := v.aliasValidators[t]; ok {

				leave := v.parseTagsRecursive(cTag, tagsVal, fieldName, t, true)

				if leave {
					return leave
				}

				continue
			}
		}

		if t == diveTag {
			cTag.diveTag = tag
			tVals := &tagVals{tagVals: [][]string{{t}}}
			cTag.tags = append(cTag.tags, tVals)
			return true
		}

		if t == omitempty {
			cTag.isOmitEmpty = true
		}

		// if a pipe character is needed within the param you must use the utf8Pipe representation "0x7C"
		orVals := strings.Split(t, orSeparator)
		tagVal := &tagVals{isAlias: isAlias, isOrVal: len(orVals) > 1, tagVals: make([][]string, len(orVals))}
		cTag.tags = append(cTag.tags, tagVal)

		var key string
		var param string

		for i, val := range orVals {
			vals := strings.SplitN(val, tagKeySeparator, 2)
			key = vals[0]

			tagVal.tag = key

			if isAlias {
				tagVal.tag = alias
			}

			if len(key) == 0 {
				panic(strings.TrimSpace(fmt.Sprintf(invalidValidation, fieldName)))
			}

			if len(vals) > 1 {
				param = strings.Replace(strings.Replace(vals[1], utf8HexComma, ",", -1), utf8Pipe, "|", -1)
			}

			tagVal.tagVals[i] = []string{key, param}
		}
	}

	return false
}
