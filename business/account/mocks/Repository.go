package mocks

import (
	"link-test/business/account"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m Repository) FindBalanceByAccNo(accNo string) (*account.Account, error) {
	ret := _m.Called(accNo)

	var r0 *account.Account
	if rf, ok := ret.Get(0).(func(accNo string) *account.Account); ok {
		r0 = rf(accNo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accNo string) error); ok {
		r1 = rf(accNo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m Repository) TransBalance(tr account.TransferRequest) error {
	ret := _m.Called(tr)

	var r0 error
	if rf, ok := ret.Get(0).(func(account.TransferRequest) error); ok {
		r0 = rf(tr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
