package login_service

import (
	"errors"
	"fmt"
	"test_kredit_plus/dto"
	"test_kredit_plus/modal/modal_user"
	"test_kredit_plus/repository/login_repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginServiceImpl struct {
	LoginRepository login_repository.LoginRepository
	DB              *gorm.DB
}

func NewLoginService(loginRepositori login_repository.LoginRepository, db *gorm.DB) LoginService {
	return &LoginServiceImpl{
		LoginRepository: loginRepositori,
		DB:              db,
	}
}

func (service *LoginServiceImpl) Auth(login dto.AuthLogin) (map[string]interface{}, error) {
	user := modal_user.User{}
	responResult, err := service.LoginRepository.Auth(login.Uername, user, service.DB)
	err = bcrypt.CompareHashAndPassword([]byte(responResult.Password), []byte(login.Password))
	fmt.Println(responResult.Password)
	if err != nil {
		return nil, errors.New("password salah")
	}

	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": responResult.UserName,
		"exp":      exp,
		"iat":      time.Now().Unix(),
		"user": jwt.MapClaims{
			"username": responResult.UserName,
			"id":       responResult.ID,
		},
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}

	data := map[string]interface{}{
		"token": tokenString,
	}

	return data, nil
}
