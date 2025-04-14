package main

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
	"net/http"
	"strings"
)

var uni *ut.UniversalTranslator

func main() {
	validate := validator.New()
	en := en.New()
	uni = ut.New(en, en, zh.New(), zh_Hant.New())

	validate = validator.New()
	enTrans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, enTrans)
	zhTrans, _ := uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(validate, zhTrans)
	zhHantTrans, _ := uni.GetTranslator("zh_Hant")
	zh_tw_translations.RegisterDefaultTranslations(validate, zhHantTrans)

	type User struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// ... fill user value
		var user User

		// Header Accept-Language value is en or zh
		trans, _ := uni.GetTranslator(strings.Replace(r.Header.Get("Accept-Language"), "-", "_", -1))
		if err := validate.Struct(&user); err != nil {
			var errs validator.ValidationErrors
			var httpErrors []validator.ValidationErrorsTranslations
			if errors.As(err, &errs) {
				httpErrors = append(httpErrors, errs.Translate(trans))
			}
			r, _ := json.Marshal(httpErrors)
			w.Write(r)
		}
	})

	http.ListenAndServe(":8081", nil)
}
