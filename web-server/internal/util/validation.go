package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func CreateValidationErrorResponse(err error) (map[string]string, bool) {
	errorMap := make(map[string]string)

	if validatonErrors, ok := err.(validator.ValidationErrors); ok {
		fmt.Println(validatonErrors)
		for _, validatonError := range validatonErrors {
			errorMap[validatonError.Field()] = validatonError.ActualTag()
		}
	} else {
		return errorMap, false
	}

	responseMap := make(map[string]string)

	for key, value := range errorMap {
		var message string

		switch value {
		case "required":
			message = "Field is missing"
		case "email":
			message = "Email format is incorrect"
		default:
			message = "Default error message, missing implementation"
		}

		responseMap[key] = message
	}

	return responseMap, true
}
