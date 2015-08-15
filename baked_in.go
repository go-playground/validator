package validator

import (
	"fmt"
	"net"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// BakedInValidators is the default map of ValidationFunc
// you can add, remove or even replace items to suite your needs,
// or even disregard and use your own map if so desired.
var BakedInValidators = map[string]Func{
	"required":     hasValue,
	"len":          hasLengthOf,
	"min":          hasMinOf,
	"max":          hasMaxOf,
	"eq":           isEq,
	"ne":           isNe,
	"lt":           isLt,
	"lte":          isLte,
	"gt":           isGt,
	"gte":          isGte,
	"eqfield":      isEqField,
	"eqcsfield":    isEqCrossStructField,
	"nefield":      isNeField,
	"gtefield":     isGteField,
	"gtfield":      isGtField,
	"ltefield":     isLteField,
	"ltfield":      isLtField,
	"alpha":        isAlpha,
	"alphanum":     isAlphanum,
	"numeric":      isNumeric,
	"number":       isNumber,
	"hexadecimal":  isHexadecimal,
	"hexcolor":     isHexcolor,
	"rgb":          isRgb,
	"rgba":         isRgba,
	"hsl":          isHsl,
	"hsla":         isHsla,
	"email":        isEmail,
	"url":          isURL,
	"uri":          isURI,
	"base64":       isBase64,
	"contains":     contains,
	"containsany":  containsAny,
	"containsrune": containsRune,
	"excludes":     excludes,
	"excludesall":  excludesAll,
	"excludesrune": excludesRune,
	"isbn":         isISBN,
	"isbn10":       isISBN10,
	"isbn13":       isISBN13,
	"uuid":         isUUID,
	"uuid3":        isUUID3,
	"uuid4":        isUUID4,
	"uuid5":        isUUID5,
	"ascii":        isASCII,
	"printascii":   isPrintableASCII,
	"multibyte":    hasMultiByteCharacter,
	"datauri":      isDataURI,
	"latitude":     isLatitude,
	"longitude":    isLongitude,
	"ssn":          isSSN,
	"ipv4":         isIPv4,
	"ipv6":         isIPv6,
	"ip":           isIP,
	"mac":          isMac,
}

func isMac(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	_, err := net.ParseMAC(field.String())
	return err == nil
}

func isIPv4(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	ip := net.ParseIP(field.String())

	return ip != nil && ip.To4() != nil
}

func isIPv6(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	ip := net.ParseIP(field.String())

	return ip != nil && ip.To4() == nil
}

func isIP(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	ip := net.ParseIP(field.String())

	return ip != nil
}

func isSSN(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if field.Len() != 11 {
		return false
	}

	return matchesRegex(sSNRegex, field.String())
}

func isLongitude(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(longitudeRegex, field.String())
}

func isLatitude(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(latitudeRegex, field.String())
}

func isDataURI(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	uri := strings.SplitN(field.String(), ",", 2)

	if len(uri) != 2 {
		return false
	}

	if !matchesRegex(dataURIRegex, uri[0]) {
		return false
	}

	fld := reflect.ValueOf(uri[1])

	return isBase64(v, topStruct, currentStruct, fld, fld.Type(), fld.Kind(), param)
}

func hasMultiByteCharacter(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if field.Len() == 0 {
		return true
	}

	return matchesRegex(multibyteRegex, field.String())
}

func isPrintableASCII(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(printableASCIIRegex, field.String())
}

func isASCII(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(aSCIIRegex, field.String())
}

func isUUID5(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(uUID5Regex, field.String())
}

func isUUID4(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(uUID4Regex, field.String())
}

func isUUID3(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(uUID3Regex, field.String())
}

func isUUID(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(uUIDRegex, field.String())
}

func isISBN(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return isISBN10(v, topStruct, currentStruct, field, fieldType, fieldKind, param) || isISBN13(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
}

func isISBN13(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	s := strings.Replace(strings.Replace(field.String(), "-", "", 4), " ", "", 4)

	if !matchesRegex(iSBN13Regex, s) {
		return false
	}

	var checksum int32
	var i int32

	factor := []int32{1, 3}

	for i = 0; i < 12; i++ {
		checksum += factor[i%2] * int32(s[i]-'0')
	}

	if (int32(s[12]-'0'))-((10-(checksum%10))%10) == 0 {
		return true
	}

	return false
}

func isISBN10(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	s := strings.Replace(strings.Replace(field.String(), "-", "", 3), " ", "", 3)

	if !matchesRegex(iSBN10Regex, s) {
		return false
	}

	var checksum int32
	var i int32

	for i = 0; i < 9; i++ {
		checksum += (i + 1) * int32(s[i]-'0')
	}

	if s[9] == 'X' {
		checksum += 10 * 10
	} else {
		checksum += 10 * int32(s[9]-'0')
	}

	if checksum%11 == 0 {
		return true
	}

	return false
}

func excludesRune(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return !containsRune(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
}

func excludesAll(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return !containsAny(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
}

func excludes(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return !contains(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
}

func containsRune(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	r, _ := utf8.DecodeRuneInString(param)

	return strings.ContainsRune(field.String(), r)
}

func containsAny(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return strings.ContainsAny(field.String(), param)
}

func contains(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return strings.Contains(field.String(), param)
}

func isNeField(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return !isEqField(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
}

func isNe(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return !isEq(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
}

func isEqCrossStructField(v *Validate, topStruct reflect.Value, current reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	topField, topKind, ok := v.getStructFieldOK(topStruct, param)
	if !ok || topKind != fieldKind {
		return false
	}

	switch fieldKind {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return topField.Int() == field.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return topField.Uint() == field.Uint()

	case reflect.Float32, reflect.Float64:
		return topField.Float() == field.Float()

	case reflect.Slice, reflect.Map, reflect.Array:
		return int64(topField.Len()) == int64(field.Len())

	case reflect.Struct:

		// Not Same underlying type i.e. struct and time
		if fieldType != topField.Type() {
			return false
		}

		if fieldType == timeType {

			t := field.Interface().(time.Time)
			fieldTime := topField.Interface().(time.Time)

			return fieldTime.Equal(t)
		}
	}

	// default reflect.String:
	return topField.String() == current.String()
}

func isEqField(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	currentField, currentKind, ok := v.getStructFieldOK(currentStruct, param)
	if !ok || currentKind != fieldKind {
		return false
	}

	switch fieldKind {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() == currentField.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return field.Uint() == currentField.Uint()

	case reflect.Float32, reflect.Float64:
		return field.Float() == currentField.Float()

	case reflect.Slice, reflect.Map, reflect.Array:
		return int64(field.Len()) == int64(currentField.Len())

	case reflect.Struct:

		// Not Same underlying type i.e. struct and time
		if fieldType != currentField.Type() {
			return false
		}

		if fieldType == timeType {

			t := currentField.Interface().(time.Time)
			fieldTime := field.Interface().(time.Time)

			return fieldTime.Equal(t)
		}

	}

	// default reflect.String:
	return field.String() == currentField.String()
}

func isEq(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		return field.String() == param

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() == p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isBase64(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(base64Regex, field.String())
}

func isURI(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		_, err := url.ParseRequestURI(field.String())

		return err == nil
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isURL(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		url, err := url.ParseRequestURI(field.String())

		if err != nil {
			return false
		}

		if len(url.Scheme) == 0 {
			return false
		}

		return err == nil
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isEmail(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(emailRegex, field.String())
}

func isHsla(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(hslaRegex, field.String())
}

func isHsl(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(hslRegex, field.String())
}

func isRgba(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(rgbaRegex, field.String())
}

func isRgb(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(rgbRegex, field.String())
}

func isHexcolor(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(hexcolorRegex, field.String())
}

func isHexadecimal(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(hexadecimalRegex, field.String())
}

func isNumber(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(numberRegex, field.String())
}

func isNumeric(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(numericRegex, field.String())
}

func isAlphanum(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(alphaNumericRegex, field.String())
}

func isAlpha(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return matchesRegex(alphaRegex, field.String())
}

func hasValue(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(fieldType).Interface()
	}
}

func isGteField(v *Validate, topStruct reflect.Value, current reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if !current.IsValid() {
		panic("struct not passed for cross validation")
	}

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	switch current.Kind() {

	case reflect.Struct:

		if current.Type() == timeType || current.Type() == timePtrType {
			break
		}

		current = current.FieldByName(param)

		if current.Kind() == reflect.Invalid {
			panic(fmt.Sprintf("Field \"%s\" not found in struct", param))
		}
	}

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	switch fieldKind {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return field.Int() >= current.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:

		return field.Uint() >= current.Uint()

	case reflect.Float32, reflect.Float64:

		return field.Float() >= current.Float()

	case reflect.Struct:

		if field.Type() == timeType || field.Type() == timePtrType {

			if current.Type() != timeType && current.Type() != timePtrType {
				panic("Bad Top Level field type")
			}

			t := current.Interface().(time.Time)
			fieldTime := field.Interface().(time.Time)

			return fieldTime.After(t) || fieldTime.Equal(t)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// func isEqField(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

// 	currentField, currentKind, ok := v.getStructFieldOK(currentStruct, param)
// 	if !ok || currentKind != fieldKind {
// 		return false
// 	}

// 	switch fieldKind {

// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		return field.Int() == currentField.Int()

// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
// 		return field.Uint() == currentField.Uint()

// 	case reflect.Float32, reflect.Float64:
// 		return field.Float() == currentField.Float()

// 	case reflect.Slice, reflect.Map, reflect.Array:
// 		return int64(field.Len()) == int64(currentField.Len())

// 	case reflect.Struct:

// 		// Not Same underlying type i.e. struct and time
// 		if fieldType != currentField.Type() {
// 			return false
// 		}

// 		if fieldType == timeType {

// 			t := currentField.Interface().(time.Time)
// 			fieldTime := field.Interface().(time.Time)

// 			return fieldTime.Equal(t)
// 		}

// 	}

// 	// default reflect.String:
// 	return field.String() == currentField.String()
// }

func isGtField(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	currentField, currentKind, ok := v.getStructFieldOK(currentStruct, param)
	if !ok || currentKind != fieldKind {
		return false
	}

	switch fieldKind {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return field.Int() > currentField.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:

		return field.Uint() > currentField.Uint()

	case reflect.Float32, reflect.Float64:

		return field.Float() > currentField.Float()

	case reflect.Struct:

		// Not Same underlying type i.e. struct and time
		if fieldType != currentField.Type() {
			return false
		}

		if fieldType == timeType {

			t := currentField.Interface().(time.Time)
			fieldTime := field.Interface().(time.Time)

			return fieldTime.After(t)
		}
	}

	// default reflect.String
	return len(field.String()) > len(currentField.String())
}

func isGte(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) >= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) >= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() >= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() >= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() >= p

	case reflect.Struct:

		if fieldType == timeType || fieldType == timePtrType {

			now := time.Now().UTC()
			t := field.Interface().(time.Time)

			return t.After(now) || t.Equal(now)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isGt(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) > p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) > p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() > p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() > p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() > p
	case reflect.Struct:

		if field.Type() == timeType || field.Type() == timePtrType {

			return field.Interface().(time.Time).After(time.Now().UTC())
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// length tests whether a variable's length is equal to a given
// value. For strings it tests the number of characters whereas
// for maps and slices it tests the number of items.
func hasLengthOf(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) == p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() == p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// min tests whether a variable value is larger or equal to a given
// number. For number types, it's a simple lesser-than test; for
// strings it tests the number of characters whereas for maps
// and slices it tests the number of items.
func hasMinOf(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	return isGte(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
}

func isLteField(v *Validate, topStruct reflect.Value, current reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if !current.IsValid() {
		panic("struct not passed for cross validation")
	}

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	switch current.Kind() {

	case reflect.Struct:

		if current.Type() == timeType || current.Type() == timePtrType {
			break
		}

		current = current.FieldByName(param)

		if current.Kind() == reflect.Invalid {
			panic(fmt.Sprintf("Field \"%s\" not found in struct", param))
		}
	}

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	switch fieldKind {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return field.Int() <= current.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:

		return field.Uint() <= current.Uint()

	case reflect.Float32, reflect.Float64:

		return field.Float() <= current.Float()

	case reflect.Struct:

		if field.Type() == timeType || field.Type() == timePtrType {

			if current.Type() != timeType && current.Type() != timePtrType {
				panic("Bad Top Level field type")
			}

			t := current.Interface().(time.Time)
			fieldTime := field.Interface().(time.Time)

			return fieldTime.Before(t) || fieldTime.Equal(t)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isLtField(v *Validate, topStruct reflect.Value, current reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if !current.IsValid() {
		panic("struct not passed for cross validation")
	}

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	switch current.Kind() {

	case reflect.Struct:

		if current.Type() == timeType || current.Type() == timePtrType {
			break
		}

		current = current.FieldByName(param)

		if current.Kind() == reflect.Invalid {
			panic(fmt.Sprintf("Field \"%s\" not found in struct", param))
		}
	}

	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	switch fieldKind {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return field.Int() < current.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:

		return field.Uint() < current.Uint()

	case reflect.Float32, reflect.Float64:

		return field.Float() < current.Float()

	case reflect.Struct:

		if field.Type() == timeType || field.Type() == timePtrType {

			if current.Type() != timeType && current.Type() != timePtrType {
				panic("Bad Top Level field type")
			}

			t := current.Interface().(time.Time)
			fieldTime := field.Interface().(time.Time)

			return fieldTime.Before(t)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isLte(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) <= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) <= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() <= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() <= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() <= p

	case reflect.Struct:

		if fieldType == timeType || fieldType == timePtrType {

			now := time.Now().UTC()
			t := field.Interface().(time.Time)

			return t.Before(now) || t.Equal(now)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isLt(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	switch fieldKind {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) < p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) < p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() < p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() < p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() < p

	case reflect.Struct:

		if field.Type() == timeType || field.Type() == timePtrType {

			return field.Interface().(time.Time).Before(time.Now().UTC())
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// max tests whether a variable value is lesser than a given
// value. For numbers, it's a simple lesser-than test; for
// strings it tests the number of characters whereas for maps
// and slices it tests the number of items.
func hasMaxOf(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	return isLte(v, topStruct, currentStruct, field, fieldType, fieldKind, param)
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
