package auth_controller

import (
	"test_kredit_plus/dto"
	"test_kredit_plus/helper"
	"test_kredit_plus/service/login_service"

	"github.com/gin-gonic/gin"
)

type AuthControllerImple struct {
	LoginService login_service.LoginService
}

func NewAuthController(loginservice login_service.LoginService) AuthController {
	return &AuthControllerImple{
		LoginService: loginservice,
	}
}

func (controller *AuthControllerImple) Login(c *gin.Context) {
	var payloadLogin dto.AuthLogin
	err := c.ShouldBindJSON(&payloadLogin)
	if err != nil {
		c.JSON(400, gin.H{
			"error_validasi": err.Error(),
		})
		return
	}

	data, err := controller.LoginService.Auth(payloadLogin)
	if err != nil {
		helper.SendInternalServerError(c, err)
		return
	}

	helper.SendStatusOk(c, "login success", data)
	return
}
