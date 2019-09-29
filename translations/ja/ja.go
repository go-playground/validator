package ja

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
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
			translation: "{0}は必須フィールドです",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("len-string", "{0}の長さは{1}でなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("len-string-character", "{0}文字", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-string-character", "{0}文字", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0}は{1}と等しくなければなりません", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0}は{1}を含まなければなりません", false); err != nil {
					return
				}
				// if err = ut.AddCardinal("len-items-item", "{0}つの項目", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-items-item", "{0}つの項目", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("min-string", "{0}の長さは少なくとも{1}はなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("min-string-character", "{0}文字", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-string-character", "{0}文字", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0}は{1}かより大きくなければなりません", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0}は少なくとも{1}を含まなければなりません", false); err != nil {
					return
				}
				// if err = ut.AddCardinal("min-items-item", "{0}つの項目", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-items-item", "{0}つの項目", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("max-string", "{0}の長さは最大でも{1}でなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("max-string-character", "{0}文字", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-string-character", "{0}文字", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0}は{1}かより小さくなければなりません", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0}は最大でも{1}を含まなければなりません", false); err != nil {
					return
				}
				// if err = ut.AddCardinal("max-items-item", "{0}つの項目", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-items-item", "{0}つの項目", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0}は{1}と等しくありません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "ne",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("ne-items", "{0}の項目の数は{1}と異ならなければなりません", false); err != nil {
					fmt.Printf("ne customRegisFunc #1 error because of %v\n", err)
					return
				}
				// if err = ut.AddCardinal("ne-items-item", "{0}個", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("ne-items-item", "{0}個", locales.PluralRuleOther, false); err != nil {
					return
				}
				if err = ut.Add("ne", "{0}は{1}と異ならなければなりません", false); err != nil {
					fmt.Printf("ne customRegisFunc #2 error because of %v\n", err)
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
				case reflect.Slice:
					var c string
					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("ne-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}
					t, err = ut.T("ne-items", fe.Field(), c)
				default:
					t, err = ut.T("ne", fe.Field(), fe.Param())
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lt-string", "{0}の長さは{1}よりも少なくなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-string-character", "{0}文字", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-string-character", "{0}文字", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0}は{1}よりも小さくなければなりません", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0}は{1}よりも少ない項目を含まなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-items-item", "{0}つの項目", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-items-item", "{0}つの項目", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0}は現時刻よりも前でなければなりません", false); err != nil {
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
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
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
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lte-string", "{0}の長さは最大でも{1}でなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-string-character", "{0}文字", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-string-character", "{0}文字", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0}は{1}かより小さくなければなりません", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0}は最大でも{1}を含まなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-items-item", "{0}つの項目", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-items-item", "{0}つの項目", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0}は現時刻以前でなければなりません", false); err != nil {
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
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
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
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gt-string", "{0}の長さは{1}よりも多くなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-string-character", "{0}文字", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-string-character", "{0}文字", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0}は{1}よりも大きくなければなりません", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0}は{1}よりも多い項目を含まなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-items-item", "{0}つの項目", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-items-item", "{0}つの項目", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0}は現時刻よりも後でなければなりません", false); err != nil {
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
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
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
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gte-string", "{0}の長さは少なくとも{1}以上はなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-string-character", "{0}文字", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-string-character", "{0}文字", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0}は{1}かより大きくなければなりません", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0}は少なくとも{1}を含まなければなりません", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-items-item", "{0}つの項目", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-items-item", "{0}つの項目", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0}は現時刻以降でなければなりません", false); err != nil {
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
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
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
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0}は{1}と等しくなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0}は{1}と等しくなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0}は{1}とは異ならなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0}は{1}よりも大きくなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0}は{1}以上でなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0}は{1}よりも小さくなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0}は{1}以下でなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0}は{1}とは異ならなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0}は{1}よりも大きくなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0}は{1}以上でなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0}は{1}よりも小さくなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0}は{1}以下でなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0}はアルファベットのみを含むことができます",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0}はアルファベットと数字のみを含むことができます",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0}は正しい数字でなければなりません",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0}は正しい数でなければなりません",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0}は正しい16進表記でなければなりません",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0}は正しいHEXカラーコードでなければなりません",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0}は正しいRGBカラーコードでなければなりません",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0}は正しいRGBAカラーコードでなければなりません",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0}は正しいHSLカラーコードでなければなりません",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0}は正しいHSLAカラーコードでなければなりません",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0}は正しいメールアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0}は正しいURLでなければなりません",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0}は正しいURIでなければなりません",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0}は正しいBase64文字列でなければなりません",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0}は'{1}'を含まなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "{0}は'{1}'の少なくとも1つを含まなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0}には'{1}'というテキストを含むことはできません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0}には'{1}'のどれも含めることはできません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0}には'{1}'を含めることはできません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "{0}は正しいISBN番号でなければなりません",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0}は正しいISBN-10番号でなければなりません",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0}は正しいISBN-13番号でなければなりません",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0}は正しいUUIDでなければなりません",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0}はバージョンが3の正しいUUIDでなければなりません",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0}はバージョンが4の正しいUUIDでなければなりません",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0}はバージョンが4の正しいUUIDでなければなりません",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0}はASCII文字のみを含まなければなりません",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0}は印刷可能なASCII文字のみを含まなければなりません",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0}はマルチバイト文字を含まなければなりません",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0}は正しいデータURIを含まなければなりません",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0}は正しい緯度の座標を含まなければなりません",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0}は正しい経度の座標を含まなければなりません",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0}は正しい社会保障番号でなければなりません",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0}は正しいIPv4アドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0}は正しいIPv6アドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0}は正しいIPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0}は正しいCIDR表記を含まなければなりません",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0}はIPv4アドレスの正しいCIDR表記を含まなければなりません",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0}はIPv6アドレスの正しいCIDR表記を含まなければなりません",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0}は正しいTCPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0}は正しいIPv4のTCPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0}は正しいIPv6のTCPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0}は正しいUDPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0}は正しいIPv4のUDPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0}は正しいIPv6のUDPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0}は解決可能なIPアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0}は解決可能なIPv4アドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0}は解決可能なIPv6アドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0}は解決可能なUNIXアドレスでなければなりません",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0}は正しいMACアドレスを含まなければなりません",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0}は正しい色でなければなりません",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0}は[{1}]のうちのいずれかでなければなりません",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
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
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
