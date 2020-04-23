package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string
	Status int
	Error string
	Causes []interface{}
}

func NewError(msg string)error  {
	return errors.New(msg)
}

func NewRestError(msg string,status int,error string)*RestErr  {
	return &RestErr{
		Message: msg,
		Status:    status,
		Error: error,
	}

}

func NewBadRequestError(message string)*RestErr  {
	return &RestErr{
		Message: message,
		Status:    http.StatusBadRequest,
		Error:   "Bad Request",
	}
}
func NewNotFoundError(message string)*RestErr  {
	return &RestErr{
		Message: message,
		Status:    http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(message string,err error)*RestErr  {
	result:= &RestErr{
		Message: message,
		Status:    http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
	if err!=nil {
		result.Causes=append(result.Causes,err.Error())
	}
	return result
}