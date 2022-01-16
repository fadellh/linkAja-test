package common

import "net/http"

type SuccessResponseCode string

//List of success response status
const (
	Success SuccessResponseCode = "200"
	Created SuccessResponseCode = "201"
)

//SuccessResponse default payload response
type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

type CreateResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
}

//NewSuccessResponse create new success payload
func NewSuccessResponse(data interface{}) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		"Success",
		data,
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponseCreated() (int, CreateResponse) {
	return http.StatusCreated, CreateResponse{
		Created,
		"Success",
	}
}

func NewSuccessResponseDataNull() (int, SuccessResponse) {
	return http.StatusCreated, SuccessResponse{
		Created,
		"Success",
		map[string]interface{}{},
	}
}
