package validator

import (
	"bytes"
	sql "database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func BenchmarkFieldSuccess(b *testing.B) {
	validate := New()
	s := "1"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "len=1")
	}
}

func BenchmarkFieldSuccessParallel(b *testing.B) {
	validate := New()
	s := "1"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(&s, "len=1")
		}
	})
}

func BenchmarkFieldFailure(b *testing.B) {
	validate := New()
	s := "12"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "len=1")
	}
}

func BenchmarkFieldFailureParallel(b *testing.B) {
	validate := New()
	s := "12"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(&s, "len=1")
		}
	})
}

func BenchmarkFieldArrayDiveSuccess(b *testing.B) {
	validate := New()
	m := []string{"val1", "val2", "val3"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = validate.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldArrayDiveSuccessParallel(b *testing.B) {
	validate := New()
	m := []string{"val1", "val2", "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldArrayDiveFailure(b *testing.B) {
	validate := New()
	m := []string{"val1", "", "val3"}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldArrayDiveFailureParallel(b *testing.B) {
	validate := New()
	m := []string{"val1", "", "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldMapDiveSuccess(b *testing.B) {
	validate := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = validate.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldMapDiveSuccessParallel(b *testing.B) {
	validate := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldMapDiveFailure(b *testing.B) {
	validate := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldMapDiveFailureParallel(b *testing.B) {
	validate := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldMapDiveWithKeysSuccess(b *testing.B) {
	validate := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = validate.Var(m, "required,dive,keys,required,endkeys,required")
	}
}

func BenchmarkFieldMapDiveWithKeysSuccessParallel(b *testing.B) {
	validate := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(m, "required,dive,keys,required,endkeys,required")
		}
	})
}

func BenchmarkFieldMapDiveWithKeysFailure(b *testing.B) {
	validate := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(m, "required,dive,keys,required,endkeys,required")
	}
}

func BenchmarkFieldMapDiveWithKeysFailureParallel(b *testing.B) {
	validate := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(m, "required,dive,keys,required,endkeys,required")
		}
	})
}

func BenchmarkFieldCustomTypeSuccess(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{
		Name: "1",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(val, "len=1")
	}
}

func BenchmarkFieldCustomTypeSuccessParallel(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{
		Name: "1",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(val, "len=1")
		}
	})
}

func BenchmarkFieldCustomTypeFailure(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(val, "len=1")
	}
}

func BenchmarkFieldCustomTypeFailureParallel(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(val, "len=1")
		}
	})
}

func BenchmarkFieldOrTagSuccess(b *testing.B) {
	validate := New()
	s := "rgba(0,0,0,1)"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(s, "rgb|rgba")
	}
}

func BenchmarkFieldOrTagSuccessParallel(b *testing.B) {
	validate := New()
	s := "rgba(0,0,0,1)"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(s, "rgb|rgba")
		}
	})
}

func BenchmarkFieldOrTagFailure(b *testing.B) {
	validate := New()
	s := "#000"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(s, "rgb|rgba")
	}
}

func BenchmarkFieldOrTagFailureParallel(b *testing.B) {
	validate := New()
	s := "#000"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(s, "rgb|rgba")
		}
	})
}

func BenchmarkStructLevelValidationSuccess(b *testing.B) {
	validate := New()
	validate.RegisterStructValidation(StructValidationTestStructSuccess, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(tst)
	}
}

func BenchmarkStructLevelValidationSuccessParallel(b *testing.B) {
	validate := New()
	validate.RegisterStructValidation(StructValidationTestStructSuccess, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(tst)
		}
	})
}

func BenchmarkStructLevelValidationFailure(b *testing.B) {
	validate := New()
	validate.RegisterStructValidation(StructValidationTestStruct, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(tst)
	}
}

func BenchmarkStructLevelValidationFailureParallel(b *testing.B) {
	validate := New()
	validate.RegisterStructValidation(StructValidationTestStruct, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(tst)
		}
	})
}

