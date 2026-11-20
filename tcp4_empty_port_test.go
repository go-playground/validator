package validator

import "testing"

func TestTCP4AddrRequiresPort(t *testing.T) {
	validate := New()
	for _, tc := range []struct {
		name  string
		value string
		valid bool
	}{
		{"empty_port", "127.0.0.1:", false},
		{"empty_port_unspecified_host", "0.0.0.0:", false},
		{"missing_port", "127.0.0.1", false},
		{"explicit_zero", "127.0.0.1:0", true},
		{"numeric_port", "127.0.0.1:8080", true},
		{"maximum_port", "127.0.0.1:65535", true},
		{"overflow_port", "127.0.0.1:65536", false},
		{"negative_port", "127.0.0.1:-1", false},
		{"invalid_ip", "999.0.0.1:8080", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := validate.Var(tc.value, "tcp4_addr"); (err == nil) != tc.valid {
				t.Fatalf("Var(%q, tcp4_addr) = %v; valid = %v", tc.value, err, tc.valid)
			}
			value := struct {
				Address string `validate:"tcp4_addr"`
			}{tc.value}
			if err := validate.Struct(value); (err == nil) != tc.valid {
				t.Fatalf("Struct(%q) = %v; valid = %v", tc.value, err, tc.valid)
			}
		})
	}
}
