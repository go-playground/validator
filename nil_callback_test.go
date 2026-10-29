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
	type converted struct{}
	v := New()
	v.RegisterCustomTypeFunc(func(reflect.Value) any { return nil }, converted{})
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
	v.RegisterAlias("nilacceptalias", "ordinary|accept")
	v.RegisterAlias("nilrejectalias", "ordinary|reject")
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
		{"accept|ordinary", "accept", ""},
		{"ordinary|accept", "accept", ""},
		{"ordinary|reject|accept", "reject,accept", ""},
		{"ordinary|reject", "reject", "ordinary|reject"},
		{"ordinary|ordinary", "", "ordinary|ordinary"},
		{"ordinary,accept", "", "ordinary"},
		{"ordinary|accept,reject", "accept,reject", "reject"},
		{"accept,ordinary|accept", "accept,accept", ""},
		{"nilacceptalias", "accept", ""},
		{"nilrejectalias", "reject", "nilrejectalias"},
		{"nilacceptalias,reject", "accept,reject", "reject"},
		{"reject|ordinary", "reject", "reject|ordinary"},
		{"reject|reject", "reject,reject", "reject|reject"},
		{"nilalias", "reject,reject", "nilalias"},
		{"accept,required", "accept", "required"},
		{"accept,omitempty,ordinary", "accept", ""},
		{"accept,omitzero,ordinary", "accept", ""},
		{"accept,isdefault", "accept", ""},
		{"accept,isdefault,reject", "accept,reject", "reject"},
		{"accept,omitnil,ordinary", "accept", ""},
		{"omitnil,ordinary", "", ""},
		{"omitzero,ordinary", "", ""},
		{"isdefault", "", ""},
		{"", "", ""},
		{"-", "", ""},
		{"omitempty,ordinary", "", ""},
	} {
		for source, value := range []any{nil, converted{}, &converted{}, nilCallbackValuer{}, &nilCallbackValuer{}} {
			t.Run(fmt.Sprintf("%s/source=%d", tc.tag, source), func(t *testing.T) {
				calls = nil
				err := v.Var(value, tc.tag)
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

func TestNilCallbackStructContext(t *testing.T) {
	type contextKey struct{}
	type converted struct{ Value any }
	type payload struct {
		Values []converted `validate:"dive,nilalias" json:"values"`
	}
	for _, accepted := range []bool{false, true} {
		t.Run(fmt.Sprintf("accepted=%t", accepted), func(t *testing.T) {
			v := New()
			v.RegisterCustomTypeFunc(func(field reflect.Value) any { return field.Interface().(converted).Value }, converted{})
			v.RegisterTagNameFunc(func(field reflect.StructField) string { return field.Tag.Get("json") })
			v.RegisterAlias("nilalias", "ordinary|nilcheck=expected")
			if err := v.RegisterValidation("ordinary", func(fl FieldLevel) bool {
				if !fl.Field().IsValid() {
					t.Fatal("ordinary validator called on invalid value")
				}
				return true
			}); err != nil {
				t.Fatal(err)
			}
			input := payload{Values: []converted{{Value: nil}, {Value: "present"}}}
			calls := 0
			if err := v.RegisterValidationCtx("nilcheck", func(ctx context.Context, fl FieldLevel) bool {
				calls++
				if ctx.Value(contextKey{}) != accepted || fl.Field().IsValid() || fl.Param() != "expected" {
					t.Error("callback context, value, or parameter is incorrect")
				}
				if fl.FieldName() != "values[0]" || fl.StructFieldName() != "Values[0]" || !reflect.DeepEqual(fl.Parent().Interface(), input) || !reflect.DeepEqual(fl.Top().Interface(), input) {
					t.Error("callback field names, parent, or top is incorrect")
				}
				return accepted
			}, true); err != nil {
				t.Fatal(err)
			}
			err := v.StructCtx(context.WithValue(context.Background(), contextKey{}, accepted), input)
			if calls != 1 {
				t.Fatalf("callback calls = %d, want 1", calls)
			}
			if accepted {
				if err != nil {
					t.Fatal(err)
				}
				return
			}
			errs, ok := err.(ValidationErrors)
			if !ok || len(errs) != 1 {
				t.Fatalf("expected one validation error, got %v", err)
			}
			fe := errs[0]
			if fe.Tag() != "nilalias" || fe.ActualTag() != "ordinary|nilcheck=expected" || fe.Param() != "expected" || fe.Namespace() != "payload.values[0]" || fe.StructNamespace() != "payload.Values[0]" || fe.Field() != "values[0]" || fe.StructField() != "Values[0]" || fe.Kind() != reflect.Invalid || fe.Type() != nil || fe.Value() != nil || fe.Error() == "" {
				t.Errorf("incorrect alias error metadata: %#v", fe)
			}
		})
	}
}
