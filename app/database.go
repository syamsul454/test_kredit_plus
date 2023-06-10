package app

import (
	"test_kredit_plus/helper"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=kredit_plus port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		helper.PanicIfError(err)
	}
	return db
}

// test speed
