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
	transService      account.Service
	accountRepository accountMock.Repository
	transRepository   accountMock.Repository

	accountData account.Account
	accData     []account.Account

	// mockDB = map[string]*User{
	// 	"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
	// }
	// userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	accountData.AccNumber = "550001"
	accountData.Name = "Bob Martin"
	accountData.Balance = 10000

	accData = append(accData, account.Account{AccNumber: "550002", Name: "Linux Martin", Balance: 15000})
	accData = append(accData, account.Account{AccNumber: "550001", Name: "Bob Martin", Balance: 10000})
	// fmt.Println(accData)

	accountService = account.NewService(&accountRepository)
	transService = account.NewService(&transRepository)
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
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		account, err := accountService.FindBalanceByAccNo(accountData.AccNumber)

		assert.Nil(t, account)
		assert.NotNil(t, err)

	})

	t.Run("Expect Error invalid spec", func(t *testing.T) {
		accountRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(nil, business.ErrInvalidSpec).Once()
		account, err := accountService.FindBalanceByAccNo("")

		assert.Nil(t, account)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInvalidSpec)

	})
}

func TestTransBalance(t *testing.T) {

	t.Run("Expect Create Data", func(t *testing.T) {
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[0], nil).Once()
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[1], nil).Once()
		transRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(nil).Once()

		err := transService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555002",
			Amount:    100,
		})

		assert.Nil(t, err)
	})

	t.Run("Expect Error not found", func(t *testing.T) {
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		transRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(business.ErrNotFound).Once()

		err := transService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555003",
			Amount:    100,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect Balance not enough", func(t *testing.T) {
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[0], nil).Once()
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[1], nil).Once()
		transRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(nil).Once()

		err := transService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555002",
			Amount:    100000,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrBalanceNotEnough)
	})

	t.Run("Expect Error invalid spec", func(t *testing.T) {
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[0], nil).Once()
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[1], nil).Once()
		transRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(nil).Once()

		err := transService.TransBalance(account.TransferRequest{
			FromAccNo: "",
			ToAccNo:   "",
			Amount:    0,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInvalidSpec)
	})

	t.Run("Expect Error Transer", func(t *testing.T) {
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[0], nil).Once()
		transRepository.On("FindBalanceByAccNo", mock.AnythingOfType("string")).Return(&accData[1], nil).Once()
		transRepository.On("TransBalance", mock.AnythingOfType("account.TransferRequest")).Return(business.ErrUpdateBalance)

		err := transService.TransBalance(account.TransferRequest{
			FromAccNo: "555001",
			ToAccNo:   "555002",
			Amount:    1000,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrUpdateBalance)
	})
}
