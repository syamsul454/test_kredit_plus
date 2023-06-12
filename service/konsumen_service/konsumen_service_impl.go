package konsumen_service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/mail"
	"test_kredit_plus/dto"
	"test_kredit_plus/modal/modal_konsumen"
	"test_kredit_plus/modal/modal_user"
	"test_kredit_plus/repository/konsumen_repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type KonsumenServiceImpl struct {
	KonsumenRepository konsumen_repository.KonsumenRepository
	DB                 *gorm.DB
}

func NewKonsumenService(konsumenRepository konsumen_repository.KonsumenRepository, DB *gorm.DB) KonsumenService {
	return &KonsumenServiceImpl{
		KonsumenRepository: konsumenRepository,
		DB:                 DB,
	}
}

func (konsomen_service *KonsumenServiceImpl) Register(ctx context.Context, register dto.RegisterKonsumen) (map[string]interface{}, error) {
	_, err := mail.ParseAddress(register.Email)
	if err != nil {
		return map[string]interface{}{}, errors.New("email Is valid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Email), bcrypt.DefaultCost)
	if err != nil {
		return map[string]interface{}{}, errors.New("password failed")
	}

	db := konsomen_service.KonsumenRepository.BeginTransaction(konsomen_service.DB)
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	user := modal_user.User{
		ID:           uuid.New().String(),
		UserName:     register.Username,
		Email:        register.Email,
		Password:     string(hashedPassword),
		StatusActive: "verifikasi",
	}

	konsumenUser, err := konsomen_service.KonsumenRepository.SaveDataUser(ctx, db, user)

	if err != nil {
		db.Rollback()
		return map[string]interface{}{}, errors.New("username sudah ada")
	}

	data_konsumen := modal_konsumen.DataKonsumen{
		ID:           uuid.New().String(),
		UserId:       konsumenUser.ID,
		Nik:          register.Nik,
		FullName:     register.Full_name,
		LegalName:    register.Username,
		TempatLahir:  register.TempatLahir,
		TanggalLahir: register.TanggalLahir,
		Gaji:         uint64(register.Gaji),
		KonsumenTenors: []modal_konsumen.KonsumenTenor{
			{
				ID:          uuid.New().String(),
				Name:        "tenor 1",
				JangkaWaktu: 1,
				JumlahLimit: math.Round(float64(register.Gaji) * 0.10 * 1000),
			},
			{
				ID:          uuid.New().String(),
				Name:        "tenor 2",
				JangkaWaktu: 2,
				JumlahLimit: math.Round(float64(register.Gaji) * 0.15 * 1000),
			},
			{
				ID:          uuid.New().String(),
				Name:        "tenor 3",
				JangkaWaktu: 3,
				JumlahLimit: math.Round(float64(register.Gaji) * 0.25 * 1000),
			},
			{
				ID:          uuid.New().String(),
				Name:        "tenor 4",
				JangkaWaktu: 3,
				JumlahLimit: math.Round(float64(register.Gaji) * 0.50 * 1000),
			},
		},
	}

	data_konsumen_result, err := konsomen_service.KonsumenRepository.SaveDataKonsumen(ctx, db, data_konsumen)
	if err != nil {
		db.Rollback()
		return map[string]interface{}{}, errors.New("nik sudah ada")
	}
	fmt.Println(data_konsumen_result)
	db.Commit()
	respon_data := map[string]interface{}{
		"full_name": register.Full_name,
		"username":  register.Username,
	}
	return respon_data, nil
}

func (konsomen_service *KonsumenServiceImpl) DetailKonsumen(ctx context.Context, id string) (map[string]interface{}, error) {
	user, err := konsomen_service.KonsumenRepository.DetailUser(ctx, konsomen_service.DB, id)
	detail, err := konsomen_service.KonsumenRepository.DetailKonsumen(ctx, konsomen_service.DB, id)
	if err != nil {
		return map[string]interface{}{}, errors.New("data tidak ada")
	}

	respon_data := map[string]interface{}{
		"full_name":      detail.FullName,
		"nik":            detail.Nik,
		"tempat_lahir":   detail.TempatLahir,
		"tanggal_lahir":  detail.TanggalLahir,
		"konsumen_tenor": detail.KonsumenTenors,
		"email":          user.Email,
		"username":       user.UserName,
		"status_active":  user.StatusActive,
	}

	return respon_data, nil
}
