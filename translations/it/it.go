package it

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

// RegisterDefaultTranslations registers a set of default translations for all built in tag's in validator; you may add your own as desired.
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
			translation: "{0} è un campo obbligatorio",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} è un campo obbligatorio",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "{0} deve essere {1} di lunghezza", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} carattere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} caratteri", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} deve essere uguale a {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} deve contenere {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("len-items-item", "{0} elemento", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} elementi", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err2 error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err2 := strconv.ParseFloat(fe.Param(), 64)
				if err2 != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err2 = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("len-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err2 = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("len-items", fe.Field(), c)

				default:
					t, err2 = ut.T("len-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %s", err2)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("min-string", "{0} deve essere almeno {1} di lunghezza", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} carattere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} caratteri", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} deve essere {1} o maggiore", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} deve contenere almeno {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("min-items-item", "{0} elemento", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} elementi", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err2 error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err2 := strconv.ParseFloat(fe.Param(), 64)
				if err2 != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err2 = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("min-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err2 = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("min-items", fe.Field(), c)

				default:
					t, err2 = ut.T("min-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %s", err2)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("max-string", "{0} deve avere una lunghezza massima di {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} carattere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} caratteri", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} deve essere {1} o inferiore", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} deve contenere al massimo {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("max-items-item", "{0} elemento", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} elementi", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err2 error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err2 := strconv.ParseFloat(fe.Param(), 64)
				if err2 != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err2 = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("max-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err2 = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("max-items", fe.Field(), c)

				default:
					t, err2 = ut.T("max-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %s", err2)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0} non è uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "{0} non dovrebbe essere uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "{0} deve essere inferiore a {1} di lunghezza", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} carattere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} caratteri", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} deve essere inferiore a {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} deve contenere meno di {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} elemento", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} elementi", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} deve essere inferiore alla data e ora correnti", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err2 error
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

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("lt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("lt-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err2 = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err2 = ut.T("lt-datetime", fe.Field())

				default:
					err2 = fn()
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("lt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %s", err2)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lte-string", "{0} deve avere una lunghezza massima di {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} carattere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} caratteri", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} deve essere {1} o inferiore", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} deve contenere al massimo {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} elemento", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} elementi", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} deve essere minore o uguale alla data e ora correnti", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err2 error
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

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("lte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("lte-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err2 = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err2 = ut.T("lte-datetime", fe.Field())

				default:
					err2 = fn()
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("lte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %s", err2)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gt-string", "{0} deve essere maggiore di {1} di lunghezza", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} carattere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} caratteri", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} deve essere maggiore di {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} deve contenere più di {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} elemento", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} elementi", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} deve essere maggiore della data e dell'ora correnti", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err2 error
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

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("gt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("gt-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err2 = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err2 = ut.T("gt-datetime", fe.Field())

				default:
					err2 = fn()
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("gt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %s", err2)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gte-string", "{0} deve essere almeno {1} di lunghezza", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} carattere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} caratteri", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} deve essere {1} o maggiore", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} deve contenere almeno {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} elemento", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} elementi", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} deve essere maggiore o uguale alla data e ora correnti", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err2 error
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

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("gte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err2 = fn()
					if err2 != nil {
						goto END
					}

					c, err2 = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("gte-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err2 = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err2 = ut.T("gte-datetime", fe.Field())

				default:
					err2 = fn()
					if err2 != nil {
						goto END
					}

					t, err2 = ut.T("gte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err2 != nil {
					fmt.Printf("warning: error translating FieldError: %s", err2)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0} deve essere uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0} deve essere uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0} non può essere uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0} deve essere maggiore di {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0} deve essere maggiore o uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0} deve essere minore di {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0} deve essere minore o uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0} non può essere uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0} deve essere maggiore di {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0} deve essere maggiore o uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0} deve essere minore di {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0} deve essere minore o uguale a {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0} può contenere solo caratteri alfabetici",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} può contenere solo caratteri alfanumerici",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} deve essere un valore numerico valido",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} deve essere un numero valido",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} deve essere un esadecimale valido",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} deve essere un colore HEX valido",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} deve essere un colore RGB valido",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} deve essere un colore RGBA valido",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} deve essere un colore HSL valido",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} deve essere un colore HSLA valido",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} deve essere un numero di telefono valido in formato E.164",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} deve essere un indirizzo email valido",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} deve essere un URL valido",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} deve essere un URI valido",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} deve essere una stringa Base64 valida",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0} deve contenere il testo '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "{0} deve contenere almeno uno dei seguenti caratteri '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0} non può contenere il testo '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0} non può contenere nessuno dei seguenti caratteri '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0} non può contenere '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "{0} deve essere un numero ISBN valido",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} deve essere un numero ISBN-10 valido",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} deve essere un numero ISBN-13 valido",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} deve essere un UUID valido",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} deve essere un UUID versione 3 valido",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} deve essere un UUID versione 4 valido",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} deve essere un UUID versione 5 valido",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} deve contenere solo caratteri ASCII",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} deve contenere solo caratteri ASCII stampabili",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} deve contenere caratteri multibyte",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} deve contenere un URI dati valido",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} deve contenere coordinate di latitudine valide",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} deve contenere coordinate di longitudine valide",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} deve essere un numero SSN valido",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} deve essere un indirizzo IPv4 valido",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} deve essere un indirizzo IPv6 valido",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} deve essere un indirizzo IP valido",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} deve contenere una notazione CIDR valida",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} deve contenere una notazione CIDR valida per un indirizzo IPv4",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} deve contenere una notazione CIDR valida per un indirizzo IPv6",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} deve essere un indirizzo TCP valido",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} deve essere un indirizzo TCP IPv4 valido",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} deve essere un indirizzo TCP IPv6 valido",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} deve essere un indirizzo UDP valido",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} deve essere un indirizzo UDP IPv4 valido",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} deve essere un indirizzo UDP IPv6 valido",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} deve essere un indirizzo IP risolvibile",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} deve essere un indirizzo IPv4 risolvibile",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} deve essere un indirizzo IPv6 risolvibile",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} deve essere un indirizzo UNIX risolvibile",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} deve contenere un indirizzo MAC valido",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} deve contenere valori univoci",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} deve essere un colore valido",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0} deve essere uno tra [{1}]",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
		{
			tag:         "json",
			translation: "{0} deve essere una stringa json valida",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} deve essere una stringa minuscola",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} deve essere una stringa maiuscola",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} non corrisponde al formato {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err2 := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err2 != nil {
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
