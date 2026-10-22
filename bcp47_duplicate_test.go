package validator

import "testing"

func TestBCP47StrictDuplicateVariants(t *testing.T) {
	validate := New()
	for _, value := range []string{
		"de-1901-1901", "de-DE-1901-1901", "sl-rozaj-rozaj",
		"sl-rozaj-biske-rozaj", "sl-ROZAJ-rozaj",
	} {
		t.Run(value, func(t *testing.T) {
			if err := validate.Var(value, "bcp47_strict_language_tag"); err == nil {
				t.Errorf("duplicate variant accepted: %s", value)
			}
		})
	}
	for _, value := range []string{
		"de-DE-1901", "sl-rozaj-biske", "sl-ROZAJ-biske", "de-1901-x-1901", "de-x-rozaj-rozaj",
	} {
		t.Run(value, func(t *testing.T) {
			if err := validate.Var(value, "bcp47_strict_language_tag"); err != nil {
				t.Errorf("valid tag rejected: %s: %v", value, err)
			}
		})
	}
}
