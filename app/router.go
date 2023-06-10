package app

import (
	konsumen_controller "test_kredit_plus/controller/konsumen_controler"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func NewRouter(controller konsumen_controller.KonsumenControler) *gin.Engine {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("required", func(fl validator.FieldLevel) bool {
			return fl.Field().String() != ""
		})
	}
	r.POST("/register", controller.Register)
	return r
}
