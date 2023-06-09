package app

import (
	"test_kredit_plus/controller/auth_controller"
	konsumen_controller "test_kredit_plus/controller/konsumen_controler"
	"test_kredit_plus/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, controller konsumen_controller.KonsumenControler) *gin.Engine {
	r.POST("/register", controller.Register)
	r.Use(middleware.JWTChatMiddleware())
	r.GET("/get-user", controller.GetUser)
	return r
}

func RouterLogin(r *gin.Engine, cont auth_controller.AuthController) *gin.Engine {
	r.POST("/login", cont.Login)
	return r
}
