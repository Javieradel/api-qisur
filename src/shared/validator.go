package shared

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

type XValidator struct {
	validator *validator.Validate
}

func (v *XValidator) Validate(data interface{}) []ErrorResponse {
	var errors []ErrorResponse
	err := v.validator.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ErrorResponse
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Message = err.Error()
			errors = append(errors, el)
		}
	}
	return errors
}

func NewValidator() *XValidator {
	validate := validator.New()

	validate.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if val, ok := field.Interface().(decimal.Decimal); ok {
			v, _ := val.Float64()
			return v
		}
		return nil
	}, decimal.Decimal{})

	return &XValidator{
		validator: validate,
	}
}
