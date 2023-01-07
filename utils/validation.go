package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This is not valid email address"
	}
	return "Unknow error"
}

func HandleValidationError(err error) []ValidationErrorMsg {
	var ve validator.ValidationErrors
	var responseError []ValidationErrorMsg
	if errors.As(err, &ve) {
		out := make([]ValidationErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ValidationErrorMsg{fe.Field(), getErrorMessage(fe)}
		}
		responseError = out
	}
	return responseError
}
