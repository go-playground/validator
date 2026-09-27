package validator

import "testing"

func TestAnyFieldWrongDynamicTypeReturnsErrorNotPanic(t *testing.T) {
	v := New()
	type Profile struct {
		Website any `validate:"url"`
		Locale  any `validate:"bcp47_language_tag"`
		Zone    any `validate:"timezone"`
	}

	// null-like / missing dynamic values already error; wrong kinds must too
	cases := []Profile{
		{Website: 12345},
		{Locale: true},
		{Zone: map[string]any{"tz": "UTC"}},
		{Website: "https://example.com", Locale: "en-US", Zone: 1},
	}
	for i, c := range cases {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("case %d panicked: %v", i, r)
				}
			}()
			err := v.Struct(c)
			if err == nil {
				t.Fatalf("case %d: expected validation error, got nil", i)
			}
		}()
	}

	// valid strings still pass
	ok := Profile{Website: "https://example.com", Locale: "en-US", Zone: "UTC"}
	if err := v.Struct(ok); err != nil {
		t.Fatalf("valid profile failed: %v", err)
	}
}

func TestStaticWrongTypeStillPanics(t *testing.T) {
	v := New()
	type Bad struct {
		Website int `validate:"url"`
	}
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for mistyped static field")
		}
	}()
	_ = v.Struct(Bad{Website: 1})
}
