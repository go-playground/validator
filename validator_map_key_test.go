package validator

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestMapKeyStringPointerFormatting(t *testing.T) {
	for _, key := range []any{
		&struct{ ID int }{1}, &[2]int{1, 2}, &[]int{1, 2}, &map[string]int{"id": 1},
	} {
		for _, value := range []reflect.Value{reflect.ValueOf(key), reflect.ValueOf(map[any]int{key: 0}).MapKeys()[0]} {
			if got, want := mapKeyString(value), fmt.Sprintf("%v", value); got != want {
				t.Errorf("key type %T (%s): got %q, want %q", key, value.Kind(), got, want)
			}
		}
	}
}

func TestMapDiveInterfacePointerKeyNamespaces(t *testing.T) {
	for _, tc := range []struct {
		name          string
		first, second any
	}{
		{"struct", &struct{ ID int }{1}, &struct{ ID int }{1}},
		{"array", &[2]int{1, 2}, &[2]int{1, 2}},
		{"slice", &[]int{1, 2}, &[]int{1, 2}},
		{"map", &map[string]int{"id": 1}, &map[string]int{"id": 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := New().Var(map[any]int{tc.first: 0, tc.second: 0}, "dive,gt=0")
			if err == nil {
				t.Fatal("expected validation errors")
			}
			errs := err.(ValidationErrors)
			want := []string{fmt.Sprintf("[%p]", tc.first), fmt.Sprintf("[%p]", tc.second)}
			var got []string
			for _, fe := range errs {
				got = append(got, fe.Namespace())
				if fe.StructNamespace() != fe.Namespace() {
					t.Errorf("StructNamespace() = %q, want %q", fe.StructNamespace(), fe.Namespace())
				}
			}
			slices.Sort(got)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("namespaces = %q, want %q", got, want)
			}
			if translated := errs.Translate(nil); len(translated) != 2 {
				t.Errorf("Translate() returned %d entries, want 2", len(translated))
			}
		})
	}
}
