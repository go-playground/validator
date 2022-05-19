package ko

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
			translation: "{0}필요한필드입니다",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "{0}필요한필드입니다",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "{0}필요한필드입니다",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0}필요한필드입니다",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0}필요한필드입니다",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0}필요한필드입니다",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0}필요한필드입니다",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("len-string", "{0}길이는{1}이어야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("len-string-character", "{0}字符", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("len-string-character", "{0}자", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0}은{1}과같아야합니다", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0}은{1}을포함해야합니다", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("len-items-item", "{0}항목", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("len-items-item", "{0}항목", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("경고: 번역필드오류: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("min-string", "{0}길이는{1}이상이어야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("min-string-character", "{0}个字符", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("min-string-character", "{0}자", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0}는{1}이상이어야합니다", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0}은적어도{1}을포함해야합니다", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("min-items-item", "{0}항목", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("min-items-item", "{0}항목", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("경고: 번역필드오류: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("max-string", "{0}길이는{1}을초과할수없습니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("max-string-character", "{0}자", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("max-string-character", "{0}자", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0}은{1}보다작거나같아야합니다", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0}은최대{1}만포함할수있습니다", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("max-items-item", "{0}항목", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("max-items-item", "{0}항목", locales.PluralRuleOther, false); err != nil {
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
					fmt.Printf("경고: 번역필드오류: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0}은{1}과같지않습니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "{0}은{1}과같지않아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lt-string", "{0}길이는{1}보다작아야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lt-string-character", "{0}자", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lt-string-character", "{0}자", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0}은{1}보다작아야합니다", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0}은{1}미만을포함해야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lt-items-item", "{0}항목", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lt-items-item", "{0}항목", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0}은현재날짜와시간보다작아야합니다", false); err != nil {
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
						err = fmt.Errorf("tag '%s'는유형struct유형에사용할수없습니다.", fe.Tag())
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
					fmt.Printf("경고: 번역필드오류: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lte-string", "{0}길이는{1}을초과할수없습니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lte-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lte-string-character", "{0}자", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0}은{1}보다작거나같아야합니다", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0}은최대{1}만포함할수있습니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lte-items-item", "{0}항목", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0}은현재날짜및시간보다작거나동일해야합니다", false); err != nil {
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
						err = fmt.Errorf("tag '%s'는유형struct유형에사용할수없습니다.", fe.Tag())
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
					fmt.Printf("경고: 번역필드오류: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gt-string", "{0}길이는{1}보다커야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gt-string-character", "{0}个字符", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gt-string-character", "{0}자", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0}은{1}보다커야합니다", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0}은{1}보다커야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gt-items-item", "{0}항목", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gt-items-item", "{0}항목", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0}은현재날짜와시간보다커야합니다", false); err != nil {
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
						err = fmt.Errorf("tag '%s'는유형struct유형에사용할수없습니다.", fe.Tag())
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
					fmt.Printf("경고: 번역필드오류: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gte-string", "{0}길이는{1}이상이어야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gte-string-character", "{0}个字符", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gte-string-character", "{0}자", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0}은{1}보다크거나같아야합니다", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0}은적어도{1}을포함해야합니다", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gte-items-item", "{0}항목", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gte-items-item", "{0}항목", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0}은현재날짜및시간보다크거나동일해야합니다", false); err != nil {
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
						err = fmt.Errorf("tag '%s'는유형struct유형에사용할수없습니다.", fe.Tag())
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
					fmt.Printf("경고: 번역필드오류: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0}은{1}과같아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0}은{1}과같아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0}은{1}과같지않아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0}은{1}보다커야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0}은{1}보다크거나같아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0}은{1}보다작아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0}은{1}보다작거나같아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0}은{1}과같지않아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0}은{1}보다커야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0}은{1}보다크거나같아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0}은{1}보다작아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0}은{1}보다작거나같아야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0}에는문자만포함할수있습니다",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0}에는문자와숫자만포함할수있습니다",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0}은문자,숫자및Unicode문자만포함할수있습니다",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0}은문자와Unicode문자만포함할수있습니다",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0}은유효숫자값이어야합니다",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0}은유효숫자여야합니다",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0}은효과적인16진수여야합니다",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0}은효과적인16진수색상이어야합니다",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0}은효과적인RGB색상이어야합니다",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0}은효과적인RGBA색상이어야합니다",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0}은효과적인HSL색상이어야합니다",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0}은효과적인HSLA색상이어야합니다",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0}은효과적인E.164휴대폰번호여야합니다",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0}은효과적인사서함이어야합니다",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0}은효과적인URL이어야합니다",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0}은효과적인URI여야합니다",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0}은효과적인Base64문자열이어야합니다",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0}은텍스트를포함해야합니다'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "{0}은하나이상의문자를포함해야합니다'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsrune",
			translation: "{0}은문자를포함해야합니다'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0}은텍스트를포함할수없습니다'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0}은다음문자중하나를포함할수없습니다'{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0}은'{1}'을포함할수없습니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "endswith",
			translation: "{0}텍스트'{1}'으로끝나야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "startswith",
			translation: "{0}텍스트'{1}'으로시작해야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "{0}은유효ISBN번호여야합니다",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0}은효과적인ISBN-10번호여야합니다",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0}은효과적인ISBN-13번호여야합니다",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0}은효과적인UUID여야합니다",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0}은효과적인V3 UUID여야합니다",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0}은효과적인V4 UUID여야합니다",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0}은효과적인V5 UUID여야합니다",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0}은효과적인ULID여야합니다",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0}에는ASCII문자만포함해야합니다",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0}에는인쇄가능한ASCII문자만포함해야합니다",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0}에는멀티바이트문자가포함되어야합니다",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0}에는효과적인데이터URI가포함되어야합니다",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0}에는효과적인위도좌표가포함되어야합니다",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0}에는효과적인종방향좌표가포함되어야합니다",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0}은효과적인사회보장번호(SSN)여야합니다",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0}은효과적인IPv4주소여야합니다",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0}은효과적인IPv6주소여야합니다",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0}은효과적인IP주소여야합니다",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0}은효과적인CIDR이어야합니다",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0}은IPv4를포함하는CIDR이어야합니다",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0}은IPv6를포함하는CIDR이어야합니다",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0}은효과적인TCP주소여야합니다",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0}은효과적인IPv4 TCP주소여야합니다",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0}은효과적인IPv6 TCP주소여야합니다",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0}은효과적인UDP주소여야합니다",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0}은효과적인IPv4 UDP주소여야합니다",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0}은효과적인IPv6 UDP주소여야합니다",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0}은효과적인IP주소여야합니다",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0}은효과적인IPv4주소여야합니다",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0}은효과적인IPv6주소여야합니다",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0}은효과적인UNIX주소여야합니다",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0}은효과적인MAC주소여야합니다",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0}필드의값은독특해야합니다",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0}은효과적인색상이어야합니다",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0}은[{1}]중하나여야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
		{
			tag:         "json",
			translation: "{0}은효과적인JSON문자열이어야합니다",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0}은효과적인JWT문자열이어야합니다",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0}은소문자여야합니다",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0}은대문자여야합니다",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0}의형식은{1}이어야합니다",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("경고: 번역필드오류: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "boolean",
			translation: "{0}은효과적인부울값이어야합니다",
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
		log.Printf("경고: 번역필드오류: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
