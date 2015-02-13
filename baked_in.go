package validator

import (
	"log"
	"reflect"
)

var bakedInValidators = map[string]ValidationFunc{
	"required": isRequired,
}

func isRequired(field interface{}, param string) bool {

	log.Printf("Required:%s Valid:%t\n", field, field != nil && field != reflect.Zero(reflect.TypeOf(field)).Interface())
	return field != nil && field != reflect.Zero(reflect.TypeOf(field)).Interface()

	// return true
}
