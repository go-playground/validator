package validator

import (
	"bytes"
	sql "database/sql/driver"
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
