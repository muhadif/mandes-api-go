package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhadif/mandes/pkg/fault"
	"log"
	"net/http"
)

// Response is the standard response model
type Response struct {
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// Pagination contains all pagination-related info
type Pagination struct {
	CurrentPage int `json:"CurrentPage,omitempty"`
	TotalItems  int `json:"TotalItems,omitempty"`
	TotalPage   int `json:"TotalPage,omitempty"`
}

// Error implements the Error interface
func (e ErrorResponse) Error() string {
	if len(e.ErrorDescription) > 0 {
		return fmt.Sprintf("Error %v (%v): %v - %v", e.HTTPCode, e.ErrorType, e.ErrorDescription[0].Field, e.ErrorDescription[0].Message)
	}
	return fmt.Sprintf("Error %v (%v)", e.HTTPCode, e.ErrorType)
}

// ErrorResponse is the standard response format for non-2xx responses
type ErrorResponse struct {
	Status           string        `json:"status,omitempty"`
	ErrorType        string        `json:"error,omitempty"`
	HTTPCode         int           `json:"-"`
	ErrorCode        int           `json:"errorCode,omitempty"`
	ErrorDescription []ErrorDetail `json:"errorDescription,omitempty"`
}

// ErrorDetail contains the field and message
type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func ResponseFromError(err error) ErrorResponse {
	log.Printf("Response Error: %v\n", err)

	switch v := err.(type) {

	case ErrorResponse:
		return v
	case fault.UnauthorizedError:
		return ErrorResponse{
			Status:    "error",
			ErrorType: "UnauthorizedError",
			HTTPCode:  http.StatusUnauthorized,
			ErrorCode: http.StatusUnauthorized,
		}
	case fault.MissingParamError:
		customErr := err.(fault.MissingParamError)
		return ErrorResponse{
			Status:    "error",
			ErrorType: "MissingParamError",
			HTTPCode:  http.StatusBadRequest,
			ErrorCode: http.StatusBadRequest,
			ErrorDescription: []ErrorDetail{
				{
					Field:   customErr.Field,
					Message: customErr.Error(),
				},
			},
		}
	case fault.InvalidValueError:
		customErr := err.(fault.InvalidTypeError)
		return ErrorResponse{
			Status:    "error",
			ErrorType: "InvalidValueError",
			HTTPCode:  http.StatusBadRequest,
			ErrorCode: http.StatusBadRequest,
			ErrorDescription: []ErrorDetail{
				{
					Field:   customErr.Field,
					Message: customErr.Error(),
				},
			},
		}

	default:
		return ErrorResponse{
			HTTPCode:  http.StatusInternalServerError,
			ErrorCode: http.StatusInternalServerError,
			Status:    "error",
			ErrorType: "InternalServerError",
			ErrorDescription: []ErrorDetail{
				{
					Field:   "InternalServerError",
					Message: err.Error(),
				},
			},
		}
	}
}

func ResponseSuccess(c *gin.Context, status int, data interface{}) {
	resp := Response{
		Status: "success",
		Data:   data,
	}

	c.JSON(status, resp)
}

func ResponseFailed(c *gin.Context, err error) {
	resp := ResponseFromError(err)
	c.JSON(resp.HTTPCode, resp)
	c.Abort()
}

