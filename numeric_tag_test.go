package validator

import "testing"

func TestNumericTagParametersUseDecimalUnlessPrefixed(t *testing.T) {
	validate := New()

	type maxValue struct {
		Value int `validate:"max=010"`
	}
	for _, value := range []int{9, 10} {
		if err := validate.Struct(maxValue{Value: value}); err != nil {
			t.Fatalf("max=010 rejected %d: %v", value, err)
		}
	}
	if err := validate.Struct(maxValue{Value: 11}); err == nil {
		t.Fatal("max=010 accepted 11")
	}

	type minValue struct {
		Value int `validate:"min=010"`
	}
	if err := validate.Struct(minValue{Value: 9}); err == nil {
		t.Fatal("min=010 accepted 9")
	}
	if err := validate.Struct(minValue{Value: 10}); err != nil {
		t.Fatalf("min=010 rejected 10: %v", err)
	}

	type length struct {
		Value string `validate:"len=08"`
	}
	if err := validate.Struct(length{Value: "12345678"}); err != nil {
		t.Fatalf("len=08 rejected an eight-character value: %v", err)
	}
	if err := validate.Struct(length{Value: "1234567"}); err == nil {
		t.Fatal("len=08 accepted a seven-character value")
	}

	// An explicit prefix still selects the non-decimal base.
	type hex struct {
		Value int `validate:"max=0x10"`
	}
	if err := validate.Struct(hex{Value: 16}); err != nil {
		t.Fatalf("max=0x10 rejected 16: %v", err)
	}
	if err := validate.Struct(hex{Value: 17}); err == nil {
		t.Fatal("max=0x10 accepted 17")
	}
}