func BenchmarkStructSimpleCustomTypeSuccess(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})

	val := valuer{
		Name: "1",
	}

	type Foo struct {
		Valuer   valuer `validate:"len=1"`
		IntValue int    `validate:"min=5,max=10"`
	}

	validFoo := &Foo{Valuer: val, IntValue: 7}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(validFoo)
	}
}

func BenchmarkStructSimpleCustomTypeSuccessParallel(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{
		Name: "1",
	}

	type Foo struct {
		Valuer   valuer `validate:"len=1"`
		IntValue int    `validate:"min=5,max=10"`
	}
	validFoo := &Foo{Valuer: val, IntValue: 7}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(validFoo)
		}
	})
}

func BenchmarkStructSimpleCustomTypeFailure(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})

	val := valuer{}

	type Foo struct {
		Valuer   valuer `validate:"len=1"`
		IntValue int    `validate:"min=5,max=10"`
	}
	validFoo := &Foo{Valuer: val, IntValue: 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(validFoo)
	}
}

func BenchmarkStructSimpleCustomTypeFailureParallel(b *testing.B) {
	validate := New()
	validate.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})

	val := valuer{}

	type Foo struct {
		Valuer   valuer `validate:"len=1"`
		IntValue int    `validate:"min=5,max=10"`
	}
	validFoo := &Foo{Valuer: val, IntValue: 3}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(validate.Struct(validFoo))
		}
	})
}

func BenchmarkStructFilteredSuccess(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}
	byts := []byte("Name")
	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.StructFiltered(test, fn)
	}
}

func BenchmarkStructFilteredSuccessParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}
	byts := []byte("Name")
	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.StructFiltered(test, fn)
		}
	})
}

func BenchmarkStructFilteredFailure(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	byts := []byte("NickName")

	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.StructFiltered(test, fn)
	}
}

func BenchmarkStructFilteredFailureParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}
	byts := []byte("NickName")
	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.StructFiltered(test, fn)
		}
	})
}

func BenchmarkStructPartialSuccess(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.StructPartial(test, "Name")
	}
}

func BenchmarkStructPartialSuccessParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.StructPartial(test, "Name")
		}
	})
}

func BenchmarkStructPartialFailure(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.StructPartial(test, "NickName")
	}
}

func BenchmarkStructPartialFailureParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.StructPartial(test, "NickName")
		}
	})
}

func BenchmarkStructExceptSuccess(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.StructExcept(test, "Nickname")
	}
}

func BenchmarkStructExceptSuccessParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.StructExcept(test, "NickName")
		}
	})
}

func BenchmarkStructExceptFailure(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.StructExcept(test, "Name")
	}
}

func BenchmarkStructExceptFailureParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Name     string `validate:"required"`
		NickName string `validate:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.StructExcept(test, "Name")
		}
	})
}

func BenchmarkStructSimpleCrossFieldSuccess(b *testing.B) {
	validate := New()

	type Test struct {
		Start time.Time
		End   time.Time `validate:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)
	test := &Test{
		Start: now,
		End:   then,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(test)
	}
}

func BenchmarkStructSimpleCrossFieldSuccessParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Start time.Time
		End   time.Time `validate:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)
	test := &Test{
		Start: now,
		End:   then,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(test)
		}
	})
}

func BenchmarkStructSimpleCrossFieldFailure(b *testing.B) {
	validate := New()

	type Test struct {
		Start time.Time
		End   time.Time `validate:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * -5)

	test := &Test{
		Start: now,
		End:   then,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(test)
	}
}

func BenchmarkStructSimpleCrossFieldFailureParallel(b *testing.B) {
	validate := New()

	type Test struct {
		Start time.Time
		End   time.Time `validate:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * -5)
	test := &Test{
		Start: now,
		End:   then,
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(test)
		}
	})
}

func BenchmarkStructSimpleCrossStructCrossFieldSuccess(b *testing.B) {
	validate := New()

	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `validate:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	inner := &Inner{
		Start: now,
	}
	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(outer)
	}
}

