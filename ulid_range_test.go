package validator_test

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestULID128BitRange(t *testing.T) {
	validate := validator.New()
	// The ULID specification limits the first Base32 digit to 0-7.
	for _, first := range "0123456789ABCDEFGHJKMNPQRSTVWXYZabcdefghjkmnpqrstvwxyz" {
		value := string(first) + strings.Repeat("Z", 25)
		wantValid := first >= '0' && first <= '7'
		t.Run(string(first), func(t *testing.T) {
			for name, err := range map[string]error{
				"Var": validate.Var(value, "ulid"),
				"Struct": validate.Struct(struct {
					ID string `validate:"ulid"`
				}{ID: value}),
			} {
				if (err == nil) != wantValid {
					t.Errorf("%s(%q): error = %v, want valid = %v", name, value, err, wantValid)
				}
			}
		})
	}
	for _, tc := range []struct {
		name  string
		value string
		valid bool
	}{
		{"minimum", strings.Repeat("0", 26), true},
		{"maximum_lowercase", "7" + strings.Repeat("z", 25), true},
		{"mixed_case", "01bX5zzKBkACTav9weVGEmMvrZ", true},
		{"short", "7" + strings.Repeat("Z", 24), false},
		{"long", "7" + strings.Repeat("Z", 26), false},
		{"excluded_I", "7" + strings.Repeat("Z", 24) + "I", false},
		{"excluded_L", "7" + strings.Repeat("Z", 24) + "L", false},
		{"excluded_O", "7" + strings.Repeat("Z", 24) + "O", false},
		{"excluded_U", "7" + strings.Repeat("Z", 24) + "U", false},
		{"empty", "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for name, err := range map[string]error{
				"Var": validate.Var(tc.value, "ulid"),
				"Struct": validate.Struct(struct {
					ID string `validate:"ulid"`
				}{ID: tc.value}),
			} {
				if (err == nil) != tc.valid {
					t.Errorf("%s(%q): error = %v, want valid = %v", name, tc.value, err, tc.valid)
				}
			}
		})
	}
}
