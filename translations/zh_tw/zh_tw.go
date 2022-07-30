package zh_tw

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
			translation: "{0}為必填欄位",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("len-string", "{0}長度必須為{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("len-string-character", "{0}字元", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("len-string-character", "{0}個字元", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0}必須等於{1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0}必須包含{1}", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("len-items-item", "{0}項", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("len-items-item", "{0}項", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("警告: 翻譯欄位錯誤: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("min-string", "{0}長度必須至少為{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("min-string-character", "{0}個字元", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("min-string-character", "{0}個字元", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0}最小只能為{1}", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0}必須至少包含{1}", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("min-items-item", "{0}項", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("min-items-item", "{0}項", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("警告: 翻譯欄位錯誤: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("max-string", "{0}長度不能超過{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("max-string-character", "{0}個字元", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("max-string-character", "{0}個字元", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0}必須小於或等於{1}", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0}最多只能包含{1}", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("max-items-item", "{0}項", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("max-items-item", "{0}項", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("警告: 翻譯欄位錯誤: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0}不等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "{0}不能等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lt-string", "{0}長度必須小於{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lt-string-character", "{0}個字元", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lt-string-character", "{0}個字元", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0}必須小於{1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0}必須包含少於{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lt-items-item", "{0}項", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lt-items-item", "{0}項", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0}必須小於目前日期和時間", false); err != nil {
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
						err = fmt.Errorf("tag '%s'不能用於struct類型.", fe.Tag())
					} else {
						t, err = ut.T("lt-datetime", fe.Field())
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("警告: 翻譯欄位錯誤: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lte-string", "{0}長度不能超過{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lte-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lte-string-character", "{0}個字元", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0}必須小於或等於{1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0}最多只能包含{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lte-items-item", "{0}項", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0}必須小於或等於目前日期和時間", false); err != nil {
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
						err = fmt.Errorf("tag '%s'不能用於struct類型.", fe.Tag())
					} else {
						t, err = ut.T("lte-datetime", fe.Field())
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("警告: 翻譯欄位錯誤: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gt-string", "{0}長度必須大於{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gt-string-character", "{0}個字元", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gt-string-character", "{0}個字元", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0}必須大於{1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0}必須大於{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gt-items-item", "{0}項", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gt-items-item", "{0}項", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0}必須大於目前日期和時間", false); err != nil {
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
						err = fmt.Errorf("tag '%s'不能用於struct類型.", fe.Tag())
					} else {
						t, err = ut.T("gt-datetime", fe.Field())
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("警告: 翻譯欄位錯誤: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gte-string", "{0}長度必須至少為{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gte-string-character", "{0}個字元", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gte-string-character", "{0}個字元", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0}必須大於或等於{1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0}必須至少包含{1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gte-items-item", "{0}項", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gte-items-item", "{0}項", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0}必須大於或等於目前日期和時間", false); err != nil {
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
						err = fmt.Errorf("tag '%s'不能用於struct類型.", fe.Tag())
					} else {
						t, err = ut.T("gte-datetime", fe.Field())
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("警告: 翻譯欄位錯誤: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0}必須等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0}必須等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0}不能等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0}必須大於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0}必須大於或等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0}必須小於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0}必須小於或等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0}不能等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0}必須大於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0}必須大於或等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0}必須小於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0}必須小於或等於{1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0}只能包含字母",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0}只能包含字母和數字",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0}必須是一個有效的數值",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0}必須是一個有效的數字",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0}必須是一個有效的十六進制",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0}必須是一個有效的十六進制顏色",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0}必須是一個有效的RGB顏色",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0}必須是一個有效的RGBA顏色",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0}必須是一個有效的HSL顏色",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0}必須是一個有效的HSLA顏色",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0}必須是一個有效的信箱",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0}必須是一個有效的URL",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0}必須是一個有效的URI",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0}必須是一個有效的Base64字元串",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0}必須包含文字'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "{0}必須包含至少一個以下字元'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0}不能包含文字'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0}不能包含以下任何字元'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0}不能包含'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "{0}必須是一個有效的ISBN編號",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0}必須是一個有效的ISBN-10編號",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0}必須是一個有效的ISBN-13編號",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0}必須是一個有效的UUID",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0}必須是一個有效的V3 UUID",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0}必須是一個有效的V4 UUID",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0}必須是一個有效的V5 UUID",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0}必須是一個有效的ULID",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0}必須只包含ascii字元",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0}必須只包含可輸出的ascii字元",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0}必須包含多個字元",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0}必須包含有效的數據URI",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0}必須包含有效的緯度座標",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0}必須包含有效的經度座標",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0}必須是一個有效的社會安全編號(SSN)",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0}必須是一個有效的IPv4地址",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0}必須是一個有效的IPv6地址",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0}必須是一個有效的IP地址",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0}必須是一個有效的無類別域間路由(CIDR)",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0}必須是一个包含IPv4地址的有效無類別域間路由(CIDR)",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0}必須是一个包含IPv6地址的有效無類別域間路由(CIDR)",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0}必須是一個有效的TCP地址",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0}必須是一個有效的IPv4 TCP地址",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0}必須是一個有效的IPv6 TCP地址",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0}必須是一個有效的UDP地址",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0}必須是一個有效的IPv4 UDP地址",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0}必須是一個有效的IPv6 UDP地址",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0}必須是一個有效的IP地址",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0}必須是一個有效的IPv4地址",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0}必須是一個有效的IPv6地址",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0}必須是一個有效的UNIX地址",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0}必須是一個有效的MAC地址",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0}必須是一個有效的顏色",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0}必須是[{1}]中的一個",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
		{
			tag:         "datetime",
			translation: "{0}與{1}格式不匹配",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "image",
			translation: "{0} 必須是有效圖像",
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
		log.Printf("警告: 翻譯欄位錯誤: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
