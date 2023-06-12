package konsumen_repository

import (
	"context"
	"test_kredit_plus/modal/modal_konsumen"
	"test_kredit_plus/modal/modal_user"

	"gorm.io/gorm"
)

type KonsumenRepository interface {
	SaveDataUser(ctx context.Context, db *gorm.DB, User modal_user.User) (modal_user.User, error)
	SaveDataKonsumen(ctx context.Context, db *gorm.DB, Konsumen modal_konsumen.DataKonsumen) (modal_konsumen.DataKonsumen, error)
	SaveDataTenor(ctx context.Context, db *gorm.DB, Tenor modal_konsumen.KonsumenTenor) modal_konsumen.KonsumenTenor
	ValidasiCheckUserexist(ctx context.Context, db *gorm.DB, user, nik string) error
	DetailKonsumen(ctx context.Context, db *gorm.DB, id string) (modal_konsumen.DataKonsumen, error)
	DetailUser(ctx context.Context, db *gorm.DB, id string) (modal_user.User, error)
	BeginTransaction(db *gorm.DB) *gorm.DB
}
