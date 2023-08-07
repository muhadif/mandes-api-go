package api

import (
	"fmt"
	"net/http"

	validator "gopkg.in/go-playground/validator.v8"
)

func InternalServerError(err error) ErrorResponse {
	result := ErrorResponse{
		Status:    "error",
		ErrorType: "InternalServerError",
		ErrorDescription: []ErrorDetail{
			{
				Field:   "json",
				Message: fmt.Sprintf("%v", err),
			},
		},
	}
	return result
}

func BadRequestError(field string, message string) ErrorResponse {
	result := ErrorResponse{
		Status:    "error",
		ErrorType: "BadRequest",
		HTTPCode:  http.StatusBadRequest,
		ErrorDescription: []ErrorDetail{
			{
				Field:   field,
				Message: message,
			},
		},
	}
	return result
}

func UnauthorizedError(field string, message string) ErrorResponse {
	result := ErrorResponse{
		Status:    "error",
		ErrorType: "Unauthorized",
		HTTPCode:  http.StatusUnauthorized,
		ErrorDescription: []ErrorDetail{
			{
				Field:   field,
				Message: message,
			},
		},
	}
	return result
}

func NotFoundError(entity string) ErrorResponse {
	result := ErrorResponse{
		Status:    "error",
		ErrorType: "NotFound",
		HTTPCode:  http.StatusNotFound,
		ErrorDescription: []ErrorDetail{
			{
				Field:   entity,
				Message: fmt.Sprintf("%v not found", entity),
			},
		},
	}
	return result
}

func ValidationError(err error) error {
	var errDetails []ErrorDetail

	valErrors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, valError := range valErrors {
			errDetails = append(errDetails, ErrorDetail{
				// TODO: extract field name from string tag
				Field:   valError.Field,
				Message: valError.Tag,
			})
		}
	} else {
		errDetails = append(errDetails, ErrorDetail{
			Field:   "unknown",
			Message: err.Error(),
		})
	}

	result := ErrorResponse{
		Status:           "error",
		ErrorDescription: errDetails,
		ErrorType:        "BadRequest",
	}

	return result
}
