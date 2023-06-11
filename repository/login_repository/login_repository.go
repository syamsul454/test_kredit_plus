package login_repository

import (
	"test_kredit_plus/modal/modal_user"

	"gorm.io/gorm"
)

type LoginRepository interface {
	Auth(username string, user modal_user.User, db *gorm.DB) (modal_user.User, error)
}
