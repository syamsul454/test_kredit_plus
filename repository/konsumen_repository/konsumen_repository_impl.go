package konsumen_repository

import (
	"context"
	"fmt"
	"test_kredit_plus/modal/modal_konsumen"
	"test_kredit_plus/modal/modal_user"

	"github.com/getsentry/sentry-go"
	"gorm.io/gorm"
)

type KonsumenRepositroyImple struct{}

func NewKonsumenRepositroy() KonsumenRepository {
	return &KonsumenRepositroyImple{}

}

func (repository KonsumenRepositroyImple) SaveDataUser(ctx context.Context, db *gorm.DB, User modal_user.User) (modal_user.User, error) {
	result := db.Create(&User)
	if result.Error != nil {
		// tracking error sentry.io error
		errCostum := fmt.Errorf("error save data user %v", result.Error)
		sentry.CaptureException(errCostum)
	}
	return User, result.Error

}

func (repository KonsumenRepositroyImple) SaveDataTenor(ctx context.Context, db *gorm.DB, Tenor modal_konsumen.KonsumenTenor) modal_konsumen.KonsumenTenor {
	result := db.Create(&Tenor)
	fmt.Println(result.Error)
	// database commit
	db.Commit()
	return modal_konsumen.KonsumenTenor{}
}

func (repository KonsumenRepositroyImple) SaveDataKonsumen(ctx context.Context, db *gorm.DB, Konsumen modal_konsumen.DataKonsumen) (modal_konsumen.DataKonsumen, error) {
	result := db.Create(&Konsumen)
	if result.Error != nil {
		// tracking error sentry.io error
		errCostum := fmt.Errorf("error save data konsumen %v", result.Error)
		sentry.CaptureException(errCostum)

	}
	// database commit
	db.Commit()
	return modal_konsumen.DataKonsumen{}, result.Error
}

func (repository KonsumenRepositroyImple) ValidasiCheckUserexist(ctx context.Context, db *gorm.DB, user, nik string) error {

	return nil

}

func (repository KonsumenRepositroyImple) DetailKonsumen(ctx context.Context, db *gorm.DB, id string) (modal_konsumen.DataKonsumen, error) {
	var dataKonsumen modal_konsumen.DataKonsumen
	// err := db.First(&dataKonsumen, "user_id").Error
	// one to many
	// db join table
	// err := db.Where("user_id = ?", id).First(&dataKonsumen).Error
	err := db.Joins("JOIN konsumen_tenors ON data_konsumens.id = konsumen_tenors.konsumen_id").Where("user_id = ?", id).Preload("KonsumenTenors").First(&dataKonsumen).Error

	if err != nil {
		errCostum := fmt.Errorf("error get detail konsumen %v", err)
		sentry.CaptureException(errCostum)
		return dataKonsumen, err
	}

	return dataKonsumen, nil
}

func (repository KonsumenRepositroyImple) DetailUser(ctx context.Context, db *gorm.DB, id string) (modal_user.User, error) {
	var dataUser modal_user.User
	err := db.Where("id = ?", id).First(&dataUser).Error
	if err != nil {
		errCostum := fmt.Errorf("error get detail user %v", err)
		sentry.CaptureException(errCostum)
		return dataUser, err
	}

	return dataUser, nil
}

func (repository KonsumenRepositroyImple) BeginTransaction(db *gorm.DB) *gorm.DB {
	return db.Begin()
}
