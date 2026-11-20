package validator

import "testing"

func TestIso4217NumericIntegerRange(t *testing.T) {
	type signedCurrency int64
	type unsignedCurrency uint64
	v := New()
	cases := []struct {
		name  string
		value any
		valid bool
	}{
		{"int", int(840), true},
		{"int8", int8(36), true},
		{"int16", int16(840), true},
		{"int32", int32(840), true},
		{"int64", int64(840), true},
		{"uint", uint(840), true},
		{"uint8", uint8(36), true},
		{"uint16", uint16(840), true},
		{"uint32", uint32(840), true},
		{"uint64", uint64(840), true},
		{"named_signed", signedCurrency(840), true},
		{"named_unsigned", unsignedCurrency(840), true},
		{"unknown", int64(13), false},
		{"zero", int64(0), false},
		{"highest_code", int64(999), true},
		{"highest_unsigned_code", uint64(999), true},
		{"negative", int64(-840), false},
		{"above_three_digits", int64(1000), false},
		{"unsigned_above_three_digits", uint64(1000), false},
		{"signed_wrap", int64(1<<32 + 840), false},
		{"unsigned_wrap", uint64(1<<32 + 840), false},
		{"negative_wrap", int64(-(1 << 32) + 840), false},
		{"named_signed_wrap", signedCurrency(1<<32 + 840), false},
		{"named_unsigned_wrap", unsignedCurrency(1<<32 + 840), false},
		{"max_int64", int64(1<<63 - 1), false},
		{"min_int64", int64(-1 << 63), false},
		{"max_uint64", uint64(1<<64 - 1), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("Var", func(t *testing.T) {
				err := v.Var(tc.value, "iso4217_numeric")
				if (err == nil) != tc.valid {
					t.Fatalf("value=%v accepted=%t want=%t", tc.value, err == nil, tc.valid)
				}
			})
			t.Run("Struct", func(t *testing.T) {
				input := struct {
					Currency any `validate:"iso4217_numeric"`
				}{Currency: tc.value}
				err := v.Struct(input)
				if (err == nil) != tc.valid {
					t.Fatalf("value=%v accepted=%t want=%t", tc.value, err == nil, tc.valid)
				}
			})
		})
	}
}
