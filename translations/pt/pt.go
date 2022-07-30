package pt

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
			translation: "{0} é obrigatório",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("len-string", "{0} deve ter {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} caractere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} caracteres", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} deve ser igual a {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} deve conter {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("len-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
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

				if err = ut.Add("min-string", "{0} deve ter pelo menos {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} caractere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} caracteres", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} deve ser {1} ou superior", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} deve conter pelo menos {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("min-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
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

				if err = ut.Add("max-string", "{0} deve ter no máximo {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} caractere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} caracteres", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} deve ser {1} ou menos", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} deve conter no máximo {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("max-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
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
			translation: "{0} não é igual a {1}",
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
			translation: "{0} não deve ser igual a {1}",
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

				if err = ut.Add("lt-string", "{0} deve ter menos de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} caractere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} caracteres", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} deve ser menor que {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} deve conter menos de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} deve ser anterior à data / hora atual", false); err != nil {
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

				if err = ut.Add("lte-string", "{0} deve ter no máximo {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} caractere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} caracteres", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} deve ser menor ou igual a {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} deve conter no máximo {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} deve ser anterior ou igual à data/hora atual", false); err != nil {
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

				if err = ut.Add("gt-string", "{0} deve conter mais de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} caractere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} caracteres", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} deve ser maior que {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} deve conter mais de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} deve ser posterior à data/hora atual", false); err != nil {
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

				if err = ut.Add("gte-string", "{0} deve ter pelo menos {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} caractere", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} caracteres", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} deve ser maior ou igual a {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} deve conter pelo menos {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} deve ser posterior ou igual à data/hora atual", false); err != nil {
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
			translation: "{0} deve ser igual a {1}",
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
			translation: "{0} deve ser igual a {1}",
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
			translation: "{0} não deve ser igual a {1}",
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
			translation: "{0} deve ser maior que {1}",
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
			translation: "{0} deve ser maior ou igual que {1}",
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
			translation: "{0} deve ser menor que {1}",
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
			translation: "{0} deve ser menor ou igual que {1}",
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
			translation: "{0} não deve ser igual a {1}",
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
			translation: "{0} deve ser maior que {1}",
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
			translation: "{0} deve ser maior ou igual que {1}",
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
			translation: "{0} deve ser menor que {1}",
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
			translation: "{0} deve ser menor ou igual que {1}",
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
			translation: "{0} deve conter apenas caracteres alfabéticos",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} deve conter apenas caracteres alfanuméricos",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} deve ser um valor numérico válido",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} deve ser um número válido",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} deve ser um hexadecimal válido",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} deve ser uma cor HEX válida",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} deve ser uma cor RGB válida",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} deve ser uma cor RGBA válida",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} deve ser uma cor HSL válida",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} deve ser uma cor HSLA válida",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} deve ser um número de telefone válido no formato E.164",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} deve ser um endereço de e-mail válido",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} deve ser um URL válido",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} deve ser um URI válido",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} deve ser uma string Base64 válida",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0} deve conter o texto '{1}'",
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
			translation: "{0} deve conter pelo menos um dos seguintes caracteres '{1}'",
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
			translation: "{0} não deve conter o texto '{1}'",
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
			translation: "{0} não deve conter os seguintes caracteres '{1}'",
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
			translation: "{0} não pode conter o seguinte '{1}'",
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
			translation: "{0} deve ser um número de ISBN válido",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} deve ser um número ISBN-10 válido",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} deve ser um número ISBN-13 válido",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} deve ser um UUID válido",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} deve ser um UUID versão 3 válido",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} deve ser um UUID versão 4 válido",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} deve ser um UUID versão 5 válido",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} deve ser um ULID válido",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} deve conter apenas caracteres ascii",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} deve conter apenas caracteres ascii imprimíveis",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} deve conter caracteres multibyte",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} deve conter um Data URI válido",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} deve conter uma coordenada de latitude válida",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} deve conter uma coordenada de longitude válida",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} deve ser um número SSN válido",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} deve ser um endereço IPv4 válido",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} deve ser um endereço IPv6 válido",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} deve ser um endereço IP válido",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} não respeita a notação CIDR",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} não respeita a notação CIDR para um endereço IPv4",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} não respeita a notação CIDR para um endereço IPv6",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} deve ser um endereço TCP válido",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} deve ser um endereço TCP IPv4 válido",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} deve ser um endereço TCP IPv6 válido",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} deve ser um endereço UDP válido",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} deve ser um endereço UDP IPv4 válido",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} deve ser um endereço UDP IPv6 válido",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} deve ser um endereço IP resolvível",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} deve ser um endereço IPv4 resolvível",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} deve ser um endereço IPv6 resolvível",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} deve ser um endereço UNIX resolvível",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} deve conter um endereço MAC válido",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} deve conter valores únicos",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} deve ser uma cor válida",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0} deve ser um de [{1}]",
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
			translation: "{0} deve ser uma string json válida",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} deve estar em minuscúlas",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} deve estar em maiúsculas",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} não está no formato {1}",
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
			tag:         "image",
			translation: "{0} deve ser uma imagem válida",
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
