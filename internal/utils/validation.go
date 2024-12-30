package utils

import (
	"github.com/go-playground/validator/v10"
)

var customErrorMessages = map[string]string{
	"required":     "This field is required",
	"min":          "The value is too short",
	"max":          "The value is too long",
	"email":        "Please enter a valid email address",
	"gte":          "Value should be greater than or equal to the specified minimum",
	"lte":          "Value should be less than or equal to the specified maximum",
	"email_exists": "An account with this email already exists",
}

var validate = validator.New()

func ValidateStruct(input interface{}) (map[string]string, error) {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			message, exists := customErrorMessages[err.Tag()]
			if !exists {
				message = "Invalid value for " + err.Field()
			}

			validationErrors[err.Field()] = message
		}

		return validationErrors, err
	}

	return nil, nil
}
