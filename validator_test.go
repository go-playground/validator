package validator_test

import (
	"fmt"
	"testing"

	"github.com/joeybloggs/go-validate-yourself"
	. "gopkg.in/check.v1"
)

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
	Required  string `validate:"required"`
	Len       string `validate:"len=10"`
	Min       string `validate:"min=1"`
	Max       string `validate:"max=10"`
	MinMax    string `validate:"min=1,max=10"`
	OmitEmpty string `validate:"omitempty,min=1,max=10"`
	Sub       *SubTest
	SubIgnore *SubTest `validate:"-"`
	Anonymous struct {
		A string `validate:"required"`
	}
	Iface I
}

type TestInt32 struct {
	Required  int `validate:"required"`
	Len       int `validate:"len=10"`
	Min       int `validate:"min=1"`
	Max       int `validate:"max=10"`
	MinMax    int `validate:"min=1,max=10"`
	OmitEmpty int `validate:"omitempty,min=1,max=10"`
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
	Required  int64 `validate:"required"`
	Len       int64 `validate:"len=10"`
	Min       int64 `validate:"min=1"`
	Max       int64 `validate:"max=10"`
	MinMax    int64 `validate:"min=1,max=10"`
	OmitEmpty int64 `validate:"omitempty,min=1,max=10"`
}

type TestSlice struct {
	Required  []int `validate:"required"`
	Len       []int `validate:"len=10"`
	Min       []int `validate:"min=1"`
	Max       []int `validate:"max=10"`
	MinMax    []int `validate:"min=1,max=10"`
	OmitEmpty []int `validate:"omitempty,min=1,max=10"`
}

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func AssetStruct(s *validator.StructValidationErrors, structFieldName string, expectedStructName string, c *C) *validator.StructValidationErrors {

	val, ok := s.StructErrors[structFieldName]
	c.Assert(ok, Equals, true)
	c.Assert(val, NotNil)
	c.Assert(val.Struct, Equals, expectedStructName)

	return val
}

func AssertFieldError(s *validator.StructValidationErrors, field string, expectedTag string, c *C) {

	val, ok := s.Errors[field]
	c.Assert(ok, Equals, true)
	c.Assert(val, NotNil)
	c.Assert(val.Field, Equals, field)
	c.Assert(val.ErrorTag, Equals, expectedTag)
}

func AssertMapFieldError(s map[string]*validator.FieldValidationError, field string, expectedTag string, c *C) {

	val, ok := s[field]
	c.Assert(ok, Equals, true)
	c.Assert(val, NotNil)
	c.Assert(val.Field, Equals, field)
	c.Assert(val.ErrorTag, Equals, expectedTag)
}

func (ms *MySuite) TestAlphaNumeric(c *C) {

	s := "abcd123"
	err := validator.ValidateFieldByTag(s, "alphanum")
	c.Assert(err, IsNil)

	s = "abc!23"
	err = validator.ValidateFieldByTag(s, "alphanum")
	c.Assert(err.Error(), Equals, "alphanum")

	c.Assert(func() { validator.ValidateFieldByTag(1, "alphanum") }, PanicMatches, "Bad field type int")
}

func (ms *MySuite) TestAlpha(c *C) {

	s := "abcd"
	err := validator.ValidateFieldByTag(s, "alpha")
	c.Assert(err, IsNil)

	s = "abc1"
	err = validator.ValidateFieldByTag(s, "alpha")
	c.Assert(err.Error(), Equals, "alpha")

	c.Assert(func() { validator.ValidateFieldByTag(1, "alpha") }, PanicMatches, "Bad field type int")
}

