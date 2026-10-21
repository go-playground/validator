package validator

import "testing"

func TestUniqueUnhashableValues(t *testing.T) {
	type item struct{ Tags []string }
	first := []byte{1}
	second := []byte{1}
	for _, tc := range []struct {
		name   string
		value  any
		tag    string
		unique bool
	}{
		{"slices distinct", [][]string{{"a"}, {"b"}}, "unique", true},
		{"slices repeated", [][]string{{"a"}, {"a"}}, "unique", false},
		{"arrays", [2][]byte{{1}, {1}}, "unique", false},
		{"bytes distinct", [][]byte{{1}, {2}}, "unique", true},
		{"structs", []item{{[]string{"a"}}, {[]string{"a"}}}, "unique", false},
		{"struct fields", []item{{[]string{"a"}}, {[]string{"b"}}}, "unique=Tags", true},
		{"struct fields repeated", []item{{[]string{"a"}}, {[]string{"a"}}}, "unique=Tags", false},
		{"slice pointers", []*[]byte{&first, &second}, "unique", false},
		{"nil interfaces", []any{nil, nil}, "unique", false},
		{"maps", []map[string]string{{"a": "1"}, {"a": "1"}}, "unique", false},
		{"map values", map[string][]byte{"a": {1}, "b": {2}}, "unique", true},
		{"map values repeated", map[string][]byte{"a": {1}, "b": {1}}, "unique", false},
		{"interfaces", []any{"a", []string{"a"}, 1}, "unique", true},
		{"interfaces repeated", []any{[]string{"a"}, "a", []string{"a"}}, "unique", false},
		{"struct with interface", []struct{ Value any }{{[]string{"a"}}, {[]string{"a"}}}, "unique", false},
		{"nil and empty", [][]byte{nil, {}}, "unique", true},
		{"nil repeated", [][]byte{nil, nil}, "unique", false},
		{"comparable", []string{"a", "b"}, "unique", true},
		{"comparable repeated", []string{"a", "a"}, "unique", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Fatalf("Var() panicked: %v", recovered)
				}
			}()
			v := New()
			err := v.Var(tc.value, tc.tag)
			if (err == nil) != tc.unique {
				t.Fatalf("Var() error = %v, want unique %v", err, tc.unique)
			}
			if err != nil {
				if _, ok := err.(ValidationErrors); !ok {
					t.Fatalf("Var() error type = %T, want ValidationErrors", err)
				}
			}
		})
	}
}

func TestStructUniqueUnhashableValues(t *testing.T) {
	v := New()
	type payload struct {
		Values [][]string `validate:"unique"`
	}
	if err := v.Struct(payload{[][]string{{"a"}, {"b"}}}); err != nil {
		t.Fatalf("Struct() error = %v, want nil", err)
	}
	if err := v.Struct(payload{[][]string{{"a"}, {"a"}}}); err == nil {
		t.Fatal("Struct() error = nil, want duplicate validation error")
	}
}

func BenchmarkUniqueComparableValues(b *testing.B) {
	v := New()
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ReportAllocs()
	for b.Loop() {
		if err := v.Var(values, "unique"); err != nil {
			b.Fatal(err)
		}
	}
}
