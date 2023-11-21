package valid

import "github.com/go-playground/validator/v10"

type Validator struct {
	valid *validator.Validate
}

var validatorInstance Validator

func GetValidator() Validator {
	if validatorInstance.valid == nil {
		validatorInstance.valid = validator.New()
	}

	return validatorInstance
}

func (Validator) Validate(i interface{}) error {
	err := validatorInstance.valid.Struct(i)
	if err != nil {
		return err
	}
	return nil
}
