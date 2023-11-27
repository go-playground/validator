package vi

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterDefaultTranslations registers a set of default translations
// for all built in tag's in validator; you may add your own as desired.
func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "required",
			translation: "{0} không được bỏ trống",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "{0} phải có độ dài là {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("len-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} phải bằng {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} phải chứa {1}", false); err != nil {
					return
				}
				// if err = ut.AddCardinal("len-items-item", "{0} phần tử", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-items-item", "{0} phần tử", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-items", fe.Field(), c)

				default:
					t, err = ut.T("len-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("min-string", "{0} phải chứa ít nhất {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("min-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} phải bằng {1} hoặc lớn hơn", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} phải chứa ít nhất {1}", false); err != nil {
					return
				}
				// if err = ut.AddCardinal("min-items-item", "{0} phần tử", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-items-item", "{0} phần tử", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-items", fe.Field(), c)

				default:
					t, err = ut.T("min-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("max-string", "{0} chỉ được chứa tối đa {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("max-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} phải là {1} hoặc nhỏ hơn", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} chỉ được chứa tối đa {1}", false); err != nil {
					return
				}
				// if err = ut.AddCardinal("max-items-item", "{0} phần tử", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-items-item", "{0} phần tử", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-items", fe.Field(), c)

				default:
					t, err = ut.T("max-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0} không bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "{0} không được bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "{0} phải có độ dài nhỏ hơn {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} phải nhỏ hơn {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} chỉ được chứa ít hơn {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-items-item", "{0} phần tử", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-items-item", "{0} phần tử", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} phải nhỏ hơn Ngày & Giờ hiện tại", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' không thể dùng trên kiểu struct", fe.Tag())
						goto END
					}

					t, err = ut.T("lt-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lte-string", "{0} chỉ được có độ dài tối đa là {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} phải là {1} hoặc nhỏ hơn", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} chỉ được chứa nhiều nhất {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-items-item", "{0} phần tử", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-items-item", "{0} phần tử", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} chỉ được nhỏ hơn hoặc bằng Ngày & Giờ hiện tại", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' không thể dùng trên kiểu struct", fe.Tag())
						goto END
					}

					t, err = ut.T("lte-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gt-string", "{0} phải có độ dài lớn hơn {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} phải lớn hơn {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} phải chứa nhiều hơn {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-items-item", "{0} phần tử", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-items-item", "{0} phần tử", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} phải lớn hơn Ngày & Giờ hiện tại", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' không thể dùng trên kiểu struct", fe.Tag())
						goto END
					}

					t, err = ut.T("gt-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gte-string", "{0} phải có độ dài ít nhất {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} phải là {1} hoặc lớn hơn", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} phải chứa ít nhất {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-items-item", "{0} phần tử", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-items-item", "{0} phần tử", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} phải lớn hơn hoặc bằng Ngày & Giờ hiện tại", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' không thể dùng trên kiểu struct", fe.Tag())
						goto END
					}

					t, err = ut.T("gte-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0} phải bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0} phải bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0} không được phép bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0} phải lớn hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0} phải lớn hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0} chỉ được nhỏ hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0} chỉ được nhỏ hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0} không được phép bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0} phải lớn hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0} phải lớn hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0} chỉ được nhỏ hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0} chỉ được nhỏ hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0} chỉ được chứa ký tự dạng alphabetic",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} chỉ được chứa ký tự dạng alphanumeric",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} chỉ được chứa giá trị số hoặc số dưới dạng chữ",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} chỉ được chứa giá trị số",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} phải là giá trị hexadecimal",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} phải là giá trị HEX color",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} phải là giá trị RGB color",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} phải là giá trị RGBA color",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} phải là giá trị HSL color",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} phải là giá trị HSLA color",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} phải là giá trị số điện thoại theo định dạng E.164",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} phải là giá trị email address",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} phải là giá trị URL",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} phải là giá trị URI",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} phải là giá trị chuỗi Base64",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0} phải chứa chuỗi '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "{0} phải chứa ít nhất 1 trong cách ký tự sau '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0} không được chứa chuỗi '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0} không được chứa bất kỳ ký tự nào trong nhóm ký tự '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0} không được chứa '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "{0} phải là số ISBN",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} phải là số ISBN-10",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} phải là số ISBN-13",
			override:    false,
		},
		{
			tag:         "issn",
			translation: "{0} phải là số ISSN",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} phải là giá trị UUID",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} phải là giá trị UUID phiên bản 3",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} phải là giá trị UUID phiên bản 4",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} phải là giá trị UUID phiên bản 5",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} chỉ được chứa ký tự ASCII",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} chỉ được chứa ký tự ASCII có thể in ấn",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} chỉ được chứa ký tự multibyte",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} chỉ được chứa Data URI",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} chỉ được chứa latitude (vỹ độ)",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} chỉ được chứa longitude (kinh độ)",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} phải là SSN number",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} phải là địa chỉ IPv4",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} phải là địa chỉ IPv6",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} phải là địa chỉ IP",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} chỉ được chứa CIDR notation",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} chỉ được chứa CIDR notation của một địa chỉ IPv4",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} chỉ được chứa CIDR notation của một địa chỉ IPv6",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} phải là địa chỉ TCP",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} phải là địa chỉ IPv4 TCP",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} phải là địa chỉ IPv6 TCP",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} phải là địa chỉ UDP",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} phải là địa chỉ IPv4 UDP",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} phải là địa chỉ IPv6 UDP",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} phải là địa chỉ IP có thể phân giải",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} phải là địa chỉ IPv4 có thể phân giải",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} phải là địa chỉ IPv6 có thể phân giải",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} phải là địa chỉ UNIX có thể phân giải",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} chỉ được chứa địa chỉ MAC",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} chỉ được chứa những giá trị không trùng lặp",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} phải là màu sắc hợp lệ",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0} phải là trong những giá trị [{1}]",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
		{
			tag:         "json",
			translation: "{0} phải là một chuỗi json hợp lệ",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} phải là một chuỗi jwt hợp lệ",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} phải được viết thường",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} phải được viết hoa",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} không trùng định dạng ngày tháng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0} sai định dạng postcode của quốc gia {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "postcode_iso3166_alpha2_field",
			translation: "{0} sai định dạng postcode của quốc gia tương ứng thuộc trường {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "image",
			translation: "{0} phải là một hình ảnh hợp lệ",
			override:    false,
		},
	}

	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
		} else if t.customTransFunc != nil && t.customRegisFunc == nil {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
		} else if t.customTransFunc == nil && t.customRegisFunc != nil {
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)
		} else {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("cảnh báo: lỗi chuyển ngữ FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
