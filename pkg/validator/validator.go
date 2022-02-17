package validator

import (
	"context"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator
var trans ut.Translator

func init() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")

	validate = validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatal(err)
	}

}

func ValidateStruct(ctx context.Context, s interface{}) error {
	loadTranslations()
	return validate.StructCtx(ctx, s)
}
