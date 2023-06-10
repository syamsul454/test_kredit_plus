package modal_konsumen

type DataKonsumen struct {
	ID             string
	UserId         string
	Nik            string `gorm:"unique"`
	FullName       string
	LegalName      string
	TempatLahir    string
	TanggalLahir   string
	Gaji           uint64
	PhotoSelpi     string
	PhotoKtp       string
	KonsumenTenors []KonsumenTenor `gorm:"foreignKey:KonsumenId;references:ID"`
	Created        int64           `gorm:"autoCreateTime"`
	Updated        int64           `gorm:"autoUpdateTime"`
}
