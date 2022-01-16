package common

import "net/http"

type errorControllerResponseCode string

const (
	ErrBadRequest errorControllerResponseCode = "400"
	ErrForbidden  errorControllerResponseCode = "403"
)

type ControllerResponse struct {
	Code    errorControllerResponseCode `json:"code"`
	Message string                      `json:"message"`
	Data    interface{}                 `json:"data"`
}

func NewBadRequestResponse(msg string) (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		ErrBadRequest,
		"Bad request " + msg,
		map[string]interface{}{},
	}
}

func NewForbiddenResponse() (int, ControllerResponse) {
	return http.StatusForbidden, ControllerResponse{
		ErrForbidden,
		"Forbidden",
		map[string]interface{}{},
	}
}
