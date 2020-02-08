package validator

import (
	"github.com/go-playground/validator"
)

type requestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *requestValidator {
	return &requestValidator{validator: validator.New()}
}

func (v *requestValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
