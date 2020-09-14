package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pantsmann/validator"
	//"gopkg.in/go-playground/validator.v9"
)

// MyStruct ..
type MyStruct struct {
	Title string
	A     int `validate:"group=sum"`
	B     int `validate:"group=sum"`
	C     int `validate:"gte-sum=sum 10"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterValidation("gte-sum", ValidateFieldSumGreaterThanEqual, validator.VFlagCoDependentErr)

	s := MyStruct{
		Title: "Only 5",
		A:     2,
		C:     3,
	}

	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	s.Title = "Eleven"
	s.B = 6
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateFieldSumGreaterThanEqual implements validator.Func
func ValidateFieldSumGreaterThanEqual(fl validator.FieldLevel) bool {
	vals := strings.Split(fl.Param(), " ")
	if len(vals) != 2 {
		panic("The 'gte-sum' validation tag takes two parameters: GroupName and Number")
	}
	limit, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err.Error())
	}

	validator.CoDependentGroups.AddGroupField(vals[0], fl).CloseGroup(vals[0])

	var sum int64
	for _, f := range validator.CoDependentGroups.Fields(vals[0]) {
		sum += f.Field().Int()
	}

	return sum >= int64(limit)
}
