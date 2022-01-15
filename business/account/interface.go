package account

type Service interface {
	FindBalanceByAccNo(accNo string) (*Account, error)
	TransBalance(tr TransferRequest) error
}

type Repository interface {
	FindBalanceByAccNo(accNo string) (*Account, error)
	TransBalance(tr TransferRequest) error
}
