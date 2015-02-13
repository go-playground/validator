package validator

import "reflect"

var bakedInValidators = map[string]ValidationFunc{
	"required": isRequired,
}

func isRequired(field interface{}, param string) bool {

	return field != nil && field != reflect.Zero(reflect.TypeOf(field)).Interface()
}
