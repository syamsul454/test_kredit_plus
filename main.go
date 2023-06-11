package main

import (
	"test_kredit_plus/app"
	"test_kredit_plus/controller/auth_controller"
	konsumen_controller "test_kredit_plus/controller/konsumen_controler"
	"test_kredit_plus/modal/modal_konsumen"
	"test_kredit_plus/modal/modal_user"
	"test_kredit_plus/repository/konsumen_repository"
	"test_kredit_plus/repository/login_repository"
	"test_kredit_plus/service/konsumen_service"
	"test_kredit_plus/service/login_service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func main() {
	db := app.ConnectDB()
	// konsumen
	// db.Migrator().DropTable(&modal_user.User{}, &modal_konsumen.DataKonsumen{}, &modal_konsumen.KonsumenTenor{})
	db.AutoMigrate(&modal_user.User{})
	db.AutoMigrate(&modal_konsumen.DataKonsumen{})
	db.AutoMigrate(&modal_konsumen.KonsumenTenor{})
	konsumenRepository := konsumen_repository.NewKonsumenRepositroy()
	konsumenService := konsumen_service.NewKonsumenService(konsumenRepository, db)
	konsumenController := konsumen_controller.NewKonsumenController(konsumenService)
	// repository login
	loginRepositori := login_repository.NewLoginRepository()
	loginService := login_service.NewLoginService(loginRepositori, db)
	loginController := auth_controller.NewAuthController(loginService)

	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("required", func(fl validator.FieldLevel) bool {
			return fl.Field().String() != ""
		})
	}
	router := app.RouterLogin(r, loginController)
	router = app.NewRouter(r, konsumenController)
	router.Run(":8080")
}
