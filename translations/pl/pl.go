package pl

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
			translation: "{0} jest wymaganym polem",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "{0} jest wymaganym polem",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "{0} jest wymaganym polem",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} jest wymaganym polem",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0} jest wymaganym polem",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} jest wymaganym polem",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0} jest wymaganym polem",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0} jest wykluczonym polem",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0} jest wykluczonym polem",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0} jest wykluczonym polem",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0} jest wykluczonym polem",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0} jest wykluczonym polem",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0} jest wykluczonym polem",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "{0} musi być domyślną wartością",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "{0} musi mieć długość na {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} musi być równe {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} musi zawierać {1}", false); err != nil {
					return
				}

				if err = registerCardinals(ut, "len"); err != nil {
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
				if err = ut.Add("min-string", "{0} musi mieć długość przynajmniej na {1}", false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} musi być równe {1} lub więcej", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} musi zawierać przynajmniej {1}", false); err != nil {
					return
				}

				if err = registerCardinals(ut, "min"); err != nil {
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
				if err = ut.Add("max-string", "{0} musi mieć długość maksymalnie na {1}", false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} musi być równe {1} lub mniej", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} musi zawierać maksymalnie {1}", false); err != nil {
					return
				}

				if err = registerCardinals(ut, "max"); err != nil {
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
			tag:             "eq",
			translation:     "{0} nie równa się {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ne",
			translation:     "{0} nie powinien być równy {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "{0} musi mieć długość mniejszą niż {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} musi być mniejsze niż {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} musi zawierać mniej niż {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} musi być mniejsze niż obecny dzień i godzina", false); err != nil {
					return
				}

				if err = registerCardinals(ut, "lt"); err != nil {
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
				if err = ut.Add("lte-string", "{0} musi mieć długość maksymalnie na {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} musi być równe {1} lub mniej", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} musi zawierać maksymalnie {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} musi być mniejsze lub równe niż obecny dzień i godzina", false); err != nil {
					return
				}

				if err = registerCardinals(ut, "lte"); err != nil {
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
				if err = ut.Add("gt-string", "{0} musi mieć długość większą niż {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} musi być większe niż {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} musi zawierać więcej niż {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} musi być większe niż obecny dzień i godzina", false); err != nil {
					return
				}

				if err = registerCardinals(ut, "gt"); err != nil {
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
				if err = ut.Add("gte-string", "{0} musi mieć długość przynajmniej na {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} musi być równe {1} lub większe", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} musi zawierać co najmniej {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} musi być większe lub równe niż obecny dzień i godzina", false); err != nil {
					return
				}

				if err = registerCardinals(ut, "gte"); err != nil {
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
			tag:             "eqfield",
			translation:     "{0} musi być równe {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "eqcsfield",
			translation:     "{0} musi być równe {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "necsfield",
			translation:     "{0} nie może być równe {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtcsfield",
			translation:     "{0} musi być większe niż {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtecsfield",
			translation:     "{0} musi być większe lub równe niż {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltcsfield",
			translation:     "{0} musi być mniejsze niż {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltecsfield",
			translation:     "{0} musi być mniejsze lub równe {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "nefield",
			translation:     "{0} nie może być równe {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtfield",
			translation:     "{0} musi być większe niż {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtefield",
			translation:     "{0} musi być większe lub równe {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltfield",
			translation:     "{0} musi być mniejsze niż {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltefield",
			translation:     "{0} musi być mniejsze lub równe {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "alpha",
			translation: "{0} może zawierać wyłącznie znaki alfabetu",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} może zawierać wyłącznie znaki alfanumeryczne",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} musi być poprawną wartością numeryczną",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} musi być poprawną liczbą",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} musi być poprawną wartością heksadecymalną",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} musi być poprawnym kolorem w formacie HEX",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} musi być poprawnym kolorem w formacie RGB",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} musi być poprawnym kolorem w formacie RGBA",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} musi być poprawnym kolorem w formacie HSL",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} musi być poprawnym kolorem w formacie HSLA",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} musi być poprawnym numerem telefonu w formacie E.164",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} musi być poprawnym adresem email",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} musi być poprawnym adresem URL",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} musi być poprawnym adresem URI",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} musi być ciągiem znaków zakodowanym w formacie Base64",
			override:    false,
		},
		{
			tag:             "contains",
			translation:     "{0} musi zawierać tekst '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "containsany",
			translation:     "{0} musi zawierać przynajmniej jeden z następujących znaków '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "excludes",
			translation:     "{0} nie może zawierać tekstu '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "excludesall",
			translation:     "{0} nie może zawierać żadnych z następujących znaków '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "excludesrune",
			translation:     "{0} nie może zawierać następujących znaków '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "isbn",
			translation: "{0} musi być poprawnym numerem ISBN",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} musi być poprawnym numerem ISBN-10",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} musi być poprawnym numerem ISBN-13",
			override:    false,
		},
		{
			tag:         "issn",
			translation: "{0} musi być poprawnym numerem ISSN",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} musi być poprawnym identyfikatorem UUID",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} musi być poprawnym identyfikatorem UUID w wersji 3",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} musi być poprawnym identyfikatorem UUID w wersji 4",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} musi być poprawnym identyfikatorem UUID w wersji 5",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} musi być poprawnym identyfikatorem ULID",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} może zawierać wyłącznie znaki ASCII",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} może zawierać wyłącznie drukowalne znaki ASCII",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} musi zawierać znaki wielobajtowe",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} musi zawierać poprawnie zakodowane dane w formie URI",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} musi zawierać poprawną szerokość geograficzną",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} musi zawierać poprawną długość geograficzną",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} musi zawierać poprawny numer SSN",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} musi zawierać poprawny adres IPv4",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} musi zawierać poprawny adres IPv6",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} musi zawierać poprawny adres IP",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} musi zawierać adres zapisany metodą CIDR",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} musi zawierać adres IPv4 zapisany metodą CIDR",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} musi zawierać adres IPv6 zapisany metodą CIDR",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} musi być poprawnym adresem TCP",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} musi być poprawnym adresem IPv4 TCP",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} musi być poprawnym adresem IPv6 TCP",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} musi być poprawnym adresem UDP",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} musi być poprawnym adresem IPv4 UDP",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} musi być poprawnym adresem IPv6 UDP",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} musi być rozpoznawalnym adresem IP",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} musi być rozpoznawalnym adresem IPv4",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} musi być rozpoznawalnym adresem IPv6",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} musi być rozpoznawalnym adresem UNIX",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} musi zawierać poprawny MAC adres",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0} musi być poprawnym FQDN",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} musi zawierać unikalne wartości",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} musi być prawdziwym kolorem",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0} musi być prawdziwym wyrażeniem cron",
			override:    false,
		},
		{
			tag:             "oneof",
			translation:     "{0} musi być jednym z [{1}]",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "json",
			translation: "{0} musi być ciągiem znaków w formacie JSON",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} musi być ciągiem znaków w formacie JWT",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} musi zawierać wyłącznie małe litery",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} musi zawierać wyłącznie duże litery",
			override:    false,
		},
		{
			tag:             "datetime",
			translation:     "{0} nie spełnia formatu {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "postcode_iso3166_alpha2",
			translation:     "{0} nie spełnia formatu kodu pocztowego kraju {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "postcode_iso3166_alpha2_field",
			translation:     "{0} nie spełnia formatu kodu pocztowego kraju z pola {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "boolean",
			translation: "{0} musi być wartością logiczną",
			override:    false,
		},
		{
			tag:         "image",
			translation: "{0} musi być obrazem",
			override:    false,
		},
		{
			tag:         "cve",
			translation: "{0} musi być poprawnym identyfikatorem CVE",
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

func translateFuncWithParam(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

func registerCardinals(ut ut.Translator, prefix string) (err error) {
	var (
		stringCharacterKey = fmt.Sprintf("%s-string-character", prefix)
		itemsItemKey       = fmt.Sprintf("%s-items-item", prefix)
	)

	if err = ut.AddCardinal(stringCharacterKey, "{0} znak", locales.PluralRuleOne, false); err != nil {
		return
	}

	if err = ut.AddCardinal(stringCharacterKey, "{0} znaki", locales.PluralRuleFew, false); err != nil {
		return
	}

	if err = ut.AddCardinal(stringCharacterKey, "{0} znaków", locales.PluralRuleMany, false); err != nil {
		return
	}

	if err = ut.AddCardinal(stringCharacterKey, "{0} znaków", locales.PluralRuleOther, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} element", locales.PluralRuleOne, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} elementy", locales.PluralRuleFew, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} elementów", locales.PluralRuleMany, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} elementów", locales.PluralRuleOther, false); err != nil {
		return
	}

	return
}
