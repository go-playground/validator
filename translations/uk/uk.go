package uk

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
			translation: "{0} обов'язкове поле",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "{0} обов'язкове поле",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "{0} обов'язкове поле",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} обов'язкове поле",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0} обов'язкове поле",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} обов'язкове поле",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0} обов'язкове поле",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0} є виключеним полем",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0} є виключеним полем",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0} є виключеним полем",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0} є виключеним полем",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0} є виключеним полем",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0} є виключеним полем",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "{0} має бути значенням за замовчуванням",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "{0} має бути довжиною в {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} має дорівнювати {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} має містити {1}", false); err != nil {
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
				if err = ut.Add("min-string", "{0} має містити щонайменше {1}", false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} має бути більше чи дорівнювати {1}", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} має містити щонайменше {1}", false); err != nil {
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
				if err = ut.Add("max-string", "{0} має містити максимум {1}", false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} має бути менше чи дорівнювати {1}", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} має містити максимум {1}", false); err != nil {
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
			translation:     "{0} не дорівнює {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ne",
			translation:     "{0} має не дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "{0} має мати менше за {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} має бути менше {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} має містити менше ніж {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} має бути менше поточної дати й часу", false); err != nil {
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
				if err = ut.Add("lte-string", "{0} має містити максимум {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} має бути менше чи дорівнювати {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} має містити максимум {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} має бути менше чи дорівнювати поточній даті та часу", false); err != nil {
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
				if err = ut.Add("gt-string", "{0} має бути довше за {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} має бути більше {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} має містити більше ніж {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} має бути пізніше поточного моменту", false); err != nil {
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
				if err = ut.Add("gte-string", "{0} має містити щонайменше {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} має бути більше чи дорівнювати {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} має містити щонайменше {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} має бути пізніше чи дорівнювати теперішньому моменту", false); err != nil {
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
			translation:     "{0} має дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "eqcsfield",
			translation:     "{0} має дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "necsfield",
			translation:     "{0} не має дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtcsfield",
			translation:     "{0} має бути більше {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtecsfield",
			translation:     "{0} має бути більше чи дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltcsfield",
			translation:     "{0} має бути менше {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltecsfield",
			translation:     "{0} має бути менше чи дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "nefield",
			translation:     "{0} не має дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtfield",
			translation:     "{0} має бути більше {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "gtefield",
			translation:     "{0} має бути більше чи дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltfield",
			translation:     "{0} має бути менше {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "ltefield",
			translation:     "{0} має бути менше чи дорівнювати {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "alpha",
			translation: "{0} має містити тільки літери",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} має містити тільки літери та цифри",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} має бути цифровим значенням",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} має бути цифрою",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} має бути шістнадцятковим рядком",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} має бути HEX кольором",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} має бути RGB кольором",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} має бути RGBA кольором",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} має бути HSL кольором",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} має бути HSLA кольором",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} має бути номером телефону у форматі E.164",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} має бути email адресою",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} має бути URL",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} має бути URI",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} має бути Base64 рядком",
			override:    false,
		},
		{
			tag:             "contains",
			translation:     "{0} має містити текст '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "containsany",
			translation:     "{0} має містити щонайменше один із символів '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "excludes",
			translation:     "{0} не має містити текст '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "excludesall",
			translation:     "{0} не має містити символи '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "excludesrune",
			translation:     "{0} не має містити '{1}'",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "isbn",
			translation: "{0} має бути ISBN номером",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} має бути ISBN-10 номером",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} має бути ISBN-13 номером",
			override:    false,
		},
		{
			tag:         "issn",
			translation: "{0} має бути ISSN номером",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} має бути UUID",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} має бути UUID 3 версії",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} має бути UUID 4 версії",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} має бути UUID 5 версії",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} має бути ULID",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} має містити тільки ascii символи",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} має містити тільки доступні для друку ascii символи",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} має містити мультібайтні символи",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} має містити Data URI",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} має містити координати широти",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} має містити координати довготи",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} має бути SSN номером",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} має бути IPv4 адресою",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} має бути IPv6 адресою",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} має бути IP адресою",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} має містити CIDR позначення",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} має містити CIDR позначення для IPv4 адреси",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} має містити CIDR позначення для IPv6 адреси",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} має бути TCP адресою",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} має бути IPv4 TCP адресою",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} має бути IPv6 TCP адресою",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} має бути UDP адресою",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} має бути IPv4 UDP адресою",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} має бути IPv6 UDP адресою",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} має бути розпізнаваною IP адресою",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} має бути розпізнаваною IPv4 адресою",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} має бути розпізнаваною IPv6 адресою",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} має бути розпізнаваною UNIX адресою",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} має містити MAC адресу",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0} має містити FQDN",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} має містити унікальні значення",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} має бути кольором",
			override:    false,
		},
		{
			tag:             "oneof",
			translation:     "{0} має бути одним з [{1}]",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "json",
			translation: "{0} має бути json рядком",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} має бути jwt рядком",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} має бути рядком у нижньому регістрі",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} має бути рядком у верхньому регістрі",
			override:    false,
		},
		{
			tag:             "datetime",
			translation:     "{0} не відповідає {1} формату",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "postcode_iso3166_alpha2",
			translation:     "{0} не відповідає формату поштового індексу країни {1}",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:             "postcode_iso3166_alpha2_field",
			translation:     "{0} не відповідає формату поштового індексу країни в {1} полі",
			override:        false,
			customTransFunc: translateFuncWithParam,
		},
		{
			tag:         "boolean",
			translation: "{0} має бути булевим значенням",
			override:    false,
		},
		{
			tag:         "image",
			translation: "{0} має бути допустимим зображенням",
			override:    false,
		},
		{
			tag:         "cve",
			translation: "{0} має бути cve ідентифікатором",
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

	if err = ut.AddCardinal(stringCharacterKey, "{0} символ", locales.PluralRuleOne, false); err != nil {
		return
	}

	if err = ut.AddCardinal(stringCharacterKey, "{0} символи", locales.PluralRuleFew, false); err != nil {
		return
	}

	if err = ut.AddCardinal(stringCharacterKey, "{0} символів", locales.PluralRuleMany, false); err != nil {
		return
	}

	if err = ut.AddCardinal(stringCharacterKey, "{0} символи", locales.PluralRuleOther, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} елемент", locales.PluralRuleOne, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} елементи", locales.PluralRuleFew, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} елементів", locales.PluralRuleMany, false); err != nil {
		return
	}

	if err = ut.AddCardinal(itemsItemKey, "{0} елементи", locales.PluralRuleOther, false); err != nil {
		return
	}

	return
}
