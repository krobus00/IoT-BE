package validator

import (
	"github.com/go-playground/validator/v10"

	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
)

type (
	CustomValidator interface {
		UniqueValidator(fl validator.FieldLevel) bool
		ExistValidator(fl validator.FieldLevel) bool
	}
	customValidator struct {
		db kro_pkg.Database
	}
)

func New(db kro_pkg.Database) CustomValidator {
	return &customValidator{
		db: db,
	}
}
