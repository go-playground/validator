package validator

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	. "gopkg.in/go-playground/assert.v1"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/nl"
	ut "github.com/go-playground/universal-translator"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//
//
// go test -cpuprofile cpu.out
// ./validator.test -test.bench=. -test.cpuprofile=cpu.prof
// go tool pprof validator.test cpu.prof
//
//
// go test -memprofile mem.out

type I interface {
	Foo() string
}

type Impl struct {
	F string `validate:"len=3"`
}

func (i *Impl) Foo() string {
	return i.F
}

type SubTest struct {
	Test string `validate:"required"`
}

type TestInterface struct {
	Iface I
}

type TestString struct {
	BlankTag  string `validate:""`
	Required  string `validate:"required"`
	Len       string `validate:"len=10"`
	Min       string `validate:"min=1"`
	Max       string `validate:"max=10"`
	MinMax    string `validate:"min=1,max=10"`
	Lt        string `validate:"lt=10"`
	Lte       string `validate:"lte=10"`
	Gt        string `validate:"gt=10"`
	Gte       string `validate:"gte=10"`
	OmitEmpty string `validate:"omitempty,min=1,max=10"`
	Sub       *SubTest
	SubIgnore *SubTest `validate:"-"`
	Anonymous struct {
		A string `validate:"required"`
	}
	Iface I
}

type TestUint64 struct {
	Required  uint64 `validate:"required"`
	Len       uint64 `validate:"len=10"`
	Min       uint64 `validate:"min=1"`
	Max       uint64 `validate:"max=10"`
	MinMax    uint64 `validate:"min=1,max=10"`
	OmitEmpty uint64 `validate:"omitempty,min=1,max=10"`
}

type TestFloat64 struct {
	Required  float64 `validate:"required"`
	Len       float64 `validate:"len=10"`
	Min       float64 `validate:"min=1"`
	Max       float64 `validate:"max=10"`
	MinMax    float64 `validate:"min=1,max=10"`
	Lte       float64 `validate:"lte=10"`
	OmitEmpty float64 `validate:"omitempty,min=1,max=10"`
}

type TestSlice struct {
	Required  []int `validate:"required"`
	Len       []int `validate:"len=10"`
	Min       []int `validate:"min=1"`
	Max       []int `validate:"max=10"`
	MinMax    []int `validate:"min=1,max=10"`
	OmitEmpty []int `validate:"omitempty,min=1,max=10"`
}

func AssertError(t *testing.T, err error, nsKey, structNsKey, field, structField, expectedTag string) {

	errs := err.(ValidationErrors)

	found := false
	var fe FieldError

	for i := 0; i < len(errs); i++ {
		if errs[i].Namespace() == nsKey && errs[i].StructNamespace() == structNsKey {
			found = true
			fe = errs[i]
			break
		}
	}

	EqualSkip(t, 2, found, true)
	NotEqualSkip(t, 2, fe, nil)
	EqualSkip(t, 2, fe.Field(), field)
	EqualSkip(t, 2, fe.StructField(), structField)
	EqualSkip(t, 2, fe.Tag(), expectedTag)
}

func getError(err error, nsKey, structNsKey string) FieldError {

	errs := err.(ValidationErrors)

	var fe FieldError

	for i := 0; i < len(errs); i++ {
		if errs[i].Namespace() == nsKey && errs[i].StructNamespace() == structNsKey {
			fe = errs[i]
			break
		}
	}

	return fe
}

type valuer struct {
	Name string
}

func (v valuer) Value() (driver.Value, error) {

	if v.Name == "errorme" {
		panic("SQL Driver Valuer error: some kind of error")
		// return nil, errors.New("some kind of error")
	}

	if len(v.Name) == 0 {
		return nil, nil
	}

	return v.Name, nil
}

type MadeUpCustomType struct {
	FirstName string
	LastName  string
}

func ValidateCustomType(field reflect.Value) interface{} {

	if cust, ok := field.Interface().(MadeUpCustomType); ok {

		if len(cust.FirstName) == 0 || len(cust.LastName) == 0 {
			return ""
		}

		return cust.FirstName + " " + cust.LastName
	}

	return ""
}

func OverrideIntTypeForSomeReason(field reflect.Value) interface{} {

	if i, ok := field.Interface().(int); ok {
		if i == 1 {
			return "1"
		}

		if i == 2 {
			return "12"
		}
	}

	return ""
}

type CustomMadeUpStruct struct {
	MadeUp        MadeUpCustomType `validate:"required"`
	OverriddenInt int              `validate:"gt=1"`
}

func ValidateValuerType(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err != nil {
			// handle the error how you want
			return nil
		}

		return val
	}

	return nil
}

type TestPartial struct {
	NoTag     string
	BlankTag  string     `validate:""`
	Required  string     `validate:"required"`
	SubSlice  []*SubTest `validate:"required,dive"`
	Sub       *SubTest
	SubIgnore *SubTest `validate:"-"`
	Anonymous struct {
		A         string     `validate:"required"`
		ASubSlice []*SubTest `validate:"required,dive"`

		SubAnonStruct []struct {
			Test      string `validate:"required"`
			OtherTest string `validate:"required"`
		} `validate:"required,dive"`
	}
}

type TestStruct struct {
	String string `validate:"required" json:"StringVal"`
}

func StructValidationTestStructSuccess(sl StructLevel) {

	st := sl.Current().Interface().(TestStruct)

	if st.String != "good value" {
		sl.ReportError(st.String, "StringVal", "String", "badvalueteststruct", "good value")
	}
}

func StructValidationTestStruct(sl StructLevel) {

	st := sl.Current().Interface().(TestStruct)

	if st.String != "bad value" {
		sl.ReportError(st.String, "StringVal", "String", "badvalueteststruct", "bad value")
	}
}

func StructValidationNoTestStructCustomName(sl StructLevel) {

	st := sl.Current().Interface().(TestStruct)

	if st.String != "bad value" {
		sl.ReportError(st.String, "String", "", "badvalueteststruct", "bad value")
	}
}

func StructValidationTestStructInvalid(sl StructLevel) {

	st := sl.Current().Interface().(TestStruct)

	if st.String != "bad value" {
		sl.ReportError(nil, "StringVal", "String", "badvalueteststruct", "bad value")
	}
}

func StructValidationTestStructReturnValidationErrors(sl StructLevel) {

	s := sl.Current().Interface().(TestStructReturnValidationErrors)

	errs := sl.Validator().Struct(s.Inner1.Inner2)
	if errs == nil {
		return
	}

	sl.ReportValidationErrors("Inner1.", "Inner1.", errs.(ValidationErrors))
}

func StructValidationTestStructReturnValidationErrors2(sl StructLevel) {

	s := sl.Current().Interface().(TestStructReturnValidationErrors)

	errs := sl.Validator().Struct(s.Inner1.Inner2)
	if errs == nil {
		return
	}

	sl.ReportValidationErrors("Inner1JSON.", "Inner1.", errs.(ValidationErrors))
}

type TestStructReturnValidationErrorsInner2 struct {
	String string `validate:"required" json:"JSONString"`
}

type TestStructReturnValidationErrorsInner1 struct {
	Inner2 *TestStructReturnValidationErrorsInner2
}

type TestStructReturnValidationErrors struct {
	Inner1 *TestStructReturnValidationErrorsInner1 `json:"Inner1JSON"`
}

type StructLevelInvalidErr struct {
	Value string
}

func StructLevelInvalidError(sl StructLevel) {

	top := sl.Top().Interface().(StructLevelInvalidErr)
	s := sl.Current().Interface().(StructLevelInvalidErr)

	if top.Value == s.Value {
		sl.ReportError(nil, "Value", "Value", "required", "")
	}
}

func TestStructLevelInvalidError(t *testing.T) {

	validate := New()
	validate.RegisterStructValidation(StructLevelInvalidError, StructLevelInvalidErr{})

	var test StructLevelInvalidErr

	err := validate.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(ValidationErrors)
	Equal(t, ok, true)

	fe := errs[0]
	Equal(t, fe.Field(), "Value")
	Equal(t, fe.StructField(), "Value")
	Equal(t, fe.Namespace(), "StructLevelInvalidErr.Value")
	Equal(t, fe.StructNamespace(), "StructLevelInvalidErr.Value")
	Equal(t, fe.Tag(), "required")
	Equal(t, fe.ActualTag(), "required")
	Equal(t, fe.Kind(), reflect.Invalid)
	Equal(t, fe.Type(), reflect.TypeOf(nil))
}

func TestNameNamespace(t *testing.T) {

	type Inner2Namespace struct {
		String []string `validate:"dive,required" json:"JSONString"`
	}

	type Inner1Namespace struct {
		Inner2 *Inner2Namespace `json:"Inner2JSON"`
	}

	type Namespace struct {
		Inner1 *Inner1Namespace `json:"Inner1JSON"`
	}

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	i2 := &Inner2Namespace{String: []string{"ok", "ok", "ok"}}
	i1 := &Inner1Namespace{Inner2: i2}
	ns := &Namespace{Inner1: i1}

	errs := validate.Struct(ns)
	Equal(t, errs, nil)

	i2.String[1] = ""

	errs = validate.Struct(ns)
	NotEqual(t, errs, nil)

	ve := errs.(ValidationErrors)
	Equal(t, len(ve), 1)
	AssertError(t, errs, "Namespace.Inner1JSON.Inner2JSON.JSONString[1]", "Namespace.Inner1.Inner2.String[1]", "JSONString[1]", "String[1]", "required")

	fe := getError(ve, "Namespace.Inner1JSON.Inner2JSON.JSONString[1]", "Namespace.Inner1.Inner2.String[1]")
	NotEqual(t, fe, nil)
	Equal(t, fe.Field(), "JSONString[1]")
	Equal(t, fe.StructField(), "String[1]")
	Equal(t, fe.Namespace(), "Namespace.Inner1JSON.Inner2JSON.JSONString[1]")
	Equal(t, fe.StructNamespace(), "Namespace.Inner1.Inner2.String[1]")
}

func TestAnonymous(t *testing.T) {

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type Test struct {
		Anonymous struct {
			A string `validate:"required" json:"EH"`
		}
		AnonymousB struct {
			B string `validate:"required" json:"BEE"`
		}
		anonymousC struct {
			c string `validate:"required"`
		}
	}

	tst := &Test{
		Anonymous: struct {
			A string `validate:"required" json:"EH"`
		}{
			A: "1",
		},
		AnonymousB: struct {
			B string `validate:"required" json:"BEE"`
		}{
			B: "",
		},
		anonymousC: struct {
			c string `validate:"required"`
		}{
			c: "",
		},
	}

	err := validate.Struct(tst)
	NotEqual(t, err, nil)

	errs := err.(ValidationErrors)

	Equal(t, len(errs), 1)
	AssertError(t, errs, "Test.AnonymousB.BEE", "Test.AnonymousB.B", "BEE", "B", "required")

	fe := getError(errs, "Test.AnonymousB.BEE", "Test.AnonymousB.B")
	NotEqual(t, fe, nil)
	Equal(t, fe.Field(), "BEE")
	Equal(t, fe.StructField(), "B")

	s := struct {
		c string `validate:"required"`
	}{
		c: "",
	}

	err = validate.Struct(s)
	Equal(t, err, nil)
}

func TestAnonymousSameStructDifferentTags(t *testing.T) {

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type Test struct {
		A interface{}
	}

	tst := &Test{
		A: struct {
			A string `validate:"required"`
		}{
			A: "",
		},
	}

	err := validate.Struct(tst)
	NotEqual(t, err, nil)

	errs := err.(ValidationErrors)

	Equal(t, len(errs), 1)
	AssertError(t, errs, "Test.A.A", "Test.A.A", "A", "A", "required")

	tst = &Test{
		A: struct {
			A string `validate:"omitempty,required"`
		}{
			A: "",
		},
	}

	err = validate.Struct(tst)
	Equal(t, err, nil)
}

func TestStructLevelReturnValidationErrors(t *testing.T) {

	validate := New()
	validate.RegisterStructValidation(StructValidationTestStructReturnValidationErrors, TestStructReturnValidationErrors{})

	inner2 := &TestStructReturnValidationErrorsInner2{
		String: "I'm HERE",
	}

	inner1 := &TestStructReturnValidationErrorsInner1{
		Inner2: inner2,
	}

	val := &TestStructReturnValidationErrors{
		Inner1: inner1,
	}

	errs := validate.Struct(val)
	Equal(t, errs, nil)

	inner2.String = ""

	errs = validate.Struct(val)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 2)
	AssertError(t, errs, "TestStructReturnValidationErrors.Inner1.Inner2.String", "TestStructReturnValidationErrors.Inner1.Inner2.String", "String", "String", "required")
	// this is an extra error reported from struct validation
	AssertError(t, errs, "TestStructReturnValidationErrors.Inner1.TestStructReturnValidationErrorsInner2.String", "TestStructReturnValidationErrors.Inner1.TestStructReturnValidationErrorsInner2.String", "String", "String", "required")
}

func TestStructLevelReturnValidationErrorsWithJSON(t *testing.T) {

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	validate.RegisterStructValidation(StructValidationTestStructReturnValidationErrors2, TestStructReturnValidationErrors{})

	inner2 := &TestStructReturnValidationErrorsInner2{
		String: "I'm HERE",
	}

	inner1 := &TestStructReturnValidationErrorsInner1{
		Inner2: inner2,
	}

	val := &TestStructReturnValidationErrors{
		Inner1: inner1,
	}

	errs := validate.Struct(val)
	Equal(t, errs, nil)

	inner2.String = ""

	errs = validate.Struct(val)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 2)
	AssertError(t, errs, "TestStructReturnValidationErrors.Inner1JSON.Inner2.JSONString", "TestStructReturnValidationErrors.Inner1.Inner2.String", "JSONString", "String", "required")
	// this is an extra error reported from struct validation, it's a badly formatted one, but on purpose
	AssertError(t, errs, "TestStructReturnValidationErrors.Inner1JSON.TestStructReturnValidationErrorsInner2.JSONString", "TestStructReturnValidationErrors.Inner1.TestStructReturnValidationErrorsInner2.String", "JSONString", "String", "required")

	fe := getError(errs, "TestStructReturnValidationErrors.Inner1JSON.Inner2.JSONString", "TestStructReturnValidationErrors.Inner1.Inner2.String")
	NotEqual(t, fe, nil)

	// check for proper JSON namespace
	Equal(t, fe.Field(), "JSONString")
	Equal(t, fe.StructField(), "String")
	Equal(t, fe.Namespace(), "TestStructReturnValidationErrors.Inner1JSON.Inner2.JSONString")
	Equal(t, fe.StructNamespace(), "TestStructReturnValidationErrors.Inner1.Inner2.String")

	fe = getError(errs, "TestStructReturnValidationErrors.Inner1JSON.TestStructReturnValidationErrorsInner2.JSONString", "TestStructReturnValidationErrors.Inner1.TestStructReturnValidationErrorsInner2.String")
	NotEqual(t, fe, nil)

	// check for proper JSON namespace
	Equal(t, fe.Field(), "JSONString")
	Equal(t, fe.StructField(), "String")
	Equal(t, fe.Namespace(), "TestStructReturnValidationErrors.Inner1JSON.TestStructReturnValidationErrorsInner2.JSONString")
	Equal(t, fe.StructNamespace(), "TestStructReturnValidationErrors.Inner1.TestStructReturnValidationErrorsInner2.String")
}

func TestStructLevelValidations(t *testing.T) {

	v1 := New()
	v1.RegisterStructValidation(StructValidationTestStruct, TestStruct{})

	tst := &TestStruct{
		String: "good value",
	}

	errs := v1.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestStruct.StringVal", "TestStruct.String", "StringVal", "String", "badvalueteststruct")

	v2 := New()
	v2.RegisterStructValidation(StructValidationNoTestStructCustomName, TestStruct{})

	errs = v2.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestStruct.String", "TestStruct.String", "String", "String", "badvalueteststruct")

	v3 := New()
	v3.RegisterStructValidation(StructValidationTestStructInvalid, TestStruct{})

	errs = v3.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestStruct.StringVal", "TestStruct.String", "StringVal", "String", "badvalueteststruct")

	v4 := New()
	v4.RegisterStructValidation(StructValidationTestStructSuccess, TestStruct{})

	errs = v4.Struct(tst)
	Equal(t, errs, nil)
}

func TestAliasTags(t *testing.T) {

	validate := New()
	validate.RegisterAlias("iscoloralias", "hexcolor|rgb|rgba|hsl|hsla")

	s := "rgb(255,255,255)"
	errs := validate.Var(s, "iscoloralias")
	Equal(t, errs, nil)

	s = ""
	errs = validate.Var(s, "omitempty,iscoloralias")
	Equal(t, errs, nil)

	s = "rgb(255,255,0)"
	errs = validate.Var(s, "iscoloralias,len=5")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "len")

	type Test struct {
		Color string `validate:"iscoloralias"`
	}

	tst := &Test{
		Color: "#000",
	}

	errs = validate.Struct(tst)
	Equal(t, errs, nil)

	tst.Color = "cfvre"
	errs = validate.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Color", "Test.Color", "Color", "Color", "iscoloralias")

	fe := getError(errs, "Test.Color", "Test.Color")
	NotEqual(t, fe, nil)
	Equal(t, fe.ActualTag(), "hexcolor|rgb|rgba|hsl|hsla")

	validate.RegisterAlias("req", "required,dive,iscoloralias")
	arr := []string{"val1", "#fff", "#000"}

	errs = validate.Var(arr, "req")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "[0]", "[0]", "[0]", "[0]", "iscoloralias")

	PanicMatches(t, func() { validate.RegisterAlias("exists!", "gt=5,lt=10") }, "Alias 'exists!' either contains restricted characters or is the same as a restricted tag needed for normal operation")
}

func TestNilValidator(t *testing.T) {

	type TestStruct struct {
		Test string `validate:"required"`
	}

	ts := TestStruct{}

	var val *Validate

	fn := func(fl FieldLevel) bool {

		return fl.Parent().String() == fl.Field().String()
	}

	PanicMatches(t, func() { val.RegisterCustomTypeFunc(ValidateCustomType, MadeUpCustomType{}) }, "runtime error: invalid memory address or nil pointer dereference")
	PanicMatches(t, func() { val.RegisterValidation("something", fn) }, "runtime error: invalid memory address or nil pointer dereference")
	PanicMatches(t, func() { val.Var(ts.Test, "required") }, "runtime error: invalid memory address or nil pointer dereference")
	PanicMatches(t, func() { val.VarWithValue("test", ts.Test, "required") }, "runtime error: invalid memory address or nil pointer dereference")
	PanicMatches(t, func() { val.Struct(ts) }, "runtime error: invalid memory address or nil pointer dereference")
	PanicMatches(t, func() { val.StructExcept(ts, "Test") }, "runtime error: invalid memory address or nil pointer dereference")
	PanicMatches(t, func() { val.StructPartial(ts, "Test") }, "runtime error: invalid memory address or nil pointer dereference")
}

