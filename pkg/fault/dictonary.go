package fault

import (
	"fmt"
	"net/http"
)

type ErrorCode int

type CustomError struct {
	HTTPCode      ErrorCode
	Message       string
	AdditionalErr []*AdditionalErr
}

type AdditionalErr struct {
	Key     string
	Message string
}

const (
	HTTPUnauthorizedError       ErrorCode = http.StatusUnauthorized
	HTTPForbiddenRequestError   ErrorCode = http.StatusForbidden
	HTTPInternalServiceError    ErrorCode = http.StatusInternalServerError
	HTTPNotFound                ErrorCode = http.StatusNotFound
	HTTPBadGatewayError         ErrorCode = http.StatusBadGateway
	HTTPBadRequestError         ErrorCode = http.StatusBadRequest
	HTTPNotImplemented          ErrorCode = http.StatusNotImplemented
	HTTPPreconditionFailedError ErrorCode = http.StatusPreconditionFailed
	HTTPConflictError           ErrorCode = http.StatusConflict
)

func (err CustomError) Error() string {
	errorMsg := fmt.Sprintf("[ERROR] %v : (%v) - %v \n", err.HTTPCode, err.Message, err.AdditionalErr)
	return errorMsg
}

func GetErrorCode(err error) ErrorCode {
	e, ok := err.(CustomError)
	if !ok {
		return HTTPInternalServiceError
	}

	return e.HTTPCode
}

func GetCustomError(err error) CustomError {
	e, ok := err.(CustomError)
	if !ok {
		return CustomError{}
	}

	return e
}

func ErrorDictionary(customErrorCode ErrorCode, msg string, opts ...func(*CustomError)) CustomError {
	err := &CustomError{
		HTTPCode:      customErrorCode,
		Message:       msg,
		AdditionalErr: nil,
	}
	for _, o := range opts {
		o(err)
	}
	return *err
}

func WithAdditionalMsg(addon *AdditionalErr) func(*CustomError) {
	return func(err *CustomError) {
		if err.AdditionalErr == nil {
			err.AdditionalErr = []*AdditionalErr{addon}
		} else {
			err.AdditionalErr = append(err.AdditionalErr, addon)
		}
	}
}
