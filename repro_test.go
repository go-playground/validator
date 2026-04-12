package validator

import (
	"testing"
)

// TestNilPointerComparisonPanic tests that comparison validators (gte, gt, lte, lt, eq, len)
// do not panic when the field is a nil pointer. This can occur when a conditional
// required tag (e.g. required_if) passes but the field is nil, and a subsequent
// comparison tag runs on the nil pointer value.
// See: https://github.com/go-playground/validator/issues/907
func TestNilPointerComparisonPanic(t *testing.T) {
	validate := New()

	t.Run("nil *int64 with gte", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,gte=2"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("nil *int64 with gt", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,gt=0"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("nil *int64 with lte", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,lte=100"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("nil *int64 with lt", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,lt=100"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("nil *int64 with eq", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,eq=5"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("nil *int64 with len", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,len=5"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("nil *float64 with gte", func(t *testing.T) {
		type S struct {
			Type  string   `validate:"required"`
			Price *float64 `validate:"required_if=Type special,gte=0.01"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("nil *string with gte", func(t *testing.T) {
		type S struct {
			Type string  `validate:"required"`
			Name *string `validate:"required_if=Type special,gte=3"`
		}
		s := &S{Type: "notspecial"}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("non-nil *int64 with gte passes", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,gte=2"`
		}
		q := int64(5)
		s := &S{Type: "notspecial", Quantity: &q}
		err := validate.Struct(s)
		if err != nil {
			t.Fatalf("unexpected validation error: %v", err)
		}
	})

	t.Run("non-nil *int64 with gte fails", func(t *testing.T) {
		type S struct {
			Type     string `validate:"required"`
			Quantity *int64 `validate:"required_if=Type special,gte=10"`
		}
		q := int64(5)
		s := &S{Type: "notspecial", Quantity: &q}
		err := validate.Struct(s)
		if err == nil {
			t.Fatal("expected validation error")
		}
	})
}
