package business

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")

	ErrHasBeenModified = errors.New("Data has been modified")

	ErrNotFound = errors.New("Account was not found")

	ErrInvalidSpec = errors.New("Given spec is not valid")

	ErrBalanceNotEnough = errors.New("Balance is not enough")
)
