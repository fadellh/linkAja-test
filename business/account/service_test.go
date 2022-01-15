package account_test

import (
	"link-test/business"
	"link-test/business/account"
	accountMock "link-test/business/account/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	accountService    account.Service
	accountRepository accountMock.Repository

	accountData account.Account

	// mockDB = map[string]*User{
	// 	"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
	// }
	// userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindBalanceByAccNo(t *testing.T) {
	t.Run("Expect Found Data", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accountData, nil).Once()

		account, err := accountService.FindBalanceByAccNo(accountData.AccNumber)

		assert.NotNil(t, account)
		assert.Nil(t, err)
		assert.Equal(t, accountData.Balance, account.Balance)

	})

	t.Run("Expect Account Not Found", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound)

		account, err := accountService.FindBalanceByAccNo(accountData.AccNumber)

		assert.Nil(t, account)
		assert.NotNil(t, err)

	})

	t.Run("Expect Error invalid spec", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(nil, business.ErrInvalidSpec)

		account, err := accountService.FindBalanceByAccNo("")

		assert.Nil(t, account)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInvalidSpec)

	})
}

func TestTransBalance(t *testing.T) {
	t.Run("Expect Create Data", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accountData, nil)
		accountRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(nil).Once()

		err := accountService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555002",
			Amount:    100,
		})

		assert.Nil(t, err)
	})

	t.Run("Expect Error not found", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound)
		accountRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(business.ErrNotFound)

		err := accountService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555003",
			Amount:    100,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect Balance not enough", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accountData, nil)
		accountRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(nil)

		err := accountService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555002",
			Amount:    100000,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrBalanceNotEnough)
	})

	t.Run("Expect Error invalid spec", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accountData, nil)
		accountRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(nil)

		err := accountService.TransBalance(account.TransferRequest{
			FromAccNo: "",
			ToAccNo:   "",
			Amount:    0,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInvalidSpec)
	})

	t.Run("Expect Error Transer", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accountData, nil)
		accountRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(new(error))

		err := accountService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555002",
			Amount:    1000,
		})

		assert.NotNil(t, err)
	})
}

func setup() {

	accountData.AccNumber = "550001"
	accountData.Name = "Bob Martin"
	accountData.Balance = 10000

	accountService = account.NewService(&accountRepository)
}
