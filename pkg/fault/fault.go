package fault

import "fmt"

type UnauthorizedError struct {
	EntityName string
	Message    string
}

func (err UnauthorizedError) Error() string {
	return err.Message
}

type EntityNotExist struct {
	EntityName string
}

func (err EntityNotExist) Error() string {
	return fmt.Sprintf("%s doesn't exist", err.EntityName)
}

type RecordNotFoundError struct {
	EntityName string
}

func (e RecordNotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.EntityName)
}

type JSONDecodeError struct {
	Field string
}

func (err JSONDecodeError) Error() string {
	return fmt.Sprintf("Invalid json")
}

type MissingParamError struct {
	Field string
}

func (err MissingParamError) Error() string {
	return fmt.Sprintf("%s is required", err.Field)
}

type InvalidTypeError struct {
	Field string
	Type  string
}

func (err InvalidTypeError) Error() string {
	return fmt.Sprintf("%s must be in %s type", err.Field, err.Type)
}

type DuplicatedEntity struct {
	EntityName string
}

func (err DuplicatedEntity) Error() string {
	return fmt.Sprintf("%s already exists.", err.EntityName)
}

type InvalidValueError struct {
	EntityName string
	Message    string
}

func (err InvalidValueError) Error() string {
	return err.Message
}

type APIClientResponseError struct {
	HTTPCode  int
	ErrorCode int
	URL       string
	Message   string
}

func (err APIClientResponseError) Error() string {
	errorMsg := fmt.Sprintf("Message : %v", err.Message)
	return errorMsg
}