func (ms *MySuite) TestFlattening(c *C) {

	tSuccess := &TestString{
		Required:  "Required",
		Len:       "length==10",
		Min:       "min=1",
		Max:       "1234567890",
		MinMax:    "12345",
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

	err1 := validator.ValidateStruct(tSuccess).Flatten()
	c.Assert(err1, IsNil)

	tFail := &TestString{
		Required:  "",
		Len:       "",
		Min:       "",
		Max:       "12345678901",
		MinMax:    "",
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

	err2 := validator.ValidateStruct(tFail).Flatten()

	// Assert Top Level
	c.Assert(err2, NotNil)

	// Assert Fields
	AssertMapFieldError(err2, "Len", "len", c)

	// Assert Struct Field
	AssertMapFieldError(err2, "Sub.Test", "required", c)

	// Assert Anonymous Struct Field
	AssertMapFieldError(err2, "Anonymous.A", "required", c)

	// Assert Interface Field
	AssertMapFieldError(err2, "Iface.F", "len", c)
}

func (ms *MySuite) TestStructStringValidation(c *C) {

	tSuccess := &TestString{
		Required:  "Required",
		Len:       "length==10",
		Min:       "min=1",
		Max:       "1234567890",
		MinMax:    "12345",
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

	err := validator.ValidateStruct(tSuccess)
	c.Assert(err, IsNil)

	tFail := &TestString{
		Required:  "",
		Len:       "",
		Min:       "",
		Max:       "12345678901",
		MinMax:    "",
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

	err = validator.ValidateStruct(tFail)

	// Assert Top Level
	c.Assert(err, NotNil)
	c.Assert(err.Struct, Equals, "TestString")
	c.Assert(len(err.Errors), Equals, 6)
	c.Assert(len(err.StructErrors), Equals, 3)

	// Assert Fields
	AssertFieldError(err, "Required", "required", c)
	AssertFieldError(err, "Len", "len", c)
	AssertFieldError(err, "Min", "min", c)
	AssertFieldError(err, "Max", "max", c)
	AssertFieldError(err, "MinMax", "min", c)
	AssertFieldError(err, "OmitEmpty", "max", c)

	// Assert Anonymous embedded struct
	AssetStruct(err, "Anonymous", "", c)

	// Assert SubTest embedded struct
	val := AssetStruct(err, "Sub", "SubTest", c)
	c.Assert(len(val.Errors), Equals, 1)
	c.Assert(len(val.StructErrors), Equals, 0)

	AssertFieldError(val, "Test", "required", c)
}

func (ms *MySuite) TestStructInt32Validation(c *C) {

	tSuccess := &TestInt32{
		Required:  1,
		Len:       10,
		Min:       1,
		Max:       10,
		MinMax:    5,
		OmitEmpty: 0,
	}

	err := validator.ValidateStruct(tSuccess)
	c.Assert(err, IsNil)

	tFail := &TestInt32{
		Required:  0,
		Len:       11,
		Min:       -1,
		Max:       11,
		MinMax:    -1,
		OmitEmpty: 11,
	}

	err = validator.ValidateStruct(tFail)

	// Assert Top Level
	c.Assert(err, NotNil)
	c.Assert(err.Struct, Equals, "TestInt32")
	c.Assert(len(err.Errors), Equals, 6)
	c.Assert(len(err.StructErrors), Equals, 0)

	// Assert Fields
	AssertFieldError(err, "Required", "required", c)
	AssertFieldError(err, "Len", "len", c)
	AssertFieldError(err, "Min", "min", c)
	AssertFieldError(err, "Max", "max", c)
	AssertFieldError(err, "MinMax", "min", c)
	AssertFieldError(err, "OmitEmpty", "max", c)
}

func (ms *MySuite) TestStructUint64Validation(c *C) {

	tSuccess := &TestUint64{
		Required:  1,
		Len:       10,
		Min:       1,
		Max:       10,
		MinMax:    5,
		OmitEmpty: 0,
	}

	err := validator.ValidateStruct(tSuccess)
	c.Assert(err, IsNil)

	tFail := &TestUint64{
		Required:  0,
		Len:       11,
		Min:       0,
		Max:       11,
		MinMax:    0,
		OmitEmpty: 11,
	}

	err = validator.ValidateStruct(tFail)

	// Assert Top Level
	c.Assert(err, NotNil)
	c.Assert(err.Struct, Equals, "TestUint64")
	c.Assert(len(err.Errors), Equals, 6)
	c.Assert(len(err.StructErrors), Equals, 0)

	// Assert Fields
	AssertFieldError(err, "Required", "required", c)
	AssertFieldError(err, "Len", "len", c)
	AssertFieldError(err, "Min", "min", c)
	AssertFieldError(err, "Max", "max", c)
	AssertFieldError(err, "MinMax", "min", c)
	AssertFieldError(err, "OmitEmpty", "max", c)
}

func (ms *MySuite) TestStructFloat64Validation(c *C) {

	tSuccess := &TestFloat64{
		Required:  1,
		Len:       10,
		Min:       1,
		Max:       10,
		MinMax:    5,
		OmitEmpty: 0,
	}

	err := validator.ValidateStruct(tSuccess)
	c.Assert(err, IsNil)

	tFail := &TestFloat64{
		Required:  0,
		Len:       11,
		Min:       0,
		Max:       11,
		MinMax:    0,
		OmitEmpty: 11,
	}

	err = validator.ValidateStruct(tFail)

	// Assert Top Level
	c.Assert(err, NotNil)
	c.Assert(err.Struct, Equals, "TestFloat64")
	c.Assert(len(err.Errors), Equals, 6)
	c.Assert(len(err.StructErrors), Equals, 0)

	// Assert Fields
	AssertFieldError(err, "Required", "required", c)
	AssertFieldError(err, "Len", "len", c)
	AssertFieldError(err, "Min", "min", c)
	AssertFieldError(err, "Max", "max", c)
	AssertFieldError(err, "MinMax", "min", c)
	AssertFieldError(err, "OmitEmpty", "max", c)
}

func (ms *MySuite) TestStructSliceValidation(c *C) {

	tSuccess := &TestSlice{
		Required:  []int{1},
		Len:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		Min:       []int{1, 2},
		Max:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		MinMax:    []int{1, 2, 3, 4, 5},
		OmitEmpty: []int{},
	}

	err := validator.ValidateStruct(tSuccess)
	c.Assert(err, IsNil)

	tFail := &TestSlice{
		Required:  []int{},
		Len:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
		Min:       []int{},
		Max:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
		MinMax:    []int{},
		OmitEmpty: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
	}

	err = validator.ValidateStruct(tFail)

	// Assert Top Level
	c.Assert(err, NotNil)
	c.Assert(err.Struct, Equals, "TestSlice")
	c.Assert(len(err.Errors), Equals, 6)
	c.Assert(len(err.StructErrors), Equals, 0)

	// Assert Fields
	AssertFieldError(err, "Required", "required", c)
	AssertFieldError(err, "Len", "len", c)
	AssertFieldError(err, "Min", "min", c)
	AssertFieldError(err, "Max", "max", c)
	AssertFieldError(err, "MinMax", "min", c)
	AssertFieldError(err, "OmitEmpty", "max", c)
}

func (ms *MySuite) TestInvalidStruct(c *C) {
	s := &SubTest{
		Test: "1",
	}

	c.Assert(func() { validator.ValidateStruct(s.Test) }, PanicMatches, "interface passed for validation is not a struct")
}

func (ms *MySuite) TestInvalidField(c *C) {
	s := &SubTest{
		Test: "1",
	}

	c.Assert(func() { validator.ValidateFieldByTag(s, "required") }, PanicMatches, "Invalid field passed to ValidateFieldWithTag")
}

func (ms *MySuite) TestInvalidTagField(c *C) {
	s := &SubTest{
		Test: "1",
	}

	c.Assert(func() { validator.ValidateFieldByTag(s.Test, "") }, PanicMatches, fmt.Sprintf("Invalid validation tag on field %s", ""))
}

func (ms *MySuite) TestInvalidValidatorFunction(c *C) {
	s := &SubTest{
		Test: "1",
	}

	c.Assert(func() { validator.ValidateFieldByTag(s.Test, "zzxxBadFunction") }, PanicMatches, fmt.Sprintf("Undefined validation function on field %s", ""))
}
