package modal_user

import (
	"test_kredit_plus/modal/modal_konsumen"
)

type User struct {
	ID           string `gorm:"index"`
	UserName     string `gorm:"unique"`
	Email        string
	Password     string
	StatusActive string
	Konsumen     modal_konsumen.DataKonsumen `gorm:"foreignKey:UserId;references:ID"`
	Created      int64                       `gorm:"autoCreateTime"`
	Updated      int64                       `gorm:"autoUpdateTime"`
}
