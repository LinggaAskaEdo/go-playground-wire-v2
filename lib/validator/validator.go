package validator

import (
	"github.com/go-playground/validator/v10"
)

func InitValidator() *validator.Validate {
	return validator.New()
}
