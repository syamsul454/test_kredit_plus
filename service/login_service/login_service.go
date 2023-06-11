package login_service

import "test_kredit_plus/dto"

type LoginService interface {
	Auth(login dto.AuthLogin) (map[string]interface{}, error)
}
