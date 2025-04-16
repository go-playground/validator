package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

var uni *ut.UniversalTranslator

// This example showcases how to use the Validator and UniversalTranslator with both Simplified and Traditional Chinese languages.
// To run the example:
// Step 1: go run _examples/http-transalations/main.go
// Step 2 - Simplified Chinese: curl -d '{"first_name":"foo"}' -H "Accept-Language: zh" -H "Content-Type: application/json" -X POST http://localhost:8081/users
// Step 3 - Traditional Chinese: curl -d '{"first_name":"foo"}' -H "Accept-Language: zh-Hant-TW" -H "Content-Type: application/json" -X POST http://localhost:8081/users
func main() {
	validate := validator.New()
	en := en.New()
	uni = ut.New(en, en, zh.New(), zh_Hant_TW.New())

	validate = validator.New()
	enTrans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, enTrans)
	zhTrans, _ := uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(validate, zhTrans)
	zhHantTrans, _ := uni.GetTranslator("zh_Hant_TW")
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
