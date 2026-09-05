package validator

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNilConverterCallback(t *testing.T) {
	type custom struct{}
	type contextKey struct{}
	ctx := context.WithValue(context.Background(), contextKey{}, "request-892")
	for _, source := range []string{"typed", "plain", "custom", "valuer", "valuer_pointer", "custom_pointer"} {
		converted := source != "typed"
		for _, enabled := range []bool{false, true} {
			for _, accepted := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/enabled=%t/accepted=%t", source, enabled, accepted), func(t *testing.T) {
					v := New()
					v.RegisterCustomTypeFunc(func(reflect.Value) interface{} { return nil }, custom{})
					calls := 0
					err := v.RegisterValidationCtx("nilcheck", func(got context.Context, fl FieldLevel) bool {
						calls++
						if got.Value(contextKey{}) != "request-892" {
							t.Error("context value did not reach callback")
						}
						wantKind := reflect.Ptr
						if converted {
							wantKind = reflect.Invalid
						}
						if fl.Field().Kind() != wantKind {
							t.Errorf("field kind = %v, want %v", fl.Field().Kind(), wantKind)
						}
						return accepted
					}, enabled)
					if err != nil {
						t.Fatal(err)
					}
					var value interface{} = (*string)(nil)
					switch source {
					case "plain":
						value = nil
					case "custom":
						value = custom{}
					case "custom_pointer":
						value = &custom{}
					case "valuer":
						value = nilCallbackValuer{}
					case "valuer_pointer":
						value = &nilCallbackValuer{}
					}
					err = v.VarCtx(ctx, value, "nilcheck")
					wantCalls := 0
					if enabled {
						wantCalls = 1
					}
					if calls != wantCalls {
						t.Errorf("callback calls = %d, want %d; error = %v", calls, wantCalls, err)
					}
					if enabled && accepted {
						if err != nil {
							t.Errorf("accepted value: %v", err)
						}
						return
					}
					errs, ok := err.(ValidationErrors)
					if !ok || len(errs) != 1 {
						t.Fatalf("expected one validation error, got %v", err)
					}
					fe := errs[0]
					if fe.Tag() != "nilcheck" || fe.ActualTag() != "nilcheck" || fe.Error() == "" {
						t.Errorf("unusable validation error: %v", fe)
					}
					if converted {
						if fe.Kind() != reflect.Invalid || fe.Type() != nil || fe.Value() != nil {
							t.Errorf("unexpected converted-nil error metadata: kind=%v type=%v value=%v", fe.Kind(), fe.Type(), fe.Value())
						}
					} else if fe.Kind() != reflect.Ptr || fe.Type() != reflect.TypeOf((*string)(nil)) || fe.Value() != (*string)(nil) {
						t.Errorf("unexpected typed-nil metadata: %v", fe)
					}
				})
			}
		}
	}
}

type nilCallbackValuer struct{}

func (nilCallbackValuer) ValidatorValue() any { return nil }

func TestNilCallbackChains(t *testing.T) {
	v := New()
	calls := []string{}
	for _, tag := range []string{"accept", "reject", "ordinary"} {
		err := v.RegisterValidation(tag, func(fl FieldLevel) bool {
			calls = append(calls, fl.GetTag())
			return fl.GetTag() != "reject"
		}, tag != "ordinary")
		if err != nil {
			t.Fatal(err)
		}
	}
	v.RegisterAlias("nilalias", "reject|reject")
	for _, tc := range []struct {
		tag       string
		wantCalls string
		wantTag   string
	}{
		{"accept,accept", "accept,accept", ""},
		{"accept,reject", "accept,reject", "reject"},
		{"reject,accept", "reject", "reject"},
		{"accept,ordinary", "accept", "ordinary"},
		{"reject|accept", "reject,accept", ""},
		{"reject|ordinary", "reject", "reject|ordinary"},
		{"reject|reject", "reject,reject", "reject|reject"},
		{"nilalias", "reject,reject", "nilalias"},
		{"accept,required", "accept", "required"},
		{"accept,omitempty,ordinary", "accept", ""},
		{"accept,omitzero,ordinary", "accept", ""},
		{"omitempty,ordinary", "", ""},
	} {
		t.Run(tc.tag, func(t *testing.T) {
			calls = nil
			err := v.Var(nilCallbackValuer{}, tc.tag)
			if strings.Join(calls, ",") != tc.wantCalls {
				t.Errorf("calls = %v, want %s", calls, tc.wantCalls)
			}
			if tc.wantTag == "" {
				if err != nil {
					t.Fatal(err)
				}
				return
			}
			errs, ok := err.(ValidationErrors)
			if !ok || len(errs) != 1 {
				t.Fatalf("expected one error, got %v", err)
			}
			fe := errs[0]
			if fe.Tag() != tc.wantTag || fe.Kind() != reflect.Invalid || fe.Type() != nil || fe.Value() != nil || fe.Error() == "" {
				t.Errorf("invalid error: %v", fe)
			}
		})
	}
}

func TestNilConditionalCallbacks(t *testing.T) {
	type converted struct{}
	v := New()
	v.RegisterCustomTypeFunc(func(reflect.Value) any { return nil }, converted{})
	for _, tc := range []struct {
		tag       string
		absentOK  bool
		presentOK bool
	}{
		{"required_if=Other yes", true, false},
		{"required_unless=Other yes", false, true},
		{"required_with=Other", true, false},
		{"required_with_all=Other", true, false},
		{"required_without=Other", false, true},
		{"required_without_all=Other", false, true},
		{"excluded_if=Other yes", true, true},
		{"excluded_unless=Other yes", true, true},
		{"excluded_with=Other", true, true},
		{"excluded_with_all=Other", true, true},
		{"excluded_without=Other", true, true},
		{"excluded_without_all=Other", true, true},
		{"skip_unless=Other yes", true, false},
	} {
		for _, present := range []bool{false, true} {
			for source, value := range []any{(*string)(nil), nil, converted{}, &converted{}, nilCallbackValuer{}, &nilCallbackValuer{}} {
				t.Run(fmt.Sprintf("%s/present=%t/source=%d", tc.tag, present, source), func(t *testing.T) {
					st := reflect.StructOf([]reflect.StructField{
						{Name: "Other", Type: reflect.TypeOf("")},
						{Name: "Value", Type: reflect.TypeOf((*any)(nil)).Elem(), Tag: reflect.StructTag(`validate:"` + tc.tag + `" json:"value"`)},
					})
					instance := reflect.New(st).Elem()
					if present {
						instance.Field(0).SetString("yes")
					}
					if value != nil {
						instance.Field(1).Set(reflect.ValueOf(value))
					}
					err := v.Struct(instance.Interface())
					wantOK := tc.absentOK
					if present {
						wantOK = tc.presentOK
					}
					if (err == nil) != wantOK {
						t.Fatalf("error = %v, want success %t", err, wantOK)
					}
					if err != nil {
						fe := err.(ValidationErrors)[0]
						if fe.Field() != "Value" || fe.StructField() != "Value" || fe.Namespace() != "Value" || fe.StructNamespace() != "Value" || fe.Param() != strings.SplitN(tc.tag, "=", 2)[1] || fe.Error() == "" {
							t.Errorf("invalid error metadata: %v", fe)
						}
					}
				})
			}
		}
	}
}
