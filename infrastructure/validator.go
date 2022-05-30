package infrastructure

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	cv "github.com/krobus00/iot-be/infrastructure/validator"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
	"github.com/labstack/echo/v4"
)

type (
	ValidationUtil struct {
		validator *validator.Validate
	}
)

func registerTagNameWithLabel(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func NewTranslator() *ut.UniversalTranslator {
	english := en.New()
	indo := id.New()
	uni := ut.New(indo, english)
	return uni
}

func NewValidator(db kro_pkg.Database, trans *ut.UniversalTranslator) echo.Validator {
	validate := validator.New()
	registerTagNameWithLabel(validate)

	id, _ := trans.GetTranslator("id")
	en, _ := trans.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, en)
	_ = idTranslations.RegisterDefaultTranslations(validate, id)
	registerCustomValidation(db, validate, trans)

	return &ValidationUtil{validator: validate}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func registerCustomValidation(db kro_pkg.Database, val *validator.Validate, trans *ut.UniversalTranslator) {
	id, _ := trans.GetTranslator("id")
	en, _ := trans.GetTranslator("en")

	customValidation := cv.New(db)

	val.RegisterValidation("custom", func(fl validator.FieldLevel) bool {
		if fl.Field().Float() > 10 {
			return false
		}
		return true
	})
	registerTranslation(val, en, "custom", "{0} must be less than 10!")
	registerTranslation(val, id, "custom", "{0} harus kurang dari 10!")

	val.RegisterValidation("uniquedb", customValidation.UniqueValidator())
	registerTranslation(val, en, "existdb", "{0} already exist!")
	registerTranslation(val, id, "existdb", "{0} sudah digunakan!")

	val.RegisterValidation("existdb", customValidation.ExistValidator())
	registerTranslation(val, en, "existdb", "{0} not exist!")
	registerTranslation(val, id, "existdb", "{0} tidak ditemukan!")

}

func registerTranslation(v *validator.Validate, trans ut.Translator, tag string, message string) {
	_ = v.RegisterTranslation(tag, trans,
		func(ut ut.Translator) error {
			return ut.Add(tag, message, true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(tag, fe.Field())
			return t
		},
	)
}
