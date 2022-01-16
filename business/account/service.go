package account

import (
	"link-test/business"
)

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}

func (s *service) FindBalanceByAccNo(accNo string) (*Account, error) {

	if accNo == "" {
		return nil, business.ErrInvalidSpec
	}

	account, err := s.repository.FindBalanceByAccNo(accNo)

	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *service) TransBalance(tr TransferRequest) error {

	if tr.FromAccNo == "" || tr.Amount <= 0 || tr.ToAccNo == "" {
		return business.ErrInvalidSpec
	}

	fromAcc, err := s.repository.FindBalanceByAccNo(tr.FromAccNo)
	toAcc, err := s.repository.FindBalanceByAccNo(tr.ToAccNo)

	if err != nil {
		return err
	}

	if fromAcc.Balance < tr.Amount {
		return business.ErrBalanceNotEnough
	}

	tr.FromAccNoBalance = fromAcc.Balance
	tr.ToAccNoBalance = toAcc.Balance

	err = s.repository.TransBalance(tr)

	if err != nil {
		return business.ErrUpdateBalance
	}

	return nil
}
