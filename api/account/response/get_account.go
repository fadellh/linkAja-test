package response

import "link-test/business/account"

type GetAccountResponse struct {
	AccNumber string `json:"account_number"`
	Name      string `json:"customer_name"`
	Balance   int64  `json:"balance"`
}

func NewGetAccountResponse(acc account.Account) *GetAccountResponse {

	accountResponse := GetAccountResponse{
		AccNumber: acc.AccNumber,
		Name:      acc.Name,
		Balance:   acc.Balance,
	}

	return &accountResponse
}
