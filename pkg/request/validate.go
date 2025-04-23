package request

import "github.com/go-playground/validator"

func isValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
