package request

type TransferRequest struct {
	ToAccNumber string `json:"to_account_number"`
	Amount      int64  `json:"amount"`
}
