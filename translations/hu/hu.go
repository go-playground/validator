package hu

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
			translation: "{0} kötelező mező",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "{0} kötelező mező",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "{0} kötelező mező",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} kötelező mező",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0} kötelező mező",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} kötelező mező",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0} kötelező mező",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0} kizárt mező",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0} kizárt mező",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0} kizárt mező",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0} kizárt mező",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0} kizárt mező",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0} kizárt mező",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "{0} értékének alapértelmezettnek kell lennie",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "{0} értékének {1} hosszúnak kell lennie", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} értékének egyenlőnek kell lennie a következővel: {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} értékének {1} kell tartalmaznia", false); err != nil {
					return
				}
				if err = ut.AddCardinal("len-items-item", "{0} elemet", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} elemet", locales.PluralRuleOther, false); err != nil {
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
				if err = ut.Add("min-string", "{0} értékének legalább {1} hosszúnak kell lennie", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} értéke nem lehet kisebb, mint {1}", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} értékének legalább {1} kell tartalmaznia", false); err != nil {
					return
				}
				if err = ut.AddCardinal("min-items-item", "{0} elemet", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} elemet", locales.PluralRuleOther, false); err != nil {
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

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-items", fe.Field(), c)

				default:
					if fe.Type() == reflect.TypeOf(time.Duration(0)) {
						t, err = ut.T("min-number", fe.Field(), fe.Param())
						goto END
					}

					err = fn()
					if err != nil {
						goto END
					}

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
				if err = ut.Add("max-string", "{0} értéke legfeljebb {1} hosszú lehet", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} értéke nem lehet nagyobb, mint {1}", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} értéke legfeljebb {1} tartalmazhat", false); err != nil {
					return
				}
				if err = ut.AddCardinal("max-items-item", "{0} elemet", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} elemet", locales.PluralRuleOther, false); err != nil {
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

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-items", fe.Field(), c)

				default:
					if fe.Type() == reflect.TypeOf(time.Duration(0)) {
						t, err = ut.T("max-number", fe.Field(), fe.Param())
						goto END
					}

					err = fn()
					if err != nil {
						goto END
					}

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
			translation: "{0} értékének egyenlőnek kell lennie ezzel: {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fmtNumberParam(ut, fe))
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "{0} értéke nem lehet egyenlő a következővel: {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fmtNumberParam(ut, fe))
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
				if err = ut.Add("lt-string", "{0} értékének rövidebbnek kell lennie, mint {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} értékének kisebbnek kell lennie, mint {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} értékének kevesebb mint {1} kell tartalmaznia", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} elemet", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} elemet", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} értékének korábbinak kell lennie, mint az aktuális dátum és idő", false); err != nil {
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
				if err = ut.Add("lte-string", "{0} értéke legfeljebb {1} hosszú lehet", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} értéke nem lehet nagyobb, mint {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} értéke legfeljebb {1} tartalmazhat", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} elemet", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} elemet", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} értéke nem lehet későbbi, mint az aktuális dátum és idő", false); err != nil {
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
				if err = ut.Add("gt-string", "{0} értékének hosszabbnak kell lennie, mint {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} értékének nagyobbnak kell lennie, mint {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} értékének több mint {1} kell tartalmaznia", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} elemet", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} elemet", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} értékének későbbinek kell lennie, mint az aktuális dátum és idő", false); err != nil {
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
				if err = ut.Add("gte-string", "{0} értékének legalább {1} hosszúnak kell lennie", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} értéke nem lehet kisebb, mint {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} értékének legalább {1} kell tartalmaznia", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} elemet", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} elemet", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} értéke nem lehet korábbi, mint az aktuális dátum és idő", false); err != nil {
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
			translation: "{0} értékének meg kell egyeznie {1} értékével",
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
			translation: "{0} értékének meg kell egyeznie {1} értékével",
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
			translation: "{0} értéke nem egyezhet meg {1} értékével",
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
			translation: "{0} értékének nagyobbnak kell lennie, mint {1} értéke",
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
			translation: "{0} értéke nem lehet kisebb, mint {1} értéke",
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
			translation: "{0} értékének kisebbnek kell lennie, mint {1} értéke",
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
			translation: "{0} értéke nem lehet nagyobb, mint {1} értéke",
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
			translation: "{0} értéke nem egyezhet meg {1} értékével",
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
			translation: "{0} értékének nagyobbnak kell lennie, mint {1} értéke",
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
			translation: "{0} értéke nem lehet kisebb, mint {1} értéke",
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
			translation: "{0} értékének kisebbnek kell lennie, mint {1} értéke",
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
			translation: "{0} értéke nem lehet nagyobb, mint {1} értéke",
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
			translation: "{0} értéke csak betűket tartalmazhat",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} értéke csak alfanumerikus karaktereket tartalmazhat",
			override:    false,
		},
		{
			tag:         "alphaspace",
			translation: "{0} értéke csak betűket és szóközöket tartalmazhat",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "{0} értéke csak alfanumerikus karaktereket és szóközöket tartalmazhat",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0} értéke csak Unicode-betűket tartalmazhat",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0} értéke csak alfanumerikus Unicode-karaktereket tartalmazhat",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} értékének érvényes numerikus értéknek kell lennie",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} értékének érvényes számnak kell lennie",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} értékének érvényes hexadecimális számnak kell lennie",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} értékének érvényes HEX-színkódnak kell lennie",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} értékének érvényes RGB-színkódnak kell lennie",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} értékének érvényes RGBA-színkódnak kell lennie",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} értékének érvényes HSL-színkódnak kell lennie",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} értékének érvényes HSLA-színkódnak kell lennie",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} értékének érvényes, E.164 formátumú telefonszámnak kell lennie",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} értékének érvényes e-mail-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} értékének érvényes URL-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} értékének érvényes URI-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} értékének érvényes Base64-karakterláncnak kell lennie",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0} értékének tartalmaznia kell a következő szöveget: '{1}'",
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
			translation: "{0} értékének tartalmaznia kell legalább egyet a következő karakterek közül: '{1}'",
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
			tag:         "startswith",
			translation: "{0} értékének a következővel kell kezdődnie: '{1}'",
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
			tag:         "endswith",
			translation: "{0} értékének a következővel kell végződnie: '{1}'",
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
			tag:         "startsnotwith",
			translation: "{0} értéke nem kezdődhet a következővel: '{1}'",
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
			tag:         "endsnotwith",
			translation: "{0} értéke nem végződhet a következővel: '{1}'",
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
			translation: "{0} értéke nem tartalmazhatja a következő szöveget: '{1}'",
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
			translation: "{0} értéke nem tartalmazhatja a következő karakterek egyikét sem: '{1}'",
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
			translation: "{0} értéke nem tartalmazhatja a következőt: '{1}'",
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
			translation: "{0} értékének érvényes ISBN-számnak kell lennie",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} értékének érvényes ISBN-10 számnak kell lennie",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} értékének érvényes ISBN-13 számnak kell lennie",
			override:    false,
		},
		{
			tag:         "issn",
			translation: "{0} értékének érvényes ISSN-számnak kell lennie",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "{0} értékének érvényes, RFC 2141 szerinti URN-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} értékének érvényes UUID-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} értékének érvényes, 3-as verziójú UUID-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} értékének érvényes, 4-es verziójú UUID-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} értékének érvényes, 5-ös verziójú UUID-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} értékének érvényes ULID-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} értéke csak ASCII-karaktereket tartalmazhat",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} értéke csak nyomtatható ASCII-karaktereket tartalmazhat",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} értékének több bájtból álló karaktereket kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} értékének érvényes Data URI-t kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} értékének érvényes földrajzi szélességi koordinátát kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} értékének érvényes földrajzi hosszúsági koordinátát kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} értékének érvényes társadalombiztosítási számnak (SSN) kell lennie",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} értékének érvényes IPv4-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} értékének érvényes IPv6-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} értékének érvényes IP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} értékének érvényes CIDR-jelölést kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} értékének érvényes, IPv4-címhez tartozó CIDR-jelölést kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} értékének érvényes, IPv6-címhez tartozó CIDR-jelölést kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} értékének érvényes TCP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} értékének érvényes IPv4-es TCP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} értékének érvényes IPv6-os TCP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} értékének érvényes UDP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} értékének érvényes IPv4-es UDP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} értékének érvényes IPv6-os UDP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} értékének feloldható IP-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} értékének feloldható IPv4-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} értékének feloldható IPv6-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} értékének feloldható UNIX-címnek kell lennie",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} értékének érvényes MAC-címet kell tartalmaznia",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0} értékének érvényes teljes tartománynévnek (FQDN) kell lennie",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} értéke csak egyedi elemeket tartalmazhat",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} értékének érvényes színkódnak kell lennie",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0} értékének érvényes cron-kifejezésnek kell lennie",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0} értékének a következők egyikének kell lennie: [{1}]",
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
			translation: "{0} értékének érvényes JSON-karakterláncnak kell lennie",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} értékének érvényes JWT-karakterláncnak kell lennie",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} értékének kisbetűs karakterláncnak kell lennie",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} értékének nagybetűs karakterláncnak kell lennie",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} értéke nem felel meg a következő formátumnak: {1}",
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
			tag:         "timezone",
			translation: "{0} értékének érvényes időzónának kell lennie",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0} értéke nem felel meg a következő ország irányítószám-formátumának: {1}",
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
			translation: "{0} értéke nem felel meg a következő mezőben megadott ország irányítószám-formátumának: {1}",
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
			tag:         "boolean",
			translation: "{0} értékének érvényes logikai értéknek kell lennie",
			override:    false,
		},
		{
			tag:         "image",
			translation: "{0} értékének érvényes képnek kell lennie",
			override:    false,
		},
		{
			tag:         "mimetype",
			translation: "{0} értékének érvényes MIME-típusnak kell lennie",
			override:    false,
		},
		{
			tag:         "cve",
			translation: "{0} értékének érvényes CVE-azonosítónak kell lennie",
			override:    false,
		},
		{
			tag:         "bcp47_strict_language_tag",
			translation: "{0} értékének érvényes BCP 47 nyelvi címkének kell lennie",
			override:    false,
		},
		{
			tag:         "validateFn",
			translation: "{0} értékének érvényes objektumnak kell lennie",
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
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

// fmtNumberParam formats the tag param of numeric fields in the translator's
// number format (e.g. "2.33" -> "2,33"); any other param is returned as is.
func fmtNumberParam(trans ut.Translator, fe validator.FieldError) string {
	typ := fe.Type()
	if typ != nil && typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ == nil || typ == reflect.TypeOf(time.Duration(0)) {
		return fe.Param()
	}

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
	default:
		return fe.Param()
	}

	f64, err := strconv.ParseFloat(fe.Param(), 64)
	if err != nil {
		return fe.Param()
	}

	var digits uint64
	if idx := strings.Index(fe.Param(), "."); idx != -1 {
		digits = uint64(len(fe.Param()[idx+1:]))
	}

	return trans.FmtNumber(f64, digits)
}
