package fa

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
			translation: "فیلد {0} اجباری میباشد",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "طول {0} باید {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} کاراکتر", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} کاراکتر", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "طول {0} باید برابر {1} باشد", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "تعداد {0} باید برابر {1} باشد", false); err != nil {
					return
				}
				if err = ut.AddCardinal("len-items-item", "{0} آیتم", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} آیتم", locales.PluralRuleOther, false); err != nil {
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
				if err = ut.Add("min-string", "طول {0} باید حداقل {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} کاراکتر", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} کاراکتر", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} باید بزرگتر یا برابر {1} باشد", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} باید شامل حداقل {1} باشد", false); err != nil {
					return
				}
				if err = ut.AddCardinal("min-items-item", "{0} آیتم", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} آیتم", locales.PluralRuleOther, false); err != nil {
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
				if err = ut.Add("max-string", "طول {0} باید حداکثر {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} کاراکتر", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} کاراکتر", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} باید کمتر یا برابر {1} باشد", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} باید شامل حداکثر {1} باشد", false); err != nil {
					return
				}
				if err = ut.AddCardinal("max-items-item", "{0} آیتم", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} آیتم", locales.PluralRuleOther, false); err != nil {
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
			translation: "{0} برابر {1} نمیباشد",
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
			tag:         "ne",
			translation: "{0} نباید برابر {1} باشد",
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
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "طول {0} باید کمتر از {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} کاراکتر", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} کاراکتر", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} باید کمتر از {1} باشد", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} باید دارای کمتر از {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} آیتم", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} آیتم", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} باید قبل از تاریخ و زمان کنونی باشد", false); err != nil {
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
				if err = ut.Add("lte-string", "طول {0} باید حداکثر {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} کاراکتر", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} کاراکتر", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} باید کمتر یا برابر {1} باشد", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} باید حداکثر شامل {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} آیتم", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} آیتم", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} باید قبل یا برابر تاریخ و زمان کنونی باشد", false); err != nil {
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
				if err = ut.Add("gt-string", "طول {0} باید بیشتر از {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} کاراکتر", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} کاراکتر", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} باید بیشتر از {1} باشد", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} باید دارای بیشتر از {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} آیتم", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} آیتم", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} باید بعد از تاریخ و زمان کنونی باشد", false); err != nil {
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
				if err = ut.Add("gte-string", "طول {0} باید حداقل {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} کاراکتر", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} کاراکتر", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} باید بیشتر یا برابر {1} باشد", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} باید شامل حداقل {1} باشد", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} آیتم", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} آیتم", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} باید بعد یا برابر تاریخ و زمان کنونی باشد", false); err != nil {
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
			translation: "{0} باید برابر {1} باشد",
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
			translation: "{0} باید برابر {1} باشد",
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
			translation: "{0} نمیتواند برابر {1} باشد",
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
			translation: "طول {0} باید بیشتر از {1} باشد",
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
			translation: "طول {0} باید بیشتر یا برابر {1} باشد",
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
			translation: "طول {0} باید کمتر از {1} باشد",
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
			translation: "طول {0} باید کمتر یا برابر {1} باشد",
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
			translation: "{0} نمیتواند برابر {1} باشد",
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
			translation: "طول {0} باید بیشتر از {1} باشد",
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
			translation: "طول {0} باید بیشتر یا برابر {1} باشد",
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
			translation: "طول {0} باید کمتر از {1} باشد",
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
			translation: "طول {0} باید کمتر یا برابر {1} باشد",
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
			translation: "{0} میتواند فقط شامل حروف باشد",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} میتواند فقط شامل حروف و اعداد باشد",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} باید یک عدد معتبر باشد",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} باید یک عدد معتبر باشد",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} باید یک عدد درمبنای16 باشد",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} باید یک کد رنگ HEX باشد",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} باید یک کد رنگ RGB باشد",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} باید یک کد رنگ RGBA باشد",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} باید یک کد رنگ HSL باشد",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} باید یک کد رنگ HSLA باشد",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} باید یک شماره‌تلفن معتبر با فرمت E.164 باشد",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} باید یک ایمیل معتبر باشد",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} باید یک آدرس اینترنتی معتبر باشد",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} باید یک URI معتبر باشد",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} باید یک متن درمبنای64 معتبر باشد",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0} باید شامل '{1}' باشد",
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
			translation: "{0} باید شامل کاراکترهای '{1}' باشد",
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
			translation: "{0} نمیتواند شامل '{1}' باشد",
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
			translation: "{0} نمیتواند شامل کاراکترهای '{1}' باشد",
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
			translation: "{0} نمیتواند شامل '{1}' باشد",
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
			translation: "{0} باید یک شابک معتبر باشد",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} باید یک شابک(ISBN-10) معتبر باشد",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} باید یک شابک(ISBN-13) معتبر باشد",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} باید یک UUID معتبر باشد",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} باید یک UUID نسخه 3 معتبر باشد",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} باید یک UUID نسخه 4 معتبر باشد",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} باید یک UUID نسخه 5 معتبر باشد",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} باید یک ULID معتبر باشد",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} باید فقط شامل کاراکترهای اسکی باشد",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} باید فقط شامل کاراکترهای اسکی قابل چاپ باشد",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} باید شامل کاراکترهای چندبایته باشد",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} باید یک Data URI معتبر باشد",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} باید یک عرض جغرافیایی معتبر باشد",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} باید یک طول جغرافیایی معتبر باشد",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} باید یک شماره SSN معتبر باشد",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} باید یک آدرس آی‌پی IPv4 معتبر باشد",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} باید یک آدرس آی‌پی IPv6 معتبر باشد",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} باید یک آدرس آی‌پی معتبر باشد",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} باید یک نشانه‌گذاری CIDR معتبر باشد",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} باید یک نشانه‌گذاری CIDR معتبر برای آدرس آی‌پی IPv4 باشد",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} باید یک نشانه‌گذاری CIDR معتبر برای آدرس آی‌پی IPv6 باشد",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} باید یک آدرس TCP معتبر باشد",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} باید یک آدرس TCP IPv4 معتبر باشد",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} باید یک آدرس TCP IPv6 معتبر باشد",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} باید یک آدرس UDP معتبر باشد",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} باید یک آدرس UDP IPv4 معتبر باشد",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} باید یک آدرس UDP IPv6 معتبر باشد",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} باید یک آدرس آی‌پی قابل دسترس باشد",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} باید یک آدرس آی‌پی IPv4 قابل دسترس باشد",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} باید یک آدرس آی‌پی IPv6 قابل دسترس باشد",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} باید یک آدرس UNIX معتبر باشد",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} باید یک مک‌آدرس معتبر باشد",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} باید شامل مقادیر منحصربفرد باشد",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} باید یک رنگ معتبر باشد",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0} باید یکی از مقادیر [{1}] باشد",
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
		{
			tag:         "json",
			translation: "{0} باید یک json معتبر باشد",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} باید یک متن با حروف کوچک باشد",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} باید یک متن با حروف بزرگ باشد",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "فرمت {0} با {1} سازگار نیست",
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
			tag:         "postcode_iso3166_alpha2",
			translation: "{0} یک کدپستی معتبر کشور {1} نیست",
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
			tag:         "postcode_iso3166_alpha2_field",
			translation: "{0} یک کدپستی معتبر کشور فیلد {1} نیست",
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
