package login_repository

import (
	"errors"
	"test_kredit_plus/modal/modal_user"

	"gorm.io/gorm"
)

type LoginRepositoryImpl struct{}

func NewLoginRepository() LoginRepository {
	return &LoginRepositoryImpl{}
}

func (repo *LoginRepositoryImpl) Auth(username string, user modal_user.User, db *gorm.DB) (modal_user.User, error) {
	err := db.First(&user, "user_name = ?", username).Error
	if err != nil {
		return user, errors.New("username salah")
	}

	return user, nil
}
