package konsumen_service

import (
	"context"
	"test_kredit_plus/dto"
)

type KonsumenService interface {
	Register(ctx context.Context, register dto.RegisterKonsumen) (map[string]interface{}, error)
	DetailKonsumen(ctx context.Context, id string) (map[string]interface{}, error)
}
