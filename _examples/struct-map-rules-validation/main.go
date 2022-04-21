package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Data struct {
	Name    string
	Email   string
	Details *Details
}

type Details struct {
	FamilyMembers *FamilyMembers
	Salary        string
}

type FamilyMembers struct {
	FatherName string
	MotherName string
}

type Data2 struct {
	Name string
	Age  uint32
}

var validate = validator.New()

func main() {
	validateStruct()
	// output
	// Key: 'Data2.Name' Error:Field validation for 'Name' failed on the 'min' tag
	// Key: 'Data2.Age' Error:Field validation for 'Age' failed on the 'max' tag

	validateStructNested()
	// output
	// Key: 'Data.Name' Error:Field validation for 'Name' failed on the 'max' tag
	// Key: 'Data.Details.FamilyMembers' Error:Field validation for 'FamilyMembers' failed on the 'required' tag
}

func validateStruct() {
	data := Data2{
		Name: "leo",
		Age:  1000,
	}

	rules := map[string]string{
		"Name": "min=4,max=6",
		"Age":  "min=4,max=6",
	}

	validate.RegisterStructValidationMapRules(rules, Data2{})

	err := validate.Struct(data)
	fmt.Println(err)
	fmt.Println()
}

func validateStructNested() {
	data := Data{
		Name:  "11sdfddd111",
		Email: "zytel3301@mail.com",
		Details: &Details{
			Salary: "1000",
		},
	}

	rules1 := map[string]string{
		"Name":    "min=4,max=6",
		"Email":   "required,email",
		"Details": "required",
	}

	rules2 := map[string]string{
		"Salary":        "number",
		"FamilyMembers": "required",
	}

	rules3 := map[string]string{
		"FatherName": "required,min=4,max=32",
		"MotherName": "required,min=4,max=32",
	}

	validate.RegisterStructValidationMapRules(rules1, Data{})
	validate.RegisterStructValidationMapRules(rules2, Details{})
	validate.RegisterStructValidationMapRules(rules3, FamilyMembers{})
	err := validate.Struct(data)

	fmt.Println(err)
}
