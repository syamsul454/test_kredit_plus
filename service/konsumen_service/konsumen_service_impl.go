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

	user := modal_user.User{
		ID:           uuid.New().String(),
		UserName:     register.Username,
		Email:        register.Email,
		Password:     string(hashedPassword),
		StatusActive: "verifikasi",
	}

	konsumenUser, err := konsomen_service.KonsumenRepository.SaveDataUser(ctx, konsomen_service.DB, user)

	if err != nil {
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

	data_konsumen_result, err := konsomen_service.KonsumenRepository.SaveDataKonsumen(ctx, konsomen_service.DB, data_konsumen)
	if err != nil {
		return map[string]interface{}{}, errors.New("nik sudah ada")
	}
	fmt.Println(data_konsumen_result)
	respon_data := map[string]interface{}{
		"full_name": register.Full_name,
		"username":  register.Username,
	}
	return respon_data, nil
}
