package hy

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
			translation: "«{0}» դաշտը պարտադիր է",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "«{0}» դաշտը պարտադիր է",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "«{0}» դաշտը պարտադիր է",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "«{0}» դաշտը պարտադիր է",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "«{0}» դաշտը պարտադիր է",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "«{0}» դաշտը պարտադիր է",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "«{0}» դաշտը պարտադիր է",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "«{0}» դաշտը պետք է բացակայի",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "«{0}» դաշտը պետք է բացակայի",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "«{0}» դաշտը պետք է բացակայի",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "«{0}» դաշտը պետք է բացակայի",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "«{0}» դաշտը պետք է բացակայի",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "«{0}» դաշտը պետք է բացակայի",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "«{0}» դաշտը պետք է ունենա լռելյայն արժեք",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "«{0}» դաշտի երկարությունը պետք է լինի {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} նիշ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} նիշ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "«{0}» դաշտի արժեքը պետք է լինի {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "«{0}» դաշտը պետք է պարունակի {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("len-items-item", "{0} տարր", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} տարր", locales.PluralRuleOther, false); err != nil {
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
				if err = ut.Add("min-string", "«{0}» դաշտը պետք է պարունակի առնվազն {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} նիշ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} նիշ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "«{0}» դաշտի արժեքը պետք է լինի առնվազն {1}", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "«{0}» դաշտը պետք է պարունակի առնվազն {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("min-items-item", "{0} տարր", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} տարր", locales.PluralRuleOther, false); err != nil {
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
				if err = ut.Add("max-string", "«{0}» դաշտը պետք է պարունակի առավելագույնը {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} նիշ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} նիշ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "«{0}» դաշտի արժեքը չպետք է գերազանցի {1}", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "«{0}» դաշտը պետք է պարունակի առավելագույնը {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("max-items-item", "{0} տարր", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} տարր", locales.PluralRuleOther, false); err != nil {
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
			translation: "«{0}» դաշտի արժեքը պետք է լինի {1}",
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
			translation: "«{0}» դաշտի արժեքը չպետք է լինի {1}",
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
				if err = ut.Add("lt-string", "«{0}» դաշտի երկարությունը պետք է փոքր լինի {1}-ից", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} նիշ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} նիշ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "«{0}» դաշտի արժեքը պետք է փոքր լինի {1}-ից", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "«{0}» դաշտը պետք է պարունակի {1}-ից պակաս", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} տարր", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} տարր", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "«{0}» դաշտի արժեքը պետք է ընթացիկ պահից ավելի վաղ լինի", false); err != nil {
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
				if err = ut.Add("lte-string", "«{0}» դաշտի երկարությունը չպետք է գերազանցի {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} նիշ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} նիշ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "«{0}» դաշտի արժեքը չպետք է գերազանցի {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "«{0}» դաշտը պետք է պարունակի առավելագույնը {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} տարր", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} տարր", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "«{0}» դաշտի արժեքը պետք է ընթացիկ պահից ոչ ուշ լինի", false); err != nil {
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
				if err = ut.Add("gt-string", "«{0}» դաշտի երկարությունը պետք է մեծ լինի {1}-ից", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} նիշ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} նիշ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "«{0}» դաշտի արժեքը պետք է մեծ լինի {1}-ից", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "«{0}» դաշտը պետք է պարունակի {1}-ից ավելի", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} տարր", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} տարր", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "«{0}» դաշտի արժեքը պետք է ընթացիկ պահից ավելի ուշ լինի", false); err != nil {
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
				if err = ut.Add("gte-string", "«{0}» դաշտը պետք է պարունակի առնվազն {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} նիշ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} նիշ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "«{0}» դաշտի արժեքը պետք է լինի առնվազն {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "«{0}» դաշտը պետք է պարունակի առնվազն {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} տարր", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} տարր", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "«{0}» դաշտի արժեքը պետք է ընթացիկ պահից ոչ վաղ լինի", false); err != nil {
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
			translation: "«{0}» և «{1}» դաշտերի արժեքները պետք է համընկնեն",
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
			translation: "«{0}» և «{1}» դաշտերի արժեքները պետք է համընկնեն",
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
			translation: "«{0}» և «{1}» դաշտերի արժեքները չպետք է համընկնեն",
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
			translation: "«{0}» դաշտի արժեքը պետք է մեծ լինի «{1}» դաշտի արժեքից",
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
			translation: "«{0}» դաշտի արժեքը պետք է մեծ կամ հավասար լինի «{1}» դաշտի արժեքին",
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
			translation: "«{0}» դաշտի արժեքը պետք է փոքր լինի «{1}» դաշտի արժեքից",
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
			translation: "«{0}» դաշտի արժեքը պետք է փոքր կամ հավասար լինի «{1}» դաշտի արժեքին",
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
			translation: "«{0}» և «{1}» դաշտերի արժեքները չպետք է համընկնեն",
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
			translation: "«{0}» դաշտի արժեքը պետք է մեծ լինի «{1}» դաշտի արժեքից",
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
			translation: "«{0}» դաշտի արժեքը պետք է մեծ կամ հավասար լինի «{1}» դաշտի արժեքին",
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
			translation: "«{0}» դաշտի արժեքը պետք է փոքր լինի «{1}» դաշտի արժեքից",
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
			translation: "«{0}» դաշտի արժեքը պետք է փոքր կամ հավասար լինի «{1}» դաշտի արժեքին",
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
			translation: "«{0}» դաշտը կարող է պարունակել միայն տառեր",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "«{0}» դաշտը կարող է պարունակել միայն տառեր և թվանշաններ",
			override:    false,
		},
		{
			tag:         "alphaspace",
			translation: "«{0}» դաշտը կարող է պարունակել միայն տառեր և բացատներ",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "«{0}» դաշտը կարող է պարունակել միայն տառեր, թվանշաններ և բացատներ",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "«{0}» դաշտը կարող է պարունակել միայն Unicode տառային նիշեր",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "«{0}» դաշտը կարող է պարունակել միայն Unicode տառային և թվային նիշեր",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր թվային արժեք",
			override:    false,
		},
		{
			tag:         "number",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր թիվ",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր տասնվեցական արժեք",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր HEX գույն",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր RGB գույն",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր RGBA գույն",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր HSL գույն",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր HSLA գույն",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "«{0}» դաշտը պետք է պարունակի E.164 ձևաչափին համապատասխան հեռախոսահամար",
			override:    false,
		},
		{
			tag:         "email",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր էլեկտրոնային փոստի հասցե",
			override:    false,
		},
		{
			tag:         "url",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր URL",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր URI",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր Base64 տող",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "«{0}» դաշտը պետք է պարունակի «{1}» տեքստը",
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
			translation: "«{0}» դաշտը պետք է պարունակի հետևյալ նիշերից առնվազն մեկը՝ «{1}»",
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
			translation: "«{0}» դաշտը չպետք է պարունակի «{1}» տեքստը",
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
			translation: "«{0}» դաշտը չպետք է պարունակի հետևյալ նիշերից որևէ մեկը՝ «{1}»",
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
			translation: "«{0}» դաշտը չպետք է պարունակի «{1}» նիշը",
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
			translation: "«{0}» դաշտը պետք է պարունակի վավեր ISBN համար",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր ISBN-10 համար",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր ISBN-13 համար",
			override:    false,
		},
		{
			tag:         "issn",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր ISSN համար",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "«{0}» դաշտը պետք է պարունակի RFC 2141 ստանդարտին համապատասխան URN",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր UUID",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր 3-րդ տարբերակի UUID",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր 4-րդ տարբերակի UUID",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր 5-րդ տարբերակի UUID",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր ULID",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "«{0}» դաշտը կարող է պարունակել միայն ASCII նիշեր",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "«{0}» դաշտը կարող է պարունակել միայն տպվող ASCII նիշեր",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "«{0}» դաշտը պետք է պարունակի բազմաբայթ նիշեր",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր Data URI",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր լայնության կոորդինատներ",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր երկայնության կոորդինատներ",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր SSN համար",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր IPv4 հասցե",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր IPv6 հասցե",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր IP հասցե",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր CIDR գրառում",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "«{0}» դաշտը պետք է պարունակի IPv4 հասցեի վավեր CIDR գրառում",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "«{0}» դաշտը պետք է պարունակի IPv6 հասցեի վավեր CIDR գրառում",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր TCP հասցե",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր IPv4 TCP հասցե",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր IPv6 TCP հասցե",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր UDP հասցե",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր IPv4 UDP հասցե",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր IPv6 UDP հասցե",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "«{0}» դաշտը պետք է պարունակի լուծելի IP հասցե",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "«{0}» դաշտը պետք է պարունակի լուծելի IPv4 հասցե",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "«{0}» դաշտը պետք է պարունակի լուծելի IPv6 հասցե",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "«{0}» դաշտը պետք է պարունակի լուծելի UNIX հասցե",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր MAC հասցե",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր FQDN",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "«{0}» դաշտը պետք է պարունակի միայն եզակի արժեքներ",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր գույն",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր cron արտահայտություն",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "«{0}» դաշտի արժեքը պետք է լինի հետևյալներից մեկը՝ [{1}]",
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
			translation: "«{0}» դաշտը պետք է պարունակի վավեր JSON տող",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր JWT տող",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "«{0}» դաշտի արժեքը պետք է գրված լինի փոքրատառերով",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "«{0}» դաշտի արժեքը պետք է գրված լինի մեծատառերով",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "«{0}» դաշտը չի համապատասխանում {1} ձևաչափին",
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
			translation: "«{0}» դաշտը պետք է պարունակի վավեր ժամային գոտի",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "«{0}» դաշտը չի համապատասխանում {1} երկրի փոստային ինդեքսի ձևաչափին",
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
			translation: "«{0}» դաշտը չի համապատասխանում «{1}» դաշտում նշված երկրի փոստային ինդեքսի ձևաչափին",
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
			translation: "«{0}» դաշտը պետք է պարունակի վավեր տրամաբանական արժեք",
			override:    false,
		},
		{
			tag:         "image",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր պատկեր",
			override:    false,
		},
		{
			tag:         "mimetype",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր MIME տեսակ",
			override:    false,
		},
		{
			tag:         "cve",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր CVE նույնացուցիչ",
			override:    false,
		},
		{
			tag:         "bcp47_strict_language_tag",
			translation: "«{0}» դաշտը պետք է պարունակի BCP 47 ստանդարտին համապատասխան լեզվային պիտակ",
			override:    false,
		},
		{
			tag:         "validateFn",
			translation: "«{0}» դաշտը պետք է պարունակի վավեր օբյեկտ",
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