func TestStructPartial(t *testing.T) {

	p1 := []string{
		"NoTag",
		"Required",
	}

	p2 := []string{
		"SubSlice[0].Test",
		"Sub",
		"SubIgnore",
		"Anonymous.A",
	}

	p3 := []string{
		"SubTest.Test",
	}

	p4 := []string{
		"A",
	}

	tPartial := &TestPartial{
		NoTag:    "NoTag",
		Required: "Required",

		SubSlice: []*SubTest{
			{

				Test: "Required",
			},
			{

				Test: "Required",
			},
		},

		Sub: &SubTest{
			Test: "1",
		},
		SubIgnore: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A             string     `validate:"required"`
			ASubSlice     []*SubTest `validate:"required,dive"`
			SubAnonStruct []struct {
				Test      string `validate:"required"`
				OtherTest string `validate:"required"`
			} `validate:"required,dive"`
		}{
			A: "1",
			ASubSlice: []*SubTest{
				{
					Test: "Required",
				},
				{
					Test: "Required",
				},
			},

			SubAnonStruct: []struct {
				Test      string `validate:"required"`
				OtherTest string `validate:"required"`
			}{
				{"Required", "RequiredOther"},
				{"Required", "RequiredOther"},
			},
		},
	}

	validate := New()

	// the following should all return no errors as everything is valid in
	// the default state
	errs := validate.StructPartialCtx(context.Background(), tPartial, p1...)
	Equal(t, errs, nil)

	errs = validate.StructPartial(tPartial, p2...)
	Equal(t, errs, nil)

	// this isn't really a robust test, but is ment to illustrate the ANON CASE below
	errs = validate.StructPartial(tPartial.SubSlice[0], p3...)
	Equal(t, errs, nil)

	errs = validate.StructExceptCtx(context.Background(), tPartial, p1...)
	Equal(t, errs, nil)

	errs = validate.StructExcept(tPartial, p2...)
	Equal(t, errs, nil)

	// mod tParial for required feild and re-test making sure invalid fields are NOT required:
	tPartial.Required = ""

	errs = validate.StructExcept(tPartial, p1...)
	Equal(t, errs, nil)

	errs = validate.StructPartial(tPartial, p2...)
	Equal(t, errs, nil)

	// inversion and retesting Partial to generate failures:
	errs = validate.StructPartial(tPartial, p1...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.Required", "TestPartial.Required", "Required", "Required", "required")

	errs = validate.StructExcept(tPartial, p2...)
	AssertError(t, errs, "TestPartial.Required", "TestPartial.Required", "Required", "Required", "required")

	// reset Required field, and set nested struct
	tPartial.Required = "Required"
	tPartial.Anonymous.A = ""

	// will pass as unset feilds is not going to be tested
	errs = validate.StructPartial(tPartial, p1...)
	Equal(t, errs, nil)

	errs = validate.StructExcept(tPartial, p2...)
	Equal(t, errs, nil)

	// ANON CASE the response here is strange, it clearly does what it is being told to
	errs = validate.StructExcept(tPartial.Anonymous, p4...)
	Equal(t, errs, nil)

	// will fail as unset feild is tested
	errs = validate.StructPartial(tPartial, p2...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.Anonymous.A", "TestPartial.Anonymous.A", "A", "A", "required")

	errs = validate.StructExcept(tPartial, p1...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.Anonymous.A", "TestPartial.Anonymous.A", "A", "A", "required")

	// reset nested struct and unset struct in slice
	tPartial.Anonymous.A = "Required"
	tPartial.SubSlice[0].Test = ""

	// these will pass as unset item is NOT tested
	errs = validate.StructPartial(tPartial, p1...)
	Equal(t, errs, nil)

	errs = validate.StructExcept(tPartial, p2...)
	Equal(t, errs, nil)

	// these will fail as unset item IS tested
	errs = validate.StructExcept(tPartial, p1...)
	AssertError(t, errs, "TestPartial.SubSlice[0].Test", "TestPartial.SubSlice[0].Test", "Test", "Test", "required")
	Equal(t, len(errs.(ValidationErrors)), 1)

	errs = validate.StructPartial(tPartial, p2...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.SubSlice[0].Test", "TestPartial.SubSlice[0].Test", "Test", "Test", "required")
	Equal(t, len(errs.(ValidationErrors)), 1)

	// Unset second slice member concurrently to test dive behavior:
	tPartial.SubSlice[1].Test = ""

	errs = validate.StructPartial(tPartial, p1...)
	Equal(t, errs, nil)

	// NOTE: When specifying nested items, it is still the users responsibility
	// to specify the dive tag, the library does not override this.
	errs = validate.StructExcept(tPartial, p2...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.SubSlice[1].Test", "TestPartial.SubSlice[1].Test", "Test", "Test", "required")

	errs = validate.StructExcept(tPartial, p1...)
	Equal(t, len(errs.(ValidationErrors)), 2)
	AssertError(t, errs, "TestPartial.SubSlice[0].Test", "TestPartial.SubSlice[0].Test", "Test", "Test", "required")
	AssertError(t, errs, "TestPartial.SubSlice[1].Test", "TestPartial.SubSlice[1].Test", "Test", "Test", "required")

	errs = validate.StructPartial(tPartial, p2...)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "TestPartial.SubSlice[0].Test", "TestPartial.SubSlice[0].Test", "Test", "Test", "required")

	// reset struct in slice, and unset struct in slice in unset posistion
	tPartial.SubSlice[0].Test = "Required"

	// these will pass as the unset item is NOT tested
	errs = validate.StructPartial(tPartial, p1...)
	Equal(t, errs, nil)

	errs = validate.StructPartial(tPartial, p2...)
	Equal(t, errs, nil)

	// testing for missing item by exception, yes it dives and fails
	errs = validate.StructExcept(tPartial, p1...)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "TestPartial.SubSlice[1].Test", "TestPartial.SubSlice[1].Test", "Test", "Test", "required")

	errs = validate.StructExcept(tPartial, p2...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.SubSlice[1].Test", "TestPartial.SubSlice[1].Test", "Test", "Test", "required")

	tPartial.SubSlice[1].Test = "Required"

	tPartial.Anonymous.SubAnonStruct[0].Test = ""
	// these will pass as the unset item is NOT tested
	errs = validate.StructPartial(tPartial, p1...)
	Equal(t, errs, nil)

	errs = validate.StructPartial(tPartial, p2...)
	Equal(t, errs, nil)

	errs = validate.StructExcept(tPartial, p1...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.Anonymous.SubAnonStruct[0].Test", "TestPartial.Anonymous.SubAnonStruct[0].Test", "Test", "Test", "required")

	errs = validate.StructExcept(tPartial, p2...)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.Anonymous.SubAnonStruct[0].Test", "TestPartial.Anonymous.SubAnonStruct[0].Test", "Test", "Test", "required")

}

func TestCrossStructLteFieldValidation(t *testing.T) {

	type Inner struct {
		CreatedAt *time.Time
		String    string
		Int       int
		Uint      uint
		Float     float64
		Array     []string
	}

	type Test struct {
		Inner     *Inner
		CreatedAt *time.Time `validate:"ltecsfield=Inner.CreatedAt"`
		String    string     `validate:"ltecsfield=Inner.String"`
		Int       int        `validate:"ltecsfield=Inner.Int"`
		Uint      uint       `validate:"ltecsfield=Inner.Uint"`
		Float     float64    `validate:"ltecsfield=Inner.Float"`
		Array     []string   `validate:"ltecsfield=Inner.Array"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)

	inner := &Inner{
		CreatedAt: &then,
		String:    "abcd",
		Int:       13,
		Uint:      13,
		Float:     1.13,
		Array:     []string{"val1", "val2"},
	}

	test := &Test{
		Inner:     inner,
		CreatedAt: &now,
		String:    "abc",
		Int:       12,
		Uint:      12,
		Float:     1.12,
		Array:     []string{"val1"},
	}

	validate := New()
	errs := validate.Struct(test)
	Equal(t, errs, nil)

	test.CreatedAt = &then
	test.String = "abcd"
	test.Int = 13
	test.Uint = 13
	test.Float = 1.13
	test.Array = []string{"val1", "val2"}

	errs = validate.Struct(test)
	Equal(t, errs, nil)

	after := now.Add(time.Hour * 10)

	test.CreatedAt = &after
	test.String = "abce"
	test.Int = 14
	test.Uint = 14
	test.Float = 1.14
	test.Array = []string{"val1", "val2", "val3"}

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "ltecsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "ltecsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "ltecsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "ltecsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "ltecsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "ltecsfield")

	errs = validate.VarWithValueCtx(context.Background(), 1, "", "ltecsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltecsfield")

	// this test is for the WARNING about unforseen validation issues.
	errs = validate.VarWithValue(test, now, "ltecsfield")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 6)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "ltecsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "ltecsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "ltecsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "ltecsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "ltecsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "ltecsfield")

	type Other struct {
		Value string
	}

	type Test2 struct {
		Value Other
		Time  time.Time `validate:"ltecsfield=Value"`
	}

	tst := Test2{
		Value: Other{Value: "StringVal"},
		Time:  then,
	}

	errs = validate.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test2.Time", "Test2.Time", "Time", "Time", "ltecsfield")
}

func TestCrossStructLtFieldValidation(t *testing.T) {

	type Inner struct {
		CreatedAt *time.Time
		String    string
		Int       int
		Uint      uint
		Float     float64
		Array     []string
	}

	type Test struct {
		Inner     *Inner
		CreatedAt *time.Time `validate:"ltcsfield=Inner.CreatedAt"`
		String    string     `validate:"ltcsfield=Inner.String"`
		Int       int        `validate:"ltcsfield=Inner.Int"`
		Uint      uint       `validate:"ltcsfield=Inner.Uint"`
		Float     float64    `validate:"ltcsfield=Inner.Float"`
		Array     []string   `validate:"ltcsfield=Inner.Array"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)

	inner := &Inner{
		CreatedAt: &then,
		String:    "abcd",
		Int:       13,
		Uint:      13,
		Float:     1.13,
		Array:     []string{"val1", "val2"},
	}

	test := &Test{
		Inner:     inner,
		CreatedAt: &now,
		String:    "abc",
		Int:       12,
		Uint:      12,
		Float:     1.12,
		Array:     []string{"val1"},
	}

	validate := New()
	errs := validate.Struct(test)
	Equal(t, errs, nil)

	test.CreatedAt = &then
	test.String = "abcd"
	test.Int = 13
	test.Uint = 13
	test.Float = 1.13
	test.Array = []string{"val1", "val2"}

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "ltcsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "ltcsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "ltcsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "ltcsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "ltcsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "ltcsfield")

	errs = validate.VarWithValue(1, "", "ltcsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltcsfield")

	// this test is for the WARNING about unforseen validation issues.
	errs = validate.VarWithValue(test, now, "ltcsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "ltcsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "ltcsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "ltcsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "ltcsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "ltcsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "ltcsfield")

	type Other struct {
		Value string
	}

	type Test2 struct {
		Value Other
		Time  time.Time `validate:"ltcsfield=Value"`
	}

	tst := Test2{
		Value: Other{Value: "StringVal"},
		Time:  then,
	}

	errs = validate.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test2.Time", "Test2.Time", "Time", "Time", "ltcsfield")
}

func TestCrossStructGteFieldValidation(t *testing.T) {

	type Inner struct {
		CreatedAt *time.Time
		String    string
		Int       int
		Uint      uint
		Float     float64
		Array     []string
	}

	type Test struct {
		Inner     *Inner
		CreatedAt *time.Time `validate:"gtecsfield=Inner.CreatedAt"`
		String    string     `validate:"gtecsfield=Inner.String"`
		Int       int        `validate:"gtecsfield=Inner.Int"`
		Uint      uint       `validate:"gtecsfield=Inner.Uint"`
		Float     float64    `validate:"gtecsfield=Inner.Float"`
		Array     []string   `validate:"gtecsfield=Inner.Array"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * -5)

	inner := &Inner{
		CreatedAt: &then,
		String:    "abcd",
		Int:       13,
		Uint:      13,
		Float:     1.13,
		Array:     []string{"val1", "val2"},
	}

	test := &Test{
		Inner:     inner,
		CreatedAt: &now,
		String:    "abcde",
		Int:       14,
		Uint:      14,
		Float:     1.14,
		Array:     []string{"val1", "val2", "val3"},
	}

	validate := New()
	errs := validate.Struct(test)
	Equal(t, errs, nil)

	test.CreatedAt = &then
	test.String = "abcd"
	test.Int = 13
	test.Uint = 13
	test.Float = 1.13
	test.Array = []string{"val1", "val2"}

	errs = validate.Struct(test)
	Equal(t, errs, nil)

	before := now.Add(time.Hour * -10)

	test.CreatedAt = &before
	test.String = "abc"
	test.Int = 12
	test.Uint = 12
	test.Float = 1.12
	test.Array = []string{"val1"}

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "gtecsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "gtecsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "gtecsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "gtecsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "gtecsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "gtecsfield")

	errs = validate.VarWithValue(1, "", "gtecsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtecsfield")

	// this test is for the WARNING about unforseen validation issues.
	errs = validate.VarWithValue(test, now, "gtecsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "gtecsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "gtecsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "gtecsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "gtecsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "gtecsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "gtecsfield")

	type Other struct {
		Value string
	}

	type Test2 struct {
		Value Other
		Time  time.Time `validate:"gtecsfield=Value"`
	}

	tst := Test2{
		Value: Other{Value: "StringVal"},
		Time:  then,
	}

	errs = validate.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test2.Time", "Test2.Time", "Time", "Time", "gtecsfield")
}

func TestCrossStructGtFieldValidation(t *testing.T) {

	type Inner struct {
		CreatedAt *time.Time
		String    string
		Int       int
		Uint      uint
		Float     float64
		Array     []string
	}

	type Test struct {
		Inner     *Inner
		CreatedAt *time.Time `validate:"gtcsfield=Inner.CreatedAt"`
		String    string     `validate:"gtcsfield=Inner.String"`
		Int       int        `validate:"gtcsfield=Inner.Int"`
		Uint      uint       `validate:"gtcsfield=Inner.Uint"`
		Float     float64    `validate:"gtcsfield=Inner.Float"`
		Array     []string   `validate:"gtcsfield=Inner.Array"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * -5)

	inner := &Inner{
		CreatedAt: &then,
		String:    "abcd",
		Int:       13,
		Uint:      13,
		Float:     1.13,
		Array:     []string{"val1", "val2"},
	}

	test := &Test{
		Inner:     inner,
		CreatedAt: &now,
		String:    "abcde",
		Int:       14,
		Uint:      14,
		Float:     1.14,
		Array:     []string{"val1", "val2", "val3"},
	}

	validate := New()
	errs := validate.Struct(test)
	Equal(t, errs, nil)

	test.CreatedAt = &then
	test.String = "abcd"
	test.Int = 13
	test.Uint = 13
	test.Float = 1.13
	test.Array = []string{"val1", "val2"}

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "gtcsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "gtcsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "gtcsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "gtcsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "gtcsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "gtcsfield")

	errs = validate.VarWithValue(1, "", "gtcsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtcsfield")

	// this test is for the WARNING about unforseen validation issues.
	errs = validate.VarWithValue(test, now, "gtcsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "gtcsfield")
	AssertError(t, errs, "Test.String", "Test.String", "String", "String", "gtcsfield")
	AssertError(t, errs, "Test.Int", "Test.Int", "Int", "Int", "gtcsfield")
	AssertError(t, errs, "Test.Uint", "Test.Uint", "Uint", "Uint", "gtcsfield")
	AssertError(t, errs, "Test.Float", "Test.Float", "Float", "Float", "gtcsfield")
	AssertError(t, errs, "Test.Array", "Test.Array", "Array", "Array", "gtcsfield")

	type Other struct {
		Value string
	}

	type Test2 struct {
		Value Other
		Time  time.Time `validate:"gtcsfield=Value"`
	}

	tst := Test2{
		Value: Other{Value: "StringVal"},
		Time:  then,
	}

	errs = validate.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test2.Time", "Test2.Time", "Time", "Time", "gtcsfield")
}

func TestCrossStructNeFieldValidation(t *testing.T) {

	type Inner struct {
		CreatedAt *time.Time
	}

	type Test struct {
		Inner     *Inner
		CreatedAt *time.Time `validate:"necsfield=Inner.CreatedAt"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)

	inner := &Inner{
		CreatedAt: &then,
	}

	test := &Test{
		Inner:     inner,
		CreatedAt: &now,
	}

	validate := New()
	errs := validate.Struct(test)
	Equal(t, errs, nil)

	test.CreatedAt = &then

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "necsfield")

	var j uint64
	var k float64
	var j2 uint64
	var k2 float64
	s := "abcd"
	i := 1
	j = 1
	k = 1.543
	arr := []string{"test"}

	s2 := "abcd"
	i2 := 1
	j2 = 1
	k2 = 1.543
	arr2 := []string{"test"}
	arr3 := []string{"test", "test2"}
	now2 := now

	errs = validate.VarWithValue(s, s2, "necsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "necsfield")

	errs = validate.VarWithValue(i2, i, "necsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "necsfield")

	errs = validate.VarWithValue(j2, j, "necsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "necsfield")

	errs = validate.VarWithValue(k2, k, "necsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "necsfield")

	errs = validate.VarWithValue(arr2, arr, "necsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "necsfield")

	errs = validate.VarWithValue(now2, now, "necsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "necsfield")

	errs = validate.VarWithValue(arr3, arr, "necsfield")
	Equal(t, errs, nil)

	type SInner struct {
		Name string
	}

	type TStruct struct {
		Inner     *SInner
		CreatedAt *time.Time `validate:"necsfield=Inner"`
	}

	sinner := &SInner{
		Name: "NAME",
	}

	test2 := &TStruct{
		Inner:     sinner,
		CreatedAt: &now,
	}

	errs = validate.Struct(test2)
	Equal(t, errs, nil)

	test2.Inner = nil
	errs = validate.Struct(test2)
	Equal(t, errs, nil)

	errs = validate.VarWithValue(nil, 1, "necsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "necsfield")
}

func TestCrossStructEqFieldValidation(t *testing.T) {

	type Inner struct {
		CreatedAt *time.Time
	}

	type Test struct {
		Inner     *Inner
		CreatedAt *time.Time `validate:"eqcsfield=Inner.CreatedAt"`
	}

	now := time.Now().UTC()

	inner := &Inner{
		CreatedAt: &now,
	}

	test := &Test{
		Inner:     inner,
		CreatedAt: &now,
	}

	validate := New()
	errs := validate.Struct(test)
	Equal(t, errs, nil)

	newTime := time.Now().UTC()
	test.CreatedAt = &newTime

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.CreatedAt", "Test.CreatedAt", "CreatedAt", "CreatedAt", "eqcsfield")

	var j uint64
	var k float64
	s := "abcd"
	i := 1
	j = 1
	k = 1.543
	arr := []string{"test"}

	var j2 uint64
	var k2 float64
	s2 := "abcd"
	i2 := 1
	j2 = 1
	k2 = 1.543
	arr2 := []string{"test"}
	arr3 := []string{"test", "test2"}
	now2 := now

	errs = validate.VarWithValue(s, s2, "eqcsfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(i2, i, "eqcsfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(j2, j, "eqcsfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(k2, k, "eqcsfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(arr2, arr, "eqcsfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(now2, now, "eqcsfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(arr3, arr, "eqcsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "eqcsfield")

	type SInner struct {
		Name string
	}

	type TStruct struct {
		Inner     *SInner
		CreatedAt *time.Time `validate:"eqcsfield=Inner"`
	}

	sinner := &SInner{
		Name: "NAME",
	}

	test2 := &TStruct{
		Inner:     sinner,
		CreatedAt: &now,
	}

	errs = validate.Struct(test2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TStruct.CreatedAt", "TStruct.CreatedAt", "CreatedAt", "CreatedAt", "eqcsfield")

	test2.Inner = nil
	errs = validate.Struct(test2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TStruct.CreatedAt", "TStruct.CreatedAt", "CreatedAt", "CreatedAt", "eqcsfield")

	errs = validate.VarWithValue(nil, 1, "eqcsfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "eqcsfield")
}

func TestCrossNamespaceFieldValidation(t *testing.T) {

	type SliceStruct struct {
		Name string
	}

	type MapStruct struct {
		Name string
	}

	type Inner struct {
		CreatedAt        *time.Time
		Slice            []string
		SliceStructs     []*SliceStruct
		SliceSlice       [][]string
		SliceSliceStruct [][]*SliceStruct
		SliceMap         []map[string]string
		Map              map[string]string
		MapMap           map[string]map[string]string
		MapStructs       map[string]*SliceStruct
		MapMapStruct     map[string]map[string]*SliceStruct
		MapSlice         map[string][]string
		MapInt           map[int]string
		MapInt8          map[int8]string
		MapInt16         map[int16]string
		MapInt32         map[int32]string
		MapInt64         map[int64]string
		MapUint          map[uint]string
		MapUint8         map[uint8]string
		MapUint16        map[uint16]string
		MapUint32        map[uint32]string
		MapUint64        map[uint64]string
		MapFloat32       map[float32]string
		MapFloat64       map[float64]string
		MapBool          map[bool]string
	}

	type Test struct {
		Inner     *Inner
		CreatedAt *time.Time
	}

	now := time.Now()

	inner := &Inner{
		CreatedAt:        &now,
		Slice:            []string{"val1", "val2", "val3"},
		SliceStructs:     []*SliceStruct{{Name: "name1"}, {Name: "name2"}, {Name: "name3"}},
		SliceSlice:       [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
		SliceSliceStruct: [][]*SliceStruct{{{Name: "name1"}, {Name: "name2"}, {Name: "name3"}}, {{Name: "name4"}, {Name: "name5"}, {Name: "name6"}}, {{Name: "name7"}, {Name: "name8"}, {Name: "name9"}}},
		SliceMap:         []map[string]string{{"key1": "val1", "key2": "val2", "key3": "val3"}, {"key4": "val4", "key5": "val5", "key6": "val6"}},
		Map:              map[string]string{"key1": "val1", "key2": "val2", "key3": "val3"},
		MapStructs:       map[string]*SliceStruct{"key1": {Name: "name1"}, "key2": {Name: "name2"}, "key3": {Name: "name3"}},
		MapMap:           map[string]map[string]string{"key1": {"key1-1": "val1"}, "key2": {"key2-1": "val2"}, "key3": {"key3-1": "val3"}},
		MapMapStruct:     map[string]map[string]*SliceStruct{"key1": {"key1-1": {Name: "name1"}}, "key2": {"key2-1": {Name: "name2"}}, "key3": {"key3-1": {Name: "name3"}}},
		MapSlice:         map[string][]string{"key1": {"1", "2", "3"}, "key2": {"4", "5", "6"}, "key3": {"7", "8", "9"}},
		MapInt:           map[int]string{1: "val1", 2: "val2", 3: "val3"},
		MapInt8:          map[int8]string{1: "val1", 2: "val2", 3: "val3"},
		MapInt16:         map[int16]string{1: "val1", 2: "val2", 3: "val3"},
		MapInt32:         map[int32]string{1: "val1", 2: "val2", 3: "val3"},
		MapInt64:         map[int64]string{1: "val1", 2: "val2", 3: "val3"},
		MapUint:          map[uint]string{1: "val1", 2: "val2", 3: "val3"},
		MapUint8:         map[uint8]string{1: "val1", 2: "val2", 3: "val3"},
		MapUint16:        map[uint16]string{1: "val1", 2: "val2", 3: "val3"},
		MapUint32:        map[uint32]string{1: "val1", 2: "val2", 3: "val3"},
		MapUint64:        map[uint64]string{1: "val1", 2: "val2", 3: "val3"},
		MapFloat32:       map[float32]string{1.01: "val1", 2.02: "val2", 3.03: "val3"},
		MapFloat64:       map[float64]string{1.01: "val1", 2.02: "val2", 3.03: "val3"},
		MapBool:          map[bool]string{true: "val1", false: "val2"},
	}

	test := &Test{
		Inner:     inner,
		CreatedAt: &now,
	}

	val := reflect.ValueOf(test)

	vd := New()
	v := &validate{
		v: vd,
	}

	current, kind, ok := v.getStructFieldOKInternal(val, "Inner.CreatedAt")
	Equal(t, ok, true)
	Equal(t, kind, reflect.Struct)
	tm, ok := current.Interface().(time.Time)
	Equal(t, ok, true)
	Equal(t, tm, now)

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.Slice[1]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.CrazyNonExistantField")
	Equal(t, ok, false)

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.Slice[101]")
	Equal(t, ok, false)

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.Map[key3]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val3")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapMap[key2][key2-1]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapStructs[key2].Name")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "name2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapMapStruct[key3][key3-1].Name")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "name3")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.SliceSlice[2][0]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "7")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.SliceSliceStruct[2][1].Name")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "name8")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.SliceMap[1][key5]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val5")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapSlice[key3][2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "9")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapInt[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapInt8[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapInt16[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapInt32[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapInt64[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapUint[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapUint8[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapUint16[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapUint32[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapUint64[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapFloat32[3.03]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val3")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapFloat64[2.02]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val2")

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.MapBool[true]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.String)
	Equal(t, current.String(), "val1")

	inner = &Inner{
		CreatedAt:        &now,
		Slice:            []string{"val1", "val2", "val3"},
		SliceStructs:     []*SliceStruct{{Name: "name1"}, {Name: "name2"}, nil},
		SliceSlice:       [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
		SliceSliceStruct: [][]*SliceStruct{{{Name: "name1"}, {Name: "name2"}, {Name: "name3"}}, {{Name: "name4"}, {Name: "name5"}, {Name: "name6"}}, {{Name: "name7"}, {Name: "name8"}, {Name: "name9"}}},
		SliceMap:         []map[string]string{{"key1": "val1", "key2": "val2", "key3": "val3"}, {"key4": "val4", "key5": "val5", "key6": "val6"}},
		Map:              map[string]string{"key1": "val1", "key2": "val2", "key3": "val3"},
		MapStructs:       map[string]*SliceStruct{"key1": {Name: "name1"}, "key2": {Name: "name2"}, "key3": {Name: "name3"}},
		MapMap:           map[string]map[string]string{"key1": {"key1-1": "val1"}, "key2": {"key2-1": "val2"}, "key3": {"key3-1": "val3"}},
		MapMapStruct:     map[string]map[string]*SliceStruct{"key1": {"key1-1": {Name: "name1"}}, "key2": {"key2-1": {Name: "name2"}}, "key3": {"key3-1": {Name: "name3"}}},
		MapSlice:         map[string][]string{"key1": {"1", "2", "3"}, "key2": {"4", "5", "6"}, "key3": {"7", "8", "9"}},
	}

	test = &Test{
		Inner:     inner,
		CreatedAt: nil,
	}

	val = reflect.ValueOf(test)

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.SliceStructs[2]")
	Equal(t, ok, true)
	Equal(t, kind, reflect.Ptr)
	Equal(t, current.String(), "<*validator.SliceStruct Value>")
	Equal(t, current.IsNil(), true)

	current, kind, ok = v.getStructFieldOKInternal(val, "Inner.SliceStructs[2].Name")
	Equal(t, ok, false)
	Equal(t, kind, reflect.Ptr)
	Equal(t, current.String(), "<*validator.SliceStruct Value>")
	Equal(t, current.IsNil(), true)

	PanicMatches(t, func() { v.getStructFieldOKInternal(reflect.ValueOf(1), "crazyinput") }, "Invalid field namespace")
}

func TestExistsValidation(t *testing.T) {

	jsonText := "{ \"truthiness2\": true }"

	type Thing struct {
		Truthiness *bool `json:"truthiness" validate:"required"`
	}

	var ting Thing

	err := json.Unmarshal([]byte(jsonText), &ting)
	Equal(t, err, nil)
	NotEqual(t, ting, nil)
	Equal(t, ting.Truthiness, nil)

	validate := New()
	errs := validate.Struct(ting)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Thing.Truthiness", "Thing.Truthiness", "Truthiness", "Truthiness", "required")

	jsonText = "{ \"truthiness\": true }"

	err = json.Unmarshal([]byte(jsonText), &ting)
	Equal(t, err, nil)
	NotEqual(t, ting, nil)
	Equal(t, ting.Truthiness, true)

	errs = validate.Struct(ting)
	Equal(t, errs, nil)
}

func TestSQLValue2Validation(t *testing.T) {

	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, valuer{}, (*driver.Valuer)(nil), sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
	validate.RegisterCustomTypeFunc(ValidateCustomType, MadeUpCustomType{})
	validate.RegisterCustomTypeFunc(OverrideIntTypeForSomeReason, 1)

	val := valuer{
		Name: "",
	}

	errs := validate.Var(val, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	val.Name = "Valid Name"
	errs = validate.VarCtx(context.Background(), val, "required")
	Equal(t, errs, nil)

	val.Name = "errorme"

	PanicMatches(t, func() { validate.Var(val, "required") }, "SQL Driver Valuer error: some kind of error")

	type myValuer valuer

	myVal := valuer{
		Name: "",
	}

	errs = validate.Var(myVal, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	cust := MadeUpCustomType{
		FirstName: "Joey",
		LastName:  "Bloggs",
	}

	c := CustomMadeUpStruct{MadeUp: cust, OverriddenInt: 2}

	errs = validate.Struct(c)
	Equal(t, errs, nil)

	c.MadeUp.FirstName = ""
	c.OverriddenInt = 1

	errs = validate.Struct(c)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 2)
	AssertError(t, errs, "CustomMadeUpStruct.MadeUp", "CustomMadeUpStruct.MadeUp", "MadeUp", "MadeUp", "required")
	AssertError(t, errs, "CustomMadeUpStruct.OverriddenInt", "CustomMadeUpStruct.OverriddenInt", "OverriddenInt", "OverriddenInt", "gt")
}

func TestSQLValueValidation(t *testing.T) {

	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*driver.Valuer)(nil), valuer{})
	validate.RegisterCustomTypeFunc(ValidateCustomType, MadeUpCustomType{})
	validate.RegisterCustomTypeFunc(OverrideIntTypeForSomeReason, 1)

	val := valuer{
		Name: "",
	}

	errs := validate.Var(val, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	val.Name = "Valid Name"
	errs = validate.Var(val, "required")
	Equal(t, errs, nil)

	val.Name = "errorme"

	PanicMatches(t, func() { errs = validate.Var(val, "required") }, "SQL Driver Valuer error: some kind of error")

	type myValuer valuer

	myVal := valuer{
		Name: "",
	}

	errs = validate.Var(myVal, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	cust := MadeUpCustomType{
		FirstName: "Joey",
		LastName:  "Bloggs",
	}

	c := CustomMadeUpStruct{MadeUp: cust, OverriddenInt: 2}

	errs = validate.Struct(c)
	Equal(t, errs, nil)

	c.MadeUp.FirstName = ""
	c.OverriddenInt = 1

	errs = validate.Struct(c)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 2)
	AssertError(t, errs, "CustomMadeUpStruct.MadeUp", "CustomMadeUpStruct.MadeUp", "MadeUp", "MadeUp", "required")
	AssertError(t, errs, "CustomMadeUpStruct.OverriddenInt", "CustomMadeUpStruct.OverriddenInt", "OverriddenInt", "OverriddenInt", "gt")
}

func TestMACValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"3D:F2:C9:A6:B3:4F", true},
		{"3D-F2-C9-A6-B3:4F", false},
		{"123", false},
		{"", false},
		{"abacaba", false},
		{"00:25:96:FF:FE:12:34:56", true},
		{"0025:96FF:FE12:3456", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "mac")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d mac failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d mac failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "mac" {
					t.Fatalf("Index: %d mac failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestIPValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"192.168.0.1", true},
		{"192.168.255.254", true},
		{"192.168.255.256", false},
		{"172.16.255.254", true},
		{"172.16.256.255", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", true},
		{"2001:cdba:0:0:0:0:3257:9652", true},
		{"2001:cdba::3257:9652", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "ip")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ip failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ip failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ip" {
					t.Fatalf("Index: %d ip failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestIPv6Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"10.0.0.1", false},
		{"172.16.0.1", false},
		{"192.168.0.1", false},
		{"192.168.255.254", false},
		{"192.168.255.256", false},
		{"172.16.255.254", false},
		{"172.16.256.255", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", true},
		{"2001:cdba:0:0:0:0:3257:9652", true},
		{"2001:cdba::3257:9652", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "ipv6")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ipv6 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ipv6 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ipv6" {
					t.Fatalf("Index: %d ipv6 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestIPv4Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"192.168.0.1", true},
		{"192.168.255.254", true},
		{"192.168.255.256", false},
		{"172.16.255.254", true},
		{"172.16.256.255", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
		{"2001:cdba:0:0:0:0:3257:9652", false},
		{"2001:cdba::3257:9652", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "ipv4")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ipv4 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ipv4 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ipv4" {
					t.Fatalf("Index: %d ipv4 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestCIDRValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"10.0.0.0/0", true},
		{"10.0.0.1/8", true},
		{"172.16.0.1/16", true},
		{"192.168.0.1/24", true},
		{"192.168.255.254/24", true},
		{"192.168.255.254/48", false},
		{"192.168.255.256/24", false},
		{"172.16.255.254/16", true},
		{"172.16.256.255/16", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652/64", true},
		{"2001:cdba:0000:0000:0000:0000:3257:9652/256", false},
		{"2001:cdba:0:0:0:0:3257:9652/32", true},
		{"2001:cdba::3257:9652/16", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "cidr")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d cidr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d cidr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "cidr" {
					t.Fatalf("Index: %d cidr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestCIDRv6Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"10.0.0.0/0", false},
		{"10.0.0.1/8", false},
		{"172.16.0.1/16", false},
		{"192.168.0.1/24", false},
		{"192.168.255.254/24", false},
		{"192.168.255.254/48", false},
		{"192.168.255.256/24", false},
		{"172.16.255.254/16", false},
		{"172.16.256.255/16", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652/64", true},
		{"2001:cdba:0000:0000:0000:0000:3257:9652/256", false},
		{"2001:cdba:0:0:0:0:3257:9652/32", true},
		{"2001:cdba::3257:9652/16", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "cidrv6")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d cidrv6 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d cidrv6 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "cidrv6" {
					t.Fatalf("Index: %d cidrv6 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestCIDRv4Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"10.0.0.0/0", true},
		{"10.0.0.1/8", true},
		{"172.16.0.1/16", true},
		{"192.168.0.1/24", true},
		{"192.168.255.254/24", true},
		{"192.168.255.254/48", false},
		{"192.168.255.256/24", false},
		{"172.16.255.254/16", true},
		{"172.16.256.255/16", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652/64", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652/256", false},
		{"2001:cdba:0:0:0:0:3257:9652/32", false},
		{"2001:cdba::3257:9652/16", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "cidrv4")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d cidrv4 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d cidrv4 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "cidrv4" {
					t.Fatalf("Index: %d cidrv4 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestTCPAddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{":80", false},
		{"127.0.0.1:80", true},
		{"[::1]:80", true},
		{"256.0.0.0:1", false},
		{"[::1]", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "tcp_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d tcp_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d tcp_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "tcp_addr" {
					t.Fatalf("Index: %d tcp_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestTCP6AddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{":80", false},
		{"127.0.0.1:80", false},
		{"[::1]:80", true},
		{"256.0.0.0:1", false},
		{"[::1]", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "tcp6_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d tcp6_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d tcp6_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "tcp6_addr" {
					t.Fatalf("Index: %d tcp6_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestTCP4AddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{":80", false},
		{"127.0.0.1:80", true},
		{"[::1]:80", false}, // https://github.com/golang/go/issues/14037
		{"256.0.0.0:1", false},
		{"[::1]", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "tcp4_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d tcp4_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Log(test.param, IsEqual(errs, nil))
				t.Fatalf("Index: %d tcp4_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "tcp4_addr" {
					t.Fatalf("Index: %d tcp4_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUDPAddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{":80", false},
		{"127.0.0.1:80", true},
		{"[::1]:80", true},
		{"256.0.0.0:1", false},
		{"[::1]", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "udp_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d udp_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d udp_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "udp_addr" {
					t.Fatalf("Index: %d udp_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUDP6AddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{":80", false},
		{"127.0.0.1:80", false},
		{"[::1]:80", true},
		{"256.0.0.0:1", false},
		{"[::1]", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "udp6_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d udp6_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d udp6_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "udp6_addr" {
					t.Fatalf("Index: %d udp6_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUDP4AddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{":80", false},
		{"127.0.0.1:80", true},
		{"[::1]:80", false}, // https://github.com/golang/go/issues/14037
		{"256.0.0.0:1", false},
		{"[::1]", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "udp4_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d udp4_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Log(test.param, IsEqual(errs, nil))
				t.Fatalf("Index: %d udp4_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "udp4_addr" {
					t.Fatalf("Index: %d udp4_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestIPAddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"127.0.0.1:80", false},
		{"::1", true},
		{"256.0.0.0", false},
		{"localhost", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "ip_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ip_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ip_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ip_addr" {
					t.Fatalf("Index: %d ip_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestIP6AddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", false}, // https://github.com/golang/go/issues/14037
		{"127.0.0.1:80", false},
		{"::1", true},
		{"0:0:0:0:0:0:0:1", true},
		{"256.0.0.0", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "ip6_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ip6_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ip6_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ip6_addr" {
					t.Fatalf("Index: %d ip6_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestIP4AddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"127.0.0.1:80", false},
		{"::1", false}, // https://github.com/golang/go/issues/14037
		{"256.0.0.0", false},
		{"localhost", false},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "ip4_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ip4_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Log(test.param, IsEqual(errs, nil))
				t.Fatalf("Index: %d ip4_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ip4_addr" {
					t.Fatalf("Index: %d ip4_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUnixAddrValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"v.sock", true},
	}

	validate := New()

	for i, test := range tests {
		errs := validate.Var(test.param, "unix_addr")
		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d unix_addr failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Log(test.param, IsEqual(errs, nil))
				t.Fatalf("Index: %d unix_addr failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "unix_addr" {
					t.Fatalf("Index: %d unix_addr failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestSliceMapArrayChanFuncPtrInterfaceRequiredValidation(t *testing.T) {

	validate := New()

	var m map[string]string

	errs := validate.Var(m, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	m = map[string]string{}
	errs = validate.Var(m, "required")
	Equal(t, errs, nil)

	var arr [5]string
	errs = validate.Var(arr, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	arr[0] = "ok"
	errs = validate.Var(arr, "required")
	Equal(t, errs, nil)

	var s []string
	errs = validate.Var(s, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	s = []string{}
	errs = validate.Var(s, "required")
	Equal(t, errs, nil)

	var c chan string
	errs = validate.Var(c, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	c = make(chan string)
	errs = validate.Var(c, "required")
	Equal(t, errs, nil)

	var tst *int
	errs = validate.Var(tst, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	one := 1
	tst = &one
	errs = validate.Var(tst, "required")
	Equal(t, errs, nil)

	var iface interface{}

	errs = validate.Var(iface, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	errs = validate.Var(iface, "omitempty,required")
	Equal(t, errs, nil)

	errs = validate.Var(iface, "")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(nil, iface, "")
	Equal(t, errs, nil)

	var f func(string)

	errs = validate.Var(f, "required")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "required")

	f = func(name string) {}

	errs = validate.Var(f, "required")
	Equal(t, errs, nil)
}

func TestDatePtrValidationIssueValidation(t *testing.T) {

	type Test struct {
		LastViewed *time.Time
		Reminder   *time.Time
	}

	test := &Test{}

	validate := New()
	errs := validate.Struct(test)
	Equal(t, errs, nil)
}

func TestCommaAndPipeObfuscationValidation(t *testing.T) {
	s := "My Name Is, |joeybloggs|"

	validate := New()

	errs := validate.Var(s, "excludesall=0x2C")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "excludesall")

	errs = validate.Var(s, "excludesall=0x7C")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "excludesall")
}

func TestBadKeyValidation(t *testing.T) {
	type Test struct {
		Name string `validate:"required, "`
	}

	tst := &Test{
		Name: "test",
	}

	validate := New()

	PanicMatches(t, func() { validate.Struct(tst) }, "Undefined validation function ' ' on field 'Name'")

	type Test2 struct {
		Name string `validate:"required,,len=2"`
	}

	tst2 := &Test2{
		Name: "test",
	}

	PanicMatches(t, func() { validate.Struct(tst2) }, "Invalid validation tag on field 'Name'")
}

func TestInterfaceErrValidation(t *testing.T) {

	var v1 interface{}
	var v2 interface{}

	v2 = 1
	v1 = v2

	validate := New()
	errs := validate.Var(v1, "len=1")
	Equal(t, errs, nil)

	errs = validate.Var(v2, "len=1")
	Equal(t, errs, nil)

	type ExternalCMD struct {
		Userid string      `json:"userid"`
		Action uint32      `json:"action"`
		Data   interface{} `json:"data,omitempty" validate:"required"`
	}

	s := &ExternalCMD{
		Userid: "123456",
		Action: 10000,
		// Data:   1,
	}

	errs = validate.Struct(s)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "ExternalCMD.Data", "ExternalCMD.Data", "Data", "Data", "required")

	type ExternalCMD2 struct {
		Userid string      `json:"userid"`
		Action uint32      `json:"action"`
		Data   interface{} `json:"data,omitempty" validate:"len=1"`
	}

	s2 := &ExternalCMD2{
		Userid: "123456",
		Action: 10000,
		// Data:   1,
	}

	errs = validate.Struct(s2)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "ExternalCMD2.Data", "ExternalCMD2.Data", "Data", "Data", "len")

	s3 := &ExternalCMD2{
		Userid: "123456",
		Action: 10000,
		Data:   2,
	}

	errs = validate.Struct(s3)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "ExternalCMD2.Data", "ExternalCMD2.Data", "Data", "Data", "len")

	type Inner struct {
		Name string `validate:"required"`
	}

	inner := &Inner{
		Name: "",
	}

	s4 := &ExternalCMD{
		Userid: "123456",
		Action: 10000,
		Data:   inner,
	}

	errs = validate.Struct(s4)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "ExternalCMD.Data.Name", "ExternalCMD.Data.Name", "Name", "Name", "required")

	type TestMapStructPtr struct {
		Errs map[int]interface{} `validate:"gt=0,dive,len=2"`
	}

	mip := map[int]interface{}{0: &Inner{"ok"}, 3: nil, 4: &Inner{"ok"}}

	msp := &TestMapStructPtr{
		Errs: mip,
	}

	errs = validate.Struct(msp)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "TestMapStructPtr.Errs[3]", "TestMapStructPtr.Errs[3]", "Errs[3]", "Errs[3]", "len")

	type TestMultiDimensionalStructs struct {
		Errs [][]interface{} `validate:"gt=0,dive,dive"`
	}

	var errStructArray [][]interface{}

	errStructArray = append(errStructArray, []interface{}{&Inner{"ok"}, &Inner{""}, &Inner{""}})
	errStructArray = append(errStructArray, []interface{}{&Inner{"ok"}, &Inner{""}, &Inner{""}})

	tms := &TestMultiDimensionalStructs{
		Errs: errStructArray,
	}

	errs = validate.Struct(tms)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 4)
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[0][1].Name", "TestMultiDimensionalStructs.Errs[0][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[0][2].Name", "TestMultiDimensionalStructs.Errs[0][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[1][1].Name", "TestMultiDimensionalStructs.Errs[1][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[1][2].Name", "TestMultiDimensionalStructs.Errs[1][2].Name", "Name", "Name", "required")

	type TestMultiDimensionalStructsPtr2 struct {
		Errs [][]*Inner `validate:"gt=0,dive,dive,required"`
	}

	var errStructPtr2Array [][]*Inner

	errStructPtr2Array = append(errStructPtr2Array, []*Inner{{"ok"}, {""}, {""}})
	errStructPtr2Array = append(errStructPtr2Array, []*Inner{{"ok"}, {""}, {""}})
	errStructPtr2Array = append(errStructPtr2Array, []*Inner{{"ok"}, {""}, nil})

	tmsp2 := &TestMultiDimensionalStructsPtr2{
		Errs: errStructPtr2Array,
	}

	errs = validate.Struct(tmsp2)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 6)
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[0][1].Name", "TestMultiDimensionalStructsPtr2.Errs[0][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[0][2].Name", "TestMultiDimensionalStructsPtr2.Errs[0][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[1][1].Name", "TestMultiDimensionalStructsPtr2.Errs[1][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[1][2].Name", "TestMultiDimensionalStructsPtr2.Errs[1][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[2][1].Name", "TestMultiDimensionalStructsPtr2.Errs[2][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[2][2]", "TestMultiDimensionalStructsPtr2.Errs[2][2]", "Errs[2][2]", "Errs[2][2]", "required")

	m := map[int]interface{}{0: "ok", 3: "", 4: "ok"}

	errs = validate.Var(m, "len=3,dive,len=2")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "[3]", "[3]", "[3]", "[3]", "len")

	errs = validate.Var(m, "len=2,dive,required")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "", "", "", "", "len")

	arr := []interface{}{"ok", "", "ok"}

	errs = validate.Var(arr, "len=3,dive,len=2")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "[1]", "[1]", "[1]", "[1]", "len")

	errs = validate.Var(arr, "len=2,dive,required")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "", "", "", "", "len")

	type MyStruct struct {
		A, B string
		C    interface{}
	}

	var a MyStruct

	a.A = "value"
	a.C = "nu"

	errs = validate.Struct(a)
	Equal(t, errs, nil)
}

func TestMapDiveValidation(t *testing.T) {

	validate := New()

	n := map[int]interface{}{0: nil}
	errs := validate.Var(n, "omitempty,required")
	Equal(t, errs, nil)

	m := map[int]string{0: "ok", 3: "", 4: "ok"}

	errs = validate.Var(m, "len=3,dive,required")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "[3]", "[3]", "[3]", "[3]", "required")

	errs = validate.Var(m, "len=2,dive,required")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "", "", "", "", "len")

	type Inner struct {
		Name string `validate:"required"`
	}

	type TestMapStruct struct {
		Errs map[int]Inner `validate:"gt=0,dive"`
	}

	mi := map[int]Inner{0: {"ok"}, 3: {""}, 4: {"ok"}}

	ms := &TestMapStruct{
		Errs: mi,
	}

	errs = validate.Struct(ms)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "TestMapStruct.Errs[3].Name", "TestMapStruct.Errs[3].Name", "Name", "Name", "required")

	// for full test coverage
	s := fmt.Sprint(errs.Error())
	NotEqual(t, s, "")

	type TestMapTimeStruct struct {
		Errs map[int]*time.Time `validate:"gt=0,dive,required"`
	}

	t1 := time.Now().UTC()

	mta := map[int]*time.Time{0: &t1, 3: nil, 4: nil}

	mt := &TestMapTimeStruct{
		Errs: mta,
	}

	errs = validate.Struct(mt)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 2)
	AssertError(t, errs, "TestMapTimeStruct.Errs[3]", "TestMapTimeStruct.Errs[3]", "Errs[3]", "Errs[3]", "required")
	AssertError(t, errs, "TestMapTimeStruct.Errs[4]", "TestMapTimeStruct.Errs[4]", "Errs[4]", "Errs[4]", "required")

	type TestMapStructPtr struct {
		Errs map[int]*Inner `validate:"gt=0,dive,required"`
	}

	mip := map[int]*Inner{0: {"ok"}, 3: nil, 4: {"ok"}}

	msp := &TestMapStructPtr{
		Errs: mip,
	}

	errs = validate.Struct(msp)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "TestMapStructPtr.Errs[3]", "TestMapStructPtr.Errs[3]", "Errs[3]", "Errs[3]", "required")

	type TestMapStructPtr2 struct {
		Errs map[int]*Inner `validate:"gt=0,dive,omitempty,required"`
	}

	mip2 := map[int]*Inner{0: {"ok"}, 3: nil, 4: {"ok"}}

	msp2 := &TestMapStructPtr2{
		Errs: mip2,
	}

	errs = validate.Struct(msp2)
	Equal(t, errs, nil)

	v2 := New()
	v2.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type MapDiveJSONTest struct {
		Map map[string]string `validate:"required,gte=1,dive,gte=1" json:"MyName"`
	}

	mdjt := &MapDiveJSONTest{
		Map: map[string]string{
			"Key1": "Value1",
			"Key2": "",
		},
	}

	err := v2.Struct(mdjt)
	NotEqual(t, err, nil)

	errs = err.(ValidationErrors)
	fe := getError(errs, "MapDiveJSONTest.MyName[Key2]", "MapDiveJSONTest.Map[Key2]")
	NotEqual(t, fe, nil)
	Equal(t, fe.Tag(), "gte")
	Equal(t, fe.ActualTag(), "gte")
	Equal(t, fe.Field(), "MyName[Key2]")
	Equal(t, fe.StructField(), "Map[Key2]")
}

func TestArrayDiveValidation(t *testing.T) {

	validate := New()

	arr := []string{"ok", "", "ok"}

	errs := validate.Var(arr, "len=3,dive,required")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "[1]", "[1]", "[1]", "[1]", "required")

	errs = validate.Var(arr, "len=2,dive,required")
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "", "", "", "", "len")

	type BadDive struct {
		Name string `validate:"dive"`
	}

	bd := &BadDive{
		Name: "TEST",
	}

	PanicMatches(t, func() { validate.Struct(bd) }, "dive error! can't dive on a non slice or map")

	type Test struct {
		Errs []string `validate:"gt=0,dive,required"`
	}

	test := &Test{
		Errs: []string{"ok", "", "ok"},
	}

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "Test.Errs[1]", "Test.Errs[1]", "Errs[1]", "Errs[1]", "required")

	test = &Test{
		Errs: []string{"ok", "ok", ""},
	}

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "Test.Errs[2]", "Test.Errs[2]", "Errs[2]", "Errs[2]", "required")

	type TestMultiDimensional struct {
		Errs [][]string `validate:"gt=0,dive,dive,required"`
	}

	var errArray [][]string

	errArray = append(errArray, []string{"ok", "", ""})
	errArray = append(errArray, []string{"ok", "", ""})

	tm := &TestMultiDimensional{
		Errs: errArray,
	}

	errs = validate.Struct(tm)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 4)
	AssertError(t, errs, "TestMultiDimensional.Errs[0][1]", "TestMultiDimensional.Errs[0][1]", "Errs[0][1]", "Errs[0][1]", "required")
	AssertError(t, errs, "TestMultiDimensional.Errs[0][2]", "TestMultiDimensional.Errs[0][2]", "Errs[0][2]", "Errs[0][2]", "required")
	AssertError(t, errs, "TestMultiDimensional.Errs[1][1]", "TestMultiDimensional.Errs[1][1]", "Errs[1][1]", "Errs[1][1]", "required")
	AssertError(t, errs, "TestMultiDimensional.Errs[1][2]", "TestMultiDimensional.Errs[1][2]", "Errs[1][2]", "Errs[1][2]", "required")

	type Inner struct {
		Name string `validate:"required"`
	}

	type TestMultiDimensionalStructs struct {
		Errs [][]Inner `validate:"gt=0,dive,dive"`
	}

	var errStructArray [][]Inner

	errStructArray = append(errStructArray, []Inner{{"ok"}, {""}, {""}})
	errStructArray = append(errStructArray, []Inner{{"ok"}, {""}, {""}})

	tms := &TestMultiDimensionalStructs{
		Errs: errStructArray,
	}

	errs = validate.Struct(tms)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 4)
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[0][1].Name", "TestMultiDimensionalStructs.Errs[0][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[0][2].Name", "TestMultiDimensionalStructs.Errs[0][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[1][1].Name", "TestMultiDimensionalStructs.Errs[1][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructs.Errs[1][2].Name", "TestMultiDimensionalStructs.Errs[1][2].Name", "Name", "Name", "required")

	type TestMultiDimensionalStructsPtr struct {
		Errs [][]*Inner `validate:"gt=0,dive,dive"`
	}

	var errStructPtrArray [][]*Inner

	errStructPtrArray = append(errStructPtrArray, []*Inner{{"ok"}, {""}, {""}})
	errStructPtrArray = append(errStructPtrArray, []*Inner{{"ok"}, {""}, {""}})
	errStructPtrArray = append(errStructPtrArray, []*Inner{{"ok"}, {""}, nil})

	tmsp := &TestMultiDimensionalStructsPtr{
		Errs: errStructPtrArray,
	}

	errs = validate.Struct(tmsp)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 5)
	AssertError(t, errs, "TestMultiDimensionalStructsPtr.Errs[0][1].Name", "TestMultiDimensionalStructsPtr.Errs[0][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr.Errs[0][2].Name", "TestMultiDimensionalStructsPtr.Errs[0][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr.Errs[1][1].Name", "TestMultiDimensionalStructsPtr.Errs[1][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr.Errs[1][2].Name", "TestMultiDimensionalStructsPtr.Errs[1][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr.Errs[2][1].Name", "TestMultiDimensionalStructsPtr.Errs[2][1].Name", "Name", "Name", "required")

	// for full test coverage
	s := fmt.Sprint(errs.Error())
	NotEqual(t, s, "")

	type TestMultiDimensionalStructsPtr2 struct {
		Errs [][]*Inner `validate:"gt=0,dive,dive,required"`
	}

	var errStructPtr2Array [][]*Inner

	errStructPtr2Array = append(errStructPtr2Array, []*Inner{{"ok"}, {""}, {""}})
	errStructPtr2Array = append(errStructPtr2Array, []*Inner{{"ok"}, {""}, {""}})
	errStructPtr2Array = append(errStructPtr2Array, []*Inner{{"ok"}, {""}, nil})

	tmsp2 := &TestMultiDimensionalStructsPtr2{
		Errs: errStructPtr2Array,
	}

	errs = validate.Struct(tmsp2)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 6)
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[0][1].Name", "TestMultiDimensionalStructsPtr2.Errs[0][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[0][2].Name", "TestMultiDimensionalStructsPtr2.Errs[0][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[1][1].Name", "TestMultiDimensionalStructsPtr2.Errs[1][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[1][2].Name", "TestMultiDimensionalStructsPtr2.Errs[1][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[2][1].Name", "TestMultiDimensionalStructsPtr2.Errs[2][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr2.Errs[2][2]", "TestMultiDimensionalStructsPtr2.Errs[2][2]", "Errs[2][2]", "Errs[2][2]", "required")

	type TestMultiDimensionalStructsPtr3 struct {
		Errs [][]*Inner `validate:"gt=0,dive,dive,omitempty"`
	}

	var errStructPtr3Array [][]*Inner

	errStructPtr3Array = append(errStructPtr3Array, []*Inner{{"ok"}, {""}, {""}})
	errStructPtr3Array = append(errStructPtr3Array, []*Inner{{"ok"}, {""}, {""}})
	errStructPtr3Array = append(errStructPtr3Array, []*Inner{{"ok"}, {""}, nil})

	tmsp3 := &TestMultiDimensionalStructsPtr3{
		Errs: errStructPtr3Array,
	}

	errs = validate.Struct(tmsp3)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 5)
	AssertError(t, errs, "TestMultiDimensionalStructsPtr3.Errs[0][1].Name", "TestMultiDimensionalStructsPtr3.Errs[0][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr3.Errs[0][2].Name", "TestMultiDimensionalStructsPtr3.Errs[0][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr3.Errs[1][1].Name", "TestMultiDimensionalStructsPtr3.Errs[1][1].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr3.Errs[1][2].Name", "TestMultiDimensionalStructsPtr3.Errs[1][2].Name", "Name", "Name", "required")
	AssertError(t, errs, "TestMultiDimensionalStructsPtr3.Errs[2][1].Name", "TestMultiDimensionalStructsPtr3.Errs[2][1].Name", "Name", "Name", "required")

	type TestMultiDimensionalTimeTime struct {
		Errs [][]*time.Time `validate:"gt=0,dive,dive,required"`
	}

	var errTimePtr3Array [][]*time.Time

	t1 := time.Now().UTC()
	t2 := time.Now().UTC()
	t3 := time.Now().UTC().Add(time.Hour * 24)

	errTimePtr3Array = append(errTimePtr3Array, []*time.Time{&t1, &t2, &t3})
	errTimePtr3Array = append(errTimePtr3Array, []*time.Time{&t1, &t2, nil})
	errTimePtr3Array = append(errTimePtr3Array, []*time.Time{&t1, nil, nil})

	tmtp3 := &TestMultiDimensionalTimeTime{
		Errs: errTimePtr3Array,
	}

	errs = validate.Struct(tmtp3)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 3)
	AssertError(t, errs, "TestMultiDimensionalTimeTime.Errs[1][2]", "TestMultiDimensionalTimeTime.Errs[1][2]", "Errs[1][2]", "Errs[1][2]", "required")
	AssertError(t, errs, "TestMultiDimensionalTimeTime.Errs[2][1]", "TestMultiDimensionalTimeTime.Errs[2][1]", "Errs[2][1]", "Errs[2][1]", "required")
	AssertError(t, errs, "TestMultiDimensionalTimeTime.Errs[2][2]", "TestMultiDimensionalTimeTime.Errs[2][2]", "Errs[2][2]", "Errs[2][2]", "required")

	type TestMultiDimensionalTimeTime2 struct {
		Errs [][]*time.Time `validate:"gt=0,dive,dive,required"`
	}

	var errTimeArray [][]*time.Time

	t1 = time.Now().UTC()
	t2 = time.Now().UTC()
	t3 = time.Now().UTC().Add(time.Hour * 24)

	errTimeArray = append(errTimeArray, []*time.Time{&t1, &t2, &t3})
	errTimeArray = append(errTimeArray, []*time.Time{&t1, &t2, nil})
	errTimeArray = append(errTimeArray, []*time.Time{&t1, nil, nil})

	tmtp := &TestMultiDimensionalTimeTime2{
		Errs: errTimeArray,
	}

	errs = validate.Struct(tmtp)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 3)
	AssertError(t, errs, "TestMultiDimensionalTimeTime2.Errs[1][2]", "TestMultiDimensionalTimeTime2.Errs[1][2]", "Errs[1][2]", "Errs[1][2]", "required")
	AssertError(t, errs, "TestMultiDimensionalTimeTime2.Errs[2][1]", "TestMultiDimensionalTimeTime2.Errs[2][1]", "Errs[2][1]", "Errs[2][1]", "required")
	AssertError(t, errs, "TestMultiDimensionalTimeTime2.Errs[2][2]", "TestMultiDimensionalTimeTime2.Errs[2][2]", "Errs[2][2]", "Errs[2][2]", "required")
}

func TestNilStructPointerValidation(t *testing.T) {
	type Inner struct {
		Data string
	}

	type Outer struct {
		Inner *Inner `validate:"omitempty"`
	}

	inner := &Inner{
		Data: "test",
	}

	outer := &Outer{
		Inner: inner,
	}

	validate := New()
	errs := validate.Struct(outer)
	Equal(t, errs, nil)

	outer = &Outer{
		Inner: nil,
	}

	errs = validate.Struct(outer)
	Equal(t, errs, nil)

	type Inner2 struct {
		Data string
	}

	type Outer2 struct {
		Inner2 *Inner2 `validate:"required"`
	}

	inner2 := &Inner2{
		Data: "test",
	}

	outer2 := &Outer2{
		Inner2: inner2,
	}

	errs = validate.Struct(outer2)
	Equal(t, errs, nil)

	outer2 = &Outer2{
		Inner2: nil,
	}

	errs = validate.Struct(outer2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Outer2.Inner2", "Outer2.Inner2", "Inner2", "Inner2", "required")

	type Inner3 struct {
		Data string
	}

	type Outer3 struct {
		Inner3 *Inner3
	}

	inner3 := &Inner3{
		Data: "test",
	}

	outer3 := &Outer3{
		Inner3: inner3,
	}

	errs = validate.Struct(outer3)
	Equal(t, errs, nil)

	type Inner4 struct {
		Data string
	}

	type Outer4 struct {
		Inner4 *Inner4 `validate:"-"`
	}

	inner4 := &Inner4{
		Data: "test",
	}

	outer4 := &Outer4{
		Inner4: inner4,
	}

	errs = validate.Struct(outer4)
	Equal(t, errs, nil)
}

func TestSSNValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"00-90-8787", false},
		{"66690-76", false},
		{"191 60 2869", true},
		{"191-60-2869", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "ssn")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d SSN failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d SSN failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ssn" {
					t.Fatalf("Index: %d Latitude failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestLongitudeValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"-180.000", true},
		{"180.1", false},
		{"+73.234", true},
		{"+382.3811", false},
		{"23.11111111", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "longitude")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d Longitude failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d Longitude failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "longitude" {
					t.Fatalf("Index: %d Latitude failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestLatitudeValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"-90.000", true},
		{"+90", true},
		{"47.1231231", true},
		{"+99.9", false},
		{"108", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "latitude")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d Latitude failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d Latitude failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "latitude" {
					t.Fatalf("Index: %d Latitude failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestDataURIValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"data:text/plain;base64,Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
		{"data:image/gif;base64,MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"data:image/png;base64,12345", false},
		{"", false},
		{"data:text,:;base85,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "datauri")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d DataURI failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d DataURI failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "datauri" {
					t.Fatalf("Index: %d DataURI failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestMultibyteValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc", false},
		{"123", false},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"あいうえお foobar", true},
		{"test＠example.com", true},
		{"test＠example.com", true},
		{"1234abcDEｘｙｚ", true},
		{"ｶﾀｶﾅ", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "multibyte")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d Multibyte failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d Multibyte failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "multibyte" {
					t.Fatalf("Index: %d Multibyte failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestPrintableASCIIValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"newline\n", false},
		{"\x19test\x7F", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "printascii")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d Printable ASCII failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d Printable ASCII failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "printascii" {
					t.Fatalf("Index: %d Printable ASCII failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestASCIIValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "ascii")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ASCII failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ASCII failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "ascii" {
					t.Fatalf("Index: %d ASCII failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUUID5Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{

		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"9c858901-8a57-4791-81fe-4c455b099bc9", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"987fbc97-4bed-5078-af07-9141ba07c9f3", true},
		{"987fbc97-4bed-5078-9f07-9141ba07c9f3", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "uuid5")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID5 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID5 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "uuid5" {
					t.Fatalf("Index: %d UUID5 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUUID4Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-5078-af07-9141ba07c9f3", false},
		{"934859", false},
		{"57b73598-8764-4ad0-a76a-679bb6640eb1", true},
		{"625e63f3-58f5-40b7-83a1-a72ad31acffb", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "uuid4")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID4 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID4 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "uuid4" {
					t.Fatalf("Index: %d UUID4 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUUID3Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"412452646", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-4078-8f07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "uuid3")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID3 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID3 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "uuid3" {
					t.Fatalf("Index: %d UUID3 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestUUIDValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3xxx", false},
		{"a987fbc94bed3078cf079141ba07c9f3", false},
		{"934859", false},
		{"987fbc9-4bed-3078-cf07a-9141ba07c9f3", false},
		{"aaaaaaaa-1111-1111-aaag-111111111111", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "uuid")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d UUID failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "uuid" {
					t.Fatalf("Index: %d UUID failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestISBNValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3836221195", true},
		{"1-61729-085-8", true},
		{"3 423 21412 0", true},
		{"3 401 01319 X", true},
		{"9784873113685", true},
		{"978-4-87311-368-5", true},
		{"978 3401013190", true},
		{"978-3-8362-2119-1", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "isbn")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ISBN failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ISBN failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "isbn" {
					t.Fatalf("Index: %d ISBN failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestISBN13Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3-8362-2119-5", false},
		{"01234567890ab", false},
		{"978 3 8362 2119 0", false},
		{"9784873113685", true},
		{"978-4-87311-368-5", true},
		{"978 3401013190", true},
		{"978-3-8362-2119-1", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "isbn13")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ISBN13 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ISBN13 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "isbn13" {
					t.Fatalf("Index: %d ISBN13 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestISBN10Validation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3423214121", false},
		{"978-3836221191", false},
		{"3-423-21412-1", false},
		{"3 423 21412 1", false},
		{"3836221195", true},
		{"1-61729-085-8", true},
		{"3 423 21412 0", true},
		{"3 401 01319 X", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "isbn10")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d ISBN10 failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d ISBN10 failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "isbn10" {
					t.Fatalf("Index: %d ISBN10 failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestExcludesRuneValidation(t *testing.T) {

	tests := []struct {
		Value       string `validate:"excludesrune=☻"`
		Tag         string
		ExpectedNil bool
	}{
		{Value: "a☺b☻c☹d", Tag: "excludesrune=☻", ExpectedNil: false},
		{Value: "abcd", Tag: "excludesrune=☻", ExpectedNil: true},
	}

	validate := New()

	for i, s := range tests {
		errs := validate.Var(s.Value, s.Tag)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}

		errs = validate.Struct(s)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}
	}
}

func TestExcludesAllValidation(t *testing.T) {

	tests := []struct {
		Value       string `validate:"excludesall=@!{}[]"`
		Tag         string
		ExpectedNil bool
	}{
		{Value: "abcd@!jfk", Tag: "excludesall=@!{}[]", ExpectedNil: false},
		{Value: "abcdefg", Tag: "excludesall=@!{}[]", ExpectedNil: true},
	}

	validate := New()

	for i, s := range tests {
		errs := validate.Var(s.Value, s.Tag)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}

		errs = validate.Struct(s)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}
	}

	username := "joeybloggs "

	errs := validate.Var(username, "excludesall=@ ")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "excludesall")

	excluded := ","

	errs = validate.Var(excluded, "excludesall=!@#$%^&*()_+.0x2C?")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "excludesall")

	excluded = "="

	errs = validate.Var(excluded, "excludesall=!@#$%^&*()_+.0x2C=?")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "excludesall")
}

func TestExcludesValidation(t *testing.T) {

	tests := []struct {
		Value       string `validate:"excludes=@"`
		Tag         string
		ExpectedNil bool
	}{
		{Value: "abcd@!jfk", Tag: "excludes=@", ExpectedNil: false},
		{Value: "abcdq!jfk", Tag: "excludes=@", ExpectedNil: true},
	}

	validate := New()

	for i, s := range tests {
		errs := validate.Var(s.Value, s.Tag)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}

		errs = validate.Struct(s)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}
	}
}

func TestContainsRuneValidation(t *testing.T) {

	tests := []struct {
		Value       string `validate:"containsrune=☻"`
		Tag         string
		ExpectedNil bool
	}{
		{Value: "a☺b☻c☹d", Tag: "containsrune=☻", ExpectedNil: true},
		{Value: "abcd", Tag: "containsrune=☻", ExpectedNil: false},
	}

	validate := New()

	for i, s := range tests {
		errs := validate.Var(s.Value, s.Tag)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}

		errs = validate.Struct(s)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}
	}
}

func TestContainsAnyValidation(t *testing.T) {

	tests := []struct {
		Value       string `validate:"containsany=@!{}[]"`
		Tag         string
		ExpectedNil bool
	}{
		{Value: "abcd@!jfk", Tag: "containsany=@!{}[]", ExpectedNil: true},
		{Value: "abcdefg", Tag: "containsany=@!{}[]", ExpectedNil: false},
	}

	validate := New()

	for i, s := range tests {
		errs := validate.Var(s.Value, s.Tag)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}

		errs = validate.Struct(s)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}
	}
}

func TestContainsValidation(t *testing.T) {

	tests := []struct {
		Value       string `validate:"contains=@"`
		Tag         string
		ExpectedNil bool
	}{
		{Value: "abcd@!jfk", Tag: "contains=@", ExpectedNil: true},
		{Value: "abcdq!jfk", Tag: "contains=@", ExpectedNil: false},
	}

	validate := New()

	for i, s := range tests {
		errs := validate.Var(s.Value, s.Tag)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}

		errs = validate.Struct(s)

		if (s.ExpectedNil && errs != nil) || (!s.ExpectedNil && errs == nil) {
			t.Fatalf("Index: %d failed Error: %s", i, errs)
		}
	}
}

func TestIsNeFieldValidation(t *testing.T) {

	validate := New()

	var j uint64
	var k float64
	s := "abcd"
	i := 1
	j = 1
	k = 1.543
	arr := []string{"test"}
	now := time.Now().UTC()

	var j2 uint64
	var k2 float64
	s2 := "abcdef"
	i2 := 3
	j2 = 2
	k2 = 1.5434456
	arr2 := []string{"test", "test2"}
	arr3 := []string{"test"}
	now2 := now

	errs := validate.VarWithValue(s, s2, "nefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(i2, i, "nefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(j2, j, "nefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(k2, k, "nefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(arr2, arr, "nefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(now2, now, "nefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "nefield")

	errs = validate.VarWithValue(arr3, arr, "nefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "nefield")

	type Test struct {
		Start *time.Time `validate:"nefield=End"`
		End   *time.Time
	}

	sv := &Test{
		Start: &now,
		End:   &now,
	}

	errs = validate.Struct(sv)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Start", "Test.Start", "Start", "Start", "nefield")

	now3 := time.Now().UTC()

	sv = &Test{
		Start: &now,
		End:   &now3,
	}

	errs = validate.Struct(sv)
	Equal(t, errs, nil)

	errs = validate.VarWithValue(nil, 1, "nefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "nefield")

	errs = validate.VarWithValue(sv, now, "nefield")
	Equal(t, errs, nil)

	type Test2 struct {
		Start *time.Time `validate:"nefield=NonExistantField"`
		End   *time.Time
	}

	sv2 := &Test2{
		Start: &now,
		End:   &now,
	}

	errs = validate.Struct(sv2)
	Equal(t, errs, nil)

	type Other struct {
		Value string
	}

	type Test3 struct {
		Value Other
		Time  time.Time `validate:"nefield=Value"`
	}

	tst := Test3{
		Value: Other{Value: "StringVal"},
		Time:  now,
	}

	errs = validate.Struct(tst)
	Equal(t, errs, nil)
}

func TestIsNeValidation(t *testing.T) {

	validate := New()

	var j uint64
	var k float64
	s := "abcdef"
	i := 3
	j = 2
	k = 1.5434
	arr := []string{"test"}
	now := time.Now().UTC()

	errs := validate.Var(s, "ne=abcd")
	Equal(t, errs, nil)

	errs = validate.Var(i, "ne=1")
	Equal(t, errs, nil)

	errs = validate.Var(j, "ne=1")
	Equal(t, errs, nil)

	errs = validate.Var(k, "ne=1.543")
	Equal(t, errs, nil)

	errs = validate.Var(arr, "ne=2")
	Equal(t, errs, nil)

	errs = validate.Var(arr, "ne=1")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ne")

	PanicMatches(t, func() { validate.Var(now, "ne=now") }, "Bad field type time.Time")
}

func TestIsEqFieldValidation(t *testing.T) {

	validate := New()

	var j uint64
	var k float64
	s := "abcd"
	i := 1
	j = 1
	k = 1.543
	arr := []string{"test"}
	now := time.Now().UTC()

	var j2 uint64
	var k2 float64
	s2 := "abcd"
	i2 := 1
	j2 = 1
	k2 = 1.543
	arr2 := []string{"test"}
	arr3 := []string{"test", "test2"}
	now2 := now

	errs := validate.VarWithValue(s, s2, "eqfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(i2, i, "eqfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(j2, j, "eqfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(k2, k, "eqfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(arr2, arr, "eqfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(now2, now, "eqfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(arr3, arr, "eqfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "eqfield")

	type Test struct {
		Start *time.Time `validate:"eqfield=End"`
		End   *time.Time
	}

	sv := &Test{
		Start: &now,
		End:   &now,
	}

	errs = validate.Struct(sv)
	Equal(t, errs, nil)

	now3 := time.Now().UTC()

	sv = &Test{
		Start: &now,
		End:   &now3,
	}

	errs = validate.Struct(sv)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Start", "Test.Start", "Start", "Start", "eqfield")

	errs = validate.VarWithValue(nil, 1, "eqfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "eqfield")

	channel := make(chan string)
	errs = validate.VarWithValue(5, channel, "eqfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "eqfield")

	errs = validate.VarWithValue(5, now, "eqfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "eqfield")

	type Test2 struct {
		Start *time.Time `validate:"eqfield=NonExistantField"`
		End   *time.Time
	}

	sv2 := &Test2{
		Start: &now,
		End:   &now,
	}

	errs = validate.Struct(sv2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test2.Start", "Test2.Start", "Start", "Start", "eqfield")

	type Inner struct {
		Name string
	}

	type TStruct struct {
		Inner     *Inner
		CreatedAt *time.Time `validate:"eqfield=Inner"`
	}

	inner := &Inner{
		Name: "NAME",
	}

	test := &TStruct{
		Inner:     inner,
		CreatedAt: &now,
	}

	errs = validate.Struct(test)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TStruct.CreatedAt", "TStruct.CreatedAt", "CreatedAt", "CreatedAt", "eqfield")
}

func TestIsEqValidation(t *testing.T) {

	validate := New()

	var j uint64
	var k float64
	s := "abcd"
	i := 1
	j = 1
	k = 1.543
	arr := []string{"test"}
	now := time.Now().UTC()

	errs := validate.Var(s, "eq=abcd")
	Equal(t, errs, nil)

	errs = validate.Var(i, "eq=1")
	Equal(t, errs, nil)

	errs = validate.Var(j, "eq=1")
	Equal(t, errs, nil)

	errs = validate.Var(k, "eq=1.543")
	Equal(t, errs, nil)

	errs = validate.Var(arr, "eq=1")
	Equal(t, errs, nil)

	errs = validate.Var(arr, "eq=2")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "eq")

	PanicMatches(t, func() { validate.Var(now, "eq=now") }, "Bad field type time.Time")
}

func TestBase64Validation(t *testing.T) {

	validate := New()

	s := "dW5pY29ybg=="

	errs := validate.Var(s, "base64")
	Equal(t, errs, nil)

	s = "dGhpIGlzIGEgdGVzdCBiYXNlNjQ="
	errs = validate.Var(s, "base64")
	Equal(t, errs, nil)

	s = ""
	errs = validate.Var(s, "base64")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "base64")

	s = "dW5pY29ybg== foo bar"
	errs = validate.Var(s, "base64")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "base64")
}

func TestNoStructLevelValidation(t *testing.T) {

	type Inner struct {
		Test string `validate:"len=5"`
	}

	type Outer struct {
		InnerStruct *Inner `validate:"required,nostructlevel"`
	}

	outer := &Outer{
		InnerStruct: nil,
	}

	validate := New()

	errs := validate.Struct(outer)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Outer.InnerStruct", "Outer.InnerStruct", "InnerStruct", "InnerStruct", "required")

	inner := &Inner{
		Test: "1234",
	}

	outer = &Outer{
		InnerStruct: inner,
	}

	errs = validate.Struct(outer)
	Equal(t, errs, nil)
}

func TestStructOnlyValidation(t *testing.T) {

	type Inner struct {
		Test string `validate:"len=5"`
	}

	type Outer struct {
		InnerStruct *Inner `validate:"required,structonly"`
	}

	outer := &Outer{
		InnerStruct: nil,
	}

	validate := New()

	errs := validate.Struct(outer)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Outer.InnerStruct", "Outer.InnerStruct", "InnerStruct", "InnerStruct", "required")

	inner := &Inner{
		Test: "1234",
	}

	outer = &Outer{
		InnerStruct: inner,
	}

	errs = validate.Struct(outer)
	Equal(t, errs, nil)

	// Address houses a users address information
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
		Planet string `validate:"required"`
		Phone  string `validate:"required"`
	}

	type User struct {
		FirstName      string     `json:"fname"`
		LastName       string     `json:"lname"`
		Age            uint8      `validate:"gte=0,lte=130"`
		Email          string     `validate:"required,email"`
		FavouriteColor string     `validate:"hexcolor|rgb|rgba"`
		Addresses      []*Address `validate:"required"`   // a person can have a home and cottage...
		Address        Address    `validate:"structonly"` // a person can have a home and cottage...
	}

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "Unknown",
	}

	user := &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000",
		Addresses:      []*Address{address},
		Address: Address{
			// Street: "Eavesdown Docks",
			Planet: "Persphone",
			Phone:  "none",
			City:   "Unknown",
		},
	}

	errs = validate.Struct(user)
	Equal(t, errs, nil)
}

func TestGtField(t *testing.T) {

	validate := New()

	type TimeTest struct {
		Start *time.Time `validate:"required,gt"`
		End   *time.Time `validate:"required,gt,gtfield=Start"`
	}

	now := time.Now()
	start := now.Add(time.Hour * 24)
	end := start.Add(time.Hour * 24)

	timeTest := &TimeTest{
		Start: &start,
		End:   &end,
	}

	errs := validate.Struct(timeTest)
	Equal(t, errs, nil)

	timeTest = &TimeTest{
		Start: &end,
		End:   &start,
	}

	errs = validate.Struct(timeTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest.End", "TimeTest.End", "End", "End", "gtfield")

	errs = validate.VarWithValue(&end, &start, "gtfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(&start, &end, "gtfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtfield")

	errs = validate.VarWithValue(&end, &start, "gtfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(&timeTest, &end, "gtfield")
	NotEqual(t, errs, nil)

	errs = validate.VarWithValue("test bigger", "test", "gtfield")
	Equal(t, errs, nil)

	type IntTest struct {
		Val1 int `validate:"required"`
		Val2 int `validate:"required,gtfield=Val1"`
	}

	intTest := &IntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(intTest)
	Equal(t, errs, nil)

	intTest = &IntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(intTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "IntTest.Val2", "IntTest.Val2", "Val2", "Val2", "gtfield")

	errs = validate.VarWithValue(int(5), int(1), "gtfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(int(1), int(5), "gtfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtfield")

	type UIntTest struct {
		Val1 uint `validate:"required"`
		Val2 uint `validate:"required,gtfield=Val1"`
	}

	uIntTest := &UIntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(uIntTest)
	Equal(t, errs, nil)

	uIntTest = &UIntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(uIntTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "UIntTest.Val2", "UIntTest.Val2", "Val2", "Val2", "gtfield")

	errs = validate.VarWithValue(uint(5), uint(1), "gtfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(uint(1), uint(5), "gtfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtfield")

	type FloatTest struct {
		Val1 float64 `validate:"required"`
		Val2 float64 `validate:"required,gtfield=Val1"`
	}

	floatTest := &FloatTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(floatTest)
	Equal(t, errs, nil)

	floatTest = &FloatTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(floatTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "FloatTest.Val2", "FloatTest.Val2", "Val2", "Val2", "gtfield")

	errs = validate.VarWithValue(float32(5), float32(1), "gtfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(float32(1), float32(5), "gtfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtfield")

	errs = validate.VarWithValue(nil, 1, "gtfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtfield")

	errs = validate.VarWithValue(5, "T", "gtfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtfield")

	errs = validate.VarWithValue(5, start, "gtfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtfield")

	type TimeTest2 struct {
		Start *time.Time `validate:"required"`
		End   *time.Time `validate:"required,gtfield=NonExistantField"`
	}

	timeTest2 := &TimeTest2{
		Start: &start,
		End:   &end,
	}

	errs = validate.Struct(timeTest2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest2.End", "TimeTest2.End", "End", "End", "gtfield")

	type Other struct {
		Value string
	}

	type Test struct {
		Value Other
		Time  time.Time `validate:"gtfield=Value"`
	}

	tst := Test{
		Value: Other{Value: "StringVal"},
		Time:  end,
	}

	errs = validate.Struct(tst)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Time", "Test.Time", "Time", "Time", "gtfield")
}

func TestLtField(t *testing.T) {

	validate := New()

	type TimeTest struct {
		Start *time.Time `validate:"required,lt,ltfield=End"`
		End   *time.Time `validate:"required,lt"`
	}

	now := time.Now()
	start := now.Add(time.Hour * 24 * -1 * 2)
	end := start.Add(time.Hour * 24)

	timeTest := &TimeTest{
		Start: &start,
		End:   &end,
	}

	errs := validate.Struct(timeTest)
	Equal(t, errs, nil)

	timeTest = &TimeTest{
		Start: &end,
		End:   &start,
	}

	errs = validate.Struct(timeTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest.Start", "TimeTest.Start", "Start", "Start", "ltfield")

	errs = validate.VarWithValue(&start, &end, "ltfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(&end, &start, "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	errs = validate.VarWithValue(&end, timeTest, "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	errs = validate.VarWithValue("tes", "test", "ltfield")
	Equal(t, errs, nil)

	type IntTest struct {
		Val1 int `validate:"required"`
		Val2 int `validate:"required,ltfield=Val1"`
	}

	intTest := &IntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(intTest)
	Equal(t, errs, nil)

	intTest = &IntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(intTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "IntTest.Val2", "IntTest.Val2", "Val2", "Val2", "ltfield")

	errs = validate.VarWithValue(int(1), int(5), "ltfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(int(5), int(1), "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	type UIntTest struct {
		Val1 uint `validate:"required"`
		Val2 uint `validate:"required,ltfield=Val1"`
	}

	uIntTest := &UIntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(uIntTest)
	Equal(t, errs, nil)

	uIntTest = &UIntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(uIntTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "UIntTest.Val2", "UIntTest.Val2", "Val2", "Val2", "ltfield")

	errs = validate.VarWithValue(uint(1), uint(5), "ltfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(uint(5), uint(1), "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	type FloatTest struct {
		Val1 float64 `validate:"required"`
		Val2 float64 `validate:"required,ltfield=Val1"`
	}

	floatTest := &FloatTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(floatTest)
	Equal(t, errs, nil)

	floatTest = &FloatTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(floatTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "FloatTest.Val2", "FloatTest.Val2", "Val2", "Val2", "ltfield")

	errs = validate.VarWithValue(float32(1), float32(5), "ltfield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(float32(5), float32(1), "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	errs = validate.VarWithValue(nil, 5, "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	errs = validate.VarWithValue(1, "T", "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	errs = validate.VarWithValue(1, end, "ltfield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltfield")

	type TimeTest2 struct {
		Start *time.Time `validate:"required"`
		End   *time.Time `validate:"required,ltfield=NonExistantField"`
	}

	timeTest2 := &TimeTest2{
		Start: &end,
		End:   &start,
	}

	errs = validate.Struct(timeTest2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest2.End", "TimeTest2.End", "End", "End", "ltfield")
}

func TestLteField(t *testing.T) {

	validate := New()

	type TimeTest struct {
		Start *time.Time `validate:"required,lte,ltefield=End"`
		End   *time.Time `validate:"required,lte"`
	}

	now := time.Now()
	start := now.Add(time.Hour * 24 * -1 * 2)
	end := start.Add(time.Hour * 24)

	timeTest := &TimeTest{
		Start: &start,
		End:   &end,
	}

	errs := validate.Struct(timeTest)
	Equal(t, errs, nil)

	timeTest = &TimeTest{
		Start: &end,
		End:   &start,
	}

	errs = validate.Struct(timeTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest.Start", "TimeTest.Start", "Start", "Start", "ltefield")

	errs = validate.VarWithValue(&start, &end, "ltefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(&end, &start, "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	errs = validate.VarWithValue(&end, timeTest, "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	errs = validate.VarWithValue("tes", "test", "ltefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue("test", "test", "ltefield")
	Equal(t, errs, nil)

	type IntTest struct {
		Val1 int `validate:"required"`
		Val2 int `validate:"required,ltefield=Val1"`
	}

	intTest := &IntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(intTest)
	Equal(t, errs, nil)

	intTest = &IntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(intTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "IntTest.Val2", "IntTest.Val2", "Val2", "Val2", "ltefield")

	errs = validate.VarWithValue(int(1), int(5), "ltefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(int(5), int(1), "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	type UIntTest struct {
		Val1 uint `validate:"required"`
		Val2 uint `validate:"required,ltefield=Val1"`
	}

	uIntTest := &UIntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(uIntTest)
	Equal(t, errs, nil)

	uIntTest = &UIntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(uIntTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "UIntTest.Val2", "UIntTest.Val2", "Val2", "Val2", "ltefield")

	errs = validate.VarWithValue(uint(1), uint(5), "ltefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(uint(5), uint(1), "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	type FloatTest struct {
		Val1 float64 `validate:"required"`
		Val2 float64 `validate:"required,ltefield=Val1"`
	}

	floatTest := &FloatTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(floatTest)
	Equal(t, errs, nil)

	floatTest = &FloatTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(floatTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "FloatTest.Val2", "FloatTest.Val2", "Val2", "Val2", "ltefield")

	errs = validate.VarWithValue(float32(1), float32(5), "ltefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(float32(5), float32(1), "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	errs = validate.VarWithValue(nil, 5, "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	errs = validate.VarWithValue(1, "T", "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	errs = validate.VarWithValue(1, end, "ltefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "ltefield")

	type TimeTest2 struct {
		Start *time.Time `validate:"required"`
		End   *time.Time `validate:"required,ltefield=NonExistantField"`
	}

	timeTest2 := &TimeTest2{
		Start: &end,
		End:   &start,
	}

	errs = validate.Struct(timeTest2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest2.End", "TimeTest2.End", "End", "End", "ltefield")
}

func TestGteField(t *testing.T) {

	validate := New()

	type TimeTest struct {
		Start *time.Time `validate:"required,gte"`
		End   *time.Time `validate:"required,gte,gtefield=Start"`
	}

	now := time.Now()
	start := now.Add(time.Hour * 24)
	end := start.Add(time.Hour * 24)

	timeTest := &TimeTest{
		Start: &start,
		End:   &end,
	}

	errs := validate.Struct(timeTest)
	Equal(t, errs, nil)

	timeTest = &TimeTest{
		Start: &end,
		End:   &start,
	}

	errs = validate.Struct(timeTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest.End", "TimeTest.End", "End", "End", "gtefield")

	errs = validate.VarWithValue(&end, &start, "gtefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(&start, &end, "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	errs = validate.VarWithValue(&start, timeTest, "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	errs = validate.VarWithValue("test", "test", "gtefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue("test bigger", "test", "gtefield")
	Equal(t, errs, nil)

	type IntTest struct {
		Val1 int `validate:"required"`
		Val2 int `validate:"required,gtefield=Val1"`
	}

	intTest := &IntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(intTest)
	Equal(t, errs, nil)

	intTest = &IntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(intTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "IntTest.Val2", "IntTest.Val2", "Val2", "Val2", "gtefield")

	errs = validate.VarWithValue(int(5), int(1), "gtefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(int(1), int(5), "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	type UIntTest struct {
		Val1 uint `validate:"required"`
		Val2 uint `validate:"required,gtefield=Val1"`
	}

	uIntTest := &UIntTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(uIntTest)
	Equal(t, errs, nil)

	uIntTest = &UIntTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(uIntTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "UIntTest.Val2", "UIntTest.Val2", "Val2", "Val2", "gtefield")

	errs = validate.VarWithValue(uint(5), uint(1), "gtefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(uint(1), uint(5), "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	type FloatTest struct {
		Val1 float64 `validate:"required"`
		Val2 float64 `validate:"required,gtefield=Val1"`
	}

	floatTest := &FloatTest{
		Val1: 1,
		Val2: 5,
	}

	errs = validate.Struct(floatTest)
	Equal(t, errs, nil)

	floatTest = &FloatTest{
		Val1: 5,
		Val2: 1,
	}

	errs = validate.Struct(floatTest)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "FloatTest.Val2", "FloatTest.Val2", "Val2", "Val2", "gtefield")

	errs = validate.VarWithValue(float32(5), float32(1), "gtefield")
	Equal(t, errs, nil)

	errs = validate.VarWithValue(float32(1), float32(5), "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	errs = validate.VarWithValue(nil, 1, "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	errs = validate.VarWithValue(5, "T", "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	errs = validate.VarWithValue(5, start, "gtefield")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gtefield")

	type TimeTest2 struct {
		Start *time.Time `validate:"required"`
		End   *time.Time `validate:"required,gtefield=NonExistantField"`
	}

	timeTest2 := &TimeTest2{
		Start: &start,
		End:   &end,
	}

	errs = validate.Struct(timeTest2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TimeTest2.End", "TimeTest2.End", "End", "End", "gtefield")
}

func TestValidateByTagAndValue(t *testing.T) {

	validate := New()

	val := "test"
	field := "test"
	errs := validate.VarWithValue(val, field, "required")
	Equal(t, errs, nil)

	fn := func(fl FieldLevel) bool {

		return fl.Parent().String() == fl.Field().String()
	}

	validate.RegisterValidation("isequaltestfunc", fn)

	errs = validate.VarWithValue(val, field, "isequaltestfunc")
	Equal(t, errs, nil)

	val = "unequal"

	errs = validate.VarWithValue(val, field, "isequaltestfunc")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "isequaltestfunc")
}

func TestAddFunctions(t *testing.T) {

	fn := func(fl FieldLevel) bool {

		return true
	}

	fnCtx := func(ctx context.Context, fl FieldLevel) bool {
		return true
	}

	validate := New()

	errs := validate.RegisterValidation("new", fn)
	Equal(t, errs, nil)

	errs = validate.RegisterValidation("", fn)
	NotEqual(t, errs, nil)

	validate.RegisterValidation("new", nil)
	NotEqual(t, errs, nil)

	errs = validate.RegisterValidation("new", fn)
	Equal(t, errs, nil)

	errs = validate.RegisterValidationCtx("new", fnCtx)
	Equal(t, errs, nil)

	PanicMatches(t, func() { validate.RegisterValidation("dive", fn) }, "Tag 'dive' either contains restricted characters or is the same as a restricted tag needed for normal operation")
}

func TestChangeTag(t *testing.T) {

	validate := New()
	validate.SetTagName("val")

	type Test struct {
		Name string `val:"len=4"`
	}
	s := &Test{
		Name: "TEST",
	}

	errs := validate.Struct(s)
	Equal(t, errs, nil)

	s.Name = ""

	errs = validate.Struct(s)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Name", "Test.Name", "Name", "Name", "len")
}

func TestUnexposedStruct(t *testing.T) {

	validate := New()

	type Test struct {
		Name      string
		unexposed struct {
			A string `validate:"required"`
		}
	}

	s := &Test{
		Name: "TEST",
	}

	errs := validate.Struct(s)
	Equal(t, errs, nil)
}

func TestBadParams(t *testing.T) {

	validate := New()

	i := 1
	errs := validate.Var(i, "-")
	Equal(t, errs, nil)

	PanicMatches(t, func() { validate.Var(i, "len=a") }, "strconv.ParseInt: parsing \"a\": invalid syntax")
	PanicMatches(t, func() { validate.Var(i, "len=a") }, "strconv.ParseInt: parsing \"a\": invalid syntax")

	var ui uint = 1
	PanicMatches(t, func() { validate.Var(ui, "len=a") }, "strconv.ParseUint: parsing \"a\": invalid syntax")

	f := 1.23
	PanicMatches(t, func() { validate.Var(f, "len=a") }, "strconv.ParseFloat: parsing \"a\": invalid syntax")
}

func TestLength(t *testing.T) {

	validate := New()

	i := true
	PanicMatches(t, func() { validate.Var(i, "len") }, "Bad field type bool")
}

func TestIsGt(t *testing.T) {

	validate := New()

	myMap := map[string]string{}
	errs := validate.Var(myMap, "gt=0")
	NotEqual(t, errs, nil)

	f := 1.23
	errs = validate.Var(f, "gt=5")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gt")

	var ui uint = 5
	errs = validate.Var(ui, "gt=10")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gt")

	i := true
	PanicMatches(t, func() { validate.Var(i, "gt") }, "Bad field type bool")

	tm := time.Now().UTC()
	tm = tm.Add(time.Hour * 24)

	errs = validate.Var(tm, "gt")
	Equal(t, errs, nil)

	t2 := time.Now().UTC().Add(-time.Hour)

	errs = validate.Var(t2, "gt")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gt")

	type Test struct {
		Now *time.Time `validate:"gt"`
	}
	s := &Test{
		Now: &tm,
	}

	errs = validate.Struct(s)
	Equal(t, errs, nil)

	s = &Test{
		Now: &t2,
	}

	errs = validate.Struct(s)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Now", "Test.Now", "Now", "Now", "gt")
}

func TestIsGte(t *testing.T) {

	validate := New()

	i := true
	PanicMatches(t, func() { validate.Var(i, "gte") }, "Bad field type bool")

	t1 := time.Now().UTC()
	t1 = t1.Add(time.Hour * 24)

	errs := validate.Var(t1, "gte")
	Equal(t, errs, nil)

	t2 := time.Now().UTC().Add(-time.Hour)

	errs = validate.Var(t2, "gte")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "gte")

	type Test struct {
		Now *time.Time `validate:"gte"`
	}
	s := &Test{
		Now: &t1,
	}

	errs = validate.Struct(s)
	Equal(t, errs, nil)

	s = &Test{
		Now: &t2,
	}

	errs = validate.Struct(s)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Now", "Test.Now", "Now", "Now", "gte")
}

func TestIsLt(t *testing.T) {

	validate := New()

	myMap := map[string]string{}
	errs := validate.Var(myMap, "lt=0")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "lt")

	f := 1.23
	errs = validate.Var(f, "lt=0")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "lt")

	var ui uint = 5
	errs = validate.Var(ui, "lt=0")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "lt")

	i := true
	PanicMatches(t, func() { validate.Var(i, "lt") }, "Bad field type bool")

	t1 := time.Now().UTC().Add(-time.Hour)

	errs = validate.Var(t1, "lt")
	Equal(t, errs, nil)

	t2 := time.Now().UTC()
	t2 = t2.Add(time.Hour * 24)

	errs = validate.Var(t2, "lt")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "lt")

	type Test struct {
		Now *time.Time `validate:"lt"`
	}

	s := &Test{
		Now: &t1,
	}

	errs = validate.Struct(s)
	Equal(t, errs, nil)

	s = &Test{
		Now: &t2,
	}

	errs = validate.Struct(s)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.Now", "Test.Now", "Now", "Now", "lt")
}

func TestIsLte(t *testing.T) {

	validate := New()

	i := true
	PanicMatches(t, func() { validate.Var(i, "lte") }, "Bad field type bool")

	t1 := time.Now().UTC().Add(-time.Hour)

	errs := validate.Var(t1, "lte")
	Equal(t, errs, nil)

	t2 := time.Now().UTC()
	t2 = t2.Add(time.Hour * 24)

	errs = validate.Var(t2, "lte")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "lte")

	type Test struct {
		Now *time.Time `validate:"lte"`
	}

	s := &Test{
		Now: &t1,
	}

	errs = validate.Struct(s)
	Equal(t, errs, nil)

	s = &Test{
		Now: &t2,
	}

	errs = validate.Struct(s)
	NotEqual(t, errs, nil)
}

func TestUrl(t *testing.T) {

	var tests = []struct {
		param    string
		expected bool
	}{
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", false},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"", false},
		{"xyz://foobar.com", true},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/#baz", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", true},
		{"irc://#channel@network", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "url")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d URL failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d URL failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "url" {
					t.Fatalf("Index: %d URL failed Error: %s", i, errs)
				}
			}
		}
	}

	i := 1
	PanicMatches(t, func() { validate.Var(i, "url") }, "Bad field type int")
}

func TestUri(t *testing.T) {

	var tests = []struct {
		param    string
		expected bool
	}{
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", false},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"", false},
		{"xyz://foobar.com", true},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", true},
		{"irc://#channel@network", true},
		{"/abs/test/dir", true},
		{"./rel/test/dir", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "uri")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d URI failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d URI failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "uri" {
					t.Fatalf("Index: %d URI failed Error: %s", i, errs)
				}
			}
		}
	}

	i := 1
	PanicMatches(t, func() { validate.Var(i, "uri") }, "Bad field type int")
}

func TestOrTag(t *testing.T) {

	validate := New()

	s := "rgba(0,31,255,0.5)"
	errs := validate.Var(s, "rgb|rgba")
	Equal(t, errs, nil)

	s = "rgba(0,31,255,0.5)"
	errs = validate.Var(s, "rgb|rgba|len=18")
	Equal(t, errs, nil)

	s = "this ain't right"
	errs = validate.Var(s, "rgb|rgba")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgb|rgba")

	s = "this ain't right"
	errs = validate.Var(s, "rgb|rgba|len=10")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgb|rgba|len=10")

	s = "this is right"
	errs = validate.Var(s, "rgb|rgba|len=13")
	Equal(t, errs, nil)

	s = ""
	errs = validate.Var(s, "omitempty,rgb|rgba")
	Equal(t, errs, nil)

	s = "this is right, but a blank or isn't"

	PanicMatches(t, func() { validate.Var(s, "rgb||len=13") }, "Invalid validation tag on field ''")
	PanicMatches(t, func() { validate.Var(s, "rgb|rgbaa|len=13") }, "Undefined validation function 'rgbaa' on field ''")

	v2 := New()
	v2.RegisterTagNameFunc(func(fld reflect.StructField) string {

		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type Colors struct {
		Fav string `validate:"rgb|rgba" json:"fc"`
	}

	c := Colors{Fav: "this ain't right"}

	err := v2.Struct(c)
	NotEqual(t, err, nil)

	errs = err.(ValidationErrors)
	fe := getError(errs, "Colors.fc", "Colors.Fav")
	NotEqual(t, fe, nil)
}

func TestHsla(t *testing.T) {

	validate := New()

	s := "hsla(360,100%,100%,1)"
	errs := validate.Var(s, "hsla")
	Equal(t, errs, nil)

	s = "hsla(360,100%,100%,0.5)"
	errs = validate.Var(s, "hsla")
	Equal(t, errs, nil)

	s = "hsla(0,0%,0%, 0)"
	errs = validate.Var(s, "hsla")
	Equal(t, errs, nil)

	s = "hsl(361,100%,50%,1)"
	errs = validate.Var(s, "hsla")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsla")

	s = "hsl(361,100%,50%)"
	errs = validate.Var(s, "hsla")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsla")

	s = "hsla(361,100%,50%)"
	errs = validate.Var(s, "hsla")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsla")

	s = "hsla(360,101%,50%)"
	errs = validate.Var(s, "hsla")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsla")

	s = "hsla(360,100%,101%)"
	errs = validate.Var(s, "hsla")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsla")

	i := 1
	validate.Var(i, "hsla")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsla")
}

func TestHsl(t *testing.T) {

	validate := New()

	s := "hsl(360,100%,50%)"
	errs := validate.Var(s, "hsl")
	Equal(t, errs, nil)

	s = "hsl(0,0%,0%)"
	errs = validate.Var(s, "hsl")
	Equal(t, errs, nil)

	s = "hsl(361,100%,50%)"
	errs = validate.Var(s, "hsl")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsl")

	s = "hsl(361,101%,50%)"
	errs = validate.Var(s, "hsl")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsl")

	s = "hsl(361,100%,101%)"
	errs = validate.Var(s, "hsl")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsl")

	s = "hsl(-10,100%,100%)"
	errs = validate.Var(s, "hsl")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsl")

	i := 1
	errs = validate.Var(i, "hsl")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hsl")
}

func TestRgba(t *testing.T) {

	validate := New()

	s := "rgba(0,31,255,0.5)"
	errs := validate.Var(s, "rgba")
	Equal(t, errs, nil)

	s = "rgba(0,31,255,0.12)"
	errs = validate.Var(s, "rgba")
	Equal(t, errs, nil)

	s = "rgba(12%,55%,100%,0.12)"
	errs = validate.Var(s, "rgba")
	Equal(t, errs, nil)

	s = "rgba( 0,  31, 255, 0.5)"
	errs = validate.Var(s, "rgba")
	Equal(t, errs, nil)

	s = "rgba(12%,55,100%,0.12)"
	errs = validate.Var(s, "rgba")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgba")

	s = "rgb(0,  31, 255)"
	errs = validate.Var(s, "rgba")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgba")

	s = "rgb(1,349,275,0.5)"
	errs = validate.Var(s, "rgba")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgba")

	s = "rgb(01,31,255,0.5)"
	errs = validate.Var(s, "rgba")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgba")

	i := 1
	errs = validate.Var(i, "rgba")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgba")
}

func TestRgb(t *testing.T) {

	validate := New()

	s := "rgb(0,31,255)"
	errs := validate.Var(s, "rgb")
	Equal(t, errs, nil)

	s = "rgb(0,  31, 255)"
	errs = validate.Var(s, "rgb")
	Equal(t, errs, nil)

	s = "rgb(10%,  50%, 100%)"
	errs = validate.Var(s, "rgb")
	Equal(t, errs, nil)

	s = "rgb(10%,  50%, 55)"
	errs = validate.Var(s, "rgb")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgb")

	s = "rgb(1,349,275)"
	errs = validate.Var(s, "rgb")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgb")

	s = "rgb(01,31,255)"
	errs = validate.Var(s, "rgb")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgb")

	s = "rgba(0,31,255)"
	errs = validate.Var(s, "rgb")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgb")

	i := 1
	errs = validate.Var(i, "rgb")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "rgb")
}

func TestEmail(t *testing.T) {

	validate := New()

	s := "test@mail.com"
	errs := validate.Var(s, "email")
	Equal(t, errs, nil)

	s = "Dörte@Sörensen.example.com"
	errs = validate.Var(s, "email")
	Equal(t, errs, nil)

	s = "θσερ@εχαμπλε.ψομ"
	errs = validate.Var(s, "email")
	Equal(t, errs, nil)

	s = "юзер@екзампл.ком"
	errs = validate.Var(s, "email")
	Equal(t, errs, nil)

	s = "उपयोगकर्ता@उदाहरण.कॉम"
	errs = validate.Var(s, "email")
	Equal(t, errs, nil)

	s = "用户@例子.广告"
	errs = validate.Var(s, "email")
	Equal(t, errs, nil)

	s = ""
	errs = validate.Var(s, "email")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "email")

	s = "test@email"
	errs = validate.Var(s, "email")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "email")

	s = "test@email."
	errs = validate.Var(s, "email")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "email")

	s = "@email.com"
	errs = validate.Var(s, "email")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "email")

	i := true
	errs = validate.Var(i, "email")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "email")
}

func TestHexColor(t *testing.T) {

	validate := New()

	s := "#fff"
	errs := validate.Var(s, "hexcolor")
	Equal(t, errs, nil)

	s = "#c2c2c2"
	errs = validate.Var(s, "hexcolor")
	Equal(t, errs, nil)

	s = "fff"
	errs = validate.Var(s, "hexcolor")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hexcolor")

	s = "fffFF"
	errs = validate.Var(s, "hexcolor")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hexcolor")

	i := true
	errs = validate.Var(i, "hexcolor")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hexcolor")
}

func TestHexadecimal(t *testing.T) {

	validate := New()

	s := "ff0044"
	errs := validate.Var(s, "hexadecimal")
	Equal(t, errs, nil)

	s = "abcdefg"
	errs = validate.Var(s, "hexadecimal")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hexadecimal")

	i := true
	errs = validate.Var(i, "hexadecimal")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "hexadecimal")
}

func TestNumber(t *testing.T) {

	validate := New()

	s := "1"
	errs := validate.Var(s, "number")
	Equal(t, errs, nil)

	s = "+1"
	errs = validate.Var(s, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")

	s = "-1"
	errs = validate.Var(s, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")

	s = "1.12"
	errs = validate.Var(s, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")

	s = "+1.12"
	errs = validate.Var(s, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")

	s = "-1.12"
	errs = validate.Var(s, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")

	s = "1."
	errs = validate.Var(s, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")

	s = "1.o"
	errs = validate.Var(s, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")

	i := 1
	errs = validate.Var(i, "number")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "number")
}

func TestNumeric(t *testing.T) {

	validate := New()

	s := "1"
	errs := validate.Var(s, "numeric")
	Equal(t, errs, nil)

	s = "+1"
	errs = validate.Var(s, "numeric")
	Equal(t, errs, nil)

	s = "-1"
	errs = validate.Var(s, "numeric")
	Equal(t, errs, nil)

	s = "1.12"
	errs = validate.Var(s, "numeric")
	Equal(t, errs, nil)

	s = "+1.12"
	errs = validate.Var(s, "numeric")
	Equal(t, errs, nil)

	s = "-1.12"
	errs = validate.Var(s, "numeric")
	Equal(t, errs, nil)

	s = "1."
	errs = validate.Var(s, "numeric")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "numeric")

	s = "1.o"
	errs = validate.Var(s, "numeric")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "numeric")

	i := 1
	errs = validate.Var(i, "numeric")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "numeric")
}

func TestAlphaNumeric(t *testing.T) {

	validate := New()

	s := "abcd123"
	errs := validate.Var(s, "alphanum")
	Equal(t, errs, nil)

	s = "abc!23"
	errs = validate.Var(s, "alphanum")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "alphanum")

	errs = validate.Var(1, "alphanum")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "alphanum")
}

func TestAlpha(t *testing.T) {

	validate := New()

	s := "abcd"
	errs := validate.Var(s, "alpha")
	Equal(t, errs, nil)

	s = "abc®"
	errs = validate.Var(s, "alpha")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "alpha")

	s = "abc÷"
	errs = validate.Var(s, "alpha")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "alpha")

	s = "abc1"
	errs = validate.Var(s, "alpha")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "alpha")

	s = "this is a test string"
	errs = validate.Var(s, "alpha")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "alpha")

	errs = validate.Var(1, "alpha")
	NotEqual(t, errs, nil)
	AssertError(t, errs, "", "", "", "", "alpha")

}

func TestStructStringValidation(t *testing.T) {

	validate := New()

	tSuccess := &TestString{
		Required:  "Required",
		Len:       "length==10",
		Min:       "min=1",
		Max:       "1234567890",
		MinMax:    "12345",
		Lt:        "012345678",
		Lte:       "0123456789",
		Gt:        "01234567890",
		Gte:       "0123456789",
		OmitEmpty: "",
		Sub: &SubTest{
			Test: "1",
		},
		SubIgnore: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A string `validate:"required"`
		}{
			A: "1",
		},
		Iface: &Impl{
			F: "123",
		},
	}

	errs := validate.Struct(tSuccess)
	Equal(t, errs, nil)

	tFail := &TestString{
		Required:  "",
		Len:       "",
		Min:       "",
		Max:       "12345678901",
		MinMax:    "",
		Lt:        "0123456789",
		Lte:       "01234567890",
		Gt:        "1",
		Gte:       "1",
		OmitEmpty: "12345678901",
		Sub: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A string `validate:"required"`
		}{
			A: "",
		},
		Iface: &Impl{
			F: "12",
		},
	}

	errs = validate.Struct(tFail)

	// Assert Top Level
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 13)

	// Assert Fields
	AssertError(t, errs, "TestString.Required", "TestString.Required", "Required", "Required", "required")
	AssertError(t, errs, "TestString.Len", "TestString.Len", "Len", "Len", "len")
	AssertError(t, errs, "TestString.Min", "TestString.Min", "Min", "Min", "min")
	AssertError(t, errs, "TestString.Max", "TestString.Max", "Max", "Max", "max")
	AssertError(t, errs, "TestString.MinMax", "TestString.MinMax", "MinMax", "MinMax", "min")
	AssertError(t, errs, "TestString.Lt", "TestString.Lt", "Lt", "Lt", "lt")
	AssertError(t, errs, "TestString.Lte", "TestString.Lte", "Lte", "Lte", "lte")
	AssertError(t, errs, "TestString.Gt", "TestString.Gt", "Gt", "Gt", "gt")
	AssertError(t, errs, "TestString.Gte", "TestString.Gte", "Gte", "Gte", "gte")
	AssertError(t, errs, "TestString.OmitEmpty", "TestString.OmitEmpty", "OmitEmpty", "OmitEmpty", "max")

	// Nested Struct Field Errs
	AssertError(t, errs, "TestString.Anonymous.A", "TestString.Anonymous.A", "A", "A", "required")
	AssertError(t, errs, "TestString.Sub.Test", "TestString.Sub.Test", "Test", "Test", "required")
	AssertError(t, errs, "TestString.Iface.F", "TestString.Iface.F", "F", "F", "len")
}

func TestStructInt32Validation(t *testing.T) {

	type TestInt32 struct {
		Required  int `validate:"required"`
		Len       int `validate:"len=10"`
		Min       int `validate:"min=1"`
		Max       int `validate:"max=10"`
		MinMax    int `validate:"min=1,max=10"`
		Lt        int `validate:"lt=10"`
		Lte       int `validate:"lte=10"`
		Gt        int `validate:"gt=10"`
		Gte       int `validate:"gte=10"`
		OmitEmpty int `validate:"omitempty,min=1,max=10"`
	}

	tSuccess := &TestInt32{
		Required:  1,
		Len:       10,
		Min:       1,
		Max:       10,
		MinMax:    5,
		Lt:        9,
		Lte:       10,
		Gt:        11,
		Gte:       10,
		OmitEmpty: 0,
	}

	validate := New()
	errs := validate.Struct(tSuccess)
	Equal(t, errs, nil)

	tFail := &TestInt32{
		Required:  0,
		Len:       11,
		Min:       -1,
		Max:       11,
		MinMax:    -1,
		Lt:        10,
		Lte:       11,
		Gt:        10,
		Gte:       9,
		OmitEmpty: 11,
	}

	errs = validate.Struct(tFail)

	// Assert Top Level
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 10)

	// Assert Fields
	AssertError(t, errs, "TestInt32.Required", "TestInt32.Required", "Required", "Required", "required")
	AssertError(t, errs, "TestInt32.Len", "TestInt32.Len", "Len", "Len", "len")
	AssertError(t, errs, "TestInt32.Min", "TestInt32.Min", "Min", "Min", "min")
	AssertError(t, errs, "TestInt32.Max", "TestInt32.Max", "Max", "Max", "max")
	AssertError(t, errs, "TestInt32.MinMax", "TestInt32.MinMax", "MinMax", "MinMax", "min")
	AssertError(t, errs, "TestInt32.Lt", "TestInt32.Lt", "Lt", "Lt", "lt")
	AssertError(t, errs, "TestInt32.Lte", "TestInt32.Lte", "Lte", "Lte", "lte")
	AssertError(t, errs, "TestInt32.Gt", "TestInt32.Gt", "Gt", "Gt", "gt")
	AssertError(t, errs, "TestInt32.Gte", "TestInt32.Gte", "Gte", "Gte", "gte")
	AssertError(t, errs, "TestInt32.OmitEmpty", "TestInt32.OmitEmpty", "OmitEmpty", "OmitEmpty", "max")
}

func TestStructUint64Validation(t *testing.T) {

	validate := New()

	tSuccess := &TestUint64{
		Required:  1,
		Len:       10,
		Min:       1,
		Max:       10,
		MinMax:    5,
		OmitEmpty: 0,
	}

	errs := validate.Struct(tSuccess)
	Equal(t, errs, nil)

	tFail := &TestUint64{
		Required:  0,
		Len:       11,
		Min:       0,
		Max:       11,
		MinMax:    0,
		OmitEmpty: 11,
	}

	errs = validate.Struct(tFail)

	// Assert Top Level
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 6)

	// Assert Fields
	AssertError(t, errs, "TestUint64.Required", "TestUint64.Required", "Required", "Required", "required")
	AssertError(t, errs, "TestUint64.Len", "TestUint64.Len", "Len", "Len", "len")
	AssertError(t, errs, "TestUint64.Min", "TestUint64.Min", "Min", "Min", "min")
	AssertError(t, errs, "TestUint64.Max", "TestUint64.Max", "Max", "Max", "max")
	AssertError(t, errs, "TestUint64.MinMax", "TestUint64.MinMax", "MinMax", "MinMax", "min")
	AssertError(t, errs, "TestUint64.OmitEmpty", "TestUint64.OmitEmpty", "OmitEmpty", "OmitEmpty", "max")
}

func TestStructFloat64Validation(t *testing.T) {

	validate := New()

	tSuccess := &TestFloat64{
		Required:  1,
		Len:       10,
		Min:       1,
		Max:       10,
		MinMax:    5,
		OmitEmpty: 0,
	}

	errs := validate.Struct(tSuccess)
	Equal(t, errs, nil)

	tFail := &TestFloat64{
		Required:  0,
		Len:       11,
		Min:       0,
		Max:       11,
		MinMax:    0,
		OmitEmpty: 11,
	}

	errs = validate.Struct(tFail)

	// Assert Top Level
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 6)

	// Assert Fields
	AssertError(t, errs, "TestFloat64.Required", "TestFloat64.Required", "Required", "Required", "required")
	AssertError(t, errs, "TestFloat64.Len", "TestFloat64.Len", "Len", "Len", "len")
	AssertError(t, errs, "TestFloat64.Min", "TestFloat64.Min", "Min", "Min", "min")
	AssertError(t, errs, "TestFloat64.Max", "TestFloat64.Max", "Max", "Max", "max")
	AssertError(t, errs, "TestFloat64.MinMax", "TestFloat64.MinMax", "MinMax", "MinMax", "min")
	AssertError(t, errs, "TestFloat64.OmitEmpty", "TestFloat64.OmitEmpty", "OmitEmpty", "OmitEmpty", "max")
}

func TestStructSliceValidation(t *testing.T) {

	validate := New()

	tSuccess := &TestSlice{
		Required:  []int{1},
		Len:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		Min:       []int{1, 2},
		Max:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		MinMax:    []int{1, 2, 3, 4, 5},
		OmitEmpty: nil,
	}

	errs := validate.Struct(tSuccess)
	Equal(t, errs, nil)

	tFail := &TestSlice{
		Required:  nil,
		Len:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
		Min:       []int{},
		Max:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
		MinMax:    []int{},
		OmitEmpty: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
	}

	errs = validate.Struct(tFail)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 6)

	// Assert Field Errors
	AssertError(t, errs, "TestSlice.Required", "TestSlice.Required", "Required", "Required", "required")
	AssertError(t, errs, "TestSlice.Len", "TestSlice.Len", "Len", "Len", "len")
	AssertError(t, errs, "TestSlice.Min", "TestSlice.Min", "Min", "Min", "min")
	AssertError(t, errs, "TestSlice.Max", "TestSlice.Max", "Max", "Max", "max")
	AssertError(t, errs, "TestSlice.MinMax", "TestSlice.MinMax", "MinMax", "MinMax", "min")
	AssertError(t, errs, "TestSlice.OmitEmpty", "TestSlice.OmitEmpty", "OmitEmpty", "OmitEmpty", "max")

	fe := getError(errs, "TestSlice.Len", "TestSlice.Len")
	NotEqual(t, fe, nil)
	Equal(t, fe.Field(), "Len")
	Equal(t, fe.StructField(), "Len")
	Equal(t, fe.Namespace(), "TestSlice.Len")
	Equal(t, fe.StructNamespace(), "TestSlice.Len")
	Equal(t, fe.Tag(), "len")
	Equal(t, fe.ActualTag(), "len")
	Equal(t, fe.Param(), "10")
	Equal(t, fe.Kind(), reflect.Slice)
	Equal(t, fe.Type(), reflect.TypeOf([]int{}))

	_, ok := fe.Value().([]int)
	Equal(t, ok, true)

}

func TestInvalidStruct(t *testing.T) {

	validate := New()

	s := &SubTest{
		Test: "1",
	}

	err := validate.Struct(s.Test)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "validator: (nil string)")

	err = validate.Struct(nil)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "validator: (nil)")

	err = validate.StructPartial(nil, "SubTest.Test")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "validator: (nil)")

	err = validate.StructExcept(nil, "SubTest.Test")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "validator: (nil)")
}

func TestInvalidValidatorFunction(t *testing.T) {

	validate := New()

	s := &SubTest{
		Test: "1",
	}

	PanicMatches(t, func() { validate.Var(s.Test, "zzxxBadFunction") }, "Undefined validation function 'zzxxBadFunction' on field ''")
}

func TestCustomFieldName(t *testing.T) {

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("schema"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type A struct {
		B string `schema:"b" validate:"required"`
		C string `schema:"c" validate:"required"`
		D []bool `schema:"d" validate:"required"`
		E string `schema:"-" validate:"required"`
	}

	a := &A{}

	err := validate.Struct(a)
	NotEqual(t, err, nil)

	errs := err.(ValidationErrors)
	Equal(t, len(errs), 4)
	Equal(t, getError(errs, "A.b", "A.B").Field(), "b")
	Equal(t, getError(errs, "A.c", "A.C").Field(), "c")
	Equal(t, getError(errs, "A.d", "A.D").Field(), "d")
	Equal(t, getError(errs, "A.E", "A.E").Field(), "E")

	v2 := New()
	err = v2.Struct(a)
	NotEqual(t, err, nil)

	errs = err.(ValidationErrors)
	Equal(t, len(errs), 4)
	Equal(t, getError(errs, "A.B", "A.B").Field(), "B")
	Equal(t, getError(errs, "A.C", "A.C").Field(), "C")
	Equal(t, getError(errs, "A.D", "A.D").Field(), "D")
	Equal(t, getError(errs, "A.E", "A.E").Field(), "E")
}

func TestMutipleRecursiveExtractStructCache(t *testing.T) {

	validate := New()

	type Recursive struct {
		Field *string `validate:"required,len=5,ne=string"`
	}

	var test Recursive

	current := reflect.ValueOf(test)
	name := "Recursive"
	proceed := make(chan struct{})

	sc := validate.extractStructCache(current, name)
	ptr := fmt.Sprintf("%p", sc)

	for i := 0; i < 100; i++ {

		go func() {
			<-proceed
			sc := validate.extractStructCache(current, name)
			Equal(t, ptr, fmt.Sprintf("%p", sc))
		}()
	}

	close(proceed)
}

// Thanks @robbrockbank, see https://github.com/go-playground/validator/issues/249
func TestPointerAndOmitEmpty(t *testing.T) {

	validate := New()

	type Test struct {
		MyInt *int `validate:"omitempty,gte=2,lte=255"`
	}

	val1 := 0
	val2 := 256

	t1 := Test{MyInt: &val1} // This should fail validation on gte because value is 0
	t2 := Test{MyInt: &val2} // This should fail validate on lte because value is 256
	t3 := Test{MyInt: nil}   // This should succeed validation because pointer is nil

	errs := validate.Struct(t1)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.MyInt", "Test.MyInt", "MyInt", "MyInt", "gte")

	errs = validate.Struct(t2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "Test.MyInt", "Test.MyInt", "MyInt", "MyInt", "lte")

	errs = validate.Struct(t3)
	Equal(t, errs, nil)

	type TestIface struct {
		MyInt interface{} `validate:"omitempty,gte=2,lte=255"`
	}

	ti1 := TestIface{MyInt: &val1} // This should fail validation on gte because value is 0
	ti2 := TestIface{MyInt: &val2} // This should fail validate on lte because value is 256
	ti3 := TestIface{MyInt: nil}   // This should succeed validation because pointer is nil

	errs = validate.Struct(ti1)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestIface.MyInt", "TestIface.MyInt", "MyInt", "MyInt", "gte")

	errs = validate.Struct(ti2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestIface.MyInt", "TestIface.MyInt", "MyInt", "MyInt", "lte")

	errs = validate.Struct(ti3)
	Equal(t, errs, nil)
}

func TestRequired(t *testing.T) {

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type Test struct {
		Value interface{} `validate:"required"`
	}

	var test Test

	err := validate.Struct(test)
	NotEqual(t, err, nil)
	AssertError(t, err.(ValidationErrors), "Test.Value", "Test.Value", "Value", "Value", "required")
}

func TestTranslations(t *testing.T) {
	en := en.New()
	uni := ut.New(en, en, fr.New())

	trans, _ := uni.GetTranslator("en")
	fr, _ := uni.GetTranslator("fr")

	validate := New()
	err := validate.RegisterTranslation("required", trans,
		func(ut ut.Translator) (err error) {

			// using this stype because multiple translation may have to be added for the full translation
			if err = ut.Add("required", "{0} is a required field", false); err != nil {
				return
			}

			return

		}, func(ut ut.Translator, fe FieldError) string {

			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				fmt.Printf("warning: error translating FieldError: %#v", fe.(*fieldError))
				return fe.(*fieldError).Error()
			}

			return t
		})
	Equal(t, err, nil)

	err = validate.RegisterTranslation("required", fr,
		func(ut ut.Translator) (err error) {

			// using this stype because multiple translation may have to be added for the full translation
			if err = ut.Add("required", "{0} est un champ obligatoire", false); err != nil {
				return
			}

			return

		}, func(ut ut.Translator, fe FieldError) string {

			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				fmt.Printf("warning: error translating FieldError: %#v", fe.(*fieldError))
				return fe.(*fieldError).Error()
			}

			return t
		})

	Equal(t, err, nil)

	type Test struct {
		Value interface{} `validate:"required"`
	}

	var test Test

	err = validate.Struct(test)
	NotEqual(t, err, nil)

	errs := err.(ValidationErrors)
	Equal(t, len(errs), 1)

	fe := errs[0]
	Equal(t, fe.Tag(), "required")
	Equal(t, fe.Namespace(), "Test.Value")
	Equal(t, fe.Translate(trans), fmt.Sprintf("%s is a required field", fe.Field()))
	Equal(t, fe.Translate(fr), fmt.Sprintf("%s est un champ obligatoire", fe.Field()))

	nl := nl.New()
	uni2 := ut.New(nl, nl)
	trans2, _ := uni2.GetTranslator("nl")
	Equal(t, fe.Translate(trans2), "Key: 'Test.Value' Error:Field validation for 'Value' failed on the 'required' tag")

	terrs := errs.Translate(trans)
	Equal(t, len(terrs), 1)

	v, ok := terrs["Test.Value"]
	Equal(t, ok, true)
	Equal(t, v, fmt.Sprintf("%s is a required field", fe.Field()))

	terrs = errs.Translate(fr)
	Equal(t, len(terrs), 1)

	v, ok = terrs["Test.Value"]
	Equal(t, ok, true)
	Equal(t, v, fmt.Sprintf("%s est un champ obligatoire", fe.Field()))

	type Test2 struct {
		Value string `validate:"gt=1"`
	}

	var t2 Test2

	err = validate.Struct(t2)
	NotEqual(t, err, nil)

	errs = err.(ValidationErrors)
	Equal(t, len(errs), 1)

	fe = errs[0]
	Equal(t, fe.Tag(), "gt")
	Equal(t, fe.Namespace(), "Test2.Value")
	Equal(t, fe.Translate(trans), "Key: 'Test2.Value' Error:Field validation for 'Value' failed on the 'gt' tag")
}

func TestTranslationErrors(t *testing.T) {

	en := en.New()
	uni := ut.New(en, en, fr.New())

	trans, _ := uni.GetTranslator("en")
	trans.Add("required", "{0} is a required field", false) // using translator outside of validator also

	validate := New()
	err := validate.RegisterTranslation("required", trans,
		func(ut ut.Translator) (err error) {

			// using this stype because multiple translation may have to be added for the full translation
			if err = ut.Add("required", "{0} is a required field", false); err != nil {
				return
			}

			return

		}, func(ut ut.Translator, fe FieldError) string {

			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				fmt.Printf("warning: error translating FieldError: %#v", fe.(*fieldError))
				return fe.(*fieldError).Error()
			}

			return t
		})

	NotEqual(t, err, nil)
	Equal(t, err.Error(), "error: conflicting key 'required' rule 'Unknown' with text '{0} is a required field' for locale 'en', value being ignored")
}

func TestStructFiltered(t *testing.T) {

	p1 := func(ns []byte) bool {
		if bytes.HasSuffix(ns, []byte("NoTag")) || bytes.HasSuffix(ns, []byte("Required")) {
			return false
		}

		return true
	}

	p2 := func(ns []byte) bool {
		if bytes.HasSuffix(ns, []byte("SubSlice[0].Test")) ||
			bytes.HasSuffix(ns, []byte("SubSlice[0]")) ||
			bytes.HasSuffix(ns, []byte("SubSlice")) ||
			bytes.HasSuffix(ns, []byte("Sub")) ||
			bytes.HasSuffix(ns, []byte("SubIgnore")) ||
			bytes.HasSuffix(ns, []byte("Anonymous")) ||
			bytes.HasSuffix(ns, []byte("Anonymous.A")) {
			return false
		}

		return true
	}

	p3 := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, []byte("SubTest.Test"))
	}

	// p4 := []string{
	// 	"A",
	// }

	tPartial := &TestPartial{
		NoTag:    "NoTag",
		Required: "Required",

		SubSlice: []*SubTest{
			{

				Test: "Required",
			},
			{

				Test: "Required",
			},
		},

		Sub: &SubTest{
			Test: "1",
		},
		SubIgnore: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A             string     `validate:"required"`
			ASubSlice     []*SubTest `validate:"required,dive"`
			SubAnonStruct []struct {
				Test      string `validate:"required"`
				OtherTest string `validate:"required"`
			} `validate:"required,dive"`
		}{
			A: "1",
			ASubSlice: []*SubTest{
				{
					Test: "Required",
				},
				{
					Test: "Required",
				},
			},

			SubAnonStruct: []struct {
				Test      string `validate:"required"`
				OtherTest string `validate:"required"`
			}{
				{"Required", "RequiredOther"},
				{"Required", "RequiredOther"},
			},
		},
	}

	validate := New()

	// the following should all return no errors as everything is valid in
	// the default state
	errs := validate.StructFilteredCtx(context.Background(), tPartial, p1)
	Equal(t, errs, nil)

	errs = validate.StructFiltered(tPartial, p2)
	Equal(t, errs, nil)

	// this isn't really a robust test, but is ment to illustrate the ANON CASE below
	errs = validate.StructFiltered(tPartial.SubSlice[0], p3)
	Equal(t, errs, nil)

	// mod tParial for required feild and re-test making sure invalid fields are NOT required:
	tPartial.Required = ""

	// inversion and retesting Partial to generate failures:
	errs = validate.StructFiltered(tPartial, p1)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.Required", "TestPartial.Required", "Required", "Required", "required")

	// reset Required field, and set nested struct
	tPartial.Required = "Required"
	tPartial.Anonymous.A = ""

	// will pass as unset feilds is not going to be tested
	errs = validate.StructFiltered(tPartial, p1)
	Equal(t, errs, nil)

	// will fail as unset feild is tested
	errs = validate.StructFiltered(tPartial, p2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.Anonymous.A", "TestPartial.Anonymous.A", "A", "A", "required")

	// reset nested struct and unset struct in slice
	tPartial.Anonymous.A = "Required"
	tPartial.SubSlice[0].Test = ""

	// these will pass as unset item is NOT tested
	errs = validate.StructFiltered(tPartial, p1)
	Equal(t, errs, nil)

	errs = validate.StructFiltered(tPartial, p2)
	NotEqual(t, errs, nil)
	AssertError(t, errs, "TestPartial.SubSlice[0].Test", "TestPartial.SubSlice[0].Test", "Test", "Test", "required")
	Equal(t, len(errs.(ValidationErrors)), 1)

	// Unset second slice member concurrently to test dive behavior:
	tPartial.SubSlice[1].Test = ""

	errs = validate.StructFiltered(tPartial, p1)
	Equal(t, errs, nil)

	errs = validate.StructFiltered(tPartial, p2)
	NotEqual(t, errs, nil)
	Equal(t, len(errs.(ValidationErrors)), 1)
	AssertError(t, errs, "TestPartial.SubSlice[0].Test", "TestPartial.SubSlice[0].Test", "Test", "Test", "required")

	// reset struct in slice, and unset struct in slice in unset posistion
	tPartial.SubSlice[0].Test = "Required"

	// these will pass as the unset item is NOT tested
	errs = validate.StructFiltered(tPartial, p1)
	Equal(t, errs, nil)

	errs = validate.StructFiltered(tPartial, p2)
	Equal(t, errs, nil)

	tPartial.SubSlice[1].Test = "Required"
	tPartial.Anonymous.SubAnonStruct[0].Test = ""

	// these will pass as the unset item is NOT tested
	errs = validate.StructFiltered(tPartial, p1)
	Equal(t, errs, nil)

	errs = validate.StructFiltered(tPartial, p2)
	Equal(t, errs, nil)

	dt := time.Now()
	err := validate.StructFiltered(&dt, func(ns []byte) bool { return true })
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "validator: (nil *time.Time)")
}

func TestRequiredPtr(t *testing.T) {

	type Test struct {
		Bool *bool `validate:"required"`
	}

	validate := New()

	f := false

	test := Test{
		Bool: &f,
	}

	err := validate.Struct(test)
	Equal(t, err, nil)

	tr := true

	test.Bool = &tr

	err = validate.Struct(test)
	Equal(t, err, nil)

	test.Bool = nil

	err = validate.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(ValidationErrors)
	Equal(t, ok, true)
	Equal(t, len(errs), 1)
	AssertError(t, errs, "Test.Bool", "Test.Bool", "Bool", "Bool", "required")

	type Test2 struct {
		Bool bool `validate:"required"`
	}

	var test2 Test2

	err = validate.Struct(test2)
	NotEqual(t, err, nil)

	errs, ok = err.(ValidationErrors)
	Equal(t, ok, true)
	Equal(t, len(errs), 1)
	AssertError(t, errs, "Test2.Bool", "Test2.Bool", "Bool", "Bool", "required")

	test2.Bool = true

	err = validate.Struct(test2)
	Equal(t, err, nil)

	type Test3 struct {
		Arr []string `validate:"required"`
	}

	var test3 Test3

	err = validate.Struct(test3)
	NotEqual(t, err, nil)

	errs, ok = err.(ValidationErrors)
	Equal(t, ok, true)
	Equal(t, len(errs), 1)
	AssertError(t, errs, "Test3.Arr", "Test3.Arr", "Arr", "Arr", "required")

	test3.Arr = make([]string, 0)

	err = validate.Struct(test3)
	Equal(t, err, nil)

	type Test4 struct {
		Arr *[]string `validate:"required"` // I know I know pointer to array, just making sure validation works as expected...
	}

	var test4 Test4

	err = validate.Struct(test4)
	NotEqual(t, err, nil)

	errs, ok = err.(ValidationErrors)
	Equal(t, ok, true)
	Equal(t, len(errs), 1)
	AssertError(t, errs, "Test4.Arr", "Test4.Arr", "Arr", "Arr", "required")

	arr := make([]string, 0)
	test4.Arr = &arr

	err = validate.Struct(test4)
	Equal(t, err, nil)
}

func TestAlphaUnicodeValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"abc", true},
		{"this is a test string", false},
		{"这是一个测试字符串", true},
		{"123", false},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", false},
		{"あいうえおfoobar", true},
		{"test＠example.com", false},
		{"1234abcDE", false},
		{"ｶﾀｶﾅ", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "alphaunicode")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d Alpha Unicode failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d Alpha Unicode failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "alphaunicode" {
					t.Fatalf("Index: %d Alpha Unicode failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestAlphanumericUnicodeValidation(t *testing.T) {

	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"abc", true},
		{"this is a test string", false},
		{"这是一个测试字符串", true},
		{"\u0031\u0032\u0033", true}, // unicode 5
		{"123", true},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", false},
		{"あいうえおfoobar", true},
		{"test＠example.com", false},
		{"1234abcDE", true},
		{"ｶﾀｶﾅ", true},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "alphanumunicode")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d Alphanum Unicode failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d Alphanum Unicode failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "alphanumunicode" {
					t.Fatalf("Index: %d Alphanum Unicode failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestArrayStructNamespace(t *testing.T) {

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type child struct {
		Name string `json:"name" validate:"required"`
	}
	var input struct {
		Children []child `json:"children" validate:"required,gt=0,dive"`
	}
	input.Children = []child{{"ok"}, {""}}

	errs := validate.Struct(input)
	NotEqual(t, errs, nil)

	ve := errs.(ValidationErrors)
	Equal(t, len(ve), 1)
	AssertError(t, errs, "children[1].name", "Children[1].Name", "name", "Name", "required")
}

func TestMapStructNamespace(t *testing.T) {

	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type child struct {
		Name string `json:"name" validate:"required"`
	}
	var input struct {
		Children map[int]child `json:"children" validate:"required,gt=0,dive"`
	}
	input.Children = map[int]child{
		0: {Name: "ok"},
		1: {Name: ""},
	}

	errs := validate.Struct(input)
	NotEqual(t, errs, nil)

	ve := errs.(ValidationErrors)
	Equal(t, len(ve), 1)
	AssertError(t, errs, "children[1].name", "Children[1].Name", "name", "Name", "required")
}

func TestFieldLevelName(t *testing.T) {
	type Test struct {
		String string            `validate:"custom1"      json:"json1"`
		Array  []string          `validate:"dive,custom2" json:"json2"`
		Map    map[string]string `validate:"dive,custom3" json:"json3"`
		Array2 []string          `validate:"custom4"      json:"json4"`
		Map2   map[string]string `validate:"custom5"      json:"json5"`
	}

	var res1, res2, res3, res4, res5, alt1, alt2, alt3, alt4, alt5 string
	validate := New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	validate.RegisterValidation("custom1", func(fl FieldLevel) bool {
		res1 = fl.FieldName()
		alt1 = fl.StructFieldName()
		return true
	})
	validate.RegisterValidation("custom2", func(fl FieldLevel) bool {
		res2 = fl.FieldName()
		alt2 = fl.StructFieldName()
		return true
	})
	validate.RegisterValidation("custom3", func(fl FieldLevel) bool {
		res3 = fl.FieldName()
		alt3 = fl.StructFieldName()
		return true
	})
	validate.RegisterValidation("custom4", func(fl FieldLevel) bool {
		res4 = fl.FieldName()
		alt4 = fl.StructFieldName()
		return true
	})
	validate.RegisterValidation("custom5", func(fl FieldLevel) bool {
		res5 = fl.FieldName()
		alt5 = fl.StructFieldName()
		return true
	})

	test := Test{
		String: "test",
		Array:  []string{"1"},
		Map:    map[string]string{"test": "test"},
	}

	errs := validate.Struct(test)
	Equal(t, errs, nil)
	Equal(t, res1, "json1")
	Equal(t, alt1, "String")
	Equal(t, res2, "json2[0]")
	Equal(t, alt2, "Array[0]")
	Equal(t, res3, "json3[test]")
	Equal(t, alt3, "Map[test]")
	Equal(t, res4, "json4")
	Equal(t, alt4, "Array2")
	Equal(t, res5, "json5")
	Equal(t, alt5, "Map2")
}

func TestValidateStructRegisterCtx(t *testing.T) {

	var ctxVal string

	fnCtx := func(ctx context.Context, fl FieldLevel) bool {
		ctxVal = ctx.Value(&ctxVal).(string)
		return true
	}

	var ctxSlVal string
	slFn := func(ctx context.Context, sl StructLevel) {
		ctxSlVal = ctx.Value(&ctxSlVal).(string)
	}

	type Test struct {
		Field string `validate:"val"`
	}

	var tst Test

	validate := New()
	validate.RegisterValidationCtx("val", fnCtx)
	validate.RegisterStructValidationCtx(slFn, Test{})

	ctx := context.WithValue(context.Background(), &ctxVal, "testval")
	ctx = context.WithValue(ctx, &ctxSlVal, "slVal")
	errs := validate.StructCtx(ctx, tst)
	Equal(t, errs, nil)
	Equal(t, ctxVal, "testval")
	Equal(t, ctxSlVal, "slVal")
}

func TestHostnameValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"test.example.com", true},
		{"example.com", true},
		{"example24.com", true},
		{"test.example24.com", true},
		{"test24.example24.com", true},
		{"example", true},
		{"test.example.com.", false},
		{"example.com.", false},
		{"example24.com.", false},
		{"test.example24.com.", false},
		{"test24.example24.com.", false},
		{"example.", false},
		{"192.168.0.1", false},
		{"email@example.com", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
		{"2001:cdba:0:0:0:0:3257:9652", false},
		{"2001:cdba::3257:9652", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "hostname")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d hostname failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d hostname failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "hostname" {
					t.Fatalf("Index: %d hostname failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestFQDNValidation(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{"test.example.com", true},
		{"example.com", true},
		{"example24.com", true},
		{"test.example24.com", true},
		{"test24.example24.com", true},
		{"test.example.com.", true},
		{"example.com.", true},
		{"example24.com.", true},
		{"test.example24.com.", true},
		{"test24.example24.com.", true},
		{"test24.example24.com..", false},
		{"example", false},
		{"192.168.0.1", false},
		{"email@example.com", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
		{"2001:cdba:0:0:0:0:3257:9652", false},
		{"2001:cdba::3257:9652", false},
		{"", false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "fqdn")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d fqdn failed Error: %s", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d fqdn failed Error: %s", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "fqdn" {
					t.Fatalf("Index: %d fqdn failed Error: %s", i, errs)
				}
			}
		}
	}
}

func TestIsDefault(t *testing.T) {

	validate := New()

	type Inner struct {
		String string `validate:"isdefault"`
	}
	type Test struct {
		String string `validate:"isdefault"`
		Inner  *Inner `validate:"isdefault"`
	}

	var tt Test

	errs := validate.Struct(tt)
	Equal(t, errs, nil)

	tt.Inner = &Inner{String: ""}
	errs = validate.Struct(tt)
	NotEqual(t, errs, nil)

	fe := errs.(ValidationErrors)[0]
	Equal(t, fe.Field(), "Inner")
	Equal(t, fe.Namespace(), "Test.Inner")
	Equal(t, fe.Tag(), "isdefault")

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	type Inner2 struct {
		String string `validate:"isdefault"`
	}

	type Test2 struct {
		Inner Inner2 `validate:"isdefault" json:"inner"`
	}

	var t2 Test2
	errs = validate.Struct(t2)
	Equal(t, errs, nil)

	t2.Inner.String = "Changed"
	errs = validate.Struct(t2)
	NotEqual(t, errs, nil)

	fe = errs.(ValidationErrors)[0]
	Equal(t, fe.Field(), "inner")
	Equal(t, fe.Namespace(), "Test2.inner")
	Equal(t, fe.Tag(), "isdefault")
}

func TestUniqueValidation(t *testing.T) {
	tests := []struct {
		param    interface{}
		expected bool
	}{
		{[]string{"a", "b"}, true},
		{[]int{1, 2}, true},
		{[]float64{1, 2}, true},
		{[]interface{}{"a", "b"}, true},
		{[]interface{}{"a", 1}, true},
		{[]float64{1, 1}, false},
		{[]int{1, 1}, false},
		{[]string{"a", "a"}, false},
		{[]interface{}{"a", "a"}, false},
		{[]interface{}{"a", 1, "b", 1}, false},
	}

	validate := New()

	for i, test := range tests {

		errs := validate.Var(test.param, "unique")

		if test.expected {
			if !IsEqual(errs, nil) {
				t.Fatalf("Index: %d unique failed Error: %v", i, errs)
			}
		} else {
			if IsEqual(errs, nil) {
				t.Fatalf("Index: %d unique failed Error: %v", i, errs)
			} else {
				val := getError(errs, "", "")
				if val.Tag() != "unique" {
					t.Fatalf("Index: %d unique failed Error: %v", i, errs)
				}
			}
		}
	}
	PanicMatches(t, func() { validate.Var(1.0, "unique") }, "Bad field type float64")
}
