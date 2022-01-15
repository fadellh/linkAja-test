package modules

import (
	"link-test/modules/account"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {

	err := db.AutoMigrate(&account.AccountTable{}, &account.CustomerTable{})

	if err == nil && db.Migrator().HasTable(&account.AccountTable{}) {
		err = db.FirstOrCreate(&account.AccountTable{
			AccNumber:  "555001",
			CustNumber: "1001",
			Balance:    10000,
		}).Error
		err = db.FirstOrCreate(&account.AccountTable{
			AccNumber:  "555002",
			CustNumber: "1002",
			Balance:    15000,
		}).Error
	}

	if err == nil && db.Migrator().HasTable(&account.CustomerTable{}) {
		err = db.FirstOrCreate(&account.CustomerTable{
			CustNumber: "1001",
			Name:       "Bob Martin",
		}).Error
		err = db.FirstOrCreate(&account.CustomerTable{
			CustNumber: "1002",
			Name:       "Linus Torvalds",
		}).Error
	}
}
