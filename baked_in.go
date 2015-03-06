package validator

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// BakedInValidators is the map of ValidationFunc used internally
// but can be used with any new Validator if desired
var BakedInValidators = map[string]ValidationFunc{
	"required":    hasValue,
	"len":         hasLengthOf,
	"min":         hasMinOf,
	"max":         hasMaxOf,
	"lt":          isLt,
	"lte":         isLte,
	"gt":          isGt,
	"gte":         isGte,
	"alpha":       isAlpha,
	"alphanum":    isAlphanum,
	"numeric":     isNumeric,
	"number":      isNumber,
	"hexadecimal": isHexadecimal,
	"hexcolor":    isHexcolor,
	"rgb":         isRgb,
	"rgba":        isRgba,
	"hsl":         isHsl,
	"hsla":        isHsla,
	"email":       isEmail,
	"url":         isURL,
	"uri":         isURI,
}

func isURI(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		_, err := url.ParseRequestURI(field.(string))

		return err == nil
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isURL(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		url, err := url.ParseRequestURI(field.(string))

		if err != nil {
			return false
		}

		if len(url.Scheme) == 0 {
			return false
		}

		return err == nil

	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isEmail(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return emailRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isHsla(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hslaRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isHsl(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hslRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isRgba(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return rgbaRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isRgb(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return rgbRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isHexcolor(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hexcolorRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isHexadecimal(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hexadecimalRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isNumber(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return numberRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isNumeric(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return numericRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isAlphanum(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return alphaNumericRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isAlpha(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return alphaRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func hasValue(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.Slice, reflect.Map, reflect.Array:
		return field != nil && int64(st.Len()) > 0

	default:
		return field != nil && field != reflect.Zero(reflect.TypeOf(field)).Interface()
	}
}

func isGte(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) >= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) >= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() >= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() >= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() >= p

	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isGt(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) > p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) > p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() > p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() > p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() > p

	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

// length tests whether a variable's length is equal to a given
// value. For strings it tests the number of characters whereas
// for maps and slices it tests the number of items.
func hasLengthOf(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) == p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() == p

	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

// min tests whether a variable value is larger or equal to a given
// number. For number types, it's a simple lesser-than test; for
// strings it tests the number of characters whereas for maps
// and slices it tests the number of items.
func hasMinOf(val interface{}, field interface{}, param string) bool {

	return isGte(val, field, param)
}

func isLte(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) <= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) <= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() <= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() <= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() <= p

	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func isLt(val interface{}, field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) < p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) < p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() < p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() < p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() < p

	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

// max tests whether a variable value is lesser than a given
// value. For numbers, it's a simple lesser-than test; for
// strings it tests the number of characters whereas for maps
// and slices it tests the number of items.
func hasMaxOf(val interface{}, field interface{}, param string) bool {

	return isLte(val, field, param)
}

// asInt retuns the parameter as a int64
// or panics if it can't convert
func asInt(param string) int64 {

	i, err := strconv.ParseInt(param, 0, 64)

	if err != nil {
		panic(err.Error())
	}

	return i
}

// asUint returns the parameter as a uint64
// or panics if it can't convert
func asUint(param string) uint64 {

	i, err := strconv.ParseUint(param, 0, 64)

	if err != nil {
		panic(err.Error())
	}

	return i
}

// asFloat returns the parameter as a float64
// or panics if it can't convert
func asFloat(param string) float64 {

	i, err := strconv.ParseFloat(param, 64)

	if err != nil {
		panic(err.Error())
	}

	return i
}
