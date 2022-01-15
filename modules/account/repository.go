package account

import (
	"link-test/business"
	"link-test/business/account"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormDBRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

type AccountTable struct {
	AccNumber  string `gorm:"column:account_number;primaryKey"`
	CustNumber string `gorm:"column:customer_number"`
	Balance    int64  `gorm:"column:balance"`
}

type CustomerTable struct {
	CustNumber string `gorm:"column:customer_number;primaryKey"`
	Name       string `gorm:"column:name"`
}

type AccTableJoinCustTable struct {
	AccNumber string `gorm:"column:account_number"`
	Balance   int64  `gorm:"column:balance"`
	Name      string `gorm:"column:name"`
}

// func (col *AccTableJoinCustTable) ToAccount() account.Account {
// 	var acc account.Account

// 	acc.AccNumber = col.AccNumber
// 	acc.Balance = col.Balance
// 	acc.Name = col.Name

// 	return acc
// }

func (r Repository) FindBalanceByAccNo(accNo string) (*account.Account, error) {
	var accJoin AccTableJoinCustTable
	var acc account.Account

	err := r.DB.Table("account_tables").Select(
		"account_number, balance,customer_tables.name").Joins(
		"left join customer_tables on customer_tables.customer_number = account_tables.customer_number").Where(
		"account_tables.account_number = ?", accNo).Scan(&accJoin).Error

	if err != nil {
		return nil, err
	}

	if accJoin.Name == "" {
		return nil, business.ErrNotFound
	}

	acc.AccNumber = accJoin.AccNumber
	acc.Balance = accJoin.Balance
	acc.Name = accJoin.Name

	return &acc, nil
}

func (r Repository) TransBalance(tr account.TransferRequest) error {

	var accTable AccountTable

	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(&accTable).Where("account_number = ?", tr.FromAccNo).Update("balance", tr.FromAccNoBalance-tr.Amount).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&accTable).Where("account_number = ?", tr.ToAccNo).Update("balance", tr.ToAccNoBalance+tr.Amount).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Error; err != nil {
		return gorm.ErrRecordNotFound
	}

	return tx.Commit().Error

}
