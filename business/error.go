package business

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")

	ErrNotFound = errors.New("Account was not found")

	ErrInvalidSpec = errors.New("Given spec is not valid")

	ErrBalanceNotEnough = errors.New("Balance is not enough")

	ErrUpdateBalance = errors.New("Error update balance")
)
