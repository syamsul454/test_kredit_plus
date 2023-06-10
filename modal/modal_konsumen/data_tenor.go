package modal_konsumen

type KonsumenTenor struct {
	ID          string
	KonsumenId  string
	Name        string
	JangkaWaktu uint
	JumlahLimit float64
	Created     int64 `gorm:"autoCreateTime"`
	Updated     int64 `gorm:"autoUpdateTime"`
}
