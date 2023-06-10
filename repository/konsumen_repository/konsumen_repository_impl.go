package konsumen_repository

import (
	"context"
	"fmt"
	"test_kredit_plus/modal/modal_konsumen"
	"test_kredit_plus/modal/modal_user"

	"gorm.io/gorm"
)

type KonsumenRepositroyImple struct{}

func NewKonsumenRepositroy() KonsumenRepository {
	return &KonsumenRepositroyImple{}

}

func (repository KonsumenRepositroyImple) SaveDataUser(ctx context.Context, db *gorm.DB, User modal_user.User) (modal_user.User, error) {
	result := db.Create(&User)
	return User, result.Error

}

func (repository KonsumenRepositroyImple) SaveDataTenor(ctx context.Context, db *gorm.DB, Tenor modal_konsumen.KonsumenTenor) modal_konsumen.KonsumenTenor {
	result := db.Create(&Tenor)
	fmt.Println(result.Error)
	return modal_konsumen.KonsumenTenor{}
}

func (repository KonsumenRepositroyImple) SaveDataKonsumen(ctx context.Context, db *gorm.DB, Konsumen modal_konsumen.DataKonsumen) (modal_konsumen.DataKonsumen, error) {
	result := db.Create(&Konsumen)
	fmt.Println(result.Error)
	return modal_konsumen.DataKonsumen{}, result.Error
}

func (repository KonsumenRepositroyImple) ValidasiCheckUserexist(ctx context.Context, db *gorm.DB, user, nik string) error {

	return nil

}
