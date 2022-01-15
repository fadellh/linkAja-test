package account

type Account struct {
	AccNumber  string
	Name       string
	CustNumber string
	Balance    int64
}

type TransferRequest struct {
	FromAccNo        string
	ToAccNo          string
	Amount           int64
	FromAccNoBalance int64
	ToAccNoBalance   int64
}
