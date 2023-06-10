package main

import (
	"test_kredit_plus/app"
	konsumen_controller "test_kredit_plus/controller/konsumen_controler"
	"test_kredit_plus/modal/modal_konsumen"
	"test_kredit_plus/modal/modal_user"
	"test_kredit_plus/repository/konsumen_repository"
	"test_kredit_plus/service/konsumen_service"
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
	router := app.NewRouter(konsumenController)
	router.Run(":8080")

}