func BenchmarkStructSimpleCrossStructCrossFieldSuccessParallel(b *testing.B) {
	validate := New()

	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `validate:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	inner := &Inner{
		Start: now,
	}
	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(outer)
		}
	})
}

func BenchmarkStructSimpleCrossStructCrossFieldFailure(b *testing.B) {
	validate := New()
	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `validate:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)

	inner := &Inner{
		Start: then,
	}

	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(outer)
	}
}

func BenchmarkStructSimpleCrossStructCrossFieldFailureParallel(b *testing.B) {
	validate := New()

	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `validate:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)

	inner := &Inner{
		Start: then,
	}

	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(outer)
		}
	})
}

func BenchmarkStructSimpleSuccess(b *testing.B) {
	validate := New()
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(validFoo)
	}
}

func BenchmarkStructSimpleSuccessParallel(b *testing.B) {
	validate := New()
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}
	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(validFoo)
		}
	})
}

func BenchmarkStructSimpleFailure(b *testing.B) {
	validate := New()
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}

	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(invalidFoo)
	}
}

func BenchmarkStructSimpleFailureParallel(b *testing.B) {
	validate := New()
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}

	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(invalidFoo)
		}
	})
}

func BenchmarkStructComplexSuccess(b *testing.B) {
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

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(tSuccess)
	}
}

func BenchmarkStructComplexSuccessParallel(b *testing.B) {
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

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(tSuccess)
		}
	})
}

func BenchmarkStructComplexFailure(b *testing.B) {
	validate := New()
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

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(tFail)
	}
}

func BenchmarkStructComplexFailureParallel(b *testing.B) {
	validate := New()
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

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(tFail)
		}
	})
}

type TestOneof struct {
	Color string `validate:"oneof=red green"`
}

func BenchmarkOneof(b *testing.B) {
	w := &TestOneof{Color: "green"}
	val := New()
	for i := 0; i < b.N; i++ {
		_ = val.Struct(w)
	}
}

func BenchmarkOneofParallel(b *testing.B) {
	w := &TestOneof{Color: "green"}
	val := New()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = val.Struct(w)
		}
	})
}

type TestNoneOf struct {
	Color string `validate:"noneof=red green"`
}

func BenchmarkNoneOf(b *testing.B) {
	w := &TestNoneOf{Color: "blue"}
	val := New()
	for i := 0; i < b.N; i++ {
		_ = val.Struct(w)
	}
}

func BenchmarkNoneOfParallel(b *testing.B) {
	w := &TestNoneOf{Color: "blue"}
	val := New()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = val.Struct(w)
		}
	})
}

func BenchmarkFieldBooleanValid(b *testing.B) {
	validate := New()
	s := "true"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "boolean")
	}
}

func BenchmarkFieldBooleanValidParallel(b *testing.B) {
	validate := New()
	s := "true"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(&s, "boolean")
		}
	})
}

func BenchmarkFieldBooleanInvalid(b *testing.B) {
	validate := New()
	s := "not-a-boolean"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "boolean")
	}
}

func BenchmarkFieldBooleanInvalidParallel(b *testing.B) {
	validate := New()
	s := "not-a-boolean"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(&s, "boolean")
		}
	})
}

func BenchmarkFieldOmitEmptyPointer(b *testing.B) {
	validate := New()
	s := "value"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "omitempty")
	}
}

func BenchmarkFieldOmitEmptyPointerParallel(b *testing.B) {
	validate := New()
	s := "value"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(&s, "omitempty")
		}
	})
}

func BenchmarkFieldOneOfHit(b *testing.B) {
	validate := New()
	s := "green"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "oneof=red green blue")
	}
}

func BenchmarkFieldOneOfHitParallel(b *testing.B) {
	validate := New()
	s := "green"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(&s, "oneof=red green blue")
		}
	})
}

func BenchmarkFieldOneOfMiss(b *testing.B) {
	validate := New()
	s := "black"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "oneof=red green blue")
	}
}

func BenchmarkFieldOneOfMissParallel(b *testing.B) {
	validate := New()
	s := "black"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(&s, "oneof=red green blue")
		}
	})
}

type T struct{}

func (*T) Validate() error { return errors.New("ops") }

func BenchmarkValidateFnSequential(b *testing.B) {
	validate := New()

	type Test struct {
		T T `validate:"validateFn"`
	}

	test := &Test{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(test)
	}
}

func BenchmarkValidateFnParallel(b *testing.B) {
	validate := New()

	type Test struct {
		T T `validate:"validateFn"`
	}

	test := &Test{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(test)
		}
	})
}

func BenchmarkVarString(b *testing.B) {
	validate := New()
	s := "1"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(s, "len=1")
	}
}

func BenchmarkVarStringParallel(b *testing.B) {
	validate := New()
	s := "1"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Var(s, "len=1")
		}
	})
}

func BenchmarkStructCacheHit(b *testing.B) {
	validate := New()
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}
	_ = validate.Struct(validFoo) // warm the struct cache

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Struct(validFoo)
	}
}

func BenchmarkStructCacheHitParallel(b *testing.B) {
	validate := New()
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}
	_ = validate.Struct(validFoo) // warm the struct cache

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = validate.Struct(validFoo)
		}
	})
}

func BenchmarkMapKeyString(b *testing.B) {
	private := reflect.ValueOf(&struct{ data map[any]int }{map[any]int{"key": 0}}).Elem().Field(0).MapKeys()[0]
	for _, tc := range []struct {
		name string
		key  reflect.Value
	}{
		{"string", reflect.ValueOf("key")},
		{"int", reflect.ValueOf(123)},
		{"stringer", reflect.ValueOf(mapStringerKey("key"))},
		{"formatter", reflect.ValueOf(mapFormatterKey(7))},
		{"struct", reflect.ValueOf(struct{ ID int }{1})},
		{"private-interface", private},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for b.Loop() {
				_ = mapKeyString(tc.key)
			}
		})
	}
}

func BenchmarkVarWithKeyFailure(b *testing.B) {
	v := New()
	for b.Loop() {
		_ = v.VarWithKey("name", "", "required")
	}
}

func BenchmarkOneOfOptions(b *testing.B) {
	options := make([]string, 0, 129)
	for i := range 128 {
		options = append(options, "option"+strconv.Itoa(i))
	}
	options = append(options, "blue")
	for _, tc := range []struct {
		name  string
		param string
		value string
	}{
		{"short", "red green blue", "blue"},
		{"short-first", "red green blue", "red"},
		{"many", strings.Join(options, " "), "blue"},
		{"many-first", strings.Join(options, " "), options[0]},
		{"many-miss", strings.Join(options, " "), "missing"},
		{"long-token", strings.Repeat("x", 1024) + " green blue", "blue"},
	} {
		b.Run(tc.name, func(b *testing.B) {
			v := New()
			tag := "oneof=" + tc.param
			var value any = tc.value
			for b.Loop() {
				_ = v.Var(value, tag)
			}
		})
	}
}

type benchmarkSignupRequest struct {
	Name         string   `validate:"required,min=2,max=80"`
	Email        string   `validate:"required,email"`
	Age          int      `validate:"gte=18,lte=120"`
	Password     string   `validate:"required,min=12"`
	Confirmation string   `validate:"eqfield=Password"`
	Country      string   `validate:"oneof=US GB DE IN CA"`
	Marketing    *bool    `validate:"omitempty"`
	Topics       []string `validate:"max=5,dive,oneof=news releases offers"`
}

type benchmarkOrderRequest struct {
	CustomerID string `validate:"required,uuid4"`
	Currency   string `validate:"oneof=USD EUR GBP"`
	Shipping   struct {
		Street  string `validate:"required,max=120"`
		City    string `validate:"required,max=80"`
		Postal  string `validate:"required,max=12"`
		Country string `validate:"len=2"`
	}
	Items []struct {
		SKU      string  `validate:"required,alphanum"`
		Quantity int     `validate:"gte=1,lte=100"`
		Price    float64 `validate:"gt=0"`
	} `validate:"required,min=1,max=100,dive"`
	Metadata    map[string]string `validate:"max=10,dive,keys,required,endkeys,required,max=80"`
	GiftMessage *string           `validate:"omitempty,max=200"`
}

func BenchmarkRequestSignup(b *testing.B) {
	benchmarkRequest[benchmarkSignupRequest](b, []byte(`{
		"Name":"Sample User","Email":"sample@example.com","Age":32,
		"Password":"sample-password","Confirmation":"sample-password",
		"Country":"US","Marketing":true,"Topics":["news","releases"]
	}`), func(r *benchmarkSignupRequest) {
		r.Email = "invalid"
		r.Age = 15
		r.Confirmation = "different"
	})
}

func BenchmarkRequestOrder(b *testing.B) {
	benchmarkRequest[benchmarkOrderRequest](b, []byte(`{
		"CustomerID":"550e8400-e29b-41d4-a716-446655440000","Currency":"USD",
		"Shipping":{"Street":"123 Sample Street","City":"London","Postal":"SW1A 1AA","Country":"GB"},
		"Items":[
			{"SKU":"SKU001","Quantity":2,"Price":19.95},
			{"SKU":"SKU002","Quantity":1,"Price":49.50},
			{"SKU":"SKU003","Quantity":3,"Price":5.25},
			{"SKU":"SKU004","Quantity":1,"Price":99.00},
			{"SKU":"SKU005","Quantity":2,"Price":12.00},
			{"SKU":"SKU006","Quantity":1,"Price":8.75},
			{"SKU":"SKU007","Quantity":4,"Price":3.50},
			{"SKU":"SKU008","Quantity":1,"Price":24.99}
		],
		"Metadata":{"source":"web","campaign":"summer","locale":"en-GB"},
		"GiftMessage":"Happy birthday"
	}`), func(r *benchmarkOrderRequest) {
		r.Shipping.Postal = ""
		r.Items[0].Quantity = 0
		r.Items[3].Price = -1
		r.Metadata["source"] = ""
	})
}

// benchmarkRequest models a reused validator with 90% valid requests and 10%
// invalid requests. DecodeValidate includes JSON decoding, but no HTTP or I/O.
func benchmarkRequest[T any](b *testing.B, validJSON []byte, invalidate func(*T)) {
	b.Helper()
	v := New(WithRequiredStructEnabled())
	var requests [10]T
	var payloads [10][]byte
	for i := range requests {
		if err := json.Unmarshal(validJSON, &requests[i]); err != nil {
			b.Fatal(err)
		}
		if i == len(requests)-1 {
			invalidate(&requests[i])
		}
		payload, err := json.Marshal(&requests[i])
		if err != nil {
			b.Fatal(err)
		}
		payloads[i] = payload
		if err := v.Struct(&requests[i]); (err != nil) != (i == len(requests)-1) {
			b.Fatalf("request %d has unexpected validation result: %v", i, err)
		}
	}

	for _, mode := range []string{"Validate", "DecodeValidate"} {
		b.Run(mode, func(b *testing.B) {
			run := func(i int) {
				if mode == "DecodeValidate" {
					var request T
					if err := json.Unmarshal(payloads[i], &request); err != nil {
						b.Error(err)
						return
					}
					_ = v.Struct(&request)
					return
				}
				_ = v.Struct(&requests[i])
			}
			b.Run("Serial", func(b *testing.B) {
				b.ReportAllocs()
				i := 0
				for b.Loop() {
					run(i % len(requests))
					i++
				}
			})
			b.Run("Parallel", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					i := 0
					for pb.Next() {
						run(i % len(requests))
						i++
					}
				})
			})
		})
	}
}
