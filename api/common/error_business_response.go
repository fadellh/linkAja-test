package common

import (
	"link-test/business"
	"net/http"
)

type errorBusinessResponseCode string

const (
	errInternalServerError errorBusinessResponseCode = "500"
	errNotFound            errorBusinessResponseCode = "404"
	errInvalidSpec         errorBusinessResponseCode = "400"
	errBalance             errorBusinessResponseCode = "400"
	errUpdate              errorBusinessResponseCode = "503"
)

//BusinessResponse default payload response
type BusinessResponse struct {
	Code    errorBusinessResponseCode `json:"code"`
	Message string                    `json:"message"`
	Data    interface{}               `json:"data"`
}

//NewErrorBusinessResponse Response return choosen http status like 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewErrorBusinessResponse(err error) (int, BusinessResponse) {
	return errorMapping(err)
}

//errorMapping error for missing header key with given value
func errorMapping(err error) (int, BusinessResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()
	case business.ErrNotFound:
		return newNotFoundResponse()
	case business.ErrInvalidSpec:
		return newValidationResponse(err.Error())
	case business.ErrBalanceNotEnough:
		return newBalanceNotEnough(err.Error())
	case business.ErrUpdateBalance:
		return newErrUpdateBalance(err.Error())
	}
}

func newInternalServerErrorResponse() (int, BusinessResponse) {
	return http.StatusInternalServerError, BusinessResponse{
		errInternalServerError,
		"Internal server error",
		map[string]interface{}{},
	}
}

func newNotFoundResponse() (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errNotFound,
		"Account Not found",
		map[string]interface{}{},
	}
}

func newValidationResponse(message string) (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errInvalidSpec,
		"Validation failed " + message,
		map[string]interface{}{},
	}
}

func newBalanceNotEnough(message string) (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errBalance,
		message,
		map[string]interface{}{},
	}
}

func newErrUpdateBalance(message string) (int, BusinessResponse) {
	return http.StatusServiceUnavailable, BusinessResponse{
		errUpdate,
		message,
		map[string]interface{}{},
	}
}
